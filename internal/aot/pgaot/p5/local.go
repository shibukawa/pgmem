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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v58 int64
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v82 int64
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v230 int64
	_ = v230
	var v233 int64
	_ = v233
	var v240 int32
	_ = v240
	var v241 int64
	_ = v241
	var v244 int64
	_ = v244
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v260 int32
	_ = v260
	var v264 int64
	_ = v264
	var v268 int64
	_ = v268
	var v279 int32
	_ = v279
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
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[2]))
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[3]))
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[4]))
	v25 = v20
	v27 = v16
	v28 = v22
	v29 = v18
	goto L6
L3:
	;
	v230 = int64(0)
	v233 = base.AtomicRmwCmpxchg64(m, v54, int32(24), v230, v230)
	if v233&int64(8388608) != v230 {
		goto L45
	} else {
		goto L46
	}
L4:
	;
	v209 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[5])) = v207 + v209
	v212 = int32(_a_F_GetLocalVictimBuffer_0)
	v214 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[6])) = v214 + v209
	*(*int32)(unsafe.Add(mBase, uint32(v204+v205<<(uint(int32(2))%32)))) = v206 + v207<<(uint(int32(13))%32)
	goto L3
L5:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[7]))
	if v153 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L6:
	;
	v33 = v25
	v34 = v28
	v35 = v27
	v36 = v28
	v37 = v29
	goto L8
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[3])) = v43
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L27
	}
L8:
	;
	v40 = v33 + int32(1)
	if v40 < v36 {
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
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v35+v33<<(uint(int32(2))%32))))
	if v47 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v50 = int32(_a_F_GetLocalVictimBuffer_1)
	*(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[3])) = v43
	v54 = v37 + v33*int32(56)
	v55 = int64(0)
	v58 = base.AtomicRmwCmpxchg64(m, v54, int32(24), v55, v55)
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[1]))
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[2]))
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[4]))
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[3]))
	if v58&int64(3932160) != v55 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v133 = v34 - int32(1)
	if v133 != 0 {
		v33 = v43
		v34 = v133
		goto L8
	} else {
		goto L26
	}
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v54)+24)) = v58 - int64(262144)
	v25 = v66
	v27 = v60
	v28 = v64
	v29 = v62
	goto L6
L17:
	;
	goto L18
L18:
	;
	if v58&int64(262143) != int64(0) {
		v33 = v66
		v35 = v60
		v36 = v64
		v37 = v62
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	v79 = int64(0)
	v82 = base.AtomicRmwCmpxchg64(m, v54, int32(24), v79, v79)
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[1]))
	v89 = v84 + (int32(-2)-v78)<<(uint(int32(2))%32)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	if v90 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v93 = int32(_a_F_GetLocalVictimBuffer_2)
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[8])) = v95 + int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v54)+24)) = v82 + int64(1)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v103 = v102
	goto L22
L21:
	;
	v103 = v90
	goto L22
L22:
	;
	v104 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = v103 + v104
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[0]))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	F_ResourceOwnerRemember(m, v108, base.I64_extend_i32_s(v109+v104), int32(_a_F_GetLocalVictimBuffer_3))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[9]))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	v120 = int32(-2) - v119
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v117+v120<<(uint(int32(2))%32))))
	if v124 != 0 {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[5]))
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[10]))
	if v128 <= v126 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[11]))
	v204 = v117
	v205 = v120
	v206 = v131
	v207 = v126
	goto L4
L26:
	;
	goto L9
L27:
	;
	F_errcode(m, int32(197))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errmsg(m, int32(_a_F_GetLocalVictimBuffer_4), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_GetLocalVictimBuffer_5), int32(274), int32(_a_F_GetLocalVictimBuffer_6))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
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
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[12]))
	v163 = F_AllocSetContextCreateInternal(m, v158, int32(_a_F_GetLocalVictimBuffer_7), int32(0), int32(_a_F_GetLocalVictimBuffer_8), int32(_a_F_GetLocalVictimBuffer_9))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	v168 = v153
	v169 = v128
	goto L33
L33:
	;
	v172 = int32(16)
	v174 = v169 << (uint(int32(1)) % 32)
	if v174 <= v172 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[7])) = v163
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[10]))
	v168 = v163
	v169 = v167
	goto L33
L35:
	;
	v177 = v172
	goto L37
