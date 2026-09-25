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
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v132 int64
	_ = v132
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
	return v132
L7:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_BufferGetLSNAtomic[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = int32(_a_F_BufferGetLSNAtomic_0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(_a_F_BufferGetLSNAtomic_1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(_a_F_BufferGetLSNAtomic_2)
	v54 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v54
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(0)
	v62 = v47 + l0<<(uint(int32(6))%32) - int32(40)
	v63 = int32(_a_F_BufferGetLSNAtomic_3)
	v65 = base.AtomicRmwOr32(m, v62, v54, v63)
	if v65&v63 != 0 {
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
	v132 = v41 | v42<<(uint(int64(32))%64)
	goto L6
L11:
	;
	goto L10
L12:
	;
	goto L15
L13:
	;
	v87 = v65
	goto L14
L14:
	;
	v94 = int32(_a_F_BufferGetLSNAtomic_4)
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_BufferGetLSNAtomic[5]))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(8))+8))
	if v97 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L15:
	;
	F_perform_spin_delay(m, v9+int32(8))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v87 = v82
	goto L14
L17:
	;
	return int64(0)
L18:
	;
	v80 = int32(_a_F_BufferGetLSNAtomic_3)
	v82 = base.AtomicRmwOr32(m, v62, int32(0), v80)
	if v82&v80 != 0 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v114 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v28)+4)))
	v115 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v28))))
	v116 = int32(0)
	v119 = base.AtomicRmwOr32(m, v116, int32(_a_F_BufferGetLSNAtomic_5), v116)
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v87 & int32(-4194305)
	v132 = v114 | v115<<(uint(int64(32))%64)
	goto L6
L21:
	;
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BufferGetLSNAtomic[5])) = v112
	goto L21
L23:
	;
	if int32(999) < v95 {
		goto L21
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if v95 < int32(11) {
		goto L21
	} else {
		goto L30
	}
L26:
	;
	v102 = int32(900)
	if v102 <= v95 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v105 = v102
	goto L29
L28:
	;
	v105 = v95
	goto L29
L29:
	;
	v112 = v105 + int32(100)
	goto L22
L30:
	;
	v112 = v95 - int32(1)
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
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v150 int32
	_ = v150
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
	return v150
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_IsBufferCleanupOK[0]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v12+(l0^int32(-1))<<(uint(int32(2))%32))))
	v150 = base.B2i32(v18 == int32(1))
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
		v150 = v2
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
		v150 = v2
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
		v150 = v2
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
	v82 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = int64(0)
	v90 = v75 + l0<<(uint(int32(6))%32) - int32(40)
	v91 = int32(_a_F_IsBufferCleanupOK_11)
	v93 = base.AtomicRmwOr32(m, v90, v82, v91)
	if v93&v91 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	goto L38
L36:
	;
	v109 = v93
	goto L37
L37:
	;
	v116 = int32(_a_F_IsBufferCleanupOK_12)
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_IsBufferCleanupOK[12]))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(8))+8))
	if v119 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L38:
	;
	F_perform_spin_delay(m, v7+int32(8))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L31
	} else {
		goto L40
	}
L39:
	;
	v109 = v106
	goto L37
L40:
	;
	v104 = int32(_a_F_IsBufferCleanupOK_11)
	v106 = base.AtomicRmwOr32(m, v90, int32(0), v104)
	if v106&v104 != 0 {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v136 = int32(0)
	v139 = base.AtomicRmwOr32(m, v136, int32(_a_F_IsBufferCleanupOK_13), v136)
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = v109 & int32(-4194305)
	v150 = base.B2i32(v109&int32(_a_F_IsBufferCleanupOK_14) == int32(1))
	goto L1
L43:
	;
	goto L42
L44:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IsBufferCleanupOK[12])) = v134
	goto L43
