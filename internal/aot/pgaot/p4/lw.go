package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LWLockAcquire(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockAcquire[0]))
	if v10 < int32(200) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = int32(_a_F_LWLockAcquire_0)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockAcquire[1]))
	v16 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_LWLockAcquire[1])) = v15 + v16
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockAcquire[2]))
	v27 = int32(0)
	v29 = v16
	goto L4
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L22
	} else {
		goto L49
	}
L4:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v33 = v30
	goto L6
L6:
	;
	if l1 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	if base.B2i32(v53 == int32(0)) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L8:
	;
	v58 = base.AtomicRmwCmpxchg32(m, l0, int32(4), v33, v54)
	if v33 != v58 {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	v44 = v33 & int32(_a_F_LWLockAcquire_1)
	if v44 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v47 = v33 & int32(_a_F_LWLockAcquire_2)
	v53 = v47
	v54 = base.B2i32(int32(base.Ui32(v47)>>(uint(int32(18))%32)) == int32(0)) + v33
	goto L8
L12:
	;
	v45 = v33
	goto L14
L13:
	;
	v45 = v33 | int32(_a_F_LWLockAcquire_2)
	goto L14
L14:
	;
	v53 = v44
	v54 = v45
	goto L8
L15:
	;
	v33 = v58
	goto L6
L16:
	;
	goto L17
L17:
	;
	goto L7
L18:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockAcquire[3]))
	v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	*(*int32)(unsafe.Add(mBase, uint32(v149))) = v150 | int32(16777216)
	v159 = v27
	goto L45
L19:
	;
	F_LWLockQueueSelf(m, l0, l1)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v108 = int32(_a_F_LWLockAcquire_3)
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockAcquire[0]))
	v111 = v109 << (uint(int32(3)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v111)+uint32(_c_F_LWLockAcquire[4]))) = l0
	*(*int32)(unsafe.Add(mBase, _c_F_LWLockAcquire[0])) = v109 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v111)+uint32(_c_F_LWLockAcquire[5]))) = l1
	if int32(0) < v27 {
		goto L38
	} else {
		goto L39
	}
L22:
	;
	return int32(0)
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v69 = v66
	goto L24
L24:
	;
	if l1 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	if base.B2i32(v89 == int32(0)) == int32(0) {
		goto L18
	} else {
		goto L36
	}
L26:
	;
	v94 = base.AtomicRmwCmpxchg32(m, l0, int32(4), v69, v90)
	if v69 != v94 {
		goto L33
	} else {
		goto L34
	}
L27:
	;
	v80 = v69 & int32(_a_F_LWLockAcquire_1)
	if v80 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v83 = v69 & int32(_a_F_LWLockAcquire_2)
	v89 = v83
	v90 = base.B2i32(int32(base.Ui32(v83)>>(uint(int32(18))%32)) == int32(0)) + v69
	goto L26
L30:
	;
	v81 = v69
	goto L32
L31:
	;
	v81 = v69 | int32(_a_F_LWLockAcquire_2)
	goto L32
L32:
	;
	v89 = v80
	v90 = v81
	goto L26
L33:
	;
	v69 = v94
	goto L24
L34:
	;
	goto L35
L35:
	;
	goto L25
L36:
	;
	F_LWLockDequeueSelf(m, l0)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L22
	} else {
		goto L37
	}
L37:
	;
	goto L21
L38:
	;
	v129 = v27
	goto L41
L39:
	;
	goto L40
L40:
	;
	return v29
L41:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	F_PGSemaphoreUnlock(m, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L22
	} else {
		goto L43
	}
L42:
	;
	goto L40
L43:
	;
	v135 = int32(1)
	if base.Ui32(v135) < base.Ui32(v129) {
		v129 = v129 - v135
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	F_PGSemaphoreLock(m, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L22
	} else {
		goto L47
	}
L46:
	;
	v170 = base.AtomicRmwOr32(m, l0, int32(4), int32(1073741824))
	v171 = int32(0)
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockAcquire[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v173))) = v171
	v27 = v159
	v29 = v171
	goto L4
