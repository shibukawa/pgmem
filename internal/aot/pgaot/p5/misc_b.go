package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
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
	v12 = *(*int32)(unsafe.Add(mBase, _consts[8]))
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
	v21 = *(*int32)(unsafe.Add(mBase, _consts[513]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v22 <= int32(0) {
		v51 = v2
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _consts[8]))
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
	v29 = *(*int32)(unsafe.Add(mBase, _consts[514]))
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
	var v7 int32
	_ = v7
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
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
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 float64
	_ = v120
	var v124 int64
	_ = v124
	var v126 int64
	_ = v126
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 float64
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v278 int32
	_ = v278
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v392 int32
	_ = v392
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
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
	var v444 int32
	_ = v444
	var v452 int32
	_ = v452
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
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
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 float32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v623 int32
	_ = v623
	var v630 int32
	_ = v630
	var v631 float64
	_ = v631
	var v633 float64
	_ = v633
	var v639 int64
	_ = v639
	var v641 int64
	_ = v641
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v665 int32
	_ = v665
	var v671 int32
	_ = v671
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	v7 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(16)
	m.G0 = v29
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v31 == v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l5)+160))
	v35 = int32(4536272)
	v36 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l5)+184))
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l5)+204))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	v45 = l5 + int32(48)
	v46 = F_HnswFormIndexValue(m, v29+int32(12), l2, l3, v43, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v29 + int32(16)
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v36
	v705 = *(*int32)(unsafe.Add(mBase, uint32(l5)+184))
	F_MemoryContextReset(m, v705)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L5
	} else {
		goto L180
	}
L5:
	;
	return
L6:
	;
	if v46 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v51 = v34 + int32(76)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v53 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v81 = F_LWLockAcquire(m, v51, int32(1))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L5
	} else {
		goto L17
	}
L9:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	if base.Ui32((v57-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v79 = int32(6)
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v71 = int32(1)
	if v53&v71 != 0 {
		v79 = int32(base.Ui32(v53) >> (uint(v71) % 32))
		goto L8
	} else {
		goto L16
	}
L12:
	;
	v64 = int32(18)
	if v57&int32(255) == v64 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v70 = v64
	goto L15
L14:
	;
	v70 = int32(2)
	goto L15
L15:
	;
	v79 = v70
	goto L8
L16:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v79 = int32(base.Ui32(v75) >> (uint(int32(2)) % 32))
	goto L8
L17:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+92)))
	if v83 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(1)
	if v623 != 0 {
		goto L168
	} else {
		goto L169
	}
L19:
	;
	F_LWLockRelease(m, v51)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v93 = v34 + int32(52)
	v95 = F_LWLockAcquire(m, v93, int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L5
	} else {
		goto L25
	}
L22:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v90 = F_HnswInsertTupleOnDisk(m, l0, v45, v88, l1, int32(1))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	if v90 != 0 {
		goto L18
	} else {
		goto L24
	}
L24:
	;
	goto L4
L25:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v34)+68))
	v102 = F_add_size(m, v97, base.B2i32(v40 != int32(0))<<(uint(int32(20))%32))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	if base.Ui32(v104) <= base.Ui32(v102) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	F_LWLockRelease(m, v93)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L5
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	v155 = *(*float64)(unsafe.Add(mBase, uint32(l5)+168))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l5)+176))
	v158 = l5 + int32(188)
	v159 = F_HnswInitElement(m, v40, l1, v154, v155, v156, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L5
	} else {
		goto L52
	}
L30:
	;
	F_LWLockRelease(m, v51)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	v111 = F_LWLockAcquire(m, v51, int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+92)))
	if v113 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v118 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L5
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	F_LWLockRelease(m, v51)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L5
	} else {
		goto L49
	}
L36:
	;
	if v118 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v120 = *(*float64)(unsafe.Add(mBase, uint32(v34)+8))
	if base.F64_lt(base.F64_abs(v120), float64(9.223372036854776e+18)) != 0 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	goto L39
L39:
	;
	F_FlushPages(m, l5)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L5
	} else {
		goto L48
	}
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = v126
	F_errmsg(m, int32(168488), v29)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L5
	} else {
		goto L44
	}
L41:
	;
	v124 = base.I64_trunc_f64_s(v120)
	v126 = v124
	goto L40
L42:
	;
	goto L43
L43:
	;
	v126 = int64(-9223372036854775807 - 1)
	goto L40
L44:
	;
	F_errdetail(m, int32(648841), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L5
	} else {
		goto L45
	}
L45:
	;
	F_errhint(m, int32(616159), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(511686), int32(542), int32(393329))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	goto L39
L48:
	;
	goto L35
L49:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v152 = F_HnswInsertTupleOnDisk(m, l0, v45, v150, l1, int32(1))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	if v152 != 0 {
		goto L18
	} else {
		goto L51
	}
L51:
	;
	goto L4
L52:
	;
	v161 = F_HnswAlloc(m, v158, v79)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	F_LWLockRelease(m, v93)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	if v79 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if v167 != 0 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v166 = F__emscripten_memcpy_bulkmem(m, v161, v165, v79)
	mBase = m.M
	v167 = v166
	goto L58
L57:
	;
	v167 = v161
	goto L58
L58:
	;
	goto L55
L59:
	;
	v172 = v167 - v40 + int32(1)
	goto L61
L60:
	;
	v172 = int32(0)
	goto L61
L61:
	;
	if v40 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v173 = v172
	goto L64
L63:
	;
	v173 = v167
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159)+88)) = v173
	v176 = v159 + int32(92)
	v178 = *(*int32)(unsafe.Add(mBase, _consts[1491]))
	*(*uint16)(unsafe.Add(mBase, uint32(v176))) = uint16(v178)
	*(*int32)(unsafe.Add(mBase, uint32(v176)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v176)+8)) = int64(-1)
	goto L65
L65:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l5)+204))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l5)+160))
	v189 = v187 + int32(32)
	v191 = F_LWLockAcquire(m, v189, int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L5
	} else {
		goto L66
	}
L66:
	;
	F_LWLockRelease(m, v189)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L5
	} else {
		goto L67
	}
L67:
	;
	v196 = v187 + int32(16)
	v198 = F_LWLockAcquire(m, v196, int32(1))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L5
	} else {
		goto L68
	}
L68:
	;
	if v184 != 0 {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v239 = int32(0)
	F_HnswFindElementNeighbors(m, v184, v159, v238, v239, v45, v185, v186, v239)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L5
	} else {
		goto L86
	}
L70:
	;
	F_LWLockRelease(m, v196)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L5
	} else {
		goto L78
	}
L71:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+65)))
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+65)))
	if base.Ui32(v211) <= base.Ui32(v212) {
		v238 = v210
		goto L69
	} else {
		goto L77
	}
L72:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v187)+48))
	if v200 == int32(0) {
		goto L70
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v187)+48))
	if v206 == int32(0) {
		goto L70
	} else {
		goto L76
	}
L75:
	;
	v210 = v184 + v200 - int32(1)
	goto L71
L76:
	;
	v210 = v206
	goto L71
L77:
	;
	goto L70
L78:
	;
	v218 = int32(0)
	v220 = F_LWLockAcquire(m, v189, v218)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L5
	} else {
		goto L79
	}
L79:
	;
	v223 = F_LWLockAcquire(m, v196, int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L5
	} else {
		goto L80
	}
L80:
	;
	F_LWLockRelease(m, v189)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L5
	} else {
		goto L81
	}
L81:
	;
	if v184 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v187)+48))
	v238 = v229
	goto L69
L83:
	;
	goto L84
L84:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v187)+48))
	if v230 == int32(0) {
		v238 = v218
		goto L69
	} else {
		goto L85
	}
L85:
	;
	v238 = v230 + v184 - int32(1)
	goto L69
L86:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l5)+204))
	if v243 != 0 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l5)+160))
	v268 = v243 - int32(1)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	if v269 <= int32(0) {
		goto L97
	} else {
		goto L98
	}
L88:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v159)+72))
	v246 = int32(1)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v243+v244-v246)))
	if v248 != 0 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L90
L90:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v159)+88))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v159)+72))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v261)))
	v264 = v262
	v265 = v260
	goto L87
L91:
	;
	v253 = v243 + v248 - v246
	goto L93
L92:
	;
	v253 = int32(0)
	goto L93
L93:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v159)+88))
	if v254 == int32(0) {
		v264 = v253
		v265 = v7
		goto L87
	} else {
		goto L94
	}
L94:
	;
	v264 = v253
	v265 = v243 + v254 - int32(1)
	goto L87
L95:
	;
	F_LWLockRelease(m, v196)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L5
	} else {
		goto L166
	}
L96:
	;
	v552 = v159 + int32(4)
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+64)))
	v556 = v554 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v315)+64)) = uint8(v556)
	v560 = v315 + v554*int32(6)
	v561 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v552)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v560)+8)) = uint16(v561)
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v552)))
	*(*int32)(unsafe.Add(mBase, uint32(v560)+4)) = v563
	goto L164
L97:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	*(*int32)(unsafe.Add(mBase, uint32(v266))) = int32(1)
	if v362 != 0 {
		goto L114
	} else {
		goto L115
	}
L98:
	;
	v278 = int32(0)
	goto L99
L99:
	;
	v303 = v264 + int32(8) + v278*int32(12)
	if v243 != 0 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	goto L97
L101:
	;
	v318 = F_datumIsEqual(m, v265, v314, int32(0), int32(-1))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L5
	} else {
		goto L108
	}
L102:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v303)))
	v305 = v243 + v304
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)+87))
	if v306 != 0 {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	goto L104
L104:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v303)))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)+88))
	v314 = v313
	v315 = v312
	goto L101
L105:
	;
	v309 = v268 + v306
	goto L107
L106:
	;
	v309 = int32(0)
	goto L107
L107:
	;
	v314 = v309
	v315 = v305 - int32(1)
	goto L101
L108:
	;
	if v318 == int32(0) {
		goto L97
	} else {
		goto L109
	}
L109:
	;
	v323 = v315 + int32(92)
	v325 = F_LWLockAcquire(m, v323, int32(0))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L5
	} else {
		goto L110
	}
L110:
	;
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+64)))
	if v327 != int32(10) {
		goto L96
	} else {
		goto L111
	}
L111:
	;
	F_LWLockRelease(m, v323)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L5
	} else {
		goto L112
	}
L112:
	;
	v333 = v278 + int32(1)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	if v333 < v334 {
		v278 = v333
		goto L99
	} else {
		goto L113
	}
L113:
	;
	goto L100
L114:
	;
	F_s_lock(m, v266, int32(511686), int32(372), int32(14183))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L5
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v370
	*(*int32)(unsafe.Add(mBase, uint32(v266))) = int32(0)
	v376 = v159 - v243 + int32(1)
	if v243 != 0 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	goto L116
L118:
	;
	v377 = v376
	goto L120
L119:
	;
	v377 = v159
	goto L120
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v266)+4)) = v377
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+65)))
	v392 = v379
	goto L121
L121:
	;
	v410 = v185 << (uint(base.B2i32(v392 == int32(0))) % 32)
	v411 = F_mul_size(m, int32(12), v410)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L5
	} else {
		goto L123
	}
L122:
	;
	if v238 != 0 {
		goto L157
	} else {
		goto L158
	}
L123:
	;
	v413 = F_add_size(m, int32(8), v411)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L5
	} else {
		goto L124
	}
L124:
	;
	v415 = F_palloc(m, v413)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L5
	} else {
		goto L125
	}
L125:
	;
	v418 = F_LWLockAcquire(m, v176, int32(1))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L5
	} else {
		goto L126
	}
L126:
	;
	if v243 != 0 {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	if v413 != 0 {
		goto L135
	} else {
		goto L136
	}
L128:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v159)+72))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v268+v420+v392<<(uint(int32(2))%32))))
	if v425 != 0 {
		goto L131
	} else {
		goto L132
	}
L129:
	;
	goto L130
L130:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v159)+72))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v429+v392<<(uint(int32(2))%32))))
	v435 = v433
	goto L127
L131:
	;
	v428 = v268 + v425
	goto L133
L132:
	;
	v428 = int32(0)
	goto L133
L133:
	;
	v435 = v428
	goto L127
L134:
	;
	F_LWLockRelease(m, v176)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L5
	} else {
		goto L138
	}
L135:
	;
	v436 = F__emscripten_memcpy_bulkmem(m, v415, v435, v413)
	mBase = m.M
	v437 = v436
	goto L137
L136:
	;
	v437 = v415
	goto L137
L137:
	;
	goto L134
L138:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	if int32(0) < v440 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v444 = v392 << (uint(int32(2)) % 32)
	v452 = int32(0)
	goto L142
L140:
	;
	goto L141
L141:
	;
	if int32(0) < v392 {
		v392 = v392 - int32(1)
		goto L121
	} else {
		goto L156
	}
L142:
	;
	v477 = v437 + int32(8) + v452*int32(12)
	if v243 != 0 {
		goto L145
	} else {
		goto L146
	}
L143:
	;
	goto L141
L144:
	;
	v503 = *(*float32)(unsafe.Add(mBase, uint32(v477)+4))
	v504 = int32(0)
	F_HnswUpdateConnection(m, v243, v502, v159, v503, v410, v504, v504, v45)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L5
	} else {
		goto L153
	}
L145:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v477)))
	v479 = v243 + v478
	v481 = v479 + int32(91)
	v483 = F_LWLockAcquire(m, v481, int32(0))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L5
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v477)))
	v493 = v491 + int32(92)
	v495 = F_LWLockAcquire(m, v493, int32(0))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L5
	} else {
		goto L152
	}
L148:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v479)+71))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v268+v444+v485)))
	if v487 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v490 = v268 + v487
	goto L151
L150:
	;
	v490 = int32(0)
	goto L151
L151:
	;
	v500 = v481
	v502 = v490
	goto L144
L152:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v491)+72))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v497+v444)))
	v500 = v493
	v502 = v499
	goto L144
L153:
	;
	F_LWLockRelease(m, v500)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L5
	} else {
		goto L154
	}
L154:
	;
	v511 = v452 + int32(1)
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	if v511 < v512 {
		v452 = v511
		goto L142
	} else {
		goto L155
	}
L155:
	;
	goto L143
L156:
	;
	goto L122
L157:
	;
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+65)))
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+65)))
	if base.Ui32(v544) <= base.Ui32(v545) {
		goto L95
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	if v243 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	goto L159
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v266)+48)) = v159
	goto L95
L162:
	;
	goto L163
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v266)+48)) = v376
	goto L95
L164:
	;
	F_LWLockRelease(m, v323)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L5
	} else {
		goto L165
	}
L165:
	;
	goto L95
L166:
	;
	F_LWLockRelease(m, v51)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L5
	} else {
		goto L167
	}
L167:
	;
	goto L18
L168:
	;
	F_s_lock(m, v34, int32(511686), int32(601), int32(326028))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L5
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v631 = *(*float64)(unsafe.Add(mBase, uint32(v34)+8))
	v633 = base.F64_add(v631, float64(1))
	*(*float64)(unsafe.Add(mBase, uint32(v34)+8)) = v633
	if base.F64_lt(base.F64_abs(v633), float64(9.223372036854776e+18)) != 0 {
		goto L173
	} else {
		goto L174
	}
L171:
	;
	goto L170
L172:
	;
	v644 = *(*int32)(unsafe.Add(mBase, _consts[49]))
	if v644 == int32(0) {
		goto L177
	} else {
		goto L178
	}
L173:
	;
	v639 = base.I64_trunc_f64_s(v633)
	v641 = v639
	goto L172
L174:
	;
	goto L175
L175:
	;
	v641 = int64(-9223372036854775807 - 1)
	goto L172
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(0)
	goto L4
L177:
	;
	goto L176
L178:
	;
	v648 = int32(*(*uint8)(unsafe.Add(mBase, _consts[50])))
	if v648 != int32(1) {
		goto L177
	} else {
		goto L179
	}
L179:
	;
	v651 = int32(4530932)
	v653 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	v654 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v653 + v654
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v644)))
	*(*int32)(unsafe.Add(mBase, uint32(v644))) = v657 + v654
	*(*int64)(unsafe.Add(mBase, uint32(v644+int32(96))+232)) = v641
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v644)))
	*(*int32)(unsafe.Add(mBase, uint32(v644))) = v665 + v654
	v671 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v671 - v654
	goto L177
L180:
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
	F_errmsg_internal(m, int32(40646), v11)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(503210), int32(2705), int32(247171))
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
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(373689)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v10
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(993)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v14
	v18 = int32(4529176)
	v19 = *(*int32)(unsafe.Add(mBase, _consts[142]))
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v8 + int32(4)
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
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(516218)
				F_errmsg(m, int32(325489), v8)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					F_errfinish(m, int32(509569), int32(941), int32(222444))
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
			*(*int32)(unsafe.Add(mBase, _consts[142])) = v59
			m.G0 = v8 + int32(32)
			return
		}
	}
}
func F_beginmerge(m *base.Module, l0 int32) {
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
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int64
	_ = v97
	var v99 int64
	_ = v99
	var v111 int32
	_ = v111
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v14 < v15 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L10
	} else {
		goto L31
	}
L2:
	;
	v17 = v14
	goto L4
L3:
	;
	v17 = v15
	goto L4
L4:
	;
	if int32(0) < v17 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v24 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	m.G0 = v12 + int32(16)
	return
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+v24<<(uint(int32(2))%32))))
	v35 = F_LogicalTapeRead(m, v33, v12, int32(4))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
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
	if v35 != int32(4) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v39 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	m.T0[v40].(func(*base.Module, int32, int32, int32, int32))(m, l0, v12, v33, v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L10
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v111 = v24 + int32(1)
	if v111 != v17 {
		v24 = v111
		goto L8
	} else {
		goto L30
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v24
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v46 = *(*int32)(unsafe.Add(mBase, _consts[45]))
	if v46 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L10
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v49 + int32(1)
	if v49 <= int32(0) {
		v90 = v49
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L19
L21:
	;
	v96 = v44 + v90<<(uint(int32(4))%32)
	v97 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	*(*int64)(unsafe.Add(mBase, uint32(v96))) = v97
	v99 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = v99
	goto L15
L22:
	;
	v56 = v49
	goto L23
L23:
	;
	v64 = int32(1)
	v65 = v56 - v64
	v67 = int32(base.Ui32(v65) >> (uint(v64) % 32))
	v70 = v44 + v67<<(uint(int32(4))%32)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v72 = m.T0[v71].(func(*base.Module, int32, int32, int32) int32)(m, v12, v70, l0)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L10
	} else {
		goto L25
	}
L24:
	;
	v90 = v67
	goto L21
L25:
	;
	if int32(0) <= v72 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v90 = v56
	goto L21
L27:
	;
	goto L28
L28:
	;
	v78 = v44 + v56<<(uint(int32(4))%32)
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v70)))
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = v79
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v70)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v78)+8)) = v81
	if base.Ui32(int32(1)) < base.Ui32(v65) {
		v56 = v67
		goto L23
	} else {
		goto L29
	}
L29:
	;
	goto L24
L30:
	;
	goto L9
L31:
	;
	F_errmsg_internal(m, int32(380290), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L10
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(503582), int32(2862), int32(287620))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L10
	} else {
		goto L33
	}
