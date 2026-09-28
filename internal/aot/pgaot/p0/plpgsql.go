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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v45 int64
	_ = v45
	var v47 int64
	_ = v47
	var v49 int64
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v16 = F_palloc_mul(m, int32(4), v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v16
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+512))
	v20 = F_palloc(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if int32(0) < v15 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+508))
	v28 = int32(0)
	v29 = v20
	v32 = v20
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
	v37 = v28 << (uint(int32(2)) % 32)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v25+v37)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	switch v40 {
	case 0, 4:
		goto L11
	case 1, 3:
		v84 = v39
		v85 = v32
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
	*(*int32)(unsafe.Add(mBase, uint32(v24+v37))) = v84
	v89 = v28 + int32(1)
	if v89 != v15 {
		v28 = v89
		v29 = v85
		v32 = v85
		goto L7
	} else {
		goto L17
	}
L10:
	;
	v84 = v29
	v85 = v83
	goto L9
L11:
	;
	v67 = *(*int64)(unsafe.Add(mBase, uint32(v39)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+48)) = v67
	v69 = *(*int64)(unsafe.Add(mBase, uint32(v39)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+40)) = v69
	v71 = *(*int64)(unsafe.Add(mBase, uint32(v39)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+32)) = v71
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v39)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+24)) = v73
	v75 = *(*int64)(unsafe.Add(mBase, uint32(v39)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = v75
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v39)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = v77
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = v79
	v83 = v29 + int32(56)
	goto L10
L12:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_copy_plpgsql_datums_0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v39)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+32)) = v41
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v39)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+24)) = v43
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v39)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = v45
	v47 = *(*int64)(unsafe.Add(mBase, uint32(v39)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = v47
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = v49
	v83 = v29 + int32(40)
	goto L10
L14:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v57
	F_errmsg_internal(m, int32(_a_F_copy_plpgsql_datums_1), v12)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_copy_plpgsql_datums_2), int32(1398), int32(_a_F_copy_plpgsql_datums_3))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
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
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v49 = F_format_type_be(m, v48)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v49
							F_errmsg(m, int32(_a_F_plpgsql_build_datatype_arrayof_1), v8)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_plpgsql_build_datatype_arrayof_2), int32(2125), int32(_a_F_plpgsql_build_datatype_arrayof_3))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
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
				v24 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(v14))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					if v24 == int32(0) {
						F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_build_datatype_arrayof_0))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v14
							F_errmsg_internal(m, int32(_a_F_plpgsql_build_datatype_arrayof_4), v8+int32(16))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_plpgsql_build_datatype_arrayof_2), int32(1983), int32(_a_F_plpgsql_build_datatype_arrayof_5))
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v29 = F_build_datatype(m, v24, v21, v20, int32(0))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							F_ReleaseCatCache(m, v24)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								v33 = v29
								m.G0 = v8 + int32(32)
								return v33
							}
						}
					}
				}
			}
		}
	} else {
		v33 = l0
		m.G0 = v8 + int32(32)
		return v33
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
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+192))
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
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+192))
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
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v9 != 0 {
		v18 = v9 + int32(12)
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
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
					v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
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
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
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
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
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
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
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
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
		if v12 == int32(0) {
			v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
			if v102 != 0 {
				F_set_errcontext_domain(m, int32(_a_F_plpgsql_exec_error_callback_0))
				mBase = m.M
				v116 = m.ExcPending
				if v116 != 0 {
					return
				} else {
					v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+32))
					v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
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
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
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
						v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
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
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
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
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
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
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
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
func F_plpgsql_exec_function(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int64
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int64
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int64
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v317 int64
	_ = v317
	var v319 int64
	_ = v319
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int64
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int64
	_ = v347
	var v349 int64
	_ = v349
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v367 int64
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int64
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int64
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v385 int64
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int64
	_ = v390
	var v391 int32
	_ = v391
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v442 int64
	_ = v442
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	v6 = l5
	v11 = m.G0
	v13 = v11 - int32(176)
	m.G0 = v13
	v16 = v13 + int32(24)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_plpgsql_estate_setup(m, v16, l0, v17, l2, l3)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+63)) = uint8(v6)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+120)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(_a_F_plpgsql_exec_function_0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+164)) = int32(_a_F_plpgsql_exec_function_1)
	v29 = int32(_a_F_plpgsql_exec_function_2)
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_function[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_function[0])) = v13 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v16
	F_copy_plpgsql_datums(m, v16, l0)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+164)) = int32(_a_F_plpgsql_exec_function_3)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if int32(0) < v41 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v45 = l1 + int32(24)
	v51 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+164)) = int32(_a_F_plpgsql_exec_function_4)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v13)+100))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v188+v189<<(uint(int32(2))%32))))
	v195 = int32(0)
	F_assign_simple_var(m, v13+int32(24), v193, int64(0), v195, v195)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L40
	}
