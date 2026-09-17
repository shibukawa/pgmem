package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_copy_plpgsql_datums(m *base.Module, l0 int32, l1 int32) {
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v80 int64
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v17 = F_palloc(m, v14<<(uint(int32(2))%32))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v17
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+512))
	v21 = F_palloc(m, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if int32(0) < v14 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+508))
	v29 = int32(0)
	v30 = v21
	v33 = v21
	goto L7
L5:
	;
	goto L6
L6:
	;
	m.G0 = v12 + int32(16)
	return
L7:
	;
	v38 = v29 << (uint(int32(2)) % 32)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v26+v38)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	switch v41 {
	case 0, 4:
		goto L11
	case 1, 3:
		v85 = v40
		v86 = v33
		goto L9
	case 2:
		goto L13
	default:
		goto L12
	}
L8:
	;
	goto L6
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25+v38))) = v85
	v90 = v29 + int32(1)
	if v90 != v14 {
		v29 = v90
		v30 = v86
		v33 = v86
		goto L7
	} else {
		goto L17
	}
L10:
	;
	v85 = v30
	v86 = v84
	goto L9
L11:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v40)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+48)) = v68
	v70 = *(*int64)(unsafe.Add(mBase, uint32(v40)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+40)) = v70
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v40)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+32)) = v72
	v74 = *(*int64)(unsafe.Add(mBase, uint32(v40)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+24)) = v74
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v40)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+16)) = v76
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v40)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+8)) = v78
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v40)))
	*(*int64)(unsafe.Add(mBase, uint32(v30))) = v80
	v84 = v30 + int32(56)
	goto L10
L12:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_copy_plpgsql_datums_0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	v42 = *(*int64)(unsafe.Add(mBase, uint32(v40)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+32)) = v42
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v40)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+24)) = v44
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v40)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+16)) = v46
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v40)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+8)) = v48
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v40)))
	*(*int64)(unsafe.Add(mBase, uint32(v30))) = v50
	v84 = v30 + int32(40)
	goto L10
