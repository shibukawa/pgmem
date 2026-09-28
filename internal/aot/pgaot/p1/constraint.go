package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ChooseConstraintName(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
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
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v170 int32
	_ = v170
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	v14 = m.G0
	v16 = v14 - int32(208)
	m.G0 = v16
	v18 = int32(1)
	v21 = F_table_open(m, int32(2606), v18)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v170 = v159
	goto L39
L4:
	;
	v27 = v16 + int32(144)
	goto L10
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l2
	v157 = F_pg_snprintf(m, v16+int32(144), int32(64), int32(_a_F_ChooseConstraintName_0), v16+int32(16))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L38
	}
L7:
	;
	v159 = int32(0)
	goto L3
L8:
	;
	v144 = F_strlen(m, v133)
	mBase = m.M
	goto L7
L10:
	;
	goto L11
L11:
	;
	v34 = int32(63)
	if (v27^l2)&int32(3) != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v137 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v137)
	goto L8
L13:
	;
	v118 = v113
	v119 = v114
	v120 = v115
	goto L34
L14:
	;
	if v108 == int32(0) {
		v133 = v106
		v134 = v107
		goto L12
	} else {
		goto L33
	}
L15:
	;
	v106 = l2
	v107 = v27
	v108 = v34
	goto L14
L16:
	;
	goto L17
L17:
	;
	v38 = int32(0)
	if base.B2i32(l2&int32(3) == v38)|int32(0) == v38 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v74 == int32(0) {
		v133 = v71
		v134 = v72
		goto L12
	} else {
		goto L27
	}
L19:
	;
	v50 = l2
	v51 = v27
	v52 = v34
	goto L22
L20:
	;
	goto L21
L21:
	;
	v71 = l2
	v72 = v27
	v73 = v34
	v74 = int32(1)
	goto L18
L22:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	*(*uint8)(unsafe.Add(mBase, uint32(v51))) = uint8(v54)
	if v54 == int32(0) {
		v113 = v50
		v114 = v51
		v115 = v52
		goto L13
	} else {
		goto L24
	}
L23:
	;
	v71 = v65
	v72 = v59
	v73 = v61
	v74 = v63
	goto L18
L24:
	;
	v58 = int32(1)
	v59 = v51 + v58
	v61 = v52 - v58
	v62 = int32(0)
	v63 = base.B2i32(v61 != v62)
	v65 = v50 + v58
	if v65&int32(3) == v62 {
		v71 = v65
		v72 = v59
		v73 = v61
		v74 = v63
		goto L18
	} else {
		goto L25
	}
L25:
	;
	if v61 != 0 {
		v50 = v65
		v51 = v59
		v52 = v61
		goto L22
	} else {
		goto L26
	}
L26:
	;
	goto L23
L27:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if base.B2i32(v77 == int32(0))|base.B2i32(base.Ui32(v73) < base.Ui32(int32(4))) != 0 {
		v106 = v71
		v107 = v72
		v108 = v73
		goto L14
	} else {
		goto L28
	}
L28:
	;
	v84 = v71
	v85 = v72
	v86 = v73
	goto L29
L29:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v92 = int32(-2139062144)
	if (int32(16843008)-v89|v89)&v92 != v92 {
		v113 = v84
		v114 = v85
		v115 = v86
		goto L13
	} else {
		goto L31
	}
L30:
	;
	v106 = v100
	v107 = v98
	v108 = v102
	goto L14
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v89
	v97 = int32(4)
	v98 = v85 + v97
	v100 = v84 + v97
	v102 = v86 - v97
	if base.Ui32(int32(3)) < base.Ui32(v102) {
		v84 = v100
		v85 = v98
		v86 = v102
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v113 = v106
	v114 = v107
	v115 = v108
	goto L13
L34:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	*(*uint8)(unsafe.Add(mBase, uint32(v119))) = uint8(v122)
	if v122 == int32(0) {
		v133 = v118
		v134 = v119
		goto L12
	} else {
		goto L36
	}
L35:
	;
	v133 = v129
	v134 = v127
	goto L12
L36:
	;
	v126 = int32(1)
	v127 = v119 + v126
	v129 = v118 + v126
	v131 = v120 - v126
	if v131 != 0 {
		v118 = v129
		v119 = v127
		v120 = v131
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v159 = v18
	goto L3
L39:
	;
	v178 = F_makeObjectName(m, l0, l1, v16+int32(144))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	F_relation_close(m, v21, int32(1))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L69
	}
L41:
	;
	goto L40
L42:
	;
	if l4 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	F_pfree(m, v178)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L67
	}
