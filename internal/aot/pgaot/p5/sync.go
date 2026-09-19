package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RegisterSyncRequest(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterSyncRequest[0]))
	if v6 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v40
L2:
	;
	v40 = int32(1)
	goto L1
L3:
	;
	v9 = F_ForwardSyncRequest(m, l0, l1)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	F_RememberSyncRequest(m, l0, l1)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L6
	} else {
		goto L14
	}
L6:
	;
	return int32(0)
L7:
	;
	if v9|base.B2i32(l2 == int32(0)) != 0 {
		v40 = v9
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L9
L9:
	;
	v24 = F_WaitLatch(m, int32(0), int32(40), int32(10), int32(150994949))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L6
	} else {
		goto L11
	}
L10:
	;
	goto L2
L11:
	;
	v26 = F_ForwardSyncRequest(m, l0, l1)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	if v26 == int32(0) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	goto L2
}
func F_SyncOneBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int64
	_ = v141
	var v143 int64
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int64
	_ = v178
	var v180 int64
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_SyncOneBuffer[0]))
	F_ReservePrivateRefCountEntry(m)
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
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_SyncOneBuffer[1]))
	F_ResourceOwnerEnlarge(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = int32(_a_F_SyncOneBuffer_0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(_a_F_SyncOneBuffer_1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(_a_F_SyncOneBuffer_2)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(0)
	v33 = v12 + l0<<(uint(int32(6))%32)
	v34 = int32(_a_F_SyncOneBuffer_3)
	v36 = base.AtomicRmwOr32(m, v33, int32(24), v34)
	if v36&v34 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	goto L7
L5:
	;
	v58 = v36
	goto L6
L6:
	;
	v61 = v58 & int32(_a_F_SyncOneBuffer_4)
	v65 = int32(_a_F_SyncOneBuffer_5)
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_SyncOneBuffer[2]))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(8))+8))
	if v68 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L7:
	;
	F_perform_spin_delay(m, v9+int32(8))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v58 = v51
	goto L6
L9:
	;
	v49 = int32(_a_F_SyncOneBuffer_3)
	v51 = base.AtomicRmwOr32(m, v33, int32(24), v49)
	if v51&v49 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v85 = int32(0)
	if base.B2i32(l1 == v85)|base.B2i32(v61 == v85) == v85 {
		goto L23
	} else {
		goto L24
	}
L12:
	;
	goto L11
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SyncOneBuffer[2])) = v83
	goto L12
L14:
	;
	if int32(999) < v66 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if v66 < int32(11) {
		goto L12
	} else {
		goto L21
	}
L17:
	;
	v73 = int32(900)
	if v73 <= v66 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v76 = v73
	goto L20
L19:
	;
	v76 = v66
	goto L20
L20:
	;
	v83 = v76 + int32(100)
	goto L13
L21:
	;
	v83 = v66 - int32(1)
	goto L13
L22:
	;
	m.G0 = v9 + int32(32)
	return v195
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = v58 & int32(-4194305)
	v195 = int32(0)
	goto L22
L24:
	;
	goto L25
L25:
	;
	if v61 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v98 = int32(0)
	goto L28
L27:
	;
	v98 = int32(2)
	goto L28
L28:
	;
	v99 = int32(25165824)
	if v58&v99 != v99 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = v58 & int32(-4194305)
	v195 = v98
	goto L22
L30:
	;
	goto L31
L31:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	v107 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = (v106 + v107) & int32(-4194305)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	v113 = int32(_a_F_SyncOneBuffer_6)
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_SyncOneBuffer[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_SyncOneBuffer[3])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v114)+4)) = v107
	v121 = v112 + v107
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = v121
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_SyncOneBuffer[1]))
	F_ResourceOwnerRemember(m, v124, v121, int32(_a_F_SyncOneBuffer_7))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v129 = v33 + int32(48)
	v131 = F_LWLockAcquire(m, v129, int32(1))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_FlushBuffer(m, v33, int32(0), int32(3))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_LWLockRelease(m, v129)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v139
	v141 = *(*int64)(unsafe.Add(mBase, uint32(v33)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v141
	v143 = *(*int64)(unsafe.Add(mBase, uint32(v33)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v143
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_SyncOneBuffer[1]))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	F_ResourceOwnerForget(m, v146, v147+int32(1), int32(_a_F_SyncOneBuffer_7))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_UnpinBufferNoOwner(m, v33)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SyncOneBuffer[4])))
	if v156&int32(1) != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v195 = v98 | int32(1)
	goto L22
