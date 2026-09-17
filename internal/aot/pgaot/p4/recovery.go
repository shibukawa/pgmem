package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RecoveryRestartPoint(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int64
	_ = v92
	var v95 int64
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int64
	_ = v121
	var v123 int64
	_ = v123
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_RecoveryRestartPoint[0]))
	if v12 != 0 {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+412))
		if v17 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+376))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+364))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)+352))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+340))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)+328))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v15)+316))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)+304))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v15)+292))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v15)+280))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)+268))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v15)+256))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v15)+244))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)+232))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v15)+220))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v15)+208))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v15)+196))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v15)+184))
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v15)+172))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v15)+160))
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v15)+148))
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v15)+136))
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v15)+124))
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v15)+112))
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v15)+100))
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v15)+76))
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
			v80 = v18 + (v19 + (v20 + (v21 + (v22 + (v23 + (v24 + (v25 + (v26 + (v27 + (v28 + (v29 + (v30 + (v31 + (v32 + (v33 + (v34 + (v35 + (v36 + (v37 + (v38 + (v39 + (v40 + (v41 + (v42 + (v43 + (v44 + (v45 + (v46 + (v47 + (v48 + v16))))))))))))))))))))))))))))))
		} else {
			v80 = v16
		}
		if int32(0) < v80 {
			v85 = int32(1)
		} else {
			v85 = int32(0)
		}
	} else {
		v85 = int32(0)
	}
	if v85 != 0 {
		v88 = F_errstart(m, int32(13), int32(0))
		mBase = m.M
		v89 = m.ExcPending
		if v89 != 0 {
			return
		} else {
			if v88 == int32(0) {
				m.G0 = v9 + int32(16)
				return
			} else {
				v92 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				*(*uint32)(unsafe.Add(mBase, uint32(v9)+4)) = uint32(v92)
				v95 = int64(base.Ui64(v92) >> (uint(int64(32)) % 64))
				*(*uint32)(unsafe.Add(mBase, uint32(v9))) = uint32(v95)
				F_errmsg_internal(m, int32(_a_F_RecoveryRestartPoint_0), v9)
				mBase = m.M
				v99 = m.ExcPending
				if v99 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_RecoveryRestartPoint_1), int32(_a_F_RecoveryRestartPoint_2), int32(_a_F_RecoveryRestartPoint_3))
					mBase = m.M
					v104 = m.ExcPending
					if v104 != 0 {
						return
					} else {
						m.G0 = v9 + int32(16)
						return
					}
				}
			}
		}
	} else {
		v106 = *(*int32)(unsafe.Add(mBase, _c_F_RecoveryRestartPoint[1]))
		v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+440))
		*(*int32)(unsafe.Add(mBase, uint32(v106)+440)) = int32(1)
		if v107 != 0 {
			v111 = *(*int32)(unsafe.Add(mBase, _c_F_RecoveryRestartPoint[1]))
			F_s_lock(m, v111+int32(440), int32(_a_F_RecoveryRestartPoint_1), int32(_a_F_RecoveryRestartPoint_4), int32(_a_F_RecoveryRestartPoint_3))
			mBase = m.M
			v118 = m.ExcPending
			if v118 != 0 {
				return
			} else {
				v120 = *(*int32)(unsafe.Add(mBase, _c_F_RecoveryRestartPoint[1]))
				v121 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
				*(*int64)(unsafe.Add(mBase, uint32(v120)+328)) = v121
				v123 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
				*(*int64)(unsafe.Add(mBase, uint32(v120)+336)) = v123
				base.MemoryCopy(m, v120+int32(344), l0, int32(88))
				*(*int32)(unsafe.Add(mBase, uint32(v120)+440)) = int32(0)
				m.G0 = v9 + int32(16)
				return
			}
		} else {
			v120 = *(*int32)(unsafe.Add(mBase, _c_F_RecoveryRestartPoint[1]))
			v121 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
			*(*int64)(unsafe.Add(mBase, uint32(v120)+328)) = v121
			v123 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
			*(*int64)(unsafe.Add(mBase, uint32(v120)+336)) = v123
			base.MemoryCopy(m, v120+int32(344), l0, int32(88))
			*(*int32)(unsafe.Add(mBase, uint32(v120)+440)) = int32(0)
			m.G0 = v9 + int32(16)
			return
		}
	}
}
func F_ResolveRecoveryConflictWithVirtualXIDs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v53 int32
	_ = v53
	var v56 int64
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v76 int64
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v109 int64
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int64
	_ = v128
	var v132 int64
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int64
	_ = v143
	var v144 int64
	_ = v144
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int64
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
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
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v242 int32
	_ = v242
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int64
	_ = v308
	var v309 int64
	_ = v309
	var v318 int64
	_ = v318
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v355 int64
	_ = v355
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v379 int32
	_ = v379
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
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	v5 = int32(0)
	v14 = int64(0)
	v16 = m.G0
	v18 = v16 - int32(48)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v20 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v18 + int32(48)
	return