L33:
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
	var v35 float64
	_ = v35
	var v37 int32
	_ = v37
	var v39 float64
	_ = v39
	var v40 float64
	_ = v40
	var v42 float64
	_ = v42
	var v50 float64
	_ = v50
	var v54 float64
	_ = v54
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
			v35 = v8
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+24)))
			if v16 != 0 {
				v35 = v8
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
				v18 = base.F32_reinterpret_i32(v17)
				if base.F32_ge(v18, float32(0)) == int32(0) {
					v35 = v8
				} else {
					if base.F32_le(v18, float32(100)) == int32(0) {
						v35 = v8
					} else {
						if base.Ui32(int32(2139095040)) < base.Ui32(v17&int32(2147483647)) {
							v35 = v8
						} else {
							v35 = base.F64_promote_f32(base.F32_div(v18, float32(100)))
						}
					}
				}
			}
		}
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = v37
		v39 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
		v40 = base.F64_mul(v35, v39)
		v42 = float64(1e+100)
		if base.F64_gt(v40, v42) != 0 {
			v54 = v42
		} else {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v40)&int64(9223372036854775807)) {
				v54 = v42
			} else {
				v50 = float64(1)
				if base.F64_le(v40, v50) != 0 {
					v54 = v50
				} else {
					v54 = base.F64_nearest(v40)
				}
			}
		}
		*(*float64)(unsafe.Add(mBase, uint32(l4))) = v54
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+92))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+84)) = l2
	if l2 != 0 {
		v12 = F__emscripten_memcpy_bulkmem(m, v6+int32(4), l1, l2)
		mBase = m.M
	} else {
	}
	v15 = v6 + int32(68)
	if l3 != 0 {
		if v8 != 0 {
			v16 = F__emscripten_memcpy_bulkmem(m, v15, l3, v8)
			mBase = m.M
		} else {
		}
		return int32(0)
	} else {
		v22 = F__emscripten_memset_bulkmem(m, v15, base.I32_extend8_s(int32(0)), v8)
		mBase = m.M
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
	v53 = (v48 | v37<<(uint(int32(8))%32)) & int32(65535)
	v55 = v12 + int32(15)
	if base.Ui32(v53) <= base.Ui32(int32(51519)) {
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
	v117 = v115 & int32(65535)
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
	v97 = int32(2253800)
	goto L23
L25:
	;
	v97 = int32(2253796)
	goto L23
L26:
	;
	v97 = int32(2253792)
	goto L23
L27:
	;
	v97 = int32(2253788)
	goto L23
L28:
	;
	v97 = int32(2253784)
	goto L23
L29:
	;
	v97 = int32(2253780)
	goto L23
L30:
	;
	v87 = F_BinarySearchRange(m, int32(2253808), int32(46), v53)
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
	switch v53 - int32(51321) {
	case 0:
		v70 = int32(2253648)
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
	switch v53 - int32(63958) {
	case 0:
		v97 = int32(2253776)
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
	v70 = int32(2253660)
	goto L35
L37:
	;
	v70 = int32(2253656)
	goto L35
L38:
	;
	v70 = int32(2253652)
	goto L35
L39:
	;
	v66 = F_BinarySearchRange(m, int32(2253664), int32(23), v53)
	mBase = m.M
	if v66 != 0 {
		v82 = v66
		goto L31
	} else {
		goto L42
	}
L40:
	;
	if v53 == int32(51362) {
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
	if v53 != int32(51530) {
		goto L30
	} else {
		goto L44
	}
L44:
	;
	v82 = int32(17474)
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
	v122 = int32(41614)
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
	var v19 int32
	_ = v19
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
	v19 = int32(0)
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
	v24 = v19 << (uint(v23) % 32)
	v26 = v24 | v23
	v28 = v24 + int32(2)
	if v28 < v17 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+v19<<(uint(int32(2))%32)))) = l1
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
	*(*int32)(unsafe.Add(mBase, uint32(v14+v19<<(uint(int32(2))%32)))) = v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = v64
	v19 = v46
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
		*(*int32)(unsafe.Add(mBase, _consts[87])) = int32(0) - v7
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
	var v50 int32
	_ = v50
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
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 float64
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v276 int32
	_ = v276
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
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v316 int32
	_ = v316
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
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v347 int32
	_ = v347
	var v355 int32
	_ = v355
	var v367 int32
	_ = v367
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	v5 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(9184)
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
	v50 = int32(1)
	goto L11
L9:
	;
	v367 = v5
	goto L10
L10:
	;
	v375 = F_ReadBuffer(m, v22, int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L4
	} else {
		goto L102
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
	v367 = v347
	goto L10
L13:
	;
	v59 = int32(0)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v62 = F_ReadBufferExtended(m, v22, v59, v50, v59, v61)
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
	v355 = v50 + int32(1)
	if v355 != v34 {
		v49 = v347
		v50 = v355
		goto L11
	} else {
		goto L101
	}
L17:
	;
	v87 = v70 + int32(24)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v20)+1164))
	v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74))))
	v90 = int32(1)
	v96 = v88 * ((v89+v90)&int32(65535) - v90)
	if v96 <= int32(0) {
		goto L27
	} else {
		goto L28
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
	v83 = m.ExcPending
	if v83 != 0 {
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
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v347 = v49
	goto L16
L26:
	;
	if v300 != 0 {
		goto L90
	} else {
		goto L91
	}
L27:
	;
	v289 = v87
	v290 = v87
	v292 = v88
	v297 = v73
	v300 = v89
	goto L26
L28:
	;
	goto L29
L29:
	;
	v104 = v87
	v105 = v87
	goto L30
L30:
	;
	v117 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, v104, l3)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L33
	}
L31:
	;
	v282 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+16)))
	v284 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70+v282))))
	v289 = v280
	v290 = v278
	v292 = v279
	v297 = v282
	v300 = v284
	goto L26
L32:
	;
	v280 = v104 + v279
	if base.Ui32(v280) < base.Ui32(v96+v87) {
		v104 = v280
		v105 = v278
		goto L30
	} else {
		goto L86
	}
L33:
	;
	if v117 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+16)))
	v120 = v70 + v119
	v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v120))))
	v123 = v121 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v120))) = uint16(v123)
	v125 = *(*float64)(unsafe.Add(mBase, uint32(v30)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v30)+16)) = base.F64_add(v125, float64(1))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v20)+1164))
	v278 = v105
	v279 = v129
	goto L32
L35:
	;
	goto L36
L36:
	;
	if v104 != v105 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v20)+1164))
	if v105 == v104 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	goto L39
L39:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v20)+1164))
	v278 = v105 + v276
	v279 = v276
	goto L32
L40:
	;
	goto L39
L41:
	;
	goto L40
L42:
	;
	v135 = v105 + v131
	if base.Ui32(v104-v135) <= base.Ui32(int32(0)-v131<<(uint(int32(1))%32)) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v142 = F___memcpy(m, v105, v104, v131)
	mBase = m.M
	goto L40
L44:
	;
	goto L45
L45:
	;
	v145 = (v105 ^ v104) & int32(3)
	if base.Ui32(v105) < base.Ui32(v104) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	if v247 == int32(0) {
		goto L41
	} else {
		goto L82
	}
L47:
	;
	if base.Ui32(v225) <= base.Ui32(int32(3)) {
		v246 = v224
		v247 = v225
		v248 = v226
		goto L46
	} else {
		goto L78
	}
L48:
	;
	if v145 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	if v145 != 0 {
		v207 = v131
		goto L61
	} else {
		goto L62
	}
L51:
	;
	v246 = v104
	v247 = v131
	v248 = v105
	goto L46
L52:
	;
	goto L53
L53:
	;
	if v105&int32(3) == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v224 = v104
	v225 = v131
	v226 = v105
	goto L47
L55:
	;
	goto L56
L56:
	;
	v152 = v104
	v153 = v131
	v154 = v105
	goto L57
L57:
	;
	if v153 == int32(0) {
		goto L41
	} else {
		goto L59
	}
L58:
	;
	v224 = v161
	v225 = v163
	v226 = v165
	goto L47
L59:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
	*(*uint8)(unsafe.Add(mBase, uint32(v154))) = uint8(v158)
	v160 = int32(1)
	v161 = v152 + v160
	v163 = v153 - v160
	v165 = v154 + v160
	if v165&int32(3) != 0 {
		v152 = v161
		v153 = v163
		v154 = v165
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	if v207 == int32(0) {
		goto L41
	} else {
		goto L74
	}
L62:
	;
	if v135&int32(3) != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v172 = v131
	goto L66
L64:
	;
	v187 = v131
	goto L65
L65:
	;
	if base.Ui32(v187) <= base.Ui32(int32(3)) {
		v207 = v187
		goto L61
	} else {
		goto L70
	}
L66:
	;
	if v172 == int32(0) {
		goto L41
	} else {
		goto L68
	}
L67:
	;
	v187 = v178
	goto L65
L68:
	;
	v178 = v172 - int32(1)
	v179 = v105 + v178
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104+v178))))
	*(*uint8)(unsafe.Add(mBase, uint32(v179))) = uint8(v181)
	if v179&int32(3) != 0 {
		v172 = v178
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v194 = v187
	goto L71
L71:
	;
	v198 = v194 - int32(4)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v104+v198)))
	*(*int32)(unsafe.Add(mBase, uint32(v105+v198))) = v201
	if base.Ui32(int32(3)) < base.Ui32(v198) {
		v194 = v198
		goto L71
	} else {
		goto L73
	}
L72:
	;
	v207 = v198
	goto L61
L73:
	;
	goto L72
L74:
	;
	v214 = v207
	goto L75
L75:
	;
	v218 = v214 - int32(1)
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104+v218))))
	*(*uint8)(unsafe.Add(mBase, uint32(v105+v218))) = uint8(v221)
	if v218 != 0 {
		v214 = v218
		goto L75
	} else {
		goto L77
	}
L76:
	;
	goto L41
L77:
	;
	goto L76
L78:
	;
	v231 = v224
	v232 = v225
	v233 = v226
	goto L79
L79:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	*(*int32)(unsafe.Add(mBase, uint32(v233))) = v235
	v237 = int32(4)
	v238 = v231 + v237
	v240 = v233 + v237
	v242 = v232 - v237
	if base.Ui32(int32(3)) < base.Ui32(v242) {
		v231 = v238
		v232 = v242
		v233 = v240
		goto L79
	} else {
		goto L81
	}
L80:
	;
	v246 = v238
	v247 = v242
	v248 = v240
	goto L46
L81:
	;
	goto L80
L82:
	;
	v253 = v246
	v254 = v247
	v255 = v248
	goto L83
L83:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253))))
	*(*uint8)(unsafe.Add(mBase, uint32(v255))) = uint8(v257)
	v259 = int32(1)
	v264 = v254 - v259
	if v264 != 0 {
		v253 = v253 + v259
		v254 = v264
		v255 = v255 + v259
		goto L83
	} else {
		goto L85
	}
L84:
	;
	goto L41
L85:
	;
	goto L84
L86:
	;
	goto L31
L87:
	;
	F_UnlockReleaseBuffer(m, v62)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L4
	} else {
		goto L100
	}
L88:
	;
	F_pfree(m, v67)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L4
	} else {
		goto L99
	}
L89:
	;
	v326 = v290 - v70
	*(*uint16)(unsafe.Add(mBase, uint32(v70)+12)) = uint16(v326)
	F_GenericXLogFinish(m, v67)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L4
	} else {
		goto L98
	}
L90:
	;
	if base.Ui32(int32(8160)-v292*v300) < base.Ui32(v292) {
		v316 = v49
		goto L93
	} else {
		goto L94
	}
L91:
	;
	goto L92
L92:
	;
	if v289 == v290 {
		v330 = v49
		goto L88
	} else {
		goto L97
	}
L93:
	;
	if v289 != v290 {
		v325 = v316
		goto L89
	} else {
		goto L96
	}
L94:
	;
	if base.Ui32(int32(2003)) < base.Ui32(v49) {
		v316 = v49
		goto L93
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(1168)+v49<<(uint(int32(2))%32)))) = v50
	v316 = v49 + int32(1)
	goto L93
L96:
	;
	v330 = v316
	goto L88
L97:
	;
	v319 = v70 + v297
	v320 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v319)+2)))
	v322 = v320 | int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v319)+2)) = uint16(v322)
	v325 = v49
	goto L89
L98:
	;
	v334 = v325
	goto L87
L99:
	;
	v334 = v330
	goto L87
L100:
	;
	v347 = v334
	goto L16
L101:
	;
	goto L12
L102:
	;
	F_LockBuffer(m, v375, int32(2))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	v380 = F_GenericXLogStart(m, v22)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L4
	} else {
		goto L104
	}
L104:
	;
	v383 = F_GenericXLogRegisterBuffer(m, v380, v375, int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	v390 = v367 << (uint(int32(2)) % 32)
	if v390 != 0 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v383)+30)) = uint16(v367)
	v394 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v383)+28)) = uint16(v394)
	F_GenericXLogFinish(m, v380)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L4
	} else {
		goto L110
	}
L107:
	;
	v391 = F__emscripten_memcpy_bulkmem(m, v383+int32(168), v20+int32(1168), v390)
	mBase = m.M
	goto L109
L108:
	;
	goto L109
L109:
	;
	goto L106
L110:
	;
	F_UnlockReleaseBuffer(m, v375)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L4
	} else {
		goto L111
	}
L111:
	;
	m.G0 = v20 + int32(9184)
	return v30
}
func F_blcostestimate(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 float64
	_ = v47
	var v50 int32
	_ = v50
	var v51 float64
	_ = v51
	var v53 float64
	_ = v53
	var v55 float64
	_ = v55
	var v57 float64
	_ = v57
	var v59 float64
	_ = v59
	v16 = m.G0
	v18 = v16 + int32(-64)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v21 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v18)+48)) = v21
	v26 = v16 + int32(-24)
	*(*int64)(unsafe.Add(mBase, uint32(v26))) = v21
	v30 = v16 + int32(-32)
	*(*int64)(unsafe.Add(mBase, uint32(v30))) = v21
	v34 = v16 + int32(-40)
	*(*int64)(unsafe.Add(mBase, uint32(v34))) = v21
	v38 = v16 + int32(-48)
	*(*int64)(unsafe.Add(mBase, uint32(v38))) = v21
	v42 = v16 + int32(-56)
	*(*int64)(unsafe.Add(mBase, uint32(v42))) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v21
	v47 = *(*float64)(unsafe.Add(mBase, uint32(v20)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v26))) = v47
	F_genericcostestimate(m, l0, l1, l2, v18)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		return
	} else {
		v51 = *(*float64)(unsafe.Add(mBase, uint32(v18)))
		*(*float64)(unsafe.Add(mBase, uint32(l3))) = v51
		v53 = *(*float64)(unsafe.Add(mBase, uint32(v42)))
		*(*float64)(unsafe.Add(mBase, uint32(l4))) = v53
		v55 = *(*float64)(unsafe.Add(mBase, uint32(v38)))
		*(*float64)(unsafe.Add(mBase, uint32(l5))) = v55
		v57 = *(*float64)(unsafe.Add(mBase, uint32(v34)))
		*(*float64)(unsafe.Add(mBase, uint32(l6))) = v57
		v59 = *(*float64)(unsafe.Add(mBase, uint32(v30)))
		*(*float64)(unsafe.Add(mBase, uint32(l7))) = v59
		m.G0 = v18 - int32(-64)
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
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(16384)
			v15 = F_palloc(m, int32(16386))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v15
				if v15 == int32(0) {
					F_yy_fatal_error_1(m, int32(697879))
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
					v23 = *(*int32)(unsafe.Add(mBase, _consts[87]))
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
					*(*int32)(unsafe.Add(mBase, _consts[87])) = v23
					return v8
				}
			}
		} else {
			F_yy_fatal_error_1(m, int32(697879))
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
				F_yy_fatal_error_1(m, int32(698156))
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
					F_yy_fatal_error_1(m, int32(698156))
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
					*(*int64)(unsafe.Add(mBase, uint32(v34))) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+24)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+16)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+8)) = v35
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
			v21 = F_pg_snprintf(m, v10, int32(64), int32(694208), v6)
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	if l1 == int32(0) {
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v8 <= int32(0) {
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v13 = v8 * int32(48)
			if v13 != 0 {
				v14 = F__emscripten_memcpy_bulkmem(m, v11, l1, v13)
				mBase = m.M
			} else {
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
	var v55 int32
	_ = v55
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
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int64
	_ = v206
	var v210 int32
	_ = v210
	var v211 int64
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v244 int64
	_ = v244
	var v245 int32
	_ = v245
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v265 int64
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v319 int64
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v396 int32
	_ = v396
	var v406 int64
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
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
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
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
	var v603 int32
	_ = v603
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
	var v667 int32
	_ = v667
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
	var v732 int32
	_ = v732
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
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
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
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
	var v852 int32
	_ = v852
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
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v37)+40))
	if int32(0) < v245 {
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
	v219 = m.ExcPending
	if v219 != 0 {
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
		v244 = v15
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v54 = int32(0)
	v55 = v39
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
	v233 = v210
	v244 = v211
	goto L1
L13:
	;
	v213 = v54 + int32(1)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v44)+40))
	if v213 < v214 {
		v54 = v213
		v55 = v210
		v66 = v211
		goto L11
	} else {
		goto L45
	}
L14:
	;
	v206 = int64(*(*int16)(unsafe.Add(mBase, uint32(v73)+16)))
	v210 = v204
	v211 = int64(1)<<(uint(v206)%64) | v66
	goto L13
L15:
	;
	v174 = int32(0)
	v177 = F_errstart(m, int32(17), v174)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
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
	if base.Ui32(int32(65530)) < base.Ui32((v74-int32(16))&int32(65535)) {
		v204 = v55
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
	v204 = v55
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
	v204 = v55
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
	v204 = v55
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
	v204 = v55
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
		v210 = v133
		v211 = v66
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
	*(*int32)(unsafe.Add(mBase, uint32(v19)+180)) = int32(281010)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+176)) = v31
	F_errmsg(m, int32(481682), v19+int32(176))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(510632), int32(114), int32(364877))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	v210 = v133
	v211 = v66
	goto L13
L37:
	;
	if v171 != 0 {
		v204 = v55
		goto L14
	} else {
		goto L38
	}
L38:
	;
	goto L15
L39:
	;
	if v177 == int32(0) {
		v204 = v174
		goto L14
	} else {
		goto L40
	}
L40:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	v185 = F_format_procedure(m, v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v187 = int32(*(*int16)(unsafe.Add(mBase, uint32(v73)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+204)) = v187
	*(*int32)(unsafe.Add(mBase, uint32(v19)+200)) = v185
	*(*int32)(unsafe.Add(mBase, uint32(v19)+196)) = int32(281010)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+192)) = v31
	F_errmsg(m, int32(481507), v19+int32(192))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(510632), int32(130), int32(364877))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	v204 = v174
	goto L14
L45:
	;
	goto L12
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = l0
	F_errmsg_internal(m, int32(43351), v19)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(510632), int32(58), int32(364877))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
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
	v254 = int32(0)
	v255 = v233
	v265 = v15
	goto L52
L50:
	;
	v396 = v233
	v406 = v15
	goto L51
L51:
	;
	v409 = v28 + int32(8)
	v410 = int32(1)
	v411 = F_identify_opfamily_groups(m, v37, v44)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L2
	} else {
		goto L97
	}
L52:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(48)+v254<<(uint(int32(2))%32))))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)+56))
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+22)))
	v273 = v271 + v272
	v274 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v273)+16)))
	if base.Ui32((v274+int32(-64))&int32(65535)) <= base.Ui32(int32(65472)) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v396 = v387
	v406 = v319
	goto L51
L54:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273)+18)))
	if v320 == int32(115) {
		goto L66
	} else {
		goto L67
	}
L55:
	;
	v281 = int32(0)
	v284 = F_errstart(m, int32(17), v281)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L2
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v273)+8))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v273)+12))
	if v310 != v311 {
		v318 = v255
		v319 = v265
		goto L54
	} else {
		goto L64
	}
L58:
	;
	if v284 == int32(0) {
		v318 = v281
		v319 = v265
		goto L54
	} else {
		goto L59
	}
L59:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L2
	} else {
		goto L60
	}
L60:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v273)+20))
	v292 = F_format_operator(m, v291)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	v294 = int32(*(*int16)(unsafe.Add(mBase, uint32(v273)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+172)) = v294
	*(*int32)(unsafe.Add(mBase, uint32(v19)+168)) = v292
	*(*int32)(unsafe.Add(mBase, uint32(v19)+164)) = int32(281010)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+160)) = v31
	F_errmsg(m, int32(481301), v19+int32(160))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L2
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(510632), int32(152), int32(364877))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	v318 = v281
	v319 = v265
	goto L54
