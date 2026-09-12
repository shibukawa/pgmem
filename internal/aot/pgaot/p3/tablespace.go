package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_tablespace_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v11 = F_table_open(m, int32(1213), int32(1))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		F_ScanKeyInit(m, v7, int32(1), int32(3), int32(184), l0)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v21 = F_table_beginscan_catalog(m, v11, int32(1), v7)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = F_heap_getnext(m, v21)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					if v23 != 0 {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
						v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+22)))
						v30 = F_pstrdup(m, v25+v26+int32(4))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							v34 = v30
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+188))
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
							m.T0[v37].(func(*base.Module, int32))(m, v21)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								F_sequence_close(m, v11, int32(1))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									m.G0 = v7 + int32(48)
									return v34
								}
							}
						}
					} else {
						v34 = int32(0)
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+188))
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
						m.T0[v37].(func(*base.Module, int32))(m, v21)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							F_sequence_close(m, v11, int32(1))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								m.G0 = v7 + int32(48)
								return v34
							}
						}
					}
				}
			}
		}
	}
}
func F_has_tablespace_privilege_id_name(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
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
			v14 = F_text_to_cstring(m, v6)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v17 = F_get_tablespace_oid(m, v14, int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v20 = F_convert_any_priv_string(m, v11, int32(1632752))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v22 = F_object_aclcheck(m, int32(1213), v17, v4, v20)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v22 == int32(0))
						}
					}
				}
			}
		}
	}
}
func F_remove_tablespace_symlink(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	v4 = m.G0
	v6 = v4 - int32(160)
	m.G0 = v6
	v12 = F___fstatat(m, int32(-100), l0, v6-int32(-64), int32(256))
	mBase = m.M
	if v12 < int32(0) {
		v16 = *(*int32)(unsafe.Add(mBase, _consts[40]))
		if v16 == int32(44) {
			m.G0 = v6 + int32(160)
			return
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				F_errcode_for_file_access(m)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
					F_errmsg(m, int32(293313), v6)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						F_errfinish(m, int32(491840), int32(893), int32(310868))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
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
		v34 = *(*int32)(unsafe.Add(mBase, uint32(v6)+68))
		v36 = v34 & int32(61440)
		if v36 != int32(40960) {
			if v36 != int32(16384) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = l0
						F_errmsg(m, int32(310983), v6+int32(16))
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return
						} else {
							F_errfinish(m, int32(491840), int32(922), int32(310868))
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v41 = F_rmdir(m, l0)
				mBase = m.M
				if int32(0) <= v41 {
					m.G0 = v6 + int32(160)
					return
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, _consts[40]))
					if v45 == int32(44) {
						m.G0 = v6 + int32(160)
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = l0
								F_errmsg(m, int32(292168), v6+int32(32))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									F_errfinish(m, int32(491840), int32(906), int32(310868))
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
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
				}
			}
		} else {
			v65 = F_unlink(m, l0)
			mBase = m.M
			if int32(0) <= v65 {
				m.G0 = v6 + int32(160)
				return
			} else {
				v69 = *(*int32)(unsafe.Add(mBase, _consts[40]))
				if v69 == int32(44) {
					m.G0 = v6 + int32(160)
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = l0
							F_errmsg(m, int32(292961), v6+int32(48))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								F_errfinish(m, int32(491840), int32(914), int32(310868))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
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
			}
		}
	}
}