L7:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v13)+100))
	v59 = int32(2)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(72)+v51<<(uint(v59)%32))))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v58+v62<<(uint(v59)%32))))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	switch v67 {
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
	v171 = v51 + int32(1)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v171 < v172 {
		v51 = v171
		goto L7
	} else {
		goto L39
	}
L10:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v83))+2))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+8))
	F_MemoryContextSetParent(m, v157, v153)
	mBase = m.M
	goto L37
L11:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_exec_function_5))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L34
	}
L12:
	;
	v106 = v45 + v51<<(uint(int32(4))%32)
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+8)))
	if v107 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L13:
	;
	v72 = v45 + v51<<(uint(int32(4))%32)
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+8)))
	F_assign_simple_var(m, v13+int32(24), v66, v73, v74, int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+48)))
	if v78 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v66)+24))
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+12)))
	if v80 != int32(_a_F_plpgsql_exec_function_6) {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	v83 = *(*int64)(unsafe.Add(mBase, uint32(v66)+40))
	v84 = base.I32_wrap_i64(v83)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v85 != int32(1) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+20)))
	if v91 != int32(1) {
		goto L9
	} else {
		goto L19
	}
L18:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	switch v88 - int32(2) {
	case 0:
		goto L9
	case 1:
		goto L10
	default:
		goto L17
	}
L19:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
	v98 = F_expand_array(m, v83, v96, int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_assign_simple_var(m, v13+int32(24), v66, v98, int32(0), int32(1))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	goto L9
L22:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v13)+136))
	if v121 != 0 {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v106)))
	F_exec_move_row_from_datum(m, v13+int32(24), v66, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v117 = int32(0)
	F_exec_move_row(m, v13+int32(24), v66, v117, v117)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
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
	F_SPI_freetuptable(m, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v124 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+136)) = v124
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+152))
	if v126 == v124 {
		goto L9
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v126)+20))
	F_MemoryContextReset(m, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L9
L34:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+508))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v136+v51<<(uint(int32(2))%32))))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v141
	F_errmsg_internal(m, int32(_a_F_plpgsql_exec_function_7), v13)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_plpgsql_exec_function_8), int32(615), int32(_a_F_plpgsql_exec_function_9))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
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
	F_assign_simple_var(m, v13+int32(24), v66, base.I64_extend_i32_u(v156+int32(12)), int32(0), int32(1))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
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
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_function[1]))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	if v201 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_function[2]))
	if v240 != 0 {
		goto L53
	} else {
		goto L54
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+164)) = int32(0)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+156)) = v206
	v237 = v206
	goto L41
L43:
	;
	goto L44
L44:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v201)+4))
	if v208 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+12))
	if v230 == int32(0) {
		v237 = v228
		goto L41
	} else {
		goto L51
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+164)) = int32(0)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+156)) = v213
	v228 = v213
	v229 = v201
	goto L45
L47:
	;
	goto L48
