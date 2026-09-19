package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_BackendPidGetProc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	v2 = int32(0)
	if l0 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_BackendPidGetProc[0]))
	v16 = F_LWLockAcquire(m, v12+int32(512), int32(1))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_BackendPidGetProc[1]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v22 <= int32(0) {
		v51 = v2
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_BackendPidGetProc[0]))
	F_LWLockRelease(m, v56+int32(512))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L12
	}
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_BackendPidGetProc[2]))
	v31 = int32(0)
	goto L8
L8:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v21+int32(36)+v31<<(uint(int32(2))%32))))
	v42 = v29 + v39*int32(640)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+44))
	if v43 == l0 {
		v51 = v42
		goto L6
	} else {
		goto L10
	}
L9:
	;
	v51 = int32(0)
	goto L6
L10:
	;
	v46 = v31 + int32(1)
	if v46 != v22 {
		v31 = v46
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	return v51
}
func F_BuildCallback_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
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
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 float64
	_ = v115
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
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
	var v142 int32
	_ = v142
	var v143 float64
	_ = v143
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
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
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
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v263 int32
	_ = v263
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v371 int32
	_ = v371
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v433 int32
	_ = v433
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 float32
	_ = v493
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v615 float64
	_ = v615
	var v617 float64
	_ = v617
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v646 int32
	_ = v646
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	v25 = m.G0
	v27 = v25 - int32(16)
	m.G0 = v27
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v29 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l5)+160))
	v33 = int32(_a_F_BuildCallback_1_0)
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_BuildCallback_1[0]))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l5)+184))
	*(*int32)(unsafe.Add(mBase, _c_F_BuildCallback_1[0])) = v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l5)+204))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	v43 = l5 + int32(48)
	v44 = F_HnswFormIndexValue(m, v27+int32(12), l2, l3, v41, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v27 + int32(16)
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BuildCallback_1[0])) = v34
	v685 = *(*int32)(unsafe.Add(mBase, uint32(l5)+184))
	F_MemoryContextReset(m, v685)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L5
	} else {
		goto L169
	}
L5:
	;
	return
L6:
	;
	if v44 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v49 = v32 + int32(76)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v51 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v77 = F_LWLockAcquire(m, v49, int32(1))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L5
	} else {
		goto L17
	}
L9:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	if base.Ui32((v55-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v75 = int32(6)
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v67 = int32(1)
	if v51&v67 != 0 {
		v75 = int32(base.Ui32(v51) >> (uint(v67) % 32))
		goto L8
	} else {
		goto L16
	}
L12:
	;
	v62 = int32(18)
	if v55 == v62 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v66 = v62
	goto L15
L14:
	;
	v66 = int32(2)
	goto L15
L15:
	;
	v75 = v66
	goto L8
L16:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v75 = int32(base.Ui32(v71) >> (uint(int32(2)) % 32))
	goto L8
L17:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+92)))
	if v79 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v609 = base.AtomicRmwXchg32(m, v32, int32(0), int32(1))
	if v609 != 0 {
		goto L161
	} else {
		goto L162
	}
L19:
	;
	F_LWLockRelease(m, v49)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L5
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v89 = v32 + int32(52)
	v91 = F_LWLockAcquire(m, v89, int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L25
	}
L22:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v86 = F_HnswInsertTupleOnDisk(m, l0, v43, v84, l1, int32(1))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	if v86 != 0 {
		goto L18
	} else {
		goto L24
	}
L24:
	;
	goto L4
L25:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v32)+68))
	if v38 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v96 = int32(_a_F_BuildCallback_1_1)
	goto L28
L27:
	;
	v96 = int32(0)
	goto L28
L28:
	;
	v97 = F_add_size(m, v93, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v32)+72))
	if base.Ui32(v99) <= base.Ui32(v97) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	F_LWLockRelease(m, v89)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L5
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	v143 = *(*float64)(unsafe.Add(mBase, uint32(l5)+168))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l5)+176))
	v146 = l5 + int32(188)
	v147 = F_HnswInitElement(m, v38, l1, v142, v143, v144, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L5
	} else {
		goto L51
	}
L33:
	;
	F_LWLockRelease(m, v49)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v106 = F_LWLockAcquire(m, v49, int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+92)))
	if v108 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v113 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L5
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	F_LWLockRelease(m, v49)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L5
	} else {
		goto L48
	}
L39:
	;
	if v113 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v115 = *(*float64)(unsafe.Add(mBase, uint32(v32)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v27))) = base.I64_trunc_sat_f64_s(v115)
	F_errmsg(m, int32(_a_F_BuildCallback_1_2), v27)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L5
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	F_FlushPages(m, l5)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L5
	} else {
		goto L47
	}
L43:
	;
	F_errdetail(m, int32(_a_F_BuildCallback_1_3), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L5
	} else {
		goto L44
	}
L44:
	;
	F_errhint(m, int32(_a_F_BuildCallback_1_4), int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L5
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_BuildCallback_1_5), int32(542), int32(_a_F_BuildCallback_1_6))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	goto L42
L47:
	;
	goto L38
L48:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v140 = F_HnswInsertTupleOnDisk(m, l0, v43, v138, l1, int32(1))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L5
	} else {
		goto L49
	}
L49:
	;
	if v140 != 0 {
		goto L18
	} else {
		goto L50
	}
L50:
	;
	goto L4
L51:
	;
	v149 = F_HnswAlloc(m, v146, v75)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L5
	} else {
		goto L52
	}
L52:
	;
	F_LWLockRelease(m, v89)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	if v75 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	base.MemoryCopy(m, v149, v153, v75)
	goto L56
L55:
	;
	goto L56
L56:
	;
	if v149 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v159 = v149 - v38 + int32(1)
	goto L59
L58:
	;
	v159 = int32(0)
	goto L59
L59:
	;
	if v38 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v160 = v159
	goto L62
L61:
	;
	v160 = v149
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+88)) = v160
	v163 = v147 + int32(92)
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_BuildCallback_1[1]))
	*(*uint16)(unsafe.Add(mBase, uint32(v163))) = uint16(v165)
	*(*int32)(unsafe.Add(mBase, uint32(v163)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v163)+8)) = int64(-1)
	goto L63
L63:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l5)+204))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l5)+160))
	v176 = v174 + int32(32)
	v178 = F_LWLockAcquire(m, v176, int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L5
	} else {
		goto L64
	}
L64:
	;
	F_LWLockRelease(m, v176)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	v183 = v174 + int32(16)
	v185 = F_LWLockAcquire(m, v183, int32(1))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L5
	} else {
		goto L66
	}
L66:
	;
	if v171 != 0 {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	v226 = int32(0)
	F_HnswFindElementNeighbors(m, v171, v147, v225, v226, v43, v172, v173, v226)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L5
	} else {
		goto L84
	}
L68:
	;
	F_LWLockRelease(m, v183)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L5
	} else {
		goto L76
	}
L69:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+65)))
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+65)))
	if base.Ui32(v198) <= base.Ui32(v199) {
		v225 = v197
		goto L67
	} else {
		goto L75
	}
L70:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v174)+48))
	if v187 == int32(0) {
		goto L68
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v174)+48))
	if v193 == int32(0) {
		goto L68
	} else {
		goto L74
	}
L73:
	;
	v197 = v171 + v187 - int32(1)
	goto L69
L74:
	;
	v197 = v193
	goto L69
L75:
	;
	goto L68
L76:
	;
	v205 = int32(0)
	v207 = F_LWLockAcquire(m, v176, v205)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L5
	} else {
		goto L77
	}
L77:
	;
	v210 = F_LWLockAcquire(m, v183, int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L5
	} else {
		goto L78
	}
L78:
	;
	F_LWLockRelease(m, v176)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L5
	} else {
		goto L79
	}
L79:
	;
	if v171 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v174)+48))
	v225 = v216
	goto L67
L81:
	;
	goto L82
L82:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v174)+48))
	if v217 == int32(0) {
		v225 = v205
		goto L67
	} else {
		goto L83
	}
L83:
	;
	v225 = v171 + v217 - int32(1)
	goto L67
L84:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l5)+204))
	if v231 != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l5)+160))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	if v255 <= int32(0) {
		goto L95
	} else {
		goto L96
	}
L86:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v147)+72))
	v234 = int32(1)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v231+v232-v234)))
	if v236 != 0 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v147)+88))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v147)+72))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	v252 = v248
	v253 = v250
	goto L85
L89:
	;
	v241 = v231 + v236 - v234
	goto L91
L90:
	;
	v241 = int32(0)
	goto L91
L91:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v147)+88))
	if v242 == int32(0) {
		v252 = v226
		v253 = v241
		goto L85
	} else {
		goto L92
	}
L92:
	;
	v252 = v242 + v231 - int32(1)
	v253 = v241
	goto L85
L93:
	;
	F_LWLockRelease(m, v183)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L5
	} else {
		goto L159
	}
L94:
	;
	v540 = v147 + int32(4)
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300)+64)))
	v544 = v542 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v300)+64)) = uint8(v544)
	v548 = v300 + v542*int32(6)
	v549 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v540)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v548)+8)) = uint16(v549)
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v540)))
	*(*int32)(unsafe.Add(mBase, uint32(v548)+4)) = v551
	goto L157
L95:
	;
	v348 = base.AtomicRmwXchg32(m, v254, int32(0), int32(1))
	if v348 != 0 {
		goto L111
	} else {
		goto L112
	}
L96:
	;
	v263 = int32(0)
	goto L97
L97:
	;
	v287 = v253 + int32(8) + v263*int32(12)
	if v231 != 0 {
		goto L101
	} else {
		goto L102
	}
L98:
	;
	goto L95
L99:
	;
	v304 = F_datumIsEqual(m, v252, v301, int32(0), int32(-1))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L5
	} else {
		goto L105
	}
L100:
	;
	v300 = v291
	v301 = v292 + v231 - int32(1)
	goto L99
L101:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)))
	v289 = v231 + v288
	v291 = v289 - int32(1)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v289)+87))
	if v292 != 0 {
		goto L100
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v287)))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v294)+88))
	v300 = v294
	v301 = v295
	goto L99
L104:
	;
	v300 = v291
	v301 = int32(0)
	goto L99
L105:
	;
	if v304 == int32(0) {
		goto L95
	} else {
		goto L106
	}
L106:
	;
	v309 = v300 + int32(92)
	v311 = F_LWLockAcquire(m, v309, int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300)+64)))
	if v313 != int32(10) {
		goto L94
	} else {
		goto L108
	}
L108:
	;
	F_LWLockRelease(m, v309)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L5
	} else {
		goto L109
	}
L109:
	;
	v319 = v263 + int32(1)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	if v319 < v320 {
		v263 = v319
		goto L97
	} else {
		goto L110
	}
L110:
	;
	goto L98
L111:
	;
	F_s_lock(m, v254, int32(_a_F_BuildCallback_1_5), int32(372), int32(_a_F_BuildCallback_1_7))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L5
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v254)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = v354
	v358 = v147 - v231 + int32(1)
	if v231 != 0 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	goto L113
L115:
	;
	v359 = v358
	goto L117
L116:
	;
	v359 = v147
	goto L117
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v254)+4)) = v359
	v361 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v254))), uint32(v361))
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+65)))
	v371 = v364
	goto L118
L118:
	;
	v393 = v172 << (uint(base.B2i32(v371 == int32(0))) % 32)
	v394 = F_mul_size(m, int32(12), v393)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L5
	} else {
		goto L120
	}
L119:
	;
	if v225 != 0 {
		goto L150
	} else {
		goto L151
	}
L120:
	;
	v396 = F_add_size(m, int32(8), v394)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L5
	} else {
		goto L121
	}
L121:
	;
	v398 = F_palloc(m, v396)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L5
	} else {
		goto L122
	}
L122:
	;
	v401 = F_LWLockAcquire(m, v163, int32(1))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L5
	} else {
		goto L123
	}
L123:
	;
	if v231 != 0 {
		goto L126
	} else {
		goto L127
	}
L124:
	;
	if v396 != 0 {
		goto L130
	} else {
		goto L131
	}
L125:
	;
	v421 = v410 + v231 - int32(1)
	goto L124
L126:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v147)+72))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v231+v403+v371<<(uint(int32(2))%32)-int32(1))))
	if v410 != 0 {
		goto L125
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v147)+72))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v412+v371<<(uint(int32(2))%32))))
	v421 = v416
	goto L124
L129:
	;
	v421 = int32(0)
	goto L124
L130:
	;
	base.MemoryCopy(m, v398, v421, v396)
	goto L132
L131:
	;
	goto L132
L132:
	;
	F_LWLockRelease(m, v163)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L5
	} else {
		goto L133
	}
L133:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v398)))
	if int32(0) < v425 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v433 = int32(0)
	goto L137
L135:
	;
	goto L136
L136:
	;
	if int32(0) < v371 {
		v371 = v371 - int32(1)
		goto L118
	} else {
		goto L149
	}
L137:
	;
	v457 = v398 + int32(8) + v433*int32(12)
	if v231 != 0 {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	goto L136
L139:
	;
	v493 = *(*float32)(unsafe.Add(mBase, uint32(v457)+4))
	v494 = int32(0)
	F_HnswUpdateConnection(m, v231, v490, v147, v493, v393, v494, v494, v43)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L5
	} else {
		goto L146
	}
L140:
	;
	v458 = int32(0)
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v457)))
	v460 = v231 + v459
	v462 = v460 + int32(91)
	v464 = F_LWLockAcquire(m, v462, v458)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L5
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v457)))
	v481 = v479 + int32(92)
	v483 = F_LWLockAcquire(m, v481, int32(0))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L5
	} else {
		goto L145
	}
L143:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v460)+71))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v231+v466+v371<<(uint(int32(2))%32)-int32(1))))
	if v473 == int32(0) {
		v490 = v458
		v491 = v462
		goto L139
	} else {
		goto L144
	}
L144:
	;
	v490 = v231 + v473 - int32(1)
	v491 = v462
	goto L139
L145:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v479)+72))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v485+v371<<(uint(int32(2))%32))))
	v490 = v489
	v491 = v481
	goto L139
L146:
	;
	F_LWLockRelease(m, v491)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L5
	} else {
		goto L147
	}
L147:
	;
	v501 = v433 + int32(1)
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v398)))
	if v501 < v502 {
		v433 = v501
		goto L137
	} else {
		goto L148
	}
L148:
	;
	goto L138
L149:
	;
	goto L119
L150:
	;
	v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+65)))
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+65)))
	if base.Ui32(v532) <= base.Ui32(v533) {
		goto L93
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	if v231 == int32(0) {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	goto L152
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v254)+48)) = v147
	goto L93
L155:
	;
	goto L156
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v254)+48)) = v358
	goto L93
L157:
	;
	F_LWLockRelease(m, v309)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L5
	} else {
		goto L158
	}
L158:
	;
	goto L93
L159:
	;
	F_LWLockRelease(m, v49)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L5
	} else {
		goto L160
	}
L160:
	;
	goto L18
L161:
	;
	F_s_lock(m, v32, int32(_a_F_BuildCallback_1_5), int32(601), int32(_a_F_BuildCallback_1_8))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L5
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	v615 = *(*float64)(unsafe.Add(mBase, uint32(v32)+8))
	v617 = base.F64_add(v615, float64(1))
	*(*float64)(unsafe.Add(mBase, uint32(v32)+8)) = v617
	v623 = *(*int32)(unsafe.Add(mBase, _c_F_BuildCallback_1[2]))
	if v623 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L164:
	;
	goto L163
L165:
	;
	v656 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v32))), uint32(v656))
	goto L4
L166:
	;
	goto L165
L167:
	;
	v627 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BuildCallback_1[3])))
	if v627&int32(1) == int32(0) {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v632 = int32(_a_F_BuildCallback_1_9)
	v634 = *(*int32)(unsafe.Add(mBase, _c_F_BuildCallback_1[4]))
	v635 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BuildCallback_1[4])) = v634 + v635
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v623)))
	*(*int32)(unsafe.Add(mBase, uint32(v623))) = v638 + v635
	*(*int64)(unsafe.Add(mBase, uint32(v623+int32(96))+232)) = base.I64_trunc_sat_f64_s(v617)
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v623)))
	*(*int32)(unsafe.Add(mBase, uint32(v623))) = v646 + v635
	v652 = *(*int32)(unsafe.Add(mBase, _c_F_BuildCallback_1[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_BuildCallback_1[4])) = v652 - v635
	goto L166
L169:
	;
	goto L3
}
func F_BuildSpeculativeIndexInfo(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+10)))
	v16 = v14 << (uint(int32(2)) % 32)
	v17 = F_palloc(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+104)) = v17
	v20 = F_palloc(m, v16)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+108)) = v20
	v25 = F_palloc(m, v14<<(uint(int32(1))%32))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+112)) = v25
	if int32(0) < v14 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L16
	}
L6:
	;
	v34 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	m.G0 = v11 + int32(16)
	return
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+84))
	v42 = v34 << (uint(int32(2)) % 32)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42+v43)))
	v47 = F_IndexAmTranslateCompareType(m, int32(3), v40, v45, int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v50 = v34 << (uint(int32(1)) % 32)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	*(*uint16)(unsafe.Add(mBase, uint32(v50+v51))) = uint16(v47)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v54+v42)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v57+v42)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	v62 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60+v50))))
	v63 = F_get_opfamily_member(m, v56, v59, v59, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v65+v42))) = v63
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v68+v42)))
	if v70 == int32(0) {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v73 = F_get_opcode(m, v70)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v75+v42))) = v73
	v79 = v34 + int32(1)
	if v79 != v14 {
		v34 = v79
		goto L9
	} else {
		goto L15
	}
L15:
	;
	goto L10
L16:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v96+v34<<(uint(int32(1))%32)))))
	v102 = v34 << (uint(int32(2)) % 32)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v102+v103)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v106+v102)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v100
	F_errmsg_internal(m, int32(_a_F_BuildSpeculativeIndexInfo_0), v11)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_BuildSpeculativeIndexInfo_1), int32(2705), int32(_a_F_BuildSpeculativeIndexInfo_2))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_begin_prepare_cb_wrapper(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int64
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_begin_prepare_cb_wrapper_0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v10
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(993)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v14
	v18 = int32(_a_F_begin_prepare_cb_wrapper_1)
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_begin_prepare_cb_wrapper[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_begin_prepare_cb_wrapper[0])) = v8 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v8 + int32(16)
	v28 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+147)) = uint8(v28)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+160)) = v30
	v32 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+164)) = uint8(v3)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+152)) = v32
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	if v36 == v3 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_begin_prepare_cb_wrapper_2)
				F_errmsg(m, int32(_a_F_begin_prepare_cb_wrapper_3), v8)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_begin_prepare_cb_wrapper_4), int32(941), int32(_a_F_begin_prepare_cb_wrapper_5))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
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
		m.T0[v36].(func(*base.Module, int32, int32))(m, v10, l1)
		mBase = m.M
		v57 = m.ExcPending
		if v57 != 0 {
			return
		} else {
			v59 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			*(*int32)(unsafe.Add(mBase, _c_F_begin_prepare_cb_wrapper[0])) = v59
			m.G0 = v8 + int32(32)
			return
		}
	}
}
func F_beginmerge(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v94 int64
	_ = v94
	var v105 int32
	_ = v105
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v13 < v14 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L10
	} else {
		goto L29
	}
L2:
	;
	v16 = v13
	goto L4
L3:
	;
	v16 = v14
	goto L4
L4:
	;
	if int32(0) < v16 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v23 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	m.G0 = v11 + int32(16)
	return
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27+v23<<(uint(int32(2))%32))))
	v33 = F_LogicalTapeRead(m, v31, v11, int32(4))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	return
L11:
	;
	if v33 != int32(4) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v37 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	m.T0[v38].(func(*base.Module, int32, int32, int32, int32))(m, l0, v11, v31, v37)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L10
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v105 = v23 + int32(1)
	if v105 != v16 {
		v23 = v105
		goto L8
	} else {
		goto L28
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v23
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_beginmerge[0]))
	if v44 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L10
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v47 + int32(1)
	if v47 <= int32(0) {
		v82 = v47
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L19
L21:
	;
	v91 = v42 + v82<<(uint(int32(4))%32)
	v92 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v91)+8)) = v92
	v94 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	*(*int64)(unsafe.Add(mBase, uint32(v91))) = v94
	goto L15
L22:
	;
	v54 = v47
	goto L23
L23:
	;
	v61 = int32(1)
	v64 = int32(base.Ui32(v54-v61) >> (uint(v61) % 32))
	v67 = v42 + v64<<(uint(int32(4))%32)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v69 = m.T0[v68].(func(*base.Module, int32, int32, int32) int32)(m, v11, v67, l0)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L10
	} else {
		goto L25
	}
L24:
	;
	v82 = int32(0)
	goto L21
L25:
	;
	if int32(0) <= v69 {
		v82 = v54
		goto L21
	} else {
		goto L26
	}
L26:
	;
	v75 = v42 + v54<<(uint(int32(4))%32)
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v67)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v75)+8)) = v76
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v67)))
	*(*int64)(unsafe.Add(mBase, uint32(v75))) = v78
	if v64 != 0 {
		v54 = v64
		goto L23
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	goto L9
L29:
	;
	F_errmsg_internal(m, int32(_a_F_beginmerge_0), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L10
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_beginmerge_1), int32(2862), int32(_a_F_beginmerge_2))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_bernoulli_samplescangetsamplesize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 float64
	_ = v8
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 float32
	_ = v18
	var v21 int32
	_ = v21
	var v37 float64
	_ = v37
	var v39 int32
	_ = v39
	var v41 float64
	_ = v41
	var v42 float64
	_ = v42
	var v43 float64
	_ = v43
	var v52 float64
	_ = v52
	var v56 float64
	_ = v56
	v8 = float64(0.10000000149011612)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v11 = F_estimate_expression_value(m, l0, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		if v13 != int32(7) {
			v37 = v8
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+24)))
			if v16 != 0 {
				v37 = v8
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
				v18 = base.F32_reinterpret_i32(v17)
				v21 = int32(0)
				if base.B2i32(base.F32_ge(v18, float32(0)) == v21)|base.B2i32(base.F32_le(v18, float32(100)) == v21)|base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v17&int32(2147483647))) != 0 {
					v37 = v8
				} else {
					v37 = base.F64_promote_f32(base.F32_div(v18, float32(100)))
				}
			}
		}
		v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = v39
		v41 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
		v42 = base.F64_mul(v37, v41)
		v43 = float64(1e+100)
		if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v42)&int64(9223372036854775807)))|base.F64_gt(v42, v43) != 0 {
			v56 = v43
		} else {
			v52 = float64(1)
			if base.F64_le(v42, v52) != 0 {
				v56 = v52
			} else {
				v56 = base.F64_nearest(v42)
			}
		}
		*(*float64)(unsafe.Add(mBase, uint32(l4))) = v56
		return
	}
}
func F_bf_init(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+92))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+84)) = l2
	if l2 != 0 {
		base.MemoryCopy(m, v6+int32(4), l1, l2)
	} else {
	}
	v14 = v6 + int32(68)
	if l3 != 0 {
		if v8 == int32(0) {
			return int32(0)
		} else {
			base.MemoryCopy(m, v14, l3, v8)
			return int32(0)
		}
	} else {
		if v8 == int32(0) {
		} else {
			base.MemoryFill(m, v14, int32(0), v8)
		}
		return int32(0)
	}
}
func F_big5_to_euc_tw(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v17, v18, v19, int32(36), int32(4))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v19 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v192 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v184))) = uint8(v192)
	m.G0 = v12 + int32(16)
	return v183 - v16
L4:
	;
	v183 = v16
	v184 = v15
	goto L3
L5:
	;
	goto L6
L6:
	;
	v28 = v16
	v29 = v15
	v30 = v19
	goto L7
L7:
	;
	v37 = int32(*(*int8)(unsafe.Add(mBase, uint32(v28))))
	if v37 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v183 = v180
	v184 = v175
	goto L3
L9:
	;
	if int32(0) < v176 {
		v28 = v180
		v29 = v175
		v30 = v176
		goto L7
	} else {
		goto L60
	}
L10:
	;
	v41 = F_pg_encoding_verifymbchar(m, int32(36), v28, v30)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if v37 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L13:
	;
	if v41 < int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v14 != 0 {
		v183 = v28
		v184 = v29
		goto L3
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	v53 = (v48 | v37<<(uint(int32(8))%32)) & int32(_a_F_big5_to_euc_tw_0)
	v55 = v12 + int32(15)
	if base.Ui32(v53) <= base.Ui32(int32(_a_F_big5_to_euc_tw_1)) {
		goto L32
	} else {
		goto L33
	}
L17:
	;
	F_report_invalid_encoding(m, int32(36), v28, v30)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L19:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	switch v118 - int32(149) {
	case 0:
		goto L47
	case 1:
		goto L49
	default:
		goto L48
	}
L20:
	;
	v117 = v115 & int32(_a_F_big5_to_euc_tw_0)
	goto L19
L21:
	;
	v110 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v110)
	v115 = int32(63)
	goto L20
L22:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v104)
	v115 = v103 | int32(-32640)
	goto L20
L23:
	;
	v98 = int32(246)
	*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v98)
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97)+2)))
	v115 = v100 | int32(-32640)
	goto L20
L24:
	;
	v97 = int32(_a_F_big5_to_euc_tw_2)
	goto L23
L25:
	;
	v97 = int32(_a_F_big5_to_euc_tw_3)
	goto L23
L26:
	;
	v97 = int32(_a_F_big5_to_euc_tw_4)
	goto L23
