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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int64
	_ = v205
	var v210 int32
	_ = v210
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int64
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int64
	_ = v463
	var v464 int64
	_ = v464
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v516 int32
	_ = v516
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v567 int32
	_ = v567
	var v577 int32
	_ = v577
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v652 int32
	_ = v652
	var v657 int32
	_ = v657
	var v663 int32
	_ = v663
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v740 int32
	_ = v740
	var v743 int64
	_ = v743
	var v745 int64
	_ = v745
	var v747 int64
	_ = v747
	var v749 int64
	_ = v749
	var v751 int32
	_ = v751
	var v758 int32
	_ = v758
	var v767 int32
	_ = v767
	var v772 int32
	_ = v772
	var v778 int32
	_ = v778
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v807 int32
	_ = v807
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v955 int32
	_ = v955
	var v969 int32
	_ = v969
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v1000 int32
	_ = v1000
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1034 int32
	_ = v1034
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1079 int32
	_ = v1079
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1111 int32
	_ = v1111
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1127 int32
	_ = v1127
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1156 int32
	_ = v1156
	var v1173 int32
	_ = v1173
	var v1177 int32
	_ = v1177
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1187 int32
	_ = v1187
	var v1192 int32
	_ = v1192
	var v1196 int32
	_ = v1196
	var v1214 int32
	_ = v1214
	var v1216 int32
	_ = v1216
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1231 int32
	_ = v1231
	var v1236 int32
	_ = v1236
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1243 int32
	_ = v1243
	var v1247 int32
	_ = v1247
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1257 int32
	_ = v1257
	var v1267 int32
	_ = v1267
	var v1269 int32
	_ = v1269
	var v1277 int32
	_ = v1277
	var v1282 int32
	_ = v1282
	var v1284 int32
	_ = v1284
	var v1286 int32
	_ = v1286
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v21 = F_ShmemGetRequestedSize(m)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v23 = F_add_size(m, int32(_a_F_CreateSharedMemoryAndSemaphores_0), v21)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[0]))
	v27 = F_add_size(m, v23, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v33 = F_add_size(m, v27, int32(_a_F_CreateSharedMemoryAndSemaphores_1)-v27&int32(_a_F_CreateSharedMemoryAndSemaphores_2))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v37 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v37 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v33
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_3), v18)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v48 = m.G0
	v50 = v48 - int32(352)
	m.G0 = v50
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[1]))
	v58 = F___fstatat(m, int32(-100), v53, v50+int32(192), int32(0))
	mBase = m.M
	goto L13
L10:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_4), int32(130), int32(_a_F_CreateSharedMemoryAndSemaphores_5))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	F_InitShmemAllocator(m, v751)
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L1
	} else {
		goto L207
	}
L13:
	;
	if int32(0) <= v58 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[2]))
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[3]))
	if base.B2i32(v64 == int32(1))&base.B2i32(v68 != int32(2)) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L1
	} else {
		goto L203
	}
L17:
	;
	if v68 == int32(2) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	goto L19
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L1
	} else {
		goto L199
	}
L20:
	;
	v205 = *(*int64)(unsafe.Add(mBase, uint32(v50)+280))
	v210 = base.I32_wrap_i64(v205)
	goto L60
L21:
	;
	v77 = int32(1)
	if base.Ui32(v77) < base.Ui32(v64-v77) {
		v133 = v33
		v134 = int32(-1)
		v135 = int32(0)
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	F_SetConfigOption(m, int32(_a_F_CreateSharedMemoryAndSemaphores_6), int32(_a_F_CreateSharedMemoryAndSemaphores_7), int32(0), int32(1))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L59
	}
L24:
	;
	if v134 == int32(-1) {
		goto L41
	} else {
		goto L42
	}
L25:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[4]))
	if v82 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v86 = v82 << (uint(int32(10)) % 32)
	goto L28
L27:
	;
	v86 = int32(_a_F_CreateSharedMemoryAndSemaphores_8)
	goto L28
L28:
	;
	if v86 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v97 = int32(_a_F_CreateSharedMemoryAndSemaphores_9) - base.I32_wrap_i64(base.I64_clz(base.I64_extend_i32_u(v86)-int64(1)))<<(uint(int32(26))%32)
	goto L31
L30:
	;
	v97 = int32(_a_F_CreateSharedMemoryAndSemaphores_9)
	goto L31