L64:
	;
	v318 = v255
	v319 = int64(1)<<(uint(base.I64_extend_i32_u(v274))%64) | v265
	goto L54
L65:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v273)+20))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v273)+8))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v273)+12))
	v358 = F_check_amop_signature(m, v354, int32(16), v356, v357)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L2
	} else {
		goto L77
	}
L66:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v273)+28))
	if v323 == int32(0) {
		v353 = v318
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v326 = int32(0)
	v329 = F_errstart(m, int32(17), v326)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L2
	} else {
		goto L70
	}
L69:
	;
	goto L68
L70:
	;
	if v329 == int32(0) {
		v353 = v326
		goto L65
	} else {
		goto L71
	}
L71:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v273)+20))
	v337 = F_format_operator(m, v336)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+152)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v19)+148)) = int32(281010)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+144)) = v31
	F_errmsg(m, int32(184991), v19+int32(144))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L2
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(510632), int32(180), int32(364877))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L2
	} else {
		goto L75
	}
L75:
	;
	v353 = v326
	goto L65
L76:
	;
	v389 = v254 + int32(1)
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v37)+40))
	if v389 < v390 {
		v254 = v389
		v255 = v387
		v265 = v319
		goto L52
	} else {
		goto L85
	}
L77:
	;
	if v358 != 0 {
		v387 = v353
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v360 = int32(0)
	v363 = F_errstart(m, int32(17), v360)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L2
	} else {
		goto L79
	}
L79:
	;
	if v363 == int32(0) {
		v387 = v360
		goto L76
	} else {
		goto L80
	}
L80:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L2
	} else {
		goto L81
	}
L81:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v273)+20))
	v371 = F_format_operator(m, v370)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+136)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v19)+132)) = int32(281010)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+128)) = v31
	F_errmsg(m, int32(371318), v19+int32(128))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L2
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(510632), int32(193), int32(364877))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	v387 = v360
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
		v852 = v792
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
		v792 = v732
		v796 = v37
		v797 = v409
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
		v732 = v667
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
		v667 = v603
		goto L92
	} else {
		goto L152
	}
L95:
	;
	v599 = v532 + int32(16)
	v603 = v531
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
	if v411 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v554 = int32(0)
	v560 = v410
	goto L96
L99:
	;
	goto L100
L100:
	;
	v416 = int32(0)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v411)+4))
	if v416 < v417 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v424 = int32(0)
	v425 = v396
	v426 = v416
	goto L104
L102:
	;
	v531 = v396
	v532 = v416
	goto L103
L103:
	;
	if v532 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L104:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v411)+12))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v437+v424<<(uint(int32(2))%32))))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v441)))
	if v29 == v442 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v531 = v521
	v532 = v447
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
	v447 = v426
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
	v446 = v426
	goto L111
L111:
	;
	v447 = v446
	goto L108
L112:
	;
	v524 = v424 + int32(1)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v411)+4))
	if v524 < v525 {
		v424 = v524
		v425 = v521
		v426 = v447
		goto L104
	} else {
		goto L134
	}
L113:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v441)+4))
	if v442 != v451 {
		v521 = v425
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
	if v453 == v406 {
		v486 = v425
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
	if v488 == v244 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v19)+116)) = int32(281010)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+112)) = v31
	F_errmsg(m, int32(200523), v19+int32(112))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L2
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(510632), int32(232), int32(364877))
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
	*(*int32)(unsafe.Add(mBase, uint32(v19)+100)) = int32(281010)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v31
	F_errmsg(m, int32(200607), v19+int32(96))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L2
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(510632), int32(242), int32(364877))
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
	v554 = int32(0)
	v560 = v410
	goto L96
L136:
	;
	goto L137
L137:
	;
	v546 = *(*int64)(unsafe.Add(mBase, uint32(v532)+8))
	if v546 == v406 {
		goto L95
	} else {
		goto L138
	}
L138:
	;
	v554 = v532
	v560 = int32(0)
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
	v592 = v554 + int32(16)
	if v560 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+84)) = int32(281010)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v409
	F_errmsg(m, int32(688405), v19+int32(80))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L2
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(510632), int32(253), int32(364877))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L2
	} else {
		goto L145
	}
L145:
	;
	v586 = v554 + int32(16)
	if v560 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v599 = v586
	v603 = int32(0)
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
	v603 = int32(0)
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
		v667 = v636
		goto L92
	} else {
		goto L161
	}
L157:
	;
	v644 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = int32(281010)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v409
	F_errmsg(m, int32(482809), v19-int32(-64))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L2
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(510632), int32(264), int32(364877))
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
	v667 = v636
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
		v732 = v699
		goto L90
	} else {
		goto L171
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = int32(281010)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v409
	F_errmsg(m, int32(482809), v19+int32(48))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L2
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(510632), int32(264), int32(364877))
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
		v732 = v699
		goto L90
	} else {
		goto L170
	}
L170:
	;
	v748 = v683
	v749 = v19
	v756 = v37
	v757 = v409
	v758 = v44
	v760 = v22
	v764 = int32(1)
	goto L89
L171:
	;
	v748 = v683
	v749 = v19
	v756 = v37
	v757 = v409
	v758 = v44
	v760 = v22
	v764 = int32(1)
	goto L89
L172:
	;
	v748 = v728
	v749 = v19
	v756 = v37
	v757 = v409
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
		v815 = v756
		v816 = v757
		v817 = v758
		v819 = v760
		goto L87
	} else {
		goto L180
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v749)+40)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v749)+36)) = int32(281010)
	*(*int32)(unsafe.Add(mBase, uint32(v749)+32)) = v757
	F_errmsg(m, int32(482809), v749+int32(32))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L2
	} else {
		goto L178
	}
L178:
	;
	F_errfinish(m, int32(510632), int32(264), int32(364877))
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
	v792 = v765
	v796 = v756
	v797 = v757
	v798 = v758
	v800 = v760
	goto L88
L181:
	;
	v808 = v789
	v815 = v796
	v816 = v797
	v817 = v798
	v819 = v800
	goto L87
L182:
	;
	if v826 == int32(0) {
		v849 = v808
		v852 = v823
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
	*(*int32)(unsafe.Add(mBase, uint32(v808)+20)) = int32(281010)
	*(*int32)(unsafe.Add(mBase, uint32(v808)+16)) = v816
	F_errmsg(m, int32(482809), v808+int32(16))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L2
	} else {
		goto L185
	}
L185:
	;
	F_errfinish(m, int32(510632), int32(264), int32(364877))
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
	v852 = v823
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
	return v852
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	v36 = v15 - v34
	v37 = int32(1)
	if base.Ui32(v37) < base.Ui32(v36) {
		v14 = v35
		v15 = int32(base.Ui32(v36) >> (uint(v37) % 32))
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v61 int32
	_ = v61
	var v67 float64
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 float64
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 float64
	_ = v95
	var v97 float64
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v139 float64
	_ = v139
	var v141 float64
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 float64
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 float32
	_ = v161
	var v163 float32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v190 float64
	_ = v190
	var v191 float64
	_ = v191
	var v195 int32
	_ = v195
	var v196 float64
	_ = v196
	var v199 float64
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 float64
	_ = v206
	var v209 int32
	_ = v209
	var v216 float64
	_ = v216
	var v219 float64
	_ = v219
	var v220 int32
	_ = v220
	var v225 float64
	_ = v225
	var v226 int32
	_ = v226
	var v232 float64
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 float32
	_ = v252
	var v253 float64
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 float64
	_ = v257
	var v260 int32
	_ = v260
	var v263 float64
	_ = v263
	var v267 int32
	_ = v267
	var v268 float64
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 float64
	_ = v274
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 float64
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v358 int32
	_ = v358
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 float64
	_ = v410
	var v411 int32
	_ = v411
	var v415 float64
	_ = v415
	var v416 float64
	_ = v416
	var v419 float64
	_ = v419
	var v445 float64
	_ = v445
	var v451 float64
	_ = v451
	var v452 float64
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v518 float64
	_ = v518
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v547 float64
	_ = v547
	var v549 float64
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v579 float64
	_ = v579
	var v581 float64
	_ = v581
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v613 float64
	_ = v613
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 float64
	_ = v633
	var v634 int32
	_ = v634
	var v638 float64
	_ = v638
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v657 int32
	_ = v657
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 float64
	_ = v667
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 float64
	_ = v688
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v729 float64
	_ = v729
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 float64
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v762 float64
	_ = v762
	var v764 float64
	_ = v764
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 float64
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v793 float64
	_ = v793
	var v795 float64
	_ = v795
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v885 int32
	_ = v885
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v938 float64
	_ = v938
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v941 float64
	_ = v941
	var v943 int32
	_ = v943
	var v947 float64
	_ = v947
	var v949 float64
	_ = v949
	var v950 float64
	_ = v950
	var v953 float64
	_ = v953
	var v980 float64
	_ = v980
	var v985 float64
	_ = v985
	var v991 int32
	_ = v991
	var v993 float64
	_ = v993
	var v994 float64
	_ = v994
	var v995 float64
	_ = v995
	var v996 float64
	_ = v996
	var v1001 float64
	_ = v1001
	var v1002 float64
	_ = v1002
	var v1006 float64
	_ = v1006
	var v1009 float64
	_ = v1009
	var v1011 float64
	_ = v1011
	var v1013 float64
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1023 float64
	_ = v1023
	var v1024 float64
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1036 int32
	_ = v1036
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1044 float32
	_ = v1044
	var v1045 float64
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1049 float64
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1055 float64
	_ = v1055
	var v1059 int32
	_ = v1059
	var v1060 float64
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1068 float64
	_ = v1068
	var v1074 float64
	_ = v1074
	var v1080 float64
	_ = v1080
	var v1083 float64
	_ = v1083
	v9 = int32(0)
	v25 = float64(0)
	v30 = m.G0
	v32 = v30 - int32(176)
	m.G0 = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v35 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+128)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v32)+120)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v32)+112)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v32)+104)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v32)+96)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v32)+88)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v32)+80)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v32)+72)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v32-int32(-64)))) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v32)+56)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v32)+48)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v32)+40)) = v35
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v61 == v9 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v798 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v780)+101)))
	if v798 == int32(1) {
		goto L145
	} else {
		goto L146
	}
L2:
	;
	v769 = l0
	v770 = l1
	v771 = l2
	v772 = l3
	v773 = l4
	v774 = l5
	v775 = l6
	v776 = l7
	v777 = v32
	v778 = int32(1)
	v780 = v34
	v783 = v9
	v785 = v9
	v786 = v9
	v788 = v9
	v789 = v9
	v793 = float64(1)
	v795 = v25
	goto L1
L3:
	;
	goto L4
L4:
	;
	v67 = float64(1)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v68 <= int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v769 = l0
	v770 = l1
	v771 = l2
	v772 = l3
	v773 = l4
	v774 = l5
	v775 = l6
	v776 = l7
	v777 = v32
	v778 = int32(1)
	v780 = v34
	v783 = v9
	v785 = v9
	v786 = v9
	v788 = v9
	v789 = v9
	v793 = v67
	v795 = v25
	goto L1
L6:
	;
	goto L7
L7:
	;
	v71 = l0
	v72 = l1
	v73 = l2
	v74 = l3
	v75 = l4
	v76 = l5
	v77 = l6
	v78 = l7
	v79 = v32
	v81 = v9
	v82 = v34
	v83 = v9
	v85 = v9
	v87 = v9
	v88 = v9
	v90 = v9
	v91 = v9
	v92 = v61
	v93 = v9
	v94 = v9
	v95 = v67
	v97 = v25
	goto L8
L8:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100+v93<<(uint(int32(2))%32))))
	v105 = int32(*(*int16)(unsafe.Add(mBase, uint32(v104)+14)))
	if v105 <= v85 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v769 = v738
	v770 = v739
	v771 = v740
	v772 = v741
	v773 = v742
	v774 = v743
	v775 = v744
	v776 = v745
	v777 = v746
	v778 = v748 ^ int32(1)
	v780 = v749
	v783 = v752
	v785 = v754
	v786 = v755
	v788 = v757
	v789 = v758
	v793 = v762
	v795 = v764
	goto L1
L10:
	;
	goto L9
L11:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	if v582 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L12:
	;
	v563 = v81
	v565 = v83
	v567 = v85
	v570 = v88
	v573 = v91
	v579 = v97
	v581 = v95
	goto L11
L13:
	;
	goto L14
L14:
	;
	if v94 != 0 {
		v738 = v71
		v739 = v72
		v740 = v73
		v741 = v74
		v742 = v75
		v743 = v76
		v744 = v77
		v745 = v78
		v746 = v79
		v748 = v81
		v749 = v82
		v752 = v85
		v754 = v87
		v755 = v88
		v757 = v90
		v758 = v91
		v762 = v95
		v764 = v97
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v109 = v81 & int32(1)
	if v109 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v110 = int32(0)
	goto L18
L17:
	;
	v110 = v83
	goto L18
L18:
	;
	v111 = v109 + v85
	if v111 < v105 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v550 = int32(0)
	v551 = int32(*(*int16)(unsafe.Add(mBase, uint32(v104)+14)))
	if v535 == v551 {
		v563 = v550
		v565 = v533
		v567 = v535
		v570 = v538
		v573 = v541
		v579 = v547
		v581 = v549
		goto L11
	} else {
		goto L112
	}
L20:
	;
	v125 = v110
	v127 = v111
	v133 = v91
	v139 = v97
	v141 = v95
	goto L23
L21:
	;
	v504 = v110
	v506 = v111
	v509 = v88
	v512 = v91
	v518 = v97
	goto L22
L22:
	;
	v533 = v504
	v535 = v506
	v538 = v509
	v541 = v512
	v547 = v518
	v549 = v95
	goto L19
L23:
	;
	v142 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v79)+39)) = uint8(v142)
	F_examine_indexcol_variable(m, v71, v82, v127, v79+int32(40))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v504 = v125
	v506 = v127
	v509 = int32(1)
	v512 = v282
	v518 = v283
	goto L22
L25:
	;
	return
L26:
	;
	v149 = v79 + int32(40)
	v151 = v79 + int32(39)
	v152 = int32(0)
	v153 = float64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v151))) = uint8(v152)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v149)+8))
	if v157 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v79)+48))
	if v127 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L28:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+28)))
	if v195 != 0 {
		goto L42
	} else {
		goto L43
	}
L29:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+22)))
	v160 = v158 + v159
	v161 = *(*float32)(unsafe.Add(mBase, uint32(v160)+8))
	v163 = *(*float32)(unsafe.Add(mBase, uint32(v160)+16))
	v190 = base.F64_promote_f32(v163)
	v191 = base.F64_promote_f32(v161)
	goto L28
L30:
	;
	goto L31
L31:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v149)+16))
	if v165 == int32(16) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v190 = float64(2)
	v191 = v153
	goto L28
L33:
	;
	goto L34
L34:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	if v169 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	if v176 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v169)+76))
	if v172 != int32(5) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v190 = float64(-1)
	v191 = v153
	goto L28
L38:
	;
	v190 = float64(0)
	v191 = v153
	goto L28
L39:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	if v179 != int32(6) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v176)+8)))
	switch v183 - int32(65530) {
	case 0:
		goto L41
	default:
		goto L38
	case 5:
		v190 = float64(-1)
		v191 = v153
		goto L28
	}
L41:
	;
	v190 = float64(1)
	v191 = v153
	goto L28
L42:
	;
	v196 = base.F64_neg(base.F64_sub(float64(1), v191))
	goto L44
L43:
	;
	v196 = v190
	goto L44
L44:
	;
	if base.F64_gt(v196, float64(0)) != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v199 = F_clamp_row_est(m, v196)
	mBase = m.M
	v225 = v199
	goto L27
L46:
	;
	goto L47
L47:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	if v200 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v203 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v151))) = uint8(v203)
	v225 = float64(200)
	goto L27
L49:
	;
	goto L50
L50:
	;
	v206 = *(*float64)(unsafe.Add(mBase, uint32(v200)+120))
	if base.F64_le(v206, float64(0)) != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v209 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v151))) = uint8(v209)
	v225 = float64(200)
	goto L27
L52:
	;
	goto L53
L53:
	;
	if base.F64_lt(v196, float64(0)) != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v216 = F_clamp_row_est(m, base.F64_mul(v206, base.F64_neg(v196)))
	mBase = m.M
	v225 = v216
	goto L27
L55:
	;
	goto L56
L56:
	;
	if base.F64_lt(v206, float64(200)) != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v219 = F_clamp_row_est(m, v206)
	mBase = m.M
	v225 = v219
	goto L27
L58:
	;
	goto L59
L59:
	;
	v220 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v151))) = uint8(v220)
	v225 = float64(200)
	goto L27
L60:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+39)))
	if v285 != 0 {
		goto L81
	} else {
		goto L82
	}
L61:
	;
	if v226 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v272 = v226
	v273 = v133
	v274 = v139
	goto L63
L63:
	;
	if v272 == int32(0) {
		v282 = v273
		v283 = v274
		goto L60
	} else {
		goto L79
	}
L64:
	;
	v282 = int32(1)
	v283 = v139
	goto L60
L65:
	;
	goto L66
L66:
	;
	v232 = float64(0)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v82)+52))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v82)+56))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	v238 = F_get_opfamily_member(m, v234, v236, v236, int32(1))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L25
	} else {
		goto L68
	}
L67:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v79)+48))
	v272 = v271
	v273 = int32(1)
	v274 = v268
	goto L63
L68:
	;
	if v238 == int32(0) {
		v268 = v232
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v79)+48))
	v247 = F_get_attstatsslot(m, v79+int32(140), v244, int32(3), v238, int32(2))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L25
	} else {
		goto L70
	}
L70:
	;
	if v247 == int32(0) {
		v268 = v232
		goto L67
	} else {
		goto L71
	}
L71:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v79)+160))
	v252 = *(*float32)(unsafe.Add(mBase, uint32(v251)))
	v253 = base.F64_promote_f32(v252)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v82)+64))
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255))))
	if v256 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v257 = base.F64_neg(v253)
	goto L74
L73:
	;
	v257 = v253
	goto L74
L74:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v82)+40))
	if int32(1) < v260 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v263 = base.F64_mul(v257, float64(0.75))
	goto L77
L76:
	;
	v263 = v257
	goto L77
L77:
	;
	F_free_attstatsslot(m, v79+int32(140))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L25
	} else {
		goto L78
	}
L78:
	;
	v268 = v263
	goto L67
L79:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v79)+52))
	m.T0[v278].(func(*base.Module, int32))(m, v272)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L25
	} else {
		goto L80
	}
L80:
	;
	v282 = v273
	v283 = v274
	goto L60
L81:
	;
	goto L24
L82:
	;
	if v125 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v82)+88))
	if v286 != 0 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	v445 = v225
	goto L85
L85:
	;
	if v125 != 0 {
		goto L107
	} else {
		goto L108
	}
L86:
	;
	v287 = int32(0)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v286)+4))
	if v287 < v288 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	v405 = v125
	goto L88
L88:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v406)+68))
	v408 = int32(0)
	v410 = F_clauselist_selectivity(m, v71, v405, v407, v408, v408)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L25
	} else {
		goto L102
	}
L89:
	;
	v301 = int32(0)
	v305 = v287
	goto L92
L90:
	;
	v358 = v287
	goto L91
L91:
	;
	v374 = F_list_concat(m, v358, v125)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L25
	} else {
		goto L101
	}
L92:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v286)+12))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v321+v301<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+32)) = v325
	*(*int32)(unsafe.Add(mBase, uint32(v79)+140)) = v325
	v331 = F_list_make1_impl(m, int32(1), v79+int32(32))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L25
	} else {
		goto L94
	}
