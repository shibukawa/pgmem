package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_has_table_privilege_name_id(m *base.Module, l0 int32) int64 {
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
	var v20 int32
	_ = v20
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v38 int64
	_ = v38
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v18 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v18)
		v20 = F_get_role_oid_or_public(m, v12)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			v23 = F_convert_any_priv_string(m, v14, int32(_a_F_has_table_privilege_name_id_0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int64(0)
			} else {
				v27 = F_pg_class_aclcheck_ext(m, v11, v20, v23, v9+int32(15))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int64(0)
				} else {
					v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
					if v29 == int32(1) {
						v32 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v32)
						v38 = int64(0)
					} else {
						v38 = base.I64_extend_i32_u(base.B2i32(v27 == int32(0)))
					}
					m.G0 = v9 + int32(16)
					return v38
				}
			}
		}
	}
}
func F_has_table_privilege_name_name(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			v13 = F_get_role_oid_or_public(m, v4)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				v15 = F_textToQualifiedNameList(m, v6)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int64(0)
				} else {
					v17 = F_makeRangeVarFromNameList(m, v15)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int64(0)
					} else {
						v19 = int32(0)
						v23 = F_RangeVarGetRelidExtended(m, v17, v19, v19, v19, v19)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int64(0)
						} else {
							v26 = F_convert_any_priv_string(m, v11, int32(_a_F_has_table_privilege_name_name_0))
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return int64(0)
							} else {
								v28 = F_pg_class_aclcheck(m, v23, v13, v26)
								mBase = m.M
								v29 = m.ExcPending
								if v29 != 0 {
									return int64(0)
								} else {
									return base.I64_extend_i32_u(base.B2i32(v28 == int32(0)))
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_table_block_parallelscan_reinitialize(m *base.Module, l0 int32, l1 int32) {
	var v5 int64
	_ = v5
	v5 = base.AtomicRmwXchg64(m, l1, int32(40), int64(0))
	return
}
func F_table_openrv(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v3 = F_relation_openrv(m, l0, l1)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		F_validate_relation_as_table(m, v3)
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v3
		}
	}
}
func F_table_openrv_extended(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	v4 = F_relation_openrv_extended(m, l0, l1, l2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 != 0 {
			F_validate_relation_as_table(m, v4)
			v9 = m.ExcPending
			if v9 != 0 {
				return int32(0)
			} else {
				return v4
			}
		} else {
			return v4
		}
	}
}
func F_table_to_xml_and_xmlschema(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int64(0)
	} else {
		v20 = F_text_to_cstring(m, v16)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			v22 = base.I32_wrap_i64(v14)
			v24 = F_table_open(m, v22, int32(1))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int64(0)
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+52))
				v28 = base.B2i32(v13 != int64(0))
				v29 = F_map_sql_table_to_xmlschema(m, v26, v22, v28, v20)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int64(0)
				} else {
					F_relation_close(m, v24, int32(0))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int64(0)
					} else {
						v35 = v11 + int32(16)
						F_initStringInfo(m, v35)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int64(0)
						} else {
							v42 = F_DirectFunctionCall1Coll(m, int32(1760), int32(0), v14&int64(4294967295))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int64(0)
							} else {
								*(*uint32)(unsafe.Add(mBase, uint32(v11))) = uint32(v42)
								F_appendStringInfo(m, v35, int32(_a_F_table_to_xml_and_xmlschema_0), v11)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int64(0)
								} else {
									v48 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
									v49 = F_get_rel_name(m, v22)
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return int64(0)
									} else {
										v51 = F_query_to_xml_internal(m, v48, v49, v29, v28, v20)
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int64(0)
										} else {
											v53 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
											v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
											v55 = F_cstring_to_text_with_len(m, v53, v54)
											mBase = m.M
											v56 = m.ExcPending
											if v56 != 0 {
												return int64(0)
											} else {
												m.G0 = v11 + int32(32)
												return base.I64_extend_i32_u(v55)
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
}
func F_table_to_xmlschema(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = F_text_to_cstring(m, v8)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v15 = F_table_open(m, v6, int32(1))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int64(0)
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
				v20 = F_map_sql_table_to_xmlschema(m, v17, v6, base.B2i32(v5 != int64(0)), v12)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int64(0)
				} else {
					F_relation_close(m, v15, int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int64(0)
					} else {
						v25 = F_cstring_to_text(m, v20)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(v25)
						}
					}
				}
			}
		}
	}
}
func F_try_table_open(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v3 = F_try_relation_open(m, l0, l1)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 != 0 {
			F_validate_relation_as_table(m, v3)
			v8 = m.ExcPending
			if v8 != 0 {
				return int32(0)
			} else {
				return v3
			}
		} else {
			return v3
		}
	}
}
