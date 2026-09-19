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
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int64
	_ = v148
	var v150 int64
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int64
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v245 int64
	_ = v245
	var v248 int64
	_ = v248
	var v249 int32
	_ = v249
	var v251 int64
	_ = v251
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v274 int32
	_ = v274
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
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
	v287 = *(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[8]))
	if v287 != 0 {
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
	v186 = v76
	goto L14
L14:
	;
	if v186 < int32(0) {
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
	F_s_lock(m, v83+int32(12), int32(_a_F_ReceiveSharedInvalidMessages_2), int32(511), int32(_a_F_ReceiveSharedInvalidMessages_3))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L6
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v116 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v83)+12)), uint32(v116))
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+uint32(_c_F_ReceiveSharedInvalidMessages[10]))))
	if v119 == v116 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	goto L18
L20:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[9]))
	F_LWLockRelease(m, v179+int32(640))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L6
	} else {
		goto L33
	}
L21:
	;
	v131 = v76
	v133 = v122
	goto L26
L22:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v88)+uint32(_c_F_ReceiveSharedInvalidMessages[11])))
	goto L21
L23:
	;
	goto L24
L24:
	;
	v123 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v88)+uint32(_c_F_ReceiveSharedInvalidMessages[10]))) = uint16(v123)
	*(*int32)(unsafe.Add(mBase, uint32(v88)+uint32(_c_F_ReceiveSharedInvalidMessages[11]))) = v115
	v171 = int32(-1)
	goto L20
L25:
	;
	if v115 <= v162 {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	if v115 <= v133 {
		v161 = v131
		v162 = v133
		goto L25
	} else {
		goto L28
	}
L27:
	;
	v161 = int32(32)
	v162 = v154
	goto L25
L28:
	;
	v139 = int32(4)
	v140 = v131 << (uint(v139) % 32)
	v144 = base.I32_rem_s(v133, int32(_a_F_ReceiveSharedInvalidMessages_4))
	v147 = v83 + int32(16) + v144<<(uint(v139)%32)
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v147)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v140)+uint32(_c_F_ReceiveSharedInvalidMessages[2]))) = v148
	v150 = *(*int64)(unsafe.Add(mBase, uint32(v147)))
	*(*int64)(unsafe.Add(mBase, uint32(v140)+uint32(_c_F_ReceiveSharedInvalidMessages[3]))) = v150
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v88)+uint32(_c_F_ReceiveSharedInvalidMessages[11])))
	v153 = int32(1)
	v154 = v152 + v153
	*(*int32)(unsafe.Add(mBase, uint32(v88)+uint32(_c_F_ReceiveSharedInvalidMessages[11]))) = v154
	v157 = v131 + v153
	if v157 != int32(32) {
		v131 = v157
		v133 = v154
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v165 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v88)+uint32(_c_F_ReceiveSharedInvalidMessages[12]))) = uint8(v165)
	v171 = v161
	goto L20
L31:
	;
	goto L32
L32:
	;
	v167 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v88)+uint32(_c_F_ReceiveSharedInvalidMessages[7]))) = uint8(v167)
	v171 = v161
	goto L20
L33:
	;
	v186 = v171
	goto L14
L34:
	;
	v197 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L6
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v216 = int32(_a_F_ReceiveSharedInvalidMessages_0)
	*(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[0])) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[1])) = v186
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[0]))
	if v222 < v186 {
		goto L44
	} else {
		goto L45
	}
L37:
	;
	if v197 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_errmsg_internal(m, int32(_a_F_ReceiveSharedInvalidMessages_5), int32(0))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L6
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v208 = int32(_a_F_ReceiveSharedInvalidMessages_1)
	v210 = *(*int64)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[4]))
	*(*int64)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[4])) = v210 + int64(1)
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L6
	} else {
		goto L43
	}
L41:
	;
	F_errfinish(m, int32(_a_F_ReceiveSharedInvalidMessages_6), int32(103), int32(_a_F_ReceiveSharedInvalidMessages_7))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
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
	v274 = *(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[1]))
	if v274 == int32(32) {
		goto L9
	} else {
		goto L51
	}