L93:
	;
	v358 = v340
	goto L91
L94:
	;
	v334 = F_predicate_implied_by(m, v331, v125, int32(0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L25
	} else {
		goto L95
	}
L95:
	;
	if v334 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v338 = F_list_concat(m, v305, v331)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L25
	} else {
		goto L99
	}
L97:
	;
	v340 = v305
	goto L98
L98:
	;
	v342 = v301 + int32(1)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v286)+4))
	if v342 < v343 {
		v301 = v342
		v305 = v340
		goto L92
	} else {
		goto L100
	}
L99:
	;
	v340 = v338
	goto L98
L100:
	;
	goto L93
L101:
	;
	v405 = v374
	goto L88
L102:
	;
	if base.F64_lt(v410, float64(0.005)) != 0 {
		goto L81
	} else {
		goto L103
	}
L103:
	;
	v415 = base.F64_nearest(base.F64_mul(v225, v410))
	v416 = float64(1)
	if base.F64_gt(v415, v416) != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v419 = v415
	goto L106
L105:
	;
	v419 = v416
	goto L106
L106:
	;
	v445 = v419
	goto L85
L107:
	;
	v451 = v445
	goto L109
L108:
	;
	v451 = base.F64_add(v445, float64(1))
	goto L109
L109:
	;
	v452 = base.F64_mul(v141, v451)
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v82)+16))
	if base.F64_gt(v452, base.F64_convert_i32_u(v453)) != 0 {
		goto L81
	} else {
		goto L110
	}
L110:
	;
	v456 = int32(1)
	v457 = int32(0)
	v459 = v127 + v456
	v460 = int32(*(*int16)(unsafe.Add(mBase, uint32(v104)+14)))
	if v460 <= v459 {
		v533 = v457
		v535 = v459
		v538 = v456
		v541 = v282
		v547 = v283
		v549 = v452
		goto L19
	} else {
		goto L111
	}
L111:
	;
	v125 = v457
	v127 = v459
	v133 = v282
	v139 = v283
	v141 = v452
	goto L23
L112:
	;
	v738 = v71
	v739 = v72
	v740 = v73
	v741 = v74
	v742 = v75
	v743 = v76
	v744 = v77
	v745 = v78
	v746 = v79
	v748 = v550
	v749 = v82
	v752 = v535
	v754 = v87
	v755 = v538
	v757 = v90
	v758 = v541
	v762 = v549
	v764 = v547
	goto L10
L113:
	;
	v735 = v93 + int32(1)
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	if v735 < v736 {
		v81 = v715
		v83 = v717
		v85 = v567
		v87 = v721
		v88 = v722
		v90 = v724
		v91 = v573
		v93 = v735
		v94 = v728
		v95 = v729
		v97 = v579
		goto L8
	} else {
		goto L143
	}
L114:
	;
	v715 = v563
	v717 = v565
	v721 = v87
	v722 = v570
	v724 = v90
	v728 = v94
	v729 = v581
	goto L113
L115:
	;
	goto L116
L116:
	;
	v585 = int32(0)
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v582)+4))
	if v586 <= v585 {
		v715 = v563
		v717 = v565
		v721 = v87
		v722 = v570
		v724 = v90
		v728 = v94
		v729 = v581
		goto L113
	} else {
		goto L117
	}
L117:
	;
	v598 = v585
	v599 = v563
	v601 = v565
	v605 = v87
	v606 = v570
	v608 = v90
	v612 = v94
	v613 = v581
	goto L118
L118:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v582)+12))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v618+v598<<(uint(int32(2))%32))))
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v622)+4))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v623)))
	switch v624 - int32(17) {
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
	v715 = v683
	v717 = v700
	v721 = v690
	v722 = v685
	v724 = v686
	v728 = v687
	v729 = v688
	goto L113
L120:
	;
	v690 = F_lappend(m, v605, v622)
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L25
	} else {
		goto L136
	}
L121:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v669)))
	if v670 == int32(0) {
		v683 = v599
		v685 = v665
		v686 = v608
		v687 = v666
		v688 = v667
		goto L120
	} else {
		goto L134
	}
L122:
	;
	v665 = v606
	v666 = v612
	v667 = v613
	v669 = v623 + int32(4)
	goto L121
L123:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L25
	} else {
		goto L131
	}
L124:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v623)+8))
	v644 = base.B2i32(v642 == int32(0))
	v683 = v644 | v599
	v685 = v606
	v686 = v644 | v608
	v687 = v612
	v688 = v613
	goto L120
L125:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v623)+28))
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v630)+12))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v631)+4))
	v633 = F_estimate_array_length(m, v71, v632)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L25
	} else {
		goto L127
	}
L126:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v623)+8))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v628)+12))
	v665 = v606
	v666 = int32(1)
	v667 = v613
	v669 = v629
	goto L121
L127:
	;
	if base.F64_gt(v633, float64(1)) != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v638 = base.F64_mul(v613, v633)
	goto L130
L129:
	;
	v638 = v613
	goto L130
L130:
	;
	v665 = int32(1)
	v666 = v612
	v667 = v638
	v669 = v623 + int32(4)
	goto L121
L131:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v623)))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+16)) = v651
	F_errmsg_internal(m, int32(495628), v79+int32(16))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L25
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(506228), int32(7640), int32(362754))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L25
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
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v82)+52))
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v673+v567<<(uint(int32(2))%32))))
	v678 = F_get_op_opfamily_strategy(m, v670, v677)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L25
	} else {
		goto L135
	}
L135:
	;
	v683 = base.B2i32(v678 == int32(3)) | v599
	v685 = v665
	v686 = v608
	v687 = v666
	v688 = v667
	goto L120
L136:
	;
	if v683&int32(1) != 0 {
		v700 = v601
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v702 = v598 + int32(1)
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v582)+4))
	if v702 < v703 {
		v598 = v702
		v599 = v683
		v601 = v700
		v605 = v690
		v606 = v685
		v608 = v686
		v612 = v687
		v613 = v688
		goto L118
	} else {
		goto L142
	}
L138:
	;
	if v687 != 0 {
		v700 = v601
		goto L137
	} else {
		goto L139
	}
L139:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v82)+40))
	if v694-int32(1) <= v567 {
		v700 = v601
		goto L137
	} else {
		goto L140
	}
L140:
	;
	v698 = F_lappend(m, v601, v622)
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L25
	} else {
		goto L141
	}
L141:
	;
	v700 = v698
	goto L137
L142:
	;
	goto L119
L143:
	;
	v738 = v71
	v739 = v72
	v740 = v73
	v741 = v74
	v742 = v75
	v743 = v76
	v744 = v77
	v745 = v78
	v746 = v79
	v748 = v715
	v749 = v82
	v752 = v567
	v754 = v721
	v755 = v722
	v757 = v724
	v758 = v573
	v762 = v729
	v764 = v579
	goto L10
L144:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v777)+128)) = v980
	*(*float64)(unsafe.Add(mBase, uint32(v777)+112)) = v985
	F_genericcostestimate(m, v769, v770, v771, v777+int32(72))
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L25
	} else {
		goto L173
	}
L145:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v780)+40))
	v803 = int32(1)
	if (base.B2i32(v783 != v802-v803)|v778|v786|v788)&v803 == int32(0) {
		v980 = v793
		v985 = float64(1)
		goto L144
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v780)+88))
	if v813 != 0 {
		goto L149
	} else {
		goto L150
	}
L148:
	;
	goto L147
L149:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v813)+4))
	if v814 <= int32(0) {
		goto L153
	} else {
		goto L154
	}
L150:
	;
	v933 = v785
	goto L151
L151:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v780)+12))
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v934)+68))
	v936 = int32(0)
	v938 = F_clauselist_selectivity(m, v769, v933, v935, v936, v936)
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L25
	} else {
		goto L166
	}
L152:
	;
	v902 = F_list_concat(m, v885, v785)
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L25
	} else {
		goto L165
	}
L153:
	;
	v885 = int32(0)
	goto L152
L154:
	;
	goto L155
L155:
	;
	v818 = int32(0)
	v829 = v818
	v832 = v818
	goto L156
L156:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v813)+12))
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v849+v829<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v777)+12)) = v853
	*(*int32)(unsafe.Add(mBase, uint32(v777)+140)) = v853
	v859 = F_list_make1_impl(m, int32(1), v777+int32(12))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L25
	} else {
		goto L158
	}
L157:
	;
	v885 = v868
	goto L152
L158:
	;
	v862 = F_predicate_implied_by(m, v859, v785, int32(0))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L25
	} else {
		goto L159
	}
L159:
	;
	if v862 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v866 = F_list_concat(m, v832, v859)
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L25
	} else {
		goto L163
	}
L161:
	;
	v868 = v832
	goto L162
L162:
	;
	v870 = v829 + int32(1)
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v813)+4))
	if v870 < v871 {
		v829 = v870
		v832 = v868
		goto L156
	} else {
		goto L164
	}
L163:
	;
	v868 = v866
	goto L162
L164:
	;
	goto L157
L165:
	;
	v933 = v902
	goto L151
L166:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v780)+12))
	v941 = *(*float64)(unsafe.Add(mBase, uint32(v940)+120))
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v780)+16))
	v947 = base.F64_ceil(base.F64_mul(base.F64_convert_i32_u(v943), float64(0.3333333)))
	if base.F64_lt(v793, v947) != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v949 = v793
	goto L169
L168:
	;
	v949 = v947
	goto L169
L169:
	;
	v950 = float64(1)
	if base.F64_gt(v949, v950) != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v953 = v949
	goto L172
L171:
	;
	v953 = v950
	goto L172
L172:
	;
	v980 = v953
	v985 = base.F64_nearest(base.F64_div(base.F64_mul(v938, v941), v953))
	goto L144
L173:
	;
	v993 = *(*float64)(unsafe.Add(mBase, _consts[385]))
	v994 = *(*float64)(unsafe.Add(mBase, uint32(v777)+128))
	v995 = *(*float64)(unsafe.Add(mBase, uint32(v777)+72))
	v996 = *(*float64)(unsafe.Add(mBase, uint32(v780)+24))
	if base.F64_gt(v996, float64(1)) == int32(0) {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v780)+32))
	if v789 != 0 {
		v1068 = v795
		goto L178
	} else {
		goto L179
	}
L175:
	;
	v1001 = *(*float64)(unsafe.Add(mBase, uint32(v777)+80))
	v1011 = v995
	v1013 = v1001
	goto L174
L176:
	;
	goto L177
L177:
	;
	v1002 = F_log(m, v996)
	mBase = m.M
	v1006 = base.F64_mul(base.F64_ceil(base.F64_div(v1002, float64(0.6931471805599453))), v993)
	v1009 = *(*float64)(unsafe.Add(mBase, uint32(v777)+80))
	v1011 = base.F64_add(v995, v1006)
	v1013 = base.F64_add(base.F64_mul(v994, v1006), v1009)
	goto L174
L178:
	;
	v1074 = base.F64_mul(v993, base.F64_mul(base.F64_convert_i32_s(v1014+int32(1)), float64(50)))
	*(*float64)(unsafe.Add(mBase, uint32(v772))) = base.F64_add(v1011, v1074)
	*(*float64)(unsafe.Add(mBase, uint32(v773))) = base.F64_add(base.F64_mul(v994, v1074), v1013)
	v1080 = *(*float64)(unsafe.Add(mBase, uint32(v777)+88))
	*(*float64)(unsafe.Add(mBase, uint32(v774))) = v1080
	*(*float64)(unsafe.Add(mBase, uint32(v775))) = v1068
	v1083 = *(*float64)(unsafe.Add(mBase, uint32(v777)+104))
	*(*float64)(unsafe.Add(mBase, uint32(v776))) = v1083
	m.G0 = v777 + int32(176)
	return
L179:
	;
	F_examine_indexcol_variable(m, v769, v780, int32(0), v777+int32(40))
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L25
	} else {
		goto L180
	}
L180:
	;
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v777)+48))
	if v1020 == int32(0) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v1023 = *(*float64)(unsafe.Add(mBase, uint32(v777)+96))
	v1068 = v1023
	goto L178
L182:
	;
	goto L183
L183:
	;
	v1024 = float64(0)
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v780)+52))
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v1025)))
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v780)+56))
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v1027)))
	v1030 = F_get_opfamily_member(m, v1026, v1028, v1028, int32(1))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L25
	} else {
		goto L185
	}
L184:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v777)+48))
	if v1061 == int32(0) {
		v1068 = v1060
		goto L178
	} else {
		goto L196
	}
L185:
	;
	if v1030 == int32(0) {
		v1060 = v1024
		goto L184
	} else {
		goto L186
	}
L186:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v777)+48))
	v1039 = F_get_attstatsslot(m, v777+int32(140), v1036, int32(3), v1030, int32(2))
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L25
	} else {
		goto L187
	}
L187:
	;
	if v1039 == int32(0) {
		v1060 = v1024
		goto L184
	} else {
		goto L188
	}
L188:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v777)+160))
	v1044 = *(*float32)(unsafe.Add(mBase, uint32(v1043)))
	v1045 = base.F64_promote_f32(v1044)
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v780)+64))
	v1048 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047))))
	if v1048 != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v1049 = base.F64_neg(v1045)
	goto L191
L190:
	;
	v1049 = v1045
	goto L191
L191:
	;
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v780)+40))
	if int32(1) < v1052 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v1055 = base.F64_mul(v1049, float64(0.75))
	goto L194
L193:
	;
	v1055 = v1049
	goto L194
L194:
	;
	F_free_attstatsslot(m, v777+int32(140))
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L25
	} else {
		goto L195
	}
L195:
	;
	v1060 = v1055
	goto L184
L196:
	;
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v777)+52))
	m.T0[v1064].(func(*base.Module, int32))(m, v1061)
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L25
	} else {
		goto L197
	}
L197:
	;
	v1068 = v1060
	goto L178
}
func F_btfloat4sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(1311)
	return int32(0)
}
func F_btfloat8sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(1312)
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
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int64
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v514 int32
	_ = v514
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v588 int32
	_ = v588
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v651 int32
	_ = v651
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v736 int32
	_ = v736
	var v742 int32
	_ = v742
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v764 int32
	_ = v764
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v780 int32
	_ = v780
	var v789 int32
	_ = v789
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v879 int32
	_ = v879
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v921 int32
	_ = v921
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v966 int32
	_ = v966
	var v989 int32
	_ = v989
	var v995 int32
	_ = v995
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1025 int32
	_ = v1025
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1041 int64
	_ = v1041
	var v1043 int64
	_ = v1043
	var v1044 int64
	_ = v1044
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1075 int32
	_ = v1075
	var v1081 int32
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1101 int32
	_ = v1101
	var v1104 int32
	_ = v1104
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1115 int32
	_ = v1115
	var v1123 int32
	_ = v1123
	var v1128 int32
	_ = v1128
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1144 int32
	_ = v1144
	var v1151 int32
	_ = v1151
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1160 int32
	_ = v1160
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1206 int32
	_ = v1206
	var v1213 int32
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1230 int32
	_ = v1230
	var v1236 int32
	_ = v1236
	var v1238 int32
	_ = v1238
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1285 int32
	_ = v1285
	var v1288 int32
	_ = v1288
	var v1294 int32
	_ = v1294
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1337 int32
	_ = v1337
	var v1341 int32
	_ = v1341
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1356 int32
	_ = v1356
	var v1390 int32
	_ = v1390
	var v1392 int32
	_ = v1392
	var v1397 int32
	_ = v1397
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
	*(*int32)(unsafe.Add(mBase, uint32(v44)+72)) = (v60&int32(8191) + int32(7)) & int32(16376)
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
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	if l5 != int32(3) {
		goto L210
	} else {
		goto L211
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
	v855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v855 != int32(1) {
		v879 = v851
		goto L7
	} else {
		goto L206
	}
L10:
	;
	if v57 == int32(0) {
		v879 = v58
		goto L7
	} else {
		goto L48
	}
L11:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v44)+76))
	v225 = F__bt_search(m, l0, l4, v223, v44+int32(80), int32(2))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L47
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
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v189 != 0 {
		goto L39
	} else {
		goto L40
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
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L38
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
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	F__bt_relbuf(m, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L37
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
	v129 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v129+(v125^int32(-1))<<(uint(int32(2))%32))))
	v143 = v135
	goto L22
L24:
	;
	goto L25
L25:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _consts[13]))
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
	if base.Ui32(v163) < base.Ui32(int32(25)) {
		goto L21
	} else {
		goto L33
	}
L33:
	;
	if (v163+int32(262120))&int32(262140) == int32(0) {
		goto L21
	} else {
		goto L34
	}
L34:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v44)+76))
	v175 = F__bt_compare(m, l0, v173, v143, int32(1))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	if int32(0) < v175 {
		v229 = int32(0)
		goto L10
	} else {
		goto L36
	}
L36:
	;
	goto L21
L37:
	;
	goto L16
L38:
	;
	goto L16
L39:
	;
	v217 = v189
	goto L41
L40:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v44-int32(-64)))) = v193
	v195 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v44)+56)) = v195
	v199 = F_smgropen(m, v44+int32(56), v190)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217)+16)) = int32(-1)
	goto L11
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v199
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v199)+72))
	if v203 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v217 = v215
	goto L41
L44:
	;
	v211 = v203
	goto L46
L45:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v199)+76))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v199)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v204)+4)) = v205
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v199)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v205))) = v207
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v199)+72))
	v211 = v209
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v199)+72)) = v211 + int32(1)
	goto L43
L47:
	;
	v229 = v225
	goto L10
L48:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v44)+76))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v44)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+280)) = int32(4)
	v236 = int32(0)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	if v237 < v236 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v256 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v255)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v256) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	v241 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v241+(v237^int32(-1))<<(uint(int32(2))%32))))
	v255 = v247
	goto L49
L51:
	;
	goto L52
L52:
	;
	v249 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v255 = v249 + v237<<(uint(int32(13))%32) + int32(-8192)
	goto L49
L53:
	;
	v264 = int32(base.Ui32(v256+int32(262120)) >> (uint(int32(2)) % 32))
	goto L55
L54:
	;
	v264 = int32(0)
	goto L55
L55:
	;
	v265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v255)+16)))
	v270 = F__bt_binsrch_insert(m, l0, v44+int32(68))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v272 = int32(0)
	v278 = v272
	v284 = v272
	v285 = v236
	v289 = v255
	v292 = v255 + v265
	v297 = int32(1)
	v298 = v270
	v300 = v272
	v302 = v264
	goto L63
L57:
	;
	goto L9
L58:
	;
	if v229 == int32(0) {
		goto L8
	} else {
		goto L204
	}
L59:
	;
	F_XactLockTableWait(m, v677, l0, v34, int32(5))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L1
	} else {
		goto L203
	}
L60:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L1
	} else {
		goto L197
	}
L61:
	;
	if base.B2i32(v789 == int32(0))&base.B2i32(l5 == int32(3)) != 0 {
		goto L60
	} else {
		goto L194
	}
L62:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v44)+284))
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v44)+288))
	if v675 != 0 {
		goto L158
	} else {
		goto L159
	}
L63:
	;
	v312 = v278
	v318 = v284
	v331 = v297
	v332 = v298
	v334 = v300
	goto L66
L64:
	;
	if l5 != int32(2) {
		goto L62
	} else {
		goto L153
	}
L65:
	;
	goto L64
L66:
	;
	v342 = v332 & int32(65535)
	v347 = v342<<(uint(int32(2))%32) + (v289 + int32(24)) - int32(4)
	v348 = int32(0)
	v350 = v348
	v353 = v312
	v359 = v318
	v362 = v348
	v372 = v331
	v375 = v334
	goto L68
L67:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v292)+4))
	if v571 == int32(0) {
		v789 = v567
		goto L61
	} else {
		goto L131
	}
L68:
	;
	v383 = v302 & int32(65535)
	if base.Ui32(v342) <= base.Ui32(v383) {
		goto L76
	} else {
		goto L77
	}
