package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BufferBeginSetHintBits(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	if l0 < int32(0) {
		v23 = int32(1)
		m.G0 = v5 + int32(16)
		return v23
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_BufferBeginSetHintBits[0]))
		v12 = int32(56)
		v19 = F_SharedBufferBeginSetHintBits(m, l0, v11+l0*v12-v12, v5+int32(8))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = v19
			m.G0 = v5 + int32(16)
			return v23
		}
	}
}
func F_BufferFinishSetHintBits(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v15 int64
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int64
	_ = v26
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int64
	_ = v49
	var v52 int64
	_ = v52
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l1 == int32(0) {
		m.G0 = v7 + int32(16)
		return
	} else {
		if l0 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(0)
				F_errmsg_internal(m, int32(_a_F_BufferFinishSetHintBits_0), v7)
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_BufferFinishSetHintBits_1), int32(_a_F_BufferFinishSetHintBits_2), int32(_a_F_BufferFinishSetHintBits_3))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			if l0 < int32(0) {
				v15 = int64(0)
				v17 = *(*int32)(unsafe.Add(mBase, _c_F_BufferFinishSetHintBits[0]))
				v22 = v17 + (l0^int32(-1))*int32(56)
				v26 = base.AtomicRmwCmpxchg64(m, v22, int32(24), v15, v15)
				if v26&int64(8388608) == v15 {
					v31 = int32(_a_F_BufferFinishSetHintBits_4)
					v33 = *(*int64)(unsafe.Add(mBase, _c_F_BufferFinishSetHintBits[1]))
					*(*int64)(unsafe.Add(mBase, _c_F_BufferFinishSetHintBits[1])) = v33 + int64(1)
				} else {
				}
				*(*int64)(unsafe.Add(mBase, uint32(v22)+24)) = v26 | int64(8388608)
				m.G0 = v7 + int32(16)
				return
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, _c_F_BufferFinishSetHintBits[2]))
				v42 = int32(56)
				v44 = v41 + l0*v42
				v49 = int64(0)
				v52 = base.AtomicRmwCmpxchg64(m, v44-int32(32), int32(0), v49, v49)
				F_MarkSharedBufferDirtyHint(m, l0, v44-v42, v52, l2)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return
				} else {
					m.G0 = v7 + int32(16)
					return
				}
			}
		}
	}
}
func F_BufferGetLSNAtomic(m *base.Module, l0 int32) int64 {
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
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v65 int64
	_ = v65
	var v74 int64
	_ = v74
	var v90 int32
	_ = v90
	var v91 int64
	_ = v91
	var v94 int64
	_ = v94
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v127 int64
	_ = v127
	var v129 int64
	_ = v129
	var v139 int64
	_ = v139
	var v142 int64
	_ = v142
	var v150 int64
	_ = v150
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BufferGetLSNAtomic[0])))
	if l0 < int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v8 + int32(32)
	return v150
L2:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_BufferGetLSNAtomic[1]))
	v53 = v48 + l0*int32(56) - int32(32)
	v54 = int64(4194304)
	v56 = base.AtomicRmwOr64(m, v53, int32(0), v54)
	if v56&v54 != int64(0) {
		goto L12
	} else {
		goto L13
	}
L3:
	;
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
	v150 = base.I64_rotl(v44, int64(32))
	goto L1
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_BufferGetLSNAtomic[2]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v15+(l0^int32(-1))<<(uint(int32(2))%32))))
	if v11&int32(1) != 0 {
		v43 = v21
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_BufferGetLSNAtomic[3]))
	v33 = v30 + l0<<(uint(int32(13))%32)
	if v11&int32(1) != 0 {
		goto L2
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v43 = v21
	goto L3
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_BufferGetLSNAtomic[4]))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+268))
	goto L10