L47:
	;
	v235 = int32(_a_F_ReceiveSharedInvalidMessages_0)
	v237 = *(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[0])) = v237 + int32(1)
	v242 = v237 << (uint(int32(4)) % 32)
	v245 = *(*int64)(unsafe.Add(mBase, uint32(v242)+uint32(_c_F_ReceiveSharedInvalidMessages[2])))
	v248 = *(*int64)(unsafe.Add(mBase, uint32(v242)+uint32(_c_F_ReceiveSharedInvalidMessages[3])))
	v249 = int32(_a_F_ReceiveSharedInvalidMessages_1)
	v251 = *(*int64)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[4]))
	*(*int64)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[4])) = v251 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v248
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v245
	F_LocalExecuteInvalidationMessage(m, v12)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L6
	} else {
		goto L49
	}
L48:
	;
	goto L46
L49:
	;
	v260 = *(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[0]))
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[1]))
	if v260 < v262 {
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
	v289 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ReceiveSharedInvalidMessages[8])) = v289
	v293 = F_errstart(m, int32(11), v289)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
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
	if v293 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	F_errmsg_internal(m, int32(_a_F_ReceiveSharedInvalidMessages_8), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L6
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v304 = int32(0)
	F_SICleanupQueue(m, v304, v304)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L6
	} else {
		goto L61
	}
L59:
	;
	F_errfinish(m, int32(_a_F_ReceiveSharedInvalidMessages_6), int32(138), int32(_a_F_ReceiveSharedInvalidMessages_7))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
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
	var v12 int32
	_ = v12
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
			F_on_dsm_detach(m, l1, int32(1096), l0)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				return
			}
		} else {
			return
		}
	}
}
func F_SharedFileSetOnDetach(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	v5 = base.AtomicRmwXchg32(m, l1, int32(44), int32(1))
	if v5 != 0 {
		F_s_lock(m, l1+int32(44), int32(_a_F_SharedFileSetOnDetach_0), int32(101), int32(_a_F_SharedFileSetOnDetach_1))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
			v15 = v13 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v15
			v17 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l1)+44)), uint32(v17))
			if v15 == v17 {
				F_FileSetDeleteAll(m, l1)
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
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
		v15 = v13 - int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v15
		v17 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l1)+44)), uint32(v17))
		if v15 == v17 {
			F_FileSetDeleteAll(m, l1)
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v157 int32
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
	var v172 int32
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
	var v199 int32
	_ = v199
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
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v306 int64
	_ = v306
	var v310 int32
	_ = v310
	v5 = int32(0)
	v28 = m.G0
	v30 = v28 - int32(32)
	m.G0 = v30
	v32 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v32
	v36 = base.I32_wrap_i64(v32) & int32(448)
	v38 = l1 + int32(104)
	goto L1
L1:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v30+int32(23)))) = uint8(v41)
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_shared_buffer_readv_complete[0]))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	goto L2
L2:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+23)))
	if v50 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v36 == int32(256) {
		goto L50
	} else {
		goto L51
	}
L4:
	;
	v53 = int32(0)
	v243 = v53
	v247 = v5
	v248 = v5
	v251 = v5
	v253 = v5
	v254 = v5
	v267 = v53
	goto L3
L5:
	;
	goto L6
L6:
	;
	v68 = int32(0)
	v74 = v5
	v75 = v5
	v78 = v5
	v79 = v5
	v80 = v5
	v81 = v5
	v85 = v5
	goto L7
L7:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_shared_buffer_readv_complete[1]))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v45+v46<<(uint(int32(3))%32)+v68<<(uint(int32(3))%32))))
	v104 = v97 + v101<<(uint(int32(6))%32)
	if v101 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v232 = int32(255)
	v243 = v227 & v232 << (uint(int32(18)) % 32)
	v247 = v222
	v248 = v225
	v251 = v212
	v253 = v216
	v254 = v220
	v267 = base.B2i32(v221&v232 != int32(0))
	goto L3
L9:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v104-int32(48))))
	v128 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+22)) = uint8(v128)
	if base.B2i32(v36 == int32(256))|base.B2i32(base.I32_wrap_i64(int64(base.Ui64(v32)>>(uint(int64(32))%64))) <= v68) != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_shared_buffer_readv_complete[2]))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v110+(v101^int32(-1))<<(uint(int32(2))%32))))
	v124 = v116
	goto L9
L11:
	;
	goto L12
L12:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_shared_buffer_readv_complete[3]))
	v124 = v118 + v101<<(uint(int32(13))%32) + int32(-8192)
	goto L9
L13:
	;
	v204 = int32(0)
	F_TerminateBufferIO(m, v104+int32(-64), v204, v198, v204, int32(1))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L17
	} else {
		goto L30
	}
L14:
	;
	v196 = int32(1)
	v197 = v128
	v198 = int32(134217728)
	v199 = int32(0)
	goto L13