L69:
	;
	if base.Ui32(v342) < base.Ui32(v383) {
		goto L128
	} else {
		goto L129
	}
L70:
	;
	goto L69
L71:
	;
	v550 = int32(1)
	v553 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v545)+4)))
	if v543 < v553&int32(4095)-v550 {
		v350 = v543 + v550
		v353 = v544
		v359 = v545
		v362 = v550
		v372 = v547
		v375 = v549
		goto L68
	} else {
		goto L127
	}
L72:
	;
	v452 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v449)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+276)) = uint16(v452)
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+272)) = v454
	if base.B2i32(l5 != int32(3)) == int32(0) {
		goto L101
	} else {
		goto L102
	}
L73:
	;
	v448 = v350
	v449 = v447
	v450 = v372
	v451 = v446
	goto L72
L74:
	;
	v446 = v362
	v447 = v409
	goto L73
L75:
	;
	if v375|base.B2i32(l5 != int32(3)) != 0 {
		v851 = int32(1)
		goto L57
	} else {
		goto L98
	}
L76:
	;
	if v285 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	goto L78
L78:
	;
	if v362&int32(1) == int32(0) {
		v561 = v353
		v562 = v359
		v565 = v372
		v567 = v375
		goto L70
	} else {
		goto L97
	}
L79:
	;
	v387 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+88)))
	if v342 == v387 {
		goto L75
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	if v362&int32(1) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	goto L81
L83:
	;
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409)+7)))
	if v410&int32(32) == int32(0) {
		goto L74
	} else {
		goto L92
	}
L84:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v347)))
	v394 = int32(98304)
	if v393&v394 == v394 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	goto L86
L86:
	;
	v406 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+271)) = uint8(v406)
	v408 = v353
	v409 = v359
	goto L83
L87:
	;
	v561 = v347
	v562 = v359
	v565 = v372
	v567 = v375
	goto L70
L88:
	;
	goto L89
L89:
	;
	v398 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+271)) = uint8(v398)
	v400 = F__bt_compare(m, l0, v232, v289, v342)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	if v400 != 0 {
		v789 = v375
		goto L61
	} else {
		goto L91
	}
L91:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v347)))
	v408 = v347
	v409 = v289 + v402&int32(32767)
	goto L83
L92:
	;
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409)+5)))
	if v415&int32(32) == int32(0) {
		goto L74
	} else {
		goto L93
	}
L93:
	;
	v420 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v409)+2)))
	v421 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v409))))
	v425 = v409 + (v420 | v421<<(uint(int32(16))%32))
	v426 = int32(1)
	v427 = int32(0)
	if v362&v426 == v427 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v448 = v427
	v449 = v425
	v450 = int32(1)
	v451 = v426
	goto L72
L95:
	;
	goto L96
L96:
	;
	v446 = v426
	v447 = v425 + v350*int32(6)
	goto L73
L97:
	;
	v543 = v350
	v544 = v353
	v545 = v359
	v547 = v372
	v549 = v375
	goto L71
L98:
	;
	goto L60
L99:
	;
	v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+271)))
	v535 = int32(1)
	v538 = (v534 | (v451 ^ v535)) & v450
	if v451&v535 == int32(0) {
		v561 = v408
		v562 = v409
		v565 = v538
		v567 = v533
		goto L70
	} else {
		goto L126
	}
L100:
	;
	v505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+271)))
	if v505 != int32(1) {
		v533 = v375
		goto L99
	} else {
		goto L116
	}
L101:
	;
	v459 = v44 + int32(272)
	v463 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v459)+2)))
	v464 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v459))))
	v465 = int32(16)
	v467 = v463 | v464<<(uint(v465)%32)
	v468 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v233)+2)))
	v469 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v233))))
	v472 = v468 | v469<<(uint(v465)%32)
	if base.Ui32(v467) < base.Ui32(v472) {
		v483 = int32(-1)
		goto L105
	} else {
		goto L106
	}
L102:
	;
	goto L103
L103:
	;
	v503 = F_table_index_fetch_tuple_check(m, l4, v44+int32(272), v44+int32(280), v44+int32(271))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L114
	}
L104:
	;
	if v483 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L105:
	;
	goto L104
L106:
	;
	if base.Ui32(v472) < base.Ui32(v467) {
		v483 = int32(1)
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v477 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v459)+4)))
	v478 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v233)+4)))
	if base.Ui32(v477) < base.Ui32(v478) {
		v483 = int32(-1)
		goto L105
	} else {
		goto L108
	}
L108:
	;
	v483 = base.B2i32(base.Ui32(v478) < base.Ui32(v477))
	goto L105
L109:
	;
	v533 = int32(1)
	goto L99
L110:
	;
	goto L111
L111:
	;
	v493 = F_table_index_fetch_tuple_check(m, l4, v44+int32(272), v44+int32(280), v44+int32(271))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	if v493 == int32(0) {
		goto L100
	} else {
		goto L113
	}
L113:
	;
	goto L62
L114:
	;
	if v503 != 0 {
		goto L65
	} else {
		goto L115
	}
L115:
	;
	goto L100
L116:
	;
	if v451&int32(1) != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	if v450&int32(1) == int32(0) {
		v543 = v448
		v544 = v408
		v545 = v409
		v547 = v450
		v549 = v375
		goto L71
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v408)))
	*(*int32)(unsafe.Add(mBase, uint32(v408))) = v520 | int32(98304)
	v524 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v292)+12)))
	v526 = v524 | int32(64)
	*(*uint16)(unsafe.Add(mBase, uint32(v292)+12)) = uint16(v526)
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	if v285 != 0 {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	v514 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v409)+4)))
	if v448 != v514&int32(4095)-int32(1) {
		v543 = v448
		v544 = v408
		v545 = v409
		v547 = v450
		v549 = v375
		goto L71
	} else {
		goto L121
	}
L121:
	;
	goto L119
L122:
	;
	v529 = v285
	goto L124
L123:
	;
	v529 = v528
	goto L124
L124:
	;
	F_MarkBufferDirtyHint(m, v529, int32(1))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	v533 = v375
	goto L99
L126:
	;
	v543 = v448
	v544 = v408
	v545 = v409
	v547 = v538
	v549 = v533
	goto L71
L127:
	;
	v561 = v544
	v562 = v545
	v565 = v547
	v567 = v549
	goto L70
L128:
	;
	v312 = v561
	v318 = v562
	v331 = v565
	v332 = v332 + int32(1)
	v334 = v567
	goto L66
L129:
	;
	goto L130
L130:
	;
	goto L67
L131:
	;
	v575 = F__bt_compare(m, l0, v232, v289, int32(1))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	if v575 != 0 {
		v789 = v567
		goto L61
	} else {
		goto L133
	}
L133:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v292)+4))
	v579 = v577
	v588 = v285
	goto L135
L134:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v632)+4))
	if v659 != 0 {
		goto L147
	} else {
		goto L148
	}
L135:
	;
	v611 = F__bt_relandgetbuf(m, l0, v588, v579, int32(1))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L1
	} else {
		goto L138
	}
L136:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L1
	} else {
		goto L144
	}
L137:
	;
	v631 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v630)+16)))
	v632 = v631 + v630
	v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v632)+12)))
	if v633&int32(20) == int32(0) {
		goto L134
	} else {
		goto L142
	}
L138:
	;
	if v611 < int32(0) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v616 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v616+(v611^int32(-1))<<(uint(int32(2))%32))))
	v630 = v622
	goto L137
L140:
	;
	goto L141
L141:
	;
	v624 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v630 = v624 + v611<<(uint(int32(13))%32) + int32(-8192)
	goto L137
L142:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v632)+4))
	if v638 != 0 {
		v579 = v638
		v588 = v611
		goto L135
	} else {
		goto L143
	}
L143:
	;
	goto L136
L144:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = v643 + int32(4)
	F_errmsg_internal(m, int32(710441), v44+int32(16))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(503817), int32(743), int32(352307))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L147:
	;
	v660 = int32(2)
	goto L149
L148:
	;
	v660 = int32(1)
	goto L149
L149:
	;
	v661 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v630)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v661) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v669 = int32(base.Ui32(v661+int32(262120)) >> (uint(int32(2)) % 32))
	goto L152
L151:
	;
	v669 = int32(0)
	goto L152
L152:
	;
	v278 = v561
	v284 = v562
	v285 = v611
	v289 = v630
	v292 = v632
	v297 = v565
	v298 = v660
	v300 = v567
	v302 = v669
	goto L63
L153:
	;
	if v285 != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	F__bt_relbuf(m, v285)
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L1
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v851 = int32(0)
	goto L57
L157:
	;
	goto L156
L158:
	;
	v677 = v675
	goto L160
L159:
	;
	v677 = v676
	goto L160
L160:
	;
	if v677 != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	if v285 != 0 {
		goto L164
	} else {
		goto L165
	}
L162:
	;
	goto L163
L163:
	;
	v692 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v233)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+276)) = uint16(v692)
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+272)) = v694
	v700 = F_table_index_fetch_tuple_check(m, l4, v44+int32(272), int32(4194344), int32(0))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L1
	} else {
		goto L171
	}
L164:
	;
	F__bt_relbuf(m, v285)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L1
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v680 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+84)) = uint8(v680)
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v44)+316))
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	F__bt_relbuf(m, v683)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L168
	}
L167:
	;
	goto L166
L168:
	;
	v686 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+80)) = v686
	if v682 == v686 {
		goto L59
	} else {
		goto L169
	}
L169:
	;
	F_SpeculativeInsertionWait(m, v677, v682)
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	goto L58
L171:
	;
	if v700 == int32(0) {
		v789 = v375
		goto L61
	} else {
		goto L172
	}
L172:
	;
	v704 = int32(0)
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	if v705 < v704 {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	F_CheckForSerializableConflictIn(m, l0, v704, v724)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L1
	} else {
		goto L177
	}
L174:
	;
	v709 = *(*int32)(unsafe.Add(mBase, _consts[16]))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v709+(v705^int32(-1))<<(uint(int32(6))%32))+16))
	v724 = v715
	goto L173
L175:
	;
	goto L176
L176:
	;
	v717 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v717+v705<<(uint(int32(6))%32)+int32(-64))+16))
	v724 = v723
	goto L173
L177:
	;
	if v285 != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	F__bt_relbuf(m, v285)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L1
	} else {
		goto L181
	}
L179:
	;
	goto L180
L180:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	F__bt_relbuf(m, v729)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L1
	} else {
		goto L182
	}
L181:
	;
	goto L180
L182:
	;
	v732 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+84)) = uint8(v732)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+80)) = v732
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_index_deform_tuple(m, v233, v736, v44+int32(128), v44+int32(96))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	v747 = F_BuildIndexValueDescription(m, l0, v44+int32(128), v44+int32(96))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	F_errcode(m, int32(83906754))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+48)) = v756 + int32(4)
	F_errmsg(m, int32(713133), v44+int32(48))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	if v747 != 0 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+32)) = v747
	F_errdetail(m, int32(597014), v44+int32(32))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L1
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_errtableconstraint(m, l4, v771+int32(4))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L1
	} else {
		goto L192
	}
L191:
	;
	goto L190
L192:
	;
	F_errfinish(m, int32(503817), int32(673), int32(352307))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L194:
	;
	v795 = int32(1)
	if v285 == int32(0) {
		v851 = v795
		goto L57
	} else {
		goto L195
	}
L195:
	;
	F__bt_relbuf(m, v285)
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	v851 = v795
	goto L57
L197:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v817 + int32(4)
	F_errmsg(m, int32(707345), v44)
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	F_errhint(m, int32(631377), int32(0))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_errtableconstraint(m, l4, v828+int32(4))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	F_errfinish(m, int32(503817), int32(766), int32(352307))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
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
	goto L58
L204:
	;
	F__bt_freestack(m, v229)
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	goto L8
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = v34
	v879 = v851
	goto L7
L207:
	;
	if v229 != 0 {
		goto L315
	} else {
		goto L316
	}
L208:
	;
	v1329 = F__bt_binsrch_insert(m, l0, v44+int32(68))
	mBase = m.M
	v1330 = m.ExcPending
	if v1330 != 0 {
		goto L1
	} else {
		goto L308
	}
L209:
	;
	if v57 == int32(0) {
		goto L267
	} else {
		goto L268
	}
L210:
	;
	v894 = int32(0)
	if v891 < v894 {
		goto L214
	} else {
		goto L215
	}
L211:
	;
	goto L212
L212:
	;
	F__bt_relbuf(m, v891)
	mBase = m.M
	v1104 = m.ExcPending
	if v1104 != 0 {
		goto L1
	} else {
		goto L265
	}
L213:
	;
	F_CheckForSerializableConflictIn(m, l0, v894, v913)
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L1
	} else {
		goto L217
	}
L214:
	;
	v898 = *(*int32)(unsafe.Add(mBase, _consts[16]))
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v898+(v891^int32(-1))<<(uint(int32(6))%32))+16))
	v913 = v904
	goto L213
L215:
	;
	goto L216
L216:
	;
	v906 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v906+v891<<(uint(int32(6))%32)+int32(-64))+16))
	v913 = v912
	goto L213
L217:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v44)+76))
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	if v917 < int32(0) {
		goto L219
	} else {
		goto L220
	}
L218:
	;
	v936 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v935)+16)))
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v44)+72))
	if base.Ui32(int32(2705)) <= base.Ui32(v937) {
		goto L222
	} else {
		goto L223
	}
L219:
	;
	v921 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v921+(v917^int32(-1))<<(uint(int32(2))%32))))
	v935 = v927
	goto L218
L220:
	;
	goto L221
L221:
	;
	v929 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v935 = v929 + v917<<(uint(int32(13))%32) + int32(-8192)
	goto L218
L222:
	;
	v940 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v916))))
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v44)+68))
	F__bt_check_third_page(m, l0, l4, v940, v935, v941)
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L1
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	v944 = v935 + v936
	v945 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v916))))
	if v945 != 0 {
		goto L209
	} else {
		goto L226
	}
L225:
	;
	goto L224
L226:
	;
	v946 = int32(4)
	v947 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v935)+14)))
	v948 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v935)+12)))
	v949 = v947 - v948
	if v949 <= v946 {
		goto L228
	} else {
		goto L229
	}
L227:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v44)+72))
	if base.Ui32(v955) <= base.Ui32(v952-int32(4)) {
		goto L208
	} else {
		goto L231
	}
L228:
	;
	v952 = v946
	goto L230
L229:
	;
	v952 = v949
	goto L230
L230:
	;
	goto L227
L231:
	;
	v958 = v935
	v966 = v944
	goto L232
L232:
	;
	v989 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v966)+12)))
	if v989&int32(64) != 0 {
		goto L234
	} else {
		goto L235
	}
L233:
	;
	goto L208
L234:
	;
	v995 = int32(0)
	F__bt_delete_or_dedup_one_page(m, l0, l4, v44+int32(68), int32(1), v995, v995, v995)
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L1
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	v1011 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+84)))
	if v1011 != int32(1) {
		goto L243
	} else {
		goto L244
	}
L237:
	;
	v1000 = int32(4)
	v1001 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v958)+14)))
	v1002 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v958)+12)))
	v1003 = v1001 - v1002
	if v1003 <= v1000 {
		goto L239
	} else {
		goto L240
	}
L238:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v44)+72))
	if base.Ui32(v1009) <= base.Ui32(v1006-int32(4)) {
		goto L208
	} else {
		goto L242
	}
L239:
	;
	v1006 = v1000
	goto L241
L240:
	;
	v1006 = v1003
	goto L241
L241:
	;
	goto L238
L242:
	;
	goto L236
L243:
	;
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v966)+4))
	if v1031 == int32(0) {
		goto L208
	} else {
		goto L250
	}
L244:
	;
	v1014 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+88)))
	v1015 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+86)))
	if base.Ui32(v1014) < base.Ui32(v1015) {
		goto L243
	} else {
		goto L245
	}
L245:
	;
	v1017 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v958)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1017) {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v1025 = int32(base.Ui32(v1017+int32(262120)) >> (uint(int32(2)) % 32))
	goto L248
L247:
	;
	v1025 = int32(0)
	goto L248
L248:
	;
	if base.Ui32(v1014) <= base.Ui32(v1025&int32(65535)) {
		goto L208
	} else {
		goto L249
	}
L249:
	;
	goto L243
L250:
	;
	v1035 = F__bt_compare(m, l0, v916, v958, int32(1))
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	if v1035 != 0 {
		goto L208
	} else {
		goto L252
	}
L252:
	;
	v1039 = int32(4619784)
	v1040 = int32(4619776)
	v1041 = *(*int64)(unsafe.Add(mBase, _consts[58]))
	v1043 = *(*int64)(unsafe.Add(mBase, _consts[59]))
	v1044 = v1041 ^ v1043
	*(*int64)(unsafe.Add(mBase, _consts[59])) = base.I64_rotl(v1044, int64(37))
	*(*int64)(unsafe.Add(mBase, _consts[58])) = v1044<<(uint(int64(16))%64) ^ base.I64_rotl(v1041, int64(24)) ^ v1044
	goto L253
L253:
	;
	if base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_rotl(v1041*int64(5), int64(7))*int64(9))>>(uint(int64(32))%64)))) < base.Ui32(int32(42949673)) {
		goto L208
	} else {
		goto L254
	}
L254:
	;
	F__bt_stepright(m, l0, l4, v44+int32(68), v229)
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	if v1071 < int32(0) {
		goto L257
	} else {
		goto L258
	}
L256:
	;
	v1090 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1089)+16)))
	v1092 = int32(4)
	v1093 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1089)+14)))
	v1094 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1089)+12)))
	v1095 = v1093 - v1094
	if v1095 <= v1092 {
		goto L261
	} else {
		goto L262
	}
L257:
	;
	v1075 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v1075+(v1071^int32(-1))<<(uint(int32(2))%32))))
	v1089 = v1081
	goto L256
L258:
	;
	goto L259
L259:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v1089 = v1083 + v1071<<(uint(int32(13))%32) + int32(-8192)
	goto L256
L260:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v44)+72))
	if base.Ui32(v1098-int32(4)) < base.Ui32(v1101) {
		v958 = v1089
		v966 = v1090 + v1089
		goto L232
	} else {
		goto L264
	}
L261:
	;
	v1098 = v1092
	goto L263
L262:
	;
	v1098 = v1095
	goto L263
L263:
	;
	goto L260
L264:
	;
	goto L233
L265:
	;
	goto L207
L266:
	;
	v1279 = int32(4)
	v1280 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1248)+14)))
	v1281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1248)+12)))
	v1282 = v1280 - v1281
	if v1282 <= v1279 {
		goto L303
	} else {
		goto L304
	}
L267:
	;
	v1247 = l6
	v1248 = v935
	goto L266
L268:
	;
	goto L269
L269:
	;
	v1107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+86)))
	v1108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+88)))
	v1110 = l6 | base.B2i32(base.Ui32(v1107) < base.Ui32(v1108))
	v1111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+84)))
	if v1111 != int32(1) {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v944)+4))
	if v1128 == int32(0) {
		v1247 = v1110
		v1248 = v935
		goto L266
	} else {
		goto L277
	}
L271:
	;
	if base.Ui32(v1108) < base.Ui32(v1107) {
		goto L270
	} else {
		goto L272
	}
L272:
	;
	v1115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v935)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1115) {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v1123 = int32(base.Ui32(v1115+int32(262120)) >> (uint(int32(2)) % 32))
	goto L275
L274:
	;
	v1123 = int32(0)
	goto L275
L275:
	;
	if base.Ui32(v1108) <= base.Ui32(v1123&int32(65535)) {
		v1247 = v1110
		v1248 = v935
		goto L266
	} else {
		goto L276
	}
L276:
	;
	goto L270
L277:
	;
	v1132 = F__bt_compare(m, l0, v916, v935, int32(1))
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L1
	} else {
		goto L278
	}
