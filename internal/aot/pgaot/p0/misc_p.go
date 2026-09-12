package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ParseAbortRecord(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 int64
	_ = v265
	var v266 int64
	_ = v266
	v10 = F__emscripten_memset_bulkmem(m, l2, base.I32_extend8_s(int32(0)), int32(264))
	mBase = m.M
	goto L1
L1:
	;
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v11
	if int32(0) <= base.I32_extend8_s(l0) {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v16
	if v16&int32(1) != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v22
	v28 = l1 + int32(20)
	goto L6
L5:
	;
	v28 = l1 + int32(12)
	goto L6
L6:
	;
	if v16&int32(2) != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v33 = v28 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v31
	v39 = v33 + v31<<(uint(int32(2))%32)
	goto L9
L8:
	;
	v39 = v28
	goto L9
L9:
	;
	if v16&int32(4) != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v45 = v39 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v43
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v52 = v45 + v48*int32(12)
	goto L12
L11:
	;
	v52 = v39
	goto L12
L12:
	;
	if v16&int32(256) != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v58 = int32(4)
	v59 = v52 + v58
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v57
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v66 = v59 + v62<<(uint(v58)%32)
	goto L15
L14:
	;
	v66 = v52
	goto L15
L15:
	;
	if v16&int32(16) == int32(0) {
		v259 = v16
		v260 = v66
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if v259&int32(32) == int32(0) {
		goto L2
	} else {
		goto L68
	}
L17:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v73
	v76 = v66 + int32(4)
	if v16&int32(128) == int32(0) {
		v259 = v16
		v260 = v76
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v82 = v10 + int32(48)
	goto L22
L19:
	;
	if v76&int32(3) == int32(0) {
		v221 = v76
		goto L53
	} else {
		goto L54
	}
L20:
	;
	v195 = F_strlen(m, v184)
	mBase = m.M
	goto L19
L22:
	;
	goto L23
L23:
	;
	v89 = int32(199)
	if (v82^v76)&int32(3) != 0 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v188 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v185))) = uint8(v188)
	goto L20
L25:
	;
	v169 = v164
	v170 = v165
	v171 = v166
	goto L47
L26:
	;
	if v159 == int32(0) {
		v184 = v157
		v185 = v158
		goto L24
	} else {
		goto L46
	}
L27:
	;
	v157 = v76
	v158 = v82
	v159 = v89
	goto L26
L28:
	;
	goto L29
L29:
	;
	if v76&int32(3) == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	if v126 == int32(0) {
		v184 = v123
		v185 = v124
		goto L24
	} else {
		goto L39
	}
L31:
	;
	v123 = v76
	v124 = v82
	v125 = v89
	v126 = int32(1)
	goto L30
L32:
	;
	goto L33
L33:
	;
	v102 = v76
	v103 = v82
	v104 = v89
	goto L34
L34:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	*(*uint8)(unsafe.Add(mBase, uint32(v103))) = uint8(v106)
	if v106 == int32(0) {
		v164 = v102
		v165 = v103
		v166 = v104
		goto L25
	} else {
		goto L36
	}
L35:
	;
	v123 = v117
	v124 = v111
	v125 = v113
	v126 = v115
	goto L30
L36:
	;
	v110 = int32(1)
	v111 = v103 + v110
	v113 = v104 - v110
	v114 = int32(0)
	v115 = base.B2i32(v113 != v114)
	v117 = v102 + v110
	if v117&int32(3) == v114 {
		v123 = v117
		v124 = v111
		v125 = v113
		v126 = v115
		goto L30
	} else {
		goto L37
	}
L37:
	;
	if v113 != 0 {
		v102 = v117
		v103 = v111
		v104 = v113
		goto L34
	} else {
		goto L38
	}
L38:
	;
	goto L35
L39:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	if v129 == int32(0) {
		v157 = v123
		v158 = v124
		v159 = v125
		goto L26
	} else {
		goto L40
	}
L40:
	;
	if base.Ui32(v125) < base.Ui32(int32(4)) {
		v157 = v123
		v158 = v124
		v159 = v125
		goto L26
	} else {
		goto L41
	}
L41:
	;
	v135 = v123
	v136 = v124
	v137 = v125
	goto L42
L42:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	v143 = int32(-2139062144)
	if (int32(16843008)-v140|v140)&v143 != v143 {
		v164 = v135
		v165 = v136
		v166 = v137
		goto L25
	} else {
		goto L44
	}
L43:
	;
	v157 = v151
	v158 = v149
	v159 = v153
	goto L26
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = v140
	v148 = int32(4)
	v149 = v136 + v148
	v151 = v135 + v148
	v153 = v137 - v148
	if base.Ui32(int32(3)) < base.Ui32(v153) {
		v135 = v151
		v136 = v149
		v137 = v153
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v164 = v157
	v165 = v158
	v166 = v159
	goto L25
L47:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v173)
	if v173 == int32(0) {
		v184 = v169
		v185 = v170
		goto L24
	} else {
		goto L49
	}
L48:
	;
	v184 = v180
	v185 = v178
	goto L24
L49:
	;
	v177 = int32(1)
	v178 = v170 + v177
	v180 = v169 + v177
	v182 = v171 - v177
	if v182 != 0 {
		v169 = v180
		v170 = v178
		v171 = v182
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v259 = v258
	v260 = v254 + v76 + int32(1)
	goto L16
L52:
	;
	v254 = v246 - v76
	goto L51
L53:
	;
	v225 = v221
	goto L62
L54:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v205 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v254 = int32(0)
	goto L51
L56:
	;
	goto L57
L57:
	;
	v210 = v76
	goto L58
L58:
	;
	v214 = v210 + int32(1)
	if v214&int32(3) == int32(0) {
		v221 = v214
		goto L53
	} else {
		goto L60
	}
L59:
	;
	v246 = v214
	goto L52
L60:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214))))
	if v219 != 0 {
		v210 = v214
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	v234 = int32(-2139062144)
	if (int32(16843008)-v231|v231)&v234 == v234 {
		v225 = v225 + int32(4)
		goto L62
	} else {
		goto L64
	}
L63:
	;
	v240 = v225
	goto L65
L64:
	;
	goto L63
L65:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240))))
	if v244 != 0 {
		v240 = v240 + int32(1)
		goto L65
	} else {
		goto L67
	}
L66:
	;
	v246 = v240
	goto L52
L67:
	;
	goto L66
L68:
	;
	v265 = *(*int64)(unsafe.Add(mBase, uint32(v260)))
	v266 = *(*int64)(unsafe.Add(mBase, uint32(v260)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+256)) = v266
	*(*int64)(unsafe.Add(mBase, uint32(v10)+248)) = v265
	goto L2
}
func F_PrepareInvalidationState(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v51 int64
	_ = v51
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	v5 = *(*int32)(unsafe.Add(mBase, _consts[875]))
	if v5 == int32(0) {
		v17 = *(*int32)(unsafe.Add(mBase, _consts[76]))
		v19 = F_MemoryContextAllocZero(m, v17, int32(44))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, _consts[875]))
			*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v24
			v27 = *(*int32)(unsafe.Add(mBase, _consts[37]))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
			*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v28
			v31 = *(*int32)(unsafe.Add(mBase, _consts[875]))
			if v31 != 0 {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
				v37 = v31 + int32(12)
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
				if v32-v33 != v35-v38 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(168566), int32(0))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(492080), int32(717), int32(349815))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v32
					*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v32
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
					*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v43
					*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v43
					*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v32
					*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v43
					*(*int32)(unsafe.Add(mBase, uint32(v19))) = v32
					*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v43
					*(*int32)(unsafe.Add(mBase, _consts[875])) = v19
					return v19
				}
			} else {
				v51 = int64(0)
				*(*int64)(unsafe.Add(mBase, _consts[877])) = v51
				*(*int64)(unsafe.Add(mBase, _consts[876])) = v51
				*(*int32)(unsafe.Add(mBase, _consts[875])) = v19
				return v19
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)+40))
		v10 = *(*int32)(unsafe.Add(mBase, _consts[37]))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
		if v8 != v11 {
			v17 = *(*int32)(unsafe.Add(mBase, _consts[76]))
			v19 = F_MemoryContextAllocZero(m, v17, int32(44))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, _consts[875]))
				*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v24
				v27 = *(*int32)(unsafe.Add(mBase, _consts[37]))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
				*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v28
				v31 = *(*int32)(unsafe.Add(mBase, _consts[875]))
				if v31 != 0 {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
					v37 = v31 + int32(12)
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
					if v32-v33 != v35-v38 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(168566), int32(0))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(492080), int32(717), int32(349815))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v32
						*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v32
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
						*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v43
						*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v43
						*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v32
						*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v43
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v32
						*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v43
						*(*int32)(unsafe.Add(mBase, _consts[875])) = v19
						return v19
					}
				} else {
					v51 = int64(0)
					*(*int64)(unsafe.Add(mBase, _consts[877])) = v51
					*(*int64)(unsafe.Add(mBase, _consts[876])) = v51
					*(*int32)(unsafe.Add(mBase, _consts[875])) = v19
					return v19
				}
			}
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, _consts[875]))
			return v14
		}
	}
}
func F_PrepareSortSupportFromGistIndexRel(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+84))
	if v11 == int32(783) {
		v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+10)))
		v18 = v14<<(uint(int32(2))%32) - int32(4)
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v18+v19)))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v22+v18)))
		v25 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)) = uint8(v25)
		v28 = F_get_opfamily_proc(m, v24, v21, v21, int32(11))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return
		} else {
			if v28 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v24
					*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v21
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v21
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(11)
					F_errmsg_internal(m, int32(39469), v8)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						F_errfinish(m, int32(487450), int32(205), int32(305128))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v33 = F_OidFunctionCall1Coll(m, v28, int32(0), l1)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					m.G0 = v8 + int32(32)
					return
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return
		} else {
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+84))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v43
			F_errmsg_internal(m, int32(58757), v8+int32(16))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return
			} else {
				F_errfinish(m, int32(487450), int32(194), int32(305128))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
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
func F_p_isdigit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v3 == int32(1) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v6 != 0 {
			v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v6+v8<<(uint(int32(2))%32))))
			return base.B2i32(base.Ui32(v12-int32(48)) < base.Ui32(int32(10)))
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v18+v20<<(uint(int32(2))%32))))
			return base.B2i32(base.Ui32(v24-int32(48)) < base.Ui32(int32(10)))
		}
	} else {
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
		v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+v32))))
		return base.B2i32(base.Ui32((v34-int32(48))&int32(255)) < base.Ui32(int32(10)))
	}
}
func F_pairingheap_allocate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_palloc(m, int32(12))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
		return v5
	}
}
func F_pairingheap_remove_first(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
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
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v11 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	return v9
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v75
	v83 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v75)+8)) = v83
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = v83
	return v9
L5:
	;
	v75 = v10
	goto L4
L6:
	;
	goto L7
L7:
	;
	v15 = int32(0)
	v16 = v10
	goto L8
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v22 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if v15 == int32(0) {
		v75 = v44
		goto L4
	} else {
		goto L26
	}
L10:
	;
	goto L9
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v15
	v44 = v16
	goto L10
L12:
	;
	goto L13
L13:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = m.T0[v28].(func(*base.Module, int32, int32, int32) int32)(m, v16, v22, v27)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	v34 = base.B2i32(v29 < int32(0))
	if v29 < int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v35 = v16
	goto L18
L17:
	;
	v35 = v22
	goto L18
L18:
	;
	if v29 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v36 = v22
	goto L21
L20:
	;
	v36 = v16
	goto L21
L21:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v37 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = v35
	goto L24
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v36
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v35
	if v26 != 0 {
		v15 = v36
		v16 = v26
		goto L8
	} else {
		goto L25
	}
L25:
	;
	v44 = v36
	goto L10
L26:
	;
	v52 = v44
	v54 = v15
	goto L27
L27:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v62 = m.T0[v61].(func(*base.Module, int32, int32, int32) int32)(m, v52, v54, v60)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L14
	} else {
		goto L29
	}
L28:
	;
	v75 = v67
	goto L4
L29:
	;
	v65 = base.B2i32(v62 < int32(0))
	if v62 < int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v66 = v52
	goto L32
L31:
	;
	v66 = v54
	goto L32
L32:
	;
	if v62 < int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v67 = v54
	goto L35
L34:
	;
	v67 = v52
	goto L35
L35:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v68 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68)+8)) = v66
	goto L38
L37:
	;
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66)+8)) = v67
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v66
	if v59 != 0 {
		v52 = v67
		v54 = v59
		goto L27
	} else {
		goto L39
	}
L39:
	;
	goto L28
}
func F_parse_dispatch_option(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	v2 = int32(315003)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, _consts[374])))
	if v6 == int32(0) {
		v25 = v5
		v26 = v6
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v26-v25 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	goto L1
L3:
	;
	if v5 != v6 {
		v25 = v5
		v26 = v6
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v10 = v2
	v11 = l0
	goto L5
L5:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v15 == int32(0) {
		v25 = v14
		v26 = v15
		goto L2
	} else {
		goto L7
	}
L6:
	;
	v25 = v14
	v26 = v15
	goto L2
L7:
	;
	v18 = int32(1)
	if v14 == v15 {
		v10 = v10 + v18
		v11 = v11 + v18
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	return int32(0)
L10:
	;
	goto L11
L11:
	;
	v32 = int32(83844)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, _consts[375])))
	if v36 == int32(0) {
		v55 = v35
		v56 = v36
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v56-v55 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L13:
	;
	goto L12
L14:
	;
	if v35 != v36 {
		v55 = v35
		v56 = v36
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v40 = v32
	v41 = l0
	goto L16
L16:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)))
	if v45 == int32(0) {
		v55 = v44
		v56 = v45
		goto L13
	} else {
		goto L18
	}
L17:
	;
	v55 = v44
	v56 = v45
	goto L13
L18:
	;
	v48 = int32(1)
	if v44 == v45 {
		v40 = v40 + v48
		v41 = v41 + v48
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	return int32(1)
L21:
	;
	goto L22
L22:
	;
	v62 = int32(333768)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, _consts[376])))
	if v66 == int32(0) {
		v85 = v65
		v86 = v66
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v86-v85 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L24:
	;
	goto L23
L25:
	;
	if v65 != v66 {
		v85 = v65
		v86 = v66
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v70 = v62
	v71 = l0
	goto L27
L27:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	if v75 == int32(0) {
		v85 = v74
		v86 = v75
		goto L24
	} else {
		goto L29
	}
L28:
	;
	v85 = v74
	v86 = v75
	goto L24
L29:
	;
	v78 = int32(1)
	if v74 == v75 {
		v70 = v70 + v78
		v71 = v71 + v78
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	return int32(3)
L32:
	;
	goto L33
L33:
	;
	v94 = int32(385949)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, _consts[377])))
	if v98 == int32(0) {
		v117 = v97
		v118 = v98
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v118-v117 != 0 {
		goto L42
	} else {
		goto L43
	}
L35:
	;
	goto L34
L36:
	;
	if v97 != v98 {
		v117 = v97
		v118 = v98
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v102 = v94
	v103 = l0
	goto L38
L38:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+1)))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+1)))
	if v107 == int32(0) {
		v117 = v106
		v118 = v107
		goto L35
	} else {
		goto L40
	}
L39:
	;
	v117 = v106
	v118 = v107
	goto L35
L40:
	;
	v110 = int32(1)
	if v106 == v107 {
		v102 = v102 + v110
		v103 = v103 + v110
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v120 = int32(5)
	goto L44
L43:
	;
	v120 = int32(4)
	goto L44
L44:
	;
	return v120
}
func F_parse_scalar(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v8 == int32(11) {
		if v7 == int32(0) {
			v28 = F_json_lex(m, l0)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				return v28
			}
		} else {
			if v8 == int32(1) {
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
				if v35 != int32(1) {
					v67 = F_json_lex(m, l0)
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						if v67 != 0 {
							v86 = v67
							return v86
						} else {
							v69 = int32(0)
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v72 = m.T0[v7].(func(*base.Module, int32, int32, int32) int32)(m, v71, v69, v8)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								if v69 == int32(0) {
									v86 = v72
									return v86
								} else {
									v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v76&int32(4) == int32(0) {
										v86 = v72
										return v86
									} else {
										v81 = v69
										v82 = v72
										F_pfree(m, v81)
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return int32(0)
										} else {
											v86 = v82
											return v86
										}
									}
								}
							}
						}
					}
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
					v40 = F_pstrdup(m, v39)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						if v40 != 0 {
							v61 = v40
							v63 = F_json_lex(m, l0)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								if v63 == int32(0) {
									v69 = v61
									v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v72 = m.T0[v7].(func(*base.Module, int32, int32, int32) int32)(m, v71, v69, v8)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										if v69 == int32(0) {
											v86 = v72
											return v86
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if v76&int32(4) == int32(0) {
												v86 = v72
												return v86
											} else {
												v81 = v69
												v82 = v72
												F_pfree(m, v81)
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return int32(0)
												} else {
													v86 = v82
													return v86
												}
											}
										}
									}
								} else {
									v81 = v61
									v82 = v63
									F_pfree(m, v81)
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return int32(0)
									} else {
										v86 = v82
										return v86
									}
								}
							}
						} else {
							return int32(16)
						}
					}
				}
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v46 = v44 - v45
				v49 = F_palloc(m, v46+int32(1))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					if v49 == int32(0) {
						return int32(16)
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v46 != 0 {
							v56 = F__emscripten_memcpy_bulkmem(m, v49, v55, v46)
							mBase = m.M
							v57 = v56
						} else {
							v57 = v49
						}
						v59 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v57+v46))) = uint8(v59)
						v61 = v49
						v63 = F_json_lex(m, l0)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							if v63 == int32(0) {
								v69 = v61
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v72 = m.T0[v7].(func(*base.Module, int32, int32, int32) int32)(m, v71, v69, v8)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									if v69 == int32(0) {
										v86 = v72
										return v86
									} else {
										v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v76&int32(4) == int32(0) {
											v86 = v72
											return v86
										} else {
											v81 = v69
											v82 = v72
											F_pfree(m, v81)
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
												return int32(0)
											} else {
												v86 = v82
												return v86
											}
										}
									}
								}
							} else {
								v81 = v61
								v82 = v63
								F_pfree(m, v81)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									v86 = v82
									return v86
								}
							}
						}
					}
				}
			}
		}
	} else {
		if base.Ui32(int32(-3)) < base.Ui32(v8&int32(-9)-int32(3)) {
			if v7 == int32(0) {
				v28 = F_json_lex(m, l0)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					return v28
				}
			} else {
				if v8 == int32(1) {
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
					if v35 != int32(1) {
						v67 = F_json_lex(m, l0)
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							if v67 != 0 {
								v86 = v67
								return v86
							} else {
								v69 = int32(0)
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v72 = m.T0[v7].(func(*base.Module, int32, int32, int32) int32)(m, v71, v69, v8)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									if v69 == int32(0) {
										v86 = v72
										return v86
									} else {
										v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v76&int32(4) == int32(0) {
											v86 = v72
											return v86
										} else {
											v81 = v69
											v82 = v72
											F_pfree(m, v81)
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
												return int32(0)
											} else {
												v86 = v82
												return v86
											}
										}
									}
								}
							}
						}
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
						v40 = F_pstrdup(m, v39)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							if v40 != 0 {
								v61 = v40
								v63 = F_json_lex(m, l0)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									if v63 == int32(0) {
										v69 = v61
										v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v72 = m.T0[v7].(func(*base.Module, int32, int32, int32) int32)(m, v71, v69, v8)
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int32(0)
										} else {
											if v69 == int32(0) {
												v86 = v72
												return v86
											} else {
												v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												if v76&int32(4) == int32(0) {
													v86 = v72
													return v86
												} else {
													v81 = v69
													v82 = v72
													F_pfree(m, v81)
													mBase = m.M
													v84 = m.ExcPending
													if v84 != 0 {
														return int32(0)
													} else {
														v86 = v82
														return v86
													}
												}
											}
										}
									} else {
										v81 = v61
										v82 = v63
										F_pfree(m, v81)
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return int32(0)
										} else {
											v86 = v82
											return v86
										}
									}
								}
							} else {
								return int32(16)
							}
						}
					}
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v46 = v44 - v45
					v49 = F_palloc(m, v46+int32(1))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						if v49 == int32(0) {
							return int32(16)
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v46 != 0 {
								v56 = F__emscripten_memcpy_bulkmem(m, v49, v55, v46)
								mBase = m.M
								v57 = v56
							} else {
								v57 = v49
							}
							v59 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v57+v46))) = uint8(v59)
							v61 = v49
							v63 = F_json_lex(m, l0)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								if v63 == int32(0) {
									v69 = v61
									v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v72 = m.T0[v7].(func(*base.Module, int32, int32, int32) int32)(m, v71, v69, v8)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										if v69 == int32(0) {
											v86 = v72
											return v86
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if v76&int32(4) == int32(0) {
												v86 = v72
												return v86
											} else {
												v81 = v69
												v82 = v72
												F_pfree(m, v81)
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return int32(0)
												} else {
													v86 = v82
													return v86
												}
											}
										}
									}
								} else {
									v81 = v61
									v82 = v63
									F_pfree(m, v81)
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return int32(0)
									} else {
										v86 = v82
										return v86
									}
								}
							}
						}
					}
				}
			}
		} else {
			v17 = int32(11)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v20 != 0 {
				v21 = int32(10)
			} else {
				v21 = v17
			}
			if v8 == int32(12) {
				v24 = v17
			} else {
				v24 = v21
			}
			return v24
		}
	}
}
func F_partitioned_table_reloptions(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	if l0 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			F_errcode(m, int32(151027844))
			v8 = m.ExcPending
			if v8 != 0 {
				return
			} else {
				F_errmsg(m, int32(390709), int32(0))
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					F_errhint(m, int32(630157), int32(0))
					v16 = m.ExcPending
					if v16 != 0 {
						return
					} else {
						F_errfinish(m, int32(488388), int32(2026), int32(135679))
						v21 = m.ExcPending
						if v21 != 0 {
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
	} else {
		return
	}
}
func F_pattern_fixed_prefix(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	switch l1 - int32(1) {
	case 0:
		v15 = F_like_fixed_prefix(m, l0, int32(1), l2, l3, l4)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v15
		}
	case 1:
		v19 = F_regex_fixed_prefix(m, l0, int32(0), l2, l3, l4)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			return v19
		}
	case 2:
		v23 = F_regex_fixed_prefix(m, l0, int32(1), l2, l3, l4)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			return v23
		}
	case 3:
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
		v32 = F_datumCopy(m, v30, v31, v29)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
			v36 = F_makeConst(m, v26, v27, v28, v29, v32, v34, v35)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v36
				if l4 != 0 {
					*(*int64)(unsafe.Add(mBase, uint32(l4))) = int64(4607182418800017408)
				} else {
				}
				return int32(1)
			}
		}
	default:
		v9 = F_like_fixed_prefix(m, l0, int32(0), l2, l3, l4)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return v9
		}
	}
}
func F_perform_default_encoding_conversion(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
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
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	if l2 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L15
	} else {
		goto L45
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L15
	} else {
		goto L40
	}