L27:
	;
	v97 = int32(_a_F_big5_to_euc_tw_5)
	goto L23
L28:
	;
	v97 = int32(_a_F_big5_to_euc_tw_6)
	goto L23
L29:
	;
	v97 = int32(_a_F_big5_to_euc_tw_7)
	goto L23
L30:
	;
	v87 = F_BinarySearchRange(m, int32(_a_F_big5_to_euc_tw_8), int32(46), v53)
	mBase = m.M
	if v87 == int32(0) {
		goto L21
	} else {
		goto L45
	}
L31:
	;
	v103 = v82
	v104 = int32(149)
	goto L22
L32:
	;
	switch v53 - int32(_a_F_big5_to_euc_tw_9) {
	case 0:
		v70 = int32(_a_F_big5_to_euc_tw_10)
		goto L35
	case 1, 3:
		goto L39
	case 2:
		goto L38
	case 4:
		goto L37
	default:
		goto L40
	}
L33:
	;
	goto L34
L34:
	;
	switch v53 - int32(_a_F_big5_to_euc_tw_11) {
	case 0:
		v97 = int32(_a_F_big5_to_euc_tw_12)
		goto L23
	case 1:
		goto L29
	case 2:
		goto L28
	case 3:
		goto L27
	case 4:
		goto L26
	case 5:
		goto L25
	case 6:
		goto L24
	default:
		goto L43
	}
L35:
	;
	v71 = int32(247)
	*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v71)
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+2)))
	v115 = v73 | int32(-32640)
	goto L20
L36:
	;
	v70 = int32(_a_F_big5_to_euc_tw_13)
	goto L35
L37:
	;
	v70 = int32(_a_F_big5_to_euc_tw_14)
	goto L35
L38:
	;
	v70 = int32(_a_F_big5_to_euc_tw_15)
	goto L35
L39:
	;
	v66 = F_BinarySearchRange(m, int32(_a_F_big5_to_euc_tw_16), int32(23), v53)
	mBase = m.M
	if v66 != 0 {
		v82 = v66
		goto L31
	} else {
		goto L42
	}
L40:
	;
	if v53 == int32(_a_F_big5_to_euc_tw_17) {
		goto L36
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	goto L21
L43:
	;
	if v53 != int32(_a_F_big5_to_euc_tw_18) {
		goto L30
	} else {
		goto L44
	}
L44:
	;
	v82 = int32(_a_F_big5_to_euc_tw_19)
	goto L31
L45:
	;
	v103 = v87
	v104 = int32(150)
	goto L22
L46:
	;
	v175 = v160
	v176 = v30 - v41
	v180 = v28 + v41
	goto L9
L47:
	;
	v151 = int32(8)
	v155 = v117<<(uint(v151)%32) | int32(base.Ui32(v117)>>(uint(v151)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v29))) = uint16(v155)
	v160 = v29 + int32(2)
	goto L46
L48:
	;
	if base.Ui32((v118+int32(10))&int32(255)) <= base.Ui32(int32(4)) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+3)) = uint8(v117)
	v122 = int32(_a_F_big5_to_euc_tw_20)
	*(*uint16)(unsafe.Add(mBase, uint32(v29))) = uint16(v122)
	v125 = int32(base.Ui32(v117) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+2)) = uint8(v125)
	v160 = v29 + int32(4)
	goto L46
L50:
	;
	v135 = int32(142)
	*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v135)
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+3)) = uint8(v117)
	v140 = int32(base.Ui32(v117) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+2)) = uint8(v140)
	v143 = v137 - int32(83)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)) = uint8(v143)
	v160 = v29 + int32(4)
	goto L46
L51:
	;
	goto L52
L52:
	;
	if v14 != 0 {
		v183 = v28
		v184 = v29
		goto L3
	} else {
		goto L53
	}
L53:
	;
	F_report_untranslatable_char(m, int32(36), int32(4), v28, v30)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	if v14 != 0 {
		v183 = v28
		v184 = v29
		goto L3
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v37)
	v169 = int32(1)
	v175 = v29 + v169
	v176 = v30 - v169
	v180 = v28 + v169
	goto L9
L58:
	;
	F_report_invalid_encoding(m, int32(36), v28, v30)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	goto L8
}
func F_binaryheap_replace_first(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = l1
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if int32(2) <= v10 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = l0 + int32(20)
	v17 = v10
	v20 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v23 = int32(1)
	v24 = v20 << (uint(v23) % 32)
	v26 = v24 | v23
	v28 = v24 + int32(2)
	if v28 < v17 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+v20<<(uint(int32(2))%32)))) = l1
	goto L3
L6:
	;
	goto L5
L7:
	;
	v30 = int32(2)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v14+v26<<(uint(v30)%32))))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v14+v28<<(uint(v30)%32))))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v40 = m.T0[v39].(func(*base.Module, int32, int32, int32) int32)(m, v33, v37, v38)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v46 = v26
	v47 = v17
	goto L9
L9:
	;
	if v47 <= v26 {
		goto L6
	} else {
		goto L15
	}
L10:
	;
	return
L11:
	;
	if v40 < int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v44 = v28
	goto L14
L13:
	;
	v44 = v26
	goto L14
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v46 = v44
	v47 = v45
	goto L9
L15:
	;
	v51 = v14 + v46<<(uint(int32(2))%32)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v55 = m.T0[v54].(func(*base.Module, int32, int32, int32) int32)(m, l1, v52, v53)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	if int32(0) <= v55 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	*(*int32)(unsafe.Add(mBase, uint32(v14+v20<<(uint(int32(2))%32)))) = v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = v64
	v20 = v46
	goto L4
}
func F_binaryheap_reset(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	return
}
func F_bind(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	v4 = int32(0)
	v7 = m.Env.X__syscall_bind(m, l0, l1, l2, v4, v4, v4)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v7) {
		*(*int32)(unsafe.Add(mBase, _c_F_bind[0])) = int32(0) - v7
		v15 = int32(-1)
	} else {
		v15 = v7
	}
	return v15
}
func F_bitcat(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = F_pg_detoast_datum(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = F_bit_catenate(m, v3, v8)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		}
	}
}
func F_biteq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v14 == v15 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v17 = int32(8)
	v18 = v7 + v17
	v20 = v12 + v17
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v22 = int32(2)
	v23 = int32(base.Ui32(v21) >> (uint(v22) % 32))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v26 = int32(base.Ui32(v24) >> (uint(v22) % 32))
	if base.Ui32(v23) < base.Ui32(v26) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v95 = int32(0)
	goto L6
L6:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v97 != v7 {
		goto L28
	} else {
		goto L29
	}
L7:
	;
	v28 = v23
	goto L9
L8:
	;
	v28 = v26
	goto L9
L9:
	;
	v30 = v28 - int32(8)
	if base.Ui32(int32(4)) <= base.Ui32(v30) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v95 = base.B2i32(v92 == int32(0))
	goto L6
L11:
	;
	v92 = int32(0)
	goto L10
L12:
	;
	v66 = v61
	v67 = v62
	v68 = v63
	goto L22
L13:
	;
	if (v18|v20)&int32(3) != 0 {
		v61 = v18
		v62 = v20
		v63 = v30
		goto L12
	} else {
		goto L16
	}
L14:
	;
	v54 = v18
	v55 = v20
	v56 = v30
	goto L15
L15:
	;
	if v56 == int32(0) {
		goto L11
	} else {
		goto L21
	}
L16:
	;
	v38 = v18
	v39 = v20
	v40 = v30
	goto L17
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v43 != v44 {
		v61 = v38
		v62 = v39
		v63 = v40
		goto L12
	} else {
		goto L19
	}
L18:
	;
	v54 = v49
	v55 = v47
	v56 = v51
	goto L15
L19:
	;
	v46 = int32(4)
	v47 = v39 + v46
	v49 = v38 + v46
	v51 = v40 - v46
	if base.Ui32(int32(3)) < base.Ui32(v51) {
		v38 = v49
		v39 = v47
		v40 = v51
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v61 = v54
	v62 = v55
	v63 = v56
	goto L12
L22:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v71 == v72 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v92 = v71 - v72
	goto L10
L24:
	;
	v74 = int32(1)
	v79 = v68 - v74
	if v79 != 0 {
		v66 = v66 + v74
		v67 = v67 + v74
		v68 = v79
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	goto L23
L27:
	;
	goto L11
L28:
	;
	F_pfree(m, v7)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v101 != v12 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L30
L32:
	;
	F_pfree(m, v12)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	return v95
L35:
	;
	goto L34
}
func F_bitge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v102 != v8 {
		goto L31
	} else {
		goto L32
	}
L2:
	;
	return int32(0)
L3:
	;
	v13 = v8 + int32(8)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v18 = v15 + int32(8)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v20 = int32(2)
	v21 = int32(base.Ui32(v19) >> (uint(v20) % 32))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v24 = int32(base.Ui32(v22) >> (uint(v20) % 32))
	if base.Ui32(v21) < base.Ui32(v24) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v26 = v21
	goto L7
L6:
	;
	v26 = v24
	goto L7
L7:
	;
	v28 = v26 - int32(8)
	if base.Ui32(int32(4)) <= base.Ui32(v28) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	if v90 != 0 {
		v99 = v90
		goto L1
	} else {
		goto L26
	}
L9:
	;
	v90 = int32(0)
	goto L8
L10:
	;
	v64 = v59
	v65 = v60
	v66 = v61
	goto L20
L11:
	;
	if (v13|v18)&int32(3) != 0 {
		v59 = v13
		v60 = v18
		v61 = v28
		goto L10
	} else {
		goto L14
	}
L12:
	;
	v52 = v13
	v53 = v18
	v54 = v28
	goto L13
L13:
	;
	if v54 == int32(0) {
		goto L9
	} else {
		goto L19
	}
L14:
	;
	v36 = v13
	v37 = v18
	v38 = v28
	goto L15
L15:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v41 != v42 {
		v59 = v36
		v60 = v37
		v61 = v38
		goto L10
	} else {
		goto L17
	}
L16:
	;
	v52 = v47
	v53 = v45
	v54 = v49
	goto L13
L17:
	;
	v44 = int32(4)
	v45 = v37 + v44
	v47 = v36 + v44
	v49 = v38 - v44
	if base.Ui32(int32(3)) < base.Ui32(v49) {
		v36 = v47
		v37 = v45
		v38 = v49
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v59 = v52
	v60 = v53
	v61 = v54
	goto L10
L20:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v69 == v70 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v90 = v69 - v70
	goto L8
L22:
	;
	v72 = int32(1)
	v77 = v66 - v72
	if v77 != 0 {
		v64 = v64 + v72
		v65 = v65 + v72
		v66 = v77
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	goto L21
L25:
	;
	goto L9
L26:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v92 == v93 {
		v99 = int32(0)
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v92 < v93 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v98 = int32(-1)
	goto L30
L29:
	;
	v98 = int32(1)
	goto L30
L30:
	;
	v99 = v98
	goto L1
L31:
	;
	F_pfree(m, v8)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L2
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v106 != v15 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L33
L35:
	;
	F_pfree(m, v15)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L2
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	return int32(base.Ui32(v99^int32(-1)) >> (uint(int32(31)) % 32))
L38:
	;
	goto L37
}
func F_bitoverlay(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = F_pg_detoast_datum(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v12 = F_bit_overlay(m, v3, v8, v10, v11)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return v12
			}
		}
	}
}
func F_bitoverlay_no_len(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v9 = F_pg_detoast_datum(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			v13 = F_bit_overlay(m, v4, v9, v11, v12)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				return v13
			}
		}
	}
}
func F_blbuildempty(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_BloomInitMetapage(m, l0, int32(3))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_blbulkdelete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
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
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 float64
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
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
	var v142 int32
	_ = v142
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v222 int32
	_ = v222
	var v230 int32
	_ = v230
	var v241 int32
	_ = v241
	var v249 int32
	_ = v249
	var v260 int32
	_ = v260
	var v268 int32
	_ = v268
	var v280 int32
	_ = v280
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	v5 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(_a_F_blbulkdelete_0)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l1 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v26 = F_palloc0(m, int32(40))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v30 = l1
	goto L3
L3:
	;
	F_initBloomState(m, v20, v22)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return int32(0)
L5:
	;
	v30 = v26
	goto L3
L6:
	;
	v34 = F_RelationGetNumberOfBlocksInFork(m, v22, int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v34) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v49 = v5
	v52 = int32(1)
	goto L11
L9:
	;
	v280 = v5
	goto L10
L10:
	;
	v288 = F_ReadBuffer(m, v22, int32(0))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L4
	} else {
		goto L56
	}
L11:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L4
	} else {
		goto L13
	}
L12:
	;
	v280 = v260
	goto L10
L13:
	;
	v59 = int32(0)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v62 = F_ReadBufferExtended(m, v22, v59, v52, v59, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	F_LockBuffer(m, v62, int32(2))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v67 = F_GenericXLogStart(m, v22)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	v268 = v52 + int32(1)
	if v268 != v34 {
		v49 = v260
		v52 = v268
		goto L11
	} else {
		goto L55
	}
L17:
	;
	v86 = v70 + int32(24)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v20)+1164))
	v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74))))
	v89 = int32(1)
	v95 = v87 * ((v88+v89)&int32(_a_F_blbulkdelete_1) - v89)
	if v95 != 0 {
		goto L31
	} else {
		goto L32
	}
L18:
	;
	v70 = F_GenericXLogRegisterBuffer(m, v67, v62, int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+14)))
	if v72 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+16)))
	v74 = v70 + v73
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+2)))
	if v75&int32(2) == int32(0) {
		goto L17
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	F_UnlockReleaseBuffer(m, v62)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	F_pfree(m, v67)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v260 = v49
	goto L16
L26:
	;
	F_UnlockReleaseBuffer(m, v62)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L4
	} else {
		goto L54
	}
L27:
	;
	F_pfree(m, v67)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L4
	} else {
		goto L53
	}
L28:
	;
	v208 = v196 - v70
	*(*uint16)(unsafe.Add(mBase, uint32(v70)+12)) = uint16(v208)
	F_GenericXLogFinish(m, v67)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L4
	} else {
		goto L52
	}
L29:
	;
	if v138 == v135 {
		v222 = v49
		goto L27
	} else {
		goto L51
	}
L30:
	;
	if base.B2i32(base.Ui32(int32(_a_F_blbulkdelete_2)-v160*v161) < base.Ui32(v160))|base.B2i32(base.Ui32(int32(2003)) < base.Ui32(v49)) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L31:
	;
	v101 = v86
	v102 = v86
	goto L34
L32:
	;
	goto L33
L33:
	;
	if v88 == int32(0) {
		v222 = v49
		goto L27
	} else {
		goto L46
	}
L34:
	;
	v114 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, v101, l3)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+16)))
	v141 = v70 + v140
	v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141))))
	if v142 == int32(0) {
		goto L29
	} else {
		goto L45
	}
L36:
	;
	v138 = v101 + v137
	if base.Ui32(v138) < base.Ui32(v95+v86) {
		v101 = v138
		v102 = v135
		goto L34
	} else {
		goto L44
	}
L37:
	;
	if v114 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+16)))
	v117 = v70 + v116
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117))))
	v120 = v118 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v117))) = uint16(v120)
	v122 = *(*float64)(unsafe.Add(mBase, uint32(v30)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v30)+16)) = base.F64_add(v122, float64(1))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v20)+1164))
	v135 = v102
	v137 = v126
	goto L36
L39:
	;
	goto L40
L40:
	;
	if v101 == v102 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v20)+1164))
	v135 = v102 + v133
	v137 = v133
	goto L36
L42:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v20)+1164))
	if v128 == int32(0) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	base.MemoryCopy(m, v102, v101, v128)
	goto L41
L44:
	;
	goto L35
L45:
	;
	v153 = base.B2i32(v138 == v135)
	v154 = v135
	v160 = v137
	v161 = v142
	goto L30
L46:
	;
	v153 = int32(1)
	v154 = v86
	v160 = v87
	v161 = v88
	goto L30
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(1168)+v49<<(uint(int32(2))%32)))) = v52
	v183 = v49 + int32(1)
	goto L49
L48:
	;
	v183 = v49
	goto L49
L49:
	;
	if v153 == int32(0) {
		v196 = v154
		v201 = v183
		goto L28
	} else {
		goto L50
	}
L50:
	;
	v222 = v183
	goto L27
L51:
	;
	v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141)+2)))
	v189 = v187 | int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v141)+2)) = uint16(v189)
	v196 = v135
	v201 = v49
	goto L28
L52:
	;
	v241 = v201
	goto L26
L53:
	;
	v241 = v222
	goto L26
L54:
	;
	v260 = v241
	goto L16
L55:
	;
	goto L12
L56:
	;
	F_LockBuffer(m, v288, int32(2))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	v293 = F_GenericXLogStart(m, v22)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	v296 = F_GenericXLogRegisterBuffer(m, v293, v288, int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	v299 = v280 << (uint(int32(2)) % 32)
	if v299 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	base.MemoryCopy(m, v296+int32(168), v20+int32(1168), v299)
	goto L62
L61:
	;
	goto L62
L62:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v296)+30)) = uint16(v280)
	v306 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v296)+28)) = uint16(v306)
	F_GenericXLogFinish(m, v293)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	F_UnlockReleaseBuffer(m, v288)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	m.G0 = v20 + int32(_a_F_blbulkdelete_0)
	return v30
}
func F_blcostestimate(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v32 float64
	_ = v32
	var v35 int32
	_ = v35
	var v36 float64
	_ = v36
	var v38 float64
	_ = v38
	var v40 float64
	_ = v40
	var v42 float64
	_ = v42
	var v44 float64
	_ = v44
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v16 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+56)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v16
	v32 = *(*float64)(unsafe.Add(mBase, uint32(v15)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v13)+40)) = v32
	F_genericcostestimate(m, l0, l1, l2, v13)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		return
	} else {
		v36 = *(*float64)(unsafe.Add(mBase, uint32(v13)))
		*(*float64)(unsafe.Add(mBase, uint32(l3))) = v36
		v38 = *(*float64)(unsafe.Add(mBase, uint32(v13)+8))
		*(*float64)(unsafe.Add(mBase, uint32(l4))) = v38
		v40 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
		*(*float64)(unsafe.Add(mBase, uint32(l5))) = v40
		v42 = *(*float64)(unsafe.Add(mBase, uint32(v13)+24))
		*(*float64)(unsafe.Add(mBase, uint32(l6))) = v42
		v44 = *(*float64)(unsafe.Add(mBase, uint32(v13)+32))
		*(*float64)(unsafe.Add(mBase, uint32(l7))) = v44
		m.G0 = v13 - int32(-64)
		return
	}
}
func F_boolgt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	return base.B2i32(v2 == v3) & base.B2i32(v5 != v3)
}
func F_boot_yy_create_buffer(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	v8 = F_palloc(m, int32(48))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(_a_F_boot_yy_create_buffer_0)
			v15 = F_palloc(m, int32(_a_F_boot_yy_create_buffer_1))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v15
				if v15 == int32(0) {
					F_yy_fatal_error_1(m, int32(_a_F_boot_yy_create_buffer_2))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v20 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v20
					v23 = *(*int32)(unsafe.Add(mBase, _c_F_boot_yy_create_buffer[0]))
					v24 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v24
					*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v24)
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)) = uint8(v24)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v24
					*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v20
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v35
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					if v37 == v24 {
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						v43 = v37 + v40<<(uint(int32(2))%32)
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
						if v8 != v44 {
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v46
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v49
							*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v49
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v53
							v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v55)
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					if v62 != 0 {
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v62+v63<<(uint(int32(2))%32))))
						if v8 == v67 {
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = int64(1)
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = int64(1)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(0)
					*(*int32)(unsafe.Add(mBase, _c_F_boot_yy_create_buffer[0])) = v23
					return v8
				}
			}
		} else {
			F_yy_fatal_error_1(m, int32(_a_F_boot_yy_create_buffer_2))
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_boot_yyensure_buffer_stack(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v4 == int32(0) {
		v8 = F_palloc(m, int32(4))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v8
			if v8 == int32(0) {
				F_yy_fatal_error_1(m, int32(_a_F_boot_yyensure_buffer_stack_0))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = int64(4294967296)
				return
			}
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if base.Ui32(v18-int32(1)) <= base.Ui32(v17) {
			v23 = v18 + int32(8)
			v26 = F_repalloc(m, v4, v23<<(uint(int32(2))%32))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v26
				if v26 == int32(0) {
					F_yy_fatal_error_1(m, int32(_a_F_boot_yyensure_buffer_stack_0))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v34 = v26 + v31<<(uint(int32(2))%32)
					v35 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v34)+24)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+16)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+8)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34))) = v35
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v23
					return
				}
			}
		} else {
			return
		}
	}
}
func F_bpcharin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_strlen(m, v3)
	mBase = m.M
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = F_bpchar_input(m, v3, v4, v5, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_bpchartruelen(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	v5 = int32(-1)
	v7 = l1 - int32(1)
	if v5 <= v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = v5
	goto L3
L2:
	;
	v10 = v7
	goto L3
L3:
	;
	v14 = l1
	goto L4
L4:
	;
	v18 = v14 - int32(1)
	if v18 < int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	return v14
L6:
	;
	return v10 + int32(1)
L7:
	;
	goto L8
L8:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v18))))
	if v23 == int32(32) {
		v14 = v18
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
}
func F_bpchartypmodout(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_palloc(m, int32(64))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if int32(5) <= v8 {
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8 - int32(4)
			v21 = F_pg_snprintf(m, v10, int32(64), int32(_a_F_bpchartypmodout_0), v6)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v10
			}
		} else {
			v23 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v10))) = uint8(v23)
			m.G0 = v6 + int32(16)
			return v10
		}
	}
}
func F_brinbuildCallbackParallel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 float64
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	v9 = m.G0
	v10 = int32(16)
	v11 = v9 - v10
	m.G0 = v11
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v17 = v13 | v14<<(uint(v10)%32)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l5)+32))
	if base.Ui32(v18) <= base.Ui32(v17) {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
		if base.Ui32(v17) <= base.Ui32(v18+v20-int32(1)) {
			v54 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
			v55 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
			v56 = F_add_values_to_range(m, l0, v54, v55, l2, l3)
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return
			} else {
				m.G0 = v11 + int32(16)
				return
			}
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
			if v26 == int32(0) {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
				v32 = F_brin_form_tuple(m, v29, v18, v25, v11+int32(12))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l5)+72))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
					F_tuplesort_putbrintuple(m, v34, v32, v35)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						v38 = *(*float64)(unsafe.Add(mBase, uint32(l5)+8))
						*(*float64)(unsafe.Add(mBase, uint32(l5)+8)) = base.F64_add(v38, float64(1))
						F_pfree(m, v32)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
							v45 = v44
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
							v47 = base.I32_rem_u_s(v17, v46)
							*(*int32)(unsafe.Add(mBase, uint32(l5)+32)) = v17 - v47
							v50 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
							F_brin_memtuple_initialize(m, v45, v50)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								v54 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
								v56 = F_add_values_to_range(m, l0, v54, v55, l2, l3)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return
								} else {
									m.G0 = v11 + int32(16)
									return
								}
							}
						}
					}
				}
			} else {
				v45 = v25
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
				v47 = base.I32_rem_u_s(v17, v46)
				*(*int32)(unsafe.Add(mBase, uint32(l5)+32)) = v17 - v47
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
				F_brin_memtuple_initialize(m, v45, v50)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
					v56 = F_add_values_to_range(m, l0, v54, v55, l2, l3)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						m.G0 = v11 + int32(16)
						return
					}
				}
			}
		}
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
		v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
		if v26 == int32(0) {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
			v32 = F_brin_form_tuple(m, v29, v18, v25, v11+int32(12))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l5)+72))
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
				F_tuplesort_putbrintuple(m, v34, v32, v35)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					v38 = *(*float64)(unsafe.Add(mBase, uint32(l5)+8))
					*(*float64)(unsafe.Add(mBase, uint32(l5)+8)) = base.F64_add(v38, float64(1))
					F_pfree(m, v32)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
						v45 = v44
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
						v47 = base.I32_rem_u_s(v17, v46)
						*(*int32)(unsafe.Add(mBase, uint32(l5)+32)) = v17 - v47
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
						F_brin_memtuple_initialize(m, v45, v50)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
							v56 = F_add_values_to_range(m, l0, v54, v55, l2, l3)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return
							} else {
								m.G0 = v11 + int32(16)
								return
							}
						}
					}
				}
			}
		} else {
			v45 = v25
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
			v47 = base.I32_rem_u_s(v17, v46)
			*(*int32)(unsafe.Add(mBase, uint32(l5)+32)) = v17 - v47
			v50 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
			F_brin_memtuple_initialize(m, v45, v50)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return
			} else {
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
				v56 = F_add_values_to_range(m, l0, v54, v55, l2, l3)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					m.G0 = v11 + int32(16)
					return
				}
			}
		}
	}
}
func F_brinrescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	if l1 == int32(0) {
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v8 <= int32(0) {
		} else {
			v12 = v8 * int32(48)
			if v12 == int32(0) {
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				base.MemoryCopy(m, v15, l1, v12)
			}
		}
	}
	return
}
func F_brinvalidate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int64
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v66 int64
	_ = v66
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
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int64
	_ = v205
	var v209 int32
	_ = v209
	var v210 int64
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v243 int64
	_ = v243
	var v244 int32
	_ = v244
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v264 int64
	_ = v264
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v318 int64
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v396 int32
	_ = v396
	var v405 int64
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int64
	_ = v448
	var v451 int32
	_ = v451
	var v453 int64
	_ = v453
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int64
	_ = v488
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v546 int64
	_ = v546
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v592 int32
	_ = v592
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	var v728 int32
	_ = v728
	var v733 int32
	_ = v733
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v782 int32
	_ = v782
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v832 int32
	_ = v832
	var v842 int32
	_ = v842
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	v15 = int64(0)
	v17 = m.G0
	v19 = v17 - int32(272)
	m.G0 = v19
	v22 = F_SearchSysCache1(m, int32(14), l0)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v37)+40))
	if int32(0) < v244 {
		goto L49
	} else {
		goto L50
	}
