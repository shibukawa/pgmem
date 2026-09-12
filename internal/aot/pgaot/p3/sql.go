package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_check_sql_fn_retval(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	if l0 == int32(0) {
		v9 = F_check_sql_stmt_retval(m, int32(0), l1, l2, l3, l4)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return v9
		}
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v14+v15<<(uint(int32(2))%32)-int32(4))))
		v22 = F_check_sql_stmt_retval(m, v21, l1, l2, l3, l4)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			return v22
		}
	}
}
func F_map_sql_identifier_to_xml_name(m *base.Module) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	F_errstart_cold(m, int32(21), int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			F_errmsg(m, int32(345830), int32(0))
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				F_errdetail(m, int32(538925), int32(0))
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(474287), int32(2425), int32(360442))
					v22 = m.ExcPending
					if v22 != 0 {
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
func F_sql_exec_error_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = F_geterrposition(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		if v9 <= int32(0) {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			F_set_errcontext_domain(m, int32(0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+32))
				if int32(0) < v26 {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v34
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v31
					F_errcontext_msg(m, int32(446639), v7)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						m.G0 = v7 + int32(32)
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v31
					F_errcontext_msg(m, int32(220583), v7+int32(16))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						m.G0 = v7 + int32(32)
						return
					}
				}
			}
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
			if v14 == int32(0) {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				F_set_errcontext_domain(m, int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+32))
					if int32(0) < v26 {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v34
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v31
						F_errcontext_msg(m, int32(446639), v7)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							m.G0 = v7 + int32(32)
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v31
						F_errcontext_msg(m, int32(220583), v7+int32(16))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							m.G0 = v7 + int32(32)
							return
						}
					}
				}
			} else {
				v18 = F_errposition(m, int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					F_internalerrposition(m, v9)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
						v24 = F_internalerrquery(m, v23)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
							F_set_errcontext_domain(m, int32(0))
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return
							} else {
								v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+32))
								if int32(0) < v26 {
									v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
									*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v34
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = v31
									F_errcontext_msg(m, int32(446639), v7)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return
									} else {
										m.G0 = v7 + int32(32)
										return
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v31
									F_errcontext_msg(m, int32(220583), v7+int32(16))
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return
									} else {
										m.G0 = v7 + int32(32)
										return
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
func F_sql_fn_param_ref(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
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
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v6 <= v3 {
		v43 = v3
		return v43
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		if v10 < v6 {
			v43 = v3
			return v43
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v14 = F_palloc0(m, int32(28))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v6
				*(*int64)(unsafe.Add(mBase, uint32(v14))) = int64(8)
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+v6<<(uint(int32(2))%32)-int32(4))))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v27
				v31 = F_get_typcollation(m, v27)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v12
					*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v31
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
					if v35 == int32(0) {
						v43 = v14
					} else {
						if v31 == int32(0) {
							v43 = v14
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v35
							v43 = v14
						}
					}
					return v43
				}
			}
		}
	}
}
func F_sql_fn_parser_setup(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(678)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = int32(679)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	return
}
func F_sql_inline_error_callback(m *base.Module, l0 int32) {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = F_geterrposition(m)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		if int32(0) < v8 {
			v13 = F_errposition(m, int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				F_internalerrposition(m, v8)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v18 = F_internalerrquery(m, v17)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						F_set_errcontext_domain(m, int32(0))
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return
						} else {
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = v23
							F_errcontext_msg(m, int32(318639), v6)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								m.G0 = v6 + int32(16)
								return
							}
						}
					}
				}
			}
		} else {
			F_set_errcontext_domain(m, int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v23
				F_errcontext_msg(m, int32(318639), v6)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					m.G0 = v6 + int32(16)
					return
				}
			}
		}
	}
}