L14:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v58
	F_errmsg_internal(m, int32(_a_F_copy_plpgsql_datums_1), v12)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_copy_plpgsql_datums_2), int32(1367), int32(_a_F_copy_plpgsql_datums_3))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L17:
	;
	goto L8
}
func F_plpgsql_adddatum(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_adddatum[0]))
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_adddatum[1]))
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_adddatum[2]))
	if v7 == v9 {
		*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_adddatum[2])) = v7 << (uint(int32(1)) % 32)
		v18 = F_repalloc(m, v5, v7<<(uint(int32(3))%32))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_adddatum[0])) = v18
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_adddatum[1]))
			v23 = v22
			v24 = v18
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v23
			*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_adddatum[1])) = v23 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v24+v23<<(uint(int32(2))%32)))) = l0
			return
		}
	} else {
		v23 = v7
		v24 = v5
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v23
		*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_adddatum[1])) = v23 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v24+v23<<(uint(int32(2))%32)))) = l0
		return
	}
}
func F_plpgsql_build_datatype_arrayof(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v10 == int32(0) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v14 = F_get_array_type(m, v13)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			if v14 == int32(0) {
				F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_build_datatype_arrayof_0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v48 = F_format_type_be(m, v47)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v48
							F_errmsg(m, int32(_a_F_plpgsql_build_datatype_arrayof_1), v8)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_plpgsql_build_datatype_arrayof_2), int32(2102), int32(_a_F_plpgsql_build_datatype_arrayof_3))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
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
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v23 = F_SearchSysCache1(m, int32(82), v14)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					if v23 == int32(0) {
						F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_build_datatype_arrayof_0))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v14
							F_errmsg_internal(m, int32(_a_F_plpgsql_build_datatype_arrayof_4), v8+int32(16))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_plpgsql_build_datatype_arrayof_2), int32(1960), int32(_a_F_plpgsql_build_datatype_arrayof_5))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v28 = F_build_datatype(m, v23, v21, v20, int32(0))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							F_ReleaseCatCache(m, v23)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								v32 = v28
								m.G0 = v8 + int32(32)
								return v32
							}
						}
					}
				}
			}
		}
	} else {
		v32 = l0
		m.G0 = v8 + int32(32)
		return v32
	}
}
func F_plpgsql_compile_error_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v9 != 0 {
		v10 = F_function_parse_error_transpose(m, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			if v10 != 0 {
				m.G0 = v6 + int32(16)
				return
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_error_callback[0]))
				if v13 != 0 {
					F_set_errcontext_domain(m, int32(_a_F_plpgsql_compile_error_callback_0))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_error_callback[0]))
						v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+196))
						*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v24
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v22
						F_errcontext_msg(m, int32(_a_F_plpgsql_compile_error_callback_1), v6)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							m.G0 = v6 + int32(16)
							return
						}
					}
				} else {
					m.G0 = v6 + int32(16)
					return
				}
			}
		}
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_error_callback[0]))
		if v15 == int32(0) {
			m.G0 = v6 + int32(16)
			return
		} else {
			F_set_errcontext_domain(m, int32(_a_F_plpgsql_compile_error_callback_0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_error_callback[0]))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+196))
				*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v24
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v22
				F_errcontext_msg(m, int32(_a_F_plpgsql_compile_error_callback_1), v6)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					m.G0 = v6 + int32(16)
					return
				}
			}
		}
	}
}
func F_plpgsql_exec_error_callback(m *base.Module, l0 int32) {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v9 != 0 {
		v18 = v9 + int32(12)
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
		if v20 != 0 {
			if v19 <= int32(0) {
				F_set_errcontext_domain(m, int32(_a_F_plpgsql_exec_error_callback_0))
				mBase = m.M
				v116 = m.ExcPending
				if v116 != 0 {
					return
				} else {
					v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+32))
					v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v119
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v118
					F_errcontext_msg(m, int32(_a_F_plpgsql_exec_error_callback_1), v5+int32(-48))
					mBase = m.M
					v126 = m.ExcPending
					if v126 != 0 {
						return
					} else {
						m.G0 = v7 - int32(-64)
						return
					}
				}
			} else {
				F_set_errcontext_domain(m, int32(_a_F_plpgsql_exec_error_callback_0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+56)) = v28
					*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = v19
					*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v27
					F_errcontext_msg(m, int32(_a_F_plpgsql_exec_error_callback_2), v5+int32(-16))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						m.G0 = v7 - int32(-64)
						return
					}
				}
			}
		} else {
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
			v38 = int32(0)
			if base.B2i32(v37 == v38)|base.B2i32(v19 <= v38) != 0 {
				F_set_errcontext_domain(m, int32(_a_F_plpgsql_exec_error_callback_0))
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return
				} else {
					v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+32))
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v108
					F_errcontext_msg(m, int32(_a_F_plpgsql_exec_error_callback_3), v7)
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return
					} else {
						m.G0 = v7 - int32(-64)
						return
					}
				}
			} else {
				F_set_errcontext_domain(m, int32(_a_F_plpgsql_exec_error_callback_0))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+32))
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
					switch v51 {
					case 0:
						v91 = int32(_a_F_plpgsql_exec_error_callback_4)
						v93 = v91
					case 1:
						v93 = int32(_a_F_plpgsql_exec_error_callback_5)
					case 2:
						v93 = int32(_a_F_plpgsql_exec_error_callback_6)
					case 3:
						v93 = int32(_a_F_plpgsql_exec_error_callback_7)
					case 4:
						v93 = int32(_a_F_plpgsql_exec_error_callback_8)
					case 5:
						v93 = int32(_a_F_plpgsql_exec_error_callback_9)
					case 6:
						v93 = int32(_a_F_plpgsql_exec_error_callback_10)
					case 7:
						v93 = int32(_a_F_plpgsql_exec_error_callback_11)
					case 8:
						v93 = int32(_a_F_plpgsql_exec_error_callback_12)
					case 9:
						v93 = int32(_a_F_plpgsql_exec_error_callback_13)
					case 10:
						v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+12)))
						if v63 != 0 {
							v64 = int32(_a_F_plpgsql_exec_error_callback_14)
						} else {
							v64 = int32(_a_F_plpgsql_exec_error_callback_15)
						}
						v93 = v64
					case 11:
						v93 = int32(_a_F_plpgsql_exec_error_callback_16)
					case 12:
						v93 = int32(_a_F_plpgsql_exec_error_callback_17)
					case 13:
						v93 = int32(_a_F_plpgsql_exec_error_callback_18)
					case 14:
						v93 = int32(_a_F_plpgsql_exec_error_callback_19)
					case 15:
						v93 = int32(_a_F_plpgsql_exec_error_callback_20)
					case 16:
						v93 = int32(_a_F_plpgsql_exec_error_callback_21)
					case 17:
						v93 = int32(_a_F_plpgsql_exec_error_callback_22)
					case 18:
						v93 = int32(_a_F_plpgsql_exec_error_callback_23)
					case 19:
						v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+12)))
						if v75 != 0 {
							v76 = int32(_a_F_plpgsql_exec_error_callback_24)
						} else {
							v76 = int32(_a_F_plpgsql_exec_error_callback_25)
						}
						v93 = v76
					case 20:
						v93 = int32(_a_F_plpgsql_exec_error_callback_26)
					case 21:
						v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+32)))
						if v80 != 0 {
							v81 = int32(_a_F_plpgsql_exec_error_callback_27)
						} else {
							v81 = int32(_a_F_plpgsql_exec_error_callback_28)
						}
						v93 = v81
					case 22:
						v93 = int32(_a_F_plpgsql_exec_error_callback_29)
					case 23:
						v93 = int32(_a_F_plpgsql_exec_error_callback_30)
					case 24:
						v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+16)))
						if v86 != 0 {
							v87 = int32(_a_F_plpgsql_exec_error_callback_31)
						} else {
							v87 = int32(_a_F_plpgsql_exec_error_callback_32)
						}
						v93 = v87
					case 25:
						v93 = int32(_a_F_plpgsql_exec_error_callback_33)
					case 26:
						v93 = int32(_a_F_plpgsql_exec_error_callback_34)
					default:
						v91 = int32(_a_F_plpgsql_exec_error_callback_35)
						v93 = v91
					}
					*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v93
					*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v19
					*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v47
					F_errcontext_msg(m, int32(_a_F_plpgsql_exec_error_callback_36), v5+int32(-32))
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return
					} else {
						m.G0 = v7 - int32(-64)
						return
					}
				}
			}
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
		if v12 == int32(0) {
			v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
			if v102 != 0 {
				F_set_errcontext_domain(m, int32(_a_F_plpgsql_exec_error_callback_0))
				mBase = m.M
				v116 = m.ExcPending
				if v116 != 0 {
					return
				} else {
					v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+32))
					v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v119
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v118
					F_errcontext_msg(m, int32(_a_F_plpgsql_exec_error_callback_1), v5+int32(-48))
					mBase = m.M
					v126 = m.ExcPending
					if v126 != 0 {
						return
					} else {
						m.G0 = v7 - int32(-64)
						return
					}
				}
			} else {
				F_set_errcontext_domain(m, int32(_a_F_plpgsql_exec_error_callback_0))
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return
				} else {
					v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+32))
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v108
					F_errcontext_msg(m, int32(_a_F_plpgsql_exec_error_callback_3), v7)
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return
					} else {
						m.G0 = v7 - int32(-64)
						return
					}
				}
			}
		} else {
			v18 = v12 + int32(4)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
			if v20 != 0 {
				if v19 <= int32(0) {
					F_set_errcontext_domain(m, int32(_a_F_plpgsql_exec_error_callback_0))
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return
					} else {
						v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+32))
						v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v119
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v118
						F_errcontext_msg(m, int32(_a_F_plpgsql_exec_error_callback_1), v5+int32(-48))
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return
						} else {
							m.G0 = v7 - int32(-64)
							return
						}
					}
				} else {
					F_set_errcontext_domain(m, int32(_a_F_plpgsql_exec_error_callback_0))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+56)) = v28
						*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = v19
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v27
						F_errcontext_msg(m, int32(_a_F_plpgsql_exec_error_callback_2), v5+int32(-16))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							m.G0 = v7 - int32(-64)
							return
						}
					}
				}
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
				v38 = int32(0)
				if base.B2i32(v37 == v38)|base.B2i32(v19 <= v38) != 0 {
					F_set_errcontext_domain(m, int32(_a_F_plpgsql_exec_error_callback_0))
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return
					} else {
						v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+32))
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v108
						F_errcontext_msg(m, int32(_a_F_plpgsql_exec_error_callback_3), v7)
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return
						} else {
							m.G0 = v7 - int32(-64)
							return
						}
					}
				} else {
					F_set_errcontext_domain(m, int32(_a_F_plpgsql_exec_error_callback_0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+32))
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
						switch v51 {
						case 0:
							v91 = int32(_a_F_plpgsql_exec_error_callback_4)
							v93 = v91
						case 1:
							v93 = int32(_a_F_plpgsql_exec_error_callback_5)
						case 2:
							v93 = int32(_a_F_plpgsql_exec_error_callback_6)
						case 3:
							v93 = int32(_a_F_plpgsql_exec_error_callback_7)
						case 4:
							v93 = int32(_a_F_plpgsql_exec_error_callback_8)
						case 5:
							v93 = int32(_a_F_plpgsql_exec_error_callback_9)
						case 6:
							v93 = int32(_a_F_plpgsql_exec_error_callback_10)
						case 7:
							v93 = int32(_a_F_plpgsql_exec_error_callback_11)
						case 8:
							v93 = int32(_a_F_plpgsql_exec_error_callback_12)
						case 9:
							v93 = int32(_a_F_plpgsql_exec_error_callback_13)
						case 10:
							v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+12)))
							if v63 != 0 {
								v64 = int32(_a_F_plpgsql_exec_error_callback_14)
							} else {
								v64 = int32(_a_F_plpgsql_exec_error_callback_15)
							}
							v93 = v64
						case 11:
							v93 = int32(_a_F_plpgsql_exec_error_callback_16)
						case 12:
							v93 = int32(_a_F_plpgsql_exec_error_callback_17)
						case 13:
							v93 = int32(_a_F_plpgsql_exec_error_callback_18)
						case 14:
							v93 = int32(_a_F_plpgsql_exec_error_callback_19)
						case 15:
							v93 = int32(_a_F_plpgsql_exec_error_callback_20)
						case 16:
							v93 = int32(_a_F_plpgsql_exec_error_callback_21)
						case 17:
							v93 = int32(_a_F_plpgsql_exec_error_callback_22)
						case 18:
							v93 = int32(_a_F_plpgsql_exec_error_callback_23)
						case 19:
							v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+12)))
							if v75 != 0 {
								v76 = int32(_a_F_plpgsql_exec_error_callback_24)
							} else {
								v76 = int32(_a_F_plpgsql_exec_error_callback_25)
							}
							v93 = v76
						case 20:
							v93 = int32(_a_F_plpgsql_exec_error_callback_26)
						case 21:
							v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+32)))
							if v80 != 0 {
								v81 = int32(_a_F_plpgsql_exec_error_callback_27)
							} else {
								v81 = int32(_a_F_plpgsql_exec_error_callback_28)
							}
							v93 = v81
						case 22:
							v93 = int32(_a_F_plpgsql_exec_error_callback_29)
						case 23:
							v93 = int32(_a_F_plpgsql_exec_error_callback_30)
						case 24:
							v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+16)))
							if v86 != 0 {
								v87 = int32(_a_F_plpgsql_exec_error_callback_31)
							} else {
								v87 = int32(_a_F_plpgsql_exec_error_callback_32)
							}
							v93 = v87
						case 25:
							v93 = int32(_a_F_plpgsql_exec_error_callback_33)
						case 26:
							v93 = int32(_a_F_plpgsql_exec_error_callback_34)
						default:
							v91 = int32(_a_F_plpgsql_exec_error_callback_35)
							v93 = v91
						}
						*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v93
						*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v19
						*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v47
						F_errcontext_msg(m, int32(_a_F_plpgsql_exec_error_callback_36), v5+int32(-32))
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return
						} else {
							m.G0 = v7 - int32(-64)
							return
						}
					}
				}
			}
		}
	}
}
func F_plpgsql_exec_function(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	v6 = l5
	v10 = m.G0
	v12 = v10 - int32(176)
	m.G0 = v12
	v15 = v12 + int32(32)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_plpgsql_estate_setup(m, v15, l0, v16, l2, l3)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+63)) = uint8(v6)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+120)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(_a_F_plpgsql_exec_function_0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+164)) = int32(_a_F_plpgsql_exec_function_1)
	v28 = int32(_a_F_plpgsql_exec_function_2)
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_function[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_function[0])) = v12 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v15
	F_copy_plpgsql_datums(m, v15, l0)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+164)) = int32(_a_F_plpgsql_exec_function_3)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if int32(0) < v40 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v44 = l1 + int32(20)
	v50 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+164)) = int32(_a_F_plpgsql_exec_function_4)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v12)+100))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v180+v181<<(uint(int32(2))%32))))
	v186 = int32(0)
	F_assign_simple_var(m, v12+int32(32), v185, v186, v186, v186)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L40
	}