L2:
	;
	return int32(0)
L3:
	;
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+22)))
	v28 = v26 + v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+84))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+80))
	v31 = F_get_opfamily_name(m, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
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
	v218 = m.ExcPending
	if v218 != 0 {
		goto L2
	} else {
		goto L46
	}
L7:
	;
	v35 = int32(0)
	v37 = F_SearchSysCacheList(m, int32(4), int32(1), v30, v35, v35)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v39 = int32(1)
	v42 = int32(0)
	v44 = F_SearchSysCacheList(m, int32(5), v39, v30, v42, v42)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+40))
	if v46 <= int32(0) {
		v233 = v39
		v243 = v15
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v54 = int32(0)
	v56 = v39
	v66 = v15
	goto L11
L11:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v44+int32(48)+v54<<(uint(int32(2))%32))))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+56))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+22)))
	v73 = v71 + v72
	v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73)+16)))
	switch v74 - int32(1) {
	case 0:
		goto L16
	case 1:
		goto L21
	case 2:
		goto L20
	case 3:
		goto L19
	case 4:
		goto L18
	default:
		goto L17
	}
L12:
	;
	v233 = v209
	v243 = v210
	goto L1
L13:
	;
	v212 = v54 + int32(1)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v44)+40))
	if v212 < v213 {
		v54 = v212
		v56 = v209
		v66 = v210
		goto L11
	} else {
		goto L45
	}
L14:
	;
	v205 = int64(*(*int16)(unsafe.Add(mBase, uint32(v73)+16)))
	v209 = v203
	v210 = int64(1)<<(uint(v205)%64) | v66
	goto L13
L15:
	;
	v173 = int32(0)
	v176 = F_errstart(m, int32(17), v173)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L2
	} else {
		goto L39
	}
L16:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	v163 = int32(2281)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+208)) = v163
	v166 = int32(1)
	v171 = F_check_amproc_signature(m, v162, v163, v166, v166, v166, v19+int32(208))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L2
	} else {
		goto L37
	}
L17:
	;
	if base.Ui32(int32(_a_F_brinvalidate_0)) < base.Ui32((v74-int32(16))&int32(_a_F_brinvalidate_1)) {
		v203 = v56
		goto L14
	} else {
		goto L30
	}
L18:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	v123 = F_check_amoptsproc_signature(m, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L2
	} else {
		goto L28
	}
L19:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+264)) = int32(2281)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+256)) = int64(9796820404457)
	v114 = int32(3)
	v118 = F_check_amproc_signature(m, v107, int32(16), int32(1), v114, v114, v19+int32(256))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L2
	} else {
		goto L26
	}
L20:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+248)) = int64(98784250089)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+240)) = int64(9796820404457)
	v103 = F_check_amproc_signature(m, v92, int32(16), int32(1), int32(3), int32(4), v19+int32(240))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L2
	} else {
		goto L24
	}
L21:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	v78 = int64(9796820404457)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+232)) = v78
	*(*int64)(unsafe.Add(mBase, uint32(v19)+224)) = v78
	v84 = int32(4)
	v88 = F_check_amproc_signature(m, v77, int32(16), int32(1), v84, v84, v19+int32(224))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	if v88 == int32(0) {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	v203 = v56
	goto L14
L24:
	;
	if v103 == int32(0) {
		goto L15
	} else {
		goto L25
	}
L25:
	;
	v203 = v56
	goto L14
L26:
	;
	if v118 == int32(0) {
		goto L15
	} else {
		goto L27
	}
L27:
	;
	v203 = v56
	goto L14
L28:
	;
	if v123 == int32(0) {
		goto L15
	} else {
		goto L29
	}
L29:
	;
	v203 = v56
	goto L14
L30:
	;
	v133 = int32(0)
	v136 = F_errstart(m, int32(17), v133)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	if v136 == int32(0) {
		v209 = v133
		v210 = v66
		goto L13
	} else {
		goto L32
	}
L32:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	v144 = F_format_procedure(m, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	v146 = int32(*(*int16)(unsafe.Add(mBase, uint32(v73)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+188)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v19)+184)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v19)+180)) = int32(_a_F_brinvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+176)) = v31
	F_errmsg(m, int32(_a_F_brinvalidate_3), v19+int32(176))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_brinvalidate_4), int32(114), int32(_a_F_brinvalidate_5))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	v209 = v133
	v210 = v66
	goto L13
L37:
	;
	if v171 != 0 {
		v203 = v56
		goto L14
	} else {
		goto L38
	}
L38:
	;
	goto L15
L39:
	;
	if v176 == int32(0) {
		v203 = v173
		goto L14
	} else {
		goto L40
	}
L40:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	v184 = F_format_procedure(m, v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v186 = int32(*(*int16)(unsafe.Add(mBase, uint32(v73)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+204)) = v186
	*(*int32)(unsafe.Add(mBase, uint32(v19)+200)) = v184
	*(*int32)(unsafe.Add(mBase, uint32(v19)+196)) = int32(_a_F_brinvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+192)) = v31
	F_errmsg(m, int32(_a_F_brinvalidate_6), v19+int32(192))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_brinvalidate_4), int32(130), int32(_a_F_brinvalidate_5))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	v203 = v173
	goto L14
L45:
	;
	goto L12
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = l0
	F_errmsg_internal(m, int32(_a_F_brinvalidate_7), v19)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_brinvalidate_4), int32(58), int32(_a_F_brinvalidate_5))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L2
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	v253 = int32(0)
	v255 = v233
	v264 = v15
	goto L52
L50:
	;
	v396 = v233
	v405 = v15
	goto L51
L51:
	;
	v408 = v28 + int32(8)
	v409 = F_identify_opfamily_groups(m, v37, v44)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L2
	} else {
		goto L97
	}
L52:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(48)+v253<<(uint(int32(2))%32))))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+56))
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270)+22)))
	v272 = v270 + v271
	v273 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v272)+16)))
	if base.Ui32((v273+int32(-64))&int32(_a_F_brinvalidate_1)) <= base.Ui32(int32(_a_F_brinvalidate_8)) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v396 = v386
	v405 = v318
	goto L51
L54:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+18)))
	if v319 == int32(115) {
		goto L66
	} else {
		goto L67
	}
L55:
	;
	v280 = int32(0)
	v283 = F_errstart(m, int32(17), v280)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L2
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v272)+8))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v272)+12))
	if v309 != v310 {
		v317 = v255
		v318 = v264
		goto L54
	} else {
		goto L64
	}
L58:
	;
	if v283 == int32(0) {
		v317 = v280
		v318 = v264
		goto L54
	} else {
		goto L59
	}
L59:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L2
	} else {
		goto L60
	}
L60:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v272)+20))
	v291 = F_format_operator(m, v290)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	v293 = int32(*(*int16)(unsafe.Add(mBase, uint32(v272)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+172)) = v293
	*(*int32)(unsafe.Add(mBase, uint32(v19)+168)) = v291
	*(*int32)(unsafe.Add(mBase, uint32(v19)+164)) = int32(_a_F_brinvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+160)) = v31
	F_errmsg(m, int32(_a_F_brinvalidate_9), v19+int32(160))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L2
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_brinvalidate_4), int32(152), int32(_a_F_brinvalidate_5))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	v317 = v280
	v318 = v264
	goto L54
L64:
	;
	v317 = v255
	v318 = int64(1)<<(uint(base.I64_extend_i32_u(v273))%64) | v264
	goto L54
L65:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v272)+20))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v272)+8))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v272)+12))
	v357 = F_check_amop_signature(m, v353, int32(16), v355, v356)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L2
	} else {
		goto L77
	}
L66:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v272)+28))
	if v322 == int32(0) {
		v352 = v317
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v325 = int32(0)
	v328 = F_errstart(m, int32(17), v325)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L2
	} else {
		goto L70
	}
L69:
	;
	goto L68
L70:
	;
	if v328 == int32(0) {
		v352 = v325
		goto L65
	} else {
		goto L71
	}
L71:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v272)+20))
	v336 = F_format_operator(m, v335)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+152)) = v336
	*(*int32)(unsafe.Add(mBase, uint32(v19)+148)) = int32(_a_F_brinvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+144)) = v31
	F_errmsg(m, int32(_a_F_brinvalidate_10), v19+int32(144))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L2
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_brinvalidate_4), int32(180), int32(_a_F_brinvalidate_5))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L2
	} else {
		goto L75
	}
L75:
	;
	v352 = v325
	goto L65
L76:
	;
	v388 = v253 + int32(1)
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v37)+40))
	if v388 < v389 {
		v253 = v388
		v255 = v386
		v264 = v318
		goto L52
	} else {
		goto L85
	}
L77:
	;
	if v357 != 0 {
		v386 = v352
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v359 = int32(0)
	v362 = F_errstart(m, int32(17), v359)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L2
	} else {
		goto L79
	}
L79:
	;
	if v362 == int32(0) {
		v386 = v359
		goto L76
	} else {
		goto L80
	}
L80:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L2
	} else {
		goto L81
	}
L81:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v272)+20))
	v370 = F_format_operator(m, v369)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+136)) = v370
	*(*int32)(unsafe.Add(mBase, uint32(v19)+132)) = int32(_a_F_brinvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+128)) = v31
	F_errmsg(m, int32(_a_F_brinvalidate_11), v19+int32(128))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L2
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_brinvalidate_4), int32(193), int32(_a_F_brinvalidate_5))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	v386 = v359
	goto L76
L85:
	;
	goto L53
L86:
	;
	F_ReleaseCatCacheList(m, v858)
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L2
	} else {
		goto L187
	}
L87:
	;
	v823 = int32(0)
	v826 = F_errstart(m, int32(17), v823)
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L2
	} else {
		goto L182
	}
L88:
	;
	v804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v788))))
	if v804&int32(16) != 0 {
		v849 = v789
		v853 = v793
		v856 = v796
		v858 = v798
		v860 = v800
		goto L86
	} else {
		goto L181
	}
L89:
	;
	v765 = int32(0)
	v768 = F_errstart(m, int32(17), v765)
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L2
	} else {
		goto L173
	}
L90:
	;
	v744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v728))))
	if v744&int32(8) != 0 {
		v788 = v728
		v789 = v19
		v793 = v733
		v794 = v408
		v796 = v37
		v798 = v44
		v800 = v22
		goto L88
	} else {
		goto L172
	}
L91:
	;
	v699 = int32(0)
	v702 = F_errstart(m, int32(17), v699)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L2
	} else {
		goto L163
	}
L92:
	;
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v663))))
	if v679&int32(4) != 0 {
		v728 = v663
		v733 = v668
		goto L90
	} else {
		goto L162
	}
L93:
	;
	v636 = int32(0)
	v639 = F_errstart(m, int32(17), v636)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L2
	} else {
		goto L153
	}
L94:
	;
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599))))
	if v615&int32(2) != 0 {
		v663 = v599
		v668 = v604
		goto L92
	} else {
		goto L152
	}
L95:
	;
	v599 = v531 + int32(16)
	v604 = v532
	goto L94
L96:
	;
	v567 = F_errstart(m, int32(17), int32(0))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L2
	} else {
		goto L139
	}
L97:
	;
	if v409 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v551 = int32(1)
	v553 = int32(0)
	goto L96
L99:
	;
	goto L100
L100:
	;
	v416 = int32(0)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v409)+4))
	if v416 < v417 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v424 = int32(0)
	v425 = v416
	v426 = v396
	goto L104
L102:
	;
	v531 = v416
	v532 = v396
	goto L103
L103:
	;
	if v531 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L104:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v409)+12))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v437+v424<<(uint(int32(2))%32))))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v441)))
	if v29 == v442 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v531 = v447
	v532 = v521
	goto L103
L106:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v441)+4))
	if v444 == v29 {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	v447 = v425
	goto L108
L108:
	;
	v448 = *(*int64)(unsafe.Add(mBase, uint32(v441)+16))
	if v448 == int64(0) {
		goto L113
	} else {
		goto L114
	}
L109:
	;
	v446 = v441
	goto L111
L110:
	;
	v446 = v425
	goto L111
L111:
	;
	v447 = v446
	goto L108
L112:
	;
	v524 = v424 + int32(1)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v409)+4))
	if v524 < v525 {
		v424 = v524
		v425 = v447
		v426 = v521
		goto L104
	} else {
		goto L134
	}
L113:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v441)+4))
	if v442 != v451 {
		v521 = v426
		goto L112
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v453 = *(*int64)(unsafe.Add(mBase, uint32(v441)+8))
	if v453 == v405 {
		v486 = v426
		goto L117
	} else {
		goto L118
	}
L116:
	;
	goto L115
L117:
	;
	v488 = *(*int64)(unsafe.Add(mBase, uint32(v441)+16))
	if v488 == v243 {
		v521 = v486
		goto L112
	} else {
		goto L126
	}
L118:
	;
	v455 = int32(0)
	v458 = F_errstart(m, int32(17), v455)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L2
	} else {
		goto L119
	}
L119:
	;
	if v458 == int32(0) {
		v486 = v455
		goto L117
	} else {
		goto L120
	}
L120:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L2
	} else {
		goto L121
	}
L121:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v441)))
	v466 = F_format_type_be(m, v465)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L2
	} else {
		goto L122
	}
L122:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v441)+4))
	v469 = F_format_type_be(m, v468)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L2
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+124)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v19)+120)) = v466
	*(*int32)(unsafe.Add(mBase, uint32(v19)+116)) = int32(_a_F_brinvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+112)) = v31
	F_errmsg(m, int32(_a_F_brinvalidate_12), v19+int32(112))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L2
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(_a_F_brinvalidate_4), int32(232), int32(_a_F_brinvalidate_5))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L2
	} else {
		goto L125
	}
L125:
	;
	v486 = v455
	goto L117
L126:
	;
	v490 = int32(0)
	v493 = F_errstart(m, int32(17), v490)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L2
	} else {
		goto L127
	}
L127:
	;
	if v493 == int32(0) {
		v521 = v490
		goto L112
	} else {
		goto L128
	}
L128:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L2
	} else {
		goto L129
	}
L129:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v441)))
	v501 = F_format_type_be(m, v500)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L2
	} else {
		goto L130
	}
L130:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v441)+4))
	v504 = F_format_type_be(m, v503)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L2
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+108)) = v504
	*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = v501
	*(*int32)(unsafe.Add(mBase, uint32(v19)+100)) = int32(_a_F_brinvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v31
	F_errmsg(m, int32(_a_F_brinvalidate_13), v19+int32(96))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L2
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_brinvalidate_4), int32(242), int32(_a_F_brinvalidate_5))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L2
	} else {
		goto L133
	}
L133:
	;
	v521 = v490
	goto L112
L134:
	;
	goto L105
L135:
	;
	v551 = int32(1)
	v553 = int32(0)
	goto L96
L136:
	;
	goto L137
L137:
	;
	v546 = *(*int64)(unsafe.Add(mBase, uint32(v531)+8))
	if v546 == v405 {
		goto L95
	} else {
		goto L138
	}
L138:
	;
	v551 = int32(0)
	v553 = v531
	goto L96
L139:
	;
	if v567 != 0 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L2
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v592 = v553 + int32(16)
	if v551 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+84)) = int32(_a_F_brinvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v408
	F_errmsg(m, int32(_a_F_brinvalidate_14), v19+int32(80))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L2
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_brinvalidate_4), int32(253), int32(_a_F_brinvalidate_5))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L2
	} else {
		goto L145
	}
L145:
	;
	v586 = v553 + int32(16)
	if v551 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v599 = v586
	v604 = int32(0)
	goto L94
L147:
	;
	goto L148
L148:
	;
	v619 = v586
	v635 = int32(1)
	goto L93
L149:
	;
	v599 = v592
	v604 = int32(0)
	goto L94
L150:
	;
	goto L151
L151:
	;
	v619 = v592
	v635 = int32(1)
	goto L93
L152:
	;
	v619 = v599
	v635 = int32(0)
	goto L93
L153:
	;
	if v639 != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L2
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	if v635 == int32(0) {
		v663 = v619
		v668 = v636
		goto L92
	} else {
		goto L161
	}
L157:
	;
	v644 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = int32(_a_F_brinvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v408
	F_errmsg(m, int32(_a_F_brinvalidate_15), v19-int32(-64))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L2
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(_a_F_brinvalidate_4), int32(264), int32(_a_F_brinvalidate_5))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L2
	} else {
		goto L159
	}
L159:
	;
	if v635 != 0 {
		v683 = v619
		v686 = v644
		goto L91
	} else {
		goto L160
	}
L160:
	;
	v663 = v619
	v668 = v636
	goto L92
L161:
	;
	v683 = v619
	v686 = int32(1)
	goto L91
L162:
	;
	v683 = v663
	v686 = int32(0)
	goto L91
L163:
	;
	if v702 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L2
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	if v686 == int32(0) {
		v728 = v683
		v733 = v699
		goto L90
	} else {
		goto L171
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = int32(_a_F_brinvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v408
	F_errmsg(m, int32(_a_F_brinvalidate_15), v19+int32(48))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L2
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_brinvalidate_4), int32(264), int32(_a_F_brinvalidate_5))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L2
	} else {
		goto L169
	}
L169:
	;
	if v686 == int32(0) {
		v728 = v683
		v733 = v699
		goto L90
	} else {
		goto L170
	}
L170:
	;
	v748 = v683
	v749 = v19
	v754 = v408
	v756 = v37
	v758 = v44
	v760 = v22
	v764 = int32(1)
	goto L89
L171:
	;
	v748 = v683
	v749 = v19
	v754 = v408
	v756 = v37
	v758 = v44
	v760 = v22
	v764 = int32(1)
	goto L89
L172:
	;
	v748 = v728
	v749 = v19
	v754 = v408
	v756 = v37
	v758 = v44
	v760 = v22
	v764 = int32(0)
	goto L89
L173:
	;
	if v768 != 0 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L2
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	if v764 != 0 {
		v808 = v749
		v813 = v754
		v815 = v756
		v817 = v758
		v819 = v760
		goto L87
	} else {
		goto L180
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v749)+40)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v749)+36)) = int32(_a_F_brinvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v749)+32)) = v754
	F_errmsg(m, int32(_a_F_brinvalidate_15), v749+int32(32))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L2
	} else {
		goto L178
	}
L178:
	;
	F_errfinish(m, int32(_a_F_brinvalidate_4), int32(264), int32(_a_F_brinvalidate_5))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L2
	} else {
		goto L179
	}
L179:
	;
	goto L176
L180:
	;
	v788 = v748
	v789 = v749
	v793 = v765
	v794 = v754
	v796 = v756
	v798 = v758
	v800 = v760
	goto L88
L181:
	;
	v808 = v789
	v813 = v794
	v815 = v796
	v817 = v798
	v819 = v800
	goto L87
L182:
	;
	if v826 == int32(0) {
		v849 = v808
		v853 = v823
		v856 = v815
		v858 = v817
		v860 = v819
		goto L86
	} else {
		goto L183
	}
L183:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L2
	} else {
		goto L184
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v808)+24)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v808)+20)) = int32(_a_F_brinvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v808)+16)) = v813
	F_errmsg(m, int32(_a_F_brinvalidate_15), v808+int32(16))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L2
	} else {
		goto L185
	}
L185:
	;
	F_errfinish(m, int32(_a_F_brinvalidate_4), int32(264), int32(_a_F_brinvalidate_5))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L2
	} else {
		goto L186
	}
L186:
	;
	v849 = v808
	v853 = v823
	v856 = v815
	v858 = v817
	v860 = v819
	goto L86
L187:
	;
	F_ReleaseCatCacheList(m, v856)
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L2
	} else {
		goto L188
	}
L188:
	;
	F_ReleaseCatCache(m, v860)
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L2
	} else {
		goto L189
	}
L189:
	;
	m.G0 = v849 + int32(272)
	return v853
}
func F_bsearch_arg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	if l2 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v14 = l1
	v15 = l2
	goto L4
L4:
	;
	v24 = v14 + int32(base.Ui32(v15)>>(uint(int32(1))%32))*l3
	v25 = m.T0[l4].(func(*base.Module, int32, int32, int32) int32)(m, l0, v24, l5)
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	return int32(0)
L6:
	;
	return int32(0)
L7:
	;
	if v25 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return v24
L9:
	;
	goto L10
L10:
	;
	v34 = base.B2i32(int32(0) < v25)
	if int32(0) < v25 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v35 = l3 + v24
	goto L13
L12:
	;
	v35 = v14
	goto L13
L13:
	;
	v38 = int32(base.Ui32(v15-v34) >> (uint(int32(1)) % 32))
	if v38 != 0 {
		v14 = v35
		v15 = v38
		goto L4
	} else {
		goto L14
	}
L14:
	;
	goto L5
}
func F_btcostestimate(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v25 float64
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v58 int32
	_ = v58
	var v62 float64
	_ = v62
	var v63 int32
	_ = v63
	var v66 float64
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 float64
	_ = v69
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 float64
	_ = v91
	var v93 float64
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v133 float64
	_ = v133
	var v134 float64
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 float64
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 float32
	_ = v151
	var v153 float32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v180 float64
	_ = v180
	var v181 float64
	_ = v181
	var v185 int32
	_ = v185
	var v186 float64
	_ = v186
	var v189 float64
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 float64
	_ = v196
	var v199 int32
	_ = v199
	var v206 float64
	_ = v206
	var v209 float64
	_ = v209
	var v210 int32
	_ = v210
	var v215 float64
	_ = v215
	var v216 int32
	_ = v216
	var v222 float64
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 float32
	_ = v242
	var v243 float64
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 float64
	_ = v247
	var v250 int32
	_ = v250
	var v253 float64
	_ = v253
	var v255 int32
	_ = v255
	var v257 float64
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 float64
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 float64
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v344 int32
	_ = v344
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 float64
	_ = v396
	var v397 int32
	_ = v397
	var v401 float64
	_ = v401
	var v402 float64
	_ = v402
	var v405 float64
	_ = v405
	var v433 float64
	_ = v433
	var v436 float64
	_ = v436
	var v437 float64
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v497 int32
	_ = v497
	var v502 float64
	_ = v502
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v525 int32
	_ = v525
	var v529 float64
	_ = v529
	var v530 float64
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
	var v560 float64
	_ = v560
	var v561 float64
	_ = v561
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v594 float64
	_ = v594
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 float64
	_ = v613
	var v614 int32
	_ = v614
	var v618 float64
	_ = v618
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 float64
	_ = v647
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 float64
	_ = v668
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v710 float64
	_ = v710
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 float64
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v742 float64
	_ = v742
	var v744 float64
	_ = v744
	var v746 int32
	_ = v746
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v834 int32
	_ = v834
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v885 float64
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v888 float64
	_ = v888
	var v890 int32
	_ = v890
	var v894 float64
	_ = v894
	var v896 float64
	_ = v896
	var v897 float64
	_ = v897
	var v900 float64
	_ = v900
	var v927 float64
	_ = v927
	var v931 float64
	_ = v931
	var v937 int32
	_ = v937
	var v939 float64
	_ = v939
	var v940 float64
	_ = v940
	var v941 float64
	_ = v941
	var v942 float64
	_ = v942
	var v947 float64
	_ = v947
	var v948 float64
	_ = v948
	var v952 float64
	_ = v952
	var v955 float64
	_ = v955
	var v957 float64
	_ = v957
	var v959 float64
	_ = v959
	var v960 int32
	_ = v960
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v969 float64
	_ = v969
	var v970 float64
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v990 float32
	_ = v990
	var v991 float64
	_ = v991
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 float64
	_ = v995
	var v998 int32
	_ = v998
	var v1001 float64
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1005 float64
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1014 float64
	_ = v1014
	var v1020 float64
	_ = v1020
	var v1026 float64
	_ = v1026
	var v1029 float64
	_ = v1029
	v9 = int32(0)
	v25 = float64(0)
	v29 = m.G0
	v31 = v29 - int32(176)
	m.G0 = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v34 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+128)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v31)+120)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v31)+112)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v31)+104)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v31)+96)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v31)+88)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v31)+80)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v31)+72)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v31)+64)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v31)+56)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v31)+48)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v31)+40)) = v34
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v58 == v9 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v728)+101)))
	if v746 == int32(1) {
		goto L144
	} else {
		goto L145
	}
L2:
	;
	v67 = l0
	v68 = l1
	v69 = l2
	v70 = l3
	v71 = l4
	v72 = l5
	v73 = l6
	v74 = l7
	v75 = v31
	v77 = v33
	v78 = v9
	v79 = v9
	v80 = v9
	v82 = v9
	v84 = v9
	v85 = v58
	v86 = v9
	v87 = v9
	v88 = v9
	v90 = v9
	v91 = v62
	v93 = v25
	goto L8
L3:
	;
	v718 = l0
	v719 = l1
	v720 = l2
	v721 = l3
	v722 = l4
	v723 = l5
	v724 = l6
	v725 = l7
	v726 = v31
	v728 = v33
	v729 = v9
	v731 = v9
	v733 = v9
	v735 = v9
	v739 = v9
	v741 = v9
	v742 = v66
	v744 = v25
	goto L1