L48:
	;
	m.T0[v208].(func(*base.Module, int32, int32))(m, v13+int32(24), l0)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v220 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_function[1]))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	v222 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+164)) = v222
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+156)) = v224
	if v221 == v222 {
		v237 = v224
		goto L41
	} else {
		goto L50
	}
L50:
	;
	v228 = v224
	v229 = v221
	goto L45
L51:
	;
	m.T0[v230].(func(*base.Module, int32, int32))(m, v13+int32(24), v228)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v237 = v228
	goto L41
L53:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v244 = v13 + int32(24)
	v245 = F_exec_stmt_block(m, v244, v237)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	v248 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_function[1]))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
	if v249 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+156)) = int32(0)
	if v245 == int32(2) {
		goto L64
	} else {
		goto L65
	}
L59:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v249)+16))
	if v252 == int32(0) {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	m.T0[v252].(func(*base.Module, int32, int32))(m, v244, v237)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
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
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L125
	}
L63:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_exec_function_5))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L1
	} else {
		goto L121
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+164)) = int32(_a_F_plpgsql_exec_function_10)
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+48)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)) = uint8(v264)
	v267 = l1 + int32(16)
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+61)))
	if v268 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+164)) = int32(0)
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_exec_function_5))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L117
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+164)) = int32(_a_F_plpgsql_exec_function_11)
	v400 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_function[1]))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v400)))
	if v401 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L68:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v13)+88))
	if v271 == int32(0) {
		goto L63
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	if v264&int32(1) == int32(0) {
		goto L78
	} else {
		goto L79
	}
L71:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	if v274 != int32(389) {
		goto L63
	} else {
		goto L72
	}
L72:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+12)))
	if v277&int32(2) == int32(0) {
		goto L62
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v271)+16)) = int32(2)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	if v284 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v271)+24)) = v284
	v286 = int32(_a_F_plpgsql_exec_function_12)
	v287 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_function[3]))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v13)+80))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_function[3])) = v289
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
	v292 = F_CreateTupleDescCopy(m, v291)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = int64(0)
	v300 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v267))) = uint8(v300)
	goto L67
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v271)+28)) = v292
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_function[3])) = v287
	goto L76
L78:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)))
	if v307 == int32(1) {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+62)))
	if v380 != int32(1) {
		goto L67
	} else {
		goto L104
	}
L81:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if base.B2i32(v310 == int32(2249))|base.B2i32(v310 != v306) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	goto L83
L83:
	;
	v367 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
	v368 = int32(-1)
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v371 = F_exec_cast_value(m, v13+int32(24), v367, v267, v306, v368, v369, v368)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L100
	}
L84:
	;
	v317 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
	v319 = F_SPI_datumTransfer(m, v317, int32(-1))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v326 = F_get_call_result_type(m, l1, v13+int32(8), v13+int32(4))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L92
	}
L87:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v319
	goto L67
L88:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_exec_function_5))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L97
	}
L89:
	;
	v347 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
	v349 = F_SPI_datumTransfer(m, v347, int32(-1))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L96
	}
L90:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	F_coerce_function_result_tuple(m, v13+int32(24), v337)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L94
	}
L91:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	F_coerce_function_result_tuple(m, v13+int32(24), v332)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L93
	}
L92:
	;
	switch v326 - int32(1) {
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
	v340 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
	v341 = int32(0)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	F_domain_check(m, v340, v341, v342, v341, v341)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	goto L67
L96:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v349
	goto L67
L97:
	;
	F_errmsg_internal(m, int32(_a_F_plpgsql_exec_function_13), int32(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_plpgsql_exec_function_8), int32(747), int32(_a_F_plpgsql_exec_function_9))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
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
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v371
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267))))
	if v374 != 0 {
		goto L67
	} else {
		goto L101
	}
L101:
	;
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+60)))
	if v375 != 0 {
		goto L67
	} else {
		goto L102
	}