L2:
	;
	if l3 == int32(0) {
		v56 = v14
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v58 = l0
	v61 = int32(0)
	v65 = v5
	goto L11
L4:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[0])))
	if v26 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[1])))
	if v30&int32(1) == int32(0) {
		v56 = v14
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v38 = m.G0
	v39 = int32(16)
	v40 = v38 - v39
	m.G0 = v40
	F_gettimeofday(m, v40)
	mBase = m.M
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v40)))
	v44 = int64(*(*int32)(unsafe.Add(mBase, uint32(v40)+8)))
	m.G0 = v40 + v39
	goto L9
L8:
	;
	goto L7
L9:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v53 == int32(0) {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v56 = v44 + v43*int64(1000000) - int64(946684800000000)
	goto L3
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[2])) = int32(1000)
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v58)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v76
	v81 = F_VirtualXactLock(m, v18+int32(24), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v371 != 0 {
		goto L76
	} else {
		goto L77
	}
L13:
	;
	return
L14:
	;
	if v81 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v88 = v61
	v92 = v65
	goto L18
L16:
	;
	v367 = v61
	v371 = v65
	goto L17
L17:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	if v379 != 0 {
		v58 = v58 + int32(8)
		v61 = v367
		v65 = v371
		goto L11
	} else {
		goto L74
	}
L18:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[3]))
	if v101 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v367 = v350
	v371 = v353
	goto L17
L20:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L13
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v109 = *(*int64)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v18+int32(40)))) = v109
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[5]))
	*(*uint8)(unsafe.Add(mBase, uint32(v18+int32(39)))) = uint8(base.B2i32(v112 == int32(3)))
	goto L24
L23:
	;
	goto L22
L24:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+39)))
	if v116 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	if base.B2i32(v56 == int64(0))|v88&v92 != 0 {
		v350 = v88
		v353 = v92
		goto L56
	} else {
		goto L57
	}
L26:
	;
	v177 = *(*int64)(unsafe.Add(mBase, uint32(v58)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v177
	v180 = v18 + int32(16)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[6]))
	v186 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[7]))
	v190 = F_LWLockAcquire(m, v186+int32(512), int32(1))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L13
	} else {
		goto L41
	}
L27:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v157))) = l2
	v160 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[2]))
	F_pg_usleep(m, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L13
	} else {
		goto L37
	}
L28:
	;
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v18)+40))
	v132 = v128 + base.I64_extend_i32_u(v127)*int64(1000)
	if v132 == int64(0) {
		goto L27
	} else {
		goto L34
	}
L29:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[9]))
	if int32(0) <= v120 {
		v127 = v120
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[10]))
	if v124 < int32(0) {
		goto L27
	} else {
		goto L33
	}
L32:
	;
	goto L27
