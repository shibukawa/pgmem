package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RunObjectDropHook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l3
	v15 = *(*int32)(unsafe.Add(mBase, _consts[380]))
	m.T0[v15].(func(*base.Module, int32, int32, int32, int32, int32))(m, int32(1), l0, l1, l2, v8+int32(12))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		m.G0 = v8 + int32(16)
		return
	}
}
func F_add_object_address(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	if v10 <= v9 {
		*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v10 << (uint(int32(1)) % 32)
		v17 = F_repalloc(m, v8, v10*int32(24))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v17
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
			v21 = v17
			v22 = v20
			v25 = v22*int32(12) + v21
			*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v25))) = l0
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v29 + int32(1)
			return
		}
	} else {
		v21 = v8
		v22 = v9
		v25 = v22*int32(12) + v21
		*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v25))) = l0
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v29 + int32(1)
		return
	}
}
func F_getObjectDescriptionOids(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v3 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l0
	v15 = F_getObjectDescription(m, v6+int32(4), v3)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return v15
	}
}
func F_get_object_attnum_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[381]))
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L23
	} else {
		goto L24
	}
L2:
	;
	v67 = int32(*(*int16)(unsafe.Add(mBase, uint32(v64)+22)))
	m.G0 = v7 + int32(16)
	return v67
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v11 == l0 {
		v64 = v10
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v16 = int32(0)
	goto L8
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[381])) = v61
	v64 = v61
	goto L2
L8:
	;
	v19 = v16 * int32(40)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[382])))
	if l0 != v22 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v61 = v19 + int32(747920)
	goto L7
L10:
	;
	if v16 == int32(36) {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	goto L9
L13:
	;
	v29 = (v16 | int32(1)) * int32(40)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[382])))
	if l0 == v32 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v61 = v29 + int32(747920)
	goto L7
L15:
	;
	goto L16
L16:
	;
	v39 = (v16 | int32(2)) * int32(40)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[382])))
	if l0 == v42 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v61 = v39 + int32(747920)
	goto L7
L18:
	;
	goto L19
L19:
	;
	v49 = (v16 | int32(3)) * int32(40)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+uint32(_consts[382])))
	if l0 == v52 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v61 = v49 + int32(747920)
	goto L7
L21:
	;
	v16 = v16 + int32(4)
	goto L8
L23:
	;
	return int32(0)
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	F_errmsg_internal(m, int32(59243), v7)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(490604), int32(2777), int32(500758))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_object_catcache_oid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[381]))
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L23
	} else {
		goto L24
	}
L2:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	m.G0 = v7 + int32(16)
	return v67
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v11 == l0 {
		v64 = v10
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v16 = int32(0)
	goto L8
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[381])) = v61
	v64 = v61
	goto L2
L8:
	;
	v19 = v16 * int32(40)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[382])))
	if l0 != v22 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v61 = v19 + int32(747920)
	goto L7
L10:
	;
	if v16 == int32(36) {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	goto L9
L13:
	;
	v29 = (v16 | int32(1)) * int32(40)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[382])))
	if l0 == v32 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v61 = v29 + int32(747920)
	goto L7
L15:
	;
	goto L16
L16:
	;
	v39 = (v16 | int32(2)) * int32(40)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[382])))
	if l0 == v42 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v61 = v39 + int32(747920)
	goto L7
L18:
	;
	goto L19
L19:
	;
	v49 = (v16 | int32(3)) * int32(40)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+uint32(_consts[382])))
	if l0 == v52 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v61 = v49 + int32(747920)
	goto L7
L21:
	;
	v16 = v16 + int32(4)
	goto L8
L23:
	;
	return int32(0)
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	F_errmsg_internal(m, int32(59243), v7)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(490604), int32(2777), int32(500758))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_object_oid_index(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[381]))
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L23
	} else {
		goto L24
	}
L2:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	m.G0 = v7 + int32(16)
	return v67
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v11 == l0 {
		v64 = v10
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v16 = int32(0)
	goto L8
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[381])) = v61
	v64 = v61
	goto L2
L8:
	;
	v19 = v16 * int32(40)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[382])))
	if l0 != v22 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v61 = v19 + int32(747920)
	goto L7
L10:
	;
	if v16 == int32(36) {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	goto L9
L13:
	;
	v29 = (v16 | int32(1)) * int32(40)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[382])))
	if l0 == v32 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v61 = v29 + int32(747920)
	goto L7
L15:
	;
	goto L16
L16:
	;
	v39 = (v16 | int32(2)) * int32(40)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[382])))
	if l0 == v42 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v61 = v39 + int32(747920)
	goto L7
L18:
	;
	goto L19
L19:
	;
	v49 = (v16 | int32(3)) * int32(40)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+uint32(_consts[382])))
	if l0 == v52 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v61 = v49 + int32(747920)
	goto L7
L21:
	;
	v16 = v16 + int32(4)
	goto L8