L102:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v377 = F_SPI_datumTransfer(m, v371, v376)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v377
	goto L67
L104:
	;
	v385 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	v387 = int32(-1)
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v390 = F_exec_cast_value(m, v13+int32(24), v385, v267, v386, v387, v388, v387)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v390
	goto L67
L106:
	;
	v413 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_function[4]))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v413)+8))
	F_pfree(m, v413)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L110
	}
L107:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v401)+8))
	if v404 == int32(0) {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	m.T0[v404].(func(*base.Module, int32, int32))(m, v13+int32(24), l0)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	goto L106
L110:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_function[4])) = v414
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v13)+152))
	F_FreeExprContext(m, v419, int32(1))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v423 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+152)) = v423
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v13)+136))
	if v425 == v423 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_exec_function[0])) = v440
	v442 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
	m.G0 = v13 + int32(176)
	return v442
L113:
	;
	F_SPI_freetuptable(m, v425)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v430 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+136)) = v430
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v13)+152))
	if v432 == v430 {
		goto L112
	} else {
		goto L115
	}
L115:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v432)+20))
	F_MemoryContextReset(m, v435)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
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
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	F_errmsg(m, int32(_a_F_plpgsql_exec_function_14), int32(0))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_plpgsql_exec_function_8), int32(642), int32(_a_F_plpgsql_exec_function_9))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
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
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	F_errmsg(m, int32(_a_F_plpgsql_exec_function_15), int32(0))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_plpgsql_exec_function_8), int32(660), int32(_a_F_plpgsql_exec_function_9))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
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
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	F_errmsg(m, int32(_a_F_plpgsql_exec_function_16), int32(0))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_plpgsql_exec_function_8), int32(665), int32(_a_F_plpgsql_exec_function_9))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
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
func F_plpgsql_inline_handler(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v61 int64
	_ = v61
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v140 int64
	_ = v140
	var v141 int32
	_ = v141
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v172 int32
	_ = v172
	var v173 int64
	_ = v173
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v227 int32
	_ = v227
	var v239 int32
	_ = v239
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v274 int32
	_ = v274
	var v283 int32
	_ = v283
	var v292 int32
	_ = v292
	var v293 int64
	_ = v293
	var v305 int32
	_ = v305
	var v314 int32
	_ = v314
	var v324 int32
	_ = v324
	var v325 int64
	_ = v325
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(272)
	m.G0 = v15
	v20 = v2
	v21 = v2
	v22 = v2
	v23 = v2
	v24 = v2
	v25 = v2
	v26 = v2
	v27 = int32(-1)
	v28 = v2
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
	m.G0 = v15 + int32(272)
	return v140
L4:
	;
	goto L3
L5:
	;
	if v27 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v324 = int32(m.ExcTag)
	v325 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v324 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+13)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v15)+252)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v15)+260)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v15)+264)) = v21
	v41 = v32 + int32(13)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+268)) = v41
	F_SPI_connect_ext(m, v33^int32(1))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	v115 = v20
	v116 = v21
	v117 = v22
	v118 = v23
	v119 = v24
	v120 = v25
	v121 = v26
	v123 = v28
	goto L10
L10:
	;
	if v123 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v15)+252)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v15)+260)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v15)+264)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v15)+268)) = v41
	v55 = F_plpgsql_compile_inline(m, v47)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v55)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v55)+24)) = v57 + int64(1)
	v61 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+200)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v15)+216)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v15)+232)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v15)+224)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v15)+184)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v15)+192)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v15)+208)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v15)+252)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = v22
	v80 = v55 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+260)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v15)+264)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v15)+268)) = v41
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_inline_handler[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+204)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v15)+216)) = v15 + int32(184)
	v90 = F_CreateExecutorState(m)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v15)+252)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v15)+260)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v15)+264)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v15)+268)) = v41
	v101 = F_ResourceOwnerCreate(m, int32(0), int32(_a_F_plpgsql_inline_handler_0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_inline_handler[1]))
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_inline_handler[2]))
	goto L15