L36:
	;
	v177 = v174
	goto L37
L37:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[4]))
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[6]))
	v182 = v179 - v181
	if v177 < v182 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v184 = v177
	goto L40
L39:
	;
	v184 = v182
	goto L40
L40:
	;
	if base.Ui32(int32(_a_F_GetLocalVictimBuffer_10)) <= base.Ui32(v184) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v187 = int32(_a_F_GetLocalVictimBuffer_10)
	goto L43
L42:
	;
	v187 = v184
	goto L43
L43:
	;
	v192 = F_MemoryContextAllocAligned(m, v168, v187<<(uint(int32(13))%32), int32(_a_F_GetLocalVictimBuffer_11), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[10])) = v187
	*(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[11])) = v192
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	v202 = *(*int32)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[9]))
	v204 = v202
	v205 = int32(-2) - v199
	v206 = v192
	v207 = int32(0)
	goto L4
L45:
	;
	F_FlushLocalBuffer(m, v54, int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v241 = int64(0)
	v244 = base.AtomicRmwCmpxchg64(m, v54, int32(24), v241, v241)
	if v244&int64(33554432) != v241 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L47
L49:
	;
	F_InvalidateLocalBuffer(m, v54, int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	return v279 + int32(1)
L52:
	;
	v252 = int32(1)
	v260 = int32(512)
	v264 = *(*int64)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[13]))
	*(*int64)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[13])) = v264 + int64(1)
	v268 = *(*int64)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[14]))
	*(*int64)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[14])) = v268
	F_pgstat_count_backend_io_op(m, v252, int32(3), int32(0), v252, int64(0))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[15])) = uint8(v252)
	*(*uint8)(unsafe.Add(mBase, _c_F_GetLocalVictimBuffer[16])) = uint8(v252)
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
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v544 int32
	_ = v544
	var v552 int32
	_ = v552
	v18 = m.G0
	v20 = v18 - int32(32)
	m.G0 = v20
	if base.B2i32(l7 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(l7)) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v552 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v537))) = uint8(v552)
	m.G0 = v20 + int32(32)
	return v544 - l0
L2:
	;
	v50 = l1
	v51 = l2
	v58 = l0
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
	v36 = m.ExcPending
	if v36 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v537 = l2
	v544 = l0
	goto L1
L7:
	;
	return int32(0)
L8:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = l7
	F_errmsg(m, int32(_a_F_LocalToUtf_0), v20)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	F_errfinish(m, int32(_a_F_LocalToUtf_1), int32(495), int32(_a_F_LocalToUtf_2))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
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
	v66 = int32(*(*int8)(unsafe.Add(mBase, uint32(v58))))
	if v66 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v537 = v523
	v544 = v531
	goto L1
L14:
	;
	v532 = v50 - v526
	if int32(0) < v532 {
		v50 = v532
		v51 = v523
		v58 = v531
		goto L12
	} else {
		goto L127
	}
L15:
	;
	v523 = v518 + int32(1)
	v526 = v77
	v531 = v128
	goto L14
L16:
	;
	if l8 != 0 {
		v537 = v51
		v544 = v58
		goto L1
	} else {
		goto L125
	}
L17:
	;
	if int32(0) <= v66 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v51))) = uint8(v66)
	v72 = int32(1)
	v523 = v51 + v72
	v526 = v72
	v531 = v58 + v72
	goto L14
L19:
	;
	goto L20
L20:
	;
	v77 = F_pg_encoding_verifymbchar(m, l7, v58, v50)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	if v77 < int32(0) {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	v81 = int32(0)
	switch v77 - int32(1) {
	case 0:
		v113 = v58
		v114 = v81
		v115 = v81
		v116 = v81
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
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v117 | (v115<<(uint(int32(16))%32) | v116<<(uint(int32(24))%32) | v114<<(uint(int32(8))%32))
	v128 = v58 + v77
	if l3 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L24:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	v113 = v58 + int32(1)
	v114 = v112
	v115 = v81
	v116 = v81
	goto L23
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L7
	} else {
		goto L28
	}
L26:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+2)))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	v113 = v58 + int32(3)
	v114 = v92
	v115 = v93
	v116 = v94
	goto L23
