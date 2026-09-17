package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetLocalVictimBuffer(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v233 int32
	_ = v233
	var v237 int64
	_ = v237
	var v241 int64
	_ = v241
	var v252 int32
	_ = v252
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[0]))
	F_ResourceOwnerEnlarge(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[1]))
	v18 = v16
	goto L6
L3:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v52)+24))
	if v213&int32(_a_F_GetLocalVictimBuffer_0) != 0 {
		goto L45
	} else {
		goto L46
	}
L4:
	;
	v193 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[2])) = v191 + v193
	v196 = int32(_a_F_GetLocalVictimBuffer_1)
	v198 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[3])) = v198 + v193
	*(*int32)(unsafe.Add(mBase, uint32(v189+v190<<(uint(int32(2))%32)))) = v188 + v191<<(uint(int32(13))%32)
	goto L3
L5:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[4]))
	if v134 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[5]))
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[6]))
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[7]))
	v32 = v18
	v37 = v30
	goto L8
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L27
	}
L8:
	;
	v40 = v32 + int32(1)
	if v40 < v30 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	v43 = v40
	goto L12
L11:
	;
	v43 = int32(0)
	goto L12
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v26+v32<<(uint(int32(2))%32))))
	if v47 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v52 = v28 + v32<<(uint(int32(6))%32)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+24))
	if v53&int32(_a_F_GetLocalVictimBuffer_2) != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[1])) = v43
	v116 = v37 - int32(1)
	if v116 != 0 {
		v32 = v43
		v37 = v116
		goto L8
	} else {
		goto L26
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[1])) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v52)+24)) = v53 - int32(_a_F_GetLocalVictimBuffer_3)
	v18 = v43
	goto L6
L17:
	;
	goto L18
L18:
	;
	if v53&int32(_a_F_GetLocalVictimBuffer_4) != 0 {
		v32 = v43
		goto L8
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[1])) = v43
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v52)+24))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
	v71 = v26 + (int32(-2)-v67)<<(uint(int32(2))%32)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	if v72 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v75 = int32(_a_F_GetLocalVictimBuffer_5)
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[8]))
	v78 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[8])) = v77 + v78
	*(*int32)(unsafe.Add(mBase, uint32(v52)+24)) = v65 + v78
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v85 = v84
	goto L22
L21:
	;
	v85 = v72
	goto L22
L22:
	;
	v86 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v85 + v86
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[0]))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
	F_ResourceOwnerRemember(m, v90, v91+v86, int32(_a_F_GetLocalVictimBuffer_6))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[9]))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
	v101 = int32(-2) - v100
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v98+v101<<(uint(int32(2))%32))))
	if v105 != 0 {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[2]))
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[10]))
	if v109 <= v107 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[11]))
	v188 = v112
	v189 = v98
	v190 = v101
	v191 = v107
	goto L4
L26:
	;
	goto L9
L27:
	;
	F_errcode(m, int32(197))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errmsg(m, int32(_a_F_GetLocalVictimBuffer_7), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_GetLocalVictimBuffer_8), int32(272), int32(_a_F_GetLocalVictimBuffer_9))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[12]))
	v144 = F_AllocSetContextCreateInternal(m, v139, int32(_a_F_GetLocalVictimBuffer_10), int32(0), int32(_a_F_GetLocalVictimBuffer_11), int32(_a_F_GetLocalVictimBuffer_0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	v149 = v109
	v150 = v134
	goto L33
L33:
	;
	v153 = int32(16)
	v155 = v149 << (uint(int32(1)) % 32)
	if v155 <= v153 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[4])) = v144
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[10]))
	v149 = v148
	v150 = v144
	goto L33
L35:
	;
	v158 = v153
	goto L37
L36:
	;
	v158 = v155
	goto L37
L37:
	;
	v160 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[7]))
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[3]))
	v163 = v160 - v162
	if v158 < v163 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v165 = v158
	goto L40
L39:
	;
	v165 = v163
	goto L40
L40:
	;
	if base.Ui32(int32(_a_F_GetLocalVictimBuffer_12)) <= base.Ui32(v165) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v168 = int32(_a_F_GetLocalVictimBuffer_12)
	goto L43
L42:
	;
	v168 = v165
	goto L43
