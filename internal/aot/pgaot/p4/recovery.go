package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
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
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v82 int64
	_ = v82
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int64
	_ = v94
	var v97 int64
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int64
	_ = v119
	var v121 int64
	_ = v121
	var v127 int32
	_ = v127
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_RecoveryRestartPoint[0]))
	if v12 != 0 {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v16 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
		v17 = *(*int64)(unsafe.Add(mBase, uint32(v15)+808))
		if v17 != int64(0) {
			v20 = *(*int64)(unsafe.Add(mBase, uint32(v15)+752))
			v21 = *(*int64)(unsafe.Add(mBase, uint32(v15)+728))
			v22 = *(*int64)(unsafe.Add(mBase, uint32(v15)+704))
			v23 = *(*int64)(unsafe.Add(mBase, uint32(v15)+680))
			v24 = *(*int64)(unsafe.Add(mBase, uint32(v15)+656))
			v25 = *(*int64)(unsafe.Add(mBase, uint32(v15)+632))
			v26 = *(*int64)(unsafe.Add(mBase, uint32(v15)+608))
			v27 = *(*int64)(unsafe.Add(mBase, uint32(v15)+584))
			v28 = *(*int64)(unsafe.Add(mBase, uint32(v15)+560))
			v29 = *(*int64)(unsafe.Add(mBase, uint32(v15)+536))
			v30 = *(*int64)(unsafe.Add(mBase, uint32(v15)+512))
			v31 = *(*int64)(unsafe.Add(mBase, uint32(v15)+488))
			v32 = *(*int64)(unsafe.Add(mBase, uint32(v15)+464))
			v33 = *(*int64)(unsafe.Add(mBase, uint32(v15)+440))
			v34 = *(*int64)(unsafe.Add(mBase, uint32(v15)+416))
			v35 = *(*int64)(unsafe.Add(mBase, uint32(v15)+392))
			v36 = *(*int64)(unsafe.Add(mBase, uint32(v15)+368))
			v37 = *(*int64)(unsafe.Add(mBase, uint32(v15)+344))
			v38 = *(*int64)(unsafe.Add(mBase, uint32(v15)+320))
			v39 = *(*int64)(unsafe.Add(mBase, uint32(v15)+296))
			v40 = *(*int64)(unsafe.Add(mBase, uint32(v15)+272))
			v41 = *(*int64)(unsafe.Add(mBase, uint32(v15)+248))
			v42 = *(*int64)(unsafe.Add(mBase, uint32(v15)+224))
			v43 = *(*int64)(unsafe.Add(mBase, uint32(v15)+200))
			v44 = *(*int64)(unsafe.Add(mBase, uint32(v15)+176))
			v45 = *(*int64)(unsafe.Add(mBase, uint32(v15)+152))
			v46 = *(*int64)(unsafe.Add(mBase, uint32(v15)+128))
			v47 = *(*int64)(unsafe.Add(mBase, uint32(v15)+104))
			v48 = *(*int64)(unsafe.Add(mBase, uint32(v15)+80))
			v49 = *(*int64)(unsafe.Add(mBase, uint32(v15)+56))
			v50 = *(*int64)(unsafe.Add(mBase, uint32(v15)+32))
			v82 = v20 + (v21 + (v22 + (v23 + (v24 + (v25 + (v26 + (v27 + (v28 + (v29 + (v30 + (v31 + (v32 + (v33 + (v34 + (v35 + (v36 + (v37 + (v38 + (v39 + (v40 + (v41 + (v42 + (v43 + (v44 + (v45 + (v46 + (v47 + (v48 + (v49 + (v50 + v16))))))))))))))))))))))))))))))
		} else {
			v82 = v16
		}
		if int64(0) < v82 {
			v87 = int32(1)
		} else {
			v87 = int32(0)
		}
	} else {
		v87 = int32(0)
	}
	if v87 != 0 {
		v90 = F_errstart(m, int32(13), int32(0))
		mBase = m.M
		v91 = m.ExcPending
		if v91 != 0 {
			return
		} else {
			if v90 == int32(0) {
				m.G0 = v9 + int32(16)
				return
			} else {
				v94 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				*(*uint32)(unsafe.Add(mBase, uint32(v9)+4)) = uint32(v94)
				v97 = int64(base.Ui64(v94) >> (uint(int64(32)) % 64))
				*(*uint32)(unsafe.Add(mBase, uint32(v9))) = uint32(v97)
				F_errmsg_internal(m, int32(_a_F_RecoveryRestartPoint_0), v9)
				mBase = m.M
				v101 = m.ExcPending
				if v101 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_RecoveryRestartPoint_1), int32(_a_F_RecoveryRestartPoint_2), int32(_a_F_RecoveryRestartPoint_3))
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return
					} else {
						m.G0 = v9 + int32(16)
						return
					}
				}
			}
		}
	} else {
		v108 = *(*int32)(unsafe.Add(mBase, _c_F_RecoveryRestartPoint[1]))
		v111 = base.AtomicRmwXchg32(m, v108, int32(440), int32(1))
		if v111 != 0 {
			F_s_lock(m, v108+int32(440), int32(_a_F_RecoveryRestartPoint_4))
			mBase = m.M
			v116 = m.ExcPending
			if v116 != 0 {
				return
			} else {
				v118 = *(*int32)(unsafe.Add(mBase, _c_F_RecoveryRestartPoint[1]))
				v119 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
				*(*int64)(unsafe.Add(mBase, uint32(v118)+320)) = v119
				v121 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
				*(*int64)(unsafe.Add(mBase, uint32(v118)+328)) = v121
				base.MemoryCopy(m, v118+int32(336), l0, int32(96))
				v127 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v118)+440)), uint32(v127))
				m.G0 = v9 + int32(16)
				return
			}
		} else {
			v118 = *(*int32)(unsafe.Add(mBase, _c_F_RecoveryRestartPoint[1]))
			v119 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
			*(*int64)(unsafe.Add(mBase, uint32(v118)+320)) = v119
			v121 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
			*(*int64)(unsafe.Add(mBase, uint32(v118)+328)) = v121
			base.MemoryCopy(m, v118+int32(336), l0, int32(96))
			v127 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v118)+440)), uint32(v127))
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
	var v9 int64
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v48 int32
	_ = v48
	var v51 int64
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int64
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v94 int64
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int64
	_ = v113
	var v117 int64
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int64
	_ = v128
	var v129 int64
	_ = v129
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int64
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int64
	_ = v202
	var v203 int64
	_ = v203
	var v212 int64
	_ = v212
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v249 int64
	_ = v249
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int64
	_ = v279
	var v280 int64
	_ = v280
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	v5 = int32(0)
	v9 = int64(0)
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(48)
	return
