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
	var v35 int32
	_ = v35
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int64
	_ = v143
	var v145 int64
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int64
	_ = v180
	var v182 int64
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
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
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	v35 = int32(_a_F_SyncOneBuffer_3)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = v34 | v35
	if v34&v35 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	goto L7
L5:
	;
	v59 = v34
	goto L6
L6:
	;
	v63 = v59 & int32(_a_F_SyncOneBuffer_4)
	v67 = int32(_a_F_SyncOneBuffer_5)
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_SyncOneBuffer[2]))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(8))+8))
	if v70 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L7:
	;
	F_perform_spin_delay(m, v9+int32(8))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v59 = v50
	goto L6
L9:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	v51 = int32(_a_F_SyncOneBuffer_3)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = v50 | v51
	if v50&v51 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v87 = int32(0)
	if base.B2i32(l1 == v87)|base.B2i32(v63 == v87) == v87 {
		goto L23
	} else {
		goto L24
	}
L12:
	;
	goto L11
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SyncOneBuffer[2])) = v85
	goto L12
L14:
	;
	if int32(999) < v68 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if v68 < int32(11) {
		goto L12
	} else {
		goto L21
	}
L17:
	;
	v75 = int32(900)
	if v75 <= v68 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v78 = v75
	goto L20
L19:
	;
	v78 = v68
	goto L20
L20:
	;
	v85 = v78 + int32(100)
	goto L13
L21:
	;
	v85 = v68 - int32(1)
	goto L13
L22:
	;
	m.G0 = v9 + int32(32)
	return v197
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = v59 & int32(-4194305)
	v197 = int32(0)
	goto L22
L24:
	;
	goto L25
L25:
	;
	if v63 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v100 = int32(0)
	goto L28
L27:
	;
	v100 = int32(2)
	goto L28
L28:
	;
	v101 = int32(25165824)
	if v59&v101 != v101 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = v59 & int32(-4194305)
	v197 = v100
	goto L22
L30:
	;
	goto L31
L31:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	v109 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = (v108 + v109) & int32(-4194305)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	v115 = int32(_a_F_SyncOneBuffer_6)
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_SyncOneBuffer[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_SyncOneBuffer[3])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v116)+4)) = v109
	v123 = v114 + v109
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v123
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_SyncOneBuffer[1]))
	F_ResourceOwnerRemember(m, v126, v123, int32(_a_F_SyncOneBuffer_7))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v131 = v33 + int32(48)
	v133 = F_LWLockAcquire(m, v131, int32(1))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_FlushBuffer(m, v33, int32(0), int32(3))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_LWLockRelease(m, v131)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v141
	v143 = *(*int64)(unsafe.Add(mBase, uint32(v33)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v143
	v145 = *(*int64)(unsafe.Add(mBase, uint32(v33)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v145
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_SyncOneBuffer[1]))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	F_ResourceOwnerForget(m, v148, v149+int32(1), int32(_a_F_SyncOneBuffer_7))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_UnpinBufferNoOwner(m, v33)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SyncOneBuffer[4])))
	if v158&int32(1) != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v197 = v100 | int32(1)
	goto L22
L39:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SyncOneBuffer[5])))
	if v162&int32(1) == int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	if int32(0) < v168 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v171 + int32(1)
	v177 = l2 + v171*int32(20)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v177)+24)) = v178
	v180 = *(*int64)(unsafe.Add(mBase, uint32(v9)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v177)+16)) = v180
	v182 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v177)+8)) = v182
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v187 = v185
	goto L43
L42:
	;
	v187 = v168
	goto L43
L43:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v188 < v187 {
		goto L38
	} else {
		goto L44
	}
L44:
	;
	F_IssuePendingWritebacks(m, l2, int32(3))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
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
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
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
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[1]))
	F_LWLockRelease(m, v168+int32(_a_F_SyncRepUpdateSyncStandbysDefined_0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L10
	} else {
		goto L37
	}
L3:
	;
	if v146&int32(1) != 0 {
		goto L1
	} else {
		goto L35
	}
L4:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[2]))
	*(*uint8)(unsafe.Add(mBase, uint32(v144)+48)) = uint8(v142)
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
		v146 = v11
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
		v146 = v25
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
		v142 = int32(3)
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
	v78 = v46
	goto L17
L17:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	if v80 == int32(0) {
		v110 = v78
		goto L22
	} else {
		goto L23
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
	F_SetLatch(m, v54-int32(120))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L10
	} else {
		goto L20
	}
L19:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[2]))
	v78 = v75
	goto L17
L20:
	;
	if v46 != v59 {
		v54 = v59
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v112 = int32(1)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	if v113 == int32(0) {
		v142 = v112
		goto L4
	} else {
		goto L29
	}
L23:
	;
	v84 = v78 + int32(8)
	if v80 == v84 {
		v110 = v78
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v86 = v80
	goto L25
L25:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v90)+4)) = v91
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = v93
	*(*int64)(unsafe.Add(mBase, uint32(v86))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v86-int32(4)))) = int32(2)
	F_SetLatch(m, v86-int32(120))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L10
	} else {
		goto L27
	}
L26:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[2]))
	v110 = v107
	goto L22
L27:
	;
	if v84 != v91 {
		v86 = v91
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v117 = v110 + int32(16)
	if v113 == v117 {
		v142 = v112
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v119 = v113
	goto L31
L31:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v123)+4)) = v124
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	*(*int32)(unsafe.Add(mBase, uint32(v124))) = v126
	*(*int64)(unsafe.Add(mBase, uint32(v119))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v119-int32(4)))) = int32(2)
	F_SetLatch(m, v119-int32(120))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L10
	} else {
		goto L33
	}
L32:
	;
	v142 = v112
	goto L4
L33:
	;
	if v117 != v124 {
		v119 = v124
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[1]))
	v155 = F_LWLockAcquire(m, v151+int32(_a_F_SyncRepUpdateSyncStandbysDefined_0), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L10
	} else {
		goto L36
	}
L36:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[2]))
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+48)))
	v161 = v159 | int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v158)+48)) = uint8(v161)
	goto L2
L37:
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