L47:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+74)))
	if v167 != 0 {
		v159 = v159 + int32(1)
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	F_errmsg_internal(m, int32(_a_F_LWLockAcquire_4), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L22
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_LWLockAcquire_5), int32(1212), int32(_a_F_LWLockAcquire_6))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L22
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_LWLockConditionalAcquire(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockConditionalAcquire[0]))
	if v8 < int32(200) {
		v11 = int32(_a_F_LWLockConditionalAcquire_0)
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockConditionalAcquire[1]))
		*(*int32)(unsafe.Add(mBase, _c_F_LWLockConditionalAcquire[1])) = v13 + int32(1)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v20 = v17
		for {
			if l1 == int32(0) {
				v29 = v20 & int32(_a_F_LWLockConditionalAcquire_1)
				if v29 != 0 {
					v30 = v20
				} else {
					v30 = v20 | int32(_a_F_LWLockConditionalAcquire_2)
				}
				v38 = v29
				v39 = v30
			} else {
				v32 = v20 & int32(_a_F_LWLockConditionalAcquire_2)
				v38 = v32
				v39 = base.B2i32(int32(base.Ui32(v32)>>(uint(int32(18))%32)) == int32(0)) + v20
			}
			v41 = base.B2i32(v38 == int32(0))
			v43 = base.AtomicRmwCmpxchg32(m, l0, int32(4), v20, v39)
			if v20 != v43 {
				v20 = v43
				continue
			} else {
				break
			}
			break
		}
		if v41 == int32(0) {
			v47 = int32(_a_F_LWLockConditionalAcquire_0)
			v49 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockConditionalAcquire[1]))
			*(*int32)(unsafe.Add(mBase, _c_F_LWLockConditionalAcquire[1])) = v49 - int32(1)
			return v41
		} else {
			v54 = int32(_a_F_LWLockConditionalAcquire_3)
			v55 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockConditionalAcquire[0]))
			v57 = v55 << (uint(int32(3)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v57)+uint32(_c_F_LWLockConditionalAcquire[2]))) = l0
			*(*int32)(unsafe.Add(mBase, _c_F_LWLockConditionalAcquire[0])) = v55 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v57)+uint32(_c_F_LWLockConditionalAcquire[3]))) = l1
			return v41
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v74 = m.ExcPending
		if v74 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_LWLockConditionalAcquire_4), int32(0))
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_LWLockConditionalAcquire_5), int32(1361), int32(_a_F_LWLockConditionalAcquire_6))
				mBase = m.M
				v83 = m.ExcPending
				if v83 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_LWLockReleaseAll(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseAll[0]))
	if int32(0) < v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v6 = v3
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v7 = int32(_a_F_LWLockReleaseAll_0)
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseAll[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseAll[1])) = v9 + int32(1)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v6<<(uint(int32(3))%32))+uint32(_c_F_LWLockReleaseAll[2])))
	v18 = F_LWLockDisownInternal(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
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
	F_LWLockReleaseInternal(m, v17, v18)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v22 = int32(_a_F_LWLockReleaseAll_0)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseAll[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseAll[1])) = v24 - int32(1)
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseAll[0]))
	if int32(0) < v29 {
		v6 = v29
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
}
func F_LWLockReleaseInternal(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v320 int32
	_ = v320
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	if l1 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v16 + int32(32)
	return
L2:
	;
	v20 = int32(-1)
	goto L4
L3:
	;
	v20 = int32(-262144)
	goto L4
L4:
	;
	if l1 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v23 = int32(1)
	goto L7
L6:
	;
	v23 = int32(_a_F_LWLockReleaseInternal_0)
	goto L7
L7:
	;
	v25 = base.AtomicRmwSub32(m, l0, int32(4), v23)
	if (v20+v25)&int32(-1073217537) != int32(-1073741824) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v31 = int32(536870912)
	v33 = base.AtomicRmwOr32(m, l0, int32(4), v31)
	if v33&v31 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v37 = v33
	goto L12
L10:
	;
	goto L11
L11:
	;
	v135 = int32(-1)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v136 == v135 {
		goto L35
	} else {
		goto L36
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = int32(_a_F_LWLockReleaseInternal_1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = int32(888)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = int32(_a_F_LWLockReleaseInternal_2)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = int64(0)
	if v37&int32(536870912) != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L11
L14:
	;
	goto L17
L15:
	;
	goto L16
L16:
	;
	v97 = int32(_a_F_LWLockReleaseInternal_3)
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseInternal[0]))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(8))+8))
	if v100 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	F_perform_spin_delay(m, v16+int32(8))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L16
L19:
	;
	return
L20:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v78&int32(536870912) != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	v117 = int32(536870912)
	v119 = base.AtomicRmwOr32(m, l0, int32(4), v117)
	if v119&v117 != 0 {
		v37 = v119
		goto L12
	} else {
		goto L33
	}
L23:
	;
	goto L22
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseInternal[0])) = v115
	goto L23
L25:
	;
	if int32(999) < v98 {
		goto L23
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	if v98 < int32(11) {
		goto L23
	} else {
		goto L32
	}
L28:
	;
	v105 = int32(900)
	if v105 <= v98 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v108 = v105
	goto L31
L30:
	;
	v108 = v98
	goto L31
L31:
	;
	v115 = v108 + int32(100)
	goto L24
L32:
	;
	v115 = v98 - int32(1)
	goto L24
L33:
	;
	goto L13
L34:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v276 = base.AtomicRmwCmpxchg32(m, l0, int32(4), v270, (v270&int32(-1610612737)|v259)&v269)
	if v270 != v276 {
		goto L65
	} else {
		goto L66
	}
L35:
	;
	v257 = v135
	v259 = int32(1073741824)
	v269 = int32(1610612735)
	goto L34
L36:
	;
	goto L37
L37:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseInternal[1]))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v150 = v135
	v151 = v136
	v152 = v143 + v136*int32(640)
	v153 = v142
	v157 = int32(-1)
	v158 = v143
	v159 = int32(1)
	v160 = int32(0)
	goto L38