L2:
	;
	if l3 == int32(0) {
		v51 = v9
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v53 = l0
	v56 = int32(0)
	v59 = v5
	goto L11
L4:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[0])))
	if v21 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[1])))
	if v25&int32(1) == int32(0) {
		v51 = v9
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v33 = m.G0
	v34 = int32(16)
	v35 = v33 - v34
	m.G0 = v35
	F_gettimeofday(m, v35)
	mBase = m.M
	v38 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
	v39 = int64(*(*int32)(unsafe.Add(mBase, uint32(v35)+8)))
	m.G0 = v35 + v34
	goto L9
L8:
	;
	goto L7
L9:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v48 == int32(0) {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v51 = v39 + v38*int64(1000000) - int64(946684800000000)
	goto L3
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[2])) = int32(1000)
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v53)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v66
	v71 = F_VirtualXactLock(m, v13+int32(24), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v264 != 0 {
		goto L62
	} else {
		goto L63
	}
L13:
	;
	return
L14:
	;
	if v71 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v78 = v56
	v81 = v59
	goto L18
L16:
	;
	v261 = v56
	v264 = v59
	goto L17
L17:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	if v268 != 0 {
		v53 = v53 + int32(8)
		v56 = v261
		v59 = v264
		goto L11
	} else {
		goto L60
	}
L18:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[3]))
	if v86 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v261 = v244
	v264 = v246
	goto L17
L20:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L13
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v94 = *(*int64)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(40)))) = v94
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[5]))
	*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(39)))) = uint8(base.B2i32(v97 == int32(3)))
	goto L24
L23:
	;
	goto L22
L24:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+39)))
	if v101 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	if base.B2i32(v51 == int64(0))|v78&v81 != 0 {
		v244 = v78
		v246 = v81
		goto L42
	} else {
		goto L43
	}