L4:
	;
	v66 = float64(1)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v62 = float64(1)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if int32(0) < v63 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v66 = v62
	goto L3
L8:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v95+v86<<(uint(int32(2))%32))))
	v100 = int32(*(*int16)(unsafe.Add(mBase, uint32(v99)+14)))
	if v100 <= v80 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v718 = v67
	v719 = v68
	v720 = v69
	v721 = v70
	v722 = v71
	v723 = v72
	v724 = v73
	v725 = v74
	v726 = v75
	v728 = v77
	v729 = v697
	v731 = v548
	v733 = v701
	v735 = v703
	v739 = v556
	v741 = v709
	v742 = v710
	v744 = v561
	goto L1
L10:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v99)+8))
	if v563 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L11:
	;
	v546 = v78
	v547 = v79
	v548 = v80
	v550 = v82
	v556 = v88
	v560 = v91
	v561 = v93
	goto L10
L12:
	;
	goto L13
L13:
	;
	if v87 != 0 {
		v718 = v67
		v719 = v68
		v720 = v69
		v721 = v70
		v722 = v71
		v723 = v72
		v724 = v73
		v725 = v74
		v726 = v75
		v728 = v77
		v729 = v78
		v731 = v80
		v733 = v82
		v735 = v84
		v739 = v88
		v741 = v90
		v742 = v91
		v744 = v93
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v104 = v78 & int32(1)
	if v104 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v105 = int32(0)
	goto L17
L16:
	;
	v105 = v79
	goto L17
L17:
	;
	v106 = v104 + v80
	if v106 < v100 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v532 = int32(0)
	v533 = int32(*(*int16)(unsafe.Add(mBase, uint32(v99)+14)))
	if v517 == v533 {
		v546 = v532
		v547 = v516
		v548 = v517
		v550 = v519
		v556 = v525
		v560 = v529
		v561 = v530
		goto L10
	} else {
		goto L112
	}
L19:
	;
	v120 = v105
	v121 = v106
	v129 = v88
	v133 = v91
	v134 = v93
	goto L22
L20:
	;
	v488 = v105
	v489 = v106
	v491 = v82
	v497 = v88
	v502 = v93
	goto L21
L21:
	;
	v516 = v488
	v517 = v489
	v519 = v491
	v525 = v497
	v529 = v91
	v530 = v502
	goto L18
L22:
	;
	v137 = v75 + int32(40)
	F_examine_indexcol_variable(m, v67, v77, v121, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v488 = v120
	v489 = v121
	v491 = int32(1)
	v497 = v271
	v502 = v272
	goto L21
L24:
	;
	return
L25:
	;
	v141 = v75 + int32(39)
	v142 = int32(0)
	v143 = float64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v141))) = uint8(v142)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v137)+8))
	if v147 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v75)+48))
	if v121 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L27:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+28)))
	if v185 != 0 {
		goto L41
	} else {
		goto L42
	}
L28:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+16))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+22)))
	v150 = v148 + v149
	v151 = *(*float32)(unsafe.Add(mBase, uint32(v150)+8))
	v153 = *(*float32)(unsafe.Add(mBase, uint32(v150)+16))
	v180 = base.F64_promote_f32(v153)
	v181 = base.F64_promote_f32(v151)
	goto L27
L29:
	;
	goto L30
L30:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
	if v155 == int32(16) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v180 = float64(2)
	v181 = v143
	goto L27
L32:
	;
	goto L33
L33:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if v159 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	if v166 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v159)+76))
	if v162 != int32(5) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v180 = float64(-1)
	v181 = v143
	goto L27
L37:
	;
	v180 = float64(0)
	v181 = v143
	goto L27
L38:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	if v169 != int32(6) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166)+8)))
	switch v173 - int32(_a_F_btcostestimate_0) {
	case 0:
		goto L40
	default:
		goto L37
	case 5:
		v180 = float64(-1)
		v181 = v143
		goto L27
	}
L40:
	;
	v180 = float64(1)
	v181 = v143
	goto L27
L41:
	;
	v186 = base.F64_neg(base.F64_sub(float64(1), v181))
	goto L43
L42:
	;
	v186 = v180
	goto L43
L43:
	;
	if base.F64_gt(v186, float64(0)) != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v189 = F_clamp_row_est(m, v186)
	mBase = m.M
	v215 = v189
	goto L26
L45:
	;
	goto L46
L46:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if v190 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v193 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v141))) = uint8(v193)
	v215 = float64(200)
	goto L26
L48:
	;
	goto L49
L49:
	;
	v196 = *(*float64)(unsafe.Add(mBase, uint32(v190)+120))
	if base.F64_le(v196, float64(0)) != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v199 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v141))) = uint8(v199)
	v215 = float64(200)
	goto L26
L51:
	;
	goto L52
L52:
	;
	if base.F64_lt(v186, float64(0)) != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v206 = F_clamp_row_est(m, base.F64_mul(v196, base.F64_neg(v186)))
	mBase = m.M
	v215 = v206
	goto L26
L54:
	;
	goto L55
L55:
	;
	if base.F64_lt(v196, float64(200)) != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v209 = F_clamp_row_est(m, v196)
	mBase = m.M
	v215 = v209
	goto L26
L57:
	;
	goto L58
L58:
	;
	v210 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v141))) = uint8(v210)
	v215 = float64(200)
	goto L26
L59:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+39)))
	if v273 != 0 {
		goto L80
	} else {
		goto L81
	}
L60:
	;
	if v216 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v260 = v216
	v262 = v129
	v263 = v134
	goto L62
L62:
	;
	if v260 == int32(0) {
		v271 = v262
		v272 = v263
		goto L59
	} else {
		goto L78
	}
L63:
	;
	v271 = int32(1)
	v272 = v134
	goto L59
L64:
	;
	goto L65
L65:
	;
	v222 = float64(0)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v77)+52))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v77)+56))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	v228 = F_get_opfamily_member(m, v224, v226, v226, int32(1))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L24
	} else {
		goto L67
	}
L66:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v75)+48))
	v260 = v259
	v262 = int32(1)
	v263 = v257
	goto L62
L67:
	;
	if v228 == int32(0) {
		v257 = v222
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v233 = v75 + int32(140)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v75)+48))
	v237 = F_get_attstatsslot(m, v233, v234, int32(3), v228, int32(2))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L24
	} else {
		goto L69
	}
L69:
	;
	if v237 == int32(0) {
		v257 = v222
		goto L66
	} else {
		goto L70
	}
L70:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v75)+160))
	v242 = *(*float32)(unsafe.Add(mBase, uint32(v241)))
	v243 = base.F64_promote_f32(v242)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v77)+64))
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	if v246 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v247 = base.F64_neg(v243)
	goto L73
L72:
	;
	v247 = v243
	goto L73
L73:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v77)+40))
	if int32(1) < v250 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v253 = base.F64_mul(v247, float64(0.75))
	goto L76
L75:
	;
	v253 = v247
	goto L76
L76:
	;
	F_free_attstatsslot(m, v233)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L24
	} else {
		goto L77
	}
L77:
	;
	v257 = v253
	goto L66
L78:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v75)+52))
	m.T0[v266].(func(*base.Module, int32))(m, v260)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L24
	} else {
		goto L79
	}
L79:
	;
	v271 = v262
	v272 = v263
	goto L59
L80:
	;
	goto L23
L81:
	;
	if v120 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v77)+88))
	if v274 != 0 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	v433 = v215
	goto L84
L84:
	;
	if v120 != 0 {
		goto L107
	} else {
		goto L108
	}
L85:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+4))
	if v275 <= int32(0) {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	v391 = v120
	goto L87
L87:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v392)+68))
	v394 = int32(0)
	v396 = F_clauselist_selectivity(m, v67, v391, v393, v394, v394)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L24
	} else {
		goto L102
	}
L88:
	;
	v361 = F_list_concat(m, v344, v120)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L24
	} else {
		goto L101
	}
L89:
	;
	v344 = int32(0)
	goto L88
L90:
	;
	goto L91
L91:
	;
	v279 = int32(0)
	v290 = v279
	v292 = v279
	goto L92
L92:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v274)+12))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v309+v290<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v75)+32)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v75)+140)) = v313
	v319 = F_list_make1_impl(m, int32(1), v75+int32(32))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L24
	} else {
		goto L94
	}
L93:
	;
	v344 = v328
	goto L88
L94:
	;
	v322 = F_predicate_implied_by(m, v319, v120, int32(0))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L24
	} else {
		goto L95
	}
L95:
	;
	if v322 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v326 = F_list_concat(m, v292, v319)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L24
	} else {
		goto L99
	}
L97:
	;
	v328 = v292
	goto L98
L98:
	;
	v330 = v290 + int32(1)
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v274)+4))
	if v330 < v331 {
		v290 = v330
		v292 = v328
		goto L92
	} else {
		goto L100
	}
L99:
	;
	v328 = v326
	goto L98
L100:
	;
	goto L93
L101:
	;
	v391 = v361
	goto L87
L102:
	;
	if base.F64_lt(v396, float64(0.005)) != 0 {
		goto L80
	} else {
		goto L103
	}
L103:
	;
	v401 = base.F64_nearest(base.F64_mul(v215, v396))
	v402 = float64(1)
	if base.F64_gt(v401, v402) != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v405 = v401
	goto L106
L105:
	;
	v405 = v402
	goto L106
L106:
	;
	v433 = v405
	goto L84
L107:
	;
	v436 = v433
	goto L109
L108:
	;
	v436 = base.F64_add(v433, float64(1))
	goto L109
L109:
	;
	v437 = base.F64_mul(v133, v436)
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	if base.F64_gt(v437, base.F64_convert_i32_u(v438)) != 0 {
		goto L80
	} else {
		goto L110
	}
L110:
	;
	v441 = int32(1)
	v442 = int32(0)
	v444 = v121 + v441
	v445 = int32(*(*int16)(unsafe.Add(mBase, uint32(v99)+14)))
	if v445 <= v444 {
		v516 = v442
		v517 = v444
		v519 = v441
		v525 = v271
		v529 = v437
		v530 = v272
		goto L18
	} else {
		goto L111
	}
L111:
	;
	v120 = v442
	v121 = v444
	v129 = v271
	v133 = v437
	v134 = v272
	goto L22
L112:
	;
	v718 = v67
	v719 = v68
	v720 = v69
	v721 = v70
	v722 = v71
	v723 = v72
	v724 = v73
	v725 = v74
	v726 = v75
	v728 = v77
	v729 = v532
	v731 = v517
	v733 = v519
	v735 = v84
	v739 = v525
	v741 = v90
	v742 = v529
	v744 = v530
	goto L1
L113:
	;
	v715 = v86 + int32(1)
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	if v715 < v716 {
		v78 = v697
		v79 = v698
		v80 = v548
		v82 = v701
		v84 = v703
		v86 = v715
		v87 = v706
		v88 = v556
		v90 = v709
		v91 = v710
		v93 = v561
		goto L8
	} else {
		goto L142
	}
L114:
	;
	v697 = v546
	v698 = v547
	v701 = v550
	v703 = v84
	v706 = v87
	v709 = v90
	v710 = v560
	goto L113
L115:
	;
	goto L116
L116:
	;
	v566 = int32(0)
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v563)+4))
	if v567 <= v566 {
		v697 = v546
		v698 = v547
		v701 = v550
		v703 = v84
		v706 = v87
		v709 = v90
		v710 = v560
		goto L113
	} else {
		goto L117
	}
L117:
	;
	v579 = v566
	v581 = v546
	v582 = v547
	v585 = v550
	v587 = v84
	v590 = v87
	v593 = v90
	v594 = v560
	goto L118
L118:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v563)+12))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v598+v579<<(uint(int32(2))%32))))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v602)+4))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v603)))
	switch v604 - int32(17) {
	case 0:
		goto L122
	default:
		goto L123
	case 3:
		goto L125
	case 20:
		goto L126
	case 35:
		goto L124
	}
L119:
	;
	v697 = v663
	v698 = v681
	v701 = v665
	v703 = v670
	v706 = v666
	v709 = v667
	v710 = v668
	goto L113
L120:
	;
	v670 = F_lappend(m, v587, v602)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L24
	} else {
		goto L136
	}
L121:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v649)))
	if v650 == int32(0) {
		v663 = v581
		v665 = v645
		v666 = v646
		v667 = v593
		v668 = v647
		goto L120
	} else {
		goto L134
	}
L122:
	;
	v645 = v585
	v646 = v590
	v647 = v594
	v649 = v603 + int32(4)
	goto L121
L123:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L24
	} else {
		goto L131
	}
L124:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v603)+8))
	v624 = base.B2i32(v622 == int32(0))
	v663 = v624 | v581
	v665 = v585
	v666 = v590
	v667 = v624 | v593
	v668 = v594
	goto L120
L125:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v603)+28))
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v610)+12))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v611)+4))
	v613 = F_estimate_array_length(m, v67, v612)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L24
	} else {
		goto L127
	}
L126:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v603)+8))
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v608)+12))
	v645 = v585
	v646 = int32(1)
	v647 = v594
	v649 = v609
	goto L121
L127:
	;
	if base.F64_gt(v613, float64(1)) != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v618 = base.F64_mul(v594, v613)
	goto L130
L129:
	;
	v618 = v594
	goto L130
L130:
	;
	v645 = int32(1)
	v646 = v590
	v647 = v618
	v649 = v603 + int32(4)
	goto L121
L131:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v603)))
	*(*int32)(unsafe.Add(mBase, uint32(v75)+16)) = v631
	F_errmsg_internal(m, int32(_a_F_btcostestimate_1), v75+int32(16))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L24
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_btcostestimate_2), int32(_a_F_btcostestimate_3), int32(_a_F_btcostestimate_4))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L24
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L134:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v77)+52))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v653+v548<<(uint(int32(2))%32))))
	v658 = F_get_op_opfamily_strategy(m, v650, v657)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L24
	} else {
		goto L135
	}
L135:
	;
	v663 = base.B2i32(v658 == int32(3)) | v581
	v665 = v645
	v666 = v646
	v667 = v593
	v668 = v647
	goto L120
L136:
	;
	if v663&int32(1)|v666 != 0 {
		v681 = v582
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v683 = v579 + int32(1)
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v563)+4))
	if v683 < v684 {
		v579 = v683
		v581 = v663
		v582 = v681
		v585 = v665
		v587 = v670
		v590 = v666
		v593 = v667
		v594 = v668
		goto L118
	} else {
		goto L141
	}
L138:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v77)+40))
	if v675-int32(1) <= v548 {
		v681 = v582
		goto L137
	} else {
		goto L139
	}
L139:
	;
	v679 = F_lappend(m, v582, v602)
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L24
	} else {
		goto L140
	}
L140:
	;
	v681 = v679
	goto L137
L141:
	;
	goto L119
L142:
	;
	goto L9
L143:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v726)+128)) = v927
	*(*float64)(unsafe.Add(mBase, uint32(v726)+112)) = v931
	F_genericcostestimate(m, v718, v719, v720, v726+int32(72))
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L24
	} else {
		goto L172
	}
L144:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v728)+40))
	v753 = int32(1)
	if (v729^int32(-1)|base.B2i32(v731 != v752-v753)|v733|v741)&v753 == int32(0) {
		v927 = v742
		v931 = float64(1)
		goto L143
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v728)+88))
	if v763 != 0 {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	goto L146
L148:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v763)+4))
	if v764 <= int32(0) {
		goto L152
	} else {
		goto L153
	}
L149:
	;
	v880 = v735
	goto L150
L150:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v728)+12))
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v881)+68))
	v883 = int32(0)
	v885 = F_clauselist_selectivity(m, v718, v880, v882, v883, v883)
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L24
	} else {
		goto L165
	}
L151:
	;
	v850 = F_list_concat(m, v834, v735)
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L24
	} else {
		goto L164
	}
L152:
	;
	v834 = int32(0)
	goto L151
L153:
	;
	goto L154
L154:
	;
	v768 = int32(0)
	v779 = v768
	v782 = v768
	goto L155
L155:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v763)+12))
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v798+v779<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v726)+12)) = v802
	*(*int32)(unsafe.Add(mBase, uint32(v726)+140)) = v802
	v808 = F_list_make1_impl(m, int32(1), v726+int32(12))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L24
	} else {
		goto L157
	}
L156:
	;
	v834 = v817
	goto L151
L157:
	;
	v811 = F_predicate_implied_by(m, v808, v735, int32(0))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L24
	} else {
		goto L158
	}
L158:
	;
	if v811 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v815 = F_list_concat(m, v782, v808)
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L24
	} else {
		goto L162
	}
L160:
	;
	v817 = v782
	goto L161
L161:
	;
	v819 = v779 + int32(1)
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v763)+4))
	if v819 < v820 {
		v779 = v819
		v782 = v817
		goto L155
	} else {
		goto L163
	}
L162:
	;
	v817 = v815
	goto L161
L163:
	;
	goto L156
L164:
	;
	v880 = v850
	goto L150
L165:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v728)+12))
	v888 = *(*float64)(unsafe.Add(mBase, uint32(v887)+120))
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v728)+16))
	v894 = base.F64_ceil(base.F64_mul(base.F64_convert_i32_u(v890), float64(0.3333333)))
	if base.F64_lt(v742, v894) != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v896 = v742
	goto L168
L167:
	;
	v896 = v894
	goto L168
L168:
	;
	v897 = float64(1)
	if base.F64_gt(v896, v897) != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v900 = v896
	goto L171
L170:
	;
	v900 = v897
	goto L171
L171:
	;
	v927 = v900
	v931 = base.F64_nearest(base.F64_div(base.F64_mul(v885, v888), v900))
	goto L143
L172:
	;
	v939 = *(*float64)(unsafe.Add(mBase, _c_F_btcostestimate[0]))
	v940 = *(*float64)(unsafe.Add(mBase, uint32(v726)+128))
	v941 = *(*float64)(unsafe.Add(mBase, uint32(v726)+72))
	v942 = *(*float64)(unsafe.Add(mBase, uint32(v728)+24))
	if base.F64_gt(v942, float64(1)) == int32(0) {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v728)+32))
	if v739 != 0 {
		v1014 = v744
		goto L177
	} else {
		goto L178
	}
L174:
	;
	v947 = *(*float64)(unsafe.Add(mBase, uint32(v726)+80))
	v957 = v941
	v959 = v947
	goto L173
L175:
	;
	goto L176
L176:
	;
	v948 = F_log(m, v942)
	mBase = m.M
	v952 = base.F64_mul(base.F64_ceil(base.F64_div(v948, float64(0.6931471805599453))), v939)
	v955 = *(*float64)(unsafe.Add(mBase, uint32(v726)+80))
	v957 = base.F64_add(v941, v952)
	v959 = base.F64_add(base.F64_mul(v940, v952), v955)
	goto L173
L177:
	;
	v1020 = base.F64_mul(v939, base.F64_mul(base.F64_convert_i32_s(v960+int32(1)), float64(50)))
	*(*float64)(unsafe.Add(mBase, uint32(v721))) = base.F64_add(v957, v1020)
	*(*float64)(unsafe.Add(mBase, uint32(v722))) = base.F64_add(base.F64_mul(v940, v1020), v959)
	v1026 = *(*float64)(unsafe.Add(mBase, uint32(v726)+88))
	*(*float64)(unsafe.Add(mBase, uint32(v723))) = v1026
	*(*float64)(unsafe.Add(mBase, uint32(v724))) = v1014
	v1029 = *(*float64)(unsafe.Add(mBase, uint32(v726)+104))
	*(*float64)(unsafe.Add(mBase, uint32(v725))) = v1029
	m.G0 = v726 + int32(176)
	return
L178:
	;
	F_examine_indexcol_variable(m, v718, v728, int32(0), v726+int32(40))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L24
	} else {
		goto L179
	}
L179:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v726)+48))
	if v966 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v969 = *(*float64)(unsafe.Add(mBase, uint32(v726)+96))
	v1014 = v969
	goto L177
L181:
	;
	goto L182
L182:
	;
	v970 = float64(0)
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v728)+52))
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v971)))
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v728)+56))
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v973)))
	v976 = F_get_opfamily_member(m, v972, v974, v974, int32(1))
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L24
	} else {
		goto L184
	}
L183:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v726)+48))
	if v1006 == int32(0) {
		v1014 = v1005
		goto L177
	} else {
		goto L195
	}
L184:
	;
	if v976 == int32(0) {
		v1005 = v970
		goto L183
	} else {
		goto L185
	}
L185:
	;
	v981 = v726 + int32(140)
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v726)+48))
	v985 = F_get_attstatsslot(m, v981, v982, int32(3), v976, int32(2))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L24
	} else {
		goto L186
	}
L186:
	;
	if v985 == int32(0) {
		v1005 = v970
		goto L183
	} else {
		goto L187
	}
L187:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v726)+160))
	v990 = *(*float32)(unsafe.Add(mBase, uint32(v989)))
	v991 = base.F64_promote_f32(v990)
	v993 = *(*int32)(unsafe.Add(mBase, uint32(v728)+64))
	v994 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v993))))
	if v994 != 0 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v995 = base.F64_neg(v991)
	goto L190
L189:
	;
	v995 = v991
	goto L190
L190:
	;
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v728)+40))
	if int32(1) < v998 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v1001 = base.F64_mul(v995, float64(0.75))
	goto L193
L192:
	;
	v1001 = v995
	goto L193
L193:
	;
	F_free_attstatsslot(m, v981)
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L24
	} else {
		goto L194
	}
L194:
	;
	v1005 = v1001
	goto L183
L195:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v726)+52))
	m.T0[v1009].(func(*base.Module, int32))(m, v1006)
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L24
	} else {
		goto L196
	}
L196:
	;
	v1014 = v1005
	goto L177
}
func F_btfloat4sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(1295)
	return int32(0)
}
func F_btfloat8sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(1296)
	return int32(0)
}
func F_btinsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int64
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v372 int32
	_ = v372
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v505 int32
	_ = v505
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
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
	var v544 int32
	_ = v544
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v579 int32
	_ = v579
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v642 int32
	_ = v642
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v751 int32
	_ = v751
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v767 int32
	_ = v767
	var v776 int32
	_ = v776
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v831 int32
	_ = v831
	var v837 int32
	_ = v837
	var v842 int32
	_ = v842
	var v863 int32
	_ = v863
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v885 int32
	_ = v885
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v953 int32
	_ = v953
	var v976 int32
	_ = v976
	var v982 int32
	_ = v982
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1012 int32
	_ = v1012
	var v1018 int32
	_ = v1018
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1036 int32
	_ = v1036
	var v1042 int32
	_ = v1042
	var v1044 int32
	_ = v1044
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1079 int32
	_ = v1079
	var v1087 int32
	_ = v1087
	var v1092 int32
	_ = v1092
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1108 int32
	_ = v1108
	var v1115 int32
	_ = v1115
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1162 int32
	_ = v1162
	var v1170 int32
	_ = v1170
	var v1177 int32
	_ = v1177
	var v1180 int32
	_ = v1180
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1194 int32
	_ = v1194
	var v1200 int32
	_ = v1200
	var v1202 int32
	_ = v1202
	var v1208 int32
	_ = v1208
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1253 int32
	_ = v1253
	var v1259 int32
	_ = v1259
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1300 int32
	_ = v1300
	var v1304 int32
	_ = v1304
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1317 int32
	_ = v1317
	var v1351 int32
	_ = v1351
	var v1353 int32
	_ = v1353
	var v1358 int32
	_ = v1358
	v9 = int32(0)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v34 = F_index_form_tuple(m, v33, l1, l2)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+4)) = uint16(v38)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v40
	v42 = m.G0
	v44 = v42 - int32(352)
	m.G0 = v44
	v46 = F__bt_mkscankey(m, l0, v34)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v48 = int32(0)
	if l5 == v48 {
		v57 = v9
		v58 = v48
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+68)) = v34
	v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34)+6)))
	v61 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+84)) = uint8(v61)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v44)+92)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v44)+80)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v44)+72)) = (v60&int32(_a_F_btinsert_0) + int32(7)) & int32(_a_F_btinsert_1)
	goto L8
L5:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+2)))
	if v52 != 0 {
		v57 = v9
		v58 = int32(1)
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v53 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = v53
	v57 = int32(1)
	v58 = v53
	goto L4
L7:
	;
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	if l5 != int32(3) {
		goto L209
	} else {
		goto L210
	}
L8:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v111 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v842 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v842 != int32(1) {
		v863 = v837
		goto L7
	} else {
		goto L205
	}
L10:
	;
	if v57 == int32(0) {
		v863 = v58
		goto L7
	} else {
		goto L47
	}
L11:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v44)+76))
	v224 = F__bt_search(m, l0, l4, v222, v44+int32(80), int32(2))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L46
	}
L12:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	if v114 == int32(-1) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v117 = F_ReadBuffer(m, l0, v114)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+80)) = v117
	v120 = F_ConditionalLockBuffer(m, v117)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	if v120 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v190 != 0 {
		goto L38
	} else {
		goto L39
	}
L17:
	;
	F__bt_checkpage(m, l0, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	F_ReleaseBuffer(m, v122)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L37
	}
L20:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	if v125 < int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	F__bt_relbuf(m, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L36
	}
L22:
	;
	v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143)+16)))
	v145 = v144 + v143
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
	if v146 != 0 {
		goto L21
	} else {
		goto L26
	}
L23:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[0]))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v129+(v125^int32(-1))<<(uint(int32(2))%32))))
	v143 = v135
	goto L22
L24:
	;
	goto L25
L25:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[1]))
	v143 = v137 + v125<<(uint(int32(13))%32) + int32(-8192)
	goto L22