L44:
	;
	v253 = v16 + int32(32)
	F_ScanKeyInit(m, v253, int32(2), int32(3), int32(62), base.I64_extend_i32_u(v178))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L61
	}
L45:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v182 <= int32(0) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v185 = int32(0)
	if v185 < v182 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v188 = v182
	goto L49
L48:
	;
	v188 = v185
	goto L49
L49:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v194 = int32(0)
	goto L50
L50:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v189+v194<<(uint(int32(2))%32))))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	if base.B2i32(v210 == int32(0))|base.B2i32(v210 != v213) != 0 {
		v231 = v210
		v232 = v213
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L44
L52:
	;
	if v231-v232 == int32(0) {
		goto L43
	} else {
		goto L59
	}
L53:
	;
	goto L52
L54:
	;
	v216 = v207
	v217 = v178
	goto L55
L55:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+1)))
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+1)))
	if v221 == int32(0) {
		v231 = v221
		v232 = v220
		goto L53
	} else {
		goto L57
	}
L56:
	;
	v231 = v221
	v232 = v220
	goto L53
L57:
	;
	v224 = int32(1)
	if v221 == v220 {
		v216 = v216 + v224
		v217 = v217 + v224
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v237 = v194 + int32(1)
	if v237 != v188 {
		v194 = v237
		goto L50
	} else {
		goto L60
	}
L60:
	;
	goto L51
L61:
	;
	v260 = int32(3)
	F_ScanKeyInit(m, v16+int32(88), v260, v260, int32(184), base.I64_extend_i32_u(l3))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v269 = F_systable_beginscan(m, v21, int32(2664), int32(1), int32(0), int32(2), v253)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v271 = F_systable_getnext(m, v269)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_systable_endscan(m, v269)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	if v271 == int32(0) {
		goto L41
	} else {
		goto L66
	}
L66:
	;
	goto L43
L67:
	;
	v293 = v170 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v293
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l2
	v300 = F_pg_snprintf(m, v16+int32(144), int32(64), int32(_a_F_ChooseConstraintName_0), v16)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v170 = v293
	goto L39
L69:
	;
	m.G0 = v16 + int32(208)
	return v178
}
func F_ConstraintNameIsUsed(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	v6 = m.G0
	v8 = v6 - int32(176)
	m.G0 = v8
	v12 = F_table_open(m, int32(2606), int32(1))
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if l0 != 0 {
			v20 = int32(0)
		} else {
			v20 = l1
		}
		F_ScanKeyInit(m, v8, int32(9), int32(3), int32(184), base.I64_extend_i32_u(v20))
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			if l0 == int32(1) {
				v32 = l1
			} else {
				v32 = int32(0)
			}
			F_ScanKeyInit(m, v8+int32(56), int32(10), int32(3), int32(184), base.I64_extend_i32_u(v32))
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				F_ScanKeyInit(m, v8+int32(112), int32(2), int32(3), int32(62), base.I64_extend_i32_u(l2))
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					v48 = F_systable_beginscan(m, v12, int32(2665), int32(1), int32(0), int32(3), v8)
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						v50 = F_systable_getnext(m, v48)
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							F_systable_endscan(m, v48)
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								F_relation_close(m, v12, int32(1))
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									m.G0 = v8 + int32(176)
									return base.B2i32(v50 != int32(0))
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_get_constraint_type(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = F_SearchSysCache1(m, int32(19), base.I64_extend_i32_u(l0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(_a_F_get_constraint_type_0), v6)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_get_constraint_type_1), int32(1374), int32(_a_F_get_constraint_type_2))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
			v32 = int32(*(*int8)(unsafe.Add(mBase, uint32(v29+v30)+72)))
			F_ReleaseCatCache(m, v10)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v32
			}
		}
	}
}
