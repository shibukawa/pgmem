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
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
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
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v227 float64
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	if l0 == v2 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L17
	} else {
		goto L86
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L17
	} else {
		goto L83
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L17
	} else {
		goto L80
	}
L4:
	;
	m.G0 = v12 - int32(-64)
	return v255
L5:
	;
	v255 = v2
	goto L4
L6:
	;
	goto L7
L7:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v16 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v95 = F_palloc0(m, int32(12))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L17
	} else {
		goto L32
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	if base.Ui32(int32(11)) < base.Ui32(v42) {
		goto L8
	} else {
		goto L16
	}
L11:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if base.Ui32((v19-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v30 = int32(1)
	if v16&v30 != 0 {
		v42 = int32(base.Ui32(v16)>>(uint(v30)%32)) - v30
		goto L10
	} else {
		goto L15
	}
L14:
	;
	v42 = base.B2i32(v19 == int32(18)) << (uint(int32(4)) % 32)
	goto L10
L15:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v42 = int32(base.Ui32(v36)>>(uint(int32(2))%32)) - int32(4)
	goto L10
L16:
	;
	goto L9
L17:
	;
	return int32(0)
L18:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v52 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v82
	F_errmsg_internal(m, int32(621201), v12)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L17
	} else {
		goto L30
	}
L20:
	;
	v55 = int32(4)
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v57&int32(254) == int32(2) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v70 = int32(1)
	if v52&v70 != 0 {
		v82 = int32(base.Ui32(v52)>>(uint(v70)%32)) - v70
		goto L19
	} else {
		goto L29
	}
L23:
	;
	v66 = v55
	goto L25
L24:
	;
	v66 = base.B2i32(v57 == int32(18)) << (uint(v55) % 32)
	goto L25
L25:
	;
	if v57 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v69 = v55
	goto L28
L27:
	;
	v69 = v66
	goto L28
L28:
	;
	v82 = v69
	goto L19
L29:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v82 = int32(base.Ui32(v76)>>(uint(int32(2))%32)) - int32(4)
	goto L19
L30:
	;
	F_errfinish(m, int32(463877), int32(511), int32(320551))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L17
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	v97 = int32(1)
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v101 = v99 & v97
	if v101 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v102 = v97
	goto L35
L34:
	;
	v102 = int32(4)
	goto L35
L35:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0+v102)))
	*(*int32)(unsafe.Add(mBase, uint32(v95))) = v104
	v107 = l0 + int32(1)
	if v101 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v110 = v107
	goto L38
L37:
	;
	v110 = l0 + int32(4)
	goto L38
L38:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+4)) = v111
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v110)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+8)) = v113
	if v104 != int32(-1269523924) {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	if v111 != int32(1) {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	if v113 == int32(0) {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v121 = int32(1)
	v124 = v113<<(uint(v121)%32) + int32(10)
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v125 == v121 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	if base.Ui32(v155) < base.Ui32(v124) {
		goto L53
	} else {
		goto L54
	}
L43:
	;
	v128 = int32(4)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v130&int32(254) == int32(2) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	v143 = int32(1)
	if v125&v143 != 0 {
		v155 = int32(base.Ui32(v125)>>(uint(v143)%32)) - v143
		goto L42
	} else {
		goto L52
	}
L46:
	;
	v139 = v128
	goto L48
L47:
	;
	v139 = base.B2i32(v130 == int32(18)) << (uint(v128) % 32)
	goto L48
L48:
	;
	if v130 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v142 = v128
	goto L51
L50:
	;
	v142 = v139
	goto L51
L51:
	;
	v155 = v142
	goto L42
L52:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v155 = int32(base.Ui32(v149)>>(uint(int32(2))%32)) - int32(4)
	goto L42
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L17
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v208 = F_repalloc(m, v95, v113<<(uint(int32(2))%32)+int32(12))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L17
	} else {
		goto L70
	}
L56:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v161 == int32(1) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v191
	F_errmsg_internal(m, int32(621147), v10+int32(-48))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L17
	} else {
		goto L68
	}
L58:
	;
	v164 = int32(4)
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v166&int32(254) == int32(2) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	goto L60
L60:
	;
	v179 = int32(1)
	if v161&v179 != 0 {
		v191 = int32(base.Ui32(v161)>>(uint(v179)%32)) - v179
		goto L57
	} else {
		goto L67
	}
L61:
	;
	v175 = v164
	goto L63
L62:
	;
	v175 = base.B2i32(v166 == int32(18)) << (uint(v164) % 32)
	goto L63
L63:
	;
	if v166 == int32(1) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v178 = v164
	goto L66
L65:
	;
	v178 = v175
	goto L66
L66:
	;
	v191 = v178
	goto L57
L67:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v191 = int32(base.Ui32(v185)>>(uint(int32(2))%32)) - int32(4)
	goto L57
L68:
	;
	F_errfinish(m, int32(463877), int32(543), int32(320551))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L17
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v208)+8))
	if v210 == int32(0) {
		v255 = v208
		goto L4
	} else {
		goto L71
	}
L71:
	;
	v213 = int32(12)
	v218 = v110 + v213
	v222 = int32(0)
	goto L72
L72:
	;
	v227 = *(*float64)(unsafe.Add(mBase, uint32(v218)))
	v228 = int32(*(*int16)(unsafe.Add(mBase, uint32(v218)+8)))
	v230 = v228 << (uint(int32(1)) % 32)
	v233 = F_palloc0(m, v230+int32(10))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L17
	} else {
		goto L74
	}
L73:
	;
	v255 = v208
	goto L4
L74:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v233)+8)) = uint16(v228)
	*(*float64)(unsafe.Add(mBase, uint32(v233))) = v227
	v237 = int32(10)
	v240 = v218 + v237
	if v230 != 0 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208+v213+v222<<(uint(int32(2))%32)))) = v233
	v249 = v222 + int32(1)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v208)+8))
	if base.Ui32(v249) < base.Ui32(v250) {
		v218 = v240 + v230
		v222 = v249
		goto L72
	} else {
		goto L79
	}
L76:
	;
	v241 = F__emscripten_memcpy_bulkmem(m, v233+v237, v240, v230)
	mBase = m.M
	goto L78
L77:
	;
	goto L78
L78:
	;
	goto L75
L79:
	;
	goto L73
L80:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = int32(-1269523924)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v269
	F_errmsg_internal(m, int32(631285), v10+int32(-16))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L17
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(463877), int32(529), int32(320551))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L17
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
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v287
	F_errmsg_internal(m, int32(631204), v10+int32(-32))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L17
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(463877), int32(533), int32(320551))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L17
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
	F_errmsg_internal(m, int32(157016), int32(0))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L17
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(463877), int32(536), int32(320551))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L17
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
						F_errmsg_internal(m, int32(38778), v7+int32(16))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(462360), int32(575), int32(435930))
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
				F_errmsg_internal(m, int32(38897), v7)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(462360), int32(567), int32(435930))
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
