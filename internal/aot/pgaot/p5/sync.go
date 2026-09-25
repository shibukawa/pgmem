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
	var v57 int32
	_ = v57
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
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int64
	_ = v153
	var v155 int64
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int64
	_ = v190
	var v192 int64
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
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
	v57 = v36
	goto L6
L6:
	;
	v61 = v57 & int32(_a_F_SyncOneBuffer_4)
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
	v57 = v51
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
	return v207
L23:
	;
	v92 = int32(0)
	v96 = base.AtomicRmwOr32(m, v92, int32(_a_F_SyncOneBuffer_6), v92)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = v57 & int32(-4194305)
	v207 = v92
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
	v102 = int32(0)
	goto L28
L27:
	;
	v102 = int32(2)
	goto L28
L28:
	;
	v103 = int32(25165824)
	if v57&v103 != v103 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v107 = int32(0)
	v110 = base.AtomicRmwOr32(m, v107, int32(_a_F_SyncOneBuffer_6), v107)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = v57 & int32(-4194305)
	v207 = v102
	goto L22
L30:
	;
	goto L31
L31:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	v115 = int32(0)
	v118 = base.AtomicRmwOr32(m, v115, int32(_a_F_SyncOneBuffer_6), v115)
	v119 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = (v114 + v119) & int32(-4194305)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	v125 = int32(_a_F_SyncOneBuffer_7)
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_SyncOneBuffer[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_SyncOneBuffer[3])) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v126)+4)) = v119
	v133 = v124 + v119
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = v133
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_SyncOneBuffer[1]))
	F_ResourceOwnerRemember(m, v136, v133, int32(_a_F_SyncOneBuffer_8))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v141 = v33 + int32(48)
	v143 = F_LWLockAcquire(m, v141, int32(1))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_FlushBuffer(m, v33, int32(0), int32(3))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_LWLockRelease(m, v141)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v151
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v33)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v153
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v33)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v155
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_SyncOneBuffer[1]))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	F_ResourceOwnerForget(m, v158, v159+int32(1), int32(_a_F_SyncOneBuffer_8))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_UnpinBufferNoOwner(m, v33)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SyncOneBuffer[4])))
	if v168&int32(1) != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v207 = v102 | int32(1)
	goto L22
L39:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SyncOneBuffer[5])))
	if v172&int32(1) == int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	if int32(0) < v178 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v181 + int32(1)
	v187 = l2 + v181*int32(20)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v187)+24)) = v188
	v190 = *(*int64)(unsafe.Add(mBase, uint32(v9)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v187)+16)) = v190
	v192 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v187)+8)) = v192
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	v197 = v195
	goto L43
L42:
	;
	v197 = v178
	goto L43
L43:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v198 < v197 {
		goto L38
	} else {
		goto L44
	}
L44:
	;
	F_IssuePendingWritebacks(m, l2, int32(3))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
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
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
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
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
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
	v327 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[1]))
	F_LWLockRelease(m, v327+int32(_a_F_SyncRepUpdateSyncStandbysDefined_0))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L10
	} else {
		goto L76
	}
L3:
	;
	if v305&int32(1) != 0 {
		goto L1
	} else {
		goto L74
	}
L4:
	;
	v303 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[2]))
	*(*uint8)(unsafe.Add(mBase, uint32(v303)+48)) = uint8(v301)
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
		v305 = v11
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
		v305 = v25
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
		v301 = int32(3)
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
	v131 = v46
	goto L17
L17:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v131)+12))
	if v133 == int32(0) {
		v216 = v131
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
	v65 = int32(0)
	v68 = base.AtomicRmwOr32(m, v65, int32(_a_F_SyncRepUpdateSyncStandbysDefined_1), v65)
	*(*int32)(unsafe.Add(mBase, uint32(v54-int32(4)))) = int32(2)
	v74 = v54 - int32(120)
	v78 = base.AtomicRmwOr32(m, v65, int32(_a_F_SyncRepUpdateSyncStandbysDefined_2), v65)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	if v79 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[2]))
	v131 = v128
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
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(1)
	v82 = int32(0)
	v85 = base.AtomicRmwOr32(m, v82, int32(_a_F_SyncRepUpdateSyncStandbysDefined_2), v82)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if v86 == v82 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	if v89 == int32(0) {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[3]))
	if v93 == v89 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v95 = m.G0
	v97 = v95 - int32(16)
	m.G0 = v97
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[4]))
	if v100 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v123 = F_pgmem_kill(m, v89, int32(23))
	mBase = m.M
	goto L21