L31:
	;
	v98 = base.I32_rem_u_s(v33, v86)
	if v98 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v100 = F_add_size(m, v33, v86-v98)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	v102 = v33
	goto L34
L34:
	;
	v103 = int32(-1)
	v104 = F_mmap(m, v102, v97, v103)
	mBase = m.M
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[5]))
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[2]))
	if base.B2i32(v108 != int32(2))|base.B2i32(v104 != v103) != 0 {
		v133 = v102
		v134 = v104
		v135 = v106
		goto L24
	} else {
		goto L36
	}
L35:
	;
	v102 = v100
	goto L34
L36:
	;
	v114 = int32(-1)
	v117 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	if v117 == int32(0) {
		v133 = v102
		v134 = v114
		v135 = v106
		goto L24
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+176)) = v102
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_10), v50+int32(176))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_11), int32(629), int32(_a_F_CreateSharedMemoryAndSemaphores_12))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v133 = v102
	v134 = v114
	v135 = v106
	goto L24
L41:
	;
	v141 = int32(_a_F_CreateSharedMemoryAndSemaphores_7)
	goto L43
L42:
	;
	v141 = int32(_a_F_CreateSharedMemoryAndSemaphores_13)
	goto L43
L43:
	;
	F_SetConfigOption(m, int32(_a_F_CreateSharedMemoryAndSemaphores_6), v141, int32(0), int32(1))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if v134 != int32(-1) {
		v157 = v133
		v158 = v134
		v159 = v135
		goto L45
	} else {
		goto L46
	}
L45:
	;
	if v158 == int32(-1) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[2]))
	if v149 == int32(1) {
		v157 = v133
		v158 = v134
		v159 = v135
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v154 = F_mmap(m, v33, int32(33), int32(-1))
	mBase = m.M
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[5]))
	v157 = v33
	v158 = v154
	v159 = v156
	goto L45
L48:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[5])) = v159
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[6])) = v157
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[7])) = v158
	F_on_shmem_exit(m, int32(962), int64(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L58
	}
L51:
	;
	F_errmsg(m, int32(_a_F_CreateSharedMemoryAndSemaphores_14), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	if v159 == int32(48) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = v157
	F_errhint(m, int32(_a_F_CreateSharedMemoryAndSemaphores_15), v50+int32(16))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_11), int32(665), int32(_a_F_CreateSharedMemoryAndSemaphores_12))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	v201 = v157
	v203 = int32(32)
	goto L20
L59:
	;
	v201 = v33
	v203 = v33
	goto L20
L60:
	;
	v222 = int32(0)
	v223 = int32(_a_F_CreateSharedMemoryAndSemaphores_16)
	v229 = F___strchrnul(m, v223, int32(61))
	mBase = m.M
	if v223 == v229 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	if v271 != 0 {
		goto L78
	} else {
		goto L79
	}
L63:
	;
	v271 = int32(0)
	goto L62
L64:
	;
	goto L65
L65:
	;
	v232 = v229 - v223
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232)+uint32(_c_F_CreateSharedMemoryAndSemaphores[8]))))
	if v234 != 0 {
		v265 = v222
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v271 = v265
	goto L62
L67:
	;
	v236 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[9]))
	if v236 == int32(0) {
		v265 = v222
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v236)))
	if v239 == int32(0) {
		v265 = v222
		goto L66
	} else {
		goto L69
	}
L69:
	;
	v243 = v236
	v244 = v239
	goto L70
L70:
	;
	v247 = F_strncmp(m, v223, v244, v232)
	mBase = m.M
	if v247 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v265 = v251 + int32(1)
	goto L66
L72:
	;
	goto L71
L73:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	v251 = v250 + v232
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
	if v252 == int32(61) {
		goto L72
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v243)+4))
	if v256 != 0 {
		v243 = v243 + int32(4)
		v244 = v256
		goto L70
	} else {
		goto L77
	}
L76:
	;
	goto L75
L77:
	;
	v265 = v222
	goto L66
L78:
	;
	v272 = int32(0)
	v275 = F_strtox_2(m, v271, v272, v272, int64(4294967295))
	mBase = m.M
	goto L81
L79:
	;
	v277 = v222
	goto L80
L80:
	;
	v279 = F_pgmem_shmget(m, v210, v203, int32(1920))
	mBase = m.M
	if v279 < int32(0) {
		goto L87
	} else {
		goto L88
	}
L81:
	;
	v277 = base.I32_wrap_i64(v275)
	goto L80