L10:
	;
	if v40 != int32(0) {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v43 = v33 + int32(-8192)
	goto L3
L12:
	;
	v65 = v56
	goto L15
L13:
	;
	goto L14
L14:
	;
	v139 = *(*int64)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_BufferGetLSNAtomic[5])))
	v142 = base.AtomicRmwSub64(m, v53, int32(0), int64(4194304))
	v150 = base.I64_rotl(v139, int64(32))
	goto L1
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = int32(_a_F_BufferGetLSNAtomic_0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(_a_F_BufferGetLSNAtomic_1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_BufferGetLSNAtomic_2)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(0)
	v74 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v74
	if v65&int64(4194304) != v74 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L14
L17:
	;
	goto L20
L18:
	;
	goto L19
L19:
	;
	v107 = int32(_a_F_BufferGetLSNAtomic_3)
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_BufferGetLSNAtomic[6]))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(8))+8))
	if v110 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L20:
	;
	F_perform_spin_delay(m, v8+int32(8))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L19
L22:
	;
	return int64(0)
L23:
	;
	v91 = int64(0)
	v94 = base.AtomicRmwCmpxchg64(m, v53, int32(0), v91, v91)
	if v94&int64(4194304) != v91 {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v127 = int64(4194304)
	v129 = base.AtomicRmwOr64(m, v53, int32(0), v127)
	if v129&v127 != int64(0) {
		v65 = v129
		goto L15
	} else {
		goto L36
	}
L26:
	;
	goto L25
L27:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BufferGetLSNAtomic[6])) = v125
	goto L26
L28:
	;
	if int32(999) < v108 {
		goto L26
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	if v108 < int32(11) {
		goto L26
	} else {
		goto L35
	}
L31:
	;
	v115 = int32(900)
	if v115 <= v108 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v118 = v115
	goto L34
L33:
	;
	v118 = v108
	goto L34
L34:
	;
	v125 = v118 + int32(100)
	goto L27
L35:
	;
	v125 = v108 - int32(1)
	goto L27
L36:
	;
	goto L16
}
func F_IsBufferCleanupOK(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v62 int64
	_ = v62
	var v71 int64
	_ = v71
	var v85 int32
	_ = v85
	var v86 int64
	_ = v86
	var v89 int64
	_ = v89
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
	var v122 int64
	_ = v122
	var v124 int64
	_ = v124
	var v133 int64
	_ = v133
	var v136 int64
	_ = v136
	var v144 int32
	_ = v144
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	if l0 < v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(32)
	return v144
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_IsBufferCleanupOK[0]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13+(l0^int32(-1))<<(uint(int32(2))%32))))
	v144 = base.B2i32(v19 == int32(1))
	goto L1
L3:
	;
	goto L4
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_IsBufferCleanupOK[1]))
	if v23 != int32(-1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	if v41 != int32(1) {
		v144 = v2
		goto L1
	} else {
		goto L13
	}
L6:
	;
	v27 = v23 << (uint(int32(4)) % 32)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_IsBufferCleanupOK[2])))
	if v30 == l0 {
		v40 = v27 + int32(_a_F_IsBufferCleanupOK_0)
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v34 = F_GetPrivateRefCountEntrySlow(m, l0, int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L8
L10:
	;
	return int32(0)
L11:
	;
	if v34 == int32(0) {
		v144 = v2
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v40 = v34
	goto L5
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_IsBufferCleanupOK[3]))
	v50 = v45 + l0*int32(56) - int32(32)
	v51 = int64(4194304)
	v53 = base.AtomicRmwOr64(m, v50, int32(0), v51)
	if v53&v51 != int64(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v62 = v53
	goto L17
L15:
	;
	v133 = v53
	goto L16
L16:
	;
	v136 = base.AtomicRmwSub64(m, v50, int32(0), int64(4194304))
	v144 = base.B2i32(v133&int64(262143) == int64(1))
	goto L1
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = int32(_a_F_IsBufferCleanupOK_1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(_a_F_IsBufferCleanupOK_2)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_IsBufferCleanupOK_3)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(0)
	v71 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v71
	if v62&int64(4194304) != v71 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v133 = v124
	goto L16
L19:
	;
	goto L22
L20:
	;
	goto L21
L21:
	;
	v102 = int32(_a_F_IsBufferCleanupOK_4)
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_IsBufferCleanupOK[4]))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(8))+8))
	if v105 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L22:
	;
	F_perform_spin_delay(m, v8+int32(8))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L10
	} else {
		goto L24
	}