L26:
	;
	v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145)+12)))
	if v147&int32(21) != int32(1) {
		goto L21
	} else {
		goto L27
	}
L27:
	;
	v152 = int32(4)
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143)+14)))
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143)+12)))
	v155 = v153 - v154
	if v155 <= v152 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v44)+72))
	if base.Ui32(v158-int32(4)) <= base.Ui32(v161) {
		goto L21
	} else {
		goto L32
	}
L29:
	;
	v158 = v152
	goto L31
L30:
	;
	v158 = v155
	goto L31
L31:
	;
	goto L28
L32:
	;
	v163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143)+12)))
	if base.B2i32(base.Ui32(v163) < base.Ui32(int32(25)))|base.B2i32((v163+int32(_a_F_btinsert_2))&int32(_a_F_btinsert_3) == int32(0)) != 0 {
		goto L21
	} else {
		goto L33
	}
L33:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v44)+76))
	v176 = F__bt_compare(m, l0, v174, v143, int32(1))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	if int32(0) < v176 {
		v228 = int32(0)
		goto L10
	} else {
		goto L35
	}
L35:
	;
	goto L21
L36:
	;
	goto L16
L37:
	;
	goto L16
L38:
	;
	v216 = v190
	goto L40
L39:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+64)) = v192
	v194 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v44)+56)) = v194
	v198 = F_smgropen(m, v44+int32(56), v191)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+16)) = int32(-1)
	goto L11
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v198
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v198)+72))
	if v202 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v216 = v214
	goto L40
L43:
	;
	v210 = v202
	goto L45
L44:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v198)+76))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v198)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+4)) = v204
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v198)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v204))) = v206
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v198)+72))
	v210 = v208
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v198)+72)) = v210 + int32(1)
	goto L42
L46:
	;
	v228 = v224
	goto L10
L47:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v44)+76))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v44)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+280)) = int32(4)
	v235 = int32(0)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	if v236 < v235 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v255 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v254)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v255) {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[0]))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v240+(v236^int32(-1))<<(uint(int32(2))%32))))
	v254 = v246
	goto L48
L50:
	;
	goto L51
L51:
	;
	v248 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[1]))
	v254 = v248 + v236<<(uint(int32(13))%32) + int32(-8192)
	goto L48
L52:
	;
	v263 = int32(base.Ui32(v255+int32(_a_F_btinsert_2)) >> (uint(int32(2)) % 32))
	goto L54
L53:
	;
	v263 = int32(0)
	goto L54
L54:
	;
	v264 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v254)+16)))
	v269 = F__bt_binsrch_insert(m, l0, v44+int32(68))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v271 = int32(0)
	v277 = v271
	v283 = v271
	v284 = v271
	v286 = v254
	v289 = v254 + v264
	v293 = int32(1)
	v294 = v269
	v299 = v235
	v301 = v263
	goto L62
L56:
	;
	goto L9
L57:
	;
	if v228 == int32(0) {
		goto L8
	} else {
		goto L203
	}
L58:
	;
	F_XactLockTableWait(m, v668, l0, v34, int32(5))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L1
	} else {
		goto L202
	}
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L1
	} else {
		goto L196
	}
L60:
	;
	if base.B2i32(v776 == int32(0))&base.B2i32(l5 == int32(3)) != 0 {
		goto L59
	} else {
		goto L193
	}
L61:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v44)+284))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v44)+288))
	if v666 != 0 {
		goto L157
	} else {
		goto L158
	}
L62:
	;
	v311 = v277
	v317 = v283
	v327 = v293
	v328 = v294
	v333 = v299
	goto L65
L63:
	;
	if l5 != int32(2) {
		goto L61
	} else {
		goto L152
	}
L64:
	;
	goto L63
L65:
	;
	v341 = v328 & int32(_a_F_btinsert_4)
	v344 = v286 + int32(20) + v341<<(uint(int32(2))%32)
	v345 = int32(0)
	v347 = v345
	v350 = v311
	v356 = v317
	v363 = v345
	v366 = v327
	v372 = v333
	goto L67
L66:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v289)+4))
	if v562 == int32(0) {
		v776 = v558
		goto L60
	} else {
		goto L130
	}
L67:
	;
	v380 = v301 & int32(_a_F_btinsert_4)
	if base.Ui32(v341) <= base.Ui32(v380) {
		goto L75
	} else {
		goto L76
	}
L68:
	;
	if base.Ui32(v341) < base.Ui32(v380) {
		goto L127
	} else {
		goto L128
	}
L69:
	;
	goto L68
L70:
	;
	v541 = int32(1)
	v544 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v536)+4)))
	if v534 < v544&int32(4095)-v541 {
		v347 = v534 + v541
		v350 = v535
		v356 = v536
		v363 = v541
		v366 = v537
		v372 = v540
		goto L67
	} else {
		goto L126
	}
L71:
	;
	v443 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v441)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+276)) = uint16(v443)
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v441)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+272)) = v445
	if base.B2i32(l5 != int32(3)) == int32(0) {
		goto L100
	} else {
		goto L101
	}
L72:
	;
	v439 = v347
	v440 = v366
	v441 = v438
	v442 = v437
	goto L71
L73:
	;
	v437 = v363
	v438 = v404
	goto L72
L74:
	;
	if v372|base.B2i32(l5 != int32(3)) != 0 {
		v837 = int32(1)
		goto L56
	} else {
		goto L97
	}
L75:
	;
	if v284 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	goto L77
L77:
	;
	if v363 == int32(0) {
		v552 = v350
		v553 = v356
		v555 = v366
		v558 = v372
		goto L69
	} else {
		goto L96
	}
L78:
	;
	v384 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+88)))
	if v341 == v384 {
		goto L74
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	if v363 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	goto L80
L82:
	;
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404)+7)))
	if v405&int32(32) == int32(0) {
		goto L73
	} else {
		goto L91
	}
L83:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	v389 = int32(_a_F_btinsert_5)
	if v388&v389 == v389 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L85
L85:
	;
	v401 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+271)) = uint8(v401)
	v403 = v350
	v404 = v356
	goto L82
L86:
	;
	v552 = v344
	v553 = v356
	v555 = v366
	v558 = v372
	goto L69
L87:
	;
	goto L88
L88:
	;
	v393 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+271)) = uint8(v393)
	v395 = F__bt_compare(m, l0, v231, v286, v341)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	if v395 != 0 {
		v776 = v372
		goto L60
	} else {
		goto L90
	}
L90:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	v403 = v344
	v404 = v286 + v397&int32(_a_F_btinsert_6)
	goto L82
L91:
	;
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404)+5)))
	if v410&int32(32) == int32(0) {
		goto L73
	} else {
		goto L92
	}
L92:
	;
	v415 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v404)+2)))
	v416 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v404))))
	v420 = v415 + (v404 + v416<<(uint(int32(16))%32))
	v421 = int32(1)
	v422 = int32(0)
	if v363 == v422 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v439 = v422
	v440 = int32(1)
	v441 = v420
	v442 = v421
	goto L71
L94:
	;
	goto L95
L95:
	;
	v437 = v421
	v438 = v420 + v347*int32(6)
	goto L72
L96:
	;
	v534 = v347
	v535 = v350
	v536 = v356
	v537 = v366
	v540 = v372
	goto L70
L97:
	;
	goto L59
L98:
	;
	if v442&int32(1) == int32(0) {
		v552 = v403
		v553 = v404
		v555 = v440
		v558 = v524
		goto L69
	} else {
		goto L125
	}
L99:
	;
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+271)))
	if v496 != int32(1) {
		v524 = v372
		goto L98
	} else {
		goto L115
	}
L100:
	;
	v450 = v44 + int32(272)
	v454 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v450)+2)))
	v455 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v450))))
	v456 = int32(16)
	v458 = v454 | v455<<(uint(v456)%32)
	v459 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v232)+2)))
	v460 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v232))))
	v463 = v459 | v460<<(uint(v456)%32)
	if base.Ui32(v458) < base.Ui32(v463) {
		v474 = int32(-1)
		goto L104
	} else {
		goto L105
	}
L101:
	;
	goto L102
L102:
	;
	v494 = F_table_index_fetch_tuple_check(m, l4, v44+int32(272), v44+int32(280), v44+int32(271))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L113
	}
L103:
	;
	if v474 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L104:
	;
	goto L103
L105:
	;
	if base.Ui32(v463) < base.Ui32(v458) {
		v474 = int32(1)
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v468 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v450)+4)))
	v469 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v232)+4)))
	if base.Ui32(v468) < base.Ui32(v469) {
		v474 = int32(-1)
		goto L104
	} else {
		goto L107
	}
L107:
	;
	v474 = base.B2i32(base.Ui32(v469) < base.Ui32(v468))
	goto L104
L108:
	;
	v524 = int32(1)
	goto L98
L109:
	;
	goto L110
L110:
	;
	v484 = F_table_index_fetch_tuple_check(m, l4, v44+int32(272), v44+int32(280), v44+int32(271))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	if v484 == int32(0) {
		goto L99
	} else {
		goto L112
	}
L112:
	;
	goto L61
L113:
	;
	if v494 != 0 {
		goto L64
	} else {
		goto L114
	}
L114:
	;
	goto L99
L115:
	;
	if v442&int32(1) != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	if v440&int32(1) == int32(0) {
		v534 = v439
		v535 = v403
		v536 = v404
		v537 = v440
		v540 = v372
		goto L70
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v403)))
	*(*int32)(unsafe.Add(mBase, uint32(v403))) = v511 | int32(_a_F_btinsert_5)
	v515 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v289)+12)))
	v517 = v515 | int32(64)
	*(*uint16)(unsafe.Add(mBase, uint32(v289)+12)) = uint16(v517)
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	if v284 != 0 {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	v505 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v404)+4)))
	if v439 != v505&int32(4095)-int32(1) {
		v534 = v439
		v535 = v403
		v536 = v404
		v537 = v440
		v540 = v372
		goto L70
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	v520 = v284
	goto L123
L122:
	;
	v520 = v519
	goto L123
L123:
	;
	F_MarkBufferDirtyHint(m, v520, int32(1))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	v524 = v372
	goto L98
L125:
	;
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+271)))
	v534 = v439
	v535 = v403
	v536 = v404
	v537 = (v529 | (v442 ^ int32(1))) & v440
	v540 = v524
	goto L70
L126:
	;
	v552 = v535
	v553 = v536
	v555 = v537
	v558 = v540
	goto L69
L127:
	;
	v311 = v552
	v317 = v553
	v327 = v555
	v328 = v328 + int32(1)
	v333 = v558
	goto L65
L128:
	;
	goto L129
L129:
	;
	goto L66
L130:
	;
	v566 = F__bt_compare(m, l0, v231, v286, int32(1))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	if v566 != 0 {
		v776 = v558
		goto L60
	} else {
		goto L132
	}
L132:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v289)+4))
	v570 = v568
	v579 = v284
	goto L134
L133:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v623)+4))
	if v650 != 0 {
		goto L146
	} else {
		goto L147
	}
L134:
	;
	v602 = F__bt_relandgetbuf(m, l0, v579, v570, int32(1))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L1
	} else {
		goto L137
	}
L135:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L1
	} else {
		goto L143
	}
L136:
	;
	v622 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v621)+16)))
	v623 = v622 + v621
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v623)+12)))
	if v624&int32(20) == int32(0) {
		goto L133
	} else {
		goto L141
	}
L137:
	;
	if v602 < int32(0) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v607 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[0]))
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v607+(v602^int32(-1))<<(uint(int32(2))%32))))
	v621 = v613
	goto L136
L139:
	;
	goto L140
L140:
	;
	v615 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[1]))
	v621 = v615 + v602<<(uint(int32(13))%32) + int32(-8192)
	goto L136
L141:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v623)+4))
	if v629 != 0 {
		v570 = v629
		v579 = v602
		goto L134
	} else {
		goto L142
	}
L142:
	;
	goto L135
L143:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = v634 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btinsert_7), v44+int32(16))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_btinsert_8), int32(743), int32(_a_F_btinsert_9))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L146:
	;
	v651 = int32(2)
	goto L148
L147:
	;
	v651 = int32(1)
	goto L148
L148:
	;
	v652 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v621)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v652) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v660 = int32(base.Ui32(v652+int32(_a_F_btinsert_2)) >> (uint(int32(2)) % 32))
	goto L151
L150:
	;
	v660 = int32(0)
	goto L151
L151:
	;
	v277 = v552
	v283 = v553
	v284 = v602
	v286 = v621
	v289 = v623
	v293 = v555
	v294 = v651
	v299 = v558
	v301 = v660
	goto L62
L152:
	;
	if v284 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	F__bt_relbuf(m, v284)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L1
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v837 = int32(0)
	goto L56
L156:
	;
	goto L155
L157:
	;
	v668 = v666
	goto L159
L158:
	;
	v668 = v667
	goto L159
L159:
	;
	if v668 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	if v284 != 0 {
		goto L163
	} else {
		goto L164
	}
L161:
	;
	goto L162
L162:
	;
	v683 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v232)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+276)) = uint16(v683)
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v232)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+272)) = v685
	v691 = F_table_index_fetch_tuple_check(m, l4, v44+int32(272), int32(_a_F_btinsert_10), int32(0))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L1
	} else {
		goto L170
	}
L163:
	;
	F__bt_relbuf(m, v284)
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v671 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+84)) = uint8(v671)
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v44)+316))
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	F__bt_relbuf(m, v674)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L1
	} else {
		goto L167
	}
L166:
	;
	goto L165
L167:
	;
	v677 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+80)) = v677
	if v673 == v677 {
		goto L58
	} else {
		goto L168
	}
L168:
	;
	F_SpeculativeInsertionWait(m, v668, v673)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	goto L57
L170:
	;
	if v691 == int32(0) {
		v776 = v372
		goto L60
	} else {
		goto L171
	}
L171:
	;
	v695 = int32(0)
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	if v696 < v695 {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	F_CheckForSerializableConflictIn(m, l0, v695, v715)
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L1
	} else {
		goto L176
	}
L173:
	;
	v700 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[2]))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v700+(v696^int32(-1))<<(uint(int32(6))%32))+16))
	v715 = v706
	goto L172
L174:
	;
	goto L175
L175:
	;
	v708 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[3]))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v708+v696<<(uint(int32(6))%32)+int32(-64))+16))
	v715 = v714
	goto L172
L176:
	;
	if v284 != 0 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	F__bt_relbuf(m, v284)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L1
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	F__bt_relbuf(m, v720)
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L1
	} else {
		goto L181
	}
L180:
	;
	goto L179
L181:
	;
	v723 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+84)) = uint8(v723)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+80)) = v723
	v727 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v729 = v44 + int32(128)
	v731 = v44 + int32(96)
	F_index_deform_tuple(m, v232, v727, v729, v731)
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	v734 = F_BuildIndexValueDescription(m, l0, v729, v731)
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	F_errcode(m, int32(83906754))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+48)) = v743 + int32(4)
	F_errmsg(m, int32(_a_F_btinsert_11), v44+int32(48))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	if v734 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+32)) = v734
	F_errdetail(m, int32(_a_F_btinsert_12), v44+int32(32))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L1
	} else {
		goto L190
	}
L188:
	;
	goto L189
L189:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_errtableconstraint(m, l4, v758+int32(4))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L1
	} else {
		goto L191
	}
L190:
	;
	goto L189
L191:
	;
	F_errfinish(m, int32(_a_F_btinsert_8), int32(673), int32(_a_F_btinsert_9))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
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
	v782 = int32(1)
	if v284 == int32(0) {
		v837 = v782
		goto L56
	} else {
		goto L194
	}
L194:
	;
	F__bt_relbuf(m, v284)
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	v837 = v782
	goto L56
L196:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v804 + int32(4)
	F_errmsg(m, int32(_a_F_btinsert_13), v44)
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	F_errhint(m, int32(_a_F_btinsert_14), int32(0))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_errtableconstraint(m, l4, v815+int32(4))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(_a_F_btinsert_8), int32(766), int32(_a_F_btinsert_9))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L202:
	;
	goto L57
L203:
	;
	F__bt_freestack(m, v228)
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	goto L8
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = v34
	v863 = v837
	goto L7
L206:
	;
	if v228 != 0 {
		goto L314
	} else {
		goto L315
	}
L207:
	;
	v1293 = v44 + int32(68)
	v1294 = F__bt_binsrch_insert(m, l0, v1293)
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L1
	} else {
		goto L307
	}
L208:
	;
	if v57 == int32(0) {
		goto L266
	} else {
		goto L267
	}
L209:
	;
	v881 = int32(0)
	if v878 < v881 {
		goto L213
	} else {
		goto L214
	}
L210:
	;
	goto L211
L211:
	;
	F__bt_relbuf(m, v878)
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L1
	} else {
		goto L264
	}
L212:
	;
	F_CheckForSerializableConflictIn(m, l0, v881, v900)
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L1
	} else {
		goto L216
	}
L213:
	;
	v885 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[2]))
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v885+(v878^int32(-1))<<(uint(int32(6))%32))+16))
	v900 = v891
	goto L212
L214:
	;
	goto L215
L215:
	;
	v893 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[3]))
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v893+v878<<(uint(int32(6))%32)+int32(-64))+16))
	v900 = v899
	goto L212
L216:
	;
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v44)+76))
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	if v904 < int32(0) {
		goto L218
	} else {
		goto L219
	}
L217:
	;
	v923 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v922)+16)))
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v44)+72))
	if base.Ui32(int32(2705)) <= base.Ui32(v924) {
		goto L221
	} else {
		goto L222
	}
L218:
	;
	v908 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[0]))
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v908+(v904^int32(-1))<<(uint(int32(2))%32))))
	v922 = v914
	goto L217
L219:
	;
	goto L220
L220:
	;
	v916 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[1]))
	v922 = v916 + v904<<(uint(int32(13))%32) + int32(-8192)
	goto L217
L221:
	;
	v927 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v903))))
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v44)+68))
	F__bt_check_third_page(m, l0, l4, v927, v922, v928)
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L1
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	v931 = v923 + v922
	v932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v903))))
	if v932 != 0 {
		goto L208
	} else {
		goto L225
	}
L224:
	;
	goto L223
L225:
	;
	v933 = int32(4)
	v934 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v922)+14)))
	v935 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v922)+12)))
	v936 = v934 - v935
	if v936 <= v933 {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v44)+72))
	if base.Ui32(v942) <= base.Ui32(v939-int32(4)) {
		goto L207
	} else {
		goto L230
	}
L227:
	;
	v939 = v933
	goto L229
L228:
	;
	v939 = v936
	goto L229
L229:
	;
	goto L226
L230:
	;
	v945 = v922
	v953 = v931
	goto L231
L231:
	;
	v976 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v953)+12)))
	if v976&int32(64) != 0 {
		goto L233
	} else {
		goto L234
	}
L232:
	;
	goto L207
L233:
	;
	v982 = int32(0)
	F__bt_delete_or_dedup_one_page(m, l0, l4, v44+int32(68), int32(1), v982, v982, v982)
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L1
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	v998 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+84)))
	if v998 != int32(1) {
		goto L242
	} else {
		goto L243
	}
L236:
	;
	v987 = int32(4)
	v988 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v945)+14)))
	v989 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v945)+12)))
	v990 = v988 - v989
	if v990 <= v987 {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v44)+72))
	if base.Ui32(v996) <= base.Ui32(v993-int32(4)) {
		goto L207
	} else {
		goto L241
	}
L238:
	;
	v993 = v987
	goto L240
L239:
	;
	v993 = v990
	goto L240
L240:
	;
	goto L237
L241:
	;
	goto L235
L242:
	;
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v953)+4))
	if v1018 == int32(0) {
		goto L207
	} else {
		goto L249
	}
L243:
	;
	v1001 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+88)))
	v1002 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+86)))
	if base.Ui32(v1001) < base.Ui32(v1002) {
		goto L242
	} else {
		goto L244
	}
L244:
	;
	v1004 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v945)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1004) {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v1012 = int32(base.Ui32(v1004+int32(_a_F_btinsert_2)) >> (uint(int32(2)) % 32))
	goto L247
L246:
	;
	v1012 = int32(0)
	goto L247
L247:
	;
	if base.Ui32(v1001) <= base.Ui32(v1012&int32(_a_F_btinsert_4)) {
		goto L207
	} else {
		goto L248
	}
L248:
	;
	goto L242
L249:
	;
	v1022 = F__bt_compare(m, l0, v903, v945, int32(1))
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L1
	} else {
		goto L250
	}
L250:
	;
	if v1022 != 0 {
		goto L207
	} else {
		goto L251
	}
L251:
	;
	v1025 = Fn13986(m, int64(32))
	mBase = m.M
	goto L252
L252:
	;
	if base.Ui32(v1025) < base.Ui32(int32(42949673)) {
		goto L207
	} else {
		goto L253
	}
L253:
	;
	F__bt_stepright(m, l0, l4, v44+int32(68), v228)
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L1
	} else {
		goto L254
	}
L254:
	;
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	if v1032 < int32(0) {
		goto L256
	} else {
		goto L257
	}
L255:
	;
	v1051 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1050)+16)))
	v1053 = int32(4)
	v1054 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1050)+14)))
	v1055 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1050)+12)))
	v1056 = v1054 - v1055
	if v1056 <= v1053 {
		goto L260
	} else {
		goto L261
	}
L256:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[0]))
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1036+(v1032^int32(-1))<<(uint(int32(2))%32))))
	v1050 = v1042
	goto L255
L257:
	;
	goto L258
L258:
	;
	v1044 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[1]))
	v1050 = v1044 + v1032<<(uint(int32(13))%32) + int32(-8192)
	goto L255
L259:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v44)+72))
	if base.Ui32(v1059-int32(4)) < base.Ui32(v1062) {
		v945 = v1050
		v953 = v1051 + v1050
		goto L231
	} else {
		goto L263
	}
L260:
	;
	v1059 = v1053
	goto L262
L261:
	;
	v1059 = v1056
	goto L262
L262:
	;
	goto L259
L263:
	;
	goto L232
L264:
	;
	goto L206
L265:
	;
	v1244 = int32(4)
	v1245 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1213)+14)))
	v1246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1213)+12)))
	v1247 = v1245 - v1246
	if v1247 <= v1244 {
		goto L302
	} else {
		goto L303
	}
L266:
	;
	v1212 = l6
	v1213 = v922
	goto L265
L267:
	;
	goto L268
L268:
	;
	v1068 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+86)))
	v1069 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+88)))
	v1071 = l6 | base.B2i32(base.Ui32(v1068) < base.Ui32(v1069))
	v1072 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+84)))
	if base.B2i32(v1072 != int32(1))|base.B2i32(base.Ui32(v1069) < base.Ui32(v1068)) == int32(0) {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v1079 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v922)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1079) {
		goto L272
	} else {
		goto L273
	}
L270:
	;
	goto L271
L271:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v931)+4))
	if v1092 == int32(0) {
		v1212 = v1071
		v1213 = v922
		goto L265
	} else {
		goto L276
	}
L272:
	;
	v1087 = int32(base.Ui32(v1079+int32(_a_F_btinsert_2)) >> (uint(int32(2)) % 32))
	goto L274
L273:
	;
	v1087 = int32(0)
	goto L274
L274:
	;
	if base.Ui32(v1069) <= base.Ui32(v1087&int32(_a_F_btinsert_4)) {
		v1212 = v1071
		v1213 = v922
		goto L265
	} else {
		goto L275
	}
L275:
	;
	goto L271
L276:
	;
	v1096 = F__bt_compare(m, l0, v903, v922, int32(1))
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L1
	} else {
		goto L277
	}
L277:
	;
	if v1096 <= int32(0) {
		v1212 = v1071
		v1213 = v922
		goto L265
	} else {
		goto L278
	}
L278:
	;
	F__bt_stepright(m, l0, l4, v44+int32(68), v228)
	mBase = m.M
	v1103 = m.ExcPending
	if v1103 != 0 {
		goto L1
	} else {
		goto L279
	}
L279:
	;
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	if int32(0) <= v1104 {
		goto L281
	} else {
		goto L282
	}
L280:
	;
	v1124 = v1122
	goto L284
L281:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[1]))
	v1122 = v1108 + v1104<<(uint(int32(13))%32) + int32(-8192)
	goto L280
L282:
	;
	goto L283
L283:
	;
	v1115 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[0]))
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v1115+(v1104^int32(-1))<<(uint(int32(2))%32))))
	v1122 = v1121
	goto L280
L284:
	;
	v1155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1124)+16)))
	v1156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+84)))
	if v1156 != int32(1) {
		goto L287
	} else {
		goto L288
	}
L285:
	;
	v1212 = int32(1)
	v1213 = v1124
	goto L265
L286:
	;
	goto L285
L287:
	;
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(v1155+v1124)+4))
	if v1177 == int32(0) {
		goto L286
	} else {
		goto L294
	}
L288:
	;
	v1159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+88)))
	v1160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+86)))
	if base.Ui32(v1159) < base.Ui32(v1160) {
		goto L287
	} else {
		goto L289
	}
L289:
	;
	v1162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1124)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1162) {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v1170 = int32(base.Ui32(v1162+int32(_a_F_btinsert_2)) >> (uint(int32(2)) % 32))
	goto L292
L291:
	;
	v1170 = int32(0)
	goto L292
L292:
	;
	if base.Ui32(v1159) <= base.Ui32(v1170&int32(_a_F_btinsert_4)) {
		goto L286
	} else {
		goto L293
	}
L293:
	;
	goto L287
L294:
	;
	v1180 = int32(1)
	v1182 = F__bt_compare(m, l0, v903, v1124, v1180)
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L1
	} else {
		goto L295
	}