L3:
	;
	m.G0 = v9 + int32(32)
	return v105
L4:
	;
	v13 = int32(4474148)
	goto L6
L5:
	;
	v13 = int32(4474152)
	goto L6
L6:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v14 == int32(0) {
		v105 = l0
		goto L3
	} else {
		goto L7
	}
L7:
	;
	if base.Ui32(int32(536870911)) <= base.Ui32(l1) {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[950]))
	v22 = *(*int32)(unsafe.Add(mBase, _consts[109]))
	if l2 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v23 = v20
	goto L11
L10:
	;
	v23 = v22
	goto L11
L11:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if l2 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v25 = v22
	goto L14
L13:
	;
	v25 = v20
	goto L14
L14:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v28 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v33 = F_MemoryContextAllocHuge(m, v28, l1<<(uint(int32(2))%32)|int32(1))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	v38 = F_FunctionCall6Coll(m, v14, v24, v26, l0, v33, l1, int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	if base.Ui32(l1) < base.Ui32(int32(1000001)) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v105 = v33
	goto L3
L19:
	;
	goto L20
L20:
	;
	if v33&int32(3) == int32(0) {
		v65 = v33
		goto L23
	} else {
		goto L24
	}
L21:
	;
	if base.Ui32(int32(1073741823)) <= base.Ui32(v98) {
		goto L1
	} else {
		goto L38
	}
L22:
	;
	v98 = v90 - v33
	goto L21
L23:
	;
	v69 = v65
	goto L32
L24:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v49 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v98 = int32(0)
	goto L21
L26:
	;
	goto L27
L27:
	;
	v54 = v33
	goto L28
L28:
	;
	v58 = v54 + int32(1)
	if v58&int32(3) == int32(0) {
		v65 = v58
		goto L23
	} else {
		goto L30
	}
L29:
	;
	v90 = v58
	goto L22
L30:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v63 != 0 {
		v54 = v58
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v78 = int32(-2139062144)
	if (int32(16843008)-v75|v75)&v78 == v78 {
		v69 = v69 + int32(4)
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v84 = v69
	goto L35
L34:
	;
	goto L33
L35:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v88 != 0 {
		v84 = v84 + int32(1)
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v90 = v84
	goto L22
L37:
	;
	goto L36
L38:
	;
	v103 = F_repalloc(m, v33, v98+int32(1))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L15
	} else {
		goto L39
	}
L39:
	;
	v105 = v103
	goto L3
L40:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L15
	} else {
		goto L41
	}
L41:
	;
	F_errmsg(m, int32(13845), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L15
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
	F_errdetail(m, int32(596471), v9)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L15
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(488525), int32(825), int32(268313))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L15
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L15
	} else {
		goto L46
	}
L46:
	;
	F_errmsg(m, int32(13845), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L15
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
	F_errdetail(m, int32(596471), v9+int32(16))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L15
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(488525), int32(852), int32(268313))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L15
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgarch_call_module_shutdown_cb(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, _consts[447]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	if v5 != 0 {
		v7 = *(*int32)(unsafe.Add(mBase, _consts[448]))
		m.T0[v5].(func(*base.Module, int32))(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_pgarch_die(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, _consts[443]))
	*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(-1)
	return
}
func F_pktreader_pull(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v9 != int32(3) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = l0 + int32(4)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v14 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	v64 = F_pullf_read(m, l1, l2, l3)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L11
	} else {
		goto L22
	}
L4:
	;
	return v62
L5:
	;
	v21 = v9
	goto L8
L6:
	;
	v44 = v14
	goto L7
L7:
	;
	if l2 < v44 {
		goto L17
	} else {
		goto L18
	}
L8:
	;
	if v21 == int32(1) {
		v62 = int32(0)
		goto L4
	} else {
		goto L10
	}
L9:
	;
	v44 = v35
	goto L7
L10:
	;
	v27 = F_parse_new_len(m, l1, v13)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	if v27 < int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return v27
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v27
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v35 == int32(0) {
		v21 = v27
		goto L8
	} else {
		goto L16
	}
L16:
	;
	goto L9
L17:
	;
	v47 = l2
	goto L19
L18:
	;
	v47 = v44
	goto L19
L19:
	;
	v48 = F_pullf_read(m, l1, v47, l3)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L11
	} else {
		goto L20
	}
L20:
	;
	if v48 <= int32(0) {
		v62 = v48
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v52 - v48
	v62 = v48
	goto L4
L22:
	;
	return v64
}
func F_plainnode(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	F_check_stack_depth(m)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v8 == v9 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v8 << (uint(int32(1)) % 32)
			v16 = F_repalloc(m, v7, v8*int32(24))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v16
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v20 = v16
				v21 = v19
				v24 = v21*int32(12) + v20
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v26 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
				*(*int64)(unsafe.Add(mBase, uint32(v24))) = v26
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v28
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
				if v31 == int32(1) {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v34 + int32(1)
					F_pfree(m, l1)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						return
					}
				} else {
					v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
					if v40 == int32(1) {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v48 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v43+v44*int32(12))+4)) = v48
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v50 + v48
						v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						F_plainnode(m, l0, v54)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							F_pfree(m, l1)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v59 + int32(1)
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						F_plainnode(m, l0, v63)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v66+v59*int32(12))+4)) = v70 - v59
							v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							F_plainnode(m, l0, v73)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return
							} else {
								F_pfree(m, l1)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			}
		} else {
			v20 = v7
			v21 = v8
			v24 = v21*int32(12) + v20
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v26 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
			*(*int64)(unsafe.Add(mBase, uint32(v24))) = v26
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v28
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
			if v31 == int32(1) {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v34 + int32(1)
				F_pfree(m, l1)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					return
				}
			} else {
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
				if v40 == int32(1) {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v48 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v43+v44*int32(12))+4)) = v48
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v50 + v48
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					F_plainnode(m, l0, v54)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						F_pfree(m, l1)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v59 + int32(1)
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					F_plainnode(m, l0, v63)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v66+v59*int32(12))+4)) = v70 - v59
						v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						F_plainnode(m, l0, v73)
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return
						} else {
							F_pfree(m, l1)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_pnstrdup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	v6 = F_memchr(m, l0, int32(0), l1)
	mBase = m.M
	if v6 != 0 {
		v8 = v6 - l0
	} else {
		v8 = l1
	}
	v10 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v11 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+4)) = uint8(v11)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v18 = m.T0[v17].(func(*base.Module, int32, int32, int32) int32)(m, v10, v8+int32(1), v11)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		if v8 != 0 {
			v22 = F__emscripten_memcpy_bulkmem(m, v18, l0, v8)
			mBase = m.M
			v23 = v22
		} else {
			v23 = v18
		}
		v25 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v8+v23))) = uint8(v25)
		return v23
	}
}
func F_policy_role_list_to_array(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v83 int32
	_ = v83
	if l0 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(0)
	return v83
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
	v12 = F_palloc(m, int32(4))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v16
	v20 = F_palloc(m, v16<<(uint(int32(2))%32))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L7
	}
L5:
	;
	return int32(0)
L6:
	;
	v83 = v12
	goto L1
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v22 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v29 = int32(0)
	goto L11
L9:
	;
	goto L10
L10:
	;
	return v20
L11:
	;
	v33 = v29 << (uint(int32(2)) % 32)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33+v34)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v37 == int32(4) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L10
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v40 == int32(1) {
		v83 = v20
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v67 = F_get_rolespec_oid(m, v36, int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L5
	} else {
		goto L25
	}
L16:
	;
	v45 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	if v45 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
	v83 = v20
	goto L1
L21:
	;
	F_errmsg(m, int32(538333), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	F_errhint(m, int32(615851), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(487019), int32(170), int32(24127))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+v33))) = v67
	v71 = v29 + int32(1)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v71 < v72 {
		v29 = v71
		goto L11
	} else {
		goto L26
	}
L26:
	;
	goto L12
}
func F_portuguese_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v189 int32
	_ = v189
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v308 int32
	_ = v308
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v376 int32
	_ = v376
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v427 int32
	_ = v427
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v453 int32
	_ = v453
	var v460 int32
	_ = v460
	var v471 int32
	_ = v471
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v548 int32
	_ = v548
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v574 int32
	_ = v574
	var v586 int32
	_ = v586
	var v593 int32
	_ = v593
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v615 int32
	_ = v615
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v666 int32
	_ = v666
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v692 int32
	_ = v692
	var v700 int32
	_ = v700
	var v711 int32
	_ = v711
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v789 int32
	_ = v789
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v815 int32
	_ = v815
	var v826 int32
	_ = v826
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v878 int32
	_ = v878
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v907 int32
	_ = v907
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v927 int32
	_ = v927
	var v933 int32
	_ = v933
	var v944 int32
	_ = v944
	var v951 int32
	_ = v951
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v975 int32
	_ = v975
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v997 int32
	_ = v997
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1026 int32
	_ = v1026
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1046 int32
	_ = v1046
	var v1052 int32
	_ = v1052
	var v1059 int32
	_ = v1059
	var v1070 int32
	_ = v1070
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1118 int32
	_ = v1118
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1134 int32
	_ = v1134
	var v1147 int32
	_ = v1147
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1167 int32
	_ = v1167
	var v1173 int32
	_ = v1173
	var v1185 int32
	_ = v1185
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1202 int32
	_ = v1202
	var v1204 int32
	_ = v1204
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1218 int32
	_ = v1218
	var v1221 int32
	_ = v1221
	var v1225 int32
	_ = v1225
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1247 int32
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1282 int32
	_ = v1282
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1295 int32
	_ = v1295
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1304 int32
	_ = v1304
	var v1314 int32
	_ = v1314
	var v1316 int32
	_ = v1316
	var v1320 int32
	_ = v1320
	var v1333 int32
	_ = v1333
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1353 int32
	_ = v1353
	var v1359 int32
	_ = v1359
	var v1366 int32
	_ = v1366
	var v1377 int32
	_ = v1377
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1404 int32
	_ = v1404
	var v1411 int32
	_ = v1411
	var v1413 int32
	_ = v1413
	var v1417 int32
	_ = v1417
	var v1420 int32
	_ = v1420
	var v1422 int32
	_ = v1422
	var v1426 int32
	_ = v1426
	var v1436 int32
	_ = v1436
	var v1438 int32
	_ = v1438
	var v1442 int32
	_ = v1442
	var v1455 int32
	_ = v1455
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1475 int32
	_ = v1475
	var v1481 int32
	_ = v1481
	var v1489 int32
	_ = v1489
	var v1500 int32
	_ = v1500
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1506 int32
	_ = v1506
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1529 int32
	_ = v1529
	var v1536 int32
	_ = v1536
	var v1538 int32
	_ = v1538
	var v1542 int32
	_ = v1542
	var v1545 int32
	_ = v1545
	var v1547 int32
	_ = v1547
	var v1551 int32
	_ = v1551
	var v1561 int32
	_ = v1561
	var v1563 int32
	_ = v1563
	var v1567 int32
	_ = v1567
	var v1580 int32
	_ = v1580
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1600 int32
	_ = v1600
	var v1606 int32
	_ = v1606
	var v1613 int32
	_ = v1613
	var v1624 int32
	_ = v1624
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1642 int32
	_ = v1642
	var v1643 int32
	_ = v1643
	var v1651 int32
	_ = v1651
	var v1658 int32
	_ = v1658
	var v1660 int32
	_ = v1660
	var v1664 int32
	_ = v1664
	var v1667 int32
	_ = v1667
	var v1669 int32
	_ = v1669
	var v1673 int32
	_ = v1673
	var v1683 int32
	_ = v1683
	var v1685 int32
	_ = v1685
	var v1689 int32
	_ = v1689
	var v1702 int32
	_ = v1702
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1722 int32
	_ = v1722
	var v1728 int32
	_ = v1728
	var v1736 int32
	_ = v1736
	var v1747 int32
	_ = v1747
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1756 int32
	_ = v1756
	var v1762 int32
	_ = v1762
	var v1766 int32
	_ = v1766
	var v1779 int32
	_ = v1779
	var v1780 int32
	_ = v1780
	var v1783 int32
	_ = v1783
	var v1787 int32
	_ = v1787
	var v1788 int32
	_ = v1788
	var v1790 int32
	_ = v1790
	var v1791 int32
	_ = v1791
	var v1794 int32
	_ = v1794
	var v1795 int32
	_ = v1795
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1828 int32
	_ = v1828
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1834 int32
	_ = v1834
	var v1836 int32
	_ = v1836
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1853 int32
	_ = v1853
	var v1855 int32
	_ = v1855
	var v1856 int32
	_ = v1856
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1864 int32
	_ = v1864
	var v1866 int32
	_ = v1866
	var v1868 int32
	_ = v1868
	var v1871 int32
	_ = v1871
	var v1874 int32
	_ = v1874
	var v1877 int32
	_ = v1877
	var v1881 int32
	_ = v1881
	var v1884 int32
	_ = v1884
	var v1886 int32
	_ = v1886
	var v1887 int32
	_ = v1887
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1893 int32
	_ = v1893
	var v1894 int32
	_ = v1894
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1900 int32
	_ = v1900
	var v1902 int32
	_ = v1902
	var v1906 int32
	_ = v1906
	var v1910 int32
	_ = v1910
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1919 int32
	_ = v1919
	var v1921 int32
	_ = v1921
	var v1922 int32
	_ = v1922
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1928 int32
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1935 int32
	_ = v1935
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1941 int32
	_ = v1941
	var v1943 int32
	_ = v1943
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1960 int32
	_ = v1960
	var v1962 int32
	_ = v1962
	var v1963 int32
	_ = v1963
	var v1965 int32
	_ = v1965
	var v1966 int32
	_ = v1966
	var v1969 int32
	_ = v1969
	var v1970 int32
	_ = v1970
	var v1972 int32
	_ = v1972
	var v1973 int32
	_ = v1973
	var v1976 int32
	_ = v1976
	var v1978 int32
	_ = v1978
	var v1980 int32
	_ = v1980
	var v1983 int32
	_ = v1983
	var v1986 int32
	_ = v1986
	var v1989 int32
	_ = v1989
	var v1993 int32
	_ = v1993
	var v1996 int32
	_ = v1996
	var v1998 int32
	_ = v1998
	var v1999 int32
	_ = v1999
	var v2001 int32
	_ = v2001
	var v2002 int32
	_ = v2002
	var v2005 int32
	_ = v2005
	var v2006 int32
	_ = v2006
	var v2008 int32
	_ = v2008
	var v2010 int32
	_ = v2010
	var v2014 int32
	_ = v2014
	var v2022 int32
	_ = v2022
	var v2023 int32
	_ = v2023
	var v2028 int32
	_ = v2028
	var v2030 int32
	_ = v2030
	var v2031 int32
	_ = v2031
	var v2034 int32
	_ = v2034
	var v2038 int32
	_ = v2038
	var v2039 int32
	_ = v2039
	var v2041 int32
	_ = v2041
	var v2042 int32
	_ = v2042
	var v2048 int32
	_ = v2048
	var v2049 int32
	_ = v2049
	var v2052 int32
	_ = v2052
	var v2054 int32
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2057 int32
	_ = v2057
	var v2058 int32
	_ = v2058
	var v2061 int32
	_ = v2061
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2071 int32
	_ = v2071
	var v2074 int32
	_ = v2074
	var v2076 int32
	_ = v2076
	var v2077 int32
	_ = v2077
	var v2080 int32
	_ = v2080
	var v2084 int32
	_ = v2084
	var v2090 int32
	_ = v2090
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2096 int32
	_ = v2096
	var v2097 int32
	_ = v2097
	var v2104 int32
	_ = v2104
	var v2109 int32
	_ = v2109
	var v2110 int32
	_ = v2110
	var v2113 int32
	_ = v2113
	var v2117 int32
	_ = v2117
	var v2118 int32
	_ = v2118
	var v2120 int32
	_ = v2120
	var v2121 int32
	_ = v2121
	var v2124 int32
	_ = v2124
	var v2126 int32
	_ = v2126
	var v2128 int32
	_ = v2128
	var v2129 int32
	_ = v2129
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2136 int32
	_ = v2136
	var v2142 int32
	_ = v2142
	var v2147 int32
	_ = v2147
	var v2151 int32
	_ = v2151
	var v2157 int32
	_ = v2157
	var v2160 int32
	_ = v2160
	var v2162 int32
	_ = v2162
	var v2163 int32
	_ = v2163
	var v2165 int32
	_ = v2165
	var v2166 int32
	_ = v2166
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2180 int32
	_ = v2180
	var v2182 int32
	_ = v2182
	var v2184 int32
	_ = v2184
	var v2185 int32
	_ = v2185
	var v2191 int32
	_ = v2191
	var v2193 int32
	_ = v2193
	var v2195 int32
	_ = v2195
	var v2201 int32
	_ = v2201
	var v2202 int32
	_ = v2202
	var v2203 int32
	_ = v2203
	var v2209 int32
	_ = v2209
	var v2210 int32
	_ = v2210
	var v2215 int32
	_ = v2215
	var v2216 int32
	_ = v2216
	var v2219 int32
	_ = v2219
	var v2220 int32
	_ = v2220
	var v2221 int32
	_ = v2221
	var v2222 int32
	_ = v2222
	var v2229 int32
	_ = v2229
	var v2231 int32
	_ = v2231
	var v2236 int32
	_ = v2236
	var v2238 int32
	_ = v2238
	var v2245 int32
	_ = v2245
	var v2248 int32
	_ = v2248
	var v2252 int32
	_ = v2252
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2274 int32
	_ = v2274
	var v2280 int32
	_ = v2280
	var v2281 int32
	_ = v2281
	var v2285 int32
	_ = v2285
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = v7
	goto L2