L28:
	;
	m.G0 = v97 + int32(16)
	goto L20
L29:
	;
	v103 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v97)+15)) = uint8(v103)
	goto L30
L30:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[5]))
	v111 = F_write(m, v107, v97+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v111 {
		goto L28
	} else {
		goto L32
	}
L31:
	;
	goto L28
L32:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[6]))
	if v115 == int32(27) {
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
	v218 = int32(1)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v216)+20))
	if v219 == int32(0) {
		v301 = v218
		goto L4
	} else {
		goto L55
	}
L36:
	;
	v137 = v131 + int32(8)
	if v133 == v137 {
		v216 = v131
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v139 = v133
	goto L38
L38:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v143)+4)) = v144
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	*(*int32)(unsafe.Add(mBase, uint32(v144))) = v146
	*(*int64)(unsafe.Add(mBase, uint32(v139))) = int64(0)
	v150 = int32(0)
	v153 = base.AtomicRmwOr32(m, v150, int32(_a_F_SyncRepUpdateSyncStandbysDefined_1), v150)
	*(*int32)(unsafe.Add(mBase, uint32(v139-int32(4)))) = int32(2)
	v159 = v139 - int32(120)
	v163 = base.AtomicRmwOr32(m, v150, int32(_a_F_SyncRepUpdateSyncStandbysDefined_2), v150)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	if v164 != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v213 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[2]))
	v216 = v213
	goto L35
L40:
	;
	if v137 != v144 {
		v139 = v144
		goto L38
	} else {
		goto L54
	}
L41:
	;
	goto L40
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = int32(1)
	v167 = int32(0)
	v170 = base.AtomicRmwOr32(m, v167, int32(_a_F_SyncRepUpdateSyncStandbysDefined_2), v167)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	if v171 == v167 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v159)+12))
	if v174 == int32(0) {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[3]))
	if v178 == v174 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v180 = m.G0
	v182 = v180 - int32(16)
	m.G0 = v182
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[4]))
	if v185 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	v208 = F_pgmem_kill(m, v174, int32(23))
	mBase = m.M
	goto L41
L48:
	;
	m.G0 = v182 + int32(16)
	goto L40
L49:
	;
	v188 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v182)+15)) = uint8(v188)
	goto L50
L50:
	;
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[5]))
	v196 = F_write(m, v192, v182+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v196 {
		goto L48
	} else {
		goto L52
	}
L51:
	;
	goto L48
L52:
	;
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[6]))
	if v200 == int32(27) {
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
	v223 = v216 + int32(16)
	if v219 == v223 {
		v301 = v218
		goto L4
	} else {
		goto L56
	}
L56:
	;
	v225 = v219
	goto L57
L57:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v225)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v229)+4)) = v230
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	*(*int32)(unsafe.Add(mBase, uint32(v230))) = v232
	*(*int64)(unsafe.Add(mBase, uint32(v225))) = int64(0)
	v236 = int32(0)
	v239 = base.AtomicRmwOr32(m, v236, int32(_a_F_SyncRepUpdateSyncStandbysDefined_1), v236)
	*(*int32)(unsafe.Add(mBase, uint32(v225-int32(4)))) = int32(2)
	v245 = v225 - int32(120)
	v249 = base.AtomicRmwOr32(m, v236, int32(_a_F_SyncRepUpdateSyncStandbysDefined_2), v236)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	if v250 != 0 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v301 = v218
	goto L4
L59:
	;
	if v223 != v230 {
		v225 = v230
		goto L57
	} else {
		goto L73
	}
L60:
	;
	goto L59
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v245))) = int32(1)
	v253 = int32(0)
	v256 = base.AtomicRmwOr32(m, v253, int32(_a_F_SyncRepUpdateSyncStandbysDefined_2), v253)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v245)+4))
	if v257 == v253 {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v245)+12))
	if v260 == int32(0) {
		goto L60
	} else {
		goto L63
	}
L63:
	;
	v264 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[3]))
	if v264 == v260 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v266 = m.G0
	v268 = v266 - int32(16)
	m.G0 = v268
	v271 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[4]))
	if v271 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	v294 = F_pgmem_kill(m, v260, int32(23))
	mBase = m.M
	goto L60
L67:
	;
	m.G0 = v268 + int32(16)
	goto L59
L68:
	;
	v274 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v268)+15)) = uint8(v274)
	goto L69