L45:
	;
	if int32(999) < v117 {
		goto L43
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	if v117 < int32(11) {
		goto L43
	} else {
		goto L52
	}
L48:
	;
	v124 = int32(900)
	if v124 <= v117 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v127 = v124
	goto L51
L50:
	;
	v127 = v117
	goto L51
L51:
	;
	v134 = v127 + int32(100)
	goto L44
L52:
	;
	v134 = v117 - int32(1)
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
	var v58 int32
	_ = v58
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int64
	_ = v135
	var v136 int64
	_ = v136
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int64
	_ = v192
	var v193 int64
	_ = v193
	var v201 int64
	_ = v201
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int64
	_ = v233
	var v234 int64
	_ = v234
	var v243 int32
	_ = v243
	var v244 int64
	_ = v244
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v256 int64
	_ = v256
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v271 int64
	_ = v271
	var v278 int32
	_ = v278
	var v281 int64
	_ = v281
	var v287 int64
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int64
	_ = v296
	var v297 int64
	_ = v297
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v499 int32
	_ = v499
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v518 int64
	_ = v518
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v588 int32
	_ = v588
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
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
	v612 = int32(0)
	v615 = base.AtomicRmwOr32(m, v612, int32(_a_F_LockBufferForCleanup_0), v612)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v82 & int32(-4194305)
	F_LockBuffer(m, l0, v612)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L106
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = int32(_a_F_LockBufferForCleanup_1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = int32(_a_F_LockBufferForCleanup_2)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = int32(_a_F_LockBufferForCleanup_3)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = int64(0)
	v56 = int32(_a_F_LockBufferForCleanup_4)
	v58 = base.AtomicRmwOr32(m, v26, int32(24), v56)
	if v58&v56 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	goto L11
L9:
	;
	v82 = v58
	goto L10
L10:
	;
	v95 = int32(_a_F_LockBufferForCleanup_5)
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[1]))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(8))+8))
	if v98 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L11:
	;
	F_perform_spin_delay(m, v14+int32(8))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	v82 = v78
	goto L10
L13:
	;
	v76 = int32(_a_F_LockBufferForCleanup_4)
	v78 = base.AtomicRmwOr32(m, v26, int32(24), v76)
	if v78&v76 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	if v82&int32(_a_F_LockBufferForCleanup_6) == int32(1) {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[1])) = v113
	goto L16
L18:
	;
	if int32(999) < v96 {
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v96 < int32(11) {
		goto L16
	} else {
		goto L25
	}
L21:
	;
	v103 = int32(900)
	if v103 <= v96 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v106 = v103
	goto L24
L23:
	;
	v106 = v96
	goto L24
L24:
	;
	v113 = v106 + int32(100)
	goto L17
L25:
	;
	v113 = v96 - int32(1)
	goto L17
L26:
	;
	v119 = int32(0)
	v122 = base.AtomicRmwOr32(m, v119, int32(_a_F_LockBufferForCleanup_0), v119)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v82 & int32(-4456447)
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
	if v82&int32(536870912) == int32(0) {
		goto L34
	} else {
		goto L35
	}
L29:
	;
	v130 = m.G0
	v131 = int32(16)
	v132 = v130 - v131
	m.G0 = v132
	F_gettimeofday(m, v132)
	mBase = m.M
	v135 = *(*int64)(unsafe.Add(mBase, uint32(v132)))
	v136 = int64(*(*int32)(unsafe.Add(mBase, uint32(v132)+8)))
	m.G0 = v132 + v131
	goto L32
L30:
	;
	goto L31
L31:
	;
	goto L3
L32:
	;
	v145 = int32(0)
	F_LogRecoveryConflict(m, int32(12), v35, v136+v135*int64(1000000)-int64(946684800000000), v145, v145)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v154
	*(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[3])) = v26
	v158 = int32(0)
	v161 = base.AtomicRmwOr32(m, v158, int32(_a_F_LockBufferForCleanup_0), v158)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v82&int32(-541065217) | int32(536870912)
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[0]))
	F_LWLockRelease(m, v168+v39-int32(16))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
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
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[4]))
	if base.Ui32(int32(2)) <= base.Ui32(v175) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = int32(_a_F_LockBufferForCleanup_1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = int32(_a_F_LockBufferForCleanup_2)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = int32(_a_F_LockBufferForCleanup_3)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = int64(0)
	v531 = int32(_a_F_LockBufferForCleanup_4)
	v533 = base.AtomicRmwOr32(m, v26, int32(24), v531)
	if v533&v531 != 0 {
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
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L81
	}
L42:
	;
	v247 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v247)+72)) = v23
	goto L53
L43:
	;
	v187 = m.G0
	v188 = int32(16)
	v189 = v187 - v188
	m.G0 = v189
	F_gettimeofday(m, v189)
	mBase = m.M
	v192 = *(*int64)(unsafe.Add(mBase, uint32(v189)))
	v193 = int64(*(*int32)(unsafe.Add(mBase, uint32(v189)+8)))
	m.G0 = v189 + v188
	v201 = v193 + v192*int64(1000000) - int64(946684800000000)
	goto L46
