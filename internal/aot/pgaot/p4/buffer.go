package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BufferGetLSNAtomic(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v47 int32
	_ = v47
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v116 int64
	_ = v116
	var v117 int64
	_ = v117
	var v130 int64
	_ = v130
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	if l0 < int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_BufferGetLSNAtomic[0]))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+252))
	goto L5
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_BufferGetLSNAtomic[1]))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v14+(l0^int32(-1))<<(uint(int32(2))%32))))
	v28 = v20
	goto L1
L3:
	;
	goto L4
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_BufferGetLSNAtomic[2]))
	v28 = v22 + l0<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	if int32(0) <= l0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	m.G0 = v9 + int32(32)
	return v130
L7:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_BufferGetLSNAtomic[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = int32(_a_F_BufferGetLSNAtomic_0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(_a_F_BufferGetLSNAtomic_1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(_a_F_BufferGetLSNAtomic_2)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(0)
	v62 = v47 + l0<<(uint(int32(6))%32) - int32(40)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v64 = int32(_a_F_BufferGetLSNAtomic_3)
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v63 | v64
	if v63&v64 != 0 {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BufferGetLSNAtomic[4])))
	if (base.B2i32(v31 != int32(0))|v37)&int32(1) != 0 {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v41 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v28)+4)))
	v42 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v28))))
	v130 = v41 | v42<<(uint(int64(32))%64)
	goto L6
L11:
	;
	goto L10
L12:
	;
	goto L15
L13:
	;
	v89 = v63
	goto L14
L14:
	;
	v96 = int32(_a_F_BufferGetLSNAtomic_4)
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_BufferGetLSNAtomic[5]))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(8))+8))
	if v99 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L15:
	;
	F_perform_spin_delay(m, v9+int32(8))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v89 = v81
	goto L14
L17:
	;
	return int64(0)
L18:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v82 = int32(_a_F_BufferGetLSNAtomic_3)
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v81 | v82
	if v81&v82 != 0 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v116 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v28)+4)))
	v117 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v28))))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v89 & int32(-4194305)
	v130 = v116 | v117<<(uint(int64(32))%64)
	goto L6
L21:
	;
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BufferGetLSNAtomic[5])) = v114
	goto L21
L23:
	;
	if int32(999) < v97 {
		goto L21
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if v97 < int32(11) {
		goto L21
	} else {
		goto L30
	}
L26:
	;
	v104 = int32(900)
	if v104 <= v97 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v107 = v104
	goto L29
L28:
	;
	v107 = v97
	goto L29
L29:
	;
	v114 = v107 + int32(100)
	goto L22
L30:
	;
	v114 = v97 - int32(1)
	goto L22
}
func F_IsBufferCleanupOK(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v148 int32
	_ = v148
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if l0 < v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v7 + int32(32)
	return v148
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_IsBufferCleanupOK[0]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v12+(l0^int32(-1))<<(uint(int32(2))%32))))
	v148 = base.B2i32(v18 == int32(1))
	goto L1
L3:
	;
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l0
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_IsBufferCleanupOK[1]))
	if l0 == v23 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if v71 != int32(1) {
		v148 = v2
		goto L1
	} else {
		goto L34
	}
L6:
	;
	v70 = int32(_a_F_IsBufferCleanupOK_0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_IsBufferCleanupOK[2]))
	if l0 == v27 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v70 = int32(_a_F_IsBufferCleanupOK_1)
	goto L5
L10:
	;
	goto L11
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_IsBufferCleanupOK[3]))
	if l0 == v31 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v70 = int32(_a_F_IsBufferCleanupOK_2)
	goto L5
L13:
	;
	goto L14
L14:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_IsBufferCleanupOK[4]))
	if l0 == v35 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v70 = int32(_a_F_IsBufferCleanupOK_3)
	goto L5
L16:
	;
	goto L17
L17:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_IsBufferCleanupOK[5]))
	if l0 == v39 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v70 = int32(_a_F_IsBufferCleanupOK_4)
	goto L5
L19:
	;
	goto L20
L20:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_IsBufferCleanupOK[6]))
	if l0 == v43 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v70 = int32(_a_F_IsBufferCleanupOK_5)
	goto L5
L22:
	;
	goto L23
