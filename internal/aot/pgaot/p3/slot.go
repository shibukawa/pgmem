package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v311 int64
	_ = v311
	var v313 int64
	_ = v313
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	v8 = m.G0
	v10 = v8 - int32(2352)
	m.G0 = v10
	v14 = base.AtomicRmwXchg32(m, l0, int32(0), int32(1))
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_s_lock(m, l0, int32(_a_F_SaveSlotToPath_0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v18 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)) = uint8(v18)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v18))
	if v20 != int32(1) {
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
	m.G0 = v10 + int32(2352)
	return
L7:
	;
	v27 = l0 + int32(208)
	v29 = F_LWLockAcquire(m, v27, int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	base.MemoryFill(m, v10+int32(120), int32(0), int32(184))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = l1
	v38 = v10 + int32(1328)
	v42 = F_pg_sprintf(m, v38, int32(_a_F_SaveSlotToPath_1), v10+int32(96))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = l1
	v50 = F_pg_sprintf(m, v10+int32(304), int32(_a_F_SaveSlotToPath_2), v10+int32(80))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v53 = F_OpenTransientFile(m, v38, int32(193))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	if v53 < int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[0]))
	F_LWLockRelease(m, v27)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+112)) = int64(790273982469)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+104)) = int64(-4277855071)
	v85 = base.AtomicRmwXchg32(m, l0, int32(0), int32(1))
	if v85 != 0 {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[0])) = v58
	v64 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	if v64 == int32(0) {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v38
	F_errmsg(m, int32(_a_F_SaveSlotToPath_3), v10)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_SaveSlotToPath_4), int32(2564), int32(_a_F_SaveSlotToPath_5))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	goto L6
L21:
	;
	F_s_lock(m, l0, int32(_a_F_SaveSlotToPath_0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	base.MemoryCopy(m, v10+int32(120), l0+int32(24), int32(184))
	v95 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v95))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v10)+108))
	v102 = m.Env.Pgmem_crc32c(m, v98, v10+int32(112), int32(192))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v10)+108)) = v102 ^ int32(-1)
	*(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[0])) = v95
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v110))) = int32(167772211)
	v115 = int32(200)
	v116 = F_write(m, v53, v10+int32(104), v115)
	mBase = m.M
	if v116 != v115 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L23
L25:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[0]))
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = int32(0)
	v125 = F_CloseTransientFile(m, v53)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v154 = int32(_a_F_SaveSlotToPath_6)
	v155 = *(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[1]))
	v156 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v155))) = v156
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = int32(167772210)
	v164 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SaveSlotToPath[2])))
	if v164 != int32(1) {
		v178 = v156
		goto L39
	} else {
		goto L40
	}
L28:
	;
	v128 = v10 + int32(1328)
	v129 = F_unlink(m, v128)
	mBase = m.M
	F_LWLockRelease(m, v27)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	if v120 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v134 = v120
	goto L32
L31:
	;
	v134 = int32(51)
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[0])) = v134
	v137 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	if v137 == int32(0) {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v128
	F_errmsg(m, int32(_a_F_SaveSlotToPath_7), v10-int32(-64))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_SaveSlotToPath_4), int32(2600), int32(_a_F_SaveSlotToPath_5))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	goto L6
L38:
	;
	if v178 != 0 {
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
	v169 = F_fsync(m, v53)
	mBase = m.M
	if v169 != int32(-1) {
		v178 = v169
		goto L39
	} else {
		goto L43
	}
L42:
	;
	v178 = int32(-1)
	goto L39
L43:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[0]))
	if v173 == int32(27) {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v180 = *(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[0]))
	v182 = *(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v182))) = int32(0)
	v185 = F_CloseTransientFile(m, v53)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L4
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v213 = *(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v213))) = int32(0)
	v216 = F_CloseTransientFile(m, v53)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L4
	} else {
		goto L55
	}
L48:
	;
	v188 = v10 + int32(1328)
	v189 = F_unlink(m, v188)
	mBase = m.M
	F_LWLockRelease(m, v27)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[0])) = v180
	v195 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	if v195 == int32(0) {
		goto L6
	} else {
		goto L51
	}
L51:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v188
	F_errmsg(m, int32(_a_F_SaveSlotToPath_8), v10+int32(48))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_SaveSlotToPath_4), int32(2620), int32(_a_F_SaveSlotToPath_5))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	goto L6