L1:
	;
	return v2285
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v9
	v16 = v9 + int32(1)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v17 <= v16 {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v110)+4)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v110)+8)) = v111
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v110))) = v114
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L46
L4:
	;
	goto L3
L5:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = v108
	goto L2
L6:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L21
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v9
	v48 = v9
	v49 = v17
	goto L6
L8:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+v16))))
	switch v21 - int32(163) {
	case 0, 18:
		goto L9
	default:
		goto L7
	}
L9:
	;
	v26 = F_find_among(m, l0, int32(4288752), int32(3))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v30
	switch v26 - int32(1) {
	case 0:
		goto L13
	case 1:
		goto L12
	case 2:
		goto L14
	default:
		goto L5
	}
L12:
	;
	v43 = F_slice_from_s(m, l0, int32(2), int32(2185460))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L10
	} else {
		goto L17
	}
L13:
	;
	v37 = F_slice_from_s(m, l0, int32(2), int32(2185458))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L10
	} else {
		goto L15
	}
L14:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v48 = v30
	v49 = v34
	goto L6
L15:
	;
	if int32(0) <= v37 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v2285 = v37
	goto L1
L17:
	;
	if int32(0) <= v43 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v2285 = v43
	goto L1
L19:
	;
	if v102 < int32(0) {
		goto L4
	} else {
		goto L39
	}
L21:
	;
	goto L22
L22:
	;
	goto L23
L23:
	;
	v57 = v48
	v59 = int32(1)
	goto L26
L25:
	;
	v102 = v87
	goto L19
L26:
	;
	if v49 <= v57 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L25
L28:
	;
	v102 = int32(-1)
	goto L19
L29:
	;
	goto L30
L30:
	;
	v64 = v57 + int32(1)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v57))))
	if base.Ui32(v66) < base.Ui32(int32(192)) {
		v87 = v64
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v88 = int32(1)
	if v88 < v59 {
		v57 = v87
		v59 = v59 - v88
		goto L26
	} else {
		goto L38
	}
L32:
	;
	if v49 <= v64 {
		v87 = v64
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v73 = v64
	goto L34
L34:
	;
	v76 = int32(*(*int8)(unsafe.Add(mBase, uint32(v50+v73))))
	if int32(-65) < v76 {
		v87 = v73
		goto L31
	} else {
		goto L36
	}
L35:
	;
	v87 = v49
	goto L31
L36:
	;
	v80 = v73 + int32(1)
	if v80 != v49 {
		v73 = v80
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	goto L27
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v102
	goto L5
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v116
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1282 = v116
	goto L302
L41:
	;
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1256)+8)) = v1254
	goto L40
L42:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1254 = v1252 + v1250
	goto L41
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v116
	v729 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v730 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L175
L44:
	;
	if v234 != 0 {
		goto L43
	} else {
		goto L68
	}
L45:
	;
	v234 = v227
	goto L44
L46:
	;
	if v129 <= v116 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v227 = int32(0)
	goto L45
L48:
	;
	v234 = int32(-1)
	goto L44
L49:
	;
	goto L50
L50:
	;
	v145 = int32(1)
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116+v130))))
	if base.Ui32(v147) < base.Ui32(int32(192)) {
		v204 = v147
		v205 = v145
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if int32(250) < v204 {
		v227 = v205
		goto L45
	} else {
		goto L64
	}
L52:
	;
	v151 = v116 + int32(1)
	if v151 == v129 {
		v204 = v147
		v205 = v145
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v130))))
	v156 = v154 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v147) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160+v130))))
	v172 = v170 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v147) {
		goto L60
	} else {
		goto L61
	}
L55:
	;
	v160 = v116 + int32(2)
	if v160 != v129 {
		goto L54
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v204 = v147<<(uint(int32(6))%32)&int32(1984) | v156
	v205 = int32(2)
	goto L51
L58:
	;
	goto L57
L59:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130+v176))))
	v204 = v189&int32(63) | (v147<<(uint(int32(18))%32)&int32(1835008) | v156<<(uint(int32(12))%32) | v172<<(uint(int32(6))%32))
	v205 = int32(4)
	goto L51
L60:
	;
	v176 = v116 + int32(3)
	if v176 != v129 {
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v204 = v147<<(uint(int32(12))%32)&int32(61440) | v156<<(uint(int32(6))%32) | v172
	v205 = int32(3)
	goto L51
L63:
	;
	goto L62
L64:
	;
	v209 = v204 - int32(97)
	if v209 < int32(0) {
		v227 = v205
		goto L45
	} else {
		goto L65
	}
L65:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v209)>>(uint(int32(3))%32)))+uint32(_consts[1067]))))
	if int32(base.Ui32(v215)>>(uint(v209&int32(7))%32))&int32(1) == int32(0) {
		v227 = v205
		goto L45
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v205 + v116
	goto L67
L67:
	;
	goto L47
L68:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L71
L69:
	;
	if v352 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L70:
	;
	v352 = v345
	goto L69
L71:
	;
	if v248 <= v235 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v345 = int32(0)
	goto L70
L73:
	;
	v352 = int32(-1)
	goto L69
L74:
	;
	goto L75
L75:
	;
	v264 = int32(1)
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235+v249))))
	if base.Ui32(v266) < base.Ui32(int32(192)) {
		v323 = v266
		v324 = v264
		goto L76
	} else {
		goto L77
	}
L76:
	;
	if int32(250) < v323 {
		goto L89
	} else {
		goto L90
	}
L77:
	;
	v270 = v235 + int32(1)
	if v270 == v248 {
		v323 = v266
		v324 = v264
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270+v249))))
	v275 = v273 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v266) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279+v249))))
	v291 = v289 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v266) {
		goto L85
	} else {
		goto L86
	}
L80:
	;
	v279 = v235 + int32(2)
	if v279 != v248 {
		goto L79
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v323 = v266<<(uint(int32(6))%32)&int32(1984) | v275
	v324 = int32(2)
	goto L76
L83:
	;
	goto L82
L84:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249+v295))))
	v323 = v308&int32(63) | (v266<<(uint(int32(18))%32)&int32(1835008) | v275<<(uint(int32(12))%32) | v291<<(uint(int32(6))%32))
	v324 = int32(4)
	goto L76
L85:
	;
	v295 = v235 + int32(3)
	if v295 != v248 {
		goto L84
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v323 = v266<<(uint(int32(12))%32)&int32(61440) | v275<<(uint(int32(6))%32) | v291
	v324 = int32(3)
	goto L76
L88:
	;
	goto L87
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v324 + v235
	goto L93
L90:
	;
	v328 = v323 - int32(97)
	if v328 < int32(0) {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v328)>>(uint(int32(3))%32)))+uint32(_consts[1067]))))
	if int32(base.Ui32(v334)>>(uint(v328&int32(7))%32))&int32(1) != 0 {
		v345 = v324
		goto L70
	} else {
		goto L92
	}
L92:
	;
	goto L89
L93:
	;
	goto L72
L94:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v376 = v366
	goto L99
L95:
	;
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v235
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L125
L97:
	;
	if int32(0) <= v471 {
		v1250 = v471
		goto L42
	} else {
		goto L122
	}
L98:
	;
	v471 = v443
	goto L97
L99:
	;
	if v367 <= v376 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v471 = int32(-1)
	goto L97
L102:
	;
	goto L103
L103:
	;
	v383 = int32(1)
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376+v368))))
	if base.Ui32(v385) < base.Ui32(int32(192)) {
		v442 = v385
		v443 = v383
		goto L104
	} else {
		goto L105
	}
L104:
	;
	if int32(250) < v442 {
		goto L117
	} else {
		goto L118
	}
L105:
	;
	v389 = v376 + int32(1)
	if v389 == v367 {
		v442 = v385
		v443 = v383
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389+v368))))
	v394 = v392 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v385) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v398+v368))))
	v410 = v408 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v385) {
		goto L113
	} else {
		goto L114
	}
L108:
	;
	v398 = v376 + int32(2)
	if v398 != v367 {
		goto L107
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v442 = v385<<(uint(int32(6))%32)&int32(1984) | v394
	v443 = int32(2)
	goto L104
L111:
	;
	goto L110
L112:
	;
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368+v414))))
	v442 = v427&int32(63) | (v385<<(uint(int32(18))%32)&int32(1835008) | v394<<(uint(int32(12))%32) | v410<<(uint(int32(6))%32))
	v443 = int32(4)
	goto L104
L113:
	;
	v414 = v376 + int32(3)
	if v414 != v367 {
		goto L112
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v442 = v385<<(uint(int32(12))%32)&int32(61440) | v394<<(uint(int32(6))%32) | v410
	v443 = int32(3)
	goto L104
L116:
	;
	goto L115
L117:
	;
	v460 = v443 + v376
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460
	v376 = v460
	goto L99
L118:
	;
	v447 = v442 - int32(97)
	if v447 < int32(0) {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v447)>>(uint(int32(3))%32)))+uint32(_consts[1067]))))
	if int32(base.Ui32(v453)>>(uint(v447&int32(7))%32))&int32(1) != 0 {
		goto L98
	} else {
		goto L120
	}
L120:
	;
	goto L117
L122:
	;
	goto L96
L123:
	;
	if v593 != 0 {
		goto L43
	} else {
		goto L147
	}
L124:
	;
	v593 = v586
	goto L123
L125:
	;
	if v488 <= v235 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v586 = int32(0)
	goto L124
L127:
	;
	v593 = int32(-1)
	goto L123
L128:
	;
	goto L129
L129:
	;
	v504 = int32(1)
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235+v489))))
	if base.Ui32(v506) < base.Ui32(int32(192)) {
		v563 = v506
		v564 = v504
		goto L130
	} else {
		goto L131
	}
L130:
	;
	if int32(250) < v563 {
		v586 = v564
		goto L124
	} else {
		goto L143
	}
L131:
	;
	v510 = v235 + int32(1)
	if v510 == v488 {
		v563 = v506
		v564 = v504
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510+v489))))
	v515 = v513 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v506) {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519+v489))))
	v531 = v529 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v506) {
		goto L139
	} else {
		goto L140
	}
L134:
	;
	v519 = v235 + int32(2)
	if v519 != v488 {
		goto L133
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v563 = v506<<(uint(int32(6))%32)&int32(1984) | v515
	v564 = int32(2)
	goto L130
L137:
	;
	goto L136
L138:
	;
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489+v535))))
	v563 = v548&int32(63) | (v506<<(uint(int32(18))%32)&int32(1835008) | v515<<(uint(int32(12))%32) | v531<<(uint(int32(6))%32))
	v564 = int32(4)
	goto L130
L139:
	;
	v535 = v235 + int32(3)
	if v535 != v488 {
		goto L138
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v563 = v506<<(uint(int32(12))%32)&int32(61440) | v515<<(uint(int32(6))%32) | v531
	v564 = int32(3)
	goto L130
L142:
	;
	goto L141
L143:
	;
	v568 = v563 - int32(97)
	if v568 < int32(0) {
		v586 = v564
		goto L124
	} else {
		goto L144
	}
L144:
	;
	v574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v568)>>(uint(int32(3))%32)))+uint32(_consts[1067]))))
	if int32(base.Ui32(v574)>>(uint(v568&int32(7))%32))&int32(1) == int32(0) {
		v586 = v564
		goto L124
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v564 + v235
	goto L146
L146:
	;
	goto L126
L147:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v606 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v615 = v605
	goto L150
L148:
	;
	if int32(0) <= v711 {
		v1250 = v711
		goto L42
	} else {
		goto L172
	}
L149:
	;
	v711 = v682
	goto L148
L150:
	;
	if v606 <= v615 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v711 = int32(-1)
	goto L148
L153:
	;
	goto L154
L154:
	;
	v622 = int32(1)
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v615+v607))))
	if base.Ui32(v624) < base.Ui32(int32(192)) {
		v681 = v624
		v682 = v622
		goto L155
	} else {
		goto L156
	}
L155:
	;
	if int32(250) < v681 {
		goto L149
	} else {
		goto L168
	}
L156:
	;
	v628 = v615 + int32(1)
	if v628 == v606 {
		v681 = v624
		v682 = v622
		goto L155
	} else {
		goto L157
	}
L157:
	;
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v628+v607))))
	v633 = v631 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v624) {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637+v607))))
	v649 = v647 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v624) {
		goto L164
	} else {
		goto L165
	}
L159:
	;
	v637 = v615 + int32(2)
	if v637 != v606 {
		goto L158
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v681 = v624<<(uint(int32(6))%32)&int32(1984) | v633
	v682 = int32(2)
	goto L155
L162:
	;
	goto L161
L163:
	;
	v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v607+v653))))
	v681 = v666&int32(63) | (v624<<(uint(int32(18))%32)&int32(1835008) | v633<<(uint(int32(12))%32) | v649<<(uint(int32(6))%32))
	v682 = int32(4)
	goto L155
L164:
	;
	v653 = v615 + int32(3)
	if v653 != v606 {
		goto L163
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v681 = v624<<(uint(int32(12))%32)&int32(61440) | v633<<(uint(int32(6))%32) | v649
	v682 = int32(3)
	goto L155
L167:
	;
	goto L166
L168:
	;
	v686 = v681 - int32(97)
	if v686 < int32(0) {
		goto L149
	} else {
		goto L169
	}
L169:
	;
	v692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v686)>>(uint(int32(3))%32)))+uint32(_consts[1067]))))
	if int32(base.Ui32(v692)>>(uint(v686&int32(7))%32))&int32(1) == int32(0) {
		goto L149
	} else {
		goto L170
	}
L170:
	;
	v700 = v682 + v615
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v700
	v615 = v700
	goto L150
L172:
	;
	goto L43
L173:
	;
	if v833 != 0 {
		goto L40
	} else {
		goto L198
	}
L174:
	;
	v833 = v826
	goto L173
L175:
	;
	if v729 <= v116 {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	v826 = int32(0)
	goto L174
L177:
	;
	v833 = int32(-1)
	goto L173
L178:
	;
	goto L179
L179:
	;
	v745 = int32(1)
	v747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116+v730))))
	if base.Ui32(v747) < base.Ui32(int32(192)) {
		v804 = v747
		v805 = v745
		goto L180
	} else {
		goto L181
	}
L180:
	;
	if int32(250) < v804 {
		goto L193
	} else {
		goto L194
	}
L181:
	;
	v751 = v116 + int32(1)
	if v751 == v729 {
		v804 = v747
		v805 = v745
		goto L180
	} else {
		goto L182
	}
L182:
	;
	v754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v751+v730))))
	v756 = v754 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v747) {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v760+v730))))
	v772 = v770 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v747) {
		goto L189
	} else {
		goto L190
	}
L184:
	;
	v760 = v116 + int32(2)
	if v760 != v729 {
		goto L183
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v804 = v747<<(uint(int32(6))%32)&int32(1984) | v756
	v805 = int32(2)
	goto L180
L187:
	;
	goto L186
L188:
	;
	v789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v730+v776))))
	v804 = v789&int32(63) | (v747<<(uint(int32(18))%32)&int32(1835008) | v756<<(uint(int32(12))%32) | v772<<(uint(int32(6))%32))
	v805 = int32(4)
	goto L180
L189:
	;
	v776 = v116 + int32(3)
	if v776 != v729 {
		goto L188
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	v804 = v747<<(uint(int32(12))%32)&int32(61440) | v756<<(uint(int32(6))%32) | v772
	v805 = int32(3)
	goto L180
L192:
	;
	goto L191
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v805 + v116
	goto L197
L194:
	;
	v809 = v804 - int32(97)
	if v809 < int32(0) {
		goto L193
	} else {
		goto L195
	}
L195:
	;
	v815 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v809)>>(uint(int32(3))%32)))+uint32(_consts[1067]))))
	if int32(base.Ui32(v815)>>(uint(v809&int32(7))%32))&int32(1) != 0 {
		v826 = v805
		goto L174
	} else {
		goto L196
	}
L196:
	;
	goto L193
L197:
	;
	goto L176
L198:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v847 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v848 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L201
L199:
	;
	if v951 == int32(0) {
		goto L224
	} else {
		goto L225
	}
L200:
	;
	v951 = v944
	goto L199
L201:
	;
	if v847 <= v834 {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	v944 = int32(0)
	goto L200
L203:
	;
	v951 = int32(-1)
	goto L199
L204:
	;
	goto L205
L205:
	;
	v863 = int32(1)
	v865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v834+v848))))
	if base.Ui32(v865) < base.Ui32(int32(192)) {
		v922 = v865
		v923 = v863
		goto L206
	} else {
		goto L207
	}
L206:
	;
	if int32(250) < v922 {
		goto L219
	} else {
		goto L220
	}
L207:
	;
	v869 = v834 + int32(1)
	if v869 == v847 {
		v922 = v865
		v923 = v863
		goto L206
	} else {
		goto L208
	}
L208:
	;
	v872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v869+v848))))
	v874 = v872 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v865) {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	v888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v878+v848))))
	v890 = v888 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v865) {
		goto L215
	} else {
		goto L216
	}
L210:
	;
	v878 = v834 + int32(2)
	if v878 != v847 {
		goto L209
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v922 = v865<<(uint(int32(6))%32)&int32(1984) | v874
	v923 = int32(2)
	goto L206
L213:
	;
	goto L212
L214:
	;
	v907 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v848+v894))))
	v922 = v907&int32(63) | (v865<<(uint(int32(18))%32)&int32(1835008) | v874<<(uint(int32(12))%32) | v890<<(uint(int32(6))%32))
	v923 = int32(4)
	goto L206
L215:
	;
	v894 = v834 + int32(3)
	if v894 != v847 {
		goto L214
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	v922 = v865<<(uint(int32(12))%32)&int32(61440) | v874<<(uint(int32(6))%32) | v890
	v923 = int32(3)
	goto L206
L218:
	;
	goto L217
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v923 + v834
	goto L223
L220:
	;
	v927 = v922 - int32(97)
	if v927 < int32(0) {
		goto L219
	} else {
		goto L221
	}
L221:
	;
	v933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v927)>>(uint(int32(3))%32)))+uint32(_consts[1067]))))
	if int32(base.Ui32(v933)>>(uint(v927&int32(7))%32))&int32(1) != 0 {
		v944 = v923
		goto L200
	} else {
		goto L222
	}
L222:
	;
	goto L219
L223:
	;
	goto L202
L224:
	;
	v965 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v966 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v967 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v975 = v965
	goto L229
L225:
	;
	goto L226
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v834
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L255
L227:
	;
	if int32(0) <= v1070 {
		v1250 = v1070
		goto L42
	} else {
		goto L252
	}
L228:
	;
	v1070 = v1042
	goto L227
L229:
	;
	if v966 <= v975 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v1070 = int32(-1)
	goto L227
L232:
	;
	goto L233
L233:
	;
	v982 = int32(1)
	v984 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v975+v967))))
	if base.Ui32(v984) < base.Ui32(int32(192)) {
		v1041 = v984
		v1042 = v982
		goto L234
	} else {
		goto L235
	}
L234:
	;
	if int32(250) < v1041 {
		goto L247
	} else {
		goto L248
	}
L235:
	;
	v988 = v975 + int32(1)
	if v988 == v966 {
		v1041 = v984
		v1042 = v982
		goto L234
	} else {
		goto L236
	}
L236:
	;
	v991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v988+v967))))
	v993 = v991 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v984) {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	v1007 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v997+v967))))
	v1009 = v1007 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v984) {
		goto L243
	} else {
		goto L244
	}
L238:
	;
	v997 = v975 + int32(2)
	if v997 != v966 {
		goto L237
	} else {
		goto L241
	}
L239:
	;
	goto L240
L240:
	;
	v1041 = v984<<(uint(int32(6))%32)&int32(1984) | v993
	v1042 = int32(2)
	goto L234
L241:
	;
	goto L240