L23:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_IsBufferCleanupOK[7]))
	if l0 == v47 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v70 = int32(_a_F_IsBufferCleanupOK_6)
	goto L5
L25:
	;
	goto L26
L26:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_IsBufferCleanupOK[8]))
	if l0 == v51 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v70 = int32(_a_F_IsBufferCleanupOK_7)
	goto L5
L28:
	;
	goto L29
L29:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_IsBufferCleanupOK[9]))
	if v55 == int32(0) {
		v148 = v2
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_IsBufferCleanupOK[10]))
	v62 = int32(0)
	v64 = F_hash_search(m, v59, v7+int32(8), v62, v62)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	return int32(0)
L32:
	;
	if v64 == int32(0) {
		v148 = v2
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v70 = v64
	goto L5
L34:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_IsBufferCleanupOK[11]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(_a_F_IsBufferCleanupOK_8)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = int32(_a_F_IsBufferCleanupOK_9)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = int32(_a_F_IsBufferCleanupOK_10)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = int64(0)
	v90 = v75 + l0<<(uint(int32(6))%32) - int32(40)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v92 = int32(_a_F_IsBufferCleanupOK_11)
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = v91 | v92
	if v91&v92 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	goto L38
L36:
	;
	v111 = v91
	goto L37
L37:
	;
	v118 = int32(_a_F_IsBufferCleanupOK_12)
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_IsBufferCleanupOK[12]))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(8))+8))
	if v121 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L38:
	;
	F_perform_spin_delay(m, v7+int32(8))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L31
	} else {
		goto L40
	}
L39:
	;
	v111 = v105
	goto L37
L40:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v106 = int32(_a_F_IsBufferCleanupOK_11)
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = v105 | v106
	if v105&v106 != 0 {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = v111 & int32(-4194305)
	v148 = base.B2i32(v111&int32(_a_F_IsBufferCleanupOK_13) == int32(1))
	goto L1
L43:
	;
	goto L42
L44:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IsBufferCleanupOK[12])) = v136
	goto L43
L45:
	;
	if int32(999) < v119 {
		goto L43
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	if v119 < int32(11) {
		goto L43
	} else {
		goto L52
	}
L48:
	;
	v126 = int32(900)
	if v126 <= v119 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v129 = v126
	goto L51
L50:
	;
	v129 = v119
	goto L51
L51:
	;
	v136 = v129 + int32(100)
	goto L44
L52:
	;
	v136 = v119 - int32(1)
	goto L44
}
func F_LockBufferForCleanup(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int64
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int64
	_ = v133
	var v134 int64
	_ = v134
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int64
	_ = v186
	var v187 int64
	_ = v187
	var v195 int64
	_ = v195
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int64
	_ = v227
	var v228 int64
	_ = v228
	var v237 int32
	_ = v237
	var v238 int64
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v250 int64
	_ = v250
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v265 int64
	_ = v265
	var v272 int32
	_ = v272
	var v275 int64
	_ = v275
	var v281 int64
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int64
	_ = v290
	var v291 int64
	_ = v291
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v493 int32
	_ = v493
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v512 int64
	_ = v512
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v553 int32
	_ = v553
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v584 int32
	_ = v584
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	F_CheckBufferIsPinnedOnce(m, l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l0 < int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v14 + int32(32)
	return
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[0]))
	v23 = l0 - int32(1)
	v26 = v21 + v23<<(uint(int32(6))%32)
	v28 = v21
	v32 = int32(0)
	v35 = int64(0)
	goto L5
L5:
	;
	v39 = l0 << (uint(int32(6)) % 32)
	v44 = F_LWLockAcquire(m, v28+v39-int32(16), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v84 & int32(-4194305)
	F_LockBuffer(m, l0, int32(0))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L1
	} else {
		goto L106
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = int32(_a_F_LockBufferForCleanup_0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = int32(_a_F_LockBufferForCleanup_1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = int32(_a_F_LockBufferForCleanup_2)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = int64(0)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	v57 = int32(_a_F_LockBufferForCleanup_3)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v56 | v57
	if v56&v57 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	goto L11
L9:
	;
	v84 = v56
	goto L10
L10:
	;
	v97 = int32(_a_F_LockBufferForCleanup_4)
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[1]))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(8))+8))
	if v100 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L11:
	;
	F_perform_spin_delay(m, v14+int32(8))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	v84 = v77
	goto L10