L44:
	;
	goto L45
L45:
	;
	if v35 != int64(0) {
		v243 = v32
		v244 = v35
		goto L42
	} else {
		goto L50
	}
L46:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[6]))
	goto L47
L47:
	;
	if base.B2i32(base.I64_extend_i32_s(v203)*int64(1000) <= v201-v35) == int32(0) {
		v243 = int32(0)
		v244 = v35
		goto L42
	} else {
		goto L48
	}
L48:
	;
	v211 = int32(1)
	F_LogRecoveryConflict(m, int32(12), v35, v201, int32(0), v211)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v243 = v211
	v244 = v35
	goto L42
L50:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[7])))
	if v220&int32(1) == int32(0) {
		v243 = v32
		v244 = v35
		goto L42
	} else {
		goto L51
	}
L51:
	;
	v228 = m.G0
	v229 = int32(16)
	v230 = v228 - v229
	m.G0 = v230
	F_gettimeofday(m, v230)
	mBase = m.M
	v233 = *(*int64)(unsafe.Add(mBase, uint32(v230)))
	v234 = int64(*(*int32)(unsafe.Add(mBase, uint32(v230)+8)))
	m.G0 = v230 + v229
	goto L52
L52:
	;
	v243 = v32
	v244 = v234 + v233*int64(1000000) - int64(946684800000000)
	goto L42
L53:
	;
	v249 = m.G0
	v251 = v249 + int32(-64)
	m.G0 = v251
	v256 = *(*int64)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[8]))
	*(*int64)(unsafe.Add(mBase, uint32(v251))) = v256
	v259 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[9]))
	*(*uint8)(unsafe.Add(mBase, uint32(v249+int32(-1)))) = uint8(base.B2i32(v259 == int32(3)))
	goto L54
L54:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+63)))
	if v263 == int32(1) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v291 = m.G0
	v292 = int32(16)
	v293 = v291 - v292
	m.G0 = v293
	F_gettimeofday(m, v293)
	mBase = m.M
	v296 = *(*int64)(unsafe.Add(mBase, uint32(v293)))
	v297 = int64(*(*int32)(unsafe.Add(mBase, uint32(v293)+8)))
	m.G0 = v293 + v292
	goto L61
L56:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[10]))
	if v268 < int32(0) {
		v287 = int64(0)
		goto L55
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v278 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[11]))
	if v278 < int32(0) {
		v287 = int64(0)
		goto L55
	} else {
		goto L60
	}
L59:
	;
	v271 = *(*int64)(unsafe.Add(mBase, uint32(v251)))
	v287 = v271 + base.I64_extend_i32_u(v268)*int64(1000)
	goto L55
L60:
	;
	v281 = *(*int64)(unsafe.Add(mBase, uint32(v251)))
	v287 = v281 + base.I64_extend_i32_u(v278)*int64(1000)
	goto L55
L61:
	;
	v307 = base.B2i32(v287 == int64(0))
	if v307|base.B2i32(v297+v296*int64(1000000)-int64(946684800000000) < v287) == int32(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	F_ProcWaitForSignal(m, int32(67108864))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L72
	}
L63:
	;
	v312 = int32(0)
	F_CancelDBBackends(m, v312, int32(12), v312)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	if v287 == int64(0) {
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
	*(*int64)(unsafe.Add(mBase, uint32(v325))) = int64(4)
	v332 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v325)+8)) = v332
	F_enable_timeouts(m, v251, v324)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L71
	}
L68:
	;
	v324 = int32(1)
	v325 = v251
	goto L67
L69:
	;
	goto L70
L70:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v251)+16)) = v287
	*(*int64)(unsafe.Add(mBase, uint32(v251))) = int64(4294967301)
	v324 = int32(2)
	v325 = v249 + int32(-40)
	goto L67
L71:
	;
	goto L62
L72:
	;
	v343 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[13]))
	if v343 != 0 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v355 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[14])) = v355
	*(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[15])) = v355
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[16])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[17])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[18])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[19])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[20])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[21])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[22])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[23])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[24])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[25])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[26])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[27])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[28])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[29])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[30])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[31])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[32])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[33])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[34])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[35])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[36])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[37])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[38])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[39])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[40])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[41])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[42])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[43])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[44])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[45])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[46])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[47])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[48])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[49])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[50])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[51])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[52])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[53])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[54])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[55])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[56])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[57])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[58])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[59])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[60])) = uint8(v355)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[61])) = uint8(v355)
	goto L79
