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
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v307 int32
	_ = v307
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
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v334 int64
	_ = v334
	var v336 int64
	_ = v336
	var v341 int32
	_ = v341
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
	F_s_lock(m, l0, int32(484206), int32(2328), int32(316502))
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
	v38 = F__emscripten_memset_bulkmem(m, v11+int32(120), base.I32_extend8_s(int32(0)), int32(184))
	mBase = m.M
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = l1
	v45 = F_pg_sprintf(m, v11+int32(1328), int32(231688), v11+int32(96))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = l1
	v53 = F_pg_sprintf(m, v11+int32(304), int32(345673), v11+int32(80))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v58 = F_OpenTransientFile(m, v11+int32(1328), int32(193))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	if v58 < int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	F_LWLockRelease(m, v29)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+112)) = int64(790273982469)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+104)) = int64(-4277855071)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	if v90 != 0 {
		goto L22
	} else {
		goto L23
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = v63
	v69 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	if v69 == int32(0) {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v11 + int32(1328)
	F_errmsg(m, int32(294270), v11)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(484206), int32(2361), int32(316502))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	goto L6
L22:
	;
	F_s_lock(m, l0, int32(484206), int32(2370), int32(316502))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	goto L27
L25:
	;
	goto L24
L26:
	;
	v107 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v107
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	v111 = m.Env.Pgmem_crc32c(m, v109, v11+int32(112), int32(192))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v11)+108)) = v111 ^ int32(-1)
	*(*int32)(unsafe.Add(mBase, _consts[40])) = v107
	v119 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = int32(167772209)
	v124 = int32(200)
	v125 = F_write(m, v58, v11+int32(104), v124)
	mBase = m.M
	if v125 != v124 {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v105 = F__emscripten_memcpy_bulkmem(m, v11+int32(120), l0+int32(24), int32(184))
	mBase = m.M
	goto L29
L29:
	;
	goto L26
L30:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	v131 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = int32(0)
	v134 = F_CloseTransientFile(m, v58)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L4
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v165 = int32(4089132)
	v166 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v167 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v166))) = v167
	v170 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = int32(167772208)
	v175 = int32(*(*uint8)(unsafe.Add(mBase, _consts[41])))
	if v175 != int32(1) {
		v189 = v167
		goto L44
	} else {
		goto L45
	}
L33:
	;
	v138 = F_unlink(m, v11+int32(1328))
	mBase = m.M
	F_LWLockRelease(m, v29)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	if v129 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v143 = v129
	goto L37
L36:
	;
	v143 = int32(51)
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = v143
	v146 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	if v146 == int32(0) {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v11 + int32(1328)
	F_errmsg(m, int32(293274), v11-int32(-64))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(484206), int32(2397), int32(316502))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	goto L6
L43:
	;
	if v189 != 0 {
		goto L50
	} else {
		goto L51
	}
L44:
	;
	goto L43
L45:
	;
	goto L46
L46:
	;
	v180 = F_fsync(m, v58)
	mBase = m.M
	if v180 != int32(-1) {
		v189 = v180
		goto L44
	} else {
		goto L48
	}
L47:
	;
	v189 = int32(-1)
	goto L44
L48:
	;
	v184 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v184 == int32(27) {
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v191 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	v193 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	*(*int32)(unsafe.Add(mBase, uint32(v193))) = int32(0)
	v196 = F_CloseTransientFile(m, v58)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L4
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v226 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	*(*int32)(unsafe.Add(mBase, uint32(v226))) = int32(0)
	v229 = F_CloseTransientFile(m, v58)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L4
	} else {
		goto L60
	}
L53:
	;
	v200 = F_unlink(m, v11+int32(1328))
	mBase = m.M
	F_LWLockRelease(m, v29)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = v191
	v206 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	if v206 == int32(0) {
		goto L6
	} else {
		goto L56
	}
L56:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v11 + int32(1328)
	F_errmsg(m, int32(294539), v11+int32(48))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(484206), int32(2417), int32(316502))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	goto L6