L23:
	;
	goto L21
L24:
	;
	v86 = int64(0)
	v89 = base.AtomicRmwCmpxchg64(m, v50, int32(0), v86, v86)
	if v89&int64(4194304) != v86 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v122 = int64(4194304)
	v124 = base.AtomicRmwOr64(m, v50, int32(0), v122)
	if v124&v122 != int64(0) {
		v62 = v124
		goto L17
	} else {
		goto L37
	}
L27:
	;
	goto L26
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IsBufferCleanupOK[4])) = v120
	goto L27
L29:
	;
	if int32(999) < v103 {
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if v103 < int32(11) {
		goto L27
	} else {
		goto L36
	}
L32:
	;
	v110 = int32(900)
	if v110 <= v103 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v113 = v110
	goto L35
L34:
	;
	v113 = v103
	goto L35
L35:
	;
	v120 = v113 + int32(100)
	goto L28
L36:
	;
	v120 = v103 - int32(1)
	goto L28
L37:
	;
	goto L18
}
func F_LockBufferForCleanup(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int64
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v56 int64
	_ = v56
	var v75 int64
	_ = v75
	var v96 int32
	_ = v96
	var v97 int64
	_ = v97
	var v100 int64
	_ = v100
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v140 int64
	_ = v140
	var v142 int64
	_ = v142
	var v148 int64
	_ = v148
	var v165 int64
	_ = v165
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int64
	_ = v177
	var v178 int64
	_ = v178
	var v181 int64
	_ = v181
	var v189 int64
	_ = v189
	var v192 int64
	_ = v192
	var v206 int64
	_ = v206
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int64
	_ = v244
	var v245 int64
	_ = v245
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v260 int64
	_ = v260
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int64
	_ = v286
	var v287 int64
	_ = v287
	var v295 int64
	_ = v295
	var v297 int32
	_ = v297
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int64
	_ = v327
	var v328 int64
	_ = v328
	var v338 int64
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v350 int64
	_ = v350
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v365 int64
	_ = v365
	var v372 int32
	_ = v372
	var v375 int64
	_ = v375
	var v381 int64
	_ = v381
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int64
	_ = v390
	var v391 int64
	_ = v391
	var v401 int32
	_ = v401
	var v409 int32
	_ = v409
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v591 int32
	_ = v591
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v608 int64
	_ = v608
	var v611 int32
	_ = v611
	var v613 int64
	_ = v613
	var v615 int64
	_ = v615
	var v621 int64
	_ = v621
	var v640 int64
	_ = v640
	var v661 int32
	_ = v661
	var v662 int64
	_ = v662
	var v665 int64
	_ = v665
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v703 int32
	_ = v703
	var v705 int64
	_ = v705
	var v707 int64
	_ = v707
	var v713 int64
	_ = v713
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v735 int64
	_ = v735
	var v736 int64
	_ = v736
	var v738 int64
	_ = v738
	var v743 int64
	_ = v743
	var v748 int64
	_ = v748
	var v761 int64
	_ = v761
	var v779 int32
	_ = v779
	var v782 int64
	_ = v782
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v806 int32
	_ = v806
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	F_CheckBufferIsPinnedOnce(m, l0)
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
	if l0 < int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v15 + int32(32)
	return
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[0]))
	v24 = l0 - int32(1)
	v27 = v22 + v24*int32(56)
	v31 = int64(0)
	v35 = v22
	v37 = int32(0)
	goto L5