L39:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SyncOneBuffer[5])))
	if v160&int32(1) == int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	if int32(0) < v166 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v169 + int32(1)
	v175 = l2 + v169*int32(20)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v175)+24)) = v176
	v178 = *(*int64)(unsafe.Add(mBase, uint32(v9)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v175)+16)) = v178
	v180 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v175)+8)) = v180
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	v185 = v183
	goto L43
L42:
	;
	v185 = v166
	goto L43
L43:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v186 < v185 {
		goto L38
	} else {
		goto L44
	}
L44:
	;
	F_IssuePendingWritebacks(m, l2, int32(3))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	goto L38
}
func F_SyncRepUpdateSyncStandbysDefined(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[0]))
	if v6 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	return
L2:
	;
	v291 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[1]))
	F_LWLockRelease(m, v291+int32(_a_F_SyncRepUpdateSyncStandbysDefined_0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L10
	} else {
		goto L76
	}
L3:
	;
	if v269&int32(1) != 0 {
		goto L1
	} else {
		goto L74
	}
L4:
	;
	v267 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[2]))
	*(*uint8)(unsafe.Add(mBase, uint32(v267)+48)) = uint8(v265)
	goto L2
L5:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[2]))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v48 = int32(0)
	if base.B2i32(v47 == v48)|base.B2i32(v47 == v46) == v48 {
		goto L15
	} else {
		goto L16
	}
L6:
	;
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[2]))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+48)))
	if v11&int32(2) == int32(0) {
		v269 = v11
		goto L3
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[2]))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+48)))
	v28 = int32(0)
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if base.B2i32(v25&int32(2) == v28)^base.B2i32(v30 != v28) != 0 {
		v269 = v25
		goto L3
	} else {
		goto L12
	}
L9:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[1]))
	v21 = F_LWLockAcquire(m, v17+int32(_a_F_SyncRepUpdateSyncStandbysDefined_0), int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	goto L5
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[1]))
	v39 = F_LWLockAcquire(m, v35+int32(_a_F_SyncRepUpdateSyncStandbysDefined_0), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	if v30 != 0 {
		v265 = int32(3)
		goto L4
	} else {
		goto L14
	}
L14:
	;
	goto L5
L15:
	;
	v54 = v47
	goto L18
L16:
	;
	v119 = v46
	goto L17
L17:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v119)+12))
	if v121 == int32(0) {
		v192 = v119
		goto L35
	} else {
		goto L36
	}
L18:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v59
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v54))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v54-int32(4)))) = int32(2)
	v70 = v54 - int32(120)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	if v71 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[2]))
	v119 = v116
	goto L17
L20:
	;
	if v46 != v59 {
		v54 = v59
		goto L18
	} else {
		goto L34
	}
L21:
	;
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = int32(1)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if v74 == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
	if v77 == int32(0) {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[3]))
	if v81 == v77 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v83 = m.G0
	v85 = v83 - int32(16)
	m.G0 = v85
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[4]))
	if v88 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v111 = F_pgmem_kill(m, v77, int32(23))
	mBase = m.M
	goto L21
L28:
	;
	m.G0 = v85 + int32(16)
	goto L20
L29:
	;
	v91 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v85)+15)) = uint8(v91)
	goto L30
L30:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[5]))
	v99 = F_write(m, v95, v85+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v99 {
		goto L28
	} else {
		goto L32
	}
L31:
	;
	goto L28
L32:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[6]))
	if v103 == int32(27) {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	goto L19
L35:
	;
	v194 = int32(1)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v192)+20))
	if v195 == int32(0) {
		v265 = v194
		goto L4
	} else {
		goto L55
	}
L36:
	;
	v125 = v119 + int32(8)
	if v121 == v125 {
		v192 = v119
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v127 = v121
	goto L38
L38:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v131)+4)) = v132
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	*(*int32)(unsafe.Add(mBase, uint32(v132))) = v134
	*(*int64)(unsafe.Add(mBase, uint32(v127))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v127-int32(4)))) = int32(2)
	v143 = v127 - int32(120)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	if v144 != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[2]))
	v192 = v189
	goto L35
L40:
	;
	if v125 != v132 {
		v127 = v132
		goto L38
	} else {
		goto L54
	}
L41:
	;
	goto L40
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v143))) = int32(1)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	if v147 == int32(0) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v143)+12))
	if v150 == int32(0) {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[3]))
	if v154 == v150 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v156 = m.G0
	v158 = v156 - int32(16)
	m.G0 = v158
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[4]))
	if v161 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	v184 = F_pgmem_kill(m, v150, int32(23))
	mBase = m.M
	goto L41
L48:
	;
	m.G0 = v158 + int32(16)
	goto L40
