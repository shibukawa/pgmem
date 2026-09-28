package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_LockSharedObject(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	v4 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(264)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+14)) = uint16(v9)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v4
	v19 = F_LockAcquire(m, v7, l2, v4, v4)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		F_ReceiveSharedInvalidMessages(m)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			m.G0 = v7 + int32(16)
			return
		}
	}
}
func F_ReceiveSharedInvalidMessages(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int64
	_ = v38
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v76 int32
	_ = v76
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
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int64
	_ = v146
	var v148 int64
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int64
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v243 int64
	_ = v243
	var v246 int64
	_ = v246
	var v247 int32
	_ = v247
	var v249 int64
	_ = v249
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v272 int32
	_ = v272
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[0]))
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[1]))
	if v15 < v17 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	goto L4
L2:
	;
	goto L3
L3:
	;
	goto L9
L4:
	;
	v28 = int32(_a_F_ReceiveSharedInvalidMessages_0)
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[0])) = v30 + int32(1)
	v35 = v30 << (uint(int32(4)) % 32)
	v38 = *(*int64)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_ReceiveSharedInvalidMessages[2])))
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_ReceiveSharedInvalidMessages[3])))
	v42 = int32(_a_F_ReceiveSharedInvalidMessages_1)
	v44 = *(*int64)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[4]))
	*(*int64)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[4])) = v44 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v38
	F_LocalExecuteInvalidationMessage(m, v12)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[0]))
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[1]))
	if v53 < v55 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	v76 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[1])) = v76
	*(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[0])) = v76
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[5]))
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[6]))
	v88 = v83 + v85<<(uint(int32(4))%32)
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+uint32(_c_F_ReceiveSharedInvalidMessages[7]))))
	if v91 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v285 = *(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[8]))
	if v285 != 0 {
		goto L52
	} else {
		goto L53
	}
L11:
	;
	goto L10
L12:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[9]))
	v99 = F_LWLockAcquire(m, v95+int32(640), int32(1))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L6
	} else {
		goto L15
	}
L13:
	;
	v184 = v76
	goto L14
L14:
	;
	if v184 < int32(0) {
		goto L34
	} else {
		goto L35
	}
L15:
	;
	v103 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v88)+uint32(_c_F_ReceiveSharedInvalidMessages[7]))) = uint8(v103)
	v107 = base.AtomicRmwXchg32(m, v83, int32(12), int32(1))
	if v107 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	F_s_lock(m, v83+int32(12), int32(_a_F_ReceiveSharedInvalidMessages_2))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L6
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v114 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v83)+12)), uint32(v114))
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+uint32(_c_F_ReceiveSharedInvalidMessages[10]))))
	if v117 == v114 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	goto L18
L20:
	;
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[9]))
	F_LWLockRelease(m, v177+int32(640))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L6
	} else {
		goto L33
	}
L21:
	;
	v129 = v76
	v131 = v120
	goto L26
L22:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v88)+uint32(_c_F_ReceiveSharedInvalidMessages[11])))
	goto L21
L23:
	;
	goto L24
L24:
	;
	v121 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v88)+uint32(_c_F_ReceiveSharedInvalidMessages[10]))) = uint16(v121)
	*(*int32)(unsafe.Add(mBase, uint32(v88)+uint32(_c_F_ReceiveSharedInvalidMessages[11]))) = v113
	v169 = int32(-1)
	goto L20
L25:
	;
	if v113 <= v160 {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	if v113 <= v131 {
		v159 = v129
		v160 = v131
		goto L25
	} else {
		goto L28
	}
L27:
	;
	v159 = int32(32)
	v160 = v152
	goto L25
L28:
	;
	v137 = int32(4)
	v138 = v129 << (uint(v137) % 32)
	v142 = base.I32_rem_s(v131, int32(_a_F_ReceiveSharedInvalidMessages_3))
	v145 = v83 + int32(16) + v142<<(uint(v137)%32)
	v146 = *(*int64)(unsafe.Add(mBase, uint32(v145)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v138)+uint32(_c_F_ReceiveSharedInvalidMessages[2]))) = v146
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v145)))
	*(*int64)(unsafe.Add(mBase, uint32(v138)+uint32(_c_F_ReceiveSharedInvalidMessages[3]))) = v148
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v88)+uint32(_c_F_ReceiveSharedInvalidMessages[11])))
	v151 = int32(1)
	v152 = v150 + v151
	*(*int32)(unsafe.Add(mBase, uint32(v88)+uint32(_c_F_ReceiveSharedInvalidMessages[11]))) = v152
	v155 = v129 + v151
	if v155 != int32(32) {
		v129 = v155
		v131 = v152
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v163 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v88)+uint32(_c_F_ReceiveSharedInvalidMessages[12]))) = uint8(v163)
	v169 = v159
	goto L20
