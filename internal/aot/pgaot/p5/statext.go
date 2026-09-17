package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_statext_dependencies_deserialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v225 float64
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	if l0 == v2 {
		v250 = v2
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L18
	} else {
		goto L86
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L18
	} else {
		goto L83
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L18
	} else {
		goto L80
	}
L4:
	;
	m.G0 = v12 - int32(-64)
	return v250
L5:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v16 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v95 = F_palloc0(m, int32(12))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L18
	} else {
		goto L33
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L18
	} else {
		goto L19
	}
L8:
	;
	if base.Ui32(int32(11)) < base.Ui32(v43) {
		goto L6
	} else {
		goto L17
	}
L9:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if base.Ui32((v19-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v31 = int32(1)
	if v16&v31 != 0 {
		v43 = int32(base.Ui32(v16)>>(uint(v31)%32)) - v31
		goto L8
	} else {
		goto L16
	}
L12:
	;
	if v19 == int32(18) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v30 = int32(16)
	goto L15
L14:
	;
	v30 = int32(0)
	goto L15
L15:
	;
	v43 = v30
	goto L8
L16:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v43 = int32(base.Ui32(v37)>>(uint(int32(2))%32)) - int32(4)
	goto L8
L17:
	;
	goto L7
L18:
	;
	return int32(0)
L19:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v53 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v82
	F_errmsg_internal(m, int32(_a_F_statext_dependencies_deserialize_0), v12)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L18
	} else {
		goto L31
	}
L21:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v59 == int32(18) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v70 = int32(1)
	if v53&v70 != 0 {
		v82 = int32(base.Ui32(v53)>>(uint(v70)%32)) - v70
		goto L20
	} else {
		goto L30
	}
L24:
	;
	v62 = int32(16)
	goto L26
L25:
	;
	v62 = int32(0)
	goto L26
L26:
	;
	if base.Ui32((v59-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v69 = int32(4)
	goto L29
L28:
	;
	v69 = v62
	goto L29
L29:
	;
	v82 = v69
	goto L20
L30:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v82 = int32(base.Ui32(v76)>>(uint(int32(2))%32)) - int32(4)
	goto L20
L31:
	;
	F_errfinish(m, int32(_a_F_statext_dependencies_deserialize_1), int32(511), int32(_a_F_statext_dependencies_deserialize_2))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L18
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	v97 = int32(1)
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v101 = v99 & v97
	if v101 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v102 = v97
	goto L36
L35:
	;
	v102 = int32(4)
	goto L36
L36:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0+v102)))
	*(*int32)(unsafe.Add(mBase, uint32(v95))) = v104
	v107 = l0 + int32(1)
	if v101 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v110 = v107
	goto L39
L38:
	;
	v110 = l0 + int32(4)
	goto L39