L242:
	;
	v1026 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v967+v1013))))
	v1041 = v1026&int32(63) | (v984<<(uint(int32(18))%32)&int32(1835008) | v993<<(uint(int32(12))%32) | v1009<<(uint(int32(6))%32))
	v1042 = int32(4)
	goto L234
L243:
	;
	v1013 = v975 + int32(3)
	if v1013 != v966 {
		goto L242
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	v1041 = v984<<(uint(int32(12))%32)&int32(61440) | v993<<(uint(int32(6))%32) | v1009
	v1042 = int32(3)
	goto L234
L246:
	;
	goto L245
L247:
	;
	v1059 = v1042 + v975
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1059
	v975 = v1059
	goto L229
L248:
	;
	v1046 = v1041 - int32(97)
	if v1046 < int32(0) {
		goto L247
	} else {
		goto L249
	}
L249:
	;
	v1052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1046)>>(uint(int32(3))%32)))+uint32(_consts[1067]))))
	if int32(base.Ui32(v1052)>>(uint(v1046&int32(7))%32))&int32(1) != 0 {
		goto L228
	} else {
		goto L250
	}
L250:
	;
	goto L247
L252:
	;
	goto L226
L253:
	;
	if v1192 != 0 {
		goto L40
	} else {
		goto L277
	}
L254:
	;
	v1192 = v1185
	goto L253
L255:
	;
	if v1087 <= v834 {
		goto L257
	} else {
		goto L258
	}
L256:
	;
	v1185 = int32(0)
	goto L254
L257:
	;
	v1192 = int32(-1)
	goto L253
L258:
	;
	goto L259
L259:
	;
	v1103 = int32(1)
	v1105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v834+v1088))))
	if base.Ui32(v1105) < base.Ui32(int32(192)) {
		v1162 = v1105
		v1163 = v1103
		goto L260
	} else {
		goto L261
	}
L260:
	;
	if int32(250) < v1162 {
		v1185 = v1163
		goto L254
	} else {
		goto L273
	}
L261:
	;
	v1109 = v834 + int32(1)
	if v1109 == v1087 {
		v1162 = v1105
		v1163 = v1103
		goto L260
	} else {
		goto L262
	}
L262:
	;
	v1112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1109+v1088))))
	v1114 = v1112 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1105) {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	v1128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1118+v1088))))
	v1130 = v1128 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1105) {
		goto L269
	} else {
		goto L270
	}
L264:
	;
	v1118 = v834 + int32(2)
	if v1118 != v1087 {
		goto L263
	} else {
		goto L267
	}
L265:
	;
	goto L266
L266:
	;
	v1162 = v1105<<(uint(int32(6))%32)&int32(1984) | v1114
	v1163 = int32(2)
	goto L260
L267:
	;
	goto L266
L268:
	;
	v1147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1088+v1134))))
	v1162 = v1147&int32(63) | (v1105<<(uint(int32(18))%32)&int32(1835008) | v1114<<(uint(int32(12))%32) | v1130<<(uint(int32(6))%32))
	v1163 = int32(4)
	goto L260
L269:
	;
	v1134 = v834 + int32(3)
	if v1134 != v1087 {
		goto L268
	} else {
		goto L272
	}
L270:
	;
	goto L271
L271:
	;
	v1162 = v1105<<(uint(int32(12))%32)&int32(61440) | v1114<<(uint(int32(6))%32) | v1130
	v1163 = int32(3)
	goto L260
L272:
	;
	goto L271
L273:
	;
	v1167 = v1162 - int32(97)
	if v1167 < int32(0) {
		v1185 = v1163
		goto L254
	} else {
		goto L274
	}
L274:
	;
	v1173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1167)>>(uint(int32(3))%32)))+uint32(_consts[1067]))))
	if int32(base.Ui32(v1173)>>(uint(v1167&int32(7))%32))&int32(1) == int32(0) {
		v1185 = v1163
		goto L254
	} else {
		goto L275
	}
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1163 + v834
	goto L276
L276:
	;
	goto L256
L277:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L280
L278:
	;
	if int32(0) <= v1247 {
		v1254 = v1247
		goto L41
	} else {
		goto L298
	}
L280:
	;
	goto L281
L281:
	;
	goto L282
L282:
	;
	v1202 = v1194
	v1204 = int32(1)
	goto L285
L284:
	;
	v1247 = v1232
	goto L278
L285:
	;
	if v1195 <= v1202 {
		goto L287
	} else {
		goto L288
	}
L286:
	;
	goto L284
L287:
	;
	v1247 = int32(-1)
	goto L278
L288:
	;
	goto L289
L289:
	;
	v1209 = v1202 + int32(1)
	v1211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1193+v1202))))
	if base.Ui32(v1211) < base.Ui32(int32(192)) {
		v1232 = v1209
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v1233 = int32(1)
	if v1233 < v1204 {
		v1202 = v1232
		v1204 = v1204 - v1233
		goto L285
	} else {
		goto L297
	}
L291:
	;
	if v1195 <= v1209 {
		v1232 = v1209
		goto L290
	} else {
		goto L292
	}
L292:
	;
	v1218 = v1209
	goto L293
L293:
	;
	v1221 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1193+v1218))))
	if int32(-65) < v1221 {
		v1232 = v1218
		goto L290
	} else {
		goto L295
	}
L294:
	;
	v1232 = v1195
	goto L290
L295:
	;
	v1225 = v1218 + int32(1)
	if v1225 != v1195 {
		v1218 = v1225
		goto L293
	} else {
		goto L296
	}
L296:
	;
	goto L294
L297:
	;
	goto L286
L298:
	;
	goto L40
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v116
	v1756 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1756
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1756
	if v1756-int32(2) <= v116 {
		goto L404
	} else {
		goto L405
	}
L300:
	;
	if v1377 < int32(0) {
		goto L299
	} else {
		goto L325
	}
L301:
	;
	v1377 = v1349
	goto L300
L302:
	;
	if v1273 <= v1282 {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v1377 = int32(-1)
	goto L300
L305:
	;
	goto L306
L306:
	;
	v1289 = int32(1)
	v1291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1282+v1274))))
	if base.Ui32(v1291) < base.Ui32(int32(192)) {
		v1348 = v1291
		v1349 = v1289
		goto L307
	} else {
		goto L308
	}
L307:
	;
	if int32(250) < v1348 {
		goto L320
	} else {
		goto L321
	}
L308:
	;
	v1295 = v1282 + int32(1)
	if v1295 == v1273 {
		v1348 = v1291
		v1349 = v1289
		goto L307
	} else {
		goto L309
	}
L309:
	;
	v1298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1295+v1274))))
	v1300 = v1298 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1291) {
		goto L311
	} else {
		goto L312
	}
L310:
	;
	v1314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1304+v1274))))
	v1316 = v1314 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1291) {
		goto L316
	} else {
		goto L317
	}
L311:
	;
	v1304 = v1282 + int32(2)
	if v1304 != v1273 {
		goto L310
	} else {
		goto L314
	}
L312:
	;
	goto L313
L313:
	;
	v1348 = v1291<<(uint(int32(6))%32)&int32(1984) | v1300
	v1349 = int32(2)
	goto L307
L314:
	;
	goto L313
L315:
	;
	v1333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1274+v1320))))
	v1348 = v1333&int32(63) | (v1291<<(uint(int32(18))%32)&int32(1835008) | v1300<<(uint(int32(12))%32) | v1316<<(uint(int32(6))%32))
	v1349 = int32(4)
	goto L307
L316:
	;
	v1320 = v1282 + int32(3)
	if v1320 != v1273 {
		goto L315
	} else {
		goto L319
	}
L317:
	;
	goto L318
L318:
	;
	v1348 = v1291<<(uint(int32(12))%32)&int32(61440) | v1300<<(uint(int32(6))%32) | v1316
	v1349 = int32(3)
	goto L307
L319:
	;
	goto L318
L320:
	;
	v1366 = v1349 + v1282
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1366
	v1282 = v1366
	goto L302
L321:
	;
	v1353 = v1348 - int32(97)
	if v1353 < int32(0) {
		goto L320
	} else {
		goto L322
	}
L322:
	;
	v1359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1353)>>(uint(int32(3))%32)))+uint32(_consts[1067]))))
	if int32(base.Ui32(v1359)>>(uint(v1353&int32(7))%32))&int32(1) != 0 {
		goto L301
	} else {
		goto L323
	}
L323:
	;
	goto L320
L325:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1381 = v1380 + v1377
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1381
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1404 = v1381
	goto L328
L326:
	;
	if v1500 < int32(0) {
		goto L299
	} else {
		goto L350
	}
L327:
	;
	v1500 = v1471
	goto L326
L328:
	;
	if v1395 <= v1404 {
		goto L330
	} else {
		goto L331
	}
L330:
	;
	v1500 = int32(-1)
	goto L326
L331:
	;
	goto L332
L332:
	;
	v1411 = int32(1)
	v1413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1404+v1396))))
	if base.Ui32(v1413) < base.Ui32(int32(192)) {
		v1470 = v1413
		v1471 = v1411
		goto L333
	} else {
		goto L334
	}
L333:
	;
	if int32(250) < v1470 {
		goto L327
	} else {
		goto L346
	}
L334:
	;
	v1417 = v1404 + int32(1)
	if v1417 == v1395 {
		v1470 = v1413
		v1471 = v1411
		goto L333
	} else {
		goto L335
	}
L335:
	;
	v1420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1417+v1396))))
	v1422 = v1420 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1413) {
		goto L337
	} else {
		goto L338
	}
L336:
	;
	v1436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1426+v1396))))
	v1438 = v1436 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1413) {
		goto L342
	} else {
		goto L343
	}
L337:
	;
	v1426 = v1404 + int32(2)
	if v1426 != v1395 {
		goto L336
	} else {
		goto L340
	}
L338:
	;
	goto L339
L339:
	;
	v1470 = v1413<<(uint(int32(6))%32)&int32(1984) | v1422
	v1471 = int32(2)
	goto L333
L340:
	;
	goto L339
L341:
	;
	v1455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1396+v1442))))
	v1470 = v1455&int32(63) | (v1413<<(uint(int32(18))%32)&int32(1835008) | v1422<<(uint(int32(12))%32) | v1438<<(uint(int32(6))%32))
	v1471 = int32(4)
	goto L333
L342:
	;
	v1442 = v1404 + int32(3)
	if v1442 != v1395 {
		goto L341
	} else {
		goto L345
	}
L343:
	;
	goto L344
L344:
	;
	v1470 = v1413<<(uint(int32(12))%32)&int32(61440) | v1422<<(uint(int32(6))%32) | v1438
	v1471 = int32(3)
	goto L333
L345:
	;
	goto L344
L346:
	;
	v1475 = v1470 - int32(97)
	if v1475 < int32(0) {
		goto L327
	} else {
		goto L347
	}
L347:
	;
	v1481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1475)>>(uint(int32(3))%32)))+uint32(_consts[1067]))))
	if int32(base.Ui32(v1481)>>(uint(v1475&int32(7))%32))&int32(1) == int32(0) {
		goto L327
	} else {
		goto L348
	}
L348:
	;
	v1489 = v1471 + v1404
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1489
	v1404 = v1489
	goto L328
L350:
	;
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1504 = v1503 + v1500
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1504
	v1506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1506)+4)) = v1504
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1529 = v1519
	goto L353
L351:
	;
	if v1624 < int32(0) {
		goto L299
	} else {
		goto L376
	}
L352:
	;
	v1624 = v1596
	goto L351
L353:
	;
	if v1520 <= v1529 {
		goto L355
	} else {
		goto L356
	}
L355:
	;
	v1624 = int32(-1)
	goto L351
L356:
	;
	goto L357
L357:
	;
	v1536 = int32(1)
	v1538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1529+v1521))))
	if base.Ui32(v1538) < base.Ui32(int32(192)) {
		v1595 = v1538
		v1596 = v1536
		goto L358
	} else {
		goto L359
	}
L358:
	;
	if int32(250) < v1595 {
		goto L371
	} else {
		goto L372
	}
L359:
	;
	v1542 = v1529 + int32(1)
	if v1542 == v1520 {
		v1595 = v1538
		v1596 = v1536
		goto L358
	} else {
		goto L360
	}
L360:
	;
	v1545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1542+v1521))))
	v1547 = v1545 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1538) {
		goto L362
	} else {
		goto L363
	}
L361:
	;
	v1561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1551+v1521))))
	v1563 = v1561 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1538) {
		goto L367
	} else {
		goto L368
	}
L362:
	;
	v1551 = v1529 + int32(2)
	if v1551 != v1520 {
		goto L361
	} else {
		goto L365
	}
L363:
	;
	goto L364
L364:
	;
	v1595 = v1538<<(uint(int32(6))%32)&int32(1984) | v1547
	v1596 = int32(2)
	goto L358
L365:
	;
	goto L364
L366:
	;
	v1580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1521+v1567))))
	v1595 = v1580&int32(63) | (v1538<<(uint(int32(18))%32)&int32(1835008) | v1547<<(uint(int32(12))%32) | v1563<<(uint(int32(6))%32))
	v1596 = int32(4)
	goto L358
L367:
	;
	v1567 = v1529 + int32(3)
	if v1567 != v1520 {
		goto L366
	} else {
		goto L370
	}
L368:
	;
	goto L369
L369:
	;
	v1595 = v1538<<(uint(int32(12))%32)&int32(61440) | v1547<<(uint(int32(6))%32) | v1563
	v1596 = int32(3)
	goto L358
L370:
	;
	goto L369
L371:
	;
	v1613 = v1596 + v1529
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1613
	v1529 = v1613
	goto L353
L372:
	;
	v1600 = v1595 - int32(97)
	if v1600 < int32(0) {
		goto L371
	} else {
		goto L373
	}
L373:
	;
	v1606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1600)>>(uint(int32(3))%32)))+uint32(_consts[1067]))))
	if int32(base.Ui32(v1606)>>(uint(v1600&int32(7))%32))&int32(1) != 0 {
		goto L352
	} else {
		goto L374
	}
L374:
	;
	goto L371
L376:
	;
	v1627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1628 = v1627 + v1624
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1628
	v1642 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1643 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1651 = v1628
	goto L379
L377:
	;
	if v1747 < int32(0) {
		goto L299
	} else {
		goto L401
	}
L378:
	;
	v1747 = v1718
	goto L377
L379:
	;
	if v1642 <= v1651 {
		goto L381
	} else {
		goto L382
	}
L381:
	;
	v1747 = int32(-1)
	goto L377
L382:
	;
	goto L383
L383:
	;
	v1658 = int32(1)
	v1660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1651+v1643))))
	if base.Ui32(v1660) < base.Ui32(int32(192)) {
		v1717 = v1660
		v1718 = v1658
		goto L384
	} else {
		goto L385
	}
L384:
	;
	if int32(250) < v1717 {
		goto L378
	} else {
		goto L397
	}
L385:
	;
	v1664 = v1651 + int32(1)
	if v1664 == v1642 {
		v1717 = v1660
		v1718 = v1658
		goto L384
	} else {
		goto L386
	}
L386:
	;
	v1667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1664+v1643))))
	v1669 = v1667 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1660) {
		goto L388
	} else {
		goto L389
	}
L387:
	;
	v1683 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1673+v1643))))
	v1685 = v1683 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1660) {
		goto L393
	} else {
		goto L394
	}
L388:
	;
	v1673 = v1651 + int32(2)
	if v1673 != v1642 {
		goto L387
	} else {
		goto L391
	}
L389:
	;
	goto L390
L390:
	;
	v1717 = v1660<<(uint(int32(6))%32)&int32(1984) | v1669
	v1718 = int32(2)
	goto L384
L391:
	;
	goto L390
L392:
	;
	v1702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1643+v1689))))
	v1717 = v1702&int32(63) | (v1660<<(uint(int32(18))%32)&int32(1835008) | v1669<<(uint(int32(12))%32) | v1685<<(uint(int32(6))%32))
	v1718 = int32(4)
	goto L384
L393:
	;
	v1689 = v1651 + int32(3)
	if v1689 != v1642 {
		goto L392
	} else {
		goto L396
	}
L394:
	;
	goto L395
L395:
	;
	v1717 = v1660<<(uint(int32(12))%32)&int32(61440) | v1669<<(uint(int32(6))%32) | v1685
	v1718 = int32(3)
	goto L384
L396:
	;
	goto L395
L397:
	;
	v1722 = v1717 - int32(97)
	if v1722 < int32(0) {
		goto L378
	} else {
		goto L398
	}
L398:
	;
	v1728 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1722)>>(uint(int32(3))%32)))+uint32(_consts[1067]))))
	if int32(base.Ui32(v1728)>>(uint(v1722&int32(7))%32))&int32(1) == int32(0) {
		goto L378
	} else {
		goto L399
	}
L399:
	;
	v1736 = v1718 + v1651
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1736
	v1651 = v1736
	goto L379
L401:
	;
	v1750 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1751 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1750))) = v1751 + v1747
	goto L299
L402:
	;
	v2104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2104
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2104
	v2109 = F_find_among_b(m, l0, int32(4292480), int32(4))
	mBase = m.M
	v2110 = m.ExcPending
	if v2110 != 0 {
		goto L10
	} else {
		goto L509
	}
L403:
	;
	v2071 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2071
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2071
	v2074 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2071 <= v2074 {
		goto L402
	} else {
		goto L501
	}
L404:
	;
	v2028 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2028
	v2030 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2031 = *(*int32)(unsafe.Add(mBase, uint32(v2030)+8))
	if v2031 <= v2028 {
		goto L489
	} else {
		goto L490
	}
L405:
	;
	v1762 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1762+v1756-int32(1)))))
	if v1766&int32(224) != int32(96) {
		goto L404
	} else {
		goto L406
	}
L406:
	;
	if int32(1)<<(uint(v1766)%32)&int32(823330) == int32(0) {
		goto L404
	} else {
		goto L407
	}
L407:
	;
	v1779 = F_find_among_b(m, l0, int32(4288816), int32(45))
	mBase = m.M
	v1780 = m.ExcPending
	if v1780 != 0 {
		goto L10
	} else {
		goto L408
	}
L408:
	;
	if v1779 == int32(0) {
		goto L404
	} else {
		goto L409
	}
L409:
	;
	v1783 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1783
	switch v1779 - int32(1) {
	case 0:
		goto L418
	case 1:
		goto L417
	case 2:
		goto L416
	case 3:
		goto L415
	case 4:
		goto L414
	case 5:
		goto L413
	case 6:
		goto L412
	case 7:
		goto L411
	case 8:
		goto L410
	default:
		goto L403
	}
L410:
	;
	v2005 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2006 = *(*int32)(unsafe.Add(mBase, uint32(v2005)+8))
	if v1783 < v2006 {
		goto L404
	} else {
		goto L483
	}
L411:
	;
	v1969 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1970 = *(*int32)(unsafe.Add(mBase, uint32(v1969)))
	if v1783 < v1970 {
		goto L404
	} else {
		goto L472
	}
L412:
	;
	v1928 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1929 = *(*int32)(unsafe.Add(mBase, uint32(v1928)))
	if v1783 < v1929 {
		goto L404
	} else {
		goto L461
	}
L413:
	;
	v1893 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1894 = *(*int32)(unsafe.Add(mBase, uint32(v1893)))
	if v1783 < v1894 {
		goto L404
	} else {
		goto L451
	}
L414:
	;
	v1821 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1822 = *(*int32)(unsafe.Add(mBase, uint32(v1821)+4))
	if v1783 < v1822 {
		goto L404
	} else {
		goto L431
	}
L415:
	;
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1813 = *(*int32)(unsafe.Add(mBase, uint32(v1812)))
	if v1783 < v1813 {
		goto L404
	} else {
		goto L428
	}
L416:
	;
	v1803 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1804 = *(*int32)(unsafe.Add(mBase, uint32(v1803)))
	if v1783 < v1804 {
		goto L404
	} else {
		goto L425
	}
L417:
	;
	v1794 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1795 = *(*int32)(unsafe.Add(mBase, uint32(v1794)))
	if v1783 < v1795 {
		goto L404
	} else {
		goto L422
	}
L418:
	;
	v1787 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1788 = *(*int32)(unsafe.Add(mBase, uint32(v1787)))
	if v1783 < v1788 {
		goto L404
	} else {
		goto L419
	}