L31:
	;
	goto L32
L32:
	;
	v165 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v88)+uint32(_c_F_ReceiveSharedInvalidMessages[7]))) = uint8(v165)
	v169 = v159
	goto L20
L33:
	;
	v184 = v169
	goto L14
L34:
	;
	v195 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L6
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v214 = int32(_a_F_ReceiveSharedInvalidMessages_0)
	*(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[0])) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[1])) = v184
	v220 = *(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[0]))
	if v220 < v184 {
		goto L44
	} else {
		goto L45
	}
L37:
	;
	if v195 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_errmsg_internal(m, int32(_a_F_ReceiveSharedInvalidMessages_4), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L6
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v206 = int32(_a_F_ReceiveSharedInvalidMessages_1)
	v208 = *(*int64)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[4]))
	*(*int64)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[4])) = v208 + int64(1)
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L6
	} else {
		goto L43
	}
L41:
	;
	F_errfinish(m, int32(_a_F_ReceiveSharedInvalidMessages_5), int32(103), int32(_a_F_ReceiveSharedInvalidMessages_6))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	goto L11
L44:
	;
	goto L47
L45:
	;
	goto L46
L46:
	;
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[1]))
	if v272 == int32(32) {
		goto L9
	} else {
		goto L51
	}
L47:
	;
	v233 = int32(_a_F_ReceiveSharedInvalidMessages_0)
	v235 = *(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[0])) = v235 + int32(1)
	v240 = v235 << (uint(int32(4)) % 32)
	v243 = *(*int64)(unsafe.Add(mBase, uint32(v240)+uint32(_c_F_ReceiveSharedInvalidMessages[2])))
	v246 = *(*int64)(unsafe.Add(mBase, uint32(v240)+uint32(_c_F_ReceiveSharedInvalidMessages[3])))
	v247 = int32(_a_F_ReceiveSharedInvalidMessages_1)
	v249 = *(*int64)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[4]))
	*(*int64)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[4])) = v249 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v246
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v243
	F_LocalExecuteInvalidationMessage(m, v12)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L6
	} else {
		goto L49
	}
L48:
	;
	goto L46
L49:
	;
	v258 = *(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[0]))
	v260 = *(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[1]))
	if v258 < v260 {
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	goto L11
L52:
	;
	v287 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[8])) = v287
	v291 = F_errstart(m, int32(11), v287)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L6
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	m.G0 = v12 + int32(16)
	return
L55:
	;
	if v291 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	F_errmsg_internal(m, int32(_a_F_ReceiveSharedInvalidMessages_7), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L6
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v302 = int32(0)
	F_SICleanupQueue(m, v302, v302)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L6
	} else {
		goto L61
	}