L5:
	;
	v40 = int32(56)
	v41 = l0 * v40
	F_BufferLockAcquire(m, l0, v35+v41-v40, int32(3))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v782 = base.AtomicRmwSub64(m, v27, int32(24), int64(4194304))
	if int32(0) <= l0 {
		goto L132
	} else {
		goto L133
	}
L7:
	;
	v48 = int64(4194304)
	v50 = base.AtomicRmwOr64(m, v27, int32(24), v48)
	if v50&v48 != int64(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v56 = v50
	goto L11
L9:
	;
	v148 = v50
	goto L10
L10:
	;
	if v148&int64(262143) == int64(1) {
		goto L35
	} else {
		goto L36
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = int32(_a_F_LockBufferForCleanup_0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = int32(_a_F_LockBufferForCleanup_1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(_a_F_LockBufferForCleanup_2)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(0)
	v75 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v75
	if v56&int64(4194304) != v75 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v148 = v142
	goto L10
L13:
	;
	goto L16
L14:
	;
	goto L15
L15:
	;
	v120 = int32(_a_F_LockBufferForCleanup_3)
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[1]))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(8))+8))
	if v123 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L16:
	;
	F_perform_spin_delay(m, v15+int32(8))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	goto L15
L18:
	;
	v97 = int64(0)
	v100 = base.AtomicRmwCmpxchg64(m, v27, int32(24), v97, v97)
	if v100&int64(4194304) != v97 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v140 = int64(4194304)
	v142 = base.AtomicRmwOr64(m, v27, int32(24), v140)
	if v142&v140 != int64(0) {
		v56 = v142
		goto L11
	} else {
		goto L31
	}
L21:
	;
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[1])) = v138
	goto L21
L23:
	;
	if int32(999) < v121 {
		goto L21
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if v121 < int32(11) {
		goto L21
	} else {
		goto L30
	}
L26:
	;
	v128 = int32(900)
	if v128 <= v121 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v131 = v128
	goto L29
L28:
	;
	v131 = v121
	goto L29
L29:
	;
	v138 = v131 + int32(100)
	goto L22
L30:
	;
	v138 = v121 - int32(1)
	goto L22
L31:
	;
	goto L12
L32:
	;
	goto L6
L33:
	;
	v260 = base.AtomicRmwSub64(m, v27, int32(24), int64(4194304))
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[0]))
	F_BufferLockUnlock(m, l0, v262+v41-int32(56))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L51
	}
L34:
	;
	if v37 != 0 {
		goto L46
	} else {
		goto L47
	}
L35:
	;
	v165 = base.AtomicRmwSub64(m, v27, int32(24), int64(4194304))
	goto L34
L36:
	;
	goto L37
L37:
	;
	if v148&int64(536870912) != int64(0) {
		goto L32
	} else {
		goto L38
	}
L38:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v171
	*(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[3])) = v27
	v176 = int32(24)
	v177 = base.AtomicRmwOr64(m, v27, v176, int64(536870912))
	v178 = int64(0)
	v181 = base.AtomicRmwCmpxchg64(m, v27, v176, v178, v178)
	if v181&int64(262143) != int64(1) {
		goto L33
	} else {
		goto L39
	}
L39:
	;
	v189 = base.AtomicRmwCmpxchg64(m, v27, int32(24), v181, v181&int64(-541327359))
	if v181 != v189 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v192 = v189
	goto L43
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[3])) = int32(0)
	goto L34
L43:
	;
	v206 = base.AtomicRmwCmpxchg64(m, v27, int32(24), v192, v192&int64(-541065217))
	if v192 != v206 {
		v192 = v206
		goto L43
	} else {
		goto L45
	}
L44:
	;
	goto L42
L45:
	;
	goto L44
L46:
	;
	v239 = m.G0
	v240 = int32(16)
	v241 = v239 - v240
	m.G0 = v241
	F_gettimeofday(m, v241)
	mBase = m.M
	v244 = *(*int64)(unsafe.Add(mBase, uint32(v241)))
	v245 = int64(*(*int32)(unsafe.Add(mBase, uint32(v241)+8)))
	m.G0 = v241 + v240
	goto L49