L26:
	;
	v161 = *(*int64)(unsafe.Add(mBase, uint32(v53)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v161
	v165 = F_SignalRecoveryConflictWithVirtualXID(m, v13+int32(16), l1)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L13
	} else {
		goto L40
	}
L27:
	;
	v141 = int32(_a_F_ResolveRecoveryConflictWithVirtualXIDs_0)
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v142))) = l2
	v144 = int32(_a_F_ResolveRecoveryConflictWithVirtualXIDs_1)
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[2]))
	F_pg_usleep(m, v145)
	mBase = m.M
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = int32(0)
	v152 = int32(_a_F_ResolveRecoveryConflictWithVirtualXIDs_2)
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[2]))
	v156 = v154 << (uint(int32(1)) % 32)
	if v152 <= v156 {
		goto L37
	} else {
		goto L38
	}
L28:
	;
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
	v117 = v113 + base.I64_extend_i32_u(v112)*int64(1000)
	if v117 == int64(0) {
		goto L27
	} else {
		goto L34
	}
L29:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[7]))
	if int32(0) <= v105 {
		v112 = v105
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[8]))
	if v109 < int32(0) {
		goto L27
	} else {
		goto L33
	}
L32:
	;
	goto L27
L33:
	;
	v112 = v109
	goto L28
L34:
	;
	v123 = m.G0
	v124 = int32(16)
	v125 = v123 - v124
	m.G0 = v125
	F_gettimeofday(m, v125)
	mBase = m.M
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v125)))
	v129 = int64(*(*int32)(unsafe.Add(mBase, uint32(v125)+8)))
	m.G0 = v125 + v124
	goto L35
L35:
	;
	if v117 <= v129+v128*int64(1000000)-int64(946684800000000) {
		goto L26
	} else {
		goto L36
	}
L36:
	;
	goto L27
L37:
	;
	v159 = v152
	goto L39
L38:
	;
	v159 = v156
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[2])) = v159
	goto L25
L40:
	;
	if v165 == int32(0) {
		goto L25
	} else {
		goto L41
	}
L41:
	;
	F_pg_usleep(m, int32(_a_F_ResolveRecoveryConflictWithVirtualXIDs_3))
	mBase = m.M
	goto L25
L42:
	;
	v249 = *(*int64)(unsafe.Add(mBase, uint32(v53)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v249
	v254 = F_VirtualXactLock(m, v13+int32(8), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L13
	} else {
		goto L58
	}
L43:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[1])))
	v181 = v178 & (v78 ^ int32(1))
	v183 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[0])))
	v184 = int32(0)
	v186 = v183 & base.B2i32(v81 == v184)
	if v186 == v184 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v181&int32(1) == int32(0) {
		v227 = v78
		goto L50
	} else {
		goto L51
	}
L45:
	;
	if v181&int32(1) == int32(0) {
		v212 = int64(0)
		goto L44
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v197 = m.G0
	v198 = int32(16)
	v199 = v197 - v198
	m.G0 = v199
	F_gettimeofday(m, v199)
	mBase = m.M
	v202 = *(*int64)(unsafe.Add(mBase, uint32(v199)))
	v203 = int64(*(*int32)(unsafe.Add(mBase, uint32(v199)+8)))
	m.G0 = v199 + v198
	goto L49
L48:
	;
	goto L47
L49:
	;
	v212 = v203 + v202*int64(1000000) - int64(946684800000000)
	goto L44
L50:
	;
	if v186 == int32(0) {
		v244 = v227
		v246 = v81
		goto L42
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	if base.B2i32(base.I64_extend_i32_s(int32(500))*int64(1000) <= v212-v51) == int32(0) {
		v227 = int32(0)
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v227 = int32(1)
	goto L50
L54:
	;
	v232 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithVirtualXIDs[9]))
	goto L55
L55:
	;
	if base.B2i32(base.I64_extend_i32_s(v232)*int64(1000) <= v212-v51) == int32(0) {
		v244 = v227
		v246 = int32(0)
		goto L42
	} else {
		goto L56
	}
L56:
	;
	v240 = int32(1)
	F_LogRecoveryConflict(m, l1, v51, v212, v53, v240)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L13
	} else {
		goto L57
	}
L57:
	;
	v244 = v227
	v246 = v240
	goto L42
L58:
	;
	if v254 == int32(0) {
		v78 = v244
		v81 = v246
		goto L18
	} else {
		goto L59
	}
L59:
	;
	goto L19
L60:
	;
	goto L12
L61:
	;
	goto L1
L62:
	;
	v274 = m.G0
	v275 = int32(16)
	v276 = v274 - v275
	m.G0 = v276
	F_gettimeofday(m, v276)
	mBase = m.M
	v279 = *(*int64)(unsafe.Add(mBase, uint32(v276)))
	v280 = int64(*(*int32)(unsafe.Add(mBase, uint32(v276)+8)))
	m.G0 = v276 + v275
	goto L65