L55:
	;
	if v216 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v219 = *(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[0]))
	v221 = v10 + int32(1328)
	v222 = F_unlink(m, v221)
	mBase = m.M
	F_LWLockRelease(m, v27)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L4
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v246 = v10 + int32(1328)
	v248 = v10 + int32(304)
	v249 = F_rename(m, v246, v248)
	mBase = m.M
	if v249 != 0 {
		goto L65
	} else {
		goto L66
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[0])) = v219
	v228 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	if v228 == int32(0) {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v221
	F_errmsg(m, int32(_a_F_SaveSlotToPath_9), v10+int32(32))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_SaveSlotToPath_4), int32(2636), int32(_a_F_SaveSlotToPath_5))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	goto L6
L65:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[0]))
	v252 = F_unlink(m, v246)
	mBase = m.M
	F_LWLockRelease(m, v27)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L4
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v276 = int32(_a_F_SaveSlotToPath_10)
	v278 = *(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[3])) = v278 + int32(1)
	F_fsync_fname(m, v10+int32(304), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L4
	} else {
		goto L74
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[0])) = v251
	v258 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	if v258 == int32(0) {
		goto L6
	} else {
		goto L70
	}
L70:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v246
	F_errmsg(m, int32(_a_F_SaveSlotToPath_11), v10+int32(16))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_SaveSlotToPath_4), int32(2652), int32(_a_F_SaveSlotToPath_5))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
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
	v289 = m.ExcPending
	if v289 != 0 {
		goto L4
	} else {
		goto L75
	}
L75:
	;
	F_fsync_fname(m, int32(_a_F_SaveSlotToPath_12), int32(1))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L4
	} else {
		goto L76
	}