L59:
	;
	F_errfinish(m, int32(_a_F_ReceiveSharedInvalidMessages_5), int32(138), int32(_a_F_ReceiveSharedInvalidMessages_6))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L6
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	goto L54
}
func F_SharedFileSetInit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	v3 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0)+44)), uint32(v3))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
	F_FileSetInit(m, l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		if l1 != 0 {
			F_on_dsm_detach(m, l1, int32(1185), base.I64_extend_i32_u(l0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				return
			}
		} else {
			return
		}
	}
}
func F_SharedFileSetOnDetach(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	v4 = base.I32_wrap_i64(l1)
	v7 = base.AtomicRmwXchg32(m, v4, int32(44), int32(1))
	if v7 != 0 {
		F_s_lock(m, v4+int32(44), int32(_a_F_SharedFileSetOnDetach_0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v4)+48))
			v15 = v13 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+48)) = v15
			v17 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v4)+44)), uint32(v17))
			if v15 == v17 {
				F_FileSetDeleteAll(m, v4)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					return
				}
			} else {
				return
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v4)+48))
		v15 = v13 - int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+48)) = v15
		v17 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v4)+44)), uint32(v17))
		if v15 == v17 {
			F_FileSetDeleteAll(m, v4)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				return
			}
		} else {
			return
		}
	}
}
func F_shared_buffer_readv_complete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int64
	_ = v157
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int64
	_ = v172
	var v173 int32
	_ = v173
	var v188 int64
	_ = v188
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int64
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v305 int64
	_ = v305
	var v309 int32
	_ = v309
	v5 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(32)
	m.G0 = v29
	v31 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v31
	v35 = base.I32_wrap_i64(v31) & int32(448)
	v37 = l1 + int32(104)
	goto L1
L1:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v29+int32(23)))) = uint8(v40)
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_shared_buffer_readv_complete[0]))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	goto L2
L2:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+23)))
	if v49 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v35 == int32(256) {
		goto L53
	} else {
		goto L54
	}
L4:
	;
	v52 = int32(0)
	v243 = v52
	v247 = v5
	v248 = v5
	v250 = v5
	v252 = v5
	v253 = v5
	v266 = v52
	goto L3
L5:
	;
	goto L6
L6:
	;
	v62 = l3 & int32(1)
	if v62 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v63 = int32(10)
	goto L9
L8:
	;
	v63 = int32(2)
	goto L9
L9:
	;
	v69 = int32(0)
	v75 = v5
	v76 = v5
	v78 = v5
	v79 = v5
	v80 = v5
	v81 = v5
	v85 = v5
	goto L10
L10:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_shared_buffer_readv_complete[1]))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v44+v45<<(uint(int32(3))%32)+v69<<(uint(int32(3))%32))))
	v104 = v97 + v101*int32(56)
	if v101 < int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v232 = int32(255)
	v243 = v227 & v232 << (uint(int32(18)) % 32)
	v247 = v222
	v248 = v225
	v250 = v212
	v252 = v216
	v253 = v220
	v266 = base.B2i32(v221&v232 != int32(0))
	goto L3
L12:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v104-int32(40))))
	v128 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)) = uint8(v128)
	if base.B2i32(v35 == int32(256))|base.B2i32(base.I32_wrap_i64(int64(base.Ui64(v31)>>(uint(int64(32))%64))) <= v69) != 0 {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_shared_buffer_readv_complete[2]))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v110+(v101^int32(-1))<<(uint(int32(2))%32))))
	v124 = v116
	goto L12
L14:
	;
	goto L15
L15:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_shared_buffer_readv_complete[3]))
	v124 = v118 + v101<<(uint(int32(13))%32) + int32(-8192)
	goto L12
L16:
	;
	v204 = int32(0)
	F_TerminateBufferIO(m, v104-int32(56), v204, v203, v204, int32(1))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L20
	} else {
		goto L33
	}
L17:
	;
	v196 = int32(1)
	v197 = v128
	v198 = int32(0)
	v203 = int64(134217728)
	goto L16
L18:
	;
	goto L19
L19:
	;
	v136 = F_PageIsVerified(m, v124, v127, l3&int32(4)|v63, v29+int32(22))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return