L82:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v50)+188))
	if v790 == int32(0) {
		v210 = v778
		goto L60
	} else {
		goto L193
	}
L83:
	;
	v778 = v210 + int32(1)
	goto L82
L84:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L1
	} else {
		goto L190
	}
L85:
	;
	v722 = m.Env.Pgmem_getpid(m)
	mBase = m.M
	v723 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v385)+16)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v385))) = int32(679834894)
	*(*int32)(unsafe.Add(mBase, uint32(v385)+4)) = v722
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v50)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v385)+24)) = v205
	*(*int32)(unsafe.Add(mBase, uint32(v385)+20)) = v728
	*(*int32)(unsafe.Add(mBase, uint32(v385)+12)) = int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v385)+8)) = v201
	*(*int32)(unsafe.Add(mBase, uint32(v18+int32(12)))) = v385
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[10])) = v210
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[11])) = v385
	v740 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[7]))
	if v740 == v723 {
		goto L187
	} else {
		goto L188
	}
L86:
	;
	v407 = int32(0)
	v408 = F_pgmem_shmget(m, v210, int32(32), v407)
	mBase = m.M
	if v408 < v407 {
		goto L119
	} else {
		goto L120
	}
L87:
	;
	v283 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[5]))
	switch v283 - int32(2) {
	case 0, 18, 22:
		goto L86
	default:
		goto L92
	case 26:
		goto L93
	}
L88:
	;
	goto L89
L89:
	;
	F_on_shmem_exit(m, int32(963), base.I64_extend_i32_u(v279))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L113
	}
L90:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_11), int32(249), int32(_a_F_CreateSharedMemoryAndSemaphores_17))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L112
	}
L91:
	;
	F_errhint(m, v370, int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L111
	}
L92:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L107
	}
L93:
	;
	v286 = int32(0)
	v288 = F_pgmem_shmget(m, v210, v286, int32(1920))
	mBase = m.M
	if v288 < v286 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[5])) = int32(28)
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L104
	}
L95:
	;
	v292 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[5]))
	if base.B2i32(base.Ui32(int32(24)) < base.Ui32(v292))|base.B2i32(int32(1)<<(uint(v292)%32)&int32(17825796) == int32(0)) != 0 {
		goto L94
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v302 = int32(0)
	v304 = F_pgmem_shmctl(m, v288, v302, v302)
	mBase = m.M
	if v302 <= v304 {
		goto L94
	} else {
		goto L99
	}
L98:
	;
	goto L86
L99:
	;
	v309 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	if v309 == int32(0) {
		goto L94
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+132)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+128)) = v288
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_18), v50+int32(128))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_11), int32(210), int32(_a_F_CreateSharedMemoryAndSemaphores_17))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	goto L94
L104:
	;
	F_errmsg(m, int32(_a_F_CreateSharedMemoryAndSemaphores_19), int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+120)) = int32(1920)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+116)) = v203
	*(*int32)(unsafe.Add(mBase, uint32(v50)+112)) = v210
	v345 = F_errdetail(m, int32(_a_F_CreateSharedMemoryAndSemaphores_20), v50+int32(112))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v370 = int32(_a_F_CreateSharedMemoryAndSemaphores_21)
	goto L91
L107:
	;
	F_errmsg(m, int32(_a_F_CreateSharedMemoryAndSemaphores_19), int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+40)) = int32(1920)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+36)) = v203
	*(*int32)(unsafe.Add(mBase, uint32(v50)+32)) = v210
	v363 = F_errdetail(m, int32(_a_F_CreateSharedMemoryAndSemaphores_20), v50+int32(32))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	switch v283 - int32(48) {
	case 0:
		v370 = int32(_a_F_CreateSharedMemoryAndSemaphores_22)
		goto L91
	default:
		goto L90
	case 3:
		goto L110
	}
L110:
	;
	v370 = int32(_a_F_CreateSharedMemoryAndSemaphores_23)
	goto L91
L111:
	;
	goto L90
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	v385 = F_pgmem_shmat(m, v279, v277)
	mBase = m.M
	if v385 == int32(-1) {
		goto L84
	} else {
		goto L114
	}
L114:
	;
	F_on_shmem_exit(m, int32(964), base.I64_extend_i32_u(v385))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+160)) = v210
	*(*int32)(unsafe.Add(mBase, uint32(v50)+164)) = v279
	v395 = v50 + int32(288)
	v399 = F_pg_sprintf(m, v395, int32(_a_F_CreateSharedMemoryAndSemaphores_24), v50+int32(160))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	F_AddToDataDirLockFile(m, int32(7), v395)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	if v385 != 0 {
		goto L85
	} else {
		goto L118
	}
