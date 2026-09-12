package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_any_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(18532)
			F_errmsg(m, int32(192537), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(494416), int32(365), int32(279164))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func F_has_any_column_privilege_id(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v18)
		v21 = *(*int32)(unsafe.Add(mBase, _consts[168]))
		v23 = F_convert_any_priv_string(m, v14, int32(1656768))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v27 = F_pg_class_aclcheck_ext(m, v12, v21, v23, v10+int32(15))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				if v27 == int32(0) {
					v48 = int32(0)
					v52 = base.B2i32(v48 == int32(0))
					m.G0 = v10 + int32(16)
					return v52
				} else {
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
					if v31 == int32(1) {
						v34 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
						v52 = int32(0)
						m.G0 = v10 + int32(16)
						return v52
					} else {
						v40 = F_pg_attribute_aclcheck_all_ext(m, v12, v21, v23, int32(1), v10+int32(15))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
							if v42 != int32(1) {
								v48 = v40
								v52 = base.B2i32(v48 == int32(0))
							} else {
								v45 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
								v52 = int32(0)
							}
							m.G0 = v10 + int32(16)
							return v52
						}
					}
				}
			}
		}
	}
}
func F_has_any_column_privilege_id_id(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v19)
		v22 = F_convert_any_priv_string(m, v15, int32(1656768))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v26 = F_pg_class_aclcheck_ext(m, v12, v13, v22, v10+int32(15))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v26 == int32(0) {
					v47 = int32(0)
					v51 = base.B2i32(v47 == int32(0))
					m.G0 = v10 + int32(16)
					return v51
				} else {
					v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
					if v30 == int32(1) {
						v33 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v33)
						v51 = int32(0)
						m.G0 = v10 + int32(16)
						return v51
					} else {
						v39 = F_pg_attribute_aclcheck_all_ext(m, v12, v13, v22, int32(1), v10+int32(15))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
							if v41 != int32(1) {
								v47 = v39
								v51 = base.B2i32(v47 == int32(0))
							} else {
								v44 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v44)
								v51 = int32(0)
							}
							m.G0 = v10 + int32(16)
							return v51
						}
					}
				}
			}
		}
	}
}
func F_has_any_column_privilege_name_id(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v19)
		v21 = F_get_role_oid_or_public(m, v13)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v24 = F_convert_any_priv_string(m, v15, int32(1656768))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v28 = F_pg_class_aclcheck_ext(m, v12, v21, v24, v10+int32(15))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					if v28 == int32(0) {
						v49 = int32(0)
						v53 = base.B2i32(v49 == int32(0))
						m.G0 = v10 + int32(16)
						return v53
					} else {
						v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
						if v32 == int32(1) {
							v35 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v35)
							v53 = int32(0)
							m.G0 = v10 + int32(16)
							return v53
						} else {
							v41 = F_pg_attribute_aclcheck_all_ext(m, v12, v21, v24, int32(1), v10+int32(15))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
								if v43 != int32(1) {
									v49 = v41
									v53 = base.B2i32(v49 == int32(0))
								} else {
									v46 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v46)
									v53 = int32(0)
								}
								m.G0 = v10 + int32(16)
								return v53
							}
						}
					}
				}
			}
		}
	}
}