L13:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	v78 = int32(_a_F_LockBufferForCleanup_3)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v77 | v78
	if v77&v78 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	if v84&int32(_a_F_LockBufferForCleanup_5) == int32(1) {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[1])) = v115
	goto L16
L18:
	;
	if int32(999) < v98 {
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v98 < int32(11) {
		goto L16
	} else {
		goto L25
	}
L21:
	;
	v105 = int32(900)
	if v105 <= v98 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v108 = v105
	goto L24
L23:
	;
	v108 = v98
	goto L24
L24:
	;
	v115 = v108 + int32(100)
	goto L17
L25:
	;
	v115 = v98 - int32(1)
	goto L17
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v84 & int32(-4456447)
	if v32 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	if v84&int32(536870912) == int32(0) {
		goto L34
	} else {
		goto L35
	}
L29:
	;
	v128 = m.G0
	v129 = int32(16)
	v130 = v128 - v129
	m.G0 = v130
	F_gettimeofday(m, v130)
	mBase = m.M
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v130)))
	v134 = int64(*(*int32)(unsafe.Add(mBase, uint32(v130)+8)))
	m.G0 = v130 + v129
	goto L32
L30:
	;
	goto L31
L31:
	;
	goto L3
L32:
	;
	v143 = int32(0)
	F_LogRecoveryConflict(m, int32(12), v35, v134+v133*int64(1000000)-int64(946684800000000), v143, v143)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v152
	*(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[3])) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v84&int32(-541065217) | int32(536870912)
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[0]))
	F_LWLockRelease(m, v162+v39-int32(16))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	goto L6
L37:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[4]))
	if base.Ui32(int32(2)) <= base.Ui32(v169) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = int32(_a_F_LockBufferForCleanup_0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = int32(_a_F_LockBufferForCleanup_1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = int32(_a_F_LockBufferForCleanup_2)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = int64(0)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	v526 = int32(_a_F_LockBufferForCleanup_3)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v525 | v526
	if v525&v526 != 0 {
		goto L82
	} else {
		goto L83
	}
L39:
	;
	if v32|base.B2i32(v35 == int64(0)) == int32(0) {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	goto L41
L41:
	;
	F_ProcWaitForSignal(m, int32(67108864))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L81
	}
L42:
	;
	v241 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v241)+72)) = v23
	goto L53
L43:
	;
	v181 = m.G0
	v182 = int32(16)
	v183 = v181 - v182
	m.G0 = v183
	F_gettimeofday(m, v183)
	mBase = m.M
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v183)))
	v187 = int64(*(*int32)(unsafe.Add(mBase, uint32(v183)+8)))
	m.G0 = v183 + v182
	v195 = v187 + v186*int64(1000000) - int64(946684800000000)
	goto L46
L44:
	;
	goto L45
L45:
	;
	if v35 != int64(0) {
		v237 = v32
		v238 = v35
		goto L42
	} else {
		goto L50
	}
L46:
	;
	v197 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[6]))
	goto L47
L47:
	;
	if base.B2i32(base.I64_extend_i32_s(v197)*int64(1000) <= v195-v35) == int32(0) {
		v237 = int32(0)
		v238 = v35
		goto L42
	} else {
		goto L48
	}
L48:
	;
	v205 = int32(1)
	F_LogRecoveryConflict(m, int32(12), v35, v195, int32(0), v205)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v237 = v205
	v238 = v35
	goto L42
L50:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[7])))
	if v214&int32(1) == int32(0) {
		v237 = v32
		v238 = v35
		goto L42
	} else {
		goto L51
	}
L51:
	;
	v222 = m.G0
	v223 = int32(16)
	v224 = v222 - v223
	m.G0 = v224
	F_gettimeofday(m, v224)
	mBase = m.M
	v227 = *(*int64)(unsafe.Add(mBase, uint32(v224)))
	v228 = int64(*(*int32)(unsafe.Add(mBase, uint32(v224)+8)))
	m.G0 = v224 + v223
	goto L52
L52:
	;
	v237 = v32
	v238 = v228 + v227*int64(1000000) - int64(946684800000000)
	goto L42
