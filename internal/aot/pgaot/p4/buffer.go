package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BufferGetLSNAtomic(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v51 int32
	_ = v51
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v124 int64
	_ = v124
	var v125 int64
	_ = v125
	var v139 int64
	_ = v139
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	if l0 < int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v10 + int32(32)
	return v139
L2:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(241628)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(520922)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = int64(0)
	v66 = v51 + l0<<(uint(int32(6))%32) - int32(40)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v68 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v67 | v68
	if v67&v68 != 0 {
		goto L11
	} else {
		goto L12
	}
L3:
	;
	v45 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v43)+4)))
	v46 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v43))))
	v139 = v45 | v46<<(uint(int64(32))%64)
	goto L1
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v15+(l0^int32(-1))<<(uint(int32(2))%32))))
	goto L7
L5:
	;
	goto L6
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v31 = v28 + l0<<(uint(int32(13))%32)
	v35 = *(*int32)(unsafe.Add(mBase, _consts[284]))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+252))
	goto L8
L7:
	;
	v43 = v21
	goto L3
L8:
	;
	if v36 != int32(0) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, _consts[335])))
	if v40&int32(1) != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v43 = v31 + int32(-8192)
	goto L3
L11:
	;
	goto L14
L12:
	;
	v94 = v67
	goto L13
L13:
	;
	v102 = int32(4164732)
	v103 = *(*int32)(unsafe.Add(mBase, _consts[423]))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(8))+8))
	if v105 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L14:
	;
	F_perform_spin_delay(m, v10+int32(8))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v94 = v86
	goto L13
L16:
	;
	return int64(0)
L17:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v87 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v86 | v87
	if v86&v87 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	v124 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v31-int32(8188)))))
	v125 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v31)+uint32(_consts[919]))))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v94 & int32(-4194305)
	v139 = v124 | v125<<(uint(int64(32))%64)
	goto L1
L20:
	;
	goto L19
L21:
	;
	*(*int32)(unsafe.Add(mBase, _consts[423])) = v120
	goto L20
L22:
	;
	if int32(999) < v103 {
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v103 < int32(11) {
		goto L20
	} else {
		goto L29
	}
L25:
	;
	v110 = int32(900)
	if v110 <= v103 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v113 = v110
	goto L28
L27:
	;
	v113 = v103
	goto L28
L28:
	;
	v120 = v113 + int32(100)
	goto L21
L29:
	;
	v120 = v103 - int32(1)
	goto L21
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
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v153 int32
	_ = v153
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
	return v153
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _consts[938]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v12+(l0^int32(-1))<<(uint(int32(2))%32))))
	v153 = base.B2i32(v18 == int32(1))
	goto L1
L3:
	;
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l0
	v23 = *(*int32)(unsafe.Add(mBase, _consts[928]))
	if l0 == v23 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if v71 != int32(1) {
		v153 = v2
		goto L1
	} else {
		goto L34
	}
L6:
	;
	v70 = int32(4477904)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _consts[929]))
	if l0 == v27 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v70 = int32(4477912)
	goto L5
L10:
	;
	goto L11
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _consts[930]))
	if l0 == v31 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v70 = int32(4477920)
	goto L5
L13:
	;
	goto L14
L14:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _consts[931]))
	if l0 == v35 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v70 = int32(4477928)
	goto L5
L16:
	;
	goto L17
L17:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[932]))
	if l0 == v39 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v70 = int32(4477936)
	goto L5
L19:
	;
	goto L20
L20:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _consts[933]))
	if l0 == v43 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v70 = int32(4477944)
	goto L5
L22:
	;
	goto L23
L23:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[934]))
	if l0 == v47 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v70 = int32(4477952)
	goto L5
L25:
	;
	goto L26
L26:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _consts[935]))
	if l0 == v51 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v70 = int32(4477960)
	goto L5
L28:
	;
	goto L29