L419:
	;
	v1790 = F_slice_del(m, l0)
	mBase = m.M
	v1791 = m.ExcPending
	if v1791 != 0 {
		goto L10
	} else {
		goto L420
	}
L420:
	;
	if int32(0) <= v1790 {
		goto L403
	} else {
		goto L421
	}
L421:
	;
	v2285 = v1790
	goto L1
L422:
	;
	v1799 = F_slice_from_s(m, l0, int32(3), int32(2185492))
	mBase = m.M
	v1800 = m.ExcPending
	if v1800 != 0 {
		goto L10
	} else {
		goto L423
	}
L423:
	;
	if int32(0) <= v1799 {
		goto L403
	} else {
		goto L424
	}
L424:
	;
	v2285 = v1799
	goto L1
L425:
	;
	v1808 = F_slice_from_s(m, l0, int32(1), int32(2185495))
	mBase = m.M
	v1809 = m.ExcPending
	if v1809 != 0 {
		goto L10
	} else {
		goto L426
	}
L426:
	;
	if int32(0) <= v1808 {
		goto L403
	} else {
		goto L427
	}
L427:
	;
	v2285 = v1808
	goto L1
L428:
	;
	v1817 = F_slice_from_s(m, l0, int32(4), int32(2185496))
	mBase = m.M
	v1818 = m.ExcPending
	if v1818 != 0 {
		goto L10
	} else {
		goto L429
	}
L429:
	;
	if int32(0) <= v1817 {
		goto L403
	} else {
		goto L430
	}
L430:
	;
	v2285 = v1817
	goto L1
L431:
	;
	v1824 = F_slice_del(m, l0)
	mBase = m.M
	v1825 = m.ExcPending
	if v1825 != 0 {
		goto L10
	} else {
		goto L432
	}
L432:
	;
	if v1824 < int32(0) {
		v2285 = v1824
		goto L1
	} else {
		goto L433
	}
L433:
	;
	v1828 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1828
	v1831 = v1828 - int32(1)
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1831 <= v1832 {
		goto L403
	} else {
		goto L434
	}
L434:
	;
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1834+v1831))))
	if v1836&int32(224) != int32(96) {
		goto L403
	} else {
		goto L435
	}
L435:
	;
	if int32(1)<<(uint(v1836)%32)&int32(4718616) == int32(0) {
		goto L403
	} else {
		goto L436
	}
L436:
	;
	v1849 = F_find_among_b(m, l0, int32(4289728), int32(4))
	mBase = m.M
	v1850 = m.ExcPending
	if v1850 != 0 {
		goto L10
	} else {
		goto L437
	}
L437:
	;
	if v1849 == int32(0) {
		goto L403
	} else {
		goto L438
	}
L438:
	;
	v1853 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1853
	v1855 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1856 = *(*int32)(unsafe.Add(mBase, uint32(v1855)))
	if v1853 < v1856 {
		goto L403
	} else {
		goto L439
	}
L439:
	;
	v1858 = F_slice_del(m, l0)
	mBase = m.M
	v1859 = m.ExcPending
	if v1859 != 0 {
		goto L10
	} else {
		goto L440
	}
L440:
	;
	if v1858 < int32(0) {
		v2285 = v1858
		goto L1
	} else {
		goto L441
	}
L441:
	;
	if v1849 != int32(1) {
		goto L403
	} else {
		goto L442
	}
L442:
	;
	v1864 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1864
	v1866 = int32(2)
	v1868 = int32(0)
	v1871 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1864-v1871 < v1866 {
		v1881 = v1868
		goto L444
	} else {
		goto L445
	}
L443:
	;
	if v1881 == int32(0) {
		goto L403
	} else {
		goto L447
	}
L444:
	;
	goto L443
L445:
	;
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1877 = F_memcmp(m, v1874+v1864-v1866, int32(2185500), v1866)
	mBase = m.M
	if v1877 != 0 {
		v1881 = v1868
		goto L444
	} else {
		goto L446
	}
L446:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1864 - v1866
	v1881 = int32(1)
	goto L444
L447:
	;
	v1884 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1884
	v1886 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1887 = *(*int32)(unsafe.Add(mBase, uint32(v1886)))
	if v1884 < v1887 {
		goto L403
	} else {
		goto L448
	}
L448:
	;
	v1889 = F_slice_del(m, l0)
	mBase = m.M
	v1890 = m.ExcPending
	if v1890 != 0 {
		goto L10
	} else {
		goto L449
	}
L449:
	;
	if int32(0) <= v1889 {
		goto L403
	} else {
		goto L450
	}
L450:
	;
	v2285 = v1889
	goto L1
L451:
	;
	v1896 = F_slice_del(m, l0)
	mBase = m.M
	v1897 = m.ExcPending
	if v1897 != 0 {
		goto L10
	} else {
		goto L452
	}
L452:
	;
	if v1896 < int32(0) {
		v2285 = v1896
		goto L1
	} else {
		goto L453
	}
L453:
	;
	v1900 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1900
	v1902 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1900-int32(3) <= v1902 {
		goto L403
	} else {
		goto L454
	}
L454:
	;
	v1906 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1910 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1906+v1900-int32(1)))))
	switch v1910 - int32(101) {
	case 0, 7:
		goto L455
	default:
		goto L403
	}
L455:
	;
	v1915 = F_find_among_b(m, l0, int32(4289808), int32(3))
	mBase = m.M
	v1916 = m.ExcPending
	if v1916 != 0 {
		goto L10
	} else {
		goto L456
	}
L456:
	;
	if v1915 == int32(0) {
		goto L403
	} else {
		goto L457
	}
L457:
	;
	v1919 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1919
	v1921 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1922 = *(*int32)(unsafe.Add(mBase, uint32(v1921)))
	if v1919 < v1922 {
		goto L403
	} else {
		goto L458
	}
L458:
	;
	v1924 = F_slice_del(m, l0)
	mBase = m.M
	v1925 = m.ExcPending
	if v1925 != 0 {
		goto L10
	} else {
		goto L459
	}
L459:
	;
	if int32(0) <= v1924 {
		goto L403
	} else {
		goto L460
	}
L460:
	;
	v2285 = v1924
	goto L1
L461:
	;
	v1931 = F_slice_del(m, l0)
	mBase = m.M
	v1932 = m.ExcPending
	if v1932 != 0 {
		goto L10
	} else {
		goto L462
	}
L462:
	;
	if v1931 < int32(0) {
		v2285 = v1931
		goto L1
	} else {
		goto L463
	}
L463:
	;
	v1935 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1935
	v1938 = v1935 - int32(1)
	v1939 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1938 <= v1939 {
		goto L403
	} else {
		goto L464
	}
L464:
	;
	v1941 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1943 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1941+v1938))))
	if v1943&int32(224) != int32(96) {
		goto L403
	} else {
		goto L465
	}
L465:
	;
	if int32(1)<<(uint(v1943)%32)&int32(4198408) == int32(0) {
		goto L403
	} else {
		goto L466
	}
L466:
	;
	v1956 = F_find_among_b(m, l0, int32(4289872), int32(3))
	mBase = m.M
	v1957 = m.ExcPending
	if v1957 != 0 {
		goto L10
	} else {
		goto L467
	}
L467:
	;
	if v1956 == int32(0) {
		goto L403
	} else {
		goto L468
	}
L468:
	;
	v1960 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1960
	v1962 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1963 = *(*int32)(unsafe.Add(mBase, uint32(v1962)))
	if v1960 < v1963 {
		goto L403
	} else {
		goto L469
	}
L469:
	;
	v1965 = F_slice_del(m, l0)
	mBase = m.M
	v1966 = m.ExcPending
	if v1966 != 0 {
		goto L10
	} else {
		goto L470
	}
L470:
	;
	if int32(0) <= v1965 {
		goto L403
	} else {
		goto L471
	}
L471:
	;
	v2285 = v1965
	goto L1
L472:
	;
	v1972 = F_slice_del(m, l0)
	mBase = m.M
	v1973 = m.ExcPending
	if v1973 != 0 {
		goto L10
	} else {
		goto L473
	}
L473:
	;
	if v1972 < int32(0) {
		v2285 = v1972
		goto L1
	} else {
		goto L474
	}
L474:
	;
	v1976 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1976
	v1978 = int32(2)
	v1980 = int32(0)
	v1983 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1976-v1983 < v1978 {
		v1993 = v1980
		goto L476
	} else {
		goto L477
	}
L475:
	;
	if v1993 == int32(0) {
		goto L403
	} else {
		goto L479
	}
L476:
	;
	goto L475
L477:
	;
	v1986 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1989 = F_memcmp(m, v1986+v1976-v1978, int32(2185502), v1978)
	mBase = m.M
	if v1989 != 0 {
		v1993 = v1980
		goto L476
	} else {
		goto L478
	}
L478:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1976 - v1978
	v1993 = int32(1)
	goto L476
L479:
	;
	v1996 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1996
	v1998 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1999 = *(*int32)(unsafe.Add(mBase, uint32(v1998)))
	if v1996 < v1999 {
		goto L403
	} else {
		goto L480
	}
L480:
	;
	v2001 = F_slice_del(m, l0)
	mBase = m.M
	v2002 = m.ExcPending
	if v2002 != 0 {
		goto L10
	} else {
		goto L481
	}
L481:
	;
	if int32(0) <= v2001 {
		goto L403
	} else {
		goto L482
	}
L482:
	;
	v2285 = v2001
	goto L1
L483:
	;
	v2008 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1783 <= v2008 {
		goto L404
	} else {
		goto L484
	}
L484:
	;
	v2010 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2014 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2010+v1783-int32(1)))))
	if v2014 != int32(101) {
		goto L404
	} else {
		goto L485
	}
L485:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1783 - int32(1)
	v2022 = F_slice_from_s(m, l0, int32(2), int32(2185504))
	mBase = m.M
	v2023 = m.ExcPending
	if v2023 != 0 {
		goto L10
	} else {
		goto L486
	}
L486:
	;
	if int32(0) <= v2022 {
		goto L403
	} else {
		goto L487
	}
L487:
	;
	v2285 = v2022
	goto L1
L488:
	;
	v2061 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2061
	v2063 = F_slice_del(m, l0)
	mBase = m.M
	v2064 = m.ExcPending
	if v2064 != 0 {
		goto L10
	} else {
		goto L499
	}
L489:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2028
	v2034 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2031
	v2038 = F_find_among_b(m, l0, int32(4289936), int32(120))
	mBase = m.M
	v2039 = m.ExcPending
	if v2039 != 0 {
		goto L10
	} else {
		goto L492
	}
L490:
	;
	v2042 = v2028
	goto L491
L491:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2042
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2042
	v2048 = F_find_among_b(m, l0, int32(4292336), int32(7))
	mBase = m.M
	v2049 = m.ExcPending
	if v2049 != 0 {
		goto L10
	} else {
		goto L494
	}
L492:
	;
	if v2038 != 0 {
		goto L488
	} else {
		goto L493
	}
L493:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2034
	v2041 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2042 = v2041
	goto L491
L494:
	;
	if v2048 == int32(0) {
		goto L402
	} else {
		goto L495
	}
L495:
	;
	v2052 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2052
	v2054 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2055 = *(*int32)(unsafe.Add(mBase, uint32(v2054)+8))
	if v2052 < v2055 {
		goto L402
	} else {
		goto L496
	}
L496:
	;
	v2057 = F_slice_del(m, l0)
	mBase = m.M
	v2058 = m.ExcPending
	if v2058 != 0 {
		goto L10
	} else {
		goto L497
	}
L497:
	;
	if int32(0) <= v2057 {
		goto L402
	} else {
		goto L498
	}
L498:
	;
	v2285 = v2057
	goto L1
L499:
	;
	if v2063 < int32(0) {
		v2285 = v2063
		goto L1
	} else {
		goto L500
	}
L500:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2034
	goto L403
L501:
	;
	v2076 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2077 = v2076 + v2071
	v2080 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2077-int32(1)))))
	if v2080 != int32(105) {
		goto L402
	} else {
		goto L502
	}
L502:
	;
	v2084 = v2071 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2084
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2084
	if v2084 <= v2074 {
		goto L402
	} else {
		goto L503
	}
L503:
	;
	v2090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2077-int32(2)))))
	if v2090 != int32(99) {
		goto L402
	} else {
		goto L504
	}
L504:
	;
	v2093 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2094 = *(*int32)(unsafe.Add(mBase, uint32(v2093)+8))
	if v2071 <= v2094 {
		goto L402
	} else {
		goto L505
	}
L505:
	;
	v2096 = F_slice_del(m, l0)
	mBase = m.M
	v2097 = m.ExcPending
	if v2097 != 0 {
		goto L10
	} else {
		goto L506
	}
L506:
	;
	if v2096 < int32(0) {
		v2285 = v2096
		goto L1
	} else {
		goto L507
	}
L507:
	;
	goto L402
L508:
	;
	v2180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2180
	v2182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2184 = v2180
	v2185 = v2182
	goto L530
L509:
	;
	if v2109 == int32(0) {
		goto L508
	} else {
		goto L510
	}
L510:
	;
	v2113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2113
	switch v2109 - int32(1) {
	case 0:
		goto L512
	case 1:
		goto L511
	default:
		goto L508
	}
L511:
	;
	v2171 = F_slice_from_s(m, l0, int32(1), int32(2186308))
	mBase = m.M
	v2172 = m.ExcPending
	if v2172 != 0 {
		goto L10
	} else {
		goto L528
	}
L512:
	;
	v2117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2118 = *(*int32)(unsafe.Add(mBase, uint32(v2117)+8))
	if v2113 < v2118 {
		goto L508
	} else {
		goto L513
	}
L513:
	;
	v2120 = F_slice_del(m, l0)
	mBase = m.M
	v2121 = m.ExcPending
	if v2121 != 0 {
		goto L10
	} else {
		goto L514
	}
L514:
	;
	if v2120 < int32(0) {
		v2285 = v2120
		goto L1
	} else {
		goto L515
	}
L515:
	;
	v2124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2124
	v2126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2124 <= v2126 {
		goto L508
	} else {
		goto L516
	}
L516:
	;
	v2128 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2129 = v2128 + v2124
	v2131 = v2129 - int32(1)
	v2132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2131))))
	if v2132 != int32(117) {
		goto L518
	} else {
		goto L519
	}
L517:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2160
	v2162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2163 = *(*int32)(unsafe.Add(mBase, uint32(v2162)+8))
	if v2160 < v2163 {
		goto L508
	} else {
		goto L525
	}
L518:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2124
	v2147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2131))))
	if v2147 != int32(105) {
		goto L508
	} else {
		goto L522
	}
L519:
	;
	v2136 = v2124 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2136
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2136
	if v2136 <= v2126 {
		goto L518
	} else {
		goto L520
	}
L520:
	;
	v2142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2129-int32(2)))))
	if v2142 == int32(103) {
		v2160 = v2136
		goto L517
	} else {
		goto L521
	}
L521:
	;
	goto L518
L522:
	;
	v2151 = v2124 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2151
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2151
	if v2151 <= v2126 {
		goto L508
	} else {
		goto L523
	}
L523:
	;
	v2157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2129-int32(2)))))
	if v2157 != int32(99) {
		goto L508
	} else {
		goto L524
	}
L524:
	;
	v2160 = v2151
	goto L517
L525:
	;
	v2165 = F_slice_del(m, l0)
	mBase = m.M
	v2166 = m.ExcPending
	if v2166 != 0 {
		goto L10
	} else {
		goto L526
	}
L526:
	;
	if int32(0) <= v2165 {
		goto L508
	} else {
		goto L527
	}
L527:
	;
	v2285 = v2165
	goto L1
L528:
	;
	if v2171 < int32(0) {
		v2285 = v2171
		goto L1
	} else {
		goto L529
	}
L529:
	;
	goto L508
L530:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2184
	v2191 = v2184 + int32(1)
	if v2191 < v2185 {
		goto L536
	} else {
		goto L537
	}
L531:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2180
	v2285 = int32(1)
	goto L1
L532:
	;
	goto L531
L533:
	;
	v2280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2184 = v2281
	v2185 = v2280
	goto L530
L534:
	;
	v2222 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L550
L535:
	;
	v2201 = F_find_among(m, l0, int32(4292560), int32(3))
	mBase = m.M
	v2202 = m.ExcPending
	if v2202 != 0 {
		goto L10
	} else {
		goto L540
	}
L536:
	;
	v2193 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2193+v2191))))
	if v2195 == int32(126) {
		goto L535
	} else {
		goto L539
	}
L537:
	;
	goto L538
L538:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2184
	v2220 = v2184
	v2221 = v2185
	goto L534
L539:
	;
	goto L538
L540:
	;
	v2203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2203
	switch v2201 - int32(1) {
	case 0:
		goto L543
	case 1:
		goto L542
	case 2:
		goto L541
	default:
		goto L533
	}
L541:
	;
	v2219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2220 = v2203
	v2221 = v2219
	goto L534
L542:
	;
	v2215 = F_slice_from_s(m, l0, int32(2), int32(2186318))
	mBase = m.M
	v2216 = m.ExcPending
	if v2216 != 0 {
		goto L10
	} else {
		goto L546
	}
L543:
	;
	v2209 = F_slice_from_s(m, l0, int32(2), int32(2186316))
	mBase = m.M
	v2210 = m.ExcPending
	if v2210 != 0 {
		goto L10
	} else {
		goto L544
	}
L544:
	;
	if int32(0) <= v2209 {
		goto L533
	} else {
		goto L545
	}
L545:
	;
	v2285 = v2209
	goto L1
L546:
	;
	if int32(0) <= v2215 {
		goto L533
	} else {
		goto L547
	}
L547:
	;
	v2285 = v2215
	goto L1
L548:
	;
	if v2274 < int32(0) {
		goto L532
	} else {
		goto L568
	}
L550:
	;
	goto L551
L551:
	;
	goto L552
L552:
	;
	v2229 = v2220
	v2231 = int32(1)
	goto L555
L554:
	;
	v2274 = v2259
	goto L548
L555:
	;
	if v2221 <= v2229 {
		goto L557
	} else {
		goto L558
	}
L556:
	;
	goto L554
L557:
	;
	v2274 = int32(-1)
	goto L548
L558:
	;
	goto L559
L559:
	;
	v2236 = v2229 + int32(1)
	v2238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2222+v2229))))
	if base.Ui32(v2238) < base.Ui32(int32(192)) {
		v2259 = v2236
		goto L560
	} else {
		goto L561
	}
L560:
	;
	v2260 = int32(1)
	if v2260 < v2231 {
		v2229 = v2259
		v2231 = v2231 - v2260
		goto L555
	} else {
		goto L567
	}
L561:
	;
	if v2221 <= v2236 {
		v2259 = v2236
		goto L560
	} else {
		goto L562
	}
L562:
	;
	v2245 = v2236
	goto L563
L563:
	;
	v2248 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2222+v2245))))
	if int32(-65) < v2248 {
		v2259 = v2245
		goto L560
	} else {
		goto L565
	}
L564:
	;
	v2259 = v2221
	goto L560
L565:
	;
	v2252 = v2245 + int32(1)
	if v2252 != v2221 {
		v2245 = v2252
		goto L563
	} else {
		goto L566
	}
L566:
	;
	goto L564
L567:
	;
	goto L556
L568:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2274
	goto L533
}
func F_posix_fadvise(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32 {
	var v6 int32
	_ = v6
	v6 = m.Env.X__syscall_fadvise64(m, l0, l1, l2, l3)
	return int32(0) - v6
}
func F_preadv(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = m.Wasi_snapshot_preview1.Fd_pread(m, l0, l1, l2, l3, v8+int32(12))
	mBase = m.M
	if v12 == int32(0) {
		v19 = int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[166])) = v12
		v19 = int32(-1)
	}
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	m.G0 = v8 + int32(16)
	if v19 != 0 {
		v25 = int32(-1)
	} else {
		v25 = v20
	}
	return v25
}
func F_prefixsel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 float64
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v10 = F_patternsel_common(m, v2, v3, v4, v5, v6, v7, int32(4), v4)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_Float8GetDatum(m, v10)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			return v14
		}
	}
}
func F_preprocess_pubobj_list(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	v3 = int32(0)
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L20
	} else {
		goto L52
	}