L7:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v12)+100))
	v57 = int32(2)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(72)+v50<<(uint(v57)%32))))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v56+v60<<(uint(v57)%32))))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	switch v65 {
	case 0:
		goto L13
	default:
		goto L11
	case 2:
		goto L12
	}
L8:
	;
	goto L6
L9:
	;
	v164 = v50 + int32(1)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v164 < v165 {
		v50 = v164
		goto L7
	} else {
		goto L39
	}
L10:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v12)+104))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v81)+2))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+8))
	F_MemoryContextSetParent(m, v152, v150)
	mBase = m.M
	goto L37
L11:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_exec_function_5))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L34
	}
L12:
	;
	v103 = v44 + v50<<(uint(int32(3))%32)
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+4)))
	if v104 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L13:
	;
	v70 = v44 + v50<<(uint(int32(3))%32)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+4)))
	F_assign_simple_var(m, v12+int32(32), v64, v71, v72, int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+44)))
	if v76 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+12)))
	if v78 != int32(_a_F_plpgsql_exec_function_6) {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v64)+40))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v82 != int32(1) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+20)))
	if v88 != int32(1) {
		goto L9
	} else {
		goto L19
	}
L18:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+1)))
	switch v85 - int32(2) {
	case 0:
		goto L9
	case 1:
		goto L10
	default:
		goto L17
	}
