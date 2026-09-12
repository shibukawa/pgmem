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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	v6 = *(*int32)(unsafe.Add(mBase, _consts[803]))
	if v6 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v39
L2:
	;
	v39 = int32(1)
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
	v30 = m.ExcPending
	if v30 != 0 {
		goto L6
	} else {
		goto L15
	}
L6:
	;
	return int32(0)
L7:
	;
	if v9 != 0 {
		v39 = v9
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if l2 == int32(0) {
		v39 = v9
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L10
L10:
	;
	v23 = F_WaitLatch(m, int32(0), int32(40), int32(10), int32(150994949))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L6
	} else {
		goto L12
	}
L11:
	;
	goto L2
L12:
	;
	v25 = F_ForwardSyncRequest(m, l0, l1)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	if v25 == int32(0) {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
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
	var v60 int32
	_ = v60
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
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int64
	_ = v176
	var v178 int64
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[17]))
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
	v18 = *(*int32)(unsafe.Add(mBase, _consts[11]))
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
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = int32(239002)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(514763)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(0)
	v33 = v12 + l0<<(uint(int32(6))%32)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	v35 = int32(4194304)
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
	v60 = v34
	goto L6
L6:
	;
	v63 = v60 & int32(4194303)
	v67 = int32(4155052)
	v68 = *(*int32)(unsafe.Add(mBase, _consts[602]))
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
	v60 = v50
	goto L6
L9:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	v51 = int32(4194304)
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
	if l1 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L12:
	;
	goto L11
L13:
	;
	*(*int32)(unsafe.Add(mBase, _consts[602])) = v85
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
	return v195
L23:
	;
	v98 = base.B2i32(v63 == int32(0)) << (uint(int32(1)) % 32)
	v99 = int32(25165824)
	if v60&v99 != v99 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	if v63 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = v60 & int32(-4194305)
	v195 = int32(0)
	goto L22
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = v60 & int32(-4194305)
	v195 = v98
	goto L22
L27:
	;
	goto L28
L28:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	v107 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = (v106 + v107) & int32(-4194305)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	v113 = int32(4464680)
	v114 = *(*int32)(unsafe.Add(mBase, _consts[619]))
	*(*int32)(unsafe.Add(mBase, _consts[619])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v114)+4)) = v107
	v121 = v112 + v107
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = v121
	v124 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	F_ResourceOwnerRemember(m, v124, v121, int32(1656256))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v129 = v33 + int32(48)
	v131 = F_LWLockAcquire(m, v129, int32(1))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_FlushBuffer(m, v33, int32(0), int32(3))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_LWLockRelease(m, v129)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v139
	v141 = *(*int64)(unsafe.Add(mBase, uint32(v33)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v141
	v143 = *(*int64)(unsafe.Add(mBase, uint32(v33)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v143
	v146 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	F_ResourceOwnerForget(m, v146, v147+int32(1), int32(1656256))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_UnpinBufferNoOwner(m, v33)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, _consts[622])))
	if v156&int32(1) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v195 = v98 | int32(1)
	goto L22
L36:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, _consts[623])))
	if v160 != int32(1) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	if int32(0) < v164 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v167 + int32(1)
	v173 = l2 + v167*int32(20)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+24)) = v174
	v176 = *(*int64)(unsafe.Add(mBase, uint32(v9)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v173)+16)) = v176
	v178 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v173)+8)) = v178
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	v183 = v181
	goto L40
L39:
	;
	v183 = v164
	goto L40
L40:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v184 < v183 {
		goto L35
	} else {
		goto L41
	}
L41:
	;
	F_IssuePendingWritebacks(m, l2, int32(3))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	goto L35
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	v6 = *(*int32)(unsafe.Add(mBase, _consts[563]))
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
	v165 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v165+int32(4096))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L10
	} else {
		goto L37
	}
L3:
	;
	if v143&int32(1) != 0 {
		goto L1
	} else {
		goto L35
	}
L4:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _consts[561]))
	*(*uint8)(unsafe.Add(mBase, uint32(v141)+48)) = uint8(v139)
	goto L2
L5:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[561]))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v47 == int32(0) {
		v75 = v46
		goto L15
	} else {
		goto L16
	}