L39:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+4)) = v111
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v110)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+8)) = v113
	if v104 != int32(-1269523924) {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	if v111 != int32(1) {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	if v113 == int32(0) {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v121 = int32(1)
	v124 = v113<<(uint(v121)%32) + int32(10)
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v125 == v121 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if base.Ui32(v154) < base.Ui32(v124) {
		goto L54
	} else {
		goto L55
	}
L44:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v131 == int32(18) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	v142 = int32(1)
	if v125&v142 != 0 {
		v154 = int32(base.Ui32(v125)>>(uint(v142)%32)) - v142
		goto L43
	} else {
		goto L53
	}
L47:
	;
	v134 = int32(16)
	goto L49
L48:
	;
	v134 = int32(0)
	goto L49
L49:
	;
	if base.Ui32((v131-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v141 = int32(4)
	goto L52
L51:
	;
	v141 = v134
	goto L52
L52:
	;
	v154 = v141
	goto L43
L53:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v154 = int32(base.Ui32(v148)>>(uint(int32(2))%32)) - int32(4)
	goto L43
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L18
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v206 = F_repalloc(m, v95, v113<<(uint(int32(2))%32)+int32(12))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L18
	} else {
		goto L71
	}
L57:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v160 == int32(1) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v189
	F_errmsg_internal(m, int32(_a_F_statext_dependencies_deserialize_3), v10+int32(-48))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L18
	} else {
		goto L69
	}
L59:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v166 == int32(18) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	goto L61
L61:
	;
	v177 = int32(1)
	if v160&v177 != 0 {
		v189 = int32(base.Ui32(v160)>>(uint(v177)%32)) - v177
		goto L58
	} else {
		goto L68
	}
L62:
	;
	v169 = int32(16)
	goto L64
L63:
	;
	v169 = int32(0)
	goto L64
L64:
	;
	if base.Ui32((v166-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v176 = int32(4)
	goto L67
L66:
	;
	v176 = v169
	goto L67
L67:
	;
	v189 = v176
	goto L58
L68:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v189 = int32(base.Ui32(v183)>>(uint(int32(2))%32)) - int32(4)
	goto L58
L69:
	;
	F_errfinish(m, int32(_a_F_statext_dependencies_deserialize_1), int32(543), int32(_a_F_statext_dependencies_deserialize_2))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L18
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v206)+8))
	if v208 == int32(0) {
		v250 = v206
		goto L4
	} else {
		goto L72
	}
L72:
	;
	v211 = int32(12)
	v216 = v110 + v211
	v219 = int32(0)
	goto L73
L73:
	;
	v225 = *(*float64)(unsafe.Add(mBase, uint32(v216)))
	v226 = int32(*(*int16)(unsafe.Add(mBase, uint32(v216)+8)))
	v228 = v226 << (uint(int32(1)) % 32)
	v231 = F_palloc0(m, v228+int32(10))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L18
	} else {
		goto L75
	}
L74:
	;
	v250 = v206
	goto L4
L75:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v231)+8)) = uint16(v226)
	*(*float64)(unsafe.Add(mBase, uint32(v231))) = v225
	v236 = v216 + int32(10)
	if v228 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	base.MemoryCopy(m, v231+int32(10), v236, v228)
	goto L78
L77:
	;
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v206+v211+v219<<(uint(int32(2))%32)))) = v231
	v246 = v219 + int32(1)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v206)+8))
	if base.Ui32(v246) < base.Ui32(v247) {
		v216 = v236 + v228
		v219 = v246
		goto L73
	} else {
		goto L79
	}
L79:
	;
	goto L74
L80:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = int32(-1269523924)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v266
	F_errmsg_internal(m, int32(_a_F_statext_dependencies_deserialize_4), v10+int32(-16))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L18
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_statext_dependencies_deserialize_1), int32(529), int32(_a_F_statext_dependencies_deserialize_2))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L18
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v284
	F_errmsg_internal(m, int32(_a_F_statext_dependencies_deserialize_5), v10+int32(-32))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L18
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_statext_dependencies_deserialize_1), int32(533), int32(_a_F_statext_dependencies_deserialize_2))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L18
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	F_errmsg_internal(m, int32(_a_F_statext_dependencies_deserialize_6), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L18
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_statext_dependencies_deserialize_1), int32(536), int32(_a_F_statext_dependencies_deserialize_2))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L18
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_statext_mcv_load(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = F_SearchSysCache2(m, int32(62), l0, l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 != 0 {
			v18 = F_SysCacheGetAttr(m, int32(62), v10, int32(5), v7+int32(31))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+31)))
				if v20 == int32(1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(109)
						F_errmsg_internal(m, int32(_a_F_statext_mcv_load_0), v7+int32(16))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_statext_mcv_load_1), int32(575), int32(_a_F_statext_mcv_load_2))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v23 = F_pg_detoast_datum(m, v18)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = F_statext_mcv_deserialize(m, v23)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							F_ReleaseCatCache(m, v10)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								m.G0 = v7 + int32(32)
								return v25
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
				F_errmsg_internal(m, int32(_a_F_statext_mcv_load_3), v7)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_statext_mcv_load_1), int32(567), int32(_a_F_statext_mcv_load_2))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
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
}
