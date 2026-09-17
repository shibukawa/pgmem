package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecFetchSlotHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
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
	var v23 int32
	_ = v23
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
	if l1 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+28))
		m.T0[v6].(func(*base.Module, int32))(m, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+36))
			if v12 == int32(0) {
				if l2 != 0 {
					v15 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v15)
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v18 = v17
				} else {
					v18 = v11
				}
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
				v20 = m.T0[v19].(func(*base.Module, int32) int32)(m, l0)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					return v20
				}
			} else {
				if l2 != 0 {
					v23 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v23)
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
					v27 = v26
				} else {
					v27 = v12
				}
				v28 = m.T0[v27].(func(*base.Module, int32) int32)(m, l0)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					return v28
				}
			}
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+36))
		if v12 == int32(0) {
			if l2 != 0 {
				v15 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v15)
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v18 = v17
			} else {
				v18 = v11
			}
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
			v20 = m.T0[v19].(func(*base.Module, int32) int32)(m, l0)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				return v20
			}
		} else {
			if l2 != 0 {
				v23 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v23)
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
				v27 = v26
			} else {
				v27 = v12
			}
			v28 = m.T0[v27].(func(*base.Module, int32) int32)(m, l0)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				return v28
			}
		}
	}
}
func F_SaveSlotToPath(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v316 int64
	_ = v316
	var v318 int64
	_ = v318
	var v323 int32
	_ = v323
	v9 = m.G0
	v11 = v9 - int32(2352)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_s_lock(m, l0, int32(_a_F_SaveSlotToPath_0), int32(2328), int32(_a_F_SaveSlotToPath_1))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v21 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)) = uint8(v21)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v21
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	if v25 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	m.G0 = v11 + int32(2352)
	return
L7:
	;
	v29 = l0 + int32(208)
	v31 = F_LWLockAcquire(m, v29, int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	base.MemoryFill(m, v11+int32(120), int32(0), int32(184))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = l1
	v40 = v11 + int32(1328)
	v44 = F_pg_sprintf(m, v40, int32(_a_F_SaveSlotToPath_2), v11+int32(96))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = l1
	v52 = F_pg_sprintf(m, v11+int32(304), int32(_a_F_SaveSlotToPath_3), v11+int32(80))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v55 = F_OpenTransientFile(m, v40, int32(193))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	if v55 < int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[0]))
	F_LWLockRelease(m, v29)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+112)) = int64(790273982469)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+104)) = int64(-4277855071)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	if v85 != 0 {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[0])) = v60
	v66 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	if v66 == int32(0) {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v40
	F_errmsg(m, int32(_a_F_SaveSlotToPath_4), v11)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_SaveSlotToPath_0), int32(2361), int32(_a_F_SaveSlotToPath_1))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	goto L6
L21:
	;
	F_s_lock(m, l0, int32(_a_F_SaveSlotToPath_0), int32(2370), int32(_a_F_SaveSlotToPath_1))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	base.MemoryCopy(m, v11+int32(120), l0+int32(24), int32(184))
	v99 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	v105 = m.Env.Pgmem_crc32c(m, v101, v11+int32(112), int32(192))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v11)+108)) = v105 ^ int32(-1)
	*(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[0])) = v99
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = int32(167772209)
	v118 = int32(200)
	v119 = F_write(m, v55, v11+int32(104), v118)
	mBase = m.M
	if v119 != v118 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L23
L25:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[0]))
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v125))) = int32(0)
	v128 = F_CloseTransientFile(m, v55)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L4
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v157 = int32(_a_F_SaveSlotToPath_5)
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[1]))
	v159 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v158))) = v159
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v162))) = int32(167772208)
	v167 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SaveSlotToPath[2])))
	if v167 != int32(1) {
		v181 = v159
		goto L39
	} else {
		goto L40
	}
L28:
	;
	v131 = v11 + int32(1328)
	v132 = F_unlink(m, v131)
	mBase = m.M
	F_LWLockRelease(m, v29)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	if v123 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v137 = v123
	goto L32
L31:
	;
	v137 = int32(51)
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[0])) = v137
	v140 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	if v140 == int32(0) {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v131
	F_errmsg(m, int32(_a_F_SaveSlotToPath_6), v11-int32(-64))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_SaveSlotToPath_0), int32(2397), int32(_a_F_SaveSlotToPath_1))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	goto L6
