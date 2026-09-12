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
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int64
	_ = v160
	var v162 int32
	_ = v162
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
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
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L51
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
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L47
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
	v64 = F_strlen(m, v43)
	mBase = m.M
	if v64 <= int32(0) {
		v94 = v64
		goto L19
	} else {
		goto L20
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
	v102 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v43+v94))) = uint8(v102)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v104 == v102 {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	v71 = v64
	goto L21
L21:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+(v43-int32(1))))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v79-int32(9)))&base.B2i32(v79 != int32(32)) != 0 {
		v94 = v71
		goto L19
	} else {
		goto L23
	}
L22:
	;
	v94 = int32(0)
	goto L19
L23:
	;
	v87 = int32(1)
	if v87 < v71 {
		v71 = v71 - v87
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	goto L5
L26:
	;
	v109 = l1
	v110 = v104
	goto L27
L27:
	;
	v118 = v110
	v119 = v43
	goto L30
L28:
	;
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v109)+8))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	if v162 != 0 {
		v21 = v42
		v26 = v160 | v26
		goto L4
	} else {
		goto L46
	}
L29:
	;
	if v156 != 0 {
		goto L42
	} else {
		goto L43
	}
L30:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	if v122 == v123 {
		v145 = v122
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v156 = int32(0)
	goto L29
L32:
	;
	v147 = int32(1)
	if v145 != 0 {
		v118 = v118 + v147
		v119 = v119 + v147
		goto L30
	} else {
		goto L41
	}
L33:
	;
	if base.Ui32((v122-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v133 = v122 | int32(32)
	goto L36
L35:
	;
	v133 = v122
	goto L36
L36:
	;
	if base.Ui32((v123-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v142 = v123 | int32(32)
	goto L39
L38:
	;
	v142 = v123
	goto L39
L39:
	;
	if v133 == v142 {
		v145 = v133
		goto L32
	} else {
		goto L40
	}
L40:
	;
	v156 = v133 - v142
	goto L29
L41:
	;
	goto L31
L42:
	;
	v158 = v109 + int32(16)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v159 != 0 {
		v109 = v158
		v110 = v159
		goto L27
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	goto L28
L45:
	;
	goto L25
L46:
	;
	goto L25
L47:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v43
	F_errmsg(m, int32(761175), v12)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(522474), int32(1726), int32(346722))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
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
			v14 = *(*int32)(unsafe.Add(mBase, _consts[279]))
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
						v26 = F_convert_any_priv_string(m, v11, int32(1693632))
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
							v28 = F_convert_any_priv_string(m, v13, int32(1693632))
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