L76:
	;
	v294 = int32(_a_F_SaveSlotToPath_10)
	v296 = *(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[3]))
	v297 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_SaveSlotToPath[3])) = v296 - v297
	v302 = base.AtomicRmwXchg32(m, l0, int32(0), v297)
	if v302 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	F_s_lock(m, l0, int32(_a_F_SaveSlotToPath_0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L4
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v306 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	goto L79
L81:
	;
	v309 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v309)
	goto L83
L82:
	;
	goto L83
L83:
	;
	v311 = *(*int64)(unsafe.Add(mBase, uint32(v10)+216))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+264)) = v311
	v313 = *(*int64)(unsafe.Add(mBase, uint32(v10)+200))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+280)) = v313
	v315 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v315))
	F_LWLockRelease(m, v27)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	goto L6
}
func F_SlotSyncShmemInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v2 = int32(_a_F_SlotSyncShmemInit_0)
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_SlotSyncShmemInit[0]))
	v4 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3))) = v4
	*(*int64)(unsafe.Add(mBase, uint32(v3)+16)) = v4
	*(*int64)(unsafe.Add(mBase, uint32(v3)+8)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v3))) = int32(-1)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_SlotSyncShmemInit[0]))
	v14 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v13)+16)), uint32(v14))
	return
}
func F_slot_getmissingattrs(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int64
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if l2 <= v13 {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
		if v15 == int32(0) {
			if l2 <= l1 {
			} else {
				v91 = (l2 - l1) & int32(3)
				if v91 != 0 {
					v96 = l1
					v97 = int32(0)
					for {
						v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v100+v96<<(uint(int32(3))%32)))) = int64(0)
						v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v108 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v106+v96))) = uint8(v108)
						v111 = v96 + v108
						v113 = v97 + v108
						if v113 != v91 {
							v96 = v111
							v97 = v113
							continue
						} else {
							break
						}
						break
					}
					v118 = v111
				} else {
					v118 = l1
				}
				if base.Ui32(int32(-4)) < base.Ui32(l1-l2) {
				} else {
					v128 = v118
					for {
						v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v133 = int32(3)
						v136 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v132+v128<<(uint(v133)%32)))) = v136
						v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v140 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v138+v128))) = uint8(v140)
						v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v144 = v128 + v140
						*(*int64)(unsafe.Add(mBase, uint32(v142+v144<<(uint(v133)%32)))) = v136
						v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						*(*uint8)(unsafe.Add(mBase, uint32(v150+v144))) = uint8(v140)
						v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v156 = v128 + int32(2)
						*(*int64)(unsafe.Add(mBase, uint32(v154+v156<<(uint(v133)%32)))) = v136
						v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						*(*uint8)(unsafe.Add(mBase, uint32(v162+v156))) = uint8(v140)
						v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v168 = v128 + v133
						*(*int64)(unsafe.Add(mBase, uint32(v166+v168<<(uint(v133)%32)))) = v136
						v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						*(*uint8)(unsafe.Add(mBase, uint32(v174+v168))) = uint8(v140)
						v179 = v128 + int32(4)
						if v179 != l2 {
							v128 = v179
							continue
						} else {
							break
						}
						break
					}
				}
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
			if v18 == int32(0) {
				if l2 <= l1 {
				} else {
					v91 = (l2 - l1) & int32(3)
					if v91 != 0 {
						v96 = l1
						v97 = int32(0)
						for {
							v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v100+v96<<(uint(int32(3))%32)))) = int64(0)
							v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v108 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v106+v96))) = uint8(v108)
							v111 = v96 + v108
							v113 = v97 + v108
							if v113 != v91 {
								v96 = v111
								v97 = v113
								continue
							} else {
								break
							}
							break
						}
						v118 = v111
					} else {
						v118 = l1
					}
					if base.Ui32(int32(-4)) < base.Ui32(l1-l2) {
					} else {
						v128 = v118
						for {
							v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v133 = int32(3)
							v136 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v132+v128<<(uint(v133)%32)))) = v136
							v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v140 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v138+v128))) = uint8(v140)
							v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v144 = v128 + v140
							*(*int64)(unsafe.Add(mBase, uint32(v142+v144<<(uint(v133)%32)))) = v136
							v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							*(*uint8)(unsafe.Add(mBase, uint32(v150+v144))) = uint8(v140)
							v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v156 = v128 + int32(2)
							*(*int64)(unsafe.Add(mBase, uint32(v154+v156<<(uint(v133)%32)))) = v136
							v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							*(*uint8)(unsafe.Add(mBase, uint32(v162+v156))) = uint8(v140)
							v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v168 = v128 + v133
							*(*int64)(unsafe.Add(mBase, uint32(v166+v168<<(uint(v133)%32)))) = v136
							v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							*(*uint8)(unsafe.Add(mBase, uint32(v174+v168))) = uint8(v140)
							v179 = v128 + int32(4)
							if v179 != l2 {
								v128 = v179
								continue
							} else {
								break
							}
							break
						}
					}
				}
			} else {
				if l2 <= l1 {
				} else {
					v22 = int32(1)
					v23 = l1 + v22
					if (l2-l1)&v22 != 0 {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v33 = v18 + l1<<(uint(int32(4))%32)
						v34 = *(*int64)(unsafe.Add(mBase, uint32(v33)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v27+l1<<(uint(int32(3))%32)))) = v34
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
						v40 = v38 ^ int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v36+l1))) = uint8(v40)
						v42 = v23
					} else {
						v42 = l1
					}
					if l2 == v23 {
					} else {
						v46 = v42
						for {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v53 = int32(3)
							v56 = int32(4)
							v58 = v18 + v46<<(uint(v56)%32)
							v59 = *(*int64)(unsafe.Add(mBase, uint32(v58)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v52+v46<<(uint(v53)%32)))) = v59
							v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
							v64 = int32(1)
							v65 = v63 ^ v64
							*(*uint8)(unsafe.Add(mBase, uint32(v61+v46))) = uint8(v65)
							v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v69 = v46 + v64
							v75 = v18 + v69<<(uint(v56)%32)
							v76 = *(*int64)(unsafe.Add(mBase, uint32(v75)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v67+v69<<(uint(v53)%32)))) = v76
							v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
							v82 = v80 ^ v64
							*(*uint8)(unsafe.Add(mBase, uint32(v78+v69))) = uint8(v82)
							v85 = v46 + int32(2)
							if v85 != l2 {
								v46 = v85
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
		m.G0 = v10 + int32(16)
		return
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v194 = m.ExcPending
		if v194 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = l2
			F_errmsg_internal(m, int32(_a_F_slot_getmissingattrs_0), v10)
			mBase = m.M
			v198 = m.ExcPending
			if v198 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_slot_getmissingattrs_1), int32(2157), int32(_a_F_slot_getmissingattrs_2))
				mBase = m.M
				v203 = m.ExcPending
				if v203 != 0 {
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