L118:
	;
	goto L86
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+188)) = int32(0)
	goto L83
L120:
	;
	goto L121
L121:
	;
	v413 = int32(0)
	v415 = v50 + int32(188)
	v418 = m.G0
	v420 = v418 - int32(192)
	m.G0 = v420
	*(*int32)(unsafe.Add(mBase, uint32(v415))) = v413
	v424 = int32(2)
	v428 = F_pgmem_shmctl(m, v408, v424, v420+int32(104))
	mBase = m.M
	if v428 < v413 {
		goto L127
	} else {
		goto L128
	}
L122:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v50)+188))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v522)+16))
	if v523 != 0 {
		goto L153
	} else {
		goto L154
	}
L123:
	;
	v506 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L149
	}
L124:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L1
	} else {
		goto L144
	}
L125:
	;
	switch v471 - int32(2) {
	case 0:
		goto L123
	case 1:
		goto L83
	case 2:
		goto L122
	default:
		goto L124
	}
L126:
	;
	m.G0 = v420 + int32(192)
	goto L125
L127:
	;
	v432 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[5]))
	switch v432 - int32(2) {
	case 0:
		goto L131
	default:
		goto L130
	case 22, 26:
		v471 = v424
		goto L126
	}
L128:
	;
	goto L129
L129:
	;
	v437 = int32(0)
	v439 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[1]))
	v442 = F_stat(m, v439, v420+int32(8))
	mBase = m.M
	if v442 < v437 {
		v471 = v437
		goto L126
	} else {
		goto L132
	}
L130:
	;
	v471 = int32(0)
	goto L126
L131:
	;
	v471 = int32(3)
	goto L126
L132:
	;
	v445 = F_pgmem_shmat(m, v408, v413)
	mBase = m.M
	if v445 == int32(-1) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v448 = int32(2)
	v450 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[5]))
	switch v450 - v448 {
	case 0:
		goto L137
	default:
		goto L136
	case 22, 26:
		v471 = v448
		goto L126
	}
L134:
	;
	goto L135
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v415))) = v445
	v456 = int32(3)
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v445)))
	if v457 != int32(679834894) {
		v471 = v456
		goto L126
	} else {
		goto L138
	}
L136:
	;
	v471 = int32(0)
	goto L126
L137:
	;
	v471 = int32(3)
	goto L126
L138:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v445)+20))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v420)+8))
	if v460 != v461 {
		v471 = v456
		goto L126
	} else {
		goto L139
	}
L139:
	;
	v463 = *(*int64)(unsafe.Add(mBase, uint32(v445)+24))
	v464 = *(*int64)(unsafe.Add(mBase, uint32(v420)+96))
	if v463 != v464 {
		v471 = v456
		goto L126
	} else {
		goto L140
	}
L140:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v420)+176))
	if v468 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v469 = int32(1)
	goto L143
L142:
	;
	v469 = int32(4)
	goto L143
L143:
	;
	v471 = v469
	goto L126
L144:
	;
	F_errcode(m, int32(16777238))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+84)) = v408
	*(*int32)(unsafe.Add(mBase, uint32(v50)+80)) = v210
	F_errmsg(m, int32(_a_F_CreateSharedMemoryAndSemaphores_25), v50+int32(80))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	v492 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v50)+64)) = v492
	F_errhint(m, int32(_a_F_CreateSharedMemoryAndSemaphores_26), v50-int32(-64))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_11), int32(804), int32(_a_F_CreateSharedMemoryAndSemaphores_27))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L149:
	;
	if v506 == int32(0) {
		v778 = v210
		goto L82
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+100)) = v408
	*(*int32)(unsafe.Add(mBase, uint32(v50)+96)) = v210
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_28), v50+int32(96))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_11), int32(816), int32(_a_F_CreateSharedMemoryAndSemaphores_27))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	v778 = v210
	goto L82