L43:
	;
	v173 = F_MemoryContextAlloc(m, v150, v168<<(uint(int32(13))%32)|int32(_a_F_GetLocalVictimBuffer_13))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[10])) = v168
	v181 = (v173 + int32(4095)) & int32(-4096)
	*(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[11])) = v181
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
	v187 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[9]))
	v188 = v181
	v189 = v187
	v190 = int32(-2) - v184
	v191 = int32(0)
	goto L4
L45:
	;
	F_FlushLocalBuffer(m, v52, int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v52)+24))
	if v219&int32(33554432) != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L47
L49:
	;
	F_InvalidateLocalBuffer(m, v52, int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
	return v252 + int32(1)
L52:
	;
	v225 = int32(1)
	v233 = int32(512)
	v237 = *(*int64)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[13]))
	*(*int64)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[13])) = v237 + int64(1)
	v241 = *(*int64)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[14]))
	*(*int64)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[14])) = v241
	F_pgstat_count_backend_io_op(m, v225, int32(3), int32(0), v225, int64(0))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[15])) = uint8(v225)
	*(*uint8)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[16])) = uint8(v225)
	goto L53
L53:
	;
	goto L51
}
func F_GetNextLocalTransactionId(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v2 = int32(_a_F_GetNextLocalTransactionId_0)
	v3 = int32(1)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_GetNextLocalTransactionId[0]))
	if base.Ui32(v5) <= base.Ui32(v3) {
		v8 = v3
	} else {
		v8 = v5
	}
	*(*int32)(unsafe.Add(mBase, _c_F_GetNextLocalTransactionId[0])) = v8 + int32(1)
	return v8
}
func F_LocalToUtf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v539 int32
	_ = v539
	var v547 int32
	_ = v547
	v18 = m.G0
	v20 = v18 - int32(32)
	m.G0 = v20
	if base.Ui32(l7) <= base.Ui32(int32(41)) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v547 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v532))) = uint8(v547)
	m.G0 = v20 + int32(32)
	return v539 - l0
L2:
	;
	v47 = l1
	v48 = l2
	v55 = l0
	goto L12
L3:
	;
	if int32(0) < l1 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v532 = l2
	v539 = l0
	goto L1
L7:
	;
	return int32(0)
L8:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = l7
	F_errmsg(m, int32(_a_F_LocalToUtf_0), v20+int32(16))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	F_errfinish(m, int32(_a_F_LocalToUtf_1), int32(733), int32(_a_F_LocalToUtf_2))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L12:
	;
	v63 = int32(*(*int8)(unsafe.Add(mBase, uint32(v55))))
	if v63 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v532 = v518
	v539 = v526
	goto L1
L14:
	;
	v527 = v47 - v521
	if int32(0) < v527 {
		v47 = v527
		v48 = v518
		v55 = v526
		goto L12
	} else {
		goto L127
	}
L15:
	;
	v518 = v513 + int32(1)
	v521 = v74
	v526 = v123
	goto L14
L16:
	;
	if l8 != 0 {
		v532 = v48
		v539 = v55
		goto L1
	} else {
		goto L125
	}
L17:
	;
	if int32(0) <= v63 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v48))) = uint8(v63)
	v69 = int32(1)
	v518 = v48 + v69
	v521 = v69
	v526 = v55 + v69
	goto L14
L19:
	;
	goto L20
L20:
	;
	v74 = F_pg_encoding_verifymbchar(m, l7, v55, v47)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	if v74 < int32(0) {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	v78 = int32(0)
	switch v74 - int32(1) {
	case 0:
		v108 = v55
		v109 = v78
		v110 = v78
		v111 = v78
		goto L23
	case 1:
		goto L24
	case 2:
		goto L27
	case 3:
		goto L26
	default:
		goto L25
	}
L23:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v112 | (v110<<(uint(int32(16))%32) | v111<<(uint(int32(24))%32) | v109<<(uint(int32(8))%32))
	v123 = v55 + v74
	if l3 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L24:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	v108 = v55 + int32(1)
	v109 = v107
	v110 = v78
	v111 = v78
	goto L23
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L7
	} else {
		goto L28
	}
L26:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+2)))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	v108 = v55 + int32(3)
	v109 = v89
	v110 = v90
	v111 = v91
	goto L23