L53:
	;
	v243 = m.G0
	v245 = v243 + int32(-64)
	m.G0 = v245
	v250 = *(*int64)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[8]))
	*(*int64)(unsafe.Add(mBase, uint32(v245))) = v250
	v253 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[9]))
	*(*uint8)(unsafe.Add(mBase, uint32(v243+int32(-1)))) = uint8(base.B2i32(v253 == int32(3)))
	goto L54
L54:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+63)))
	if v257 == int32(1) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v285 = m.G0
	v286 = int32(16)
	v287 = v285 - v286
	m.G0 = v287
	F_gettimeofday(m, v287)
	mBase = m.M
	v290 = *(*int64)(unsafe.Add(mBase, uint32(v287)))
	v291 = int64(*(*int32)(unsafe.Add(mBase, uint32(v287)+8)))
	m.G0 = v287 + v286
	goto L61
L56:
	;
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[10]))
	if v262 < int32(0) {
		v281 = int64(0)
		goto L55
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[11]))
	if v272 < int32(0) {
		v281 = int64(0)
		goto L55
	} else {
		goto L60
	}
L59:
	;
	v265 = *(*int64)(unsafe.Add(mBase, uint32(v245)))
	v281 = v265 + base.I64_extend_i32_u(v262)*int64(1000)
	goto L55
L60:
	;
	v275 = *(*int64)(unsafe.Add(mBase, uint32(v245)))
	v281 = v275 + base.I64_extend_i32_u(v272)*int64(1000)
	goto L55
L61:
	;
	v301 = base.B2i32(v281 == int64(0))
	if v301|base.B2i32(v291+v290*int64(1000000)-int64(946684800000000) < v281) == int32(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	F_ProcWaitForSignal(m, int32(67108864))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L72
	}
L63:
	;
	v306 = int32(0)
	F_CancelDBBackends(m, v306, int32(12), v306)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	if v281 == int64(0) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L62
L67:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[12])) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v319))) = int64(4)
	v326 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v319)+8)) = v326
	F_enable_timeouts(m, v245, v318)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L71
	}
L68:
	;
	v318 = int32(1)
	v319 = v245
	goto L67
L69:
	;
	goto L70
L70:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v245)+16)) = v281
	*(*int64)(unsafe.Add(mBase, uint32(v245))) = int64(4294967301)
	v318 = int32(2)
	v319 = v243 + int32(-40)
	goto L67
L71:
	;
	goto L62
L72:
	;
	v337 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[13]))
	if v337 != 0 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v349 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[14])) = v349
	*(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[15])) = v349
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[16])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[17])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[18])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[19])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[20])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[21])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[22])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[23])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[24])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[25])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[26])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[27])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[28])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[29])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[30])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[31])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[32])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[33])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[34])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[35])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[36])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[37])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[38])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[39])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[40])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[41])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[42])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[43])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[44])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[45])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[46])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[47])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[48])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[49])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[50])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[51])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[52])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[53])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[54])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[55])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[56])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[57])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[58])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[59])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[60])) = uint8(v349)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[61])) = uint8(v349)
	goto L79
L74:
	;
	v344 = int32(12)
	goto L76
L75:
	;
	v340 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[12]))
	if v340 == int32(0) {
		goto L73
	} else {
		goto L77
	}
L76:
	;
	F_CancelDBBackends(m, int32(0), v344, int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L78
	}
L77:
	;
	v344 = int32(13)
	goto L76
L78:
	;
	goto L73
L79:
	;
	v493 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[13])) = v493
	*(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[12])) = v493
	m.G0 = v245 - int32(-64)
	v503 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v503)+72)) = int32(-1)
	goto L80
L80:
	;
	v510 = v237
	v512 = v238
	goto L38
L81:
	;
	v510 = v32
	v512 = v35
	goto L38
L82:
	;
	goto L85
L83:
	;
	v553 = v525
	goto L84
L84:
	;
	v566 = int32(_a_F_LockBufferForCleanup_4)
	v567 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[1]))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(8))+8))
	if v569 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L85:
	;
	F_perform_spin_delay(m, v14+int32(8))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L87
	}
L86:
	;
	v553 = v546
	goto L84
L87:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	v547 = int32(_a_F_LockBufferForCleanup_3)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v546 | v547
	if v546&v547 != 0 {
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	if v553&int32(536870912) != 0 {
		goto L100
	} else {
		goto L101
	}
L90:
	;
	goto L89
L91:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[1])) = v584
	goto L90