L153:
	;
	v524 = m.G0
	v526 = v524 - int32(48)
	m.G0 = v526
	v528 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v526)+44)) = v528
	*(*int32)(unsafe.Add(mBase, uint32(v526)+40)) = v528
	*(*int32)(unsafe.Add(mBase, uint32(v526)+36)) = v528
	*(*int32)(unsafe.Add(mBase, uint32(v526)+32)) = v528
	*(*int32)(unsafe.Add(mBase, uint32(v526)+28)) = v528
	*(*int32)(unsafe.Add(mBase, uint32(v526)+24)) = v528
	v549 = F_dsm_impl_op(m, int32(1), v523, v528, v526+int32(36), v526+int32(44), v526+int32(28), int32(14))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v716 = int32(0)
	v718 = F_pgmem_shmctl(m, v408, v716, v716)
	mBase = m.M
	v778 = int32(base.Ui32(v718)>>(uint(int32(31))%32)) + v210
	goto L82
L156:
	;
	if v549 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v551 = int32(2)
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v526)+28))
	if base.Ui32(v552) < base.Ui32(int32(12)) {
		v663 = v551
		goto L160
	} else {
		goto L161
	}
L158:
	;
	goto L159
L159:
	;
	m.G0 = v526 + int32(48)
	goto L155
L160:
	;
	v681 = F_dsm_impl_op(m, v663, v523, int32(0), v526+int32(36), v526+int32(44), v526+int32(28), int32(15))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L185
	}
L161:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v526)+44))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v555)))
	if v556 != int32(-1706017486) {
		v663 = v551
		goto L160
	} else {
		goto L162
	}
L162:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v555)+8))
	if base.Ui64(base.I64_extend_i32_u(v552)) < base.Ui64(base.I64_extend_i32_u(v560)*int64(24)+int64(12)) {
		v663 = v551
		goto L160
	} else {
		goto L163
	}
L163:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v555)+4))
	if base.Ui32(v560) < base.Ui32(v567) {
		v663 = v551
		goto L160
	} else {
		goto L164
	}
L164:
	;
	if v567 != 0 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v577 = int32(0)
	goto L168
L166:
	;
	goto L167
L167:
	;
	v642 = int32(3)
	v645 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L1
	} else {
		goto L181
	}
L168:
	;
	v589 = v555 + int32(12) + v577*int32(24)
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v589)+4))
	if v590 == int32(0) {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	goto L167
L170:
	;
	v625 = v577 + int32(1)
	if v625 != v567 {
		v577 = v625
		goto L168
	} else {
		goto L180
	}
L171:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v589)))
	if v593&int32(1) != 0 {
		goto L170
	} else {
		goto L172
	}
L172:
	;
	v598 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	if v598 != 0 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v526)+20)) = v590
	*(*int32)(unsafe.Add(mBase, uint32(v526)+16)) = v593
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_29), v526+int32(16))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L1
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	v621 = F_dsm_impl_op(m, int32(3), v593, int32(0), v526+int32(32), v526+int32(40), v526+int32(24), int32(15))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L179
	}
L177:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_30), int32(304), int32(_a_F_CreateSharedMemoryAndSemaphores_31))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	goto L176
L179:
	;
	goto L170
L180:
	;
	goto L169
L181:
	;
	if v645 == int32(0) {
		v663 = v642
		goto L160
	} else {
		goto L182
	}
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v526))) = v523
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_32), v526)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_30), int32(314), int32(_a_F_CreateSharedMemoryAndSemaphores_31))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	v663 = v642
	goto L160
L185:
	;
	goto L159
L186:
	;
	m.G0 = v50 + int32(352)
	goto L12
L187:
	;
	v751 = v385
	goto L186
L188:
	;
	goto L189
L189:
	;
	v743 = *(*int64)(unsafe.Add(mBase, uint32(v385)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v740)+24)) = v743
	v745 = *(*int64)(unsafe.Add(mBase, uint32(v385)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v740)+16)) = v745
	v747 = *(*int64)(unsafe.Add(mBase, uint32(v385)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v740)+8)) = v747
	v749 = *(*int64)(unsafe.Add(mBase, uint32(v385)))
	*(*int64)(unsafe.Add(mBase, uint32(v740))) = v749
	v751 = v740
	goto L186
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+152)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+148)) = v277
	*(*int32)(unsafe.Add(mBase, uint32(v50)+144)) = v279
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_33), v50+int32(144))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_11), int32(260), int32(_a_F_CreateSharedMemoryAndSemaphores_17))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L193:
	;
	v793 = F_pgmem_shmdt(m, v790)
	mBase = m.M
	if int32(0) <= v793 {
		v210 = v778
		goto L60
	} else {
		goto L194
	}