L19:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v12)+104))
	v95 = F_expand_array(m, v81, v93, int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_assign_simple_var(m, v12+int32(32), v64, v95, int32(0), int32(1))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	goto L9
L22:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v12)+136))
	if v118 != 0 {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	F_exec_move_row_from_datum(m, v12+int32(32), v64, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v114 = int32(0)
	F_exec_move_row(m, v12+int32(32), v64, v114, v114)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L27
	}
L26:
	;
	goto L22
L27:
	;
	goto L22
L28:
	;
	F_SPI_freetuptable(m, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v121 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+136)) = v121
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v12)+152))
	if v123 == v121 {
		goto L9
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v123)+20))
	F_MemoryContextReset(m, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L9
L34:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+508))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v133+v50<<(uint(int32(2))%32))))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v138
	F_errmsg_internal(m, int32(_a_F_plpgsql_exec_function_7), v12)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_plpgsql_exec_function_8), int32(614), int32(_a_F_plpgsql_exec_function_9))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	F_assign_simple_var(m, v12+int32(32), v64, v151+int32(12), int32(0), int32(1))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	goto L9
L39:
	;
	goto L8
L40:
	;
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_function[1]))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	if v193 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v232 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_function[2]))
	if v232 != 0 {
		goto L53
	} else {
		goto L54
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+164)) = int32(0)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+156)) = v198
	v229 = v198
	goto L41
L43:
	;
	goto L44
L44:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	if v200 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
	if v222 == int32(0) {
		v229 = v220
		goto L41
	} else {
		goto L51
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+164)) = int32(0)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+156)) = v205
	v220 = v205
	v221 = v193
	goto L45
L47:
	;
	goto L48
L48:
	;
	m.T0[v200].(func(*base.Module, int32, int32))(m, v12+int32(32), l0)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_function[1]))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	v214 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+164)) = v214
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+156)) = v216
	if v213 == v214 {
		v229 = v216
		goto L41
	} else {
		goto L50
	}
L50:
	;
	v220 = v216
	v221 = v213
	goto L45
L51:
	;
	m.T0[v222].(func(*base.Module, int32, int32))(m, v12+int32(32), v220)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v229 = v220
	goto L41
L53:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v236 = v12 + int32(32)
	v237 = F_exec_stmt_block(m, v236, v229)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_function[1]))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)))
	if v241 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+156)) = int32(0)
	if v237 == int32(2) {
		goto L64
	} else {
		goto L65
	}