L27:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	v108 = v55 + int32(2)
	v109 = v85
	v110 = v86
	v111 = v78
	goto L23
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v74
	F_errmsg_internal(m, int32(_a_F_LocalToUtf_3), v20)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_LocalToUtf_1), int32(781), int32(_a_F_LocalToUtf_2))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	if l6 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L32:
	;
	v126 = int32(0)
	switch v74 - int32(1) {
	case 0:
		goto L35
	case 1:
		goto L36
	case 2:
		goto L37
	case 3:
		goto L38
	default:
		v357 = v126
		goto L34
	}
L33:
	;
	if v368 != 0 {
		goto L71
	} else {
		goto L72
	}
L34:
	;
	v368 = v357
	goto L33
L35:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+12)))
	if base.Ui32(v112) < base.Ui32(v328) {
		v357 = v126
		goto L34
	} else {
		goto L66
	}
L36:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+20)))
	if base.Ui32(v109) < base.Ui32(v283) {
		v357 = v126
		goto L34
	} else {
		goto L59
	}
L37:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+28)))
	if base.Ui32(v110) < base.Ui32(v218) {
		v357 = v126
		goto L34
	} else {
		goto L50
	}
L38:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+40)))
	if base.Ui32(v111) < base.Ui32(v133) {
		v357 = v126
		goto L34
	} else {
		goto L39
	}
L39:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+41)))
	if base.Ui32(v135) < base.Ui32(v111) {
		v357 = v126
		goto L34
	} else {
		goto L40
	}
L40:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+42)))
	if base.Ui32(v110) < base.Ui32(v137) {
		v357 = v126
		goto L34
	} else {
		goto L41
	}
L41:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+43)))
	if base.Ui32(v139) < base.Ui32(v110) {
		v357 = v126
		goto L34
	} else {
		goto L42
	}
L42:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+44)))
	if base.Ui32(v109) < base.Ui32(v141) {
		v357 = v126
		goto L34
	} else {
		goto L43
	}
L43:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+45)))
	if base.Ui32(v143) < base.Ui32(v109) {
		v357 = v126
		goto L34
	} else {
		goto L44
	}
L44:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+46)))
	if base.Ui32(v112) < base.Ui32(v145) {
		v357 = v126
		goto L34
	} else {
		goto L45
	}
L45:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+47)))
	if base.Ui32(v147) < base.Ui32(v112) {
		v357 = v126
		goto L34
	} else {
		goto L46
	}
L46:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l3)+36))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v150 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v152 = int32(2)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v150+(v111-v133)<<(uint(v152)%32)+v149<<(uint(v152)%32))))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v150+(v110-v137)<<(uint(v152)%32)+v170<<(uint(v152)%32))))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v150+(v109-v141)<<(uint(v152)%32)+v174<<(uint(v152)%32))))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v150+(v112-v145)<<(uint(v152)%32)+v178<<(uint(v152)%32))))
	v368 = v182
	goto L33
L48:
	;
	goto L49
L49:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v185 = int32(1)
	v205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183+(v111-v133)<<(uint(v185)%32)+v149&int32(_a_F_LocalToUtf_4)<<(uint(v185)%32)))))
	v209 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183+(v110-v137)<<(uint(v185)%32)+v205<<(uint(v185)%32)))))
	v213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183+(v109-v141)<<(uint(v185)%32)+v209<<(uint(v185)%32)))))
	v217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183+(v112-v145)<<(uint(v185)%32)+v213<<(uint(v185)%32)))))
	v368 = v217
	goto L33
L50:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+29)))
	if base.Ui32(v220) < base.Ui32(v110) {
		v357 = v126
		goto L34
	} else {
		goto L51
	}
L51:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+30)))
	if base.Ui32(v109) < base.Ui32(v222) {
		v357 = v126
		goto L34
	} else {
		goto L52
	}
L52:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+31)))
	if base.Ui32(v224) < base.Ui32(v109) {
		v357 = v126
		goto L34
	} else {
		goto L53
	}
L53:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+32)))
	if base.Ui32(v112) < base.Ui32(v226) {
		v357 = v126
		goto L34
	} else {
		goto L54
	}
L54:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+33)))
	if base.Ui32(v228) < base.Ui32(v112) {
		v357 = v126
		goto L34
	} else {
		goto L55
	}