L27:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	v113 = v58 + int32(2)
	v114 = v88
	v115 = v89
	v116 = v81
	goto L23
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v77
	F_errmsg_internal(m, int32(_a_F_LocalToUtf_3), v20+int32(16))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_LocalToUtf_1), int32(543), int32(_a_F_LocalToUtf_2))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
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
	v131 = int32(0)
	switch v77 - int32(1) {
	case 0:
		goto L35
	case 1:
		goto L36
	case 2:
		goto L37
	case 3:
		goto L38
	default:
		v362 = v131
		goto L34
	}
L33:
	;
	if v373 != 0 {
		goto L71
	} else {
		goto L72
	}
L34:
	;
	v373 = v362
	goto L33
L35:
	;
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+12)))
	if base.Ui32(v117) < base.Ui32(v333) {
		v362 = v131
		goto L34
	} else {
		goto L66
	}
L36:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+20)))
	if base.Ui32(v114) < base.Ui32(v288) {
		v362 = v131
		goto L34
	} else {
		goto L59
	}
L37:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+28)))
	if base.Ui32(v115) < base.Ui32(v223) {
		v362 = v131
		goto L34
	} else {
		goto L50
	}
L38:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+40)))
	if base.Ui32(v116) < base.Ui32(v138) {
		v362 = v131
		goto L34
	} else {
		goto L39
	}
L39:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+41)))
	if base.Ui32(v140) < base.Ui32(v116) {
		v362 = v131
		goto L34
	} else {
		goto L40
	}
L40:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+42)))
	if base.Ui32(v115) < base.Ui32(v142) {
		v362 = v131
		goto L34
	} else {
		goto L41
	}
L41:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+43)))
	if base.Ui32(v144) < base.Ui32(v115) {
		v362 = v131
		goto L34
	} else {
		goto L42
	}
L42:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+44)))
	if base.Ui32(v114) < base.Ui32(v146) {
		v362 = v131
		goto L34
	} else {
		goto L43
	}
L43:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+45)))
	if base.Ui32(v148) < base.Ui32(v114) {
		v362 = v131
		goto L34
	} else {
		goto L44
	}
L44:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+46)))
	if base.Ui32(v117) < base.Ui32(v150) {
		v362 = v131
		goto L34
	} else {
		goto L45
	}
L45:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+47)))
	if base.Ui32(v152) < base.Ui32(v117) {
		v362 = v131
		goto L34
	} else {
		goto L46
	}
L46:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l3)+36))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v155 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v157 = int32(2)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v155+(v116-v138)<<(uint(v157)%32)+v154<<(uint(v157)%32))))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v155+(v115-v142)<<(uint(v157)%32)+v175<<(uint(v157)%32))))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v155+(v114-v146)<<(uint(v157)%32)+v179<<(uint(v157)%32))))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v155+(v117-v150)<<(uint(v157)%32)+v183<<(uint(v157)%32))))
	v373 = v187
	goto L33
L48:
	;
	goto L49
L49:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v190 = int32(1)
	v210 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188+(v116-v138)<<(uint(v190)%32)+v154&int32(_a_F_LocalToUtf_4)<<(uint(v190)%32)))))
	v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188+(v115-v142)<<(uint(v190)%32)+v210<<(uint(v190)%32)))))
	v218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188+(v114-v146)<<(uint(v190)%32)+v214<<(uint(v190)%32)))))
	v222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188+(v117-v150)<<(uint(v190)%32)+v218<<(uint(v190)%32)))))
	v373 = v222
	goto L33
L50:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+29)))
	if base.Ui32(v225) < base.Ui32(v115) {
		v362 = v131
		goto L34
	} else {
		goto L51
	}
L51:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+30)))
	if base.Ui32(v114) < base.Ui32(v227) {
		v362 = v131
		goto L34
	} else {
		goto L52
	}
L52:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+31)))
	if base.Ui32(v229) < base.Ui32(v114) {
		v362 = v131
		goto L34
	} else {
		goto L53
	}
L53:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+32)))
	if base.Ui32(v117) < base.Ui32(v231) {
		v362 = v131
		goto L34
	} else {
		goto L54
	}
L54:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+33)))
	if base.Ui32(v233) < base.Ui32(v117) {
		v362 = v131
		goto L34
	} else {
		goto L55
	}