L2:
	;
	return
L3:
	;
	v9 = int32(3)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v12 == v9 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 <= int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v21 = v9
	v22 = v3
	goto L6
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+v22<<(uint(int32(2))%32))))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v29 == int32(3) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L2
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v21
	v33 = v21
	goto L10
L9:
	;
	v33 = v29
	goto L10
L10:
	;
	switch v33 {
	case 0:
		goto L17
	case 1, 2:
		goto L16
	default:
		v141 = v33
		goto L11
	}
L11:
	;
	v143 = v22 + int32(1)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v143 < v144 {
		v21 = v141
		v22 = v143
		goto L6
	} else {
		goto L51
	}
L12:
	;
	v138 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v138
	v141 = v138
	goto L11
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L20
	} else {
		goto L46
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L20
	} else {
		goto L41
	}
L15:
	;
	v86 = F_palloc0(m, int32(16))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L20
	} else {
		goto L39
	}
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	if v56 != 0 {
		goto L27
	} else {
		goto L28
	}
L17:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	if v34 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	if v36 != 0 {
		v141 = int32(0)
		goto L11
	} else {
		goto L19
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return
L21:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	F_errmsg(m, int32(377693), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	F_scanner_errposition(m, v48, l1)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(26882), int32(19607), int32(73386))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L20
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L20
	} else {
		goto L34
	}
L27:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	if v57 != 0 {
		goto L14
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	if v62 != 0 {
		goto L12
	} else {
		goto L33
	}
L30:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v58 != 0 {
		goto L13
	} else {
		goto L31
	}
L31:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	if v59 == int32(0) {
		goto L26
	} else {
		goto L32
	}
L32:
	;
	goto L12
L33:
	;
	v63 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v63
	v141 = v63
	goto L11
L34:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L20
	} else {
		goto L35
	}
L35:
	;
	F_errmsg(m, int32(377712), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L20
	} else {
		goto L36
	}
L36:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	F_scanner_errposition(m, v77, l1)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L20
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(26882), int32(19649), int32(73386))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L20
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = int32(259)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v93 = F_makeRangeVar(m, int32(0), v91, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L20
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+4)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v86
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v141 = v99
	goto L11
L41:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L20
	} else {
		goto L42
	}
L42:
	;
	F_errmsg(m, int32(500694), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L20
	} else {
		goto L43
	}
L43:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	F_scanner_errposition(m, v111, l1)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L20
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(26882), int32(19628), int32(73386))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L20
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L20
	} else {
		goto L47
	}
L47:
	;
	F_errmsg(m, int32(500650), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L20
	} else {
		goto L48
	}
L48:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	F_scanner_errposition(m, v130, l1)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L20
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(26882), int32(19635), int32(73386))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L20
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	goto L7
L52:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L20
	} else {
		goto L53
	}
L53:
	;
	F_errmsg(m, int32(73855), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L20
	} else {
		goto L54
	}
L54:
	;
	F_errdetail(m, int32(614582), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L20
	} else {
		goto L55
	}
L55:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	F_scanner_errposition(m, v167, l1)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L20
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(26882), int32(19591), int32(73386))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L20
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_process_postgres_switches(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	var v616 int32
	_ = v616
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v644 int32
	_ = v644
	v9 = m.G0
	v11 = v9 - int32(112)
	m.G0 = v11
	if l2 != int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _consts[668])) = int32(0)
	goto L22
L2:
	;
	v52 = l0
	v53 = l1
	v54 = int32(9)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v16 = int32(4)
	if l0 < int32(2) {
		v52 = l0
		v53 = l1
		v54 = v16
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v20 = l1 + int32(4)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v22 = int32(385947)
	v25 = int32(*(*uint8)(unsafe.Add(mBase, _consts[669])))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v26 == int32(0) {
		v45 = v25
		v46 = v26
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v47 != 0 {
		goto L14
	} else {
		goto L15
	}
L7:
	;
	v47 = v46 - v45
	goto L6
L8:
	;
	if v25 != v26 {
		v45 = v25
		v46 = v26
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v30 = v21
	v31 = v22
	goto L10
L10:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
	if v35 == int32(0) {
		v45 = v34
		v46 = v35
		goto L7
	} else {
		goto L12
	}
L11:
	;
	v45 = v34
	v46 = v35
	goto L7
L12:
	;
	v38 = int32(1)
	if v34 == v35 {
		v30 = v30 + v38
		v31 = v31 + v38
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v48 = l1
	goto L16
L15:
	;
	v48 = v20
	goto L16
L16:
	;
	v52 = l0 - base.B2i32(v47 == int32(0))
	v53 = v48
	v54 = v16
	goto L1
L17:
	;
	v623 = *(*int32)(unsafe.Add(mBase, _consts[670]))
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v53+v623<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v627
	F_errmsg(m, int32(198298), v11+int32(16))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L50
	} else {
		goto L195
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v127
	F_errmsg(m, int32(342954), v11-int32(-64))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L50
	} else {
		goto L193
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L50
	} else {
		goto L189
	}
L20:
	;
	v553 = int32(*(*uint8)(unsafe.Add(mBase, _consts[123])))
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L50
	} else {
		goto L183
	}
L21:
	;
	v544 = int32(4369908)
	v546 = *(*int32)(unsafe.Add(mBase, _consts[670]))
	*(*int32)(unsafe.Add(mBase, _consts[670])) = v546 - int32(1)
	goto L20
L22:
	;
	v69 = F_getopt(m, v52, v53, int32(541379))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L50
	} else {
		goto L51
	}
L23:
	;
	if l3 == int32(0) {
		goto L174
	} else {
		goto L175
	}
L24:
	;
	goto L23
L25:
	;
	v498 = *(*int32)(unsafe.Add(mBase, _consts[671]))
	F_SetConfigOption(m, int32(26507), v498, l2, v54)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L50
	} else {
		goto L173
	}
L26:
	;
	if l2 != int32(1) {
		goto L22
	} else {
		goto L156
	}
L27:
	;
	v429 = *(*int32)(unsafe.Add(mBase, _consts[671]))
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429))))
	switch v430 - int32(101) {
	case 0:
		v438 = int32(124496)
		goto L152
	default:
		goto L21
	case 11:
		goto L153
	}
L28:
	;
	F_SetConfigOption(m, int32(124455), int32(340544), l2, v54)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L50
	} else {
		goto L151
	}
L29:
	;
	v420 = *(*int32)(unsafe.Add(mBase, _consts[671]))
	F_SetConfigOption(m, int32(288337), v420, l2, v54)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L50
	} else {
		goto L150
	}
L30:
	;
	if l2 != int32(1) {
		goto L22
	} else {
		goto L117
	}
L31:
	;
	v295 = *(*int32)(unsafe.Add(mBase, _consts[671]))
	F_SetConfigOption(m, int32(80430), v295, l2, v54)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L50
	} else {
		goto L116
	}
L32:
	;
	F_SetConfigOption(m, int32(155933), int32(340544), l2, v54)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L50
	} else {
		goto L115
	}
L33:
	;
	F_SetConfigOption(m, int32(170507), int32(340544), l2, v54)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L50
	} else {
		goto L114
	}
L34:
	;
	v282 = *(*int32)(unsafe.Add(mBase, _consts[671]))
	F_SetConfigOption(m, int32(140101), v282, l2, v54)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L50
	} else {
		goto L113
	}
L35:
	;
	F_SetConfigOption(m, int32(297075), int32(340544), l2, v54)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L50
	} else {
		goto L112
	}
L36:
	;
	v273 = *(*int32)(unsafe.Add(mBase, _consts[671]))
	F_SetConfigOption(m, int32(166387), v273, l2, v54)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L50
	} else {
		goto L111
	}
L37:
	;
	if l2 != int32(1) {
		goto L22
	} else {
		goto L110
	}
L38:
	;
	F_SetConfigOption(m, int32(159534), int32(648495), l2, v54)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L50
	} else {
		goto L109
	}
L39:
	;
	v259 = *(*int32)(unsafe.Add(mBase, _consts[671]))
	F_SetConfigOption(m, int32(159534), v259, l2, v54)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L50
	} else {
		goto L108
	}
L40:
	;
	v233 = *(*int32)(unsafe.Add(mBase, _consts[671]))
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233))))
	v236 = v234 - int32(98)
	v238 = v236 & int32(255)
	if base.Ui32(int32(18)) < base.Ui32(v238) {
		goto L21
	} else {
		goto L105
	}
L41:
	;
	F_SetConfigOption(m, int32(483747), int32(357375), l2, v54)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L50
	} else {
		goto L104
	}
L42:
	;
	F_SetConfigOption(m, int32(378284), int32(237184), l2, v54)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L50
	} else {
		goto L103
	}
L43:
	;
	if l2 != int32(1) {
		goto L22
	} else {
		goto L102
	}
L44:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _consts[671]))
	v172 = v168
	goto L86
L45:
	;
	if l2 != int32(1) {
		goto L22
	} else {
		goto L80
	}
L46:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _consts[671]))
	F_ParseLongOption(m, v109, v11+int32(108), v11+int32(104))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L50
	} else {
		goto L68
	}
L47:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _consts[671]))
	v86 = F_strcmp(m, int32(315003), v84)
	mBase = m.M
	if v86 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L48:
	;
	if l2 != int32(1) {
		goto L22
	} else {
		goto L53
	}
L49:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _consts[671]))
	F_SetConfigOption(m, int32(133801), v75, l2, v54)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L50
	} else {
		goto L52
	}
L50:
	;
	return
L51:
	;
	switch v69 + int32(1) {
	case 0:
		goto L24
	default:
		goto L21
	case 46:
		goto L47
	case 67:
		goto L49
	case 68, 85, 111:
		goto L22
	case 69:
		goto L45
	case 70:
		goto L43
	case 71:
		goto L41
	case 79:
		goto L34
	case 80:
		goto L33
	case 81:
		goto L32
	case 84:
		goto L29
	case 88:
		goto L25
	case 99:
		goto L48
	case 100:
		goto L46
	case 101:
		goto L44
	case 102:
		goto L42
	case 103:
		goto L40
	case 105:
		goto L39
	case 106:
		goto L38
	case 107:
		goto L37
	case 108:
		goto L36
	case 109:
		goto L35
	case 113:
		goto L31
	case 115:
		goto L30
	case 116:
		goto L28
	case 117:
		goto L27
	case 119:
		goto L26
	}
L52:
	;
	goto L22
L53:
	;
	v81 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[200])) = uint8(v81)
	goto L22
L54:
	;
	if v105 != int32(5) {
		goto L19
	} else {
		goto L67
	}
L55:
	;
	v105 = int32(0)
	goto L54
L56:
	;
	goto L57
L57:
	;
	v91 = F_strcmp(m, int32(83844), v84)
	mBase = m.M
	if v91 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v105 = int32(1)
	goto L54
L59:
	;
	goto L60
L60:
	;
	v96 = F_strcmp(m, int32(333768), v84)
	mBase = m.M
	if v96 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v105 = int32(3)
	goto L54
L62:
	;
	goto L63
L63:
	;
	v103 = F_strcmp(m, int32(385949), v84)
	mBase = m.M
	if v103 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v104 = int32(5)
	goto L66
L65:
	;
	v104 = int32(4)
	goto L66
L66:
	;
	v105 = v104
	goto L54
L67:
	;
	goto L46
L68:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
	if v116 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L50
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	F_SetConfigOption(m, v141, v116, l2, v54)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L50
	} else {
		goto L77
	}
L72:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L50
	} else {
		goto L73
	}
L73:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _consts[671]))
	if v69 == int32(45) {
		goto L18
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v127
	F_errmsg(m, int32(342976), v11+int32(80))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L50
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(488689), int32(3985), int32(167054))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L50
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	F_pfree(m, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L50
	} else {
		goto L78
	}
L78:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
	F_pfree(m, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L50
	} else {
		goto L79
	}
L79:
	;
	goto L22
L80:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _consts[671]))
	v157 = F_strlen(m, v154)
	mBase = m.M
	v159 = v157 + int32(1)
	v160 = F_emscripten_builtin_malloc(m, v159)
	mBase = m.M
	if v160 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, _consts[672])) = v165
	goto L22
L82:
	;
	v165 = int32(0)
	goto L81
L83:
	;
	goto L84
L84:
	;
	v164 = F___memcpy(m, v160, v154, v159)
	mBase = m.M
	v165 = v164
	goto L81
L85:
	;
	F_set_debug_options(m, v216, l2, v54)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L50
	} else {
		goto L101
	}
L86:
	;
	v177 = v172 + int32(1)
	v178 = int32(*(*int8)(unsafe.Add(mBase, uint32(v172))))
	v179 = F___isspace(m, v178)
	mBase = m.M
	if v179 != 0 {
		v172 = v177
		goto L86
	} else {
		goto L88
	}
L87:
	;
	v180 = int32(1)
	switch v178&int32(255) - int32(43) {
	case 0:
		v186 = v180
		goto L90
	default:
		v188 = v178
		v189 = v172
		v190 = v180
		goto L89
	case 2:
		goto L91
	}
L88:
	;
	goto L87
L89:
	;
	v191 = int32(0)
	v193 = v188 - int32(48)
	if base.Ui32(v193) <= base.Ui32(int32(9)) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	v187 = int32(*(*int8)(unsafe.Add(mBase, uint32(v177))))
	v188 = v187
	v189 = v177
	v190 = v186
	goto L89
L91:
	;
	v186 = int32(0)
	goto L90
L92:
	;
	v196 = v191
	v197 = v193
	v198 = v189
	goto L95
L93:
	;
	v210 = v191
	goto L94
L94:
	;
	if v190 != 0 {
		goto L98
	} else {
		goto L99
	}
L95:
	;
	v200 = int32(10)
	v202 = v196*v200 - v197
	v203 = int32(*(*int8)(unsafe.Add(mBase, uint32(v198)+1)))
	v207 = v203 - int32(48)
	if base.Ui32(v207) < base.Ui32(v200) {
		v196 = v202
		v197 = v207
		v198 = v198 + int32(1)
		goto L95
	} else {
		goto L97
	}
L96:
	;
	v210 = v202
	goto L94
L97:
	;
	goto L96
L98:
	;
	v216 = int32(0) - v210
	goto L100
L99:
	;
	v216 = v210
	goto L100
L100:
	;
	goto L85
L101:
	;
	goto L22
L102:
	;
	v222 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[673])) = uint8(v222)
	goto L22
L103:
	;
	goto L22
L104:
	;
	goto L22
L105:
	;
	if int32(base.Ui32(int32(407745))>>(uint(v238)%32))&int32(1) == int32(0) {
		goto L21
	} else {
		goto L106
	}
L106:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v236&int32(255)<<(uint(int32(2))%32))+uint32(_consts[674])))
	F_SetConfigOption(m, v253, int32(357375), l2, v54)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L50
	} else {
		goto L107
	}
L107:
	;
	goto L22
L108:
	;
	goto L22
L109:
	;
	goto L22
L110:
	;
	v269 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[675])) = uint8(v269)
	goto L22
L111:
	;
	goto L22
L112:
	;
	goto L22
L113:
	;
	goto L22
L114:
	;
	goto L22
L115:
	;
	goto L22
L116:
	;
	goto L22
L117:
	;
	v300 = int32(4470976)
	v302 = *(*int32)(unsafe.Add(mBase, _consts[671]))
	goto L121
L118:
	;
	goto L22
L119:
	;
	v415 = F_strlen(m, v404)
	mBase = m.M
	goto L118
L121:
	;
	goto L122
L122:
	;
	v309 = int32(1023)
	if (v300^v302)&int32(3) != 0 {
		goto L126
	} else {
		goto L127
	}
L123:
	;
	v408 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v405))) = uint8(v408)
	goto L119
L124:
	;
	v389 = v384
	v390 = v385
	v391 = v386
	goto L146
L125:
	;
	if v379 == int32(0) {
		v404 = v377
		v405 = v378
		goto L123
	} else {
		goto L145
	}
L126:
	;
	v377 = v302
	v378 = v300
	v379 = v309
	goto L125
L127:
	;
	goto L128
L128:
	;
	if v302&int32(3) == int32(0) {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	if v346 == int32(0) {
		v404 = v343
		v405 = v344
		goto L123
	} else {
		goto L138
	}
L130:
	;
	v343 = v302
	v344 = v300
	v345 = v309
	v346 = int32(1)
	goto L129
L131:
	;
	goto L132
L132:
	;
	v322 = v302
	v323 = v300
	v324 = v309
	goto L133
L133:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322))))
	*(*uint8)(unsafe.Add(mBase, uint32(v323))) = uint8(v326)
	if v326 == int32(0) {
		v384 = v322
		v385 = v323
		v386 = v324
		goto L124
	} else {
		goto L135
	}
L134:
	;
	v343 = v337
	v344 = v331
	v345 = v333
	v346 = v335
	goto L129
L135:
	;
	v330 = int32(1)
	v331 = v323 + v330
	v333 = v324 - v330
	v334 = int32(0)
	v335 = base.B2i32(v333 != v334)
	v337 = v322 + v330
	if v337&int32(3) == v334 {
		v343 = v337
		v344 = v331
		v345 = v333
		v346 = v335
		goto L129
	} else {
		goto L136
	}
L136:
	;
	if v333 != 0 {
		v322 = v337
		v323 = v331
		v324 = v333
		goto L133
	} else {
		goto L137
	}
L137:
	;
	goto L134
L138:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343))))
	if v349 == int32(0) {
		v377 = v343
		v378 = v344
		v379 = v345
		goto L125
	} else {
		goto L139
	}
L139:
	;
	if base.Ui32(v345) < base.Ui32(int32(4)) {
		v377 = v343
		v378 = v344
		v379 = v345
		goto L125
	} else {
		goto L140
	}
L140:
	;
	v355 = v343
	v356 = v344
	v357 = v345
	goto L141
L141:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
	v363 = int32(-2139062144)
	if (int32(16843008)-v360|v360)&v363 != v363 {
		v384 = v355
		v385 = v356
		v386 = v357
		goto L124
	} else {
		goto L143
	}
L142:
	;
	v377 = v371
	v378 = v369
	v379 = v373
	goto L125
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v356))) = v360
	v368 = int32(4)
	v369 = v356 + v368
	v371 = v355 + v368
	v373 = v357 - v368
	if base.Ui32(int32(3)) < base.Ui32(v373) {
		v355 = v371
		v356 = v369
		v357 = v373
		goto L141
	} else {
		goto L144
	}
L144:
	;
	goto L142
L145:
	;
	v384 = v377
	v385 = v378
	v386 = v379
	goto L124
L146:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389))))
	*(*uint8)(unsafe.Add(mBase, uint32(v390))) = uint8(v393)
	if v393 == int32(0) {
		v404 = v389
		v405 = v390
		goto L123
	} else {
		goto L148
	}
L147:
	;
	v404 = v400
	v405 = v398
	goto L123
L148:
	;
	v397 = int32(1)
	v398 = v390 + v397
	v400 = v389 + v397
	v402 = v391 - v397
	if v402 != 0 {
		v389 = v400
		v390 = v398
		v391 = v402
		goto L146
	} else {
		goto L149
	}
L149:
	;
	goto L147
L150:
	;
	goto L22
L151:
	;
	goto L22
L152:
	;
	F_SetConfigOption(m, v438, int32(340544), l2, v54)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L50
	} else {
		goto L155
	}
L153:
	;
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429)+1)))
	switch v434 - int32(97) {
	case 0:
		v438 = int32(124538)
		goto L152
	default:
		goto L21
	case 11:
		goto L154
	}
L154:
	;
	v438 = int32(124555)
	goto L152
L155:
	;
	goto L22
L156:
	;
	v446 = *(*int32)(unsafe.Add(mBase, _consts[671]))
	v450 = v446
	goto L158
L157:
	;
	*(*int32)(unsafe.Add(mBase, _consts[648])) = v494
	goto L22
L158:
	;
	v455 = v450 + int32(1)
	v456 = int32(*(*int8)(unsafe.Add(mBase, uint32(v450))))
	v457 = F___isspace(m, v456)
	mBase = m.M
	if v457 != 0 {
		v450 = v455
		goto L158
	} else {
		goto L160
	}
