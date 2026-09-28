package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__brin_end_parallel(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v34 int32
	_ = v34
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v70 int32
	_ = v70
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v76 int32
	_ = v76
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v82 int32
	_ = v82
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v88 int32
	_ = v88
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v94 int32
	_ = v94
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v100 int32
	_ = v100
	var v102 int64
	_ = v102
	var v103 int64
	_ = v103
	var v106 int32
	_ = v106
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v112 int32
	_ = v112
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v118 int32
	_ = v118
	var v120 int64
	_ = v120
	var v121 int64
	_ = v121
	var v124 int32
	_ = v124
	var v126 int64
	_ = v126
	var v127 int64
	_ = v127
	var v130 int32
	_ = v130
	var v132 int64
	_ = v132
	var v133 int64
	_ = v133
	var v136 int32
	_ = v136
	var v138 int64
	_ = v138
	var v139 int64
	_ = v139
	var v142 int32
	_ = v142
	var v144 int64
	_ = v144
	var v145 int64
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_WaitForParallelWorkersToFinish(m, v4)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
		if int32(0) < v8 {
			v12 = int32(0)
			for {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v17 = v14 + v12<<(uint(int32(7))%32)
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v21 = v18 + v12*int32(40)
				v22 = int32(_a_F__brin_end_parallel_0)
				v24 = *(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[0]))
				v25 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
				*(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[0])) = v24 + v25
				v28 = int32(_a_F__brin_end_parallel_1)
				v30 = *(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[1]))
				v31 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
				*(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[1])) = v30 + v31
				v34 = int32(_a_F__brin_end_parallel_2)
				v36 = *(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[2]))
				v37 = *(*int64)(unsafe.Add(mBase, uint32(v17)+16))
				*(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[2])) = v36 + v37
				v40 = int32(_a_F__brin_end_parallel_3)
				v42 = *(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[3]))
				v43 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
				*(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[3])) = v42 + v43
				v46 = int32(_a_F__brin_end_parallel_4)
				v48 = *(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[4]))
				v49 = *(*int64)(unsafe.Add(mBase, uint32(v17)+32))
				*(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[4])) = v48 + v49
				v52 = int32(_a_F__brin_end_parallel_5)
				v54 = *(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[5]))
				v55 = *(*int64)(unsafe.Add(mBase, uint32(v17)+40))
				*(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[5])) = v54 + v55
				v58 = int32(_a_F__brin_end_parallel_6)
				v60 = *(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[6]))
				v61 = *(*int64)(unsafe.Add(mBase, uint32(v17)+48))
				*(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[6])) = v60 + v61
				v64 = int32(_a_F__brin_end_parallel_7)
				v66 = *(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[7]))
				v67 = *(*int64)(unsafe.Add(mBase, uint32(v17)+56))
				*(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[7])) = v66 + v67
				v70 = int32(_a_F__brin_end_parallel_8)
				v72 = *(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[8]))
				v73 = *(*int64)(unsafe.Add(mBase, uint32(v17)+64))
				*(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[8])) = v72 + v73
				v76 = int32(_a_F__brin_end_parallel_9)
				v78 = *(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[9]))
				v79 = *(*int64)(unsafe.Add(mBase, uint32(v17)+72))
				*(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[9])) = v78 + v79
				v82 = int32(_a_F__brin_end_parallel_10)
				v84 = *(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[10]))
				v85 = *(*int64)(unsafe.Add(mBase, uint32(v17)+80))
				*(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[10])) = v84 + v85
				v88 = int32(_a_F__brin_end_parallel_11)
				v90 = *(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[11]))
				v91 = *(*int64)(unsafe.Add(mBase, uint32(v17)+88))
				*(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[11])) = v90 + v91
				v94 = int32(_a_F__brin_end_parallel_12)
				v96 = *(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[12]))
				v97 = *(*int64)(unsafe.Add(mBase, uint32(v17)+96))
				*(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[12])) = v96 + v97
				v100 = int32(_a_F__brin_end_parallel_13)
				v102 = *(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[13]))
				v103 = *(*int64)(unsafe.Add(mBase, uint32(v17)+104))
				*(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[13])) = v102 + v103
				v106 = int32(_a_F__brin_end_parallel_14)
				v108 = *(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[14]))
				v109 = *(*int64)(unsafe.Add(mBase, uint32(v17)+112))
				*(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[14])) = v108 + v109
				v112 = int32(_a_F__brin_end_parallel_15)
				v114 = *(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[15]))
				v115 = *(*int64)(unsafe.Add(mBase, uint32(v17)+120))
				*(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[15])) = v114 + v115
				v118 = int32(_a_F__brin_end_parallel_16)
				v120 = *(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[16]))
				v121 = *(*int64)(unsafe.Add(mBase, uint32(v21)+16))
				*(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[16])) = v120 + v121
				v124 = int32(_a_F__brin_end_parallel_17)
				v126 = *(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[17]))
				v127 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
				*(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[17])) = v126 + v127
				v130 = int32(_a_F__brin_end_parallel_18)
				v132 = *(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[18]))
				v133 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
				*(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[18])) = v132 + v133
				v136 = int32(_a_F__brin_end_parallel_19)
				v138 = *(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[19]))
				v139 = *(*int64)(unsafe.Add(mBase, uint32(v21)+24))
				*(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[19])) = v138 + v139
				v142 = int32(_a_F__brin_end_parallel_20)
				v144 = *(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[20]))
				v145 = *(*int64)(unsafe.Add(mBase, uint32(v21)+32))
				*(*int64)(unsafe.Add(mBase, _c_F__brin_end_parallel[20])) = v144 + v145
				v149 = v12 + int32(1)
				v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+20))
				if v149 < v151 {
					v12 = v149
					continue
				} else {
					break
				}
				break
			}
			v155 = v150
		} else {
			v155 = v7
		}
		v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
		if v157 != 0 {
			v161 = v155
			F_DestroyParallelContext(m, v161)
			mBase = m.M
			v163 = m.ExcPending
			if v163 != 0 {
				return
			} else {
				v166 = *(*int32)(unsafe.Add(mBase, _c_F__brin_end_parallel[21]))
				v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+72))
				*(*int32)(unsafe.Add(mBase, uint32(v166)+72)) = v167 - int32(1)
				return
			}
		} else {
			F_UnregisterSnapshot(m, v156)
			mBase = m.M
			v159 = m.ExcPending
			if v159 != 0 {
				return
			} else {
				v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v161 = v160
				F_DestroyParallelContext(m, v161)
				mBase = m.M
				v163 = m.ExcPending
				if v163 != 0 {
					return
				} else {
					v166 = *(*int32)(unsafe.Add(mBase, _c_F__brin_end_parallel[21]))
					v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+72))
					*(*int32)(unsafe.Add(mBase, uint32(v166)+72)) = v167 - int32(1)
					return
				}
			}
		}
	}
}
func F_brin_deform_tuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v327 int64
	_ = v327
	var v328 int64
	_ = v328
	var v329 int64
	_ = v329
	var v330 int64
	_ = v330
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v345 int64
	_ = v345
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v412 int32
	_ = v412
	var v419 int32
	_ = v419
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v503 int64
	_ = v503
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int64
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v580 int32
	_ = v580
	v4 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(16)
	m.G0 = v22
	if l2 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v131&int32(64) != 0 {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	F_MemoryContextReset(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v110 = F_brin_new_memtuple(m, l0)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L5
	} else {
		goto L13
	}
L5:
	;
	return int32(0)
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if int32(0) < v30 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v35 = int32(24)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v49 = v4
	v51 = l2 + (v30*v35+int32(31))&int32(-8)
	goto L10
L8:
	;
	goto L9
L9:
	;
	v108 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)) = uint8(v108)
	v114 = l2
	goto L1
L10:
	;
	v66 = l2 + v35 + v49*int32(24)
	v68 = v49 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v66))) = uint16(v68)
	*(*int32)(unsafe.Add(mBase, uint32(v66)+20)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v66)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = v51
	v75 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v66)+2)) = uint16(v75)
	*(*int32)(unsafe.Add(mBase, uint32(v66)+16)) = v44
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(20)+v49<<(uint(int32(2))%32))))
	v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81))))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	if v68 < v87 {
		v49 = v68
		v51 = v51 + v82<<(uint(int32(3))%32)
		goto L10
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	goto L11
L13:
	;
	v114 = v110
	goto L1
L14:
	;
	v134 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v114))) = uint8(v134)
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	v137 = v136
	goto L16
L15:
	;
	v137 = v131
	goto L16