L59:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v241)+16))
	if v244 == int32(0) {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	m.T0[v244].(func(*base.Module, int32, int32))(m, v236, v229)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	goto L58
L62:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_exec_function_5))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L125
	}
L63:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_exec_function_5))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L121
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+164)) = int32(_a_F_plpgsql_exec_function_10)
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+48)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)) = uint8(v256)
	v259 = l1 + int32(16)
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+61)))
	if v260 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+164)) = int32(0)
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_exec_function_5))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L117
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+164)) = int32(_a_F_plpgsql_exec_function_11)
	v391 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_function[1]))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	if v392 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L68:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v12)+88))
	if v263 == int32(0) {
		goto L63
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	if v256&int32(1) == int32(0) {
		goto L78
	} else {
		goto L79
	}
L71:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	if v266 != int32(383) {
		goto L63
	} else {
		goto L72
	}
L72:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263)+12)))
	if v269&int32(2) == int32(0) {
		goto L62
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v263)+16)) = int32(2)
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
	if v276 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v263)+24)) = v276
	v278 = int32(_a_F_plpgsql_exec_function_12)
	v279 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_function[3]))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_function[3])) = v281
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
	v284 = F_CreateTupleDescCopy(m, v283)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(0)
	v292 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v259))) = uint8(v292)
	goto L67
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v263)+28)) = v284
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_function[3])) = v279
	goto L76
L78:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+60)))
	if v299 == int32(1) {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+62)))
	if v372 != int32(1) {
		goto L67
	} else {
		goto L104
	}
L81:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if base.B2i32(v302 == int32(2249))|base.B2i32(v302 != v298) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	goto L83
L83:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v360 = int32(-1)
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v363 = F_exec_cast_value(m, v12+int32(32), v359, v259, v298, v360, v361, v360)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L100
	}
L84:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v311 = F_SPI_datumTransfer(m, v309, int32(-1))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v318 = F_get_call_result_type(m, l1, v12+int32(16), v12+int32(12))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L92
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v311
	goto L67
L88:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_exec_function_5))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L97
	}
L89:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v341 = F_SPI_datumTransfer(m, v339, int32(-1))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L96
	}
L90:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	F_coerce_function_result_tuple(m, v12+int32(32), v329)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L94
	}
L91:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	F_coerce_function_result_tuple(m, v12+int32(32), v324)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L93
	}
L92:
	;
	switch v318 - int32(1) {
	case 0:
		goto L91
	case 1:
		goto L90
	case 2:
		goto L89
	default:
		goto L88
	}
L93:
	;
	goto L67
L94:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v333 = int32(0)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	F_domain_check(m, v332, v333, v334, v333, v333)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	goto L67
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v341
	goto L67
L97:
	;
	F_errmsg_internal(m, int32(_a_F_plpgsql_exec_function_13), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_plpgsql_exec_function_8), int32(746), int32(_a_F_plpgsql_exec_function_9))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v363
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	if v366 != 0 {
		goto L67
	} else {
		goto L101
	}
L101:
	;
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+60)))
	if v367 != 0 {
		goto L67
	} else {
		goto L102
	}
L102:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v369 = F_SPI_datumTransfer(m, v363, v368)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v369
	goto L67
L104:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	v379 = int32(-1)
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v382 = F_exec_cast_value(m, v12+int32(32), v377, v259, v378, v379, v380, v379)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v382
	goto L67
L106:
	;
	v404 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_function[4]))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v404)+8))
	F_pfree(m, v404)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L110
	}
L107:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v392)+8))
	if v395 == int32(0) {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	m.T0[v395].(func(*base.Module, int32, int32))(m, v12+int32(32), l0)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	goto L106
L110:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_function[4])) = v405
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v12)+152))
	F_FreeExprContext(m, v410, int32(1))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v414 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+152)) = v414
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v12)+136))
	if v416 == v414 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_function[0])) = v431
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	m.G0 = v12 + int32(176)
	return v433
L113:
	;
	F_SPI_freetuptable(m, v416)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v421 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+136)) = v421
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v12)+152))
	if v423 == v421 {
		goto L112
	} else {
		goto L115
	}
L115:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v423)+20))
	F_MemoryContextReset(m, v426)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	goto L112
L117:
	;
	F_errcode(m, int32(83887490))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	F_errmsg(m, int32(_a_F_plpgsql_exec_function_14), int32(0))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_plpgsql_exec_function_8), int32(641), int32(_a_F_plpgsql_exec_function_9))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L121:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	F_errmsg(m, int32(_a_F_plpgsql_exec_function_15), int32(0))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_plpgsql_exec_function_8), int32(659), int32(_a_F_plpgsql_exec_function_9))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	F_errmsg(m, int32(_a_F_plpgsql_exec_function_16), int32(0))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_plpgsql_exec_function_8), int32(664), int32(_a_F_plpgsql_exec_function_9))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_plpgsql_extra_errors_assign_hook(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_extra_errors_assign_hook[0])) = v4
	return
}
func F_plpgsql_inline_handler(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v59 int64
	_ = v59
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v169 int64
	_ = v169
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v223 int32
	_ = v223
	var v235 int32
	_ = v235
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v270 int32
	_ = v270
	var v279 int32
	_ = v279
	var v288 int32
	_ = v288
	var v289 int64
	_ = v289
	var v301 int32
	_ = v301
	var v310 int32
	_ = v310
	var v320 int32
	_ = v320
	var v321 int64
	_ = v321
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(256)
	m.G0 = v14
	v19 = v2
	v20 = v2
	v21 = v2
	v22 = v2
	v23 = v2
	v24 = v2
	v25 = v2
	v26 = int32(-1)
	v27 = v2
	goto L2
L1:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2:
	;
	goto L5
L3:
	;
	m.G0 = v14 + int32(256)
	return v136
L4:
	;
	goto L3
L5:
	;
	if v26 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v320 = int32(m.ExcTag)
	v321 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v320 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+13)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+232)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v14)+228)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v14)+236)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v14)+244)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v14)+248)) = v20
	v39 = v30 + int32(13)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+252)) = v39
	F_SPI_connect_ext(m, v31^int32(1))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	v111 = v19
	v112 = v20
	v113 = v21
	v114 = v22
	v115 = v23
	v116 = v24
	v117 = v25
	v119 = v27
	goto L10