L15:
	;
	v108 = v15 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v108)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v108))) = v15 + int32(12)
	goto L18
L16:
	;
	v115 = v101
	v116 = v55
	v117 = v90
	v118 = v80
	v119 = v106
	v120 = v104
	v121 = v41
	v123 = int32(0)
	goto L10
L18:
	;
	goto L16
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_inline_handler[2])) = v15 + int32(16)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v15)+252)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v15)+260)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v15)+264)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v15)+268)) = v121
	v140 = F_plpgsql_exec_function(m, v116, v15+int32(216), v117, v115, v115, v130)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L7
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_inline_handler[1])) = v120
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_inline_handler[2])) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v15)+252)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v15)+260)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v15)+264)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v15)+268)) = v121
	v252 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_inline_handler[3]))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+8))
	goto L33
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_inline_handler[1])) = v120
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_inline_handler[2])) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v15)+252)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v15)+260)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v15)+264)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v15)+268)) = v121
	F_FreeExecutorState(m, v117)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v15)+252)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v15)+260)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v15)+264)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v15)+268)) = v121
	F_ReleaseAllPlanCacheRefsInOwner(m, v115)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v15)+252)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v15)+260)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v15)+264)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v15)+268)) = v121
	F_ResourceOwnerDelete(m, v115)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v173 = *(*int64)(unsafe.Add(mBase, uint32(v118)))
	*(*int64)(unsafe.Add(mBase, uint32(v118))) = v173 - int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v15)+252)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v15)+260)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v15)+264)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v15)+268)) = v121
	F_plpgsql_free_function_memory(m, v116)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v15)+252)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v15)+260)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v15)+264)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v15)+268)) = v121
	v193 = F_SPI_finish(m)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	if v193 == int32(2) {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v15)+252)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v15)+260)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v15)+264)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v15)+268)) = v121
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_inline_handler_1))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v15)+252)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v15)+260)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v15)+264)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v15)+268)) = v121
	v215 = F_SPI_result_code_string(m, v193)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v15)+252)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v15)+260)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v15)+264)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v15)+268)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v215
	F_errmsg_internal(m, int32(_a_F_plpgsql_inline_handler_2), v15)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v15)+252)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v15)+260)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v15)+264)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v15)+268)) = v121
	F_errfinish(m, int32(_a_F_plpgsql_inline_handler_3), int32(427), int32(_a_F_plpgsql_inline_handler_4))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	goto L1
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v15)+252)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v15)+260)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v15)+264)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v15)+268)) = v121
	v262 = int32(0)
	F_plpgsql_subxact_cb(m, int32(2), v253, v262, v262)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v15)+252)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v15)+260)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v15)+264)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v15)+268)) = v121
	F_FreeExecutorState(m, v117)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v15)+252)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v15)+260)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v15)+264)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v15)+268)) = v121
	F_ReleaseAllPlanCacheRefsInOwner(m, v115)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v15)+252)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v15)+260)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v15)+264)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v15)+268)) = v121
	F_ResourceOwnerDelete(m, v115)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	v293 = *(*int64)(unsafe.Add(mBase, uint32(v118)))
	*(*int64)(unsafe.Add(mBase, uint32(v118))) = v293 - int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v15)+252)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v15)+260)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v15)+264)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v15)+268)) = v121
	F_plpgsql_free_function_memory(m, v116)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+248)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v15)+252)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v15)+260)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v15)+264)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v15)+268)) = v121
	F_pg_re_throw(m)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	goto L1
L40:
	;
	v329 = int32(v325)
	m.G0 = v15
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v329)+4))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v329)))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	if v15+int32(12) == v335 {
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
	if v339 != 0 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v332)+4))
	v339 = v337
	goto L45
L44:
	;
	v339 = int32(0)
	goto L45