L194:
	;
	v798 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	if v798 == int32(0) {
		v210 = v778
		goto L60
	} else {
		goto L196
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+48)) = v790
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_34), v50+int32(48))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_11), int32(841), int32(_a_F_CreateSharedMemoryAndSemaphores_27))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	v210 = v778
	goto L60
L199:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	F_errmsg(m, int32(_a_F_CreateSharedMemoryAndSemaphores_35), int32(0))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_11), int32(734), int32(_a_F_CreateSharedMemoryAndSemaphores_27))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
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
	F_errcode_for_file_access(m)
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	v836 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v836
	F_errmsg(m, int32(_a_F_CreateSharedMemoryAndSemaphores_36), v50)
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_11), int32(720), int32(_a_F_CreateSharedMemoryAndSemaphores_27))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L207:
	;
	v848 = int32(0)
	v850 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[12]))
	if v850 == v848 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v897 = int32(0)
	v899 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[13]))
	if v899 == v897 {
		goto L215
	} else {
		goto L216
	}
L209:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v850)+4))
	if v853 <= int32(0) {
		goto L208
	} else {
		goto L210
	}
L210:
	;
	v857 = v848
	goto L211
L211:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v850)+12))
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v871+v857<<(uint(int32(2))%32))))
	F_InitShmemIndexEntry(m, v875)
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L1
	} else {
		goto L213
	}
L212:
	;
	goto L208
L213:
	;
	v879 = v857 + int32(1)
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v850)+4))
	if v879 < v880 {
		v857 = v879
		goto L211
	} else {
		goto L214
	}
L214:
	;
	goto L212
L215:
	;
	v950 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[12]))
	if v950 != 0 {
		goto L225
	} else {
		goto L226
	}
L216:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v899)+4))
	if v902 <= int32(0) {
		goto L215
	} else {
		goto L217
	}
L217:
	;
	v906 = v897
	goto L218
L218:
	;
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v899)+12))
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v920+v906<<(uint(int32(2))%32))))
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v924)+8))
	if v925 != 0 {
		goto L220
	} else {
		goto L221
	}
L219:
	;
	goto L215
L220:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v924)+16))
	m.T0[v925].(func(*base.Module, int32))(m, v926)
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L1
	} else {
		goto L223
	}
L221:
	;
	goto L222
L222:
	;
	v930 = v906 + int32(1)
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v899)+4))
	if v930 < v931 {
		v906 = v930
		goto L218
	} else {
		goto L224
	}
L223:
	;
	goto L222
L224:
	;
	goto L219
L225:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v950)+4))
	if int32(0) < v951 {
		goto L228
	} else {
		goto L229
	}
L226:
	;
	v1017 = int32(0)
	goto L227
L227:
	;
	F_list_free_deep(m, v1017)
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L1
	} else {
		goto L235
	}
L228:
	;
	v955 = int32(0)
	goto L231
L229:
	;
	goto L230
L230:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[12]))
	v1017 = v1000
	goto L227
L231:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v950)+12))
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v969+v955<<(uint(int32(2))%32))))
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v973)+8))
	v975 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v974)+60)) = uint8(v975)
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v973)))
	F_pfree(m, v977)
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L1
	} else {
		goto L233
	}
L232:
	;
	goto L230
L233:
	;
	v981 = v955 + int32(1)
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v950)+4))
	if v981 < v982 {
		v955 = v981
		goto L231
	} else {
		goto L234
	}
L234:
	;
	goto L232
L235:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[14])) = int32(5)
	v1024 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[12])) = v1024
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v1027 = m.G0
	v1029 = v1027 - int32(1120)
	m.G0 = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v1029)+76)) = v1024
	v1034 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[15]))
	if v1034 == int32(4) {
		goto L238
	} else {
		goto L239
	}
L236:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[16]))
	if v1284 != 0 {
		goto L298
	} else {
		goto L299
	}
L237:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1267 = m.ExcPending
	if v1267 != 0 {
		goto L1
	} else {
		goto L294
	}
L238:
	;
	v1038 = F_AllocateDir(m, int32(_a_F_CreateSharedMemoryAndSemaphores_37))
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L1
	} else {
		goto L241
	}
L239:
	;
	goto L240
L240:
	;
	v1173 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[17]))
	v1177 = v1173*int32(5) - int32(-64)
	v1180 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L1
	} else {
		goto L275
	}