L10:
	;
	if v119 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+232)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v14)+228)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v14)+236)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v14)+244)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v14)+248)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v14)+252)) = v39
	v53 = F_plpgsql_compile_inline(m, v45)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v53)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+24)) = v55 + int64(1)
	v59 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+192)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v14)+220)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v14)+212)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v14)+176)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v14)+184)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v14)+200)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+228)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v14)+232)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v14)+236)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v21
	v76 = v53 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+244)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v14)+248)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v14)+252)) = v39
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_inline_handler[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+196)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v14)+208)) = v14 + int32(176)
	v86 = F_CreateExecutorState(m)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+232)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v14)+228)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v14)+236)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v14)+244)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v14)+248)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v14)+252)) = v39
	v97 = F_ResourceOwnerCreate(m, int32(0), int32(_a_F_plpgsql_inline_handler_0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_inline_handler[1]))
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_inline_handler[2]))
	goto L15
L15:
	;
	v104 = v14 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v104)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v104))) = v14 + int32(12)
	goto L18
L16:
	;
	v111 = v97
	v112 = v53
	v113 = v86
	v114 = v76
	v115 = v102
	v116 = v100
	v117 = v39
	v119 = int32(0)
	goto L10
L18:
	;
	goto L16
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_inline_handler[2])) = v14 + int32(16)
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+232)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v14)+228)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v14)+236)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v14)+244)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v14)+248)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v14)+252)) = v117
	v136 = F_plpgsql_exec_function(m, v112, v14+int32(208), v113, v111, v111, v126)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L7
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_inline_handler[1])) = v116
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_inline_handler[2])) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v14)+228)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v14)+232)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v14)+236)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v14)+244)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v14)+248)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v14)+252)) = v117
	v248 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_inline_handler[3]))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)+8))
	goto L33
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_inline_handler[1])) = v116
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_inline_handler[2])) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v14)+228)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v14)+232)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v14)+236)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v14)+244)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v14)+248)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v14)+252)) = v117
	F_FreeExecutorState(m, v113)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+232)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v14)+228)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v14)+236)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v14)+244)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v14)+248)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v14)+252)) = v117
	F_ReleaseAllPlanCacheRefsInOwner(m, v111)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+232)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v14)+228)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v14)+236)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v14)+244)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v14)+248)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v14)+252)) = v117
	F_ResourceOwnerDelete(m, v111)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v169 = *(*int64)(unsafe.Add(mBase, uint32(v114)))
	*(*int64)(unsafe.Add(mBase, uint32(v114))) = v169 - int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+232)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v14)+228)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v14)+236)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v14)+244)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v14)+248)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v14)+252)) = v117
	F_plpgsql_free_function_memory(m, v112)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+232)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v14)+228)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v14)+236)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v14)+244)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v14)+248)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v14)+252)) = v117
	v189 = F_SPI_finish(m)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	if v189 == int32(2) {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+232)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v14)+228)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v14)+236)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v14)+244)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v14)+248)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v14)+252)) = v117
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_inline_handler_1))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+232)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v14)+228)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v14)+236)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v14)+244)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v14)+248)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v14)+252)) = v117
	v211 = F_SPI_result_code_string(m, v189)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+232)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v14)+228)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v14)+236)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v14)+244)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v14)+248)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v14)+252)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v211
	F_errmsg_internal(m, int32(_a_F_plpgsql_inline_handler_2), v14)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+232)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v14)+228)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v14)+236)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v14)+244)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v14)+248)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v14)+252)) = v117
	F_errfinish(m, int32(_a_F_plpgsql_inline_handler_3), int32(427), int32(_a_F_plpgsql_inline_handler_4))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	goto L1
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+232)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v14)+228)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v14)+236)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v14)+244)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v14)+248)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v14)+252)) = v117
	v258 = int32(0)
	F_plpgsql_subxact_cb(m, int32(2), v249, v258, v258)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+232)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v14)+228)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v14)+236)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v14)+244)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v14)+248)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v14)+252)) = v117
	F_FreeExecutorState(m, v113)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+232)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v14)+228)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v14)+236)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v14)+244)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v14)+248)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v14)+252)) = v117
	F_ReleaseAllPlanCacheRefsInOwner(m, v111)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+232)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v14)+228)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v14)+236)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v14)+244)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v14)+248)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v14)+252)) = v117
	F_ResourceOwnerDelete(m, v111)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	v289 = *(*int64)(unsafe.Add(mBase, uint32(v114)))
	*(*int64)(unsafe.Add(mBase, uint32(v114))) = v289 - int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+232)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v14)+228)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v14)+236)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v14)+244)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v14)+248)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v14)+252)) = v117
	F_plpgsql_free_function_memory(m, v112)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+232)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v14)+228)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v14)+236)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v14)+244)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v14)+248)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v14)+252)) = v117
	F_pg_re_throw(m)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	goto L1