L74:
	;
	v350 = int32(12)
	goto L76
L75:
	;
	v346 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[12]))
	if v346 == int32(0) {
		goto L73
	} else {
		goto L77
	}
L76:
	;
	F_CancelDBBackends(m, int32(0), v350, int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L78
	}
L77:
	;
	v350 = int32(13)
	goto L76
L78:
	;
	goto L73
L79:
	;
	v499 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[13])) = v499
	*(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[12])) = v499
	m.G0 = v251 - int32(-64)
	v509 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v509)+72)) = int32(-1)
	goto L80
L80:
	;
	v516 = v243
	v518 = v244
	goto L38
L81:
	;
	v516 = v32
	v518 = v35
	goto L38
L82:
	;
	goto L85
L83:
	;
	v557 = v533
	goto L84
L84:
	;
	v570 = int32(_a_F_LockBufferForCleanup_5)
	v571 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[1]))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(8))+8))
	if v573 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L85:
	;
	F_perform_spin_delay(m, v14+int32(8))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L87
	}
L86:
	;
	v557 = v553
	goto L84
L87:
	;
	v551 = int32(_a_F_LockBufferForCleanup_4)
	v553 = base.AtomicRmwOr32(m, v26, int32(24), v551)
	if v553&v551 != 0 {
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	if v557&int32(536870912) != 0 {
		goto L100
	} else {
		goto L101
	}
L90:
	;
	goto L89
L91:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[1])) = v588
	goto L90
L92:
	;
	if int32(999) < v571 {
		goto L90
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	if v571 < int32(11) {
		goto L90
	} else {
		goto L99
	}
L95:
	;
	v578 = int32(900)
	if v578 <= v571 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v581 = v578
	goto L98
L97:
	;
	v581 = v571
	goto L98
L98:
	;
	v588 = v581 + int32(100)
	goto L91
L99:
	;
	v588 = v571 - int32(1)
	goto L91
L100:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	v596 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[2]))
	if v594 == v596 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	v599 = v557
	goto L102
L102:
	;
	v600 = int32(0)
	v603 = base.AtomicRmwOr32(m, v600, int32(_a_F_LockBufferForCleanup_0), v600)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v599 & int32(-4194305)
	*(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[3])) = v600
	v611 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[0]))
	v28 = v611
	v32 = v516
	v35 = v518
	goto L5
L103:
	;
	v598 = v557 & int32(-536870913)
	goto L105
L104:
	;
	v598 = v557
	goto L105
L105:
	;
	v599 = v598
	goto L102
L106:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_errmsg_internal(m, int32(_a_F_LockBufferForCleanup_7), int32(0))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_LockBufferForCleanup_3), int32(_a_F_LockBufferForCleanup_8), int32(_a_F_LockBufferForCleanup_9))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
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
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v355 int32
	_ = v355
	var v359 int64
	_ = v359
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v383 int64
	_ = v383
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
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
	v281 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v281
	*(*int64)(unsafe.Add(mBase, uint32(v15)+40)) = int64(0)
	v285 = int32(_a_F_MarkBufferDirtyHint_8)
	v287 = base.AtomicRmwOr32(m, v63, v281, v285)
	if v287&v285 != 0 {
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
	v312 = v287
	goto L68
L68:
	;
	v326 = int32(_a_F_MarkBufferDirtyHint_14)
	v327 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[13]))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-24))+8))
	if v329 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L69:
	;
	F_perform_spin_delay(m, v13+int32(-24))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L7
	} else {
		goto L71
	}
L70:
	;
	v312 = v308
	goto L68
L71:
	;
	v306 = int32(_a_F_MarkBufferDirtyHint_8)
	v308 = base.AtomicRmwOr32(m, v63, int32(0), v306)
	if v308&v306 != 0 {
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v347 = v312 & int32(_a_F_MarkBufferDirtyHint_0)
	if v347|base.B2i32(v273 == int64(0)) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L74:
	;
	goto L73
L75:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[13])) = v344
	goto L74