L6:
	;
	v10 = *(*int32)(unsafe.Add(mBase, _consts[561]))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+48)))
	if v11&int32(2) == int32(0) {
		v143 = v11
		goto L3
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v24 = int32(0)
	v27 = *(*int32)(unsafe.Add(mBase, _consts[561]))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+48)))
	if base.B2i32(v23 != v24)^base.B2i32(v28&int32(2) == v24) != 0 {
		v143 = v28
		goto L3
	} else {
		goto L12
	}
L9:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v21 = F_LWLockAcquire(m, v17+int32(4096), int32(0))
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
	v35 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v39 = F_LWLockAcquire(m, v35+int32(4096), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	if v23 != 0 {
		v139 = int32(3)
		goto L4
	} else {
		goto L14
	}
L14:
	;
	goto L5
L15:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	if v77 == int32(0) {
		v107 = v75
		goto L22
	} else {
		goto L23
	}
L16:
	;
	if v47 == v46 {
		v75 = v46
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v51 = v47
	goto L18
L18:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v56
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v58
	*(*int64)(unsafe.Add(mBase, uint32(v51))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v51-int32(4)))) = int32(2)
	F_SetLatch(m, v51-int32(120))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L10
	} else {
		goto L20
	}
L19:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _consts[561]))
	v75 = v72
	goto L15
L20:
	;
	if v46 != v56 {
		v51 = v56
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v109 = int32(1)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v107)+20))
	if v110 == int32(0) {
		v139 = v109
		goto L4
	} else {
		goto L29
	}
L23:
	;
	v81 = v75 + int32(8)
	if v77 == v81 {
		v107 = v75
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v83 = v77
	goto L25
L25:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v87)+4)) = v88
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v90
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v83-int32(4)))) = int32(2)
	F_SetLatch(m, v83-int32(120))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L10
	} else {
		goto L27
	}
L26:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _consts[561]))
	v107 = v104
	goto L22
L27:
	;
	if v81 != v88 {
		v83 = v88
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v114 = v107 + int32(16)
	if v110 == v114 {
		v139 = v109
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v116 = v110
	goto L31
L31:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v120)+4)) = v121
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	*(*int32)(unsafe.Add(mBase, uint32(v121))) = v123
	*(*int64)(unsafe.Add(mBase, uint32(v116))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v116-int32(4)))) = int32(2)
	F_SetLatch(m, v116-int32(120))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L10
	} else {
		goto L33
	}
L32:
	;
	v139 = v109
	goto L4
L33:
	;
	if v114 != v121 {
		v116 = v121
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v152 = F_LWLockAcquire(m, v148+int32(4096), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L10
	} else {
		goto L36
	}
L36:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _consts[561]))
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+48)))
	v158 = v156 | int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v155)+48)) = uint8(v158)
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
	var v91 int32
	_ = v91
	var v97 int64
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v142 int64
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	v1 = l0
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[432]))
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
	v15 = *(*int32)(unsafe.Add(mBase, _consts[560]))
	if v15 < int32(2) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[561]))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+48)))
	if v20&int32(3) == int32(1) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _consts[562]))
	v28 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v32 = F_LWLockAcquire(m, v28+int32(4096), int32(0))
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
	v40 = *(*int32)(unsafe.Add(mBase, _consts[561]))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+48)))
	if v41&int32(1) != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _consts[293]))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+136)) = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v79)+128)) = v1
	v85 = v40 + v38<<(uint(int32(3))%32)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v86 == int32(0) {
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
	v52 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v52+int32(4096))
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
	v63 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v63+int32(4096))
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
	v69 = *(*int32)(unsafe.Add(mBase, _consts[563]))
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
	v72 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v72+int32(4096))
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
	v131 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v131+int32(4096))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L5
	} else {
		goto L44
	}
L32:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	if v114 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L33:
	;
	if v86 == v85 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v91 = v86
	goto L35
L35:
	;
	v97 = *(*int64)(unsafe.Add(mBase, uint32(v91-int32(12))))
	if base.Ui64(v97) < base.Ui64(v1) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L32
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79)+140)) = v91
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+144)) = v100
	v103 = v79 + int32(140)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+4)) = v103
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v79)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = v103
	goto L31
L38:
	;
	goto L39