L40:
	;
	v325 = int32(v321)
	m.G0 = v14
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v325)+4))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v328)))
	if v14+int32(12) == v331 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	m.ExcPending = 1
	goto L49
L42:
	;
	if v335 != 0 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v328)+4))
	v335 = v333
	goto L45
L44:
	;
	v335 = int32(0)
	goto L45
L45:
	;
	goto L42
L46:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v14)+252))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v14)+248))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v14)+244))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v14)+240))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v14)+236))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v14)+232))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v14)+228))
	v19 = v340
	v20 = v337
	v21 = v339
	v22 = v338
	v23 = v341
	v24 = v342
	v25 = v336
	v26 = v335
	v27 = v327
	goto L2
L47:
	;
	goto L48
L48:
	;
	F___wasm_longjmp(m, v328, v327)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	return int32(0)
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_plpgsql_push_back_token(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v45 int64
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v12
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+72))
	if int32(4) <= v17 {
		F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_push_back_token_0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_plpgsql_push_back_token_1), int32(0))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_plpgsql_push_back_token_2), int32(388), int32(_a_F_plpgsql_push_back_token_3))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v34 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
		*(*int32)(unsafe.Add(mBase, uint32(v16+v17<<(uint(int32(2))%32))+76)) = l0
		v39 = *(*int32)(unsafe.Add(mBase, uint32(v16)+72))
		v42 = v16 + v39*int32(24)
		v43 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v42)+100)) = v43
		v45 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
		*(*int64)(unsafe.Add(mBase, uint32(v42)+92)) = v45
		*(*int32)(unsafe.Add(mBase, uint32(v42)+112)) = v34
		*(*int32)(unsafe.Add(mBase, uint32(v42)+108)) = v33
		v49 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
		v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+72))
		*(*int32)(unsafe.Add(mBase, uint32(v49)+72)) = v50 + int32(1)
		m.G0 = v10 + int32(16)
		return
	}
}
func F_plpgsql_subxact_cb(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v5 = int32(1)
	if base.Ui32(v5) < base.Ui32(l0-v5) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_subxact_cb[0]))
	if v10 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = v10
	goto L4
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v19 != l1 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	goto L1
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	F_FreeExprContext(m, v21, base.B2i32(l0 == int32(1)))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_subxact_cb[0]))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	F_pfree(m, v25)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_subxact_cb[0])) = v26
	if v26 != 0 {
		v17 = v26
		goto L4
	} else {
		goto L10
	}
L10:
	;
	goto L5
}
func F_plpgsql_validator(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v120 int32
	_ = v120
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int64
	_ = v137
	var v151 int32
	_ = v151
	var v156 int64
	_ = v156
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(160)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_CheckFunctionValidatorAccess(m, v14, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_validator_0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L5
	} else {
		goto L73
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_validator_0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L5
	} else {
		goto L68
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_validator_0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L5
	} else {
		goto L64
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_validator_0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L5
	} else {
		goto L61
	}
L5:
	;
	return int32(0)
L6:
	;
	if v16 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v21 = F_SearchSysCache1(m, int32(47), v15)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	m.G0 = v11 + int32(160)
	return int32(0)
L10:
	;
	if v21 == int32(0) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+22)))
	v28 = v26 + v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+108))
	v30 = F_get_typtype(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L13
	}
L12:
	;
	v70 = F_get_func_arg_info(m, v21, v11+int32(156), v11+int32(152), v11+int32(148))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L5
	} else {
		goto L26
	}
L13:
	;
	if v30 != int32(112) {
		v62 = v2
		v63 = v2
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)+108))
	if v34 <= int32(3830) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v62 = int32(0)
	v63 = v59
	goto L12
L16:
	;
	v59 = int32(0)
	goto L15
L17:
	;
	switch v34 - int32(2249) {
	case 0, 28, 29, 34:
		goto L16
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 31, 32, 33:
		goto L1
	case 30:
		v62 = int32(1)
		v63 = v2
		goto L12
	default:
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if base.Ui32(v34-int32(_a_F_plpgsql_validator_1)) < base.Ui32(int32(4)) {
		goto L16
	} else {
		goto L22
	}
L20:
	;
	if base.B2i32(v34 == int32(2776))|base.B2i32(v34 == int32(3500)) != 0 {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L1
L22:
	;
	switch v34 - int32(3831) {
	case 0:
		goto L16
	case 1, 2, 3, 4, 5, 6:
		goto L1
	case 7:
		goto L24
	default:
		goto L23
	}
L23:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v34-int32(_a_F_plpgsql_validator_2)) {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	v59 = int32(1)
	goto L15
L25:
	;
	goto L16
L26:
	;
	if int32(0) < v70 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v74 = int32(0)
	goto L30
L28:
	;
	goto L29
L29:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_validator[0])))
	if v131 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L30:
	;
	v83 = v74 << (uint(int32(2)) % 32)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v11)+156))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v83+v84)))
	v87 = F_get_typtype(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L5
	} else {
		goto L33
	}