L47:
	;
	goto L48
L48:
	;
	goto L3
L49:
	;
	v254 = int32(0)
	F_LogRecoveryConflict(m, int32(5), v31, v245+v244*int64(1000000)-int64(946684800000000), v254, v254)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v269 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[4]))
	if base.Ui32(int32(2)) <= base.Ui32(v269) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v613 = int64(4194304)
	v615 = base.AtomicRmwOr64(m, v27, int32(24), v613)
	if v615&v613 != int64(0) {
		goto L96
	} else {
		goto L97
	}
L53:
	;
	if v37|base.B2i32(v31 == int64(0)) == int32(0) {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	goto L55
L55:
	;
	F_ProcWaitForSignal(m, int32(67108864))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L1
	} else {
		goto L95
	}
L56:
	;
	v341 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v341)+76)) = v24
	goto L67
L57:
	;
	v281 = m.G0
	v282 = int32(16)
	v283 = v281 - v282
	m.G0 = v283
	F_gettimeofday(m, v283)
	mBase = m.M
	v286 = *(*int64)(unsafe.Add(mBase, uint32(v283)))
	v287 = int64(*(*int32)(unsafe.Add(mBase, uint32(v283)+8)))
	m.G0 = v283 + v282
	v295 = v287 + v286*int64(1000000) - int64(946684800000000)
	goto L60
L58:
	;
	goto L59
L59:
	;
	if v31 != int64(0) {
		v338 = v31
		v339 = v37
		goto L56
	} else {
		goto L64
	}
L60:
	;
	v297 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[6]))
	goto L61
L61:
	;
	if base.B2i32(base.I64_extend_i32_s(v297)*int64(1000) <= v295-v31) == int32(0) {
		v338 = v31
		v339 = int32(0)
		goto L56
	} else {
		goto L62
	}
L62:
	;
	v305 = int32(1)
	F_LogRecoveryConflict(m, int32(5), v31, v295, int32(0), v305)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v338 = v31
	v339 = v305
	goto L56
L64:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[7])))
	if v314&int32(1) == int32(0) {
		v338 = v31
		v339 = v37
		goto L56
	} else {
		goto L65
	}
L65:
	;
	v322 = m.G0
	v323 = int32(16)
	v324 = v322 - v323
	m.G0 = v324
	F_gettimeofday(m, v324)
	mBase = m.M
	v327 = *(*int64)(unsafe.Add(mBase, uint32(v324)))
	v328 = int64(*(*int32)(unsafe.Add(mBase, uint32(v324)+8)))
	m.G0 = v324 + v323
	goto L66
L66:
	;
	v338 = v328 + v327*int64(1000000) - int64(946684800000000)
	v339 = v37
	goto L56
L67:
	;
	v343 = m.G0
	v345 = v343 + int32(-64)
	m.G0 = v345
	v350 = *(*int64)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[8]))
	*(*int64)(unsafe.Add(mBase, uint32(v345))) = v350
	v353 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[9]))
	*(*uint8)(unsafe.Add(mBase, uint32(v343+int32(-1)))) = uint8(base.B2i32(v353 == int32(3)))
	goto L68
L68:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345)+63)))
	if v357 == int32(1) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v385 = m.G0
	v386 = int32(16)
	v387 = v385 - v386
	m.G0 = v387
	F_gettimeofday(m, v387)
	mBase = m.M
	v390 = *(*int64)(unsafe.Add(mBase, uint32(v387)))
	v391 = int64(*(*int32)(unsafe.Add(mBase, uint32(v387)+8)))
	m.G0 = v387 + v386
	goto L75
L70:
	;
	v362 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[10]))
	if v362 < int32(0) {
		v381 = int64(0)
		goto L69
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v372 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[11]))
	if v372 < int32(0) {
		v381 = int64(0)
		goto L69
	} else {
		goto L74
	}