L23:
	;
	return int32(0)
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	F_errmsg_internal(m, int32(59243), v7)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(490604), int32(2777), int32(500758))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_object_aclmask_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int64
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
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
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int64
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int64
	_ = v303
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v314 int64
	_ = v314
	v11 = m.G0
	v13 = v11 - int32(80)
	m.G0 = v13
	if l0 != int32(1247) {
		if l0 != int32(2615) {
			v253 = F_superuser_arg(m, l2)
			mBase = m.M
			v254 = m.ExcPending
			if v254 != 0 {
				return int64(0)
			} else {
				if v253 != 0 {
					v314 = l3
					m.G0 = v13 + int32(80)
					return v314
				} else {
					v255 = F_get_object_catcache_oid(m, l0)
					mBase = m.M
					v256 = m.ExcPending
					if v256 != 0 {
						return int64(0)
					} else {
						v257 = F_SearchSysCache1(m, v255, l1)
						mBase = m.M
						v258 = m.ExcPending
						if v258 != 0 {
							return int64(0)
						} else {
							if v257 == int32(0) {
								if l4 != 0 {
									v261 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v261)
									v314 = int64(0)
									m.G0 = v13 + int32(80)
									return v314
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v267 = m.ExcPending
									if v267 != 0 {
										return int64(0)
									} else {
										v268 = F_get_object_class_descr(m, l0)
										mBase = m.M
										v269 = m.ExcPending
										if v269 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = l1
											*(*int32)(unsafe.Add(mBase, uint32(v13))) = v268
											F_errmsg_internal(m, int32(42474), v13)
											mBase = m.M
											v274 = m.ExcPending
											if v274 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(494511), int32(3107), int32(63388))
												mBase = m.M
												v279 = m.ExcPending
												if v279 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								}
							} else {
								v280 = F_get_object_attnum_owner(m, l0)
								mBase = m.M
								v281 = m.ExcPending
								if v281 != 0 {
									return int64(0)
								} else {
									v282 = F_SysCacheGetAttrNotNull(m, v255, v257, v280)
									mBase = m.M
									v283 = m.ExcPending
									if v283 != 0 {
										return int64(0)
									} else {
										v284 = F_get_object_attnum_acl(m, l0)
										mBase = m.M
										v285 = m.ExcPending
										if v285 != 0 {
											return int64(0)
										} else {
											v288 = F_SysCacheGetAttr(m, v255, v257, v284, v13+int32(79))
											mBase = m.M
											v289 = m.ExcPending
											if v289 != 0 {
												return int64(0)
											} else {
												v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+79)))
												if v290 == int32(1) {
													v294 = F_get_object_type(m, l0, l1)
													mBase = m.M
													v295 = m.ExcPending
													if v295 != 0 {
														return int64(0)
													} else {
														v296 = F_acldefault(m, v294, v282)
														mBase = m.M
														v297 = m.ExcPending
														if v297 != 0 {
															return int64(0)
														} else {
															v300 = int32(0)
															v301 = v296
															v303 = F_aclmask(m, v301, l2, v282, l3, int32(1))
															mBase = m.M
															v304 = m.ExcPending
															if v304 != 0 {
																return int64(0)
															} else {
																if v301 == int32(0) {
																	F_ReleaseCatCache(m, v257)
																	mBase = m.M
																	v311 = m.ExcPending
																	if v311 != 0 {
																		return int64(0)
																	} else {
																		v314 = v303
																		m.G0 = v13 + int32(80)
																		return v314
																	}
																} else {
																	if v301 == v300 {
																		F_ReleaseCatCache(m, v257)
																		mBase = m.M
																		v311 = m.ExcPending
																		if v311 != 0 {
																			return int64(0)
																		} else {
																			v314 = v303
																			m.G0 = v13 + int32(80)
																			return v314
																		}
																	} else {
																		F_pfree(m, v301)
																		mBase = m.M
																		v309 = m.ExcPending
																		if v309 != 0 {
																			return int64(0)
																		} else {
																			F_ReleaseCatCache(m, v257)
																			mBase = m.M
																			v311 = m.ExcPending
																			if v311 != 0 {
																				return int64(0)
																			} else {
																				v314 = v303
																				m.G0 = v13 + int32(80)
																				return v314
																			}
																		}
																	}
																}
															}
														}
													}
												} else {
													v298 = F_pg_detoast_datum(m, v288)
													mBase = m.M
													v299 = m.ExcPending
													if v299 != 0 {
														return int64(0)
													} else {
														v300 = v288
														v301 = v298
														v303 = F_aclmask(m, v301, l2, v282, l3, int32(1))
														mBase = m.M
														v304 = m.ExcPending
														if v304 != 0 {
															return int64(0)
														} else {
															if v301 == int32(0) {
																F_ReleaseCatCache(m, v257)
																mBase = m.M
																v311 = m.ExcPending
																if v311 != 0 {
																	return int64(0)
																} else {
																	v314 = v303
																	m.G0 = v13 + int32(80)
																	return v314
																}
															} else {
																if v301 == v300 {
																	F_ReleaseCatCache(m, v257)
																	mBase = m.M
																	v311 = m.ExcPending
																	if v311 != 0 {
																		return int64(0)
																	} else {
																		v314 = v303
																		m.G0 = v13 + int32(80)
																		return v314
																	}
																} else {
																	F_pfree(m, v301)
																	mBase = m.M
																	v309 = m.ExcPending
																	if v309 != 0 {
																		return int64(0)
																	} else {
																		F_ReleaseCatCache(m, v257)
																		mBase = m.M
																		v311 = m.ExcPending
																		if v311 != 0 {
																			return int64(0)
																		} else {
																			v314 = v303
																			m.G0 = v13 + int32(80)
																			return v314
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v19 = F_superuser_arg(m, l2)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int64(0)
			} else {
				if v19 != 0 {
					v314 = l3
					m.G0 = v13 + int32(80)
					return v314
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, _consts[251]))
					if base.B2i32(v25 != int32(0))&base.B2i32(l1 == v25) != 0 {
						v32 = *(*int32)(unsafe.Add(mBase, _consts[100]))
						v34 = F_object_aclmask_ext(m, int32(1262), v32, l2, int64(1024), l4)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int64(0)
						} else {
							if v34 != int64(0) {
								v314 = l3 & int64(768)
							} else {
								v314 = l3 & int64(256)
							}
							m.G0 = v13 + int32(80)
							return v314
						}
					} else {
						v43 = F_SearchSysCache1(m, int32(38), l1)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int64(0)
						} else {
							if v43 == int32(0) {
								if l4 != 0 {
									v47 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v47)
									v314 = int64(0)
									m.G0 = v13 + int32(80)
									return v314
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(1411))
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l1
											F_errmsg(m, int32(69119), v13+int32(16))
											mBase = m.M
											v62 = m.ExcPending
											if v62 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(494511), int32(3667), int32(63473))
												mBase = m.M
												v67 = m.ExcPending
												if v67 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								}
							} else {
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
								v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+22)))
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v68+v69)+68))
								v76 = F_SysCacheGetAttr(m, int32(38), v43, int32(4), v13+int32(79))
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int64(0)
								} else {
									v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+79)))
									if v78 == int32(1) {
										v83 = F_acldefault(m, int32(36), v71)
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return int64(0)
										} else {
											v87 = int32(0)
											v88 = v83
											v90 = F_aclmask(m, v88, l2, v71, l3, int32(1))
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
												return int64(0)
											} else {
												if v88 == int32(0) {
													F_ReleaseCatCache(m, v43)
													mBase = m.M
													v98 = m.ExcPending
													if v98 != 0 {
														return int64(0)
													} else {
														if l3&int64(256) == int64(0) {
															v314 = v90
															m.G0 = v13 + int32(80)
															return v314
														} else {
															if v90&int64(256) != int64(0) {
																v314 = v90
																m.G0 = v13 + int32(80)
																return v314
															} else {
																v108 = F_has_privs_of_role(m, l2, int32(6181))
																mBase = m.M
																v109 = m.ExcPending
																if v109 != 0 {
																	return int64(0)
																} else {
																	if v108 != 0 {
																		v314 = v90 | int64(256)
																		m.G0 = v13 + int32(80)
																		return v314
																	} else {
																		v111 = F_has_privs_of_role(m, l2, int32(6182))
																		mBase = m.M
																		v112 = m.ExcPending
																		if v112 != 0 {
																			return int64(0)
																		} else {
																			if v111 != 0 {
																				v314 = v90 | int64(256)
																			} else {
																				v314 = v90
																			}
																			m.G0 = v13 + int32(80)
																			return v314
																		}
																	}
																}
															}
														}
													}
												} else {
													if v88 == v87 {
														F_ReleaseCatCache(m, v43)
														mBase = m.M
														v98 = m.ExcPending
														if v98 != 0 {
															return int64(0)
														} else {
															if l3&int64(256) == int64(0) {
																v314 = v90
																m.G0 = v13 + int32(80)
																return v314
															} else {
																if v90&int64(256) != int64(0) {
																	v314 = v90
																	m.G0 = v13 + int32(80)
																	return v314
																} else {
																	v108 = F_has_privs_of_role(m, l2, int32(6181))
																	mBase = m.M
																	v109 = m.ExcPending
																	if v109 != 0 {
																		return int64(0)
																	} else {
																		if v108 != 0 {
																			v314 = v90 | int64(256)
																			m.G0 = v13 + int32(80)
																			return v314
																		} else {
																			v111 = F_has_privs_of_role(m, l2, int32(6182))
																			mBase = m.M
																			v112 = m.ExcPending
																			if v112 != 0 {
																				return int64(0)
																			} else {
																				if v111 != 0 {
																					v314 = v90 | int64(256)
																				} else {
																					v314 = v90
																				}
																				m.G0 = v13 + int32(80)
																				return v314
																			}
																		}
																	}
																}
															}
														}
													} else {
														F_pfree(m, v88)
														mBase = m.M
														v96 = m.ExcPending
														if v96 != 0 {
															return int64(0)
														} else {
															F_ReleaseCatCache(m, v43)
															mBase = m.M
															v98 = m.ExcPending
															if v98 != 0 {
																return int64(0)
															} else {
																if l3&int64(256) == int64(0) {
																	v314 = v90
																	m.G0 = v13 + int32(80)
																	return v314
																} else {
																	if v90&int64(256) != int64(0) {
																		v314 = v90
																		m.G0 = v13 + int32(80)
																		return v314
																	} else {
																		v108 = F_has_privs_of_role(m, l2, int32(6181))
																		mBase = m.M
																		v109 = m.ExcPending
																		if v109 != 0 {
																			return int64(0)
																		} else {
																			if v108 != 0 {
																				v314 = v90 | int64(256)
																				m.G0 = v13 + int32(80)
																				return v314
																			} else {
																				v111 = F_has_privs_of_role(m, l2, int32(6182))
																				mBase = m.M
																				v112 = m.ExcPending
																				if v112 != 0 {
																					return int64(0)
																				} else {
																					if v111 != 0 {
																						v314 = v90 | int64(256)
																					} else {
																						v314 = v90
																					}
																					m.G0 = v13 + int32(80)
																					return v314
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									} else {
										v85 = F_pg_detoast_datum(m, v76)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return int64(0)
										} else {
											v87 = v76
											v88 = v85
											v90 = F_aclmask(m, v88, l2, v71, l3, int32(1))
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
												return int64(0)
											} else {
												if v88 == int32(0) {
													F_ReleaseCatCache(m, v43)
													mBase = m.M
													v98 = m.ExcPending
													if v98 != 0 {
														return int64(0)
													} else {
														if l3&int64(256) == int64(0) {
															v314 = v90
															m.G0 = v13 + int32(80)
															return v314
														} else {
															if v90&int64(256) != int64(0) {
																v314 = v90
																m.G0 = v13 + int32(80)
																return v314
															} else {
																v108 = F_has_privs_of_role(m, l2, int32(6181))
																mBase = m.M
																v109 = m.ExcPending
																if v109 != 0 {
																	return int64(0)
																} else {
																	if v108 != 0 {
																		v314 = v90 | int64(256)
																		m.G0 = v13 + int32(80)
																		return v314
																	} else {
																		v111 = F_has_privs_of_role(m, l2, int32(6182))
																		mBase = m.M
																		v112 = m.ExcPending
																		if v112 != 0 {
																			return int64(0)
																		} else {
																			if v111 != 0 {
																				v314 = v90 | int64(256)
																			} else {
																				v314 = v90
																			}
																			m.G0 = v13 + int32(80)
																			return v314
																		}
																	}
																}
															}
														}
													}
												} else {
													if v88 == v87 {
														F_ReleaseCatCache(m, v43)
														mBase = m.M
														v98 = m.ExcPending
														if v98 != 0 {
															return int64(0)
														} else {
															if l3&int64(256) == int64(0) {
																v314 = v90
																m.G0 = v13 + int32(80)
																return v314
															} else {
																if v90&int64(256) != int64(0) {
																	v314 = v90
																	m.G0 = v13 + int32(80)
																	return v314
																} else {
																	v108 = F_has_privs_of_role(m, l2, int32(6181))
																	mBase = m.M
																	v109 = m.ExcPending
																	if v109 != 0 {
																		return int64(0)
																	} else {
																		if v108 != 0 {
																			v314 = v90 | int64(256)
																			m.G0 = v13 + int32(80)
																			return v314
																		} else {
																			v111 = F_has_privs_of_role(m, l2, int32(6182))
																			mBase = m.M
																			v112 = m.ExcPending
																			if v112 != 0 {
																				return int64(0)
																			} else {
																				if v111 != 0 {
																					v314 = v90 | int64(256)
																				} else {
																					v314 = v90
																				}
																				m.G0 = v13 + int32(80)
																				return v314
																			}
																		}
																	}
																}
															}
														}
													} else {
														F_pfree(m, v88)
														mBase = m.M
														v96 = m.ExcPending
														if v96 != 0 {
															return int64(0)
														} else {
															F_ReleaseCatCache(m, v43)
															mBase = m.M
															v98 = m.ExcPending
															if v98 != 0 {
																return int64(0)
															} else {
																if l3&int64(256) == int64(0) {
																	v314 = v90
																	m.G0 = v13 + int32(80)
																	return v314
																} else {
																	if v90&int64(256) != int64(0) {
																		v314 = v90
																		m.G0 = v13 + int32(80)
																		return v314
																	} else {
																		v108 = F_has_privs_of_role(m, l2, int32(6181))
																		mBase = m.M
																		v109 = m.ExcPending
																		if v109 != 0 {
																			return int64(0)
																		} else {
																			if v108 != 0 {
																				v314 = v90 | int64(256)
																				m.G0 = v13 + int32(80)
																				return v314
																			} else {
																				v111 = F_has_privs_of_role(m, l2, int32(6182))
																				mBase = m.M
																				v112 = m.ExcPending
																				if v112 != 0 {
																					return int64(0)
																				} else {
																					if v111 != 0 {
																						v314 = v90 | int64(256)
																					} else {
																						v314 = v90
																					}
																					m.G0 = v13 + int32(80)
																					return v314
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v115 = F_superuser_arg(m, l2)
		mBase = m.M
		v116 = m.ExcPending
		if v116 != 0 {
			return int64(0)
		} else {
			if v115 != 0 {
				v314 = l3
				m.G0 = v13 + int32(80)
				return v314
			} else {
				v118 = F_SearchSysCache1(m, int32(82), l1)
				mBase = m.M
				v119 = m.ExcPending
				if v119 != 0 {
					return int64(0)
				} else {
					if v118 == int32(0) {
						if l4 != 0 {
							v122 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v122)
							v314 = int64(0)
							m.G0 = v13 + int32(80)
							return v314
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v128 = m.ExcPending
							if v128 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(67137668))
								mBase = m.M
								v131 = m.ExcPending
								if v131 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = l1
									F_errmsg(m, int32(69017), v13+int32(32))
									mBase = m.M
									v137 = m.ExcPending
									if v137 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(494511), int32(3742), int32(63453))
										mBase = m.M
										v142 = m.ExcPending
										if v142 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						}
					} else {
						v143 = *(*int32)(unsafe.Add(mBase, uint32(v118)+16))
						v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+22)))
						v145 = v143 + v144
						v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+92))
						if v146 == int32(0) {
							v183 = v118
							v184 = v145
							v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+79)))
							if v185 == int32(109) {
								v188 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
								v189 = F_get_multirange_range(m, v188)
								mBase = m.M
								v190 = m.ExcPending
								if v190 != 0 {
									return int64(0)
								} else {
									F_ReleaseCatCache(m, v183)
									mBase = m.M
									v192 = m.ExcPending
									if v192 != 0 {
										return int64(0)
									} else {
										v194 = F_SearchSysCache1(m, int32(82), v189)
										mBase = m.M
										v195 = m.ExcPending
										if v195 != 0 {
											return int64(0)
										} else {
											if v194 == int32(0) {
												if l4 != 0 {
													v198 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v198)
													v314 = int64(0)
													m.G0 = v13 + int32(80)
													return v314
												} else {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v204 = m.ExcPending
													if v204 != 0 {
														return int64(0)
													} else {
														F_errcode(m, int32(67137668))
														mBase = m.M
														v207 = m.ExcPending
														if v207 != 0 {
															return int64(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v189
															F_errmsg(m, int32(69017), v13+int32(48))
															mBase = m.M
															v213 = m.ExcPending
															if v213 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(494511), int32(3798), int32(63453))
																mBase = m.M
																v218 = m.ExcPending
																if v218 != 0 {
																	return int64(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												}
											} else {
												v219 = *(*int32)(unsafe.Add(mBase, uint32(v194)+16))
												v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+22)))
												v222 = v194
												v224 = v219 + v220
												v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+72))
												v230 = F_SysCacheGetAttr(m, int32(82), v222, int32(32), v13+int32(79))
												mBase = m.M
												v231 = m.ExcPending
												if v231 != 0 {
													return int64(0)
												} else {
													v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+79)))
													if v232 == int32(1) {
														v237 = F_acldefault(m, int32(49), v225)
														mBase = m.M
														v238 = m.ExcPending
														if v238 != 0 {
															return int64(0)
														} else {
															v241 = int32(0)
															v242 = v237
															v244 = F_aclmask(m, v242, l2, v225, l3, int32(1))
															mBase = m.M
															v245 = m.ExcPending
															if v245 != 0 {
																return int64(0)
															} else {
																if v242 == int32(0) {
																	F_ReleaseCatCache(m, v222)
																	mBase = m.M
																	v252 = m.ExcPending
																	if v252 != 0 {
																		return int64(0)
																	} else {
																		v314 = v244
																		m.G0 = v13 + int32(80)
																		return v314
																	}
																} else {
																	if v242 == v241 {
																		F_ReleaseCatCache(m, v222)
																		mBase = m.M
																		v252 = m.ExcPending
																		if v252 != 0 {
																			return int64(0)
																		} else {
																			v314 = v244
																			m.G0 = v13 + int32(80)
																			return v314
																		}
																	} else {
																		F_pfree(m, v242)
																		mBase = m.M
																		v250 = m.ExcPending
																		if v250 != 0 {
																			return int64(0)
																		} else {
																			F_ReleaseCatCache(m, v222)
																			mBase = m.M
																			v252 = m.ExcPending
																			if v252 != 0 {
																				return int64(0)
																			} else {
																				v314 = v244
																				m.G0 = v13 + int32(80)
																				return v314
																			}
																		}
																	}
																}
															}
														}
													} else {
														v239 = F_pg_detoast_datum(m, v230)
														mBase = m.M
														v240 = m.ExcPending
														if v240 != 0 {
															return int64(0)
														} else {
															v241 = v230
															v242 = v239
															v244 = F_aclmask(m, v242, l2, v225, l3, int32(1))
															mBase = m.M
															v245 = m.ExcPending
															if v245 != 0 {
																return int64(0)
															} else {
																if v242 == int32(0) {
																	F_ReleaseCatCache(m, v222)
																	mBase = m.M
																	v252 = m.ExcPending
																	if v252 != 0 {
																		return int64(0)
																	} else {
																		v314 = v244
																		m.G0 = v13 + int32(80)
																		return v314
																	}
																} else {
																	if v242 == v241 {
																		F_ReleaseCatCache(m, v222)
																		mBase = m.M
																		v252 = m.ExcPending
																		if v252 != 0 {
																			return int64(0)
																		} else {
																			v314 = v244
																			m.G0 = v13 + int32(80)
																			return v314
																		}
																	} else {
																		F_pfree(m, v242)
																		mBase = m.M
																		v250 = m.ExcPending
																		if v250 != 0 {
																			return int64(0)
																		} else {
																			F_ReleaseCatCache(m, v222)
																			mBase = m.M
																			v252 = m.ExcPending
																			if v252 != 0 {
																				return int64(0)
																			} else {
																				v314 = v244
																				m.G0 = v13 + int32(80)
																				return v314
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								v222 = v183
								v224 = v184
								v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+72))
								v230 = F_SysCacheGetAttr(m, int32(82), v222, int32(32), v13+int32(79))
								mBase = m.M
								v231 = m.ExcPending
								if v231 != 0 {
									return int64(0)
								} else {
									v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+79)))
									if v232 == int32(1) {
										v237 = F_acldefault(m, int32(49), v225)
										mBase = m.M
										v238 = m.ExcPending
										if v238 != 0 {
											return int64(0)
										} else {
											v241 = int32(0)
											v242 = v237
											v244 = F_aclmask(m, v242, l2, v225, l3, int32(1))
											mBase = m.M
											v245 = m.ExcPending
											if v245 != 0 {
												return int64(0)
											} else {
												if v242 == int32(0) {
													F_ReleaseCatCache(m, v222)
													mBase = m.M
													v252 = m.ExcPending
													if v252 != 0 {
														return int64(0)
													} else {
														v314 = v244
														m.G0 = v13 + int32(80)
														return v314
													}
												} else {
													if v242 == v241 {
														F_ReleaseCatCache(m, v222)
														mBase = m.M
														v252 = m.ExcPending
														if v252 != 0 {
															return int64(0)
														} else {
															v314 = v244
															m.G0 = v13 + int32(80)
															return v314
														}
													} else {
														F_pfree(m, v242)
														mBase = m.M
														v250 = m.ExcPending
														if v250 != 0 {
															return int64(0)
														} else {
															F_ReleaseCatCache(m, v222)
															mBase = m.M
															v252 = m.ExcPending
															if v252 != 0 {
																return int64(0)
															} else {
																v314 = v244
																m.G0 = v13 + int32(80)
																return v314
															}
														}
													}
												}
											}
										}
									} else {
										v239 = F_pg_detoast_datum(m, v230)
										mBase = m.M
										v240 = m.ExcPending
										if v240 != 0 {
											return int64(0)
										} else {
											v241 = v230
											v242 = v239
											v244 = F_aclmask(m, v242, l2, v225, l3, int32(1))
											mBase = m.M
											v245 = m.ExcPending
											if v245 != 0 {
												return int64(0)
											} else {
												if v242 == int32(0) {
													F_ReleaseCatCache(m, v222)
													mBase = m.M
													v252 = m.ExcPending
													if v252 != 0 {
														return int64(0)
													} else {
														v314 = v244
														m.G0 = v13 + int32(80)
														return v314
													}
												} else {
													if v242 == v241 {
														F_ReleaseCatCache(m, v222)
														mBase = m.M
														v252 = m.ExcPending
														if v252 != 0 {
															return int64(0)
														} else {
															v314 = v244
															m.G0 = v13 + int32(80)
															return v314
														}
													} else {
														F_pfree(m, v242)
														mBase = m.M
														v250 = m.ExcPending
														if v250 != 0 {
															return int64(0)
														} else {
															F_ReleaseCatCache(m, v222)
															mBase = m.M
															v252 = m.ExcPending
															if v252 != 0 {
																return int64(0)
															} else {
																v314 = v244
																m.G0 = v13 + int32(80)
																return v314
															}
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							v149 = *(*int32)(unsafe.Add(mBase, uint32(v145)+88))
							if v149 != int32(6179) {
								v183 = v118
								v184 = v145
								v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+79)))
								if v185 == int32(109) {
									v188 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
									v189 = F_get_multirange_range(m, v188)
									mBase = m.M
									v190 = m.ExcPending
									if v190 != 0 {
										return int64(0)
									} else {
										F_ReleaseCatCache(m, v183)
										mBase = m.M
										v192 = m.ExcPending
										if v192 != 0 {
											return int64(0)
										} else {
											v194 = F_SearchSysCache1(m, int32(82), v189)
											mBase = m.M
											v195 = m.ExcPending
											if v195 != 0 {
												return int64(0)
											} else {
												if v194 == int32(0) {
													if l4 != 0 {
														v198 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v198)
														v314 = int64(0)
														m.G0 = v13 + int32(80)
														return v314
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v204 = m.ExcPending
														if v204 != 0 {
															return int64(0)
														} else {
															F_errcode(m, int32(67137668))
															mBase = m.M
															v207 = m.ExcPending
															if v207 != 0 {
																return int64(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v189
																F_errmsg(m, int32(69017), v13+int32(48))
																mBase = m.M
																v213 = m.ExcPending
																if v213 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(494511), int32(3798), int32(63453))
																	mBase = m.M
																	v218 = m.ExcPending
																	if v218 != 0 {
																		return int64(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														}
													}
												} else {
													v219 = *(*int32)(unsafe.Add(mBase, uint32(v194)+16))
													v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+22)))
													v222 = v194
													v224 = v219 + v220
													v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+72))
													v230 = F_SysCacheGetAttr(m, int32(82), v222, int32(32), v13+int32(79))
													mBase = m.M
													v231 = m.ExcPending
													if v231 != 0 {
														return int64(0)
													} else {
														v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+79)))
														if v232 == int32(1) {
															v237 = F_acldefault(m, int32(49), v225)
															mBase = m.M
															v238 = m.ExcPending
															if v238 != 0 {
																return int64(0)
															} else {
																v241 = int32(0)
																v242 = v237
																v244 = F_aclmask(m, v242, l2, v225, l3, int32(1))
																mBase = m.M
																v245 = m.ExcPending
																if v245 != 0 {
																	return int64(0)
																} else {
																	if v242 == int32(0) {
																		F_ReleaseCatCache(m, v222)
																		mBase = m.M
																		v252 = m.ExcPending
																		if v252 != 0 {
																			return int64(0)
																		} else {
																			v314 = v244
																			m.G0 = v13 + int32(80)
																			return v314
																		}
																	} else {
																		if v242 == v241 {
																			F_ReleaseCatCache(m, v222)
																			mBase = m.M
																			v252 = m.ExcPending
																			if v252 != 0 {
																				return int64(0)
																			} else {
																				v314 = v244
																				m.G0 = v13 + int32(80)
																				return v314
																			}
																		} else {
																			F_pfree(m, v242)
																			mBase = m.M
																			v250 = m.ExcPending
																			if v250 != 0 {
																				return int64(0)
																			} else {
																				F_ReleaseCatCache(m, v222)
																				mBase = m.M
																				v252 = m.ExcPending
																				if v252 != 0 {
																					return int64(0)
																				} else {
																					v314 = v244
																					m.G0 = v13 + int32(80)
																					return v314
																				}
																			}
																		}
																	}
																}
															}
														} else {
															v239 = F_pg_detoast_datum(m, v230)
															mBase = m.M
															v240 = m.ExcPending
															if v240 != 0 {
																return int64(0)
															} else {
																v241 = v230
																v242 = v239
																v244 = F_aclmask(m, v242, l2, v225, l3, int32(1))
																mBase = m.M
																v245 = m.ExcPending
																if v245 != 0 {
																	return int64(0)
																} else {
																	if v242 == int32(0) {
																		F_ReleaseCatCache(m, v222)
																		mBase = m.M
																		v252 = m.ExcPending
																		if v252 != 0 {
																			return int64(0)
																		} else {
																			v314 = v244
																			m.G0 = v13 + int32(80)
																			return v314
																		}
																	} else {
																		if v242 == v241 {
																			F_ReleaseCatCache(m, v222)
																			mBase = m.M
																			v252 = m.ExcPending
																			if v252 != 0 {
																				return int64(0)
																			} else {
																				v314 = v244
																				m.G0 = v13 + int32(80)
																				return v314
																			}
																		} else {
																			F_pfree(m, v242)
																			mBase = m.M
																			v250 = m.ExcPending
																			if v250 != 0 {
																				return int64(0)
																			} else {
																				F_ReleaseCatCache(m, v222)
																				mBase = m.M
																				v252 = m.ExcPending
																				if v252 != 0 {
																					return int64(0)
																				} else {
																					v314 = v244
																					m.G0 = v13 + int32(80)
																					return v314
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								} else {
									v222 = v183
									v224 = v184
									v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+72))
									v230 = F_SysCacheGetAttr(m, int32(82), v222, int32(32), v13+int32(79))
									mBase = m.M
									v231 = m.ExcPending
									if v231 != 0 {
										return int64(0)
									} else {
										v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+79)))
										if v232 == int32(1) {
											v237 = F_acldefault(m, int32(49), v225)
											mBase = m.M
											v238 = m.ExcPending
											if v238 != 0 {
												return int64(0)
											} else {
												v241 = int32(0)
												v242 = v237
												v244 = F_aclmask(m, v242, l2, v225, l3, int32(1))
												mBase = m.M
												v245 = m.ExcPending
												if v245 != 0 {
													return int64(0)
												} else {
													if v242 == int32(0) {
														F_ReleaseCatCache(m, v222)
														mBase = m.M
														v252 = m.ExcPending
														if v252 != 0 {
															return int64(0)
														} else {
															v314 = v244
															m.G0 = v13 + int32(80)
															return v314
														}
													} else {
														if v242 == v241 {
															F_ReleaseCatCache(m, v222)
															mBase = m.M
															v252 = m.ExcPending
															if v252 != 0 {
																return int64(0)
															} else {
																v314 = v244
																m.G0 = v13 + int32(80)
																return v314
															}
														} else {
															F_pfree(m, v242)
															mBase = m.M
															v250 = m.ExcPending
															if v250 != 0 {
																return int64(0)
															} else {
																F_ReleaseCatCache(m, v222)
																mBase = m.M
																v252 = m.ExcPending
																if v252 != 0 {
																	return int64(0)
																} else {
																	v314 = v244
																	m.G0 = v13 + int32(80)
																	return v314
																}
															}
														}
													}
												}
											}
										} else {
											v239 = F_pg_detoast_datum(m, v230)
											mBase = m.M
											v240 = m.ExcPending
											if v240 != 0 {
												return int64(0)
											} else {
												v241 = v230
												v242 = v239
												v244 = F_aclmask(m, v242, l2, v225, l3, int32(1))
												mBase = m.M
												v245 = m.ExcPending
												if v245 != 0 {
													return int64(0)
												} else {
													if v242 == int32(0) {
														F_ReleaseCatCache(m, v222)
														mBase = m.M
														v252 = m.ExcPending
														if v252 != 0 {
															return int64(0)
														} else {
															v314 = v244
															m.G0 = v13 + int32(80)
															return v314
														}
													} else {
														if v242 == v241 {
															F_ReleaseCatCache(m, v222)
															mBase = m.M
															v252 = m.ExcPending
															if v252 != 0 {
																return int64(0)
															} else {
																v314 = v244
																m.G0 = v13 + int32(80)
																return v314
															}
														} else {
															F_pfree(m, v242)
															mBase = m.M
															v250 = m.ExcPending
															if v250 != 0 {
																return int64(0)
															} else {
																F_ReleaseCatCache(m, v222)
																mBase = m.M
																v252 = m.ExcPending
																if v252 != 0 {
																	return int64(0)
																} else {
																	v314 = v244
																	m.G0 = v13 + int32(80)
																	return v314
																}
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								F_ReleaseCatCache(m, v118)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return int64(0)
								} else {
									v155 = F_SearchSysCache1(m, int32(82), v146)
									mBase = m.M
									v156 = m.ExcPending
									if v156 != 0 {
										return int64(0)
									} else {
										if v155 == int32(0) {
											if l4 != 0 {
												v159 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v159)
												v314 = int64(0)
												m.G0 = v13 + int32(80)
												return v314
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v165 = m.ExcPending
												if v165 != 0 {
													return int64(0)
												} else {
													F_errcode(m, int32(67137668))
													mBase = m.M
													v168 = m.ExcPending
													if v168 != 0 {
														return int64(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v146
														F_errmsg(m, int32(69017), v13-int32(-64))
														mBase = m.M
														v174 = m.ExcPending
														if v174 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(494511), int32(3769), int32(63453))
															mBase = m.M
															v179 = m.ExcPending
															if v179 != 0 {
																return int64(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											}
										} else {
											v180 = *(*int32)(unsafe.Add(mBase, uint32(v155)+16))
											v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+22)))
											v183 = v155
											v184 = v180 + v181
											v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+79)))
											if v185 == int32(109) {
												v188 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
												v189 = F_get_multirange_range(m, v188)
												mBase = m.M
												v190 = m.ExcPending
												if v190 != 0 {
													return int64(0)
												} else {
													F_ReleaseCatCache(m, v183)
													mBase = m.M
													v192 = m.ExcPending
													if v192 != 0 {
														return int64(0)
													} else {
														v194 = F_SearchSysCache1(m, int32(82), v189)
														mBase = m.M
														v195 = m.ExcPending
														if v195 != 0 {
															return int64(0)
														} else {
															if v194 == int32(0) {
																if l4 != 0 {
																	v198 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v198)
																	v314 = int64(0)
																	m.G0 = v13 + int32(80)
																	return v314
																} else {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v204 = m.ExcPending
																	if v204 != 0 {
																		return int64(0)
																	} else {
																		F_errcode(m, int32(67137668))
																		mBase = m.M
																		v207 = m.ExcPending
																		if v207 != 0 {
																			return int64(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v189
																			F_errmsg(m, int32(69017), v13+int32(48))
																			mBase = m.M
																			v213 = m.ExcPending
																			if v213 != 0 {
																				return int64(0)
																			} else {
																				F_errfinish(m, int32(494511), int32(3798), int32(63453))
																				mBase = m.M
																				v218 = m.ExcPending
																				if v218 != 0 {
																					return int64(0)
																				} else {
																					base.Wasm_trap_unreachable()
																					for {
																					}
																				}
																			}
																		}
																	}
																}
															} else {
																v219 = *(*int32)(unsafe.Add(mBase, uint32(v194)+16))
																v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+22)))
																v222 = v194
																v224 = v219 + v220
																v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+72))
																v230 = F_SysCacheGetAttr(m, int32(82), v222, int32(32), v13+int32(79))
																mBase = m.M
																v231 = m.ExcPending
																if v231 != 0 {
																	return int64(0)
																} else {
																	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+79)))
																	if v232 == int32(1) {
																		v237 = F_acldefault(m, int32(49), v225)
																		mBase = m.M
																		v238 = m.ExcPending
																		if v238 != 0 {
																			return int64(0)
																		} else {
																			v241 = int32(0)
																			v242 = v237
																			v244 = F_aclmask(m, v242, l2, v225, l3, int32(1))
																			mBase = m.M
																			v245 = m.ExcPending
																			if v245 != 0 {
																				return int64(0)
																			} else {
																				if v242 == int32(0) {
																					F_ReleaseCatCache(m, v222)
																					mBase = m.M
																					v252 = m.ExcPending
																					if v252 != 0 {
																						return int64(0)
																					} else {
																						v314 = v244
																						m.G0 = v13 + int32(80)
																						return v314
																					}
																				} else {
																					if v242 == v241 {
																						F_ReleaseCatCache(m, v222)
																						mBase = m.M
																						v252 = m.ExcPending
																						if v252 != 0 {
																							return int64(0)
																						} else {
																							v314 = v244
																							m.G0 = v13 + int32(80)
																							return v314
																						}
																					} else {
																						F_pfree(m, v242)
																						mBase = m.M
																						v250 = m.ExcPending
																						if v250 != 0 {
																							return int64(0)
																						} else {
																							F_ReleaseCatCache(m, v222)
																							mBase = m.M
																							v252 = m.ExcPending
																							if v252 != 0 {
																								return int64(0)
																							} else {
																								v314 = v244
																								m.G0 = v13 + int32(80)
																								return v314
																							}
																						}
																					}
																				}
																			}
																		}
																	} else {
																		v239 = F_pg_detoast_datum(m, v230)
																		mBase = m.M
																		v240 = m.ExcPending
																		if v240 != 0 {
																			return int64(0)
																		} else {
																			v241 = v230
																			v242 = v239
																			v244 = F_aclmask(m, v242, l2, v225, l3, int32(1))
																			mBase = m.M
																			v245 = m.ExcPending
																			if v245 != 0 {
																				return int64(0)
																			} else {
																				if v242 == int32(0) {
																					F_ReleaseCatCache(m, v222)
																					mBase = m.M
																					v252 = m.ExcPending
																					if v252 != 0 {
																						return int64(0)
																					} else {
																						v314 = v244
																						m.G0 = v13 + int32(80)
																						return v314
																					}
																				} else {
																					if v242 == v241 {
																						F_ReleaseCatCache(m, v222)
																						mBase = m.M
																						v252 = m.ExcPending
																						if v252 != 0 {
																							return int64(0)
																						} else {
																							v314 = v244
																							m.G0 = v13 + int32(80)
																							return v314
																						}
																					} else {
																						F_pfree(m, v242)
																						mBase = m.M
																						v250 = m.ExcPending
																						if v250 != 0 {
																							return int64(0)
																						} else {
																							F_ReleaseCatCache(m, v222)
																							mBase = m.M
																							v252 = m.ExcPending
																							if v252 != 0 {
																								return int64(0)
																							} else {
																								v314 = v244
																								m.G0 = v13 + int32(80)
																								return v314
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												v222 = v183
												v224 = v184
												v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+72))
												v230 = F_SysCacheGetAttr(m, int32(82), v222, int32(32), v13+int32(79))
												mBase = m.M
												v231 = m.ExcPending
												if v231 != 0 {
													return int64(0)
												} else {
													v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+79)))
													if v232 == int32(1) {
														v237 = F_acldefault(m, int32(49), v225)
														mBase = m.M
														v238 = m.ExcPending
														if v238 != 0 {
															return int64(0)
														} else {
															v241 = int32(0)
															v242 = v237
															v244 = F_aclmask(m, v242, l2, v225, l3, int32(1))
															mBase = m.M
															v245 = m.ExcPending
															if v245 != 0 {
																return int64(0)
															} else {
																if v242 == int32(0) {
																	F_ReleaseCatCache(m, v222)
																	mBase = m.M
																	v252 = m.ExcPending
																	if v252 != 0 {
																		return int64(0)
																	} else {
																		v314 = v244
																		m.G0 = v13 + int32(80)
																		return v314
																	}
																} else {
																	if v242 == v241 {
																		F_ReleaseCatCache(m, v222)
																		mBase = m.M
																		v252 = m.ExcPending
																		if v252 != 0 {
																			return int64(0)
																		} else {
																			v314 = v244
																			m.G0 = v13 + int32(80)
																			return v314
																		}
																	} else {
																		F_pfree(m, v242)
																		mBase = m.M
																		v250 = m.ExcPending
																		if v250 != 0 {
																			return int64(0)
																		} else {
																			F_ReleaseCatCache(m, v222)
																			mBase = m.M
																			v252 = m.ExcPending
																			if v252 != 0 {
																				return int64(0)
																			} else {
																				v314 = v244
																				m.G0 = v13 + int32(80)
																				return v314
																			}
																		}
																	}
																}
															}
														}
													} else {
														v239 = F_pg_detoast_datum(m, v230)
														mBase = m.M
														v240 = m.ExcPending
														if v240 != 0 {
															return int64(0)
														} else {
															v241 = v230
															v242 = v239
															v244 = F_aclmask(m, v242, l2, v225, l3, int32(1))
															mBase = m.M
															v245 = m.ExcPending
															if v245 != 0 {
																return int64(0)
															} else {
																if v242 == int32(0) {
																	F_ReleaseCatCache(m, v222)
																	mBase = m.M
																	v252 = m.ExcPending
																	if v252 != 0 {
																		return int64(0)
																	} else {
																		v314 = v244
																		m.G0 = v13 + int32(80)
																		return v314
																	}
																} else {
																	if v242 == v241 {
																		F_ReleaseCatCache(m, v222)
																		mBase = m.M
																		v252 = m.ExcPending
																		if v252 != 0 {
																			return int64(0)
																		} else {
																			v314 = v244
																			m.G0 = v13 + int32(80)
																			return v314
																		}
																	} else {
																		F_pfree(m, v242)
																		mBase = m.M
																		v250 = m.ExcPending
																		if v250 != 0 {
																			return int64(0)
																		} else {
																			F_ReleaseCatCache(m, v222)
																			mBase = m.M
																			v252 = m.ExcPending
																			if v252 != 0 {
																				return int64(0)
																			} else {
																				v314 = v244
																				m.G0 = v13 + int32(80)
																				return v314
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_parse_object_field(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v8 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = int32(11)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v20 = int32(0)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v21|v22 == v20 {
		v37 = v20
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v15 = int32(14)
	goto L6
L5:
	;
	v15 = v11
	goto L6
L6:
	;
	if v8 == int32(12) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v18 = v11
	goto L9
L8:
	;
	v18 = v15
	goto L9
L9:
	;
	return v18
L10:
	;
	v38 = F_json_lex(m, l0)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L13
	} else {
		goto L18
	}
L11:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v26 != int32(1) {
		v37 = v20
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v31 = F_pstrdup(m, v30)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	if v31 != 0 {
		v37 = v31
		goto L10
	} else {
		goto L15
	}
L15:
	;
	return int32(16)
L16:
	;
	return v88
L17:
	;
	F_pfree(m, v37)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L13
	} else {
		goto L56
	}
L18:
	;
	if v38 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v37 != 0 {
		v84 = v38
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v40 != int32(8) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v88 = v38
	goto L16
L23:
	;
	v43 = int32(11)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v46 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v51 = F_json_lex(m, l0)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L13
	} else {
		goto L33
	}
L26:
	;
	v47 = int32(8)
	goto L28
L27:
	;
	v47 = v43
	goto L28
L28:
	;
	if v40 == int32(12) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v50 = v43
	goto L31
L30:
	;
	v50 = v47
	goto L31
L31:
	;
	if v37 != 0 {
		v84 = v50
		goto L17
	} else {
		goto L32
	}
L32:
	;
	v88 = v50
	goto L16
L33:
	;
	if v51 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	if v37 != 0 {
		v84 = v51
		goto L17
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v21 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v88 = v51
	goto L16
L38:
	;
	if v37 == int32(0) {
		v88 = v76
		goto L16
	} else {
		goto L54
	}
L39:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v57 = m.T0[v21].(func(*base.Module, int32, int32, int32) int32)(m, v54, v37, base.B2i32(v53 == int32(11)))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L13
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	switch v53 - int32(3) {
	case 0:
		goto L47
	default:
		goto L45
	case 2:
		goto L46
	}
L42:
	;
	if v57 != 0 {
		v76 = v57
		goto L38
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	if v68 != 0 {
		v76 = v68
		goto L38
	} else {
		goto L51
	}
L45:
	;
	v66 = F_parse_scalar(m, l0, l1)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L13
	} else {
		goto L50
	}
L46:
	;
	v64 = F_parse_array(m, l0, l1)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L13
	} else {
		goto L49
	}
L47:
	;
	v62 = F_parse_object(m, l0, l1)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L13
	} else {
		goto L48
	}
L48:
	;
	v68 = v62
	goto L44
L49:
	;
	v68 = v64
	goto L44
L50:
	;
	v68 = v66
	goto L44
L51:
	;
	if v22 == int32(0) {
		v76 = v68
		goto L38
	} else {
		goto L52
	}
L52:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v74 = m.T0[v22].(func(*base.Module, int32, int32, int32) int32)(m, v71, v37, base.B2i32(v53 == int32(11)))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L13
	} else {
		goto L53
	}
L53:
	;
	v76 = v74
	goto L38
L54:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v79&int32(4) == int32(0) {
		v88 = v76
		goto L16
	} else {
		goto L55
	}
L55:
	;
	v84 = v76
	goto L17
L56:
	;
	v88 = v84
	goto L16
}
func F_storeObjectDescription(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	v6 = m.G0
	v8 = v6 - int32(128)
	m.G0 = v8
	v11 = F_getObjectDescription(m, l2, int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		if v11 != 0 {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v13 != 0 {
				F_appendStringInfoChar(m, l0, int32(10))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					if base.Ui32(int32(2)) <= base.Ui32(l1) {
						if l1 != int32(2) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
								F_errmsg_internal(m, int32(481751), v8)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return
								} else {
									F_errfinish(m, int32(496455), int32(1322), int32(245557))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+116)) = v11
							*(*int32)(unsafe.Add(mBase, uint32(v8)+112)) = int32(0)
							F_appendStringInfo(m, l0, int32(184216), v8+int32(112))
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return
							} else {
								F_pfree(m, v11)
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return
								} else {
									m.G0 = v8 + int32(128)
									return
								}
							}
						}
					} else {
						switch l3 - int32(97) {
						case 0:
							*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v11
							F_appendStringInfo(m, l0, int32(180495), v8+int32(48))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return
							} else {
								F_pfree(m, v11)
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return
								} else {
									m.G0 = v8 + int32(128)
									return
								}
							}
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l3
								F_errmsg_internal(m, int32(481376), v8+int32(16))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return
								} else {
									F_errfinish(m, int32(496455), int32(1310), int32(245557))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						case 8:
							*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v11
							F_appendStringInfo(m, l0, int32(180487), v8-int32(-64))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return
							} else {
								F_pfree(m, v11)
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return
								} else {
									m.G0 = v8 + int32(128)
									return
								}
							}
						case 14:
							*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v11
							F_appendStringInfo(m, l0, int32(185731), v8+int32(32))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return
							} else {
								F_pfree(m, v11)
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return
								} else {
									m.G0 = v8 + int32(128)
									return
								}
							}
						case 17:
							*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v11
							F_appendStringInfo(m, l0, int32(185702), v8+int32(80))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								F_pfree(m, v11)
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return
								} else {
									m.G0 = v8 + int32(128)
									return
								}
							}
						case 19:
							*(*int32)(unsafe.Add(mBase, uint32(v8)+96)) = v11
							F_appendStringInfo(m, l0, int32(180659), v8+int32(96))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								F_pfree(m, v11)
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return
								} else {
									m.G0 = v8 + int32(128)
									return
								}
							}
						}
					}
				}
			} else {
				if base.Ui32(int32(2)) <= base.Ui32(l1) {
					if l1 != int32(2) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
							F_errmsg_internal(m, int32(481751), v8)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return
							} else {
								F_errfinish(m, int32(496455), int32(1322), int32(245557))
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+116)) = v11
						*(*int32)(unsafe.Add(mBase, uint32(v8)+112)) = int32(0)
						F_appendStringInfo(m, l0, int32(184216), v8+int32(112))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							F_pfree(m, v11)
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return
							} else {
								m.G0 = v8 + int32(128)
								return
							}
						}
					}
				} else {
					switch l3 - int32(97) {
					case 0:
						*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v11
						F_appendStringInfo(m, l0, int32(180495), v8+int32(48))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							F_pfree(m, v11)
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return
							} else {
								m.G0 = v8 + int32(128)
								return
							}
						}
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l3
							F_errmsg_internal(m, int32(481376), v8+int32(16))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								F_errfinish(m, int32(496455), int32(1310), int32(245557))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 8:
						*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v11
						F_appendStringInfo(m, l0, int32(180487), v8-int32(-64))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							F_pfree(m, v11)
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return
							} else {
								m.G0 = v8 + int32(128)
								return
							}
						}
					case 14:
						*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v11
						F_appendStringInfo(m, l0, int32(185731), v8+int32(32))
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return
						} else {
							F_pfree(m, v11)
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return
							} else {
								m.G0 = v8 + int32(128)
								return
							}
						}
					case 17:
						*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v11
						F_appendStringInfo(m, l0, int32(185702), v8+int32(80))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							F_pfree(m, v11)
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return
							} else {
								m.G0 = v8 + int32(128)
								return
							}
						}
					case 19:
						*(*int32)(unsafe.Add(mBase, uint32(v8)+96)) = v11
						F_appendStringInfo(m, l0, int32(180659), v8+int32(96))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							F_pfree(m, v11)
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return
							} else {
								m.G0 = v8 + int32(128)
								return
							}
						}
					}
				}
			}
		} else {
			m.G0 = v8 + int32(128)
			return
		}
	}
}