L92:
	;
	if int32(999) < v567 {
		goto L90
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	if v567 < int32(11) {
		goto L90
	} else {
		goto L99
	}
L95:
	;
	v574 = int32(900)
	if v574 <= v567 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v577 = v574
	goto L98
L97:
	;
	v577 = v567
	goto L98
L98:
	;
	v584 = v577 + int32(100)
	goto L91
L99:
	;
	v584 = v567 - int32(1)
	goto L91
L100:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	v592 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[2]))
	if v590 == v592 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	v595 = v553
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v595 & int32(-4194305)
	*(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[3])) = int32(0)
	v603 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[0]))
	v28 = v603
	v32 = v510
	v35 = v512
	goto L5
L103:
	;
	v594 = v553 & int32(-536870913)
	goto L105
L104:
	;
	v594 = v553
	goto L105
L105:
	;
	v595 = v594
	goto L102
L106:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_errmsg_internal(m, int32(_a_F_LockBufferForCleanup_6), int32(0))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_LockBufferForCleanup_2), int32(_a_F_LockBufferForCleanup_7), int32(_a_F_LockBufferForCleanup_8))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_MarkBufferDirtyHint(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int64
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
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
	var v110 int64
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
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
	var v128 int64
	_ = v128
	var v129 int32
	_ = v129
	var v130 int64
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v201 int64
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int64
	_ = v226
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v246 int64
	_ = v246
	var v247 int32
	_ = v247
	var v251 int64
	_ = v251
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v273 int64
	_ = v273
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v357 int32
	_ = v357
	var v361 int64
	_ = v361
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v381 int64
	_ = v381
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	v3 = int32(0)
	v11 = int64(0)
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	if v3 <= l0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v15 - int32(-64)
	return
L2:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[0]))
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[1]))
	v61 = v58 + l0<<(uint(int32(6))%32)
	v63 = v61 - int32(40)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v65 = int32(276824064)
	if v64&v65 == v65 {
		goto L1
	} else {
		goto L15
	}
L3:
	;
	if l0 != 0 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[2]))
	v40 = v35 + (l0^int32(-1))<<(uint(int32(6))%32)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+24))
	if v41&int32(_a_F_MarkBufferDirtyHint_0) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(0)
	F_errmsg_internal(m, int32(_a_F_MarkBufferDirtyHint_1), v15)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	F_errfinish(m, int32(_a_F_MarkBufferDirtyHint_2), int32(_a_F_MarkBufferDirtyHint_3), int32(_a_F_MarkBufferDirtyHint_4))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L11:
	;
	goto L1
L12:
	;
	v46 = int32(_a_F_MarkBufferDirtyHint_5)
	v48 = *(*int64)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[3]))
	*(*int64)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[3])) = v48 + int64(1)
	goto L14
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+24)) = v41 | int32(_a_F_MarkBufferDirtyHint_0)
	goto L11
L15:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[4]))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+252))
	goto L17
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = int32(_a_F_MarkBufferDirtyHint_6)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = int32(_a_F_MarkBufferDirtyHint_7)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = int32(_a_F_MarkBufferDirtyHint_2)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+40)) = int64(0)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v286 = int32(_a_F_MarkBufferDirtyHint_8)
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v285 | v286
	if v285&v286 != 0 {
		goto L66
	} else {
		goto L67
	}
L17:
	;
	if base.B2i32(v71 != int32(0)) == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[5])))
	if v77&int32(1) == int32(0) {
		v272 = v3
		v273 = v11
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if int32(0) <= v82 {
		v272 = v3
		v273 = v11
		goto L16
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[6])))
	if v87 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v97 != 0 {
		goto L1
	} else {
		goto L27
	}
L24:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[7]))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+316))
	v95 = base.B2i32(v93 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[6])) = uint8(v95)
	v97 = v95
	goto L26
L25:
	;
	v97 = int32(0)
	goto L26
L26:
	;
	goto L23
L27:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-64))))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v100
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v61-int32(60))))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v104
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v61-int32(56))))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v108
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v15)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v108
	v115 = F_RelFileLocatorSkippingWAL(m, v13+int32(-48))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	if v115 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v117 = int32(1)
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[8]))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v119)+120)) = v120 | v117
	v124 = m.G0
	v126 = v124 - int32(_a_F_MarkBufferDirtyHint_9)
	m.G0 = v126
	v128 = F_GetRedoRecPtr(m)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L7
	} else {
		goto L31
	}