L29:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _consts[936]))
	if v55 == int32(0) {
		v153 = v2
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _consts[937]))
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
		v153 = v2
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
	v75 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(241628)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = int32(520922)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = int64(0)
	v90 = v75 + l0<<(uint(int32(6))%32) - int32(40)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v92 = int32(4194304)
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
	v119 = int32(4164732)
	v120 = *(*int32)(unsafe.Add(mBase, _consts[423]))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(8))+8))
	if v122 == int32(0) {
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
	v106 = int32(4194304)
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
	if v111&int32(262143) == int32(1) {
		goto L53
	} else {
		goto L54
	}
L43:
	;
	goto L42
L44:
	;
	*(*int32)(unsafe.Add(mBase, _consts[423])) = v137
	goto L43
L45:
	;
	if int32(999) < v120 {
		goto L43
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	if v120 < int32(11) {
		goto L43
	} else {
		goto L52
	}
L48:
	;
	v127 = int32(900)
	if v127 <= v120 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v130 = v127
	goto L51
L50:
	;
	v130 = v120
	goto L51
L51:
	;
	v137 = v130 + int32(100)
	goto L44
L52:
	;
	v137 = v120 - int32(1)
	goto L44
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = v111 & int32(-4456447)
	v153 = int32(1)
	goto L1
L54:
	;
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = v111 & int32(-4194305)
	v153 = int32(0)
	goto L1
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
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int64
	_ = v190
	var v191 int64
	_ = v191
	var v199 int64
	_ = v199
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int64
	_ = v232
	var v233 int64
	_ = v233
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
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v496 int32
	_ = v496
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v515 int64
	_ = v515
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v587 int32
	_ = v587
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v606 int32
	_ = v606
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
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
	v21 = *(*int32)(unsafe.Add(mBase, _consts[4]))
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
	v612 = m.ExcPending
	if v612 != 0 {
		goto L1
	} else {
		goto L106
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = int32(241628)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = int32(520922)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = int64(0)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	v57 = int32(4194304)
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
	v97 = int32(4164732)
	v98 = *(*int32)(unsafe.Add(mBase, _consts[423]))
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
	v78 = int32(4194304)
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
	if v84&int32(262143) == int32(1) {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, _consts[423])) = v115
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
	if v32&int32(1) != 0 {
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
	v130 = m.G0
	v131 = int32(16)
	v132 = v130 - v131
	m.G0 = v132
	F___gettimeofday(m, v132)
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
	v154 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v154
	*(*int32)(unsafe.Add(mBase, _consts[921])) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v84&int32(-541065217) | int32(536870912)
	v164 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	F_LWLockRelease(m, v164+v39-int32(16))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
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
	v171 = *(*int32)(unsafe.Add(mBase, _consts[266]))
	if base.Ui32(int32(2)) <= base.Ui32(v171) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = int32(241628)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = int32(520922)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = int64(0)
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	v529 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v528 | v529
	if v528&v529 != 0 {
		goto L82
	} else {
		goto L83
	}
L39:
	;
	if (base.B2i32(v35 == int64(0))|v32)&int32(1) == int32(0) {
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
	v510 = m.ExcPending
	if v510 != 0 {
		goto L1
	} else {
		goto L81
	}
L42:
	;
	v247 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	*(*int32)(unsafe.Add(mBase, uint32(v247)+72)) = v23
	goto L53
L43:
	;
	v185 = m.G0
	v186 = int32(16)
	v187 = v185 - v186
	m.G0 = v187
	F___gettimeofday(m, v187)
	mBase = m.M
	v190 = *(*int64)(unsafe.Add(mBase, uint32(v187)))
	v191 = int64(*(*int32)(unsafe.Add(mBase, uint32(v187)+8)))
	m.G0 = v187 + v186
	v199 = v191 + v190*int64(1000000) - int64(946684800000000)
	goto L46
L44:
	;
	goto L45
L45:
	;
	v216 = base.B2i32(v35 != int64(0))
	v217 = v216 | v32
	if v35 != int64(0) {
		v243 = v217
		v244 = v35
		goto L42
	} else {
		goto L50
	}
L46:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _consts[922]))
	goto L47
L47:
	;
	if base.B2i32(base.I64_extend_i32_s(v201)*int64(1000) <= v199-v35) == int32(0) {
		v243 = int32(0)
		v244 = v35
		goto L42
	} else {
		goto L48
	}
L48:
	;
	v209 = int32(1)
	F_LogRecoveryConflict(m, int32(12), v35, v199, int32(0), v209)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v243 = v209
	v244 = v35
	goto L42
L50:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, _consts[923])))
	if v219&int32(1) == int32(0) {
		v243 = v217
		v244 = v35
		goto L42
	} else {
		goto L51
	}