L33:
	;
	v127 = v124
	goto L28
L34:
	;
	v138 = m.G0
	v139 = int32(16)
	v140 = v138 - v139
	m.G0 = v140
	F_gettimeofday(m, v140)
	mBase = m.M
	v143 = *(*int64)(unsafe.Add(mBase, uint32(v140)))
	v144 = int64(*(*int32)(unsafe.Add(mBase, uint32(v140)+8)))
	m.G0 = v140 + v139
	goto L35
L35:
	;
	if v132 <= v144+v143*int64(1000000)-int64(946684800000000) {
		goto L26
	} else {
		goto L36
	}
L36:
	;
	goto L27
L37:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v164))) = int32(0)
	v167 = int32(_a_F_ResolveRecoveryConflictWithVirtualXIDs_0)
	v168 = int32(_a_F_ResolveRecoveryConflictWithVirtualXIDs_1)
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[2]))
	v172 = v170 << (uint(int32(1)) % 32)
	if v168 <= v172 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v175 = v168
	goto L40
L39:
	;
	v175 = v172
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[2])) = v175
	goto L25
L41:
	;
	v192 = int32(0)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	if v193 <= v192 {
		v242 = v192
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v254 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[7]))
	F_LWLockRelease(m, v254+int32(512))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L13
	} else {
		goto L53
	}
L43:
	;
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[11]))
	v204 = v192
	goto L44
L44:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v184+int32(36)+v204<<(uint(int32(2))%32))))
	v221 = v199 + v218*int32(640)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+52))
	if v222 != v182 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v242 = int32(0)
	goto L42
L46:
	;
	goto L45
L47:
	;
	v234 = v204 + int32(1)
	if v234 != v193 {
		v204 = v234
		goto L44
	} else {
		goto L52
	}
L48:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v221)+56))
	if v224 != v181 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v226 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v221)+73)) = uint8(v226)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v221)+44))
	if v228 == int32(0) {
		goto L46
	} else {
		goto L50
	}
L50:
	;
	v231 = F_SendProcSignal(m, v228, l1, v182)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L13
	} else {
		goto L51
	}
L51:
	;
	v242 = v228
	goto L42
L52:
	;
	goto L46
L53:
	;
	if v242 == int32(0) {
		goto L25
	} else {
		goto L54
	}
L54:
	;
	F_pg_usleep(m, int32(_a_F_ResolveRecoveryConflictWithVirtualXIDs_2))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L13
	} else {
		goto L55
	}
L55:
	;
	goto L25
L56:
	;
	v355 = *(*int64)(unsafe.Add(mBase, uint32(v58)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v355
	v360 = F_VirtualXactLock(m, v18+int32(8), int32(0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L13
	} else {
		goto L72
	}
L57:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[1])))
	v287 = v284 & (v88 ^ int32(1))
	v289 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[0])))
	v290 = int32(0)
	v292 = v289 & base.B2i32(v92 == v290)
	if v292 == v290 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	if v287&int32(1) == int32(0) {
		v333 = v88
		goto L64
	} else {
		goto L65
	}
L59:
	;
	if v287&int32(1) == int32(0) {
		v318 = int64(0)
		goto L58
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v303 = m.G0
	v304 = int32(16)
	v305 = v303 - v304
	m.G0 = v305
	F_gettimeofday(m, v305)
	mBase = m.M
	v308 = *(*int64)(unsafe.Add(mBase, uint32(v305)))
	v309 = int64(*(*int32)(unsafe.Add(mBase, uint32(v305)+8)))
	m.G0 = v305 + v304
	goto L63
L62:
	;
	goto L61
L63:
	;
	v318 = v309 + v308*int64(1000000) - int64(946684800000000)
	goto L58
L64:
	;
	if v292 == int32(0) {
		v350 = v333
		v353 = v92
		goto L56
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	if base.B2i32(base.I64_extend_i32_s(int32(500))*int64(1000) <= v318-v56) == int32(0) {
		v333 = int32(0)
		goto L64
	} else {
		goto L67
	}
L67:
	;
	v333 = int32(1)
	goto L64
L68:
	;
	v338 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[12]))
	goto L69