L30:
	;
	v272 = v117
	v273 = v251
	goto L16
L31:
	;
	v130 = F_BufferGetLSNAtomic(m, l0)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L7
	} else {
		goto L33
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L7
	} else {
		goto L63
	}
L33:
	;
	if base.Ui64(v130) <= base.Ui64(v128) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	if l0 < int32(0) {
		goto L40
	} else {
		goto L41
	}
L35:
	;
	v251 = v11
	goto L36
L36:
	;
	m.G0 = v126 + int32(_a_F_MarkBufferDirtyHint_9)
	goto L30
L37:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L7
	} else {
		goto L49
	}
L38:
	;
	base.MemoryCopy(m, v126+int32(32), v167, int32(_a_F_MarkBufferDirtyHint_10))
	goto L37
L39:
	;
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152)+14)))
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152)+12)))
	if v154 != 0 {
		goto L45
	} else {
		goto L46
	}
L40:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[9]))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v136+(l0^int32(-1))<<(uint(int32(2))%32))))
	if l1 != 0 {
		v152 = v142
		goto L39
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[0]))
	v149 = v144 + l0<<(uint(int32(13))%32) + int32(-8192)
	if l1 == int32(0) {
		v167 = v149
		goto L38
	} else {
		goto L44
	}
L43:
	;
	v167 = v142
	goto L38
L44:
	;
	v152 = v149
	goto L39
L45:
	;
	base.MemoryCopy(m, v126+int32(32), v152, v154)
	goto L47
L46:
	;
	goto L47
L47:
	;
	v159 = int32(_a_F_MarkBufferDirtyHint_10) - v153
	if v159 == int32(0) {
		goto L37
	} else {
		goto L48
	}
L48:
	;
	base.MemoryCopy(m, v126+int32(32)+v153, v152+v153, v159)
	goto L37
L49:
	;
	v178 = v126 + int32(20)
	if l0 < int32(0) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[10]))
	if v210 <= int32(0) {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	v201 = *(*int64)(unsafe.Add(mBase, uint32(v200)))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v178)+8)) = v202
	*(*int64)(unsafe.Add(mBase, uint32(v178))) = v201
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v200)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v126+int32(16)))) = v205
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v200)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v126+int32(12)))) = v207
	goto L50
L52:
	;
	v187 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[2]))
	v200 = v187 + (l0^int32(-1))<<(uint(int32(6))%32)
	goto L51
L53:
	;
	goto L54
L54:
	;
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[1]))
	v200 = v194 + l0<<(uint(int32(6))%32) + int32(-64)
	goto L51
L55:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[10])) = int32(1)
	goto L57
L56:
	;
	goto L57
L57:
	;
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[11]))
	if v217 <= int32(0) {
		goto L32
	} else {
		goto L58
	}
L58:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[12]))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v126)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v223)+12)) = v224
	v226 = *(*int64)(unsafe.Add(mBase, uint32(v126)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v223)+4)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v223)+20)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v223)+16)) = v221
	if l1 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v232 = int32(8)
	goto L61
L60:
	;
	v232 = int32(0)
	goto L61
L61:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v223)+1)) = uint8(v232)
	v234 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+28)) = v234
	v236 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v223))) = uint8(v236)
	v238 = int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+36)) = v223 + v238
	*(*int32)(unsafe.Add(mBase, uint32(v223)+24)) = v126 + v238
	v246 = F_XLogInsert(m, v234, int32(160))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L7
	} else {
		goto L62
	}
L62:
	;
	v251 = v246
	goto L36
L63:
	;
	F_errmsg_internal(m, int32(_a_F_MarkBufferDirtyHint_11), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L7
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_MarkBufferDirtyHint_12), int32(320), int32(_a_F_MarkBufferDirtyHint_13))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L7
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	goto L69
L67:
	;
	v314 = v285
	goto L68
L68:
	;
	v328 = int32(_a_F_MarkBufferDirtyHint_14)
	v329 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[13]))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-24))+8))
	if v331 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L69:
	;
	F_perform_spin_delay(m, v13+int32(-24))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L7
	} else {
		goto L71
	}
L70:
	;
	v314 = v307
	goto L68