L38:
	;
	if v181 != 0 {
		goto L45
	} else {
		goto L46
	}
L39:
	;
	goto L38
L40:
	;
	goto L41
L41:
	;
	v172 = F_fsync(m, v55)
	mBase = m.M
	if v172 != int32(-1) {
		v181 = v172
		goto L39
	} else {
		goto L43
	}
L42:
	;
	v181 = int32(-1)
	goto L39
L43:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[0]))
	if v176 == int32(27) {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[0]))
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v185))) = int32(0)
	v188 = F_CloseTransientFile(m, v55)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L4
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v216))) = int32(0)
	v219 = F_CloseTransientFile(m, v55)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L4
	} else {
		goto L55
	}
L48:
	;
	v191 = v11 + int32(1328)
	v192 = F_unlink(m, v191)
	mBase = m.M
	F_LWLockRelease(m, v29)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[0])) = v183
	v198 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	if v198 == int32(0) {
		goto L6
	} else {
		goto L51
	}
L51:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v191
	F_errmsg(m, int32(_a_F_SaveSlotToPath_7), v11+int32(48))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_SaveSlotToPath_0), int32(2417), int32(_a_F_SaveSlotToPath_1))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	goto L6
L55:
	;
	if v219 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[0]))
	v224 = v11 + int32(1328)
	v225 = F_unlink(m, v224)
	mBase = m.M
	F_LWLockRelease(m, v29)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L4
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v249 = v11 + int32(1328)
	v251 = v11 + int32(304)
	v252 = F_rename(m, v249, v251)
	mBase = m.M
	if v252 != 0 {
		goto L65
	} else {
		goto L66
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[0])) = v222
	v231 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	if v231 == int32(0) {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v224
	F_errmsg(m, int32(_a_F_SaveSlotToPath_8), v11+int32(32))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_SaveSlotToPath_0), int32(2433), int32(_a_F_SaveSlotToPath_1))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	goto L6
L65:
	;
	v254 = *(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[0]))
	v255 = F_unlink(m, v249)
	mBase = m.M
	F_LWLockRelease(m, v29)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L4
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v279 = int32(_a_F_SaveSlotToPath_9)
	v281 = *(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[3])) = v281 + int32(1)
	F_fsync_fname(m, v11+int32(304), int32(0))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L4
	} else {
		goto L74
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[0])) = v254
	v261 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	if v261 == int32(0) {
		goto L6
	} else {
		goto L70
	}
L70:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v249
	F_errmsg(m, int32(_a_F_SaveSlotToPath_10), v11+int32(16))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_SaveSlotToPath_0), int32(2449), int32(_a_F_SaveSlotToPath_1))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	goto L6
L74:
	;
	F_fsync_fname(m, l1, int32(1))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L4
	} else {
		goto L75
	}
L75:
	;
	F_fsync_fname(m, int32(_a_F_SaveSlotToPath_11), int32(1))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L4
	} else {
		goto L76
	}
L76:
	;
	v297 = int32(_a_F_SaveSlotToPath_9)
	v299 = *(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[3]))
	v300 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[3])) = v299 - v300
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v300
	if v303 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	F_s_lock(m, l0, int32(_a_F_SaveSlotToPath_0), int32(2468), int32(_a_F_SaveSlotToPath_1))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L4
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v311 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	goto L79
L81:
	;
	v314 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v314)
	goto L83
L82:
	;
	goto L83