L51:
	;
	v227 = m.G0
	v228 = int32(16)
	v229 = v227 - v228
	m.G0 = v229
	F___gettimeofday(m, v229)
	mBase = m.M
	v232 = *(*int64)(unsafe.Add(mBase, uint32(v229)))
	v233 = int64(*(*int32)(unsafe.Add(mBase, uint32(v229)+8)))
	m.G0 = v229 + v228
	goto L52
L52:
	;
	v243 = v217
	v244 = v233 + v232*int64(1000000) - int64(946684800000000)
	goto L42
L53:
	;
	v249 = m.G0
	v251 = v249 + int32(-64)
	m.G0 = v251
	v256 = *(*int64)(unsafe.Add(mBase, _consts[308]))
	*(*int64)(unsafe.Add(mBase, uint32(v251))) = v256
	v259 = *(*int32)(unsafe.Add(mBase, _consts[306]))
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
	F___gettimeofday(m, v293)
	mBase = m.M
	v296 = *(*int64)(unsafe.Add(mBase, uint32(v293)))
	v297 = int64(*(*int32)(unsafe.Add(mBase, uint32(v293)+8)))
	m.G0 = v293 + v292
	goto L61
L56:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _consts[924]))
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
	v278 = *(*int32)(unsafe.Add(mBase, _consts[925]))
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
	if v287 == int64(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	F_ProcWaitForSignal(m, int32(67108864))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L72
	}
L63:
	;
	if v287 == int64(0) {
		goto L68
	} else {
		goto L69
	}
L64:
	;
	if v297+v296*int64(1000000)-int64(946684800000000) < v287 {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v309 = int32(0)
	F_CancelDBBackends(m, v309, int32(12), v309)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	goto L62
L67:
	;
	*(*int32)(unsafe.Add(mBase, _consts[926])) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v322))) = int64(4)
	v329 = *(*int32)(unsafe.Add(mBase, _consts[922]))
	*(*int32)(unsafe.Add(mBase, uint32(v322)+8)) = v329
	F_enable_timeouts(m, v251, v321)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L71
	}
L68:
	;
	v321 = int32(1)
	v322 = v251
	goto L67
L69:
	;
	goto L70
L70:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v251)+16)) = v287
	*(*int64)(unsafe.Add(mBase, uint32(v251))) = int64(4294967301)
	v321 = int32(2)
	v322 = v249 + int32(-40)
	goto L67
L71:
	;
	goto L62
L72:
	;
	v340 = *(*int32)(unsafe.Add(mBase, _consts[927]))
	if v340 != 0 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v352 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[643])) = v352
	*(*int32)(unsafe.Add(mBase, _consts[644])) = v352
	*(*uint8)(unsafe.Add(mBase, _consts[645])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[651])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[659])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[660])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[661])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[662])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[663])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[664])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[665])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[666])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[667])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[668])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[669])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[670])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[671])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[672])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[673])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[674])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[675])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[676])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[677])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[678])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[679])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[680])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[681])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[682])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[683])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[684])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[685])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[686])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[687])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[688])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[689])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[690])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[691])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[692])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[693])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[694])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[695])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[696])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[697])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[698])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[699])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[700])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[701])) = uint8(v352)
	*(*uint8)(unsafe.Add(mBase, _consts[702])) = uint8(v352)
	goto L79