L69:
	;
	v278 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[5]))
	v282 = F_write(m, v278, v268+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v282 {
		goto L67
	} else {
		goto L71
	}
L70:
	;
	goto L67
L71:
	;
	v286 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[6]))
	if v286 == int32(27) {
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
	v310 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[1]))
	v314 = F_LWLockAcquire(m, v310+int32(_a_F_SyncRepUpdateSyncStandbysDefined_0), int32(0))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L10
	} else {
		goto L75
	}
L75:
	;
	v317 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepUpdateSyncStandbysDefined[2]))
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317)+48)))
	v320 = v318 | int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v317)+48)) = uint8(v320)
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
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
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
	v306 = int32(0)
	v309 = base.AtomicRmwOr32(m, v306, int32(_a_F_SyncRepWaitForLSN_2), v306)
	v311 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[5]))
	*(*int64)(unsafe.Add(mBase, uint32(v311)+128)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v311)+136)) = v306
	goto L1
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v294)+136)) = int32(0)
	v299 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[4]))
	F_LWLockRelease(m, v299+int32(_a_F_SyncRepWaitForLSN_0))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L5
	} else {
		goto L84
	}
L51:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[8]))
	v159 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v158))) = v159
	v164 = base.AtomicRmwOr32(m, v159, int32(_a_F_SyncRepWaitForLSN_3), v159)
	goto L53
L52:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[9])) = int32(1)
	v272 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[10])) = v272
	v275 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[4]))
	v279 = F_LWLockAcquire(m, v275+int32(_a_F_SyncRepWaitForLSN_0), v272)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L5
	} else {
		goto L82
	}
L53:
	;
	v166 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[5]))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+136))
	if v167 == int32(2) {
		goto L49
	} else {
		goto L54
	}
L54:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[9]))
	if v171 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v174 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L5
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[11]))
	if v216 != 0 {
		goto L68
	} else {
		goto L69
	}
L58:
	;
	if v174 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L5
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v193 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[10])) = v193
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[4]))
	v200 = F_LWLockAcquire(m, v196+int32(_a_F_SyncRepWaitForLSN_0), v193)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L5
	} else {
		goto L66
	}
L62:
	;
	F_errmsg(m, int32(_a_F_SyncRepWaitForLSN_4), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L5
	} else {
		goto L63
	}
L63:
	;
	F_errdetail(m, int32(_a_F_SyncRepWaitForLSN_5), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L5
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_SyncRepWaitForLSN_6), int32(305), int32(_a_F_SyncRepWaitForLSN_7))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	goto L61
L66:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[5]))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+144))
	if v204 == int32(0) {
		v294 = v203
		goto L50
	} else {
		goto L67
	}
L67:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v203)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v207)+4)) = v204
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v203)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v204))) = v209
	*(*int64)(unsafe.Add(mBase, uint32(v203)+140)) = int64(0)
	v214 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[5]))
	v294 = v214
	goto L50
L68:
	;
	v218 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[11])) = v218
	v222 = F_errstart(m, int32(19), v218)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L5
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v258 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[8]))
	v262 = F_WaitLatch(m, v258, int32(17), int32(-1), int32(134217780))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L5
	} else {
		goto L80
	}
L71:
	;
	if v222 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	F_errmsg(m, int32(_a_F_SyncRepWaitForLSN_8), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L5
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v238 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[4]))
	v242 = F_LWLockAcquire(m, v238+int32(_a_F_SyncRepWaitForLSN_0), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L5
	} else {
		goto L78
	}
L75:
	;
	F_errdetail(m, int32(_a_F_SyncRepWaitForLSN_5), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L5
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_SyncRepWaitForLSN_6), int32(322), int32(_a_F_SyncRepWaitForLSN_7))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L5
	} else {
		goto L77
	}
L77:
	;
	goto L74
L78:
	;
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[5]))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)+144))
	if v246 == int32(0) {
		v294 = v245
		goto L50
	} else {
		goto L79
	}
L79:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v245)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v249)+4)) = v246
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v245)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v246))) = v251
	*(*int64)(unsafe.Add(mBase, uint32(v245)+140)) = int64(0)
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[5]))
	v294 = v256
	goto L50
L80:
	;
	if v262&int32(16) == int32(0) {
		goto L51
	} else {
		goto L81
	}
L81:
	;
	goto L52
L82:
	;
	v282 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[5]))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v282)+144))
	if v283 == int32(0) {
		v294 = v282
		goto L50
	} else {
		goto L83
	}
L83:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v282)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v286)+4)) = v283
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v282)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v283))) = v288
	*(*int64)(unsafe.Add(mBase, uint32(v282)+140)) = int64(0)
	v293 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[5]))
	v294 = v293
	goto L50
L84:
	;
	goto L49
}