L60:
	;
	if v229 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v232 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	v235 = F_unlink(m, v11+int32(1328))
	mBase = m.M
	F_LWLockRelease(m, v29)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L4
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v264 = F_rename(m, v11+int32(1328), v11+int32(304))
	mBase = m.M
	if v264 != 0 {
		goto L70
	} else {
		goto L71
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = v232
	v241 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	if v241 == int32(0) {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v11 + int32(1328)
	F_errmsg(m, int32(294334), v11+int32(32))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(484206), int32(2433), int32(316502))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	goto L6
L70:
	;
	v266 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	v269 = F_unlink(m, v11+int32(1328))
	mBase = m.M
	F_LWLockRelease(m, v29)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L4
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v297 = int32(4465060)
	v299 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v299 + int32(1)
	F_fsync_fname(m, v11+int32(304), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L4
	} else {
		goto L79
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = v266
	v275 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	if v275 == int32(0) {
		goto L6
	} else {
		goto L75
	}
L75:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L4
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v11 + int32(304)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v11 + int32(1328)
	F_errmsg(m, int32(292340), v11+int32(16))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(484206), int32(2449), int32(316502))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	goto L6
L79:
	;
	F_fsync_fname(m, l1, int32(1))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	F_fsync_fname(m, int32(83530), int32(1))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	v315 = int32(4465060)
	v317 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	v318 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v317 - v318
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v318
	if v321 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	F_s_lock(m, l0, int32(484206), int32(2468), int32(316502))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L4
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v329 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	goto L84
L86:
	;
	v332 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v332)
	goto L88
L87:
	;
	goto L88
L88:
	;
	v334 = *(*int64)(unsafe.Add(mBase, uint32(v11)+216))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+264)) = v334
	v336 = *(*int64)(unsafe.Add(mBase, uint32(v11)+200))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+280)) = v336
	F_LWLockRelease(m, v29)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
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
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	if v8 == int32(0) {
		v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v87 = int32(2)
		v91 = l2 - l1
		v95 = F__emscripten_memset_bulkmem(m, v86+l1<<(uint(v87)%32), base.I32_extend8_s(int32(0)), v91<<(uint(v87)%32))
		mBase = m.M
		v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v100 = F__emscripten_memset_bulkmem(m, v96+l1, base.I32_extend8_s(int32(1)), v91)
		mBase = m.M
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
		if v11 == int32(0) {
			v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v87 = int32(2)
			v91 = l2 - l1
			v95 = F__emscripten_memset_bulkmem(m, v86+l1<<(uint(v87)%32), base.I32_extend8_s(int32(0)), v91<<(uint(v87)%32))
			mBase = m.M
			v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v100 = F__emscripten_memset_bulkmem(m, v96+l1, base.I32_extend8_s(int32(1)), v91)
			mBase = m.M
			return
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
			return
		}
	}
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
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
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
					v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v101 = int32(2)
					v105 = v2 - v16
					v108 = F___memset(m, v100+v16<<(uint(v101)%32), int32(0), v105<<(uint(v101)%32))
					mBase = m.M
					v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v112 = F___memset(m, v109+v16, int32(1), v105)
					mBase = m.M
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
					if v25 == int32(0) {
						v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v101 = int32(2)
						v105 = v2 - v16
						v108 = F___memset(m, v100+v16<<(uint(v101)%32), int32(0), v105<<(uint(v101)%32))
						mBase = m.M
						v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v112 = F___memset(m, v109+v16, int32(1), v105)
						mBase = m.M
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
		v126 = m.ExcPending
		if v126 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v2
			F_errmsg_internal(m, int32(463592), v7)
			mBase = m.M
			v130 = m.ExcPending
			if v130 != 0 {
				return
			} else {
				F_errfinish(m, int32(485304), int32(2098), int32(89770))
				mBase = m.M
				v135 = m.ExcPending
				if v135 != 0 {
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
