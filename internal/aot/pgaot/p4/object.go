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
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_RunObjectDropHook[0]))
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_get_object_attnum_name[0]))
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	v52 = int32(*(*int16)(unsafe.Add(mBase, uint32(v48)+22)))
	m.G0 = v8 + int32(16)
	return v52
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v12 == l0 {
		v48 = v11
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v18 = int32(0)
	goto L11
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_get_object_attnum_name[0])) = v45
	v48 = v45
	goto L2
L8:
	;
	v45 = v21 + int32(_a_F_get_object_attnum_name_0)
	goto L7
L9:
	;
	v45 = v21 + int32(_a_F_get_object_attnum_name_1)
	goto L7
L10:
	;
	v45 = v21 + int32(_a_F_get_object_attnum_name_2)
	goto L7
L11:
	;
	v21 = v18 * int32(40)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_attnum_name[1])))
	if l0 != v22 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v45 = v21 + int32(_a_F_get_object_attnum_name_3)
	goto L7
L13:
	;
	if v18 == int32(36) {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	goto L12
L16:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_attnum_name[2])))
	if v28 == l0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_attnum_name[3])))
	if v30 == l0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_attnum_name[4])))
	if v32 == l0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	v18 = v18 + int32(4)
	goto L11
L20:
	;
	return int32(0)
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	F_errmsg_internal(m, int32(_a_F_get_object_attnum_name_4), v8)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_get_object_attnum_name_5), int32(2777), int32(_a_F_get_object_attnum_name_6))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_object_catcache_oid(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_get_object_catcache_oid[0]))
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	m.G0 = v8 + int32(16)
	return v52
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v12 == l0 {
		v48 = v11
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v18 = int32(0)
	goto L11
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_get_object_catcache_oid[0])) = v45
	v48 = v45
	goto L2
L8:
	;
	v45 = v21 + int32(_a_F_get_object_catcache_oid_0)
	goto L7
L9:
	;
	v45 = v21 + int32(_a_F_get_object_catcache_oid_1)
	goto L7
L10:
	;
	v45 = v21 + int32(_a_F_get_object_catcache_oid_2)
	goto L7
L11:
	;
	v21 = v18 * int32(40)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_catcache_oid[1])))
	if l0 != v22 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v45 = v21 + int32(_a_F_get_object_catcache_oid_3)
	goto L7
L13:
	;
	if v18 == int32(36) {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	goto L12
L16:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_catcache_oid[2])))
	if v28 == l0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_catcache_oid[3])))
	if v30 == l0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_catcache_oid[4])))
	if v32 == l0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	v18 = v18 + int32(4)
	goto L11
L20:
	;
	return int32(0)
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	F_errmsg_internal(m, int32(_a_F_get_object_catcache_oid_4), v8)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_get_object_catcache_oid_5), int32(2777), int32(_a_F_get_object_catcache_oid_6))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_object_oid_index(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_get_object_oid_index[0]))
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	m.G0 = v8 + int32(16)
	return v52
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v12 == l0 {
		v48 = v11
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v18 = int32(0)
	goto L11
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_get_object_oid_index[0])) = v45
	v48 = v45
	goto L2
L8:
	;
	v45 = v21 + int32(_a_F_get_object_oid_index_0)
	goto L7
L9:
	;
	v45 = v21 + int32(_a_F_get_object_oid_index_1)
	goto L7
L10:
	;
	v45 = v21 + int32(_a_F_get_object_oid_index_2)
	goto L7
L11:
	;
	v21 = v18 * int32(40)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_oid_index[1])))
	if l0 != v22 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v45 = v21 + int32(_a_F_get_object_oid_index_3)
	goto L7
L13:
	;
	if v18 == int32(36) {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	goto L12
L16:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_oid_index[2])))
	if v28 == l0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_oid_index[3])))
	if v30 == l0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_oid_index[4])))
	if v32 == l0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	v18 = v18 + int32(4)
	goto L11