L278:
	;
	if v1132 <= int32(0) {
		v1247 = v1110
		v1248 = v935
		goto L266
	} else {
		goto L279
	}
L279:
	;
	F__bt_stepright(m, l0, l4, v44+int32(68), v229)
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	if int32(0) <= v1140 {
		goto L282
	} else {
		goto L283
	}
L281:
	;
	v1160 = v1158
	goto L285
L282:
	;
	v1144 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v1158 = v1144 + v1140<<(uint(int32(13))%32) + int32(-8192)
	goto L281
L283:
	;
	goto L284
L284:
	;
	v1151 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(v1151+(v1140^int32(-1))<<(uint(int32(2))%32))))
	v1158 = v1157
	goto L281
L285:
	;
	v1191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1160)+16)))
	v1192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+84)))
	if v1192 != int32(1) {
		goto L288
	} else {
		goto L289
	}
L286:
	;
	v1247 = int32(1)
	v1248 = v1160
	goto L266
L287:
	;
	goto L286
L288:
	;
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(v1160+v1191)+4))
	if v1213 == int32(0) {
		goto L287
	} else {
		goto L295
	}
L289:
	;
	v1195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+88)))
	v1196 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+86)))
	if base.Ui32(v1195) < base.Ui32(v1196) {
		goto L288
	} else {
		goto L290
	}
L290:
	;
	v1198 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1160)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1198) {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v1206 = int32(base.Ui32(v1198+int32(262120)) >> (uint(int32(2)) % 32))
	goto L293
L292:
	;
	v1206 = int32(0)
	goto L293
L293:
	;
	if base.Ui32(v1195) <= base.Ui32(v1206&int32(65535)) {
		goto L287
	} else {
		goto L294
	}
L294:
	;
	goto L288
L295:
	;
	v1216 = int32(1)
	v1218 = F__bt_compare(m, l0, v916, v1160, v1216)
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L1
	} else {
		goto L296
	}
L296:
	;
	if v1218 <= int32(0) {
		v1247 = v1216
		v1248 = v1160
		goto L266
	} else {
		goto L297
	}
L297:
	;
	F__bt_stepright(m, l0, l4, v44+int32(68), v229)
	mBase = m.M
	v1225 = m.ExcPending
	if v1225 != 0 {
		goto L1
	} else {
		goto L298
	}
L298:
	;
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	if v1226 < int32(0) {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v1230 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(v1230+(v1226^int32(-1))<<(uint(int32(2))%32))))
	v1160 = v1236
	goto L285
L300:
	;
	v1238 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v1160 = v1238 + v1226<<(uint(int32(13))%32) + int32(-8192)
	goto L285
L302:
	;
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v44)+72))
	if base.Ui32(v1288) <= base.Ui32(v1285-int32(4)) {
		goto L208
	} else {
		goto L306
	}
L303:
	;
	v1285 = v1279
	goto L305
L304:
	;
	v1285 = v1282
	goto L305
L305:
	;
	goto L302
L306:
	;
	F__bt_delete_or_dedup_one_page(m, l0, l4, v44+int32(68), int32(0), v57, v1247, l6)
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	goto L208
L308:
	;
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v44)+92))
	if v1331 == int32(-1) {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	v1337 = int32(0)
	F__bt_delete_or_dedup_one_page(m, l0, l4, v44+int32(68), int32(1), v1337, v1337, v1337)
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L1
	} else {
		goto L312
	}
L310:
	;
	v1349 = v1329
	v1350 = v1331
	goto L311
L311:
	;
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	v1352 = int32(0)
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(v44)+72))
	F__bt_insertonpg(m, l0, l4, v46, v1351, v1352, v229, v34, v1353, v1349, v1350, v1352)
	mBase = m.M
	v1356 = m.ExcPending
	if v1356 != 0 {
		goto L1
	} else {
		goto L314
	}
L312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+92)) = int32(0)
	v1346 = F__bt_binsrch_insert(m, l0, v44+int32(68))
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L1
	} else {
		goto L313
	}
L313:
	;
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v44)+92))
	v1349 = v1346
	v1350 = v1348
	goto L311
L314:
	;
	goto L207
L315:
	;
	F__bt_freestack(m, v229)
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L1
	} else {
		goto L318
	}
L316:
	;
	goto L317
L317:
	;
	F_pfree(m, v46)
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L1
	} else {
		goto L319
	}
L318:
	;
	goto L317
L319:
	;
	m.G0 = v44 + int32(352)
	F_pfree(m, v34)
	mBase = m.M
	v1397 = m.ExcPending
	if v1397 != 0 {
		goto L1
	} else {
		goto L320
	}
L320:
	;
	return v879
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
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
				v29 = int32(4)
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
				if v31&int32(254) == int32(2) {
					v40 = v29
				} else {
					v40 = base.B2i32(v31 == int32(18)) << (uint(v29) % 32)
				}
				if v31 == int32(1) {
					v43 = v29
				} else {
					v43 = v40
				}
				v54 = v43
			} else {
				v44 = int32(1)
				if v25 != 0 {
					v54 = int32(base.Ui32(v23)>>(uint(v44)%32)) - v44
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v54 = int32(base.Ui32(v48)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v55 = int32(1)
			v56 = v17 + v55
			if v19&v55 != 0 {
				v61 = v56
			} else {
				v61 = v17 + int32(4)
			}
			if v19 == int32(1) {
				v64 = int32(4)
				v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
				if v66&int32(254) == int32(2) {
					v75 = v64
				} else {
					v75 = base.B2i32(v66 == int32(18)) << (uint(v64) % 32)
				}
				if v66 == int32(1) {
					v78 = v64
				} else {
					v78 = v75
				}
				v91 = v78
			} else {
				v79 = int32(1)
				if v19&v79 != 0 {
					v91 = int32(base.Ui32(v19)>>(uint(v79)%32)) - v79
				} else {
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
					v91 = int32(base.Ui32(v85)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v92 = F_varstr_cmp(m, v26, v54, v61, v91, v20)
			mBase = m.M
			v93 = m.ExcPending
			if v93 != 0 {
				return int32(0)
			} else {
				v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v94 != v10 {
					F_pfree(m, v10)
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return int32(0)
					} else {
						v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v98 != v17 {
							F_pfree(m, v17)
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int32(0)
							} else {
								return v92
							}
						} else {
							return v92
						}
					}
				} else {
					v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v98 != v17 {
						F_pfree(m, v17)
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return int32(0)
						} else {
							return v92
						}
					} else {
						return v92
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
				F_errmsg(m, int32(250667), int32(0))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					F_errhint(m, int32(588396), int32(0))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(512786), int32(1648), int32(108358))
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
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
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
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	v6 = int32(0)
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v15 = v12 << (uint(int32(5)) % 32)
	goto L3
L2:
	;
	v15 = v6
	goto L3
L3:
	;
	v16 = F_palloc0(m, v15)
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
	v25 = v6
	goto L6
L6:
	;
	v31 = int32(0)
	if l2 == v31 {
		v42 = v31
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if l3 == int32(0) {
		v51 = v31
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v36 <= v25 {
		v42 = int32(0)
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v42 = v38 + v25<<(uint(int32(2))%32)
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
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v45 <= v25 {
		v51 = v31
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v51 = v47 + v25<<(uint(int32(2))%32)
	goto L11
L14:
	;
	v80 = v16 + v25<<(uint(int32(5))%32)
	v82 = v25 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v80)+4)) = uint16(v82)
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = l1
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+8)) = v85
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+12)) = v87
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	*(*uint16)(unsafe.Add(mBase, uint32(v80)+28)) = uint16(v82)
	*(*int32)(unsafe.Add(mBase, uint32(v80)+24)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v80)+16)) = v89
	v25 = v82
	goto L6
L15:
	;
	v66 = F_palloc(m, int32(28))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L21
	}
L16:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v54 <= v25 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	if v42 == int32(0) {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	if v51 == int32(0) {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v63 = v60 + v25<<(uint(int32(2))%32)
	if v63 != 0 {
		goto L14
	} else {
		goto L20
	}
L20:
	;
	goto L15
L21:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v66)+20)) = int64(16777473)
	*(*int32)(unsafe.Add(mBase, uint32(v66)+16)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v66)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v66)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v68
	return v66
}
func F_build_backup_content(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v39 int64
	_ = v39
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v52 int64
	_ = v52
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v73 int64
	_ = v73
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int64
	_ = v83
	var v86 int64
	_ = v86
	var v95 int32
	_ = v95
	var v99 int64
	_ = v99
	var v102 int64
	_ = v102
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int64
	_ = v170
	var v175 int64
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	v8 = m.G0
	v10 = v8 - int32(528)
	m.G0 = v10
	v12 = F_makeStringInfo(m)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, _consts[145]))
		v24 = F_pg_localtime(m, l0+int32(1056), v23)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v26 = F_pg_strftime(m, v10+int32(400), int32(128), int32(520983), v24)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1032))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1040))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+192)) = v29
				v32 = int64(*(*int32)(unsafe.Add(mBase, _consts[117])))
				v33 = base.I64_div_u_s(v28, v32)
				v35 = base.I64_div_u_s(int64(4294967296), v32)
				v36 = base.I64_div_u_s(v33, v35)
				*(*uint32)(unsafe.Add(mBase, uint32(v10)+196)) = uint32(v36)
				v39 = v33 - v35*v36
				*(*uint32)(unsafe.Add(mBase, uint32(v10)+200)) = uint32(v39)
				v47 = F_pg_snprintf(m, v10+int32(336), int32(64), int32(522846), v10+int32(192))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					v49 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1032))
					*(*uint32)(unsafe.Add(mBase, uint32(v10)+180)) = uint32(v49)
					v52 = int64(base.Ui64(v49) >> (uint(int64(32)) % 64))
					*(*uint32)(unsafe.Add(mBase, uint32(v10)+176)) = uint32(v52)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v10 + int32(336)
					F_appendStringInfo(m, v12, int32(770226), v10+int32(176))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						if l1 != 0 {
							v62 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1088))
							v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1096))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+160)) = v63
							v66 = int64(*(*int32)(unsafe.Add(mBase, _consts[117])))
							v67 = base.I64_div_u_s(v62, v66)
							v69 = base.I64_div_u_s(int64(4294967296), v66)
							v70 = base.I64_div_u_s(v67, v69)
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+164)) = uint32(v70)
							v73 = v67 - v69*v70
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+168)) = uint32(v73)
							v81 = F_pg_snprintf(m, v10+int32(208), int32(64), int32(522846), v10+int32(160))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								v83 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1088))
								*(*uint32)(unsafe.Add(mBase, uint32(v10)+148)) = uint32(v83)
								v86 = int64(base.Ui64(v83) >> (uint(int64(32)) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v10)+144)) = uint32(v86)
								*(*int32)(unsafe.Add(mBase, uint32(v10)+152)) = v10 + int32(208)
								F_appendStringInfo(m, v12, int32(770263), v10+int32(144))
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return int32(0)
								} else {
									v99 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1048))
									*(*uint32)(unsafe.Add(mBase, uint32(v10)+132)) = uint32(v99)
									v102 = int64(base.Ui64(v99) >> (uint(int64(32)) % 64))
									*(*uint32)(unsafe.Add(mBase, uint32(v10)+128)) = uint32(v102)
									F_appendStringInfo(m, v12, int32(766287), v10+int32(128))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return int32(0)
									} else {
										F_appendStringInfoString(m, v12, int32(765135))
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
											return int32(0)
										} else {
											v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1064)))
											if v114 != 0 {
												v115 = int32(24067)
											} else {
												v115 = int32(18284)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v10)+112)) = v115
											F_appendStringInfo(m, v12, int32(762481), v10+int32(112))
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = v10 + int32(400)
												F_appendStringInfo(m, v12, int32(762509), v10+int32(96))
												mBase = m.M
												v129 = m.ExcPending
												if v129 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = l0
													F_appendStringInfo(m, v12, int32(762498), v10+int32(80))
													mBase = m.M
													v135 = m.ExcPending
													if v135 != 0 {
														return int32(0)
													} else {
														v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1040))
														*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v136
														F_appendStringInfo(m, v12, int32(760949), v10-int32(-64))
														mBase = m.M
														v142 = m.ExcPending
														if v142 != 0 {
															return int32(0)
														} else {
															if l1 != 0 {
																v150 = *(*int32)(unsafe.Add(mBase, _consts[145]))
																v151 = F_pg_localtime(m, l0+int32(1104), v150)
																mBase = m.M
																v152 = m.ExcPending
																if v152 != 0 {
																	return int32(0)
																} else {
																	v153 = F_pg_strftime(m, v10+int32(208), int32(128), int32(520983), v151)
																	mBase = m.M
																	v154 = m.ExcPending
																	if v154 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v10 + int32(208)
																		F_appendStringInfo(m, v12, int32(762525), v10+int32(48))
																		mBase = m.M
																		v162 = m.ExcPending
																		if v162 != 0 {
																			return int32(0)
																		} else {
																			v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1096))
																			*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v163
																			F_appendStringInfo(m, v12, int32(760969), v10+int32(32))
																			mBase = m.M
																			v169 = m.ExcPending
																			if v169 != 0 {
																				return int32(0)
																			} else {
																				v170 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1072))
																				if v170 != int64(0) {
																					*(*uint32)(unsafe.Add(mBase, uint32(v10)+20)) = uint32(v170)
																					v175 = int64(base.Ui64(v170) >> (uint(int64(32)) % 64))
																					*(*uint32)(unsafe.Add(mBase, uint32(v10)+16)) = uint32(v175)
																					F_appendStringInfo(m, v12, int32(766258), v10+int32(16))
																					mBase = m.M
																					v181 = m.ExcPending
																					if v181 != 0 {
																						return int32(0)
																					} else {
																						v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1080))
																						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v182
																						F_appendStringInfo(m, v12, int32(760923), v10)
																						mBase = m.M
																						v186 = m.ExcPending
																						if v186 != 0 {
																							return int32(0)
																						} else {
																							v187 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
																							F_pfree(m, v12)
																							mBase = m.M
																							v189 = m.ExcPending
																							if v189 != 0 {
																								return int32(0)
																							} else {
																								m.G0 = v10 + int32(528)
																								return v187
																							}
																						}
																					}
																				} else {
																					v187 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
																					F_pfree(m, v12)
																					mBase = m.M
																					v189 = m.ExcPending
																					if v189 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v10 + int32(528)
																						return v187
																					}
																				}
																			}
																		}
																	}
																}
															} else {
																v170 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1072))
																if v170 != int64(0) {
																	*(*uint32)(unsafe.Add(mBase, uint32(v10)+20)) = uint32(v170)
																	v175 = int64(base.Ui64(v170) >> (uint(int64(32)) % 64))
																	*(*uint32)(unsafe.Add(mBase, uint32(v10)+16)) = uint32(v175)
																	F_appendStringInfo(m, v12, int32(766258), v10+int32(16))
																	mBase = m.M
																	v181 = m.ExcPending
																	if v181 != 0 {
																		return int32(0)
																	} else {
																		v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1080))
																		*(*int32)(unsafe.Add(mBase, uint32(v10))) = v182
																		F_appendStringInfo(m, v12, int32(760923), v10)
																		mBase = m.M
																		v186 = m.ExcPending
																		if v186 != 0 {
																			return int32(0)
																		} else {
																			v187 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
																			F_pfree(m, v12)
																			mBase = m.M
																			v189 = m.ExcPending
																			if v189 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v10 + int32(528)
																				return v187
																			}
																		}
																	}
																} else {
																	v187 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
																	F_pfree(m, v12)
																	mBase = m.M
																	v189 = m.ExcPending
																	if v189 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v10 + int32(528)
																		return v187
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
							v99 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1048))
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+132)) = uint32(v99)
							v102 = int64(base.Ui64(v99) >> (uint(int64(32)) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+128)) = uint32(v102)
							F_appendStringInfo(m, v12, int32(766287), v10+int32(128))
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return int32(0)
							} else {
								F_appendStringInfoString(m, v12, int32(765135))
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int32(0)
								} else {
									v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1064)))
									if v114 != 0 {
										v115 = int32(24067)
									} else {
										v115 = int32(18284)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v10)+112)) = v115
									F_appendStringInfo(m, v12, int32(762481), v10+int32(112))
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = v10 + int32(400)
										F_appendStringInfo(m, v12, int32(762509), v10+int32(96))
										mBase = m.M
										v129 = m.ExcPending
										if v129 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = l0
											F_appendStringInfo(m, v12, int32(762498), v10+int32(80))
											mBase = m.M
											v135 = m.ExcPending
											if v135 != 0 {
												return int32(0)
											} else {
												v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1040))
												*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v136
												F_appendStringInfo(m, v12, int32(760949), v10-int32(-64))
												mBase = m.M
												v142 = m.ExcPending
												if v142 != 0 {
													return int32(0)
												} else {
													if l1 != 0 {
														v150 = *(*int32)(unsafe.Add(mBase, _consts[145]))
														v151 = F_pg_localtime(m, l0+int32(1104), v150)
														mBase = m.M
														v152 = m.ExcPending
														if v152 != 0 {
															return int32(0)
														} else {
															v153 = F_pg_strftime(m, v10+int32(208), int32(128), int32(520983), v151)
															mBase = m.M
															v154 = m.ExcPending
															if v154 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v10 + int32(208)
																F_appendStringInfo(m, v12, int32(762525), v10+int32(48))
																mBase = m.M
																v162 = m.ExcPending
																if v162 != 0 {
																	return int32(0)
																} else {
																	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1096))
																	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v163
																	F_appendStringInfo(m, v12, int32(760969), v10+int32(32))
																	mBase = m.M
																	v169 = m.ExcPending
																	if v169 != 0 {
																		return int32(0)
																	} else {
																		v170 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1072))
																		if v170 != int64(0) {
																			*(*uint32)(unsafe.Add(mBase, uint32(v10)+20)) = uint32(v170)
																			v175 = int64(base.Ui64(v170) >> (uint(int64(32)) % 64))
																			*(*uint32)(unsafe.Add(mBase, uint32(v10)+16)) = uint32(v175)
																			F_appendStringInfo(m, v12, int32(766258), v10+int32(16))
																			mBase = m.M
																			v181 = m.ExcPending
																			if v181 != 0 {
																				return int32(0)
																			} else {
																				v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1080))
																				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v182
																				F_appendStringInfo(m, v12, int32(760923), v10)
																				mBase = m.M
																				v186 = m.ExcPending
																				if v186 != 0 {
																					return int32(0)
																				} else {
																					v187 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
																					F_pfree(m, v12)
																					mBase = m.M
																					v189 = m.ExcPending
																					if v189 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v10 + int32(528)
																						return v187
																					}
																				}
																			}
																		} else {
																			v187 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
																			F_pfree(m, v12)
																			mBase = m.M
																			v189 = m.ExcPending
																			if v189 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v10 + int32(528)
																				return v187
																			}
																		}
																	}
																}
															}
														}
													} else {
														v170 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1072))
														if v170 != int64(0) {
															*(*uint32)(unsafe.Add(mBase, uint32(v10)+20)) = uint32(v170)
															v175 = int64(base.Ui64(v170) >> (uint(int64(32)) % 64))
															*(*uint32)(unsafe.Add(mBase, uint32(v10)+16)) = uint32(v175)
															F_appendStringInfo(m, v12, int32(766258), v10+int32(16))
															mBase = m.M
															v181 = m.ExcPending
															if v181 != 0 {
																return int32(0)
															} else {
																v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1080))
																*(*int32)(unsafe.Add(mBase, uint32(v10))) = v182
																F_appendStringInfo(m, v12, int32(760923), v10)
																mBase = m.M
																v186 = m.ExcPending
																if v186 != 0 {
																	return int32(0)
																} else {
																	v187 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
																	F_pfree(m, v12)
																	mBase = m.M
																	v189 = m.ExcPending
																	if v189 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v10 + int32(528)
																		return v187
																	}
																}
															}
														} else {
															v187 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
															F_pfree(m, v12)
															mBase = m.M
															v189 = m.ExcPending
															if v189 != 0 {
																return int32(0)
															} else {
																m.G0 = v10 + int32(528)
																return v187
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int64
	_ = v128
	var v130 int64
	_ = v130
	var v134 int64
	_ = v134
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
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
	F_errstart_cold(m, int32(21), int32(569208))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
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
	F_errstart_cold(m, int32(21), int32(569208))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v59
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+12)) = uint16(v61)
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+14)) = uint8(v63)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+79)))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+15)) = uint8(v65)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v15)+144))
	if v67 != 0 {
		goto L21
	} else {
		goto L22
	}