L49:
	;
	v164 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v158)+15)) = uint8(v164)
	goto L50
L50:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[5]))
	v172 = F_write(m, v168, v158+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v172 {
		goto L48
	} else {
		goto L52
	}
L51:
	;
	goto L48
L52:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[6]))
	if v176 == int32(27) {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	goto L39
L55:
	;
	v199 = v192 + int32(16)
	if v195 == v199 {
		v265 = v194
		goto L4
	} else {
		goto L56
	}
L56:
	;
	v201 = v195
	goto L57
L57:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v201)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v205)+4)) = v206
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	*(*int32)(unsafe.Add(mBase, uint32(v206))) = v208
	*(*int64)(unsafe.Add(mBase, uint32(v201))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v201-int32(4)))) = int32(2)
	v217 = v201 - int32(120)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	if v218 != 0 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v265 = v194
	goto L4
L59:
	;
	if v199 != v206 {
		v201 = v206
		goto L57
	} else {
		goto L73
	}
L60:
	;
	goto L59
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217))) = int32(1)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
	if v221 == int32(0) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v217)+12))
	if v224 == int32(0) {
		goto L60
	} else {
		goto L63
	}
L63:
	;
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[3]))
	if v228 == v224 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v230 = m.G0
	v232 = v230 - int32(16)
	m.G0 = v232
	v235 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[4]))
	if v235 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	v258 = F_pgmem_kill(m, v224, int32(23))
	mBase = m.M
	goto L60
L67:
	;
	m.G0 = v232 + int32(16)
	goto L59
L68:
	;
	v238 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v232)+15)) = uint8(v238)
	goto L69
L69:
	;
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[5]))
	v246 = F_write(m, v242, v232+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v246 {
		goto L67
	} else {
		goto L71
	}
L70:
	;
	goto L67
L71:
	;
	v250 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[6]))
	if v250 == int32(27) {
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	goto L58
L74:
	;
	v274 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[1]))
	v278 = F_LWLockAcquire(m, v274+int32(_a_F_SyncRepUpdateSyncStandbysDefined_0), int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L10
	} else {
		goto L75
	}
L75:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[2]))
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281)+48)))
	v284 = v282 | int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v281)+48)) = uint8(v284)
	goto L2
L76:
	;
	goto L1
}
func F_SyncRepWaitForLSN(m *base.Module, l0 int64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v1 int64
	_ = v1
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int64
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int64
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v100 int64
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v145 int64
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	v1 = l0
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[0]))
	if v11 <= int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(48)
	return
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[1]))
	if v15 < int32(2) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[2]))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+48)))
	if v20&int32(3) == int32(1) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[3]))
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[4]))
	v32 = F_LWLockAcquire(m, v28+int32(_a_F_SyncRepWaitForLSN_0), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	if int32(0) < v26 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v37 = int32(1)
	goto L9
L8:
	;
	v37 = v26
	goto L9
L9:
	;
	if l1 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v38 = v26
	goto L12
L11:
	;
	v38 = v37
	goto L12
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[2]))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+48)))
	if v41&int32(1) != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+136)) = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v79)+128)) = v1
	v85 = v40 + v38<<(uint(int32(3))%32)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v87 = int32(0)
	if base.B2i32(v86 == v87)|base.B2i32(v86 == v85) == v87 {
		goto L32
	} else {
		goto L33
	}
L14:
	;
	if v41&int32(2) != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v40+v38<<(uint(int32(3))%32))+24))
	if base.Ui64(v1) <= base.Ui64(v60) {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v40+v38<<(uint(int32(3))%32))+24))
	if base.Ui64(v49) < base.Ui64(v1) {
		goto L13
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[4]))
	F_LWLockRelease(m, v52+int32(_a_F_SyncRepWaitForLSN_0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L5
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	goto L1
L22:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[4]))
	F_LWLockRelease(m, v63+int32(_a_F_SyncRepWaitForLSN_0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[6]))
	if v69 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L1
L26:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v70 != 0 {
		goto L13
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[4]))
	F_LWLockRelease(m, v72+int32(_a_F_SyncRepWaitForLSN_0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L5
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	goto L1
L31:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[4]))
	F_LWLockRelease(m, v134+int32(_a_F_SyncRepWaitForLSN_0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L5
	} else {
		goto L44
	}
L32:
	;
	v94 = v86
	goto L35
L33:
	;
	goto L34
L34:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	if v117 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L35:
	;
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v94-int32(12))))
	if base.Ui64(v100) < base.Ui64(v1) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L34
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79)+140)) = v94
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+144)) = v103
	v106 = v79 + int32(140)
	*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v106
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v79)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v108))) = v106
	goto L31