L45:
	;
	goto L42
L46:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v15)+268))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v15)+264))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v15)+260))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v15)+256))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v15)+252))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v15)+248))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v15)+244))
	v20 = v344
	v21 = v341
	v22 = v343
	v23 = v342
	v24 = v345
	v25 = v346
	v26 = v340
	v27 = v339
	v28 = v331
	goto L2
L47:
	;
	goto L48
L48:
	;
	F___wasm_longjmp(m, v332, v331)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	return int64(0)
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
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+68))
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
		v34 = *(*int32)(unsafe.Add(mBase, uint32(v16)+60))
		*(*int32)(unsafe.Add(mBase, uint32(v16+v17<<(uint(int32(2))%32))+72)) = l0
		v39 = *(*int32)(unsafe.Add(mBase, uint32(v16)+68))
		v42 = v16 + v39*int32(24)
		v43 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v42)+96)) = v43
		v45 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
		*(*int64)(unsafe.Add(mBase, uint32(v42)+88)) = v45
		*(*int32)(unsafe.Add(mBase, uint32(v42)+108)) = v34
		*(*int32)(unsafe.Add(mBase, uint32(v42)+104)) = v33
		v49 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
		v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+68))
		*(*int32)(unsafe.Add(mBase, uint32(v49)+68)) = v50 + int32(1)
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
func F_plpgsql_validator(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v125 int32
	_ = v125
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int64
	_ = v143
	var v159 int32
	_ = v159
	var v164 int64
	_ = v164
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(176)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = base.I32_wrap_i64(v16)
	v18 = F_CheckFunctionValidatorAccess(m, v15, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_validator_0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L5
	} else {
		goto L73
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_validator_0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L5
	} else {
		goto L68
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_validator_0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L5
	} else {
		goto L64
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_validator_0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L5
	} else {
		goto L61
	}
L5:
	;
	return int64(0)
L6:
	;
	if v18 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v25 = F_SearchSysCache1(m, int32(47), v16&int64(4294967295))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	m.G0 = v12 + int32(176)
	return int64(0)
L10:
	;
	if v25 == int32(0) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+22)))
	v32 = v30 + v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+108))
	v34 = F_get_typtype(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L5
	} else {
		goto L13
	}
L12:
	;
	v74 = F_get_func_arg_info(m, v25, v12+int32(172), v12+int32(168), v12+int32(164))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L5
	} else {
		goto L26
	}
L13:
	;
	if v34 != int32(112) {
		v66 = v2
		v67 = v2
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)+108))
	if v38 <= int32(3830) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v66 = int32(0)
	v67 = v63
	goto L12
L16:
	;
	v63 = int32(0)
	goto L15
L17:
	;
	switch v38 - int32(2249) {
	case 0, 28, 29, 34:
		goto L16
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 31, 32, 33:
		goto L1
	case 30:
		v66 = int32(1)
		v67 = v2
		goto L12
	default:
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if base.Ui32(v38-int32(_a_F_plpgsql_validator_1)) < base.Ui32(int32(4)) {
		goto L16
	} else {
		goto L22
	}
L20:
	;
	if base.B2i32(v38 == int32(2776))|base.B2i32(v38 == int32(3500)) != 0 {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L1
L22:
	;
	switch v38 - int32(3831) {
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
	if base.Ui32(int32(2)) <= base.Ui32(v38-int32(_a_F_plpgsql_validator_2)) {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	v63 = int32(1)
	goto L15
L25:
	;
	goto L16
L26:
	;
	if int32(0) < v74 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v78 = int32(0)
	goto L30
L28:
	;
	goto L29
L29:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_validator[0])))
	if v137 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L30:
	;
	v88 = v78 << (uint(int32(2)) % 32)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v12)+172))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v88+v89)))
	v92 = F_get_typtype(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L5
	} else {
		goto L33
	}
L31:
	;
	goto L29