L241:
	;
	v1041 = F_ReadDir(m, v1038, int32(_a_F_CreateSharedMemoryAndSemaphores_37))
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	if v1041 != 0 {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v1043 = v1041
	goto L246
L244:
	;
	goto L245
L245:
	;
	F_FreeDir(m, v1038)
	mBase = m.M
	v1156 = m.ExcPending
	if v1156 != 0 {
		goto L1
	} else {
		goto L274
	}
L246:
	;
	v1059 = v1043 + int32(19)
	v1060 = int32(_a_F_CreateSharedMemoryAndSemaphores_38)
	goto L250
L247:
	;
	goto L245
L248:
	;
	if v1098-v1099 == int32(0) {
		goto L261
	} else {
		goto L262
	}
L250:
	;
	goto L251
L251:
	;
	v1067 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1059))))
	if v1067 != 0 {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v1068 = v1059
	v1069 = v1060
	v1070 = int32(5)
	v1071 = v1067
	goto L256
L253:
	;
	v1094 = v1060
	v1098 = int32(0)
	goto L254
L254:
	;
	v1099 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1094))))
	goto L248
L255:
	;
	v1094 = v1089
	v1098 = v1091
	goto L254
L256:
	;
	v1073 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1069))))
	if base.B2i32(v1071 != v1073)|base.B2i32(v1073 == int32(0)) != 0 {
		v1089 = v1069
		v1091 = v1071
		goto L255
	} else {
		goto L258
	}
L257:
	;
	v1089 = v1083
	v1091 = int32(0)
	goto L255
L258:
	;
	v1079 = v1070 - int32(1)
	if v1079 == int32(0) {
		v1089 = v1069
		v1091 = v1071
		goto L255
	} else {
		goto L259
	}
L259:
	;
	v1082 = int32(1)
	v1083 = v1069 + v1082
	v1084 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1068)+1)))
	if v1084 != 0 {
		v1068 = v1068 + v1082
		v1069 = v1083
		v1070 = v1079
		v1071 = v1084
		goto L256
	} else {
		goto L260
	}
L260:
	;
	goto L257
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1029)+64)) = v1059
	v1111 = v1029 + int32(80)
	v1116 = F_pg_snprintf(m, v1111, int32(1036), int32(_a_F_CreateSharedMemoryAndSemaphores_39), v1029-int32(-64))
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L1
	} else {
		goto L264
	}
L262:
	;
	goto L263
L263:
	;
	v1138 = F_ReadDir(m, v1038, int32(_a_F_CreateSharedMemoryAndSemaphores_37))
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L1
	} else {
		goto L272
	}
L264:
	;
	v1120 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L1
	} else {
		goto L265
	}
L265:
	;
	if v1120 != 0 {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1029)+48)) = v1111
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_40), v1029+int32(48))
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L1
	} else {
		goto L269
	}
L267:
	;
	goto L268
L268:
	;
	v1135 = F_unlink(m, v1029+int32(80))
	mBase = m.M
	if v1135 != 0 {
		goto L237
	} else {
		goto L271
	}
L269:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_30), int32(347), int32(_a_F_CreateSharedMemoryAndSemaphores_41))
	mBase = m.M
	v1132 = m.ExcPending
	if v1132 != 0 {
		goto L1
	} else {
		goto L270
	}
L270:
	;
	goto L268
L271:
	;
	goto L263
L272:
	;
	if v1138 != 0 {
		v1043 = v1138
		goto L246
	} else {
		goto L273
	}
L273:
	;
	goto L247
L274:
	;
	goto L240
L275:
	;
	if v1180 != 0 {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1029)+16)) = v1177
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_42), v1029+int32(16))
	mBase = m.M
	v1187 = m.ExcPending
	if v1187 != 0 {
		goto L1
	} else {
		goto L279
	}
L277:
	;
	goto L278
L278:
	;
	v1196 = v1177*int32(24) + int32(12)
	goto L281
L279:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_30), int32(208), int32(_a_F_CreateSharedMemoryAndSemaphores_43))
	mBase = m.M
	v1192 = m.ExcPending
	if v1192 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	goto L278
L281:
	;
	v1214 = Fn14349(m, int64(32))
	mBase = m.M
	goto L283
L282:
	;
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v1029)+76))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[18])) = v1231
	F_on_shmem_exit(m, int32(1186), base.I64_extend_i32_u(v1026))
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L1
	} else {
		goto L287
	}
L283:
	;
	v1216 = v1214 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[19])) = v1216
	if v1216 == int32(0) {
		goto L281
	} else {
		goto L284
	}