L73:
	;
	v365 = *(*int64)(unsafe.Add(mBase, uint32(v345)))
	v381 = v365 + base.I64_extend_i32_u(v362)*int64(1000)
	goto L69
L74:
	;
	v375 = *(*int64)(unsafe.Add(mBase, uint32(v345)))
	v381 = v375 + base.I64_extend_i32_u(v372)*int64(1000)
	goto L69
L75:
	;
	v401 = base.B2i32(v381 == int64(0))
	if v401|base.B2i32(v391+v390*int64(1000000)-int64(946684800000000) < v381) == int32(0) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	F_ProcWaitForSignal(m, int32(67108864))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L86
	}
L77:
	;
	F_SignalRecoveryConflictWithDatabase(m, int32(0), int32(5))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	if v381 == int64(0) {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L76
L81:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[12])) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v418))) = int64(4)
	v425 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v418)+8)) = v425
	F_enable_timeouts(m, v345, v417)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L85
	}
L82:
	;
	v417 = int32(1)
	v418 = v345
	goto L81
L83:
	;
	goto L84
L84:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v345)+16)) = v381
	*(*int64)(unsafe.Add(mBase, uint32(v345))) = int64(4294967301)
	v417 = int32(2)
	v418 = v343 + int32(-40)
	goto L81
L85:
	;
	goto L76
L86:
	;
	v436 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[13]))
	if v436 != 0 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v447 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[14])) = v447
	*(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[15])) = v447
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[16])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[17])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[18])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[19])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[20])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[21])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[22])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[23])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[24])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[25])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[26])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[27])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[28])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[29])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[30])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[31])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[32])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[33])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[34])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[35])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[36])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[37])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[38])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[39])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[40])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[41])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[42])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[43])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[44])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[45])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[46])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[47])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[48])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[49])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[50])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[51])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[52])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[53])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[54])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[55])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[56])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[57])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[58])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[59])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[60])) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[61])) = uint8(v447)
	goto L93
L88:
	;
	v443 = int32(5)
	goto L90
L89:
	;
	v439 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[12]))
	if v439 == int32(0) {
		goto L87
	} else {
		goto L91
	}
L90:
	;
	F_SignalRecoveryConflictWithDatabase(m, int32(0), v443)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L92
	}
L91:
	;
	v443 = int32(7)
	goto L90
L92:
	;
	goto L87
L93:
	;
	v591 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[13])) = v591
	*(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[12])) = v591
	m.G0 = v345 - int32(-64)
	v601 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v601)+76)) = int32(-1)
	goto L94
L94:
	;
	v608 = v338
	v611 = v339
	goto L52
L95:
	;
	v608 = v31
	v611 = v37
	goto L52
L96:
	;
	v621 = v615
	goto L99
L97:
	;
	v713 = v615
	goto L98
L98:
	;
	if v713&int64(536870912) != int64(0) {
		goto L120
	} else {
		goto L121
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = int32(_a_F_LockBufferForCleanup_0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = int32(_a_F_LockBufferForCleanup_1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(_a_F_LockBufferForCleanup_2)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(0)
	v640 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v640
	if v621&int64(4194304) != v640 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v713 = v707
	goto L98
L101:
	;
	goto L104
L102:
	;
	goto L103
L103:
	;
	v685 = int32(_a_F_LockBufferForCleanup_3)
	v686 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[1]))
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(8))+8))
	if v688 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L104:
	;
	F_perform_spin_delay(m, v15+int32(8))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L1
	} else {
		goto L106
	}
L105:
	;
	goto L103
L106:
	;
	v662 = int64(0)
	v665 = base.AtomicRmwCmpxchg64(m, v27, int32(24), v662, v662)
	if v665&int64(4194304) != v662 {
		goto L104
	} else {
		goto L107
	}
L107:
	;
	goto L105