L69:
	;
	if base.B2i32(base.I64_extend_i32_s(v338)*int64(1000) <= v318-v56) == int32(0) {
		v350 = v333
		v353 = int32(0)
		goto L56
	} else {
		goto L70
	}
L70:
	;
	v346 = int32(1)
	F_LogRecoveryConflict(m, l1, v56, v318, v58, v346)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L13
	} else {
		goto L71
	}
L71:
	;
	v350 = v333
	v353 = v346
	goto L56
L72:
	;
	if v360 == int32(0) {
		v88 = v350
		v92 = v353
		goto L18
	} else {
		goto L73
	}
L73:
	;
	goto L19
L74:
	;
	goto L12
L75:
	;
	goto L1
L76:
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
	goto L79
L77:
	;
	goto L78
L78:
	;
	if v367&int32(1) == int32(0) {
		goto L1
	} else {
		goto L82
	}
L79:
	;
	v400 = int32(0)
	F_LogRecoveryConflict(m, l1, v56, v391+v390*int64(1000000)-int64(946684800000000), v400, v400)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L13
	} else {
		goto L80
	}
L80:
	;
	if v367&int32(1) != 0 {
		goto L75
	} else {
		goto L81
	}
L81:
	;
	goto L1
L82:
	;
	goto L75
}
func F_check_recovery_prefetch(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return int32(1)
}
func F_check_recovery_target(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	v4 = int32(1)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = int32(_a_F_check_recovery_target_0)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_recovery_target[0])))
	if base.B2i32(v9 == int32(0))|base.B2i32(v9 != v12) != 0 {
		v30 = v9
		v31 = v12
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v51
L2:
	;
	if v30-v31 == int32(0) {
		v51 = v4
		goto L1
	} else {
		goto L9
	}
L3:
	;
	goto L2
L4:
	;
	v15 = v5
	v16 = v6
	goto L5
L5:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	if v20 == int32(0) {
		v30 = v20
		v31 = v19
		goto L3
	} else {
		goto L7
	}
L6:
	;
	v30 = v20
	v31 = v19
	goto L3
L7:
	;
	v23 = int32(1)
	if v20 == v19 {
		v15 = v15 + v23
		v16 = v16 + v23
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v35 == int32(0) {
		v51 = v4
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_check_recovery_target[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_recovery_target[2])) = v40
	goto L11
L11:
	;
	v46 = F_format_elog_string(m, int32(_a_F_check_recovery_target_1), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_recovery_target[3])) = v46
	v51 = int32(0)
	goto L1
}
func F_check_recovery_target_lsn(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int64
	_ = v127
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v218 int64
	_ = v218
	var v226 int64
	_ = v226
	var v232 int64
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v250 int32
	_ = v250
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v14 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v250
L2:
	;
	v15 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v15)
	v19 = v11 + int32(15)
	*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v15)
	v22 = int32(_a_F_check_recovery_target_lsn_0)
	v26 = m.G0
	v28 = v26 - int32(32)
	v29 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+24)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v28)+16)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v28))) = v29
	v37 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_recovery_target_lsn[0])))
	if v37 == v15 {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	goto L4
L4:
	;
	v250 = int32(1)
	goto L1
L5:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
	if v233 != 0 {
		v250 = v15
		goto L1
	} else {
		goto L56
	}
L6:
	;
	v119 = v110 + int32(1)
	v120 = int32(_a_F_check_recovery_target_lsn_0)
	v124 = m.G0
	v126 = v124 - int32(32)
	v127 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v126)+24)) = v127
	*(*int64)(unsafe.Add(mBase, uint32(v126)+16)) = v127
	*(*int64)(unsafe.Add(mBase, uint32(v126)+8)) = v127
	*(*int64)(unsafe.Add(mBase, uint32(v126))) = v127
	v135 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_recovery_target_lsn[0])))
	if v135 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L7:
	;
	if base.Ui32(int32(-8)) <= base.Ui32(v105-int32(9)) {
		goto L26
	} else {
		goto L27
	}
