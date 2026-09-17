package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_convert_any_priv_string(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int64
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int64
	_ = v154
	var v156 int32
	_ = v156
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = F_text_to_cstring(m, l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v19 = v13
	v24 = int64(0)
	goto L4
L3:
	;
	F_pfree(m, v13)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L51
	}
L4:
	;
	if v19 == int32(0) {
		goto L3
	} else {
		goto L6
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L47
	}
L6:
	;
	v27 = int32(44)
	v28 = F___strchrnul(m, v19, v27)
	mBase = m.M
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v30 == v27 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v34 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v34 = v28
	goto L10
L9:
	;
	v34 = int32(0)
	goto L10
L10:
	;
	goto L7
L11:
	;
	v35 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v34))) = uint8(v35)
	v40 = v34 + int32(1)
	goto L13
L12:
	;
	v40 = int32(0)
	goto L13
L13:
	;
	v41 = v19
	goto L14
L14:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v49-int32(9)))&base.B2i32(v49 != int32(32)) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v61 = F_strlen(m, v41)
	mBase = m.M
	if v61 <= int32(0) {
		v91 = v61
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v41 = v41 + int32(1)
	goto L14
L17:
	;
	goto L18
L18:
	;
	goto L15
L19:
	;
	v97 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v41+v91))) = uint8(v97)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v99 == v97 {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	v67 = v61
	goto L21
L21:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v67-int32(1)))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v75-int32(9)))&base.B2i32(v75 != int32(32)) != 0 {
		v91 = v67
		goto L19
	} else {
		goto L23
	}
L22:
	;
	v91 = int32(0)
	goto L19
L23:
	;
	v83 = int32(1)
	if v83 < v67 {
		v67 = v67 - v83
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
	v104 = v99
	v105 = l1
	goto L27
L27:
	;
	v112 = v104
	v113 = v41
	goto L30
L28:
	;
	v154 = *(*int64)(unsafe.Add(mBase, uint32(v105)+8))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	if v156 != 0 {
		v19 = v40
		v24 = v154 | v24
		goto L4
	} else {
		goto L46
	}
L29:
	;
	if v150 != 0 {
		goto L42
	} else {
		goto L43
	}
L30:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	if v116 == v117 {
		v139 = v116
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v150 = int32(0)
	goto L29
L32:
	;
	v141 = int32(1)
	if v139 != 0 {
		v112 = v112 + v141
		v113 = v113 + v141
		goto L30
	} else {
		goto L41
	}
L33:
	;
	if base.Ui32((v116-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v127 = v116 | int32(32)
	goto L36
L35:
	;
	v127 = v116
	goto L36
L36:
	;
	if base.Ui32((v117-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v136 = v117 | int32(32)
	goto L39
L38:
	;
	v136 = v117
	goto L39
L39:
	;
	if v127 == v136 {
		v139 = v127
		goto L32
	} else {
		goto L40
	}
L40:
	;
	v150 = v127 - v136
	goto L29
L41:
	;
	goto L31
L42:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v105)+16))
	if v151 != 0 {
		v104 = v151
		v105 = v105 + int32(16)
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
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v41
	F_errmsg(m, int32(_a_F_convert_any_priv_string_0), v11)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_convert_any_priv_string_1), int32(1726), int32(_a_F_convert_any_priv_string_2))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
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
	m.G0 = v11 + int32(16)
	return v24
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
			v14 = *(*int32)(unsafe.Add(mBase, _c_F_has_any_column_privilege_name[0]))
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
						v26 = F_convert_any_priv_string(m, v11, int32(_a_F_has_any_column_privilege_name_0))
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
							v28 = F_convert_any_priv_string(m, v13, int32(_a_F_has_any_column_privilege_name_name_0))
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
