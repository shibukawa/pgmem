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
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	v14 = m.G0
	v16 = v14 - int32(192)
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
	v27 = v16 + int32(128)
	goto L10
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l2
	v157 = F_pg_snprintf(m, v16+int32(128), int32(64), int32(_a_F_ChooseConstraintName_0), v16+int32(16))
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
	v177 = F_makeObjectName(m, l0, l1, v16+int32(128))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	F_relation_close(m, v21, int32(1))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
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
	F_pfree(m, v177)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L67
	}
L44:
	;
	v252 = v16 + int32(32)
	F_ScanKeyInit(m, v252, int32(2), int32(3), int32(62), v177)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L61
	}
L45:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v181 <= int32(0) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v184 = int32(0)
	if v184 < v181 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v187 = v181
	goto L49
L48:
	;
	v187 = v184
	goto L49
L49:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v196 = int32(0)
	goto L50
L50:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v188+v196<<(uint(int32(2))%32))))
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	if base.B2i32(v209 == int32(0))|base.B2i32(v209 != v212) != 0 {
		v230 = v209
		v231 = v212
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L44
L52:
	;
	if v230-v231 == int32(0) {
		goto L43
	} else {
		goto L59
	}
L53:
	;
	goto L52
L54:
	;
	v215 = v206
	v216 = v177
	goto L55
L55:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+1)))
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+1)))
	if v220 == int32(0) {
		v230 = v220
		v231 = v219
		goto L53
	} else {
		goto L57
	}
L56:
	;
	v230 = v220
	v231 = v219
	goto L53
L57:
	;
	v223 = int32(1)
	if v220 == v219 {
		v215 = v215 + v223
		v216 = v216 + v223
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v236 = v196 + int32(1)
	if v236 != v187 {
		v196 = v236
		goto L50
	} else {
		goto L60
	}
L60:
	;
	goto L51
L61:
	;
	v258 = int32(3)
	F_ScanKeyInit(m, v16+int32(80), v258, v258, int32(184), l3)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v267 = F_systable_beginscan(m, v21, int32(2664), int32(1), int32(0), int32(2), v252)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v269 = F_systable_getnext(m, v267)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_systable_endscan(m, v267)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	if v269 == int32(0) {
		goto L41
	} else {
		goto L66
	}
L66:
	;
	goto L43
L67:
	;
	v291 = v170 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v291
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l2
	v298 = F_pg_snprintf(m, v16+int32(128), int32(64), int32(_a_F_ChooseConstraintName_0), v16)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v170 = v291
	goto L39
L69:
	;
	m.G0 = v16 + int32(192)
	return v177
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
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	v6 = m.G0
	v8 = v6 - int32(144)
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
		F_ScanKeyInit(m, v8, int32(9), int32(3), int32(184), v20)
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			if l0 == int32(1) {
				v31 = l1
			} else {
				v31 = int32(0)
			}
			F_ScanKeyInit(m, v8+int32(48), int32(10), int32(3), int32(184), v31)
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				F_ScanKeyInit(m, v8+int32(96), int32(2), int32(3), int32(62), l2)
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					v45 = F_systable_beginscan(m, v12, int32(2665), int32(1), int32(0), int32(3), v8)
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						v47 = F_systable_getnext(m, v45)
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							F_systable_endscan(m, v45)
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								F_relation_close(m, v12, int32(1))
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									m.G0 = v8 + int32(144)
									return base.B2i32(v47 != int32(0))
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
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_SearchSysCache1(m, int32(19), l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(_a_F_get_constraint_type_0), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_get_constraint_type_1), int32(1242), int32(_a_F_get_constraint_type_2))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
			v31 = int32(*(*int8)(unsafe.Add(mBase, uint32(v28+v29)+72)))
			F_ReleaseCatCache(m, v9)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v31
			}
		}
	}
}
