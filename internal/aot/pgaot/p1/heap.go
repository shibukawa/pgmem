package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_HeapTupleHeaderIsOnlyLocked(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
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
	var v141 int32
	_ = v141
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	v4 = int32(1)
	v5 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	if v5&int32(2176) != 0 {
		v159 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v159
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v8 == int32(0) {
		v159 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = int32(0)
	if v5&int32(_a_F_HeapTupleHeaderIsOnlyLocked_0) == v11 {
		v159 = v11
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v16 = F_HeapTupleGetUpdateXid(m, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	if base.Ui32(v16) < base.Ui32(int32(3)) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v151 != 0 {
		v159 = v11
		goto L1
	} else {
		goto L47
	}
L8:
	;
	v151 = int32(0)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleHeaderIsOnlyLocked[0]))
	if v31 == v16 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v151 = int32(1)
	goto L7
L12:
	;
	goto L13
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleHeaderIsOnlyLocked[1]))
	if v35 <= int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v151 = v141
	goto L7
L15:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleHeaderIsOnlyLocked[2]))
	if v39 == int32(0) {
		v141 = int32(0)
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleHeaderIsOnlyLocked[3]))
	v111 = int32(0)
	v114 = v35 - int32(1)
	goto L37
L18:
	;
	v44 = v39
	goto L19
L19:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
	if v50 == int32(4) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v141 = int32(0)
	goto L14
L21:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	if v104 != 0 {
		v44 = v104
		goto L19
	} else {
		goto L36
	}
L22:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v53 == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v56 = int32(1)
	if v16 == v53 {
		v141 = v56
		goto L14
	} else {
		goto L24
	}
L24:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v44)+52))
	v60 = v58 - int32(1)
	if v60 < int32(0) {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	v66 = int32(0)
	v69 = v60
	goto L26
L26:
	;
	v74 = int32(2)
	v75 = base.I32_div_s(v69-v66, v74)
	v76 = v75 + v66
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v63+v76<<(uint(v74)%32))))
	if v80 == v16 {
		v141 = v56
		goto L14
	} else {
		goto L28
	}
L27:
	;
	goto L21
L28:
	;
	v89 = base.B2i32(v80-v16 < int32(0)) | base.B2i32(base.Ui32(v80) < base.Ui32(int32(3)))
	if v89 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v90 = v76 + int32(1)
	goto L31
L30:
	;
	v90 = v66
	goto L31
L31:
	;
	if v89 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v93 = v69
	goto L34
L33:
	;
	v93 = v76 - int32(1)
	goto L34
L34:
	;
	if v90 <= v93 {
		v66 = v90
		v69 = v93
		goto L26
	} else {
		goto L35
	}
L35:
	;
	goto L27
L36:
	;
	goto L20
L37:
	;
	v119 = int32(2)
	v120 = base.I32_div_s(v114-v111, v119)
	v121 = v120 + v111
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v109+v121<<(uint(v119)%32))))
	v126 = base.B2i32(v125 == v16)
	if v125 == v16 {
		v141 = v126
		goto L14
	} else {
		goto L39
	}
L38:
	;
	v141 = v126
	goto L14
L39:
	;
	v129 = base.B2i32(base.Ui32(v125) < base.Ui32(v16))
	if base.Ui32(v125) < base.Ui32(v16) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v130 = v121 + int32(1)
	goto L42
L41:
	;
	v130 = v111
	goto L42
L42:
	;
	if base.Ui32(v125) < base.Ui32(v16) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v133 = v114
	goto L45
L44:
	;
	v133 = v121 - int32(1)
	goto L45
L45:
	;
	if v130 <= v133 {
		v111 = v130
		v114 = v133
		goto L37
	} else {
		goto L46
	}
L46:
	;
	goto L38
L47:
	;
	v152 = F_TransactionIdIsInProgress(m, v16)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	if v152 != 0 {
		v159 = v11
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v154 = F_TransactionIdDidCommit(m, v16)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	v159 = v154 ^ int32(1)
	goto L1
}
func F_HeapTupleSatisfiesVacuumHorizon(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
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
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
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
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	v4 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v4
	v11 = v7 + int32(20)
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)))
	if v12&int32(256) == v4 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v427 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11))))
	F_BufferSetHintBits16(m, v11, v427|int32(2048), l1)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L9
	} else {
		goto L156
	}
L2:
	;
	v419 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11))))
	F_BufferSetHintBits16(m, v11, v419|int32(2048), l1)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L9
	} else {
		goto L155
	}
L3:
	;
	v412 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11))))
	F_BufferSetHintBits16(m, v11, v412|int32(512), l1)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L9
	} else {
		goto L154
	}
L4:
	;
	return v410
L5:
	;
	if v12&int32(512) != 0 {
		v410 = v4
		goto L4
	} else {
		goto L8
	}
L6:
	;
	v332 = v12
	goto L7
L7:
	;
	v333 = int32(1)
	if v332&int32(2048) != 0 {
		v410 = v333
		goto L4
	} else {
		goto L113
	}
L8:
	;
	v19 = F_HeapTupleCleanMoved(m, v7, l1)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	if v19 == int32(0) {
		v410 = v4
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if base.Ui32(v25) < base.Ui32(int32(3)) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v157 != 0 {
		goto L52
	} else {
		goto L53
	}
L13:
	;
	v157 = int32(0)
	goto L12
L14:
	;
	goto L15
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVacuumHorizon[0]))
	if v37 == v25 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v157 = int32(1)
	goto L12
L17:
	;
	goto L18
L18:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVacuumHorizon[1]))
	if v41 <= int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v157 = v147
	goto L12
L20:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVacuumHorizon[2]))
	if v45 == int32(0) {
		v147 = int32(0)
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVacuumHorizon[3]))
	v117 = int32(0)
	v120 = v41 - int32(1)
	goto L42
L23:
	;
	v50 = v45
	goto L24
L24:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
	if v56 == int32(4) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v147 = int32(0)
	goto L19
L26:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v50)+80))
	if v110 != 0 {
		v50 = v110
		goto L24
	} else {
		goto L41
	}
L27:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v59 == int32(0) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v62 = int32(1)
	if v25 == v59 {
		v147 = v62
		goto L19
	} else {
		goto L29
	}
L29:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v50)+52))
	v66 = v64 - int32(1)
	if v66 < int32(0) {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v50)+48))
	v72 = int32(0)
	v75 = v66
	goto L31
L31:
	;
	v80 = int32(2)
	v81 = base.I32_div_s(v75-v72, v80)
	v82 = v81 + v72
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v69+v82<<(uint(v80)%32))))
	if v86 == v25 {
		v147 = v62
		goto L19
	} else {
		goto L33
	}
L32:
	;
	goto L26
L33:
	;
	v95 = base.B2i32(v86-v25 < int32(0)) | base.B2i32(base.Ui32(v86) < base.Ui32(int32(3)))
	if v95 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v96 = v82 + int32(1)
	goto L36
L35:
	;
	v96 = v72
	goto L36
L36:
	;
	if v95 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v99 = v75
	goto L39
L38:
	;
	v99 = v82 - int32(1)
	goto L39
L39:
	;
	if v96 <= v99 {
		v72 = v96
		v75 = v99
		goto L31
	} else {
		goto L40
	}
L40:
	;
	goto L32
L41:
	;
	goto L25
L42:
	;
	v125 = int32(2)
	v126 = base.I32_div_s(v120-v117, v125)
	v127 = v126 + v117
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v115+v127<<(uint(v125)%32))))
	v132 = base.B2i32(v131 == v25)
	if v131 == v25 {
		v147 = v132
		goto L19
	} else {
		goto L44
	}
L43:
	;
	v147 = v132
	goto L19
L44:
	;
	v135 = base.B2i32(base.Ui32(v131) < base.Ui32(v25))
	if base.Ui32(v131) < base.Ui32(v25) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v136 = v127 + int32(1)
	goto L47
L46:
	;
	v136 = v117
	goto L47
L47:
	;
	if base.Ui32(v131) < base.Ui32(v25) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v139 = v120
	goto L50
L49:
	;
	v139 = v127 - int32(1)
	goto L50
L50:
	;
	if v136 <= v139 {
		v117 = v136
		v120 = v139
		goto L42
	} else {
		goto L51
	}
L51:
	;
	goto L43
L52:
	;
	v158 = int32(3)
	v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11))))
	if v159&int32(2048)|v159&int32(128)|base.B2i32(v159&int32(_a_F_HeapTupleSatisfiesVacuumHorizon_0) == int32(64)) != 0 {
		v410 = v158
		goto L4
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v318 = F_TransactionIdIsInProgress(m, v317)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L9
	} else {
		goto L106
	}
L55:
	;
	v170 = F_HeapTupleHeaderIsOnlyLocked(m, v7)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L9
	} else {
		goto L56
	}
L56:
	;
	if v170 != 0 {
		v410 = v158
		goto L4
	} else {
		goto L57
	}
L57:
	;
	v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11))))
	if v174&int32(_a_F_HeapTupleSatisfiesVacuumHorizon_1) == int32(_a_F_HeapTupleSatisfiesVacuumHorizon_2) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	if base.Ui32(v182) < base.Ui32(int32(3)) {
		goto L64
	} else {
		goto L65
	}
L59:
	;
	v179 = F_HeapTupleGetUpdateXid(m, v7)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L9
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v182 = v181
	goto L58
L62:
	;
	v182 = v179
	goto L58
L63:
	;
	if v314 != 0 {
		goto L103
	} else {
		goto L104
	}
L64:
	;
	v314 = int32(0)
	goto L63
L65:
	;
	goto L66
L66:
	;
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVacuumHorizon[0]))
	if v194 == v182 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v314 = int32(1)
	goto L63
L68:
	;
	goto L69
L69:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVacuumHorizon[1]))
	if v198 <= int32(0) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v314 = v304
	goto L63
L71:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVacuumHorizon[2]))
	if v202 == int32(0) {
		v304 = int32(0)
		goto L70
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVacuumHorizon[3]))
	v274 = int32(0)
	v277 = v198 - int32(1)
	goto L93
L74:
	;
	v207 = v202
	goto L75
L75:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v207)+20))
	if v213 == int32(4) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v304 = int32(0)
	goto L70
L77:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v207)+80))
	if v267 != 0 {
		v207 = v267
		goto L75
	} else {
		goto L92
	}
L78:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
	if v216 == int32(0) {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v219 = int32(1)
	if v182 == v216 {
		v304 = v219
		goto L70
	} else {
		goto L80
	}
L80:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v207)+52))
	v223 = v221 - int32(1)
	if v223 < int32(0) {
		goto L77
	} else {
		goto L81
	}
L81:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v207)+48))
	v229 = int32(0)
	v232 = v223
	goto L82
L82:
	;
	v237 = int32(2)
	v238 = base.I32_div_s(v232-v229, v237)
	v239 = v238 + v229
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v226+v239<<(uint(v237)%32))))
	if v243 == v182 {
		v304 = v219
		goto L70
	} else {
		goto L84
	}
L83:
	;
	goto L77
L84:
	;
	v252 = base.B2i32(v243-v182 < int32(0)) | base.B2i32(base.Ui32(v243) < base.Ui32(int32(3)))
	if v252 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v253 = v239 + int32(1)
	goto L87
L86:
	;
	v253 = v229
	goto L87
L87:
	;
	if v252 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v256 = v232
	goto L90
L89:
	;
	v256 = v239 - int32(1)
	goto L90
L90:
	;
	if v253 <= v256 {
		v229 = v253
		v232 = v256
		goto L82
	} else {
		goto L91
	}
L91:
	;
	goto L83
L92:
	;
	goto L76
L93:
	;
	v282 = int32(2)
	v283 = base.I32_div_s(v277-v274, v282)
	v284 = v283 + v274
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v272+v284<<(uint(v282)%32))))
	v289 = base.B2i32(v288 == v182)
	if v288 == v182 {
		v304 = v289
		goto L70
	} else {
		goto L95
	}
L94:
	;
	v304 = v289
	goto L70
L95:
	;
	v292 = base.B2i32(base.Ui32(v288) < base.Ui32(v182))
	if base.Ui32(v288) < base.Ui32(v182) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v293 = v284 + int32(1)
	goto L98
L97:
	;
	v293 = v274
	goto L98
L98:
	;
	if base.Ui32(v288) < base.Ui32(v182) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v296 = v277
	goto L101
L100:
	;
	v296 = v284 - int32(1)
	goto L101
L101:
	;
	if v293 <= v296 {
		v274 = v293
		v277 = v296
		goto L93
	} else {
		goto L102
	}
L102:
	;
	goto L94
L103:
	;
	v315 = int32(4)
	goto L105
L104:
	;
	v315 = int32(3)
	goto L105
L105:
	;
	return v315
L106:
	;
	if v318 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	return int32(3)
L108:
	;
	goto L109
L109:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v323 = F_TransactionIdDidCommit(m, v322)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L9
	} else {
		goto L110
	}
L110:
	;
	if v323 == int32(0) {
		goto L3
	} else {
		goto L111
	}
L111:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	F_HeapTupleSetHintBits(m, v7, l1, int32(256), v328)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L9
	} else {
		goto L112
	}
L112:
	;
	v331 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+20)))
	v332 = v331
	goto L7
L113:
	;
	v338 = int32(0)
	if base.B2i32(v332&int32(128) == v338)&base.B2i32(v332&int32(_a_F_HeapTupleSatisfiesVacuumHorizon_0) != int32(64)) == v338 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	if v332&int32(1024) != 0 {
		v410 = v333
		goto L4
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	if v332&int32(_a_F_HeapTupleSatisfiesVacuumHorizon_2) != 0 {
		goto L130
	} else {
		goto L131
	}
L117:
	;
	if v332&int32(_a_F_HeapTupleSatisfiesVacuumHorizon_2) != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	if v332&int32(_a_F_HeapTupleSatisfiesVacuumHorizon_3) != int32(_a_F_HeapTupleSatisfiesVacuumHorizon_4) {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	goto L120
L120:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v370 = F_TransactionIdIsInProgress(m, v369)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L9
	} else {
		goto L127
	}
L121:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v357 = F_MultiXactIdIsRunning(m, v355, int32(1))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L9
	} else {
		goto L124
	}
L122:
	;
	v360 = v332
	goto L123
L123:
	;
	F_BufferSetHintBits16(m, v11, (v360|int32(2048))&int32(_a_F_HeapTupleSatisfiesVacuumHorizon_5), l1)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L9
	} else {
		goto L126
	}
L124:
	;
	if v357 != 0 {
		v410 = v333
		goto L4
	} else {
		goto L125
	}
L125:
	;
	v359 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11))))
	v360 = v359
	goto L123
L126:
	;
	return int32(1)
L127:
	;
	if v370 != 0 {
		v410 = v333
		goto L4
	} else {
		goto L128
	}
L128:
	;
	goto L1
L129:
	;
	v410 = int32(2)
	goto L4
L130:
	;
	v374 = F_HeapTupleGetUpdateXid(m, v7)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L9
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	if v332&int32(1024) == int32(0) {
		goto L144
	} else {
		goto L145
	}
L133:
	;
	v376 = F_TransactionIdIsInProgress(m, v374)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L9
	} else {
		goto L134
	}
L134:
	;
	if v376 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	return int32(4)
L136:
	;
	goto L137
L137:
	;
	v380 = F_TransactionIdDidCommit(m, v374)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L9
	} else {
		goto L138
	}
L138:
	;
	if v380 != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v374
	goto L129
L140:
	;
	goto L141
L141:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v385 = F_MultiXactIdIsRunning(m, v383, int32(0))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L9
	} else {
		goto L142
	}
L142:
	;
	if v385 != 0 {
		v410 = v333
		goto L4
	} else {
		goto L143
	}
L143:
	;
	goto L1
L144:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v392 = F_TransactionIdIsInProgress(m, v391)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L9
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v405
	goto L129
L147:
	;
	if v392 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	return int32(4)
L149:
	;
	goto L150
L150:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v397 = F_TransactionIdDidCommit(m, v396)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L9
	} else {
		goto L151
	}
L151:
	;
	if v397 == int32(0) {
		goto L2
	} else {
		goto L152
	}
L152:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	F_HeapTupleSetHintBits(m, v7, l1, int32(1024), v402)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L9
	} else {
		goto L153
	}
L153:
	;
	goto L146
L154:
	;
	return int32(0)
L155:
	;
	return int32(1)
L156:
	;
	return int32(1)
}
func F_HeapTupleSatisfiesVisibility(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v306 int32
	_ = v306
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v441 int32
	_ = v441
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v616 int32
	_ = v616
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v756 int32
	_ = v756
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v875 int32
	_ = v875
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v935 int32
	_ = v935
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v972 int32
	_ = v972
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1010 int32
	_ = v1010
	var v1014 int32
	_ = v1014
	var v1018 int32
	_ = v1018
	var v1023 int32
	_ = v1023
	var v1029 int32
	_ = v1029
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1048 int32
	_ = v1048
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1059 int32
	_ = v1059
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1072 int32
	_ = v1072
	var v1083 int32
	_ = v1083
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1093 int32
	_ = v1093
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1120 int32
	_ = v1120
	var v1130 int32
	_ = v1130
	var v1133 int32
	_ = v1133
	var v1145 int32
	_ = v1145
	var v1149 int32
	_ = v1149
	var v1153 int32
	_ = v1153
	var v1158 int32
	_ = v1158
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1174 int32
	_ = v1174
	var v1177 int32
	_ = v1177
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1194 int32
	_ = v1194
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1218 int32
	_ = v1218
	var v1223 int32
	_ = v1223
	var v1225 int32
	_ = v1225
	var v1228 int32
	_ = v1228
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1247 int32
	_ = v1247
	var v1255 int32
	_ = v1255
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1284 int32
	_ = v1284
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1293 int32
	_ = v1293
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1302 int32
	_ = v1302
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1332 int32
	_ = v1332
	var v1336 int32
	_ = v1336
	var v1340 int32
	_ = v1340
	var v1345 int32
	_ = v1345
	var v1351 int32
	_ = v1351
	var v1354 int32
	_ = v1354
	var v1357 int32
	_ = v1357
	var v1359 int32
	_ = v1359
	var v1361 int32
	_ = v1361
	var v1364 int32
	_ = v1364
	var v1367 int32
	_ = v1367
	var v1370 int32
	_ = v1370
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1381 int32
	_ = v1381
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1394 int32
	_ = v1394
	var v1405 int32
	_ = v1405
	var v1410 int32
	_ = v1410
	var v1412 int32
	_ = v1412
	var v1415 int32
	_ = v1415
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1434 int32
	_ = v1434
	var v1442 int32
	_ = v1442
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1461 int32
	_ = v1461
	var v1473 int32
	_ = v1473
	var v1477 int32
	_ = v1477
	var v1481 int32
	_ = v1481
	var v1486 int32
	_ = v1486
	var v1492 int32
	_ = v1492
	var v1495 int32
	_ = v1495
	var v1498 int32
	_ = v1498
	var v1500 int32
	_ = v1500
	var v1502 int32
	_ = v1502
	var v1505 int32
	_ = v1505
	var v1508 int32
	_ = v1508
	var v1511 int32
	_ = v1511
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1522 int32
	_ = v1522
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1535 int32
	_ = v1535
	var v1546 int32
	_ = v1546
	var v1551 int32
	_ = v1551
	var v1553 int32
	_ = v1553
	var v1556 int32
	_ = v1556
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1575 int32
	_ = v1575
	var v1583 int32
	_ = v1583
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1615 int32
	_ = v1615
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1626 int32
	_ = v1626
	var v1629 int32
	_ = v1629
	var v1641 int32
	_ = v1641
	var v1643 int32
	_ = v1643
	var v1645 int32
	_ = v1645
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1660 int32
	_ = v1660
	var v1665 int32
	_ = v1665
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1672 int32
	_ = v1672
	var v1677 int32
	_ = v1677
	var v1680 int32
	_ = v1680
	var v1681 int32
	_ = v1681
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1688 int32
	_ = v1688
	var v1690 int32
	_ = v1690
	var v1700 int32
	_ = v1700
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1707 int32
	_ = v1707
	var v1708 int32
	_ = v1708
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1721 int32
	_ = v1721
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1733 int32
	_ = v1733
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1752 int32
	_ = v1752
	var v1753 int32
	_ = v1753
	var v1759 int32
	_ = v1759
	var v1773 int32
	_ = v1773
	var v1774 int32
	_ = v1774
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1782 int32
	_ = v1782
	var v1785 int32
	_ = v1785
	var v1786 int32
	_ = v1786
	var v1789 int32
	_ = v1789
	var v1791 int32
	_ = v1791
	var v1793 int32
	_ = v1793
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1800 int32
	_ = v1800
	var v1803 int32
	_ = v1803
	var v1806 int32
	_ = v1806
	var v1808 int32
	_ = v1808
	var v1818 int32
	_ = v1818
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1825 int32
	_ = v1825
	var v1827 int32
	_ = v1827
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1856 int32
	_ = v1856
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v16 {
	case 0:
		goto L12
	case 1:
		goto L11
	case 2:
		goto L10
	case 3:
		goto L9
	case 4:
		goto L8
	case 5:
		goto L7
	case 6:
		goto L6
	default:
		v1856 = v4
		goto L1
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v1856
L2:
	;
	v1759 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1648)+20)))
	if v1759&int32(2048) != 0 {
		goto L591
	} else {
		goto L592
	}
L3:
	;
	v1731 = int32(3)
	v1733 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.B2i32(base.Ui32(v1657) < base.Ui32(v1731))|base.B2i32(base.Ui32(v1733) < base.Ui32(v1731)) == int32(0) {
		goto L583
	} else {
		goto L584
	}
L4:
	;
	v1726 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v838))))
	F_BufferSetHintBits16(m, v838, v1726|int32(512), l2)
	mBase = m.M
	v1730 = m.ExcPending
	if v1730 != 0 {
		goto L13
	} else {
		goto L581
	}
L5:
	;
	v1721 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24))))
	F_BufferSetHintBits16(m, v24, v1721|int32(512), l2)
	mBase = m.M
	v1725 = m.ExcPending
	if v1725 != 0 {
		goto L13
	} else {
		goto L580
	}
L6:
	;
	v1707 = F_HeapTupleSatisfiesVacuumHorizon(m, l0, l2, v14+int32(12))
	mBase = m.M
	v1708 = m.ExcPending
	if v1708 != 0 {
		goto L13
	} else {
		goto L572
	}
L7:
	;
	v1648 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1649 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1648)+20)))
	v1650 = int32(768)
	v1651 = v1649 & v1650
	if v1651 != v1650 {
		goto L552
	} else {
		goto L553
	}
L8:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v833 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v833
	*(*int64)(unsafe.Add(mBase, uint32(l1)+4)) = int64(0)
	v838 = v832 + int32(20)
	v839 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v832)+20)))
	if v839&int32(256) == v833 {
		goto L285
	} else {
		goto L286
	}
L9:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v814 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v813)+20)))
	if v814&int32(256) != 0 {
		goto L278
	} else {
		goto L279
	}
L10:
	;
	v1856 = int32(1)
	goto L1
L11:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v24 = v22 + int32(20)
	v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+20)))
	if v25&int32(256) == int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v18 = F_HeapTupleSatisfiesMVCC(m, l0, l1, l2, int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	v1856 = v18
	goto L1
L15:
	;
	if v25&int32(512) != 0 {
		v1856 = v4
		goto L1
	} else {
		goto L18
	}
L16:
	;
	v471 = v25
	goto L17
L17:
	;
	if v471&int32(2048) != 0 {
		goto L164
	} else {
		goto L165
	}
L18:
	;
	v32 = F_HeapTupleCleanMoved(m, v22, l2)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	if v32 == int32(0) {
		v1856 = v4
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if base.Ui32(v36) < base.Ui32(int32(3)) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v168 != 0 {
		goto L61
	} else {
		goto L62
	}
L22:
	;
	v168 = int32(0)
	goto L21
L23:
	;
	goto L24
L24:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v48 == v36 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v168 = int32(1)
	goto L21
L26:
	;
	goto L27
L27:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v52 <= int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v168 = v158
	goto L21
L29:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v56 == int32(0) {
		v158 = int32(0)
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v128 = int32(0)
	v131 = v52 - int32(1)
	goto L51
L32:
	;
	v61 = v56
	goto L33
L33:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
	if v67 == int32(4) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v158 = int32(0)
	goto L28
L35:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v61)+80))
	if v121 != 0 {
		v61 = v121
		goto L33
	} else {
		goto L50
	}
L36:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v70 == int32(0) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v73 = int32(1)
	if v36 == v70 {
		v158 = v73
		goto L28
	} else {
		goto L38
	}
L38:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v61)+52))
	v77 = v75 - int32(1)
	if v77 < int32(0) {
		goto L35
	} else {
		goto L39
	}
L39:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v61)+48))
	v83 = int32(0)
	v86 = v77
	goto L40
L40:
	;
	v91 = int32(2)
	v92 = base.I32_div_s(v86-v83, v91)
	v93 = v92 + v83
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v80+v93<<(uint(v91)%32))))
	if v97 == v36 {
		v158 = v73
		goto L28
	} else {
		goto L42
	}
L41:
	;
	goto L35
L42:
	;
	v106 = base.B2i32(v97-v36 < int32(0)) | base.B2i32(base.Ui32(v97) < base.Ui32(int32(3)))
	if v106 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v107 = v93 + int32(1)
	goto L45
L44:
	;
	v107 = v83
	goto L45
L45:
	;
	if v106 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v110 = v86
	goto L48
L47:
	;
	v110 = v93 - int32(1)
	goto L48
L48:
	;
	if v107 <= v110 {
		v83 = v107
		v86 = v110
		goto L40
	} else {
		goto L49
	}
L49:
	;
	goto L41
L50:
	;
	goto L34
L51:
	;
	v136 = int32(2)
	v137 = base.I32_div_s(v131-v128, v136)
	v138 = v137 + v128
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v126+v138<<(uint(v136)%32))))
	v143 = base.B2i32(v142 == v36)
	if v142 == v36 {
		v158 = v143
		goto L28
	} else {
		goto L53
	}
L52:
	;
	v158 = v143
	goto L28
L53:
	;
	v146 = base.B2i32(base.Ui32(v142) < base.Ui32(v36))
	if base.Ui32(v142) < base.Ui32(v36) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v147 = v138 + int32(1)
	goto L56
L55:
	;
	v147 = v128
	goto L56
L56:
	;
	if base.Ui32(v142) < base.Ui32(v36) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v150 = v131
	goto L59
L58:
	;
	v150 = v138 - int32(1)
	goto L59
L59:
	;
	if v147 <= v150 {
		v128 = v147
		v131 = v150
		goto L51
	} else {
		goto L60
	}
L60:
	;
	goto L52
L61:
	;
	v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24))))
	if v169&int32(2048) != 0 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L63
L63:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v459 = F_TransactionIdIsInProgress(m, v458)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L13
	} else {
		goto L159
	}
L64:
	;
	v1856 = int32(1)
	goto L1
L65:
	;
	goto L66
L66:
	;
	if v169&int32(128) != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v1856 = int32(1)
	goto L1
L68:
	;
	goto L69
L69:
	;
	if v169&int32(_a_F_HeapTupleSatisfiesVisibility_0) == int32(64) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v1856 = int32(1)
	goto L1
L71:
	;
	goto L72
L72:
	;
	if v169&int32(_a_F_HeapTupleSatisfiesVisibility_1) != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v183 = F_HeapTupleGetUpdateXid(m, v22)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L13
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if base.Ui32(v319) < base.Ui32(int32(3)) {
		goto L118
	} else {
		goto L119
	}
L76:
	;
	if base.Ui32(v183) < base.Ui32(int32(3)) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v1856 = v316 ^ int32(1)
	goto L1
L78:
	;
	v316 = int32(0)
	goto L77
L79:
	;
	goto L80
L80:
	;
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v196 == v183 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v316 = int32(1)
	goto L77
L82:
	;
	goto L83
L83:
	;
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v200 <= int32(0) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v316 = v306
	goto L77
L85:
	;
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v204 == int32(0) {
		v306 = int32(0)
		goto L84
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v274 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v276 = int32(0)
	v279 = v200 - int32(1)
	goto L107
L88:
	;
	v209 = v204
	goto L89
L89:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v209)+20))
	if v215 == int32(4) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v306 = int32(0)
	goto L84
L91:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v209)+80))
	if v269 != 0 {
		v209 = v269
		goto L89
	} else {
		goto L106
	}
L92:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	if v218 == int32(0) {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v221 = int32(1)
	if v183 == v218 {
		v306 = v221
		goto L84
	} else {
		goto L94
	}
L94:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v209)+52))
	v225 = v223 - int32(1)
	if v225 < int32(0) {
		goto L91
	} else {
		goto L95
	}
L95:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v209)+48))
	v231 = int32(0)
	v234 = v225
	goto L96
L96:
	;
	v239 = int32(2)
	v240 = base.I32_div_s(v234-v231, v239)
	v241 = v240 + v231
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v228+v241<<(uint(v239)%32))))
	if v245 == v183 {
		v306 = v221
		goto L84
	} else {
		goto L98
	}
L97:
	;
	goto L91
L98:
	;
	v254 = base.B2i32(v245-v183 < int32(0)) | base.B2i32(base.Ui32(v245) < base.Ui32(int32(3)))
	if v254 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v255 = v241 + int32(1)
	goto L101
L100:
	;
	v255 = v231
	goto L101
L101:
	;
	if v254 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v258 = v234
	goto L104
L103:
	;
	v258 = v241 - int32(1)
	goto L104
L104:
	;
	if v255 <= v258 {
		v231 = v255
		v234 = v258
		goto L96
	} else {
		goto L105
	}
L105:
	;
	goto L97
L106:
	;
	goto L90
L107:
	;
	v284 = int32(2)
	v285 = base.I32_div_s(v279-v276, v284)
	v286 = v285 + v276
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v274+v286<<(uint(v284)%32))))
	v291 = base.B2i32(v290 == v183)
	if v290 == v183 {
		v306 = v291
		goto L84
	} else {
		goto L109
	}
L108:
	;
	v306 = v291
	goto L84
L109:
	;
	v294 = base.B2i32(base.Ui32(v290) < base.Ui32(v183))
	if base.Ui32(v290) < base.Ui32(v183) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v295 = v286 + int32(1)
	goto L112
L111:
	;
	v295 = v276
	goto L112
L112:
	;
	if base.Ui32(v290) < base.Ui32(v183) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v298 = v279
	goto L115
L114:
	;
	v298 = v286 - int32(1)
	goto L115
L115:
	;
	if v295 <= v298 {
		v276 = v295
		v279 = v298
		goto L107
	} else {
		goto L116
	}
L116:
	;
	goto L108
L117:
	;
	if v451 != 0 {
		v1856 = v4
		goto L1
	} else {
		goto L157
	}
L118:
	;
	v451 = int32(0)
	goto L117
L119:
	;
	goto L120
L120:
	;
	v331 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v331 == v319 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v451 = int32(1)
	goto L117
L122:
	;
	goto L123
L123:
	;
	v335 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v335 <= int32(0) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v451 = v441
	goto L117
L125:
	;
	v339 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v339 == int32(0) {
		v441 = int32(0)
		goto L124
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v409 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v411 = int32(0)
	v414 = v335 - int32(1)
	goto L147
L128:
	;
	v344 = v339
	goto L129
L129:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v344)+20))
	if v350 == int32(4) {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v441 = int32(0)
	goto L124
L131:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v344)+80))
	if v404 != 0 {
		v344 = v404
		goto L129
	} else {
		goto L146
	}
L132:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	if v353 == int32(0) {
		goto L131
	} else {
		goto L133
	}
L133:
	;
	v356 = int32(1)
	if v319 == v353 {
		v441 = v356
		goto L124
	} else {
		goto L134
	}
L134:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v344)+52))
	v360 = v358 - int32(1)
	if v360 < int32(0) {
		goto L131
	} else {
		goto L135
	}
L135:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v344)+48))
	v366 = int32(0)
	v369 = v360
	goto L136
L136:
	;
	v374 = int32(2)
	v375 = base.I32_div_s(v369-v366, v374)
	v376 = v375 + v366
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v363+v376<<(uint(v374)%32))))
	if v380 == v319 {
		v441 = v356
		goto L124
	} else {
		goto L138
	}
L137:
	;
	goto L131
L138:
	;
	v389 = base.B2i32(v380-v319 < int32(0)) | base.B2i32(base.Ui32(v380) < base.Ui32(int32(3)))
	if v389 != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v390 = v376 + int32(1)
	goto L141
L140:
	;
	v390 = v366
	goto L141
L141:
	;
	if v389 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v393 = v369
	goto L144
L143:
	;
	v393 = v376 - int32(1)
	goto L144
L144:
	;
	if v390 <= v393 {
		v366 = v390
		v369 = v393
		goto L136
	} else {
		goto L145
	}
L145:
	;
	goto L137
L146:
	;
	goto L130
L147:
	;
	v419 = int32(2)
	v420 = base.I32_div_s(v414-v411, v419)
	v421 = v420 + v411
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v409+v421<<(uint(v419)%32))))
	v426 = base.B2i32(v425 == v319)
	if v425 == v319 {
		v441 = v426
		goto L124
	} else {
		goto L149
	}
L148:
	;
	v441 = v426
	goto L124
L149:
	;
	v429 = base.B2i32(base.Ui32(v425) < base.Ui32(v319))
	if base.Ui32(v425) < base.Ui32(v319) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v430 = v421 + int32(1)
	goto L152
L151:
	;
	v430 = v411
	goto L152
L152:
	;
	if base.Ui32(v425) < base.Ui32(v319) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v433 = v414
	goto L155
L154:
	;
	v433 = v421 - int32(1)
	goto L155
L155:
	;
	if v430 <= v433 {
		v411 = v430
		v414 = v433
		goto L147
	} else {
		goto L156
	}
L156:
	;
	goto L148
L157:
	;
	v452 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24))))
	F_BufferSetHintBits16(m, v24, v452|int32(2048), l2)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L13
	} else {
		goto L158
	}
L158:
	;
	v1856 = int32(1)
	goto L1
L159:
	;
	if v459 != 0 {
		v1856 = v4
		goto L1
	} else {
		goto L160
	}
L160:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v462 = F_TransactionIdDidCommit(m, v461)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L13
	} else {
		goto L161
	}
L161:
	;
	if v462 == int32(0) {
		goto L5
	} else {
		goto L162
	}
L162:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	F_HeapTupleSetHintBits(m, v22, l2, int32(256), v467)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L13
	} else {
		goto L163
	}
L163:
	;
	v470 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+20)))
	v471 = v470
	goto L17
L164:
	;
	v1856 = int32(1)
	goto L1
L165:
	;
	goto L166
L166:
	;
	v476 = v471 & int32(_a_F_HeapTupleSatisfiesVisibility_2)
	if v476&int32(1024) != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v1856 = int32(base.Ui32(v476&int32(128))>>(uint(int32(7))%32)) | base.B2i32(v476&int32(_a_F_HeapTupleSatisfiesVisibility_0) == int32(64))
	goto L1
L168:
	;
	goto L169
L169:
	;
	if v476&int32(_a_F_HeapTupleSatisfiesVisibility_1) != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	if v476&int32(128) != 0 {
		goto L173
	} else {
		goto L174
	}
L171:
	;
	goto L172
L172:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if base.Ui32(v634) < base.Ui32(int32(3)) {
		goto L222
	} else {
		goto L223
	}
L173:
	;
	v1856 = int32(1)
	goto L1
L174:
	;
	goto L175
L175:
	;
	v493 = F_HeapTupleGetUpdateXid(m, v22)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L13
	} else {
		goto L176
	}
L176:
	;
	if base.Ui32(v493) < base.Ui32(int32(3)) {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	if v626 != 0 {
		v1856 = v4
		goto L1
	} else {
		goto L217
	}
L178:
	;
	v626 = int32(0)
	goto L177
L179:
	;
	goto L180
L180:
	;
	v506 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v506 == v493 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v626 = int32(1)
	goto L177
L182:
	;
	goto L183
L183:
	;
	v510 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v510 <= int32(0) {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	v626 = v616
	goto L177
L185:
	;
	v514 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v514 == int32(0) {
		v616 = int32(0)
		goto L184
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v584 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v586 = int32(0)
	v589 = v510 - int32(1)
	goto L207
L188:
	;
	v519 = v514
	goto L189
L189:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v519)+20))
	if v525 == int32(4) {
		goto L191
	} else {
		goto L192
	}
L190:
	;
	v616 = int32(0)
	goto L184
L191:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v519)+80))
	if v579 != 0 {
		v519 = v579
		goto L189
	} else {
		goto L206
	}
L192:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v519)))
	if v528 == int32(0) {
		goto L191
	} else {
		goto L193
	}
L193:
	;
	v531 = int32(1)
	if v493 == v528 {
		v616 = v531
		goto L184
	} else {
		goto L194
	}
L194:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v519)+52))
	v535 = v533 - int32(1)
	if v535 < int32(0) {
		goto L191
	} else {
		goto L195
	}
L195:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v519)+48))
	v541 = int32(0)
	v544 = v535
	goto L196
L196:
	;
	v549 = int32(2)
	v550 = base.I32_div_s(v544-v541, v549)
	v551 = v550 + v541
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v538+v551<<(uint(v549)%32))))
	if v555 == v493 {
		v616 = v531
		goto L184
	} else {
		goto L198
	}
L197:
	;
	goto L191
L198:
	;
	v564 = base.B2i32(v555-v493 < int32(0)) | base.B2i32(base.Ui32(v555) < base.Ui32(int32(3)))
	if v564 != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v565 = v551 + int32(1)
	goto L201
L200:
	;
	v565 = v541
	goto L201
L201:
	;
	if v564 != 0 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v568 = v544
	goto L204
L203:
	;
	v568 = v551 - int32(1)
	goto L204
L204:
	;
	if v565 <= v568 {
		v541 = v565
		v544 = v568
		goto L196
	} else {
		goto L205
	}
L205:
	;
	goto L197
L206:
	;
	goto L190
L207:
	;
	v594 = int32(2)
	v595 = base.I32_div_s(v589-v586, v594)
	v596 = v595 + v586
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v584+v596<<(uint(v594)%32))))
	v601 = base.B2i32(v600 == v493)
	if v600 == v493 {
		v616 = v601
		goto L184
	} else {
		goto L209
	}
L208:
	;
	v616 = v601
	goto L184
L209:
	;
	v604 = base.B2i32(base.Ui32(v600) < base.Ui32(v493))
	if base.Ui32(v600) < base.Ui32(v493) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v605 = v596 + int32(1)
	goto L212
L211:
	;
	v605 = v586
	goto L212
L212:
	;
	if base.Ui32(v600) < base.Ui32(v493) {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v608 = v589
	goto L215
L214:
	;
	v608 = v596 - int32(1)
	goto L215
L215:
	;
	if v605 <= v608 {
		v586 = v605
		v589 = v608
		goto L207
	} else {
		goto L216
	}
L216:
	;
	goto L208
L217:
	;
	v628 = F_TransactionIdIsInProgress(m, v493)
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L13
	} else {
		goto L218
	}
L218:
	;
	if v628 != 0 {
		v1856 = int32(1)
		goto L1
	} else {
		goto L219
	}
L219:
	;
	v630 = F_TransactionIdDidCommit(m, v493)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L13
	} else {
		goto L220
	}
L220:
	;
	v1856 = v630 ^ int32(1)
	goto L1
L221:
	;
	if v766 != 0 {
		goto L261
	} else {
		goto L262
	}
L222:
	;
	v766 = int32(0)
	goto L221
L223:
	;
	goto L224
L224:
	;
	v646 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v646 == v634 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v766 = int32(1)
	goto L221
L226:
	;
	goto L227
L227:
	;
	v650 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v650 <= int32(0) {
		goto L229
	} else {
		goto L230
	}
L228:
	;
	v766 = v756
	goto L221
L229:
	;
	v654 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v654 == int32(0) {
		v756 = int32(0)
		goto L228
	} else {
		goto L232
	}
L230:
	;
	goto L231
L231:
	;
	v724 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v726 = int32(0)
	v729 = v650 - int32(1)
	goto L251
L232:
	;
	v659 = v654
	goto L233
L233:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v659)+20))
	if v665 == int32(4) {
		goto L235
	} else {
		goto L236
	}
L234:
	;
	v756 = int32(0)
	goto L228
L235:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v659)+80))
	if v719 != 0 {
		v659 = v719
		goto L233
	} else {
		goto L250
	}
L236:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v659)))
	if v668 == int32(0) {
		goto L235
	} else {
		goto L237
	}
L237:
	;
	v671 = int32(1)
	if v634 == v668 {
		v756 = v671
		goto L228
	} else {
		goto L238
	}
L238:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v659)+52))
	v675 = v673 - int32(1)
	if v675 < int32(0) {
		goto L235
	} else {
		goto L239
	}
L239:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v659)+48))
	v681 = int32(0)
	v684 = v675
	goto L240
L240:
	;
	v689 = int32(2)
	v690 = base.I32_div_s(v684-v681, v689)
	v691 = v690 + v681
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v678+v691<<(uint(v689)%32))))
	if v695 == v634 {
		v756 = v671
		goto L228
	} else {
		goto L242
	}
L241:
	;
	goto L235
L242:
	;
	v704 = base.B2i32(v695-v634 < int32(0)) | base.B2i32(base.Ui32(v695) < base.Ui32(int32(3)))
	if v704 != 0 {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v705 = v691 + int32(1)
	goto L245
L244:
	;
	v705 = v681
	goto L245
L245:
	;
	if v704 != 0 {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v708 = v684
	goto L248
L247:
	;
	v708 = v691 - int32(1)
	goto L248
L248:
	;
	if v705 <= v708 {
		v681 = v705
		v684 = v708
		goto L240
	} else {
		goto L249
	}
L249:
	;
	goto L241
L250:
	;
	goto L234
L251:
	;
	v734 = int32(2)
	v735 = base.I32_div_s(v729-v726, v734)
	v736 = v735 + v726
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v724+v736<<(uint(v734)%32))))
	v741 = base.B2i32(v740 == v634)
	if v740 == v634 {
		v756 = v741
		goto L228
	} else {
		goto L253
	}
L252:
	;
	v756 = v741
	goto L228
L253:
	;
	v744 = base.B2i32(base.Ui32(v740) < base.Ui32(v634))
	if base.Ui32(v740) < base.Ui32(v634) {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v745 = v736 + int32(1)
	goto L256
L255:
	;
	v745 = v726
	goto L256
L256:
	;
	if base.Ui32(v740) < base.Ui32(v634) {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v748 = v729
	goto L259
L258:
	;
	v748 = v736 - int32(1)
	goto L259
L259:
	;
	if v745 <= v748 {
		v726 = v745
		v729 = v748
		goto L251
	} else {
		goto L260
	}
L260:
	;
	goto L252
L261:
	;
	v767 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24))))
	v1856 = int32(base.Ui32(v767&int32(128))>>(uint(int32(7))%32)) | base.B2i32(v767&int32(_a_F_HeapTupleSatisfiesVisibility_0) == int32(64))
	goto L1
L262:
	;
	goto L263
L263:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v778 = F_TransactionIdIsInProgress(m, v777)
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L13
	} else {
		goto L264
	}
L264:
	;
	if v778 != 0 {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v1856 = int32(1)
	goto L1
L266:
	;
	goto L267
L267:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v782 = F_TransactionIdDidCommit(m, v781)
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L13
	} else {
		goto L268
	}
L268:
	;
	v784 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+20)))
	if v782 == int32(0) {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	F_BufferSetHintBits16(m, v24, v784|int32(2048), l2)
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L13
	} else {
		goto L272
	}
L270:
	;
	goto L271
L271:
	;
	v794 = int32(0)
	if base.B2i32(v784&int32(128) == v794)&base.B2i32(v784&int32(_a_F_HeapTupleSatisfiesVisibility_0) != int32(64)) == v794 {
		goto L273
	} else {
		goto L274
	}
L272:
	;
	v1856 = int32(1)
	goto L1
L273:
	;
	F_BufferSetHintBits16(m, v24, v784|int32(2048), l2)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L13
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	F_HeapTupleSetHintBits(m, v22, l2, int32(1024), v809)
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L13
	} else {
		goto L277
	}
L276:
	;
	v1856 = int32(1)
	goto L1
L277:
	;
	v1856 = v4
	goto L1
L278:
	;
	v1856 = int32(1)
	goto L1
L279:
	;
	if v814&int32(512) != 0 {
		v1856 = v4
		goto L1
	} else {
		goto L280
	}
L280:
	;
	v819 = F_HeapTupleCleanMoved(m, v813, l2)
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L13
	} else {
		goto L281
	}
L281:
	;
	if v819 == int32(0) {
		v1856 = v4
		goto L1
	} else {
		goto L282
	}
L282:
	;
	v823 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v813)+20)))
	v824 = int32(768)
	if v823&v824 == v824 {
		goto L278
	} else {
		goto L283
	}
L283:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v813)))
	if v828 == int32(0) {
		v1856 = v4
		goto L1
	} else {
		goto L284
	}
L284:
	;
	goto L278
L285:
	;
	if v839&int32(512) != 0 {
		v1856 = v4
		goto L1
	} else {
		goto L288
	}
L286:
	;
	v1297 = v839
	goto L287
L287:
	;
	v1298 = int32(1)
	if v1297&int32(2048) != 0 {
		v1856 = v1298
		goto L1
	} else {
		goto L439
	}
L288:
	;
	v846 = F_HeapTupleCleanMoved(m, v832, l2)
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L13
	} else {
		goto L289
	}
L289:
	;
	if v846 == int32(0) {
		v1856 = v4
		goto L1
	} else {
		goto L290
	}
L290:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v832)))
	if base.Ui32(v850) < base.Ui32(int32(3)) {
		goto L292
	} else {
		goto L293
	}
L291:
	;
	if v982 != 0 {
		goto L331
	} else {
		goto L332
	}
L292:
	;
	v982 = int32(0)
	goto L291
L293:
	;
	goto L294
L294:
	;
	v862 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v862 == v850 {
		goto L295
	} else {
		goto L296
	}
L295:
	;
	v982 = int32(1)
	goto L291
L296:
	;
	goto L297
L297:
	;
	v866 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v866 <= int32(0) {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	v982 = v972
	goto L291
L299:
	;
	v870 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v870 == int32(0) {
		v972 = int32(0)
		goto L298
	} else {
		goto L302
	}
L300:
	;
	goto L301
L301:
	;
	v940 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v942 = int32(0)
	v945 = v866 - int32(1)
	goto L321
L302:
	;
	v875 = v870
	goto L303
L303:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v875)+20))
	if v881 == int32(4) {
		goto L305
	} else {
		goto L306
	}
L304:
	;
	v972 = int32(0)
	goto L298
L305:
	;
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v875)+80))
	if v935 != 0 {
		v875 = v935
		goto L303
	} else {
		goto L320
	}
L306:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v875)))
	if v884 == int32(0) {
		goto L305
	} else {
		goto L307
	}
L307:
	;
	v887 = int32(1)
	if v850 == v884 {
		v972 = v887
		goto L298
	} else {
		goto L308
	}
L308:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v875)+52))
	v891 = v889 - int32(1)
	if v891 < int32(0) {
		goto L305
	} else {
		goto L309
	}
L309:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v875)+48))
	v897 = int32(0)
	v900 = v891
	goto L310
L310:
	;
	v905 = int32(2)
	v906 = base.I32_div_s(v900-v897, v905)
	v907 = v906 + v897
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v894+v907<<(uint(v905)%32))))
	if v911 == v850 {
		v972 = v887
		goto L298
	} else {
		goto L312
	}
L311:
	;
	goto L305
L312:
	;
	v920 = base.B2i32(v911-v850 < int32(0)) | base.B2i32(base.Ui32(v911) < base.Ui32(int32(3)))
	if v920 != 0 {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	v921 = v907 + int32(1)
	goto L315
L314:
	;
	v921 = v897
	goto L315
L315:
	;
	if v920 != 0 {
		goto L316
	} else {
		goto L317
	}
L316:
	;
	v924 = v900
	goto L318
L317:
	;
	v924 = v907 - int32(1)
	goto L318
L318:
	;
	if v921 <= v924 {
		v897 = v921
		v900 = v924
		goto L310
	} else {
		goto L319
	}
L319:
	;
	goto L311
L320:
	;
	goto L304
L321:
	;
	v950 = int32(2)
	v951 = base.I32_div_s(v945-v942, v950)
	v952 = v951 + v942
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v940+v952<<(uint(v950)%32))))
	v957 = base.B2i32(v956 == v850)
	if v956 == v850 {
		v972 = v957
		goto L298
	} else {
		goto L323
	}
L322:
	;
	v972 = v957
	goto L298
L323:
	;
	v960 = base.B2i32(base.Ui32(v956) < base.Ui32(v850))
	if base.Ui32(v956) < base.Ui32(v850) {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	v961 = v952 + int32(1)
	goto L326
L325:
	;
	v961 = v942
	goto L326
L326:
	;
	if base.Ui32(v956) < base.Ui32(v850) {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	v964 = v945
	goto L329
L328:
	;
	v964 = v952 - int32(1)
	goto L329
L329:
	;
	if v961 <= v964 {
		v942 = v961
		v945 = v964
		goto L321
	} else {
		goto L330
	}
L330:
	;
	goto L322
L331:
	;
	v983 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v838))))
	if v983&int32(2048) != 0 {
		goto L334
	} else {
		goto L335
	}
L332:
	;
	goto L333
L333:
	;
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v832)))
	v1273 = F_TransactionIdIsInProgress(m, v1272)
	mBase = m.M
	v1274 = m.ExcPending
	if v1274 != 0 {
		goto L13
	} else {
		goto L429
	}
L334:
	;
	v1856 = int32(1)
	goto L1
L335:
	;
	goto L336
L336:
	;
	if v983&int32(128) != 0 {
		goto L337
	} else {
		goto L338
	}
L337:
	;
	v1856 = int32(1)
	goto L1
L338:
	;
	goto L339
L339:
	;
	if v983&int32(_a_F_HeapTupleSatisfiesVisibility_0) == int32(64) {
		goto L340
	} else {
		goto L341
	}
L340:
	;
	v1856 = int32(1)
	goto L1
L341:
	;
	goto L342
L342:
	;
	if v983&int32(_a_F_HeapTupleSatisfiesVisibility_1) != 0 {
		goto L343
	} else {
		goto L344
	}
L343:
	;
	v997 = F_HeapTupleGetUpdateXid(m, v832)
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L13
	} else {
		goto L346
	}
L344:
	;
	goto L345
L345:
	;
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v832)+4))
	if base.Ui32(v1133) < base.Ui32(int32(3)) {
		goto L388
	} else {
		goto L389
	}
L346:
	;
	if base.Ui32(v997) < base.Ui32(int32(3)) {
		goto L348
	} else {
		goto L349
	}
L347:
	;
	v1856 = v1130 ^ int32(1)
	goto L1
L348:
	;
	v1130 = int32(0)
	goto L347
L349:
	;
	goto L350
L350:
	;
	v1010 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v1010 == v997 {
		goto L351
	} else {
		goto L352
	}
L351:
	;
	v1130 = int32(1)
	goto L347
L352:
	;
	goto L353
L353:
	;
	v1014 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v1014 <= int32(0) {
		goto L355
	} else {
		goto L356
	}
L354:
	;
	v1130 = v1120
	goto L347
L355:
	;
	v1018 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v1018 == int32(0) {
		v1120 = int32(0)
		goto L354
	} else {
		goto L358
	}
L356:
	;
	goto L357
L357:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v1090 = int32(0)
	v1093 = v1014 - int32(1)
	goto L377
L358:
	;
	v1023 = v1018
	goto L359
L359:
	;
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v1023)+20))
	if v1029 == int32(4) {
		goto L361
	} else {
		goto L362
	}
L360:
	;
	v1120 = int32(0)
	goto L354
L361:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v1023)+80))
	if v1083 != 0 {
		v1023 = v1083
		goto L359
	} else {
		goto L376
	}
L362:
	;
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(v1023)))
	if v1032 == int32(0) {
		goto L361
	} else {
		goto L363
	}
L363:
	;
	v1035 = int32(1)
	if v997 == v1032 {
		v1120 = v1035
		goto L354
	} else {
		goto L364
	}
L364:
	;
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v1023)+52))
	v1039 = v1037 - int32(1)
	if v1039 < int32(0) {
		goto L361
	} else {
		goto L365
	}
L365:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1023)+48))
	v1045 = int32(0)
	v1048 = v1039
	goto L366
L366:
	;
	v1053 = int32(2)
	v1054 = base.I32_div_s(v1048-v1045, v1053)
	v1055 = v1054 + v1045
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v1042+v1055<<(uint(v1053)%32))))
	if v1059 == v997 {
		v1120 = v1035
		goto L354
	} else {
		goto L368
	}
L367:
	;
	goto L361
L368:
	;
	v1068 = base.B2i32(v1059-v997 < int32(0)) | base.B2i32(base.Ui32(v1059) < base.Ui32(int32(3)))
	if v1068 != 0 {
		goto L369
	} else {
		goto L370
	}
L369:
	;
	v1069 = v1055 + int32(1)
	goto L371
L370:
	;
	v1069 = v1045
	goto L371
L371:
	;
	if v1068 != 0 {
		goto L372
	} else {
		goto L373
	}
L372:
	;
	v1072 = v1048
	goto L374
L373:
	;
	v1072 = v1055 - int32(1)
	goto L374
L374:
	;
	if v1069 <= v1072 {
		v1045 = v1069
		v1048 = v1072
		goto L366
	} else {
		goto L375
	}
L375:
	;
	goto L367
L376:
	;
	goto L360
L377:
	;
	v1098 = int32(2)
	v1099 = base.I32_div_s(v1093-v1090, v1098)
	v1100 = v1099 + v1090
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v1088+v1100<<(uint(v1098)%32))))
	v1105 = base.B2i32(v1104 == v997)
	if v1104 == v997 {
		v1120 = v1105
		goto L354
	} else {
		goto L379
	}
L378:
	;
	v1120 = v1105
	goto L354
L379:
	;
	v1108 = base.B2i32(base.Ui32(v1104) < base.Ui32(v997))
	if base.Ui32(v1104) < base.Ui32(v997) {
		goto L380
	} else {
		goto L381
	}
L380:
	;
	v1109 = v1100 + int32(1)
	goto L382
L381:
	;
	v1109 = v1090
	goto L382
L382:
	;
	if base.Ui32(v1104) < base.Ui32(v997) {
		goto L383
	} else {
		goto L384
	}
L383:
	;
	v1112 = v1093
	goto L385
L384:
	;
	v1112 = v1100 - int32(1)
	goto L385
L385:
	;
	if v1109 <= v1112 {
		v1090 = v1109
		v1093 = v1112
		goto L377
	} else {
		goto L386
	}
L386:
	;
	goto L378
L387:
	;
	if v1265 != 0 {
		v1856 = v4
		goto L1
	} else {
		goto L427
	}
L388:
	;
	v1265 = int32(0)
	goto L387
L389:
	;
	goto L390
L390:
	;
	v1145 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v1145 == v1133 {
		goto L391
	} else {
		goto L392
	}
L391:
	;
	v1265 = int32(1)
	goto L387
L392:
	;
	goto L393
L393:
	;
	v1149 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v1149 <= int32(0) {
		goto L395
	} else {
		goto L396
	}
L394:
	;
	v1265 = v1255
	goto L387
L395:
	;
	v1153 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v1153 == int32(0) {
		v1255 = int32(0)
		goto L394
	} else {
		goto L398
	}
L396:
	;
	goto L397
L397:
	;
	v1223 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v1225 = int32(0)
	v1228 = v1149 - int32(1)
	goto L417
L398:
	;
	v1158 = v1153
	goto L399
L399:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+20))
	if v1164 == int32(4) {
		goto L401
	} else {
		goto L402
	}
L400:
	;
	v1255 = int32(0)
	goto L394
L401:
	;
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+80))
	if v1218 != 0 {
		v1158 = v1218
		goto L399
	} else {
		goto L416
	}
L402:
	;
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v1158)))
	if v1167 == int32(0) {
		goto L401
	} else {
		goto L403
	}
L403:
	;
	v1170 = int32(1)
	if v1133 == v1167 {
		v1255 = v1170
		goto L394
	} else {
		goto L404
	}
L404:
	;
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+52))
	v1174 = v1172 - int32(1)
	if v1174 < int32(0) {
		goto L401
	} else {
		goto L405
	}
L405:
	;
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+48))
	v1180 = int32(0)
	v1183 = v1174
	goto L406
L406:
	;
	v1188 = int32(2)
	v1189 = base.I32_div_s(v1183-v1180, v1188)
	v1190 = v1189 + v1180
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1177+v1190<<(uint(v1188)%32))))
	if v1194 == v1133 {
		v1255 = v1170
		goto L394
	} else {
		goto L408
	}
L407:
	;
	goto L401
L408:
	;
	v1203 = base.B2i32(v1194-v1133 < int32(0)) | base.B2i32(base.Ui32(v1194) < base.Ui32(int32(3)))
	if v1203 != 0 {
		goto L409
	} else {
		goto L410
	}
L409:
	;
	v1204 = v1190 + int32(1)
	goto L411
L410:
	;
	v1204 = v1180
	goto L411
L411:
	;
	if v1203 != 0 {
		goto L412
	} else {
		goto L413
	}
L412:
	;
	v1207 = v1183
	goto L414
L413:
	;
	v1207 = v1190 - int32(1)
	goto L414
L414:
	;
	if v1204 <= v1207 {
		v1180 = v1204
		v1183 = v1207
		goto L406
	} else {
		goto L415
	}
L415:
	;
	goto L407
L416:
	;
	goto L400
L417:
	;
	v1233 = int32(2)
	v1234 = base.I32_div_s(v1228-v1225, v1233)
	v1235 = v1234 + v1225
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v1223+v1235<<(uint(v1233)%32))))
	v1240 = base.B2i32(v1239 == v1133)
	if v1239 == v1133 {
		v1255 = v1240
		goto L394
	} else {
		goto L419
	}
L418:
	;
	v1255 = v1240
	goto L394
L419:
	;
	v1243 = base.B2i32(base.Ui32(v1239) < base.Ui32(v1133))
	if base.Ui32(v1239) < base.Ui32(v1133) {
		goto L420
	} else {
		goto L421
	}
L420:
	;
	v1244 = v1235 + int32(1)
	goto L422
L421:
	;
	v1244 = v1225
	goto L422
L422:
	;
	if base.Ui32(v1239) < base.Ui32(v1133) {
		goto L423
	} else {
		goto L424
	}
L423:
	;
	v1247 = v1228
	goto L425
L424:
	;
	v1247 = v1235 - int32(1)
	goto L425
L425:
	;
	if v1244 <= v1247 {
		v1225 = v1244
		v1228 = v1247
		goto L417
	} else {
		goto L426
	}
L426:
	;
	goto L418
L427:
	;
	v1266 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v838))))
	F_BufferSetHintBits16(m, v838, v1266|int32(2048), l2)
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L13
	} else {
		goto L428
	}
L428:
	;
	v1856 = int32(1)
	goto L1
L429:
	;
	if v1273 != 0 {
		goto L430
	} else {
		goto L431
	}
L430:
	;
	v1275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v832)+16)))
	if v1275 == int32(_a_F_HeapTupleSatisfiesVisibility_3) {
		goto L433
	} else {
		goto L434
	}
L431:
	;
	goto L432
L432:
	;
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v832)))
	v1288 = F_TransactionIdDidCommit(m, v1287)
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L13
	} else {
		goto L436
	}
L433:
	;
	v1278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v832)+14)))
	v1279 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v832)+12)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v1278 | v1279<<(uint(int32(16))%32)
	goto L435
L434:
	;
	goto L435
L435:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v832)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v1284
	v1856 = int32(1)
	goto L1
L436:
	;
	if v1288 == int32(0) {
		goto L4
	} else {
		goto L437
	}
L437:
	;
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(v832)))
	F_HeapTupleSetHintBits(m, v832, l2, int32(256), v1293)
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L13
	} else {
		goto L438
	}
L438:
	;
	v1296 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v832)+20)))
	v1297 = v1296
	goto L287
L439:
	;
	v1302 = v1297 & int32(_a_F_HeapTupleSatisfiesVisibility_2)
	if v1302&int32(1024) != 0 {
		goto L440
	} else {
		goto L441
	}
L440:
	;
	v1856 = int32(base.Ui32(v1302&int32(128))>>(uint(int32(7))%32)) | base.B2i32(v1302&int32(_a_F_HeapTupleSatisfiesVisibility_0) == int32(64))
	goto L1
L441:
	;
	goto L442
L442:
	;
	if v1302&int32(_a_F_HeapTupleSatisfiesVisibility_1) != 0 {
		goto L443
	} else {
		goto L444
	}
L443:
	;
	if v1302&int32(128) != 0 {
		v1856 = v1298
		goto L1
	} else {
		goto L446
	}
L444:
	;
	goto L445
L445:
	;
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(v832)+4))
	if base.Ui32(v1461) < base.Ui32(int32(3)) {
		goto L495
	} else {
		goto L496
	}
L446:
	;
	v1319 = F_HeapTupleGetUpdateXid(m, v832)
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		goto L13
	} else {
		goto L447
	}
L447:
	;
	if base.Ui32(v1319) < base.Ui32(int32(3)) {
		goto L449
	} else {
		goto L450
	}
L448:
	;
	if v1452 != 0 {
		v1856 = int32(0)
		goto L1
	} else {
		goto L488
	}
L449:
	;
	v1452 = int32(0)
	goto L448
L450:
	;
	goto L451
L451:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v1332 == v1319 {
		goto L452
	} else {
		goto L453
	}
L452:
	;
	v1452 = int32(1)
	goto L448
L453:
	;
	goto L454
L454:
	;
	v1336 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v1336 <= int32(0) {
		goto L456
	} else {
		goto L457
	}
L455:
	;
	v1452 = v1442
	goto L448
L456:
	;
	v1340 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v1340 == int32(0) {
		v1442 = int32(0)
		goto L455
	} else {
		goto L459
	}
L457:
	;
	goto L458
L458:
	;
	v1410 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v1412 = int32(0)
	v1415 = v1336 - int32(1)
	goto L478
L459:
	;
	v1345 = v1340
	goto L460
L460:
	;
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(v1345)+20))
	if v1351 == int32(4) {
		goto L462
	} else {
		goto L463
	}
L461:
	;
	v1442 = int32(0)
	goto L455
L462:
	;
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(v1345)+80))
	if v1405 != 0 {
		v1345 = v1405
		goto L460
	} else {
		goto L477
	}
L463:
	;
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v1345)))
	if v1354 == int32(0) {
		goto L462
	} else {
		goto L464
	}
L464:
	;
	v1357 = int32(1)
	if v1319 == v1354 {
		v1442 = v1357
		goto L455
	} else {
		goto L465
	}
L465:
	;
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(v1345)+52))
	v1361 = v1359 - int32(1)
	if v1361 < int32(0) {
		goto L462
	} else {
		goto L466
	}
L466:
	;
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v1345)+48))
	v1367 = int32(0)
	v1370 = v1361
	goto L467
L467:
	;
	v1375 = int32(2)
	v1376 = base.I32_div_s(v1370-v1367, v1375)
	v1377 = v1376 + v1367
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v1364+v1377<<(uint(v1375)%32))))
	if v1381 == v1319 {
		v1442 = v1357
		goto L455
	} else {
		goto L469
	}
L468:
	;
	goto L462
L469:
	;
	v1390 = base.B2i32(v1381-v1319 < int32(0)) | base.B2i32(base.Ui32(v1381) < base.Ui32(int32(3)))
	if v1390 != 0 {
		goto L470
	} else {
		goto L471
	}
L470:
	;
	v1391 = v1377 + int32(1)
	goto L472
L471:
	;
	v1391 = v1367
	goto L472
L472:
	;
	if v1390 != 0 {
		goto L473
	} else {
		goto L474
	}
L473:
	;
	v1394 = v1370
	goto L475
L474:
	;
	v1394 = v1377 - int32(1)
	goto L475
L475:
	;
	if v1391 <= v1394 {
		v1367 = v1391
		v1370 = v1394
		goto L467
	} else {
		goto L476
	}
L476:
	;
	goto L468
L477:
	;
	goto L461
L478:
	;
	v1420 = int32(2)
	v1421 = base.I32_div_s(v1415-v1412, v1420)
	v1422 = v1421 + v1412
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(v1410+v1422<<(uint(v1420)%32))))
	v1427 = base.B2i32(v1426 == v1319)
	if v1426 == v1319 {
		v1442 = v1427
		goto L455
	} else {
		goto L480
	}
L479:
	;
	v1442 = v1427
	goto L455
L480:
	;
	v1430 = base.B2i32(base.Ui32(v1426) < base.Ui32(v1319))
	if base.Ui32(v1426) < base.Ui32(v1319) {
		goto L481
	} else {
		goto L482
	}
L481:
	;
	v1431 = v1422 + int32(1)
	goto L483
L482:
	;
	v1431 = v1412
	goto L483
L483:
	;
	if base.Ui32(v1426) < base.Ui32(v1319) {
		goto L484
	} else {
		goto L485
	}
L484:
	;
	v1434 = v1415
	goto L486
L485:
	;
	v1434 = v1422 - int32(1)
	goto L486
L486:
	;
	if v1431 <= v1434 {
		v1412 = v1431
		v1415 = v1434
		goto L478
	} else {
		goto L487
	}
L487:
	;
	goto L479
L488:
	;
	v1453 = F_TransactionIdIsInProgress(m, v1319)
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L13
	} else {
		goto L489
	}
L489:
	;
	if v1453 != 0 {
		goto L490
	} else {
		goto L491
	}
L490:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v1319
	v1856 = int32(1)
	goto L1
L491:
	;
	goto L492
L492:
	;
	v1457 = F_TransactionIdDidCommit(m, v1319)
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
		goto L13
	} else {
		goto L493
	}
L493:
	;
	v1856 = v1457 ^ int32(1)
	goto L1
L494:
	;
	if v1593 != 0 {
		goto L534
	} else {
		goto L535
	}
L495:
	;
	v1593 = int32(0)
	goto L494
L496:
	;
	goto L497
L497:
	;
	v1473 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[0]))
	if v1473 == v1461 {
		goto L498
	} else {
		goto L499
	}
L498:
	;
	v1593 = int32(1)
	goto L494
L499:
	;
	goto L500
L500:
	;
	v1477 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[1]))
	if v1477 <= int32(0) {
		goto L502
	} else {
		goto L503
	}
L501:
	;
	v1593 = v1583
	goto L494
L502:
	;
	v1481 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[2]))
	if v1481 == int32(0) {
		v1583 = int32(0)
		goto L501
	} else {
		goto L505
	}
L503:
	;
	goto L504
L504:
	;
	v1551 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[3]))
	v1553 = int32(0)
	v1556 = v1477 - int32(1)
	goto L524
L505:
	;
	v1486 = v1481
	goto L506
L506:
	;
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v1486)+20))
	if v1492 == int32(4) {
		goto L508
	} else {
		goto L509
	}
L507:
	;
	v1583 = int32(0)
	goto L501
L508:
	;
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v1486)+80))
	if v1546 != 0 {
		v1486 = v1546
		goto L506
	} else {
		goto L523
	}
L509:
	;
	v1495 = *(*int32)(unsafe.Add(mBase, uint32(v1486)))
	if v1495 == int32(0) {
		goto L508
	} else {
		goto L510
	}
L510:
	;
	v1498 = int32(1)
	if v1461 == v1495 {
		v1583 = v1498
		goto L501
	} else {
		goto L511
	}
L511:
	;
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(v1486)+52))
	v1502 = v1500 - int32(1)
	if v1502 < int32(0) {
		goto L508
	} else {
		goto L512
	}
L512:
	;
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v1486)+48))
	v1508 = int32(0)
	v1511 = v1502
	goto L513
L513:
	;
	v1516 = int32(2)
	v1517 = base.I32_div_s(v1511-v1508, v1516)
	v1518 = v1517 + v1508
	v1522 = *(*int32)(unsafe.Add(mBase, uint32(v1505+v1518<<(uint(v1516)%32))))
	if v1522 == v1461 {
		v1583 = v1498
		goto L501
	} else {
		goto L515
	}
L514:
	;
	goto L508
L515:
	;
	v1531 = base.B2i32(v1522-v1461 < int32(0)) | base.B2i32(base.Ui32(v1522) < base.Ui32(int32(3)))
	if v1531 != 0 {
		goto L516
	} else {
		goto L517
	}
L516:
	;
	v1532 = v1518 + int32(1)
	goto L518
L517:
	;
	v1532 = v1508
	goto L518
L518:
	;
	if v1531 != 0 {
		goto L519
	} else {
		goto L520
	}
L519:
	;
	v1535 = v1511
	goto L521
L520:
	;
	v1535 = v1518 - int32(1)
	goto L521
L521:
	;
	if v1532 <= v1535 {
		v1508 = v1532
		v1511 = v1535
		goto L513
	} else {
		goto L522
	}
L522:
	;
	goto L514
L523:
	;
	goto L507
L524:
	;
	v1561 = int32(2)
	v1562 = base.I32_div_s(v1556-v1553, v1561)
	v1563 = v1562 + v1553
	v1567 = *(*int32)(unsafe.Add(mBase, uint32(v1551+v1563<<(uint(v1561)%32))))
	v1568 = base.B2i32(v1567 == v1461)
	if v1567 == v1461 {
		v1583 = v1568
		goto L501
	} else {
		goto L526
	}
L525:
	;
	v1583 = v1568
	goto L501
L526:
	;
	v1571 = base.B2i32(base.Ui32(v1567) < base.Ui32(v1461))
	if base.Ui32(v1567) < base.Ui32(v1461) {
		goto L527
	} else {
		goto L528
	}
L527:
	;
	v1572 = v1563 + int32(1)
	goto L529
L528:
	;
	v1572 = v1553
	goto L529
L529:
	;
	if base.Ui32(v1567) < base.Ui32(v1461) {
		goto L530
	} else {
		goto L531
	}
L530:
	;
	v1575 = v1556
	goto L532
L531:
	;
	v1575 = v1563 - int32(1)
	goto L532
L532:
	;
	if v1572 <= v1575 {
		v1553 = v1572
		v1556 = v1575
		goto L524
	} else {
		goto L533
	}
L533:
	;
	goto L525
L534:
	;
	v1594 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v838))))
	v1856 = int32(base.Ui32(v1594&int32(128))>>(uint(int32(7))%32)) | base.B2i32(v1594&int32(_a_F_HeapTupleSatisfiesVisibility_0) == int32(64))
	goto L1
L535:
	;
	goto L536
L536:
	;
	v1604 = *(*int32)(unsafe.Add(mBase, uint32(v832)+4))
	v1605 = F_TransactionIdIsInProgress(m, v1604)
	mBase = m.M
	v1606 = m.ExcPending
	if v1606 != 0 {
		goto L13
	} else {
		goto L537
	}
L537:
	;
	if v1605 != 0 {
		goto L538
	} else {
		goto L539
	}
L538:
	;
	v1607 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v838))))
	if v1607&int32(128)|base.B2i32(v1607&int32(_a_F_HeapTupleSatisfiesVisibility_0) == int32(64)) != 0 {
		v1856 = v1298
		goto L1
	} else {
		goto L541
	}
L539:
	;
	goto L540
L540:
	;
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v832)+4))
	v1618 = F_TransactionIdDidCommit(m, v1617)
	mBase = m.M
	v1619 = m.ExcPending
	if v1619 != 0 {
		goto L13
	} else {
		goto L542
	}
L541:
	;
	v1615 = *(*int32)(unsafe.Add(mBase, uint32(v832)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v1615
	v1856 = v1298
	goto L1
L542:
	;
	v1620 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v832)+20)))
	if v1618 == int32(0) {
		goto L543
	} else {
		goto L544
	}
L543:
	;
	F_BufferSetHintBits16(m, v838, v1620|int32(2048), l2)
	mBase = m.M
	v1626 = m.ExcPending
	if v1626 != 0 {
		goto L13
	} else {
		goto L546
	}
L544:
	;
	goto L545
L545:
	;
	v1629 = int32(0)
	if base.B2i32(v1620&int32(128) == v1629)&base.B2i32(v1620&int32(_a_F_HeapTupleSatisfiesVisibility_0) != int32(64)) == v1629 {
		goto L547
	} else {
		goto L548
	}
L546:
	;
	v1856 = v1298
	goto L1
L547:
	;
	F_BufferSetHintBits16(m, v838, v1620|int32(2048), l2)
	mBase = m.M
	v1641 = m.ExcPending
	if v1641 != 0 {
		goto L13
	} else {
		goto L550
	}
L548:
	;
	goto L549
L549:
	;
	v1643 = *(*int32)(unsafe.Add(mBase, uint32(v832)+4))
	F_HeapTupleSetHintBits(m, v832, l2, int32(1024), v1643)
	mBase = m.M
	v1645 = m.ExcPending
	if v1645 != 0 {
		goto L13
	} else {
		goto L551
	}
L550:
	;
	v1856 = v1298
	goto L1
L551:
	;
	v1856 = int32(0)
	goto L1
L552:
	;
	if v1651 == int32(512) {
		v1856 = v4
		goto L1
	} else {
		goto L555
	}
L553:
	;
	v1657 = int32(2)
	goto L554
L554:
	;
	v1658 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v1659 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v1660 = *(*int32)(unsafe.Add(mBase, uint32(v1648)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v1657
	if v1658 == int32(0) {
		goto L556
	} else {
		goto L557
	}
L555:
	;
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(v1648)))
	v1657 = v1656
	goto L554
L556:
	;
	v1688 = int32(3)
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.B2i32(base.Ui32(v1657) < base.Ui32(v1688))|base.B2i32(base.Ui32(v1690) < base.Ui32(v1688)) == int32(0) {
		goto L564
	} else {
		goto L565
	}
L557:
	;
	v1665 = v14 + int32(12)
	v1668 = F_bsearch(m, v1665, v1659, v1658, int32(4), int32(187))
	mBase = m.M
	v1669 = m.ExcPending
	if v1669 != 0 {
		goto L13
	} else {
		goto L558
	}
L558:
	;
	if v1668 == int32(0) {
		goto L556
	} else {
		goto L559
	}
L559:
	;
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(v1648)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v1672
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = int32(-1)
	v1677 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[4]))
	v1680 = F_ResolveCminCmaxDuringDecoding(m, v1677, l1, l0, l2, v1665, v14+int32(8))
	mBase = m.M
	v1681 = m.ExcPending
	if v1681 != 0 {
		goto L13
	} else {
		goto L560
	}
L560:
	;
	if v1680 == int32(0) {
		v1856 = v4
		goto L1
	} else {
		goto L561
	}
L561:
	;
	v1684 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v1685 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if base.Ui32(v1684) < base.Ui32(v1685) {
		goto L2
	} else {
		goto L562
	}
L562:
	;
	v1856 = v4
	goto L1
L563:
	;
	v1700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1648)+21)))
	if v1700&int32(1) != 0 {
		goto L2
	} else {
		goto L569
	}
L564:
	;
	if v1657-v1690 < int32(0) {
		goto L563
	} else {
		goto L567
	}
L565:
	;
	goto L566
L566:
	;
	if base.Ui32(v1690) <= base.Ui32(v1657) {
		goto L3
	} else {
		goto L568
	}
L567:
	;
	goto L3
L568:
	;
	goto L563
L569:
	;
	v1703 = F_TransactionIdDidCommit(m, v1657)
	mBase = m.M
	v1704 = m.ExcPending
	if v1704 != 0 {
		goto L13
	} else {
		goto L570
	}
L570:
	;
	if v1703 != 0 {
		goto L2
	} else {
		goto L571
	}
L571:
	;
	v1856 = v4
	goto L1
L572:
	;
	if v1707 == int32(2) {
		goto L573
	} else {
		goto L574
	}
L573:
	;
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v1714 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v1715 = F_GlobalVisTestIsRemovableXid(m, v1713, v1714)
	mBase = m.M
	v1716 = m.ExcPending
	if v1716 != 0 {
		goto L13
	} else {
		goto L576
	}
L574:
	;
	v1718 = v1707
	goto L575
L575:
	;
	v1856 = base.B2i32(v1718 != int32(0))
	goto L1
L576:
	;
	if v1715 != 0 {
		goto L577
	} else {
		goto L578
	}
L577:
	;
	v1717 = int32(0)
	goto L579
L578:
	;
	v1717 = int32(2)
	goto L579
L579:
	;
	v1718 = v1717
	goto L575
L580:
	;
	v1856 = v4
	goto L1
L581:
	;
	v1856 = v4
	goto L1
L582:
	;
	v1743 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1744 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v1657
	if v1743 == int32(0) {
		v1856 = v4
		goto L1
	} else {
		goto L588
	}
L583:
	;
	if v1657-v1733 < int32(0) {
		goto L582
	} else {
		goto L586
	}
L584:
	;
	goto L585
L585:
	;
	if base.Ui32(v1733) <= base.Ui32(v1657) {
		v1856 = v4
		goto L1
	} else {
		goto L587
	}
L586:
	;
	v1856 = v4
	goto L1
L587:
	;
	goto L582
L588:
	;
	v1752 = F_bsearch(m, v14+int32(12), v1744, v1743, int32(4), int32(187))
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L13
	} else {
		goto L589
	}
L589:
	;
	if v1752 == int32(0) {
		v1856 = v4
		goto L1
	} else {
		goto L590
	}
L590:
	;
	goto L2
L591:
	;
	v1856 = int32(1)
	goto L1
L592:
	;
	goto L593
L593:
	;
	if v1759&int32(128) != 0 {
		goto L594
	} else {
		goto L595
	}
L594:
	;
	v1856 = int32(1)
	goto L1
L595:
	;
	goto L596
L596:
	;
	if v1759&int32(_a_F_HeapTupleSatisfiesVisibility_0) == int32(64) {
		goto L597
	} else {
		goto L598
	}
L597:
	;
	v1856 = int32(1)
	goto L1
L598:
	;
	goto L599
L599:
	;
	if v1759&int32(_a_F_HeapTupleSatisfiesVisibility_1) != 0 {
		goto L600
	} else {
		goto L601
	}
L600:
	;
	v1773 = F_HeapTupleGetUpdateXid(m, v1648)
	mBase = m.M
	v1774 = m.ExcPending
	if v1774 != 0 {
		goto L13
	} else {
		goto L603
	}
L601:
	;
	v1775 = v1660
	goto L602
L602:
	;
	v1776 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v1777 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v1775
	if v1776 == int32(0) {
		goto L604
	} else {
		goto L605
	}
L603:
	;
	v1775 = v1773
	goto L602
L604:
	;
	v1806 = int32(3)
	v1808 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.B2i32(base.Ui32(v1775) < base.Ui32(v1806))|base.B2i32(base.Ui32(v1808) < base.Ui32(v1806)) == int32(0) {
		goto L613
	} else {
		goto L614
	}
L605:
	;
	v1782 = v14 + int32(12)
	v1785 = F_bsearch(m, v1782, v1777, v1776, int32(4), int32(187))
	mBase = m.M
	v1786 = m.ExcPending
	if v1786 != 0 {
		goto L13
	} else {
		goto L606
	}
L606:
	;
	if v1785 == int32(0) {
		goto L604
	} else {
		goto L607
	}
L607:
	;
	v1789 = *(*int32)(unsafe.Add(mBase, uint32(v1648)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v1789
	v1791 = int32(1)
	v1793 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSatisfiesVisibility[4]))
	v1796 = F_ResolveCminCmaxDuringDecoding(m, v1793, l1, l0, l2, v1782, v14+int32(8))
	mBase = m.M
	v1797 = m.ExcPending
	if v1797 != 0 {
		goto L13
	} else {
		goto L608
	}
L608:
	;
	if v1796 == int32(0) {
		v1856 = v1791
		goto L1
	} else {
		goto L609
	}
L609:
	;
	v1800 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v1800 == int32(-1) {
		v1856 = v1791
		goto L1
	} else {
		goto L610
	}
L610:
	;
	v1803 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v1856 = base.B2i32(base.Ui32(v1803) <= base.Ui32(v1800))
	goto L1
L611:
	;
	v1825 = int32(3)
	v1827 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.B2i32(base.Ui32(v1775) < base.Ui32(v1825))|base.B2i32(base.Ui32(v1827) < base.Ui32(v1825)) == int32(0) {
		goto L621
	} else {
		goto L622
	}
L612:
	;
	v1818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1648)+21)))
	if v1818&int32(4) != 0 {
		v1856 = v4
		goto L1
	} else {
		goto L618
	}
L613:
	;
	if v1775-v1808 < int32(0) {
		goto L612
	} else {
		goto L616
	}
L614:
	;
	goto L615
L615:
	;
	if base.Ui32(v1808) <= base.Ui32(v1775) {
		goto L611
	} else {
		goto L617
	}
L616:
	;
	goto L611
L617:
	;
	goto L612
L618:
	;
	v1821 = F_TransactionIdDidCommit(m, v1775)
	mBase = m.M
	v1822 = m.ExcPending
	if v1822 != 0 {
		goto L13
	} else {
		goto L619
	}
L619:
	;
	v1856 = v1821 ^ int32(1)
	goto L1
L620:
	;
	v1839 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v1775
	if v1839 == int32(0) {
		goto L626
	} else {
		goto L627
	}
L621:
	;
	if v1775-v1827 < int32(0) {
		goto L620
	} else {
		goto L624
	}
L622:
	;
	goto L623
L623:
	;
	if base.Ui32(v1775) < base.Ui32(v1827) {
		goto L620
	} else {
		goto L625
	}
L624:
	;
	v1856 = int32(1)
	goto L1
L625:
	;
	v1856 = int32(1)
	goto L1
L626:
	;
	v1856 = int32(1)
	goto L1
L627:
	;
	goto L628
L628:
	;
	v1849 = F_bsearch(m, v14+int32(12), v1840, v1839, int32(4), int32(187))
	mBase = m.M
	v1850 = m.ExcPending
	if v1850 != 0 {
		goto L13
	} else {
		goto L629
	}
L629:
	;
	v1856 = base.B2i32(v1849 == int32(0))
	goto L1
}
func F_heap_attisnull(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+18)))
	if v10&int32(2047) < l1 {
		if l2 == int32(0) {
			v66 = int32(1)
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+l1<<(uint(int32(3))%32))+26)))
			if v19&int32(2) == int32(0) {
				v66 = int32(1)
			} else {
				v66 = int32(0)
			}
		}
		m.G0 = v7 + int32(16)
		return v66
	} else {
		if int32(0) < l1 {
			v27 = int32(0)
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
			if v28&int32(1) == v27 {
				v66 = v27
			} else {
				v33 = int32(1)
				v34 = l1 - v33
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+int32(base.Ui32(v34)>>(uint(int32(3))%32)))+23)))
				v66 = base.B2i32(int32(base.Ui32(v38)>>(uint(v34&int32(7))%32))&v33 == int32(0))
			}
			m.G0 = v7 + int32(16)
			return v66
		} else {
			if base.Ui32(int32(-7)) < base.Ui32(l1) {
				v66 = int32(0)
				m.G0 = v7 + int32(16)
				return v66
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
					F_errmsg_internal(m, int32(_a_F_heap_attisnull_0), v7)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_heap_attisnull_1), int32(491), int32(_a_F_heap_attisnull_2))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
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
	}
}
func F_heap_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v111 int32
	_ = v111
	var v124 int32
	_ = v124
	var v141 int32
	_ = v141
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v415 int32
	_ = v415
	var v423 int32
	_ = v423
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v483 int32
	_ = v483
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
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
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v616 int32
	_ = v616
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v630 int32
	_ = v630
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v706 int32
	_ = v706
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v756 int32
	_ = v756
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v777 int64
	_ = v777
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v804 int32
	_ = v804
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int64
	_ = v816
	var v818 int32
	_ = v818
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v847 int32
	_ = v847
	v8 = l7
	v9 = l8
	v10 = l9
	v24 = m.G0
	v26 = v24 - int32(32)
	m.G0 = v26
	if l11 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L82
	} else {
		goto L197
	}
L2:
	;
	v51 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l12))) = v51
	*(*int32)(unsafe.Add(mBase, uint32(l13))) = v51
	if v8 == int32(105) {
		goto L18
	} else {
		goto L19
	}
L3:
	;
	goto L4
L4:
	;
	if l1 == int32(11) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v33 = base.B2i32(v8 != int32(105))
	goto L7
L6:
	;
	v33 = int32(0)
	goto L7
L7:
	;
	if v33 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if l1 != int32(99) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L10
L10:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[0]))
	if v48 == int32(2) {
		goto L1
	} else {
		goto L17
	}
L11:
	;
	if v40 == int32(0) {
		goto L2
	} else {
		goto L15
	}
L12:
	;
	v38 = F_isTempToastNamespace(m, l1)
	mBase = m.M
	v40 = v38
	goto L14
L13:
	;
	v40 = int32(1)
	goto L14
L14:
	;
	goto L11
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[0]))
	if v44 != int32(2) {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	goto L1
L17:
	;
	goto L2
L18:
	;
	v58 = l2
	goto L20
L19:
	;
	v58 = v51
	goto L20
L20:
	;
	if v8 == int32(114) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v61 = l2
	goto L23
L22:
	;
	v61 = v58
	goto L23
L23:
	;
	if v8 == int32(83) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v64 = l2
	goto L26
L25:
	;
	v64 = v61
	goto L26
L26:
	;
	if v8 == int32(116) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v67 = l2
	goto L29
L28:
	;
	v67 = v64
	goto L29
L29:
	;
	if v8 == int32(109) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v70 = l2
	goto L32
L31:
	;
	v70 = v67
	goto L32
L32:
	;
	if v8 == int32(73) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v73 = l2
	goto L35
L34:
	;
	v73 = v70
	goto L35
L35:
	;
	if v8 == int32(112) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v76 = l2
	goto L38
L37:
	;
	v76 = v73
	goto L38
L38:
	;
	if v8 != int32(83) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v80 = v76
	goto L41
L40:
	;
	v80 = int32(0)
	goto L41
L41:
	;
	switch v8 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L43
	default:
		v85 = l4
		v86 = int32(0)
		goto L42
	}
L42:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[1]))
	if v80 != v89 {
		goto L48
	} else {
		goto L49
	}
L43:
	;
	if l4 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v84 = l4
	goto L46
L45:
	;
	v84 = l3
	goto L46
L46:
	;
	v85 = v84
	v86 = l14
	goto L42
L47:
	;
	if v86 != 0 {
		goto L184
	} else {
		goto L185
	}
L48:
	;
	v91 = v80
	goto L50
L49:
	;
	v91 = int32(0)
	goto L50
L50:
	;
	v92 = m.G0
	v94 = v92 - int32(48)
	m.G0 = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v99 = int32(1)
	if l3 <= int32(3591) {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	if v170 == v10 {
		goto L76
	} else {
		goto L77
	}
L52:
	;
	goto L51
L53:
	;
	v170 = int32(0)
	goto L52
L54:
	;
	if base.B2i32(base.Ui32(l3-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(l3-int32(2846)) < base.Ui32(int32(2))) != 0 {
		v170 = v99
		goto L52
	} else {
		goto L75
	}
L55:
	;
	if l3 <= int32(2670) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	if l3 <= int32(_a_F_heap_create_0) {
		goto L65
	} else {
		goto L66
	}
L58:
	;
	switch l3 - int32(1213) {
	case 0, 1, 19, 20, 47, 48, 49:
		v170 = v99
		goto L52
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
		goto L53
	default:
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v111 = l3 - int32(2671)
	if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v111))|base.B2i32(int32(1)<<(uint(v111)%32)&int32(226492515) == int32(0)) != 0 {
		goto L54
	} else {
		goto L63
	}
L61:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l3-int32(2396)) {
		goto L53
	} else {
		goto L62
	}
L62:
	;
	v170 = v99
	goto L52
L63:
	;
	v170 = v99
	goto L52
L64:
	;
	if base.Ui32(l3-int32(3592)) < base.Ui32(int32(2)) {
		v170 = v99
		goto L52
	} else {
		goto L73
	}
L65:
	;
	v124 = l3 - int32(_a_F_heap_create_1)
	if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v124))|base.B2i32(int32(1)<<(uint(v124)%32)&int32(963) == int32(0)) != 0 {
		goto L64
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	switch l3 - int32(_a_F_heap_create_2) {
	case 0, 1, 2, 3, 4, 59, 60:
		v170 = v99
		goto L52
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L53
	default:
		goto L69
	}
L68:
	;
	v170 = v99
	goto L52
L69:
	;
	if base.Ui32(l3-int32(_a_F_heap_create_3)) < base.Ui32(int32(3)) {
		v170 = v99
		goto L52
	} else {
		goto L70
	}
L70:
	;
	v141 = l3 - int32(_a_F_heap_create_4)
	if base.Ui32(int32(15)) < base.Ui32(v141) {
		goto L53
	} else {
		goto L71
	}
L71:
	;
	if int32(1)<<(uint(v141)%32)&int32(_a_F_heap_create_5) != 0 {
		v170 = v99
		goto L52
	} else {
		goto L72
	}
L72:
	;
	goto L53
L73:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l3-int32(4060)) {
		goto L53
	} else {
		goto L74
	}
L74:
	;
	v170 = v99
	goto L52
L75:
	;
	goto L53
L76:
	;
	v173 = l3 - int32(1247)
	v178 = base.B2i32(base.Ui32(v173) < base.Ui32(int32(16))) & int32(base.Ui32(int32(_a_F_heap_create_6))>>(uint(v173)%32))
	v180 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[2]))
	if v180 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	goto L78
L78:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L82
	} else {
		goto L180
	}
L79:
	;
	F_CreateCacheMemoryContext(m)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	v189 = v180
	goto L81
L81:
	;
	v190 = int32(_a_F_heap_create_7)
	v191 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_create[3])) = v189
	v195 = F_palloc0(m, int32(276))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L82
	} else {
		goto L84
	}
L82:
	;
	return int32(0)
L83:
	;
	v188 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[2]))
	v189 = v188
	goto L81
L84:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v195)+25)) = uint8(v178)
	*(*int32)(unsafe.Add(mBase, uint32(v195)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v195)+16)) = v178
	v202 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[4]))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+8))
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195)+44)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v195)+36)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v195)+32)) = v203
	v209 = F_CreateTupleDescCopy(m, l6)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L82
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195)+52)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v209)+12)) = int32(1)
	if v96 <= int32(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v316 = F_palloc0(m, int32(144))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L82
	} else {
		goto L98
	}
L88:
	;
	v231 = int32(0)
	v239 = int32(0)
	goto L89
L89:
	;
	v241 = v231 * int32(100)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v195)+52))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	v244 = int32(3)
	v247 = v241 + (v242 + v243<<(uint(v244)%32))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v252 = l6 + v248<<(uint(v244)%32) + v241
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252)+117)))
	*(*uint8)(unsafe.Add(mBase, uint32(v247)+117)) = uint8(v253)
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252)+118)))
	*(*uint8)(unsafe.Add(mBase, uint32(v247)+118)) = uint8(v255)
	v258 = v252 + int32(114)
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258))))
	*(*uint8)(unsafe.Add(mBase, uint32(v247)+114)) = uint8(v259)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v195)+52))
	F_populate_compact_attribute(m, v261, v231)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L82
	} else {
		goto L91
	}
L90:
	;
	if v264&int32(255) == int32(0) {
		goto L87
	} else {
		goto L96
	}
L91:
	;
	v264 = v259 | v239
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258))))
	if v265 == int32(1) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v269 = v231 << (uint(int32(3)) % 32)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v195)+52))
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6+v269)+35)))
	*(*uint8)(unsafe.Add(mBase, uint32(v269+v270)+35)) = uint8(v273)
	goto L94
L93:
	;
	goto L94
L94:
	;
	v276 = int32(1)
	v279 = v231 + v276
	if v279 != v96 {
		v231 = v279
		v239 = v264 & v276
		goto L89
	} else {
		goto L95
	}
L95:
	;
	goto L90
L96:
	;
	v286 = F_palloc0(m, int32(20))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L82
	} else {
		goto L97
	}
L97:
	;
	v288 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v286)+16)) = uint8(v288)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v195)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v290)+24)) = v286
	goto L87
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195)+48)) = v316
	v322 = F_strncpy(m, v316+int32(4), l0, int32(64))
	mBase = m.M
	v323 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v322)+63)) = uint8(v323)
	goto L99
L99:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v325)+68)) = l1
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v327)+119)) = uint8(v8)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*uint16)(unsafe.Add(mBase, uint32(v329)+120)) = uint16(v96)
	v331 = int32(0)
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v332)+72)) = v331
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v335)+80)) = int32(10)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v338)+118)) = uint8(v9)
	switch v9 - int32(112) {
	case 0, 5:
		v366 = int32(-1)
		v367 = v331
		goto L100
	default:
		goto L102
	case 4:
		goto L101
	}
L100:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v195)+24)) = uint8(v367)
	*(*int32)(unsafe.Add(mBase, uint32(v195)+20)) = v366
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v370)+129)) = uint8(base.B2i32(v8 != int32(109)))
	v374 = int32(110)
	goto L110
L101:
	;
	v359 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[5]))
	v361 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[6]))
	if v361 == int32(-1) {
		goto L106
	} else {
		goto L107
	}
L102:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L82
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94)+16)) = v9
	F_errmsg_internal(m, int32(_a_F_heap_create_8), v94+int32(16))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L82
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_heap_create_9), int32(3659), int32(_a_F_heap_create_10))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L82
	} else {
		goto L105
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L106:
	;
	v364 = v359
	goto L108
L107:
	;
	v364 = v361
	goto L108
L108:
	;
	v366 = v364
	v367 = int32(1)
	goto L100
L109:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v393)+130)) = uint8(v392)
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v395)+117)) = uint8(v10)
	*(*int32)(unsafe.Add(mBase, uint32(v195)+56)) = l3
	if v96 <= int32(0) {
		goto L113
	} else {
		goto L114
	}
L110:
	;
	if l1 == int32(11) {
		v392 = v374
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v378 = v8 - int32(109)
	if base.Ui32(int32(5)) < base.Ui32(v378&int32(255)) {
		v392 = v374
		goto L109
	} else {
		goto L112
	}
L112:
	;
	v392 = base.I32_wrap_i64(int64(base.Ui64(int64(110425294138980)) >> (uint(base.I64_extend_i32_u(v378<<(uint(int32(3))%32))&int64(248)) % 64)))
	goto L109
L113:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v195)+52))
	v553 = int32(0)
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v552)))
	if v553 < v562 {
		goto L126
	} else {
		goto L127
	}
L114:
	;
	v401 = v96 & int32(3)
	v402 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v96) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v415 = int32(0)
	v423 = v402
	goto L118
L116:
	;
	v483 = v402
	goto L117
L117:
	;
	v506 = v483
	v508 = v402
	goto L122
L118:
	;
	v433 = v423 * int32(100)
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v195)+52))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v434)))
	v436 = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v433+(v434+v435<<(uint(v436)%32)))+28)) = l3
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v195)+52))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v441)))
	*(*int32)(unsafe.Add(mBase, uint32(v441+v442<<(uint(v436)%32)+v433)+128)) = l3
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v195)+52))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)))
	*(*int32)(unsafe.Add(mBase, uint32(v448+v449<<(uint(v436)%32)+v433)+228)) = l3
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v195)+52))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v455)))
	*(*int32)(unsafe.Add(mBase, uint32(v433+(v455+v456<<(uint(v436)%32)))+328)) = l3
	v462 = int32(4)
	v463 = v423 + v462
	v465 = v415 + v462
	if v465 != v96&int32(2147483644) {
		v415 = v465
		v423 = v463
		goto L118
	} else {
		goto L120
	}
L119:
	;
	if v401 == int32(0) {
		goto L113
	} else {
		goto L121
	}
L120:
	;
	goto L119
L121:
	;
	v483 = v463
	goto L117
L122:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v195)+52))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v515)))
	*(*int32)(unsafe.Add(mBase, uint32(v515+v516<<(uint(int32(3))%32)+v506*int32(100))+28)) = l3
	v524 = int32(1)
	v527 = v508 + v524
	if v527 != v401 {
		v506 = v506 + v524
		v508 = v527
		goto L122
	} else {
		goto L124
	}
L123:
	;
	goto L113
L124:
	;
	goto L123
L125:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v640)+92)) = v91
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	if l10 != 0 {
		goto L145
	} else {
		goto L146
	}
L126:
	;
	v566 = v552 + int32(28)
	v573 = v553
	v574 = v562
	v576 = v553
	goto L130
L127:
	;
	v630 = v553
	v637 = v562
	goto L128
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v552)+20)) = v637
	*(*int32)(unsafe.Add(mBase, uint32(v552)+16)) = v630
	goto L125
L129:
	;
	v630 = v624
	v637 = v603
	goto L128
L130:
	;
	v582 = v566 + v562<<(uint(int32(3))%32) + v573*int32(100)
	v585 = v566 + v573<<(uint(int32(3))%32)
	if v562 != v574 {
		v603 = v574
		goto L132
	} else {
		goto L133
	}
L131:
	;
	v624 = v562
	goto L129
L132:
	;
	v604 = int32(*(*int16)(unsafe.Add(mBase, uint32(v585)+2)))
	if v604 <= int32(0) {
		v624 = v573
		goto L129
	} else {
		goto L140
	}
L133:
	;
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585)+7)))
	if v587 != int32(118) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v603 = v573
	goto L132
L135:
	;
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585)+4)))
	if v590 != int32(1) {
		goto L134
	} else {
		goto L136
	}
L136:
	;
	v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585)+6)))
	if v593&int32(6) != 0 {
		goto L134
	} else {
		goto L137
	}
L137:
	;
	v596 = int32(*(*int16)(unsafe.Add(mBase, uint32(v585)+2)))
	if v596 <= int32(0) {
		goto L134
	} else {
		goto L138
	}
L138:
	;
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582)+90)))
	if v599 != int32(118) {
		v603 = v562
		goto L132
	} else {
		goto L139
	}
L139:
	;
	goto L134
L140:
	;
	v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582)+90)))
	if v607 == int32(118) {
		v624 = v573
		goto L129
	} else {
		goto L141
	}
L141:
	;
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585)+5)))
	v616 = (v576 + v610 - int32(1)) & (int32(0) - v610)
	if int32(_a_F_heap_create_11) < v616 {
		v624 = v573
		goto L129
	} else {
		goto L142
	}
L142:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v585))) = uint16(v616)
	v622 = v573 + int32(1)
	if v622 != v562 {
		v573 = v622
		v574 = v603
		v576 = v616 + v604
		goto L130
	} else {
		goto L143
	}
L143:
	;
	goto L131
L144:
	;
	v650 = v195 + int32(56)
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v195)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v195)+60)) = v651
	v655 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[7]))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656)+117)))
	if v657 != 0 {
		goto L150
	} else {
		goto L151
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v642)+88)) = int32(0)
	F_RelationMapUpdateMap(m, l3, v85, v10, int32(1))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L82
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v642)+88)) = v85
	goto L144
L148:
	;
	goto L144
L149:
	;
	F_RelationInitPhysicalAddr(m, v195)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L82
	} else {
		goto L153
	}
L150:
	;
	v658 = int32(0)
	goto L152
L151:
	;
	v658 = v655
	goto L152
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195)+64)) = v658
	goto L149
L153:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v662)+84)) = l5
	*(*int32)(unsafe.Add(mBase, _c_F_heap_create[3])) = v191
	switch v8 - int32(83) {
	case 0, 26, 31, 33:
		goto L155
	default:
		goto L154
	}
L154:
	;
	v671 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[8]))
	v675 = F_hash_search(m, v671, v650, int32(1), v94+int32(47))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L82
	} else {
		goto L157
	}
L155:
	;
	F_RelationInitTableAccessMethod(m, v195)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L82
	} else {
		goto L156
	}
L156:
	;
	goto L154
L157:
	;
	v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+47)))
	if v677 == int32(1) {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v715 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[9]))
	if v715 <= int32(31) {
		goto L172
	} else {
		goto L173
	}
L159:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v675)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v675)+4)) = v195
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v680)+16))
	if v682 == int32(0) {
		goto L162
	} else {
		goto L163
	}
L160:
	;
	goto L161
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v675)+4)) = v195
	goto L158
L162:
	;
	F_RelationDestroyRelation(m, v680, int32(0))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L82
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	v689 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[0]))
	if v689 == int32(0) {
		goto L158
	} else {
		goto L166
	}
L165:
	;
	goto L158
L166:
	;
	v694 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L82
	} else {
		goto L167
	}
L167:
	;
	if v694 == int32(0) {
		goto L158
	} else {
		goto L168
	}
L168:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v680)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v94)+32)) = v698 + int32(4)
	F_errmsg_internal(m, int32(_a_F_heap_create_12), v94+int32(32))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L82
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(_a_F_heap_create_9), int32(3734), int32(_a_F_heap_create_10))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L82
	} else {
		goto L170
	}
L170:
	;
	goto L158
L171:
	;
	v732 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v195)+26)) = uint8(v732)
	v735 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[10]))
	F_ResourceOwnerEnlarge(m, v735)
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L82
	} else {
		goto L175
	}
L172:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v650)))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_create[9])) = v715 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v715<<(uint(int32(2))%32))+uint32(_c_F_heap_create[11]))) = v718
	goto L171
L173:
	;
	goto L174
L174:
	;
	v729 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_heap_create[12])) = uint8(v729)
	goto L171
L175:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v195)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v195)+16)) = v738 + int32(1)
	v743 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[0]))
	if v743 != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v745 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[10]))
	F_ResourceOwnerRemember(m, v745, base.I64_extend_i32_u(v195), int32(_a_F_heap_create_13))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L82
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	m.G0 = v94 + int32(48)
	goto L47
L179:
	;
	goto L178
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = l0
	F_errmsg_internal(m, int32(_a_F_heap_create_14), v94)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L82
	} else {
		goto L181
	}
L181:
	;
	F_errfinish(m, int32(_a_F_heap_create_9), int32(3560), int32(_a_F_heap_create_10))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L82
	} else {
		goto L182
	}
L182:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L183:
	;
	v812 = *(*int32)(unsafe.Add(mBase, _c_F_heap_create[7]))
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	v814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v813)+117)))
	if v814 != 0 {
		goto L193
	} else {
		goto L194
	}
L184:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	v768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v767)+119)))
	switch v768 - int32(83) {
	case 0, 22:
		goto L187
	default:
		goto L183
	case 26, 31, 33:
		goto L188
	}
L185:
	;
	goto L186
L186:
	;
	if v91 == int32(0) {
		goto L183
	} else {
		goto L191
	}
L187:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v195)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v775
	v777 = *(*int64)(unsafe.Add(mBase, uint32(v195)))
	*(*int64)(unsafe.Add(mBase, uint32(v26))) = v777
	v780 = F_RelationCreateStorage(m, v26, v9, int32(1))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L82
	} else {
		goto L190
	}
L188:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v195)+188))
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v771)+112))
	m.T0[v772].(func(*base.Module, int32, int32, int32, int32, int32))(m, v195, v195, v9, l12, l13)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L82
	} else {
		goto L189
	}
L189:
	;
	goto L183
L190:
	;
	goto L183
L191:
	;
	v784 = m.G0
	v786 = v784 - int32(32)
	m.G0 = v786
	v788 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v786)+28)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v786)+24)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v786)+20)) = int32(1259)
	*(*int32)(unsafe.Add(mBase, uint32(v786)+16)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v786)+12)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v786)+8)) = int32(1213)
	F_recordSharedDependencyOn(m, v786+int32(20), v786+int32(8), int32(116))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L82
	} else {
		goto L192
	}
L192:
	;
	m.G0 = v786 + int32(32)
	goto L183
L193:
	;
	v815 = int32(0)
	goto L195
L194:
	;
	v815 = v812
	goto L195
L195:
	;
	v816 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v195)+56)))
	F_pgstat_create_transactional(m, int32(2), v815, v816)
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L82
	} else {
		goto L196
	}
L196:
	;
	m.G0 = v26 + int32(32)
	return v195
L197:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L82
	} else {
		goto L198
	}
L198:
	;
	v830 = F_get_namespace_name(m, l1)
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L82
	} else {
		goto L199
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v830
	F_errmsg(m, int32(_a_F_heap_create_15), v26+int32(16))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L82
	} else {
		goto L200
	}
L200:
	;
	v841 = F_errdetail(m, int32(_a_F_heap_create_16), int32(0))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L82
	} else {
		goto L201
	}
L201:
	;
	F_errfinish(m, int32(_a_F_heap_create_17), int32(324), int32(_a_F_heap_create_18))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L82
	} else {
		goto L202
	}
L202:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_heap_decode(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
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
	var v19 int64
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
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
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
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
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
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
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int64
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int64
	_ = v184
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int64
	_ = v205
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v223 int64
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int64
	_ = v321
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int64
	_ = v375
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v391 int32
	_ = v391
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int64
	_ = v422
	var v428 int32
	_ = v428
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int64
	_ = v447
	var v450 int32
	_ = v450
	var v462 int64
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v563 int64
	_ = v563
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int64
	_ = v588
	var v594 int32
	_ = v594
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int64
	_ = v612
	var v615 int32
	_ = v615
	var v625 int64
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int64
	_ = v733
	var v736 int32
	_ = v736
	var v740 int64
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v829 int64
	_ = v829
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int64
	_ = v838
	var v841 int32
	_ = v841
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+96))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+48)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	v19 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	F_ReorderBufferProcessXid(m, v17, v18, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
		if v22 <= int32(0) {
			return
		} else {
			switch int32(base.Ui32(v15)>>(uint(int32(4))%32))&int32(7) - int32(1) {
			case 0:
				v462 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				v463 = F_SnapBuildProcessChange(m, v16, v18, v462)
				mBase = m.M
				v464 = m.ExcPending
				if v464 != 0 {
					return
				} else {
					if v463 == int32(0) {
						return
					} else {
						v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
						if v467 != 0 {
							return
						} else {
							v468 = int32(0)
							v470 = m.G0
							v472 = v470 - int32(16)
							m.G0 = v472
							v475 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[0]))
							if v475 == v468 {
								v518 = v468
							} else {
								v478 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
								v479 = int32(0)
								v485 = F_XLogRecGetBlockTagExtended(m, v478, v479, v472+int32(4), v479, v479, v479)
								mBase = m.M
								if v485 == v479 {
									v518 = v468
								} else {
									v488 = *(*int32)(unsafe.Add(mBase, uint32(v472)+12))
									v490 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[0]))
									if v488 != v490 {
										v501 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[1]))
										if base.B2i32(v501 == int32(0))|base.B2i32(v488 != v501) != 0 {
											v518 = int32(1)
										} else {
											v507 = *(*int32)(unsafe.Add(mBase, uint32(v472)+8))
											v509 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[2]))
											if v507 != v509 {
												v518 = int32(1)
											} else {
												v511 = *(*int32)(unsafe.Add(mBase, uint32(v472)+4))
												v513 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[3]))
												if v511 == v513 {
													v518 = int32(0)
												} else {
													v518 = int32(1)
												}
											}
										}
									} else {
										v492 = *(*int32)(unsafe.Add(mBase, uint32(v472)+8))
										v494 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[4]))
										if v492 != v494 {
											v501 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[1]))
											if base.B2i32(v501 == int32(0))|base.B2i32(v488 != v501) != 0 {
												v518 = int32(1)
											} else {
												v507 = *(*int32)(unsafe.Add(mBase, uint32(v472)+8))
												v509 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[2]))
												if v507 != v509 {
													v518 = int32(1)
												} else {
													v511 = *(*int32)(unsafe.Add(mBase, uint32(v472)+4))
													v513 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[3]))
													if v511 == v513 {
														v518 = int32(0)
													} else {
														v518 = int32(1)
													}
												}
											}
										} else {
											v496 = *(*int32)(unsafe.Add(mBase, uint32(v472)+4))
											v498 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[5]))
											if v496 == v498 {
												v518 = v468
											} else {
												v501 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[1]))
												if base.B2i32(v501 == int32(0))|base.B2i32(v488 != v501) != 0 {
													v518 = int32(1)
												} else {
													v507 = *(*int32)(unsafe.Add(mBase, uint32(v472)+8))
													v509 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[2]))
													if v507 != v509 {
														v518 = int32(1)
													} else {
														v511 = *(*int32)(unsafe.Add(mBase, uint32(v472)+4))
														v513 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[3]))
														if v511 == v513 {
															v518 = int32(0)
														} else {
															v518 = int32(1)
														}
													}
												}
											}
										}
									}
								}
							}
							m.G0 = v472 + int32(16)
							if v518 != 0 {
								return
							} else {
								v522 = m.G0
								v524 = v522 - int32(16)
								m.G0 = v524
								v526 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
								v527 = *(*int32)(unsafe.Add(mBase, uint32(v526)+96))
								v528 = *(*int32)(unsafe.Add(mBase, uint32(v527)+64))
								v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528)+7)))
								if v529&int32(32) != 0 {
									m.G0 = v524 + int32(16)
									return
								} else {
									v532 = int32(0)
									F_XLogRecGetBlockTag(m, v526, v532, v524+int32(4), v532, v532)
									mBase = m.M
									v538 = m.ExcPending
									if v538 != 0 {
										return
									} else {
										v539 = *(*int32)(unsafe.Add(mBase, uint32(v524)+8))
										v540 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v541 = *(*int32)(unsafe.Add(mBase, uint32(v540)+88))
										if v539 != v541 {
											m.G0 = v524 + int32(16)
											return
										} else {
											v543 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											if v543 != 0 {
												v544 = *(*int32)(unsafe.Add(mBase, uint32(v526)+96))
												v545 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v544)+56)))
												v546 = F_filter_by_origin_cb_wrapper(m, l0, v545)
												mBase = m.M
												v547 = m.ExcPending
												if v547 != 0 {
													return
												} else {
													if v546 != 0 {
														m.G0 = v524 + int32(16)
														return
													} else {
														v548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v549 = F_ReorderBufferAllocChange(m, v548)
														mBase = m.M
														v550 = m.ExcPending
														if v550 != 0 {
															return
														} else {
															v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528)+7)))
															if v553&int32(8) != 0 {
																v556 = int32(10)
															} else {
																v556 = int32(2)
															}
															*(*int32)(unsafe.Add(mBase, uint32(v549)+8)) = v556
															v558 = *(*int32)(unsafe.Add(mBase, uint32(v526)+96))
															v559 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v558)+56)))
															*(*uint16)(unsafe.Add(mBase, uint32(v549)+16)) = uint16(v559)
															v561 = *(*int32)(unsafe.Add(mBase, uint32(v524)+12))
															*(*int32)(unsafe.Add(mBase, uint32(v549)+28)) = v561
															v563 = *(*int64)(unsafe.Add(mBase, uint32(v524)+4))
															*(*int64)(unsafe.Add(mBase, uint32(v549)+20)) = v563
															v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528)+7)))
															if v565&int32(6) != 0 {
																v568 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																v569 = *(*int32)(unsafe.Add(mBase, uint32(v526)+96))
																v570 = *(*int32)(unsafe.Add(mBase, uint32(v569)+68))
																v572 = v570 - int32(13)
																v573 = F_ReorderBufferAllocTupleBuf(m, v568, v572)
																mBase = m.M
																v574 = m.ExcPending
																if v574 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v549)+36)) = v573
																	v576 = int32(0)
																	*(*int32)(unsafe.Add(mBase, uint32(v573)+12)) = v576
																	*(*uint16)(unsafe.Add(mBase, uint32(v573)+8)) = uint16(v576)
																	*(*int32)(unsafe.Add(mBase, uint32(v573)+4)) = int32(-1)
																	*(*int32)(unsafe.Add(mBase, uint32(v573))) = v570 + int32(10)
																	v585 = *(*int32)(unsafe.Add(mBase, uint32(v528)+8))
																	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528)+12)))
																	v587 = *(*int32)(unsafe.Add(mBase, uint32(v573)+16))
																	v588 = int64(0)
																	*(*int64)(unsafe.Add(mBase, uint32(v587)+8)) = v588
																	*(*int64)(unsafe.Add(mBase, uint32(v587)+15)) = v588
																	*(*int64)(unsafe.Add(mBase, uint32(v587))) = v588
																	if v572 != 0 {
																		v594 = *(*int32)(unsafe.Add(mBase, uint32(v573)+16))
																		base.MemoryCopy(m, v594+int32(23), v528+int32(13), v572)
																	} else {
																	}
																	*(*uint8)(unsafe.Add(mBase, uint32(v587)+22)) = uint8(v586)
																	*(*int32)(unsafe.Add(mBase, uint32(v587)+18)) = v585
																	v607 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(v549)+32)) = uint8(v607)
																	v609 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																	v610 = *(*int32)(unsafe.Add(mBase, uint32(v526)+96))
																	v611 = *(*int32)(unsafe.Add(mBase, uint32(v610)+36))
																	v612 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
																	F_ReorderBufferQueueChange(m, v609, v611, v612, v549, int32(0))
																	mBase = m.M
																	v615 = m.ExcPending
																	if v615 != 0 {
																		return
																	} else {
																		m.G0 = v524 + int32(16)
																		return
																	}
																}
															} else {
																v607 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v549)+32)) = uint8(v607)
																v609 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																v610 = *(*int32)(unsafe.Add(mBase, uint32(v526)+96))
																v611 = *(*int32)(unsafe.Add(mBase, uint32(v610)+36))
																v612 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
																F_ReorderBufferQueueChange(m, v609, v611, v612, v549, int32(0))
																mBase = m.M
																v615 = m.ExcPending
																if v615 != 0 {
																	return
																} else {
																	m.G0 = v524 + int32(16)
																	return
																}
															}
														}
													}
												}
											} else {
												v548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v549 = F_ReorderBufferAllocChange(m, v548)
												mBase = m.M
												v550 = m.ExcPending
												if v550 != 0 {
													return
												} else {
													v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528)+7)))
													if v553&int32(8) != 0 {
														v556 = int32(10)
													} else {
														v556 = int32(2)
													}
													*(*int32)(unsafe.Add(mBase, uint32(v549)+8)) = v556
													v558 = *(*int32)(unsafe.Add(mBase, uint32(v526)+96))
													v559 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v558)+56)))
													*(*uint16)(unsafe.Add(mBase, uint32(v549)+16)) = uint16(v559)
													v561 = *(*int32)(unsafe.Add(mBase, uint32(v524)+12))
													*(*int32)(unsafe.Add(mBase, uint32(v549)+28)) = v561
													v563 = *(*int64)(unsafe.Add(mBase, uint32(v524)+4))
													*(*int64)(unsafe.Add(mBase, uint32(v549)+20)) = v563
													v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528)+7)))
													if v565&int32(6) != 0 {
														v568 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v569 = *(*int32)(unsafe.Add(mBase, uint32(v526)+96))
														v570 = *(*int32)(unsafe.Add(mBase, uint32(v569)+68))
														v572 = v570 - int32(13)
														v573 = F_ReorderBufferAllocTupleBuf(m, v568, v572)
														mBase = m.M
														v574 = m.ExcPending
														if v574 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v549)+36)) = v573
															v576 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v573)+12)) = v576
															*(*uint16)(unsafe.Add(mBase, uint32(v573)+8)) = uint16(v576)
															*(*int32)(unsafe.Add(mBase, uint32(v573)+4)) = int32(-1)
															*(*int32)(unsafe.Add(mBase, uint32(v573))) = v570 + int32(10)
															v585 = *(*int32)(unsafe.Add(mBase, uint32(v528)+8))
															v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528)+12)))
															v587 = *(*int32)(unsafe.Add(mBase, uint32(v573)+16))
															v588 = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(v587)+8)) = v588
															*(*int64)(unsafe.Add(mBase, uint32(v587)+15)) = v588
															*(*int64)(unsafe.Add(mBase, uint32(v587))) = v588
															if v572 != 0 {
																v594 = *(*int32)(unsafe.Add(mBase, uint32(v573)+16))
																base.MemoryCopy(m, v594+int32(23), v528+int32(13), v572)
															} else {
															}
															*(*uint8)(unsafe.Add(mBase, uint32(v587)+22)) = uint8(v586)
															*(*int32)(unsafe.Add(mBase, uint32(v587)+18)) = v585
															v607 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v549)+32)) = uint8(v607)
															v609 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															v610 = *(*int32)(unsafe.Add(mBase, uint32(v526)+96))
															v611 = *(*int32)(unsafe.Add(mBase, uint32(v610)+36))
															v612 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
															F_ReorderBufferQueueChange(m, v609, v611, v612, v549, int32(0))
															mBase = m.M
															v615 = m.ExcPending
															if v615 != 0 {
																return
															} else {
																m.G0 = v524 + int32(16)
																return
															}
														}
													} else {
														v607 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v549)+32)) = uint8(v607)
														v609 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v610 = *(*int32)(unsafe.Add(mBase, uint32(v526)+96))
														v611 = *(*int32)(unsafe.Add(mBase, uint32(v610)+36))
														v612 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
														F_ReorderBufferQueueChange(m, v609, v611, v612, v549, int32(0))
														mBase = m.M
														v615 = m.ExcPending
														if v615 != 0 {
															return
														} else {
															m.G0 = v524 + int32(16)
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
			case 1, 3:
				v223 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				v224 = F_SnapBuildProcessChange(m, v16, v18, v223)
				mBase = m.M
				v225 = m.ExcPending
				if v225 != 0 {
					return
				} else {
					if v224 == int32(0) {
						return
					} else {
						v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
						if v228 != 0 {
							return
						} else {
							v229 = int32(0)
							v231 = m.G0
							v233 = v231 - int32(16)
							m.G0 = v233
							v236 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[0]))
							if v236 == v229 {
								v279 = v229
							} else {
								v239 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
								v240 = int32(0)
								v246 = F_XLogRecGetBlockTagExtended(m, v239, v240, v233+int32(4), v240, v240, v240)
								mBase = m.M
								if v246 == v240 {
									v279 = v229
								} else {
									v249 = *(*int32)(unsafe.Add(mBase, uint32(v233)+12))
									v251 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[0]))
									if v249 != v251 {
										v262 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[1]))
										if base.B2i32(v262 == int32(0))|base.B2i32(v249 != v262) != 0 {
											v279 = int32(1)
										} else {
											v268 = *(*int32)(unsafe.Add(mBase, uint32(v233)+8))
											v270 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[2]))
											if v268 != v270 {
												v279 = int32(1)
											} else {
												v272 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
												v274 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[3]))
												if v272 == v274 {
													v279 = int32(0)
												} else {
													v279 = int32(1)
												}
											}
										}
									} else {
										v253 = *(*int32)(unsafe.Add(mBase, uint32(v233)+8))
										v255 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[4]))
										if v253 != v255 {
											v262 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[1]))
											if base.B2i32(v262 == int32(0))|base.B2i32(v249 != v262) != 0 {
												v279 = int32(1)
											} else {
												v268 = *(*int32)(unsafe.Add(mBase, uint32(v233)+8))
												v270 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[2]))
												if v268 != v270 {
													v279 = int32(1)
												} else {
													v272 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
													v274 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[3]))
													if v272 == v274 {
														v279 = int32(0)
													} else {
														v279 = int32(1)
													}
												}
											}
										} else {
											v257 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
											v259 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[5]))
											if v257 == v259 {
												v279 = v229
											} else {
												v262 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[1]))
												if base.B2i32(v262 == int32(0))|base.B2i32(v249 != v262) != 0 {
													v279 = int32(1)
												} else {
													v268 = *(*int32)(unsafe.Add(mBase, uint32(v233)+8))
													v270 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[2]))
													if v268 != v270 {
														v279 = int32(1)
													} else {
														v272 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
														v274 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[3]))
														if v272 == v274 {
															v279 = int32(0)
														} else {
															v279 = int32(1)
														}
													}
												}
											}
										}
									}
								}
							}
							m.G0 = v233 + int32(16)
							if v279 != 0 {
								return
							} else {
								v283 = m.G0
								v284 = int32(16)
								v285 = v283 - v284
								m.G0 = v285
								v287 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
								v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+96))
								v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)+64))
								v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289)+7)))
								if v290&v284 == int32(0) {
									m.G0 = v285 + int32(16)
									return
								} else {
									v295 = int32(0)
									F_XLogRecGetBlockTag(m, v287, v295, v285+int32(4), v295, v295)
									mBase = m.M
									v301 = m.ExcPending
									if v301 != 0 {
										return
									} else {
										v302 = *(*int32)(unsafe.Add(mBase, uint32(v285)+8))
										v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v304 = *(*int32)(unsafe.Add(mBase, uint32(v303)+88))
										if v302 != v304 {
											m.G0 = v285 + int32(16)
											return
										} else {
											v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											if v306 != 0 {
												v307 = *(*int32)(unsafe.Add(mBase, uint32(v287)+96))
												v308 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v307)+56)))
												v309 = F_filter_by_origin_cb_wrapper(m, l0, v308)
												mBase = m.M
												v310 = m.ExcPending
												if v310 != 0 {
													return
												} else {
													if v309 != 0 {
														m.G0 = v285 + int32(16)
														return
													} else {
														v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v312 = F_ReorderBufferAllocChange(m, v311)
														mBase = m.M
														v313 = m.ExcPending
														if v313 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v312)+8)) = int32(1)
															v316 = *(*int32)(unsafe.Add(mBase, uint32(v287)+96))
															v317 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v316)+56)))
															*(*uint16)(unsafe.Add(mBase, uint32(v312)+16)) = uint16(v317)
															v319 = *(*int32)(unsafe.Add(mBase, uint32(v285)+12))
															*(*int32)(unsafe.Add(mBase, uint32(v312)+28)) = v319
															v321 = *(*int64)(unsafe.Add(mBase, uint32(v285)+4))
															*(*int64)(unsafe.Add(mBase, uint32(v312)+20)) = v321
															v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289)+7)))
															if v323&int32(16) != 0 {
																v326 = int32(0)
																v328 = *(*int32)(unsafe.Add(mBase, uint32(v287)+96))
																v329 = *(*int32)(unsafe.Add(mBase, uint32(v328)+72))
																if v329 < v326 {
																	v351 = v326
																	v354 = v351
																} else {
																	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328+int32(0))+76)))
																	if v334 != int32(1) {
																		v351 = v326
																		v354 = v351
																	} else {
																		v338 = v328 + int32(76)
																		v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+43)))
																		if v339 == int32(0) {
																			if v285 == int32(0) {
																				v351 = v326
																				v354 = v351
																			} else {
																				v344 = int32(0)
																				*(*int32)(unsafe.Add(mBase, uint32(v285))) = v344
																				v354 = v344
																			}
																		} else {
																			if v285 != 0 {
																				v347 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v338)+48)))
																				*(*int32)(unsafe.Add(mBase, uint32(v285))) = v347
																			} else {
																			}
																			v349 = *(*int32)(unsafe.Add(mBase, uint32(v338)+44))
																			v351 = v349
																			v354 = v351
																		}
																	}
																}
																v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																v356 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
																v359 = F_ReorderBufferAllocTupleBuf(m, v355, v356-int32(5))
																mBase = m.M
																v360 = m.ExcPending
																if v360 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v312)+40)) = v359
																	v362 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
																	v363 = int32(0)
																	*(*int32)(unsafe.Add(mBase, uint32(v359)+12)) = v363
																	*(*uint16)(unsafe.Add(mBase, uint32(v359)+8)) = uint16(v363)
																	*(*int32)(unsafe.Add(mBase, uint32(v359)+4)) = int32(-1)
																	*(*int32)(unsafe.Add(mBase, uint32(v359))) = v362 + int32(18)
																	v372 = *(*int32)(unsafe.Add(mBase, uint32(v354)))
																	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+4)))
																	v374 = *(*int32)(unsafe.Add(mBase, uint32(v359)+16))
																	v375 = int64(0)
																	*(*int64)(unsafe.Add(mBase, uint32(v374)+8)) = v375
																	*(*int64)(unsafe.Add(mBase, uint32(v374)+15)) = v375
																	*(*int64)(unsafe.Add(mBase, uint32(v374))) = v375
																	v382 = v362 - int32(5)
																	if v382 != 0 {
																		v383 = *(*int32)(unsafe.Add(mBase, uint32(v359)+16))
																		base.MemoryCopy(m, v383+int32(23), v354+int32(5), v382)
																	} else {
																	}
																	*(*uint8)(unsafe.Add(mBase, uint32(v374)+22)) = uint8(v373)
																	*(*int32)(unsafe.Add(mBase, uint32(v374)+18)) = v372
																	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289)+7)))
																	v398 = v391
																	if v398&int32(12) != 0 {
																		v401 = *(*int32)(unsafe.Add(mBase, uint32(v287)+96))
																		v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)+64))
																		v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																		v404 = *(*int32)(unsafe.Add(mBase, uint32(v401)+68))
																		v406 = v404 - int32(19)
																		v407 = F_ReorderBufferAllocTupleBuf(m, v403, v406)
																		mBase = m.M
																		v408 = m.ExcPending
																		if v408 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v312)+36)) = v407
																			v410 = int32(0)
																			*(*int32)(unsafe.Add(mBase, uint32(v407)+12)) = v410
																			*(*uint16)(unsafe.Add(mBase, uint32(v407)+8)) = uint16(v410)
																			*(*int32)(unsafe.Add(mBase, uint32(v407)+4)) = int32(-1)
																			*(*int32)(unsafe.Add(mBase, uint32(v407))) = v404 + int32(4)
																			v419 = *(*int32)(unsafe.Add(mBase, uint32(v402)+14))
																			v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402)+18)))
																			v421 = *(*int32)(unsafe.Add(mBase, uint32(v407)+16))
																			v422 = int64(0)
																			*(*int64)(unsafe.Add(mBase, uint32(v421)+8)) = v422
																			*(*int64)(unsafe.Add(mBase, uint32(v421)+15)) = v422
																			*(*int64)(unsafe.Add(mBase, uint32(v421))) = v422
																			if v406 != 0 {
																				v428 = *(*int32)(unsafe.Add(mBase, uint32(v407)+16))
																				base.MemoryCopy(m, v428+int32(23), v402+int32(19), v406)
																			} else {
																			}
																			*(*uint8)(unsafe.Add(mBase, uint32(v421)+22)) = uint8(v420)
																			*(*int32)(unsafe.Add(mBase, uint32(v421)+18)) = v419
																			v442 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v312)+32)) = uint8(v442)
																			v444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																			v445 = *(*int32)(unsafe.Add(mBase, uint32(v287)+96))
																			v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)+36))
																			v447 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
																			F_ReorderBufferQueueChange(m, v444, v446, v447, v312, int32(0))
																			mBase = m.M
																			v450 = m.ExcPending
																			if v450 != 0 {
																				return
																			} else {
																				m.G0 = v285 + int32(16)
																				return
																			}
																		}
																	} else {
																		v442 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(v312)+32)) = uint8(v442)
																		v444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																		v445 = *(*int32)(unsafe.Add(mBase, uint32(v287)+96))
																		v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)+36))
																		v447 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
																		F_ReorderBufferQueueChange(m, v444, v446, v447, v312, int32(0))
																		mBase = m.M
																		v450 = m.ExcPending
																		if v450 != 0 {
																			return
																		} else {
																			m.G0 = v285 + int32(16)
																			return
																		}
																	}
																}
															} else {
																v398 = v323
																if v398&int32(12) != 0 {
																	v401 = *(*int32)(unsafe.Add(mBase, uint32(v287)+96))
																	v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)+64))
																	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																	v404 = *(*int32)(unsafe.Add(mBase, uint32(v401)+68))
																	v406 = v404 - int32(19)
																	v407 = F_ReorderBufferAllocTupleBuf(m, v403, v406)
																	mBase = m.M
																	v408 = m.ExcPending
																	if v408 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v312)+36)) = v407
																		v410 = int32(0)
																		*(*int32)(unsafe.Add(mBase, uint32(v407)+12)) = v410
																		*(*uint16)(unsafe.Add(mBase, uint32(v407)+8)) = uint16(v410)
																		*(*int32)(unsafe.Add(mBase, uint32(v407)+4)) = int32(-1)
																		*(*int32)(unsafe.Add(mBase, uint32(v407))) = v404 + int32(4)
																		v419 = *(*int32)(unsafe.Add(mBase, uint32(v402)+14))
																		v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402)+18)))
																		v421 = *(*int32)(unsafe.Add(mBase, uint32(v407)+16))
																		v422 = int64(0)
																		*(*int64)(unsafe.Add(mBase, uint32(v421)+8)) = v422
																		*(*int64)(unsafe.Add(mBase, uint32(v421)+15)) = v422
																		*(*int64)(unsafe.Add(mBase, uint32(v421))) = v422
																		if v406 != 0 {
																			v428 = *(*int32)(unsafe.Add(mBase, uint32(v407)+16))
																			base.MemoryCopy(m, v428+int32(23), v402+int32(19), v406)
																		} else {
																		}
																		*(*uint8)(unsafe.Add(mBase, uint32(v421)+22)) = uint8(v420)
																		*(*int32)(unsafe.Add(mBase, uint32(v421)+18)) = v419
																		v442 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(v312)+32)) = uint8(v442)
																		v444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																		v445 = *(*int32)(unsafe.Add(mBase, uint32(v287)+96))
																		v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)+36))
																		v447 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
																		F_ReorderBufferQueueChange(m, v444, v446, v447, v312, int32(0))
																		mBase = m.M
																		v450 = m.ExcPending
																		if v450 != 0 {
																			return
																		} else {
																			m.G0 = v285 + int32(16)
																			return
																		}
																	}
																} else {
																	v442 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(v312)+32)) = uint8(v442)
																	v444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																	v445 = *(*int32)(unsafe.Add(mBase, uint32(v287)+96))
																	v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)+36))
																	v447 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
																	F_ReorderBufferQueueChange(m, v444, v446, v447, v312, int32(0))
																	mBase = m.M
																	v450 = m.ExcPending
																	if v450 != 0 {
																		return
																	} else {
																		m.G0 = v285 + int32(16)
																		return
																	}
																}
															}
														}
													}
												}
											} else {
												v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v312 = F_ReorderBufferAllocChange(m, v311)
												mBase = m.M
												v313 = m.ExcPending
												if v313 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v312)+8)) = int32(1)
													v316 = *(*int32)(unsafe.Add(mBase, uint32(v287)+96))
													v317 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v316)+56)))
													*(*uint16)(unsafe.Add(mBase, uint32(v312)+16)) = uint16(v317)
													v319 = *(*int32)(unsafe.Add(mBase, uint32(v285)+12))
													*(*int32)(unsafe.Add(mBase, uint32(v312)+28)) = v319
													v321 = *(*int64)(unsafe.Add(mBase, uint32(v285)+4))
													*(*int64)(unsafe.Add(mBase, uint32(v312)+20)) = v321
													v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289)+7)))
													if v323&int32(16) != 0 {
														v326 = int32(0)
														v328 = *(*int32)(unsafe.Add(mBase, uint32(v287)+96))
														v329 = *(*int32)(unsafe.Add(mBase, uint32(v328)+72))
														if v329 < v326 {
															v351 = v326
															v354 = v351
														} else {
															v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328+int32(0))+76)))
															if v334 != int32(1) {
																v351 = v326
																v354 = v351
															} else {
																v338 = v328 + int32(76)
																v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+43)))
																if v339 == int32(0) {
																	if v285 == int32(0) {
																		v351 = v326
																		v354 = v351
																	} else {
																		v344 = int32(0)
																		*(*int32)(unsafe.Add(mBase, uint32(v285))) = v344
																		v354 = v344
																	}
																} else {
																	if v285 != 0 {
																		v347 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v338)+48)))
																		*(*int32)(unsafe.Add(mBase, uint32(v285))) = v347
																	} else {
																	}
																	v349 = *(*int32)(unsafe.Add(mBase, uint32(v338)+44))
																	v351 = v349
																	v354 = v351
																}
															}
														}
														v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v356 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
														v359 = F_ReorderBufferAllocTupleBuf(m, v355, v356-int32(5))
														mBase = m.M
														v360 = m.ExcPending
														if v360 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v312)+40)) = v359
															v362 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
															v363 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v359)+12)) = v363
															*(*uint16)(unsafe.Add(mBase, uint32(v359)+8)) = uint16(v363)
															*(*int32)(unsafe.Add(mBase, uint32(v359)+4)) = int32(-1)
															*(*int32)(unsafe.Add(mBase, uint32(v359))) = v362 + int32(18)
															v372 = *(*int32)(unsafe.Add(mBase, uint32(v354)))
															v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+4)))
															v374 = *(*int32)(unsafe.Add(mBase, uint32(v359)+16))
															v375 = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(v374)+8)) = v375
															*(*int64)(unsafe.Add(mBase, uint32(v374)+15)) = v375
															*(*int64)(unsafe.Add(mBase, uint32(v374))) = v375
															v382 = v362 - int32(5)
															if v382 != 0 {
																v383 = *(*int32)(unsafe.Add(mBase, uint32(v359)+16))
																base.MemoryCopy(m, v383+int32(23), v354+int32(5), v382)
															} else {
															}
															*(*uint8)(unsafe.Add(mBase, uint32(v374)+22)) = uint8(v373)
															*(*int32)(unsafe.Add(mBase, uint32(v374)+18)) = v372
															v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289)+7)))
															v398 = v391
															if v398&int32(12) != 0 {
																v401 = *(*int32)(unsafe.Add(mBase, uint32(v287)+96))
																v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)+64))
																v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																v404 = *(*int32)(unsafe.Add(mBase, uint32(v401)+68))
																v406 = v404 - int32(19)
																v407 = F_ReorderBufferAllocTupleBuf(m, v403, v406)
																mBase = m.M
																v408 = m.ExcPending
																if v408 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v312)+36)) = v407
																	v410 = int32(0)
																	*(*int32)(unsafe.Add(mBase, uint32(v407)+12)) = v410
																	*(*uint16)(unsafe.Add(mBase, uint32(v407)+8)) = uint16(v410)
																	*(*int32)(unsafe.Add(mBase, uint32(v407)+4)) = int32(-1)
																	*(*int32)(unsafe.Add(mBase, uint32(v407))) = v404 + int32(4)
																	v419 = *(*int32)(unsafe.Add(mBase, uint32(v402)+14))
																	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402)+18)))
																	v421 = *(*int32)(unsafe.Add(mBase, uint32(v407)+16))
																	v422 = int64(0)
																	*(*int64)(unsafe.Add(mBase, uint32(v421)+8)) = v422
																	*(*int64)(unsafe.Add(mBase, uint32(v421)+15)) = v422
																	*(*int64)(unsafe.Add(mBase, uint32(v421))) = v422
																	if v406 != 0 {
																		v428 = *(*int32)(unsafe.Add(mBase, uint32(v407)+16))
																		base.MemoryCopy(m, v428+int32(23), v402+int32(19), v406)
																	} else {
																	}
																	*(*uint8)(unsafe.Add(mBase, uint32(v421)+22)) = uint8(v420)
																	*(*int32)(unsafe.Add(mBase, uint32(v421)+18)) = v419
																	v442 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(v312)+32)) = uint8(v442)
																	v444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																	v445 = *(*int32)(unsafe.Add(mBase, uint32(v287)+96))
																	v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)+36))
																	v447 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
																	F_ReorderBufferQueueChange(m, v444, v446, v447, v312, int32(0))
																	mBase = m.M
																	v450 = m.ExcPending
																	if v450 != 0 {
																		return
																	} else {
																		m.G0 = v285 + int32(16)
																		return
																	}
																}
															} else {
																v442 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v312)+32)) = uint8(v442)
																v444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																v445 = *(*int32)(unsafe.Add(mBase, uint32(v287)+96))
																v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)+36))
																v447 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
																F_ReorderBufferQueueChange(m, v444, v446, v447, v312, int32(0))
																mBase = m.M
																v450 = m.ExcPending
																if v450 != 0 {
																	return
																} else {
																	m.G0 = v285 + int32(16)
																	return
																}
															}
														}
													} else {
														v398 = v323
														if v398&int32(12) != 0 {
															v401 = *(*int32)(unsafe.Add(mBase, uint32(v287)+96))
															v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)+64))
															v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															v404 = *(*int32)(unsafe.Add(mBase, uint32(v401)+68))
															v406 = v404 - int32(19)
															v407 = F_ReorderBufferAllocTupleBuf(m, v403, v406)
															mBase = m.M
															v408 = m.ExcPending
															if v408 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v312)+36)) = v407
																v410 = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(v407)+12)) = v410
																*(*uint16)(unsafe.Add(mBase, uint32(v407)+8)) = uint16(v410)
																*(*int32)(unsafe.Add(mBase, uint32(v407)+4)) = int32(-1)
																*(*int32)(unsafe.Add(mBase, uint32(v407))) = v404 + int32(4)
																v419 = *(*int32)(unsafe.Add(mBase, uint32(v402)+14))
																v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402)+18)))
																v421 = *(*int32)(unsafe.Add(mBase, uint32(v407)+16))
																v422 = int64(0)
																*(*int64)(unsafe.Add(mBase, uint32(v421)+8)) = v422
																*(*int64)(unsafe.Add(mBase, uint32(v421)+15)) = v422
																*(*int64)(unsafe.Add(mBase, uint32(v421))) = v422
																if v406 != 0 {
																	v428 = *(*int32)(unsafe.Add(mBase, uint32(v407)+16))
																	base.MemoryCopy(m, v428+int32(23), v402+int32(19), v406)
																} else {
																}
																*(*uint8)(unsafe.Add(mBase, uint32(v421)+22)) = uint8(v420)
																*(*int32)(unsafe.Add(mBase, uint32(v421)+18)) = v419
																v442 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v312)+32)) = uint8(v442)
																v444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																v445 = *(*int32)(unsafe.Add(mBase, uint32(v287)+96))
																v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)+36))
																v447 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
																F_ReorderBufferQueueChange(m, v444, v446, v447, v312, int32(0))
																mBase = m.M
																v450 = m.ExcPending
																if v450 != 0 {
																	return
																} else {
																	m.G0 = v285 + int32(16)
																	return
																}
															}
														} else {
															v442 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v312)+32)) = uint8(v442)
															v444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															v445 = *(*int32)(unsafe.Add(mBase, uint32(v287)+96))
															v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)+36))
															v447 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
															F_ReorderBufferQueueChange(m, v444, v446, v447, v312, int32(0))
															mBase = m.M
															v450 = m.ExcPending
															if v450 != 0 {
																return
															} else {
																m.G0 = v285 + int32(16)
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
			case 2:
				v625 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				v626 = F_SnapBuildProcessChange(m, v16, v18, v625)
				mBase = m.M
				v627 = m.ExcPending
				if v627 != 0 {
					return
				} else {
					if v626 == int32(0) {
						return
					} else {
						v630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
						if v630 != 0 {
							return
						} else {
							v631 = int32(0)
							v633 = m.G0
							v635 = v633 - int32(16)
							m.G0 = v635
							v638 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[0]))
							if v638 == v631 {
								v681 = v631
							} else {
								v641 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
								v642 = int32(0)
								v648 = F_XLogRecGetBlockTagExtended(m, v641, v642, v635+int32(4), v642, v642, v642)
								mBase = m.M
								if v648 == v642 {
									v681 = v631
								} else {
									v651 = *(*int32)(unsafe.Add(mBase, uint32(v635)+12))
									v653 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[0]))
									if v651 != v653 {
										v664 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[1]))
										if base.B2i32(v664 == int32(0))|base.B2i32(v651 != v664) != 0 {
											v681 = int32(1)
										} else {
											v670 = *(*int32)(unsafe.Add(mBase, uint32(v635)+8))
											v672 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[2]))
											if v670 != v672 {
												v681 = int32(1)
											} else {
												v674 = *(*int32)(unsafe.Add(mBase, uint32(v635)+4))
												v676 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[3]))
												if v674 == v676 {
													v681 = int32(0)
												} else {
													v681 = int32(1)
												}
											}
										}
									} else {
										v655 = *(*int32)(unsafe.Add(mBase, uint32(v635)+8))
										v657 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[4]))
										if v655 != v657 {
											v664 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[1]))
											if base.B2i32(v664 == int32(0))|base.B2i32(v651 != v664) != 0 {
												v681 = int32(1)
											} else {
												v670 = *(*int32)(unsafe.Add(mBase, uint32(v635)+8))
												v672 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[2]))
												if v670 != v672 {
													v681 = int32(1)
												} else {
													v674 = *(*int32)(unsafe.Add(mBase, uint32(v635)+4))
													v676 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[3]))
													if v674 == v676 {
														v681 = int32(0)
													} else {
														v681 = int32(1)
													}
												}
											}
										} else {
											v659 = *(*int32)(unsafe.Add(mBase, uint32(v635)+4))
											v661 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[5]))
											if v659 == v661 {
												v681 = v631
											} else {
												v664 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[1]))
												if base.B2i32(v664 == int32(0))|base.B2i32(v651 != v664) != 0 {
													v681 = int32(1)
												} else {
													v670 = *(*int32)(unsafe.Add(mBase, uint32(v635)+8))
													v672 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[2]))
													if v670 != v672 {
														v681 = int32(1)
													} else {
														v674 = *(*int32)(unsafe.Add(mBase, uint32(v635)+4))
														v676 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[3]))
														if v674 == v676 {
															v681 = int32(0)
														} else {
															v681 = int32(1)
														}
													}
												}
											}
										}
									}
								}
							}
							m.G0 = v635 + int32(16)
							if v681 != 0 {
								return
							} else {
								v685 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
								v686 = *(*int32)(unsafe.Add(mBase, uint32(v685)+96))
								v687 = *(*int32)(unsafe.Add(mBase, uint32(v686)+64))
								v688 = *(*int32)(unsafe.Add(mBase, uint32(v687)))
								v689 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v690 = *(*int32)(unsafe.Add(mBase, uint32(v689)+88))
								if v688 != v690 {
									return
								} else {
									v692 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									if v692 != 0 {
										v693 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v686)+56)))
										v694 = F_filter_by_origin_cb_wrapper(m, l0, v693)
										mBase = m.M
										v695 = m.ExcPending
										if v695 != 0 {
											return
										} else {
											if v694 != 0 {
												return
											} else {
												v696 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v697 = F_ReorderBufferAllocChange(m, v696)
												mBase = m.M
												v698 = m.ExcPending
												if v698 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v697)+8)) = int32(11)
													v701 = *(*int32)(unsafe.Add(mBase, uint32(v685)+96))
													v702 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v701)+56)))
													*(*uint16)(unsafe.Add(mBase, uint32(v697)+16)) = uint16(v702)
													v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v687)+8)))
													if v704&int32(1) != 0 {
														v707 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v697)+24)) = uint8(v707)
														v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v687)+8)))
														v710 = v709
													} else {
														v710 = v704
													}
													if v710&int32(2) != 0 {
														v713 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v697)+25)) = uint8(v713)
													} else {
													}
													v715 = *(*int32)(unsafe.Add(mBase, uint32(v687)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v697)+20)) = v715
													v717 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v718 = *(*int32)(unsafe.Add(mBase, uint32(v717)+120))
													v721 = F_MemoryContextAlloc(m, v718, v715<<(uint(int32(2))%32))
													mBase = m.M
													v722 = m.ExcPending
													if v722 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v697)+28)) = v721
														v724 = *(*int32)(unsafe.Add(mBase, uint32(v687)+4))
														v726 = v724 << (uint(int32(2)) % 32)
														if v726 != 0 {
															base.MemoryCopy(m, v721, v687+int32(12), v726)
														} else {
														}
														v730 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v731 = *(*int32)(unsafe.Add(mBase, uint32(v685)+96))
														v732 = *(*int32)(unsafe.Add(mBase, uint32(v731)+36))
														v733 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
														F_ReorderBufferQueueChange(m, v730, v732, v733, v697, int32(0))
														mBase = m.M
														v736 = m.ExcPending
														if v736 != 0 {
															return
														} else {
															return
														}
													}
												}
											}
										}
									} else {
										v696 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v697 = F_ReorderBufferAllocChange(m, v696)
										mBase = m.M
										v698 = m.ExcPending
										if v698 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v697)+8)) = int32(11)
											v701 = *(*int32)(unsafe.Add(mBase, uint32(v685)+96))
											v702 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v701)+56)))
											*(*uint16)(unsafe.Add(mBase, uint32(v697)+16)) = uint16(v702)
											v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v687)+8)))
											if v704&int32(1) != 0 {
												v707 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v697)+24)) = uint8(v707)
												v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v687)+8)))
												v710 = v709
											} else {
												v710 = v704
											}
											if v710&int32(2) != 0 {
												v713 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v697)+25)) = uint8(v713)
											} else {
											}
											v715 = *(*int32)(unsafe.Add(mBase, uint32(v687)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v697)+20)) = v715
											v717 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v718 = *(*int32)(unsafe.Add(mBase, uint32(v717)+120))
											v721 = F_MemoryContextAlloc(m, v718, v715<<(uint(int32(2))%32))
											mBase = m.M
											v722 = m.ExcPending
											if v722 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v697)+28)) = v721
												v724 = *(*int32)(unsafe.Add(mBase, uint32(v687)+4))
												v726 = v724 << (uint(int32(2)) % 32)
												if v726 != 0 {
													base.MemoryCopy(m, v721, v687+int32(12), v726)
												} else {
												}
												v730 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v731 = *(*int32)(unsafe.Add(mBase, uint32(v685)+96))
												v732 = *(*int32)(unsafe.Add(mBase, uint32(v731)+36))
												v733 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
												F_ReorderBufferQueueChange(m, v730, v732, v733, v697, int32(0))
												mBase = m.M
												v736 = m.ExcPending
												if v736 != 0 {
													return
												} else {
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
			case 4:
				v740 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				v741 = F_SnapBuildProcessChange(m, v16, v18, v740)
				mBase = m.M
				v742 = m.ExcPending
				if v742 != 0 {
					return
				} else {
					if v741 == int32(0) {
						return
					} else {
						v745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
						if v745 != 0 {
							return
						} else {
							v746 = int32(0)
							v748 = m.G0
							v750 = v748 - int32(16)
							m.G0 = v750
							v753 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[0]))
							if v753 == v746 {
								v796 = v746
							} else {
								v756 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
								v757 = int32(0)
								v763 = F_XLogRecGetBlockTagExtended(m, v756, v757, v750+int32(4), v757, v757, v757)
								mBase = m.M
								if v763 == v757 {
									v796 = v746
								} else {
									v766 = *(*int32)(unsafe.Add(mBase, uint32(v750)+12))
									v768 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[0]))
									if v766 != v768 {
										v779 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[1]))
										if base.B2i32(v779 == int32(0))|base.B2i32(v766 != v779) != 0 {
											v796 = int32(1)
										} else {
											v785 = *(*int32)(unsafe.Add(mBase, uint32(v750)+8))
											v787 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[2]))
											if v785 != v787 {
												v796 = int32(1)
											} else {
												v789 = *(*int32)(unsafe.Add(mBase, uint32(v750)+4))
												v791 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[3]))
												if v789 == v791 {
													v796 = int32(0)
												} else {
													v796 = int32(1)
												}
											}
										}
									} else {
										v770 = *(*int32)(unsafe.Add(mBase, uint32(v750)+8))
										v772 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[4]))
										if v770 != v772 {
											v779 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[1]))
											if base.B2i32(v779 == int32(0))|base.B2i32(v766 != v779) != 0 {
												v796 = int32(1)
											} else {
												v785 = *(*int32)(unsafe.Add(mBase, uint32(v750)+8))
												v787 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[2]))
												if v785 != v787 {
													v796 = int32(1)
												} else {
													v789 = *(*int32)(unsafe.Add(mBase, uint32(v750)+4))
													v791 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[3]))
													if v789 == v791 {
														v796 = int32(0)
													} else {
														v796 = int32(1)
													}
												}
											}
										} else {
											v774 = *(*int32)(unsafe.Add(mBase, uint32(v750)+4))
											v776 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[5]))
											if v774 == v776 {
												v796 = v746
											} else {
												v779 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[1]))
												if base.B2i32(v779 == int32(0))|base.B2i32(v766 != v779) != 0 {
													v796 = int32(1)
												} else {
													v785 = *(*int32)(unsafe.Add(mBase, uint32(v750)+8))
													v787 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[2]))
													if v785 != v787 {
														v796 = int32(1)
													} else {
														v789 = *(*int32)(unsafe.Add(mBase, uint32(v750)+4))
														v791 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[3]))
														if v789 == v791 {
															v796 = int32(0)
														} else {
															v796 = int32(1)
														}
													}
												}
											}
										}
									}
								}
							}
							m.G0 = v750 + int32(16)
							if v796 != 0 {
								return
							} else {
								v800 = m.G0
								v802 = v800 - int32(16)
								m.G0 = v802
								v804 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
								v805 = int32(0)
								F_XLogRecGetBlockTag(m, v804, v805, v802+int32(4), v805, v805)
								mBase = m.M
								v811 = m.ExcPending
								if v811 != 0 {
									return
								} else {
									v812 = *(*int32)(unsafe.Add(mBase, uint32(v802)+8))
									v813 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v814 = *(*int32)(unsafe.Add(mBase, uint32(v813)+88))
									if v812 != v814 {
										m.G0 = v802 + int32(16)
										return
									} else {
										v816 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										if v816 != 0 {
											v817 = *(*int32)(unsafe.Add(mBase, uint32(v804)+96))
											v818 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v817)+56)))
											v819 = F_filter_by_origin_cb_wrapper(m, l0, v818)
											mBase = m.M
											v820 = m.ExcPending
											if v820 != 0 {
												return
											} else {
												if v819 != 0 {
													m.G0 = v802 + int32(16)
													return
												} else {
													v821 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v822 = F_ReorderBufferAllocChange(m, v821)
													mBase = m.M
													v823 = m.ExcPending
													if v823 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v822)+8)) = int32(9)
														v826 = *(*int32)(unsafe.Add(mBase, uint32(v804)+96))
														v827 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v826)+56)))
														*(*uint16)(unsafe.Add(mBase, uint32(v822)+16)) = uint16(v827)
														v829 = *(*int64)(unsafe.Add(mBase, uint32(v802)+4))
														*(*int64)(unsafe.Add(mBase, uint32(v822)+20)) = v829
														v831 = *(*int32)(unsafe.Add(mBase, uint32(v802)+12))
														*(*int32)(unsafe.Add(mBase, uint32(v822)+28)) = v831
														v833 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v822)+32)) = uint8(v833)
														v835 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v836 = *(*int32)(unsafe.Add(mBase, uint32(v804)+96))
														v837 = *(*int32)(unsafe.Add(mBase, uint32(v836)+36))
														v838 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
														F_ReorderBufferQueueChange(m, v835, v837, v838, v822, int32(0))
														mBase = m.M
														v841 = m.ExcPending
														if v841 != 0 {
															return
														} else {
															m.G0 = v802 + int32(16)
															return
														}
													}
												}
											}
										} else {
											v821 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v822 = F_ReorderBufferAllocChange(m, v821)
											mBase = m.M
											v823 = m.ExcPending
											if v823 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v822)+8)) = int32(9)
												v826 = *(*int32)(unsafe.Add(mBase, uint32(v804)+96))
												v827 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v826)+56)))
												*(*uint16)(unsafe.Add(mBase, uint32(v822)+16)) = uint16(v827)
												v829 = *(*int64)(unsafe.Add(mBase, uint32(v802)+4))
												*(*int64)(unsafe.Add(mBase, uint32(v822)+20)) = v829
												v831 = *(*int32)(unsafe.Add(mBase, uint32(v802)+12))
												*(*int32)(unsafe.Add(mBase, uint32(v822)+28)) = v831
												v833 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v822)+32)) = uint8(v833)
												v835 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v836 = *(*int32)(unsafe.Add(mBase, uint32(v804)+96))
												v837 = *(*int32)(unsafe.Add(mBase, uint32(v836)+36))
												v838 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
												F_ReorderBufferQueueChange(m, v835, v837, v838, v822, int32(0))
												mBase = m.M
												v841 = m.ExcPending
												if v841 != 0 {
													return
												} else {
													m.G0 = v802 + int32(16)
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
			case 5, 6:
				return
			default:
				v31 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				v32 = F_SnapBuildProcessChange(m, v16, v18, v31)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					if v32 == int32(0) {
						return
					} else {
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
						if v36 != 0 {
							return
						} else {
							v37 = int32(0)
							v39 = m.G0
							v41 = v39 - int32(16)
							m.G0 = v41
							v44 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[0]))
							if v44 == v37 {
								v87 = v37
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
								v48 = int32(0)
								v54 = F_XLogRecGetBlockTagExtended(m, v47, v48, v41+int32(4), v48, v48, v48)
								mBase = m.M
								if v54 == v48 {
									v87 = v37
								} else {
									v57 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
									v59 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[0]))
									if v57 != v59 {
										v70 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[1]))
										if base.B2i32(v70 == int32(0))|base.B2i32(v57 != v70) != 0 {
											v87 = int32(1)
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
											v78 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[2]))
											if v76 != v78 {
												v87 = int32(1)
											} else {
												v80 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
												v82 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[3]))
												if v80 == v82 {
													v87 = int32(0)
												} else {
													v87 = int32(1)
												}
											}
										}
									} else {
										v61 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
										v63 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[4]))
										if v61 != v63 {
											v70 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[1]))
											if base.B2i32(v70 == int32(0))|base.B2i32(v57 != v70) != 0 {
												v87 = int32(1)
											} else {
												v76 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
												v78 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[2]))
												if v76 != v78 {
													v87 = int32(1)
												} else {
													v80 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
													v82 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[3]))
													if v80 == v82 {
														v87 = int32(0)
													} else {
														v87 = int32(1)
													}
												}
											}
										} else {
											v65 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
											v67 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[5]))
											if v65 == v67 {
												v87 = v37
											} else {
												v70 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[1]))
												if base.B2i32(v70 == int32(0))|base.B2i32(v57 != v70) != 0 {
													v87 = int32(1)
												} else {
													v76 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
													v78 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[2]))
													if v76 != v78 {
														v87 = int32(1)
													} else {
														v80 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
														v82 = *(*int32)(unsafe.Add(mBase, _c_F_heap_decode[3]))
														if v80 == v82 {
															v87 = int32(0)
														} else {
															v87 = int32(1)
														}
													}
												}
											}
										}
									}
								}
							}
							m.G0 = v41 + int32(16)
							if v87 != 0 {
								return
							} else {
								v91 = m.G0
								v93 = v91 - int32(16)
								m.G0 = v93
								v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+96))
								v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+64))
								v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+2)))
								if v98&int32(8) == int32(0) {
									m.G0 = v93 + int32(16)
									return
								} else {
									v103 = int32(0)
									F_XLogRecGetBlockTag(m, v95, v103, v93, v103, v103)
									mBase = m.M
									v107 = m.ExcPending
									if v107 != 0 {
										return
									} else {
										v108 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
										v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+88))
										if v108 != v110 {
											m.G0 = v93 + int32(16)
											return
										} else {
											v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											if v112 != 0 {
												v113 = *(*int32)(unsafe.Add(mBase, uint32(v95)+96))
												v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v113)+56)))
												v115 = F_filter_by_origin_cb_wrapper(m, l0, v114)
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return
												} else {
													if v115 != 0 {
														m.G0 = v93 + int32(16)
														return
													} else {
														v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v118 = F_ReorderBufferAllocChange(m, v117)
														mBase = m.M
														v119 = m.ExcPending
														if v119 != 0 {
															return
														} else {
															v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+2)))
															*(*int32)(unsafe.Add(mBase, uint32(v118)+8)) = v120 << (uint(int32(1)) % 32) & int32(8)
															v126 = *(*int32)(unsafe.Add(mBase, uint32(v95)+96))
															v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126)+56)))
															*(*uint16)(unsafe.Add(mBase, uint32(v118)+16)) = uint16(v127)
															v129 = *(*int64)(unsafe.Add(mBase, uint32(v93)))
															*(*int64)(unsafe.Add(mBase, uint32(v118)+20)) = v129
															v131 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v118)+28)) = v131
															v133 = int32(0)
															v135 = v93 + int32(12)
															v137 = *(*int32)(unsafe.Add(mBase, uint32(v95)+96))
															v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+72))
															if v138 < v133 {
																v160 = v133
																v163 = v160
															} else {
																v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137+int32(0))+76)))
																if v143 != int32(1) {
																	v160 = v133
																	v163 = v160
																} else {
																	v147 = v137 + int32(76)
																	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+43)))
																	if v148 == int32(0) {
																		if v135 == int32(0) {
																			v160 = v133
																			v163 = v160
																		} else {
																			v153 = int32(0)
																			*(*int32)(unsafe.Add(mBase, uint32(v135))) = v153
																			v163 = v153
																		}
																	} else {
																		if v135 != 0 {
																			v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+48)))
																			*(*int32)(unsafe.Add(mBase, uint32(v135))) = v156
																		} else {
																		}
																		v158 = *(*int32)(unsafe.Add(mBase, uint32(v147)+44))
																		v160 = v158
																		v163 = v160
																	}
																}
															}
															v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															v165 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
															v168 = F_ReorderBufferAllocTupleBuf(m, v164, v165-int32(5))
															mBase = m.M
															v169 = m.ExcPending
															if v169 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v118)+40)) = v168
																v171 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
																v172 = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(v168)+12)) = v172
																*(*uint16)(unsafe.Add(mBase, uint32(v168)+8)) = uint16(v172)
																*(*int32)(unsafe.Add(mBase, uint32(v168)+4)) = int32(-1)
																*(*int32)(unsafe.Add(mBase, uint32(v168))) = v171 + int32(18)
																v181 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
																v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+4)))
																v183 = *(*int32)(unsafe.Add(mBase, uint32(v168)+16))
																v184 = int64(0)
																*(*int64)(unsafe.Add(mBase, uint32(v183)+8)) = v184
																*(*int64)(unsafe.Add(mBase, uint32(v183)+15)) = v184
																*(*int64)(unsafe.Add(mBase, uint32(v183))) = v184
																v191 = v171 - int32(5)
																if v191 != 0 {
																	v192 = *(*int32)(unsafe.Add(mBase, uint32(v168)+16))
																	base.MemoryCopy(m, v192+int32(23), v163+int32(5), v191)
																} else {
																}
																*(*uint8)(unsafe.Add(mBase, uint32(v183)+22)) = uint8(v182)
																*(*int32)(unsafe.Add(mBase, uint32(v183)+18)) = v181
																v200 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v118)+32)) = uint8(v200)
																v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																v203 = *(*int32)(unsafe.Add(mBase, uint32(v95)+96))
																v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+36))
																v205 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
																v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+2)))
																F_ReorderBufferQueueChange(m, v202, v204, v205, v118, int32(base.Ui32(v206&int32(16))>>(uint(int32(4))%32)))
																mBase = m.M
																v212 = m.ExcPending
																if v212 != 0 {
																	return
																} else {
																	m.G0 = v93 + int32(16)
																	return
																}
															}
														}
													}
												}
											} else {
												v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												v118 = F_ReorderBufferAllocChange(m, v117)
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
													return
												} else {
													v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+2)))
													*(*int32)(unsafe.Add(mBase, uint32(v118)+8)) = v120 << (uint(int32(1)) % 32) & int32(8)
													v126 = *(*int32)(unsafe.Add(mBase, uint32(v95)+96))
													v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126)+56)))
													*(*uint16)(unsafe.Add(mBase, uint32(v118)+16)) = uint16(v127)
													v129 = *(*int64)(unsafe.Add(mBase, uint32(v93)))
													*(*int64)(unsafe.Add(mBase, uint32(v118)+20)) = v129
													v131 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v118)+28)) = v131
													v133 = int32(0)
													v135 = v93 + int32(12)
													v137 = *(*int32)(unsafe.Add(mBase, uint32(v95)+96))
													v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+72))
													if v138 < v133 {
														v160 = v133
														v163 = v160
													} else {
														v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137+int32(0))+76)))
														if v143 != int32(1) {
															v160 = v133
															v163 = v160
														} else {
															v147 = v137 + int32(76)
															v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+43)))
															if v148 == int32(0) {
																if v135 == int32(0) {
																	v160 = v133
																	v163 = v160
																} else {
																	v153 = int32(0)
																	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v153
																	v163 = v153
																}
															} else {
																if v135 != 0 {
																	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+48)))
																	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v156
																} else {
																}
																v158 = *(*int32)(unsafe.Add(mBase, uint32(v147)+44))
																v160 = v158
																v163 = v160
															}
														}
													}
													v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													v165 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
													v168 = F_ReorderBufferAllocTupleBuf(m, v164, v165-int32(5))
													mBase = m.M
													v169 = m.ExcPending
													if v169 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v118)+40)) = v168
														v171 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
														v172 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v168)+12)) = v172
														*(*uint16)(unsafe.Add(mBase, uint32(v168)+8)) = uint16(v172)
														*(*int32)(unsafe.Add(mBase, uint32(v168)+4)) = int32(-1)
														*(*int32)(unsafe.Add(mBase, uint32(v168))) = v171 + int32(18)
														v181 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
														v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+4)))
														v183 = *(*int32)(unsafe.Add(mBase, uint32(v168)+16))
														v184 = int64(0)
														*(*int64)(unsafe.Add(mBase, uint32(v183)+8)) = v184
														*(*int64)(unsafe.Add(mBase, uint32(v183)+15)) = v184
														*(*int64)(unsafe.Add(mBase, uint32(v183))) = v184
														v191 = v171 - int32(5)
														if v191 != 0 {
															v192 = *(*int32)(unsafe.Add(mBase, uint32(v168)+16))
															base.MemoryCopy(m, v192+int32(23), v163+int32(5), v191)
														} else {
														}
														*(*uint8)(unsafe.Add(mBase, uint32(v183)+22)) = uint8(v182)
														*(*int32)(unsafe.Add(mBase, uint32(v183)+18)) = v181
														v200 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v118)+32)) = uint8(v200)
														v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v203 = *(*int32)(unsafe.Add(mBase, uint32(v95)+96))
														v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+36))
														v205 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
														v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+2)))
														F_ReorderBufferQueueChange(m, v202, v204, v205, v118, int32(base.Ui32(v206&int32(16))>>(uint(int32(4))%32)))
														mBase = m.M
														v212 = m.ExcPending
														if v212 != 0 {
															return
														} else {
															m.G0 = v93 + int32(16)
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
func F_heap_fill_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l3
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = int32(128)
	goto L3
L2:
	;
	v17 = int32(0)
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v17
	if l5 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = l5 - int32(1)
	goto L6
L5:
	;
	v22 = int32(0)
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v22
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4))))
	v26 = v24 & int32(_a_F_heap_fill_tuple_0)
	*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v26)
	if int32(0) < v14 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v38 = int32(0)
	goto L10
L8:
	;
	goto L9
L9:
	;
	m.G0 = v11 + int32(16)
	return
L10:
	;
	v42 = v38 << (uint(int32(3)) % 32)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v47 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L9
L12:
	;
	v48 = v11 + int32(8)
	goto L14
L13:
	;
	v48 = int32(0)
	goto L14
L14:
	;
	if l1 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v54 = *(*int64)(unsafe.Add(mBase, uint32(l1+v42)))
	v56 = v54
	goto L17
L16:
	;
	v56 = int64(0)
	goto L17
L17:
	;
	if l2 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v38))))
	v60 = v58
	goto L20
L19:
	;
	v60 = int32(1)
	goto L20
L20:
	;
	F_fill_val(m, l0+int32(28)+v42, v48, v11+int32(4), v11+int32(12), l4, v56, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return
L22:
	;
	v64 = v38 + int32(1)
	if v64 != v14 {
		v38 = v64
		goto L10
	} else {
		goto L23
	}
L23:
	;
	goto L11
}
func F_heap_getattr_6(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64 {
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
	var v18 int64
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v83 int64
	_ = v83
	var v84 int32
	_ = v84
	var v90 int64
	_ = v90
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+18)))
	if base.Ui32(v14&int32(2047)) < base.Ui32(l1) {
		v18 = F_getmissingattr(m, l2, l1, l3)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			v90 = v18
			m.G0 = v11 + int32(16)
			return v90
		}
	} else {
		v22 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v22)
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+20)))
		if v25&int32(1) == v22 {
			v34 = l2 + l1<<(uint(int32(3))%32) + int32(20)
			v35 = int32(*(*int16)(unsafe.Add(mBase, uint32(v34))))
			if v35 < int32(0) {
				v83 = F_nocachegetattr(m, l0, l1, l2)
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int64(0)
				} else {
					v90 = v83
					m.G0 = v11 + int32(16)
					return v90
				}
			} else {
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+22)))
				v40 = v24 + v38 + v35
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+4)))
				if v41 == int32(1) {
					v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(v34)+2)))
					if base.I32_popcnt(v44) != int32(1) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v44
							F_errmsg_internal(m, int32(_a_F_heap_getattr_6_0), v11)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_heap_getattr_6_1), int32(123), int32(_a_F_heap_getattr_6_2))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						switch base.I32_ctz(v44) {
						case 0:
							v49 = int64(*(*int8)(unsafe.Add(mBase, uint32(v40))))
							v90 = v49
							m.G0 = v11 + int32(16)
							return v90
						case 1:
							v50 = int64(*(*int16)(unsafe.Add(mBase, uint32(v40))))
							v90 = v50
							m.G0 = v11 + int32(16)
							return v90
						case 2:
							v51 = int64(*(*int32)(unsafe.Add(mBase, uint32(v40))))
							v90 = v51
							m.G0 = v11 + int32(16)
							return v90
						case 3:
							v52 = *(*int64)(unsafe.Add(mBase, uint32(v40)))
							v90 = v52
							m.G0 = v11 + int32(16)
							return v90
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = v44
								F_errmsg_internal(m, int32(_a_F_heap_getattr_6_0), v11)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_heap_getattr_6_1), int32(123), int32(_a_F_heap_getattr_6_2))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
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
				} else {
					v90 = base.I64_extend_i32_u(v40)
					m.G0 = v11 + int32(16)
					return v90
				}
			}
		} else {
			v67 = int32(1)
			v68 = l1 - v67
			v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(base.Ui32(v68)>>(uint(int32(3))%32)))+23)))
			if int32(base.Ui32(v72)>>(uint(v68&int32(7))%32))&v67 != 0 {
				v83 = F_nocachegetattr(m, l0, l1, l2)
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int64(0)
				} else {
					v90 = v83
					m.G0 = v11 + int32(16)
					return v90
				}
			} else {
				v78 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v78)
				v90 = int64(0)
				m.G0 = v11 + int32(16)
				return v90
			}
		}
	}
}
func F_heap_index_delete_tuples(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v56 int32
	_ = v56
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	var v105 int64
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v167 int64
	_ = v167
	var v173 int32
	_ = v173
	var v202 int32
	_ = v202
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v412 int32
	_ = v412
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v486 int32
	_ = v486
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v541 int32
	_ = v541
	var v553 int64
	_ = v553
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int64
	_ = v562
	var v565 int64
	_ = v565
	var v566 int64
	_ = v566
	var v567 int64
	_ = v567
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v659 int32
	_ = v659
	var v664 int32
	_ = v664
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
	var v708 int32
	_ = v708
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v736 int32
	_ = v736
	var v741 int32
	_ = v741
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v772 int32
	_ = v772
	var v777 int32
	_ = v777
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v900 int32
	_ = v900
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v913 int32
	_ = v913
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v945 int32
	_ = v945
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v979 int32
	_ = v979
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1016 int32
	_ = v1016
	var v1018 int32
	_ = v1018
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1032 int32
	_ = v1032
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1085 int32
	_ = v1085
	var v1091 int32
	_ = v1091
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1114 int32
	_ = v1114
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1124 int32
	_ = v1124
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1154 int32
	_ = v1154
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1237 int32
	_ = v1237
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1267 int32
	_ = v1267
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1281 int32
	_ = v1281
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1290 int32
	_ = v1290
	var v1293 int32
	_ = v1293
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1325 int32
	_ = v1325
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1393 int32
	_ = v1393
	var v1397 int32
	_ = v1397
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1419 int32
	_ = v1419
	var v1424 int32
	_ = v1424
	var v1428 int32
	_ = v1428
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1452 int32
	_ = v1452
	var v1457 int32
	_ = v1457
	var v1461 int32
	_ = v1461
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1485 int32
	_ = v1485
	var v1490 int32
	_ = v1490
	var v1509 int32
	_ = v1509
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1514 int32
	_ = v1514
	var v1520 int32
	_ = v1520
	var v1536 int32
	_ = v1536
	var v1543 int32
	_ = v1543
	var v1549 int32
	_ = v1549
	var v1556 int32
	_ = v1556
	var v1564 int32
	_ = v1564
	var v1571 int32
	_ = v1571
	var v1578 int32
	_ = v1578
	v3 = int32(0)
	v28 = m.G0
	v30 = v28 - int32(192)
	m.G0 = v30
	*(*int32)(unsafe.Add(mBase, uint32(v30)+188)) = v3
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+112)) = int32(6)
	v37 = F_GlobalVisHorizonKindForRel(m, l0)
	mBase = m.M
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37<<(uint(int32(2))%32))+uint32(_c_F_heap_index_delete_tuples[0])))
	goto L1
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+152)) = v40
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v56 = v3
	goto L2
L2:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v56<<(uint(int32(2))%32))+uint32(_c_F_heap_index_delete_tuples[1])))
	if v73 < v42 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v236 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L4:
	;
	v83 = v73
	goto L7
L5:
	;
	goto L6
L6:
	;
	v232 = v56 + int32(1)
	if v232 != int32(9) {
		v56 = v232
		goto L2
	} else {
		goto L25
	}
L7:
	;
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v43+v83<<(uint(int32(3))%32))))
	if v83 < v73 {
		v173 = v83
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v43+v173<<(uint(int32(3))%32)))) = v105
	v202 = v83 + int32(1)
	if v202 != v42 {
		v83 = v202
		goto L7
	} else {
		goto L24
	}
L10:
	;
	v109 = base.I32_rotl(base.I32_wrap_i64(v105), int32(16))
	v114 = base.I32_wrap_i64(int64(base.Ui64(v105)>>(uint(int64(32))%64))) & int32(_a_F_heap_index_delete_tuples_0)
	v120 = v83
	goto L11
L11:
	;
	v142 = v120 - v73
	v145 = v43 + v142<<(uint(int32(3))%32)
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145))))
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145)+2)))
	v150 = v146<<(uint(int32(16))%32) | v149
	if v109 != v150 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v173 = v142
	goto L9
L13:
	;
	if v161 < int32(0) {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	if base.Ui32(v150) < base.Ui32(v109) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145)+4)))
	v161 = base.B2i32(base.Ui32(v114) < base.Ui32(v156)) - base.B2i32(base.Ui32(v156) < base.Ui32(v114))
	goto L13
L17:
	;
	v155 = int32(-1)
	goto L19
L18:
	;
	v155 = int32(1)
	goto L19
L19:
	;
	v161 = v155
	goto L13
L20:
	;
	v173 = v120
	goto L9
L21:
	;
	goto L22
L22:
	;
	v167 = *(*int64)(unsafe.Add(mBase, uint32(v145)))
	*(*int64)(unsafe.Add(mBase, uint32(v43+v120<<(uint(int32(3))%32)))) = v167
	if v73 <= v142 {
		v120 = v142
		goto L11
	} else {
		goto L23
	}
L23:
	;
	goto L12
L24:
	;
	goto L8
L25:
	;
	goto L3
L26:
	;
	v240 = F_palloc_mul(m, int32(6), v235)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v772 = v235
	v777 = v3
	goto L28
L28:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v792 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L109
L29:
	;
	return int32(0)
L30:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v244 <= int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	F_pg_qsort(m, v240, v486, int32(6), int32(143))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L29
	} else {
		goto L72
	}
L32:
	;
	v247 = int32(0)
	v486 = v247
	v509 = v247
	goto L31
L33:
	;
	goto L34
L34:
	;
	v249 = int32(0)
	v254 = v249
	v256 = v249
	v257 = int32(-1)
	goto L35
L35:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v283 = v280 + v254<<(uint(int32(3))%32)
	v284 = int32(*(*int16)(unsafe.Add(mBase, uint32(v283)+6)))
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279+v284*int32(6))+3)))
	v289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v283)+2)))
	v290 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v283))))
	v293 = v289 | v290<<(uint(int32(16))%32)
	if v293 != v257 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v331 = int32(0)
	if v313 <= v331 {
		v486 = v313
		v509 = v331
		goto L31
	} else {
		goto L45
	}
L37:
	;
	if v288&int32(1) != 0 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v297 = v240 + v256*int32(6)
	*(*uint16)(unsafe.Add(mBase, uint32(v297)+4)) = uint16(v254)
	*(*int32)(unsafe.Add(mBase, uint32(v297))) = int32(_a_F_heap_index_delete_tuples_1)
	v313 = v256 + int32(1)
	v314 = v293
	goto L37
L39:
	;
	goto L40
L40:
	;
	v307 = v240 + v256*int32(6) - int32(4)
	v308 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v307))))
	v310 = v308 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v307))) = uint16(v310)
	v313 = v256
	v314 = v257
	goto L37
L41:
	;
	v317 = int32(6)
	v321 = v240 + v313*v317 - v317
	v322 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v321))))
	v324 = v322 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v321))) = uint16(v324)
	goto L43
L42:
	;
	goto L43
L43:
	;
	v328 = v254 + int32(1)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v328 < v329 {
		v254 = v328
		v256 = v313
		v257 = v314
		goto L35
	} else {
		goto L44
	}
L44:
	;
	goto L36
L45:
	;
	v334 = int32(0)
	if v313 != int32(1) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v486 = v313
	v509 = int32(1)
	goto L31
L47:
	;
	v345 = v334
	v350 = int32(0)
	goto L50
L48:
	;
	v412 = v334
	goto L49
L49:
	;
	v439 = v240 + v412*int32(6)
	v440 = int32(*(*int16)(unsafe.Add(mBase, uint32(v439))))
	if int32(5) <= v440 {
		goto L66
	} else {
		goto L67
	}
L50:
	;
	v372 = v240 + v345*int32(6)
	v373 = int32(*(*int16)(unsafe.Add(mBase, uint32(v372))))
	if int32(5) <= v373 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v313&int32(1) == int32(0) {
		goto L46
	} else {
		goto L65
	}
L52:
	;
	v376 = int32(1)
	if v373&(v373-v376) != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v385 = int32(4)
	goto L54
L54:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v372))) = uint16(v385)
	v387 = int32(*(*int16)(unsafe.Add(mBase, uint32(v372)+6)))
	if int32(5) <= v387 {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v384 = v376 << (uint(int32(32)-base.I32_clz(v373)) % 32)
	goto L57
L56:
	;
	v384 = v373
	goto L57
L57:
	;
	v385 = v384
	goto L54
L58:
	;
	v390 = int32(1)
	if v387&(v387-v390) != 0 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v400 = int32(4)
	goto L60
L60:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v372)+6)) = uint16(v400)
	v402 = int32(2)
	v403 = v345 + v402
	v405 = v350 + v402
	if v405 != v313&int32(2147483646) {
		v345 = v403
		v350 = v405
		goto L50
	} else {
		goto L64
	}
L61:
	;
	v398 = v390 << (uint(int32(32)-base.I32_clz(v387)) % 32)
	goto L63
L62:
	;
	v398 = v387
	goto L63
L63:
	;
	v400 = v398
	goto L60
L64:
	;
	goto L51
L65:
	;
	v412 = v403
	goto L49
L66:
	;
	v443 = int32(1)
	if v440&(v440-v443) != 0 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v452 = int32(4)
	goto L68
L68:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v439))) = uint16(v452)
	goto L46
L69:
	;
	v451 = v443 << (uint(int32(32)-base.I32_clz(v440)) % 32)
	goto L71
L70:
	;
	v451 = v440
	goto L71
L71:
	;
	v452 = v451
	goto L68
L72:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v517 = F_palloc(m, v514<<(uint(int32(3))%32))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L29
	} else {
		goto L73
	}
L73:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v509 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v756 = v736 << (uint(int32(3)) % 32)
	if v756 != 0 {
		goto L103
	} else {
		goto L104
	}
L75:
	;
	v731 = v519
	v736 = int32(0)
	v741 = v3
	goto L74
L76:
	;
	goto L77
L77:
	;
	v523 = int32(6)
	if v523 <= v486 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v526 = v523
	goto L80
L79:
	;
	v526 = v486
	goto L80
L80:
	;
	v541 = v3
	v553 = int64(-1)
	goto L82
L81:
	;
	v582 = int32(0)
	if v486 != int32(1) {
		goto L87
	} else {
		goto L88
	}
L82:
	;
	v558 = int32(*(*int16)(unsafe.Add(mBase, uint32(v240+v541*int32(6))+4)))
	v561 = v519 + v558<<(uint(int32(3))%32)
	v562 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v561))))
	v565 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v561)+2)))
	v566 = v562<<(uint(int64(16))%64) | v565
	v567 = int64(3)
	if (base.B2i32(v566 < v553-v567)|base.B2i32(v553+v567 < v566))&base.B2i32(v553 != int64(-1)) != 0 {
		v580 = v541
		goto L81
	} else {
		goto L84
	}
L83:
	;
	v580 = v526
	goto L81
L84:
	;
	v578 = v541 + int32(1)
	if v578 != v526 {
		v541 = v578
		v553 = v566
		goto L82
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v731 = v727
	v736 = v708
	v741 = v580
	goto L74
L87:
	;
	v594 = v582
	v596 = int32(0)
	v599 = v582
	goto L90
L88:
	;
	v659 = v582
	v664 = v582
	goto L89
L89:
	;
	v685 = v240 + v659*int32(6)
	v686 = int32(*(*int16)(unsafe.Add(mBase, uint32(v685)+2)))
	v688 = v686 << (uint(int32(3)) % 32)
	if v688 != 0 {
		goto L100
	} else {
		goto L101
	}
L90:
	;
	v620 = v240 + v594*int32(6)
	v621 = int32(*(*int16)(unsafe.Add(mBase, uint32(v620)+2)))
	v623 = v621 << (uint(int32(3)) % 32)
	if v623 != 0 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	if v526&int32(1) == int32(0) {
		v708 = v648
		goto L86
	} else {
		goto L99
	}
L92:
	;
	v624 = int32(3)
	v627 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v628 = int32(*(*int16)(unsafe.Add(mBase, uint32(v620)+4)))
	base.MemoryCopy(m, v517+v599<<(uint(v624)%32), v627+v628<<(uint(v624)%32), v623)
	goto L94
L93:
	;
	goto L94
L94:
	;
	v633 = int32(*(*int16)(unsafe.Add(mBase, uint32(v620)+2)))
	v634 = v599 + v633
	v635 = int32(*(*int16)(unsafe.Add(mBase, uint32(v620)+8)))
	v637 = v635 << (uint(int32(3)) % 32)
	if v637 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v638 = int32(3)
	v641 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v642 = int32(*(*int16)(unsafe.Add(mBase, uint32(v620)+10)))
	base.MemoryCopy(m, v517+v634<<(uint(v638)%32), v641+v642<<(uint(v638)%32), v637)
	goto L97
L96:
	;
	goto L97
L97:
	;
	v647 = int32(*(*int16)(unsafe.Add(mBase, uint32(v620)+8)))
	v648 = v634 + v647
	v649 = int32(2)
	v650 = v594 + v649
	v652 = v596 + v649
	if v652 != v526&int32(-2) {
		v594 = v650
		v596 = v652
		v599 = v648
		goto L90
	} else {
		goto L98
	}
L98:
	;
	goto L91
L99:
	;
	v659 = v650
	v664 = v648
	goto L89
L100:
	;
	v689 = int32(3)
	v692 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v693 = int32(*(*int16)(unsafe.Add(mBase, uint32(v685)+4)))
	base.MemoryCopy(m, v517+v664<<(uint(v689)%32), v692+v693<<(uint(v689)%32), v688)
	goto L102
L101:
	;
	goto L102
L102:
	;
	v698 = int32(*(*int16)(unsafe.Add(mBase, uint32(v685)+2)))
	v708 = v664 + v698
	goto L86
L103:
	;
	base.MemoryCopy(m, v731, v517, v756)
	goto L105
L104:
	;
	goto L105
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v736
	F_pfree(m, v517)
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L29
	} else {
		goto L106
	}
L106:
	;
	F_pfree(m, v240)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L29
	} else {
		goto L107
	}
L107:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v772 = v763
	v777 = v741
	goto L28
L108:
	;
	v802 = int32(0)
	if v772 <= v802 {
		goto L115
	} else {
		goto L116
	}
L109:
	;
	if base.Ui32(v792) < base.Ui32(int32(_a_F_heap_index_delete_tuples_2)) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v796 = *(*int32)(unsafe.Add(mBase, _c_F_heap_index_delete_tuples[2]))
	v801 = v796
	goto L108
L111:
	;
	goto L112
L112:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v797)+92))
	v799 = F_get_tablespace_maintenance_io_concurrency(m, v798)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L29
	} else {
		goto L113
	}
L113:
	;
	v801 = v799
	goto L108
L114:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if int32(0) < v900 {
		goto L135
	} else {
		goto L136
	}
L115:
	;
	v876 = int32(0)
	v877 = int32(-1)
	goto L114
L116:
	;
	if v801 < v777 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v807 = v801
	goto L119
L118:
	;
	v807 = v777
	goto L119
L119:
	;
	v808 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v808 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v809 = v807
	goto L122
L121:
	;
	v809 = v801
	goto L122
L122:
	;
	if v809 <= int32(0) {
		goto L115
	} else {
		goto L123
	}
L123:
	;
	v812 = int32(0)
	v816 = int32(-1)
	v817 = v812
	v819 = v812
	goto L124
L124:
	;
	v843 = v791 + v817<<(uint(int32(3))%32)
	v844 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v843))))
	if v816 == int32(-1) {
		goto L128
	} else {
		goto L129
	}
L125:
	;
	v876 = v867
	v877 = v864
	goto L114
L126:
	;
	v867 = v817 + int32(1)
	if v772 <= v867 {
		v876 = v867
		v877 = v864
		goto L114
	} else {
		goto L133
	}
L127:
	;
	F_PrefetchBuffer(m, v30+int32(84), l0, int32(0), v856)
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L29
	} else {
		goto L132
	}
L128:
	;
	v847 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v843)+2)))
	v856 = v847 | v844<<(uint(int32(16))%32)
	goto L127
L129:
	;
	goto L130
L130:
	;
	v851 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v843)+2)))
	v854 = v851 | v844<<(uint(int32(16))%32)
	if v854 != v816 {
		v856 = v854
		goto L127
	} else {
		goto L131
	}
L131:
	;
	v864 = v816
	v865 = v819
	goto L126
L132:
	;
	v864 = v856
	v865 = v819 + int32(1)
	goto L126
L133:
	;
	if v865 < v809 {
		v816 = v864
		v817 = v867
		v819 = v865
		goto L124
	} else {
		goto L134
	}
L134:
	;
	goto L125
L135:
	;
	v907 = v876
	v908 = v877
	v913 = int32(-1)
	v917 = v777
	v918 = v3
	v920 = v3
	v921 = v34
	v922 = v3
	v923 = v3
	v924 = v3
	v925 = v3
	v926 = v3
	v927 = v3
	v928 = v3
	goto L138
L136:
	;
	v1556 = v802
	v1564 = v3
	v1571 = v3
	goto L137
L137:
	;
	F_UnlockReleaseBuffer(m, v1564)
	mBase = m.M
	v1578 = m.ExcPending
	if v1578 != 0 {
		goto L29
	} else {
		goto L244
	}
L138:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v934 = v931 + v920<<(uint(int32(3))%32)
	v935 = int32(*(*int16)(unsafe.Add(mBase, uint32(v934)+6)))
	v936 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v913 != int32(-1) {
		goto L142
	} else {
		goto L143
	}
L139:
	;
	v1549 = *(*int32)(unsafe.Add(mBase, uint32(v30)+188))
	v1556 = v1549
	v1564 = v1536
	v1571 = v1543
	goto L137
L140:
	;
	goto L139
L141:
	;
	v1144 = v936 + v935*int32(6)
	v1145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v934)+4)))
	v1147 = v1139 & int32(_a_F_heap_index_delete_tuples_0)
	if base.Ui32(v1145) <= base.Ui32(v1147) {
		goto L185
	} else {
		goto L186
	}
L142:
	;
	v939 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v934)+2)))
	v940 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v934))))
	if v939|v940<<(uint(int32(16))%32) == v913 {
		v1118 = v907
		v1119 = v908
		v1124 = v913
		v1128 = v917
		v1129 = v918
		v1132 = v921
		v1135 = v924
		v1137 = v926
		v1138 = v927
		v1139 = v928
		goto L141
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v945 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v945 == int32(1) {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	goto L144
L146:
	;
	if v922&int32(1)|base.B2i32(v923 == v926)&base.B2i32(int32(0) < v927) != 0 {
		v1536 = v918
		v1543 = v925
		goto L140
	} else {
		goto L149
	}
L147:
	;
	v963 = v917
	v964 = v921
	v965 = v926
	goto L148
L148:
	;
	if v918 != 0 {
		goto L154
	} else {
		goto L155
	}
L149:
	;
	if int32(0) < v917 {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v963 = v961
	v964 = v962
	v965 = v923
	goto L148
L151:
	;
	v961 = v917 - int32(1)
	v962 = v921
	goto L150
L152:
	;
	goto L153
L153:
	;
	v960 = base.I32_div_s(v921, int32(2))
	v961 = v917
	v962 = v960
	goto L150
L154:
	;
	F_UnlockReleaseBuffer(m, v918)
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L29
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v968 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v934)+2)))
	v969 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v934))))
	v972 = v968 | v969<<(uint(int32(16))%32)
	v973 = F_ReadBuffer(m, l0, v972)
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L29
	} else {
		goto L158
	}
L157:
	;
	goto L156
L158:
	;
	if v907 < v772 {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	F_LockBufferInternal(m, v973, int32(1))
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L29
	} else {
		goto L173
	}
L160:
	;
	v979 = v907
	goto L163
L161:
	;
	v1032 = v907
	goto L162
L162:
	;
	v1058 = v908
	v1059 = v1032
	goto L159
L163:
	;
	v1005 = v791 + v979<<(uint(int32(3))%32)
	v1006 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1005))))
	if v908 == int32(-1) {
		goto L167
	} else {
		goto L168
	}
L164:
	;
	v1032 = v1027
	goto L162
L165:
	;
	v1027 = v979 + int32(1)
	if v1027 < v772 {
		v979 = v1027
		goto L163
	} else {
		goto L172
	}
L166:
	;
	F_PrefetchBuffer(m, v30+int32(84), l0, int32(0), v1018)
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L29
	} else {
		goto L171
	}
L167:
	;
	v1009 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1005)+2)))
	v1018 = v1009 | v1006<<(uint(int32(16))%32)
	goto L166
L168:
	;
	goto L169
L169:
	;
	v1013 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1005)+2)))
	v1016 = v1013 | v1006<<(uint(int32(16))%32)
	if v1016 == v908 {
		goto L165
	} else {
		goto L170
	}
L170:
	;
	v1018 = v1016
	goto L166
L171:
	;
	v1058 = v1018
	v1059 = v979 + int32(1)
	goto L159
L172:
	;
	goto L164
L173:
	;
	if v973 < int32(0) {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	v1106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1105)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1106) {
		goto L178
	} else {
		goto L179
	}
L175:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, _c_F_heap_index_delete_tuples[3]))
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1091+(v973^int32(-1))<<(uint(int32(2))%32))))
	v1105 = v1097
	goto L174
L176:
	;
	goto L177
L177:
	;
	v1099 = *(*int32)(unsafe.Add(mBase, _c_F_heap_index_delete_tuples[4]))
	v1105 = v1099 + v973<<(uint(int32(13))%32) + int32(-8192)
	goto L174
L178:
	;
	v1114 = int32(base.Ui32(v1106+int32(_a_F_heap_index_delete_tuples_3)) >> (uint(int32(2)) % 32))
	goto L180
L179:
	;
	v1114 = int32(0)
	goto L180
L180:
	;
	v1118 = v1059
	v1119 = v1058
	v1124 = v972
	v1128 = v963
	v1129 = v973
	v1132 = v964
	v1135 = v1105
	v1137 = v965
	v1138 = v927 + int32(1)
	v1139 = v1114
	goto L141
L181:
	;
	v1520 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v1509 < v1520 {
		v907 = v1118
		v908 = v1119
		v913 = v1124
		v917 = v1128
		v918 = v1129
		v920 = v1509
		v921 = v1132
		v922 = v1511
		v923 = v1512
		v924 = v1135
		v925 = v1514
		v926 = v1137
		v927 = v1138
		v928 = v1139
		goto L138
	} else {
		goto L243
	}
L182:
	;
	v1509 = v920 + int32(1)
	v1511 = v922
	v1512 = v923
	v1514 = v925
	goto L181
L183:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1461 = m.ExcPending
	if v1461 != 0 {
		goto L29
	} else {
		goto L239
	}
L184:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L29
	} else {
		goto L235
	}
L185:
	;
	v1150 = v1135 + int32(20)
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v1150+v1145<<(uint(int32(2))%32))))
	if v1154&int32(_a_F_heap_index_delete_tuples_4) == int32(0) {
		goto L184
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1397 = m.ExcPending
	if v1397 != 0 {
		goto L29
	} else {
		goto L231
	}
L188:
	;
	if base.Ui32(int32(_a_F_heap_index_delete_tuples_5)) <= base.Ui32(v1154) {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v1164 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1135+v1154&int32(_a_F_heap_index_delete_tuples_6))+18)))
	if v1164 < int32(0) {
		goto L183
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	v1167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1144)+2)))
	if v1167 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	goto L191
L193:
	;
	v1170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v934)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+108)) = uint16(v1170)
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v934)))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+104)) = v1172
	v1182 = F_heap_hot_search_buffer(m, v30+int32(104), l0, v1129, v30+int32(112), v30+int32(84), int32(0), int32(1))
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L29
	} else {
		goto L196
	}
L194:
	;
	v1196 = v1145
	v1197 = v922
	v1198 = v923
	goto L195
L195:
	;
	if base.Ui32(v1147) <= base.Ui32((v1196-int32(1))&int32(_a_F_heap_index_delete_tuples_0)) {
		goto L201
	} else {
		goto L202
	}
L196:
	;
	if v1182 != 0 {
		goto L182
	} else {
		goto L197
	}
L197:
	;
	v1184 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1144)+2)) = uint8(v1184)
	v1186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v1186 == v1184 {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v1189 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1144)+4)))
	v1190 = v923 + v1189
	v1193 = base.B2i32(v1132 <= v1190) | v922
	v1194 = v1190
	goto L200
L199:
	;
	v1193 = v922
	v1194 = v923
	goto L200
L200:
	;
	v1195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v934)+4)))
	v1196 = v1195
	v1197 = v1193
	v1198 = v1194
	goto L195
L201:
	;
	v1393 = v920 + int32(1)
	v1509 = v1393
	v1511 = v1197
	v1512 = v1198
	v1514 = v1393
	goto L181
L202:
	;
	v1210 = v1196
	v1211 = int32(0)
	goto L203
L203:
	;
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v1150+v1210&int32(_a_F_heap_index_delete_tuples_0)<<(uint(int32(2))%32))))
	switch int32(base.Ui32(v1237)>>(uint(int32(15))%32))&int32(3) - int32(1) {
	case 0:
		goto L206
	case 1:
		goto L207
	default:
		goto L201
	}
L204:
	;
	goto L201
L205:
	;
	if base.Ui32((v1338-int32(1))&int32(_a_F_heap_index_delete_tuples_0)) < base.Ui32(v1147) {
		v1210 = v1338
		v1211 = v1339
		goto L203
	} else {
		goto L230
	}
L206:
	;
	v1248 = v1135 + v1237&int32(_a_F_heap_index_delete_tuples_6)
	if v1211 != 0 {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	v1338 = v1237 & int32(_a_F_heap_index_delete_tuples_6)
	v1339 = v1211
	goto L205
L208:
	;
	v1249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1248)+20)))
	v1250 = int32(768)
	if v1249&v1250 != v1250 {
		goto L211
	} else {
		goto L212
	}
L209:
	;
	goto L210
L210:
	;
	F_HeapTupleHeaderAdvanceConflictHorizon(m, v1248, v30+int32(188))
	mBase = m.M
	v1261 = m.ExcPending
	if v1261 != 0 {
		goto L29
	} else {
		goto L215
	}
L211:
	;
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v1248)))
	v1256 = v1254
	goto L213
L212:
	;
	v1256 = int32(2)
	goto L213
L213:
	;
	if v1256 != v1211 {
		goto L201
	} else {
		goto L214
	}
L214:
	;
	goto L210
L215:
	;
	v1262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1248)+19)))
	if v1262&int32(64) == int32(0) {
		goto L201
	} else {
		goto L216
	}
L216:
	;
	v1267 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1248)+20)))
	if v1267&int32(2048)|base.B2i32(v1267&int32(768) == int32(512)) != 0 {
		goto L201
	} else {
		goto L217
	}
L217:
	;
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v1248)+4))
	v1276 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1248)+16)))
	if v1267&int32(_a_F_heap_index_delete_tuples_7) != int32(_a_F_heap_index_delete_tuples_8) {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v1338 = v1276
	v1339 = v1275
	goto L205
L219:
	;
	goto L220
L220:
	;
	v1281 = int32(0)
	v1285 = F_GetMultiXactIdMembers(m, v1275, v30+int32(84), v1281)
	mBase = m.M
	v1286 = m.ExcPending
	if v1286 != 0 {
		goto L29
	} else {
		goto L221
	}
L221:
	;
	if v1285 <= int32(0) {
		v1338 = v1276
		v1339 = v1281
		goto L205
	} else {
		goto L222
	}
L222:
	;
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(v30)+84))
	v1293 = int32(0)
	goto L225
L223:
	;
	F_pfree(m, v1290)
	mBase = m.M
	v1332 = m.ExcPending
	if v1332 != 0 {
		goto L29
	} else {
		goto L229
	}
L224:
	;
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	v1330 = v1328
	goto L223
L225:
	;
	v1320 = v1290 + v1293<<(uint(int32(3))%32)
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(v1320)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v1321) {
		goto L224
	} else {
		goto L227
	}
L226:
	;
	v1330 = int32(0)
	goto L223
L227:
	;
	v1325 = v1293 + int32(1)
	if v1325 != v1285 {
		v1293 = v1325
		goto L225
	} else {
		goto L228
	}
L228:
	;
	goto L226
L229:
	;
	v1338 = v1276
	v1339 = v1330
	goto L205
L230:
	;
	goto L204
L231:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1400 = m.ExcPending
	if v1400 != 0 {
		goto L29
	} else {
		goto L232
	}
L232:
	;
	v1401 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v934)+2)))
	v1402 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v934))))
	v1403 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1144))))
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1406 = *(*int32)(unsafe.Add(mBase, uint32(v1405)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v1406 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+12)) = v1404
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v1403
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v1145
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v1401 | v1402<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(_a_F_heap_index_delete_tuples_9), v30)
	mBase = m.M
	v1419 = m.ExcPending
	if v1419 != 0 {
		goto L29
	} else {
		goto L233
	}
L233:
	;
	F_errfinish(m, int32(_a_F_heap_index_delete_tuples_10), int32(_a_F_heap_index_delete_tuples_11), int32(_a_F_heap_index_delete_tuples_12))
	mBase = m.M
	v1424 = m.ExcPending
	if v1424 != 0 {
		goto L29
	} else {
		goto L234
	}
L234:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L235:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L29
	} else {
		goto L236
	}
L236:
	;
	v1432 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v934)+2)))
	v1433 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v934))))
	v1434 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1144))))
	v1435 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(v1436)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+48)) = v1437 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+44)) = v1435
	*(*int32)(unsafe.Add(mBase, uint32(v30)+40)) = v1434
	*(*int32)(unsafe.Add(mBase, uint32(v30)+36)) = v1145
	*(*int32)(unsafe.Add(mBase, uint32(v30)+32)) = v1432 | v1433<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(_a_F_heap_index_delete_tuples_13), v30+int32(32))
	mBase = m.M
	v1452 = m.ExcPending
	if v1452 != 0 {
		goto L29
	} else {
		goto L237
	}
L237:
	;
	F_errfinish(m, int32(_a_F_heap_index_delete_tuples_10), int32(_a_F_heap_index_delete_tuples_14), int32(_a_F_heap_index_delete_tuples_12))
	mBase = m.M
	v1457 = m.ExcPending
	if v1457 != 0 {
		goto L29
	} else {
		goto L238
	}
L238:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L239:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v1464 = m.ExcPending
	if v1464 != 0 {
		goto L29
	} else {
		goto L240
	}
L240:
	;
	v1465 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v934)+2)))
	v1466 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v934))))
	v1467 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1144))))
	v1468 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1470 = *(*int32)(unsafe.Add(mBase, uint32(v1469)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+80)) = v1470 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+76)) = v1468
	*(*int32)(unsafe.Add(mBase, uint32(v30)+72)) = v1467
	*(*int32)(unsafe.Add(mBase, uint32(v30)+68)) = v1145
	*(*int32)(unsafe.Add(mBase, uint32(v30)+64)) = v1465 | v1466<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(_a_F_heap_index_delete_tuples_15), v30-int32(-64))
	mBase = m.M
	v1485 = m.ExcPending
	if v1485 != 0 {
		goto L29
	} else {
		goto L241
	}
L241:
	;
	F_errfinish(m, int32(_a_F_heap_index_delete_tuples_10), int32(_a_F_heap_index_delete_tuples_16), int32(_a_F_heap_index_delete_tuples_12))
	mBase = m.M
	v1490 = m.ExcPending
	if v1490 != 0 {
		goto L29
	} else {
		goto L242
	}
L242:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L243:
	;
	v1536 = v1129
	v1543 = v1514
	goto L140
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v1571
	m.G0 = v30 + int32(192)
	return v1556
}
func F_heap_lock_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
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
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v174 int32
	_ = v174
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v318 int32
	_ = v318
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v494 int32
	_ = v494
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v527 int32
	_ = v527
	var v574 int32
	_ = v574
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v653 int32
	_ = v653
	var v664 int32
	_ = v664
	var v677 int32
	_ = v677
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v743 int32
	_ = v743
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v785 int32
	_ = v785
	var v793 int32
	_ = v793
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v818 int32
	_ = v818
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v878 int32
	_ = v878
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v915 int32
	_ = v915
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v937 int32
	_ = v937
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v973 int32
	_ = v973
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1018 int32
	_ = v1018
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1054 int32
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1064 int32
	_ = v1064
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1115 int32
	_ = v1115
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1127 int32
	_ = v1127
	var v1132 int32
	_ = v1132
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1146 int32
	_ = v1146
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1201 int32
	_ = v1201
	var v1206 int32
	_ = v1206
	var v1209 int32
	_ = v1209
	var v1228 int32
	_ = v1228
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1236 int32
	_ = v1236
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1299 int32
	_ = v1299
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1311 int32
	_ = v1311
	var v1321 int32
	_ = v1321
	var v1323 int32
	_ = v1323
	var v1325 int32
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1333 int32
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1341 int32
	_ = v1341
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1354 int32
	_ = v1354
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1360 int32
	_ = v1360
	var v1362 int32
	_ = v1362
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1376 int32
	_ = v1376
	var v1378 int32
	_ = v1378
	var v1380 int32
	_ = v1380
	var v1395 int32
	_ = v1395
	var v1401 int32
	_ = v1401
	var v1403 int32
	_ = v1403
	var v1406 int32
	_ = v1406
	var v1409 int64
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1412 int64
	_ = v1412
	var v1414 int32
	_ = v1414
	var v1418 int32
	_ = v1418
	var v1424 int32
	_ = v1424
	var v1427 int32
	_ = v1427
	var v1436 int64
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1447 int32
	_ = v1447
	var v1453 int32
	_ = v1453
	var v1455 int32
	_ = v1455
	var v1472 int32
	_ = v1472
	var v1475 int32
	_ = v1475
	var v1497 int32
	_ = v1497
	var v1499 int32
	_ = v1499
	var v1502 int32
	_ = v1502
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1526 int32
	_ = v1526
	var v1531 int32
	_ = v1531
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1544 int32
	_ = v1544
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1553 int32
	_ = v1553
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1561 int32
	_ = v1561
	var v1570 int32
	_ = v1570
	var v1590 int32
	_ = v1590
	var v1591 int32
	_ = v1591
	var v1595 int32
	_ = v1595
	var v1598 int32
	_ = v1598
	var v1600 int32
	_ = v1600
	var v1602 int32
	_ = v1602
	var v1611 int32
	_ = v1611
	var v1630 int32
	_ = v1630
	var v1633 int32
	_ = v1633
	var v1635 int32
	_ = v1635
	var v1636 int32
	_ = v1636
	var v1640 int32
	_ = v1640
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1649 int32
	_ = v1649
	var v1652 int32
	_ = v1652
	var v1655 int32
	_ = v1655
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1662 int32
	_ = v1662
	var v1665 int32
	_ = v1665
	var v1675 int32
	_ = v1675
	var v1677 int32
	_ = v1677
	var v1678 int32
	_ = v1678
	var v1681 int32
	_ = v1681
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1691 int32
	_ = v1691
	var v1694 int32
	_ = v1694
	var v1704 int32
	_ = v1704
	var v1706 int32
	_ = v1706
	var v1711 int32
	_ = v1711
	var v1713 int32
	_ = v1713
	v9 = int32(0)
	v27 = m.G0
	v29 = v27 + int32(-64)
	m.G0 = v29
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v9
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v38 = F_ReadBuffer(m, l0, v33|v34<<(uint(int32(16))%32))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v38
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v47 = v43 | v44<<(uint(int32(16))%32)
	if v38 < int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+10)))
	if v66&int32(4) != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[0]))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v51+(v38^int32(-1))<<(uint(int32(2))%32))))
	v65 = v57
	goto L3
L5:
	;
	goto L6
L6:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[1]))
	v65 = v59 + v38<<(uint(int32(13))%32) + int32(-8192)
	goto L3
L7:
	;
	F_visibilitymap_pin(m, l0, v47, v27+int32(-4))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	v74 = v38
	goto L9
L9:
	;
	v76 = l1 + int32(4)
	F_LockBufferInternal(m, v74, int32(3))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v74 = v73
	goto L9
L11:
	;
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
	v85 = v65 + v80<<(uint(int32(2))%32) + int32(20)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v65 + v86&int32(_a_F_heap_lock_tuple_0)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(base.Ui32(v91) >> (uint(int32(17)) % 32))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v95
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v99 = F_HeapTupleSatisfiesUpdate(m, l1, l2, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v1704 = *(*int32)(unsafe.Add(mBase, uint32(v1687)+60))
	if v1704 != 0 {
		goto L372
	} else {
		goto L373
	}
L13:
	;
	v1675 = *(*int32)(unsafe.Add(mBase, uint32(v1655)))
	F_UnlockBuffer(m, v1675)
	mBase = m.M
	v1677 = m.ExcPending
	if v1677 != 0 {
		goto L1
	} else {
		goto L371
	}
L14:
	;
	if v99 == int32(1) {
		v1649 = l0
		v1652 = l3
		v1655 = l6
		v1657 = int32(1)
		v1658 = v29
		v1662 = v76
		v1665 = v9
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v103 = int32(1)
	v110 = l0
	v111 = l1
	v112 = l2
	v113 = l3
	v114 = l4
	v115 = l5
	v116 = l6
	v117 = l7
	v119 = v29
	v121 = v99
	v123 = v76
	v124 = v103
	v125 = v65
	v126 = v9
	v127 = l3*int32(12) + int32(_a_F_heap_lock_tuple_1)
	v128 = v47
	v129 = v9
	v131 = l5 ^ v103
	goto L16
L16:
	;
	v137 = v121 - int32(3)
	if base.Ui32(v137) <= base.Ui32(int32(2)) {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	v1541 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v1542 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1541)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v117)+4)) = uint16(v1542)
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(v1541)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v117))) = v1544
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(v1546)+4))
	v1548 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1546)+20)))
	if v1548&int32(_a_F_heap_lock_tuple_2) != int32(_a_F_heap_lock_tuple_3) {
		goto L352
	} else {
		goto L353
	}
L18:
	;
	goto L17
L19:
	;
	v1510 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	v1511 = F_HeapTupleSatisfiesUpdate(m, v111, v112, v1510)
	mBase = m.M
	v1512 = m.ExcPending
	if v1512 != 0 {
		goto L1
	} else {
		goto L349
	}
L20:
	;
	v1497 = int32(0)
	v1499 = v1472
	v1502 = v1475
	goto L19
L21:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140)+18)))
	v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140)+20)))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v119)+48)) = uint16(v144)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v140)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v119)+44)) = v146
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	F_UnlockBuffer(m, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	v1228 = v121
	v1231 = v124
	v1233 = v126
	v1236 = v129
	goto L23
L23:
	;
	if v1228 != 0 {
		v1526 = v1228
		v1531 = v1233
		goto L18
	} else {
		goto L305
	}
L24:
	;
	if v124 == int32(0) {
		v574 = v129
		goto L30
	} else {
		goto L31
	}
L25:
	;
	v1228 = v1201
	v1231 = int32(0)
	v1233 = v1206
	v1236 = v1209
	goto L23
L26:
	;
	v793 = v142 & int32(_a_F_heap_lock_tuple_3)
	if v793 != 0 {
		goto L172
	} else {
		goto L173
	}
L27:
	;
	if v142&int32(80) != int32(16) {
		v785 = v743
		goto L26
	} else {
		goto L168
	}
L28:
	;
	if v142&int32(_a_F_heap_lock_tuple_3) == int32(0) {
		v743 = v574
		goto L27
	} else {
		goto L162
	}
L29:
	;
	v677 = int32(64)
	if base.B2i32(v142&int32(128) == int32(0))&base.B2i32(v142&int32(_a_F_heap_lock_tuple_4) != v677)|base.B2i32(v653 == v677) != 0 {
		v785 = v664
		goto L26
	} else {
		goto L159
	}
L30:
	;
	switch v113 {
	case 0:
		goto L140
	case 1:
		goto L139
	case 2:
		goto L28
	default:
		v785 = v574
		goto L26
	}
L31:
	;
	if v142&int32(_a_F_heap_lock_tuple_3) != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v1678 = v110
	v1681 = v113
	v1686 = int32(0)
	v1687 = v119
	v1691 = v123
	v1694 = v126
	goto L12
L33:
	;
	F_pfree(m, v330)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L1
	} else {
		goto L138
	}
L34:
	;
	v162 = F_GetMultiXactIdMembers(m, v143, v119+int32(56), int32(base.Ui32(v142&int32(128))>>(uint(int32(7))%32)))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if base.Ui32(v143) < base.Ui32(int32(3)) {
		goto L91
	} else {
		goto L92
	}
L37:
	;
	if int32(0) < v162 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v174 = int32(0)
	v185 = v129
	goto L41
L39:
	;
	v361 = v129
	goto L40
L40:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v119)+56))
	if v368 == int32(0) {
		v574 = v361
		goto L30
	} else {
		goto L88
	}
L41:
	;
	v192 = int32(3)
	v193 = v174 << (uint(v192) % 32)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v119)+56))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v193+v194)))
	if base.Ui32(v196) < base.Ui32(v192) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v361 = v337
	goto L40
L43:
	;
	if v328 != 0 {
		goto L83
	} else {
		goto L84
	}
L44:
	;
	v328 = int32(0)
	goto L43
L45:
	;
	goto L46
L46:
	;
	v208 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[2]))
	if v208 == v196 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v328 = int32(1)
	goto L43
L48:
	;
	goto L49
L49:
	;
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[3]))
	if v212 <= int32(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v328 = v318
	goto L43
L51:
	;
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[4]))
	if v216 == int32(0) {
		v318 = int32(0)
		goto L50
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v286 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[5]))
	v288 = int32(0)
	v291 = v212 - int32(1)
	goto L73
L54:
	;
	v221 = v216
	goto L55
L55:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
	if v227 == int32(4) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v318 = int32(0)
	goto L50
L57:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v221)+80))
	if v281 != 0 {
		v221 = v281
		goto L55
	} else {
		goto L72
	}
L58:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	if v230 == int32(0) {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v233 = int32(1)
	if v196 == v230 {
		v318 = v233
		goto L50
	} else {
		goto L60
	}
L60:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v221)+52))
	v237 = v235 - int32(1)
	if v237 < int32(0) {
		goto L57
	} else {
		goto L61
	}
L61:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v221)+48))
	v243 = int32(0)
	v246 = v237
	goto L62
L62:
	;
	v251 = int32(2)
	v252 = base.I32_div_s(v246-v243, v251)
	v253 = v252 + v243
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v240+v253<<(uint(v251)%32))))
	if v257 == v196 {
		v318 = v233
		goto L50
	} else {
		goto L64
	}
L63:
	;
	goto L57
L64:
	;
	v266 = base.B2i32(v257-v196 < int32(0)) | base.B2i32(base.Ui32(v257) < base.Ui32(int32(3)))
	if v266 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v267 = v253 + int32(1)
	goto L67
L66:
	;
	v267 = v243
	goto L67
L67:
	;
	if v266 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v270 = v246
	goto L70
L69:
	;
	v270 = v253 - int32(1)
	goto L70
L70:
	;
	if v267 <= v270 {
		v243 = v267
		v246 = v270
		goto L62
	} else {
		goto L71
	}
L71:
	;
	goto L63
L72:
	;
	goto L56
L73:
	;
	v296 = int32(2)
	v297 = base.I32_div_s(v291-v288, v296)
	v298 = v297 + v288
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v286+v298<<(uint(v296)%32))))
	v303 = base.B2i32(v302 == v196)
	if v302 == v196 {
		v318 = v303
		goto L50
	} else {
		goto L75
	}
L74:
	;
	v318 = v303
	goto L50
L75:
	;
	v306 = base.B2i32(base.Ui32(v302) < base.Ui32(v196))
	if base.Ui32(v302) < base.Ui32(v196) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v307 = v298 + int32(1)
	goto L78
L77:
	;
	v307 = v288
	goto L78
L78:
	;
	if base.Ui32(v302) < base.Ui32(v196) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v310 = v291
	goto L81
L80:
	;
	v310 = v298 - int32(1)
	goto L81
L81:
	;
	if v307 <= v310 {
		v288 = v307
		v291 = v310
		goto L73
	} else {
		goto L82
	}
L82:
	;
	goto L74
L83:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v119)+56))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v193+v330)+4))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v332<<(uint(int32(2))%32))+uint32(_c_F_heap_lock_tuple[6])))
	if base.Ui32(v113) <= base.Ui32(v335) {
		goto L33
	} else {
		goto L86
	}
L84:
	;
	v337 = v185
	goto L85
L85:
	;
	v340 = v174 + int32(1)
	if v340 != v162 {
		v174 = v340
		v185 = v337
		goto L41
	} else {
		goto L87
	}
L86:
	;
	v337 = int32(1)
	goto L85
L87:
	;
	goto L42
L88:
	;
	F_pfree(m, v368)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v574 = v361
	goto L30
L90:
	;
	if v504 == int32(0) {
		v574 = v129
		goto L30
	} else {
		goto L130
	}
L91:
	;
	v504 = int32(0)
	goto L90
L92:
	;
	goto L93
L93:
	;
	v384 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[2]))
	if v384 == v143 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v504 = int32(1)
	goto L90
L95:
	;
	goto L96
L96:
	;
	v388 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[3]))
	if v388 <= int32(0) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v504 = v494
	goto L90
L98:
	;
	v392 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[4]))
	if v392 == int32(0) {
		v494 = int32(0)
		goto L97
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v462 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[5]))
	v464 = int32(0)
	v467 = v388 - int32(1)
	goto L120
L101:
	;
	v397 = v392
	goto L102
L102:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v397)+20))
	if v403 == int32(4) {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v494 = int32(0)
	goto L97
L104:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v397)+80))
	if v457 != 0 {
		v397 = v457
		goto L102
	} else {
		goto L119
	}
L105:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v397)))
	if v406 == int32(0) {
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v409 = int32(1)
	if v143 == v406 {
		v494 = v409
		goto L97
	} else {
		goto L107
	}
L107:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v397)+52))
	v413 = v411 - int32(1)
	if v413 < int32(0) {
		goto L104
	} else {
		goto L108
	}
L108:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v397)+48))
	v419 = int32(0)
	v422 = v413
	goto L109
L109:
	;
	v427 = int32(2)
	v428 = base.I32_div_s(v422-v419, v427)
	v429 = v428 + v419
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v416+v429<<(uint(v427)%32))))
	if v433 == v143 {
		v494 = v409
		goto L97
	} else {
		goto L111
	}
L110:
	;
	goto L104
L111:
	;
	v442 = base.B2i32(v433-v143 < int32(0)) | base.B2i32(base.Ui32(v433) < base.Ui32(int32(3)))
	if v442 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v443 = v429 + int32(1)
	goto L114
L113:
	;
	v443 = v419
	goto L114
L114:
	;
	if v442 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v446 = v422
	goto L117
L116:
	;
	v446 = v429 - int32(1)
	goto L117
L117:
	;
	if v443 <= v446 {
		v419 = v443
		v422 = v446
		goto L109
	} else {
		goto L118
	}
L118:
	;
	goto L110
L119:
	;
	goto L103
L120:
	;
	v472 = int32(2)
	v473 = base.I32_div_s(v467-v464, v472)
	v474 = v473 + v464
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v462+v474<<(uint(v472)%32))))
	v479 = base.B2i32(v478 == v143)
	if v478 == v143 {
		v494 = v479
		goto L97
	} else {
		goto L122
	}
L121:
	;
	v494 = v479
	goto L97
L122:
	;
	v482 = base.B2i32(base.Ui32(v478) < base.Ui32(v143))
	if base.Ui32(v478) < base.Ui32(v143) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v483 = v474 + int32(1)
	goto L125
L124:
	;
	v483 = v464
	goto L125
L125:
	;
	if base.Ui32(v478) < base.Ui32(v143) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v486 = v467
	goto L128
L127:
	;
	v486 = v474 - int32(1)
	goto L128
L128:
	;
	if v483 <= v486 {
		v464 = v483
		v467 = v486
		goto L120
	} else {
		goto L129
	}
L129:
	;
	goto L121
L130:
	;
	switch v113 {
	case 0:
		goto L32
	case 1:
		goto L133
	case 2:
		goto L131
	case 3:
		goto L132
	default:
		v574 = v129
		goto L30
	}
L131:
	;
	if v142&int32(80) == int32(64) {
		goto L32
	} else {
		goto L137
	}
L132:
	;
	if v142&int32(80) != int32(64) {
		v785 = v129
		goto L26
	} else {
		goto L135
	}
L133:
	;
	v508 = v142 & int32(80)
	v510 = v508 + int32(-64)
	if base.B2i32(v510 == int32(0))|base.B2i32(v510 == int32(16)) != 0 {
		goto L32
	} else {
		goto L134
	}
L134:
	;
	v653 = v508
	v664 = v129
	goto L29
L135:
	;
	if v141&int32(_a_F_heap_lock_tuple_5) != 0 {
		goto L32
	} else {
		goto L136
	}
L136:
	;
	v785 = v129
	goto L26
L137:
	;
	v743 = v129
	goto L27
L138:
	;
	goto L32
L139:
	;
	v653 = v142 & int32(80)
	v664 = v574
	goto L29
L140:
	;
	if v141&int32(_a_F_heap_lock_tuple_5) != 0 {
		v785 = v574
		goto L26
	} else {
		goto L141
	}
L141:
	;
	v591 = int32(base.Ui32(v142&int32(128))>>(uint(int32(7))%32)) | base.B2i32(v142&int32(_a_F_heap_lock_tuple_4) == int32(64))
	if (v131|v591)&int32(1) != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	F_LockBufferInternal(m, v626, int32(3))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L1
	} else {
		goto L155
	}
L143:
	;
	v596 = v119 + int32(44)
	v597 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123)+2)))
	v598 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123))))
	v599 = int32(16)
	v602 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v596)+2)))
	v603 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v596))))
	if v597|v598<<(uint(v599)%32) == v602|v603<<(uint(v599)%32) {
		goto L146
	} else {
		goto L147
	}
L144:
	;
	if v613 != 0 {
		goto L142
	} else {
		goto L150
	}
L145:
	;
	goto L144
L146:
	;
	v609 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123)+4)))
	v610 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v596)+4)))
	if v609 == v610 {
		v613 = int32(1)
		goto L145
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v613 = int32(0)
	goto L145
L149:
	;
	goto L148
L150:
	;
	v614 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	v617 = F_heap_lock_updated_tuple(m, v110, v142, v143, v596, v614, int32(0))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	if v617 == int32(0) {
		goto L142
	} else {
		goto L153
	}
L153:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	F_LockBufferInternal(m, v621, int32(3))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	v1526 = v617
	v1531 = v126
	goto L18
L155:
	;
	v630 = int32(0)
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v632 = F_HeapTupleHeaderIsOnlyLocked(m, v631)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	if v632 != 0 {
		v1201 = v630
		v1206 = v126
		v1209 = v574
		goto L25
	} else {
		goto L157
	}
L157:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v634)+19)))
	if (int32(base.Ui32(v635)>>(uint(int32(5))%32))|v591)&int32(1) == int32(0) {
		v1201 = v630
		v1206 = v126
		v1209 = v574
		goto L25
	} else {
		goto L158
	}
L158:
	;
	v1472 = v126
	v1475 = v574
	goto L20
L159:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	F_LockBufferInternal(m, v683, int32(3))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v688 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v687)+20)))
	v691 = int32(64)
	if base.B2i32(v688&int32(80) == v691)|base.B2i32(v688&int32(128) == int32(0))&base.B2i32(v688&int32(_a_F_heap_lock_tuple_4) != v691) != 0 {
		v1472 = v126
		v1475 = v664
		goto L20
	} else {
		goto L161
	}
L161:
	;
	v1201 = int32(0)
	v1206 = v126
	v1209 = v664
	goto L25
L162:
	;
	v710 = F_DoesMultiXactIdConflict(m, v143, v142, int32(2), int32(0))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	if v710 != 0 {
		v785 = v574
		goto L26
	} else {
		goto L164
	}
L164:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	F_LockBufferInternal(m, v712, int32(3))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v717 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v716)+20)))
	if (v717^v142)&int32(_a_F_heap_lock_tuple_6) != 0 {
		v1472 = v126
		v1475 = v574
		goto L20
	} else {
		goto L166
	}
L166:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v716)+4))
	if v721 != v143 {
		v1472 = v126
		v1475 = v574
		goto L20
	} else {
		goto L167
	}
L167:
	;
	v1201 = int32(0)
	v1206 = v126
	v1209 = v574
	goto L25
L168:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	F_LockBufferInternal(m, v754, int32(3))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v759 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v758)+20)))
	if (v759^v142)&int32(_a_F_heap_lock_tuple_6) != 0 {
		v1472 = v126
		v1475 = v743
		goto L20
	} else {
		goto L170
	}
L170:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v758)+4))
	if v763 != v143 {
		v1472 = v126
		v1475 = v743
		goto L20
	} else {
		goto L171
	}
L171:
	;
	v1201 = int32(0)
	v1206 = v126
	v1209 = v743
	goto L25
L172:
	;
	if v137 != int32(2) {
		goto L218
	} else {
		goto L219
	}
L173:
	;
	if base.Ui32(v143) < base.Ui32(int32(3)) {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	if v925 == int32(0) {
		goto L172
	} else {
		goto L214
	}
L175:
	;
	v925 = int32(0)
	goto L174
L176:
	;
	goto L177
L177:
	;
	v805 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[2]))
	if v805 == v143 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v925 = int32(1)
	goto L174
L179:
	;
	goto L180
L180:
	;
	v809 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[3]))
	if v809 <= int32(0) {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	v925 = v915
	goto L174
L182:
	;
	v813 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[4]))
	if v813 == int32(0) {
		v915 = int32(0)
		goto L181
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	v883 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[5]))
	v885 = int32(0)
	v888 = v809 - int32(1)
	goto L204
L185:
	;
	v818 = v813
	goto L186
L186:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v818)+20))
	if v824 == int32(4) {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	v915 = int32(0)
	goto L181
L188:
	;
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v818)+80))
	if v878 != 0 {
		v818 = v878
		goto L186
	} else {
		goto L203
	}
L189:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v818)))
	if v827 == int32(0) {
		goto L188
	} else {
		goto L190
	}
L190:
	;
	v830 = int32(1)
	if v143 == v827 {
		v915 = v830
		goto L181
	} else {
		goto L191
	}
L191:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v818)+52))
	v834 = v832 - int32(1)
	if v834 < int32(0) {
		goto L188
	} else {
		goto L192
	}
L192:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v818)+48))
	v840 = int32(0)
	v843 = v834
	goto L193
L193:
	;
	v848 = int32(2)
	v849 = base.I32_div_s(v843-v840, v848)
	v850 = v849 + v840
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v837+v850<<(uint(v848)%32))))
	if v854 == v143 {
		v915 = v830
		goto L181
	} else {
		goto L195
	}
L194:
	;
	goto L188
L195:
	;
	v863 = base.B2i32(v854-v143 < int32(0)) | base.B2i32(base.Ui32(v854) < base.Ui32(int32(3)))
	if v863 != 0 {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v864 = v850 + int32(1)
	goto L198
L197:
	;
	v864 = v840
	goto L198
L198:
	;
	if v863 != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v867 = v843
	goto L201
L200:
	;
	v867 = v850 - int32(1)
	goto L201
L201:
	;
	if v864 <= v867 {
		v840 = v864
		v843 = v867
		goto L193
	} else {
		goto L202
	}
L202:
	;
	goto L194
L203:
	;
	goto L187
L204:
	;
	v893 = int32(2)
	v894 = base.I32_div_s(v888-v885, v893)
	v895 = v894 + v885
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v883+v895<<(uint(v893)%32))))
	v900 = base.B2i32(v899 == v143)
	if v899 == v143 {
		v915 = v900
		goto L181
	} else {
		goto L206
	}
L205:
	;
	v915 = v900
	goto L181
L206:
	;
	v903 = base.B2i32(base.Ui32(v899) < base.Ui32(v143))
	if base.Ui32(v899) < base.Ui32(v143) {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v904 = v895 + int32(1)
	goto L209
L208:
	;
	v904 = v885
	goto L209
L209:
	;
	if base.Ui32(v899) < base.Ui32(v143) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v907 = v888
	goto L212
L211:
	;
	v907 = v895 - int32(1)
	goto L212
L212:
	;
	if v904 <= v907 {
		v885 = v904
		v888 = v907
		goto L204
	} else {
		goto L213
	}
L213:
	;
	goto L205
L214:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	F_LockBufferInternal(m, v928, int32(3))
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v933 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v932)+20)))
	if (v933^v142)&int32(_a_F_heap_lock_tuple_6) != 0 {
		v1472 = v126
		v1475 = v785
		goto L20
	} else {
		goto L216
	}
L216:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v932)+4))
	if v937 != v143 {
		v1472 = v126
		v1475 = v785
		goto L20
	} else {
		goto L217
	}
L217:
	;
	v1201 = int32(0)
	v1206 = v126
	v1209 = v785
	goto L25
L218:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	F_LockBufferInternal(m, v942, int32(3))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L1
	} else {
		goto L221
	}
L219:
	;
	goto L220
L220:
	;
	if (v126|v785)&int32(1) != 0 {
		goto L223
	} else {
		goto L224
	}
L221:
	;
	v1201 = v121
	v1206 = v126
	v1209 = v785
	goto L25
L222:
	;
	if v793 != 0 {
		goto L241
	} else {
		goto L242
	}
L223:
	;
	v993 = base.B2i32(v785 == int32(0)) | v126
	goto L222
L224:
	;
	goto L225
L225:
	;
	v952 = int32(1)
	switch v114 {
	case 0:
		goto L226
	case 1:
		goto L227
	case 2:
		goto L228
	default:
		v993 = v952
		goto L222
	}
L226:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	F_LockTuple(m, v110, v123, v990)
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L1
	} else {
		goto L238
	}
L227:
	;
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	v981 = F_ConditionalLockTuple(m, v110, v123, v979, int32(0))
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L1
	} else {
		goto L235
	}
L228:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	v955 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_lock_tuple[7])))
	v956 = F_ConditionalLockTuple(m, v110, v123, v953, v955)
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L1
	} else {
		goto L229
	}
L229:
	;
	if v956 != 0 {
		v993 = v952
		goto L222
	} else {
		goto L230
	}
L230:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	F_errcode(m, int32(50463045))
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v110)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v119)+32)) = v965 + int32(4)
	F_errmsg(m, int32(_a_F_heap_lock_tuple_7), v119+int32(32))
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	F_errfinish(m, int32(_a_F_heap_lock_tuple_8), int32(_a_F_heap_lock_tuple_9), int32(_a_F_heap_lock_tuple_10))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L235:
	;
	if v981 != 0 {
		v993 = v952
		goto L222
	} else {
		goto L236
	}
L236:
	;
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	F_LockBufferInternal(m, v983, int32(3))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	v988 = int32(0)
	v1201 = int32(6)
	v1206 = v988
	v1209 = v988
	goto L25
L238:
	;
	v993 = v952
	goto L222
L239:
	;
	if base.B2i32(v115 == int32(0))|v142&int32(128)|base.B2i32(v142&int32(_a_F_heap_lock_tuple_4) == int32(64)) != 0 {
		goto L269
	} else {
		goto L270
	}
L240:
	;
	v1070 = int32(0)
	v1074 = F_Do_MultiXactIdWait(m, v143, v994, v142, v1070, v110, v123, int32(3), v1070, v1070)
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L1
	} else {
		goto L268
	}
L241:
	;
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	switch v114 {
	case 0:
		goto L240
	case 1:
		goto L244
	case 2:
		goto L245
	default:
		goto L239
	}
L242:
	;
	goto L243
L243:
	;
	switch v114 {
	case 0:
		goto L257
	case 1:
		goto L256
	case 2:
		goto L255
	default:
		goto L239
	}
L244:
	;
	v1025 = int32(0)
	v1029 = F_Do_MultiXactIdWait(m, v143, v994, v142, int32(1), v110, v1025, v1025, v1025, v1025)
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L1
	} else {
		goto L252
	}
L245:
	;
	v996 = int32(0)
	v1000 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_lock_tuple[7])))
	v1001 = F_Do_MultiXactIdWait(m, v143, v994, v142, int32(1), v110, v996, v996, v996, v1000)
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L1
	} else {
		goto L246
	}
L246:
	;
	if v1001 != 0 {
		goto L239
	} else {
		goto L247
	}
L247:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	F_errcode(m, int32(50463045))
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v110)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v119)+16)) = v1010 + int32(4)
	F_errmsg(m, int32(_a_F_heap_lock_tuple_7), v119+int32(16))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L1
	} else {
		goto L250
	}
L250:
	;
	F_errfinish(m, int32(_a_F_heap_lock_tuple_8), int32(_a_F_heap_lock_tuple_11), int32(_a_F_heap_lock_tuple_12))
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L252:
	;
	if v1029 != 0 {
		goto L239
	} else {
		goto L253
	}
L253:
	;
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	F_LockBufferInternal(m, v1031, int32(3))
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L1
	} else {
		goto L254
	}
L254:
	;
	v1201 = int32(6)
	v1206 = v993
	v1209 = v785
	goto L25
L255:
	;
	v1048 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_lock_tuple[7])))
	v1049 = F_ConditionalXactLockTableWait(m, v143, v1048)
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L1
	} else {
		goto L262
	}
L256:
	;
	v1040 = F_ConditionalXactLockTableWait(m, v143, int32(0))
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L1
	} else {
		goto L259
	}
L257:
	;
	F_XactLockTableWait(m, v143, v110, v123, int32(3))
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	goto L239
L259:
	;
	if v1040 != 0 {
		goto L239
	} else {
		goto L260
	}
L260:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	F_LockBufferInternal(m, v1042, int32(3))
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L1
	} else {
		goto L261
	}
L261:
	;
	v1201 = int32(6)
	v1206 = v993
	v1209 = v785
	goto L25
L262:
	;
	if v1049 != 0 {
		goto L239
	} else {
		goto L263
	}
L263:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L1
	} else {
		goto L264
	}
L264:
	;
	F_errcode(m, int32(50463045))
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L1
	} else {
		goto L265
	}
L265:
	;
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v110)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = v1058 + int32(4)
	F_errmsg(m, int32(_a_F_heap_lock_tuple_7), v119)
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L1
	} else {
		goto L266
	}
L266:
	;
	F_errfinish(m, int32(_a_F_heap_lock_tuple_8), int32(_a_F_heap_lock_tuple_13), int32(_a_F_heap_lock_tuple_12))
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L1
	} else {
		goto L267
	}
L267:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L268:
	;
	goto L239
L269:
	;
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	F_LockBufferInternal(m, v1118, int32(3))
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L1
	} else {
		goto L282
	}
L270:
	;
	v1088 = v119 + int32(44)
	v1089 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123)+2)))
	v1090 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123))))
	v1091 = int32(16)
	v1094 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1088)+2)))
	v1095 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1088))))
	if v1089|v1090<<(uint(v1091)%32) == v1094|v1095<<(uint(v1091)%32) {
		goto L273
	} else {
		goto L274
	}
L271:
	;
	if v1105 != 0 {
		goto L269
	} else {
		goto L277
	}
L272:
	;
	goto L271
L273:
	;
	v1101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123)+4)))
	v1102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1088)+4)))
	if v1101 == v1102 {
		v1105 = int32(1)
		goto L272
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	v1105 = int32(0)
	goto L272
L276:
	;
	goto L275
L277:
	;
	v1106 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L1
	} else {
		goto L278
	}
L278:
	;
	v1108 = F_heap_lock_updated_tuple(m, v110, v142, v143, v1088, v1106, v113)
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L1
	} else {
		goto L279
	}
L279:
	;
	if v1108 == int32(0) {
		goto L269
	} else {
		goto L280
	}
L280:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	F_LockBufferInternal(m, v1112, int32(3))
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		goto L1
	} else {
		goto L281
	}
L281:
	;
	v1201 = v1108
	v1206 = v993
	v1209 = v785
	goto L25
L282:
	;
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v1123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1122)+20)))
	if (v1123^v142)&int32(_a_F_heap_lock_tuple_6) != 0 {
		v1472 = v993
		v1475 = v785
		goto L20
	} else {
		goto L283
	}
L283:
	;
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v1122)+4))
	if v1127 != v143 {
		v1472 = v993
		v1475 = v785
		goto L20
	} else {
		goto L284
	}
L284:
	;
	if v793|v1123&int32(3072) != 0 {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1152 = int32(0)
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v1154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1153)+20)))
	if v1154&int32(2048)|v1154&int32(128)|base.B2i32(v1154&int32(_a_F_heap_lock_tuple_4) == int32(64)) != 0 {
		v1201 = v1152
		v1206 = v993
		v1209 = v785
		goto L25
	} else {
		goto L293
	}
L286:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	if v1123&int32(128)|base.B2i32(v1123&int32(_a_F_heap_lock_tuple_4) == int32(64)) != 0 {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	F_HeapTupleSetHintBits(m, v1122, v1132, int32(2048), int32(0))
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L1
	} else {
		goto L292
	}
L288:
	;
	v1140 = F_TransactionIdDidCommit(m, v143)
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L1
	} else {
		goto L289
	}
L289:
	;
	if v1140 == int32(0) {
		goto L287
	} else {
		goto L290
	}
L290:
	;
	F_HeapTupleSetHintBits(m, v1122, v1132, int32(1024), v143)
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L1
	} else {
		goto L291
	}
L291:
	;
	goto L285
L292:
	;
	goto L285
L293:
	;
	v1165 = F_HeapTupleHeaderIsOnlyLocked(m, v1153)
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L1
	} else {
		goto L294
	}
L294:
	;
	if v1165 != 0 {
		v1201 = v1152
		v1206 = v993
		v1209 = v785
		goto L25
	} else {
		goto L295
	}
L295:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v1171 = v1169 + int32(12)
	v1172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123)+2)))
	v1173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123))))
	v1174 = int32(16)
	v1177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1171)+2)))
	v1178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1171))))
	if v1172|v1173<<(uint(v1174)%32) == v1177|v1178<<(uint(v1174)%32) {
		goto L298
	} else {
		goto L299
	}
L296:
	;
	if v1188 != 0 {
		goto L302
	} else {
		goto L303
	}
L297:
	;
	goto L296
L298:
	;
	v1184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123)+4)))
	v1185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1171)+4)))
	if v1184 == v1185 {
		v1188 = int32(1)
		goto L297
	} else {
		goto L301
	}
L299:
	;
	goto L300
L300:
	;
	v1188 = int32(0)
	goto L297
L301:
	;
	goto L300
L302:
	;
	v1189 = int32(4)
	goto L304
L303:
	;
	v1189 = int32(3)
	goto L304
L304:
	;
	v1201 = v1189
	v1206 = v993
	v1209 = v785
	goto L25
L305:
	;
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v119)+60))
	if v1243 != 0 {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v1261 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1260)+20)))
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(v1260)+4))
	F_MultiXactIdSetOldestMember(m)
	mBase = m.M
	v1264 = m.ExcPending
	if v1264 != 0 {
		goto L1
	} else {
		goto L312
	}
L307:
	;
	v1244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+10)))
	if v1244&int32(4) == int32(0) {
		goto L306
	} else {
		goto L308
	}
L308:
	;
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	F_UnlockBuffer(m, v1249)
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L1
	} else {
		goto L309
	}
L309:
	;
	F_visibilitymap_pin(m, v110, v128, v119+int32(60))
	mBase = m.M
	v1255 = m.ExcPending
	if v1255 != 0 {
		goto L1
	} else {
		goto L310
	}
L310:
	;
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	F_LockBufferInternal(m, v1256, int32(3))
	mBase = m.M
	v1259 = m.ExcPending
	if v1259 != 0 {
		goto L1
	} else {
		goto L311
	}
L311:
	;
	v1497 = v1231
	v1499 = v1233
	v1502 = v1236
	goto L19
L312:
	;
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v1266 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1265)+18)))
	v1267 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v1268 = m.ExcPending
	if v1268 != 0 {
		goto L1
	} else {
		goto L313
	}
L313:
	;
	F_compute_new_xmax_infomask(m, v1262, v1261, v1266, v1267, v113, int32(0), v119+int32(56), v119+int32(54), v119+int32(52))
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L1
	} else {
		goto L314
	}
L314:
	;
	v1278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v125)+10)))
	v1280 = v1278 & int32(4)
	if v1280 != 0 {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v119)+60))
	F_LockBufferInternal(m, v1281, int32(3))
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		goto L1
	} else {
		goto L318
	}
L316:
	;
	goto L317
L317:
	;
	v1285 = int32(_a_F_heap_lock_tuple_14)
	v1287 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[8])) = v1287 + int32(1)
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v1292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1291)+20)))
	v1294 = v1292 & int32(_a_F_heap_lock_tuple_15)
	*(*uint16)(unsafe.Add(mBase, uint32(v1291)+20)) = uint16(v1294)
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v1297 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1296)+18)))
	v1299 = v1297 & int32(_a_F_heap_lock_tuple_16)
	*(*uint16)(unsafe.Add(mBase, uint32(v1296)+18)) = uint16(v1299)
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v1302 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v119)+54)))
	v1303 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1301)+20)))
	v1304 = v1302 | v1303
	*(*uint16)(unsafe.Add(mBase, uint32(v1301)+20)) = uint16(v1304)
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v1307 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1306)+18)))
	v1308 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v119)+52)))
	v1309 = v1307 | v1308
	*(*uint16)(unsafe.Add(mBase, uint32(v1306)+18)) = uint16(v1309)
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	if v1302&int32(128)|base.B2i32(v1302&int32(_a_F_heap_lock_tuple_4) == int32(64)) == int32(0) {
		goto L320
	} else {
		goto L321
	}
L318:
	;
	goto L317
L319:
	;
	v1338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+10)))
	if v1338&int32(4) != 0 {
		goto L323
	} else {
		goto L324
	}
L320:
	;
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(v119)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v1311)+4)) = v1321
	v1335 = v1321
	goto L319
L321:
	;
	goto L322
L322:
	;
	v1323 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1311)+18)))
	v1325 = v1323 & int32(_a_F_heap_lock_tuple_17)
	*(*uint16)(unsafe.Add(mBase, uint32(v1311)+18)) = uint16(v1325)
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(v119)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v1327)+4)) = v1328
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v1331 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1330)+16)) = uint16(v1331)
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	*(*int32)(unsafe.Add(mBase, uint32(v1330)+12)) = v1333
	v1335 = v1328
	goto L319
L323:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(v119)+60))
	v1343 = F_visibilitymap_clear(m, v128, v1341, int32(2))
	mBase = m.M
	v1344 = m.ExcPending
	if v1344 != 0 {
		goto L1
	} else {
		goto L326
	}
L324:
	;
	v1345 = int32(0)
	goto L325
L325:
	;
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	F_MarkBufferDirty(m, v1346)
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L1
	} else {
		goto L327
	}
L326:
	;
	v1345 = v1343
	goto L325
L327:
	;
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(v110)+48))
	v1350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1349)+118)))
	if v1350 != int32(112) {
		goto L328
	} else {
		goto L329
	}
L328:
	;
	v1444 = int32(0)
	v1445 = int32(_a_F_heap_lock_tuple_14)
	v1447 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[8])) = v1447 - int32(1)
	if v1280 == v1444 {
		v1649 = v110
		v1652 = v113
		v1655 = v116
		v1657 = v1444
		v1658 = v119
		v1662 = v123
		v1665 = v1233
		goto L13
	} else {
		goto L347
	}
L329:
	;
	v1354 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[9]))
	if v1354 <= int32(0) {
		goto L330
	} else {
		goto L331
	}
L330:
	;
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(v110)+32))
	if v1357 != 0 {
		goto L328
	} else {
		goto L333
	}
L331:
	;
	goto L332
L332:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v1360 = m.ExcPending
	if v1360 != 0 {
		goto L1
	} else {
		goto L335
	}
L333:
	;
	v1358 = *(*int32)(unsafe.Add(mBase, uint32(v110)+40))
	if v1358 != 0 {
		goto L328
	} else {
		goto L334
	}
L334:
	;
	goto L332
L335:
	;
	v1362 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	F_XLogRegisterBuffer(m, int32(0), v1362, int32(8))
	mBase = m.M
	v1365 = m.ExcPending
	if v1365 != 0 {
		goto L1
	} else {
		goto L336
	}
L336:
	;
	v1366 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v111)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v119)+44)) = v1335
	*(*uint16)(unsafe.Add(mBase, uint32(v119)+48)) = uint16(v1366)
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v1370 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1369)+18)))
	*(*uint8)(unsafe.Add(mBase, uint32(v119)+51)) = uint8(v1345)
	v1376 = int32(1)
	v1378 = int32(8)
	v1380 = int32(4)
	v1395 = int32(base.Ui32(v1370)>>(uint(int32(9))%32))&int32(16) | (int32(base.Ui32(v1302)>>(uint(v1376)%32))&v1378 | (int32(base.Ui32(v1302)>>(uint(v1380)%32))&v1380 | (int32(base.Ui32(v1302)>>(uint(int32(12))%32))&v1376 | int32(base.Ui32(v1302)>>(uint(int32(6))%32))&int32(2))))
	*(*uint8)(unsafe.Add(mBase, uint32(v119)+50)) = uint8(v1395)
	F_XLogRegisterData(m, v119+int32(44), v1378)
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		goto L1
	} else {
		goto L337
	}
L337:
	;
	if v1345 != 0 {
		goto L338
	} else {
		goto L339
	}
L338:
	;
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v119)+60))
	F_XLogRegisterBuffer(m, int32(1), v1403, int32(0))
	mBase = m.M
	v1406 = m.ExcPending
	if v1406 != 0 {
		goto L1
	} else {
		goto L341
	}
L339:
	;
	goto L340
L340:
	;
	v1436 = F_XLogInsert(m, int32(10), int32(96))
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		goto L1
	} else {
		goto L346
	}
L341:
	;
	v1409 = F_XLogInsert(m, int32(10), int32(96))
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L1
	} else {
		goto L342
	}
L342:
	;
	v1412 = base.I64_rotl(v1409, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v125))) = v1412
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(v119)+60))
	if v1414 < int32(0) {
		goto L343
	} else {
		goto L344
	}
L343:
	;
	v1418 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[0]))
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(v1418+(v1414^int32(-1))<<(uint(int32(2))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v1424))) = v1412
	goto L328
L344:
	;
	goto L345
L345:
	;
	v1427 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v1427+v1414<<(uint(int32(13))%32))+uint32(_c_F_heap_lock_tuple[10]))) = v1412
	goto L328
L346:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v125))) = base.I64_rotl(v1436, int64(32))
	goto L328
L347:
	;
	v1453 = *(*int32)(unsafe.Add(mBase, uint32(v119)+60))
	F_UnlockBuffer(m, v1453)
	mBase = m.M
	v1455 = m.ExcPending
	if v1455 != 0 {
		goto L1
	} else {
		goto L348
	}
L348:
	;
	v1649 = v110
	v1652 = v113
	v1655 = v116
	v1657 = v1444
	v1658 = v119
	v1662 = v123
	v1665 = v1233
	goto L13
L349:
	;
	if v1511 != int32(1) {
		v121 = v1511
		v124 = v1497
		v126 = v1499
		v129 = v1502
		goto L16
	} else {
		goto L350
	}
L350:
	;
	v1649 = v110
	v1652 = v113
	v1655 = v116
	v1657 = int32(1)
	v1658 = v119
	v1662 = v123
	v1665 = v1499
	goto L13
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117)+8)) = v1611
	v1630 = int32(2)
	if v1526 == v1630 {
		goto L364
	} else {
		goto L365
	}
L352:
	;
	v1611 = v1547
	goto L351
L353:
	;
	goto L354
L354:
	;
	v1553 = int32(0)
	v1557 = F_GetMultiXactIdMembers(m, v1547, v119+int32(44), v1553)
	mBase = m.M
	v1558 = m.ExcPending
	if v1558 != 0 {
		goto L1
	} else {
		goto L355
	}
L355:
	;
	if v1557 <= int32(0) {
		v1611 = v1553
		goto L351
	} else {
		goto L356
	}
L356:
	;
	v1561 = *(*int32)(unsafe.Add(mBase, uint32(v119)+44))
	v1570 = v1553
	goto L359
L357:
	;
	F_pfree(m, v1561)
	mBase = m.M
	v1602 = m.ExcPending
	if v1602 != 0 {
		goto L1
	} else {
		goto L363
	}
L358:
	;
	v1598 = *(*int32)(unsafe.Add(mBase, uint32(v1590)))
	v1600 = v1598
	goto L357
L359:
	;
	v1590 = v1561 + v1570<<(uint(int32(3))%32)
	v1591 = *(*int32)(unsafe.Add(mBase, uint32(v1590)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v1591) {
		goto L358
	} else {
		goto L361
	}
L360:
	;
	v1600 = int32(0)
	goto L357
L361:
	;
	v1595 = v1570 + int32(1)
	if v1595 != v1557 {
		v1570 = v1595
		goto L359
	} else {
		goto L362
	}
L362:
	;
	goto L360
L363:
	;
	v1611 = v1600
	goto L351
L364:
	;
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(v1633)+8))
	v1636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1633)+20)))
	if v1636&int32(32) != 0 {
		goto L368
	} else {
		goto L369
	}
L365:
	;
	goto L366
L366:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117)+12)) = int32(-1)
	v1649 = v110
	v1652 = v113
	v1655 = v116
	v1657 = v1526
	v1658 = v119
	v1662 = v123
	v1665 = v1531
	goto L13
L367:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117)+12)) = v1645
	v1649 = v110
	v1652 = v113
	v1655 = v116
	v1657 = v1630
	v1658 = v119
	v1662 = v123
	v1665 = v1531
	goto L13
L368:
	;
	v1640 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_tuple[11]))
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(v1640+v1635<<(uint(int32(3))%32))+4))
	v1645 = v1644
	goto L370
L369:
	;
	v1645 = v1635
	goto L370
L370:
	;
	goto L367
L371:
	;
	v1678 = v1649
	v1681 = v1652
	v1686 = v1657
	v1687 = v1658
	v1691 = v1662
	v1694 = v1665
	goto L12
L372:
	;
	F_ReleaseBuffer(m, v1704)
	mBase = m.M
	v1706 = m.ExcPending
	if v1706 != 0 {
		goto L1
	} else {
		goto L375
	}
L373:
	;
	goto L374
L374:
	;
	if v1694&int32(1) != 0 {
		goto L376
	} else {
		goto L377
	}
L375:
	;
	goto L374
L376:
	;
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(v1681*int32(12))+uint32(_c_F_heap_lock_tuple[12])))
	F_UnlockTuple(m, v1678, v1691, v1711)
	mBase = m.M
	v1713 = m.ExcPending
	if v1713 != 0 {
		goto L1
	} else {
		goto L379
	}
L377:
	;
	goto L378
L378:
	;
	m.G0 = v1687 - int32(-64)
	return v1686
L379:
	;
	goto L378
}
func F_heap_page_items(m *base.Module, l0 int32) int64 {
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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int64
	_ = v56
	var v67 int64
	_ = v67
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int64
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v123 int32
	_ = v123
	var v124 int64
	_ = v124
	var v126 int64
	_ = v126
	var v128 int64
	_ = v128
	var v134 int32
	_ = v134
	var v136 int64
	_ = v136
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
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
	var v188 int32
	_ = v188
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v313 int32
	_ = v313
	var v317 int64
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
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
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int64
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int64
	_ = v381
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v411 int64
	_ = v411
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	v17 = m.G0
	v19 = v17 - int32(128)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v22 = F_pg_detoast_datum(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int64(0)
	} else {
		v26 = F_superuser(m)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int64(0)
		} else {
			if v26 != 0 {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
				if v29 == int32(0) {
					v32 = F_init_MultiFuncCall(m, l0)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int64(0)
					} else {
						v34 = int32(_a_F_heap_page_items_0)
						v35 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_items[0]))
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
						*(*int32)(unsafe.Add(mBase, _c_F_heap_page_items[0])) = v37
						v40 = F_palloc(m, int32(12))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int64(0)
						} else {
							v45 = F_get_call_result_type(m, l0, int32(0), v19+int32(16))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int64(0)
							} else {
								if v45 != int32(1) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v435 = m.ExcPending
									if v435 != 0 {
										return int64(0)
									} else {
										F_errmsg_internal(m, int32(_a_F_heap_page_items_1), int32(0))
										mBase = m.M
										v439 = m.ExcPending
										if v439 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_heap_page_items_2), int32(154), int32(_a_F_heap_page_items_3))
											mBase = m.M
											v444 = m.ExcPending
											if v444 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v49 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
									v50 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v40)+8)) = uint16(v50)
									*(*int32)(unsafe.Add(mBase, uint32(v40))) = v49
									v53 = F_get_page_from_raw(m, v22)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = v53
										v56 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v53)+12)))
										*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v40
										if base.Ui64(int64(25)) <= base.Ui64(v56) {
											v67 = int64(base.Ui64(v56+int64(262120))>>(uint(int64(2))%64)) & int64(65535)
										} else {
											v67 = int64(0)
										}
										*(*int64)(unsafe.Add(mBase, uint32(v32)+8)) = v67
										*(*int32)(unsafe.Add(mBase, _c_F_heap_page_items[0])) = v35
										v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
										v79 = *(*int64)(unsafe.Add(mBase, uint32(v78)))
										v80 = *(*int64)(unsafe.Add(mBase, uint32(v78)+8))
										if base.Ui64(v79) < base.Ui64(v80) {
											v82 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
											v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
											v84 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(v19)+6)) = v84
											*(*int64)(unsafe.Add(mBase, uint32(v19))) = v84
											v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+8)))
											v92 = *(*int32)(unsafe.Add(mBase, uint32(v83+v88<<(uint(int32(2))%32))+20))
											*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = base.I64_extend_i32_u(v88)
											v96 = int32(base.Ui32(v92) >> (uint(int32(17)) % 32))
											*(*int64)(unsafe.Add(mBase, uint32(v19)+40)) = base.I64_extend_i32_u(v96)
											v100 = v92 & int32(_a_F_heap_page_items_4)
											*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = base.I64_extend_i32_u(v100)
											*(*int64)(unsafe.Add(mBase, uint32(v19)+32)) = base.I64_extend_i32_u(int32(base.Ui32(v92)>>(uint(int32(15))%32)) & int32(3))
											if base.B2i32(v100 != (v100+int32(7))&int32(_a_F_heap_page_items_5))|base.B2i32(base.Ui32(v92) < base.Ui32(int32(_a_F_heap_page_items_6)))|base.B2i32(base.Ui32(int32(_a_F_heap_page_items_7)) < base.Ui32(v100+v96)) == int32(0) {
												v123 = v100 + v83
												v124 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v123))))
												*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = v124
												v126 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v123)+4)))
												*(*int64)(unsafe.Add(mBase, uint32(v19)+56)) = v126
												v128 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v123)+8)))
												*(*int64)(unsafe.Add(mBase, uint32(v19)+72)) = base.I64_extend_i32_u(v123 + int32(12))
												*(*int64)(unsafe.Add(mBase, uint32(v19)+64)) = v128
												v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123)+18)))
												v136 = int64(65535)
												*(*int64)(unsafe.Add(mBase, uint32(v19)+80)) = base.I64_extend_i32_u(v134) & v136
												v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123)+20)))
												*(*int64)(unsafe.Add(mBase, uint32(v19)+88)) = base.I64_extend_i32_u(v139) & v136
												v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+22)))
												*(*int64)(unsafe.Add(mBase, uint32(v19)+96)) = base.I64_extend_i32_u(v144)
												if base.B2i32(base.Ui32(v144) < base.Ui32(int32(23)))|base.B2i32(base.Ui32(v96) < base.Ui32(v144))|base.B2i32((v144+int32(7))&int32(504) != v144) == int32(0) {
													if v139&int32(1) != 0 {
														v166 = int32(base.Ui32(v134&int32(2047)+int32(7)) >> (uint(int32(3)) % 32))
														if base.Ui32(v166) <= base.Ui32(v144-int32(23)) {
															v171 = v123 + int32(23)
															v172 = int32(0)
															v174 = v166 << (uint(int32(3)) % 32)
															v177 = F_palloc(m, v174+int32(1))
															mBase = m.M
															v178 = m.ExcPending
															if v178 != 0 {
																return int64(0)
															} else {
																if v174 == int32(0) {
																} else {
																	if v174 != int32(1) {
																		v188 = v172
																		v201 = int32(0)
																		for {
																			v208 = v171 + int32(base.Ui32(v188)>>(uint(int32(3))%32))
																			v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
																			if int32(base.Ui32(v209)>>(uint(v188&int32(6))%32))&int32(1) != 0 {
																				v215 = int32(49)
																			} else {
																				v215 = int32(48)
																			}
																			*(*uint8)(unsafe.Add(mBase, uint32(v188+v177))) = uint8(v215)
																			v217 = int32(1)
																			v218 = v188 | v217
																			v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
																			if int32(base.Ui32(v222)>>(uint(v218&int32(7))%32))&v217 != 0 {
																				v228 = int32(49)
																			} else {
																				v228 = int32(48)
																			}
																			*(*uint8)(unsafe.Add(mBase, uint32(v177+v218))) = uint8(v228)
																			v230 = int32(2)
																			v231 = v188 + v230
																			v233 = v201 + v230
																			if v233 != v174&int32(4094) {
																				v188 = v231
																				v201 = v233
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		if v174&int32(1) == int32(0) {
																		} else {
																			v238 = v231
																			v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171+int32(base.Ui32(v238)>>(uint(int32(3))%32))))))
																			if int32(base.Ui32(v259)>>(uint(v238&int32(7))%32))&int32(1) != 0 {
																				v265 = int32(49)
																			} else {
																				v265 = int32(48)
																			}
																			*(*uint8)(unsafe.Add(mBase, uint32(v238+v177))) = uint8(v265)
																		}
																	} else {
																		v238 = v172
																		v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171+int32(base.Ui32(v238)>>(uint(int32(3))%32))))))
																		if int32(base.Ui32(v259)>>(uint(v238&int32(7))%32))&int32(1) != 0 {
																			v265 = int32(49)
																		} else {
																			v265 = int32(48)
																		}
																		*(*uint8)(unsafe.Add(mBase, uint32(v238+v177))) = uint8(v265)
																	}
																}
																v284 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(v174+v177))) = uint8(v284)
																v286 = F_cstring_to_text(m, v177)
																mBase = m.M
																v287 = m.ExcPending
																if v287 != 0 {
																	return int64(0)
																} else {
																	*(*int64)(unsafe.Add(mBase, uint32(v19)+104)) = base.I64_extend_i32_u(v286)
																	v290 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123)+20)))
																	v296 = v290
																	if v296&int32(8) != 0 {
																		v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+22)))
																		v317 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v123+v313-int32(4)))))
																		*(*int64)(unsafe.Add(mBase, uint32(v19)+112)) = v317
																		v322 = v313
																	} else {
																		v319 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(v19)+12)) = uint8(v319)
																		v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+22)))
																		v322 = v321
																	}
																	v323 = v96 - v322
																	v325 = v323 + int32(4)
																	v326 = F_palloc(m, v325)
																	mBase = m.M
																	v327 = m.ExcPending
																	if v327 != 0 {
																		return int64(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v326))) = v325 << (uint(int32(2)) % 32)
																		v331 = int32(0)
																		if base.B2i32(v323 == v331)|base.B2i32(v323 <= v331) == v331 {
																			v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+22)))
																			base.MemoryCopy(m, v326+int32(4), v123+v340, v323)
																		} else {
																		}
																		*(*int64)(unsafe.Add(mBase, uint32(v19)+120)) = base.I64_extend_i32_u(v326)
																		v369 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
																		v372 = F_heap_form_tuple(m, v369, v19+int32(16), v19)
																		mBase = m.M
																		v373 = m.ExcPending
																		if v373 != 0 {
																			return int64(0)
																		} else {
																			v374 = *(*int32)(unsafe.Add(mBase, uint32(v372)+16))
																			v375 = F_HeapTupleHeaderGetDatum(m, v374)
																			mBase = m.M
																			v376 = m.ExcPending
																			if v376 != 0 {
																				return int64(0)
																			} else {
																				v377 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+8)))
																				v378 = int32(1)
																				v379 = v377 + v378
																				*(*uint16)(unsafe.Add(mBase, uint32(v82)+8)) = uint16(v379)
																				v381 = *(*int64)(unsafe.Add(mBase, uint32(v78)))
																				*(*int64)(unsafe.Add(mBase, uint32(v78))) = v381 + int64(1)
																				v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																				*(*int32)(unsafe.Add(mBase, uint32(v385)+20)) = v378
																				v411 = v375
																				m.G0 = v19 + int32(128)
																				return v411
																			}
																		}
																	}
																}
															}
														} else {
															v291 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v19)+11)) = uint8(v291)
															v296 = v139
															if v296&int32(8) != 0 {
																v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+22)))
																v317 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v123+v313-int32(4)))))
																*(*int64)(unsafe.Add(mBase, uint32(v19)+112)) = v317
																v322 = v313
															} else {
																v319 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v19)+12)) = uint8(v319)
																v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+22)))
																v322 = v321
															}
															v323 = v96 - v322
															v325 = v323 + int32(4)
															v326 = F_palloc(m, v325)
															mBase = m.M
															v327 = m.ExcPending
															if v327 != 0 {
																return int64(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v326))) = v325 << (uint(int32(2)) % 32)
																v331 = int32(0)
																if base.B2i32(v323 == v331)|base.B2i32(v323 <= v331) == v331 {
																	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+22)))
																	base.MemoryCopy(m, v326+int32(4), v123+v340, v323)
																} else {
																}
																*(*int64)(unsafe.Add(mBase, uint32(v19)+120)) = base.I64_extend_i32_u(v326)
																v369 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
																v372 = F_heap_form_tuple(m, v369, v19+int32(16), v19)
																mBase = m.M
																v373 = m.ExcPending
																if v373 != 0 {
																	return int64(0)
																} else {
																	v374 = *(*int32)(unsafe.Add(mBase, uint32(v372)+16))
																	v375 = F_HeapTupleHeaderGetDatum(m, v374)
																	mBase = m.M
																	v376 = m.ExcPending
																	if v376 != 0 {
																		return int64(0)
																	} else {
																		v377 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+8)))
																		v378 = int32(1)
																		v379 = v377 + v378
																		*(*uint16)(unsafe.Add(mBase, uint32(v82)+8)) = uint16(v379)
																		v381 = *(*int64)(unsafe.Add(mBase, uint32(v78)))
																		*(*int64)(unsafe.Add(mBase, uint32(v78))) = v381 + int64(1)
																		v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																		*(*int32)(unsafe.Add(mBase, uint32(v385)+20)) = v378
																		v411 = v375
																		m.G0 = v19 + int32(128)
																		return v411
																	}
																}
															}
														}
													} else {
														v293 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v19)+11)) = uint8(v293)
														v296 = v139
														if v296&int32(8) != 0 {
															v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+22)))
															v317 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v123+v313-int32(4)))))
															*(*int64)(unsafe.Add(mBase, uint32(v19)+112)) = v317
															v322 = v313
														} else {
															v319 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v19)+12)) = uint8(v319)
															v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+22)))
															v322 = v321
														}
														v323 = v96 - v322
														v325 = v323 + int32(4)
														v326 = F_palloc(m, v325)
														mBase = m.M
														v327 = m.ExcPending
														if v327 != 0 {
															return int64(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v326))) = v325 << (uint(int32(2)) % 32)
															v331 = int32(0)
															if base.B2i32(v323 == v331)|base.B2i32(v323 <= v331) == v331 {
																v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+22)))
																base.MemoryCopy(m, v326+int32(4), v123+v340, v323)
															} else {
															}
															*(*int64)(unsafe.Add(mBase, uint32(v19)+120)) = base.I64_extend_i32_u(v326)
															v369 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
															v372 = F_heap_form_tuple(m, v369, v19+int32(16), v19)
															mBase = m.M
															v373 = m.ExcPending
															if v373 != 0 {
																return int64(0)
															} else {
																v374 = *(*int32)(unsafe.Add(mBase, uint32(v372)+16))
																v375 = F_HeapTupleHeaderGetDatum(m, v374)
																mBase = m.M
																v376 = m.ExcPending
																if v376 != 0 {
																	return int64(0)
																} else {
																	v377 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+8)))
																	v378 = int32(1)
																	v379 = v377 + v378
																	*(*uint16)(unsafe.Add(mBase, uint32(v82)+8)) = uint16(v379)
																	v381 = *(*int64)(unsafe.Add(mBase, uint32(v78)))
																	*(*int64)(unsafe.Add(mBase, uint32(v78))) = v381 + int64(1)
																	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																	*(*int32)(unsafe.Add(mBase, uint32(v385)+20)) = v378
																	v411 = v375
																	m.G0 = v19 + int32(128)
																	return v411
																}
															}
														}
													}
												} else {
													v345 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v19)+13)) = uint8(v345)
													v347 = int32(257)
													*(*uint16)(unsafe.Add(mBase, uint32(v19)+11)) = uint16(v347)
													v369 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
													v372 = F_heap_form_tuple(m, v369, v19+int32(16), v19)
													mBase = m.M
													v373 = m.ExcPending
													if v373 != 0 {
														return int64(0)
													} else {
														v374 = *(*int32)(unsafe.Add(mBase, uint32(v372)+16))
														v375 = F_HeapTupleHeaderGetDatum(m, v374)
														mBase = m.M
														v376 = m.ExcPending
														if v376 != 0 {
															return int64(0)
														} else {
															v377 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+8)))
															v378 = int32(1)
															v379 = v377 + v378
															*(*uint16)(unsafe.Add(mBase, uint32(v82)+8)) = uint16(v379)
															v381 = *(*int64)(unsafe.Add(mBase, uint32(v78)))
															*(*int64)(unsafe.Add(mBase, uint32(v78))) = v381 + int64(1)
															v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v385)+20)) = v378
															v411 = v375
															m.G0 = v19 + int32(128)
															return v411
														}
													}
												}
											} else {
												v349 = int32(257)
												*(*uint16)(unsafe.Add(mBase, uint32(v19)+12)) = uint16(v349)
												*(*int64)(unsafe.Add(mBase, uint32(v19)+4)) = int64(72340172838076673)
												v369 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
												v372 = F_heap_form_tuple(m, v369, v19+int32(16), v19)
												mBase = m.M
												v373 = m.ExcPending
												if v373 != 0 {
													return int64(0)
												} else {
													v374 = *(*int32)(unsafe.Add(mBase, uint32(v372)+16))
													v375 = F_HeapTupleHeaderGetDatum(m, v374)
													mBase = m.M
													v376 = m.ExcPending
													if v376 != 0 {
														return int64(0)
													} else {
														v377 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+8)))
														v378 = int32(1)
														v379 = v377 + v378
														*(*uint16)(unsafe.Add(mBase, uint32(v82)+8)) = uint16(v379)
														v381 = *(*int64)(unsafe.Add(mBase, uint32(v78)))
														*(*int64)(unsafe.Add(mBase, uint32(v78))) = v381 + int64(1)
														v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v385)+20)) = v378
														v411 = v375
														m.G0 = v19 + int32(128)
														return v411
													}
												}
											}
										} else {
											F_end_MultiFuncCall(m, l0)
											mBase = m.M
											v389 = m.ExcPending
											if v389 != 0 {
												return int64(0)
											} else {
												v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v390)+20)) = int32(2)
												v393 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v393)
												v411 = int64(0)
												m.G0 = v19 + int32(128)
												return v411
											}
										}
									}
								}
							}
						}
					}
				} else {
					v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
					v79 = *(*int64)(unsafe.Add(mBase, uint32(v78)))
					v80 = *(*int64)(unsafe.Add(mBase, uint32(v78)+8))
					if base.Ui64(v79) < base.Ui64(v80) {
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
						v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
						v84 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v19)+6)) = v84
						*(*int64)(unsafe.Add(mBase, uint32(v19))) = v84
						v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+8)))
						v92 = *(*int32)(unsafe.Add(mBase, uint32(v83+v88<<(uint(int32(2))%32))+20))
						*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = base.I64_extend_i32_u(v88)
						v96 = int32(base.Ui32(v92) >> (uint(int32(17)) % 32))
						*(*int64)(unsafe.Add(mBase, uint32(v19)+40)) = base.I64_extend_i32_u(v96)
						v100 = v92 & int32(_a_F_heap_page_items_4)
						*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = base.I64_extend_i32_u(v100)
						*(*int64)(unsafe.Add(mBase, uint32(v19)+32)) = base.I64_extend_i32_u(int32(base.Ui32(v92)>>(uint(int32(15))%32)) & int32(3))
						if base.B2i32(v100 != (v100+int32(7))&int32(_a_F_heap_page_items_5))|base.B2i32(base.Ui32(v92) < base.Ui32(int32(_a_F_heap_page_items_6)))|base.B2i32(base.Ui32(int32(_a_F_heap_page_items_7)) < base.Ui32(v100+v96)) == int32(0) {
							v123 = v100 + v83
							v124 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v123))))
							*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = v124
							v126 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v123)+4)))
							*(*int64)(unsafe.Add(mBase, uint32(v19)+56)) = v126
							v128 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v123)+8)))
							*(*int64)(unsafe.Add(mBase, uint32(v19)+72)) = base.I64_extend_i32_u(v123 + int32(12))
							*(*int64)(unsafe.Add(mBase, uint32(v19)+64)) = v128
							v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123)+18)))
							v136 = int64(65535)
							*(*int64)(unsafe.Add(mBase, uint32(v19)+80)) = base.I64_extend_i32_u(v134) & v136
							v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123)+20)))
							*(*int64)(unsafe.Add(mBase, uint32(v19)+88)) = base.I64_extend_i32_u(v139) & v136
							v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+22)))
							*(*int64)(unsafe.Add(mBase, uint32(v19)+96)) = base.I64_extend_i32_u(v144)
							if base.B2i32(base.Ui32(v144) < base.Ui32(int32(23)))|base.B2i32(base.Ui32(v96) < base.Ui32(v144))|base.B2i32((v144+int32(7))&int32(504) != v144) == int32(0) {
								if v139&int32(1) != 0 {
									v166 = int32(base.Ui32(v134&int32(2047)+int32(7)) >> (uint(int32(3)) % 32))
									if base.Ui32(v166) <= base.Ui32(v144-int32(23)) {
										v171 = v123 + int32(23)
										v172 = int32(0)
										v174 = v166 << (uint(int32(3)) % 32)
										v177 = F_palloc(m, v174+int32(1))
										mBase = m.M
										v178 = m.ExcPending
										if v178 != 0 {
											return int64(0)
										} else {
											if v174 == int32(0) {
											} else {
												if v174 != int32(1) {
													v188 = v172
													v201 = int32(0)
													for {
														v208 = v171 + int32(base.Ui32(v188)>>(uint(int32(3))%32))
														v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
														if int32(base.Ui32(v209)>>(uint(v188&int32(6))%32))&int32(1) != 0 {
															v215 = int32(49)
														} else {
															v215 = int32(48)
														}
														*(*uint8)(unsafe.Add(mBase, uint32(v188+v177))) = uint8(v215)
														v217 = int32(1)
														v218 = v188 | v217
														v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
														if int32(base.Ui32(v222)>>(uint(v218&int32(7))%32))&v217 != 0 {
															v228 = int32(49)
														} else {
															v228 = int32(48)
														}
														*(*uint8)(unsafe.Add(mBase, uint32(v177+v218))) = uint8(v228)
														v230 = int32(2)
														v231 = v188 + v230
														v233 = v201 + v230
														if v233 != v174&int32(4094) {
															v188 = v231
															v201 = v233
															continue
														} else {
															break
														}
														break
													}
													if v174&int32(1) == int32(0) {
													} else {
														v238 = v231
														v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171+int32(base.Ui32(v238)>>(uint(int32(3))%32))))))
														if int32(base.Ui32(v259)>>(uint(v238&int32(7))%32))&int32(1) != 0 {
															v265 = int32(49)
														} else {
															v265 = int32(48)
														}
														*(*uint8)(unsafe.Add(mBase, uint32(v238+v177))) = uint8(v265)
													}
												} else {
													v238 = v172
													v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171+int32(base.Ui32(v238)>>(uint(int32(3))%32))))))
													if int32(base.Ui32(v259)>>(uint(v238&int32(7))%32))&int32(1) != 0 {
														v265 = int32(49)
													} else {
														v265 = int32(48)
													}
													*(*uint8)(unsafe.Add(mBase, uint32(v238+v177))) = uint8(v265)
												}
											}
											v284 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v174+v177))) = uint8(v284)
											v286 = F_cstring_to_text(m, v177)
											mBase = m.M
											v287 = m.ExcPending
											if v287 != 0 {
												return int64(0)
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v19)+104)) = base.I64_extend_i32_u(v286)
												v290 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123)+20)))
												v296 = v290
												if v296&int32(8) != 0 {
													v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+22)))
													v317 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v123+v313-int32(4)))))
													*(*int64)(unsafe.Add(mBase, uint32(v19)+112)) = v317
													v322 = v313
												} else {
													v319 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v19)+12)) = uint8(v319)
													v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+22)))
													v322 = v321
												}
												v323 = v96 - v322
												v325 = v323 + int32(4)
												v326 = F_palloc(m, v325)
												mBase = m.M
												v327 = m.ExcPending
												if v327 != 0 {
													return int64(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v326))) = v325 << (uint(int32(2)) % 32)
													v331 = int32(0)
													if base.B2i32(v323 == v331)|base.B2i32(v323 <= v331) == v331 {
														v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+22)))
														base.MemoryCopy(m, v326+int32(4), v123+v340, v323)
													} else {
													}
													*(*int64)(unsafe.Add(mBase, uint32(v19)+120)) = base.I64_extend_i32_u(v326)
													v369 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
													v372 = F_heap_form_tuple(m, v369, v19+int32(16), v19)
													mBase = m.M
													v373 = m.ExcPending
													if v373 != 0 {
														return int64(0)
													} else {
														v374 = *(*int32)(unsafe.Add(mBase, uint32(v372)+16))
														v375 = F_HeapTupleHeaderGetDatum(m, v374)
														mBase = m.M
														v376 = m.ExcPending
														if v376 != 0 {
															return int64(0)
														} else {
															v377 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+8)))
															v378 = int32(1)
															v379 = v377 + v378
															*(*uint16)(unsafe.Add(mBase, uint32(v82)+8)) = uint16(v379)
															v381 = *(*int64)(unsafe.Add(mBase, uint32(v78)))
															*(*int64)(unsafe.Add(mBase, uint32(v78))) = v381 + int64(1)
															v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v385)+20)) = v378
															v411 = v375
															m.G0 = v19 + int32(128)
															return v411
														}
													}
												}
											}
										}
									} else {
										v291 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v19)+11)) = uint8(v291)
										v296 = v139
										if v296&int32(8) != 0 {
											v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+22)))
											v317 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v123+v313-int32(4)))))
											*(*int64)(unsafe.Add(mBase, uint32(v19)+112)) = v317
											v322 = v313
										} else {
											v319 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v19)+12)) = uint8(v319)
											v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+22)))
											v322 = v321
										}
										v323 = v96 - v322
										v325 = v323 + int32(4)
										v326 = F_palloc(m, v325)
										mBase = m.M
										v327 = m.ExcPending
										if v327 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v326))) = v325 << (uint(int32(2)) % 32)
											v331 = int32(0)
											if base.B2i32(v323 == v331)|base.B2i32(v323 <= v331) == v331 {
												v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+22)))
												base.MemoryCopy(m, v326+int32(4), v123+v340, v323)
											} else {
											}
											*(*int64)(unsafe.Add(mBase, uint32(v19)+120)) = base.I64_extend_i32_u(v326)
											v369 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
											v372 = F_heap_form_tuple(m, v369, v19+int32(16), v19)
											mBase = m.M
											v373 = m.ExcPending
											if v373 != 0 {
												return int64(0)
											} else {
												v374 = *(*int32)(unsafe.Add(mBase, uint32(v372)+16))
												v375 = F_HeapTupleHeaderGetDatum(m, v374)
												mBase = m.M
												v376 = m.ExcPending
												if v376 != 0 {
													return int64(0)
												} else {
													v377 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+8)))
													v378 = int32(1)
													v379 = v377 + v378
													*(*uint16)(unsafe.Add(mBase, uint32(v82)+8)) = uint16(v379)
													v381 = *(*int64)(unsafe.Add(mBase, uint32(v78)))
													*(*int64)(unsafe.Add(mBase, uint32(v78))) = v381 + int64(1)
													v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v385)+20)) = v378
													v411 = v375
													m.G0 = v19 + int32(128)
													return v411
												}
											}
										}
									}
								} else {
									v293 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v19)+11)) = uint8(v293)
									v296 = v139
									if v296&int32(8) != 0 {
										v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+22)))
										v317 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v123+v313-int32(4)))))
										*(*int64)(unsafe.Add(mBase, uint32(v19)+112)) = v317
										v322 = v313
									} else {
										v319 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v19)+12)) = uint8(v319)
										v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+22)))
										v322 = v321
									}
									v323 = v96 - v322
									v325 = v323 + int32(4)
									v326 = F_palloc(m, v325)
									mBase = m.M
									v327 = m.ExcPending
									if v327 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v326))) = v325 << (uint(int32(2)) % 32)
										v331 = int32(0)
										if base.B2i32(v323 == v331)|base.B2i32(v323 <= v331) == v331 {
											v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+22)))
											base.MemoryCopy(m, v326+int32(4), v123+v340, v323)
										} else {
										}
										*(*int64)(unsafe.Add(mBase, uint32(v19)+120)) = base.I64_extend_i32_u(v326)
										v369 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
										v372 = F_heap_form_tuple(m, v369, v19+int32(16), v19)
										mBase = m.M
										v373 = m.ExcPending
										if v373 != 0 {
											return int64(0)
										} else {
											v374 = *(*int32)(unsafe.Add(mBase, uint32(v372)+16))
											v375 = F_HeapTupleHeaderGetDatum(m, v374)
											mBase = m.M
											v376 = m.ExcPending
											if v376 != 0 {
												return int64(0)
											} else {
												v377 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+8)))
												v378 = int32(1)
												v379 = v377 + v378
												*(*uint16)(unsafe.Add(mBase, uint32(v82)+8)) = uint16(v379)
												v381 = *(*int64)(unsafe.Add(mBase, uint32(v78)))
												*(*int64)(unsafe.Add(mBase, uint32(v78))) = v381 + int64(1)
												v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v385)+20)) = v378
												v411 = v375
												m.G0 = v19 + int32(128)
												return v411
											}
										}
									}
								}
							} else {
								v345 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v19)+13)) = uint8(v345)
								v347 = int32(257)
								*(*uint16)(unsafe.Add(mBase, uint32(v19)+11)) = uint16(v347)
								v369 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
								v372 = F_heap_form_tuple(m, v369, v19+int32(16), v19)
								mBase = m.M
								v373 = m.ExcPending
								if v373 != 0 {
									return int64(0)
								} else {
									v374 = *(*int32)(unsafe.Add(mBase, uint32(v372)+16))
									v375 = F_HeapTupleHeaderGetDatum(m, v374)
									mBase = m.M
									v376 = m.ExcPending
									if v376 != 0 {
										return int64(0)
									} else {
										v377 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+8)))
										v378 = int32(1)
										v379 = v377 + v378
										*(*uint16)(unsafe.Add(mBase, uint32(v82)+8)) = uint16(v379)
										v381 = *(*int64)(unsafe.Add(mBase, uint32(v78)))
										*(*int64)(unsafe.Add(mBase, uint32(v78))) = v381 + int64(1)
										v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v385)+20)) = v378
										v411 = v375
										m.G0 = v19 + int32(128)
										return v411
									}
								}
							}
						} else {
							v349 = int32(257)
							*(*uint16)(unsafe.Add(mBase, uint32(v19)+12)) = uint16(v349)
							*(*int64)(unsafe.Add(mBase, uint32(v19)+4)) = int64(72340172838076673)
							v369 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
							v372 = F_heap_form_tuple(m, v369, v19+int32(16), v19)
							mBase = m.M
							v373 = m.ExcPending
							if v373 != 0 {
								return int64(0)
							} else {
								v374 = *(*int32)(unsafe.Add(mBase, uint32(v372)+16))
								v375 = F_HeapTupleHeaderGetDatum(m, v374)
								mBase = m.M
								v376 = m.ExcPending
								if v376 != 0 {
									return int64(0)
								} else {
									v377 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+8)))
									v378 = int32(1)
									v379 = v377 + v378
									*(*uint16)(unsafe.Add(mBase, uint32(v82)+8)) = uint16(v379)
									v381 = *(*int64)(unsafe.Add(mBase, uint32(v78)))
									*(*int64)(unsafe.Add(mBase, uint32(v78))) = v381 + int64(1)
									v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v385)+20)) = v378
									v411 = v375
									m.G0 = v19 + int32(128)
									return v411
								}
							}
						}
					} else {
						F_end_MultiFuncCall(m, l0)
						mBase = m.M
						v389 = m.ExcPending
						if v389 != 0 {
							return int64(0)
						} else {
							v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v390)+20)) = int32(2)
							v393 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v393)
							v411 = int64(0)
							m.G0 = v19 + int32(128)
							return v411
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v419 = m.ExcPending
				if v419 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v422 = m.ExcPending
					if v422 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_heap_page_items_8), int32(0))
						mBase = m.M
						v426 = m.ExcPending
						if v426 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_heap_page_items_2), int32(140), int32(_a_F_heap_page_items_3))
							mBase = m.M
							v431 = m.ExcPending
							if v431 != 0 {
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
}
func F_heap_page_prune_and_freeze(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v117 int64
	_ = v117
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int64
	_ = v144
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
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
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v418 int32
	_ = v418
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v635 int32
	_ = v635
	var v656 int32
	_ = v656
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v707 int32
	_ = v707
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v739 int32
	_ = v739
	var v745 int32
	_ = v745
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v791 int32
	_ = v791
	var v799 int32
	_ = v799
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v848 int32
	_ = v848
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v883 int32
	_ = v883
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v899 int32
	_ = v899
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v934 int32
	_ = v934
	var v939 int32
	_ = v939
	var v942 int32
	_ = v942
	var v950 int32
	_ = v950
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1002 int32
	_ = v1002
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1017 int32
	_ = v1017
	var v1021 int32
	_ = v1021
	var v1029 int32
	_ = v1029
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1100 int32
	_ = v1100
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1137 int32
	_ = v1137
	var v1139 int32
	_ = v1139
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1185 int32
	_ = v1185
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1197 int32
	_ = v1197
	var v1202 int32
	_ = v1202
	var v1210 int32
	_ = v1210
	var v1214 int32
	_ = v1214
	var v1219 int32
	_ = v1219
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1236 int32
	_ = v1236
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1263 int32
	_ = v1263
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1274 int32
	_ = v1274
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1282 int32
	_ = v1282
	var v1290 int32
	_ = v1290
	var v1292 int32
	_ = v1292
	var v1294 int32
	_ = v1294
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1301 int32
	_ = v1301
	var v1311 int32
	_ = v1311
	var v1314 int32
	_ = v1314
	var v1338 int32
	_ = v1338
	var v1341 int32
	_ = v1341
	var v1349 int32
	_ = v1349
	var v1353 int32
	_ = v1353
	var v1397 int32
	_ = v1397
	var v1427 int32
	_ = v1427
	var v1429 int32
	_ = v1429
	var v1431 int32
	_ = v1431
	var v1499 int32
	_ = v1499
	var v1501 int32
	_ = v1501
	var v1522 int32
	_ = v1522
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1555 int32
	_ = v1555
	var v1561 int32
	_ = v1561
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1570 int32
	_ = v1570
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1582 int32
	_ = v1582
	var v1587 int32
	_ = v1587
	var v1591 int32
	_ = v1591
	var v1598 int32
	_ = v1598
	var v1605 int32
	_ = v1605
	var v1610 int32
	_ = v1610
	var v1614 int32
	_ = v1614
	var v1652 int32
	_ = v1652
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1658 int32
	_ = v1658
	var v1661 int32
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1666 int32
	_ = v1666
	var v1669 int32
	_ = v1669
	var v1672 int32
	_ = v1672
	var v1675 int32
	_ = v1675
	var v1678 int32
	_ = v1678
	var v1680 int32
	_ = v1680
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1684 int32
	_ = v1684
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1698 int32
	_ = v1698
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1708 int64
	_ = v1708
	var v1710 int32
	_ = v1710
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1723 int32
	_ = v1723
	var v1726 int32
	_ = v1726
	var v1728 int32
	_ = v1728
	var v1738 int32
	_ = v1738
	var v1744 int32
	_ = v1744
	var v1746 int32
	_ = v1746
	var v1752 int32
	_ = v1752
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1757 int64
	_ = v1757
	var v1758 int64
	_ = v1758
	var v1763 int32
	_ = v1763
	var v1772 int32
	_ = v1772
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1782 int32
	_ = v1782
	var v1785 int32
	_ = v1785
	var v1787 int32
	_ = v1787
	var v1797 int32
	_ = v1797
	var v1803 int32
	_ = v1803
	var v1805 int32
	_ = v1805
	var v1811 int32
	_ = v1811
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1816 int64
	_ = v1816
	var v1817 int64
	_ = v1817
	var v1822 int32
	_ = v1822
	var v1831 int32
	_ = v1831
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1837 int32
	_ = v1837
	var v1842 int32
	_ = v1842
	var v1848 int32
	_ = v1848
	var v1850 int32
	_ = v1850
	var v1856 int32
	_ = v1856
	var v1870 int32
	_ = v1870
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1900 int32
	_ = v1900
	var v1903 int32
	_ = v1903
	var v1904 int32
	_ = v1904
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1912 int32
	_ = v1912
	var v1914 int32
	_ = v1914
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1919 int32
	_ = v1919
	var v1922 int32
	_ = v1922
	var v1961 int32
	_ = v1961
	var v1964 int32
	_ = v1964
	var v1970 int32
	_ = v1970
	var v1975 int32
	_ = v1975
	var v1979 int32
	_ = v1979
	var v1982 int32
	_ = v1982
	var v1986 int32
	_ = v1986
	var v1991 int32
	_ = v1991
	var v1995 int32
	_ = v1995
	var v1997 int32
	_ = v1997
	var v2001 int32
	_ = v2001
	var v2015 int32
	_ = v2015
	var v2036 int32
	_ = v2036
	var v2039 int32
	_ = v2039
	var v2041 int32
	_ = v2041
	var v2044 int32
	_ = v2044
	var v2045 int32
	_ = v2045
	var v2050 int32
	_ = v2050
	var v2051 int32
	_ = v2051
	var v2057 int32
	_ = v2057
	var v2061 int32
	_ = v2061
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2074 int32
	_ = v2074
	var v2075 int64
	_ = v2075
	var v2078 int64
	_ = v2078
	var v2085 int32
	_ = v2085
	var v2088 int32
	_ = v2088
	var v2090 int32
	_ = v2090
	var v2100 int32
	_ = v2100
	var v2106 int32
	_ = v2106
	var v2108 int32
	_ = v2108
	var v2114 int32
	_ = v2114
	var v2115 int32
	_ = v2115
	var v2116 int32
	_ = v2116
	var v2119 int64
	_ = v2119
	var v2120 int64
	_ = v2120
	var v2125 int32
	_ = v2125
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2135 int32
	_ = v2135
	var v2137 int32
	_ = v2137
	var v2138 int32
	_ = v2138
	var v2139 int32
	_ = v2139
	var v2141 int32
	_ = v2141
	var v2143 int32
	_ = v2143
	var v2145 int32
	_ = v2145
	var v2149 int32
	_ = v2149
	var v2150 int32
	_ = v2150
	var v2151 int32
	_ = v2151
	var v2163 int32
	_ = v2163
	var v2164 int32
	_ = v2164
	var v2166 int32
	_ = v2166
	var v2168 int32
	_ = v2168
	var v2180 int32
	_ = v2180
	var v2181 int32
	_ = v2181
	var v2183 int32
	_ = v2183
	var v2185 int32
	_ = v2185
	var v2188 int32
	_ = v2188
	var v2189 int32
	_ = v2189
	var v2191 int32
	_ = v2191
	var v2197 int32
	_ = v2197
	var v2198 int32
	_ = v2198
	var v2200 int32
	_ = v2200
	var v2201 int32
	_ = v2201
	var v2203 int32
	_ = v2203
	var v2209 int32
	_ = v2209
	var v2212 int32
	_ = v2212
	var v2218 int32
	_ = v2218
	var v2219 int32
	_ = v2219
	var v2223 int32
	_ = v2223
	var v2229 int32
	_ = v2229
	var v2231 int32
	_ = v2231
	var v2237 int32
	_ = v2237
	var v2238 int32
	_ = v2238
	var v2239 int32
	_ = v2239
	var v2243 int32
	_ = v2243
	var v2245 int32
	_ = v2245
	var v2260 int32
	_ = v2260
	var v2261 int32
	_ = v2261
	var v2284 int32
	_ = v2284
	var v2285 int32
	_ = v2285
	var v2288 int32
	_ = v2288
	var v2289 int32
	_ = v2289
	var v2291 int32
	_ = v2291
	var v2294 int32
	_ = v2294
	var v2298 int32
	_ = v2298
	var v2305 int32
	_ = v2305
	var v2307 int32
	_ = v2307
	var v2318 int32
	_ = v2318
	var v2342 int32
	_ = v2342
	var v2346 int32
	_ = v2346
	var v2386 int32
	_ = v2386
	var v2388 int32
	_ = v2388
	var v2390 int32
	_ = v2390
	var v2403 int32
	_ = v2403
	var v2404 int32
	_ = v2404
	var v2427 int32
	_ = v2427
	var v2428 int32
	_ = v2428
	var v2431 int32
	_ = v2431
	var v2433 int32
	_ = v2433
	var v2439 int32
	_ = v2439
	var v2445 int32
	_ = v2445
	var v2452 int32
	_ = v2452
	var v2454 int32
	_ = v2454
	var v2465 int32
	_ = v2465
	var v2497 int32
	_ = v2497
	var v2498 int32
	_ = v2498
	var v2521 int32
	_ = v2521
	var v2522 int32
	_ = v2522
	var v2530 int32
	_ = v2530
	var v2566 int32
	_ = v2566
	var v2568 int32
	_ = v2568
	var v2570 int32
	_ = v2570
	var v2583 int32
	_ = v2583
	var v2584 int32
	_ = v2584
	var v2607 int32
	_ = v2607
	var v2608 int32
	_ = v2608
	var v2611 int32
	_ = v2611
	var v2613 int32
	_ = v2613
	var v2619 int32
	_ = v2619
	var v2625 int32
	_ = v2625
	var v2632 int32
	_ = v2632
	var v2634 int32
	_ = v2634
	var v2645 int32
	_ = v2645
	var v2677 int32
	_ = v2677
	var v2678 int32
	_ = v2678
	var v2701 int32
	_ = v2701
	var v2702 int32
	_ = v2702
	var v2710 int32
	_ = v2710
	var v2744 int32
	_ = v2744
	var v2776 int32
	_ = v2776
	var v2777 int32
	_ = v2777
	var v2778 int32
	_ = v2778
	var v2782 int32
	_ = v2782
	var v2788 int32
	_ = v2788
	var v2790 int32
	_ = v2790
	var v2796 int32
	_ = v2796
	var v2811 int32
	_ = v2811
	var v2836 int32
	_ = v2836
	var v2837 int32
	_ = v2837
	var v2838 int32
	_ = v2838
	var v2841 int32
	_ = v2841
	var v2844 int32
	_ = v2844
	var v2845 int32
	_ = v2845
	var v2847 int32
	_ = v2847
	var v2852 int32
	_ = v2852
	var v2853 int32
	_ = v2853
	var v2858 int32
	_ = v2858
	var v2860 int32
	_ = v2860
	var v2863 int32
	_ = v2863
	var v2927 int32
	_ = v2927
	var v2928 int32
	_ = v2928
	var v2930 int32
	_ = v2930
	var v2932 int32
	_ = v2932
	var v2935 int32
	_ = v2935
	var v2936 int64
	_ = v2936
	var v2938 int32
	_ = v2938
	var v2940 int32
	_ = v2940
	var v2941 int32
	_ = v2941
	var v2942 int32
	_ = v2942
	var v2943 int32
	_ = v2943
	var v2944 int32
	_ = v2944
	var v2946 int32
	_ = v2946
	var v2948 int32
	_ = v2948
	var v2949 int32
	_ = v2949
	var v2950 int32
	_ = v2950
	var v2951 int32
	_ = v2951
	var v2955 int32
	_ = v2955
	var v2958 int32
	_ = v2958
	var v2959 int32
	_ = v2959
	var v2960 int32
	_ = v2960
	var v2961 int32
	_ = v2961
	var v2963 int32
	_ = v2963
	var v2964 int32
	_ = v2964
	var v2966 int32
	_ = v2966
	var v2970 int32
	_ = v2970
	var v2973 int32
	_ = v2973
	var v2976 int32
	_ = v2976
	var v2979 int32
	_ = v2979
	var v2982 int32
	_ = v2982
	var v2984 int32
	_ = v2984
	var v3016 int32
	_ = v3016
	var v3018 int32
	_ = v3018
	var v3022 int32
	_ = v3022
	var v3024 int32
	_ = v3024
	var v3025 int32
	_ = v3025
	var v3027 int32
	_ = v3027
	var v3029 int32
	_ = v3029
	var v3031 int32
	_ = v3031
	var v3033 int32
	_ = v3033
	var v3035 int32
	_ = v3035
	var v3037 int32
	_ = v3037
	var v3038 int32
	_ = v3038
	var v3045 int32
	_ = v3045
	var v3050 int32
	_ = v3050
	var v3052 int32
	_ = v3052
	var v3055 int32
	_ = v3055
	var v3059 int32
	_ = v3059
	var v3064 int32
	_ = v3064
	var v3067 int32
	_ = v3067
	var v3072 int32
	_ = v3072
	var v3074 int32
	_ = v3074
	var v3076 int32
	_ = v3076
	var v3078 int32
	_ = v3078
	v6 = int32(0)
	v32 = m.G0
	v34 = v32 - int32(_a_F_heap_page_prune_and_freeze_0)
	m.G0 = v34
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v36
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	v39 = int32(1)
	v40 = v38 & v39
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+32)) = uint8(v40)
	v45 = int32(base.Ui32(v38)>>(uint(int32(3))%32)) & v39
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+34)) = uint8(v45)
	v50 = int32(base.Ui32(v38)>>(uint(v39)%32)) & v39
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+33)) = uint8(v50)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+36)) = v52
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+40)) = v54
	v57 = *(*int64)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[0]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v58 < v6 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+44)) = v77
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+48)) = v79
	if v79 < int32(0) {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[1]))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v62+(v58^int32(-1))*int32(56))+16))
	v77 = v68
	goto L1
L3:
	;
	goto L4
L4:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[2]))
	v71 = int32(56)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v70+v58*v71-v71)+16))
	v77 = v76
	goto L1
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+52)) = v98
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v101 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[3]))) = uint8(v101)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[4]))) = v100
	v106 = F_visibilitymap_get_status(m, v54, v77, v34+int32(_a_F_heap_page_prune_and_freeze_1))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[5]))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v84+(v79^int32(-1))<<(uint(int32(2))%32))))
	v98 = v90
	goto L5
L7:
	;
	goto L8
L8:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[6]))
	v98 = v92 + v79<<(uint(int32(13))%32) + int32(-8192)
	goto L5
L9:
	;
	return
L10:
	;
	v108 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[7]))) = v108
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[8]))) = uint8(v108)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[9]))) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[10]))) = v108
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[11]))) = uint8(v106)
	v117 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v34)+56)) = v117
	*(*int64)(unsafe.Add(mBase, uint32(v34)+64)) = v117
	*(*int64)(unsafe.Add(mBase, uint32(v34)+72)) = v117
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+33)))
	if v123 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[12]))) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v137))) = v138
	v141 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[13]))) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[14]))) = v138
	v144 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[15]))) = v144
	*(*int64)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[16]))) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[17]))) = l1 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[18]))) = v141
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+34)))
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[19]))) = uint8(v153)
	v155 = v153 & v123
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[20]))) = uint8(v155)
	if v106&int32(3) == v141 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[21]))) = v128
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v136 = v128
	v137 = v34 + int32(_a_F_heap_page_prune_and_freeze_2)
	v138 = v130
	goto L11
L13:
	;
	goto L14
L14:
	;
	v133 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[22]))) = v133
	v136 = v6
	v137 = v34 + int32(_a_F_heap_page_prune_and_freeze_3)
	v138 = v133
	goto L11
L15:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v171&int32(4) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+10)))
	if v162&int32(4) != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v167 = int32(0)
	F_heap_page_fix_vm_corruption(m, v34+int32(28), v167, v167)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	m.G0 = v34 + int32(_a_F_heap_page_prune_and_freeze_0)
	return
L20:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v394 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v393)+12)))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v34)+44))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v34)+40))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v396)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[23]))) = v397
	if base.Ui32(int32(25)) <= base.Ui32(v394) {
		goto L57
	} else {
		goto L58
	}
L21:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[11]))))
	if v176&int32(2) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if v176&int32(1) == int32(0) {
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188)+12)))
	base.MemoryFill(m, l1, int32(0), int32(612))
	v194 = v188 + int32(20)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v188)+20))
	if v195 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+33)))
	if v185&int32(1) != 0 {
		goto L20
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v194))) = int32(0)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	F_MarkBufferDirtyHint(m, v198, int32(1))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L9
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188)+12)))
	if base.Ui32(v202) < base.Ui32(int32(25)) {
		goto L19
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v189) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v212 = int32(base.Ui32(v189+int32(_a_F_heap_page_prune_and_freeze_4)) >> (uint(int32(2)) % 32))
	goto L34
L33:
	;
	v212 = int32(0)
	goto L34
L34:
	;
	if v212&int32(_a_F_heap_page_prune_and_freeze_5) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[24])))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v217
	goto L19
L36:
	;
	goto L37
L37:
	;
	v219 = int32(1)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[24])))
	v224 = (v212 + v219) & int32(_a_F_heap_page_prune_and_freeze_5)
	if base.Ui32(int32(3)) <= base.Ui32(v224) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v362
	goto L19
L39:
	;
	v227 = int32(2)
	if base.Ui32(v224) <= base.Ui32(v227) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v313 = v220
	v317 = v219
	goto L41
L41:
	;
	v344 = v194 + v317<<(uint(int32(2))%32)
	v345 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v344)+1)))
	if v345&int32(384) == int32(0) {
		v362 = v313
		goto L38
	} else {
		goto L55
	}
L42:
	;
	v230 = v227
	goto L44
L43:
	;
	v230 = v224
	goto L44
L44:
	;
	v231 = int32(1)
	v232 = v230 - v231
	v240 = v220
	v244 = v219
	v246 = int32(0)
	goto L45
L45:
	;
	v271 = v194 + v244<<(uint(int32(2))%32)
	v272 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v271)+1)))
	if v272&int32(384) == int32(0) {
		v287 = v240
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v232&v231 == int32(0) {
		v362 = v303
		goto L38
	} else {
		goto L54
	}
L47:
	;
	v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v271)+5)))
	if v288&int32(384) == int32(0) {
		v303 = v287
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v277 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+23)) = uint8(v277)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	if v279&int32(_a_F_heap_page_prune_and_freeze_6) != int32(_a_F_heap_page_prune_and_freeze_7) {
		v287 = v240
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v285 = v240 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[24]))) = v285
	v287 = v285
	goto L47
L50:
	;
	v304 = int32(2)
	v305 = v244 + v304
	v307 = v246 + v304
	if v307 != v232&int32(-2) {
		v240 = v303
		v244 = v305
		v246 = v307
		goto L45
	} else {
		goto L53
	}
L51:
	;
	v293 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+23)) = uint8(v293)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v271)+4))
	if v295&int32(_a_F_heap_page_prune_and_freeze_6) != int32(_a_F_heap_page_prune_and_freeze_7) {
		v303 = v287
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v301 = v287 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[24]))) = v301
	v303 = v301
	goto L50
L53:
	;
	goto L46
L54:
	;
	v313 = v303
	v317 = v305
	goto L41
L55:
	;
	v350 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+23)) = uint8(v350)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	if v352&int32(_a_F_heap_page_prune_and_freeze_6) != int32(_a_F_heap_page_prune_and_freeze_7) {
		v362 = v313
		goto L38
	} else {
		goto L56
	}
L56:
	;
	v358 = v313 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[24]))) = v358
	v362 = v358
	goto L38
L57:
	;
	v406 = int32(base.Ui32(v394+int32(_a_F_heap_page_prune_and_freeze_4)) >> (uint(int32(2)) % 32))
	goto L59
L58:
	;
	v406 = int32(0)
	goto L59
L59:
	;
	v408 = v406 & int32(_a_F_heap_page_prune_and_freeze_5)
	if v408 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v410 = int32(base.Ui32(v395) >> (uint(int32(16)) % 32))
	v418 = v34 + int32(_a_F_heap_page_prune_and_freeze_8)
	v429 = v406
	v430 = v408
	goto L63
L61:
	;
	goto L62
L62:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[10])))
	v621 = v619 - int32(1)
	if int32(0) <= v621 {
		goto L92
	} else {
		goto L93
	}
L63:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v430)
	v455 = v430 + (v34 + int32(_a_F_heap_page_prune_and_freeze_9))
	v456 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v455))) = uint8(v456)
	v458 = v430 + (v34 + int32(_a_F_heap_page_prune_and_freeze_10))
	v459 = int32(255)
	*(*uint8)(unsafe.Add(mBase, uint32(v458))) = uint8(v459)
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v393+int32(20)+v430<<(uint(int32(2))%32))))
	switch int32(base.Ui32(v464)>>(uint(int32(15))%32))&int32(3) - int32(1) {
	case 0:
		goto L66
	case 1:
		goto L67
	case 2:
		goto L68
	default:
		goto L69
	}
L64:
	;
	goto L62
L65:
	;
	v582 = int32(1)
	v585 = v429 - v582
	if v585&int32(_a_F_heap_page_prune_and_freeze_5) != 0 {
		v429 = v585
		v430 = v430 - v582
		goto L63
	} else {
		goto L91
	}
L66:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[25]))) = uint16(v430)
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[26]))) = uint16(v395)
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[27]))) = uint16(v410)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[28]))) = int32(base.Ui32(v464) >> (uint(int32(17)) % 32))
	v523 = v393 + v464&int32(_a_F_heap_page_prune_and_freeze_11)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[29]))) = v523
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	v530 = F_HeapTupleSatisfiesVacuumHorizon(m, v34+int32(_a_F_heap_page_prune_and_freeze_12), v527, v34+int32(_a_F_heap_page_prune_and_freeze_13))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L9
	} else {
		goto L76
	}
L67:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[10])))
	v508 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[10]))) = v507 + v508
	*(*uint16)(unsafe.Add(mBase, uint32(v418+v507<<(uint(v508)%32)))) = uint16(v430)
	goto L65
L68:
	;
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+32)))
	v474 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v455))) = uint8(v474)
	if v473 == v474 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v471 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v455))) = uint8(v471)
	goto L65
L70:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	v479 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v34+int32(1826)+v478<<(uint(v479)%32)))) = uint16(v430)
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+72)) = v483 + v479
	goto L65
L71:
	;
	goto L72
L72:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[13])))
	v488 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[13]))) = v487 + v488
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[17])))
	*(*uint16)(unsafe.Add(mBase, uint32(v491+v487<<(uint(v488)%32)))) = uint16(v430)
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+10)))
	if v497&int32(4) == int32(0) {
		goto L65
	} else {
		goto L73
	}
L73:
	;
	F_heap_page_fix_vm_corruption(m, v34+int32(28), v430, int32(1))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L9
	} else {
		goto L74
	}
L74:
	;
	goto L65
L75:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v458))) = uint8(v555)
	v559 = int32(*(*int16)(unsafe.Add(mBase, uint32(v523)+18)))
	if int32(0) <= v559 {
		goto L88
	} else {
		goto L89
	}
L76:
	;
	if v530 != int32(2) {
		v555 = v530
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v34)+36))
	if v534 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v34)+28))
	v552 = F_GlobalVisTestIsRemovableXid(m, v551, v547)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L9
	} else {
		goto L84
	}
L79:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[30])))
	v547 = v537
	goto L78
L80:
	;
	goto L81
L81:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[30])))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v534)+8))
	if v539 == int32(0) {
		v547 = v538
		goto L78
	} else {
		goto L82
	}
L82:
	;
	v542 = int32(0)
	if v538-v539 < v542 {
		v555 = v542
		goto L75
	} else {
		goto L83
	}
L83:
	;
	v547 = v538
	goto L78
L84:
	;
	if v552 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v554 = int32(0)
	goto L87
L86:
	;
	v554 = int32(2)
	goto L87
L87:
	;
	v555 = v554
	goto L75
L88:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[10])))
	v563 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[10]))) = v562 + v563
	*(*uint16)(unsafe.Add(mBase, uint32(v418+v562<<(uint(v563)%32)))) = uint16(v430)
	goto L65
L89:
	;
	goto L90
L90:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[9])))
	v571 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[9]))) = v570 + v571
	*(*uint16)(unsafe.Add(mBase, uint32(v34+int32(_a_F_heap_page_prune_and_freeze_14)+v570<<(uint(v571)%32)))) = uint16(v430)
	goto L65
L91:
	;
	goto L64
L92:
	;
	v625 = v34 + int32(1244)
	v627 = v34 + int32(1826)
	v629 = v34 + int32(80)
	v635 = v34 + int32(_a_F_heap_page_prune_and_freeze_9)
	v656 = v621
	goto L95
L93:
	;
	goto L94
L94:
	;
	v1499 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[9])))
	v1501 = v1499 - int32(1)
	if int32(0) <= v1501 {
		goto L189
	} else {
		goto L190
	}
L95:
	;
	v672 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34+int32(_a_F_heap_page_prune_and_freeze_8)+v656<<(uint(int32(1))%32)))))
	v673 = v635 + v672
	v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v673))))
	if v674 != 0 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	goto L94
L97:
	;
	if int32(0) < v656 {
		v656 = v656 - int32(1)
		goto L95
	} else {
		goto L188
	}
L98:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v672)
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v678 = v676 + int32(20)
	v681 = v678 + v672<<(uint(int32(2))%32)
	v682 = int32(0)
	v685 = int32(_a_F_heap_page_prune_and_freeze_5)
	v686 = v406 & v685
	if base.Ui32(v686) <= base.Ui32((v672-int32(1))&v685) {
		v824 = v682
		v825 = v682
		goto L100
	} else {
		goto L101
	}
L99:
	;
	if v908 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L100:
	;
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v681)))
	if base.B2i32(v848&int32(_a_F_heap_page_prune_and_freeze_6) != int32(_a_F_heap_page_prune_and_freeze_15))|base.B2i32(int32(1) < v824) != 0 {
		v907 = v824
		v908 = v825
		goto L99
	} else {
		goto L131
	}
L101:
	;
	v698 = v672
	v699 = v682
	v700 = v682
	v707 = v682
	goto L102
L102:
	;
	v724 = v698 & int32(_a_F_heap_page_prune_and_freeze_5)
	v726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v635+v724))))
	if v726 != 0 {
		v824 = v699
		v825 = v700
		goto L100
	} else {
		goto L104
	}
L103:
	;
	v824 = v808
	v825 = v809
	goto L100
L104:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v678+v724<<(uint(int32(2))%32))))
	if v730&int32(_a_F_heap_page_prune_and_freeze_6) == int32(_a_F_heap_page_prune_and_freeze_15) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	if base.Ui32((v807-int32(1))&int32(_a_F_heap_page_prune_and_freeze_5)) < base.Ui32(v686) {
		v698 = v807
		v699 = v808
		v700 = v809
		v707 = v811
		goto L102
	} else {
		goto L130
	}
L106:
	;
	if int32(0) < v699 {
		v824 = v699
		v825 = v700
		goto L100
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v750 = v676 + v730&int32(_a_F_heap_page_prune_and_freeze_11)
	if v707 != 0 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v739 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v34+int32(_a_F_heap_page_prune_and_freeze_13)+v699<<(uint(v739)%32)))) = uint16(v698)
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v681)))
	v807 = v745 & int32(_a_F_heap_page_prune_and_freeze_11)
	v808 = v699 + v739
	v809 = v700
	v811 = v707
	goto L105
L110:
	;
	v751 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v750)+20)))
	v752 = int32(768)
	if v751&v752 != v752 {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L112
L112:
	;
	v762 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v34+int32(_a_F_heap_page_prune_and_freeze_13)+v699<<(uint(v762)%32)))) = uint16(v698)
	v767 = v699 + v762
	v769 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v724+(v34+int32(_a_F_heap_page_prune_and_freeze_10))))))
	switch v769 {
	case 0:
		goto L118
	case 1, 3, 4:
		v907 = v767
		v908 = v700
		goto L99
	case 2:
		v785 = v700
		goto L117
	default:
		goto L119
	}
L113:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v750)))
	v758 = v756
	goto L115
L114:
	;
	v758 = int32(2)
	goto L115
L115:
	;
	if v758 != v707 {
		v824 = v699
		v825 = v700
		goto L100
	} else {
		goto L116
	}
L116:
	;
	goto L112
L117:
	;
	v786 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+19)))
	if v786&int32(64) == int32(0) {
		v907 = v767
		v908 = v785
		goto L99
	} else {
		goto L124
	}
L118:
	;
	F_HeapTupleHeaderAdvanceConflictHorizon(m, v750, v34+int32(60))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L9
	} else {
		goto L123
	}
L119:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L9
	} else {
		goto L120
	}
L120:
	;
	F_errmsg_internal(m, int32(_a_F_heap_page_prune_and_freeze_16), int32(0))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L9
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_heap_page_prune_and_freeze_17), int32(1631), int32(_a_F_heap_page_prune_and_freeze_18))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L9
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L123:
	;
	v785 = v767
	goto L117
L124:
	;
	v791 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v750)+20)))
	if v791&int32(2048)|base.B2i32(v791&int32(768) == int32(512)) != 0 {
		v907 = v767
		v908 = v785
		goto L99
	} else {
		goto L125
	}
L125:
	;
	v799 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v750)+16)))
	if v791&int32(_a_F_heap_page_prune_and_freeze_19) == int32(_a_F_heap_page_prune_and_freeze_20) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v804 = F_HeapTupleGetUpdateXid(m, v750)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L9
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v750)+4))
	v807 = v799
	v808 = v767
	v809 = v785
	v811 = v806
	goto L105
L129:
	;
	v807 = v799
	v808 = v767
	v809 = v785
	v811 = v804
	goto L105
L130:
	;
	goto L103
L131:
	;
	v856 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+32)))
	v857 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v673))) = uint8(v857)
	if v856 == v857 {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v890 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v889)+10)))
	if v890&int32(4) == int32(0) {
		goto L97
	} else {
		goto L136
	}
L133:
	;
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	v862 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v627+v861<<(uint(v862)%32)))) = uint16(v672)
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+72)) = v866 + v862
	goto L132
L134:
	;
	goto L135
L135:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v34)+68))
	v871 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v625+v870<<(uint(v871)%32)))) = uint16(v672)
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v34)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+68)) = v875 + v871
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[13])))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[13]))) = v879 + v871
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[17])))
	*(*uint16)(unsafe.Add(mBase, uint32(v883+v879<<(uint(v871)%32)))) = uint16(v672)
	goto L132
L136:
	;
	F_heap_page_fix_vm_corruption(m, v34+int32(28), v672, int32(1))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L9
	} else {
		goto L137
	}
L137:
	;
	goto L97
L138:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v681)))
	if v934&int32(_a_F_heap_page_prune_and_freeze_6) == int32(_a_F_heap_page_prune_and_freeze_15) {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	goto L140
L140:
	;
	if v907 == v908 {
		goto L149
	} else {
		goto L150
	}
L141:
	;
	v939 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v673))) = uint8(v939)
	v942 = v939
	goto L143
L142:
	;
	v942 = int32(0)
	goto L143
L143:
	;
	if v907 <= v942 {
		goto L97
	} else {
		goto L144
	}
L144:
	;
	v950 = v942
	goto L145
L145:
	;
	v982 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34+int32(_a_F_heap_page_prune_and_freeze_13)+v950<<(uint(int32(1))%32)))))
	F_heap_prune_record_unchanged_lp_normal(m, v34+int32(28), v982)
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L9
	} else {
		goto L147
	}
L146:
	;
	goto L97
L147:
	;
	v986 = v950 + int32(1)
	if v986 != v907 {
		v950 = v986
		goto L145
	} else {
		goto L148
	}
L148:
	;
	goto L146
L149:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v681)))
	v990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+32)))
	v991 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v673))) = uint8(v991)
	v994 = v989 & int32(_a_F_heap_page_prune_and_freeze_6)
	if v990 == v991 {
		goto L154
	} else {
		goto L155
	}
L150:
	;
	goto L151
L151:
	;
	v1185 = int32(1)
	v1188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34+int32(_a_F_heap_page_prune_and_freeze_13)+v908<<(uint(v1185)%32)))))
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v681)))
	*(*uint8)(unsafe.Add(mBase, uint32(v673))) = uint8(v1185)
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v34)+64))
	v1193 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v629+v1192<<(uint(v1193)%32)))) = uint16(v672)
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v34)+64))
	*(*uint16)(unsafe.Add(mBase, uint32(v629+v1197<<(uint(v1193)%32))+2)) = uint16(v1188)
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v34)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+64)) = v1202 + v1185
	if v1189&int32(_a_F_heap_page_prune_and_freeze_6) == int32(_a_F_heap_page_prune_and_freeze_7) {
		goto L171
	} else {
		goto L172
	}
L152:
	;
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v1035 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1034)+10)))
	if v1035&int32(4) != 0 {
		goto L159
	} else {
		goto L160
	}
L153:
	;
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[15])))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[15]))) = v1029 + int32(1)
	goto L152
L154:
	;
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	v998 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v627+v997<<(uint(v998)%32)))) = uint16(v672)
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+72)) = v1002 + v998
	if v994 == int32(_a_F_heap_page_prune_and_freeze_7) {
		goto L153
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v34)+68))
	v1009 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v625+v1008<<(uint(v1009)%32)))) = uint16(v672)
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v34)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+68)) = v1013 + v1009
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[13])))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[13]))) = v1017 + v1009
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[17])))
	*(*uint16)(unsafe.Add(mBase, uint32(v1021+v1017<<(uint(v1009)%32)))) = uint16(v672)
	if v994 != int32(_a_F_heap_page_prune_and_freeze_7) {
		goto L152
	} else {
		goto L158
	}
L157:
	;
	goto L152
L158:
	;
	goto L153
L159:
	;
	F_heap_page_fix_vm_corruption(m, v34+int32(28), v672, int32(1))
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L9
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	if v907 < int32(2) {
		goto L97
	} else {
		goto L163
	}
L162:
	;
	goto L161
L163:
	;
	v1045 = int32(1)
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	if v907 != int32(2) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v1049 = int32(1)
	v1050 = v907 - v1049
	v1062 = v1045
	v1063 = int32(0)
	v1064 = v1046
	goto L167
L165:
	;
	v1137 = v1045
	v1139 = v1046
	goto L166
L166:
	;
	v1164 = int32(1)
	v1167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34+int32(_a_F_heap_page_prune_and_freeze_13)+v1137<<(uint(v1164)%32)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v635+v1167))) = uint8(v1164)
	*(*uint16)(unsafe.Add(mBase, uint32(v627+v1139<<(uint(v1164)%32)))) = uint16(v1167)
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+72)) = v1175 + v1164
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[15])))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[15]))) = v1179 + v1164
	goto L97
L167:
	;
	v1089 = int32(1)
	v1091 = v34 + int32(_a_F_heap_page_prune_and_freeze_13) + v1062<<(uint(v1089)%32)
	v1092 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1091))))
	*(*uint8)(unsafe.Add(mBase, uint32(v635+v1092))) = uint8(v1089)
	*(*uint16)(unsafe.Add(mBase, uint32(v627+v1064<<(uint(v1089)%32)))) = uint16(v1092)
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[15])))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[15]))) = v1100 + v1089
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	v1106 = v1104 + v1089
	*(*int32)(unsafe.Add(mBase, uint32(v34)+72)) = v1106
	v1108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1091)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v635+v1108))) = uint8(v1089)
	*(*uint16)(unsafe.Add(mBase, uint32(v627+v1106<<(uint(v1089)%32)))) = uint16(v1108)
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	v1118 = v1116 + v1089
	*(*int32)(unsafe.Add(mBase, uint32(v34)+72)) = v1118
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[15])))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[15]))) = v1120 + v1089
	v1124 = int32(2)
	v1125 = v1062 + v1124
	v1127 = v1063 + v1124
	if v1127 != v1050&int32(-2) {
		v1062 = v1125
		v1063 = v1127
		v1064 = v1118
		goto L167
	} else {
		goto L169
	}
L168:
	;
	if v1050&v1049 == int32(0) {
		goto L97
	} else {
		goto L170
	}
L169:
	;
	goto L168
L170:
	;
	v1137 = v1125
	v1139 = v1118
	goto L166
L171:
	;
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[15])))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[15]))) = v1210 + int32(1)
	goto L173
L172:
	;
	goto L173
L173:
	;
	v1214 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[31]))) = uint8(v1214)
	if v908 < int32(2) {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	if v907 <= v908 {
		goto L97
	} else {
		goto L183
	}
L175:
	;
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	if v908 != int32(2) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v1222 = int32(1)
	v1223 = v908 - v1222
	v1236 = v1222
	v1239 = v1219
	v1240 = int32(0)
	goto L179
L177:
	;
	v1311 = int32(1)
	v1314 = v1219
	goto L178
L178:
	;
	v1338 = int32(1)
	v1341 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34+int32(_a_F_heap_page_prune_and_freeze_13)+v1311<<(uint(v1338)%32)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v635+v1341))) = uint8(v1338)
	*(*uint16)(unsafe.Add(mBase, uint32(v627+v1314<<(uint(v1338)%32)))) = uint16(v1341)
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+72)) = v1349 + v1338
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[15])))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[15]))) = v1353 + v1338
	goto L174
L179:
	;
	v1263 = int32(1)
	v1265 = v34 + int32(_a_F_heap_page_prune_and_freeze_13) + v1236<<(uint(v1263)%32)
	v1266 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1265))))
	*(*uint8)(unsafe.Add(mBase, uint32(v635+v1266))) = uint8(v1263)
	*(*uint16)(unsafe.Add(mBase, uint32(v627+v1239<<(uint(v1263)%32)))) = uint16(v1266)
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[15])))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[15]))) = v1274 + v1263
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	v1280 = v1278 + v1263
	*(*int32)(unsafe.Add(mBase, uint32(v34)+72)) = v1280
	v1282 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1265)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v635+v1282))) = uint8(v1263)
	*(*uint16)(unsafe.Add(mBase, uint32(v627+v1280<<(uint(v1263)%32)))) = uint16(v1282)
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	v1292 = v1290 + v1263
	*(*int32)(unsafe.Add(mBase, uint32(v34)+72)) = v1292
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[15])))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[15]))) = v1294 + v1263
	v1298 = int32(2)
	v1299 = v1236 + v1298
	v1301 = v1240 + v1298
	if v1301 != v1223&int32(-2) {
		v1236 = v1299
		v1239 = v1292
		v1240 = v1301
		goto L179
	} else {
		goto L181
	}
L180:
	;
	if v1223&v1222 == int32(0) {
		goto L174
	} else {
		goto L182
	}
L181:
	;
	goto L180
L182:
	;
	v1311 = v1299
	v1314 = v1292
	goto L178
L183:
	;
	v1397 = v908
	goto L184
L184:
	;
	v1427 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34+int32(_a_F_heap_page_prune_and_freeze_13)+v1397<<(uint(int32(1))%32)))))
	F_heap_prune_record_unchanged_lp_normal(m, v34+int32(28), v1427)
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L9
	} else {
		goto L186
	}
L185:
	;
	goto L97
L186:
	;
	v1431 = v1397 + int32(1)
	if v1431 != v907 {
		v1397 = v1431
		goto L184
	} else {
		goto L187
	}
L187:
	;
	goto L185
L188:
	;
	goto L96
L189:
	;
	v1522 = v1501
	goto L192
L190:
	;
	goto L191
L191:
	;
	v1652 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v1652)
	v1654 = int32(1)
	v1655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[19]))))
	if v1655 != v1654 {
		goto L210
	} else {
		goto L211
	}
L192:
	;
	v1550 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34+int32(_a_F_heap_page_prune_and_freeze_14)+v1522<<(uint(int32(1))%32)))))
	v1551 = v34 + int32(_a_F_heap_page_prune_and_freeze_9) + v1550
	v1552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1551))))
	if v1552 != 0 {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	goto L191
L194:
	;
	if int32(0) < v1522 {
		v1522 = v1522 - int32(1)
		goto L192
	} else {
		goto L209
	}
L195:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v1550)
	v1555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1550+(v34+int32(_a_F_heap_page_prune_and_freeze_10))))))
	if v1555 == int32(0) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v1561 = *(*int32)(unsafe.Add(mBase, uint32(v393+int32(20)+v1550<<(uint(int32(2))%32))))
	v1564 = v393 + v1561&int32(_a_F_heap_page_prune_and_freeze_11)
	v1565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1564)+19)))
	if v1565&int32(64) == int32(0) {
		goto L200
	} else {
		goto L201
	}
L197:
	;
	goto L198
L198:
	;
	F_heap_prune_record_unchanged_lp_normal(m, v34+int32(28), v1550)
	mBase = m.M
	v1614 = m.ExcPending
	if v1614 != 0 {
		goto L9
	} else {
		goto L208
	}
L199:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1598 = m.ExcPending
	if v1598 != 0 {
		goto L9
	} else {
		goto L205
	}
L200:
	;
	F_HeapTupleHeaderAdvanceConflictHorizon(m, v1564, v34+int32(60))
	mBase = m.M
	v1579 = m.ExcPending
	if v1579 != 0 {
		goto L9
	} else {
		goto L204
	}
L201:
	;
	v1570 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1564)+20)))
	if v1570&int32(2048) != 0 {
		goto L200
	} else {
		goto L202
	}
L202:
	;
	if v1570&int32(768) != int32(512) {
		goto L199
	} else {
		goto L203
	}
L203:
	;
	goto L200
L204:
	;
	v1580 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1551))) = uint8(v1580)
	v1582 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	*(*uint16)(unsafe.Add(mBase, uint32(v34+int32(1826)+v1582<<(uint(v1580)%32)))) = uint16(v1550)
	v1587 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+72)) = v1587 + v1580
	v1591 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[15])))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[15]))) = v1591 + v1580
	goto L194
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v1550
	*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v395
	F_errmsg_internal(m, int32(_a_F_heap_page_prune_and_freeze_21), v34+int32(16))
	mBase = m.M
	v1605 = m.ExcPending
	if v1605 != 0 {
		goto L9
	} else {
		goto L206
	}
L206:
	;
	F_errfinish(m, int32(_a_F_heap_page_prune_and_freeze_17), int32(722), int32(_a_F_heap_page_prune_and_freeze_22))
	mBase = m.M
	v1610 = m.ExcPending
	if v1610 != 0 {
		goto L9
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
	goto L194
L209:
	;
	goto L193
L210:
	;
	v1669 = *(*int32)(unsafe.Add(mBase, uint32(v34)+64))
	if int32(0) < v1669 {
		v1678 = v1654
		goto L215
	} else {
		goto L216
	}
L211:
	;
	v1658 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[18])))
	if base.Ui32(v1658) < base.Ui32(int32(3)) {
		goto L210
	} else {
		goto L212
	}
L212:
	;
	v1661 = *(*int32)(unsafe.Add(mBase, uint32(v34)+28))
	v1662 = F_GlobalVisTestXidConsideredRunning(m, v1661, v1658)
	mBase = m.M
	v1663 = m.ExcPending
	if v1663 != 0 {
		goto L9
	} else {
		goto L213
	}
L213:
	;
	if v1662 == int32(0) {
		goto L210
	} else {
		goto L214
	}
L214:
	;
	v1666 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[19]))) = uint16(v1666)
	goto L210
L215:
	;
	v1680 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(v1680)+20))
	v1682 = *(*int32)(unsafe.Add(mBase, uint32(v34)+56))
	if v1681 == v1682 {
		goto L218
	} else {
		goto L219
	}
L216:
	;
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(v34)+68))
	if int32(0) < v1672 {
		v1678 = v1654
		goto L215
	} else {
		goto L217
	}
L217:
	;
	v1675 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	v1678 = base.B2i32(int32(0) < v1675)
	goto L215
L218:
	;
	v1684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1680)+10)))
	v1689 = int32(base.Ui32(v1684&int32(2)) >> (uint(int32(1)) % 32))
	goto L220
L219:
	;
	v1689 = int32(1)
	goto L220
L220:
	;
	v1690 = int32(0)
	v1691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+33)))
	if v1691 != int32(1) {
		v2015 = v1690
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v2036 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[13])))
	if int32(0) < v2036 {
		goto L301
	} else {
		goto L302
	}
L222:
	;
	v1694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[8]))))
	if v1694 != 0 {
		goto L225
	} else {
		goto L226
	}
L223:
	;
	if v1997 <= int32(0) {
		v2015 = v1690
		goto L221
	} else {
		goto L299
	}
L224:
	;
	v1995 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
	v1997 = v1995
	goto L223
L225:
	;
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
	v1835 = m.G0
	v1837 = v1835 - int32(32)
	m.G0 = v1837
	if v1831 < int32(0) {
		goto L269
	} else {
		goto L270
	}
L226:
	;
	v1695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[20]))))
	v1698 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
	if base.B2i32(v1695 != int32(1))|base.B2i32(v1698 <= int32(0)) != 0 {
		v1997 = v1698
		goto L223
	} else {
		goto L227
	}
L227:
	;
	v1702 = *(*int32)(unsafe.Add(mBase, uint32(v34)+40))
	v1703 = *(*int32)(unsafe.Add(mBase, uint32(v1702)+48))
	v1704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1703)+118)))
	if v1704 != int32(112) {
		goto L224
	} else {
		goto L228
	}
L228:
	;
	v1708 = *(*int64)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[0]))
	v1710 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[32]))
	if v1710 <= int32(0) {
		goto L230
	} else {
		goto L231
	}
L229:
	;
	if v1678 != 0 {
		goto L237
	} else {
		goto L238
	}
L230:
	;
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(v1702)+32))
	if v1713 != 0 {
		goto L224
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	if v57 != v1708 {
		goto L225
	} else {
		goto L236
	}
L233:
	;
	v1714 = *(*int32)(unsafe.Add(mBase, uint32(v1702)+40))
	if base.B2i32(v1714 == int32(0))&base.B2i32(v57 == v1708) != 0 {
		goto L229
	} else {
		goto L234
	}
L234:
	;
	if v1714 == int32(0) {
		goto L225
	} else {
		goto L235
	}
L235:
	;
	goto L224
L236:
	;
	goto L229
L237:
	;
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	v1726 = m.G0
	v1728 = v1726 - int32(16)
	m.G0 = v1728
	F_GetFullPageWriteInfo(m, v1728+int32(8), v1728+int32(7))
	mBase = m.M
	if v1723 < int32(0) {
		goto L242
	} else {
		goto L243
	}
L238:
	;
	goto L239
L239:
	;
	if v1689 == int32(0) {
		goto L224
	} else {
		goto L251
	}
L240:
	;
	if v1763 == int32(0) {
		goto L224
	} else {
		goto L250
	}
L241:
	;
	v1753 = int32(1)
	v1754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1728)+7)))
	if v1754 == v1753 {
		goto L246
	} else {
		goto L247
	}
L242:
	;
	v1738 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[5]))
	v1744 = *(*int32)(unsafe.Add(mBase, uint32(v1738+(v1723^int32(-1))<<(uint(int32(2))%32))))
	v1752 = v1744
	goto L241
L243:
	;
	goto L244
L244:
	;
	v1746 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[6]))
	v1752 = v1746 + v1723<<(uint(int32(13))%32) + int32(-8192)
	goto L241
L245:
	;
	m.G0 = v1728 + int32(16)
	goto L240
L246:
	;
	v1757 = *(*int64)(unsafe.Add(mBase, uint32(v1728)+8))
	v1758 = *(*int64)(unsafe.Add(mBase, uint32(v1752)))
	if base.Ui64(base.I64_rotl(v1758, int64(32))) <= base.Ui64(v1757) {
		v1763 = v1753
		goto L245
	} else {
		goto L249
	}
L247:
	;
	goto L248
L248:
	;
	v1763 = int32(0)
	goto L245
L249:
	;
	goto L248
L250:
	;
	goto L225
L251:
	;
	v1772 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[33])))
	if v1772 == int32(0) {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v1776 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[34]))
	v1777 = *(*int32)(unsafe.Add(mBase, uint32(v1776)+268))
	goto L255
L253:
	;
	goto L254
L254:
	;
	v1782 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	v1785 = m.G0
	v1787 = v1785 - int32(16)
	m.G0 = v1787
	F_GetFullPageWriteInfo(m, v1787+int32(8), v1787+int32(7))
	mBase = m.M
	if v1782 < int32(0) {
		goto L259
	} else {
		goto L260
	}
L255:
	;
	if base.B2i32(v1777 != int32(0)) == int32(0) {
		goto L224
	} else {
		goto L256
	}
L256:
	;
	goto L254
L257:
	;
	if v1822 == int32(0) {
		goto L224
	} else {
		goto L267
	}
L258:
	;
	v1812 = int32(1)
	v1813 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1787)+7)))
	if v1813 == v1812 {
		goto L263
	} else {
		goto L264
	}
L259:
	;
	v1797 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[5]))
	v1803 = *(*int32)(unsafe.Add(mBase, uint32(v1797+(v1782^int32(-1))<<(uint(int32(2))%32))))
	v1811 = v1803
	goto L258
L260:
	;
	goto L261
L261:
	;
	v1805 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[6]))
	v1811 = v1805 + v1782<<(uint(int32(13))%32) + int32(-8192)
	goto L258
L262:
	;
	m.G0 = v1787 + int32(16)
	goto L257
L263:
	;
	v1816 = *(*int64)(unsafe.Add(mBase, uint32(v1787)+8))
	v1817 = *(*int64)(unsafe.Add(mBase, uint32(v1811)))
	if base.Ui64(base.I64_rotl(v1817, int64(32))) <= base.Ui64(v1816) {
		v1822 = v1812
		goto L262
	} else {
		goto L266
	}
L264:
	;
	goto L265
L265:
	;
	v1822 = int32(0)
	goto L262
L266:
	;
	goto L265
L267:
	;
	goto L225
L268:
	;
	if int32(0) < v1834 {
		goto L275
	} else {
		goto L276
	}
L269:
	;
	v1842 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[5]))
	v1848 = *(*int32)(unsafe.Add(mBase, uint32(v1842+(v1831^int32(-1))<<(uint(int32(2))%32))))
	v1856 = v1848
	goto L268
L270:
	;
	goto L271
L271:
	;
	v1850 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[6]))
	v1856 = v1850 + v1831<<(uint(int32(13))%32) + int32(-8192)
	goto L268
L272:
	;
	v2015 = int32(1)
	goto L221
L273:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1979 = m.ExcPending
	if v1979 != 0 {
		goto L9
	} else {
		goto L295
	}
L274:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1961 = m.ExcPending
	if v1961 != 0 {
		goto L9
	} else {
		goto L291
	}
L275:
	;
	v1870 = int32(0)
	goto L278
L276:
	;
	goto L277
L277:
	;
	m.G0 = v1837 + int32(32)
	goto L272
L278:
	;
	v1895 = v34 + int32(2408) + v1870*int32(12)
	v1896 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1895)+10)))
	v1900 = *(*int32)(unsafe.Add(mBase, uint32(v1856+int32(20)+v1896<<(uint(int32(2))%32))))
	v1903 = v1856 + v1900&int32(_a_F_heap_page_prune_and_freeze_11)
	v1904 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1895)+9)))
	if v1904&int32(1) != 0 {
		goto L280
	} else {
		goto L281
	}
L279:
	;
	goto L277
L280:
	;
	v1907 = *(*int32)(unsafe.Add(mBase, uint32(v1903)))
	v1908 = F_TransactionIdDidCommit(m, v1907)
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		goto L9
	} else {
		goto L283
	}
L281:
	;
	v1914 = v1904
	goto L282
L282:
	;
	if v1914&int32(2) != 0 {
		goto L285
	} else {
		goto L286
	}
L283:
	;
	if v1908 == int32(0) {
		goto L274
	} else {
		goto L284
	}
L284:
	;
	v1912 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1895)+9)))
	v1914 = v1912
	goto L282
L285:
	;
	v1917 = *(*int32)(unsafe.Add(mBase, uint32(v1903)+4))
	v1918 = F_TransactionIdDidCommit(m, v1917)
	mBase = m.M
	v1919 = m.ExcPending
	if v1919 != 0 {
		goto L9
	} else {
		goto L288
	}
L286:
	;
	goto L287
L287:
	;
	v1922 = v1870 + int32(1)
	if v1922 != v1834 {
		v1870 = v1922
		goto L278
	} else {
		goto L290
	}
L288:
	;
	if v1918 != 0 {
		goto L273
	} else {
		goto L289
	}
L289:
	;
	goto L287
L290:
	;
	goto L279
L291:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v1964 = m.ExcPending
	if v1964 != 0 {
		goto L9
	} else {
		goto L292
	}
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+16)) = v1907
	F_errmsg_internal(m, int32(_a_F_heap_page_prune_and_freeze_23), v1837+int32(16))
	mBase = m.M
	v1970 = m.ExcPending
	if v1970 != 0 {
		goto L9
	} else {
		goto L293
	}
L293:
	;
	F_errfinish(m, int32(_a_F_heap_page_prune_and_freeze_24), int32(_a_F_heap_page_prune_and_freeze_25), int32(_a_F_heap_page_prune_and_freeze_26))
	mBase = m.M
	v1975 = m.ExcPending
	if v1975 != 0 {
		goto L9
	} else {
		goto L294
	}
L294:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L295:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v1982 = m.ExcPending
	if v1982 != 0 {
		goto L9
	} else {
		goto L296
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1837))) = v1917
	F_errmsg_internal(m, int32(_a_F_heap_page_prune_and_freeze_27), v1837)
	mBase = m.M
	v1986 = m.ExcPending
	if v1986 != 0 {
		goto L9
	} else {
		goto L297
	}
L297:
	;
	F_errfinish(m, int32(_a_F_heap_page_prune_and_freeze_24), int32(_a_F_heap_page_prune_and_freeze_28), int32(_a_F_heap_page_prune_and_freeze_26))
	mBase = m.M
	v1991 = m.ExcPending
	if v1991 != 0 {
		goto L9
	} else {
		goto L298
	}
L298:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L299:
	;
	v2001 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+76)) = v2001
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[20]))) = uint8(v2001)
	v2015 = v1690
	goto L221
L300:
	;
	if v2015 != 0 {
		goto L333
	} else {
		goto L334
	}
L301:
	;
	v2145 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[19]))) = uint16(v2145)
	v2149 = v2145
	v2150 = v2145
	goto L300
L302:
	;
	v2039 = int32(0)
	v2041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+34)))
	if v2041 != int32(1) {
		v2149 = v2039
		v2150 = v2039
		goto L300
	} else {
		goto L303
	}
L303:
	;
	v2044 = int32(0)
	v2045 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[19]))))
	if v2045&int32(1) == v2044 {
		v2149 = v2039
		v2150 = v2044
		goto L300
	} else {
		goto L304
	}
L304:
	;
	v2050 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2051 = int32(0)
	if v2015|(v1678|base.B2i32(v2050 != v2051)) == v2051 {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v2057 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	if v2057 < int32(0) {
		goto L309
	} else {
		goto L310
	}
L306:
	;
	goto L307
L307:
	;
	v2131 = int32(1)
	v2132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[20]))))
	if v2132&v2131 != 0 {
		goto L324
	} else {
		goto L325
	}
L308:
	;
	v2075 = int64(0)
	v2078 = base.AtomicRmwCmpxchg64(m, v2074, int32(24), v2075, v2075)
	if int64(base.Ui64(v2078&int64(8388608))>>(uint(int64(23))%64)) == v2075 {
		goto L301
	} else {
		goto L312
	}
L309:
	;
	v2061 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[1]))
	v2074 = v2061 + (v2057^int32(-1))*int32(56)
	goto L308
L310:
	;
	goto L311
L311:
	;
	v2068 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[2]))
	v2069 = int32(56)
	v2074 = v2068 + v2057*v2069 - v2069
	goto L308
L312:
	;
	v2085 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	v2088 = m.G0
	v2090 = v2088 - int32(16)
	m.G0 = v2090
	F_GetFullPageWriteInfo(m, v2090+int32(8), v2090+int32(7))
	mBase = m.M
	if v2085 < int32(0) {
		goto L315
	} else {
		goto L316
	}
L313:
	;
	if v2125 != 0 {
		goto L301
	} else {
		goto L323
	}
L314:
	;
	v2115 = int32(1)
	v2116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2090)+7)))
	if v2116 == v2115 {
		goto L319
	} else {
		goto L320
	}
L315:
	;
	v2100 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[5]))
	v2106 = *(*int32)(unsafe.Add(mBase, uint32(v2100+(v2085^int32(-1))<<(uint(int32(2))%32))))
	v2114 = v2106
	goto L314
L316:
	;
	goto L317
L317:
	;
	v2108 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[6]))
	v2114 = v2108 + v2085<<(uint(int32(13))%32) + int32(-8192)
	goto L314
L318:
	;
	m.G0 = v2090 + int32(16)
	goto L313
L319:
	;
	v2119 = *(*int64)(unsafe.Add(mBase, uint32(v2090)+8))
	v2120 = *(*int64)(unsafe.Add(mBase, uint32(v2114)))
	if base.Ui64(base.I64_rotl(v2120, int64(32))) <= base.Ui64(v2119) {
		v2125 = v2115
		goto L318
	} else {
		goto L322
	}
L320:
	;
	goto L321
L321:
	;
	v2125 = int32(0)
	goto L318
L322:
	;
	goto L321
L323:
	;
	goto L307
L324:
	;
	v2135 = int32(3)
	goto L326
L325:
	;
	v2135 = v2131
	goto L326
L326:
	;
	v2137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[11]))))
	v2138 = base.B2i32(v2135 != v2137)
	if v2135 != v2137 {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	v2139 = v2135
	goto L329
L328:
	;
	v2139 = int32(0)
	goto L329
L329:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[3]))) = uint8(v2139)
	v2141 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[18])))
	if v2135 != v2137 {
		goto L330
	} else {
		goto L331
	}
L330:
	;
	v2143 = v2141
	goto L332
L331:
	;
	v2143 = int32(0)
	goto L332
L332:
	;
	v2149 = v2138
	v2150 = v2143
	goto L300
L333:
	;
	v2151 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[7])))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v2151))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v2150)) == int32(0) {
		goto L337
	} else {
		goto L338
	}
L334:
	;
	v2166 = v2150
	goto L335
L335:
	;
	if v1678 != 0 {
		goto L343
	} else {
		goto L344
	}
L336:
	;
	if v2163 != 0 {
		goto L340
	} else {
		goto L341
	}
L337:
	;
	v2163 = base.B2i32(base.Ui32(v2150) < base.Ui32(v2151))
	goto L336
L338:
	;
	goto L339
L339:
	;
	v2163 = base.B2i32(int32(0) < v2151-v2150)
	goto L336
L340:
	;
	v2164 = v2151
	goto L342
L341:
	;
	v2164 = v2150
	goto L342
L342:
	;
	v2166 = v2164
	goto L335
L343:
	;
	v2168 = *(*int32)(unsafe.Add(mBase, uint32(v34)+60))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v2168))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v2166)) == int32(0) {
		goto L347
	} else {
		goto L348
	}
L344:
	;
	v2183 = v2166
	goto L345
L345:
	;
	if v2149 != 0 {
		goto L353
	} else {
		goto L354
	}
L346:
	;
	if v2180 != 0 {
		goto L350
	} else {
		goto L351
	}
L347:
	;
	v2180 = base.B2i32(base.Ui32(v2166) < base.Ui32(v2168))
	goto L346
L348:
	;
	goto L349
L349:
	;
	v2180 = base.B2i32(int32(0) < v2168-v2166)
	goto L346
L350:
	;
	v2181 = v2168
	goto L352
L351:
	;
	v2181 = v2166
	goto L352
L352:
	;
	v2183 = v2181
	goto L345
L353:
	;
	v2185 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[4])))
	F_LockBufferInternal(m, v2185, int32(3))
	mBase = m.M
	v2188 = m.ExcPending
	if v2188 != 0 {
		goto L9
	} else {
		goto L356
	}
L354:
	;
	goto L355
L355:
	;
	v2189 = int32(_a_F_heap_page_prune_and_freeze_29)
	v2191 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[35]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[35])) = v2191 + int32(1)
	if v1689 == int32(0) {
		goto L357
	} else {
		goto L358
	}
L356:
	;
	goto L355
L357:
	;
	if v1678|v2015|v2149 != int32(1) {
		goto L361
	} else {
		goto L362
	}
L358:
	;
	v2197 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v2198 = *(*int32)(unsafe.Add(mBase, uint32(v34)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v2197)+20)) = v2198
	v2200 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v2201 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2200)+10)))
	v2203 = v2201 & int32(_a_F_heap_page_prune_and_freeze_30)
	*(*uint16)(unsafe.Add(mBase, uint32(v2200)+10)) = uint16(v2203)
	if (v1678|v2015|v2149)&int32(1) != 0 {
		goto L357
	} else {
		goto L359
	}
L359:
	;
	v2209 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	F_MarkBufferDirtyHint(m, v2209, int32(1))
	mBase = m.M
	v2212 = m.ExcPending
	if v2212 != 0 {
		goto L9
	} else {
		goto L360
	}
L360:
	;
	goto L357
L361:
	;
	v3016 = int32(_a_F_heap_page_prune_and_freeze_29)
	v3018 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[35]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[35])) = v3018 - int32(1)
	if v2149 != 0 {
		goto L441
	} else {
		goto L442
	}
L362:
	;
	if v1678 != 0 {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	v2218 = *(*int32)(unsafe.Add(mBase, uint32(v34)+64))
	v2219 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	if v2219 < int32(0) {
		goto L367
	} else {
		goto L368
	}
L364:
	;
	goto L365
L365:
	;
	if v2015 != 0 {
		goto L404
	} else {
		goto L405
	}
L366:
	;
	v2238 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	v2239 = *(*int32)(unsafe.Add(mBase, uint32(v34)+68))
	if v2218 <= int32(0) {
		goto L370
	} else {
		goto L371
	}
L367:
	;
	v2223 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[5]))
	v2229 = *(*int32)(unsafe.Add(mBase, uint32(v2223+(v2219^int32(-1))<<(uint(int32(2))%32))))
	v2237 = v2229
	goto L366
L368:
	;
	goto L369
L369:
	;
	v2231 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[6]))
	v2237 = v2231 + v2219<<(uint(int32(13))%32) + int32(-8192)
	goto L366
L370:
	;
	if v2239 <= int32(0) {
		goto L379
	} else {
		goto L380
	}
L371:
	;
	v2243 = v34 + int32(80)
	v2245 = v2237 + int32(20)
	if v2218 != int32(1) {
		goto L372
	} else {
		goto L373
	}
L372:
	;
	v2260 = v2243
	v2261 = int32(0)
	goto L375
L373:
	;
	v2318 = v2243
	goto L374
L374:
	;
	v2342 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2318))))
	v2346 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2318)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v2245+v2342<<(uint(int32(2))%32)))) = v2346&int32(_a_F_heap_page_prune_and_freeze_11) | int32(_a_F_heap_page_prune_and_freeze_15)
	goto L370
L375:
	;
	v2284 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2260))))
	v2285 = int32(2)
	v2288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2260)+2)))
	v2289 = int32(_a_F_heap_page_prune_and_freeze_11)
	v2291 = int32(_a_F_heap_page_prune_and_freeze_15)
	*(*int32)(unsafe.Add(mBase, uint32(v2245+v2284<<(uint(v2285)%32)))) = v2288&v2289 | v2291
	v2294 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2260)+4)))
	v2298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2260)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v2245+v2294<<(uint(v2285)%32)))) = v2298&v2289 | v2291
	v2305 = v2260 + int32(8)
	v2307 = v2261 + v2285
	if v2307 != v2218&int32(2147483646) {
		v2260 = v2305
		v2261 = v2307
		goto L375
	} else {
		goto L377
	}
L376:
	;
	if v2218&int32(1) == int32(0) {
		goto L370
	} else {
		goto L378
	}
L377:
	;
	goto L376
L378:
	;
	v2318 = v2305
	goto L374
L379:
	;
	if v2238 <= int32(0) {
		goto L391
	} else {
		goto L392
	}
L380:
	;
	v2386 = v34 + int32(1244)
	v2388 = v2239 & int32(3)
	v2390 = v2237 + int32(20)
	if base.Ui32(int32(4)) <= base.Ui32(v2239) {
		goto L381
	} else {
		goto L382
	}
L381:
	;
	v2403 = v2386
	v2404 = int32(0)
	goto L384
L382:
	;
	v2465 = v2386
	goto L383
L383:
	;
	v2497 = v2465
	v2498 = int32(0)
	goto L388
L384:
	;
	v2427 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2403))))
	v2428 = int32(2)
	v2431 = int32(_a_F_heap_page_prune_and_freeze_6)
	*(*int32)(unsafe.Add(mBase, uint32(v2390+v2427<<(uint(v2428)%32)))) = v2431
	v2433 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2403)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v2390+v2433<<(uint(v2428)%32)))) = v2431
	v2439 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2403)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v2390+v2439<<(uint(v2428)%32)))) = v2431
	v2445 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2403)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v2390+v2445<<(uint(v2428)%32)))) = v2431
	v2452 = v2403 + int32(8)
	v2454 = v2404 + int32(4)
	if v2454 != v2239&int32(2147483644) {
		v2403 = v2452
		v2404 = v2454
		goto L384
	} else {
		goto L386
	}
L385:
	;
	if v2388 == int32(0) {
		goto L379
	} else {
		goto L387
	}
L386:
	;
	goto L385
L387:
	;
	v2465 = v2452
	goto L383
L388:
	;
	v2521 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2497))))
	v2522 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v2390+v2521<<(uint(v2522)%32)))) = int32(_a_F_heap_page_prune_and_freeze_6)
	v2530 = v2498 + int32(1)
	if v2530 != v2388 {
		v2497 = v2497 + v2522
		v2498 = v2530
		goto L388
	} else {
		goto L390
	}
L389:
	;
	goto L379
L390:
	;
	goto L389
L391:
	;
	F_PageRepairFragmentation(m, v2237)
	mBase = m.M
	v2744 = m.ExcPending
	if v2744 != 0 {
		goto L9
	} else {
		goto L403
	}
L392:
	;
	v2566 = v34 + int32(1826)
	v2568 = v2238 & int32(3)
	v2570 = v2237 + int32(20)
	if base.Ui32(int32(4)) <= base.Ui32(v2238) {
		goto L393
	} else {
		goto L394
	}
L393:
	;
	v2583 = v2566
	v2584 = int32(0)
	goto L396
L394:
	;
	v2645 = v2566
	goto L395
L395:
	;
	v2677 = v2645
	v2678 = int32(0)
	goto L400
L396:
	;
	v2607 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2583))))
	v2608 = int32(2)
	v2611 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2570+v2607<<(uint(v2608)%32)))) = v2611
	v2613 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2583)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v2570+v2613<<(uint(v2608)%32)))) = v2611
	v2619 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2583)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v2570+v2619<<(uint(v2608)%32)))) = v2611
	v2625 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2583)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v2570+v2625<<(uint(v2608)%32)))) = v2611
	v2632 = v2583 + int32(8)
	v2634 = v2584 + int32(4)
	if v2634 != v2238&int32(2147483644) {
		v2583 = v2632
		v2584 = v2634
		goto L396
	} else {
		goto L398
	}
L397:
	;
	if v2568 == int32(0) {
		goto L391
	} else {
		goto L399
	}
L398:
	;
	goto L397
L399:
	;
	v2645 = v2632
	goto L395
L400:
	;
	v2701 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2677))))
	v2702 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v2570+v2701<<(uint(v2702)%32)))) = int32(0)
	v2710 = v2678 + int32(1)
	if v2710 != v2568 {
		v2677 = v2677 + v2702
		v2678 = v2710
		goto L400
	} else {
		goto L402
	}
L401:
	;
	goto L391
L402:
	;
	goto L401
L403:
	;
	goto L365
L404:
	;
	v2776 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
	v2777 = int32(0)
	v2778 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	if v2778 < v2777 {
		goto L408
	} else {
		goto L409
	}
L405:
	;
	goto L406
L406:
	;
	if v2149 != 0 {
		goto L423
	} else {
		goto L424
	}
L407:
	;
	if int32(0) < v2776 {
		goto L411
	} else {
		goto L412
	}
L408:
	;
	v2782 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[5]))
	v2788 = *(*int32)(unsafe.Add(mBase, uint32(v2782+(v2778^int32(-1))<<(uint(int32(2))%32))))
	v2796 = v2788
	goto L407
L409:
	;
	goto L410
L410:
	;
	v2790 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[6]))
	v2796 = v2790 + v2778<<(uint(int32(13))%32) + int32(-8192)
	goto L407
L411:
	;
	v2811 = v2777
	goto L414
L412:
	;
	goto L413
L413:
	;
	goto L406
L414:
	;
	v2836 = v34 + int32(2408) + v2811*int32(12)
	v2837 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2836)+10)))
	v2838 = int32(2)
	v2841 = *(*int32)(unsafe.Add(mBase, uint32(v2796+int32(20)+v2837<<(uint(v2838)%32))))
	v2844 = v2796 + v2841&int32(_a_F_heap_page_prune_and_freeze_11)
	v2845 = *(*int32)(unsafe.Add(mBase, uint32(v2836)))
	*(*int32)(unsafe.Add(mBase, uint32(v2844)+4)) = v2845
	v2847 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2836)+8)))
	if v2847&v2838 != 0 {
		goto L416
	} else {
		goto L417
	}
L415:
	;
	goto L413
L416:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2844)+8)) = int32(2)
	v2852 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2836)+8)))
	v2853 = v2852
	goto L418
L417:
	;
	v2853 = v2847
	goto L418
L418:
	;
	if v2853&int32(4) != 0 {
		goto L419
	} else {
		goto L420
	}
L419:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2844)+8)) = int32(0)
	goto L421
L420:
	;
	goto L421
L421:
	;
	v2858 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2836)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v2844)+20)) = uint16(v2858)
	v2860 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2836)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v2844)+18)) = uint16(v2860)
	v2863 = v2811 + int32(1)
	if v2863 != v2776 {
		v2811 = v2863
		goto L414
	} else {
		goto L422
	}
L422:
	;
	goto L415
L423:
	;
	v2927 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v2928 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2927)+10)))
	v2930 = v2928 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v2927)+10)) = uint16(v2930)
	v2932 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v2932)+20)) = int32(0)
	v2935 = *(*int32)(unsafe.Add(mBase, uint32(v34)+40))
	v2936 = *(*int64)(unsafe.Add(mBase, uint32(v2935)))
	*(*int64)(unsafe.Add(mBase, uint32(v34))) = v2936
	v2938 = *(*int32)(unsafe.Add(mBase, uint32(v2935)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v2938
	v2940 = *(*int32)(unsafe.Add(mBase, uint32(v34)+44))
	v2941 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[4])))
	v2942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[3]))))
	v2943 = F_visibilitymap_set(m, v2940, v2941, v2942)
	mBase = m.M
	v2944 = m.ExcPending
	if v2944 != 0 {
		goto L9
	} else {
		goto L426
	}
L424:
	;
	goto L425
L425:
	;
	v2946 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	F_MarkBufferDirty(m, v2946)
	mBase = m.M
	v2948 = m.ExcPending
	if v2948 != 0 {
		goto L9
	} else {
		goto L427
	}
L426:
	;
	goto L425
L427:
	;
	v2949 = *(*int32)(unsafe.Add(mBase, uint32(v34)+40))
	v2950 = *(*int32)(unsafe.Add(mBase, uint32(v2949)+48))
	v2951 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2950)+118)))
	if v2951 != int32(112) {
		goto L361
	} else {
		goto L428
	}
L428:
	;
	v2955 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_and_freeze[32]))
	if v2955 <= int32(0) {
		goto L429
	} else {
		goto L430
	}
L429:
	;
	v2958 = *(*int32)(unsafe.Add(mBase, uint32(v2949)+32))
	if v2958 != 0 {
		goto L361
	} else {
		goto L432
	}
L430:
	;
	goto L431
L431:
	;
	v2960 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	v2961 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[4])))
	if v2149 != 0 {
		goto L434
	} else {
		goto L435
	}
L432:
	;
	v2959 = *(*int32)(unsafe.Add(mBase, uint32(v2949)+40))
	if v2959 != 0 {
		goto L361
	} else {
		goto L433
	}
L433:
	;
	goto L431
L434:
	;
	v2963 = v2961
	goto L436
L435:
	;
	v2963 = int32(0)
	goto L436
L436:
	;
	v2964 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[3]))))
	if v2149 != 0 {
		goto L437
	} else {
		goto L438
	}
L437:
	;
	v2966 = v2964
	goto L439
L438:
	;
	v2966 = int32(0)
	goto L439
L439:
	;
	v2970 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2973 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
	v2976 = *(*int32)(unsafe.Add(mBase, uint32(v34)+64))
	v2979 = *(*int32)(unsafe.Add(mBase, uint32(v34)+68))
	v2982 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	F_log_heap_prune_and_freeze(m, v2949, v2960, v2963, v2966&int32(255), v2183, int32(1), v2970, v34+int32(2408), v2973, v34+int32(80), v2976, v34+int32(1244), v2979, v34+int32(1826), v2982)
	mBase = m.M
	v2984 = m.ExcPending
	if v2984 != 0 {
		goto L9
	} else {
		goto L440
	}
L440:
	;
	goto L361
L441:
	;
	v3022 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[4])))
	F_UnlockBuffer(m, v3022)
	mBase = m.M
	v3024 = m.ExcPending
	if v3024 != 0 {
		goto L9
	} else {
		goto L444
	}
L442:
	;
	goto L443
L443:
	;
	v3025 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[15])))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v3025
	v3027 = *(*int32)(unsafe.Add(mBase, uint32(v34)+68))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v3027
	v3029 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v3029
	v3031 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[24])))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v3031
	v3033 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[36])))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v3033
	v3035 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[31]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+23)) = uint8(v3035)
	v3037 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[13])))
	v3038 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+22)) = uint8(v3038)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v3037
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+20)) = uint16(v3038)
	if v2149 == v3038 {
		goto L445
	} else {
		goto L446
	}
L444:
	;
	goto L443
L445:
	;
	v3067 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+33)))
	if v3067 != int32(1) {
		goto L19
	} else {
		goto L453
	}
L446:
	;
	v3045 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[11]))))
	if v3045&int32(1) == int32(0) {
		goto L447
	} else {
		goto L448
	}
L447:
	;
	v3050 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v3050)
	v3052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[20]))))
	if v3052 != v3050 {
		goto L445
	} else {
		goto L450
	}
L448:
	;
	goto L449
L449:
	;
	if v3045&int32(2) != 0 {
		goto L445
	} else {
		goto L451
	}
L450:
	;
	v3055 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)) = uint8(v3055)
	goto L445
L451:
	;
	v3059 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[20]))))
	if v3059&int32(1) == int32(0) {
		goto L445
	} else {
		goto L452
	}
L452:
	;
	v3064 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+22)) = uint8(v3064)
	goto L445
L453:
	;
	if int32(0) < v3029 {
		goto L454
	} else {
		goto L455
	}
L454:
	;
	v3072 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[12])))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v3072
	v3074 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[22])))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v3074
	goto L19
L455:
	;
	goto L456
L456:
	;
	v3076 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[21])))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v3076
	v3078 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_heap_page_prune_and_freeze[14])))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v3078
	goto L19
}
func F_heap_redo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v272 int32
	_ = v272
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v355 int64
	_ = v355
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int64
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v381 int64
	_ = v381
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v472 int32
	_ = v472
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v512 int32
	_ = v512
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
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int64
	_ = v560
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v598 int32
	_ = v598
	var v605 int32
	_ = v605
	var v613 int32
	_ = v613
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v652 int64
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v664 int32
	_ = v664
	var v665 int64
	_ = v665
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
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
	var v712 int32
	_ = v712
	var v719 int32
	_ = v719
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v757 int32
	_ = v757
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v829 int32
	_ = v829
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int64
	_ = v835
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v878 int32
	_ = v878
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v904 int32
	_ = v904
	var v911 int32
	_ = v911
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v968 int32
	_ = v968
	var v972 int32
	_ = v972
	var v977 int32
	_ = v977
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v990 int32
	_ = v990
	var v995 int32
	_ = v995
	var v999 int32
	_ = v999
	var v1004 int32
	_ = v1004
	var v1008 int32
	_ = v1008
	var v1012 int32
	_ = v1012
	var v1017 int32
	_ = v1017
	var v1022 int32
	_ = v1022
	var v1026 int32
	_ = v1026
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1039 int32
	_ = v1039
	var v1044 int32
	_ = v1044
	var v1049 int32
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1058 int32
	_ = v1058
	var v1062 int32
	_ = v1062
	var v1066 int32
	_ = v1066
	var v1071 int32
	_ = v1071
	var v1076 int32
	_ = v1076
	var v1080 int32
	_ = v1080
	var v1085 int32
	_ = v1085
	var v1089 int32
	_ = v1089
	var v1093 int32
	_ = v1093
	var v1098 int32
	_ = v1098
	var v1102 int32
	_ = v1102
	var v1106 int32
	_ = v1106
	var v1111 int32
	_ = v1111
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(_a_F_heap_redo_0)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+48)))
	switch int32(base.Ui32(v20)>>(uint(int32(4))%32))&int32(7) - int32(1) {
	case 0:
		goto L18
	case 1:
		goto L17
	case 2:
		goto L12
	case 3:
		goto L16
	case 4:
		goto L15
	case 5:
		goto L14
	case 6:
		goto L13
	default:
		goto L19
	}
L1:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L20
	} else {
		goto L279
	}
L2:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L20
	} else {
		goto L276
	}
L3:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L20
	} else {
		goto L273
	}
L4:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L20
	} else {
		goto L270
	}
L5:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L20
	} else {
		goto L267
	}
L6:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L20
	} else {
		goto L264
	}
L7:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L20
	} else {
		goto L261
	}
L8:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L20
	} else {
		goto L258
	}
L9:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L20
	} else {
		goto L255
	}
L10:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L20
	} else {
		goto L252
	}
L11:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L20
	} else {
		goto L249
	}
L12:
	;
	m.G0 = v17 + int32(_a_F_heap_redo_0)
	return
L13:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	v835 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v839 = F_XLogReadBufferForRedo(m, l0, int32(0), v17+int32(92))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L20
	} else {
		goto L214
	}
L14:
	;
	v652 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v653)+7)))
	if v654&int32(1) != 0 {
		goto L179
	} else {
		goto L180
	}
L15:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	v560 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v564 = F_XLogReadBufferForRedo(m, l0, int32(0), v17+int32(92))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L20
	} else {
		goto L158
	}
L16:
	;
	F_heap_xlog_update(m, l0, int32(1))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L20
	} else {
		goto L157
	}
L17:
	;
	F_heap_xlog_update(m, l0, int32(0))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L20
	} else {
		goto L156
	}
L18:
	;
	v364 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	v366 = int32(0)
	F_XLogRecGetBlockTag(m, l0, v366, v17+int32(92), v366, v17+int32(_a_F_heap_redo_1))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L20
	} else {
		goto L109
	}
L19:
	;
	v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	v29 = int32(0)
	F_XLogRecGetBlockTag(m, l0, v29, v17+int32(80), v29, v17+int32(76))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return
L21:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28))))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+2)))
	if v39&int32(1) != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v17)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v42
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v17)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v44
	F_heap_xlog_vm_clear(m, l0, v17+int32(24), v37, int32(3))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v52 = int32(*(*int8)(unsafe.Add(mBase, uint32(v51)+48)))
	if v52 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	goto L24
L26:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_heap_redo[0])))
	if v347 != 0 {
		goto L103
	} else {
		goto L104
	}
L27:
	;
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28))))
	if v133 < int32(0) {
		goto L50
	} else {
		goto L51
	}
L28:
	;
	v56 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L20
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v130 = F_XLogReadBufferForRedo(m, l0, int32(0), v17+int32(_a_F_heap_redo_1))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L20
	} else {
		goto L47
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_heap_redo[0]))) = v56
	if v56 < int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v77 = int32(_a_F_heap_redo_2)
	v78 = int32(0)
	if v78|(v76&int32(3)|int32(1)) == v78 {
		goto L38
	} else {
		goto L39
	}
L33:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[1]))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v62+(v56^int32(-1))<<(uint(int32(2))%32))))
	v76 = v68
	goto L32
L34:
	;
	goto L35
L35:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[2]))
	v76 = v70 + v56<<(uint(int32(13))%32) + int32(-8192)
	goto L32
L36:
	;
	v133 = v56
	goto L27
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+10)) = int32(_a_F_heap_redo_3)
	v118 = int32(_a_F_heap_redo_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+18)) = uint16(v118)
	v124 = int32(_a_F_heap_redo_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+16)) = uint16(v124)
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+14)) = uint16(v124)
	goto L36
L38:
	;
	goto L41
L39:
	;
	goto L40
L40:
	;
	goto L46
L41:
	;
	v95 = v76 + v77
	v97 = v76 + int32(4)
	if base.Ui32(v97) < base.Ui32(v95) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v99 = v95
	goto L44
L43:
	;
	v99 = v97
	goto L44
L44:
	;
	v104 = (v76^int32(-1)+v99)&int32(-4) + int32(4)
	if v104 == int32(0) {
		goto L37
	} else {
		goto L45
	}
L45:
	;
	base.MemoryFill(m, v76, int32(0), v104)
	goto L37
L46:
	;
	base.MemoryFill(m, v76, int32(0), v77)
	goto L37
L47:
	;
	if v130 != 0 {
		v340 = v2
		v341 = v2
		goto L26
	} else {
		goto L48
	}
L48:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_heap_redo[0])))
	v133 = v132
	goto L27
L49:
	;
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153)+12)))
	if base.Ui32(v154) < base.Ui32(int32(25)) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[1]))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v139+(v133^int32(-1))<<(uint(int32(2))%32))))
	v153 = v145
	goto L49
L51:
	;
	goto L52
L52:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[2]))
	v153 = v147 + v133<<(uint(int32(13))%32) + int32(-8192)
	goto L49
L53:
	;
	v163 = int32(1)
	goto L55
L54:
	;
	v163 = int32(base.Ui32(v154+int32(_a_F_heap_redo_5))>>(uint(int32(2))%32)) + int32(1)
	goto L55
L55:
	;
	if base.Ui32(v163&int32(_a_F_heap_redo_6)) < base.Ui32(v134) {
		goto L11
	} else {
		goto L56
	}
L56:
	;
	v168 = int32(base.Ui32(v37) >> (uint(int32(16)) % 32))
	v169 = int32(0)
	v171 = v17 + int32(72)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+72))
	if v174 < v169 {
		v196 = v169
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v199)+2)))
	v201 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v199))))
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+96)) = int32(0)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
	v207 = v205 - int32(5)
	if v207 != 0 {
		goto L68
	} else {
		goto L69
	}
L58:
	;
	v199 = v196
	goto L57
L59:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173+int32(0))+76)))
	if v179 != int32(1) {
		v196 = v169
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v183 = v173 + int32(76)
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+43)))
	if v184 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	if v171 == int32(0) {
		v196 = v169
		goto L58
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if v171 != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v189 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v171))) = v189
	v199 = v189
	goto L57
L65:
	;
	v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v171))) = v192
	goto L67
L66:
	;
	goto L67
L67:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v183)+44))
	v196 = v194
	goto L58
L68:
	;
	base.MemoryCopy(m, v17+int32(115), v199+int32(5), v207)
	goto L70
L69:
	;
	goto L70
L70:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+114)) = uint8(v202)
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+110)) = uint16(v201)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+36))
	v218 = v200 & int32(_a_F_heap_redo_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+112)) = uint16(v218)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+100)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+92)) = v216
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+108)) = uint16(v38)
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+106)) = uint16(v37)
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+104)) = uint16(v168)
	v230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28))))
	v232 = F_PageAddItemExtended(m, v153, v17+int32(92), v205+int32(18), v230, int32(3))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L20
	} else {
		goto L71
	}
L71:
	;
	if v232 == int32(0) {
		goto L10
	} else {
		goto L72
	}
L72:
	;
	v239 = int32(4)
	v240 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153)+14)))
	v241 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153)+12)))
	v242 = v240 - v241
	if v242 <= v239 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v303)+36))
	if base.Ui32(v304) < base.Ui32(int32(3)) {
		goto L92
	} else {
		goto L93
	}
L74:
	;
	v245 = v239
	goto L76
L75:
	;
	v245 = v242
	goto L76
L76:
	;
	v247 = v245 - int32(4)
	if v247 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v302 = int32(0)
	goto L73
L78:
	;
	goto L79
L79:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v241) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v302 = v247
	goto L73
L81:
	;
	v258 = int32(base.Ui32(v241+int32(_a_F_heap_redo_5)) >> (uint(int32(2)) % 32))
	goto L83
L82:
	;
	v258 = int32(0)
	goto L83
L83:
	;
	if base.Ui32(v258&int32(_a_F_heap_redo_6)) < base.Ui32(int32(291)) {
		goto L80
	} else {
		goto L84
	}
L84:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+10)))
	if v263&int32(1) == int32(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v302 = int32(0)
	goto L73
L86:
	;
	goto L87
L87:
	;
	v272 = int32(1)
	goto L88
L88:
	;
	v281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153+int32(20)+v272&int32(_a_F_heap_redo_6)<<(uint(int32(2))%32))+1)))
	if v281&int32(384) == int32(0) {
		goto L80
	} else {
		goto L90
	}
L89:
	;
	v302 = int32(0)
	goto L73
L90:
	;
	v287 = v272 + int32(1)
	v288 = int32(_a_F_heap_redo_6)
	if base.Ui32(v287&v288) <= base.Ui32(v258&v288) {
		v272 = v287
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v153))) = base.I64_rotl(v27, int64(32))
	v326 = int32(1)
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+2)))
	if v327&v326 != 0 {
		goto L99
	} else {
		goto L100
	}
L93:
	;
	v307 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+112)))
	v308 = int32(768)
	if v307&v308 == v308 {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v153)+20))
	v313 = int32(0)
	if base.B2i32(base.Ui32(v312) < base.Ui32(int32(3)))|base.B2i32(v313 <= v304-v312) != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v320 = v312
	goto L97
L96:
	;
	v320 = v313
	goto L97
L97:
	;
	if v320 != 0 {
		goto L92
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153)+20)) = v304
	goto L92
L99:
	;
	v330 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153)+10)))
	v332 = v330 & int32(_a_F_heap_redo_8)
	*(*uint16)(unsafe.Add(mBase, uint32(v153)+10)) = uint16(v332)
	goto L101
L100:
	;
	goto L101
L101:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_heap_redo[0])))
	F_MarkBufferDirty(m, v334)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L20
	} else {
		goto L102
	}
L102:
	;
	v340 = v302
	v341 = v326
	goto L26
L103:
	;
	F_UnlockReleaseBuffer(m, v347)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L20
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	if v341&base.B2i32(base.Ui32(v340) < base.Ui32(int32(1638))) == int32(0) {
		goto L12
	} else {
		goto L107
	}
L106:
	;
	goto L105
L107:
	;
	v355 = *(*int64)(unsafe.Add(mBase, uint32(v17)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v355
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v17)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v357
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	F_XLogRecordPageWithFreeSpace(m, v17+int32(8), v361, v340)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L20
	} else {
		goto L108
	}
L108:
	;
	goto L12
L109:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_heap_redo[0])))
	v375 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v365)+4)))
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365)+7)))
	if v376&int32(1) != 0 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v17)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v379
	v381 = *(*int64)(unsafe.Add(mBase, uint32(v17)+92))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+40)) = v381
	F_heap_xlog_vm_clear(m, l0, v17+int32(40), v374, int32(3))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L20
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v391 = F_XLogReadBufferForRedo(m, l0, int32(0), v17+int32(80))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L20
	} else {
		goto L114
	}
L113:
	;
	goto L112
L114:
	;
	if v391 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v17)+80))
	if v395 < int32(0) {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	goto L117
L117:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v17)+80))
	if v548 == int32(0) {
		goto L12
	} else {
		goto L154
	}
L118:
	;
	v414 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v365)+4)))
	if v414 == int32(0) {
		goto L9
	} else {
		goto L122
	}
L119:
	;
	v399 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[1]))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v399+(v395^int32(-1))<<(uint(int32(2))%32))))
	v413 = v405
	goto L118
L120:
	;
	goto L121
L121:
	;
	v407 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[2]))
	v413 = v407 + v395<<(uint(int32(13))%32) + int32(-8192)
	goto L118
L122:
	;
	v417 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v413)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v417) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v425 = int32(base.Ui32(v417+int32(_a_F_heap_redo_5)) >> (uint(int32(2)) % 32))
	goto L125
L124:
	;
	v425 = int32(0)
	goto L125
L125:
	;
	if base.Ui32(v425&int32(_a_F_heap_redo_6)) < base.Ui32(v414) {
		goto L9
	} else {
		goto L126
	}
L126:
	;
	v430 = v413 + int32(20)
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v430+v414<<(uint(int32(2))%32))))
	if v434&int32(_a_F_heap_redo_9) != int32(_a_F_heap_redo_10) {
		goto L8
	} else {
		goto L127
	}
L127:
	;
	v441 = v413 + v434&int32(_a_F_heap_redo_11)
	v442 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v441)+20)))
	v444 = v442 & int32(_a_F_heap_redo_12)
	*(*uint16)(unsafe.Add(mBase, uint32(v441)+20)) = uint16(v444)
	v446 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v441)+18)))
	v448 = v446 & int32(-24577)
	*(*uint16)(unsafe.Add(mBase, uint32(v441)+18)) = uint16(v448)
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v441)+18)) = uint16(v448)
	*(*uint16)(unsafe.Add(mBase, uint32(v441)+20)) = uint16(v444)
	v453 = int32(1)
	v472 = v450<<(uint(v453)%32)&int32(16) | (v450<<(uint(int32(4))%32)&int32(64) | (v450<<(uint(int32(6))%32)&int32(128) | v450&v453<<(uint(int32(12))%32))) | v444
	if v450&int32(15) != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v441)+20)) = uint16(v472)
	goto L130
L129:
	;
	goto L130
L130:
	;
	if v450&int32(16) != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v479 = v448 | int32(_a_F_heap_redo_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v441)+18)) = uint16(v479)
	goto L133
L132:
	;
	goto L133
L133:
	;
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365)+7)))
	if v481&int32(8) == int32(0) {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v490 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v441)+8)) = v490
	v493 = v472 & int32(_a_F_heap_redo_13)
	*(*uint16)(unsafe.Add(mBase, uint32(v441)+20)) = uint16(v493)
	v495 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v495)+36))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v430)))
	if v497 == v490 {
		goto L139
	} else {
		goto L140
	}
L135:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v365)))
	*(*int32)(unsafe.Add(mBase, uint32(v441)+4)) = v486
	goto L134
L136:
	;
	goto L137
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v441))) = int32(0)
	goto L134
L138:
	;
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365)+7)))
	if v512&int32(1) != 0 {
		goto L147
	} else {
		goto L148
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v430))) = v496
	goto L138
L140:
	;
	v500 = int32(3)
	if base.B2i32(base.Ui32(v497) < base.Ui32(v500))|base.B2i32(base.Ui32(v496) < base.Ui32(v500)) == int32(0) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	if v496-v497 < int32(0) {
		goto L139
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	if base.Ui32(v497) <= base.Ui32(v496) {
		goto L138
	} else {
		goto L145
	}
L144:
	;
	goto L138
L145:
	;
	goto L139
L146:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v441)+14)) = uint16(v532)
	*(*int64)(unsafe.Add(mBase, uint32(v413))) = base.I64_rotl(v364, int64(32))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v17)+80))
	F_MarkBufferDirty(m, v537)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L20
	} else {
		goto L153
	}
L147:
	;
	v515 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v413)+10)))
	v517 = v515 & int32(_a_F_heap_redo_8)
	*(*uint16)(unsafe.Add(mBase, uint32(v413)+10)) = uint16(v517)
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365)+7)))
	v520 = v519
	goto L149
L148:
	;
	v520 = v512
	goto L149
L149:
	;
	if v520&int32(16) != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v523 = int32(_a_F_heap_redo_14)
	*(*uint16)(unsafe.Add(mBase, uint32(v441)+16)) = uint16(v523)
	v525 = int32(_a_F_heap_redo_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v441)+12)) = uint16(v525)
	v532 = v525
	goto L146
L151:
	;
	goto L152
L152:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v441)+16)) = uint16(v375)
	v530 = int32(base.Ui32(v374) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v441)+12)) = uint16(v530)
	v532 = v374
	goto L146
L153:
	;
	goto L117
L154:
	;
	F_UnlockReleaseBuffer(m, v548)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L20
	} else {
		goto L155
	}
L155:
	;
	goto L12
L156:
	;
	goto L12
L157:
	;
	goto L12
L158:
	;
	if v564 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	if v568 < int32(0) {
		goto L163
	} else {
		goto L164
	}
L160:
	;
	goto L161
L161:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	if v647 == int32(0) {
		goto L12
	} else {
		goto L177
	}
L162:
	;
	v587 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v559))))
	if v587 == int32(0) {
		goto L7
	} else {
		goto L166
	}
L163:
	;
	v572 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[1]))
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v572+(v568^int32(-1))<<(uint(int32(2))%32))))
	v586 = v578
	goto L162
L164:
	;
	goto L165
L165:
	;
	v580 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[2]))
	v586 = v580 + v568<<(uint(int32(13))%32) + int32(-8192)
	goto L162
L166:
	;
	v590 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v586)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v590) {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v598 = int32(base.Ui32(v590+int32(_a_F_heap_redo_5)) >> (uint(int32(2)) % 32))
	goto L169
L168:
	;
	v598 = int32(0)
	goto L169
L169:
	;
	if base.Ui32(v598&int32(_a_F_heap_redo_6)) < base.Ui32(v587) {
		goto L7
	} else {
		goto L170
	}
L170:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v586+v587<<(uint(int32(2))%32))+20))
	if v605&int32(_a_F_heap_redo_9) != int32(_a_F_heap_redo_10) {
		goto L6
	} else {
		goto L171
	}
L171:
	;
	if v568 < int32(0) {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	v631 = v586 + v605&int32(_a_F_heap_redo_11)
	*(*uint16)(unsafe.Add(mBase, uint32(v631)+16)) = uint16(v587)
	*(*uint16)(unsafe.Add(mBase, uint32(v631)+14)) = uint16(v628)
	v635 = int32(base.Ui32(v628) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v631)+12)) = uint16(v635)
	*(*int64)(unsafe.Add(mBase, uint32(v586))) = base.I64_rotl(v560, int64(32))
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	F_MarkBufferDirty(m, v640)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L20
	} else {
		goto L176
	}
L173:
	;
	v613 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[3]))
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v613+(v568^int32(-1))*int32(56))+16))
	v628 = v619
	goto L172
L174:
	;
	goto L175
L175:
	;
	v621 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[4]))
	v622 = int32(56)
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v621+v568*v622-v622)+16))
	v628 = v627
	goto L172
L176:
	;
	goto L161
L177:
	;
	F_UnlockReleaseBuffer(m, v647)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L20
	} else {
		goto L178
	}
L178:
	;
	goto L12
L179:
	;
	v657 = int32(0)
	F_XLogRecGetBlockTag(m, l0, v657, v17+int32(92), v657, v17+int32(80))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L20
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v678 = F_XLogReadBufferForRedo(m, l0, int32(0), v17+int32(92))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L20
	} else {
		goto L184
	}
L182:
	;
	v665 = *(*int64)(unsafe.Add(mBase, uint32(v17)+92))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+56)) = v665
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v17)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v667
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v17)+80))
	F_heap_xlog_vm_clear(m, l0, v17+int32(56), v671, int32(2))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L20
	} else {
		goto L183
	}
L183:
	;
	goto L181
L184:
	;
	if v678 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	if v682 < int32(0) {
		goto L189
	} else {
		goto L190
	}
L186:
	;
	goto L187
L187:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	if v829 == int32(0) {
		goto L12
	} else {
		goto L212
	}
L188:
	;
	v701 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v653)+4)))
	if v701 == int32(0) {
		goto L5
	} else {
		goto L192
	}
L189:
	;
	v686 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[1]))
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v686+(v682^int32(-1))<<(uint(int32(2))%32))))
	v700 = v692
	goto L188
L190:
	;
	goto L191
L191:
	;
	v694 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[2]))
	v700 = v694 + v682<<(uint(int32(13))%32) + int32(-8192)
	goto L188
L192:
	;
	v704 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v700)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v704) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v712 = int32(base.Ui32(v704+int32(_a_F_heap_redo_5)) >> (uint(int32(2)) % 32))
	goto L195
L194:
	;
	v712 = int32(0)
	goto L195
L195:
	;
	if base.Ui32(v712&int32(_a_F_heap_redo_6)) < base.Ui32(v701) {
		goto L5
	} else {
		goto L196
	}
L196:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v700+v701<<(uint(int32(2))%32))+20))
	if v719&int32(_a_F_heap_redo_9) != int32(_a_F_heap_redo_10) {
		goto L4
	} else {
		goto L197
	}
L197:
	;
	v726 = v700 + v719&int32(_a_F_heap_redo_11)
	v727 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v726)+20)))
	v729 = v727 & int32(_a_F_heap_redo_12)
	*(*uint16)(unsafe.Add(mBase, uint32(v726)+20)) = uint16(v729)
	v731 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v726)+18)))
	v733 = v731 & int32(-8193)
	*(*uint16)(unsafe.Add(mBase, uint32(v726)+18)) = uint16(v733)
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v653)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v726)+18)) = uint16(v733)
	*(*uint16)(unsafe.Add(mBase, uint32(v726)+20)) = uint16(v729)
	v738 = int32(1)
	v757 = v735<<(uint(v738)%32)&int32(16) | (v735<<(uint(int32(4))%32)&int32(64) | (v735<<(uint(int32(6))%32)&int32(128) | v735&v738<<(uint(int32(12))%32))) | v729
	if v735&int32(15) != 0 {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v726)+20)) = uint16(v757)
	goto L200
L199:
	;
	goto L200
L200:
	;
	if v735&int32(16) != 0 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v764 = v731 | int32(_a_F_heap_redo_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v726)+18)) = uint16(v764)
	v766 = v764
	goto L203
L202:
	;
	v766 = v733
	goto L203
L203:
	;
	v769 = int32(0)
	if base.B2i32(v757&int32(128) == v769)&base.B2i32(v757&int32(_a_F_heap_redo_15) != int32(64)) == v769 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v779 = v766 & int32(_a_F_heap_redo_16)
	*(*uint16)(unsafe.Add(mBase, uint32(v726)+18)) = uint16(v779)
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	if v781 < int32(0) {
		goto L208
	} else {
		goto L209
	}
L205:
	;
	v808 = v757
	goto L206
L206:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v653)))
	v811 = v808 & int32(_a_F_heap_redo_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v726)+20)) = uint16(v811)
	*(*int32)(unsafe.Add(mBase, uint32(v726)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v726)+4)) = v809
	*(*int64)(unsafe.Add(mBase, uint32(v700))) = base.I64_rotl(v652, int64(32))
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	F_MarkBufferDirty(m, v819)
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L20
	} else {
		goto L211
	}
L207:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v726)+16)) = uint16(v701)
	*(*uint16)(unsafe.Add(mBase, uint32(v726)+14)) = uint16(v800)
	v804 = int32(base.Ui32(v800) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v726)+12)) = uint16(v804)
	v806 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v726)+20)))
	v808 = v806
	goto L206
L208:
	;
	v785 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[3]))
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v785+(v781^int32(-1))*int32(56))+16))
	v800 = v791
	goto L207
L209:
	;
	goto L210
L210:
	;
	v793 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[4]))
	v794 = int32(56)
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v793+v781*v794-v794)+16))
	v800 = v799
	goto L207
L211:
	;
	goto L187
L212:
	;
	F_UnlockReleaseBuffer(m, v829)
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L20
	} else {
		goto L213
	}
L213:
	;
	goto L12
L214:
	;
	if v839 == int32(0) {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v843 = int32(0)
	v845 = v17 + int32(80)
	v847 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v847)+72))
	if v848 < v843 {
		v870 = v843
		goto L219
	} else {
		goto L220
	}
L216:
	;
	goto L217
L217:
	;
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	if v938 != 0 {
		goto L244
	} else {
		goto L245
	}
L218:
	;
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	if v874 < int32(0) {
		goto L230
	} else {
		goto L231
	}
L219:
	;
	v873 = v870
	goto L218
L220:
	;
	v853 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v847+int32(0))+76)))
	if v853 != int32(1) {
		v870 = v843
		goto L219
	} else {
		goto L221
	}
L221:
	;
	v857 = v847 + int32(76)
	v858 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v857)+43)))
	if v858 == int32(0) {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	if v845 == int32(0) {
		v870 = v843
		goto L219
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	if v845 != 0 {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	v863 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v845))) = v863
	v873 = v863
	goto L218
L226:
	;
	v866 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v857)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v845))) = v866
	goto L228
L227:
	;
	goto L228
L228:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v857)+44))
	v870 = v868
	goto L219
L229:
	;
	v893 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v834))))
	if v893 == int32(0) {
		goto L3
	} else {
		goto L233
	}
L230:
	;
	v878 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[1]))
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v878+(v874^int32(-1))<<(uint(int32(2))%32))))
	v892 = v884
	goto L229
L231:
	;
	goto L232
L232:
	;
	v886 = *(*int32)(unsafe.Add(mBase, _c_F_heap_redo[2]))
	v892 = v886 + v874<<(uint(int32(13))%32) + int32(-8192)
	goto L229
L233:
	;
	v896 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v892)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v896) {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v904 = int32(base.Ui32(v896+int32(_a_F_heap_redo_5)) >> (uint(int32(2)) % 32))
	goto L236
L235:
	;
	v904 = int32(0)
	goto L236
L236:
	;
	if base.Ui32(v904&int32(_a_F_heap_redo_6)) < base.Ui32(v893) {
		goto L3
	} else {
		goto L237
	}
L237:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v892+v893<<(uint(int32(2))%32))+20))
	if v911&int32(_a_F_heap_redo_9) != int32(_a_F_heap_redo_10) {
		goto L2
	} else {
		goto L238
	}
L238:
	;
	v920 = v892 + v911&int32(_a_F_heap_redo_11)
	v921 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v920)+22)))
	v922 = int32(base.Ui32(v911)>>(uint(int32(17))%32)) - v921
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v17)+80))
	if v922 != v923 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	if v922 != 0 {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	base.MemoryCopy(m, v920+v921, v873, v922)
	goto L242
L241:
	;
	goto L242
L242:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v892))) = base.I64_rotl(v835, int64(32))
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	F_MarkBufferDirty(m, v930)
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L20
	} else {
		goto L243
	}
L243:
	;
	goto L217
L244:
	;
	F_UnlockReleaseBuffer(m, v938)
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L20
	} else {
		goto L247
	}
L245:
	;
	goto L246
L246:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v834)+16))
	v944 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v834)+12)))
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v834)+4))
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v834)+8))
	F_ProcessCommittedInvalidationMessages(m, v834+int32(20), v943, v944, v945, v946)
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L20
	} else {
		goto L248
	}
L247:
	;
	goto L246
L248:
	;
	goto L12
L249:
	;
	F_errmsg_internal(m, int32(_a_F_heap_redo_17), int32(0))
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L20
	} else {
		goto L250
	}
L250:
	;
	F_errfinish(m, int32(_a_F_heap_redo_18), int32(462), int32(_a_F_heap_redo_19))
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L20
	} else {
		goto L251
	}
L251:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L252:
	;
	F_errmsg_internal(m, int32(_a_F_heap_redo_20), int32(0))
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L20
	} else {
		goto L253
	}
L253:
	;
	F_errfinish(m, int32(_a_F_heap_redo_18), int32(486), int32(_a_F_heap_redo_19))
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L20
	} else {
		goto L254
	}
L254:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L255:
	;
	F_errmsg_internal(m, int32(_a_F_heap_redo_21), int32(0))
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L20
	} else {
		goto L256
	}
L256:
	;
	F_errfinish(m, int32(_a_F_heap_redo_18), int32(363), int32(_a_F_heap_redo_22))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L20
	} else {
		goto L257
	}
L257:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L258:
	;
	F_errmsg_internal(m, int32(_a_F_heap_redo_23), int32(0))
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L20
	} else {
		goto L259
	}
L259:
	;
	F_errfinish(m, int32(_a_F_heap_redo_18), int32(366), int32(_a_F_heap_redo_22))
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L20
	} else {
		goto L260
	}
L260:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L261:
	;
	F_errmsg_internal(m, int32(_a_F_heap_redo_21), int32(0))
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L20
	} else {
		goto L262
	}
L262:
	;
	F_errfinish(m, int32(_a_F_heap_redo_18), int32(1099), int32(_a_F_heap_redo_24))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L20
	} else {
		goto L263
	}
L263:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L264:
	;
	F_errmsg_internal(m, int32(_a_F_heap_redo_23), int32(0))
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L20
	} else {
		goto L265
	}
L265:
	;
	F_errfinish(m, int32(_a_F_heap_redo_18), int32(1102), int32(_a_F_heap_redo_24))
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L20
	} else {
		goto L266
	}
L266:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L267:
	;
	F_errmsg_internal(m, int32(_a_F_heap_redo_21), int32(0))
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L20
	} else {
		goto L268
	}
L268:
	;
	F_errfinish(m, int32(_a_F_heap_redo_18), int32(1156), int32(_a_F_heap_redo_25))
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L20
	} else {
		goto L269
	}
L269:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L270:
	;
	F_errmsg_internal(m, int32(_a_F_heap_redo_23), int32(0))
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L20
	} else {
		goto L271
	}
L271:
	;
	F_errfinish(m, int32(_a_F_heap_redo_18), int32(1159), int32(_a_F_heap_redo_25))
	mBase = m.M
	v1071 = m.ExcPending
	if v1071 != 0 {
		goto L20
	} else {
		goto L272
	}
L272:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L273:
	;
	F_errmsg_internal(m, int32(_a_F_heap_redo_21), int32(0))
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L20
	} else {
		goto L274
	}
L274:
	;
	F_errfinish(m, int32(_a_F_heap_redo_18), int32(1273), int32(_a_F_heap_redo_26))
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L20
	} else {
		goto L275
	}
L275:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L276:
	;
	F_errmsg_internal(m, int32(_a_F_heap_redo_23), int32(0))
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L20
	} else {
		goto L277
	}
L277:
	;
	F_errfinish(m, int32(_a_F_heap_redo_18), int32(1276), int32(_a_F_heap_redo_26))
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		goto L20
	} else {
		goto L278
	}
L278:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L279:
	;
	F_errmsg_internal(m, int32(_a_F_heap_redo_27), int32(0))
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L20
	} else {
		goto L280
	}
L280:
	;
	F_errfinish(m, int32(_a_F_heap_redo_18), int32(1282), int32(_a_F_heap_redo_26))
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L20
	} else {
		goto L281
	}
L281:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_heap_scan_stream_read_next_serial(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+52)))
	if v6 == int32(0) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
		v10 = int32(-1)
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
		if v11 == int32(0) {
			v35 = v10
			v39 = v35
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
			if v14 == int32(0) {
				v35 = v10
				v39 = v35
			} else {
				if v9 == int32(1) {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
					v39 = v19
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v20 & int32(-129)
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
					if v14 != int32(-1) {
						v30 = base.I32_rem_u_s(v24+v14-int32(1), v11)
						v39 = v30
					} else {
						if v24 != 0 {
							v39 = v24 - int32(1)
						} else {
							v35 = v11 - int32(1)
							v39 = v35
						}
					}
				}
			}
		}
		v40 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l1)+52)) = uint8(v40)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v39
		return v39
	} else {
		v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
		v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
		if v45 == int32(1) {
			v49 = v44 + int32(1)
			v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
			if base.Ui32(v49) < base.Ui32(v51) {
				v53 = v49
			} else {
				v53 = int32(0)
			}
			v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+28)))
			if v54&int32(128) != 0 {
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				F_ss_report_location(m, v57, v53)
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
					if v62 == v53 {
						v64 = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v64
						return v64
					} else {
						v68 = int32(-1)
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
						if v69 == v68 {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v53
							return v53
						} else {
							v75 = v69 - int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v75
							if v75 == int32(0) {
								v101 = v68
								*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v101
								return v101
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v53
								return v53
							}
						}
					}
				}
			} else {
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
				if v62 == v53 {
					v64 = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v64
					return v64
				} else {
					v68 = int32(-1)
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
					if v69 == v68 {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v53
						return v53
					} else {
						v75 = v69 - int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v75
						if v75 == int32(0) {
							v101 = v68
							*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v101
							return v101
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v53
							return v53
						}
					}
				}
			}
		} else {
			v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
			if v81 == v44 {
				v83 = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v83
				return v83
			} else {
				v87 = int32(-1)
				v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
				if v88 != v87 {
					v92 = v88 - int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v92
					if v92 == int32(0) {
						v101 = v87
					} else {
						if v44 != 0 {
							v98 = v44
						} else {
							v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
							v98 = v97
						}
						v101 = v98 - int32(1)
					}
				} else {
					if v44 != 0 {
						v98 = v44
					} else {
						v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
						v98 = v97
					}
					v101 = v98 - int32(1)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v101
				return v101
			}
		}
	}
}
func F_heap_truncate_check_FKs(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
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
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
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
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(48)
	return
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v15 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v21 = v3
	v22 = v3
	goto L6
L4:
	;
	v49 = v3
	goto L5
L5:
	;
	if v49 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v22<<(uint(int32(2))%32))))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+48))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+125)))
	if v32 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v49 = v41
	goto L5
L8:
	;
	v43 = v22 + int32(1)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v43 < v44 {
		v21 = v41
		v22 = v43
		goto L6
	} else {
		goto L15
	}
L9:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+119)))
	if v35 != int32(112) {
		v41 = v21
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)+56))
	v39 = F_lappend_oid(m, v21, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	return
L14:
	;
	v41 = v39
	goto L8
L15:
	;
	goto L7
L16:
	;
	v56 = F_heap_truncate_find_FKs(m, v49)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	if v56 == int32(0) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v60 <= int32(0) {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v64 = int32(0)
	goto L20
L20:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72+v64<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v76
	v82 = F_list_make1_impl(m, int32(480), v11+int32(40))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L13
	} else {
		goto L23
	}
L21:
	;
	goto L1
L22:
	;
	v206 = v64 + int32(1)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v206 < v207 {
		v64 = v206
		goto L20
	} else {
		goto L60
	}
L23:
	;
	v84 = F_heap_truncate_find_FKs(m, v82)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L13
	} else {
		goto L24
	}
L24:
	;
	if v84 == int32(0) {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v88 = int32(0)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v89 <= v88 {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v96 = v88
	goto L27
L27:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100+v96<<(uint(int32(2))%32))))
	v105 = int32(0)
	if v49 == v105 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v148 = F_get_rel_name(m, v76)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L13
	} else {
		goto L46
	}
L29:
	;
	if v143 != 0 {
		goto L42
	} else {
		goto L43
	}
L30:
	;
	v143 = int32(0)
	goto L29
L31:
	;
	goto L32
L32:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v111 <= int32(0) {
		v137 = v105
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v143 = v137
	goto L29
L34:
	;
	v114 = int32(0)
	if v114 < v111 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v117 = v111
	goto L37
L36:
	;
	v117 = v114
	goto L37
L37:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v120 = int32(0)
	goto L38
L38:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v118+v120<<(uint(int32(2))%32))))
	v129 = base.B2i32(v128 == v104)
	if v128 == v104 {
		v137 = v129
		goto L33
	} else {
		goto L40
	}
L39:
	;
	v137 = v129
	goto L33
L40:
	;
	v131 = v120 + int32(1)
	if v131 != v117 {
		v120 = v131
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v145 = v96 + int32(1)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v145 < v146 {
		v96 = v145
		goto L27
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	goto L28
L45:
	;
	goto L22
L46:
	;
	v150 = F_get_rel_name(m, v104)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L13
	} else {
		goto L48
	}
L48:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L13
	} else {
		goto L49
	}
L49:
	;
	if l1 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	F_errmsg(m, int32(_a_F_heap_truncate_check_FKs_0), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L13
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	F_errmsg(m, int32(_a_F_heap_truncate_check_FKs_1), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L13
	} else {
		goto L57
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v150
	v170 = F_errdetail(m, int32(_a_F_heap_truncate_check_FKs_2), v11+int32(32))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L13
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v150
	F_errhint(m, int32(_a_F_heap_truncate_check_FKs_3), v11+int32(16))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L13
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_heap_truncate_check_FKs_4), int32(3786), int32(_a_F_heap_truncate_check_FKs_5))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L13
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v150
	v190 = F_errdetail(m, int32(_a_F_heap_truncate_check_FKs_6), v11)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L13
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_heap_truncate_check_FKs_4), int32(3777), int32(_a_F_heap_truncate_check_FKs_5))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L13
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
	goto L21
}
func F_heap_tuple_should_freeze(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v197 int32
	_ = v197
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	v18 = int32(768)
	if v17&v18 == v18 {
		v45 = int32(0)
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if base.Ui32(v23) < base.Ui32(int32(3)) {
			v45 = int32(0)
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v30 = int32(0)
			if base.B2i32(base.Ui32(v26) < base.Ui32(int32(3)))|base.B2i32(v30 <= v23-v26) == v30 {
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v23
			} else {
			}
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			if base.Ui32(v37) < base.Ui32(int32(3)) {
				v45 = int32(0)
			} else {
				v45 = int32(base.Ui32(v23-v37) >> (uint(int32(31)) % 32))
			}
		}
	}
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	if v48&int32(_a_F_heap_tuple_should_freeze_0) != 0 {
		v51 = int32(0)
	} else {
		v51 = v47
	}
	if base.Ui32(int32(3)) <= base.Ui32(v51) {
		v54 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v58 = int32(0)
		if base.B2i32(base.Ui32(v54) < base.Ui32(int32(3)))|base.B2i32(v58 <= v51-v54) == v58 {
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v51
		} else {
		}
		v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		v174 = v45 | base.B2i32(base.Ui32(int32(2)) < base.Ui32(v64))&base.B2i32(v51-v64 < int32(0))
		v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
		if base.Ui32(v180) < base.Ui32(int32(_a_F_heap_tuple_should_freeze_1)) {
			v197 = v174
		} else {
			v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if base.Ui32(v183) < base.Ui32(int32(3)) {
				v197 = v174
			} else {
				v186 = int32(1)
				v187 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				if base.B2i32(base.Ui32(v187) < base.Ui32(int32(3)))|base.B2i32(int32(0) <= v183-v187) != 0 {
					v197 = v186
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v183
					v197 = v186
				}
			}
		}
		m.G0 = v14 + int32(16)
		return v197 & int32(1)
	} else {
		v76 = v48 << (uint(int32(19)) % 32) >> (uint(int32(31)) % 32) & v47
		if v76 == int32(0) {
			v174 = v45
			v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
			if base.Ui32(v180) < base.Ui32(int32(_a_F_heap_tuple_should_freeze_1)) {
				v197 = v174
			} else {
				v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if base.Ui32(v183) < base.Ui32(int32(3)) {
					v197 = v174
				} else {
					v186 = int32(1)
					v187 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					if base.B2i32(base.Ui32(v187) < base.Ui32(int32(3)))|base.B2i32(int32(0) <= v183-v187) != 0 {
						v197 = v186
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v183
						v197 = v186
					}
				}
			}
			m.G0 = v14 + int32(16)
			return v197 & int32(1)
		} else {
			if v48&int32(_a_F_heap_tuple_should_freeze_2) == int32(_a_F_heap_tuple_should_freeze_3) {
				v83 = int32(1)
				v84 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
				if int32(0) <= v76-v84 {
					v174 = v83
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v76
					v174 = v83
				}
				v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
				if base.Ui32(v180) < base.Ui32(int32(_a_F_heap_tuple_should_freeze_1)) {
					v197 = v174
				} else {
					v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if base.Ui32(v183) < base.Ui32(int32(3)) {
						v197 = v174
					} else {
						v186 = int32(1)
						v187 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						if base.B2i32(base.Ui32(v187) < base.Ui32(int32(3)))|base.B2i32(int32(0) <= v183-v187) != 0 {
							v197 = v186
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v183
							v197 = v186
						}
					}
				}
				m.G0 = v14 + int32(16)
				return v197 & int32(1)
			} else {
				v89 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
				if v76-v89 < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v76
					v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
					v95 = v94
				} else {
					v95 = v48
				}
				v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				v100 = v45 | base.B2i32(v76-v96 < int32(0))
				v112 = F_GetMultiXactIdMembers(m, v76, v14+int32(12), int32(base.Ui32(v95&int32(128))>>(uint(int32(7))%32))|base.B2i32(v95&int32(_a_F_heap_tuple_should_freeze_4) == int32(64)))
				mBase = m.M
				v115 = m.ExcPending
				if v115 != 0 {
					return int32(0)
				} else {
					if v112 <= int32(0) {
						v174 = v100
						v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
						if base.Ui32(v180) < base.Ui32(int32(_a_F_heap_tuple_should_freeze_1)) {
							v197 = v174
						} else {
							v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							if base.Ui32(v183) < base.Ui32(int32(3)) {
								v197 = v174
							} else {
								v186 = int32(1)
								v187 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								if base.B2i32(base.Ui32(v187) < base.Ui32(int32(3)))|base.B2i32(int32(0) <= v183-v187) != 0 {
									v197 = v186
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = v183
									v197 = v186
								}
							}
						}
						m.G0 = v14 + int32(16)
						return v197 & int32(1)
					} else {
						v118 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						v119 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
						v124 = v118
						v126 = v100
						v127 = int32(0)
						for {
							v132 = int32(3)
							v135 = *(*int32)(unsafe.Add(mBase, uint32(v119+v127<<(uint(v132)%32))))
							v137 = base.B2i32(base.Ui32(v135) < base.Ui32(v132))
							if v137|base.B2i32(base.Ui32(v124) < base.Ui32(v132)) == int32(0) {
								if v135-v124 < int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = v135
									v148 = v135
								} else {
									v148 = v124
								}
							} else {
								if base.Ui32(v124) <= base.Ui32(v135) {
									v148 = v124
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = v135
									v148 = v135
								}
							}
							v149 = int32(0)
							v151 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
							if base.B2i32(v137 == v149)&base.B2i32(base.Ui32(int32(2)) < base.Ui32(v151)) == v149 {
								v161 = base.B2i32(base.Ui32(v135) < base.Ui32(v151))
							} else {
								v161 = int32(base.Ui32(v135-v151) >> (uint(int32(31)) % 32))
							}
							v162 = v161 | v126
							v164 = v127 + int32(1)
							if v164 != v112 {
								v124 = v148
								v126 = v162
								v127 = v164
								continue
							} else {
								break
							}
							break
						}
						v166 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
						F_pfree(m, v166)
						mBase = m.M
						v168 = m.ExcPending
						if v168 != 0 {
							return int32(0)
						} else {
							v174 = v162
							v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
							if base.Ui32(v180) < base.Ui32(int32(_a_F_heap_tuple_should_freeze_1)) {
								v197 = v174
							} else {
								v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								if base.Ui32(v183) < base.Ui32(int32(3)) {
									v197 = v174
								} else {
									v186 = int32(1)
									v187 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
									if base.B2i32(base.Ui32(v187) < base.Ui32(int32(3)))|base.B2i32(int32(0) <= v183-v187) != 0 {
										v197 = v186
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = v183
										v197 = v186
									}
								}
							}
							m.G0 = v14 + int32(16)
							return v197 & int32(1)
						}
					}
				}
			}
		}
	}
}
func F_heap_vac_scan_next_block(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
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
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v141 int32
	_ = v141
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+248))
	v18 = v16 + int32(1)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if base.Ui32(v19) <= base.Ui32(v18) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v155
L2:
	;
	v21 = int32(-1)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+260))
	if v22 == int32(0) {
		v155 = v21
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+252))
	if base.B2i32(base.Ui32(v18) <= base.Ui32(v31))&base.B2i32(v31 != int32(-1)) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	F_ReleaseBuffer(m, v22)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+260)) = int32(0)
	v155 = v21
	goto L1
L8:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l1)+248))
	v155 = v154
	goto L1
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+248)) = v129
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+256)))
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v141)
	goto L8
L10:
	;
	v127 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+64)) = uint8(v127)
	v129 = v67
	goto L9
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+260))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v38
	v43 = v31
	v46 = int32(0)
	goto L15
L12:
	;
	v112 = v31
	v116 = v18
	goto L13
L13:
	;
	if base.Ui32(v112) <= base.Ui32(v116) {
		goto L34
	} else {
		goto L35
	}
L14:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+256)) = uint8(v102)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+252)) = v67
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+260)) = v105
	v109 = base.B2i32(base.Ui32(int32(31)) < base.Ui32(v54-v16))
	if v46&v109 != 0 {
		goto L10
	} else {
		goto L30
	}
L15:
	;
	v54 = v43
	goto L17
L16:
	;
	v102 = v91 ^ int32(1)
	goto L14
L17:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v67 = v54 + int32(1)
	v70 = F_visibilitymap_get_status(m, v65, v67, v14+int32(12))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L6
	} else {
		goto L19
	}
L18:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	if v91 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L19:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+264))
	if base.Ui32(v72) <= base.Ui32(v67) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+272))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+276)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(l1)+264)) = v72 + int32(_a_F_heap_vac_scan_next_block_0)
	goto L22
L21:
	;
	goto L22
L22:
	;
	v79 = int32(0)
	if base.B2i32(v70&int32(1) == v79)|base.B2i32(v54 == v19-int32(2)) != 0 {
		v102 = v79
		goto L14
	} else {
		goto L23
	}
L23:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	if v86 != int32(1) {
		v102 = v79
		goto L14
	} else {
		goto L24
	}
L24:
	;
	if v70&int32(2) != 0 {
		v54 = v67
		goto L17
	} else {
		goto L25
	}
L25:
	;
	goto L18
L26:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+276))
	if v95 == int32(0) {
		v43 = v67
		v46 = int32(1)
		goto L15
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	goto L16
L29:
	;
	goto L28
L30:
	;
	if base.Ui32(int32(31)) < base.Ui32(v54-v16) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v111 = v67
	goto L33
L32:
	;
	v111 = v18
	goto L33
L33:
	;
	v112 = v67
	v116 = v111
	goto L13
L34:
	;
	v129 = v116
	goto L9
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+248)) = v116
	v125 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v125)
	goto L8
}
func F_heap_xlog_update(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int64
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v471 int32
	_ = v471
	var v477 int32
	_ = v477
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int64
	_ = v606
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v636 int32
	_ = v636
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v709 int32
	_ = v709
	var v718 int32
	_ = v718
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v739 int32
	_ = v739
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v796 int64
	_ = v796
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v831 int32
	_ = v831
	v3 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(_a_F_heap_xlog_update_0)
	m.G0 = v25
	v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+64))
	F_XLogRecGetBlockTag(m, l0, v3, v25+int32(_a_F_heap_xlog_update_1), v3, v25+int32(_a_F_heap_xlog_update_2))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v39 = int32(0)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+72))
	if v46 < int32(1) {
		v70 = v39
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_heap_xlog_update[0])))
	if v70 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L4:
	;
	goto L3
L5:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+int32(52))+76)))
	if v51 != int32(1) {
		v70 = v39
		goto L4
	} else {
		goto L6
	}
L6:
	;
	goto L8
L8:
	;
	goto L9
L9:
	;
	goto L11
L11:
	;
	goto L12
L12:
	;
	if v25+int32(_a_F_heap_xlog_update_3) != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v45+int32(128))+20))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_heap_xlog_update[1]))) = v62
	goto L15
L14:
	;
	goto L15
L15:
	;
	v70 = int32(1)
	goto L4
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_heap_xlog_update[1]))) = v71
	goto L19
L18:
	;
	goto L19
L19:
	;
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+12)))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+72))
	if int32(3) <= v77 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	if v87&int32(1) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L21:
	;
	if v87&int32(1) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+180)))
	v86 = v84
	v87 = v85
	goto L21
L23:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+232)))
	v84 = v80
	goto L22
L24:
	;
	goto L25
L25:
	;
	if v77 != int32(2) {
		v86 = v3
		v87 = int32(0)
		goto L21
	} else {
		goto L26
	}
L26:
	;
	v84 = v3
	goto L22
L27:
	;
	v92 = int32(0)
	if v86&int32(1) == v92 {
		v105 = v92
		goto L20
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_heap_xlog_update[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v97
	v99 = *(*int64)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_heap_xlog_update[3])))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+16)) = v99
	v103 = F_CreateFakeRelcacheEntry(m, v25+int32(16))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	v105 = v103
	goto L20
L32:
	;
	if v86&int32(1) == int32(0) {
		goto L58
	} else {
		goto L59
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = int32(0)
	v115 = F_XLogReadBufferForRedo(m, l0, int32(2), v25+int32(32))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	if v115 != 0 {
		v192 = v117
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if v192 == int32(0) {
		goto L32
	} else {
		goto L56
	}
L36:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+7)))
	if v118&int32(1) == int32(0) {
		v163 = v117
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_heap_xlog_update[0])))
	v166 = F_visibilitymap_clear(m, v164, v163, int32(3))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L50
	}
L38:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_heap_xlog_update[1])))
	if v117 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	if v131 == int32(0) {
		v163 = v132
		goto L37
	} else {
		goto L43
	}
L40:
	;
	v131 = int32(0)
	goto L39
L41:
	;
	goto L42
L42:
	;
	v127 = F_BufferGetBlockNumber(m, v117)
	mBase = m.M
	v129 = base.I32_div_u_s(v123, int32(_a_F_heap_xlog_update_4))
	v131 = base.B2i32(v127 == v129)
	goto L39
L43:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_heap_xlog_update[1])))
	v137 = F_visibilitymap_clear(m, v135, v132, int32(3))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	if v137 == int32(0) {
		v163 = v139
		goto L37
	} else {
		goto L45
	}
L45:
	;
	if v139 < int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v159))) = base.I64_rotl(v27, int64(32))
	v163 = v139
	goto L37
L47:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_heap_xlog_update[4]))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v145+(v139^int32(-1))<<(uint(int32(2))%32))))
	v159 = v151
	goto L46
L48:
	;
	goto L49
L49:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _c_F_heap_xlog_update[5]))
	v159 = v153 + v139<<(uint(int32(13))%32) + int32(-8192)
	goto L46
L50:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	if v166 == int32(0) {
		v192 = v168
		goto L35
	} else {
		goto L51
	}
L51:
	;
	if v168 < int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v188))) = base.I64_rotl(v27, int64(32))
	v192 = v168
	goto L35
L53:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_heap_xlog_update[4]))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v174+(v168^int32(-1))<<(uint(int32(2))%32))))
	v188 = v180
	goto L52
L54:
	;
	goto L55
L55:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _c_F_heap_xlog_update[5]))
	v188 = v182 + v168<<(uint(int32(13))%32) + int32(-8192)
	goto L52
L56:
	;
	F_UnlockReleaseBuffer(m, v192)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	goto L32
L58:
	;
	if v105 != 0 {
		goto L71
	} else {
		goto L72
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = int32(0)
	v208 = F_XLogReadBufferForRedo(m, l0, int32(3), v25+int32(32))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	if v208 != 0 {
		v239 = v210
		goto L61
	} else {
		goto L62
	}
L61:
	;
	if v239 == int32(0) {
		goto L58
	} else {
		goto L69
	}
L62:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_heap_xlog_update[1])))
	v213 = F_visibilitymap_clear(m, v211, v210, int32(3))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	if v213 == int32(0) {
		v239 = v215
		goto L61
	} else {
		goto L64
	}
L64:
	;
	if v215 < int32(0) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v235))) = base.I64_rotl(v27, int64(32))
	v239 = v215
	goto L61
L66:
	;
	v221 = *(*int32)(unsafe.Add(mBase, _c_F_heap_xlog_update[4]))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v221+(v215^int32(-1))<<(uint(int32(2))%32))))
	v235 = v227
	goto L65
L67:
	;
	goto L68
L68:
	;
	v229 = *(*int32)(unsafe.Add(mBase, _c_F_heap_xlog_update[5]))
	v235 = v229 + v215<<(uint(int32(13))%32) + int32(-8192)
	goto L65
L69:
	;
	F_UnlockReleaseBuffer(m, v239)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	goto L58
L71:
	;
	F_pfree(m, v105)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v249 = int32(base.Ui32(v71) >> (uint(int32(16)) % 32))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_heap_xlog_update[1])))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_heap_xlog_update[0])))
	v256 = F_XLogReadBufferForRedo(m, l0, base.B2i32(v251 != v252), v25+int32(_a_F_heap_xlog_update_5))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L81
	}
L74:
	;
	goto L73
L75:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L1
	} else {
		goto L242
	}
L76:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L1
	} else {
		goto L239
	}
L77:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_heap_xlog_update[6])))
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_heap_xlog_update[7])))
	if v779 != 0 {
		goto L226
	} else {
		goto L227
	}
L78:
	;
	v519 = int32(0)
	v521 = v25 + int32(28)
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v523)+72))
	if v524 < v519 {
		v546 = v519
		goto L148
	} else {
		goto L149
	}
L79:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L144
	}
L80:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L141
	}
L81:
	;
	if v256 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_heap_xlog_update[6])))
	if v260 < int32(0) {
		goto L86
	} else {
		goto L87
	}
L83:
	;
	v392 = int32(0)
	v398 = v3
	goto L84
L84:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_heap_xlog_update[1])))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_heap_xlog_update[0])))
	if v399 == v400 {
		goto L117
	} else {
		goto L118
	}
L85:
	;
	v279 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+4)))
	if v279 == int32(0) {
		goto L80
	} else {
		goto L89
	}
L86:
	;
	v264 = *(*int32)(unsafe.Add(mBase, _c_F_heap_xlog_update[4]))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v264+(v260^int32(-1))<<(uint(int32(2))%32))))
	v278 = v270
	goto L85
L87:
	;
	goto L88
L88:
	;
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_heap_xlog_update[5]))
	v278 = v272 + v260<<(uint(int32(13))%32) + int32(-8192)
	goto L85
L89:
	;
	v282 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v278)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v282) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v290 = int32(base.Ui32(v282+int32(_a_F_heap_xlog_update_6)) >> (uint(int32(2)) % 32))
	goto L92
L91:
	;
	v290 = int32(0)
	goto L92
L92:
	;
	if base.Ui32(v290&int32(_a_F_heap_xlog_update_7)) < base.Ui32(v279) {
		goto L80
	} else {
		goto L93
	}
L93:
	;
	v295 = v278 + int32(20)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v295+v279<<(uint(int32(2))%32))))
	if v299&int32(_a_F_heap_xlog_update_8) != int32(_a_F_heap_xlog_update_9) {
		goto L79
	} else {
		goto L94
	}
L94:
	;
	v306 = v278 + v299&int32(_a_F_heap_xlog_update_10)
	v307 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v306)+20)))
	v309 = v307 & int32(_a_F_heap_xlog_update_11)
	*(*uint16)(unsafe.Add(mBase, uint32(v306)+20)) = uint16(v309)
	v311 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v306)+18)))
	v313 = v311 & int32(-24577)
	if l1 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v316 = v313 | int32(_a_F_heap_xlog_update_12)
	goto L97
L96:
	;
	v316 = v313
	goto L97
L97:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v306)+18)) = uint16(v316)
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v306)+18)) = uint16(v316)
	*(*uint16)(unsafe.Add(mBase, uint32(v306)+20)) = uint16(v309)
	v321 = int32(1)
	v340 = v318<<(uint(v321)%32)&int32(16) | (v318<<(uint(int32(4))%32)&int32(64) | (v318<<(uint(int32(6))%32)&int32(128) | v318&v321<<(uint(int32(12))%32))) | v309
	if v318&int32(15) != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v306)+20)) = uint16(v340)
	goto L100
L99:
	;
	goto L100
L100:
	;
	if v318&int32(16) != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v347 = v316 | int32(_a_F_heap_xlog_update_13)
	*(*uint16)(unsafe.Add(mBase, uint32(v306)+18)) = uint16(v347)
	goto L103
L102:
	;
	goto L103
L103:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v351 = v340 & int32(_a_F_heap_xlog_update_14)
	*(*uint16)(unsafe.Add(mBase, uint32(v306)+20)) = uint16(v351)
	v353 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v306)+8)) = v353
	*(*int32)(unsafe.Add(mBase, uint32(v306)+4)) = v349
	*(*uint16)(unsafe.Add(mBase, uint32(v306)+16)) = uint16(v75)
	*(*uint16)(unsafe.Add(mBase, uint32(v306)+14)) = uint16(v71)
	*(*uint16)(unsafe.Add(mBase, uint32(v306)+12)) = uint16(v249)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)+36))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	if v361 == v353 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+7)))
	if v376&int32(1) != 0 {
		goto L112
	} else {
		goto L113
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v295))) = v360
	goto L104
L106:
	;
	v364 = int32(3)
	if base.B2i32(base.Ui32(v361) < base.Ui32(v364))|base.B2i32(base.Ui32(v360) < base.Ui32(v364)) == int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	if v360-v361 < int32(0) {
		goto L105
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	if base.Ui32(v361) <= base.Ui32(v360) {
		goto L104
	} else {
		goto L111
	}
L110:
	;
	goto L104
L111:
	;
	goto L105
L112:
	;
	v379 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v278)+10)))
	v381 = v379 & int32(_a_F_heap_xlog_update_15)
	*(*uint16)(unsafe.Add(mBase, uint32(v278)+10)) = uint16(v381)
	goto L114
L113:
	;
	goto L114
L114:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v278))) = base.I64_rotl(v27, int64(32))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_heap_xlog_update[6])))
	F_MarkBufferDirty(m, v388)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v392 = v306
	v398 = int32(base.Ui32(v299) >> (uint(int32(17)) % 32))
	goto L84
L116:
	;
	if v485 == int32(0) {
		goto L78
	} else {
		goto L140
	}
L117:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_heap_xlog_update[6])))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_heap_xlog_update[7]))) = v402
	v485 = v256
	goto L116
L118:
	;
	goto L119
L119:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v405 = int32(*(*int8)(unsafe.Add(mBase, uint32(v404)+48)))
	if v405 < int32(0) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v409 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v483 = F_XLogReadBufferForRedo(m, l0, int32(0), v25+int32(_a_F_heap_xlog_update_16))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L139
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_heap_xlog_update[7]))) = v409
	if v409 < int32(0) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v430 = int32(_a_F_heap_xlog_update_13)
	v431 = int32(0)
	if v431|(v429&int32(3)|int32(1)) == v431 {
		goto L130
	} else {
		goto L131
	}
L125:
	;
	v415 = *(*int32)(unsafe.Add(mBase, _c_F_heap_xlog_update[4]))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v415+(v409^int32(-1))<<(uint(int32(2))%32))))
	v429 = v421
	goto L124
L126:
	;
	goto L127
L127:
	;
	v423 = *(*int32)(unsafe.Add(mBase, _c_F_heap_xlog_update[5]))
	v429 = v423 + v409<<(uint(int32(13))%32) + int32(-8192)
	goto L124
L128:
	;
	goto L78
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v429)+10)) = int32(_a_F_heap_xlog_update_17)
	v471 = int32(_a_F_heap_xlog_update_16)
	*(*uint16)(unsafe.Add(mBase, uint32(v429)+18)) = uint16(v471)
	v477 = int32(_a_F_heap_xlog_update_13)
	*(*uint16)(unsafe.Add(mBase, uint32(v429)+16)) = uint16(v477)
	*(*uint16)(unsafe.Add(mBase, uint32(v429)+14)) = uint16(v477)
	goto L128
L130:
	;
	goto L133
L131:
	;
	goto L132
L132:
	;
	goto L138
L133:
	;
	v448 = v429 + v430
	v450 = v429 + int32(4)
	if base.Ui32(v450) < base.Ui32(v448) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v452 = v448
	goto L136
L135:
	;
	v452 = v450
	goto L136
L136:
	;
	v457 = (v429^int32(-1)+v452)&int32(-4) + int32(4)
	if v457 == int32(0) {
		goto L129
	} else {
		goto L137
	}
L137:
	;
	base.MemoryFill(m, v429, int32(0), v457)
	goto L129
L138:
	;
	base.MemoryFill(m, v429, int32(0), v430)
	goto L129
L139:
	;
	v485 = v483
	goto L116
L140:
	;
	v764 = int32(1)
	v766 = int32(0)
	goto L77
L141:
	;
	F_errmsg_internal(m, int32(_a_F_heap_xlog_update_18), int32(0))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F_heap_xlog_update_19), int32(895), int32(_a_F_heap_xlog_update_20))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L144:
	;
	F_errmsg_internal(m, int32(_a_F_heap_xlog_update_21), int32(0))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(_a_F_heap_xlog_update_19), int32(898), int32(_a_F_heap_xlog_update_20))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
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
	v550 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+12)))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_heap_xlog_update[7])))
	if v552 < int32(0) {
		goto L159
	} else {
		goto L160
	}
L148:
	;
	v549 = v546
	goto L147
L149:
	;
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v523+int32(0))+76)))
	if v529 != int32(1) {
		v546 = v519
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v533 = v523 + int32(76)
	v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+43)))
	if v534 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	if v521 == int32(0) {
		v546 = v519
		goto L148
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	if v521 != 0 {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v539 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v521))) = v539
	v549 = v539
	goto L147
L155:
	;
	v542 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v533)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v521))) = v542
	goto L157
L156:
	;
	goto L157
L157:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v533)+44))
	v546 = v544
	goto L148
L158:
	;
	v571 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v570)+12)))
	if base.Ui32(v571) < base.Ui32(int32(25)) {
		goto L162
	} else {
		goto L163
	}
L159:
	;
	v556 = *(*int32)(unsafe.Add(mBase, _c_F_heap_xlog_update[4]))
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v556+(v552^int32(-1))<<(uint(int32(2))%32))))
	v570 = v562
	goto L158
L160:
	;
	goto L161
L161:
	;
	v564 = *(*int32)(unsafe.Add(mBase, _c_F_heap_xlog_update[5]))
	v570 = v564 + v552<<(uint(int32(13))%32) + int32(-8192)
	goto L158
L162:
	;
	v580 = int32(1)
	goto L164
L163:
	;
	v580 = int32(base.Ui32(v571+int32(_a_F_heap_xlog_update_6))>>(uint(int32(2))%32)) + int32(1)
	goto L164
L164:
	;
	if base.Ui32(v580&int32(_a_F_heap_xlog_update_7)) < base.Ui32(v550) {
		goto L76
	} else {
		goto L165
	}
L165:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	v585 = int32(0)
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+7)))
	if v587&int32(32) != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v590 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v549))))
	v593 = v549 + int32(2)
	v594 = v590
	goto L168
L167:
	;
	v593 = v549
	v594 = v585
	goto L168
L168:
	;
	if v587&int32(64) != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v598 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v593))))
	v601 = v593 + int32(2)
	v602 = v598
	goto L171
L170:
	;
	v601 = v593
	v602 = v585
	goto L171
L171:
	;
	v603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601)+4)))
	v604 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v601)+2)))
	v605 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v601))))
	v606 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+47)) = v606
	*(*int64)(unsafe.Add(mBase, uint32(v25)+40)) = v606
	*(*int64)(unsafe.Add(mBase, uint32(v25)+32)) = v606
	v613 = v601 + int32(5)
	v614 = v549 + v584 - v613
	v616 = v25 + int32(55)
	if v594 != 0 {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	if v602 != 0 {
		goto L188
	} else {
		goto L189
	}
L173:
	;
	v618 = v603 - int32(23)
	if v618 != 0 {
		goto L176
	} else {
		goto L177
	}
L174:
	;
	goto L175
L175:
	;
	if v614 != 0 {
		goto L185
	} else {
		goto L186
	}
L176:
	;
	base.MemoryCopy(m, v616, v613, v618)
	goto L178
L177:
	;
	goto L178
L178:
	;
	v622 = v25 + int32(32) + v603
	if v594 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+22)))
	base.MemoryCopy(m, v622, v392+v623, v594)
	goto L181
L180:
	;
	goto L181
L181:
	;
	v626 = v594 + v622
	v627 = v614 - v618
	if v627 != 0 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	base.MemoryCopy(m, v626, v613+v618, v627)
	goto L184
L183:
	;
	goto L184
L184:
	;
	v636 = v626 + v627
	goto L172
L185:
	;
	base.MemoryCopy(m, v616, v613, v614)
	goto L187
L186:
	;
	goto L187
L187:
	;
	v636 = v614 + v616
	goto L172
L188:
	;
	base.MemoryCopy(m, v636, v392+v398-v602, v602)
	goto L190
L189:
	;
	goto L190
L190:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+54)) = uint8(v603)
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+50)) = uint16(v605)
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v642)+36))
	v645 = v604 & int32(_a_F_heap_xlog_update_22)
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+52)) = uint16(v645)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+40)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v643
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+48)) = uint16(v75)
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+46)) = uint16(v71)
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+44)) = uint16(v249)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v650
	v662 = F_PageAddItemExtended(m, v570, v25+int32(32), v594+v602+v614+int32(23), v550, int32(3))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	if v662 == int32(0) {
		goto L75
	} else {
		goto L192
	}
L192:
	;
	v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+7)))
	if v666&int32(2) != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v669 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v570)+10)))
	v671 = v669 & int32(_a_F_heap_xlog_update_15)
	*(*uint16)(unsafe.Add(mBase, uint32(v570)+10)) = uint16(v671)
	goto L195
L194:
	;
	goto L195
L195:
	;
	v676 = int32(4)
	v677 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v570)+14)))
	v678 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v570)+12)))
	v679 = v677 - v678
	if v679 <= v676 {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v570))) = base.I64_rotl(v27, int64(32))
	v743 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v743)+36))
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v570)+20))
	if v745 == int32(0) {
		goto L216
	} else {
		goto L217
	}
L197:
	;
	v682 = v676
	goto L199
L198:
	;
	v682 = v679
	goto L199
L199:
	;
	v684 = v682 - int32(4)
	if v684 == int32(0) {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v739 = int32(0)
	goto L196
L201:
	;
	goto L202
L202:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v678) {
		goto L204
	} else {
		goto L205
	}
L203:
	;
	v739 = v684
	goto L196
L204:
	;
	v695 = int32(base.Ui32(v678+int32(_a_F_heap_xlog_update_6)) >> (uint(int32(2)) % 32))
	goto L206
L205:
	;
	v695 = int32(0)
	goto L206
L206:
	;
	if base.Ui32(v695&int32(_a_F_heap_xlog_update_7)) < base.Ui32(int32(291)) {
		goto L203
	} else {
		goto L207
	}
L207:
	;
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v570)+10)))
	if v700&int32(1) == int32(0) {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v739 = int32(0)
	goto L196
L209:
	;
	goto L210
L210:
	;
	v709 = int32(1)
	goto L211
L211:
	;
	v718 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v570+int32(20)+v709&int32(_a_F_heap_xlog_update_7)<<(uint(int32(2))%32))+1)))
	if v718&int32(384) == int32(0) {
		goto L203
	} else {
		goto L213
	}
L212:
	;
	v739 = int32(0)
	goto L196
L213:
	;
	v724 = v709 + int32(1)
	v725 = int32(_a_F_heap_xlog_update_7)
	if base.Ui32(v724&v725) <= base.Ui32(v695&v725) {
		v709 = v724
		goto L211
	} else {
		goto L214
	}
L214:
	;
	goto L212
L215:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_heap_xlog_update[7])))
	F_MarkBufferDirty(m, v760)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L1
	} else {
		goto L223
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v570)+20)) = v744
	goto L215
L217:
	;
	v748 = int32(3)
	if base.B2i32(base.Ui32(v745) < base.Ui32(v748))|base.B2i32(base.Ui32(v744) < base.Ui32(v748)) == int32(0) {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	if v744-v745 < int32(0) {
		goto L216
	} else {
		goto L221
	}
L219:
	;
	goto L220
L220:
	;
	if base.Ui32(v745) <= base.Ui32(v744) {
		goto L215
	} else {
		goto L222
	}
L221:
	;
	goto L215
L222:
	;
	goto L216
L223:
	;
	v764 = l1
	v766 = v739
	goto L77
L224:
	;
	if v764|base.B2i32(base.Ui32(int32(1637)) < base.Ui32(v766)) == int32(0) {
		goto L235
	} else {
		goto L236
	}
L225:
	;
	F_UnlockReleaseBuffer(m, v787)
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L1
	} else {
		goto L234
	}
L226:
	;
	if v778 == v779 {
		goto L229
	} else {
		goto L230
	}
L227:
	;
	v784 = v778
	goto L228
L228:
	;
	if v784 == int32(0) {
		goto L224
	} else {
		goto L233
	}
L229:
	;
	v787 = v779
	goto L225
L230:
	;
	goto L231
L231:
	;
	F_UnlockReleaseBuffer(m, v779)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_heap_xlog_update[6])))
	v784 = v783
	goto L228
L233:
	;
	v787 = v784
	goto L225
L234:
	;
	goto L224
L235:
	;
	v796 = *(*int64)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_heap_xlog_update[3])))
	*(*int64)(unsafe.Add(mBase, uint32(v25))) = v796
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_heap_xlog_update[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v798
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_heap_xlog_update[0])))
	F_XLogRecordPageWithFreeSpace(m, v25, v800, v766)
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L1
	} else {
		goto L238
	}
L236:
	;
	goto L237
L237:
	;
	m.G0 = v25 + int32(_a_F_heap_xlog_update_0)
	return
L238:
	;
	goto L237
L239:
	;
	F_errmsg_internal(m, int32(_a_F_heap_xlog_update_23), int32(0))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	F_errfinish(m, int32(_a_F_heap_xlog_update_19), int32(963), int32(_a_F_heap_xlog_update_20))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L242:
	;
	F_errmsg_internal(m, int32(_a_F_heap_xlog_update_24), int32(0))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	F_errfinish(m, int32(_a_F_heap_xlog_update_19), int32(1041), int32(_a_F_heap_xlog_update_20))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_heap_xlog_vm_clear(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v13
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v15
	v17 = F_CreateFakeRelcacheEntry(m, v10)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		v19 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v19
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+72))
		if v22 <= v19 {
			F_pfree(m, v17)
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return
			} else {
				m.G0 = v10 + int32(16)
				return
			}
		} else {
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+128)))
			if v25 != int32(1) {
				F_pfree(m, v17)
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return
				} else {
					m.G0 = v10 + int32(16)
					return
				}
			} else {
				v31 = F_XLogReadBufferForRedo(m, l0, int32(1), v10+int32(12))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
					if v31 != 0 {
						v60 = v33
						if v60 == int32(0) {
							F_pfree(m, v17)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return
							} else {
								m.G0 = v10 + int32(16)
								return
							}
						} else {
							F_UnlockReleaseBuffer(m, v60)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								F_pfree(m, v17)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									m.G0 = v10 + int32(16)
									return
								}
							}
						}
					} else {
						v34 = F_visibilitymap_clear(m, l2, v33, l3)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
							if v34 == int32(0) {
								v60 = v36
							} else {
								if v36 < int32(0) {
									v42 = *(*int32)(unsafe.Add(mBase, _c_F_heap_xlog_vm_clear[0]))
									v48 = *(*int32)(unsafe.Add(mBase, uint32(v42+(v36^int32(-1))<<(uint(int32(2))%32))))
									v56 = v48
								} else {
									v50 = *(*int32)(unsafe.Add(mBase, _c_F_heap_xlog_vm_clear[1]))
									v56 = v50 + v36<<(uint(int32(13))%32) + int32(-8192)
								}
								*(*int64)(unsafe.Add(mBase, uint32(v56))) = base.I64_rotl(v12, int64(32))
								v60 = v36
							}
							if v60 == int32(0) {
								F_pfree(m, v17)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									m.G0 = v10 + int32(16)
									return
								}
							} else {
								F_UnlockReleaseBuffer(m, v60)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return
								} else {
									F_pfree(m, v17)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										m.G0 = v10 + int32(16)
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