L295:
	;
	if v1182 <= int32(0) {
		v1212 = v1180
		v1213 = v1124
		goto L265
	} else {
		goto L296
	}
L296:
	;
	F__bt_stepright(m, l0, l4, v44+int32(68), v228)
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	if v1190 < int32(0) {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	v1194 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[0]))
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v1194+(v1190^int32(-1))<<(uint(int32(2))%32))))
	v1208 = v1200
	goto L300
L299:
	;
	v1202 = *(*int32)(unsafe.Add(mBase, _c_F_btinsert[1]))
	v1208 = v1202 + v1190<<(uint(int32(13))%32) + int32(-8192)
	goto L300
L300:
	;
	v1124 = v1208
	goto L284
L301:
	;
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v44)+72))
	if base.Ui32(v1253) <= base.Ui32(v1250-int32(4)) {
		goto L207
	} else {
		goto L305
	}
L302:
	;
	v1250 = v1244
	goto L304
L303:
	;
	v1250 = v1247
	goto L304
L304:
	;
	goto L301
L305:
	;
	F__bt_delete_or_dedup_one_page(m, l0, l4, v44+int32(68), int32(0), v57, v1212, l6)
	mBase = m.M
	v1259 = m.ExcPending
	if v1259 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	goto L207
L307:
	;
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v44)+92))
	if v1296 == int32(-1) {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	v1300 = int32(0)
	F__bt_delete_or_dedup_one_page(m, l0, l4, v1293, int32(1), v1300, v1300, v1300)
	mBase = m.M
	v1304 = m.ExcPending
	if v1304 != 0 {
		goto L1
	} else {
		goto L311
	}
L309:
	;
	v1310 = v1294
	v1311 = v1296
	goto L310
L310:
	;
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	v1313 = int32(0)
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(v44)+72))
	F__bt_insertonpg(m, l0, l4, v46, v1312, v1313, v228, v34, v1314, v1310, v1311, v1313)
	mBase = m.M
	v1317 = m.ExcPending
	if v1317 != 0 {
		goto L1
	} else {
		goto L313
	}
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+92)) = int32(0)
	v1307 = F__bt_binsrch_insert(m, l0, v1293)
	mBase = m.M
	v1308 = m.ExcPending
	if v1308 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v44)+92))
	v1310 = v1307
	v1311 = v1309
	goto L310
L313:
	;
	goto L206
L314:
	;
	F__bt_freestack(m, v228)
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L1
	} else {
		goto L317
	}
L315:
	;
	goto L316
L316:
	;
	F_pfree(m, v46)
	mBase = m.M
	v1353 = m.ExcPending
	if v1353 != 0 {
		goto L1
	} else {
		goto L318
	}
L317:
	;
	goto L316
L318:
	;
	m.G0 = v44 + int32(352)
	F_pfree(m, v34)
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L1
	} else {
		goto L319
	}
L319:
	;
	return v863
}
func F_btint24cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return base.B2i32(v4 < v3) - base.B2i32(v3 < v4)
}
func F_btint28cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	v4 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	return base.B2i32(v6 < v4) - base.B2i32(v4 < v6)
}
func F_btint2fastcmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return base.I32_extend16_s(l0) - base.I32_extend16_s(l1)
}
func F_btint48cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	v4 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+20)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	return base.B2i32(v6 < v4) - base.B2i32(v4 < v6)
}
func F_btint8fastcmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	return base.B2i32(v7 < v6) - base.B2i32(v6 < v7)
}
func F_btproperty(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v11 = base.B2i32(l2 == int32(7)) & base.B2i32(l1 != int32(0))
	if v11 != 0 {
		v12 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v12)
	} else {
	}
	return v11
}
func F_btrecordcmp(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_record_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_bttextcmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v15 = v10 + int32(1)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum_packed(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
			v25 = v23 & int32(1)
			if v25 != 0 {
				v26 = v15
			} else {
				v26 = v10 + int32(4)
			}
			if v23 == int32(1) {
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
				if v32 == int32(18) {
					v35 = int32(16)
				} else {
					v35 = int32(0)
				}
				if base.Ui32((v32-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v42 = int32(4)
				} else {
					v42 = v35
				}
				v53 = v42
			} else {
				v43 = int32(1)
				if v25 != 0 {
					v53 = int32(base.Ui32(v23)>>(uint(v43)%32)) - v43
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v54 = int32(1)
			v55 = v17 + v54
			if v19&v54 != 0 {
				v60 = v55
			} else {
				v60 = v17 + int32(4)
			}
			if v19 == int32(1) {
				v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
				if v66 == int32(18) {
					v69 = int32(16)
				} else {
					v69 = int32(0)
				}
				if base.Ui32((v66-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v76 = int32(4)
				} else {
					v76 = v69
				}
				v89 = v76
			} else {
				v77 = int32(1)
				if v19&v77 != 0 {
					v89 = int32(base.Ui32(v19)>>(uint(v77)%32)) - v77
				} else {
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
					v89 = int32(base.Ui32(v83)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v90 = F_varstr_cmp(m, v26, v53, v60, v89, v20)
			mBase = m.M
			v91 = m.ExcPending
			if v91 != 0 {
				return int32(0)
			} else {
				v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v92 != v10 {
					F_pfree(m, v10)
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return int32(0)
					} else {
						v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v96 != v17 {
							F_pfree(m, v17)
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								return v90
							}
						} else {
							return v90
						}
					}
				} else {
					v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v96 != v17 {
						F_pfree(m, v17)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							return v90
						}
					} else {
						return v90
					}
				}
			}
		}
	}
}
func F_bttidcmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2)+2)))
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2))))
	v9 = int32(16)
	v11 = v7 | v8<<(uint(v9)%32)
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+2)))
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3))))
	v16 = v12 | v13<<(uint(v9)%32)
	if base.Ui32(v11) < base.Ui32(v16) {
		v27 = int32(-1)
	} else {
		if base.Ui32(v16) < base.Ui32(v11) {
			v27 = int32(1)
		} else {
			v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2)+4)))
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+4)))
			if base.Ui32(v21) < base.Ui32(v22) {
				v27 = int32(-1)
			} else {
				v27 = base.B2i32(base.Ui32(v22) < base.Ui32(v21))
			}
		}
	}
	return v27
}
func F_btvarstrequalimage(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(34209924))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_btvarstrequalimage_0), int32(0))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					F_errhint(m, int32(_a_F_btvarstrequalimage_1), int32(0))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_btvarstrequalimage_2), int32(1648), int32(_a_F_btvarstrequalimage_3))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	} else {
		v27 = F_pg_newlocale_from_collation(m, v2)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)))
			return v29
		}
	}
}
func F_buildNSItemFromLists(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v17 = v13 << (uint(int32(5)) % 32)
	goto L3
L2:
	;
	v17 = int32(0)
	goto L3
L3:
	;
	v18 = F_palloc0(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v27 = int32(0)
	goto L6
L6:
	;
	v34 = int32(0)
	if l2 == v34 {
		v45 = v34
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if l3 == int32(0) {
		v54 = v34
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v39 <= v27 {
		v45 = int32(0)
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v45 = v41 + v27<<(uint(int32(2))%32)
	goto L8
L11:
	;
	if l4 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v48 <= v27 {
		v54 = v34
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v54 = v50 + v27<<(uint(int32(2))%32)
	goto L11
L14:
	;
	v82 = v18 + v27<<(uint(int32(5))%32)
	v84 = v27 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v82)+4)) = uint16(v84)
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = l1
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = v87
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+12)) = v89
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v65+v27<<(uint(int32(2))%32))))
	*(*uint16)(unsafe.Add(mBase, uint32(v82)+28)) = uint16(v84)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+24)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v82)+16)) = v94
	v27 = v84
	goto L6
L15:
	;
	v68 = F_palloc(m, int32(28))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L19
	}
L16:
	;
	v57 = int32(0)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if base.B2i32(v54 == v57)|(base.B2i32(v45 == v57)|base.B2i32(v61 <= v27)) != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	if v65 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v68)+20)) = int64(16777473)
	*(*int32)(unsafe.Add(mBase, uint32(v68)+16)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v68)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v68)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = v70
	return v68
}
func F_build_backup_content(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v30 int32
	_ = v30
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v53 int64
	_ = v53
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v62 int32
	_ = v62
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v72 int64
	_ = v72
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v85 int64
	_ = v85
	var v92 int32
	_ = v92
	var v97 int64
	_ = v97
	var v100 int64
	_ = v100
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v167 int64
	_ = v167
	var v172 int64
	_ = v172
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	v9 = m.G0
	v11 = v9 - int32(528)
	m.G0 = v11
	v13 = F_makeStringInfo(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, _c_F_build_backup_content[0]))
		v25 = F_pg_localtime(m, l0+int32(1056), v24)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v27 = F_pg_strftime(m, v11+int32(400), int32(128), int32(_a_F_build_backup_content_0), v25)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v29 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1032))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1040))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+192)) = v30
				v33 = int64(*(*int32)(unsafe.Add(mBase, _c_F_build_backup_content[1])))
				v34 = base.I64_div_u_s(v29, v33)
				v36 = base.I64_div_u_s(int64(4294967296), v33)
				v37 = base.I64_div_u_s(v34, v36)
				*(*uint32)(unsafe.Add(mBase, uint32(v11)+196)) = uint32(v37)
				v40 = v34 - v36*v37
				*(*uint32)(unsafe.Add(mBase, uint32(v11)+200)) = uint32(v40)
				v43 = v11 + int32(336)
				v48 = F_pg_snprintf(m, v43, int32(64), int32(_a_F_build_backup_content_1), v11+int32(192))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					v50 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1032))
					*(*uint32)(unsafe.Add(mBase, uint32(v11)+180)) = uint32(v50)
					v53 = int64(base.Ui64(v50) >> (uint(int64(32)) % 64))
					*(*uint32)(unsafe.Add(mBase, uint32(v11)+176)) = uint32(v53)
					*(*int32)(unsafe.Add(mBase, uint32(v11)+184)) = v43
					F_appendStringInfo(m, v13, int32(_a_F_build_backup_content_2), v11+int32(176))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						if l1 != 0 {
							v61 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1088))
							v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1096))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+160)) = v62
							v65 = int64(*(*int32)(unsafe.Add(mBase, _c_F_build_backup_content[1])))
							v66 = base.I64_div_u_s(v61, v65)
							v68 = base.I64_div_u_s(int64(4294967296), v65)
							v69 = base.I64_div_u_s(v66, v68)
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+164)) = uint32(v69)
							v72 = v66 - v68*v69
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+168)) = uint32(v72)
							v75 = v11 + int32(208)
							v80 = F_pg_snprintf(m, v75, int32(64), int32(_a_F_build_backup_content_1), v11+int32(160))
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								v82 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1088))
								*(*uint32)(unsafe.Add(mBase, uint32(v11)+148)) = uint32(v82)
								v85 = int64(base.Ui64(v82) >> (uint(int64(32)) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v11)+144)) = uint32(v85)
								*(*int32)(unsafe.Add(mBase, uint32(v11)+152)) = v75
								F_appendStringInfo(m, v13, int32(_a_F_build_backup_content_3), v11+int32(144))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									v97 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1048))
									*(*uint32)(unsafe.Add(mBase, uint32(v11)+132)) = uint32(v97)
									v100 = int64(base.Ui64(v97) >> (uint(int64(32)) % 64))
									*(*uint32)(unsafe.Add(mBase, uint32(v11)+128)) = uint32(v100)
									F_appendStringInfo(m, v13, int32(_a_F_build_backup_content_4), v11+int32(128))
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int32(0)
									} else {
										F_appendStringInfoString(m, v13, int32(_a_F_build_backup_content_5))
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return int32(0)
										} else {
											v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1064)))
											if v112 != 0 {
												v113 = int32(_a_F_build_backup_content_6)
											} else {
												v113 = int32(_a_F_build_backup_content_7)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v113
											F_appendStringInfo(m, v13, int32(_a_F_build_backup_content_8), v11+int32(112))
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v11 + int32(400)
												F_appendStringInfo(m, v13, int32(_a_F_build_backup_content_9), v11+int32(96))
												mBase = m.M
												v127 = m.ExcPending
												if v127 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = l0
													F_appendStringInfo(m, v13, int32(_a_F_build_backup_content_10), v11+int32(80))
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
														return int32(0)
													} else {
														v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1040))
														*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v134
														F_appendStringInfo(m, v13, int32(_a_F_build_backup_content_11), v11-int32(-64))
														mBase = m.M
														v140 = m.ExcPending
														if v140 != 0 {
															return int32(0)
														} else {
															if l1 != 0 {
																v142 = v11 + int32(208)
																v148 = *(*int32)(unsafe.Add(mBase, _c_F_build_backup_content[0]))
																v149 = F_pg_localtime(m, l0+int32(1104), v148)
																mBase = m.M
																v150 = m.ExcPending
																if v150 != 0 {
																	return int32(0)
																} else {
																	v151 = F_pg_strftime(m, v142, int32(128), int32(_a_F_build_backup_content_0), v149)
																	mBase = m.M
																	v152 = m.ExcPending
																	if v152 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v142
																		F_appendStringInfo(m, v13, int32(_a_F_build_backup_content_12), v11+int32(48))
																		mBase = m.M
																		v158 = m.ExcPending
																		if v158 != 0 {
																			return int32(0)
																		} else {
																			v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1096))
																			*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v159
																			F_appendStringInfo(m, v13, int32(_a_F_build_backup_content_13), v11+int32(32))
																			mBase = m.M
																			v165 = m.ExcPending
																			if v165 != 0 {
																				return int32(0)
																			} else {
																				v167 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1072))
																				if v167 != int64(0) {
																					*(*uint32)(unsafe.Add(mBase, uint32(v11)+20)) = uint32(v167)
																					v172 = int64(base.Ui64(v167) >> (uint(int64(32)) % 64))
																					*(*uint32)(unsafe.Add(mBase, uint32(v11)+16)) = uint32(v172)
																					F_appendStringInfo(m, v13, int32(_a_F_build_backup_content_14), v11+int32(16))
																					mBase = m.M
																					v178 = m.ExcPending
																					if v178 != 0 {
																						return int32(0)
																					} else {
																						v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1080))
																						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v179
																						F_appendStringInfo(m, v13, int32(_a_F_build_backup_content_15), v11)
																						mBase = m.M
																						v183 = m.ExcPending
																						if v183 != 0 {
																							return int32(0)
																						} else {
																							v184 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																							F_pfree(m, v13)
																							mBase = m.M
																							v186 = m.ExcPending
																							if v186 != 0 {
																								return int32(0)
																							} else {
																								m.G0 = v11 + int32(528)
																								return v184
																							}
																						}
																					}
																				} else {
																					v184 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																					F_pfree(m, v13)
																					mBase = m.M
																					v186 = m.ExcPending
																					if v186 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v11 + int32(528)
																						return v184
																					}
																				}
																			}
																		}
																	}
																}
															} else {
																v167 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1072))
																if v167 != int64(0) {
																	*(*uint32)(unsafe.Add(mBase, uint32(v11)+20)) = uint32(v167)
																	v172 = int64(base.Ui64(v167) >> (uint(int64(32)) % 64))
																	*(*uint32)(unsafe.Add(mBase, uint32(v11)+16)) = uint32(v172)
																	F_appendStringInfo(m, v13, int32(_a_F_build_backup_content_14), v11+int32(16))
																	mBase = m.M
																	v178 = m.ExcPending
																	if v178 != 0 {
																		return int32(0)
																	} else {
																		v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1080))
																		*(*int32)(unsafe.Add(mBase, uint32(v11))) = v179
																		F_appendStringInfo(m, v13, int32(_a_F_build_backup_content_15), v11)
																		mBase = m.M
																		v183 = m.ExcPending
																		if v183 != 0 {
																			return int32(0)
																		} else {
																			v184 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																			F_pfree(m, v13)
																			mBase = m.M
																			v186 = m.ExcPending
																			if v186 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v11 + int32(528)
																				return v184
																			}
																		}
																	}
																} else {
																	v184 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																	F_pfree(m, v13)
																	mBase = m.M
																	v186 = m.ExcPending
																	if v186 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v11 + int32(528)
																		return v184
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
							v97 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1048))
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+132)) = uint32(v97)
							v100 = int64(base.Ui64(v97) >> (uint(int64(32)) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+128)) = uint32(v100)
							F_appendStringInfo(m, v13, int32(_a_F_build_backup_content_4), v11+int32(128))
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return int32(0)
							} else {
								F_appendStringInfoString(m, v13, int32(_a_F_build_backup_content_5))
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return int32(0)
								} else {
									v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1064)))
									if v112 != 0 {
										v113 = int32(_a_F_build_backup_content_6)
									} else {
										v113 = int32(_a_F_build_backup_content_7)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v113
									F_appendStringInfo(m, v13, int32(_a_F_build_backup_content_8), v11+int32(112))
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v11 + int32(400)
										F_appendStringInfo(m, v13, int32(_a_F_build_backup_content_9), v11+int32(96))
										mBase = m.M
										v127 = m.ExcPending
										if v127 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = l0
											F_appendStringInfo(m, v13, int32(_a_F_build_backup_content_10), v11+int32(80))
											mBase = m.M
											v133 = m.ExcPending
											if v133 != 0 {
												return int32(0)
											} else {
												v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1040))
												*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v134
												F_appendStringInfo(m, v13, int32(_a_F_build_backup_content_11), v11-int32(-64))
												mBase = m.M
												v140 = m.ExcPending
												if v140 != 0 {
													return int32(0)
												} else {
													if l1 != 0 {
														v142 = v11 + int32(208)
														v148 = *(*int32)(unsafe.Add(mBase, _c_F_build_backup_content[0]))
														v149 = F_pg_localtime(m, l0+int32(1104), v148)
														mBase = m.M
														v150 = m.ExcPending
														if v150 != 0 {
															return int32(0)
														} else {
															v151 = F_pg_strftime(m, v142, int32(128), int32(_a_F_build_backup_content_0), v149)
															mBase = m.M
															v152 = m.ExcPending
															if v152 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v142
																F_appendStringInfo(m, v13, int32(_a_F_build_backup_content_12), v11+int32(48))
																mBase = m.M
																v158 = m.ExcPending
																if v158 != 0 {
																	return int32(0)
																} else {
																	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1096))
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v159
																	F_appendStringInfo(m, v13, int32(_a_F_build_backup_content_13), v11+int32(32))
																	mBase = m.M
																	v165 = m.ExcPending
																	if v165 != 0 {
																		return int32(0)
																	} else {
																		v167 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1072))
																		if v167 != int64(0) {
																			*(*uint32)(unsafe.Add(mBase, uint32(v11)+20)) = uint32(v167)
																			v172 = int64(base.Ui64(v167) >> (uint(int64(32)) % 64))
																			*(*uint32)(unsafe.Add(mBase, uint32(v11)+16)) = uint32(v172)
																			F_appendStringInfo(m, v13, int32(_a_F_build_backup_content_14), v11+int32(16))
																			mBase = m.M
																			v178 = m.ExcPending
																			if v178 != 0 {
																				return int32(0)
																			} else {
																				v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1080))
																				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v179
																				F_appendStringInfo(m, v13, int32(_a_F_build_backup_content_15), v11)
																				mBase = m.M
																				v183 = m.ExcPending
																				if v183 != 0 {
																					return int32(0)
																				} else {
																					v184 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																					F_pfree(m, v13)
																					mBase = m.M
																					v186 = m.ExcPending
																					if v186 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v11 + int32(528)
																						return v184
																					}
																				}
																			}
																		} else {
																			v184 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																			F_pfree(m, v13)
																			mBase = m.M
																			v186 = m.ExcPending
																			if v186 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v11 + int32(528)
																				return v184
																			}
																		}
																	}
																}
															}
														}
													} else {
														v167 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1072))
														if v167 != int64(0) {
															*(*uint32)(unsafe.Add(mBase, uint32(v11)+20)) = uint32(v167)
															v172 = int64(base.Ui64(v167) >> (uint(int64(32)) % 64))
															*(*uint32)(unsafe.Add(mBase, uint32(v11)+16)) = uint32(v172)
															F_appendStringInfo(m, v13, int32(_a_F_build_backup_content_14), v11+int32(16))
															mBase = m.M
															v178 = m.ExcPending
															if v178 != 0 {
																return int32(0)
															} else {
																v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1080))
																*(*int32)(unsafe.Add(mBase, uint32(v11))) = v179
																F_appendStringInfo(m, v13, int32(_a_F_build_backup_content_15), v11)
																mBase = m.M
																v183 = m.ExcPending
																if v183 != 0 {
																	return int32(0)
																} else {
																	v184 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																	F_pfree(m, v13)
																	mBase = m.M
																	v186 = m.ExcPending
																	if v186 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v11 + int32(528)
																		return v184
																	}
																}
															}
														} else {
															v184 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
															F_pfree(m, v13)
															mBase = m.M
															v186 = m.ExcPending
															if v186 != 0 {
																return int32(0)
															} else {
																m.G0 = v11 + int32(528)
																return v184
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
func F_build_bound_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l4)+32))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+10)))
	v13 = int32(*(*int16)(unsafe.Add(mBase, uint32(l4)+8)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if l3 != 0 {
		v17 = int32(4)
	} else {
		v17 = int32(5)
	}
	if l3 != 0 {
		v20 = int32(2)
	} else {
		v20 = int32(1)
	}
	if l2 != 0 {
		v21 = v17
	} else {
		v21 = v20
	}
	v22 = F_get_opfamily_member(m, l5, v14, v14, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		if v22 == int32(0) {
			return int32(0)
		} else {
			v34 = F_makeConst(m, v14, int32(-1), v11, v13, l1, int32(0), v12&int32(1))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				v36 = F_make_opclause(m, v22, l0, v34, l6)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					return v36
				}
			}
		}
	}
}
func F_build_datatype(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
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
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int64
	_ = v123
	var v125 int64
	_ = v125
	var v129 int64
	_ = v129
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)))
	v15 = v13 + v14
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+82)))
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_build_datatype_2))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L5
	} else {
		goto L52
	}
L2:
	;
	v18 = F_palloc(m, int32(48))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_build_datatype_2))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L5
	} else {
		goto L48
	}
L5:
	;
	return int32(0)
L6:
	;
	v24 = F_pstrdup(m, v15+int32(4))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v24
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v27
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+79)))
	switch v29 - int32(98) {
	case 0, 3, 11, 16:
		goto L13
	case 1:
		goto L11
	case 2:
		goto L14
	default:
		goto L10
	case 14:
		goto L12
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v54
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+12)) = uint16(v56)
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+14)) = uint8(v58)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+79)))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+15)) = uint8(v60)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v15)+144))
	if v62 != 0 {
		goto L21
	} else {
		goto L22
	}
L9:
	;
	v54 = int32(2)
	goto L8
L10:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_build_datatype_2))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L5
	} else {
		goto L18
	}
L11:
	;
	v54 = int32(1)
	goto L8
L12:
	;
	if v27 != int32(2249) {
		goto L9
	} else {
		goto L17
	}
L13:
	;
	v54 = int32(0)
	goto L8
L14:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v15)+132))
	v33 = F_type_is_rowtype(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	if v33 != 0 {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	goto L11
L18:
	;
	v43 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15)+79)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v43
	F_errmsg_internal(m, int32(_a_F_build_datatype_7), v11)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_build_datatype_4), int32(2015), int32(_a_F_build_datatype_5))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L21:
	;
	v63 = l2
	goto L23
L22:
	;
	v63 = v62
	goto L23
L23:
	;
	if l2 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v64 = v63
	goto L26
L25:
	;
	v64 = v62
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v64
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+79)))
	switch v66 - int32(98) {
	case 0:
		goto L30
	default:
		goto L28
	case 2:
		goto L29
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = l1
	if v100 != int32(1) {
		goto L39
	} else {
		goto L40
	}
L28:
	;
	v97 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+20)) = uint8(v97)
	v100 = v54
	goto L27
L29:
	;
	v81 = int32(0)
	v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+76)))
	if v82 != int32(_a_F_build_datatype_6) {
		v94 = v81
		v95 = v54
		goto L34
	} else {
		goto L35
	}
L30:
	;
	v69 = int32(0)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+92))
	if v70 == v69 {
		v79 = v69
		goto L31
	} else {
		goto L32
	}
L31:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+20)) = uint8(v79)
	v100 = v54
	goto L27
L32:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
	if v73 != int32(_a_F_build_datatype_0) {
		v79 = v69
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+129)))
	v79 = base.B2i32(v76 != int32(112))
	goto L31
L34:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+20)) = uint8(v94)
	v100 = v95
	goto L27
L35:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+129)))
	if v85 == int32(112) {
		v94 = v81
		v95 = v54
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v15)+132))
	v89 = F_get_base_element_type(m, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v94 = base.B2i32(v89 != int32(0))
	v95 = v93
	goto L34
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+40)) = v129
	m.G0 = v11 + int32(48)
	return v18
L39:
	;
	v125 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+28)) = v125
	v129 = v125
	goto L38
L40:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v104 == int32(2249) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v108 = F_lookup_type_cache(m, v104, int32(_a_F_build_datatype_1))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+13)))
	if v110 == int32(100) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v108)+300))
	v115 = F_lookup_type_cache(m, v113, int32(256))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L5
	} else {
		goto L46
	}
L44:
	;
	v117 = v108
	goto L45