L16:
	;
	if v137&int32(32) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v142 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v114)+1)) = uint8(v142)
	goto L19
L18:
	;
	goto L19
L19:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v114)+4)) = v144
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v114)+20))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v114)+16))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	if int32(0) < v151 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v155 = l1 + int32(5)
	v156 = int32(0)
	v163 = v156
	goto L23
L21:
	;
	goto L22
L22:
	;
	v236 = F_brtuple_disk_tupdesc(m, l0)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L5
	} else {
		goto L30
	}
L23:
	;
	if base.B2i32(v156 <= base.I32_extend8_s(v146)) == int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L22
L25:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v163+v147))) = uint8(v210)
	v213 = v163 + int32(1)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
	if v213 < v215 {
		v163 = v213
		goto L23
	} else {
		goto L29
	}
L26:
	;
	v183 = int32(3)
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155+int32(base.Ui32(v163)>>(uint(v183)%32))))))
	v187 = int32(7)
	v190 = int32(1)
	v191 = int32(base.Ui32(v186)>>(uint(v163&v187)%32)) & v190
	*(*uint8)(unsafe.Add(mBase, uint32(v163+v148))) = uint8(v191)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	v195 = v194 + v163
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155+v195>>(uint(v183)%32)))))
	v210 = int32(base.Ui32(v199)>>(uint(v195&v187)%32)) & v190
	goto L25
L27:
	;
	goto L28
L28:
	;
	v206 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v163+v148))) = uint8(v206)
	v210 = v206
	goto L25
L29:
	;
	goto L24
L30:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	if int32(0) < v239 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v244 = l1 + v146&int32(31)
	v249 = int32(0)
	v254 = v249
	v256 = v238
	v257 = v249
	v260 = v239
	v261 = v4
	goto L34
L32:
	;
	v419 = v238
	goto L33
L33:
	;
	v433 = int32(_a_F_brin_deform_tuple_0)
	v434 = *(*int32)(unsafe.Add(mBase, _c_F_brin_deform_tuple[0]))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v114)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_brin_deform_tuple[0])) = v436
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v419)))
	if int32(0) < v438 {
		goto L80
	} else {
		goto L81
	}
L34:
	;
	v272 = l0 + int32(20) + v261<<(uint(int32(2))%32)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v272)))
	v274 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v273))))
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148+v261))))
	if v276 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v419 = v397
	goto L33
L36:
	;
	v412 = v261 + int32(1)
	if v412 < v401 {
		v254 = v395
		v256 = v397
		v257 = v398
		v260 = v401
		v261 = v412
		goto L34
	} else {
		goto L79
	}
L37:
	;
	v279 = int32(0)
	if v274 == v279 {
		v395 = v254
		v397 = v256
		v398 = v257
		v401 = v260
		goto L36
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v395 = v254
	v397 = v256
	v398 = v274 + v257
	v401 = v260
	goto L36
L40:
	;
	v285 = v254
	v288 = v257
	v289 = v279
	goto L41
L41:
	;
	v302 = v288 << (uint(int32(3)) % 32)
	v303 = v236 + int32(28) + v302
	v304 = int32(*(*int16)(unsafe.Add(mBase, uint32(v303)+2)))
	if v304 == int32(-1) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v389)))
	v395 = v383
	v397 = v389
	v398 = v382
	v401 = v390
	goto L36
L43:
	;
	v318 = v317 + v244
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303)+4)))
	if v320 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L44:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285+v244))))
	if v308 != 0 {
		v317 = v285
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303)+5)))
	v317 = (v285 + v309 - int32(1)) & (int32(0) - v309)
	goto L43
L47:
	;
	goto L46
L48:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v302+v149))) = v345
	v347 = int32(*(*int16)(unsafe.Add(mBase, uint32(v303)+2)))
	if int32(0) < v347 {
		v380 = v347
		goto L61
	} else {
		goto L62
	}
L49:
	;
	if base.I32_popcnt(v304) != int32(1) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	v345 = base.I64_extend_i32_u(v318)
	goto L48
L52:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L5
	} else {
		goto L58
	}
L53:
	;
	switch base.I32_ctz(v304) {
	case 0:
		goto L57
	case 1:
		goto L56
	case 2:
		goto L55
	case 3:
		goto L54
	default:
		goto L52
	}
L54:
	;
	v330 = *(*int64)(unsafe.Add(mBase, uint32(v318)))
	v345 = v330
	goto L48
L55:
	;
	v329 = int64(*(*int32)(unsafe.Add(mBase, uint32(v318))))
	v345 = v329
	goto L48
L56:
	;
	v328 = int64(*(*int16)(unsafe.Add(mBase, uint32(v318))))
	v345 = v328
	goto L48
L57:
	;
	v327 = int64(*(*int8)(unsafe.Add(mBase, uint32(v318))))
	v345 = v327
	goto L48
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v304
	F_errmsg_internal(m, int32(_a_F_brin_deform_tuple_1), v22)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L5
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_brin_deform_tuple_2), int32(123), int32(_a_F_brin_deform_tuple_3))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L5
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	v381 = int32(1)
	v382 = v288 + v381
	v383 = v317 + v380
	v385 = v289 + v381
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v272)))
	v387 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v386))))
	if base.Ui32(v385) < base.Ui32(v387) {
		v285 = v383
		v288 = v382
		v289 = v385
		goto L41
	} else {
		goto L78
	}
L62:
	;
	if v347 == int32(-1) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318))))
	if v352 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	v376 = F_strlen(m, v318)
	mBase = m.M
	v380 = v376 + int32(1)
	goto L61
L66:
	;
	v356 = int32(18)
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318)+1)))
	if v358 == v356 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L68
L68:
	;
	if v352&int32(1) != 0 {
		goto L75
	} else {
		goto L76
	}
L69:
	;
	v361 = v356
	goto L71
L70:
	;
	v361 = int32(2)
	goto L71
L71:
	;
	if base.Ui32((v358-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v368 = int32(6)
	goto L74
L73:
	;
	v368 = v361
	goto L74
L74:
	;
	v380 = v368
	goto L61
L75:
	;
	v380 = int32(base.Ui32(v352) >> (uint(int32(1)) % 32))
	goto L61
L76:
	;
	goto L77
L77:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v318)))
	v380 = int32(base.Ui32(v373) >> (uint(int32(2)) % 32))
	goto L61
L78:
	;
	goto L42
L79:
	;
	goto L35
L80:
	;
	v445 = int32(0)
	v450 = v436
	v451 = v445
	v452 = v445
	v454 = v438
	goto L83
L81:
	;
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_brin_deform_tuple[0])) = v434
	m.G0 = v22 + int32(16)
	return v114
L83:
	;
	v468 = l0 + int32(20) + v452<<(uint(int32(2))%32)
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v468)))
	v470 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v469))))
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v452+v148))))
	if v472 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L82
L85:
	;
	v580 = v452 + int32(1)
	if v580 < v567 {
		v450 = v563
		v451 = v564
		v452 = v580
		v454 = v567
		goto L83
	} else {
		goto L96
	}
L86:
	;
	if v470 != 0 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	v563 = v450
	v564 = v470 + v451
	v567 = v454
	goto L85
L89:
	;
	v484 = int32(0)
	v485 = v451
	v487 = v469
	goto L92
L90:
	;
	v528 = v450
	v529 = v451
	goto L91
L91:
	;
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v452+v147))))
	v548 = v114 + int32(24) + v452*int32(24)
	v549 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v548)+20)) = v549
	*(*int64)(unsafe.Add(mBase, uint32(v548)+8)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v548)+3)) = uint8(v549)
	*(*uint8)(unsafe.Add(mBase, uint32(v548)+2)) = uint8(v545)
	*(*int32)(unsafe.Add(mBase, uint32(v548)+16)) = v528
	v557 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v557)))
	v563 = v528
	v564 = v529
	v567 = v558
	goto L85
L92:
	;
	v503 = *(*int64)(unsafe.Add(mBase, uint32(v149+v485<<(uint(int32(3))%32))))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v487+v484<<(uint(int32(2))%32))+8))
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+10)))
	v509 = int32(*(*int16)(unsafe.Add(mBase, uint32(v507)+8)))
	v510 = F_datumCopy(m, v503, v508, v509)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L5
	} else {
		goto L94
	}
L93:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v114)+8))
	v528 = v524
	v529 = v518
	goto L91