L8:
	;
	v105 = int32(0)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_recovery_target_lsn[1])))
	if v41 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v45 = v13
	goto L14
L12:
	;
	goto L13
L13:
	;
	v55 = v22
	v56 = v37
	goto L17
L14:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v51 == v37 {
		v45 = v45 + int32(1)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v105 = v45 - v13
	goto L7
L16:
	;
	goto L15
L17:
	;
	v63 = v28 + int32(base.Ui32(v56)>>(uint(int32(3))%32))&int32(28)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v65 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v64 | v65<<(uint(v56)%32)
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	if v69 != 0 {
		v55 = v55 + v65
		v56 = v69
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v72 == int32(0) {
		v95 = v13
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	v105 = v95 - v13
	goto L7
L21:
	;
	v76 = v13
	v77 = v72
	goto L22
L22:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(base.Ui32(v77)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v85)>>(uint(v77)%32))&int32(1) == int32(0) {
		v95 = v76
		goto L20
	} else {
		goto L24
	}
L23:
	;
	v95 = v93
	goto L20
L24:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
	v93 = v76 + int32(1)
	if v91 != 0 {
		v76 = v93
		v77 = v91
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v110 = v13 + v105
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v111 == int32(47) {
		goto L6
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v115 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v115)
	v232 = int64(0)
	goto L5
L29:
	;
	goto L28
L30:
	;
	v218 = F_strtox_2(m, v13, int32(0), int32(16), int64(4294967295))
	mBase = m.M
	goto L54
L31:
	;
	if base.Ui32(int32(-8)) <= base.Ui32(v203-int32(9)) {
		goto L50
	} else {
		goto L51
	}
L32:
	;
	v203 = int32(0)
	goto L31
L33:
	;
	goto L34
L34:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_recovery_target_lsn[1])))
	if v139 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v143 = v119
	goto L38
L36:
	;
	goto L37
L37:
	;
	v153 = v120
	v154 = v135
	goto L41
L38:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	if v149 == v135 {
		v143 = v143 + int32(1)
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v203 = v143 - v119
	goto L31
L40:
	;
	goto L39
L41:
	;
	v161 = v126 + int32(base.Ui32(v154)>>(uint(int32(3))%32))&int32(28)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	v163 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v161))) = v162 | v163<<(uint(v154)%32)
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+1)))
	if v167 != 0 {
		v153 = v153 + v163
		v154 = v167
		goto L41
	} else {
		goto L43
	}
L42:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	if v170 == int32(0) {
		v193 = v119
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L42
L44:
	;
	v203 = v193 - v119
	goto L31
L45:
	;
	v174 = v119
	v175 = v170
	goto L46
L46:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v126+int32(base.Ui32(v175)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v183)>>(uint(v175)%32))&int32(1) == int32(0) {
		v193 = v174
		goto L44
	} else {
		goto L48
	}
L47:
	;
	v193 = v191
	goto L44
L48:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
	v191 = v174 + int32(1)
	if v189 != 0 {
		v174 = v191
		v175 = v189
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119+v203))))
	if v209 == int32(0) {
		goto L30
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v212 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v212)
	v232 = int64(0)
	goto L5
L53:
	;
	goto L52
L54:
	;
	v226 = F_strtox_2(m, v119, int32(0), int32(16), int64(4294967295))
	mBase = m.M
	goto L55
L55:
	;
	v232 = base.I64_extend_i32_u(base.I32_wrap_i64(v218))<<(uint(int64(32))%64) | base.I64_extend_i32_u(base.I32_wrap_i64(v226))
	goto L5