L32:
	;
	v125 = v78 + int32(1)
	if v125 != v74 {
		v78 = v125
		goto L30
	} else {
		goto L46
	}
L33:
	;
	if v92 != int32(112) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v12)+172))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v96+v88)))
	if v98 <= int32(3830) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v98 != int32(2249) {
		goto L2
	} else {
		goto L45
	}
L36:
	;
	if v98 <= int32(2775) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	if base.B2i32(base.Ui32(v98-int32(_a_F_plpgsql_validator_1)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v98-int32(_a_F_plpgsql_validator_2)) < base.Ui32(int32(2)))|base.B2i32(v98 == int32(3831)) != 0 {
		goto L32
	} else {
		goto L44
	}
L39:
	;
	switch v98 - int32(2277) {
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
	if v98 == int32(2776) {
		goto L32
	} else {
		goto L42
	}
L42:
	;
	if v98 != int32(3500) {
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
	v142 = m.ExcPending
	if v142 != 0 {
		goto L5
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	F_ReleaseCatCache(m, v25)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L5
	} else {
		goto L60
	}
L50:
	;
	v143 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+104)) = v143
	*(*int64)(unsafe.Add(mBase, uint32(v12)+120)) = v143
	*(*int64)(unsafe.Add(mBase, uint32(v12)+136)) = v143
	*(*int64)(unsafe.Add(mBase, uint32(v12)+152)) = v143
	*(*int64)(unsafe.Add(mBase, uint32(v12)+144)) = v143
	*(*int64)(unsafe.Add(mBase, uint32(v12)+112)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+108)) = v17
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_validator[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+124)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v12)+136)) = v12 + int32(104)
	if v66 != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v190 = F_plpgsql_compile(m, v12+int32(136), int32(1))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L5
	} else {
		goto L57
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+140)) = v12 + int32(60)
	goto L51
L53:
	;
	v164 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+96)) = v164
	*(*int64)(unsafe.Add(mBase, uint32(v12)+88)) = v164
	*(*int64)(unsafe.Add(mBase, uint32(v12)+80)) = v164
	*(*int64)(unsafe.Add(mBase, uint32(v12)+72)) = v164
	*(*int64)(unsafe.Add(mBase, uint32(v12)+64)) = v164
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = int32(448)
	goto L52
L54:
	;
	goto L55
L55:
	;
	if v67 == int32(0) {
		goto L51
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+64)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = int32(447)
	goto L52
L57:
	;
	v192 = F_SPI_finish(m)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L5
	} else {
		goto L58
	}
L58:
	;
	if v192 != int32(2) {
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
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v17
	F_errmsg_internal(m, int32(_a_F_plpgsql_validator_3), v12)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_plpgsql_validator_4), int32(462), int32(_a_F_plpgsql_validator_5))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
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
	v230 = F_SPI_result_code_string(m, v192)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v230
	F_errmsg_internal(m, int32(_a_F_plpgsql_validator_6), v12+int32(48))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L5
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_plpgsql_validator_4), int32(544), int32(_a_F_plpgsql_validator_5))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
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
	v249 = m.ExcPending
	if v249 != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v12)+172))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v250+v78<<(uint(int32(2))%32))))
	v255 = F_format_type_be(m, v254)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L5
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v255
	F_errmsg(m, int32(_a_F_plpgsql_validator_7), v12+int32(32))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L5
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_plpgsql_validator_4), int32(497), int32(_a_F_plpgsql_validator_5))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
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
	v275 = m.ExcPending
	if v275 != 0 {
		goto L5
	} else {
		goto L74
	}
L74:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v32)+108))
	v277 = F_format_type_be(m, v276)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L5
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v277
	F_errmsg(m, int32(_a_F_plpgsql_validator_8), v12+int32(16))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L5
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_plpgsql_validator_4), int32(481), int32(_a_F_plpgsql_validator_5))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
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