L94:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v114+v452*int32(24)+int32(28))))
	*(*int64)(unsafe.Add(mBase, uint32(v512+v484<<(uint(int32(3))%32)))) = v510
	v517 = int32(1)
	v518 = v485 + v517
	v520 = v484 + v517
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v468)))
	v522 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v521))))
	if base.Ui32(v520) < base.Ui32(v522) {
		v484 = v520
		v485 = v518
		v487 = v521
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	goto L84
}
func F_brin_initialize_empty_new_buffer(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	v3 = int32(0)
	v6 = int32(_a_F_brin_initialize_empty_new_buffer_0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_brin_initialize_empty_new_buffer[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_brin_initialize_empty_new_buffer[0])) = v8 + int32(1)
	if l1 < v3 {
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_brin_initialize_empty_new_buffer[1]))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v15+(l1^int32(-1))<<(uint(int32(2))%32))))
		v29 = v21
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, _c_F_brin_initialize_empty_new_buffer[2]))
		v29 = v23 + l1<<(uint(int32(13))%32) + int32(-8192)
	}
	v30 = int32(_a_F_brin_initialize_empty_new_buffer_1)
	v32 = int32(0)
	if v32|(v29&int32(3)|int32(1)) == v32 {
		v48 = v29 + v30
		v50 = v29 + int32(4)
		if base.Ui32(v50) < base.Ui32(v48) {
			v52 = v48
		} else {
			v52 = v50
		}
		v57 = (v29^int32(-1)+v52)&int32(-4) + int32(4)
		if v57 == int32(0) {
		} else {
			base.MemoryFill(m, v29, int32(0), v57)
		}
	} else {
		base.MemoryFill(m, v29, int32(0), v30)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v29)+10)) = int32(_a_F_brin_initialize_empty_new_buffer_2)
	v71 = int32(_a_F_brin_initialize_empty_new_buffer_3)
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+18)) = uint16(v71)
	v77 = int32(_a_F_brin_initialize_empty_new_buffer_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+16)) = uint16(v77)
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+14)) = uint16(v77)
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+16)))
	v82 = int32(_a_F_brin_initialize_empty_new_buffer_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v29+v80)+6)) = uint16(v82)
	F_MarkBufferDirty(m, l1)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		return
	} else {
		v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+118)))
		if v87 != int32(112) {
			v99 = int32(_a_F_brin_initialize_empty_new_buffer_0)
			v101 = *(*int32)(unsafe.Add(mBase, _c_F_brin_initialize_empty_new_buffer[0]))
			*(*int32)(unsafe.Add(mBase, _c_F_brin_initialize_empty_new_buffer[0])) = v101 - int32(1)
			if l1 < int32(0) {
				v108 = *(*int32)(unsafe.Add(mBase, _c_F_brin_initialize_empty_new_buffer[3]))
				v114 = *(*int32)(unsafe.Add(mBase, uint32(v108+(l1^int32(-1))*int32(56))+16))
				v123 = v114
			} else {
				v116 = *(*int32)(unsafe.Add(mBase, _c_F_brin_initialize_empty_new_buffer[4]))
				v117 = int32(56)
				v122 = *(*int32)(unsafe.Add(mBase, uint32(v116+l1*v117-v117)+16))
				v123 = v122
			}
			v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+16)))
			v125 = v29 + v124
			v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v125)+6)))
			if v126 != int32(_a_F_brin_initialize_empty_new_buffer_5) {
				v141 = v3
			} else {
				v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+4)))
				if v129&int32(1) != 0 {
					v141 = v3
				} else {
					v132 = int32(4)
					v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+14)))
					v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+12)))
					v135 = v133 - v134
					if v135 <= v132 {
						v138 = v132
					} else {
						v138 = v135
					}
					v141 = v138 - int32(4)
				}
			}
			F_RecordPageWithFreeSpace(m, l0, v123, v141)
			mBase = m.M
			v143 = m.ExcPending
			if v143 != 0 {
				return
			} else {
				return
			}
		} else {
			v91 = *(*int32)(unsafe.Add(mBase, _c_F_brin_initialize_empty_new_buffer[5]))
			if v91 <= int32(0) {
				v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				if v94 != 0 {
					v99 = int32(_a_F_brin_initialize_empty_new_buffer_0)
					v101 = *(*int32)(unsafe.Add(mBase, _c_F_brin_initialize_empty_new_buffer[0]))
					*(*int32)(unsafe.Add(mBase, _c_F_brin_initialize_empty_new_buffer[0])) = v101 - int32(1)
					if l1 < int32(0) {
						v108 = *(*int32)(unsafe.Add(mBase, _c_F_brin_initialize_empty_new_buffer[3]))
						v114 = *(*int32)(unsafe.Add(mBase, uint32(v108+(l1^int32(-1))*int32(56))+16))
						v123 = v114
					} else {
						v116 = *(*int32)(unsafe.Add(mBase, _c_F_brin_initialize_empty_new_buffer[4]))
						v117 = int32(56)
						v122 = *(*int32)(unsafe.Add(mBase, uint32(v116+l1*v117-v117)+16))
						v123 = v122
					}
					v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+16)))
					v125 = v29 + v124
					v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v125)+6)))
					if v126 != int32(_a_F_brin_initialize_empty_new_buffer_5) {
						v141 = v3
					} else {
						v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+4)))
						if v129&int32(1) != 0 {
							v141 = v3
						} else {
							v132 = int32(4)
							v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+14)))
							v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+12)))
							v135 = v133 - v134
							if v135 <= v132 {
								v138 = v132
							} else {
								v138 = v135
							}
							v141 = v138 - int32(4)
						}
					}
					F_RecordPageWithFreeSpace(m, l0, v123, v141)
					mBase = m.M
					v143 = m.ExcPending
					if v143 != 0 {
						return
					} else {
						return
					}
				} else {
					v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v95 != 0 {
						v99 = int32(_a_F_brin_initialize_empty_new_buffer_0)
						v101 = *(*int32)(unsafe.Add(mBase, _c_F_brin_initialize_empty_new_buffer[0]))
						*(*int32)(unsafe.Add(mBase, _c_F_brin_initialize_empty_new_buffer[0])) = v101 - int32(1)
						if l1 < int32(0) {
							v108 = *(*int32)(unsafe.Add(mBase, _c_F_brin_initialize_empty_new_buffer[3]))
							v114 = *(*int32)(unsafe.Add(mBase, uint32(v108+(l1^int32(-1))*int32(56))+16))
							v123 = v114
						} else {
							v116 = *(*int32)(unsafe.Add(mBase, _c_F_brin_initialize_empty_new_buffer[4]))
							v117 = int32(56)
							v122 = *(*int32)(unsafe.Add(mBase, uint32(v116+l1*v117-v117)+16))
							v123 = v122
						}
						v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+16)))
						v125 = v29 + v124
						v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v125)+6)))
						if v126 != int32(_a_F_brin_initialize_empty_new_buffer_5) {
							v141 = v3
						} else {
							v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+4)))
							if v129&int32(1) != 0 {
								v141 = v3
							} else {
								v132 = int32(4)
								v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+14)))
								v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+12)))
								v135 = v133 - v134
								if v135 <= v132 {
									v138 = v132
								} else {
									v138 = v135
								}
								v141 = v138 - int32(4)
							}
						}
						F_RecordPageWithFreeSpace(m, l0, v123, v141)
						mBase = m.M
						v143 = m.ExcPending
						if v143 != 0 {
							return
						} else {
							return
						}
					} else {
						F_log_newpage_buffer(m, l1, int32(1))
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return
						} else {
							v99 = int32(_a_F_brin_initialize_empty_new_buffer_0)
							v101 = *(*int32)(unsafe.Add(mBase, _c_F_brin_initialize_empty_new_buffer[0]))
							*(*int32)(unsafe.Add(mBase, _c_F_brin_initialize_empty_new_buffer[0])) = v101 - int32(1)
							if l1 < int32(0) {
								v108 = *(*int32)(unsafe.Add(mBase, _c_F_brin_initialize_empty_new_buffer[3]))
								v114 = *(*int32)(unsafe.Add(mBase, uint32(v108+(l1^int32(-1))*int32(56))+16))
								v123 = v114
							} else {
								v116 = *(*int32)(unsafe.Add(mBase, _c_F_brin_initialize_empty_new_buffer[4]))
								v117 = int32(56)
								v122 = *(*int32)(unsafe.Add(mBase, uint32(v116+l1*v117-v117)+16))
								v123 = v122
							}
							v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+16)))
							v125 = v29 + v124
							v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v125)+6)))
							if v126 != int32(_a_F_brin_initialize_empty_new_buffer_5) {
								v141 = v3
							} else {
								v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+4)))
								if v129&int32(1) != 0 {
									v141 = v3
								} else {
									v132 = int32(4)
									v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+14)))
									v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+12)))
									v135 = v133 - v134
									if v135 <= v132 {
										v138 = v132
									} else {
										v138 = v135
									}
									v141 = v138 - int32(4)
								}
							}
							F_RecordPageWithFreeSpace(m, l0, v123, v141)
							mBase = m.M
							v143 = m.ExcPending
							if v143 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			} else {
				F_log_newpage_buffer(m, l1, int32(1))
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return
				} else {
					v99 = int32(_a_F_brin_initialize_empty_new_buffer_0)
					v101 = *(*int32)(unsafe.Add(mBase, _c_F_brin_initialize_empty_new_buffer[0]))
					*(*int32)(unsafe.Add(mBase, _c_F_brin_initialize_empty_new_buffer[0])) = v101 - int32(1)
					if l1 < int32(0) {
						v108 = *(*int32)(unsafe.Add(mBase, _c_F_brin_initialize_empty_new_buffer[3]))
						v114 = *(*int32)(unsafe.Add(mBase, uint32(v108+(l1^int32(-1))*int32(56))+16))
						v123 = v114
					} else {
						v116 = *(*int32)(unsafe.Add(mBase, _c_F_brin_initialize_empty_new_buffer[4]))
						v117 = int32(56)
						v122 = *(*int32)(unsafe.Add(mBase, uint32(v116+l1*v117-v117)+16))
						v123 = v122
					}
					v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+16)))
					v125 = v29 + v124
					v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v125)+6)))
					if v126 != int32(_a_F_brin_initialize_empty_new_buffer_5) {
						v141 = v3
					} else {
						v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+4)))
						if v129&int32(1) != 0 {
							v141 = v3
						} else {
							v132 = int32(4)
							v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+14)))
							v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+12)))
							v135 = v133 - v134
							if v135 <= v132 {
								v138 = v132
							} else {
								v138 = v135
							}
							v141 = v138 - int32(4)
						}
					}
					F_RecordPageWithFreeSpace(m, l0, v123, v141)
					mBase = m.M
					v143 = m.ExcPending
					if v143 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_brin_minmax_multi_distance_int2(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
	v4 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	return base.I64_reinterpret_f64(base.F64_sub(base.F64_convert_i32_s(v2), base.F64_convert_i32_s(v4)))
}
func F_brin_minmax_multi_distance_numeric(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	v3 = int32(0)
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_DirectFunctionCall2Coll(m, int32(19), v3, v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = F_DirectFunctionCall1Coll(m, int32(18), v3, v8)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			return v12
		}
	}
}
func F_brin_minmax_multi_distance_pg_lsn(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	return base.I64_reinterpret_f64(base.F64_convert_i64_u(v2 - v3))
}
func F_brin_minmax_multi_distance_time(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	return base.I64_reinterpret_f64(base.F64_convert_i64_s(v2 - v3))
}
func F_brin_minmax_union(m *base.Module, l0 int32) int64 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int64
	_ = v83
	var v86 int32
	_ = v86
	var v87 int64
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v11 = int32(1)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v25 = v15 + v16<<(uint(int32(3))%32) + base.I32_extend16_s(v14)*int32(100) - int32(72)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+68))
	v28 = F_minmax_get_strategy_procinfo(m, v12, v14, v26, v11)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		return int64(0)
	} else {
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		v33 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
		v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		v35 = *(*int64)(unsafe.Add(mBase, uint32(v34)))
		v36 = F_FunctionCall2Coll(m, v28, v9, v33, v35)
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int64(0)
		} else {
			if v36 != int64(0) {
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+82)))
				if v40 == int32(0) {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
					F_pfree(m, v44)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int64(0)
					} else {
						v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+82)))
						v48 = v47
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
						v50 = *(*int64)(unsafe.Add(mBase, uint32(v49)))
						v53 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+72)))
						v54 = F_datumCopy(m, v50, v48&int32(1), v53)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int64(0)
						} else {
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
							*(*int64)(unsafe.Add(mBase, uint32(v56))) = v54
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v25)+68))
							v62 = F_minmax_get_strategy_procinfo(m, v12, v14, v60, int32(5))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int64(0)
							} else {
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
								v65 = *(*int64)(unsafe.Add(mBase, uint32(v64)+8))
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
								v67 = *(*int64)(unsafe.Add(mBase, uint32(v66)+8))
								v68 = F_FunctionCall2Coll(m, v62, v9, v65, v67)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int64(0)
								} else {
									if v68 != int64(0) {
										v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+82)))
										if v73 == int32(0) {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
											v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
											F_pfree(m, v77)
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return int64(0)
											} else {
												v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+82)))
												v81 = v80
												v82 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
												v83 = *(*int64)(unsafe.Add(mBase, uint32(v82)+8))
												v86 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+72)))
												v87 = F_datumCopy(m, v83, v81&int32(1), v86)
												mBase = m.M
												v88 = m.ExcPending
												if v88 != 0 {
													return int64(0)
												} else {
													v89 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
													*(*int64)(unsafe.Add(mBase, uint32(v89)+8)) = v87
													return int64(0)
												}
											}
										} else {
											v81 = int32(1)
											v82 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
											v83 = *(*int64)(unsafe.Add(mBase, uint32(v82)+8))
											v86 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+72)))
											v87 = F_datumCopy(m, v83, v81&int32(1), v86)
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return int64(0)
											} else {
												v89 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
												*(*int64)(unsafe.Add(mBase, uint32(v89)+8)) = v87
												return int64(0)
											}
										}
									} else {
										return int64(0)
									}
								}
							}
						}
					}
				} else {
					v48 = v11
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
					v50 = *(*int64)(unsafe.Add(mBase, uint32(v49)))
					v53 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+72)))
					v54 = F_datumCopy(m, v50, v48&int32(1), v53)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int64(0)
					} else {
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
						*(*int64)(unsafe.Add(mBase, uint32(v56))) = v54
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v25)+68))
						v62 = F_minmax_get_strategy_procinfo(m, v12, v14, v60, int32(5))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int64(0)
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
							v65 = *(*int64)(unsafe.Add(mBase, uint32(v64)+8))
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
							v67 = *(*int64)(unsafe.Add(mBase, uint32(v66)+8))
							v68 = F_FunctionCall2Coll(m, v62, v9, v65, v67)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int64(0)
							} else {
								if v68 != int64(0) {
									v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+82)))
									if v73 == int32(0) {
										v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
										v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
										F_pfree(m, v77)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int64(0)
										} else {
											v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+82)))
											v81 = v80
											v82 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
											v83 = *(*int64)(unsafe.Add(mBase, uint32(v82)+8))
											v86 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+72)))
											v87 = F_datumCopy(m, v83, v81&int32(1), v86)
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return int64(0)
											} else {
												v89 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
												*(*int64)(unsafe.Add(mBase, uint32(v89)+8)) = v87
												return int64(0)
											}
										}
									} else {
										v81 = int32(1)
										v82 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
										v83 = *(*int64)(unsafe.Add(mBase, uint32(v82)+8))
										v86 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+72)))
										v87 = F_datumCopy(m, v83, v81&int32(1), v86)
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return int64(0)
										} else {
											v89 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
											*(*int64)(unsafe.Add(mBase, uint32(v89)+8)) = v87
											return int64(0)
										}
									}
								} else {
									return int64(0)
								}
							}
						}
					}
				}
			} else {
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v25)+68))
				v62 = F_minmax_get_strategy_procinfo(m, v12, v14, v60, int32(5))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int64(0)
				} else {
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
					v65 = *(*int64)(unsafe.Add(mBase, uint32(v64)+8))
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v67 = *(*int64)(unsafe.Add(mBase, uint32(v66)+8))
					v68 = F_FunctionCall2Coll(m, v62, v9, v65, v67)
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int64(0)
					} else {
						if v68 != int64(0) {
							v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+82)))
							if v73 == int32(0) {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
								F_pfree(m, v77)
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int64(0)
								} else {
									v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+82)))
									v81 = v80
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
									v83 = *(*int64)(unsafe.Add(mBase, uint32(v82)+8))
									v86 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+72)))
									v87 = F_datumCopy(m, v83, v81&int32(1), v86)
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return int64(0)
									} else {
										v89 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
										*(*int64)(unsafe.Add(mBase, uint32(v89)+8)) = v87
										return int64(0)
									}
								}
							} else {
								v81 = int32(1)
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
								v83 = *(*int64)(unsafe.Add(mBase, uint32(v82)+8))
								v86 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+72)))
								v87 = F_datumCopy(m, v83, v81&int32(1), v86)
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return int64(0)
								} else {
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
									*(*int64)(unsafe.Add(mBase, uint32(v89)+8)) = v87
									return int64(0)
								}
							}
						} else {
							return int64(0)
						}
					}
				}
			}
		}
	}
}
func F_brin_page_items(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v220 int32
	_ = v220
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v287 int64
	_ = v287
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int64
	_ = v300
	var v302 int64
	_ = v302
	var v304 int64
	_ = v304
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v327 int64
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
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
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v370 int64
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	v17 = m.G0
	v19 = v17 - int32(96)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v22 = F_pg_detoast_datum(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v28 = F_superuser(m)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v28 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L1
	} else {
		goto L97
	}
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if int32(7) < v34 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v38 = F_index_open(m, v27, int32(1))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L92
	}
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)+48))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+84))
	if v41 == int32(3580) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v44 = F_brin_build_desc(m, v38)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L1
	} else {
		goto L88
	}