L55:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v236 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v238 = int32(2)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v236+(v115-v223)<<(uint(v238)%32)+v235<<(uint(v238)%32))))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v236+(v114-v227)<<(uint(v238)%32)+v252<<(uint(v238)%32))))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v236+(v117-v231)<<(uint(v238)%32)+v256<<(uint(v238)%32))))
	v373 = v260
	goto L33
L57:
	;
	goto L58
L58:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v263 = int32(1)
	v279 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v261+(v115-v223)<<(uint(v263)%32)+v235&int32(_a_F_LocalToUtf_4)<<(uint(v263)%32)))))
	v283 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v261+(v114-v227)<<(uint(v263)%32)+v279<<(uint(v263)%32)))))
	v287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v261+(v117-v231)<<(uint(v263)%32)+v283<<(uint(v263)%32)))))
	v373 = v287
	goto L33
L59:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+21)))
	if base.Ui32(v290) < base.Ui32(v114) {
		v362 = v131
		goto L34
	} else {
		goto L60
	}
L60:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+22)))
	if base.Ui32(v117) < base.Ui32(v292) {
		v362 = v131
		goto L34
	} else {
		goto L61
	}
L61:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+23)))
	if base.Ui32(v294) < base.Ui32(v117) {
		v362 = v131
		goto L34
	} else {
		goto L62
	}
L62:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v297 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v299 = int32(2)
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v297+(v114-v288)<<(uint(v299)%32)+v296<<(uint(v299)%32))))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v297+(v117-v292)<<(uint(v299)%32)+v309<<(uint(v299)%32))))
	v373 = v313
	goto L33
L64:
	;
	goto L65
L65:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v316 = int32(1)
	v328 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v314+(v114-v288)<<(uint(v316)%32)+v296&int32(_a_F_LocalToUtf_4)<<(uint(v316)%32)))))
	v332 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v314+(v117-v292)<<(uint(v316)%32)+v328<<(uint(v316)%32)))))
	v373 = v332
	goto L33
L66:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+13)))
	if base.Ui32(v335) < base.Ui32(v117) {
		v362 = v131
		goto L34
	} else {
		goto L67
	}
L67:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v337 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v339 = int32(2)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v337+(v117-v333)<<(uint(v339)%32)+v342<<(uint(v339)%32))))
	v373 = v346
	goto L33
L69:
	;
	goto L70
L70:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v349 = int32(1)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v356 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v347+(v117-v333)<<(uint(v349)%32)+v352<<(uint(v349)%32)))))
	v362 = v356
	goto L34
L71:
	;
	if base.Ui32(int32(16777216)) <= base.Ui32(v373) {
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
	v377 = int32(base.Ui32(v373) >> (uint(int32(24)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v51))) = uint8(v377)
	v381 = v51 + int32(1)
	goto L76
L75:
	;
	v381 = v51
	goto L76
L76:
	;
	if v373&int32(16711680) != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v385 = int32(base.Ui32(v373) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v381))) = uint8(v385)
	v389 = v381 + int32(1)
	goto L79
L78:
	;
	v389 = v381
	goto L79
L79:
	;
	if v373&int32(_a_F_LocalToUtf_5) != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v393 = int32(base.Ui32(v373) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v389))) = uint8(v393)
	v397 = v389 + int32(1)
	goto L82
L81:
	;
	v397 = v389
	goto L82
L82:
	;
	if v373&int32(255) == int32(0) {
		v523 = v397
		v526 = v77
		v531 = v128
		goto L14
	} else {
		goto L83
	}
L83:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v397))) = uint8(v373)
	v518 = v397
	goto L15
L84:
	;
	v409 = F_bsearch(m, v20+int32(28), l4, l5, int32(12), int32(1850))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L7
	} else {
		goto L85
	}
L85:
	;
	if v409 == int32(0) {
		goto L31
	} else {
		goto L86
	}
L86:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v409)+4))
	if base.Ui32(int32(16777216)) <= base.Ui32(v413) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v417 = int32(base.Ui32(v413) >> (uint(int32(24)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v51))) = uint8(v417)
	v421 = v51 + int32(1)
	goto L89
L88:
	;
	v421 = v51
	goto L89
L89:
	;
	if v413&int32(16711680) != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v425 = int32(base.Ui32(v413) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v421))) = uint8(v425)
	v429 = v421 + int32(1)
	goto L92