L20:
	;
	return int32(0)
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	F_errmsg_internal(m, int32(_a_F_get_object_oid_index_4), v8)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_get_object_oid_index_5), int32(2777), int32(_a_F_get_object_oid_index_6))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
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
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int64
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int64
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v323 int64
	_ = v323
	v11 = m.G0
	v13 = v11 - int32(80)
	m.G0 = v13
	if l0 != int32(1247) {
		if l0 != int32(2615) {
			v259 = F_superuser_arg(m, l2)
			mBase = m.M
			v260 = m.ExcPending
			if v260 != 0 {
				return int64(0)
			} else {
				if v259 != 0 {
					v323 = l3
					m.G0 = v13 + int32(80)
					return v323
				} else {
					v261 = F_get_object_catcache_oid(m, l0)
					mBase = m.M
					v262 = m.ExcPending
					if v262 != 0 {
						return int64(0)
					} else {
						v263 = F_SearchSysCache1(m, v261, l1)
						mBase = m.M
						v264 = m.ExcPending
						if v264 != 0 {
							return int64(0)
						} else {
							if v263 == int32(0) {
								if l4 != 0 {
									v267 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v267)
									v323 = int64(0)
									m.G0 = v13 + int32(80)
									return v323
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v273 = m.ExcPending
									if v273 != 0 {
										return int64(0)
									} else {
										v274 = F_get_object_class_descr(m, l0)
										mBase = m.M
										v275 = m.ExcPending
										if v275 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = l1
											*(*int32)(unsafe.Add(mBase, uint32(v13))) = v274
											F_errmsg_internal(m, int32(_a_F_object_aclmask_ext_0), v13)
											mBase = m.M
											v280 = m.ExcPending
											if v280 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_object_aclmask_ext_1), int32(3107), int32(_a_F_object_aclmask_ext_2))
												mBase = m.M
												v285 = m.ExcPending
												if v285 != 0 {
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
								v286 = F_get_object_attnum_owner(m, l0)
								mBase = m.M
								v287 = m.ExcPending
								if v287 != 0 {
									return int64(0)
								} else {
									v288 = F_SysCacheGetAttrNotNull(m, v261, v263, v286)
									mBase = m.M
									v289 = m.ExcPending
									if v289 != 0 {
										return int64(0)
									} else {
										v290 = F_get_object_attnum_acl(m, l0)
										mBase = m.M
										v291 = m.ExcPending
										if v291 != 0 {
											return int64(0)
										} else {
											v294 = F_SysCacheGetAttr(m, v261, v263, v290, v13+int32(79))
											mBase = m.M
											v295 = m.ExcPending
											if v295 != 0 {
												return int64(0)
											} else {
												v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+79)))
												if v296 == int32(1) {
													v300 = F_get_object_type(m, l0, l1)
													mBase = m.M
													v301 = m.ExcPending
													if v301 != 0 {
														return int64(0)
													} else {
														v302 = F_acldefault(m, v300, v288)
														mBase = m.M
														v303 = m.ExcPending
														if v303 != 0 {
															return int64(0)
														} else {
															v306 = int32(0)
															v307 = v302
															v309 = F_aclmask(m, v307, l2, v288, l3, int32(1))
															mBase = m.M
															v310 = m.ExcPending
															if v310 != 0 {
																return int64(0)
															} else {
																v311 = int32(0)
																if base.B2i32(v307 == v311)|base.B2i32(v307 == v306) == v311 {
																	F_pfree(m, v307)
																	mBase = m.M
																	v318 = m.ExcPending
																	if v318 != 0 {
																		return int64(0)
																	} else {
																		F_ReleaseCatCache(m, v263)
																		mBase = m.M
																		v320 = m.ExcPending
																		if v320 != 0 {
																			return int64(0)
																		} else {
																			v323 = v309
																			m.G0 = v13 + int32(80)
																			return v323
																		}
																	}
																} else {
																	F_ReleaseCatCache(m, v263)
																	mBase = m.M
																	v320 = m.ExcPending
																	if v320 != 0 {
																		return int64(0)
																	} else {
																		v323 = v309
																		m.G0 = v13 + int32(80)
																		return v323
																	}
																}
															}
														}
													}
												} else {
													v304 = F_pg_detoast_datum(m, v294)
													mBase = m.M
													v305 = m.ExcPending
													if v305 != 0 {
														return int64(0)
													} else {
														v306 = v294
														v307 = v304
														v309 = F_aclmask(m, v307, l2, v288, l3, int32(1))
														mBase = m.M
														v310 = m.ExcPending
														if v310 != 0 {
															return int64(0)
														} else {
															v311 = int32(0)
															if base.B2i32(v307 == v311)|base.B2i32(v307 == v306) == v311 {
																F_pfree(m, v307)
																mBase = m.M
																v318 = m.ExcPending
																if v318 != 0 {
																	return int64(0)
																} else {
																	F_ReleaseCatCache(m, v263)
																	mBase = m.M
																	v320 = m.ExcPending
																	if v320 != 0 {
																		return int64(0)
																	} else {
																		v323 = v309
																		m.G0 = v13 + int32(80)
																		return v323
																	}
																}
															} else {
																F_ReleaseCatCache(m, v263)
																mBase = m.M
																v320 = m.ExcPending
																if v320 != 0 {
																	return int64(0)
																} else {
																	v323 = v309
																	m.G0 = v13 + int32(80)
																	return v323
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
					v323 = l3
					m.G0 = v13 + int32(80)
					return v323
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, _c_F_object_aclmask_ext[0]))
					if base.B2i32(v25 != int32(0))&base.B2i32(l1 == v25) != 0 {
						v32 = *(*int32)(unsafe.Add(mBase, _c_F_object_aclmask_ext[1]))
						v34 = F_object_aclmask_ext(m, int32(1262), v32, l2, int64(1024), l4)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int64(0)
						} else {
							if v34 != int64(0) {
								v323 = l3 & int64(768)
							} else {
								v323 = l3 & int64(256)
							}
							m.G0 = v13 + int32(80)
							return v323
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
									v323 = int64(0)
									m.G0 = v13 + int32(80)
									return v323
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
											F_errmsg(m, int32(_a_F_object_aclmask_ext_3), v13+int32(16))
											mBase = m.M
											v62 = m.ExcPending
											if v62 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_object_aclmask_ext_1), int32(3667), int32(_a_F_object_aclmask_ext_4))
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
												v92 = int32(0)
												if base.B2i32(v88 == v92)|base.B2i32(v88 == v87) == v92 {
													F_pfree(m, v88)
													mBase = m.M
													v99 = m.ExcPending
													if v99 != 0 {
														return int64(0)
													} else {
														F_ReleaseCatCache(m, v43)
														mBase = m.M
														v101 = m.ExcPending
														if v101 != 0 {
															return int64(0)
														} else {
															if l3&int64(256) == int64(0) {
																v323 = v90
																m.G0 = v13 + int32(80)
																return v323
															} else {
																if v90&int64(256) != int64(0) {
																	v323 = v90
																	m.G0 = v13 + int32(80)
																	return v323
																} else {
																	v111 = F_has_privs_of_role(m, l2, int32(_a_F_object_aclmask_ext_5))
																	mBase = m.M
																	v112 = m.ExcPending
																	if v112 != 0 {
																		return int64(0)
																	} else {
																		if v111 != 0 {
																			v323 = v90 | int64(256)
																			m.G0 = v13 + int32(80)
																			return v323
																		} else {
																			v114 = F_has_privs_of_role(m, l2, int32(_a_F_object_aclmask_ext_6))
																			mBase = m.M
																			v115 = m.ExcPending
																			if v115 != 0 {
																				return int64(0)
																			} else {
																				if v114 != 0 {
																					v323 = v90 | int64(256)
																				} else {
																					v323 = v90
																				}
																				m.G0 = v13 + int32(80)
																				return v323
																			}
																		}
																	}
																}
															}
														}
													}
												} else {
													F_ReleaseCatCache(m, v43)
													mBase = m.M
													v101 = m.ExcPending
													if v101 != 0 {
														return int64(0)
													} else {
														if l3&int64(256) == int64(0) {
															v323 = v90
															m.G0 = v13 + int32(80)
															return v323
														} else {
															if v90&int64(256) != int64(0) {
																v323 = v90
																m.G0 = v13 + int32(80)
																return v323
															} else {
																v111 = F_has_privs_of_role(m, l2, int32(_a_F_object_aclmask_ext_5))
																mBase = m.M
																v112 = m.ExcPending
																if v112 != 0 {
																	return int64(0)
																} else {
																	if v111 != 0 {
																		v323 = v90 | int64(256)
																		m.G0 = v13 + int32(80)
																		return v323
																	} else {
																		v114 = F_has_privs_of_role(m, l2, int32(_a_F_object_aclmask_ext_6))
																		mBase = m.M
																		v115 = m.ExcPending
																		if v115 != 0 {
																			return int64(0)
																		} else {
																			if v114 != 0 {
																				v323 = v90 | int64(256)
																			} else {
																				v323 = v90
																			}
																			m.G0 = v13 + int32(80)
																			return v323
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
												v92 = int32(0)
												if base.B2i32(v88 == v92)|base.B2i32(v88 == v87) == v92 {
													F_pfree(m, v88)
													mBase = m.M
													v99 = m.ExcPending
													if v99 != 0 {
														return int64(0)
													} else {
														F_ReleaseCatCache(m, v43)
														mBase = m.M
														v101 = m.ExcPending
														if v101 != 0 {
															return int64(0)
														} else {
															if l3&int64(256) == int64(0) {
																v323 = v90
																m.G0 = v13 + int32(80)
																return v323
															} else {
																if v90&int64(256) != int64(0) {
																	v323 = v90
																	m.G0 = v13 + int32(80)
																	return v323
																} else {
																	v111 = F_has_privs_of_role(m, l2, int32(_a_F_object_aclmask_ext_5))
																	mBase = m.M
																	v112 = m.ExcPending
																	if v112 != 0 {
																		return int64(0)
																	} else {
																		if v111 != 0 {
																			v323 = v90 | int64(256)
																			m.G0 = v13 + int32(80)
																			return v323
																		} else {
																			v114 = F_has_privs_of_role(m, l2, int32(_a_F_object_aclmask_ext_6))
																			mBase = m.M
																			v115 = m.ExcPending
																			if v115 != 0 {
																				return int64(0)
																			} else {
																				if v114 != 0 {
																					v323 = v90 | int64(256)
																				} else {
																					v323 = v90
																				}
																				m.G0 = v13 + int32(80)
																				return v323
																			}
																		}
																	}
																}
															}
														}
													}
												} else {
													F_ReleaseCatCache(m, v43)
													mBase = m.M
													v101 = m.ExcPending
													if v101 != 0 {
														return int64(0)
													} else {
														if l3&int64(256) == int64(0) {
															v323 = v90
															m.G0 = v13 + int32(80)
															return v323
														} else {
															if v90&int64(256) != int64(0) {
																v323 = v90
																m.G0 = v13 + int32(80)
																return v323
															} else {
																v111 = F_has_privs_of_role(m, l2, int32(_a_F_object_aclmask_ext_5))
																mBase = m.M
																v112 = m.ExcPending
																if v112 != 0 {
																	return int64(0)
																} else {
																	if v111 != 0 {
																		v323 = v90 | int64(256)
																		m.G0 = v13 + int32(80)
																		return v323
																	} else {
																		v114 = F_has_privs_of_role(m, l2, int32(_a_F_object_aclmask_ext_6))
																		mBase = m.M
																		v115 = m.ExcPending
																		if v115 != 0 {
																			return int64(0)
																		} else {
																			if v114 != 0 {
																				v323 = v90 | int64(256)
																			} else {
																				v323 = v90
																			}
																			m.G0 = v13 + int32(80)
																			return v323
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
		v118 = F_superuser_arg(m, l2)
		mBase = m.M
		v119 = m.ExcPending
		if v119 != 0 {
			return int64(0)
		} else {
			if v118 != 0 {
				v323 = l3
				m.G0 = v13 + int32(80)
				return v323
			} else {
				v121 = F_SearchSysCache1(m, int32(82), l1)
				mBase = m.M
				v122 = m.ExcPending
				if v122 != 0 {
					return int64(0)
				} else {
					if v121 == int32(0) {
						if l4 != 0 {
							v125 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v125)
							v323 = int64(0)
							m.G0 = v13 + int32(80)
							return v323
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v131 = m.ExcPending
							if v131 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(67137668))
								mBase = m.M
								v134 = m.ExcPending
								if v134 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = l1
									F_errmsg(m, int32(_a_F_object_aclmask_ext_7), v13+int32(32))
									mBase = m.M
									v140 = m.ExcPending
									if v140 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_object_aclmask_ext_1), int32(3742), int32(_a_F_object_aclmask_ext_8))
										mBase = m.M
										v145 = m.ExcPending
										if v145 != 0 {
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
						v146 = *(*int32)(unsafe.Add(mBase, uint32(v121)+16))
						v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+22)))
						v148 = v146 + v147
						v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+92))
						if v149 == int32(0) {
							v186 = v121
							v187 = v148
							v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+79)))
							if v188 == int32(109) {
								v191 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
								v192 = F_get_multirange_range(m, v191)
								mBase = m.M
								v193 = m.ExcPending
								if v193 != 0 {
									return int64(0)
								} else {
									F_ReleaseCatCache(m, v186)
									mBase = m.M
									v195 = m.ExcPending
									if v195 != 0 {
										return int64(0)
									} else {
										v197 = F_SearchSysCache1(m, int32(82), v192)
										mBase = m.M
										v198 = m.ExcPending
										if v198 != 0 {
											return int64(0)
										} else {
											if v197 == int32(0) {
												if l4 != 0 {
													v201 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v201)
													v323 = int64(0)
													m.G0 = v13 + int32(80)
													return v323
												} else {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v207 = m.ExcPending
													if v207 != 0 {
														return int64(0)
													} else {
														F_errcode(m, int32(67137668))
														mBase = m.M
														v210 = m.ExcPending
														if v210 != 0 {
															return int64(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v192
															F_errmsg(m, int32(_a_F_object_aclmask_ext_7), v13+int32(48))
															mBase = m.M
															v216 = m.ExcPending
															if v216 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_object_aclmask_ext_1), int32(3798), int32(_a_F_object_aclmask_ext_8))
																mBase = m.M
																v221 = m.ExcPending
																if v221 != 0 {
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
												v222 = *(*int32)(unsafe.Add(mBase, uint32(v197)+16))
												v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+22)))
												v225 = v197
												v227 = v222 + v223
												v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+72))
												v233 = F_SysCacheGetAttr(m, int32(82), v225, int32(32), v13+int32(79))
												mBase = m.M
												v234 = m.ExcPending
												if v234 != 0 {
													return int64(0)
												} else {
													v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+79)))
													if v235 == int32(1) {
														v240 = F_acldefault(m, int32(49), v228)
														mBase = m.M
														v241 = m.ExcPending
														if v241 != 0 {
															return int64(0)
														} else {
															v244 = int32(0)
															v245 = v240
															v247 = F_aclmask(m, v245, l2, v228, l3, int32(1))
															mBase = m.M
															v248 = m.ExcPending
															if v248 != 0 {
																return int64(0)
															} else {
																v249 = int32(0)
																if base.B2i32(v245 == v249)|base.B2i32(v245 == v244) == v249 {
																	F_pfree(m, v245)
																	mBase = m.M
																	v256 = m.ExcPending
																	if v256 != 0 {
																		return int64(0)
																	} else {
																		F_ReleaseCatCache(m, v225)
																		mBase = m.M
																		v258 = m.ExcPending
																		if v258 != 0 {
																			return int64(0)
																		} else {
																			v323 = v247
																			m.G0 = v13 + int32(80)
																			return v323
																		}
																	}
																} else {
																	F_ReleaseCatCache(m, v225)
																	mBase = m.M
																	v258 = m.ExcPending
																	if v258 != 0 {
																		return int64(0)
																	} else {
																		v323 = v247
																		m.G0 = v13 + int32(80)
																		return v323
																	}
																}
															}
														}
													} else {
														v242 = F_pg_detoast_datum(m, v233)
														mBase = m.M
														v243 = m.ExcPending
														if v243 != 0 {
															return int64(0)
														} else {
															v244 = v233
															v245 = v242
															v247 = F_aclmask(m, v245, l2, v228, l3, int32(1))
															mBase = m.M
															v248 = m.ExcPending
															if v248 != 0 {
																return int64(0)
															} else {
																v249 = int32(0)
																if base.B2i32(v245 == v249)|base.B2i32(v245 == v244) == v249 {
																	F_pfree(m, v245)
																	mBase = m.M
																	v256 = m.ExcPending
																	if v256 != 0 {
																		return int64(0)
																	} else {
																		F_ReleaseCatCache(m, v225)
																		mBase = m.M
																		v258 = m.ExcPending
																		if v258 != 0 {
																			return int64(0)
																		} else {
																			v323 = v247
																			m.G0 = v13 + int32(80)
																			return v323
																		}
																	}
																} else {
																	F_ReleaseCatCache(m, v225)
																	mBase = m.M
																	v258 = m.ExcPending
																	if v258 != 0 {
																		return int64(0)
																	} else {
																		v323 = v247
																		m.G0 = v13 + int32(80)
																		return v323
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
								v225 = v186
								v227 = v187
								v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+72))
								v233 = F_SysCacheGetAttr(m, int32(82), v225, int32(32), v13+int32(79))
								mBase = m.M
								v234 = m.ExcPending
								if v234 != 0 {
									return int64(0)
								} else {
									v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+79)))
									if v235 == int32(1) {
										v240 = F_acldefault(m, int32(49), v228)
										mBase = m.M
										v241 = m.ExcPending
										if v241 != 0 {
											return int64(0)
										} else {
											v244 = int32(0)
											v245 = v240
											v247 = F_aclmask(m, v245, l2, v228, l3, int32(1))
											mBase = m.M
											v248 = m.ExcPending
											if v248 != 0 {
												return int64(0)
											} else {
												v249 = int32(0)
												if base.B2i32(v245 == v249)|base.B2i32(v245 == v244) == v249 {
													F_pfree(m, v245)
													mBase = m.M
													v256 = m.ExcPending
													if v256 != 0 {
														return int64(0)
													} else {
														F_ReleaseCatCache(m, v225)
														mBase = m.M
														v258 = m.ExcPending
														if v258 != 0 {
															return int64(0)
														} else {
															v323 = v247
															m.G0 = v13 + int32(80)
															return v323
														}
													}
												} else {
													F_ReleaseCatCache(m, v225)
													mBase = m.M
													v258 = m.ExcPending
													if v258 != 0 {
														return int64(0)
													} else {
														v323 = v247
														m.G0 = v13 + int32(80)
														return v323
													}
												}
											}
										}
									} else {
										v242 = F_pg_detoast_datum(m, v233)
										mBase = m.M
										v243 = m.ExcPending
										if v243 != 0 {
											return int64(0)
										} else {
											v244 = v233
											v245 = v242
											v247 = F_aclmask(m, v245, l2, v228, l3, int32(1))
											mBase = m.M
											v248 = m.ExcPending
											if v248 != 0 {
												return int64(0)
											} else {
												v249 = int32(0)
												if base.B2i32(v245 == v249)|base.B2i32(v245 == v244) == v249 {
													F_pfree(m, v245)
													mBase = m.M
													v256 = m.ExcPending
													if v256 != 0 {
														return int64(0)
													} else {
														F_ReleaseCatCache(m, v225)
														mBase = m.M
														v258 = m.ExcPending
														if v258 != 0 {
															return int64(0)
														} else {
															v323 = v247
															m.G0 = v13 + int32(80)
															return v323
														}
													}
												} else {
													F_ReleaseCatCache(m, v225)
													mBase = m.M
													v258 = m.ExcPending
													if v258 != 0 {
														return int64(0)
													} else {
														v323 = v247
														m.G0 = v13 + int32(80)
														return v323
													}
												}
											}
										}
									}
								}
							}
						} else {
							v152 = *(*int32)(unsafe.Add(mBase, uint32(v148)+88))
							if v152 != int32(_a_F_object_aclmask_ext_9) {
								v186 = v121
								v187 = v148
								v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+79)))
								if v188 == int32(109) {
									v191 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
									v192 = F_get_multirange_range(m, v191)
									mBase = m.M
									v193 = m.ExcPending
									if v193 != 0 {
										return int64(0)
									} else {
										F_ReleaseCatCache(m, v186)
										mBase = m.M
										v195 = m.ExcPending
										if v195 != 0 {
											return int64(0)
										} else {
											v197 = F_SearchSysCache1(m, int32(82), v192)
											mBase = m.M
											v198 = m.ExcPending
											if v198 != 0 {
												return int64(0)
											} else {
												if v197 == int32(0) {
													if l4 != 0 {
														v201 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v201)
														v323 = int64(0)
														m.G0 = v13 + int32(80)
														return v323
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v207 = m.ExcPending
														if v207 != 0 {
															return int64(0)
														} else {
															F_errcode(m, int32(67137668))
															mBase = m.M
															v210 = m.ExcPending
															if v210 != 0 {
																return int64(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v192
																F_errmsg(m, int32(_a_F_object_aclmask_ext_7), v13+int32(48))
																mBase = m.M
																v216 = m.ExcPending
																if v216 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(_a_F_object_aclmask_ext_1), int32(3798), int32(_a_F_object_aclmask_ext_8))
																	mBase = m.M
																	v221 = m.ExcPending
																	if v221 != 0 {
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
													v222 = *(*int32)(unsafe.Add(mBase, uint32(v197)+16))
													v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+22)))
													v225 = v197
													v227 = v222 + v223
													v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+72))
													v233 = F_SysCacheGetAttr(m, int32(82), v225, int32(32), v13+int32(79))
													mBase = m.M
													v234 = m.ExcPending
													if v234 != 0 {
														return int64(0)
													} else {
														v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+79)))
														if v235 == int32(1) {
															v240 = F_acldefault(m, int32(49), v228)
															mBase = m.M
															v241 = m.ExcPending
															if v241 != 0 {
																return int64(0)
															} else {
																v244 = int32(0)
																v245 = v240
																v247 = F_aclmask(m, v245, l2, v228, l3, int32(1))
																mBase = m.M
																v248 = m.ExcPending
																if v248 != 0 {
																	return int64(0)
																} else {
																	v249 = int32(0)
																	if base.B2i32(v245 == v249)|base.B2i32(v245 == v244) == v249 {
																		F_pfree(m, v245)
																		mBase = m.M
																		v256 = m.ExcPending
																		if v256 != 0 {
																			return int64(0)
																		} else {
																			F_ReleaseCatCache(m, v225)
																			mBase = m.M
																			v258 = m.ExcPending
																			if v258 != 0 {
																				return int64(0)
																			} else {
																				v323 = v247
																				m.G0 = v13 + int32(80)
																				return v323
																			}
																		}
																	} else {
																		F_ReleaseCatCache(m, v225)
																		mBase = m.M
																		v258 = m.ExcPending
																		if v258 != 0 {
																			return int64(0)
																		} else {
																			v323 = v247
																			m.G0 = v13 + int32(80)
																			return v323
																		}
																	}
																}
															}
														} else {
															v242 = F_pg_detoast_datum(m, v233)
															mBase = m.M
															v243 = m.ExcPending
															if v243 != 0 {
																return int64(0)
															} else {
																v244 = v233
																v245 = v242
																v247 = F_aclmask(m, v245, l2, v228, l3, int32(1))
																mBase = m.M
																v248 = m.ExcPending
																if v248 != 0 {
																	return int64(0)
																} else {
																	v249 = int32(0)
																	if base.B2i32(v245 == v249)|base.B2i32(v245 == v244) == v249 {
																		F_pfree(m, v245)
																		mBase = m.M
																		v256 = m.ExcPending
																		if v256 != 0 {
																			return int64(0)
																		} else {
																			F_ReleaseCatCache(m, v225)
																			mBase = m.M
																			v258 = m.ExcPending
																			if v258 != 0 {
																				return int64(0)
																			} else {
																				v323 = v247
																				m.G0 = v13 + int32(80)
																				return v323
																			}
																		}
																	} else {
																		F_ReleaseCatCache(m, v225)
																		mBase = m.M
																		v258 = m.ExcPending
																		if v258 != 0 {
																			return int64(0)
																		} else {
																			v323 = v247
																			m.G0 = v13 + int32(80)
																			return v323
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
									v225 = v186
									v227 = v187
									v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+72))
									v233 = F_SysCacheGetAttr(m, int32(82), v225, int32(32), v13+int32(79))
									mBase = m.M
									v234 = m.ExcPending
									if v234 != 0 {
										return int64(0)
									} else {
										v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+79)))
										if v235 == int32(1) {
											v240 = F_acldefault(m, int32(49), v228)
											mBase = m.M
											v241 = m.ExcPending
											if v241 != 0 {
												return int64(0)
											} else {
												v244 = int32(0)
												v245 = v240
												v247 = F_aclmask(m, v245, l2, v228, l3, int32(1))
												mBase = m.M
												v248 = m.ExcPending
												if v248 != 0 {
													return int64(0)
												} else {
													v249 = int32(0)
													if base.B2i32(v245 == v249)|base.B2i32(v245 == v244) == v249 {
														F_pfree(m, v245)
														mBase = m.M
														v256 = m.ExcPending
														if v256 != 0 {
															return int64(0)
														} else {
															F_ReleaseCatCache(m, v225)
															mBase = m.M
															v258 = m.ExcPending
															if v258 != 0 {
																return int64(0)
															} else {
																v323 = v247
																m.G0 = v13 + int32(80)
																return v323
															}
														}
													} else {
														F_ReleaseCatCache(m, v225)
														mBase = m.M
														v258 = m.ExcPending
														if v258 != 0 {
															return int64(0)
														} else {
															v323 = v247
															m.G0 = v13 + int32(80)
															return v323
														}
													}
												}
											}
										} else {
											v242 = F_pg_detoast_datum(m, v233)
											mBase = m.M
											v243 = m.ExcPending
											if v243 != 0 {
												return int64(0)
											} else {
												v244 = v233
												v245 = v242
												v247 = F_aclmask(m, v245, l2, v228, l3, int32(1))
												mBase = m.M
												v248 = m.ExcPending
												if v248 != 0 {
													return int64(0)
												} else {
													v249 = int32(0)
													if base.B2i32(v245 == v249)|base.B2i32(v245 == v244) == v249 {
														F_pfree(m, v245)
														mBase = m.M
														v256 = m.ExcPending
														if v256 != 0 {
															return int64(0)
														} else {
															F_ReleaseCatCache(m, v225)
															mBase = m.M
															v258 = m.ExcPending
															if v258 != 0 {
																return int64(0)
															} else {
																v323 = v247
																m.G0 = v13 + int32(80)
																return v323
															}
														}
													} else {
														F_ReleaseCatCache(m, v225)
														mBase = m.M
														v258 = m.ExcPending
														if v258 != 0 {
															return int64(0)
														} else {
															v323 = v247
															m.G0 = v13 + int32(80)
															return v323
														}
													}
												}
											}
										}
									}
								}
							} else {
								F_ReleaseCatCache(m, v121)
								mBase = m.M
								v156 = m.ExcPending
								if v156 != 0 {
									return int64(0)
								} else {
									v158 = F_SearchSysCache1(m, int32(82), v149)
									mBase = m.M
									v159 = m.ExcPending
									if v159 != 0 {
										return int64(0)
									} else {
										if v158 == int32(0) {
											if l4 != 0 {
												v162 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v162)
												v323 = int64(0)
												m.G0 = v13 + int32(80)
												return v323
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v168 = m.ExcPending
												if v168 != 0 {
													return int64(0)
												} else {
													F_errcode(m, int32(67137668))
													mBase = m.M
													v171 = m.ExcPending
													if v171 != 0 {
														return int64(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v149
														F_errmsg(m, int32(_a_F_object_aclmask_ext_7), v13-int32(-64))
														mBase = m.M
														v177 = m.ExcPending
														if v177 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_object_aclmask_ext_1), int32(3769), int32(_a_F_object_aclmask_ext_8))
															mBase = m.M
															v182 = m.ExcPending
															if v182 != 0 {
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
											v183 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
											v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+22)))
											v186 = v158
											v187 = v183 + v184
											v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+79)))
											if v188 == int32(109) {
												v191 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
												v192 = F_get_multirange_range(m, v191)
												mBase = m.M
												v193 = m.ExcPending
												if v193 != 0 {
													return int64(0)
												} else {
													F_ReleaseCatCache(m, v186)
													mBase = m.M
													v195 = m.ExcPending
													if v195 != 0 {
														return int64(0)
													} else {
														v197 = F_SearchSysCache1(m, int32(82), v192)
														mBase = m.M
														v198 = m.ExcPending
														if v198 != 0 {
															return int64(0)
														} else {
															if v197 == int32(0) {
																if l4 != 0 {
																	v201 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v201)
																	v323 = int64(0)
																	m.G0 = v13 + int32(80)
																	return v323
																} else {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v207 = m.ExcPending
																	if v207 != 0 {
																		return int64(0)
																	} else {
																		F_errcode(m, int32(67137668))
																		mBase = m.M
																		v210 = m.ExcPending
																		if v210 != 0 {
																			return int64(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v192
																			F_errmsg(m, int32(_a_F_object_aclmask_ext_7), v13+int32(48))
																			mBase = m.M
																			v216 = m.ExcPending
																			if v216 != 0 {
																				return int64(0)
																			} else {
																				F_errfinish(m, int32(_a_F_object_aclmask_ext_1), int32(3798), int32(_a_F_object_aclmask_ext_8))
																				mBase = m.M
																				v221 = m.ExcPending
																				if v221 != 0 {
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
																v222 = *(*int32)(unsafe.Add(mBase, uint32(v197)+16))
																v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+22)))
																v225 = v197
																v227 = v222 + v223
																v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+72))
																v233 = F_SysCacheGetAttr(m, int32(82), v225, int32(32), v13+int32(79))
																mBase = m.M
																v234 = m.ExcPending
																if v234 != 0 {
																	return int64(0)
																} else {
																	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+79)))
																	if v235 == int32(1) {
																		v240 = F_acldefault(m, int32(49), v228)
																		mBase = m.M
																		v241 = m.ExcPending
																		if v241 != 0 {
																			return int64(0)
																		} else {
																			v244 = int32(0)
																			v245 = v240
																			v247 = F_aclmask(m, v245, l2, v228, l3, int32(1))
																			mBase = m.M
																			v248 = m.ExcPending
																			if v248 != 0 {
																				return int64(0)
																			} else {
																				v249 = int32(0)
																				if base.B2i32(v245 == v249)|base.B2i32(v245 == v244) == v249 {
																					F_pfree(m, v245)
																					mBase = m.M
																					v256 = m.ExcPending
																					if v256 != 0 {
																						return int64(0)
																					} else {
																						F_ReleaseCatCache(m, v225)
																						mBase = m.M
																						v258 = m.ExcPending
																						if v258 != 0 {
																							return int64(0)
																						} else {
																							v323 = v247
																							m.G0 = v13 + int32(80)
																							return v323
																						}
																					}
																				} else {
																					F_ReleaseCatCache(m, v225)
																					mBase = m.M
																					v258 = m.ExcPending
																					if v258 != 0 {
																						return int64(0)
																					} else {
																						v323 = v247
																						m.G0 = v13 + int32(80)
																						return v323
																					}
																				}
																			}
																		}
																	} else {
																		v242 = F_pg_detoast_datum(m, v233)
																		mBase = m.M
																		v243 = m.ExcPending
																		if v243 != 0 {
																			return int64(0)
																		} else {
																			v244 = v233
																			v245 = v242
																			v247 = F_aclmask(m, v245, l2, v228, l3, int32(1))
																			mBase = m.M
																			v248 = m.ExcPending
																			if v248 != 0 {
																				return int64(0)
																			} else {
																				v249 = int32(0)
																				if base.B2i32(v245 == v249)|base.B2i32(v245 == v244) == v249 {
																					F_pfree(m, v245)
																					mBase = m.M
																					v256 = m.ExcPending
																					if v256 != 0 {
																						return int64(0)
																					} else {
																						F_ReleaseCatCache(m, v225)
																						mBase = m.M
																						v258 = m.ExcPending
																						if v258 != 0 {
																							return int64(0)
																						} else {
																							v323 = v247
																							m.G0 = v13 + int32(80)
																							return v323
																						}
																					}
																				} else {
																					F_ReleaseCatCache(m, v225)
																					mBase = m.M
																					v258 = m.ExcPending
																					if v258 != 0 {
																						return int64(0)
																					} else {
																						v323 = v247
																						m.G0 = v13 + int32(80)
																						return v323
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
												v225 = v186
												v227 = v187
												v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+72))
												v233 = F_SysCacheGetAttr(m, int32(82), v225, int32(32), v13+int32(79))
												mBase = m.M
												v234 = m.ExcPending
												if v234 != 0 {
													return int64(0)
												} else {
													v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+79)))
													if v235 == int32(1) {
														v240 = F_acldefault(m, int32(49), v228)
														mBase = m.M
														v241 = m.ExcPending
														if v241 != 0 {
															return int64(0)
														} else {
															v244 = int32(0)
															v245 = v240
															v247 = F_aclmask(m, v245, l2, v228, l3, int32(1))
															mBase = m.M
															v248 = m.ExcPending
															if v248 != 0 {
																return int64(0)
															} else {
																v249 = int32(0)
																if base.B2i32(v245 == v249)|base.B2i32(v245 == v244) == v249 {
																	F_pfree(m, v245)
																	mBase = m.M
																	v256 = m.ExcPending
																	if v256 != 0 {
																		return int64(0)
																	} else {
																		F_ReleaseCatCache(m, v225)
																		mBase = m.M
																		v258 = m.ExcPending
																		if v258 != 0 {
																			return int64(0)
																		} else {
																			v323 = v247
																			m.G0 = v13 + int32(80)
																			return v323
																		}
																	}
																} else {
																	F_ReleaseCatCache(m, v225)
																	mBase = m.M
																	v258 = m.ExcPending
																	if v258 != 0 {
																		return int64(0)
																	} else {
																		v323 = v247
																		m.G0 = v13 + int32(80)
																		return v323
																	}
																}
															}
														}
													} else {
														v242 = F_pg_detoast_datum(m, v233)
														mBase = m.M
														v243 = m.ExcPending
														if v243 != 0 {
															return int64(0)
														} else {
															v244 = v233
															v245 = v242
															v247 = F_aclmask(m, v245, l2, v228, l3, int32(1))
															mBase = m.M
															v248 = m.ExcPending
															if v248 != 0 {
																return int64(0)
															} else {
																v249 = int32(0)
																if base.B2i32(v245 == v249)|base.B2i32(v245 == v244) == v249 {
																	F_pfree(m, v245)
																	mBase = m.M
																	v256 = m.ExcPending
																	if v256 != 0 {
																		return int64(0)
																	} else {
																		F_ReleaseCatCache(m, v225)
																		mBase = m.M
																		v258 = m.ExcPending
																		if v258 != 0 {
																			return int64(0)
																		} else {
																			v323 = v247
																			m.G0 = v13 + int32(80)
																			return v323
																		}
																	}
																} else {
																	F_ReleaseCatCache(m, v225)
																	mBase = m.M
																	v258 = m.ExcPending
																	if v258 != 0 {
																		return int64(0)
																	} else {
																		v323 = v247
																		m.G0 = v13 + int32(80)
																		return v323
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
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
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
	return v89
L17:
	;
	F_pfree(m, v37)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L13
	} else {
		goto L55
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
		v85 = v38
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
	v89 = v38
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
		v85 = v50
		goto L17
	} else {
		goto L32
	}
L32:
	;
	v89 = v50
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
		v85 = v51
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
	v89 = v51
	goto L16
L38:
	;
	if v37 == int32(0) {
		v89 = v77
		goto L16
	} else {
		goto L53
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
		v77 = v57
		goto L38
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	if v68|base.B2i32(v22 == int32(0)) != 0 {
		v77 = v68
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
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v75 = m.T0[v22].(func(*base.Module, int32, int32, int32) int32)(m, v72, v37, base.B2i32(v53 == int32(11)))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L13
	} else {
		goto L52
	}
L52:
	;
	v77 = v75
	goto L38
L53:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v80&int32(4) == int32(0) {
		v89 = v77
		goto L16
	} else {
		goto L54
	}
L54:
	;
	v85 = v77
	goto L17
L55:
	;
	v89 = v85
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
								F_errmsg_internal(m, int32(_a_F_storeObjectDescription_0), v8)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_storeObjectDescription_1), int32(1322), int32(_a_F_storeObjectDescription_2))
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
							F_appendStringInfo(m, l0, int32(_a_F_storeObjectDescription_3), v8+int32(112))
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
							F_appendStringInfo(m, l0, int32(_a_F_storeObjectDescription_4), v8+int32(48))
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
								F_errmsg_internal(m, int32(_a_F_storeObjectDescription_5), v8+int32(16))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_storeObjectDescription_1), int32(1310), int32(_a_F_storeObjectDescription_2))
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
							F_appendStringInfo(m, l0, int32(_a_F_storeObjectDescription_6), v8-int32(-64))
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
							F_appendStringInfo(m, l0, int32(_a_F_storeObjectDescription_7), v8+int32(32))
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
							F_appendStringInfo(m, l0, int32(_a_F_storeObjectDescription_8), v8+int32(80))
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
							F_appendStringInfo(m, l0, int32(_a_F_storeObjectDescription_9), v8+int32(96))
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
							F_errmsg_internal(m, int32(_a_F_storeObjectDescription_0), v8)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_storeObjectDescription_1), int32(1322), int32(_a_F_storeObjectDescription_2))
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
						F_appendStringInfo(m, l0, int32(_a_F_storeObjectDescription_3), v8+int32(112))
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
						F_appendStringInfo(m, l0, int32(_a_F_storeObjectDescription_4), v8+int32(48))
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
							F_errmsg_internal(m, int32(_a_F_storeObjectDescription_5), v8+int32(16))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_storeObjectDescription_1), int32(1310), int32(_a_F_storeObjectDescription_2))
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
						F_appendStringInfo(m, l0, int32(_a_F_storeObjectDescription_6), v8-int32(-64))
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
						F_appendStringInfo(m, l0, int32(_a_F_storeObjectDescription_7), v8+int32(32))
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
						F_appendStringInfo(m, l0, int32(_a_F_storeObjectDescription_8), v8+int32(80))
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
						F_appendStringInfo(m, l0, int32(_a_F_storeObjectDescription_9), v8+int32(96))
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