L38:
	;
	goto L39
L39:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if v110 != v85 {
		v94 = v110
		goto L35
	} else {
		goto L40
	}
L40:
	;
	goto L36
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v85
	v121 = v85
	goto L43
L42:
	;
	v121 = v117
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79)+140)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v79)+144)) = v121
	v125 = v79 + int32(140)
	*(*int32)(unsafe.Add(mBase, uint32(v121))) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = v125
	goto L31
L44:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[7])))
	if v140 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+4)) = uint32(v1)
	v145 = int64(base.Ui64(v1) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v8))) = uint32(v145)
	v150 = F_pg_sprintf(m, v8+int32(16), int32(_a_F_SyncRepWaitForLSN_1), v8)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L5
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	goto L51
L48:
	;
	goto L47
L49:
	;
	v303 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[5]))
	*(*int64)(unsafe.Add(mBase, uint32(v303)+128)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v303)+136)) = int32(0)
	goto L1
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v290)+136)) = int32(0)
	v295 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[4]))
	F_LWLockRelease(m, v295+int32(_a_F_SyncRepWaitForLSN_0))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L5
	} else {
		goto L84
	}
L51:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v158))) = int32(0)
	goto L53
L52:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[9])) = int32(1)
	v268 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[10])) = v268
	v271 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[4]))
	v275 = F_LWLockAcquire(m, v271+int32(_a_F_SyncRepWaitForLSN_0), v268)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L5
	} else {
		goto L82
	}
L53:
	;
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[5]))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+136))
	if v163 == int32(2) {
		goto L49
	} else {
		goto L54
	}
L54:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[9]))
	if v167 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v170 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L5
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[11]))
	if v212 != 0 {
		goto L68
	} else {
		goto L69
	}
L58:
	;
	if v170 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L5
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v189 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[10])) = v189
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[4]))
	v196 = F_LWLockAcquire(m, v192+int32(_a_F_SyncRepWaitForLSN_0), v189)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L5
	} else {
		goto L66
	}
L62:
	;
	F_errmsg(m, int32(_a_F_SyncRepWaitForLSN_2), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L5
	} else {
		goto L63
	}
L63:
	;
	F_errdetail(m, int32(_a_F_SyncRepWaitForLSN_3), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L5
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_SyncRepWaitForLSN_4), int32(305), int32(_a_F_SyncRepWaitForLSN_5))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	goto L61
L66:
	;
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[5]))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)+144))
	if v200 == int32(0) {
		v290 = v199
		goto L50
	} else {
		goto L67
	}
L67:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v199)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+4)) = v200
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v199)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v205
	*(*int64)(unsafe.Add(mBase, uint32(v199)+140)) = int64(0)
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[5]))
	v290 = v210
	goto L50
L68:
	;
	v214 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[11])) = v214
	v218 = F_errstart(m, int32(19), v214)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L5
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v254 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[8]))
	v258 = F_WaitLatch(m, v254, int32(17), int32(-1), int32(134217780))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L5
	} else {
		goto L80
	}
L71:
	;
	if v218 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	F_errmsg(m, int32(_a_F_SyncRepWaitForLSN_6), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L5
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[4]))
	v238 = F_LWLockAcquire(m, v234+int32(_a_F_SyncRepWaitForLSN_0), int32(0))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L5
	} else {
		goto L78
	}
L75:
	;
	F_errdetail(m, int32(_a_F_SyncRepWaitForLSN_3), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L5
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_SyncRepWaitForLSN_4), int32(322), int32(_a_F_SyncRepWaitForLSN_5))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L5
	} else {
		goto L77
	}
L77:
	;
	goto L74
L78:
	;
	v241 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[5]))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v241)+144))
	if v242 == int32(0) {
		v290 = v241
		goto L50
	} else {
		goto L79
	}
L79:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v241)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v245)+4)) = v242
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v241)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v242))) = v247
	*(*int64)(unsafe.Add(mBase, uint32(v241)+140)) = int64(0)
	v252 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[5]))
	v290 = v252
	goto L50
L80:
	;
	if v258&int32(16) == int32(0) {
		goto L51
	} else {
		goto L81
	}
L81:
	;
	goto L52
L82:
	;
	v278 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[5]))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v278)+144))
	if v279 == int32(0) {
		v290 = v278
		goto L50
	} else {
		goto L83
	}
L83:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v278)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v282)+4)) = v279
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v278)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v279))) = v284
	*(*int64)(unsafe.Add(mBase, uint32(v278)+140)) = int64(0)
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[5]))
	v290 = v289
	goto L50
L84:
	;
	goto L49
}
