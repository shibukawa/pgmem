package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_has_table_privilege_name_id(m *base.Module, l0 int32) int32 {
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
	var v36 int32
	_ = v36
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
		v20 = F_get_role_oid_or_public(m, v12)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v23 = F_convert_any_priv_string(m, v14, int32(1607232))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v27 = F_pg_class_aclcheck_ext(m, v11, v20, v23, v9+int32(15))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
					if v29 == int32(1) {
						v32 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v32)
						v36 = int32(0)
					} else {
						v36 = base.B2i32(v27 == int32(0))
					}
					m.G0 = v9 + int32(16)
					return v36
				}
			}
		}
	}
}
func F_has_table_privilege_name_name(m *base.Module, l0 int32) int32 {
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = F_get_role_oid_or_public(m, v4)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
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
							v26 = F_convert_any_priv_string(m, v11, int32(1607232))
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return int32(0)
							} else {
								v28 = F_pg_class_aclcheck(m, v23, v13, v26)
								mBase = m.M
								v29 = m.ExcPending
								if v29 != 0 {
									return int32(0)
								} else {
									return base.B2i32(v28 == int32(0))
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
	mBase := m.M
	_ = mBase
	*(*int64)(unsafe.Add(mBase, uint32(l1)+32)) = int64(0)
	return
}
func F_table_openrv(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = F_relation_openrv(m, l0, l1)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+119)))
		switch v13 - int32(99) {
		case 0, 6:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(151027844))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v25 + int32(4)
					F_errmsg(m, int32(656442), v6)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
						v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(v32)+119)))
						F_errdetail_relkind_not_supported(m, v33)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(468396), int32(147), int32(399594))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
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
		case 1, 2, 3, 4, 5:
			m.G0 = v6 + int32(16)
			return v8
		default:
			if v13 != int32(73) {
				m.G0 = v6 + int32(16)
				return v8
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(151027844))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v25 + int32(4)
						F_errmsg(m, int32(656442), v6)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
							v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(v32)+119)))
							F_errdetail_relkind_not_supported(m, v33)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(468396), int32(147), int32(399594))
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
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
		}
	}
}
func F_table_openrv_extended(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = F_relation_openrv_extended(m, l0, l1, l2)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			m.G0 = v7 + int32(16)
			return v9
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+48))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+119)))
			switch v16 - int32(99) {
			case 0, 6:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(151027844))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v28 + int32(4)
						F_errmsg(m, int32(656442), v7)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v9)+48))
							v36 = int32(*(*int8)(unsafe.Add(mBase, uint32(v35)+119)))
							F_errdetail_relkind_not_supported(m, v36)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(468396), int32(147), int32(399594))
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
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
			case 1, 2, 3, 4, 5:
				m.G0 = v7 + int32(16)
				return v9
			default:
				if v16 != int32(73) {
					m.G0 = v7 + int32(16)
					return v9
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(151027844))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v28 + int32(4)
							F_errmsg(m, int32(656442), v7)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								v35 = *(*int32)(unsafe.Add(mBase, uint32(v9)+48))
								v36 = int32(*(*int8)(unsafe.Add(mBase, uint32(v35)+119)))
								F_errdetail_relkind_not_supported(m, v36)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(468396), int32(147), int32(399594))
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
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
			}
		}
	}
}
func F_table_to_xml_and_xmlschema(m *base.Module, l0 int32) int32 {
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
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
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = F_text_to_cstring(m, v14)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v21 = F_table_open(m, v12, int32(1))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
				v25 = base.B2i32(v11 != int32(0))
				v26 = F_map_sql_table_to_xmlschema(m, v23, v12, v25, v18)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					F_sequence_close(m, v21, int32(0))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						F_initStringInfo(m, v9+int32(16))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v37 = F_DirectFunctionCall1Coll(m, int32(1562), int32(0), v12)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v37
								F_appendStringInfo(m, v9+int32(16), int32(186078), v9)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int32(0)
								} else {
									v45 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
									v46 = F_get_rel_name(m, v12)
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return int32(0)
									} else {
										v48 = F_query_to_xml_internal(m, v45, v46, v26, v25, v18)
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return int32(0)
										} else {
											v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
											v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
											v52 = F_cstring_to_text_with_len(m, v50, v51)
											mBase = m.M
											v53 = m.ExcPending
											if v53 != 0 {
												return int32(0)
											} else {
												m.G0 = v9 + int32(32)
												return v52
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
func F_table_to_xmlschema(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = F_text_to_cstring(m, v8)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v15 = F_table_open(m, v6, int32(1))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
				v20 = F_map_sql_table_to_xmlschema(m, v17, v6, base.B2i32(v5 != int32(0)), v12)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					F_sequence_close(m, v15, int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = F_cstring_to_text(m, v20)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							return v25
						}
					}
				}
			}
		}
	}
}
func F_try_table_open(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = F_try_relation_open(m, l0, l1)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 == int32(0) {
			m.G0 = v6 + int32(16)
			return v8
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+119)))
			switch v15 - int32(99) {
			case 0, 6:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(151027844))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v27 + int32(4)
						F_errmsg(m, int32(656442), v6)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
							v35 = int32(*(*int8)(unsafe.Add(mBase, uint32(v34)+119)))
							F_errdetail_relkind_not_supported(m, v35)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(468396), int32(147), int32(399594))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
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
			case 1, 2, 3, 4, 5:
				m.G0 = v6 + int32(16)
				return v8
			default:
				if v15 != int32(73) {
					m.G0 = v6 + int32(16)
					return v8
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(151027844))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = v27 + int32(4)
							F_errmsg(m, int32(656442), v6)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
								v35 = int32(*(*int8)(unsafe.Add(mBase, uint32(v34)+119)))
								F_errdetail_relkind_not_supported(m, v35)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(468396), int32(147), int32(399594))
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
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
			}
		}
	}
}