L284:
	;
	v1226 = F_dsm_impl_op(m, int32(0), v1216, v1196, int32(_a_F_CreateSharedMemoryAndSemaphores_44), v1029+int32(76), int32(_a_F_CreateSharedMemoryAndSemaphores_45), int32(21))
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L1
	} else {
		goto L285
	}
L285:
	;
	if v1226 == int32(0) {
		goto L281
	} else {
		goto L286
	}
L286:
	;
	goto L282
L287:
	;
	v1239 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1240 = m.ExcPending
	if v1240 != 0 {
		goto L1
	} else {
		goto L288
	}
L288:
	;
	if v1239 != 0 {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1029)+4)) = v1196
	v1243 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v1029))) = v1243
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_46), v1029)
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		goto L1
	} else {
		goto L292
	}
L290:
	;
	goto L291
L291:
	;
	v1254 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v1026)+16)) = v1254
	v1257 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[18]))
	*(*int32)(unsafe.Add(mBase, uint32(v1257)+8)) = v1177
	*(*int64)(unsafe.Add(mBase, uint32(v1257))) = int64(2588949810)
	m.G0 = v1029 + int32(1120)
	goto L236
L292:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_30), int32(233), int32(_a_F_CreateSharedMemoryAndSemaphores_43))
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		goto L1
	} else {
		goto L293
	}
L293:
	;
	goto L291
L294:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1269 = m.ExcPending
	if v1269 != 0 {
		goto L1
	} else {
		goto L295
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1029)+32)) = v1029 + int32(80)
	F_errmsg(m, int32(_a_F_CreateSharedMemoryAndSemaphores_47), v1029+int32(32))
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L1
	} else {
		goto L296
	}
L296:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_30), int32(353), int32(_a_F_CreateSharedMemoryAndSemaphores_41))
	mBase = m.M
	v1282 = m.ExcPending
	if v1282 != 0 {
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
	m.T0[v1284].(func(*base.Module))(m)
	mBase = m.M
	v1286 = m.ExcPending
	if v1286 != 0 {
		goto L1
	} else {
		goto L301
	}
L299:
	;
	goto L300
L300:
	;
	m.G0 = v18 + int32(16)
	return
L301:
	;
	goto L300
}
func F_DeleteSharedComments(m *base.Module, l0 int32, l1 int32) {
	var v6 int32
	_ = v6
	Fn14211(m, l0, l1, int32(2397), int32(2396))
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	v5 = int32(44)
	v6 = l0 + v5
	v9 = base.AtomicRmwXchg32(m, l0, v5, int32(1))
	if v9 != 0 {
		F_s_lock(m, v6, int32(_a_F_SharedFileSetAttach_0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			if v13 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v13 + int32(1)
				v17 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0)+44)), uint32(v17))
				F_on_dsm_detach(m, l1, int32(1185), base.I64_extend_i32_u(l0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					return
				}
			} else {
				v24 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v6))), uint32(v24))
				F_errstart_cold(m, int32(21), v24)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_SharedFileSetAttach_1), int32(0))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_SharedFileSetAttach_2), int32(73), int32(_a_F_SharedFileSetAttach_3))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
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
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		if v13 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v13 + int32(1)
			v17 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0)+44)), uint32(v17))
			F_on_dsm_detach(m, l1, int32(1185), base.I64_extend_i32_u(l0))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				return
			}
		} else {
			v24 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v6))), uint32(v24))
			F_errstart_cold(m, int32(21), v24)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_SharedFileSetAttach_1), int32(0))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_SharedFileSetAttach_2), int32(73), int32(_a_F_SharedFileSetAttach_3))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
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
func F_SharedInvalShmemRequest(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	v3 = m.G0
	v4 = int32(16)
	v5 = v3 - v4
	m.G0 = v5
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_SharedInvalShmemRequest[0]))
	v13 = F_mul_size(m, v4, v10+int32(38))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v15 = F_add_size(m, int32(_a_F_SharedInvalShmemRequest_0), v13)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_SharedInvalShmemRequest[0]))
			v22 = F_mul_size(m, int32(4), v19+int32(38))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				v24 = F_add_size(m, v15, v22)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = int32(_a_F_SharedInvalShmemRequest_1)
					*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v24
					*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_SharedInvalShmemRequest_2)
					F_ShmemRequestStructWithOpts(m, v5)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						m.G0 = v5 + int32(16)
						return
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