L74:
	;
	v347 = int32(12)
	goto L76
L75:
	;
	v343 = *(*int32)(unsafe.Add(mBase, _consts[926]))
	if v343 == int32(0) {
		goto L73
	} else {
		goto L77
	}
L76:
	;
	F_CancelDBBackends(m, int32(0), v347, int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L78
	}
L77:
	;
	v347 = int32(13)
	goto L76
L78:
	;
	goto L73
L79:
	;
	v496 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[927])) = v496
	*(*int32)(unsafe.Add(mBase, _consts[926])) = v496
	m.G0 = v251 - int32(-64)
	v506 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	*(*int32)(unsafe.Add(mBase, uint32(v506)+72)) = int32(-1)
	goto L80
L80:
	;
	v513 = v243
	v515 = v244
	goto L38
L81:
	;
	v513 = v32
	v515 = v35
	goto L38
L82:
	;
	goto L85
L83:
	;
	v556 = v528
	goto L84
L84:
	;
	v569 = int32(4164732)
	v570 = *(*int32)(unsafe.Add(mBase, _consts[423]))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(8))+8))
	if v572 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L85:
	;
	F_perform_spin_delay(m, v14+int32(8))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L1
	} else {
		goto L87
	}
L86:
	;
	v556 = v549
	goto L84
L87:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	v550 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v549 | v550
	if v549&v550 != 0 {
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	if v556&int32(536870912) != 0 {
		goto L100
	} else {
		goto L101
	}
L90:
	;
	goto L89
L91:
	;
	*(*int32)(unsafe.Add(mBase, _consts[423])) = v587
	goto L90
L92:
	;
	if int32(999) < v570 {
		goto L90
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	if v570 < int32(11) {
		goto L90
	} else {
		goto L99
	}
L95:
	;
	v577 = int32(900)
	if v577 <= v570 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v580 = v577
	goto L98
L97:
	;
	v580 = v570
	goto L98
L98:
	;
	v587 = v580 + int32(100)
	goto L91
L99:
	;
	v587 = v570 - int32(1)
	goto L91
L100:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	v595 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v593 == v595 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	v598 = v556
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v598 & int32(-4194305)
	*(*int32)(unsafe.Add(mBase, _consts[921])) = int32(0)
	v606 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v28 = v606
	v32 = v513
	v35 = v515
	goto L5
L103:
	;
	v597 = v556 & int32(-536870913)
	goto L105
L104:
	;
	v597 = v556
	goto L105
L105:
	;
	v598 = v597
	goto L102
L106:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_errmsg_internal(m, int32(591448), int32(0))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(520922), int32(5742), int32(245401))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
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
	var v111 int64
	_ = v111
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
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
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
	var v224 int64
	_ = v224
	var v226 int32
	_ = v226
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
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
	var v275 int64
	_ = v275
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v359 int64
	_ = v359
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v379 int64
	_ = v379
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	v3 = int32(0)
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
	v58 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v60 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v63 = v60 + l0<<(uint(int32(6))%32)
	v65 = v63 - int32(40)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v67 = int32(276824064)
	if v66&v67 == v67 {
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
	v35 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v42 = v35 + (l0^int32(-1))<<(uint(int32(6))%32) + int32(24)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v43&int32(8388608) == int32(0) {
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
	F_errmsg_internal(m, int32(512111), v15)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	F_errfinish(m, int32(520922), int32(5436), int32(97626))
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
	v48 = int32(4460376)
	v50 = *(*int64)(unsafe.Add(mBase, _consts[11]))
	*(*int64)(unsafe.Add(mBase, _consts[11])) = v50 + int64(1)
	goto L14
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v43 | int32(8388608)
	goto L11
L15:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _consts[284]))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+252))
	goto L18
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = int32(241628)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = int32(520922)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+40)) = int64(0)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v287 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v286 | v287
	if v286&v287 != 0 {
		goto L73
	} else {
		goto L74
	}