L39:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if v107 != v85 {
		v91 = v107
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
	v118 = v85
	goto L43
L42:
	;
	v118 = v114
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79)+140)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v79)+144)) = v118
	v122 = v79 + int32(140)
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = v122
	goto L31
L44:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, _consts[219])))
	if v137 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+4)) = uint32(v1)
	v142 = int64(base.Ui64(v1) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v8))) = uint32(v142)
	v147 = F_pg_sprintf(m, v8+int32(16), int32(535146), v8)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
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
	v300 = *(*int32)(unsafe.Add(mBase, _consts[293]))
	*(*int64)(unsafe.Add(mBase, uint32(v300)+128)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v300)+136)) = int32(0)
	goto L1
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v287)+136)) = int32(0)
	v292 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v292+int32(4096))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L5
	} else {
		goto L84
	}
L51:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	*(*int32)(unsafe.Add(mBase, uint32(v155))) = int32(0)
	goto L53
L52:
	;
	*(*int32)(unsafe.Add(mBase, _consts[564])) = int32(1)
	v265 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[279])) = v265
	v268 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v272 = F_LWLockAcquire(m, v268+int32(4096), v265)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L5
	} else {
		goto L82
	}
L53:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _consts[293]))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)+136))
	if v160 == int32(2) {
		goto L49
	} else {
		goto L54
	}
L54:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _consts[564]))
	if v164 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v167 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L5
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v209 = *(*int32)(unsafe.Add(mBase, _consts[413]))
	if v209 != 0 {
		goto L68
	} else {
		goto L69
	}
L58:
	;
	if v167 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L5
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v186 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[279])) = v186
	v189 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v193 = F_LWLockAcquire(m, v189+int32(4096), v186)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L5
	} else {
		goto L66
	}
L62:
	;
	F_errmsg(m, int32(445069), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L5
	} else {
		goto L63
	}
L63:
	;
	F_errdetail(m, int32(598084), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L5
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(515551), int32(305), int32(548953))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	goto L61
L66:
	;
	v196 = *(*int32)(unsafe.Add(mBase, _consts[293]))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)+144))
	if v197 == int32(0) {
		v287 = v196
		goto L50
	} else {
		goto L67
	}
L67:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v196)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v197
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v196)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v197))) = v202
	*(*int64)(unsafe.Add(mBase, uint32(v196)+140)) = int64(0)
	v207 = *(*int32)(unsafe.Add(mBase, _consts[293]))
	v287 = v207
	goto L50
L68:
	;
	v211 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[413])) = v211
	v215 = F_errstart(m, int32(19), v211)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L5
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	v255 = F_WaitLatch(m, v251, int32(17), int32(-1), int32(134217780))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L5
	} else {
		goto L80
	}
L71:
	;
	if v215 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	F_errmsg(m, int32(82136), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L5
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v231 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v235 = F_LWLockAcquire(m, v231+int32(4096), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L5
	} else {
		goto L78
	}
L75:
	;
	F_errdetail(m, int32(598084), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L5
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(515551), int32(322), int32(548953))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L5
	} else {
		goto L77
	}
L77:
	;
	goto L74
L78:
	;
	v238 = *(*int32)(unsafe.Add(mBase, _consts[293]))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)+144))
	if v239 == int32(0) {
		v287 = v238
		goto L50
	} else {
		goto L79
	}
L79:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v238)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v242)+4)) = v239
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v238)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v239))) = v244
	*(*int64)(unsafe.Add(mBase, uint32(v238)+140)) = int64(0)
	v249 = *(*int32)(unsafe.Add(mBase, _consts[293]))
	v287 = v249
	goto L50
L80:
	;
	if v255&int32(16) == int32(0) {
		goto L51
	} else {
		goto L81
	}
L81:
	;
	goto L52
L82:
	;
	v275 = *(*int32)(unsafe.Add(mBase, _consts[293]))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)+144))
	if v276 == int32(0) {
		v287 = v275
		goto L50
	} else {
		goto L83
	}
L83:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v275)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v279)+4)) = v276
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v275)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v276))) = v281
	*(*int64)(unsafe.Add(mBase, uint32(v275)+140)) = int64(0)
	v286 = *(*int32)(unsafe.Add(mBase, _consts[293]))
	v287 = v286
	goto L50
L84:
	;
	goto L49
}
