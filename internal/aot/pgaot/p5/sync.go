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
	v24 = F_WaitLatch(m, int32(0), int32(40), int32(10), int32(150994950))
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v29 int64
	_ = v29
	var v41 int64
	_ = v41
	var v51 int64
	_ = v51
	var v69 int32
	_ = v69
	var v70 int64
	_ = v70
	var v73 int64
	_ = v73
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v110 int64
	_ = v110
	var v112 int64
	_ = v112
	var v124 int64
	_ = v124
	var v126 int32
	_ = v126
	var v129 int64
	_ = v129
	var v137 int64
	_ = v137
	var v143 int32
	_ = v143
	var v144 int64
	_ = v144
	var v150 int64
	_ = v150
	var v151 int64
	_ = v151
	var v153 int32
	_ = v153
	var v154 int64
	_ = v154
	var v160 int64
	_ = v160
	var v169 int64
	_ = v169
	var v176 int64
	_ = v176
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int64
	_ = v235
	var v237 int64
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int64
	_ = v273
	var v275 int64
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_SyncOneBuffer[0]))
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_SyncOneBuffer[1]))
	F_ResourceOwnerEnlarge(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = v15 + l0*int32(56)
	v27 = int64(4194304)
	v29 = base.AtomicRmwOr64(m, v26, int32(24), v27)
	if v29&v27 != int64(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v41 = v29
	goto L7
L5:
	;
	v124 = v29
	goto L6
L6:
	;
	v126 = int32(0)
	v129 = v124 & int64(4194303)
	if base.B2i32(l1 == v126)|base.B2i32(v129 == int64(0)) == v126 {
		goto L29
	} else {
		goto L30
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = int32(_a_F_SyncOneBuffer_0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(_a_F_SyncOneBuffer_1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(_a_F_SyncOneBuffer_2)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(0)
	v51 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v51
	if v41&int64(4194304) != v51 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v124 = v112
	goto L6
L9:
	;
	goto L12
L10:
	;
	goto L11
L11:
	;
	v90 = int32(_a_F_SyncOneBuffer_3)
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_SyncOneBuffer[2]))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(8))+8))
	if v93 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L12:
	;
	F_perform_spin_delay(m, v12+int32(8))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	goto L11
L14:
	;
	v70 = int64(0)
	v73 = base.AtomicRmwCmpxchg64(m, v26, int32(24), v70, v70)
	if v73&int64(4194304) != v70 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v110 = int64(4194304)
	v112 = base.AtomicRmwOr64(m, v26, int32(24), v110)
	if v112&v110 != int64(0) {
		v41 = v112
		goto L7
	} else {
		goto L27
	}
L17:
	;
	goto L16
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SyncOneBuffer[2])) = v108
	goto L17
L19:
	;
	if int32(999) < v91 {
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v91 < int32(11) {
		goto L17
	} else {
		goto L26
	}
L22:
	;
	v98 = int32(900)
	if v98 <= v91 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v101 = v98
	goto L25
L24:
	;
	v101 = v91
	goto L25
L25:
	;
	v108 = v101 + int32(100)
	goto L18
L26:
	;
	v108 = v91 - int32(1)
	goto L18
L27:
	;
	goto L8
L28:
	;
	m.G0 = v12 + int32(32)
	return v290
L29:
	;
	v137 = base.AtomicRmwSub64(m, v26, int32(24), int64(4194304))
	v290 = int32(0)
	goto L28
L30:
	;
	goto L31
L31:
	;
	if v129 == int64(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v143 = int32(2)
	goto L34
L33:
	;
	v143 = int32(0)
	goto L34
L34:
	;
	v144 = int64(25165824)
	if v124&v144 != v144 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v150 = base.AtomicRmwSub64(m, v26, int32(24), int64(4194304))
	v290 = v143
	goto L28
L36:
	;
	goto L37
L37:
	;
	v151 = int64(0)
	v153 = int32(24)
	v154 = base.AtomicRmwCmpxchg64(m, v26, v153, v151, v151)
	v160 = base.AtomicRmwCmpxchg64(m, v26, v153, v154, v154&int64(-4194305)+int64(1))
	if v154 != v160 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v169 = v160
	goto L41
L39:
	;
	goto L40
L40:
	;
	v187 = int32(_a_F_SyncOneBuffer_4)
	v188 = *(*int32)(unsafe.Add(mBase, _c_F_SyncOneBuffer[3]))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v194 = int32(1)
	v195 = v193 + v194
	*(*int32)(unsafe.Add(mBase, uint32(v188<<(uint(int32(2))%32))+uint32(_c_F_SyncOneBuffer[4]))) = v195
	v198 = v188 << (uint(int32(4)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v198)+uint32(_c_F_SyncOneBuffer[5]))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v198)+uint32(_c_F_SyncOneBuffer[6]))) = v195
	*(*int32)(unsafe.Add(mBase, _c_F_SyncOneBuffer[3])) = int32(-1)
	*(*int32)(unsafe.Add(mBase, _c_F_SyncOneBuffer[7])) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v198)+uint32(_c_F_SyncOneBuffer[8]))) = v194
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_SyncOneBuffer[1]))
	F_ResourceOwnerRemember(m, v216, base.I64_extend_i32_s(v195), int32(_a_F_SyncOneBuffer_5))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L44
	}