L15:
	;
	v48 = F_verify_brin_page(m, v22, int32(_a_F_brin_page_items_0), int32(_a_F_brin_page_items_1))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	m.G0 = v19 + int32(96)
	return int64(0)
L17:
	;
	v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+14)))
	if v50 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_brin_free_desc(m, v44)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v38)+52))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v63 = F_palloc_mul(m, int32(4), v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	F_relation_close(m, v38, int32(1))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v58 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v58)
	goto L16
L23:
	;
	v65 = int32(1)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	if int32(0) < v67 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v73 = int32(1)
	v79 = v65
	goto L27
L25:
	;
	v176 = v65
	goto L26
L26:
	;
	v190 = int32(1)
	v191 = v19 + int32(24) | v190
	v192 = int32(0)
	v195 = v192
	v197 = v190
	v198 = v192
	v201 = v176
	goto L39
L27:
	;
	v92 = (v73 - int32(1)) << (uint(int32(2)) % 32)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v44+int32(20)+v92)))
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94))))
	v100 = F_palloc(m, v95*int32(28)+int32(4))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	v176 = v165
	goto L26
L29:
	;
	v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94))))
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v102
	if v102 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v109 = int32(0)
	goto L33
L31:
	;
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v92+v63))) = v100
	v165 = v79 + int32(1)
	v166 = base.I32_extend16_s(v165)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	if v166 <= v168 {
		v73 = v166
		v79 = v165
		goto L27
	} else {
		goto L38
	}