L17:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if int32(0) <= v82 {
		v272 = v3
		v275 = int64(0)
		goto L16
	} else {
		goto L21
	}
L18:
	;
	if v73 != int32(0) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, _consts[335])))
	if v77 == int32(1) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v272 = v3
	v275 = int64(0)
	goto L16
L21:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, _consts[198])))
	if v87 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v97 != 0 {
		goto L1
	} else {
		goto L26
	}
L23:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+316))
	v95 = base.B2i32(v93 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[198])) = uint8(v95)
	v97 = v95
	goto L25
L24:
	;
	v97 = int32(0)
	goto L25
L25:
	;
	goto L22
L26:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-64))))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v100
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v63-int32(60))))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v104
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v63-int32(56))))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v108
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v15)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v111
	v115 = F_RelFileLocatorSkippingWAL(m, v13+int32(-48))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	if v115 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v117 = int32(1)
	v119 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v119)+120)) = v120 | v117
	v124 = m.G0
	v126 = v124 - int32(8224)
	m.G0 = v126
	v128 = F_GetRedoRecPtr(m)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L7
	} else {
		goto L30
	}
L29:
	;
	v272 = v117
	v275 = v251
	goto L16
L30:
	;
	v130 = F_BufferGetLSNAtomic(m, l0)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L7
	} else {
		goto L32
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L7
	} else {
		goto L70
	}
L32:
	;
	if base.Ui64(v130) <= base.Ui64(v128) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	if l0 < int32(0) {
		goto L39
	} else {
		goto L40
	}
L34:
	;
	v251 = int64(0)
	goto L35
L35:
	;
	m.G0 = v126 + int32(8224)
	goto L29
L36:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L7
	} else {
		goto L56
	}
L37:
	;
	goto L53
L38:
	;
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152)+14)))
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152)+12)))
	if v156 != 0 {
		goto L45
	} else {
		goto L46
	}
L39:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v136+(l0^int32(-1))<<(uint(int32(2))%32))))
	if l1 != 0 {
		v152 = v142
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v149 = v144 + l0<<(uint(int32(13))%32) + int32(-8192)
	if l1 == int32(0) {
		v167 = v149
		goto L37
	} else {
		goto L43
	}
L42:
	;
	v167 = v142
	goto L37
L43:
	;
	v152 = v149
	goto L38
L44:
	;
	v164 = int32(8192) - v153
	if v164 != 0 {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	v157 = F__emscripten_memcpy_bulkmem(m, v126+int32(32), v152, v156)
	mBase = m.M
	goto L47
L46:
	;
	goto L47
L47:
	;
	goto L44
L48:
	;
	goto L36
L49:
	;
	v165 = F__emscripten_memcpy_bulkmem(m, v153+(v126+int32(32)), v152+v153, v164)
	mBase = m.M
	goto L51
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
	v171 = F__emscripten_memcpy_bulkmem(m, v126+int32(32), v167, int32(8192))
	mBase = m.M
	goto L55
L55:
	;
	goto L52
L56:
	;
	v178 = v126 + int32(20)
	if l0 < int32(0) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	if v210 <= int32(0) {
		goto L62
	} else {
		goto L63
	}
L58:
	;
	v201 = *(*int64)(unsafe.Add(mBase, uint32(v200)))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v178)+8)) = v202
	*(*int64)(unsafe.Add(mBase, uint32(v178))) = v201
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v200)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v126+int32(16)))) = v205
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v200)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v126+int32(12)))) = v207
	goto L57
L59:
	;
	v187 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v200 = v187 + (l0^int32(-1))<<(uint(int32(6))%32)
	goto L58
L60:
	;
	goto L61
L61:
	;
	v194 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v200 = v194 + l0<<(uint(int32(6))%32) + int32(-64)
	goto L58
L62:
	;
	*(*int32)(unsafe.Add(mBase, _consts[130])) = int32(1)
	goto L64
L63:
	;
	goto L64
L64:
	;
	v217 = *(*int32)(unsafe.Add(mBase, _consts[131]))
	if v217 <= int32(0) {
		goto L31
	} else {
		goto L65
	}