L41:
	;
	v176 = base.AtomicRmwCmpxchg64(m, v26, int32(24), v169, v169&int64(-4194305)+int64(1))
	if v169 != v176 {
		v169 = v176
		goto L41
	} else {
		goto L43
	}
L42:
	;
	goto L40
L43:
	;
	goto L42
L44:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v223 = v221 + int32(1)
	F_BufferLockAcquire(m, v223, v26, int32(2))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_FlushBuffer(m, v26, int32(0), int32(3))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_BufferLockUnlock(m, v223, v26)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v233
	v235 = *(*int64)(unsafe.Add(mBase, uint32(v26)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v235
	v237 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v237
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_SyncOneBuffer[1]))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	F_ResourceOwnerForget(m, v240, base.I64_extend_i32_s(v241+int32(1)), int32(_a_F_SyncOneBuffer_5))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_UnpinBufferNoOwner(m, v26)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SyncOneBuffer[9])))
	if v251&int32(1) != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v290 = v143 | int32(1)
	goto L28
L51:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SyncOneBuffer[10])))
	if v255&int32(1) == int32(0) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	if int32(0) < v261 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v264 + int32(1)
	v270 = l2 + v264*int32(20)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v270)+24)) = v271
	v273 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v270)+16)) = v273
	v275 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v270)+8)) = v275
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	v280 = v278
	goto L55
L54:
	;
	v280 = v261
	goto L55
L55:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v281 < v280 {
		goto L50
	} else {
		goto L56
	}
L56:
	;
	F_IssuePendingWritebacks(m, l2, int32(3))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	goto L50
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
	v74 = v54 - int32(280)
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
	v159 = v139 - int32(280)
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
	v245 = v225 - int32(280)
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
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
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
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	v1 = l0
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[0]))
	if v11 <= int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 - int32(-64)
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
	*(*int32)(unsafe.Add(mBase, uint32(v79)+592)) = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v79)+584)) = v1
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
	*(*int32)(unsafe.Add(mBase, uint32(v79)+596)) = v94
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+600)) = v103
	v106 = v79 + int32(596)
	*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v106
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v79)+600))
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
	*(*int32)(unsafe.Add(mBase, uint32(v79)+596)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v79)+600)) = v121
	v125 = v79 + int32(596)
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
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+20)) = uint32(v1)
	v145 = int64(base.Ui64(v1) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+16)) = uint32(v145)
	v152 = F_pg_sprintf(m, v6+int32(-32), int32(_a_F_SyncRepWaitForLSN_1), v6+int32(-48))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
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
	v336 = int32(0)
	v339 = base.AtomicRmwOr32(m, v336, int32(_a_F_SyncRepWaitForLSN_2), v336)
	v341 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[5]))
	*(*int64)(unsafe.Add(mBase, uint32(v341)+584)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v341)+592)) = v336
	goto L1
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v324)+592)) = int32(0)
	v329 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[4]))
	F_LWLockRelease(m, v329+int32(_a_F_SyncRepWaitForLSN_0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L5
	} else {
		goto L92
	}
L51:
	;
	v160 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[8]))
	v161 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v161
	v166 = base.AtomicRmwOr32(m, v161, int32(_a_F_SyncRepWaitForLSN_3), v161)
	goto L53
L52:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[9])) = int32(1)
	v302 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[10])) = v302
	v305 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[4]))
	v309 = F_LWLockAcquire(m, v305+int32(_a_F_SyncRepWaitForLSN_0), v302)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L5
	} else {
		goto L90
	}