L83:
	;
	v316 = *(*int64)(unsafe.Add(mBase, uint32(v11)+216))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+264)) = v316
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v11)+200))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+280)) = v318
	F_LWLockRelease(m, v29)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	goto L6
}
func F_slot_getmissingattrs(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	if v8 == int32(0) {
		v80 = l2 - l1
		v82 = v80 << (uint(int32(2)) % 32)
		if v82 != 0 {
			v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			base.MemoryFill(m, v83+l1<<(uint(int32(2))%32), int32(0), v82)
		} else {
		}
		if v80 == int32(0) {
		} else {
			v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			base.MemoryFill(m, v91+l1, int32(1), v80)
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
		if v11 == int32(0) {
			v80 = l2 - l1
			v82 = v80 << (uint(int32(2)) % 32)
			if v82 != 0 {
				v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				base.MemoryFill(m, v83+l1<<(uint(int32(2))%32), int32(0), v82)
			} else {
			}
			if v80 == int32(0) {
			} else {
				v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				base.MemoryFill(m, v91+l1, int32(1), v80)
			}
		} else {
			if l2 <= l1 {
			} else {
				v15 = int32(1)
				v16 = l1 + v15
				if (l2-l1)&v15 != 0 {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v26 = v11 + l1<<(uint(int32(3))%32)
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v20+l1<<(uint(int32(2))%32)))) = v27
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
					v33 = v31 ^ int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v29+l1))) = uint8(v33)
					v35 = v16
				} else {
					v35 = l1
				}
				if l2 == v16 {
				} else {
					v39 = v35
					for {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v45 = int32(2)
						v48 = int32(3)
						v50 = v11 + v39<<(uint(v48)%32)
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v44+v39<<(uint(v45)%32)))) = v51
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
						v56 = int32(1)
						v57 = v55 ^ v56
						*(*uint8)(unsafe.Add(mBase, uint32(v53+v39))) = uint8(v57)
						v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v61 = v39 + v56
						v67 = v11 + v61<<(uint(v48)%32)
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v59+v61<<(uint(v45)%32)))) = v68
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
						v74 = v72 ^ v56
						*(*uint8)(unsafe.Add(mBase, uint32(v70+v61))) = uint8(v74)
						v77 = v39 + v45
						if v77 != l2 {
							v39 = v77
							continue
						} else {
							break
						}
						break
					}
				}
			}
		}
	}
	return
}
func F_slot_getsomeattrs_int(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	v2 = l1
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v2 <= v10 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
		m.T0[v13].(func(*base.Module, int32, int32))(m, l0, v2)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v16 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
			if v16 < v2 {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
				if v22 == int32(0) {
					v94 = v2 - v16
					v96 = v94 << (uint(int32(2)) % 32)
					if v96 != 0 {
						v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						base.MemoryFill(m, v97+v16<<(uint(int32(2))%32), int32(0), v96)
					} else {
					}
					if v94 == int32(0) {
					} else {
						v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						base.MemoryFill(m, v105+v16, int32(1), v94)
					}
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
					if v25 == int32(0) {
						v94 = v2 - v16
						v96 = v94 << (uint(int32(2)) % 32)
						if v96 != 0 {
							v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							base.MemoryFill(m, v97+v16<<(uint(int32(2))%32), int32(0), v96)
						} else {
						}
						if v94 == int32(0) {
						} else {
							v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							base.MemoryFill(m, v105+v16, int32(1), v94)
						}
					} else {
						if v2 <= v16 {
						} else {
							v29 = int32(1)
							v30 = v16 + v29
							if (v2-v16)&v29 != 0 {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v40 = v25 + v16<<(uint(int32(3))%32)
								v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v34+v16<<(uint(int32(2))%32)))) = v41
								v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
								v47 = v45 ^ int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v43+v16))) = uint8(v47)
								v49 = v30
							} else {
								v49 = v16
							}
							if v2 == v30 {
							} else {
								v53 = v49
								for {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v59 = int32(2)
									v62 = int32(3)
									v64 = v25 + v53<<(uint(v62)%32)
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v58+v53<<(uint(v59)%32)))) = v65
									v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
									v70 = int32(1)
									v71 = v69 ^ v70
									*(*uint8)(unsafe.Add(mBase, uint32(v67+v53))) = uint8(v71)
									v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v75 = v53 + v70
									v81 = v25 + v75<<(uint(v62)%32)
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v73+v75<<(uint(v59)%32)))) = v82
									v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
									v88 = v86 ^ v70
									*(*uint8)(unsafe.Add(mBase, uint32(v84+v75))) = uint8(v88)
									v91 = v53 + v59
									if v91 != v2 {
										v53 = v91
										continue
									} else {
										break
									}
									break
								}
							}
						}
					}
				}
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v2)
			} else {
			}
			m.G0 = v7 + int32(16)
			return
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v122 = m.ExcPending
		if v122 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v2
			F_errmsg_internal(m, int32(_a_F_slot_getsomeattrs_int_0), v7)
			mBase = m.M
			v126 = m.ExcPending
			if v126 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_slot_getsomeattrs_int_1), int32(2098), int32(_a_F_slot_getsomeattrs_int_2))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