L9:
	;
	v59 = int32(2)
	goto L8
L10:
	;
	F_errstart_cold(m, int32(21), int32(569208))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L5
	} else {
		goto L18
	}
L11:
	;
	v59 = int32(1)
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
	v59 = int32(0)
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
	v45 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15)+79)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v45
	F_errmsg_internal(m, int32(494305), v11)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(507433), int32(2015), int32(374477))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
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
	v68 = l2
	goto L23
L22:
	;
	v68 = v67
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
	v69 = v68
	goto L26
L25:
	;
	v69 = v67
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v69
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+79)))
	switch v71 - int32(98) {
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
	if v105 != int32(1) {
		goto L39
	} else {
		goto L40
	}
L28:
	;
	v102 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+20)) = uint8(v102)
	v105 = v59
	goto L27
L29:
	;
	v86 = int32(0)
	v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+76)))
	if v87 != int32(65535) {
		v99 = v86
		v100 = v59
		goto L34
	} else {
		goto L35
	}
L30:
	;
	v74 = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v15)+92))
	if v75 == v74 {
		v84 = v74
		goto L31
	} else {
		goto L32
	}
L31:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+20)) = uint8(v84)
	v105 = v59
	goto L27
L32:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
	if v78 != int32(6179) {
		v84 = v74
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+129)))
	v84 = base.B2i32(v81 != int32(112))
	goto L31
L34:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+20)) = uint8(v99)
	v105 = v100
	goto L27
L35:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+129)))
	if v90 == int32(112) {
		v99 = v86
		v100 = v59
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v15)+132))
	v94 = F_get_base_element_type(m, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v99 = base.B2i32(v94 != int32(0))
	v100 = v98
	goto L34
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+40)) = v134
	m.G0 = v11 + int32(48)
	return v18
L39:
	;
	v130 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+28)) = v130
	v134 = v130
	goto L38
L40:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v109 == int32(2249) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v113 = F_lookup_type_cache(m, v109, int32(4352))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+13)))
	if v115 == int32(100) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v113)+300))
	v120 = F_lookup_type_cache(m, v118, int32(256))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L5
	} else {
		goto L46
	}
L44:
	;
	v122 = v113
	goto L45
L45:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+188))
	if v123 == int32(0) {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	v122 = v120
	goto L45
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = l3
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v122)+192))
	v134 = v128
	goto L38
L48:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L5
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v15 + int32(4)
	F_errmsg(m, int32(310690), v11+int32(32))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(507433), int32(1984), int32(374477))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
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
	v173 = m.ExcPending
	if v173 != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v175 = F_format_type_be(m, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v175
	F_errmsg(m, int32(357237), v11+int32(16))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(507433), int32(2066), int32(374477))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
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
	var v16 int32
	_ = v16
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
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int64
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v7 == int32(0) {
		v10 = F_expanded_record_fetch_tupdesc(m, l0)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v12 = v10
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
			if v13 == int32(0) {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v21 = F_AllocSetContextCreateInternal(m, v16, int32(62140), int32(0), int32(1024), int32(8192))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v21
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
					if v27 != 0 {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+64))
						if v28 == v26 {
							v66 = v27
							*(*int32)(unsafe.Add(mBase, uint32(v66)+28)) = int32(128)
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+32)) = v71
							*(*int32)(unsafe.Add(mBase, uint32(v66)+36)) = v71
							v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+44)) = v12
							*(*int32)(unsafe.Add(mBase, uint32(v66)+40)) = v74
							v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+68)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v66)+48)) = v77
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+84)) = v81
							v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+88)) = v83
							v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+92)) = v85
							return
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v35 = F_MemoryContextAlloc(m, v30, v26*int32(5)+int32(120))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return
							} else {
								v40 = F__emscripten_memset_bulkmem(m, v35, base.I32_extend8_s(int32(0)), int32(120))
								mBase = m.M
								v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								v43 = int32(513)
								*(*uint16)(unsafe.Add(mBase, uint32(v40)+18)) = uint16(v43)
								v45 = int32(769)
								*(*uint16)(unsafe.Add(mBase, uint32(v40)+12)) = uint16(v45)
								*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = v42
								*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = int32(1677936)
								*(*int32)(unsafe.Add(mBase, uint32(v40))) = int32(-1)
								*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v40
								*(*int32)(unsafe.Add(mBase, uint32(v40)+14)) = v40
								v54 = v40 + int32(120)
								*(*int32)(unsafe.Add(mBase, uint32(v40)+56)) = v54
								*(*int32)(unsafe.Add(mBase, uint32(v40)+24)) = int32(1384727874)
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
								*(*int32)(unsafe.Add(mBase, uint32(v40)+60)) = v54 + v58<<(uint(int32(2))%32)
								v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
								*(*int32)(unsafe.Add(mBase, uint32(v40)+64)) = v63
								*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v40
								v66 = v35
								*(*int32)(unsafe.Add(mBase, uint32(v66)+28)) = int32(128)
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								*(*int32)(unsafe.Add(mBase, uint32(v66)+32)) = v71
								*(*int32)(unsafe.Add(mBase, uint32(v66)+36)) = v71
								v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v66)+44)) = v12
								*(*int32)(unsafe.Add(mBase, uint32(v66)+40)) = v74
								v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v66)+68)) = int32(0)
								*(*int64)(unsafe.Add(mBase, uint32(v66)+48)) = v77
								v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
								*(*int32)(unsafe.Add(mBase, uint32(v66)+84)) = v81
								v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
								*(*int32)(unsafe.Add(mBase, uint32(v66)+88)) = v83
								v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
								*(*int32)(unsafe.Add(mBase, uint32(v66)+92)) = v85
								return
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v35 = F_MemoryContextAlloc(m, v30, v26*int32(5)+int32(120))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							v40 = F__emscripten_memset_bulkmem(m, v35, base.I32_extend8_s(int32(0)), int32(120))
							mBase = m.M
							v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
							v43 = int32(513)
							*(*uint16)(unsafe.Add(mBase, uint32(v40)+18)) = uint16(v43)
							v45 = int32(769)
							*(*uint16)(unsafe.Add(mBase, uint32(v40)+12)) = uint16(v45)
							*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = v42
							*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = int32(1677936)
							*(*int32)(unsafe.Add(mBase, uint32(v40))) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v40
							*(*int32)(unsafe.Add(mBase, uint32(v40)+14)) = v40
							v54 = v40 + int32(120)
							*(*int32)(unsafe.Add(mBase, uint32(v40)+56)) = v54
							*(*int32)(unsafe.Add(mBase, uint32(v40)+24)) = int32(1384727874)
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
							*(*int32)(unsafe.Add(mBase, uint32(v40)+60)) = v54 + v58<<(uint(int32(2))%32)
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
							*(*int32)(unsafe.Add(mBase, uint32(v40)+64)) = v63
							*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v40
							v66 = v35
							*(*int32)(unsafe.Add(mBase, uint32(v66)+28)) = int32(128)
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+32)) = v71
							*(*int32)(unsafe.Add(mBase, uint32(v66)+36)) = v71
							v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+44)) = v12
							*(*int32)(unsafe.Add(mBase, uint32(v66)+40)) = v74
							v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+68)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v66)+48)) = v77
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+84)) = v81
							v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+88)) = v83
							v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+92)) = v85
							return
						}
					}
				}
			} else {
				F_MemoryContextReset(m, v13)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
					if v27 != 0 {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+64))
						if v28 == v26 {
							v66 = v27
							*(*int32)(unsafe.Add(mBase, uint32(v66)+28)) = int32(128)
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+32)) = v71
							*(*int32)(unsafe.Add(mBase, uint32(v66)+36)) = v71
							v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+44)) = v12
							*(*int32)(unsafe.Add(mBase, uint32(v66)+40)) = v74
							v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+68)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v66)+48)) = v77
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+84)) = v81
							v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+88)) = v83
							v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+92)) = v85
							return
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v35 = F_MemoryContextAlloc(m, v30, v26*int32(5)+int32(120))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return
							} else {
								v40 = F__emscripten_memset_bulkmem(m, v35, base.I32_extend8_s(int32(0)), int32(120))
								mBase = m.M
								v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								v43 = int32(513)
								*(*uint16)(unsafe.Add(mBase, uint32(v40)+18)) = uint16(v43)
								v45 = int32(769)
								*(*uint16)(unsafe.Add(mBase, uint32(v40)+12)) = uint16(v45)
								*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = v42
								*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = int32(1677936)
								*(*int32)(unsafe.Add(mBase, uint32(v40))) = int32(-1)
								*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v40
								*(*int32)(unsafe.Add(mBase, uint32(v40)+14)) = v40
								v54 = v40 + int32(120)
								*(*int32)(unsafe.Add(mBase, uint32(v40)+56)) = v54
								*(*int32)(unsafe.Add(mBase, uint32(v40)+24)) = int32(1384727874)
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
								*(*int32)(unsafe.Add(mBase, uint32(v40)+60)) = v54 + v58<<(uint(int32(2))%32)
								v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
								*(*int32)(unsafe.Add(mBase, uint32(v40)+64)) = v63
								*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v40
								v66 = v35
								*(*int32)(unsafe.Add(mBase, uint32(v66)+28)) = int32(128)
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								*(*int32)(unsafe.Add(mBase, uint32(v66)+32)) = v71
								*(*int32)(unsafe.Add(mBase, uint32(v66)+36)) = v71
								v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v66)+44)) = v12
								*(*int32)(unsafe.Add(mBase, uint32(v66)+40)) = v74
								v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v66)+68)) = int32(0)
								*(*int64)(unsafe.Add(mBase, uint32(v66)+48)) = v77
								v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
								*(*int32)(unsafe.Add(mBase, uint32(v66)+84)) = v81
								v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
								*(*int32)(unsafe.Add(mBase, uint32(v66)+88)) = v83
								v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
								*(*int32)(unsafe.Add(mBase, uint32(v66)+92)) = v85
								return
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v35 = F_MemoryContextAlloc(m, v30, v26*int32(5)+int32(120))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							v40 = F__emscripten_memset_bulkmem(m, v35, base.I32_extend8_s(int32(0)), int32(120))
							mBase = m.M
							v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
							v43 = int32(513)
							*(*uint16)(unsafe.Add(mBase, uint32(v40)+18)) = uint16(v43)
							v45 = int32(769)
							*(*uint16)(unsafe.Add(mBase, uint32(v40)+12)) = uint16(v45)
							*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = v42
							*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = int32(1677936)
							*(*int32)(unsafe.Add(mBase, uint32(v40))) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v40
							*(*int32)(unsafe.Add(mBase, uint32(v40)+14)) = v40
							v54 = v40 + int32(120)
							*(*int32)(unsafe.Add(mBase, uint32(v40)+56)) = v54
							*(*int32)(unsafe.Add(mBase, uint32(v40)+24)) = int32(1384727874)
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
							*(*int32)(unsafe.Add(mBase, uint32(v40)+60)) = v54 + v58<<(uint(int32(2))%32)
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
							*(*int32)(unsafe.Add(mBase, uint32(v40)+64)) = v63
							*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v40
							v66 = v35
							*(*int32)(unsafe.Add(mBase, uint32(v66)+28)) = int32(128)
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+32)) = v71
							*(*int32)(unsafe.Add(mBase, uint32(v66)+36)) = v71
							v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+44)) = v12
							*(*int32)(unsafe.Add(mBase, uint32(v66)+40)) = v74
							v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+68)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v66)+48)) = v77
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+84)) = v81
							v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+88)) = v83
							v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+92)) = v85
							return
						}
					}
				}
			}
		}
	} else {
		v12 = v7
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
		if v13 == int32(0) {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v21 = F_AllocSetContextCreateInternal(m, v16, int32(62140), int32(0), int32(1024), int32(8192))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v21
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
				if v27 != 0 {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+64))
					if v28 == v26 {
						v66 = v27
						*(*int32)(unsafe.Add(mBase, uint32(v66)+28)) = int32(128)
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						*(*int32)(unsafe.Add(mBase, uint32(v66)+32)) = v71
						*(*int32)(unsafe.Add(mBase, uint32(v66)+36)) = v71
						v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v66)+44)) = v12
						*(*int32)(unsafe.Add(mBase, uint32(v66)+40)) = v74
						v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v66)+68)) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v66)+48)) = v77
						v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
						*(*int32)(unsafe.Add(mBase, uint32(v66)+84)) = v81
						v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
						*(*int32)(unsafe.Add(mBase, uint32(v66)+88)) = v83
						v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
						*(*int32)(unsafe.Add(mBase, uint32(v66)+92)) = v85
						return
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v35 = F_MemoryContextAlloc(m, v30, v26*int32(5)+int32(120))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							v40 = F__emscripten_memset_bulkmem(m, v35, base.I32_extend8_s(int32(0)), int32(120))
							mBase = m.M
							v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
							v43 = int32(513)
							*(*uint16)(unsafe.Add(mBase, uint32(v40)+18)) = uint16(v43)
							v45 = int32(769)
							*(*uint16)(unsafe.Add(mBase, uint32(v40)+12)) = uint16(v45)
							*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = v42
							*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = int32(1677936)
							*(*int32)(unsafe.Add(mBase, uint32(v40))) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v40
							*(*int32)(unsafe.Add(mBase, uint32(v40)+14)) = v40
							v54 = v40 + int32(120)
							*(*int32)(unsafe.Add(mBase, uint32(v40)+56)) = v54
							*(*int32)(unsafe.Add(mBase, uint32(v40)+24)) = int32(1384727874)
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
							*(*int32)(unsafe.Add(mBase, uint32(v40)+60)) = v54 + v58<<(uint(int32(2))%32)
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
							*(*int32)(unsafe.Add(mBase, uint32(v40)+64)) = v63
							*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v40
							v66 = v35
							*(*int32)(unsafe.Add(mBase, uint32(v66)+28)) = int32(128)
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+32)) = v71
							*(*int32)(unsafe.Add(mBase, uint32(v66)+36)) = v71
							v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+44)) = v12
							*(*int32)(unsafe.Add(mBase, uint32(v66)+40)) = v74
							v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+68)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v66)+48)) = v77
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+84)) = v81
							v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+88)) = v83
							v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+92)) = v85
							return
						}
					}
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v35 = F_MemoryContextAlloc(m, v30, v26*int32(5)+int32(120))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						v40 = F__emscripten_memset_bulkmem(m, v35, base.I32_extend8_s(int32(0)), int32(120))
						mBase = m.M
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
						v43 = int32(513)
						*(*uint16)(unsafe.Add(mBase, uint32(v40)+18)) = uint16(v43)
						v45 = int32(769)
						*(*uint16)(unsafe.Add(mBase, uint32(v40)+12)) = uint16(v45)
						*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = v42
						*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = int32(1677936)
						*(*int32)(unsafe.Add(mBase, uint32(v40))) = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v40
						*(*int32)(unsafe.Add(mBase, uint32(v40)+14)) = v40
						v54 = v40 + int32(120)
						*(*int32)(unsafe.Add(mBase, uint32(v40)+56)) = v54
						*(*int32)(unsafe.Add(mBase, uint32(v40)+24)) = int32(1384727874)
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						*(*int32)(unsafe.Add(mBase, uint32(v40)+60)) = v54 + v58<<(uint(int32(2))%32)
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						*(*int32)(unsafe.Add(mBase, uint32(v40)+64)) = v63
						*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v40
						v66 = v35
						*(*int32)(unsafe.Add(mBase, uint32(v66)+28)) = int32(128)
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						*(*int32)(unsafe.Add(mBase, uint32(v66)+32)) = v71
						*(*int32)(unsafe.Add(mBase, uint32(v66)+36)) = v71
						v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v66)+44)) = v12
						*(*int32)(unsafe.Add(mBase, uint32(v66)+40)) = v74
						v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v66)+68)) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v66)+48)) = v77
						v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
						*(*int32)(unsafe.Add(mBase, uint32(v66)+84)) = v81
						v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
						*(*int32)(unsafe.Add(mBase, uint32(v66)+88)) = v83
						v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
						*(*int32)(unsafe.Add(mBase, uint32(v66)+92)) = v85
						return
					}
				}
			}
		} else {
			F_MemoryContextReset(m, v13)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
				if v27 != 0 {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+64))
					if v28 == v26 {
						v66 = v27
						*(*int32)(unsafe.Add(mBase, uint32(v66)+28)) = int32(128)
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						*(*int32)(unsafe.Add(mBase, uint32(v66)+32)) = v71
						*(*int32)(unsafe.Add(mBase, uint32(v66)+36)) = v71
						v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v66)+44)) = v12
						*(*int32)(unsafe.Add(mBase, uint32(v66)+40)) = v74
						v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v66)+68)) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v66)+48)) = v77
						v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
						*(*int32)(unsafe.Add(mBase, uint32(v66)+84)) = v81
						v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
						*(*int32)(unsafe.Add(mBase, uint32(v66)+88)) = v83
						v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
						*(*int32)(unsafe.Add(mBase, uint32(v66)+92)) = v85
						return
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v35 = F_MemoryContextAlloc(m, v30, v26*int32(5)+int32(120))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							v40 = F__emscripten_memset_bulkmem(m, v35, base.I32_extend8_s(int32(0)), int32(120))
							mBase = m.M
							v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
							v43 = int32(513)
							*(*uint16)(unsafe.Add(mBase, uint32(v40)+18)) = uint16(v43)
							v45 = int32(769)
							*(*uint16)(unsafe.Add(mBase, uint32(v40)+12)) = uint16(v45)
							*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = v42
							*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = int32(1677936)
							*(*int32)(unsafe.Add(mBase, uint32(v40))) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v40
							*(*int32)(unsafe.Add(mBase, uint32(v40)+14)) = v40
							v54 = v40 + int32(120)
							*(*int32)(unsafe.Add(mBase, uint32(v40)+56)) = v54
							*(*int32)(unsafe.Add(mBase, uint32(v40)+24)) = int32(1384727874)
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
							*(*int32)(unsafe.Add(mBase, uint32(v40)+60)) = v54 + v58<<(uint(int32(2))%32)
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
							*(*int32)(unsafe.Add(mBase, uint32(v40)+64)) = v63
							*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v40
							v66 = v35
							*(*int32)(unsafe.Add(mBase, uint32(v66)+28)) = int32(128)
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+32)) = v71
							*(*int32)(unsafe.Add(mBase, uint32(v66)+36)) = v71
							v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+44)) = v12
							*(*int32)(unsafe.Add(mBase, uint32(v66)+40)) = v74
							v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+68)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v66)+48)) = v77
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+84)) = v81
							v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+88)) = v83
							v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
							*(*int32)(unsafe.Add(mBase, uint32(v66)+92)) = v85
							return
						}
					}
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v35 = F_MemoryContextAlloc(m, v30, v26*int32(5)+int32(120))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						v40 = F__emscripten_memset_bulkmem(m, v35, base.I32_extend8_s(int32(0)), int32(120))
						mBase = m.M
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
						v43 = int32(513)
						*(*uint16)(unsafe.Add(mBase, uint32(v40)+18)) = uint16(v43)
						v45 = int32(769)
						*(*uint16)(unsafe.Add(mBase, uint32(v40)+12)) = uint16(v45)
						*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = v42
						*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = int32(1677936)
						*(*int32)(unsafe.Add(mBase, uint32(v40))) = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v40
						*(*int32)(unsafe.Add(mBase, uint32(v40)+14)) = v40
						v54 = v40 + int32(120)
						*(*int32)(unsafe.Add(mBase, uint32(v40)+56)) = v54
						*(*int32)(unsafe.Add(mBase, uint32(v40)+24)) = int32(1384727874)
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						*(*int32)(unsafe.Add(mBase, uint32(v40)+60)) = v54 + v58<<(uint(int32(2))%32)
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						*(*int32)(unsafe.Add(mBase, uint32(v40)+64)) = v63
						*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v40
						v66 = v35
						*(*int32)(unsafe.Add(mBase, uint32(v66)+28)) = int32(128)
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						*(*int32)(unsafe.Add(mBase, uint32(v66)+32)) = v71
						*(*int32)(unsafe.Add(mBase, uint32(v66)+36)) = v71
						v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v66)+44)) = v12
						*(*int32)(unsafe.Add(mBase, uint32(v66)+40)) = v74
						v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v66)+68)) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v66)+48)) = v77
						v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
						*(*int32)(unsafe.Add(mBase, uint32(v66)+84)) = v81
						v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
						*(*int32)(unsafe.Add(mBase, uint32(v66)+88)) = v83
						v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
						*(*int32)(unsafe.Add(mBase, uint32(v66)+92)) = v85
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
	var v31 int32
	_ = v31
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
	var v84 int32
	_ = v84
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
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
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
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l3)+212))
	if v222 == int32(0) {
		v425 = v219
		goto L65
	} else {
		goto L66
	}