L45:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+188))
	if v118 == int32(0) {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	v117 = v115
	goto L45
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = l3
	v123 = *(*int64)(unsafe.Add(mBase, uint32(v117)+192))
	v129 = v123
	goto L38
L48:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L5
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v15 + int32(4)
	F_errmsg(m, int32(_a_F_build_datatype_8), v11+int32(32))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_build_datatype_4), int32(1984), int32(_a_F_build_datatype_5))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v163 = F_format_type_be(m, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v163
	F_errmsg(m, int32(_a_F_build_datatype_3), v11+int32(16))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_build_datatype_4), int32(2066), int32(_a_F_build_datatype_5))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_build_dummy_expanded_header(m *base.Module, l0 int32) {
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int64
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v6 == int32(0) {
		v9 = F_expanded_record_fetch_tupdesc(m, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = v9
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
			if v12 == int32(0) {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v20 = F_AllocSetContextCreateInternal(m, v15, int32(_a_F_build_dummy_expanded_header_0), int32(0), int32(1024), int32(_a_F_build_dummy_expanded_header_1))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v20
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
					if v26 != 0 {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+64))
						if v27 == v25 {
							v64 = v26
							*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = int32(128)
							v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+32)) = v68
							*(*int32)(unsafe.Add(mBase, uint32(v64)+36)) = v68
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+44)) = v11
							*(*int32)(unsafe.Add(mBase, uint32(v64)+40)) = v71
							v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+68)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v64)+48)) = v74
							v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+84)) = v78
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+88)) = v80
							v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+92)) = v82
							return
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v34 = F_MemoryContextAlloc(m, v29, v25*int32(5)+int32(120))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								base.MemoryFill(m, v34, int32(0), int32(120))
								v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								v41 = int32(513)
								*(*uint16)(unsafe.Add(mBase, uint32(v34)+18)) = uint16(v41)
								v43 = int32(769)
								*(*uint16)(unsafe.Add(mBase, uint32(v34)+12)) = uint16(v43)
								*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v40
								*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = int32(_a_F_build_dummy_expanded_header_2)
								*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(-1)
								*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v34
								*(*int32)(unsafe.Add(mBase, uint32(v34)+14)) = v34
								v52 = v34 + int32(120)
								*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v52
								*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = int32(1384727874)
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
								*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v52 + v56<<(uint(int32(2))%32)
								v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
								*(*int32)(unsafe.Add(mBase, uint32(v34)+64)) = v61
								*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v34
								v64 = v34
								*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = int32(128)
								v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								*(*int32)(unsafe.Add(mBase, uint32(v64)+32)) = v68
								*(*int32)(unsafe.Add(mBase, uint32(v64)+36)) = v68
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v64)+44)) = v11
								*(*int32)(unsafe.Add(mBase, uint32(v64)+40)) = v71
								v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v64)+68)) = int32(0)
								*(*int64)(unsafe.Add(mBase, uint32(v64)+48)) = v74
								v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
								*(*int32)(unsafe.Add(mBase, uint32(v64)+84)) = v78
								v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
								*(*int32)(unsafe.Add(mBase, uint32(v64)+88)) = v80
								v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
								*(*int32)(unsafe.Add(mBase, uint32(v64)+92)) = v82
								return
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v34 = F_MemoryContextAlloc(m, v29, v25*int32(5)+int32(120))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							base.MemoryFill(m, v34, int32(0), int32(120))
							v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
							v41 = int32(513)
							*(*uint16)(unsafe.Add(mBase, uint32(v34)+18)) = uint16(v41)
							v43 = int32(769)
							*(*uint16)(unsafe.Add(mBase, uint32(v34)+12)) = uint16(v43)
							*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v40
							*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = int32(_a_F_build_dummy_expanded_header_2)
							*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v34
							*(*int32)(unsafe.Add(mBase, uint32(v34)+14)) = v34
							v52 = v34 + int32(120)
							*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v52
							*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = int32(1384727874)
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v52 + v56<<(uint(int32(2))%32)
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							*(*int32)(unsafe.Add(mBase, uint32(v34)+64)) = v61
							*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v34
							v64 = v34
							*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = int32(128)
							v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+32)) = v68
							*(*int32)(unsafe.Add(mBase, uint32(v64)+36)) = v68
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+44)) = v11
							*(*int32)(unsafe.Add(mBase, uint32(v64)+40)) = v71
							v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+68)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v64)+48)) = v74
							v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+84)) = v78
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+88)) = v80
							v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+92)) = v82
							return
						}
					}
				}
			} else {
				F_MemoryContextReset(m, v12)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
					if v26 != 0 {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+64))
						if v27 == v25 {
							v64 = v26
							*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = int32(128)
							v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+32)) = v68
							*(*int32)(unsafe.Add(mBase, uint32(v64)+36)) = v68
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+44)) = v11
							*(*int32)(unsafe.Add(mBase, uint32(v64)+40)) = v71
							v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+68)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v64)+48)) = v74
							v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+84)) = v78
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+88)) = v80
							v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+92)) = v82
							return
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v34 = F_MemoryContextAlloc(m, v29, v25*int32(5)+int32(120))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								base.MemoryFill(m, v34, int32(0), int32(120))
								v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								v41 = int32(513)
								*(*uint16)(unsafe.Add(mBase, uint32(v34)+18)) = uint16(v41)
								v43 = int32(769)
								*(*uint16)(unsafe.Add(mBase, uint32(v34)+12)) = uint16(v43)
								*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v40
								*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = int32(_a_F_build_dummy_expanded_header_2)
								*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(-1)
								*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v34
								*(*int32)(unsafe.Add(mBase, uint32(v34)+14)) = v34
								v52 = v34 + int32(120)
								*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v52
								*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = int32(1384727874)
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
								*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v52 + v56<<(uint(int32(2))%32)
								v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
								*(*int32)(unsafe.Add(mBase, uint32(v34)+64)) = v61
								*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v34
								v64 = v34
								*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = int32(128)
								v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								*(*int32)(unsafe.Add(mBase, uint32(v64)+32)) = v68
								*(*int32)(unsafe.Add(mBase, uint32(v64)+36)) = v68
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v64)+44)) = v11
								*(*int32)(unsafe.Add(mBase, uint32(v64)+40)) = v71
								v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v64)+68)) = int32(0)
								*(*int64)(unsafe.Add(mBase, uint32(v64)+48)) = v74
								v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
								*(*int32)(unsafe.Add(mBase, uint32(v64)+84)) = v78
								v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
								*(*int32)(unsafe.Add(mBase, uint32(v64)+88)) = v80
								v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
								*(*int32)(unsafe.Add(mBase, uint32(v64)+92)) = v82
								return
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v34 = F_MemoryContextAlloc(m, v29, v25*int32(5)+int32(120))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							base.MemoryFill(m, v34, int32(0), int32(120))
							v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
							v41 = int32(513)
							*(*uint16)(unsafe.Add(mBase, uint32(v34)+18)) = uint16(v41)
							v43 = int32(769)
							*(*uint16)(unsafe.Add(mBase, uint32(v34)+12)) = uint16(v43)
							*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v40
							*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = int32(_a_F_build_dummy_expanded_header_2)
							*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v34
							*(*int32)(unsafe.Add(mBase, uint32(v34)+14)) = v34
							v52 = v34 + int32(120)
							*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v52
							*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = int32(1384727874)
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v52 + v56<<(uint(int32(2))%32)
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							*(*int32)(unsafe.Add(mBase, uint32(v34)+64)) = v61
							*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v34
							v64 = v34
							*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = int32(128)
							v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+32)) = v68
							*(*int32)(unsafe.Add(mBase, uint32(v64)+36)) = v68
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+44)) = v11
							*(*int32)(unsafe.Add(mBase, uint32(v64)+40)) = v71
							v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+68)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v64)+48)) = v74
							v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+84)) = v78
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+88)) = v80
							v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+92)) = v82
							return
						}
					}
				}
			}
		}
	} else {
		v11 = v6
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
		if v12 == int32(0) {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v20 = F_AllocSetContextCreateInternal(m, v15, int32(_a_F_build_dummy_expanded_header_0), int32(0), int32(1024), int32(_a_F_build_dummy_expanded_header_1))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v20
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
				if v26 != 0 {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+64))
					if v27 == v25 {
						v64 = v26
						*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = int32(128)
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+32)) = v68
						*(*int32)(unsafe.Add(mBase, uint32(v64)+36)) = v68
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+44)) = v11
						*(*int32)(unsafe.Add(mBase, uint32(v64)+40)) = v71
						v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+68)) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v64)+48)) = v74
						v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+84)) = v78
						v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+88)) = v80
						v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+92)) = v82
						return
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v34 = F_MemoryContextAlloc(m, v29, v25*int32(5)+int32(120))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							base.MemoryFill(m, v34, int32(0), int32(120))
							v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
							v41 = int32(513)
							*(*uint16)(unsafe.Add(mBase, uint32(v34)+18)) = uint16(v41)
							v43 = int32(769)
							*(*uint16)(unsafe.Add(mBase, uint32(v34)+12)) = uint16(v43)
							*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v40
							*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = int32(_a_F_build_dummy_expanded_header_2)
							*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v34
							*(*int32)(unsafe.Add(mBase, uint32(v34)+14)) = v34
							v52 = v34 + int32(120)
							*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v52
							*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = int32(1384727874)
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v52 + v56<<(uint(int32(2))%32)
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							*(*int32)(unsafe.Add(mBase, uint32(v34)+64)) = v61
							*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v34
							v64 = v34
							*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = int32(128)
							v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+32)) = v68
							*(*int32)(unsafe.Add(mBase, uint32(v64)+36)) = v68
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+44)) = v11
							*(*int32)(unsafe.Add(mBase, uint32(v64)+40)) = v71
							v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+68)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v64)+48)) = v74
							v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+84)) = v78
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+88)) = v80
							v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+92)) = v82
							return
						}
					}
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v34 = F_MemoryContextAlloc(m, v29, v25*int32(5)+int32(120))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						base.MemoryFill(m, v34, int32(0), int32(120))
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
						v41 = int32(513)
						*(*uint16)(unsafe.Add(mBase, uint32(v34)+18)) = uint16(v41)
						v43 = int32(769)
						*(*uint16)(unsafe.Add(mBase, uint32(v34)+12)) = uint16(v43)
						*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v40
						*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = int32(_a_F_build_dummy_expanded_header_2)
						*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v34
						*(*int32)(unsafe.Add(mBase, uint32(v34)+14)) = v34
						v52 = v34 + int32(120)
						*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v52
						*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = int32(1384727874)
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
						*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v52 + v56<<(uint(int32(2))%32)
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
						*(*int32)(unsafe.Add(mBase, uint32(v34)+64)) = v61
						*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v34
						v64 = v34
						*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = int32(128)
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+32)) = v68
						*(*int32)(unsafe.Add(mBase, uint32(v64)+36)) = v68
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+44)) = v11
						*(*int32)(unsafe.Add(mBase, uint32(v64)+40)) = v71
						v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+68)) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v64)+48)) = v74
						v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+84)) = v78
						v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+88)) = v80
						v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+92)) = v82
						return
					}
				}
			}
		} else {
			F_MemoryContextReset(m, v12)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
				if v26 != 0 {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+64))
					if v27 == v25 {
						v64 = v26
						*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = int32(128)
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+32)) = v68
						*(*int32)(unsafe.Add(mBase, uint32(v64)+36)) = v68
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+44)) = v11
						*(*int32)(unsafe.Add(mBase, uint32(v64)+40)) = v71
						v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+68)) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v64)+48)) = v74
						v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+84)) = v78
						v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+88)) = v80
						v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+92)) = v82
						return
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v34 = F_MemoryContextAlloc(m, v29, v25*int32(5)+int32(120))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							base.MemoryFill(m, v34, int32(0), int32(120))
							v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
							v41 = int32(513)
							*(*uint16)(unsafe.Add(mBase, uint32(v34)+18)) = uint16(v41)
							v43 = int32(769)
							*(*uint16)(unsafe.Add(mBase, uint32(v34)+12)) = uint16(v43)
							*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v40
							*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = int32(_a_F_build_dummy_expanded_header_2)
							*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v34
							*(*int32)(unsafe.Add(mBase, uint32(v34)+14)) = v34
							v52 = v34 + int32(120)
							*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v52
							*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = int32(1384727874)
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v52 + v56<<(uint(int32(2))%32)
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							*(*int32)(unsafe.Add(mBase, uint32(v34)+64)) = v61
							*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v34
							v64 = v34
							*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = int32(128)
							v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+32)) = v68
							*(*int32)(unsafe.Add(mBase, uint32(v64)+36)) = v68
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+44)) = v11
							*(*int32)(unsafe.Add(mBase, uint32(v64)+40)) = v71
							v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+68)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v64)+48)) = v74
							v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+84)) = v78
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+88)) = v80
							v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+92)) = v82
							return
						}
					}
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v34 = F_MemoryContextAlloc(m, v29, v25*int32(5)+int32(120))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						base.MemoryFill(m, v34, int32(0), int32(120))
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
						v41 = int32(513)
						*(*uint16)(unsafe.Add(mBase, uint32(v34)+18)) = uint16(v41)
						v43 = int32(769)
						*(*uint16)(unsafe.Add(mBase, uint32(v34)+12)) = uint16(v43)
						*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v40
						*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = int32(_a_F_build_dummy_expanded_header_2)
						*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v34
						*(*int32)(unsafe.Add(mBase, uint32(v34)+14)) = v34
						v52 = v34 + int32(120)
						*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v52
						*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = int32(1384727874)
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
						*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v52 + v56<<(uint(int32(2))%32)
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
						*(*int32)(unsafe.Add(mBase, uint32(v34)+64)) = v61
						*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v34
						v64 = v34
						*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = int32(128)
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+32)) = v68
						*(*int32)(unsafe.Add(mBase, uint32(v64)+36)) = v68
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+44)) = v11
						*(*int32)(unsafe.Add(mBase, uint32(v64)+40)) = v71
						v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+68)) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v64)+48)) = v74
						v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+84)) = v78
						v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+88)) = v80
						v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+92)) = v82
						return
					}
				}
			}
		}
	}
}
func F_build_implied_join_equality(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	v8 = F_copyObjectImpl(m, l3)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = F_copyObjectImpl(m, l4)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = F_make_opclause(m, l1, v8, v12, l2)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v17 = int32(0)
				v22 = F_make_restrictinfo(m, l0, v14, int32(1), v17, v17, v17, l6, l5, v17, v17)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+10)))
					if v24 != 0 {
						v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+10)))
						if v53 != 0 {
							F_check_memoizable(m, v22)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								return v22
							}
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
							if v54 == int32(0) {
								F_check_memoizable(m, v22)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									return v22
								}
							} else {
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
								if v57 != int32(17) {
									F_check_memoizable(m, v22)
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										return v22
									}
								} else {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
									if v60 == int32(0) {
										F_check_memoizable(m, v22)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											return v22
										}
									} else {
										v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
										if v63 != int32(2) {
											F_check_memoizable(m, v22)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												return v22
											}
										} else {
											v66 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
											v67 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
											v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
											v69 = F_exprType(m, v68)
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return int32(0)
											} else {
												v71 = F_op_hashjoinable(m, v66, v69)
												mBase = m.M
												v72 = m.ExcPending
												if v72 != 0 {
													return int32(0)
												} else {
													if v71 == int32(0) {
														F_check_memoizable(m, v22)
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int32(0)
														} else {
															return v22
														}
													} else {
														v75 = F_contain_volatile_functions(m, v22)
														mBase = m.M
														v76 = m.ExcPending
														if v76 != 0 {
															return int32(0)
														} else {
															if v75 != 0 {
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v22)+124)) = v66
															}
															F_check_memoizable(m, v22)
															mBase = m.M
															v81 = m.ExcPending
															if v81 != 0 {
																return int32(0)
															} else {
																return v22
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
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
						if v25 == int32(0) {
							v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+10)))
							if v53 != 0 {
								F_check_memoizable(m, v22)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									return v22
								}
							} else {
								v54 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
								if v54 == int32(0) {
									F_check_memoizable(m, v22)
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										return v22
									}
								} else {
									v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
									if v57 != int32(17) {
										F_check_memoizable(m, v22)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											return v22
										}
									} else {
										v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
										if v60 == int32(0) {
											F_check_memoizable(m, v22)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												return v22
											}
										} else {
											v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
											if v63 != int32(2) {
												F_check_memoizable(m, v22)
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return int32(0)
												} else {
													return v22
												}
											} else {
												v66 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
												v67 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
												v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
												v69 = F_exprType(m, v68)
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return int32(0)
												} else {
													v71 = F_op_hashjoinable(m, v66, v69)
													mBase = m.M
													v72 = m.ExcPending
													if v72 != 0 {
														return int32(0)
													} else {
														if v71 == int32(0) {
															F_check_memoizable(m, v22)
															mBase = m.M
															v81 = m.ExcPending
															if v81 != 0 {
																return int32(0)
															} else {
																return v22
															}
														} else {
															v75 = F_contain_volatile_functions(m, v22)
															mBase = m.M
															v76 = m.ExcPending
															if v76 != 0 {
																return int32(0)
															} else {
																if v75 != 0 {
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v22)+124)) = v66
																}
																F_check_memoizable(m, v22)
																mBase = m.M
																v81 = m.ExcPending
																if v81 != 0 {
																	return int32(0)
																} else {
																	return v22
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
							v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
							if v28 != int32(17) {
								v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+10)))
								if v53 != 0 {
									F_check_memoizable(m, v22)
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										return v22
									}
								} else {
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
									if v54 == int32(0) {
										F_check_memoizable(m, v22)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											return v22
										}
									} else {
										v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
										if v57 != int32(17) {
											F_check_memoizable(m, v22)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												return v22
											}
										} else {
											v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
											if v60 == int32(0) {
												F_check_memoizable(m, v22)
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return int32(0)
												} else {
													return v22
												}
											} else {
												v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
												if v63 != int32(2) {
													F_check_memoizable(m, v22)
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return int32(0)
													} else {
														return v22
													}
												} else {
													v66 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
													v67 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
													v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
													v69 = F_exprType(m, v68)
													mBase = m.M
													v70 = m.ExcPending
													if v70 != 0 {
														return int32(0)
													} else {
														v71 = F_op_hashjoinable(m, v66, v69)
														mBase = m.M
														v72 = m.ExcPending
														if v72 != 0 {
															return int32(0)
														} else {
															if v71 == int32(0) {
																F_check_memoizable(m, v22)
																mBase = m.M
																v81 = m.ExcPending
																if v81 != 0 {
																	return int32(0)
																} else {
																	return v22
																}
															} else {
																v75 = F_contain_volatile_functions(m, v22)
																mBase = m.M
																v76 = m.ExcPending
																if v76 != 0 {
																	return int32(0)
																} else {
																	if v75 != 0 {
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v22)+124)) = v66
																	}
																	F_check_memoizable(m, v22)
																	mBase = m.M
																	v81 = m.ExcPending
																	if v81 != 0 {
																		return int32(0)
																	} else {
																		return v22
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
								v31 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
								if v31 == int32(0) {
									v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+10)))
									if v53 != 0 {
										F_check_memoizable(m, v22)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											return v22
										}
									} else {
										v54 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
										if v54 == int32(0) {
											F_check_memoizable(m, v22)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												return v22
											}
										} else {
											v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
											if v57 != int32(17) {
												F_check_memoizable(m, v22)
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return int32(0)
												} else {
													return v22
												}
											} else {
												v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
												if v60 == int32(0) {
													F_check_memoizable(m, v22)
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return int32(0)
													} else {
														return v22
													}
												} else {
													v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
													if v63 != int32(2) {
														F_check_memoizable(m, v22)
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int32(0)
														} else {
															return v22
														}
													} else {
														v66 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
														v67 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
														v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
														v69 = F_exprType(m, v68)
														mBase = m.M
														v70 = m.ExcPending
														if v70 != 0 {
															return int32(0)
														} else {
															v71 = F_op_hashjoinable(m, v66, v69)
															mBase = m.M
															v72 = m.ExcPending
															if v72 != 0 {
																return int32(0)
															} else {
																if v71 == int32(0) {
																	F_check_memoizable(m, v22)
																	mBase = m.M
																	v81 = m.ExcPending
																	if v81 != 0 {
																		return int32(0)
																	} else {
																		return v22
																	}
																} else {
																	v75 = F_contain_volatile_functions(m, v22)
																	mBase = m.M
																	v76 = m.ExcPending
																	if v76 != 0 {
																		return int32(0)
																	} else {
																		if v75 != 0 {
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v22)+124)) = v66
																		}
																		F_check_memoizable(m, v22)
																		mBase = m.M
																		v81 = m.ExcPending
																		if v81 != 0 {
																			return int32(0)
																		} else {
																			return v22
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
									v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
									if v34 != int32(2) {
										v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+10)))
										if v53 != 0 {
											F_check_memoizable(m, v22)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												return v22
											}
										} else {
											v54 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
											if v54 == int32(0) {
												F_check_memoizable(m, v22)
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return int32(0)
												} else {
													return v22
												}
											} else {
												v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
												if v57 != int32(17) {
													F_check_memoizable(m, v22)
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return int32(0)
													} else {
														return v22
													}
												} else {
													v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
													if v60 == int32(0) {
														F_check_memoizable(m, v22)
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int32(0)
														} else {
															return v22
														}
													} else {
														v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
														if v63 != int32(2) {
															F_check_memoizable(m, v22)
															mBase = m.M
															v81 = m.ExcPending
															if v81 != 0 {
																return int32(0)
															} else {
																return v22
															}
														} else {
															v66 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
															v67 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
															v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
															v69 = F_exprType(m, v68)
															mBase = m.M
															v70 = m.ExcPending
															if v70 != 0 {
																return int32(0)
															} else {
																v71 = F_op_hashjoinable(m, v66, v69)
																mBase = m.M
																v72 = m.ExcPending
																if v72 != 0 {
																	return int32(0)
																} else {
																	if v71 == int32(0) {
																		F_check_memoizable(m, v22)
																		mBase = m.M
																		v81 = m.ExcPending
																		if v81 != 0 {
																			return int32(0)
																		} else {
																			return v22
																		}
																	} else {
																		v75 = F_contain_volatile_functions(m, v22)
																		mBase = m.M
																		v76 = m.ExcPending
																		if v76 != 0 {
																			return int32(0)
																		} else {
																			if v75 != 0 {
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v22)+124)) = v66
																			}
																			F_check_memoizable(m, v22)
																			mBase = m.M
																			v81 = m.ExcPending
																			if v81 != 0 {
																				return int32(0)
																			} else {
																				return v22
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
										v37 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
										v38 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
										v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
										v40 = F_exprType(m, v39)
										mBase = m.M
										v41 = m.ExcPending
										if v41 != 0 {
											return int32(0)
										} else {
											v42 = F_op_mergejoinable(m, v37, v40)
											mBase = m.M
											v43 = m.ExcPending
											if v43 != 0 {
												return int32(0)
											} else {
												if v42 == int32(0) {
													v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+10)))
													if v53 != 0 {
														F_check_memoizable(m, v22)
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int32(0)
														} else {
															return v22
														}
													} else {
														v54 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
														if v54 == int32(0) {
															F_check_memoizable(m, v22)
															mBase = m.M
															v81 = m.ExcPending
															if v81 != 0 {
																return int32(0)
															} else {
																return v22
															}
														} else {
															v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
															if v57 != int32(17) {
																F_check_memoizable(m, v22)
																mBase = m.M
																v81 = m.ExcPending
																if v81 != 0 {
																	return int32(0)
																} else {
																	return v22
																}
															} else {
																v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
																if v60 == int32(0) {
																	F_check_memoizable(m, v22)
																	mBase = m.M
																	v81 = m.ExcPending
																	if v81 != 0 {
																		return int32(0)
																	} else {
																		return v22
																	}
																} else {
																	v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
																	if v63 != int32(2) {
																		F_check_memoizable(m, v22)
																		mBase = m.M
																		v81 = m.ExcPending
																		if v81 != 0 {
																			return int32(0)
																		} else {
																			return v22
																		}
																	} else {
																		v66 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
																		v67 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
																		v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
																		v69 = F_exprType(m, v68)
																		mBase = m.M
																		v70 = m.ExcPending
																		if v70 != 0 {
																			return int32(0)
																		} else {
																			v71 = F_op_hashjoinable(m, v66, v69)
																			mBase = m.M
																			v72 = m.ExcPending
																			if v72 != 0 {
																				return int32(0)
																			} else {
																				if v71 == int32(0) {
																					F_check_memoizable(m, v22)
																					mBase = m.M
																					v81 = m.ExcPending
																					if v81 != 0 {
																						return int32(0)
																					} else {
																						return v22
																					}
																				} else {
																					v75 = F_contain_volatile_functions(m, v22)
																					mBase = m.M
																					v76 = m.ExcPending
																					if v76 != 0 {
																						return int32(0)
																					} else {
																						if v75 != 0 {
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v22)+124)) = v66
																						}
																						F_check_memoizable(m, v22)
																						mBase = m.M
																						v81 = m.ExcPending
																						if v81 != 0 {
																							return int32(0)
																						} else {
																							return v22
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
													v46 = F_contain_volatile_functions(m, v22)
													mBase = m.M
													v47 = m.ExcPending
													if v47 != 0 {
														return int32(0)
													} else {
														if v46 != 0 {
															v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+10)))
															if v53 != 0 {
																F_check_memoizable(m, v22)
																mBase = m.M
																v81 = m.ExcPending
																if v81 != 0 {
																	return int32(0)
																} else {
																	return v22
																}
															} else {
																v54 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
																if v54 == int32(0) {
																	F_check_memoizable(m, v22)
																	mBase = m.M
																	v81 = m.ExcPending
																	if v81 != 0 {
																		return int32(0)
																	} else {
																		return v22
																	}
																} else {
																	v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
																	if v57 != int32(17) {
																		F_check_memoizable(m, v22)
																		mBase = m.M
																		v81 = m.ExcPending
																		if v81 != 0 {
																			return int32(0)
																		} else {
																			return v22
																		}
																	} else {
																		v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
																		if v60 == int32(0) {
																			F_check_memoizable(m, v22)
																			mBase = m.M
																			v81 = m.ExcPending
																			if v81 != 0 {
																				return int32(0)
																			} else {
																				return v22
																			}
																		} else {
																			v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
																			if v63 != int32(2) {
																				F_check_memoizable(m, v22)
																				mBase = m.M
																				v81 = m.ExcPending
																				if v81 != 0 {
																					return int32(0)
																				} else {
																					return v22
																				}
																			} else {
																				v66 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
																				v67 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
																				v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
																				v69 = F_exprType(m, v68)
																				mBase = m.M
																				v70 = m.ExcPending
																				if v70 != 0 {
																					return int32(0)
																				} else {
																					v71 = F_op_hashjoinable(m, v66, v69)
																					mBase = m.M
																					v72 = m.ExcPending
																					if v72 != 0 {
																						return int32(0)
																					} else {
																						if v71 == int32(0) {
																							F_check_memoizable(m, v22)
																							mBase = m.M
																							v81 = m.ExcPending
																							if v81 != 0 {
																								return int32(0)
																							} else {
																								return v22
																							}
																						} else {
																							v75 = F_contain_volatile_functions(m, v22)
																							mBase = m.M
																							v76 = m.ExcPending
																							if v76 != 0 {
																								return int32(0)
																							} else {
																								if v75 != 0 {
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(v22)+124)) = v66
																								}
																								F_check_memoizable(m, v22)
																								mBase = m.M
																								v81 = m.ExcPending
																								if v81 != 0 {
																									return int32(0)
																								} else {
																									return v22
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
															v48 = F_get_mergejoin_opfamilies(m, v37)
															mBase = m.M
															v49 = m.ExcPending
															if v49 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = v48
																v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+10)))
																if v53 != 0 {
																	F_check_memoizable(m, v22)
																	mBase = m.M
																	v81 = m.ExcPending
																	if v81 != 0 {
																		return int32(0)
																	} else {
																		return v22
																	}
																} else {
																	v54 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
																	if v54 == int32(0) {
																		F_check_memoizable(m, v22)
																		mBase = m.M
																		v81 = m.ExcPending
																		if v81 != 0 {
																			return int32(0)
																		} else {
																			return v22
																		}
																	} else {
																		v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
																		if v57 != int32(17) {
																			F_check_memoizable(m, v22)
																			mBase = m.M
																			v81 = m.ExcPending
																			if v81 != 0 {
																				return int32(0)
																			} else {
																				return v22
																			}
																		} else {
																			v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
																			if v60 == int32(0) {
																				F_check_memoizable(m, v22)
																				mBase = m.M
																				v81 = m.ExcPending
																				if v81 != 0 {
																					return int32(0)
																				} else {
																					return v22
																				}
																			} else {
																				v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
																				if v63 != int32(2) {
																					F_check_memoizable(m, v22)
																					mBase = m.M
																					v81 = m.ExcPending
																					if v81 != 0 {
																						return int32(0)
																					} else {
																						return v22
																					}
																				} else {
																					v66 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
																					v67 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
																					v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
																					v69 = F_exprType(m, v68)
																					mBase = m.M
																					v70 = m.ExcPending
																					if v70 != 0 {
																						return int32(0)
																					} else {
																						v71 = F_op_hashjoinable(m, v66, v69)
																						mBase = m.M
																						v72 = m.ExcPending
																						if v72 != 0 {
																							return int32(0)
																						} else {
																							if v71 == int32(0) {
																								F_check_memoizable(m, v22)
																								mBase = m.M
																								v81 = m.ExcPending
																								if v81 != 0 {
																									return int32(0)
																								} else {
																									return v22
																								}
																							} else {
																								v75 = F_contain_volatile_functions(m, v22)
																								mBase = m.M
																								v76 = m.ExcPending
																								if v76 != 0 {
																									return int32(0)
																								} else {
																									if v75 != 0 {
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(v22)+124)) = v66
																									}
																									F_check_memoizable(m, v22)
																									mBase = m.M
																									v81 = m.ExcPending
																									if v81 != 0 {
																										return int32(0)
																									} else {
																										return v22
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
func F_build_joinrel_restrictlist(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	v6 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v13 = F_bms_union(m, v11, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+212))
	if v17 == int32(0) {
		v222 = v6
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l3)+212))
	if v223 == int32(0) {
		v429 = v222
		goto L60
	} else {
		goto L61
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v20 <= int32(0) {
		v222 = v6
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v30 = v6
	v32 = v6
	goto L6
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v30<<(uint(int32(2))%32))))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v40 = int32(0)
	if v38 == v40 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v222 = v208
	goto L3
L8:
	;
	v210 = v30 + int32(1)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v210 < v211 {
		v30 = v210
		v32 = v208
		goto L6
	} else {
		goto L59
	}
L9:
	;
	if v93 == int32(0) {
		v208 = v32
		goto L8
	} else {
		goto L23
	}
L10:
	;
	v93 = int32(1)
	goto L9
L11:
	;
	goto L12
L12:
	;
	if v39 == int32(0) {
		v86 = v40
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v93 = v86
	goto L9
L14:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v50 < v49 {
		v86 = v40
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v52 = int32(1)
	if v49 <= v52 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v55 = v52
	goto L18
L17:
	;
	v55 = v49
	goto L18
L18:
	;
	v56 = int32(8)
	v61 = int32(0)
	goto L19
L19:
	;
	v68 = v61 << (uint(int32(2)) % 32)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v38+v56+v68)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v39+v56+v68)))
	v75 = v70 & (v72 ^ int32(-1))
	v77 = base.B2i32(v75 == int32(0))
	if v75 != 0 {
		v86 = v77
		goto L13
	} else {
		goto L21
	}
L20:
	;
	v86 = v77
	goto L13
L21:
	;
	v79 = v61 + int32(1)
	if v79 != v55 {
		v61 = v79
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+11)))
	if v96 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v206 = F_list_append_unique_ptr(m, v32, v37)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L58
	}
L25:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+12)))
	if v99 != int32(1) {
		goto L24
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
	v103 = int32(0)
	if v102 == v103 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L27
L29:
	;
	if v156 == int32(0) {
		v208 = v32
		goto L8
	} else {
		goto L43
	}
L30:
	;
	v156 = int32(1)
	goto L29
L31:
	;
	goto L32
L32:
	;
	if v13 == int32(0) {
		v149 = v103
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v156 = v149
	goto L29
L34:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v113 < v112 {
		v149 = v103
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v115 = int32(1)
	if v112 <= v115 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v118 = v115
	goto L38
L37:
	;
	v118 = v112
	goto L38
L38:
	;
	v119 = int32(8)
	v124 = int32(0)
	goto L39
L39:
	;
	v131 = v124 << (uint(int32(2)) % 32)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v102+v119+v131)))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v13+v119+v131)))
	v138 = v133 & (v135 ^ int32(-1))
	v140 = base.B2i32(v138 == int32(0))
	if v138 != 0 {
		v149 = v140
		goto L33
	} else {
		goto L41
	}
L40:
	;
	v149 = v140
	goto L33
L41:
	;
	v142 = v124 + int32(1)
	if v142 != v118 {
		v124 = v142
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v37)+36))
	v160 = int32(0)
	if base.B2i32(v159 == v160)|base.B2i32(v13 == v160) != 0 {
		v205 = v160
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v205 != 0 {
		v208 = v32
		goto L8
	} else {
		goto L57
	}
L45:
	;
	goto L44
L46:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v170 < v171 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v173 = v170
	goto L49
L48:
	;
	v173 = v171
	goto L49
L49:
	;
	if v173 <= int32(1) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v176 = int32(1)
	goto L52
L51:
	;
	v176 = v173
	goto L52
L52:
	;
	v177 = int32(8)
	v182 = int32(0)
	goto L53
L53:
	;
	v189 = v182 << (uint(int32(2)) % 32)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v13+v177+v189)))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v159+v177+v189)))
	v194 = v191 & v193
	v196 = base.B2i32(v194 != int32(0))
	if v194 != 0 {
		v205 = v196
		goto L45
	} else {
		goto L55
	}
L54:
	;
	v205 = v196
	goto L45
L55:
	;
	v198 = v182 + int32(1)
	if v198 != v176 {
		v182 = v198
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	goto L24
L58:
	;
	v208 = v206
	goto L8
L59:
	;
	goto L7
L60:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v432 = F_generate_join_implied_equalities(m, l0, v430, v431, l3, l4)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L117
	}
L61:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	if v226 <= int32(0) {
		v429 = v222
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v237 = int32(0)
	v239 = v222
	goto L63
L63:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v223)+12))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v240+v237<<(uint(int32(2))%32))))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+32))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v247 = int32(0)
	if v245 == v247 {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	v429 = v415
	goto L60
L65:
	;
	v417 = v237 + int32(1)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	if v417 < v418 {
		v237 = v417
		v239 = v415
		goto L63
	} else {
		goto L116
	}
L66:
	;
	if v300 == int32(0) {
		v415 = v239
		goto L65
	} else {
		goto L80
	}
L67:
	;
	v300 = int32(1)
	goto L66
L68:
	;
	goto L69
L69:
	;
	if v246 == int32(0) {
		v293 = v247
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v300 = v293
	goto L66
L71:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v245)+4))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v246)+4))
	if v257 < v256 {
		v293 = v247
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v259 = int32(1)
	if v256 <= v259 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v262 = v259
	goto L75
L74:
	;
	v262 = v256
	goto L75
L75:
	;
	v263 = int32(8)
	v268 = int32(0)
	goto L76
L76:
	;
	v275 = v268 << (uint(int32(2)) % 32)
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v245+v263+v275)))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v246+v263+v275)))
	v282 = v277 & (v279 ^ int32(-1))
	v284 = base.B2i32(v282 == int32(0))
	if v282 != 0 {
		v293 = v284
		goto L70
	} else {
		goto L78
	}