L91:
	;
	v429 = v421
	goto L92
L92:
	;
	if v413&int32(_a_F_LocalToUtf_5) != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v433 = int32(base.Ui32(v413) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v429))) = uint8(v433)
	v437 = v429 + int32(1)
	goto L95
L94:
	;
	v437 = v429
	goto L95
L95:
	;
	if v413&int32(255) != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v437))) = uint8(v413)
	v443 = v437 + int32(1)
	goto L98
L97:
	;
	v443 = v437
	goto L98
L98:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v409)+8))
	if base.Ui32(int32(16777216)) <= base.Ui32(v444) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v448 = int32(base.Ui32(v444) >> (uint(int32(24)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v443))) = uint8(v448)
	v452 = v443 + int32(1)
	goto L101
L100:
	;
	v452 = v443
	goto L101
L101:
	;
	if v444&int32(16711680) != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v456 = int32(base.Ui32(v444) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v452))) = uint8(v456)
	v460 = v452 + int32(1)
	goto L104
L103:
	;
	v460 = v452
	goto L104
L104:
	;
	if v444&int32(_a_F_LocalToUtf_5) != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v464 = int32(base.Ui32(v444) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v460))) = uint8(v464)
	v468 = v460 + int32(1)
	goto L107
L106:
	;
	v468 = v460
	goto L107
L107:
	;
	if v444&int32(255) == int32(0) {
		v523 = v468
		v526 = v77
		v531 = v128
		goto L14
	} else {
		goto L108
	}
L108:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v468))) = uint8(v444)
	v518 = v468
	goto L15
L109:
	;
	if l8 != 0 {
		v537 = v51
		v544 = v58
		goto L1
	} else {
		goto L123
	}
L110:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v478 = m.T0[l6].(func(*base.Module, int32) int32)(m, v477)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L7
	} else {
		goto L111
	}
L111:
	;
	if v478 == int32(0) {
		goto L109
	} else {
		goto L112
	}
L112:
	;
	if base.Ui32(int32(16777216)) <= base.Ui32(v478) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v485 = int32(base.Ui32(v478) >> (uint(int32(24)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v51))) = uint8(v485)
	v489 = v51 + int32(1)
	goto L115
L114:
	;
	v489 = v51
	goto L115
L115:
	;
	if v478&int32(16711680) != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v493 = int32(base.Ui32(v478) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v489))) = uint8(v493)
	v497 = v489 + int32(1)
	goto L118
L117:
	;
	v497 = v489
	goto L118
L118:
	;
	if v478&int32(_a_F_LocalToUtf_5) != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v501 = int32(base.Ui32(v478) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v497))) = uint8(v501)
	v505 = v497 + int32(1)
	goto L121
L120:
	;
	v505 = v497
	goto L121
L121:
	;
	if v478&int32(255) == int32(0) {
		v523 = v505
		v526 = v77
		v531 = v128
		goto L14
	} else {
		goto L122
	}
L122:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v505))) = uint8(v478)
	v518 = v505
	goto L15
L123:
	;
	F_report_untranslatable_char(m, l7, int32(6), v58, v50)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
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
	F_report_invalid_encoding(m, l7, v58, v50)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v31 int64
	_ = v31
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_UnpinLocalBuffer[0]))
	v7 = l0 ^ int32(-1)
	v10 = v5 + v7<<(uint(int32(2))%32)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v13 = v11 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v13
	if v13 == int32(0) {
		v17 = int32(_a_F_UnpinLocalBuffer_0)
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_UnpinLocalBuffer[1]))
		*(*int32)(unsafe.Add(mBase, _c_F_UnpinLocalBuffer[1])) = v19 - int32(1)
		v24 = *(*int32)(unsafe.Add(mBase, _c_F_UnpinLocalBuffer[2]))
		v27 = v24 + v7*int32(56)
		v28 = int64(0)
		v31 = base.AtomicRmwCmpxchg64(m, v27, int32(24), v28, v28)
		*(*int64)(unsafe.Add(mBase, uint32(v27)+24)) = v31 - int64(1)
	} else {
	}
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_UnpinLocalBuffer[3]))
	F_ResourceOwnerForget(m, v37, base.I64_extend_i32_s(l0), int32(_a_F_UnpinLocalBuffer_1))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		return
	} else {
		return
	}
}