L63:
	;
	goto L64
L64:
	;
	if v261&int32(1) == int32(0) {
		goto L1
	} else {
		goto L68
	}
L65:
	;
	v289 = int32(0)
	F_LogRecoveryConflict(m, l1, v51, v280+v279*int64(1000000)-int64(946684800000000), v289, v289)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L13
	} else {
		goto L66
	}
L66:
	;
	if v261&int32(1) != 0 {
		goto L61
	} else {
		goto L67
	}
L67:
	;
	goto L1
L68:
	;
	goto L61
}
func F_SignalRecoveryConflictWithDatabase(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	v3 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_SignalRecoveryConflictWithDatabase[0]))
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_SignalRecoveryConflictWithDatabase[1]))
	v17 = F_LWLockAcquire(m, v13+int32(512), v3)
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
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if int32(0) < v19 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_SignalRecoveryConflictWithDatabase[2]))
	v30 = v19
	v32 = v27
	v33 = v3
	goto L6
L4:
	;
	goto L5
L5:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_SignalRecoveryConflictWithDatabase[1]))
	F_LWLockRelease(m, v74+int32(512))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L16
	}
L6:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(36)+v33<<(uint(int32(2))%32))))
	v43 = v32 + v40*int32(768)
	if l0 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L5
L8:
	;
	v62 = v33 + int32(1)
	if v62 < v58 {
		v30 = v58
		v32 = v59
		v33 = v62
		goto L6
	} else {
		goto L15
	}
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	if v44 != l0 {
		v58 = v30
		v59 = v32
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	if v46 == int32(0) {
		v58 = v30
		v59 = v32
		goto L8
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v43)+40))
	v51 = base.AtomicRmwOr32(m, v43, int32(340), int32(1)<<(uint(l1)%32))
	v53 = F_SendProcSignal(m, v46, int32(9), v49)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_SignalRecoveryConflictWithDatabase[2]))
	v58 = v55
	v59 = v57
	goto L8
L15:
	;
	goto L7
L16:
	;
	return
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v11 != 0 {
		v12 = int32(0)
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_check_recovery_target_lsn[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v14
		v17 = *(*int64)(unsafe.Add(mBase, _c_F_check_recovery_target_lsn[1]))
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = v17
		v19 = F_pg_lsn_in_safe(m, v10, v8)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+4)))
			if v23 != 0 {
				v35 = v12
				m.G0 = v8 + int32(16)
				return v35
			} else {
				v25 = F_guc_malloc(m, int32(8))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					if v25 == int32(0) {
						v35 = v12
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v25))) = v19
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v25
						v35 = int32(1)
					}
					m.G0 = v8 + int32(16)
					return v35
				}
			}
		}
	} else {
		v35 = int32(1)
		m.G0 = v8 + int32(16)
		return v35
	}
}
func F_check_recovery_target_timeline(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int64
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = int32(_a_F_check_recovery_target_timeline_0)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_recovery_target_timeline[0])))
	if base.B2i32(v16 == v4)|base.B2i32(v16 != v19) != 0 {
		v37 = v16
		v38 = v19
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v10 + int32(32)
	return v135
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_recovery_target_timeline[1])) = v130
	v135 = v4
	goto L1
L3:
	;
	v123 = F_guc_malloc(m, int32(4))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L27
	} else {
		goto L34
	}
L4:
	;
	if v37-v38 == int32(0) {
		v121 = v4
		goto L3
	} else {
		goto L11
	}
L5:
	;
	goto L4
L6:
	;
	v22 = v12
	v23 = v13
	goto L7
L7:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	if v27 == int32(0) {
		v37 = v27
		v38 = v26
		goto L5
	} else {
		goto L9
	}
L8:
	;
	v37 = v27
	v38 = v26
	goto L5
L9:
	;
	v30 = int32(1)
	if v27 == v26 {
		v22 = v22 + v30
		v23 = v23 + v30
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v42 = int32(_a_F_check_recovery_target_timeline_1)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_recovery_target_timeline[2])))
	if base.B2i32(v45 == int32(0))|base.B2i32(v45 != v48) != 0 {
		v66 = v45
		v67 = v48
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v66-v67 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	goto L12
L14:
	;
	v51 = v12
	v52 = v42
	goto L15
L15:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	if v56 == int32(0) {
		v66 = v56
		v67 = v55
		goto L13
	} else {
		goto L17
	}
L16:
	;
	v66 = v56
	v67 = v55
	goto L13
L17:
	;
	v59 = int32(1)
	if v56 == v55 {
		v51 = v51 + v59
		v52 = v52 + v59
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v121 = int32(1)
	goto L3
L20:
	;
	goto L21
L21:
	;
	v73 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_check_recovery_target_timeline[3])) = v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v80 = F_strtox_2(m, v75, v10+int32(28), v73, int64(-1))
	mBase = m.M
	goto L22