L38:
	;
	v163 = v151 * int32(640)
	v164 = v158 + v163
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v152)+76))
	if v160&int32(1) == int32(0) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	if v245 != 0 {
		goto L59
	} else {
		goto L60
	}
L40:
	;
	goto L39
L41:
	;
	if v165 == int32(-1) {
		v240 = v227
		v245 = v232
		goto L40
	} else {
		goto L58
	}
L42:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v164)+76))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v164)+80))
	if v173 == int32(-1) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+75)))
	if v170 != 0 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v227 = v150
	v229 = v153
	v231 = v157
	v232 = v159
	v233 = int32(1)
	goto L41
L45:
	;
	if v172 == int32(-1) {
		goto L50
	} else {
		goto L51
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v172
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v164)+80))
	v182 = v177
	goto L45
L47:
	;
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v158+v173*int32(640))+76)) = v172
	v182 = v173
	goto L45
L49:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v164)+76)) = int64(0)
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseInternal[1]))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	v198 = v197 + v163
	if v157 == int32(-1) {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v182
	goto L49
L51:
	;
	goto L52
L52:
	;
	v187 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseInternal[1]))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	*(*int32)(unsafe.Add(mBase, uint32(v188+v172*int32(640))+80)) = v182
	goto L49
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v198)+76)) = int32(-1)
	v214 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v164)+74)) = uint8(v214)
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+75)))
	v219 = base.B2i32(v216 == v214) & v159
	if v216 == int32(0) {
		v240 = v211
		v245 = v219
		goto L40
	} else {
		goto L57
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v198)+80)) = int32(-1)
	v211 = v151
	goto L53
L55:
	;
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v198)+80)) = v157
	v205 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseInternal[1]))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	*(*int32)(unsafe.Add(mBase, uint32(v206+v157*int32(640))+76)) = v151
	v211 = v150
	goto L53
L57:
	;
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseInternal[1]))
	v227 = v211
	v229 = v223
	v231 = v151
	v232 = v219
	v233 = v160 | base.B2i32(v216 != int32(2))
	goto L41
L58:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	v150 = v227
	v151 = v165
	v152 = v236 + v165*int32(640)
	v153 = v229
	v157 = v231
	v158 = v236
	v159 = v232
	v160 = v233
	goto L38
L59:
	;
	v249 = int32(1073741824)
	goto L61
L60:
	;
	v249 = int32(0)
	goto L61
L61:
	;
	v251 = int32(-1)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v252 == v251 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v255 = int32(1610612735)
	goto L64
L63:
	;
	v255 = v251
	goto L64
L64:
	;
	v257 = v240
	v259 = v249
	v269 = v255
	goto L34
L65:
	;
	v280 = v276
	goto L68
L66:
	;
	goto L67
L67:
	;
	if v257 == int32(-1) {
		goto L1
	} else {
		goto L74
	}
L68:
	;
	v295 = int32(-1)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v296 == v295 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L67
L70:
	;
	v299 = int32(1610612735)
	goto L72
L71:
	;
	v299 = v295
	goto L72
L72:
	;
	v302 = base.AtomicRmwCmpxchg32(m, l0, int32(4), v280, (v280&int32(-1610612737)|v259)&v299)
	if v280 != v302 {
		v280 = v302
		goto L68
	} else {
		goto L73
	}
L73:
	;
	goto L69
L74:
	;
	v320 = v257
	goto L75
L75:
	;
	v333 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseInternal[1]))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)))
	v337 = v334 + v320*int32(640)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v337)+76))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v337)+80))
	if v339 != int32(-1) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v337)+76)) = int64(0)
	v364 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v337)+74)) = uint8(v364)
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v337)+12))
	F_PGSemaphoreUnlock(m, v366)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L19
	} else {
		goto L84
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v334+v339*int32(640))+76)) = v338
	goto L79
L78:
	;
	goto L79
L79:
	;
	if v338 != int32(-1) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v349 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseInternal[1]))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v349)))
	*(*int32)(unsafe.Add(mBase, uint32(v350+v338*int32(640))+80)) = v339
	*(*int64)(unsafe.Add(mBase, uint32(v337)+76)) = int64(0)
	v357 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v337)+74)) = uint8(v357)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v337)+12))
	F_PGSemaphoreUnlock(m, v359)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L19
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	goto L76
L83:
	;
	v320 = v338
	goto L75
L84:
	;
	goto L1
}