L4:
	;
	v219 = v6
	goto L3
L5:
	;
	goto L6
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v20 <= int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v219 = v6
	goto L3
L8:
	;
	goto L9
L9:
	;
	v30 = v6
	v31 = v6
	goto L10
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v31<<(uint(int32(2))%32))))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v40 = int32(0)
	if v38 == v40 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v219 = v207
	goto L3
L12:
	;
	v209 = v31 + int32(1)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v209 < v210 {
		v30 = v207
		v31 = v209
		goto L10
	} else {
		goto L64
	}
L13:
	;
	if v93 == int32(0) {
		v207 = v30
		goto L12
	} else {
		goto L27
	}
L14:
	;
	v93 = int32(1)
	goto L13
L15:
	;
	goto L16
L16:
	;
	if v39 == int32(0) {
		v84 = v40
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v93 = v84
	goto L13
L18:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v50 < v49 {
		v84 = v40
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v52 = int32(1)
	if v49 <= v52 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v55 = v52
	goto L22
L21:
	;
	v55 = v49
	goto L22
L22:
	;
	v56 = int32(8)
	v61 = int32(0)
	goto L23
L23:
	;
	v68 = v61 << (uint(int32(2)) % 32)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v38+v56+v68)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68+(v39+v56))))
	v75 = v70 & (v72 ^ int32(-1))
	v77 = base.B2i32(v75 == int32(0))
	if v75 != 0 {
		v84 = v77
		goto L17
	} else {
		goto L25
	}
L24:
	;
	v84 = v77
	goto L17
L25:
	;
	v79 = v61 + int32(1)
	if v79 != v55 {
		v61 = v79
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+11)))
	if v96 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v205 = F_list_append_unique_ptr(m, v30, v37)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L63
	}
L29:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+12)))
	if v99 != int32(1) {
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
	v103 = int32(0)
	if v102 == v103 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L31
L33:
	;
	if v156 == int32(0) {
		v207 = v30
		goto L12
	} else {
		goto L47
	}
L34:
	;
	v156 = int32(1)
	goto L33
L35:
	;
	goto L36
L36:
	;
	if v13 == int32(0) {
		v147 = v103
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v156 = v147
	goto L33
L38:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v113 < v112 {
		v147 = v103
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v115 = int32(1)
	if v112 <= v115 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v118 = v115
	goto L42
L41:
	;
	v118 = v112
	goto L42
L42:
	;
	v119 = int32(8)
	v124 = int32(0)
	goto L43
L43:
	;
	v131 = v124 << (uint(int32(2)) % 32)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v102+v119+v131)))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v131+(v13+v119))))
	v138 = v133 & (v135 ^ int32(-1))
	v140 = base.B2i32(v138 == int32(0))
	if v138 != 0 {
		v147 = v140
		goto L37
	} else {
		goto L45
	}
L44:
	;
	v147 = v140
	goto L37
L45:
	;
	v142 = v124 + int32(1)
	if v142 != v118 {
		v124 = v142
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v37)+36))
	v160 = int32(0)
	if v159 == v160 {
		v201 = v160
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if v201 != 0 {
		v207 = v30
		goto L12
	} else {
		goto L62
	}
L49:
	;
	goto L48
L50:
	;
	if v13 == int32(0) {
		v201 = v160
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v169 < v170 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v172 = v169
	goto L54
L53:
	;
	v172 = v170
	goto L54
L54:
	;
	if v172 <= int32(1) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v175 = int32(1)
	goto L57
L56:
	;
	v175 = v172
	goto L57
L57:
	;
	v176 = int32(8)
	v181 = int32(0)
	goto L58
L58:
	;
	v188 = v181 << (uint(int32(2)) % 32)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v13+v176+v188)))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v188+(v159+v176))))
	v193 = v190 & v192
	v195 = base.B2i32(v193 != int32(0))
	if v193 != 0 {
		v201 = v195
		goto L49
	} else {
		goto L60
	}
L59:
	;
	v201 = v195
	goto L49
L60:
	;
	v197 = v181 + int32(1)
	if v197 != v175 {
		v181 = v197
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	goto L28
L63:
	;
	v207 = v205
	goto L12
L64:
	;
	goto L11
L65:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v430 = F_generate_join_implied_equalities(m, l0, v428, v429, l3, l4)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L123
	}
L66:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	if v225 <= int32(0) {
		v425 = v219
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v236 = v219
	v237 = int32(0)
	goto L68
L68:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v222)+12))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v239+v237<<(uint(int32(2))%32))))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)+32))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v246 = int32(0)
	if v244 == v246 {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v425 = v413
	goto L65
L70:
	;
	v415 = v237 + int32(1)
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	if v415 < v416 {
		v236 = v413
		v237 = v415
		goto L68
	} else {
		goto L122
	}
L71:
	;
	if v299 == int32(0) {
		v413 = v236
		goto L70
	} else {
		goto L85
	}
L72:
	;
	v299 = int32(1)
	goto L71
L73:
	;
	goto L74
L74:
	;
	if v245 == int32(0) {
		v290 = v246
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v299 = v290
	goto L71
L76:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v245)+4))
	if v256 < v255 {
		v290 = v246
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v258 = int32(1)
	if v255 <= v258 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v261 = v258
	goto L80
L79:
	;
	v261 = v255
	goto L80
L80:
	;
	v262 = int32(8)
	v267 = int32(0)
	goto L81
L81:
	;
	v274 = v267 << (uint(int32(2)) % 32)
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v244+v262+v274)))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v274+(v245+v262))))
	v281 = v276 & (v278 ^ int32(-1))
	v283 = base.B2i32(v281 == int32(0))
	if v281 != 0 {
		v290 = v283
		goto L75
	} else {
		goto L83
	}
L82:
	;
	v290 = v283
	goto L75
L83:
	;
	v285 = v267 + int32(1)
	if v285 != v261 {
		v267 = v285
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+11)))
	if v302 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v411 = F_list_append_unique_ptr(m, v236, v243)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L121
	}
L87:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+12)))
	if v305 != int32(1) {
		goto L86
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v243)+32))
	v309 = int32(0)
	if v308 == v309 {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L89
L91:
	;
	if v362 == int32(0) {
		v413 = v236
		goto L70
	} else {
		goto L105
	}
L92:
	;
	v362 = int32(1)
	goto L91
L93:
	;
	goto L94
L94:
	;
	if v13 == int32(0) {
		v353 = v309
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v362 = v353
	goto L91
L96:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v319 < v318 {
		v353 = v309
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v321 = int32(1)
	if v318 <= v321 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v324 = v321
	goto L100
L99:
	;
	v324 = v318
	goto L100
L100:
	;
	v325 = int32(8)
	v330 = int32(0)
	goto L101
L101:
	;
	v337 = v330 << (uint(int32(2)) % 32)
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v308+v325+v337)))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v337+(v13+v325))))
	v344 = v339 & (v341 ^ int32(-1))
	v346 = base.B2i32(v344 == int32(0))
	if v344 != 0 {
		v353 = v346
		goto L95
	} else {
		goto L103
	}
L102:
	;
	v353 = v346
	goto L95
L103:
	;
	v348 = v330 + int32(1)
	if v348 != v324 {
		v330 = v348
		goto L101
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v243)+36))
	v366 = int32(0)
	if v365 == v366 {
		v407 = v366
		goto L107
	} else {
		goto L108
	}
L106:
	;
	if v407 != 0 {
		v413 = v236
		goto L70
	} else {
		goto L120
	}
L107:
	;
	goto L106
L108:
	;
	if v13 == int32(0) {
		v407 = v366
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v365)+4))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v375 < v376 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v378 = v375
	goto L112
L111:
	;
	v378 = v376
	goto L112
L112:
	;
	if v378 <= int32(1) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v381 = int32(1)
	goto L115
L114:
	;
	v381 = v378
	goto L115
L115:
	;
	v382 = int32(8)
	v387 = int32(0)
	goto L116
L116:
	;
	v394 = v387 << (uint(int32(2)) % 32)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v13+v382+v394)))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v394+(v365+v382))))
	v399 = v396 & v398
	v401 = base.B2i32(v399 != int32(0))
	if v399 != 0 {
		v407 = v401
		goto L107
	} else {
		goto L118
	}
L117:
	;
	v407 = v401
	goto L107
L118:
	;
	v403 = v387 + int32(1)
	if v403 != v381 {
		v387 = v403
		goto L116
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	goto L86
L121:
	;
	v413 = v411
	goto L70
L122:
	;
	goto L69
L123:
	;
	v432 = F_list_concat(m, v425, v430)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	return v432
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
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
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
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
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
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v335 int32
	_ = v335
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v384 int32
	_ = v384
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v411 int32
	_ = v411
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
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
	v153 = m.G0
	v155 = v153 - int32(16)
	m.G0 = v155
	v158 = F_palloc0(m, int32(28))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L7
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+50)))
	if v226 == int32(110) {
		goto L36
	} else {
		goto L37
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v158)+24)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v158)+16)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v158)+8)) = int64(77309411327)
	*(*int64)(unsafe.Add(mBase, uint32(v158))) = int64(4294967304)
	*(*int32)(unsafe.Add(mBase, uint32(v155)+12)) = v158
	v170 = F_palloc0(m, int32(28))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170)+24)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v170)+16)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v170)+8)) = int64(9801115369471)
	*(*int64)(unsafe.Add(mBase, uint32(v170))) = int64(4294967304)
	*(*int32)(unsafe.Add(mBase, uint32(v155)+8)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v155))) = v170
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v155)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v155)+4)) = v182
	v187 = F_list_make2_impl(m, v155+int32(4), v155)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v189 = int32(0)
	v191 = F_makeFuncExpr(m, l7, int32(2281), v187, v189, v189)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(4)))) = v191
	m.G0 = v155 + int32(16)
	v198 = l0 + int32(88)
	F_fmgr_info(m, l7, v198)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v201
	v204 = F_palloc(m, int32(36))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+220)) = v204
	*(*int32)(unsafe.Add(mBase, uint32(v204))) = v198
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	*(*int32)(unsafe.Add(mBase, uint32(v208)+4)) = l1
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	v211 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v210)+8)) = v211
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	*(*int32)(unsafe.Add(mBase, uint32(v213)+12)) = v211
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	*(*uint8)(unsafe.Add(mBase, uint32(v216)+16)) = uint8(v211)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	v220 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v219)+18)) = uint16(v220)
	goto L22
L29:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	if v392 != 0 {
		goto L68
	} else {
		goto L69
	}
L30:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	v278 = F_ExecTypeFromTL(m, v277)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L7
	} else {
		goto L47
	}
L31:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	if v268 == int32(0) {
		v384 = v266
		goto L29
	} else {
		goto L46
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v255
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)) = uint8(v254)
	if v255 <= int32(0) {
		v265 = v255
		v266 = v256
		v267 = v257
		goto L31
	} else {
		goto L45
	}
L33:
	;
	v247 = int32(0)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l3)+36))
	if v249 != 0 {
		goto L42
	} else {
		goto L43
	}
L34:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	v254 = v230 ^ int32(1)
	v255 = v246
	v256 = v246
	v257 = v229
	goto L32
L35:
	;
	if v229 == int32(0) {
		goto L33
	} else {
		goto L41
	}
L36:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+51)))
	if v230 != int32(1) {
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
	v237 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)) = uint8(v237)
	v265 = v237
	v266 = v237
	v267 = v237
	goto L31
L39:
	;
	if v229 != 0 {
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
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)+4))
	v251 = v250
	goto L44
L43:
	;
	v251 = v247
	goto L44
L44:
	;
	v254 = base.B2i32(int32(0) < v251)
	v255 = v251
	v256 = v247
	v257 = v249
	goto L32
L45:
	;
	v273 = v255
	v274 = v256
	v275 = v257
	v276 = int32(1)
	goto L30
L46:
	;
	v273 = v265
	v274 = v266
	v275 = v267
	v276 = int32(0)
	goto L30
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v278
	v282 = F_ExecInitExtraTupleSlot(m, l2, v278, int32(1632252))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v282
	if v276 == int32(0) {
		v384 = v274
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
	v308 = F_palloc(m, v273<<(uint(int32(1))%32))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L7
	} else {
		goto L57
	}
L51:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l10+v39<<(uint(int32(2))%32))))
	F_get_typlenbyval(m, v292, l0+int32(182), l0+int32(186))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L7
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if v274 <= int32(0) {
		goto L50
	} else {
		goto L55
	}
L54:
	;
	goto L50
L55:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v303 = F_ExecInitExtraTupleSlot(m, l2, v301, int32(1632252))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v303
	goto L50
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v308
	v312 = v273 << (uint(int32(2)) % 32)
	v313 = F_palloc(m, v312)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L7
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v313
	v316 = F_palloc(m, v312)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L7
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v316
	v319 = F_palloc(m, v273)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v319
	if v275 == int32(0) {
		v384 = v274
		goto L29
	} else {
		goto L61
	}
L61:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	if v324 <= int32(0) {
		v384 = v274
		goto L29
	} else {
		goto L62
	}
L62:
	;
	v335 = int32(0)
	goto L63
L63:
	;
	v345 = v335 << (uint(int32(2)) % 32)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v275)+12))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v345+v346)))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	v350 = F_get_sortgroupclause_tle(m, v348, v349)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L7
	} else {
		goto L65
	}
L64:
	;
	v384 = v274
	goto L29
L65:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v356 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v350)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v352+v335<<(uint(int32(1))%32)))) = uint16(v356)
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v348)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v358+v345))) = v360
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v350)+4))
	v363 = F_exprCollation(m, v362)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v365+v345))) = v363
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v368+v335))) = uint8(v370)
	v373 = v335 + int32(1)
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	if v373 < v374 {
		v335 = v373
		goto L63
	} else {
		goto L67
	}
L67:
	;
	goto L64
L68:
	;
	v395 = F_palloc(m, v384<<(uint(int32(2))%32))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L7
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v481 = int32(1)
	if v21 <= v481 {
		goto L86
	} else {
		goto L87
	}
L71:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	if v397 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	if v384 == int32(1) {
		goto L79
	} else {
		goto L80
	}
L73:
	;
	v400 = int32(0)
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v397)+4))
	if v401 <= v400 {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v411 = v400
	goto L75
L75:
	;
	v421 = v411 << (uint(int32(2)) % 32)
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v397)+12))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v423+v421)))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v425)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v395+v421))) = v426
	v429 = v411 + int32(1)
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v397)+4))
	if v429 < v430 {
		v411 = v429
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
	F_pfree(m, v395)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L7
	} else {
		goto L85
	}
L79:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v395)))
	v451 = F_get_opcode(m, v450)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L7
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v460 = F_execTuplesMatchPrepare(m, v457, v384, v458, v395, v459, l1)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L7
	} else {
		goto L84
	}
L82:
	;
	F_fmgr_info(m, v451, l0+int32(144))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L7
	} else {
		goto L83
	}
L83:
	;
	goto L78
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = v460
	goto L78
L85:
	;
	goto L70
L86:
	;
	v484 = v481
	goto L88
L87:
	;
	v484 = v21
	goto L88
L88:
	;
	v487 = F_palloc0(m, v484<<(uint(int32(2))%32))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L7
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v487
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
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
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
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
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
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
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v47 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v19 = int32(4)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v21&int32(254) == int32(2) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v34 = int32(1)
	if v16&v34 != 0 {
		v46 = int32(base.Ui32(v16)>>(uint(v34)%32)) - v34
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v30 = v19
	goto L10
L9:
	;
	v30 = base.B2i32(v21 == int32(18)) << (uint(v19) % 32)
	goto L10
L10:
	;
	if v21 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v33 = v19
	goto L13
L12:
	;
	v33 = v30
	goto L13
L13:
	;
	v46 = v33
	goto L4
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v78 = int32(1)
	if v16&v78 != 0 {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v50 = int32(4)
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v52&int32(254) == int32(2) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v65 = int32(1)
	if v47&v65 != 0 {
		v77 = int32(base.Ui32(v47)>>(uint(v65)%32)) - v65
		goto L15
	} else {
		goto L25
	}
L19:
	;
	v61 = v50
	goto L21
L20:
	;
	v61 = base.B2i32(v52 == int32(18)) << (uint(v50) % 32)
	goto L21
L21:
	;
	if v52 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v64 = v50
	goto L24
L23:
	;
	v64 = v61
	goto L24
L24:
	;
	v77 = v64
	goto L15
L25:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v77 = int32(base.Ui32(v71)>>(uint(int32(2))%32)) - int32(4)
	goto L15
L26:
	;
	v82 = v78
	goto L28
L27:
	;
	v82 = int32(4)
	goto L28
L28:
	;
	v83 = v9 + v82
	v84 = int32(1)
	if v47&v84 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v88 = v84
	goto L31
L30:
	;
	v88 = int32(4)
	goto L31
L31:
	;
	v89 = v14 + v88
	if v46 < v77 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v91 = v46
	goto L34
L33:
	;
	v91 = v77
	goto L34
L34:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v91) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v154 != v9 {
		goto L53
	} else {
		goto L54
	}
L36:
	;
	v153 = int32(0)
	goto L35
L37:
	;
	v127 = v122
	v128 = v123
	v129 = v124
	goto L47
L38:
	;
	if (v83|v89)&int32(3) != 0 {
		v122 = v83
		v123 = v89
		v124 = v91
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v115 = v83
	v116 = v89
	v117 = v91
	goto L40
L40:
	;
	if v117 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v99 = v83
	v100 = v89
	v101 = v91
	goto L42
L42:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v104 != v105 {
		v122 = v99
		v123 = v100
		v124 = v101
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v115 = v110
	v116 = v108
	v117 = v112
	goto L40
L44:
	;
	v107 = int32(4)
	v108 = v100 + v107
	v110 = v99 + v107
	v112 = v101 - v107
	if base.Ui32(int32(3)) < base.Ui32(v112) {
		v99 = v110
		v100 = v108
		v101 = v112
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v122 = v115
	v123 = v116
	v124 = v117
	goto L37
L47:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	if v132 == v133 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v153 = v132 - v133
	goto L35
L49:
	;
	v135 = int32(1)
	v140 = v129 - v135
	if v140 != 0 {
		v127 = v127 + v135
		v128 = v128 + v135
		v129 = v140
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
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v158 != v14 {
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
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v162 = int32(0)
	return base.B2i32(v153 == v162)&base.B2i32(v77 <= v46) | base.B2i32(v162 < v153)
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