L159:
	;
	v458 = int32(1)
	switch v456&int32(255) - int32(43) {
	case 0:
		v464 = v458
		goto L162
	default:
		v466 = v456
		v467 = v450
		v468 = v458
		goto L161
	case 2:
		goto L163
	}
L160:
	;
	goto L159
L161:
	;
	v469 = int32(0)
	v471 = v466 - int32(48)
	if base.Ui32(v471) <= base.Ui32(int32(9)) {
		goto L164
	} else {
		goto L165
	}
L162:
	;
	v465 = int32(*(*int8)(unsafe.Add(mBase, uint32(v455))))
	v466 = v465
	v467 = v455
	v468 = v464
	goto L161
L163:
	;
	v464 = int32(0)
	goto L162
L164:
	;
	v474 = v469
	v475 = v471
	v476 = v467
	goto L167
L165:
	;
	v488 = v469
	goto L166
L166:
	;
	if v468 != 0 {
		goto L170
	} else {
		goto L171
	}
L167:
	;
	v478 = int32(10)
	v480 = v474*v478 - v475
	v481 = int32(*(*int8)(unsafe.Add(mBase, uint32(v476)+1)))
	v485 = v481 - int32(48)
	if base.Ui32(v485) < base.Ui32(v478) {
		v474 = v480
		v475 = v485
		v476 = v476 + int32(1)
		goto L167
	} else {
		goto L169
	}
L168:
	;
	v488 = v480
	goto L166
L169:
	;
	goto L168
L170:
	;
	v494 = int32(0) - v488
	goto L172
L171:
	;
	v494 = v488
	goto L172
L172:
	;
	goto L157
L173:
	;
	goto L22
L174:
	;
	v531 = *(*int32)(unsafe.Add(mBase, _consts[670]))
	if v52 != v531 {
		goto L20
	} else {
		goto L182
	}
L175:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v503 != 0 {
		goto L174
	} else {
		goto L176
	}
L176:
	;
	v505 = *(*int32)(unsafe.Add(mBase, _consts[670]))
	if v52-v505 <= int32(0) {
		goto L174
	} else {
		goto L177
	}
L177:
	;
	v510 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[670])) = v505 + v510
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v53+v505<<(uint(int32(2))%32))))
	v519 = F_strlen(m, v516)
	mBase = m.M
	v521 = v519 + v510
	v522 = F_emscripten_builtin_malloc(m, v521)
	mBase = m.M
	if v522 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v527
	goto L174
L179:
	;
	v527 = int32(0)
	goto L178
L180:
	;
	goto L181
L181:
	;
	v526 = F___memcpy(m, v522, v516, v521)
	mBase = m.M
	v527 = v526
	goto L178
L182:
	;
	v534 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[676])) = v534
	*(*int32)(unsafe.Add(mBase, _consts[670])) = v534
	m.G0 = v11 + int32(112)
	return
L183:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L50
	} else {
		goto L184
	}
L184:
	;
	if v553 == int32(1) {
		goto L17
	} else {
		goto L185
	}
L185:
	;
	v564 = *(*int32)(unsafe.Add(mBase, _consts[670]))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v53+v564<<(uint(int32(2))%32))))
	v570 = *(*int32)(unsafe.Add(mBase, _consts[677]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v570
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v568
	F_errmsg(m, int32(197831), v11+int32(48))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L50
	} else {
		goto L186
	}
L186:
	;
	v579 = *(*int32)(unsafe.Add(mBase, _consts[677]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v579
	F_errhint(m, int32(595224), v11+int32(32))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L50
	} else {
		goto L187
	}
L187:
	;
	F_errfinish(m, int32(488689), int32(4138), int32(167054))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L50
	} else {
		goto L188
	}
L188:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L189:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L50
	} else {
		goto L190
	}
L190:
	;
	v599 = *(*int32)(unsafe.Add(mBase, _consts[671]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v599
	F_errmsg(m, int32(93199), v11+int32(96))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L50
	} else {
		goto L191
	}
L191:
	;
	F_errfinish(m, int32(488689), int32(3965), int32(167054))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L50
	} else {
		goto L192
	}
L192:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L193:
	;
	F_errfinish(m, int32(488689), int32(3980), int32(167054))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L50
	} else {
		goto L194
	}
L194:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L195:
	;
	v635 = *(*int32)(unsafe.Add(mBase, _consts[677]))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v635
	F_errhint(m, int32(595224), v11)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L50
	} else {
		goto L196
	}
L196:
	;
	F_errfinish(m, int32(488689), int32(4132), int32(167054))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L50
	} else {
		goto L197
	}
L197:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_procsignal_sigusr1_handler(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v229 int32
	_ = v229
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v277 int32
	_ = v277
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	v4 = *(*int32)(unsafe.Add(mBase, _consts[608]))
	if v4 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v288 = *(*int32)(unsafe.Add(mBase, _consts[80]))
	F_SetLatch(m, v288)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L6
	} else {
		goto L72
	}
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+40))
	if v7 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4)+40)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[609])) = int32(1)
	v14 = *(*int32)(unsafe.Add(mBase, _consts[80]))
	F_SetLatch(m, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v21 = v4
	goto L5
L5:
	;
	v23 = v21 + int32(44)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v24 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	return
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _consts[608]))
	if v18 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v21 = v18
	goto L5
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[610])) = int32(1)
	v31 = *(*int32)(unsafe.Add(mBase, _consts[80]))
	F_SetLatch(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L6
	} else {
		goto L12
	}
L10:
	;
	v38 = v21
	goto L11
L11:
	;
	v40 = v38 + int32(48)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v41 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _consts[608]))
	if v35 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v38 = v35
	goto L11
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = int32(0)
	v45 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[1])) = v45
	*(*int32)(unsafe.Add(mBase, _consts[611])) = v45
	v51 = *(*int32)(unsafe.Add(mBase, _consts[80]))
	F_SetLatch(m, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L6
	} else {
		goto L17
	}
L15:
	;
	v58 = v38
	goto L16
L16:
	;
	v60 = v58 + int32(52)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v61 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _consts[608]))
	if v55 == int32(0) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v58 = v55
	goto L16
L19:
	;
	v62 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v62
	v65 = *(*int32)(unsafe.Add(mBase, _consts[612]))
	if v65 == v62 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v80 = v58
	goto L21
L21:
	;
	v82 = v80 + int32(56)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v83 != 0 {
		goto L28
	} else {
		goto L29
	}
L22:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _consts[608]))
	if v77 == int32(0) {
		goto L1
	} else {
		goto L27
	}
L23:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _consts[522]))
	v71 = F_kill(m, v69, int32(15))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L6
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, _consts[622])) = int32(1)
	goto L22
L26:
	;
	goto L22
L27:
	;
	v80 = v77
	goto L21
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = int32(0)
	v87 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[1])) = v87
	*(*int32)(unsafe.Add(mBase, _consts[492])) = v87
	goto L30
L29:
	;
	goto L30
L30:
	;
	v93 = v80 + int32(60)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if v94 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = int32(0)
	v98 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[1])) = v98
	*(*int32)(unsafe.Add(mBase, _consts[494])) = v98
	goto L33
L32:
	;
	goto L33
L33:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _consts[608]))
	if v104 == int32(0) {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v108 = v104 - int32(-64)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	if v109 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108))) = int32(0)
	v113 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[1])) = v113
	*(*int32)(unsafe.Add(mBase, _consts[613])) = v113
	v119 = *(*int32)(unsafe.Add(mBase, _consts[80]))
	F_SetLatch(m, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L6
	} else {
		goto L38
	}
L36:
	;
	v126 = v104
	goto L37
L37:
	;
	v128 = v126 + int32(68)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	if v129 != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _consts[608]))
	if v123 == int32(0) {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v126 = v123
	goto L37
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128))) = int32(0)
	v137 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[614])) = v137
	*(*int32)(unsafe.Add(mBase, _consts[615])) = v137
	*(*int32)(unsafe.Add(mBase, _consts[1])) = v137
	goto L43
L41:
	;
	v149 = v126
	goto L42
L42:
	;
	v151 = v149 + int32(72)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	if v152 != 0 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v146 = *(*int32)(unsafe.Add(mBase, _consts[608]))
	if v146 == int32(0) {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v149 = v146
	goto L42
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151))) = int32(0)
	v160 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[616])) = v160
	*(*int32)(unsafe.Add(mBase, _consts[615])) = v160
	*(*int32)(unsafe.Add(mBase, _consts[1])) = v160
	goto L48
L46:
	;
	v172 = v149
	goto L47
L47:
	;
	v174 = v172 + int32(76)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	if v175 != 0 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _consts[608]))
	if v169 == int32(0) {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v172 = v169
	goto L47
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v174))) = int32(0)
	v183 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[617])) = v183
	*(*int32)(unsafe.Add(mBase, _consts[615])) = v183
	*(*int32)(unsafe.Add(mBase, _consts[1])) = v183
	goto L53
L51:
	;
	v195 = v172
	goto L52
L52:
	;
	v197 = v195 + int32(80)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	if v198 != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v192 = *(*int32)(unsafe.Add(mBase, _consts[608]))
	if v192 == int32(0) {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v195 = v192
	goto L52
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197))) = int32(0)
	v206 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[618])) = v206
	*(*int32)(unsafe.Add(mBase, _consts[615])) = v206
	*(*int32)(unsafe.Add(mBase, _consts[1])) = v206
	goto L58
L56:
	;
	v218 = v195
	goto L57
L57:
	;
	v220 = v218 + int32(84)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	if v221 != 0 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v215 = *(*int32)(unsafe.Add(mBase, _consts[608]))
	if v215 == int32(0) {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v218 = v215
	goto L57
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220))) = int32(0)
	v229 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[619])) = v229
	*(*int32)(unsafe.Add(mBase, _consts[615])) = v229
	*(*int32)(unsafe.Add(mBase, _consts[1])) = v229
	goto L63
L61:
	;
	v241 = v218
	goto L62
L62:
	;
	v243 = v241 + int32(92)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	if v244 != 0 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v238 = *(*int32)(unsafe.Add(mBase, _consts[608]))
	if v238 == int32(0) {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v241 = v238
	goto L62
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v243))) = int32(0)
	v252 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[620])) = v252
	*(*int32)(unsafe.Add(mBase, _consts[615])) = v252
	*(*int32)(unsafe.Add(mBase, _consts[1])) = v252
	goto L68
L66:
	;
	v264 = v241
	goto L67
L67:
	;
	v266 = v264 + int32(88)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	if v267 == int32(0) {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	v261 = *(*int32)(unsafe.Add(mBase, _consts[608]))
	if v261 == int32(0) {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v264 = v261
	goto L67
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v266))) = int32(0)
	v277 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[621])) = v277
	*(*int32)(unsafe.Add(mBase, _consts[615])) = v277
	*(*int32)(unsafe.Add(mBase, _consts[1])) = v277
	goto L71
L71:
	;
	goto L1
L72:
	;
	return
}
func F_prs_process_call(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v11 < v12 {
		v15 = v8 + int32(16)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v15
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v17+v11<<(uint(int32(3))%32))))
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v21
		v26 = F_pg_sprintf(m, v15, int32(482897), v8)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v30+v31<<(uint(int32(3))%32))+4))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v35
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v40 = F_BuildTupleFromCStrings(m, v37, v8+int32(40))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
				v43 = F_HeapTupleHeaderGetDatum(m, v42)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
					F_pfree(m, v45)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v48 + int32(1)
						v52 = v43
						m.G0 = v8 + int32(48)
						return v52
					}
				}
			}
		}
	} else {
		v52 = int32(0)
		m.G0 = v8 + int32(48)
		return v52
	}
}
func F_prs_setup_firstcall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = F_lookup_ts_parser_cache(m, l2)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v19 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v19
	v23 = int32(4476144)
	v24 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v26
	v29 = F_palloc(m, int32(12))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = int64(68719476736)
	v34 = F_palloc(m, int32(128))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v34
	v37 = int32(1)
	v38 = l3 + v37
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	v43 = v41 & v37
	if v43 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v44 = v38
	goto L7
L6:
	;
	v44 = l3 + int32(4)
	goto L7
L7:
	;
	if v41 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v76 = v17 + int32(56)
	v77 = int32(0)
	v79 = F_FunctionCall2Coll(m, v17+int32(28), v77, v44, v74)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L19
	}
L9:
	;
	v49 = int32(4)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v51&int32(254) == int32(2) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v64 = int32(1)
	if v43 != 0 {
		v74 = int32(base.Ui32(v41)>>(uint(v64)%32)) - v64
		goto L8
	} else {
		goto L18
	}
L12:
	;
	v60 = v49
	goto L14
L13:
	;
	v60 = base.B2i32(v51 == int32(18)) << (uint(v49) % 32)
	goto L14
L14:
	;
	if v51 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v63 = v49
	goto L17
L16:
	;
	v63 = v60
	goto L17
L17:
	;
	v74 = v63
	goto L8
L18:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v74 = int32(base.Ui32(v68)>>(uint(int32(2))%32)) - int32(4)
	goto L8
L19:
	;
	v85 = F_FunctionCall3Coll(m, v76, v77, v79, v15+int32(8), v15+int32(4))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if v85 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v90 = v85
	goto L24
L22:
	;
	goto L23
L23:
	;
	v174 = F_FunctionCall1Coll(m, v17+int32(84), int32(0), v79)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L37
	}
L24:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v99 <= v100 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L23
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v99 << (uint(int32(1)) % 32)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v108 = F_repalloc(m, v105, v99<<(uint(int32(4))%32))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v114 = F_palloc(m, v111+int32(1))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v108
	goto L28
L30:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v118 = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v116+v117<<(uint(v118)%32))+4)) = v114
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v122+v123<<(uint(v118)%32))+4))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v129 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v134 = int32(3)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v132+v133<<(uint(v134)%32))+4))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v140 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v137+v138))) = uint8(v140)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	*(*int32)(unsafe.Add(mBase, uint32(v142+v143<<(uint(v134)%32)))) = v90
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v148 + int32(1)
	v157 = F_FunctionCall3Coll(m, v76, v140, v79, v15+int32(8), v15+int32(4))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L35
	}
L32:
	;
	v130 = F__emscripten_memcpy_bulkmem(m, v127, v128, v129)
	mBase = m.M
	goto L34
L33:
	;
	goto L34
L34:
	;
	goto L31
L35:
	;
	if v157 != 0 {
		v90 = v157
		goto L24
	} else {
		goto L36
	}
L36:
	;
	goto L25
L37:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v176
	v178 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v178
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v29
	v184 = F_get_call_result_type(m, l1, v178, v15+int32(12))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v184 != int32(1) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v201
	v203 = F_TupleDescGetAttInMetadata(m, v201)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L45
	}
L42:
	;
	F_errmsg_internal(m, int32(363339), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(489831), int32(209), int32(301654))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v203
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v24
	m.G0 = v15 + int32(16)
	return
}
func F_pull_ands(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	v2 = int32(0)
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v38
L2:
	;
	v9 = v2
	v11 = v2
	goto L7
L3:
	;
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v5 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v38 = v2
	goto L1
L6:
	;
	goto L5
L7:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12+v11<<(uint(int32(2))%32))))
	if v16 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v38 = v32
	goto L1
L9:
	;
	v34 = v11 + int32(1)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v34 < v35 {
		v9 = v32
		v11 = v34
		goto L7
	} else {
		goto L18
	}
L10:
	;
	v30 = F_lappend(m, v9, v16)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L14
	} else {
		goto L17
	}
L11:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v19 != int32(21) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v22 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v24 = F_pull_ands(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	v28 = F_list_concat(m, v9, v24)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v32 = v28
	goto L9
L17:
	;
	v32 = v30
	goto L9
L18:
	;
	goto L8
}
func F_pull_varnos_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	v3 = int32(0)
	if l0 == v3 {
		v214 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v214
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v7 - int32(58) {
	case 0:
		goto L6
	case 1, 2, 3, 4, 5, 6, 7, 8:
		goto L3
	case 9:
		goto L4
	default:
		goto L7
	}
L3:
	;
	v212 = F_expression_tree_walker_impl(m, l0, int32(896), l1)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L11
	} else {
		goto L63
	}
L4:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v196 + int32(1)
	v202 = F_query_tree_walker_impl(m, l0, int32(896), l1, int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L11
	} else {
		goto L62
	}
L5:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v38 != v39 {
		goto L3
	} else {
		goto L16
	}
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v30 != 0 {
		v214 = v3
		goto L1
	} else {
		goto L14
	}
L7:
	;
	if v7 == int32(319) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	if v7 != int32(6) {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v14 != v15 {
		v214 = v3
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = F_bms_add_member(m, v17, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v19
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v25 = F_bms_add_members(m, v19, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v25
	return int32(0)
L14:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v33 = F_bms_add_member(m, v31, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v33
	return int32(0)
L16:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v41 == int32(0) {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	if v38 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v188
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v191 = F_bms_add_members(m, v188, v190)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L11
	} else {
		goto L61
	}
L19:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	v60 = int32(0)
	v67 = base.B2i32(v57|v59 == v60)
	if v57 == v60 {
		v106 = v67
		goto L26
	} else {
		goto L27
	}
L20:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v55 = F_bms_add_members(m, v53, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L11
	} else {
		goto L24
	}
L21:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+148))
	if base.Ui32(v45) <= base.Ui32(v44) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v41)+144))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47+v44<<(uint(int32(2))%32))))
	if v51 != 0 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	v188 = v55
	goto L18
L25:
	;
	if v106 != 0 {
		goto L37
	} else {
		goto L38
	}
L26:
	;
	goto L25
L27:
	;
	if v59 == int32(0) {
		v106 = v67
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v73 != v74 {
		v106 = int32(0)
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v76 = int32(1)
	if v73 <= v76 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v79 = v76
	goto L32
L31:
	;
	v79 = v73
	goto L32
L32:
	;
	v80 = int32(8)
	v85 = int32(0)
	goto L33
L33:
	;
	v93 = v85 << (uint(int32(2)) % 32)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v57+v80+v93)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v93+(v59+v80))))
	v98 = base.B2i32(v95 == v97)
	if v97 != v95 {
		v106 = v98
		goto L26
	} else {
		goto L35
	}
L34:
	;
	v106 = v98
	goto L26
L35:
	;
	v101 = v85 + int32(1)
	if v101 != v79 {
		v85 = v101
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v112 = F_bms_add_members(m, v110, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L11
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+8))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v117 = F_bms_difference(m, v115, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L11
	} else {
		goto L41
	}
L40:
	;
	v188 = v112
	goto L18
L41:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v120 = F_bms_difference(m, v119, v117)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L11
	} else {
		goto L42
	}
L42:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v123 = int32(0)
	v130 = base.B2i32(v120|v122 == v123)
	if v120 == v123 {
		v169 = v130
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v169 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L44:
	;
	goto L43
L45:
	;
	if v122 == int32(0) {
		v169 = v130
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v136 != v137 {
		v169 = int32(0)
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v139 = int32(1)
	if v136 <= v139 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v142 = v139
	goto L50
L49:
	;
	v142 = v136
	goto L50
L50:
	;
	v143 = int32(8)
	v148 = int32(0)
	goto L51
L51:
	;
	v156 = v148 << (uint(int32(2)) % 32)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v120+v143+v156)))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v156+(v122+v143))))
	v161 = base.B2i32(v158 == v160)
	if v160 != v158 {
		v169 = v161
		goto L44
	} else {
		goto L53
	}
L52:
	;
	v169 = v161
	goto L44
L53:
	;
	v164 = v148 + int32(1)
	if v164 != v142 {
		v148 = v164
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+8))
	v178 = F_bms_difference(m, v175, v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L11
	} else {
		goto L58
	}
L56:
	;
	v182 = v120
	goto L57
L57:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v184 = F_bms_join(m, v183, v182)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L11
	} else {
		goto L60
	}
L58:
	;
	v180 = F_bms_join(m, v120, v178)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L11
	} else {
		goto L59
	}
L59:
	;
	v182 = v180
	goto L57
L60:
	;
	v188 = v184
	goto L18
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v191
	return int32(0)
L62:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v204 - int32(1)
	return v202