L15:
	;
	goto L16
L16:
	;
	v136 = F_PageIsVerified(m, v124, v127, l3&int32(4)|int32(2), v30+int32(22))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return
L18:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+22)))
	if v136 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v30)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = v170<<(uint(int32(18))%32) | v173 | (v172 | (v168<<(uint(int32(9))%32) | v171)) | v68<<(uint(int32(25))%32)
	v188 = *(*int64)(unsafe.Add(mBase, uint32(v30)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+8)) = v188
	F_pgaio_result_report(m, v30+int32(8), v38, int32(16))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L17
	} else {
		goto L29
	}
L20:
	;
	v141 = int32(2048)
	if l3&int32(1) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v156 = int32(1)
	v157 = int32(16777216)
	if v138&v156 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v145 = int32(0)
	v167 = v145
	v168 = v128
	v169 = int32(134217728)
	v170 = v138
	v171 = v145
	v172 = v141
	v173 = int32(258)
	goto L19
L24:
	;
	goto L25
L25:
	;
	v148 = int32(0)
	base.MemoryFill(m, v124, v148, int32(_a_F_shared_buffer_readv_complete_0))
	v153 = int32(1)
	v167 = v153
	v168 = v153
	v169 = int32(16777216)
	v170 = v138
	v171 = v148
	v172 = v141
	v173 = int32(194)
	goto L19
L26:
	;
	v196 = v156
	v197 = v128
	v198 = v157
	v199 = int32(0)
	goto L13
L27:
	;
	goto L28
L28:
	;
	v167 = v156
	v168 = v128
	v169 = v157
	v170 = int32(1)
	v171 = int32(1024)
	v172 = int32(0)
	v173 = int32(194)
	goto L19
L29:
	;
	v196 = v167 | v168
	v197 = v168
	v198 = v169
	v199 = v136
	goto L13
L30:
	;
	if v79&int32(255) != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v211 = v78
	goto L33
L32:
	;
	v211 = v68
	goto L33
L33:
	;
	if v199 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v212 = v211
	goto L36
L35:
	;
	v212 = v78
	goto L36
L36:
	;
	if v74&int32(255) != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v215 = v80
	goto L39
L38:
	;
	v215 = v68
	goto L39
L39:
	;
	if v197 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v216 = v215
	goto L42
L41:
	;
	v216 = v80
	goto L42
L42:
	;
	if v75&int32(255) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v219 = v81
	goto L45
L44:
	;
	v219 = v68
	goto L45
L45:
	;
	if v196 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v220 = v81
	goto L48
L47:
	;
	v220 = v219
	goto L48
L48:
	;
	v221 = v199 + v79
	v222 = v197 + v74
	v223 = int32(1)
	v225 = v75 + (v196 ^ v223)
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+22)))
	v227 = v226 + v85
	v229 = v68 + v223
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+23)))
	if base.Ui32(v229) < base.Ui32(v230) {
		v68 = v229
		v74 = v222
		v75 = v225
		v78 = v212
		v79 = v221
		v80 = v216
		v81 = v220
		v85 = v227
		goto L7
	} else {
		goto L49
	}
L49:
	;
	goto L8
L50:
	;
	m.G0 = v30 + int32(32)
	return
L51:
	;
	v270 = int32(255)
	v271 = v247 & v270
	v273 = v248 & v270
	v274 = int32(0)
	if v271|(base.B2i32(v273 != v274)|v267)&int32(1) == v274 {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	if v273 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v284 = int32(258)
	goto L55
L54:
	;
	v284 = int32(194)
	goto L55
L55:
	;
	if v271 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v288 = int32(512)
	goto L58
L57:
	;
	v288 = int32(0)
	goto L58
L58:
	;
	if v267 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v291 = int32(1024)
	goto L61
L60:
	;
	v291 = int32(0)
	goto L61
L61:
	;
	if v273 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v293 = v248
	goto L64
L63:
	;
	v293 = v247
	goto L64
L64:
	;
	if v271 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v300 = v253
	goto L67
L66:
	;
	v300 = v251
	goto L67
L67:
	;
	if v273 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v301 = v254
	goto L70
L69:
	;
	v301 = v300
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v243 | v284 | (v288 | v291 | v293&int32(255)<<(uint(int32(11))%32)) | v301<<(uint(int32(25))%32)
	v306 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v30))) = v306
	F_pgaio_result_report(m, v30, v38, int32(14))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L17
	} else {
		goto L71
	}
L71:
	;
	goto L50
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