L77:
	;
	v293 = v284
	goto L70
L78:
	;
	v286 = v268 + int32(1)
	if v286 != v262 {
		v268 = v286
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+11)))
	if v303 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v413 = F_list_append_unique_ptr(m, v239, v244)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L115
	}
L82:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+12)))
	if v306 != int32(1) {
		goto L81
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v244)+32))
	v310 = int32(0)
	if v309 == v310 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	goto L84
L86:
	;
	if v363 == int32(0) {
		v415 = v239
		goto L65
	} else {
		goto L100
	}
L87:
	;
	v363 = int32(1)
	goto L86
L88:
	;
	goto L89
L89:
	;
	if v13 == int32(0) {
		v356 = v310
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v363 = v356
	goto L86
L91:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v309)+4))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v320 < v319 {
		v356 = v310
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v322 = int32(1)
	if v319 <= v322 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v325 = v322
	goto L95
L94:
	;
	v325 = v319
	goto L95
L95:
	;
	v326 = int32(8)
	v331 = int32(0)
	goto L96
L96:
	;
	v338 = v331 << (uint(int32(2)) % 32)
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v309+v326+v338)))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v13+v326+v338)))
	v345 = v340 & (v342 ^ int32(-1))
	v347 = base.B2i32(v345 == int32(0))
	if v345 != 0 {
		v356 = v347
		goto L90
	} else {
		goto L98
	}
L97:
	;
	v356 = v347
	goto L90
L98:
	;
	v349 = v331 + int32(1)
	if v349 != v325 {
		v331 = v349
		goto L96
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v244)+36))
	v367 = int32(0)
	if base.B2i32(v366 == v367)|base.B2i32(v13 == v367) != 0 {
		v412 = v367
		goto L102
	} else {
		goto L103
	}
L101:
	;
	if v412 != 0 {
		v415 = v239
		goto L65
	} else {
		goto L114
	}
L102:
	;
	goto L101
L103:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v366)+4))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v377 < v378 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v380 = v377
	goto L106
L105:
	;
	v380 = v378
	goto L106
L106:
	;
	if v380 <= int32(1) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v383 = int32(1)
	goto L109
L108:
	;
	v383 = v380
	goto L109
L109:
	;
	v384 = int32(8)
	v389 = int32(0)
	goto L110
L110:
	;
	v396 = v389 << (uint(int32(2)) % 32)
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v13+v384+v396)))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v366+v384+v396)))
	v401 = v398 & v400
	v403 = base.B2i32(v401 != int32(0))
	if v401 != 0 {
		v412 = v403
		goto L102
	} else {
		goto L112
	}
L111:
	;
	v412 = v403
	goto L102
L112:
	;
	v405 = v389 + int32(1)
	if v405 != v383 {
		v389 = v405
		goto L110
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	goto L81
L115:
	;
	v415 = v413
	goto L65
L116:
	;
	goto L64
L117:
	;
	v434 = F_list_concat(m, v429, v432)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	return v434
}
func F_build_pertrans_for_aggref(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v333 int32
	_ = v333
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v382 int32
	_ = v382
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v409 int32
	_ = v409
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	v10 = l9
	v13 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+212))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v13
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v13)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l3
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+180)) = uint8(v10)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = l8
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v29
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
	if v37 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v39 = v38
	goto L3
L2:
	;
	v39 = v13
	goto L3
L3:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	if v40 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v42 = v41
	goto L6
L5:
	;
	v42 = v13
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v42
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+49)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v48 = int32(0)
	F_build_aggregate_transfn_expr(m, l10, l11, v39, v46, l5, v47, l4, v48, v19+int32(12), v48)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	v55 = l0 + int32(32)
	F_fmgr_info(m, l4, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v58
	v61 = v45 + int32(1)
	v66 = F_palloc(m, v61<<(uint(int32(3))%32)+int32(20))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+212)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v55
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = l1
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v73 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v75)+12)) = v76
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	*(*uint8)(unsafe.Add(mBase, uint32(v78)+16)) = uint8(v73)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	*(*uint16)(unsafe.Add(mBase, uint32(v81)+18)) = uint16(v61)
	F_get_typlenbyval(m, l5, l0+int32(184), l0+int32(187))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	if l6 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v91 = m.G0
	v93 = v91 - int32(16)
	m.G0 = v93
	v96 = F_palloc0(m, int32(28))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L7
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if l7 != 0 {
		goto L20
	} else {
		goto L21
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+24)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v96)+16)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = int64(9801115369471)
	*(*int64)(unsafe.Add(mBase, uint32(v96))) = int64(4294967304)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v93)+12)) = v96
	v112 = F_list_make1_impl(m, int32(1), v93+int32(8))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v114 = int32(0)
	v116 = F_makeFuncExpr(m, l6, int32(17), v112, v114, v114)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(8)))) = v116
	m.G0 = v93 + int32(16)
	v123 = l0 + int32(60)
	F_fmgr_info(m, l6, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v126
	v129 = F_palloc(m, int32(28))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = v123
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+4)) = l1
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v136 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v135)+8)) = v136
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	*(*int32)(unsafe.Add(mBase, uint32(v138)+12)) = v136
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	*(*uint8)(unsafe.Add(mBase, uint32(v141)+16)) = uint8(v136)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v145 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v144)+18)) = uint16(v145)
	goto L14
L20:
	;
	v151 = m.G0
	v153 = v151 - int32(16)
	m.G0 = v153
	v156 = F_palloc0(m, int32(28))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L7
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+50)))
	if v222 == int32(110) {
		goto L36
	} else {
		goto L37
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v156)+24)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v156)+16)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v156)+8)) = int64(77309411327)
	*(*int64)(unsafe.Add(mBase, uint32(v156))) = int64(4294967304)
	*(*int32)(unsafe.Add(mBase, uint32(v153)+12)) = v156
	v168 = F_palloc0(m, int32(28))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168)+24)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v168)+16)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v168)+8)) = int64(9801115369471)
	*(*int64)(unsafe.Add(mBase, uint32(v168))) = int64(4294967304)
	*(*int32)(unsafe.Add(mBase, uint32(v153)+8)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = v168
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+4)) = v180
	v185 = F_list_make2_impl(m, v153+int32(4), v153)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v187 = int32(0)
	v189 = F_makeFuncExpr(m, l7, int32(2281), v185, v187, v187)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(4)))) = v189
	m.G0 = v153 + int32(16)
	v196 = l0 + int32(88)
	F_fmgr_info(m, l7, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v199
	v202 = F_palloc(m, int32(36))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+220)) = v202
	*(*int32)(unsafe.Add(mBase, uint32(v202))) = v196
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	*(*int32)(unsafe.Add(mBase, uint32(v206)+4)) = l1
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	v209 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v208)+8)) = v209
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	*(*int32)(unsafe.Add(mBase, uint32(v211)+12)) = v209
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	*(*uint8)(unsafe.Add(mBase, uint32(v214)+16)) = uint8(v209)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	v218 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v217)+18)) = uint16(v218)
	goto L22
L29:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	if v390 != 0 {
		goto L68
	} else {
		goto L69
	}
L30:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	v276 = F_ExecTypeFromTL(m, v275)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L7
	} else {
		goto L47
	}
L31:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	if v265 == int32(0) {
		v382 = v263
		goto L29
	} else {
		goto L46
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v250
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)) = uint8(v253)
	if v250 <= int32(0) {
		v262 = v250
		v263 = v251
		v264 = v252
		goto L31
	} else {
		goto L45
	}
L33:
	;
	v243 = int32(0)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l3)+36))
	if v245 != 0 {
		goto L42
	} else {
		goto L43
	}
L34:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v225)+4))
	v250 = v240
	v251 = v240
	v252 = v225
	v253 = v226 ^ int32(1)
	goto L32
L35:
	;
	if v225 == int32(0) {
		goto L33
	} else {
		goto L41
	}
L36:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+51)))
	if v226 != int32(1) {
		goto L35
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = int64(0)
	v233 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)) = uint8(v233)
	v262 = v233
	v263 = v233
	v264 = v233
	goto L31
L39:
	;
	if v225 != 0 {
		goto L34
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	goto L34
L42:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)+4))
	v247 = v246
	goto L44
L43:
	;
	v247 = v243
	goto L44
L44:
	;
	v250 = v247
	v251 = v243
	v252 = v245
	v253 = base.B2i32(int32(0) < v247)
	goto L32
L45:
	;
	v271 = v250
	v272 = v251
	v273 = v252
	v274 = int32(1)
	goto L30
L46:
	;
	v271 = v262
	v272 = v263
	v273 = v264
	v274 = int32(0)
	goto L30
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v276
	v280 = F_ExecInitExtraTupleSlot(m, l2, v276, int32(_a_F_build_pertrans_for_aggref_0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v280
	if v274 == int32(0) {
		v382 = v272
		goto L29
	} else {
		goto L49
	}
L49:
	;
	if v42 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v306 = F_palloc(m, v271<<(uint(int32(1))%32))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L7
	} else {
		goto L57
	}
L51:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l10+v39<<(uint(int32(2))%32))))
	F_get_typlenbyval(m, v290, l0+int32(182), l0+int32(186))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L7
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if v272 <= int32(0) {
		goto L50
	} else {
		goto L55
	}
L54:
	;
	goto L50
L55:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v301 = F_ExecInitExtraTupleSlot(m, l2, v299, int32(_a_F_build_pertrans_for_aggref_0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v301
	goto L50
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v306
	v310 = v271 << (uint(int32(2)) % 32)
	v311 = F_palloc(m, v310)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L7
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v311
	v314 = F_palloc(m, v310)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L7
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v314
	v317 = F_palloc(m, v271)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v317
	if v273 == int32(0) {
		v382 = v272
		goto L29
	} else {
		goto L61
	}
L61:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v273)+4))
	if v322 <= int32(0) {
		v382 = v272
		goto L29
	} else {
		goto L62
	}
L62:
	;
	v333 = int32(0)
	goto L63
L63:
	;
	v343 = v333 << (uint(int32(2)) % 32)
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v273)+12))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v343+v344)))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	v348 = F_get_sortgroupclause_tle(m, v346, v347)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L7
	} else {
		goto L65
	}
L64:
	;
	v382 = v272
	goto L29
L65:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v354 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v348)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v350+v333<<(uint(int32(1))%32)))) = uint16(v354)
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v346)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v356+v343))) = v358
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v348)+4))
	v361 = F_exprCollation(m, v360)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v363+v343))) = v361
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v366+v333))) = uint8(v368)
	v371 = v333 + int32(1)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v273)+4))
	if v371 < v372 {
		v333 = v371
		goto L63
	} else {
		goto L67
	}
L67:
	;
	goto L64
L68:
	;
	v393 = F_palloc(m, v382<<(uint(int32(2))%32))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L7
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v479 = int32(1)
	if v21 <= v479 {
		goto L86
	} else {
		goto L87
	}
L71:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	if v395 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	if v382 == int32(1) {
		goto L79
	} else {
		goto L80
	}
L73:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v395)+4))
	if v398 <= int32(0) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v409 = int32(0)
	goto L75
L75:
	;
	v419 = v409 << (uint(int32(2)) % 32)
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v395)+12))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v421+v419)))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v423)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v393+v419))) = v424
	v427 = v409 + int32(1)
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v395)+4))
	if v427 < v428 {
		v409 = v427
		goto L75
	} else {
		goto L77
	}
L76:
	;
	goto L72
L77:
	;
	goto L76
L78:
	;
	F_pfree(m, v393)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L7
	} else {
		goto L85
	}
L79:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v393)))
	v449 = F_get_opcode(m, v448)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L7
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v458 = F_execTuplesMatchPrepare(m, v455, v382, v456, v393, v457, l1)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L7
	} else {
		goto L84
	}
L82:
	;
	F_fmgr_info(m, v449, l0+int32(144))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L7
	} else {
		goto L83
	}
L83:
	;
	goto L78
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = v458
	goto L78
L85:
	;
	goto L70
L86:
	;
	v482 = v479
	goto L88
L87:
	;
	v482 = v21
	goto L88
L88:
	;
	v485 = F_palloc0(m, v482<<(uint(int32(2))%32))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L7
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v485
	m.G0 = v19 + int32(16)
	return
}
func F_byteage(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
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
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v16 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v46 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v22 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v33 = int32(1)
	if v16&v33 != 0 {
		v45 = int32(base.Ui32(v16)>>(uint(v33)%32)) - v33
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v25 = int32(16)
	goto L10
L9:
	;
	v25 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v22-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v32 = int32(4)
	goto L13
L12:
	;
	v32 = v25
	goto L13
L13:
	;
	v45 = v32
	goto L4
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v76 = int32(1)
	if v16&v76 != 0 {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v52 == int32(18) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v63 = int32(1)
	if v46&v63 != 0 {
		v75 = int32(base.Ui32(v46)>>(uint(v63)%32)) - v63
		goto L15
	} else {
		goto L25
	}
L19:
	;
	v55 = int32(16)
	goto L21
L20:
	;
	v55 = int32(0)
	goto L21
L21:
	;
	if base.Ui32((v52-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v62 = int32(4)
	goto L24
L23:
	;
	v62 = v55
	goto L24
L24:
	;
	v75 = v62
	goto L15
L25:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v75 = int32(base.Ui32(v69)>>(uint(int32(2))%32)) - int32(4)
	goto L15
L26:
	;
	v80 = v76
	goto L28
L27:
	;
	v80 = int32(4)
	goto L28
L28:
	;
	v81 = v9 + v80
	v82 = int32(1)
	if v46&v82 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v86 = v82
	goto L31
L30:
	;
	v86 = int32(4)
	goto L31
L31:
	;
	v87 = v14 + v86
	if v45 < v75 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v89 = v45
	goto L34
L33:
	;
	v89 = v75
	goto L34
L34:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v89) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v152 != v9 {
		goto L53
	} else {
		goto L54
	}
L36:
	;
	v151 = int32(0)
	goto L35
L37:
	;
	v125 = v120
	v126 = v121
	v127 = v122
	goto L47
L38:
	;
	if (v81|v87)&int32(3) != 0 {
		v120 = v81
		v121 = v87
		v122 = v89
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v113 = v81
	v114 = v87
	v115 = v89
	goto L40
L40:
	;
	if v115 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v97 = v81
	v98 = v87
	v99 = v89
	goto L42
L42:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v102 != v103 {
		v120 = v97
		v121 = v98
		v122 = v99
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v113 = v108
	v114 = v106
	v115 = v110
	goto L40
L44:
	;
	v105 = int32(4)
	v106 = v98 + v105
	v108 = v97 + v105
	v110 = v99 - v105
	if base.Ui32(int32(3)) < base.Ui32(v110) {
		v97 = v108
		v98 = v106
		v99 = v110
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v120 = v113
	v121 = v114
	v122 = v115
	goto L37
L47:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	if v130 == v131 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v151 = v130 - v131
	goto L35
L49:
	;
	v133 = int32(1)
	v138 = v127 - v133
	if v138 != 0 {
		v125 = v125 + v133
		v126 = v126 + v133
		v127 = v138
		goto L47
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	goto L36
L53:
	;
	F_pfree(m, v9)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v156 != v14 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	F_pfree(m, v14)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v160 = int32(0)
	return base.B2i32(v151 == v160)&base.B2i32(v75 <= v45) | base.B2i32(v160 < v151)
L60:
	;
	goto L59
}
func F_byteaoverlay(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v12 = F_bytea_overlay(m, v3, v8, v10, v11)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return v12
			}
		}
	}
}
func F_byteasend(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_copy(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