L55:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v231 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v233 = int32(2)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v231+(v110-v218)<<(uint(v233)%32)+v230<<(uint(v233)%32))))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v231+(v109-v222)<<(uint(v233)%32)+v247<<(uint(v233)%32))))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v231+(v112-v226)<<(uint(v233)%32)+v251<<(uint(v233)%32))))
	v368 = v255
	goto L33
L57:
	;
	goto L58
L58:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v258 = int32(1)
	v274 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v256+(v110-v218)<<(uint(v258)%32)+v230&int32(_a_F_LocalToUtf_4)<<(uint(v258)%32)))))
	v278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v256+(v109-v222)<<(uint(v258)%32)+v274<<(uint(v258)%32)))))
	v282 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v256+(v112-v226)<<(uint(v258)%32)+v278<<(uint(v258)%32)))))
	v368 = v282
	goto L33
L59:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+21)))
	if base.Ui32(v285) < base.Ui32(v109) {
		v357 = v126
		goto L34
	} else {
		goto L60
	}
L60:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+22)))
	if base.Ui32(v112) < base.Ui32(v287) {
		v357 = v126
		goto L34
	} else {
		goto L61
	}
L61:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+23)))
	if base.Ui32(v289) < base.Ui32(v112) {
		v357 = v126
		goto L34
	} else {
		goto L62
	}
L62:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v292 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v294 = int32(2)
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v292+(v109-v283)<<(uint(v294)%32)+v291<<(uint(v294)%32))))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v292+(v112-v287)<<(uint(v294)%32)+v304<<(uint(v294)%32))))
	v368 = v308
	goto L33
L64:
	;
	goto L65
L65:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v311 = int32(1)
	v323 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v309+(v109-v283)<<(uint(v311)%32)+v291&int32(_a_F_LocalToUtf_4)<<(uint(v311)%32)))))
	v327 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v309+(v112-v287)<<(uint(v311)%32)+v323<<(uint(v311)%32)))))
	v368 = v327
	goto L33
L66:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+13)))
	if base.Ui32(v330) < base.Ui32(v112) {
		v357 = v126
		goto L34
	} else {
		goto L67
	}
L67:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v332 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v334 = int32(2)
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v332+(v112-v328)<<(uint(v334)%32)+v337<<(uint(v334)%32))))
	v368 = v341
	goto L33
L69:
	;
	goto L70
L70:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v344 = int32(1)
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v351 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v342+(v112-v328)<<(uint(v344)%32)+v347<<(uint(v344)%32)))))
	v357 = v351
	goto L34
L71:
	;
	if base.Ui32(int32(16777216)) <= base.Ui32(v368) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L73
L73:
	;
	if l4 == int32(0) {
		goto L31
	} else {
		goto L84
	}
L74:
	;
	v372 = int32(base.Ui32(v368) >> (uint(int32(24)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v48))) = uint8(v372)
	v376 = v48 + int32(1)
	goto L76
L75:
	;
	v376 = v48
	goto L76
L76:
	;
	if v368&int32(16711680) != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v380 = int32(base.Ui32(v368) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v376))) = uint8(v380)
	v384 = v376 + int32(1)
	goto L79
L78:
	;
	v384 = v376
	goto L79
L79:
	;
	if v368&int32(_a_F_LocalToUtf_5) != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v388 = int32(base.Ui32(v368) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v384))) = uint8(v388)
	v392 = v384 + int32(1)
	goto L82
L81:
	;
	v392 = v384
	goto L82
L82:
	;
	if v368&int32(255) == int32(0) {
		v518 = v392
		v521 = v74
		v526 = v123
		goto L14
	} else {
		goto L83
	}
L83:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v392))) = uint8(v368)
	v513 = v392
	goto L15
L84:
	;
	v404 = F_bsearch(m, v20+int32(28), l4, l5, int32(12), int32(1635))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L7
	} else {
		goto L85
	}
L85:
	;
	if v404 == int32(0) {
		goto L31
	} else {
		goto L86
	}
L86:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v404)+4))
	if base.Ui32(int32(16777216)) <= base.Ui32(v408) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v412 = int32(base.Ui32(v408) >> (uint(int32(24)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v48))) = uint8(v412)
	v416 = v48 + int32(1)
	goto L89
L88:
	;
	v416 = v48
	goto L89
L89:
	;
	if v408&int32(16711680) != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v420 = int32(base.Ui32(v408) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v416))) = uint8(v420)
	v424 = v416 + int32(1)
	goto L92