L63:
	;
	v214 = v212
	goto L1
}
func F_pullf_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
	if v4 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		m.T0[v4].(func(*base.Module, int32))(m, v5)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v8 != 0 {
				v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v11 = F___memset(m, v8, int32(0), v10)
				mBase = m.M
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				F_pfree(m, v12)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					v17 = F___memset(m, l0, int32(0), int32(24))
					mBase = m.M
					F_pfree(m, l0)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				v17 = F___memset(m, l0, int32(0), int32(24))
				mBase = m.M
				F_pfree(m, l0)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v8 != 0 {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v11 = F___memset(m, v8, int32(0), v10)
			mBase = m.M
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			F_pfree(m, v12)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				v17 = F___memset(m, l0, int32(0), int32(24))
				mBase = m.M
				F_pfree(m, l0)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			v17 = F___memset(m, l0, int32(0), int32(24))
			mBase = m.M
			F_pfree(m, l0)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_pullf_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v7 == int32(0) {
		v10 = l0
		for {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
			if v17 == int32(0) {
				v10 = v15
				continue
			} else {
				break
			}
			break
		}
		v20 = v15
		v24 = v17
	} else {
		v20 = l0
		v24 = v7
	}
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if l1 < v27 {
		v29 = l1
	} else {
		v29 = v27
	}
	if v27 != 0 {
		v30 = v29
	} else {
		v30 = l1
	}
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v32 = m.T0[v24].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v25, v26, v30, l2, v31, v27)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		return int32(0)
	} else {
		return v32
	}
}
func F_pullup_replace_vars_callback(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v523 int32
	_ = v523
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v566 int32
	_ = v566
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	v9 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
	if v9 < int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v18 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	return int32(0)
L5:
	;
	return v12
L6:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v573 == int32(0) {
		v751 = v566
		goto L155
	} else {
		goto L156
	}
L7:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v48 = int32(0)
	v50 = F_ReplaceVarFromTargetList(m, l0, v46, v24, v47, v48, v48)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L20
	}
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v41 = int32(0)
	v43 = F_ReplaceVarFromTargetList(m, l0, v38, v39, v40, v41, v41)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L19
	}
L9:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	if v21 == int32(0) {
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v24 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v27 = v25
	goto L15
L14:
	;
	v27 = int32(0)
	goto L15
L15:
	;
	if v27 < v9 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+v9<<(uint(int32(2))%32))))
	if v33 == int32(0) {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v36 = F_copyObjectImpl(m, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v566 = v36
	goto L6
L19:
	;
	v566 = v43
	goto L6
L20:
	;
	if v9 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v549 = F_bms_make_singleton(m, v548)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L4
	} else {
		goto L148
	}
L22:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	if v54 == int32(1) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	if v50 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+124)))
	if v337 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L25:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v59 != int32(319) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	if v59 != int32(6) {
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
	if v141 != 0 {
		goto L24
	} else {
		goto L49
	}
L29:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v50)+28))
	if v64 != 0 {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+124)))
	if v66 != int32(1) {
		v566 = v50
		goto L6
	} else {
		goto L31
	}
L31:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v71 = F_bms_is_member(m, v69, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	if v71 != 0 {
		v566 = v50
		goto L6
	} else {
		goto L33
	}
L33:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v76 = int32(2)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v74+v75<<(uint(v76)%32))))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v74+v80<<(uint(v76)%32))))
	v85 = int32(0)
	if v79 == v85 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v138 == int32(0) {
		goto L21
	} else {
		goto L48
	}
L35:
	;
	v138 = int32(1)
	goto L34
L36:
	;
	goto L37
L37:
	;
	if v84 == int32(0) {
		v129 = v85
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v138 = v129
	goto L34
L39:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v95 < v94 {
		v129 = v85
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v97 = int32(1)
	if v94 <= v97 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v100 = v97
	goto L43
L42:
	;
	v100 = v94
	goto L43
L43:
	;
	v101 = int32(8)
	v106 = int32(0)
	goto L44
L44:
	;
	v113 = v106 << (uint(int32(2)) % 32)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v79+v101+v113)))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113+(v84+v101))))
	v120 = v115 & (v117 ^ int32(-1))
	v122 = base.B2i32(v120 == int32(0))
	if v120 != 0 {
		v129 = v122
		goto L38
	} else {
		goto L46
	}
L45:
	;
	v129 = v122
	goto L38
L46:
	;
	v124 = v106 + int32(1)
	if v124 != v100 {
		v106 = v124
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v566 = v50
	goto L6
L49:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+124)))
	if v143 != int32(1) {
		v566 = v50
		goto L6
	} else {
		goto L50
	}
L50:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v148 = int32(0)
	if v146 == v148 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v201 != 0 {
		v566 = v50
		goto L6
	} else {
		goto L65
	}
L52:
	;
	v201 = int32(1)
	goto L51
L53:
	;
	goto L54
L54:
	;
	if v147 == int32(0) {
		v192 = v148
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v201 = v192
	goto L51
L56:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v147)+4))
	if v158 < v157 {
		v192 = v148
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v160 = int32(1)
	if v157 <= v160 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v163 = v160
	goto L60
L59:
	;
	v163 = v157
	goto L60
L60:
	;
	v164 = int32(8)
	v169 = int32(0)
	goto L61
L61:
	;
	v176 = v169 << (uint(int32(2)) % 32)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v146+v164+v176)))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v176+(v147+v164))))
	v183 = v178 & (v180 ^ int32(-1))
	v185 = base.B2i32(v183 == int32(0))
	if v183 != 0 {
		v192 = v185
		goto L55
	} else {
		goto L63
	}
L62:
	;
	v192 = v185
	goto L55
L63:
	;
	v187 = v169 + int32(1)
	if v187 != v163 {
		v169 = v187
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v208 = int32(-1)
	goto L66
L66:
	;
	if v202 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	goto L21
L68:
	;
	if v268 < int32(0) {
		v566 = v50
		goto L6
	} else {
		goto L79
	}
L69:
	;
	v268 = base.I32_ctz(v254) | v255<<(uint(int32(5))%32)
	goto L68
L70:
	;
	v268 = int32(-2)
	goto L68
L71:
	;
	v219 = v208 + int32(1)
	v221 = base.I32_div_s(v219, int32(32))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v202)+4))
	if v222 <= v221 {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v225 = v202 + int32(8)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v225+v221<<(uint(int32(2))%32))))
	v232 = v229 & (int32(-1) << (uint(v219) % 32))
	if v232 != 0 {
		v254 = v232
		v255 = v221
		goto L69
	} else {
		goto L73
	}
L73:
	;
	v234 = v221 + int32(1)
	if v234 == v222 {
		goto L70
	} else {
		goto L74
	}
L74:
	;
	v237 = v234
	goto L75
L75:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v225+v237<<(uint(int32(2))%32))))
	if v244 != 0 {
		v254 = v244
		v255 = v237
		goto L69
	} else {
		goto L77
	}
L76:
	;
	goto L70
L77:
	;
	v246 = v237 + int32(1)
	if v246 != v222 {
		v237 = v246
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v273 = int32(2)
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v271+v272<<(uint(v273)%32))))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v271+v268<<(uint(v273)%32))))
	v281 = int32(0)
	if v276 == v281 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v334 != 0 {
		v208 = v268
		goto L66
	} else {
		goto L94
	}
L81:
	;
	v334 = int32(1)
	goto L80
L82:
	;
	goto L83
L83:
	;
	if v280 == int32(0) {
		v325 = v281
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v334 = v325
	goto L80
L85:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v280)+4))
	if v291 < v290 {
		v325 = v281
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v293 = int32(1)
	if v290 <= v293 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v296 = v293
	goto L89
L88:
	;
	v296 = v290
	goto L89
L89:
	;
	v297 = int32(8)
	v302 = int32(0)
	goto L90
L90:
	;
	v309 = v302 << (uint(int32(2)) % 32)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v276+v297+v309)))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v309+(v280+v297))))
	v316 = v311 & (v313 ^ int32(-1))
	v318 = base.B2i32(v316 == int32(0))
	if v316 != 0 {
		v325 = v318
		goto L84
	} else {
		goto L92
	}
L91:
	;
	v325 = v318
	goto L84
L92:
	;
	v320 = v302 + int32(1)
	if v320 != v296 {
		v302 = v320
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	goto L67
L95:
	;
	v535 = F_contain_nonstrict_functions_walker(m, v50, int32(0))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L4
	} else {
		goto L146
	}
L96:
	;
	v341 = F_contain_vars_of_level(m, v50, int32(0))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L4
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v344 = F_pull_varnos(m, v343, v50)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L4
	} else {
		goto L101
	}
L99:
	;
	if v341 != 0 {
		goto L95
	} else {
		goto L100
	}
L100:
	;
	goto L21
L101:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v347 = int32(0)
	if v344 == v347 {
		v388 = v347
		goto L103
	} else {
		goto L104
	}
L102:
	;
	if v388 != 0 {
		goto L95
	} else {
		goto L116
	}
L103:
	;
	goto L102
L104:
	;
	if v346 == int32(0) {
		v388 = v347
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v344)+4))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v346)+4))
	if v356 < v357 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v359 = v356
	goto L108
L107:
	;
	v359 = v357
	goto L108
L108:
	;
	if v359 <= int32(1) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v362 = int32(1)
	goto L111
L110:
	;
	v362 = v359
	goto L111
L111:
	;
	v363 = int32(8)
	v368 = int32(0)
	goto L112
L112:
	;
	v375 = v368 << (uint(int32(2)) % 32)
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v346+v363+v375)))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v375+(v344+v363))))
	v380 = v377 & v379
	v382 = base.B2i32(v380 != int32(0))
	if v380 != 0 {
		v388 = v382
		goto L103
	} else {
		goto L114
	}
L113:
	;
	v388 = v382
	goto L103
L114:
	;
	v384 = v368 + int32(1)
	if v384 != v362 {
		v368 = v384
		goto L112
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v397 = int32(-1)
	goto L117
L117:
	;
	if v344 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L118:
	;
	goto L95
L119:
	;
	if v457 < int32(0) {
		goto L21
	} else {
		goto L130
	}
L120:
	;
	v457 = base.I32_ctz(v443) | v444<<(uint(int32(5))%32)
	goto L119
L121:
	;
	v457 = int32(-2)
	goto L119
L122:
	;
	v408 = v397 + int32(1)
	v410 = base.I32_div_s(v408, int32(32))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v344)+4))
	if v411 <= v410 {
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v414 = v344 + int32(8)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v414+v410<<(uint(int32(2))%32))))
	v421 = v418 & (int32(-1) << (uint(v408) % 32))
	if v421 != 0 {
		v443 = v421
		v444 = v410
		goto L120
	} else {
		goto L124
	}
L124:
	;
	v423 = v410 + int32(1)
	if v423 == v411 {
		goto L121
	} else {
		goto L125
	}
L125:
	;
	v426 = v423
	goto L126
L126:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v414+v426<<(uint(int32(2))%32))))
	if v433 != 0 {
		v443 = v433
		v444 = v426
		goto L120
	} else {
		goto L128
	}
L127:
	;
	goto L121
L128:
	;
	v435 = v426 + int32(1)
	if v435 != v411 {
		v426 = v435
		goto L126
	} else {
		goto L129
	}
L129:
	;
	goto L127
L130:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v392)))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v462 = int32(2)
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v460+v461<<(uint(v462)%32))))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v460+v457<<(uint(v462)%32))))
	v470 = int32(0)
	if v465 == v470 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	if v523 == int32(0) {
		v397 = v457
		goto L117
	} else {
		goto L145
	}
L132:
	;
	v523 = int32(1)
	goto L131
L133:
	;
	goto L134
L134:
	;
	if v469 == int32(0) {
		v514 = v470
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v523 = v514
	goto L131
L136:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v465)+4))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v469)+4))
	if v480 < v479 {
		v514 = v470
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v482 = int32(1)
	if v479 <= v482 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v485 = v482
	goto L140
L139:
	;
	v485 = v479
	goto L140
L140:
	;
	v486 = int32(8)
	v491 = int32(0)
	goto L141
L141:
	;
	v498 = v491 << (uint(int32(2)) % 32)
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v465+v486+v498)))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v498+(v469+v486))))
	v505 = v500 & (v502 ^ int32(-1))
	v507 = base.B2i32(v505 == int32(0))
	if v505 != 0 {
		v514 = v507
		goto L135
	} else {
		goto L143
	}
L142:
	;
	v514 = v507
	goto L135
L143:
	;
	v509 = v491 + int32(1)
	if v509 != v485 {
		v491 = v509
		goto L141
	} else {
		goto L144
	}
L144:
	;
	goto L142
L145:
	;
	goto L118
L146:
	;
	if v535 == int32(0) {
		v566 = v50
		goto L6
	} else {
		goto L147
	}
L147:
	;
	goto L21
L148:
	;
	v551 = F_make_placeholder_expr(m, v547, v50, v549)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v553 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v553)+4))
	v556 = v554
	goto L152
L151:
	;
	v556 = int32(0)
	goto L152
L152:
	;
	if v556 < v9 {
		v566 = v551
		goto L6
	} else {
		goto L153
	}
L153:
	;
	v558 = F_copyObjectImpl(m, v551)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v560+v9<<(uint(int32(2))%32)))) = v558
	v566 = v551
	goto L6
L155:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v758 != 0 {
		goto L201
	} else {
		goto L202
	}
L156:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v566)))
	if v576 != int32(319) {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589)+124)))
	if v590 != int32(1) {
		v739 = v566
		goto L164
	} else {
		goto L165
	}
L158:
	;
	if v576 != int32(6) {
		goto L157
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v566)+12))
	v586 = F_bms_add_members(m, v585, v573)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L4
	} else {
		goto L163
	}
L161:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v566)+24))
	v582 = F_bms_add_members(m, v581, v573)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L4
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v566)+24)) = v582
	v751 = v566
	goto L155
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v566)+12)) = v586
	v751 = v566
	goto L155
L164:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v747 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v748 = F_add_nulling_relids(m, v739, v746, v747)
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L4
	} else {
		goto L200
	}
L165:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v595 = F_pull_varnos(m, v594, v566)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L4
	} else {
		goto L166
	}
L166:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v598 = F_bms_del_members(m, v595, v597)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L4
	} else {
		goto L167
	}
L167:
	;
	if v598 == int32(0) {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	if v656 < int32(0) {
		v739 = v566
		goto L164
	} else {
		goto L179
	}
L169:
	;
	v656 = base.I32_ctz(v642) | v643<<(uint(int32(5))%32)
	goto L168
L170:
	;
	v656 = int32(-2)
	goto L168
L171:
	;
	v609 = base.I32_div_s(int32(0), int32(32))
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v598)+4))
	if v610 <= v609 {
		goto L170
	} else {
		goto L172
	}
L172:
	;
	v613 = v598 + int32(8)
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v613+v609<<(uint(int32(2))%32))))
	v620 = v617 & int32(-1)
	if v620 != 0 {
		v642 = v620
		v643 = v609
		goto L169
	} else {
		goto L173
	}
L173:
	;
	v622 = v609 + int32(1)
	if v622 == v610 {
		goto L170
	} else {
		goto L174
	}
L174:
	;
	v625 = v622
	goto L175
L175:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v613+v625<<(uint(int32(2))%32))))
	if v632 != 0 {
		v642 = v632
		v643 = v625
		goto L169
	} else {
		goto L177
	}
L176:
	;
	goto L170
L177:
	;
	v634 = v625 + int32(1)
	if v634 != v610 {
		v625 = v634
		goto L175
	} else {
		goto L178
	}
L178:
	;
	goto L176
L179:
	;
	v660 = v566
	v663 = v656
	goto L180
L180:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v593)))
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v668+v663<<(uint(int32(2))%32))))
	v673 = F_bms_intersect(m, v667, v672)
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L4
	} else {
		goto L182
	}
L181:
	;
	v739 = v679
	goto L164
L182:
	;
	if v673 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v675 = F_bms_make_singleton(m, v663)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L4
	} else {
		goto L186
	}
L184:
	;
	v679 = v660
	goto L185
L185:
	;
	if v598 == int32(0) {
		goto L190
	} else {
		goto L191
	}
L186:
	;
	v677 = F_add_nulling_relids(m, v660, v675, v673)
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L4
	} else {
		goto L187
	}
L187:
	;
	v679 = v677
	goto L185
L188:
	;
	if int32(0) <= v735 {
		v660 = v679
		v663 = v735
		goto L180
	} else {
		goto L199
	}
L189:
	;
	v735 = base.I32_ctz(v721) | v722<<(uint(int32(5))%32)
	goto L188
L190:
	;
	v735 = int32(-2)
	goto L188
L191:
	;
	v686 = v663 + int32(1)
	v688 = base.I32_div_s(v686, int32(32))
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v598)+4))
	if v689 <= v688 {
		goto L190
	} else {
		goto L192
	}
L192:
	;
	v692 = v598 + int32(8)
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v692+v688<<(uint(int32(2))%32))))
	v699 = v696 & (int32(-1) << (uint(v686) % 32))
	if v699 != 0 {
		v721 = v699
		v722 = v688
		goto L189
	} else {
		goto L193
	}
L193:
	;
	v701 = v688 + int32(1)
	if v701 == v689 {
		goto L190
	} else {
		goto L194
	}
L194:
	;
	v704 = v701
	goto L195
L195:
	;
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v692+v704<<(uint(int32(2))%32))))
	if v711 != 0 {
		v721 = v711
		v722 = v704
		goto L189
	} else {
		goto L197
	}
L196:
	;
	goto L190
L197:
	;
	v713 = v704 + int32(1)
	if v713 != v689 {
		v704 = v713
		goto L195
	} else {
		goto L198
	}
L198:
	;
	goto L196
L199:
	;
	goto L181
L200:
	;
	v751 = v748
	goto L155
L201:
	;
	F_IncrementVarSublevelsUp(m, v751, v758, int32(0))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L4
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	return v751
L204:
	;
	goto L203
}
func F_push_old_value(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	v6 = *(*int32)(unsafe.Add(mBase, _consts[159]))
	if v6 == int32(0) {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		if v9 == int32(0) {
			v37 = *(*int32)(unsafe.Add(mBase, _consts[76]))
			v39 = F_MemoryContextAllocZero(m, v37, int32(64))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				*(*int32)(unsafe.Add(mBase, uint32(v39))) = v41
				v44 = *(*int32)(unsafe.Add(mBase, _consts[159]))
				*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v44
				if base.Ui32(l1) <= base.Ui32(int32(2)) {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(2))%32))+uint32(_consts[956])))
					*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v52
				} else {
				}
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v54
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v56
				v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v39)+24)) = v58
				F_set_stack_value(m, l0, v39+int32(32))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					if v64 == int32(0) {
						v67 = int32(4474188)
						v68 = *(*int32)(unsafe.Add(mBase, _consts[957]))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v68
						*(*int32)(unsafe.Add(mBase, _consts[957])) = l0 + int32(72)
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v39
					return
				}
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			if v12 < v6 {
				v37 = *(*int32)(unsafe.Add(mBase, _consts[76]))
				v39 = F_MemoryContextAllocZero(m, v37, int32(64))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					*(*int32)(unsafe.Add(mBase, uint32(v39))) = v41
					v44 = *(*int32)(unsafe.Add(mBase, _consts[159]))
					*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v44
					if base.Ui32(l1) <= base.Ui32(int32(2)) {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(2))%32))+uint32(_consts[956])))
						*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v52
					} else {
					}
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v54
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v56
					v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v39)+24)) = v58
					F_set_stack_value(m, l0, v39+int32(32))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
						if v64 == int32(0) {
							v67 = int32(4474188)
							v68 = *(*int32)(unsafe.Add(mBase, _consts[957]))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v68
							*(*int32)(unsafe.Add(mBase, _consts[957])) = l0 + int32(72)
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v39
						return
					}
				}
			} else {
				switch l1 {
				case 0:
					v14 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
					if v14 == int32(3) {
						F_discard_stack_value(m, l0, v9+int32(48))
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(1)
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(1)
						return
					}
				case 1:
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
					if v23 != int32(1) {
						return
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v26
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v28
						F_set_stack_value(m, l0, v9+int32(48))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(3)
							return
						}
					}
				default:
					return
				}
			}
		}
	}
}