L71:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v308 = int32(_a_F_MarkBufferDirtyHint_8)
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v307 | v308
	if v307&v308 != 0 {
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v349 = v314 & int32(_a_F_MarkBufferDirtyHint_0)
	if v349|base.B2i32(v273 == int64(0)) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L74:
	;
	goto L73
L75:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[13])) = v346
	goto L74
L76:
	;
	if int32(999) < v329 {
		goto L74
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	if v329 < int32(11) {
		goto L74
	} else {
		goto L83
	}
L79:
	;
	v336 = int32(900)
	if v336 <= v329 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v339 = v336
	goto L82
L81:
	;
	v339 = v329
	goto L82
L82:
	;
	v346 = v339 + int32(100)
	goto L75
L83:
	;
	v346 = v329 - int32(1)
	goto L75
L84:
	;
	v357 = v56 + l0<<(uint(int32(13))%32)
	v361 = int64(base.Ui64(v273) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v357)+uint32(_c_F_MarkBufferDirtyHint[14]))) = uint32(v361)
	*(*uint32)(unsafe.Add(mBase, uint32(v357-int32(_a_F_MarkBufferDirtyHint_15)))) = uint32(v273)
	goto L86
L85:
	;
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v314&int32(-281018369) | int32(276824064)
	if v272 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v373 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[8]))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v373)+120)) = v374 & int32(-2)
	goto L89
L88:
	;
	goto L89
L89:
	;
	if v349 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v379 = int32(_a_F_MarkBufferDirtyHint_16)
	v381 = *(*int64)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[15]))
	*(*int64)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[15])) = v381 + int64(1)
	v386 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[16])))
	if v386 != int32(1) {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v389 = int32(_a_F_MarkBufferDirtyHint_17)
	v391 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[17]))
	v393 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[18]))
	*(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[17])) = v391 + v393
	goto L1
}
func F_ReadBuffer(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v3 = int32(0)
	v6 = F_ReadBufferExtended(m, l0, v3, l1, v3, v3)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_StartBufferIO(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_StartBufferIO[0]))
	F_ResourceOwnerEnlarge(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L6
L3:
	;
	m.G0 = v8 + int32(32)
	return v109
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v55&int32(-71303169) | int32(67108864)
	v100 = int32(1)
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_StartBufferIO[0]))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_ResourceOwnerRemember(m, v102, v103+v100, int32(_a_F_StartBufferIO_0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L37
	}
L5:
	;
	v109 = int32(0)
	goto L3
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = int32(_a_F_StartBufferIO_1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(_a_F_StartBufferIO_2)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_StartBufferIO_3)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(0)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v32 = int32(_a_F_StartBufferIO_4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v31 | v32
	if v31&v32 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if l1 != 0 {
		goto L32
	} else {
		goto L33
	}
L8:
	;
	goto L11
L9:
	;
	v55 = v31
	goto L10
L10:
	;
	v60 = int32(_a_F_StartBufferIO_5)
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_StartBufferIO[1]))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(8))+8))
	if v63 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L11:
	;
	F_perform_spin_delay(m, v8+int32(8))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	v55 = v46
	goto L10
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v47 = int32(_a_F_StartBufferIO_4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v46 | v47
	if v46&v47 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	if v55&int32(67108864) != 0 {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartBufferIO[1])) = v78
	goto L16
L18:
	;
	if int32(999) < v61 {
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v61 < int32(11) {
		goto L16
	} else {
		goto L25
	}
L21:
	;
	v68 = int32(900)
	if v68 <= v61 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v71 = v68
	goto L24
L23:
	;
	v71 = v61
	goto L24
L24:
	;
	v78 = v71 + int32(100)
	goto L17
L25:
	;
	v78 = v61 - int32(1)
	goto L17
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v55 & int32(-4194305)
	if l2 != 0 {
		goto L5
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	goto L7
L29:
	;
	F_WaitIO(m, l0)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	goto L6
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v55 & int32(-71303169)
	goto L5
L32:
	;
	if v55&int32(16777216) != 0 {
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if v55&int32(_a_F_StartBufferIO_6) != 0 {
		goto L4
	} else {
		goto L36
	}
L35:
	;
	goto L4
L36:
	;
	goto L31
L37:
	;
	v109 = v100
	goto L3
}