L56:
	;
	v235 = F_guc_malloc(m, int32(8))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	return int32(0)
L58:
	;
	if v235 == int32(0) {
		v250 = v15
		goto L1
	} else {
		goto L59
	}
L59:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v235))) = v232
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v235
	goto L4
}
func F_check_recovery_target_timeline(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int64
	_ = v73
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	v4 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = int32(_a_F_check_recovery_target_timeline_0)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_recovery_target_timeline[0])))
	if base.B2i32(v10 == v4)|base.B2i32(v10 != v13) != 0 {
		v31 = v10
		v32 = v13
		goto L4
	} else {
		goto L5
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_recovery_target_timeline[1])) = v76
	goto L28
L2:
	;
	v86 = F_guc_malloc(m, int32(4))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L23
	} else {
		goto L24
	}
L3:
	;
	if v31-v32 == int32(0) {
		v84 = v4
		goto L2
	} else {
		goto L10
	}
L4:
	;
	goto L3
L5:
	;
	v16 = v6
	v17 = v7
	goto L6
L6:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v21 == int32(0) {
		v31 = v21
		v32 = v20
		goto L4
	} else {
		goto L8
	}
L7:
	;
	v31 = v21
	v32 = v20
	goto L4
L8:
	;
	v24 = int32(1)
	if v21 == v20 {
		v16 = v16 + v24
		v17 = v17 + v24
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v36 = int32(_a_F_check_recovery_target_timeline_1)
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_recovery_target_timeline[2])))
	if base.B2i32(v39 == int32(0))|base.B2i32(v39 != v42) != 0 {
		v60 = v39
		v61 = v42
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v60-v61 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	goto L11
L13:
	;
	v45 = v6
	v46 = v36
	goto L14
L14:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+1)))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+1)))
	if v50 == int32(0) {
		v60 = v50
		v61 = v49
		goto L12
	} else {
		goto L16
	}
L15:
	;
	v60 = v50
	v61 = v49
	goto L12
L16:
	;
	v53 = int32(1)
	if v50 == v49 {
		v45 = v45 + v53
		v46 = v46 + v53
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v84 = int32(1)
	goto L2
L19:
	;
	goto L20
L20:
	;
	v67 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_check_recovery_target_timeline[3])) = v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v73 = F_strtox_2(m, v69, v67, v67, int64(4294967295))
	mBase = m.M
	goto L21
L21:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_check_recovery_target_timeline[3]))
	if base.B2i32(v76 == int32(68))|base.B2i32(v76 == int32(28)) != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v84 = int32(2)
	goto L2
L23:
	;
	return int32(0)
L24:
	;
	if v86 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = v84
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v86
	v94 = int32(1)
	goto L27
L26:
	;
	v94 = int32(0)
	goto L27
L27:
	;
	return v94
L28:
	;
	v101 = F_format_elog_string(m, int32(_a_F_check_recovery_target_timeline_2), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_recovery_target_timeline[4])) = v101
	return int32(0)
}
func F_check_recovery_target_xid(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int64
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v7 == int32(0) {
		v37 = int32(1)
		return v37
	} else {
		v10 = int32(0)
		*(*int32)(unsafe.Add(mBase, _c_F_check_recovery_target_xid[0])) = v10
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v18 = F_strtox_2(m, v14, v10, v10, int64(-1))
		mBase = m.M
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_check_recovery_target_xid[0]))
		if base.B2i32(v20 == int32(28))|base.B2i32(v20 == int32(68)) != 0 {
			v37 = v10
			return v37
		} else {
			v27 = F_guc_malloc(m, int32(4))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				if v27 == int32(0) {
					v37 = v10
				} else {
					*(*uint32)(unsafe.Add(mBase, uint32(v27))) = uint32(v18)
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v27
					v37 = int32(1)
				}
				return v37
			}
		}
	}
}