L33:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v94+int32(8)+v109<<(uint(int32(2))%32))))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	F_getTypeOutputInfo(m, v129, v19+int32(32), v19+int32(8))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	goto L32
L35:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	F_fmgr_info(m, v136, v100+int32(4)+v109*int32(28))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v143 = v109 + int32(1)
	v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94))))
	if base.Ui32(v143) < base.Ui32(v144) {
		v109 = v143
		goto L33
	} else {
		goto L37
	}
L37:
	;
	goto L34
L38:
	;
	goto L28
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = int64(0)
	if v198 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L40:
	;
	F_brin_free_desc(m, v44)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L86
	}
L41:
	;
	v473 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v473) {
		goto L82
	} else {
		goto L83
	}
L42:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+32)) = base.I64_extend_i32_u(v197) & int64(65535)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v266+v267<<(uint(int32(3))%32))+196))
	switch v271 - int32(20) {
	case 0, 3:
		goto L51
	default:
		goto L52
	}
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+32)) = base.I64_extend_i32_u(v197) & int64(65535)
	v244 = int32(16843009)
	*(*int32)(unsafe.Add(mBase, uint32(v191)+3)) = v244
	*(*int32)(unsafe.Add(mBase, uint32(v191))) = v244
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	F_tuplestore_putvalues(m, v248, v249, v19+int32(32), v19+int32(24))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L50
	}
L44:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v48+int32(20)+v197&int32(_a_F_brin_page_items_2)<<(uint(int32(2))%32))))
	if v220&int32(_a_F_brin_page_items_3) == int32(0) {
		v239 = v201
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v232 = int32(1)
	v233 = v201 + v232
	if v195&v232 == int32(0) {
		v260 = v198
		v261 = v233
		goto L42
	} else {
		goto L49
	}
L47:
	;
	v230 = F_brin_deform_tuple(m, v44, v48+v220&int32(_a_F_brin_page_items_4), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v260 = v230
	v261 = int32(1)
	goto L42
L49:
	;
	v239 = v233
	goto L43
L50:
	;
	v256 = int32(1)
	v456 = v197 + v256
	v457 = v198
	v460 = v239
	v470 = v256
	goto L41
L51:
	;
	v287 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v260)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+40)) = v287
	*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = base.I64_extend_i32_u(v261) & int64(65535)
	v293 = base.I32_extend16_s(v261)
	v296 = v260 + v293*int32(24)
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296)+3)))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+56)) = base.I64_extend_i32_u(v297)
	v300 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v296)+2)))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+64)) = v300
	v302 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v260))))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+72)) = v302
	v304 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v260)+1)))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+80)) = v304
	if v297 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L52:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errmsg_internal(m, int32(_a_F_brin_page_items_5), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_brin_page_items_6), int32(273), int32(_a_F_brin_page_items_7))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)))
	if v293 < v446 {
		v456 = v197
		v457 = v260
		v460 = v261
		v470 = int32(0)
		goto L41
	} else {
		goto L80
	}
L57:
	;
	v309 = v19 + int32(8)
	F_initStringInfo(m, v309)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v419 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+31)) = uint8(v419)
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	F_tuplestore_putvalues(m, v421, v266, v19+int32(32), v19+int32(24))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L79
	}
L60:
	;
	F_appendStringInfoChar(m, v309, int32(123))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v319 = v63 + v293<<(uint(int32(2))%32) - int32(4)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v319)))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v320)))
	if v321 <= int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	F_appendStringInfoChar(m, v19+int32(8), int32(125))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L75
	}
L63:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v296)+4))
	v327 = *(*int64)(unsafe.Add(mBase, uint32(v326)))
	v328 = F_OutputFunctionCall(m, v320+int32(4), v327)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_appendStringInfoString(m, v309, v328)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_pfree(m, v328)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v319)))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v335)))
	if v336 < int32(2) {
		goto L62
	} else {
		goto L67
	}
L67:
	;
	v339 = int32(1)
	goto L68
L68:
	;
	v356 = v19 + int32(8)
	F_appendStringInfoString(m, v356, int32(_a_F_brin_page_items_8))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L70
	}
L69:
	;
	goto L62
L70:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v319)))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v296)+4))
	v370 = *(*int64)(unsafe.Add(mBase, uint32(v366+v339<<(uint(int32(3))%32))))
	v371 = F_OutputFunctionCall(m, v360+v339*int32(28)+int32(4), v370)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_appendStringInfoString(m, v356, v371)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_pfree(m, v371)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v378 = v339 + int32(1)
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v319)))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v379)))
	if v378 < v380 {
		v339 = v378
		goto L68
	} else {
		goto L74
	}
L74:
	;
	goto L69
L75:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v404 = F_cstring_to_text(m, v403)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+88)) = base.I64_extend_i32_u(v404)
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	F_pfree(m, v408)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	F_tuplestore_putvalues(m, v411, v412, v19+int32(32), v19+int32(24))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	goto L56
L79:
	;
	goto L56
L80:
	;
	F_pfree(m, v260)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v452 = int32(0)
	v456 = v197 + int32(1)
	v457 = v452
	v460 = v261
	v470 = v452
	goto L41
L82:
	;
	v481 = int32(base.Ui32(v473+int32(_a_F_brin_page_items_9)) >> (uint(int32(2)) % 32))
	goto L84
L83:
	;
	v481 = int32(0)
	goto L84
L84:
	;
	if base.Ui32(v456&int32(_a_F_brin_page_items_2)) <= base.Ui32(v481&int32(_a_F_brin_page_items_2)) {
		v195 = v470
		v197 = v456
		v198 = v457
		v201 = v460
		goto L39
	} else {
		goto L85
	}
L85:
	;
	goto L40
L86:
	;
	F_relation_close(m, v38, int32(1))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	goto L16