L91:
	;
	v424 = v416
	goto L92
L92:
	;
	if v408&int32(_a_F_LocalToUtf_5) != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v428 = int32(base.Ui32(v408) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v424))) = uint8(v428)
	v432 = v424 + int32(1)
	goto L95
L94:
	;
	v432 = v424
	goto L95
L95:
	;
	if v408&int32(255) != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v432))) = uint8(v408)
	v438 = v432 + int32(1)
	goto L98
L97:
	;
	v438 = v432
	goto L98
L98:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v404)+8))
	if base.Ui32(int32(16777216)) <= base.Ui32(v439) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v443 = int32(base.Ui32(v439) >> (uint(int32(24)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v438))) = uint8(v443)
	v447 = v438 + int32(1)
	goto L101
L100:
	;
	v447 = v438
	goto L101
L101:
	;
	if v439&int32(16711680) != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v451 = int32(base.Ui32(v439) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v447))) = uint8(v451)
	v455 = v447 + int32(1)
	goto L104
L103:
	;
	v455 = v447
	goto L104
L104:
	;
	if v439&int32(_a_F_LocalToUtf_5) != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v459 = int32(base.Ui32(v439) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v455))) = uint8(v459)
	v463 = v455 + int32(1)
	goto L107
L106:
	;
	v463 = v455
	goto L107
L107:
	;
	if v439&int32(255) == int32(0) {
		v518 = v463
		v521 = v74
		v526 = v123
		goto L14
	} else {
		goto L108
	}
L108:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v463))) = uint8(v439)
	v513 = v463
	goto L15
L109:
	;
	if l8 != 0 {
		v532 = v48
		v539 = v55
		goto L1
	} else {
		goto L123
	}
L110:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v473 = m.T0[l6].(func(*base.Module, int32) int32)(m, v472)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L7
	} else {
		goto L111
	}
L111:
	;
	if v473 == int32(0) {
		goto L109
	} else {
		goto L112
	}
L112:
	;
	if base.Ui32(int32(16777216)) <= base.Ui32(v473) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v480 = int32(base.Ui32(v473) >> (uint(int32(24)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v48))) = uint8(v480)
	v484 = v48 + int32(1)
	goto L115
L114:
	;
	v484 = v48
	goto L115
L115:
	;
	if v473&int32(16711680) != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v488 = int32(base.Ui32(v473) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v484))) = uint8(v488)
	v492 = v484 + int32(1)
	goto L118
L117:
	;
	v492 = v484
	goto L118
L118:
	;
	if v473&int32(_a_F_LocalToUtf_5) != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v496 = int32(base.Ui32(v473) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v492))) = uint8(v496)
	v500 = v492 + int32(1)
	goto L121
L120:
	;
	v500 = v492
	goto L121
L121:
	;
	if v473&int32(255) == int32(0) {
		v518 = v500
		v521 = v74
		v526 = v123
		goto L14
	} else {
		goto L122
	}
L122:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v500))) = uint8(v473)
	v513 = v500
	goto L15
L123:
	;
	F_report_untranslatable_char(m, l7, int32(6), v55, v47)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L7
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	F_report_invalid_encoding(m, l7, v55, v47)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L7
	} else {
		goto L126
	}
L126:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L127:
	;
	goto L13
}
func F_UnpinLocalBuffer(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_UnpinLocalBuffer[0]))
	v7 = l0 ^ int32(-1)
	v10 = v5 + v7<<(uint(int32(2))%32)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v13 = v11 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v13
	if v13 == int32(0) {
		v17 = int32(_a_F_UnpinLocalBuffer_0)
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_UnpinLocalBuffer[1]))
		v20 = int32(1)
		*(*int32)(unsafe.Add(mBase, _c_F_UnpinLocalBuffer[1])) = v19 - v20
		v24 = *(*int32)(unsafe.Add(mBase, _c_F_UnpinLocalBuffer[2]))
		v27 = v24 + v7<<(uint(int32(6))%32)
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v27)+24)) = v28 - v20
	} else {
	}
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_UnpinLocalBuffer[3]))
	F_ResourceOwnerForget(m, v34, l0, int32(_a_F_UnpinLocalBuffer_1))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		return
	} else {
		return
	}
}
