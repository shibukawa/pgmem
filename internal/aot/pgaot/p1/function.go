package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_build_function_result_tupdesc_t(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
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
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+22)))
	v12 = v10 + v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+108))
	if v13 != int32(2249) {
		v47 = v2
		m.G0 = v8 + int32(16)
		return v47
	} else {
		v18 = F_heap_attisnull(m, l0, int32(21), int32(0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			if v18 != 0 {
				v47 = v2
				m.G0 = v8 + int32(16)
				return v47
			} else {
				v24 = F_heap_attisnull(m, l0, int32(22), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					if v24 != 0 {
						v47 = v2
						m.G0 = v8 + int32(16)
						return v47
					} else {
						v28 = F_SysCacheGetAttrNotNull(m, int32(47), l0, int32(21))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							v32 = F_SysCacheGetAttrNotNull(m, int32(47), l0, int32(22))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								v38 = F_SysCacheGetAttr(m, int32(47), l0, int32(23), v8+int32(15))
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int32(0)
								} else {
									v40 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+96)))
									v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
									if v42 != 0 {
										v43 = int32(0)
									} else {
										v43 = v38
									}
									v44 = F_build_function_result_tupdesc_d(m, v40, v28, v32, v43)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int32(0)
									} else {
										v47 = v44
										m.G0 = v8 + int32(16)
										return v47
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
func F_has_function_privilege_id(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v16)
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_has_function_privilege_id[0]))
		v22 = F_convert_any_priv_string(m, v12, int32(_a_F_has_function_privilege_id_0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v26 = F_object_aclcheck_ext(m, int32(1255), v10, v20, v22, v8+int32(15))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
				if v28 == int32(1) {
					v31 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
					v35 = int32(0)
				} else {
					v35 = base.B2i32(v26 == int32(0))
				}
				m.G0 = v8 + int32(16)
				return v35
			}
		}
	}
}
func F_has_function_privilege_id_id(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v18)
		v22 = F_convert_any_priv_string(m, v14, int32(_a_F_has_function_privilege_id_id_0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v26 = F_object_aclcheck_ext(m, int32(1255), v11, v12, v22, v9+int32(15))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
				if v28 == int32(1) {
					v31 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
					v35 = int32(0)
				} else {
					v35 = base.B2i32(v26 == int32(0))
				}
				m.G0 = v9 + int32(16)
				return v35
			}
		}
	}
}
func F_has_function_privilege_name_name(m *base.Module, l0 int32) int32 {
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v17 = F_pg_detoast_datum_packed(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = F_get_role_oid_or_public(m, v10)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v23 = F_text_to_cstring(m, v12)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = F_DirectFunctionCall1Coll(m, int32(1252), int32(0), v23)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						if v25 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(52461700))
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = v23
									F_errmsg(m, int32(_a_F_has_function_privilege_name_name_0), v8)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_has_function_privilege_name_name_1), int32(3565), int32(_a_F_has_function_privilege_name_name_2))
										mBase = m.M
										v44 = m.ExcPending
										if v44 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v47 = F_convert_any_priv_string(m, v17, int32(_a_F_has_function_privilege_name_name_3))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								v49 = F_object_aclcheck(m, int32(1255), v25, v19, v47)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									m.G0 = v8 + int32(16)
									return base.B2i32(v49 == int32(0))
								}
							}
						}
					}
				}
			}
		}
	}
}