L88:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v38)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = int32(_a_F_brin_page_items_10)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v518 + int32(4)
	F_errmsg(m, int32(_a_F_brin_page_items_11), v19)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_brin_page_items_6), int32(172), int32(_a_F_brin_page_items_7))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_errmsg(m, int32(_a_F_brin_page_items_12), int32(0))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_errhint(m, int32(_a_F_brin_page_items_13), int32(0))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_brin_page_items_6), int32(164), int32(_a_F_brin_page_items_7))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errmsg(m, int32(_a_F_brin_page_items_14), int32(0))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_brin_page_items_6), int32(147), int32(_a_F_brin_page_items_7))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_brin_redo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int64
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int64
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int64
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+48)))
	v14 = v12 & int32(240)
	switch int32(base.Ui32(v14)>>(uint(int32(4))%32)) & int32(7) {
	case 0:
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
		v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v22 = F_XLogInitBufferForRedo(m, l0, int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			if v22 < int32(0) {
				v27 = *(*int32)(unsafe.Add(mBase, _c_F_brin_redo[0]))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v27+(v22^int32(-1))<<(uint(int32(2))%32))))
				v41 = v33
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, _c_F_brin_redo[1]))
				v41 = v35 + v22<<(uint(int32(13))%32) + int32(-8192)
			}
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
			F_PageInit(m, v41, int32(_a_F_brin_redo_0), int32(8))
			mBase = m.M
			v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+16)))
			v49 = int32(_a_F_brin_redo_1)
			*(*uint16)(unsafe.Add(mBase, uint32(v41+v47)+6)) = uint16(v49)
			*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v42
			*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v43
			*(*int32)(unsafe.Add(mBase, uint32(v41)+24)) = int32(-1475306246)
			v57 = int32(40)
			*(*uint16)(unsafe.Add(mBase, uint32(v41)+12)) = uint16(v57)
			*(*int64)(unsafe.Add(mBase, uint32(v41))) = base.I64_rotl(v20, int64(32))
			F_MarkBufferDirty(m, v22)
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return
			} else {
				F_UnlockReleaseBuffer(m, v22)
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					m.G0 = v9 + int32(32)
					return
				}
			}
		}
	case 1:
		v419 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
		F_brin_xlog_insert_update(m, l0, v419)
		mBase = m.M
		v421 = m.ExcPending
		if v421 != 0 {
			return
		} else {
			m.G0 = v9 + int32(32)
			return
		}
	case 2:
		v66 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
		v67 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v71 = F_XLogReadBufferForRedo(m, l0, int32(2), v9+int32(20))
		mBase = m.M
		v72 = m.ExcPending
		if v72 != 0 {
			return
		} else {
			if v71 == int32(0) {
				v75 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
				if v75 < int32(0) {
					v79 = *(*int32)(unsafe.Add(mBase, _c_F_brin_redo[0]))
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v79+(v75^int32(-1))<<(uint(int32(2))%32))))
					v93 = v85
				} else {
					v87 = *(*int32)(unsafe.Add(mBase, _c_F_brin_redo[1]))
					v93 = v87 + v75<<(uint(int32(13))%32) + int32(-8192)
				}
				v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66))))
				F_PageIndexTupleDeleteNoCompact(m, v93, v94)
				mBase = m.M
				v96 = m.ExcPending
				if v96 != 0 {
					return
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v93))) = base.I64_rotl(v67, int64(32))
					v100 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
					F_MarkBufferDirty(m, v100)
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return
					} else {
						F_brin_xlog_insert_update(m, l0, v66+int32(4))
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return
						} else {
							v108 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
							if v108 == int32(0) {
								m.G0 = v9 + int32(32)
								return
							} else {
								F_UnlockReleaseBuffer(m, v108)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return
								} else {
									m.G0 = v9 + int32(32)
									return
								}
							}
						}
					}
				}
			} else {
				F_brin_xlog_insert_update(m, l0, v66+int32(4))
				mBase = m.M
				v107 = m.ExcPending
				if v107 != 0 {
					return
				} else {
					v108 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
					if v108 == int32(0) {
						m.G0 = v9 + int32(32)
						return
					} else {
						F_UnlockReleaseBuffer(m, v108)
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return
						} else {
							m.G0 = v9 + int32(32)
							return
						}
					}
				}
			}
		}
	case 3:
		v113 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v114 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
		v118 = F_XLogReadBufferForRedo(m, l0, int32(0), v9+int32(20))
		mBase = m.M
		v119 = m.ExcPending
		if v119 != 0 {
			return
		} else {
			if v118 == int32(0) {
				v122 = int32(0)
				v124 = v9 + int32(28)
				v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
				v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+72))
				if v127 < v122 {
					v149 = v122
					v152 = v149
				} else {
					v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126+int32(0))+76)))
					if v132 != int32(1) {
						v149 = v122
						v152 = v149
					} else {
						v136 = v126 + int32(76)
						v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+43)))
						if v137 == int32(0) {
							if v124 == int32(0) {
								v149 = v122
								v152 = v149
							} else {
								v142 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v124))) = v142
								v152 = v142
							}
						} else {
							if v124 != 0 {
								v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136)+48)))
								*(*int32)(unsafe.Add(mBase, uint32(v124))) = v145
							} else {
							}
							v147 = *(*int32)(unsafe.Add(mBase, uint32(v136)+44))
							v149 = v147
							v152 = v149
						}
					}
				}
				v153 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
				if v153 < int32(0) {
					v157 = *(*int32)(unsafe.Add(mBase, _c_F_brin_redo[0]))
					v163 = *(*int32)(unsafe.Add(mBase, uint32(v157+(v153^int32(-1))<<(uint(int32(2))%32))))
					v171 = v163
				} else {
					v165 = *(*int32)(unsafe.Add(mBase, _c_F_brin_redo[1]))
					v171 = v165 + v153<<(uint(int32(13))%32) + int32(-8192)
				}
				v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114))))
				v173 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
				v174 = F_PageIndexTupleOverwrite(m, v171, v172, v152, v173)
				mBase = m.M
				v175 = m.ExcPending
				if v175 != 0 {
					return
				} else {
					if v174 == int32(0) {
						F_errstart_cold(m, int32(24), int32(0))
						mBase = m.M
						v433 = m.ExcPending
						if v433 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(_a_F_brin_redo_2), int32(0))
							mBase = m.M
							v437 = m.ExcPending
							if v437 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_brin_redo_3), int32(193), int32(_a_F_brin_redo_4))
								mBase = m.M
								v442 = m.ExcPending
								if v442 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v171))) = base.I64_rotl(v113, int64(32))
						v181 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
						F_MarkBufferDirty(m, v181)
						mBase = m.M
						v183 = m.ExcPending
						if v183 != 0 {
							return
						} else {
							v186 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
							if v186 == int32(0) {
								m.G0 = v9 + int32(32)
								return
							} else {
								F_UnlockReleaseBuffer(m, v186)
								mBase = m.M
								v190 = m.ExcPending
								if v190 != 0 {
									return
								} else {
									m.G0 = v9 + int32(32)
									return
								}
							}
						}
					}
				}
			} else {
				v186 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
				if v186 == int32(0) {
					m.G0 = v9 + int32(32)
					return
				} else {
					F_UnlockReleaseBuffer(m, v186)
					mBase = m.M
					v190 = m.ExcPending
					if v190 != 0 {
						return
					} else {
						m.G0 = v9 + int32(32)
						return
					}
				}
			}
		}
	case 4:
		v191 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v192 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
		v194 = int32(0)
		F_XLogRecGetBlockTag(m, l0, int32(1), v194, v194, v9+int32(28))
		mBase = m.M
		v199 = m.ExcPending
		if v199 != 0 {
			return
		} else {
			v203 = F_XLogReadBufferForRedo(m, l0, int32(0), v9+int32(20))
			mBase = m.M
			v204 = m.ExcPending
			if v204 != 0 {
				return
			} else {
				if v203 == int32(0) {
					v207 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
					v208 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
					if v208 < int32(0) {
						v212 = *(*int32)(unsafe.Add(mBase, _c_F_brin_redo[0]))
						v218 = *(*int32)(unsafe.Add(mBase, uint32(v212+(v208^int32(-1))<<(uint(int32(2))%32))))
						v226 = v218
					} else {
						v220 = *(*int32)(unsafe.Add(mBase, _c_F_brin_redo[1]))
						v226 = v220 + v208<<(uint(int32(13))%32) + int32(-8192)
					}
					*(*int64)(unsafe.Add(mBase, uint32(v226))) = base.I64_rotl(v191, int64(32))
					*(*int32)(unsafe.Add(mBase, uint32(v226)+36)) = v207
					v231 = int32(40)
					*(*uint16)(unsafe.Add(mBase, uint32(v226)+12)) = uint16(v231)
					F_MarkBufferDirty(m, v208)
					mBase = m.M
					v234 = m.ExcPending
					if v234 != 0 {
						return
					} else {
						v239 = F_XLogInitBufferForRedo(m, l0, int32(1))
						mBase = m.M
						v240 = m.ExcPending
						if v240 != 0 {
							return
						} else {
							if v239 < int32(0) {
								v244 = *(*int32)(unsafe.Add(mBase, _c_F_brin_redo[0]))
								v250 = *(*int32)(unsafe.Add(mBase, uint32(v244+(v239^int32(-1))<<(uint(int32(2))%32))))
								v258 = v250
							} else {
								v252 = *(*int32)(unsafe.Add(mBase, _c_F_brin_redo[1]))
								v258 = v252 + v239<<(uint(int32(13))%32) + int32(-8192)
							}
							v259 = int32(_a_F_brin_redo_5)
							F_PageInit(m, v258, int32(_a_F_brin_redo_0), int32(8))
							mBase = m.M
							v263 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v258)+16)))
							*(*uint16)(unsafe.Add(mBase, uint32(v258+v263)+6)) = uint16(v259)
							*(*int64)(unsafe.Add(mBase, uint32(v258))) = base.I64_rotl(v191, int64(32))
							F_MarkBufferDirty(m, v239)
							mBase = m.M
							v270 = m.ExcPending
							if v270 != 0 {
								return
							} else {
								F_UnlockReleaseBuffer(m, v239)
								mBase = m.M
								v272 = m.ExcPending
								if v272 != 0 {
									return
								} else {
									v273 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
									if v273 == int32(0) {
										m.G0 = v9 + int32(32)
										return
									} else {
										F_UnlockReleaseBuffer(m, v273)
										mBase = m.M
										v277 = m.ExcPending
										if v277 != 0 {
											return
										} else {
											m.G0 = v9 + int32(32)
											return
										}
									}
								}
							}
						}
					}
				} else {
					v239 = F_XLogInitBufferForRedo(m, l0, int32(1))
					mBase = m.M
					v240 = m.ExcPending
					if v240 != 0 {
						return
					} else {
						if v239 < int32(0) {
							v244 = *(*int32)(unsafe.Add(mBase, _c_F_brin_redo[0]))
							v250 = *(*int32)(unsafe.Add(mBase, uint32(v244+(v239^int32(-1))<<(uint(int32(2))%32))))
							v258 = v250
						} else {
							v252 = *(*int32)(unsafe.Add(mBase, _c_F_brin_redo[1]))
							v258 = v252 + v239<<(uint(int32(13))%32) + int32(-8192)
						}
						v259 = int32(_a_F_brin_redo_5)
						F_PageInit(m, v258, int32(_a_F_brin_redo_0), int32(8))
						mBase = m.M
						v263 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v258)+16)))
						*(*uint16)(unsafe.Add(mBase, uint32(v258+v263)+6)) = uint16(v259)
						*(*int64)(unsafe.Add(mBase, uint32(v258))) = base.I64_rotl(v191, int64(32))
						F_MarkBufferDirty(m, v239)
						mBase = m.M
						v270 = m.ExcPending
						if v270 != 0 {
							return
						} else {
							F_UnlockReleaseBuffer(m, v239)
							mBase = m.M
							v272 = m.ExcPending
							if v272 != 0 {
								return
							} else {
								v273 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
								if v273 == int32(0) {
									m.G0 = v9 + int32(32)
									return
								} else {
									F_UnlockReleaseBuffer(m, v273)
									mBase = m.M
									v277 = m.ExcPending
									if v277 != 0 {
										return
									} else {
										m.G0 = v9 + int32(32)
										return
									}
								}
							}
						}
					}
				}
			}
		}
	case 5:
		v278 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v279 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
		v283 = F_XLogReadBufferForRedo(m, l0, int32(0), v9+int32(28))
		mBase = m.M
		v284 = m.ExcPending
		if v284 != 0 {
			return
		} else {
			if v283 == int32(0) {
				v287 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v9)+24)) = uint16(v287)
				v289 = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v289
				v291 = *(*int32)(unsafe.Add(mBase, uint32(v279)+4))
				v292 = *(*int32)(unsafe.Add(mBase, uint32(v279)))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v289
				*(*uint16)(unsafe.Add(mBase, uint32(v9)+16)) = uint16(v287)
				v297 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
				v299 = v9 + int32(12)
				v302 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v299))))
				v303 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v299)+2)))
				if v297 < v287 {
					v307 = *(*int32)(unsafe.Add(mBase, _c_F_brin_redo[0]))
					v313 = *(*int32)(unsafe.Add(mBase, uint32(v307+(v297^int32(-1))<<(uint(int32(2))%32))))
					v321 = v313
				} else {
					v315 = *(*int32)(unsafe.Add(mBase, _c_F_brin_redo[1]))
					v321 = v315 + v297<<(uint(int32(13))%32) + int32(-8192)
				}
				v322 = base.I32_div_u_s(v291, v292)
				v324 = base.I32_rem_u_s(v322, int32(1360))
				v327 = v321 + v324*int32(6)
				v328 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v299)+4)))
				*(*uint16)(unsafe.Add(mBase, uint32(v327)+28)) = uint16(v328)
				if v328 != 0 {
					v331 = v303
				} else {
					v331 = int32(-1)
				}
				*(*uint16)(unsafe.Add(mBase, uint32(v327)+26)) = uint16(v331)
				if v328 != 0 {
					v334 = v302
				} else {
					v334 = int32(-1)
				}
				*(*uint16)(unsafe.Add(mBase, uint32(v327)+24)) = uint16(v334)
				v336 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
				if v336 < int32(0) {
					v340 = *(*int32)(unsafe.Add(mBase, _c_F_brin_redo[0]))
					v346 = *(*int32)(unsafe.Add(mBase, uint32(v340+(v336^int32(-1))<<(uint(int32(2))%32))))
					v354 = v346
				} else {
					v348 = *(*int32)(unsafe.Add(mBase, _c_F_brin_redo[1]))
					v354 = v348 + v336<<(uint(int32(13))%32) + int32(-8192)
				}
				*(*int64)(unsafe.Add(mBase, uint32(v354))) = base.I64_rotl(v278, int64(32))
				F_MarkBufferDirty(m, v336)
				mBase = m.M
				v359 = m.ExcPending
				if v359 != 0 {
					return
				} else {
					v362 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
					if v362 != 0 {
						F_UnlockReleaseBuffer(m, v362)
						mBase = m.M
						v364 = m.ExcPending
						if v364 != 0 {
							return
						} else {
							v368 = F_XLogReadBufferForRedo(m, l0, int32(1), v9+int32(28))
							mBase = m.M
							v369 = m.ExcPending
							if v369 != 0 {
								return
							} else {
								if v368 == int32(0) {
									v372 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
									if v372 < int32(0) {
										v376 = *(*int32)(unsafe.Add(mBase, _c_F_brin_redo[0]))
										v382 = *(*int32)(unsafe.Add(mBase, uint32(v376+(v372^int32(-1))<<(uint(int32(2))%32))))
										v390 = v382
									} else {
										v384 = *(*int32)(unsafe.Add(mBase, _c_F_brin_redo[1]))
										v390 = v384 + v372<<(uint(int32(13))%32) + int32(-8192)
									}
									v391 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v279)+8)))
									F_PageIndexTupleDeleteNoCompact(m, v390, v391)
									mBase = m.M
									v393 = m.ExcPending
									if v393 != 0 {
										return
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v390))) = base.I64_rotl(v278, int64(32))
										v397 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
										F_MarkBufferDirty(m, v397)
										mBase = m.M
										v399 = m.ExcPending
										if v399 != 0 {
											return
										} else {
											v401 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
											if v401 == int32(0) {
												m.G0 = v9 + int32(32)
												return
											} else {
												F_UnlockReleaseBuffer(m, v401)
												mBase = m.M
												v405 = m.ExcPending
												if v405 != 0 {
													return
												} else {
													m.G0 = v9 + int32(32)
													return
												}
											}
										}
									}
								} else {
									v401 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
									if v401 == int32(0) {
										m.G0 = v9 + int32(32)
										return
									} else {
										F_UnlockReleaseBuffer(m, v401)
										mBase = m.M
										v405 = m.ExcPending
										if v405 != 0 {
											return
										} else {
											m.G0 = v9 + int32(32)
											return
										}
									}
								}
							}
						}
					} else {
						v368 = F_XLogReadBufferForRedo(m, l0, int32(1), v9+int32(28))
						mBase = m.M
						v369 = m.ExcPending
						if v369 != 0 {
							return
						} else {
							if v368 == int32(0) {
								v372 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
								if v372 < int32(0) {
									v376 = *(*int32)(unsafe.Add(mBase, _c_F_brin_redo[0]))
									v382 = *(*int32)(unsafe.Add(mBase, uint32(v376+(v372^int32(-1))<<(uint(int32(2))%32))))
									v390 = v382
								} else {
									v384 = *(*int32)(unsafe.Add(mBase, _c_F_brin_redo[1]))
									v390 = v384 + v372<<(uint(int32(13))%32) + int32(-8192)
								}
								v391 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v279)+8)))
								F_PageIndexTupleDeleteNoCompact(m, v390, v391)
								mBase = m.M
								v393 = m.ExcPending
								if v393 != 0 {
									return
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v390))) = base.I64_rotl(v278, int64(32))
									v397 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
									F_MarkBufferDirty(m, v397)
									mBase = m.M
									v399 = m.ExcPending
									if v399 != 0 {
										return
									} else {
										v401 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
										if v401 == int32(0) {
											m.G0 = v9 + int32(32)
											return
										} else {
											F_UnlockReleaseBuffer(m, v401)
											mBase = m.M
											v405 = m.ExcPending
											if v405 != 0 {
												return
											} else {
												m.G0 = v9 + int32(32)
												return
											}
										}
									}
								}
							} else {
								v401 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
								if v401 == int32(0) {
									m.G0 = v9 + int32(32)
									return
								} else {
									F_UnlockReleaseBuffer(m, v401)
									mBase = m.M
									v405 = m.ExcPending
									if v405 != 0 {
										return
									} else {
										m.G0 = v9 + int32(32)
										return
									}
								}
							}
						}
					}
				}
			} else {
				v362 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
				if v362 != 0 {
					F_UnlockReleaseBuffer(m, v362)
					mBase = m.M
					v364 = m.ExcPending
					if v364 != 0 {
						return
					} else {
						v368 = F_XLogReadBufferForRedo(m, l0, int32(1), v9+int32(28))
						mBase = m.M
						v369 = m.ExcPending
						if v369 != 0 {
							return
						} else {
							if v368 == int32(0) {
								v372 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
								if v372 < int32(0) {
									v376 = *(*int32)(unsafe.Add(mBase, _c_F_brin_redo[0]))
									v382 = *(*int32)(unsafe.Add(mBase, uint32(v376+(v372^int32(-1))<<(uint(int32(2))%32))))
									v390 = v382
								} else {
									v384 = *(*int32)(unsafe.Add(mBase, _c_F_brin_redo[1]))
									v390 = v384 + v372<<(uint(int32(13))%32) + int32(-8192)
								}
								v391 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v279)+8)))
								F_PageIndexTupleDeleteNoCompact(m, v390, v391)
								mBase = m.M
								v393 = m.ExcPending
								if v393 != 0 {
									return
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v390))) = base.I64_rotl(v278, int64(32))
									v397 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
									F_MarkBufferDirty(m, v397)
									mBase = m.M
									v399 = m.ExcPending
									if v399 != 0 {
										return
									} else {
										v401 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
										if v401 == int32(0) {
											m.G0 = v9 + int32(32)
											return
										} else {
											F_UnlockReleaseBuffer(m, v401)
											mBase = m.M
											v405 = m.ExcPending
											if v405 != 0 {
												return
											} else {
												m.G0 = v9 + int32(32)
												return
											}
										}
									}
								}
							} else {
								v401 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
								if v401 == int32(0) {
									m.G0 = v9 + int32(32)
									return
								} else {
									F_UnlockReleaseBuffer(m, v401)
									mBase = m.M
									v405 = m.ExcPending
									if v405 != 0 {
										return
									} else {
										m.G0 = v9 + int32(32)
										return
									}
								}
							}
						}
					}
				} else {
					v368 = F_XLogReadBufferForRedo(m, l0, int32(1), v9+int32(28))
					mBase = m.M
					v369 = m.ExcPending
					if v369 != 0 {
						return
					} else {
						if v368 == int32(0) {
							v372 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
							if v372 < int32(0) {
								v376 = *(*int32)(unsafe.Add(mBase, _c_F_brin_redo[0]))
								v382 = *(*int32)(unsafe.Add(mBase, uint32(v376+(v372^int32(-1))<<(uint(int32(2))%32))))
								v390 = v382
							} else {
								v384 = *(*int32)(unsafe.Add(mBase, _c_F_brin_redo[1]))
								v390 = v384 + v372<<(uint(int32(13))%32) + int32(-8192)
							}
							v391 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v279)+8)))
							F_PageIndexTupleDeleteNoCompact(m, v390, v391)
							mBase = m.M
							v393 = m.ExcPending
							if v393 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v390))) = base.I64_rotl(v278, int64(32))
								v397 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
								F_MarkBufferDirty(m, v397)
								mBase = m.M
								v399 = m.ExcPending
								if v399 != 0 {
									return
								} else {
									v401 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
									if v401 == int32(0) {
										m.G0 = v9 + int32(32)
										return
									} else {
										F_UnlockReleaseBuffer(m, v401)
										mBase = m.M
										v405 = m.ExcPending
										if v405 != 0 {
											return
										} else {
											m.G0 = v9 + int32(32)
											return
										}
									}
								}
							}
						} else {
							v401 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
							if v401 == int32(0) {
								m.G0 = v9 + int32(32)
								return
							} else {
								F_UnlockReleaseBuffer(m, v401)
								mBase = m.M
								v405 = m.ExcPending
								if v405 != 0 {
									return
								} else {
									m.G0 = v9 + int32(32)
									return
								}
							}
						}
					}
				}
			}
		}
	default:
		F_errstart_cold(m, int32(24), int32(0))
		mBase = m.M
		v409 = m.ExcPending
		if v409 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v14
			F_errmsg_internal(m, int32(_a_F_brin_redo_6), v9)
			mBase = m.M
			v413 = m.ExcPending
			if v413 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_brin_redo_3), int32(334), int32(_a_F_brin_redo_7))
				mBase = m.M
				v418 = m.ExcPending
				if v418 != 0 {
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
func F_brin_revmap_data(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int64
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	v6 = F_superuser(m)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		if v6 != 0 {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			if v11 == int32(0) {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v15 = F_pg_detoast_datum(m, v14)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int64(0)
				} else {
					v17 = F_init_MultiFuncCall(m, l0)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int64(0)
					} else {
						v19 = int32(_a_F_brin_revmap_data_0)
						v20 = *(*int32)(unsafe.Add(mBase, _c_F_brin_revmap_data[0]))
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
						*(*int32)(unsafe.Add(mBase, _c_F_brin_revmap_data[0])) = v22
						v26 = F_verify_brin_page(m, v15, int32(_a_F_brin_revmap_data_1), int32(_a_F_brin_revmap_data_2))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int64(0)
						} else {
							v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+14)))
							if v28 == int32(0) {
								*(*int32)(unsafe.Add(mBase, _c_F_brin_revmap_data[0])) = v20
								v33 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v33)
								return int64(0)
							} else {
								v38 = F_palloc(m, int32(8))
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v38))) = v26 + int32(24)
									*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v38
									*(*int32)(unsafe.Add(mBase, _c_F_brin_revmap_data[0])) = v20
									v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
									if base.Ui32(v55) <= base.Ui32(int32(1359)) {
										v58 = *(*int64)(unsafe.Add(mBase, uint32(v53)))
										*(*int64)(unsafe.Add(mBase, uint32(v53))) = v58 + int64(1)
										v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v63 = int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v62)+20)) = v63
										*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v55 + v63
										v68 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
										return base.I64_extend_i32_u(v68 + v55*int32(6))
									} else {
										F_end_MultiFuncCall(m, l0)
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return int64(0)
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v76)+20)) = int32(2)
											v79 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v79)
											return int64(0)
										}
									}
								}
							}
						}
					}
				}
			} else {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
				if base.Ui32(v55) <= base.Ui32(int32(1359)) {
					v58 = *(*int64)(unsafe.Add(mBase, uint32(v53)))
					*(*int64)(unsafe.Add(mBase, uint32(v53))) = v58 + int64(1)
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v63 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v62)+20)) = v63
					*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v55 + v63
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
					return base.I64_extend_i32_u(v68 + v55*int32(6))
				} else {
					F_end_MultiFuncCall(m, l0)
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int64(0)
					} else {
						v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v76)+20)) = int32(2)
						v79 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v79)
						return int64(0)
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v86 = m.ExcPending
			if v86 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_brin_revmap_data_3), int32(0))
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_brin_revmap_data_4), int32(397), int32(_a_F_brin_revmap_data_5))
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return int64(0)
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
