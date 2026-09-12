package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_convert_any_priv_string(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int64
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v189 int32
	_ = v189
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int64
	_ = v216
	var v218 int32
	_ = v218
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = F_text_to_cstring(m, l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v21 = v14
	v26 = int64(0)
	goto L4
L3:
	;
	F_pfree(m, v14)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L68
	}
L4:
	;
	if v21 == int32(0) {
		goto L3
	} else {
		goto L6
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L64
	}
L6:
	;
	v29 = int32(44)
	v30 = F___strchrnul(m, v21, v29)
	mBase = m.M
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v32 == v29 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v36 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v36 = v30
	goto L10
L9:
	;
	v36 = int32(0)
	goto L10
L10:
	;
	goto L7
L11:
	;
	v37 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v36))) = uint8(v37)
	v42 = v36 + int32(1)
	goto L13
L12:
	;
	v42 = int32(0)
	goto L13
L13:
	;
	v43 = v21
	goto L14
L14:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v52-int32(9)))&base.B2i32(v52 != int32(32)) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v43&int32(3) == int32(0) {
		v87 = v43
		goto L22
	} else {
		goto L23
	}
L16:
	;
	v43 = v43 + int32(1)
	goto L14
L17:
	;
	goto L18
L18:
	;
	goto L15
L19:
	;
	v158 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v43+v150))) = uint8(v158)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v160 == v158 {
		goto L42
	} else {
		goto L43
	}
L20:
	;
	if v120 <= int32(0) {
		v150 = v120
		goto L19
	} else {
		goto L37
	}
L21:
	;
	v120 = v112 - v43
	goto L20
L22:
	;
	v91 = v87
	goto L31
L23:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v71 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v120 = int32(0)
	goto L20
L25:
	;
	goto L26
L26:
	;
	v76 = v43
	goto L27
L27:
	;
	v80 = v76 + int32(1)
	if v80&int32(3) == int32(0) {
		v87 = v80
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v112 = v80
	goto L21
L29:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	if v85 != 0 {
		v76 = v80
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v100 = int32(-2139062144)
	if (int32(16843008)-v97|v97)&v100 == v100 {
		v91 = v91 + int32(4)
		goto L31
	} else {
		goto L33
	}
L32:
	;
	v106 = v91
	goto L34
L33:
	;
	goto L32
L34:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	if v110 != 0 {
		v106 = v106 + int32(1)
		goto L34
	} else {
		goto L36
	}
L35:
	;
	v112 = v106
	goto L21
L36:
	;
	goto L35
L37:
	;
	v127 = v120
	goto L38
L38:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127+(v43-int32(1))))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v135-int32(9)))&base.B2i32(v135 != int32(32)) != 0 {
		v150 = v127
		goto L19
	} else {
		goto L40
	}
L39:
	;
	v150 = int32(0)
	goto L19
L40:
	;
	v143 = int32(1)
	if v143 < v127 {
		v127 = v127 - v143
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	goto L5
L43:
	;
	v165 = l1
	v166 = v160
	goto L44
L44:
	;
	v174 = v166
	v175 = v43
	goto L47
L45:
	;
	v216 = *(*int64)(unsafe.Add(mBase, uint32(v165)+8))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	if v218 != 0 {
		v21 = v42
		v26 = v216 | v26
		goto L4
	} else {
		goto L63
	}
L46:
	;
	if v212 != 0 {
		goto L59
	} else {
		goto L60
	}
L47:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	if v178 == v179 {
		v201 = v178
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v212 = int32(0)
	goto L46
L49:
	;
	v203 = int32(1)
	if v201 != 0 {
		v174 = v174 + v203
		v175 = v175 + v203
		goto L47
	} else {
		goto L58
	}
L50:
	;
	if base.Ui32((v178-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v189 = v178 | int32(32)
	goto L53
L52:
	;
	v189 = v178
	goto L53
L53:
	;
	if base.Ui32((v179-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v198 = v179 | int32(32)
	goto L56
L55:
	;
	v198 = v179
	goto L56
L56:
	;
	if v189 == v198 {
		v201 = v189
		goto L49
	} else {
		goto L57
	}
L57:
	;
	v212 = v189 - v198
	goto L46
L58:
	;
	goto L48
L59:
	;
	v214 = v165 + int32(16)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
	if v215 != 0 {
		v165 = v214
		v166 = v215
		goto L44
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	goto L45
L62:
	;
	goto L42
L63:
	;
	goto L42
L64:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v43
	F_errmsg(m, int32(697700), v12)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(486333), int32(1726), int32(322390))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	m.G0 = v12 + int32(16)
	return v26
}
func F_has_any_column_privilege_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, _consts[276]))
			v15 = F_textToQualifiedNameList(m, v6)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = F_makeRangeVarFromNameList(m, v15)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v19 = int32(0)
					v23 = F_RangeVarGetRelidExtended(m, v17, v19, v19, v19, v19)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v26 = F_convert_any_priv_string(m, v11, int32(1628096))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							v28 = F_pg_class_aclcheck(m, v23, v14, v26)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int32(0)
							} else {
								if v28 == int32(0) {
									return int32(1)
								} else {
									v35 = F_pg_attribute_aclcheck_all(m, v23, v14, v26, int32(1))
									mBase = m.M
									v36 = m.ExcPending
									if v36 != 0 {
										return int32(0)
									} else {
										return base.B2i32(v35 == int32(0))
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
func F_has_any_column_privilege_name_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v13 = F_pg_detoast_datum_packed(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = F_get_role_oid_or_public(m, v6)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = F_textToQualifiedNameList(m, v8)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v19 = F_makeRangeVarFromNameList(m, v17)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						v21 = int32(0)
						v25 = F_RangeVarGetRelidExtended(m, v19, v21, v21, v21, v21)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							v28 = F_convert_any_priv_string(m, v13, int32(1628096))
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int32(0)
							} else {
								v30 = F_pg_class_aclcheck(m, v25, v15, v28)
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return int32(0)
								} else {
									if v30 == int32(0) {
										return int32(1)
									} else {
										v37 = F_pg_attribute_aclcheck_all(m, v25, v15, v28, int32(1))
										mBase = m.M
										v38 = m.ExcPending
										if v38 != 0 {
											return int32(0)
										} else {
											return base.B2i32(v37 == int32(0))
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