L31:
	;
	goto L29
L32:
	;
	v120 = v74 + int32(1)
	if v120 != v70 {
		v74 = v120
		goto L30
	} else {
		goto L46
	}
L33:
	;
	if v87 != int32(112) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v11)+156))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v91+v83)))
	if v93 <= int32(3830) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v93 != int32(2249) {
		goto L2
	} else {
		goto L45
	}
L36:
	;
	if v93 <= int32(2775) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	if base.B2i32(base.Ui32(v93-int32(_a_F_plpgsql_validator_1)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v93-int32(_a_F_plpgsql_validator_2)) < base.Ui32(int32(2)))|base.B2i32(v93 == int32(3831)) != 0 {
		goto L32
	} else {
		goto L44
	}
L39:
	;
	switch v93 - int32(2277) {
	case 0, 6:
		goto L32
	case 1, 2, 3, 4, 5:
		goto L2
	default:
		goto L35
	}
L40:
	;
	goto L41
L41:
	;
	if v93 == int32(2776) {
		goto L32
	} else {
		goto L42
	}
L42:
	;
	if v93 != int32(3500) {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	goto L32
L44:
	;
	goto L2
L45:
	;
	goto L32
L46:
	;
	goto L31
L47:
	;
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L5
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	F_ReleaseCatCache(m, v21)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L5
	} else {
		goto L60
	}
L50:
	;
	v137 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+96)) = v137
	*(*int64)(unsafe.Add(mBase, uint32(v11)+112)) = v137
	*(*int64)(unsafe.Add(mBase, uint32(v11)+140)) = v137
	*(*int64)(unsafe.Add(mBase, uint32(v11)+132)) = v137
	*(*int64)(unsafe.Add(mBase, uint32(v11)+104)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v11)+120)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+100)) = v15
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_validator[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+116)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v11)+128)) = v11 + int32(96)
	if v62 != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v182 = F_plpgsql_compile(m, v11+int32(128), int32(1))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L5
	} else {
		goto L57
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+132)) = v11 + int32(52)
	goto L51
L53:
	;
	v156 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+88)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v11)+80)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v11)+72)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v11)+64)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v11)+56)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = int32(442)
	goto L52
L54:
	;
	goto L55
L55:
	;
	if v63 == int32(0) {
		goto L51
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+56)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = int32(441)
	goto L52
L57:
	;
	v184 = F_SPI_finish(m)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L5
	} else {
		goto L58
	}
L58:
	;
	if v184 != int32(2) {
		goto L3
	} else {
		goto L59
	}
L59:
	;
	goto L49
L60:
	;
	goto L9
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v15
	F_errmsg_internal(m, int32(_a_F_plpgsql_validator_3), v11)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_plpgsql_validator_4), int32(462), int32(_a_F_plpgsql_validator_5))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L5
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	v221 = F_SPI_result_code_string(m, v184)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v221
	F_errmsg_internal(m, int32(_a_F_plpgsql_validator_6), v11+int32(48))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L5
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_plpgsql_validator_4), int32(544), int32(_a_F_plpgsql_validator_5))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L5
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v11)+156))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v241+v74<<(uint(int32(2))%32))))
	v246 = F_format_type_be(m, v245)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L5
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v246
	F_errmsg(m, int32(_a_F_plpgsql_validator_7), v11+int32(32))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L5
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_plpgsql_validator_4), int32(497), int32(_a_F_plpgsql_validator_5))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L5
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L5
	} else {
		goto L74
	}
L74:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v28)+108))
	v268 = F_format_type_be(m, v267)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L5
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v268
	F_errmsg(m, int32(_a_F_plpgsql_validator_8), v11+int32(16))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L5
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_plpgsql_validator_4), int32(481), int32(_a_F_plpgsql_validator_5))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L5
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_plpgsql_xact_cb(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	v7 = int32(0)
	if base.B2i32(int32(1)<<(uint(l0)%32)&int32(19) == v7)|base.B2i32(base.Ui32(int32(4)) < base.Ui32(l0)) == v7 {
		*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_xact_cb[0])) = int32(0)
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_xact_cb[1]))
		if v18 != 0 {
			F_FreeExecutorState(m, v18)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v22 = int32(0)
				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_xact_cb[1])) = v22
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_xact_cb[2]))
				if v25 == v22 {
					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_xact_cb[2])) = int32(0)
					return
				} else {
					F_ReleaseAllPlanCacheRefsInOwner(m, v25)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_xact_cb[2])) = int32(0)
						return
					}
				}
			}
		} else {
			v22 = int32(0)
			*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_xact_cb[1])) = v22
			v25 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_xact_cb[2]))
			if v25 == v22 {
				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_xact_cb[2])) = int32(0)
				return
			} else {
				F_ReleaseAllPlanCacheRefsInOwner(m, v25)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_xact_cb[2])) = int32(0)
					return
				}
			}
		}
	} else {
		if l0&int32(-2) != int32(2) {
		} else {
			v35 = int32(0)
			*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_xact_cb[1])) = v35
			*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_xact_cb[0])) = v35
			*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_xact_cb[2])) = int32(0)
		}
		return
	}
}