L108:
	;
	v705 = int64(4194304)
	v707 = base.AtomicRmwOr64(m, v27, int32(24), v705)
	if v707&v705 != int64(0) {
		v621 = v707
		goto L99
	} else {
		goto L119
	}
L109:
	;
	goto L108
L110:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[1])) = v703
	goto L109
L111:
	;
	if int32(999) < v686 {
		goto L109
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	if v686 < int32(11) {
		goto L109
	} else {
		goto L118
	}
L114:
	;
	v693 = int32(900)
	if v693 <= v686 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v696 = v693
	goto L117
L116:
	;
	v696 = v686
	goto L117
L117:
	;
	v703 = v696 + int32(100)
	goto L110
L118:
	;
	v703 = v686 - int32(1)
	goto L110
L119:
	;
	goto L100
L120:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v27)+32))
	v733 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[2]))
	if v731 == v733 {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	v736 = int64(-1)
	goto L122
L122:
	;
	v738 = v713 | int64(4194304)
	v743 = base.AtomicRmwCmpxchg64(m, v27, int32(24), v738, v713&v736&int64(-4194305))
	if v743 != v738 {
		goto L126
	} else {
		goto L127
	}
L123:
	;
	v735 = int64(-536870913)
	goto L125
L124:
	;
	v735 = int64(-1)
	goto L125
L125:
	;
	v736 = v735
	goto L122
L126:
	;
	v748 = v743
	goto L129
L127:
	;
	goto L128
L128:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[3])) = int32(0)
	v779 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[0]))
	v31 = v608
	v35 = v779
	v37 = v611
	goto L5
L129:
	;
	v761 = base.AtomicRmwCmpxchg64(m, v27, int32(24), v748, v748&(v736&int64(-4194305)))
	if v748 != v761 {
		v748 = v761
		goto L129
	} else {
		goto L131
	}
L130:
	;
	goto L128
L131:
	;
	goto L130
L132:
	;
	v786 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufferForCleanup[0]))
	v787 = int32(56)
	F_BufferLockUnlock(m, l0, v786+l0*v787-v787)
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L1
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L1
	} else {
		goto L136
	}
L135:
	;
	goto L134
L136:
	;
	F_errmsg_internal(m, int32(_a_F_LockBufferForCleanup_4), int32(0))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_LockBufferForCleanup_2), int32(_a_F_LockBufferForCleanup_5), int32(_a_F_LockBufferForCleanup_6))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_MarkBufferDirtyHint(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int64
	_ = v21
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v44 int64
	_ = v44
	var v47 int64
	_ = v47
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l0 != 0 {
		if l0 < int32(0) {
			v10 = int64(0)
			v12 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[0]))
			v17 = v12 + (l0^int32(-1))*int32(56)
			v21 = base.AtomicRmwCmpxchg64(m, v17, int32(24), v10, v10)
			if v21&int64(8388608) == v10 {
				v26 = int32(_a_F_MarkBufferDirtyHint_0)
				v28 = *(*int64)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[1]))
				*(*int64)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[1])) = v28 + int64(1)
			} else {
			}
			*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v21 | int64(8388608)
			m.G0 = v6 + int32(16)
			return
		} else {
			v36 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirtyHint[2]))
			v37 = int32(56)
			v39 = v36 + l0*v37
			v44 = int64(0)
			v47 = base.AtomicRmwCmpxchg64(m, v39-int32(32), int32(0), v44, v44)
			F_MarkSharedBufferDirtyHint(m, l0, v39-v37, v47, l1)
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return
			} else {
				m.G0 = v6 + int32(16)
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v57 = m.ExcPending
		if v57 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(0)
			F_errmsg_internal(m, int32(_a_F_MarkBufferDirtyHint_1), v6)
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_MarkBufferDirtyHint_2), int32(_a_F_MarkBufferDirtyHint_3), int32(_a_F_MarkBufferDirtyHint_4))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
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