L53:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[5]))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+592))
	if v169 == int32(2) {
		goto L49
	} else {
		goto L54
	}
L54:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[9]))
	if v173 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[11]))
	v178 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L5
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v246 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[12]))
	if v246 != 0 {
		goto L76
	} else {
		goto L77
	}
L58:
	;
	if v175 != 0 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v223 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[10])) = v223
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[4]))
	v230 = F_LWLockAcquire(m, v226+int32(_a_F_SyncRepWaitForLSN_0), v223)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L5
	} else {
		goto L74
	}
L60:
	;
	F_errfinish(m, int32(_a_F_SyncRepWaitForLSN_4), v218, int32(_a_F_SyncRepWaitForLSN_5))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L5
	} else {
		goto L73
	}
L61:
	;
	if v178 == int32(0) {
		goto L59
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if v178 == int32(0) {
		goto L59
	} else {
		goto L69
	}
L64:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	F_errmsg(m, int32(_a_F_SyncRepWaitForLSN_6), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L5
	} else {
		goto L66
	}
L66:
	;
	v192 = F_errdetail(m, int32(_a_F_SyncRepWaitForLSN_7), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L5
	} else {
		goto L67
	}
L67:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[11]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v195
	v198 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[13]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v198
	F_errdetail_log(m, int32(_a_F_SyncRepWaitForLSN_8), v8)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L5
	} else {
		goto L68
	}
L68:
	;
	v218 = int32(310)
	goto L60
L69:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L5
	} else {
		goto L70
	}
L70:
	;
	F_errmsg(m, int32(_a_F_SyncRepWaitForLSN_6), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L5
	} else {
		goto L71
	}
L71:
	;
	v215 = F_errdetail(m, int32(_a_F_SyncRepWaitForLSN_7), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L5
	} else {
		goto L72
	}
L72:
	;
	v218 = int32(315)
	goto L60
L73:
	;
	goto L59
L74:
	;
	v233 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[5]))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)+600))
	if v234 == int32(0) {
		v324 = v233
		goto L50
	} else {
		goto L75
	}
L75:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v233)+596))
	*(*int32)(unsafe.Add(mBase, uint32(v237)+4)) = v234
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v233)+596))
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = v239
	*(*int64)(unsafe.Add(mBase, uint32(v233)+596)) = int64(0)
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[5]))
	v324 = v244
	goto L50
L76:
	;
	v248 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[12])) = v248
	v252 = F_errstart(m, int32(19), v248)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L5
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v288 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[8]))
	v292 = F_WaitLatch(m, v288, int32(17), int32(-1), int32(134217781))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L5
	} else {
		goto L88
	}
L79:
	;
	if v252 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	F_errmsg(m, int32(_a_F_SyncRepWaitForLSN_9), int32(0))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L5
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[4]))
	v272 = F_LWLockAcquire(m, v268+int32(_a_F_SyncRepWaitForLSN_0), int32(0))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L5
	} else {
		goto L86
	}
L83:
	;
	v260 = F_errdetail(m, int32(_a_F_SyncRepWaitForLSN_7), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L5
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_SyncRepWaitForLSN_4), int32(332), int32(_a_F_SyncRepWaitForLSN_5))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	goto L82
L86:
	;
	v275 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[5]))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)+600))
	if v276 == int32(0) {
		v324 = v275
		goto L50
	} else {
		goto L87
	}
L87:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v275)+596))
	*(*int32)(unsafe.Add(mBase, uint32(v279)+4)) = v276
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v275)+596))
	*(*int32)(unsafe.Add(mBase, uint32(v276))) = v281
	*(*int64)(unsafe.Add(mBase, uint32(v275)+596)) = int64(0)
	v286 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[5]))
	v324 = v286
	goto L50
L88:
	;
	if v292&int32(16) == int32(0) {
		goto L51
	} else {
		goto L89
	}
L89:
	;
	goto L52
L90:
	;
	v312 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[5]))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)+600))
	if v313 == int32(0) {
		v324 = v312
		goto L50
	} else {
		goto L91
	}
L91:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v312)+596))
	*(*int32)(unsafe.Add(mBase, uint32(v316)+4)) = v313
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v312)+596))
	*(*int32)(unsafe.Add(mBase, uint32(v313))) = v318
	*(*int64)(unsafe.Add(mBase, uint32(v312)+596)) = int64(0)
	v323 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepWaitForLSN[5]))
	v324 = v323
	goto L50
L92:
	;
	goto L49
}