L76:
	;
	if int32(999) < v327 {
		goto L74
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	if v327 < int32(11) {
		goto L74
	} else {
		goto L83
	}
L79:
	;
	v334 = int32(900)
	if v334 <= v327 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v337 = v334
	goto L82
L81:
	;
	v337 = v327
	goto L82
L82:
	;
	v344 = v337 + int32(100)
	goto L75
L83:
	;
	v344 = v327 - int32(1)
	goto L75
L84:
	;
	v355 = v56 + l0<<(uint(int32(13))%32)
	v359 = int64(base.Ui64(v273) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v355)+uint32(_c_F_MarkBufferDirtyHint[14]))) = uint32(v359)
	*(*uint32)(unsafe.Add(mBase, uint32(v355-int32(_a_F_MarkBufferDirtyHint_15)))) = uint32(v273)
	goto L86
L85:
	;
	goto L86
L86:
	;
	v365 = int32(0)
	v368 = base.AtomicRmwOr32(m, v365, int32(_a_F_MarkBufferDirtyHint_16), v365)
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v312&int32(-281018369) | int32(276824064)
	if v272 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v375 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[8]))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v375)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v375)+120)) = v376 & int32(-2)
	goto L89
L88:
	;
	goto L89
L89:
	;
	if v347 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v381 = int32(_a_F_MarkBufferDirtyHint_17)
	v383 = *(*int64)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[15]))
	*(*int64)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[15])) = v383 + int64(1)
	v388 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[16])))
	if v388 != int32(1) {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v391 = int32(_a_F_MarkBufferDirtyHint_18)
	v393 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[17]))
	v395 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[18]))
	*(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[17])) = v393 + v395
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
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_StartBufferIO[0]))
	F_ResourceOwnerEnlarge(m, v12)
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
	goto L4
L3:
	;
	m.G0 = v9 + int32(32)
	return v122
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = int32(_a_F_StartBufferIO_0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(_a_F_StartBufferIO_1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(_a_F_StartBufferIO_2)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(0)
	v33 = int32(_a_F_StartBufferIO_3)
	v35 = base.AtomicRmwOr32(m, l0, int32(24), v33)
	if v35&v33 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if l1 != 0 {
		goto L31
	} else {
		goto L32
	}
L6:
	;
	goto L9
L7:
	;
	v57 = v35
	goto L8
L8:
	;
	v62 = int32(_a_F_StartBufferIO_4)
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_StartBufferIO[1]))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(8))+8))
	if v65 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	F_perform_spin_delay(m, v9+int32(8))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v57 = v50
	goto L8
L11:
	;
	v48 = int32(_a_F_StartBufferIO_3)
	v50 = base.AtomicRmwOr32(m, l0, int32(24), v48)
	if v50&v48 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	if v57&int32(67108864) != 0 {
		goto L24
	} else {
		goto L25
	}
L14:
	;
	goto L13
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartBufferIO[1])) = v80
	goto L14
L16:
	;
	if int32(999) < v63 {
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if v63 < int32(11) {
		goto L14
	} else {
		goto L23
	}
L19:
	;
	v70 = int32(900)
	if v70 <= v63 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v73 = v70
	goto L22
L21:
	;
	v73 = v63
	goto L22
L22:
	;
	v80 = v73 + int32(100)
	goto L15
L23:
	;
	v80 = v63 - int32(1)
	goto L15
L24:
	;
	v84 = int32(0)
	v87 = base.AtomicRmwOr32(m, v84, int32(_a_F_StartBufferIO_5), v84)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v57 & int32(-4194305)
	if l2 != 0 {
		v122 = v4
		goto L3
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	goto L5
L27:
	;
	F_WaitIO(m, l0)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	goto L4
L29:
	;
	v104 = int32(0)
	v107 = base.AtomicRmwOr32(m, v104, int32(_a_F_StartBufferIO_5), v104)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v57&int32(-71303169) | int32(67108864)
	v113 = int32(1)
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_StartBufferIO[0]))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_ResourceOwnerRemember(m, v115, v116+v113, int32(_a_F_StartBufferIO_6))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L36
	}
L30:
	;
	v97 = int32(0)
	v100 = base.AtomicRmwOr32(m, v97, int32(_a_F_StartBufferIO_5), v97)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v57 & int32(-71303169)
	v122 = v4
	goto L3
L31:
	;
	if v57&int32(16777216) != 0 {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if v57&int32(_a_F_StartBufferIO_7) != 0 {
		goto L29
	} else {
		goto L35
	}
L34:
	;
	goto L29
L35:
	;
	goto L30
L36:
	;
	v122 = v113
	goto L3
}