L65:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
	v223 = *(*int32)(unsafe.Add(mBase, _consts[132]))
	v224 = *(*int64)(unsafe.Add(mBase, uint32(v126)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v223)+4)) = v224
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v126)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v223)+12)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v223)+20)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v223)+16)) = v221
	*(*int32)(unsafe.Add(mBase, uint32(v223)+36)) = v223 + int32(32)
	if l1 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v235 = int32(8)
	goto L68
L67:
	;
	v235 = int32(0)
	goto L68
L68:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v223)+1)) = uint8(v235)
	v237 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+28)) = v237
	v239 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v223))) = uint8(v239)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+24)) = v126 + int32(32)
	v246 = F_XLogInsert(m, v237, int32(160))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	v251 = v246
	goto L35
L70:
	;
	F_errmsg_internal(m, int32(144337), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L7
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(517998), int32(320), int32(333485))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L7
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	goto L76
L74:
	;
	v315 = v286
	goto L75
L75:
	;
	v329 = int32(4164732)
	v330 = *(*int32)(unsafe.Add(mBase, _consts[423]))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-24))+8))
	if v332 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L76:
	;
	F_perform_spin_delay(m, v13+int32(-24))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L7
	} else {
		goto L78
	}
L77:
	;
	v315 = v308
	goto L75
L78:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v309 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v308 | v309
	if v308&v309 != 0 {
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v350 = v315 & int32(8388608)
	if v350 != 0 {
		goto L91
	} else {
		goto L92
	}
L81:
	;
	goto L80
L82:
	;
	*(*int32)(unsafe.Add(mBase, _consts[423])) = v347
	goto L81
L83:
	;
	if int32(999) < v330 {
		goto L81
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	if v330 < int32(11) {
		goto L81
	} else {
		goto L90
	}
L86:
	;
	v337 = int32(900)
	if v337 <= v330 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v340 = v337
	goto L89
L88:
	;
	v340 = v330
	goto L89
L89:
	;
	v347 = v340 + int32(100)
	goto L82
L90:
	;
	v347 = v330 - int32(1)
	goto L82
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v315&int32(-281018369) | int32(276824064)
	if v272 != 0 {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	if v275 == int64(0) {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v355 = v58 + l0<<(uint(int32(13))%32)
	v359 = int64(base.Ui64(v275) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v355)+uint32(_consts[919]))) = uint32(v359)
	*(*uint32)(unsafe.Add(mBase, uint32(v355-int32(8188)))) = uint32(v275)
	goto L91
L94:
	;
	v371 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v371)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v371)+120)) = v372 & int32(-2)
	goto L96
L95:
	;
	goto L96
L96:
	;
	if v350 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v377 = int32(4460344)
	v379 = *(*int64)(unsafe.Add(mBase, _consts[7]))
	*(*int64)(unsafe.Add(mBase, _consts[7])) = v379 + int64(1)
	v384 = int32(*(*uint8)(unsafe.Add(mBase, _consts[726])))
	if v384 != int32(1) {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v387 = int32(4556816)
	v389 = *(*int32)(unsafe.Add(mBase, _consts[471]))
	v391 = *(*int32)(unsafe.Add(mBase, _consts[920]))
	*(*int32)(unsafe.Add(mBase, _consts[471])) = v389 + v391
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
	v11 = *(*int32)(unsafe.Add(mBase, _consts[182]))
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
	v102 = *(*int32)(unsafe.Add(mBase, _consts[182]))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_ResourceOwnerRemember(m, v102, v103+v100, int32(1665484))
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
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = int32(241628)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(520922)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(0)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v32 = int32(4194304)
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
	v60 = int32(4164732)
	v61 = *(*int32)(unsafe.Add(mBase, _consts[423]))
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
	v47 = int32(4194304)
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
	*(*int32)(unsafe.Add(mBase, _consts[423])) = v78
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
	if v55&int32(8388608) != 0 {
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