L21:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
	if v136 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v29)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v169<<(uint(int32(18))%32) | v173 | (v171 | (v168<<(uint(int32(9))%32) | v170)) | v69<<(uint(int32(25))%32)
	v188 = *(*int64)(unsafe.Add(mBase, uint32(v29)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = v188
	F_pgaio_result_report(m, v29+int32(8), v37, int32(16))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L20
	} else {
		goto L32
	}
L23:
	;
	v141 = int32(2048)
	if v62 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v156 = int32(1)
	v157 = int64(16777216)
	if v138&v156 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v145 = int32(0)
	v167 = v145
	v168 = v128
	v169 = v138
	v170 = v145
	v171 = v141
	v172 = int64(134217728)
	v173 = int32(258)
	goto L22
L27:
	;
	goto L28
L28:
	;
	v148 = int32(0)
	base.MemoryFill(m, v124, v148, int32(_a_F_shared_buffer_readv_complete_0))
	v153 = int32(1)
	v167 = v153
	v168 = v153
	v169 = v138
	v170 = v148
	v171 = v141
	v172 = int64(16777216)
	v173 = int32(194)
	goto L22
L29:
	;
	v196 = v156
	v197 = v128
	v198 = int32(0)
	v203 = v157
	goto L16
L30:
	;
	goto L31
L31:
	;
	v167 = v156
	v168 = v128
	v169 = int32(1)
	v170 = int32(1024)
	v171 = int32(0)
	v172 = v157
	v173 = int32(194)
	goto L22
L32:
	;
	v196 = v167 | v168
	v197 = v168
	v198 = v136
	v203 = v172
	goto L16
L33:
	;
	if v79&int32(255) != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v211 = v78
	goto L36
L35:
	;
	v211 = v69
	goto L36
L36:
	;
	if v198 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v212 = v211
	goto L39
L38:
	;
	v212 = v78
	goto L39
L39:
	;
	if v75&int32(255) != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v215 = v80
	goto L42
L41:
	;
	v215 = v69
	goto L42
L42:
	;
	if v197 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v216 = v215
	goto L45
L44:
	;
	v216 = v80
	goto L45
L45:
	;
	if v76&int32(255) != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v219 = v81
	goto L48
L47:
	;
	v219 = v69
	goto L48
L48:
	;
	if v196 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v220 = v81
	goto L51
L50:
	;
	v220 = v219
	goto L51
L51:
	;
	v221 = v198 + v79
	v222 = v197 + v75
	v223 = int32(1)
	v225 = v76 + (v196 ^ v223)
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
	v227 = v226 + v85
	v229 = v69 + v223
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+23)))
	if base.Ui32(v229) < base.Ui32(v230) {
		v69 = v229
		v75 = v222
		v76 = v225
		v78 = v212
		v79 = v221
		v80 = v216
		v81 = v220
		v85 = v227
		goto L10
	} else {
		goto L52
	}
L52:
	;
	goto L11
L53:
	;
	m.G0 = v29 + int32(32)
	return
L54:
	;
	v269 = int32(255)
	v270 = v247 & v269
	v272 = v248 & v269
	v273 = int32(0)
	if v270|(base.B2i32(v272 != v273)|v266)&int32(1) == v273 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	if v272 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v283 = int32(258)
	goto L58
L57:
	;
	v283 = int32(194)
	goto L58
L58:
	;
	if v270 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v287 = int32(512)
	goto L61
L60:
	;
	v287 = int32(0)
	goto L61
L61:
	;
	if v266 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v290 = int32(1024)
	goto L64
L63:
	;
	v290 = int32(0)
	goto L64
L64:
	;
	if v272 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v292 = v248
	goto L67
L66:
	;
	v292 = v247
	goto L67
L67:
	;
	if v270 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v299 = v252
	goto L70
L69:
	;
	v299 = v250
	goto L70
L70:
	;
	if v272 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v300 = v253
	goto L73
L72:
	;
	v300 = v299
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v243 | v283 | (v287 | v290 | v292&int32(255)<<(uint(int32(11))%32)) | v300<<(uint(int32(25))%32)
	v305 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = v305
	F_pgaio_result_report(m, v29, v37, int32(14))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L20
	} else {
		goto L74
	}
L74:
	;
	goto L53
}
func F_shared_buffer_write_error_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	if l0 != 0 {
		F_set_errcontext_domain(m, int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v14 = v7 + int32(8)
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			F_GetRelationPath(m, v14, v15, v16, v17, int32(-1), v19)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v12
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v14
				F_errcontext_msg(m, int32(_a_F_shared_buffer_write_error_callback_0), v7)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					m.G0 = v7 + int32(80)
					return
				}
			}
		}
	} else {
		m.G0 = v7 + int32(80)
		return
	}
}