L22:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_check_recovery_target_timeline[3]))
	v88 = int32(0)
	if base.B2i32(v82|base.B2i32(v84 == int32(68)) == v88)&base.B2i32(v84 != int32(28)) == v88 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_recovery_target_timeline[4])) = v84
	goto L26
L24:
	;
	goto L25
L25:
	;
	if base.Ui64(v80-int64(4294967296)) <= base.Ui64(int64(-4294967296)) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_check_recovery_target_timeline_2)
	v102 = F_format_elog_string(m, int32(_a_F_check_recovery_target_timeline_3), v10+int32(16))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	return int32(0)
L28:
	;
	v130 = v102
	goto L2
L29:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_recovery_target_timeline[4])) = v84
	goto L32
L30:
	;
	goto L31
L31:
	;
	v121 = int32(2)
	goto L3
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+4)) = int64(-4294967295)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_check_recovery_target_timeline_2)
	v117 = F_format_elog_string(m, int32(_a_F_check_recovery_target_timeline_4), v10)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L27
	} else {
		goto L33
	}
L33:
	;
	v130 = v117
	goto L2
L34:
	;
	if v123 == int32(0) {
		v135 = v4
		goto L1
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123))) = v121
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v123
	v135 = int32(1)
	goto L1
}
func F_check_recovery_target_xid(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v37 int64
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v100 int32
	_ = v100
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v11 != 0 {
		*(*int32)(unsafe.Add(mBase, _c_F_check_recovery_target_xid[0])) = int32(0)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v16 = v15
		for {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
			if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v21-int32(9)))&base.B2i32(v21 != int32(32)) == int32(0) {
				v16 = v16 + int32(1)
				continue
			} else {
				break
			}
			break
		}
		v37 = F_strtox_2(m, v16, v8+int32(28), int32(0), int64(-1))
		mBase = m.M
		v39 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
		v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
		v42 = *(*int32)(unsafe.Add(mBase, _c_F_check_recovery_target_xid[0]))
		if v40|base.B2i32(v42 == int32(28))|base.B2i32(v42 == int32(68)) == int32(0) {
			v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
			if v51 != int32(45) {
				v63 = base.I32_wrap_i64(v37)
				if base.Ui32(int32(2)) < base.Ui32(v63) {
					v83 = F_guc_malloc(m, int32(4))
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return int32(0)
					} else {
						if v83 == int32(0) {
							v100 = int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v83))) = v63
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v83
							v100 = int32(1)
						}
						m.G0 = v8 + int32(32)
						return v100
					}
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_check_recovery_target_xid[1])) = v42
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(_a_F_check_recovery_target_xid_0)
					v75 = F_format_elog_string(m, int32(_a_F_check_recovery_target_xid_1), v8+int32(16))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						v78 = v75
						*(*int32)(unsafe.Add(mBase, _c_F_check_recovery_target_xid[2])) = v78
						v100 = int32(0)
						m.G0 = v8 + int32(32)
						return v100
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_check_recovery_target_xid[1])) = v42
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_check_recovery_target_xid_0)
				v59 = F_format_elog_string(m, int32(_a_F_check_recovery_target_xid_2), v8)
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					v78 = v59
					*(*int32)(unsafe.Add(mBase, _c_F_check_recovery_target_xid[2])) = v78
					v100 = int32(0)
					m.G0 = v8 + int32(32)
					return v100
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_check_recovery_target_xid[1])) = v42
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_check_recovery_target_xid_0)
			v59 = F_format_elog_string(m, int32(_a_F_check_recovery_target_xid_2), v8)
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return int32(0)
			} else {
				v78 = v59
				*(*int32)(unsafe.Add(mBase, _c_F_check_recovery_target_xid[2])) = v78
				v100 = int32(0)
				m.G0 = v8 + int32(32)
				return v100
			}
		}
	} else {
		v100 = int32(1)
		m.G0 = v8 + int32(32)
		return v100
	}
}
