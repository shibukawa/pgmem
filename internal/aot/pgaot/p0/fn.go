package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_fn_expr_argtype(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	v3 = int32(0)
	if l0 == v3 {
		v52 = v3
		return v52
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v7 == int32(0) {
			v52 = v3
			return v52
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			v12 = v10 - int32(11)
			v19 = int32(0)
			if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v12))|base.B2i32(int32(base.Ui32(int32(977))>>(uint(v12)%32))&int32(1) == v19)|base.B2i32(l1 < v19) != 0 {
				v52 = v3
				return v52
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v12<<(uint(int32(2))%32))+uint32(_c_F_get_fn_expr_argtype[0])))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v7+v27)))
				if v29 == int32(0) {
					v52 = v3
					return v52
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
					if v32 <= l1 {
						v52 = v3
						return v52
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+l1<<(uint(int32(2))%32))))
						v39 = F_exprType(m, v38)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							if l1 != int32(1) {
								v52 = v39
								return v52
							} else {
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
								if v45 != int32(20) {
									v52 = v39
									return v52
								} else {
									v48 = F_get_base_element_type(m, v39)
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int32(0)
									} else {
										v52 = v48
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
func F_get_fn_opclass_options(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	v2 = int32(0)
	if l0 == v2 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_get_fn_opclass_options_0), int32(0))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_get_fn_opclass_options_1), int32(2075), int32(_a_F_get_fn_opclass_options_2))
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
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v5 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_get_fn_opclass_options_0), int32(0))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_get_fn_opclass_options_1), int32(2075), int32(_a_F_get_fn_opclass_options_2))
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
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			if v8 != int32(7) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_get_fn_opclass_options_0), int32(0))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_get_fn_opclass_options_1), int32(2075), int32(_a_F_get_fn_opclass_options_2))
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
			} else {
				v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
				if v11 != int32(17) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_get_fn_opclass_options_0), int32(0))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_get_fn_opclass_options_1), int32(2075), int32(_a_F_get_fn_opclass_options_2))
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
				} else {
					v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+32)))
					if v14 != 0 {
						v25 = v2
						return v25
					} else {
						v15 = *(*int32)(unsafe.Add(mBase, uint32(v5)+24))
						v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
						if v16&int32(3) == int32(0) {
							v25 = v15
							return v25
						} else {
							v21 = F_detoast_attr(m, v15)
							mBase = m.M
							v24 = m.ExcPending
							if v24 != 0 {
								return int32(0)
							} else {
								v25 = v21
								return v25
							}
						}
					}
				}
			}
		}
	}
}
func Fn14206(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v18 int64
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v18 = *(*int64)(unsafe.Add(mBase, uint32(l0-int32(8))))
		*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v18
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
		F_errmsg_internal(m, l3, v8)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			F_errfinish(m, int32(_a_Fn14206_0), l2, l1)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func Fn14213(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = F_SearchSysCache1(m, l5, base.I64_extend_i32_u(l0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
				F_errmsg_internal(m, l4, v11)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, l3, l2, l1)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v29+v30)+4))
			F_ReleaseCatCache(m, v14)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				m.G0 = v11 + int32(16)
				return v32
			}
		}
	}
}
func Fn14215(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	v2 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v2
	*(*int32)(unsafe.Add(mBase, _c_Fn14215[0])) = v2
	v8 = *(*int32)(unsafe.Add(mBase, _c_Fn14215[1]))
	v9 = int32(0)
	v12 = base.AtomicRmwOr32(m, v9, int32(_a_Fn14215_0), v9)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(1)
	v16 = int32(0)
	v19 = base.AtomicRmwOr32(m, v16, int32(_a_Fn14215_0), v16)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v20 == v16 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v23 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_Fn14215[2]))
	if v27 == v23 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v29 = m.G0
	v31 = v29 - int32(16)
	m.G0 = v31
	v34 = *(*int32)(unsafe.Add(mBase, _c_Fn14215[3]))
	if v34 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v57 = F_pgmem_kill(m, v23, int32(23))
	mBase = m.M
	goto L2
L9:
	;
	m.G0 = v31 + int32(16)
	goto L1
L10:
	;
	v37 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+15)) = uint8(v37)
	goto L11
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_Fn14215[4]))
	v45 = F_write(m, v41, v31+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v45 {
		goto L9
	} else {
		goto L13
	}
L12:
	;
	goto L9
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_Fn14215[5]))
	if v49 == int32(27) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
}
func Fn14226(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v12 = F_pg_detoast_datum(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v15 = F_array_iterator(m, v7, l1, v12, int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int64(0)
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v17 != v7 {
					F_pfree(m, v7)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int64(0)
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v21 != v12 {
							F_pfree(m, v12)
							mBase = m.M
							v24 = m.ExcPending
							if v24 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(v15)
							}
						} else {
							return base.I64_extend_i32_u(v15)
						}
					}
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v21 != v12 {
						F_pfree(m, v12)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(v15)
						}
					} else {
						return base.I64_extend_i32_u(v15)
					}
				}
			}
		}
	}
}
func Fn14233(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v50 int32
	_ = v50
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+16))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v18 == int32(0) {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v22 = base.I32_extend16_s(l1)
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+216))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+204))
		v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+6)))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v24+v26*(v22-int32(1))<<(uint(int32(2))%32)+int32(44)-int32(4))))
		if v38 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(117833860))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_Fn14233_0), int32(0))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(11)
						F_errdetail_internal(m, int32(_a_Fn14233_1), v11)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, l4, l3, l2)
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
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
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v43 = F_index_getprocinfo(m, v41, v22, int32(11))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v48 = *(*int64)(unsafe.Add(mBase, uint32(v43)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v48
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
				*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v50
				v52 = *(*int64)(unsafe.Add(mBase, uint32(v43)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v52
				v54 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
				*(*int64)(unsafe.Add(mBase, uint32(v17))) = v54
				*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v47
				*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = int32(0)
				m.G0 = v11 + int32(16)
				return v17
			}
		}
	} else {
		m.G0 = v11 + int32(16)
		return v17
	}
}
func Fn14240(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
		if v11 == int32(1) {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
			if v17 == int32(18) {
				v20 = int32(16)
			} else {
				v20 = int32(0)
			}
			if base.Ui32((v17-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v27 = int32(4)
			} else {
				v27 = v20
			}
			v40 = v27
		} else {
			v28 = int32(1)
			if v11&v28 != 0 {
				v40 = int32(base.Ui32(v11)>>(uint(v28)%32)) - v28
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v40 = int32(base.Ui32(v34)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v41 = int32(1)
		if v11&v41 != 0 {
			v45 = v41
		} else {
			v45 = int32(4)
		}
		v49 = F_dotrim(m, v7+v45, v40, int32(_a_Fn14240_0), int32(1), l2, l1)
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v49)
		}
	}
}
func Fn14246(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v36 int64
	_ = v36
	var v43 int64
	_ = v43
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = int64(63)
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v22 = int64(32)
	v23 = int64(base.Ui64(v15) >> (uint(v22) % 64))
	v25 = int64(base.Ui64(v12) >> (uint(v22) % 64))
	v28 = int64(4294967295)
	v29 = v15 & v28
	v31 = v12 & v28
	v32 = v29 * v31
	v36 = int64(base.Ui64(v32)>>(uint(v22)%64)) + v29*v25
	v43 = v31*v23 + v36&v28
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v12*(v15>>(uint(v13)%64)) + v12>>(uint(v13)%64)*v15 + v23*v25 + int64(base.Ui64(v36)>>(uint(v22)%64)) + int64(base.Ui64(v43)>>(uint(v22)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v32&v28 | v43<<(uint(v22)%64)
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	if v54 != v55>>(uint(int64(63))%64) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v64 = m.ExcPending
		if v64 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, l4, int32(0))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, l3, l2, l1)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		m.G0 = v10 + int32(16)
		return v55
	}
}
func Fn14251(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	if l0 == int32(0) {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v9 == l3 {
			return int32(1)
		} else {
			v13 = F_expression_tree_walker_impl(m, l0, l2, l1)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				return v13
			}
		}
	}
}
func Fn14257(m *base.Module, l0 int32, l1 int32) int64 {
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = F_string2ean(m, v9, v10, v7+int32(8), l1)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		if v13 == int32(0) {
			v19 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
			v23 = int64(0)
		} else {
			v22 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
			v23 = v22
		}
		m.G0 = v7 + int32(16)
		return v23
	}
}
func Fn14262(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v77 int64
	_ = v77
	v6 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v6)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+20)))
	if v17&int32(1) == v6 {
		v26 = l2 + l1<<(uint(int32(3))%32) + int32(20)
		v27 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26))))
		if v27 < int32(0) {
			v71 = F_nocachegetattr(m, l0, l1, l2)
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return int64(0)
			} else {
				v77 = v71
				m.G0 = v12 + int32(16)
				return v77
			}
		} else {
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+22)))
			v32 = v16 + v30 + v27
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+4)))
			if v33 == int32(1) {
				v36 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+2)))
				if base.I32_popcnt(v36) != int32(1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12))) = v36
						F_errmsg_internal(m, int32(_a_Fn14262_0), v12)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, l4, int32(123), int32(_a_Fn14262_1))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					switch base.I32_ctz(v36) {
					case 0:
						v41 = int64(*(*int8)(unsafe.Add(mBase, uint32(v32))))
						v77 = v41
						m.G0 = v12 + int32(16)
						return v77
					case 1:
						v42 = int64(*(*int16)(unsafe.Add(mBase, uint32(v32))))
						v77 = v42
						m.G0 = v12 + int32(16)
						return v77
					case 2:
						v43 = int64(*(*int32)(unsafe.Add(mBase, uint32(v32))))
						v77 = v43
						m.G0 = v12 + int32(16)
						return v77
					case 3:
						v44 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
						v77 = v44
						m.G0 = v12 + int32(16)
						return v77
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12))) = v36
							F_errmsg_internal(m, int32(_a_Fn14262_0), v12)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, l4, int32(123), int32(_a_Fn14262_1))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int64(0)
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
				v77 = base.I64_extend_i32_u(v32)
				m.G0 = v12 + int32(16)
				return v77
			}
		}
	} else {
		v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+23)))
		v61 = int32(1)
		if int32(base.Ui32(v60)>>(uint(l1-v61)%32))&v61 != 0 {
			v71 = F_nocachegetattr(m, l0, l1, l2)
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return int64(0)
			} else {
				v77 = v71
				m.G0 = v12 + int32(16)
				return v77
			}
		} else {
			v66 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v66)
			v77 = int64(0)
			m.G0 = v12 + int32(16)
			return v77
		}
	}
}
func Fn14277(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 float64
	_ = v28
	var v31 int32
	_ = v31
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v15 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v14 + v15
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v14
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+16)))
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21+v22)+12)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = F_gbt_num_distance(m, v9, v9+v15, v24&int32(1), l1, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		return int64(0)
	} else {
		m.G0 = v9 + int32(16)
		return base.I64_reinterpret_f64(v28)
	}
}
func Fn14282(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_palloc0(m, int32(16))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(16)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v15 = F_gbt_num_union(m, v7, v5, l1, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v15)
		}
	}
}
func Fn14284(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = base.I32_wrap_i64(l0)
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = base.I32_wrap_i64(l1)
		v20 = F_pg_detoast_datum(m, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v23 = v12 + int32(8)
			v25 = int32(4)
			v26 = v15 + v25
			*(*int32)(unsafe.Add(mBase, uint32(v23))) = v26
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
			v29 = int32(2)
			v30 = int32(base.Ui32(v28) >> (uint(v29) % 32))
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			if base.Ui32(v30+v25) < base.Ui32(int32(base.Ui32(v38)>>(uint(v29)%32))) {
				v42 = v26 + (v30+int32(3))&int32(2147483644)
			} else {
				v42 = v26
			}
			*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v42
			v45 = int32(4)
			v46 = v20 + v45
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v46
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
			v49 = int32(2)
			v50 = int32(base.Ui32(v48) >> (uint(v49) % 32))
			v58 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
			if base.Ui32(v50+v45) < base.Ui32(int32(base.Ui32(v58)>>(uint(v49)%32))) {
				v62 = v46 + (v50+int32(3))&int32(2147483644)
			} else {
				v62 = v46
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v62
			v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			v65 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v12)+8)))
			v66 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v12))))
			v67 = F_DirectFunctionCall2Coll(m, l3, v64, v65, v66)
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int32(0)
			} else {
				if v15 != v14 {
					F_pfree(m, v15)
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						if v20 != v19 {
							F_pfree(m, v20)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								m.G0 = v12 + int32(16)
								return base.I32_wrap_i64(v67)
							}
						} else {
							m.G0 = v12 + int32(16)
							return base.I32_wrap_i64(v67)
						}
					}
				} else {
					if v20 != v19 {
						F_pfree(m, v20)
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							m.G0 = v12 + int32(16)
							return base.I32_wrap_i64(v67)
						}
					} else {
						m.G0 = v12 + int32(16)
						return base.I32_wrap_i64(v67)
					}
				}
			}
		}
	}
}
func Fn14297(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	if l1 < int32(0) {
		v16 = *(*int32)(unsafe.Add(mBase, _c_Fn14297[0]))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v16+(l1^int32(-1))<<(uint(int32(2))%32))))
		v30 = v22
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, _c_Fn14297[1]))
		v30 = v24 + l1<<(uint(int32(13))%32) + int32(-8192)
	}
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+14)))
	if v31 != 0 {
		v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+19)))
		v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+16)))
		if (v32<<(uint(int32(8))%32)-v35)&int32(_a_Fn14297_0) != int32(16) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v87 = m.ExcPending
			if v87 != 0 {
				return
			} else {
				F_errcode(m, int32(33557032))
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return
				} else {
					v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					if l1 < int32(0) {
						v95 = *(*int32)(unsafe.Add(mBase, _c_Fn14297[2]))
						v101 = *(*int32)(unsafe.Add(mBase, uint32(v95+(l1^int32(-1))*int32(56))+16))
						v110 = v101
					} else {
						v103 = *(*int32)(unsafe.Add(mBase, _c_Fn14297[3]))
						v104 = int32(56)
						v109 = *(*int32)(unsafe.Add(mBase, uint32(v103+l1*v104-v104)+16))
						v110 = v109
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v110
					*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v91 + int32(4)
					F_errmsg(m, int32(_a_Fn14297_1), v11+int32(16))
					mBase = m.M
					v119 = m.ExcPending
					if v119 != 0 {
						return
					} else {
						F_errhint(m, int32(_a_Fn14297_2), int32(0))
						mBase = m.M
						v123 = m.ExcPending
						if v123 != 0 {
							return
						} else {
							F_errfinish(m, l4, l3, l2)
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
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
			m.G0 = v11 + int32(32)
			return
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v47 = m.ExcPending
		if v47 != 0 {
			return
		} else {
			F_errcode(m, int32(33557032))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				if l1 < int32(0) {
					v55 = *(*int32)(unsafe.Add(mBase, _c_Fn14297[2]))
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v55+(l1^int32(-1))*int32(56))+16))
					v70 = v61
				} else {
					v63 = *(*int32)(unsafe.Add(mBase, _c_Fn14297[3]))
					v64 = int32(56)
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v63+l1*v64-v64)+16))
					v70 = v69
				}
				*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v70
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v51 + int32(4)
				F_errmsg(m, int32(_a_Fn14297_3), v11)
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return
				} else {
					F_errhint(m, int32(_a_Fn14297_2), int32(0))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return
					} else {
						F_errfinish(m, l4, l5, l2)
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
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
func Fn14301(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v23 = F_ArrayGetIntegerTypmods(m, v17, v14+int32(12))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int64(0)
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
			if v25 == int32(1) {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				if v28 <= int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, l7, int32(0))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, l3, l6, l1)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					if base.Ui32(l9) <= base.Ui32(v28) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v14))) = l5
								F_errmsg(m, l4, v14)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, l3, l2, l1)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						m.G0 = v14 + int32(16)
						return base.I64_extend_i32_u(v28)
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_Fn14301_0), int32(0))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, l3, l8, l1)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int64(0)
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
func Fn14310(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
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
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v14 = int32(0)
	goto L3
L3:
	;
	v15 = F_list_concat_copy(m, l2, l3)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L8
	}
L4:
	;
	return
L5:
	;
	v11 = F_get_opclass_input_type(m, l1)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v14 = v11
	goto L3
L7:
	;
	return
L8:
	;
	if v15 == int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v19 <= int32(0) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v24 = l1
	v25 = int32(0)
	v29 = v14
	goto L11
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v25<<(uint(int32(2))%32))))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v36 != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L7
L13:
	;
	v63 = v25 + int32(1)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v63 < v64 {
		v24 = v59
		v25 = v63
		v29 = v61
		goto L11
	} else {
		goto L27
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = l0
	v57 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+24)) = uint16(v57)
	v59 = v24
	v61 = v29
	goto L13
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	if v37 != int32(1) {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
	if v40 != v41 {
		goto L14
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	if v40 != v29 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v44 = F_opclass_for_family_datatype(m, l4, l0, v40)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	v46 = v24
	v47 = v29
	goto L22
L22:
	;
	if v46 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v46 = v44
	v47 = v40
	goto L22
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = v46
	v49 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+24)) = uint16(v49)
	v59 = v46
	v61 = v47
	goto L13
L25:
	;
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = l0
	v52 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+24)) = uint16(v52)
	v59 = int32(0)
	v61 = v47
	goto L13
L27:
	;
	goto L12
}
func Fn14316(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	F_ean2string(m, v8, v6, l1)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v13 = F_pstrdup(m, v6)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			m.G0 = v6 + int32(32)
			return base.I64_extend_i32_u(v13)
		}
	}
}
func Fn14325(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	if l1 == int32(0) {
		v8 = F_palloc(m, int32(32))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(4)
			*(*int64)(unsafe.Add(mBase, uint32(v8))) = l2
			v16 = v8 + int32(16)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v16
			*(*int32)(unsafe.Add(mBase, uint32(v16))) = l0
			return v8
		}
	} else {
		F_new_head_cell(m, l1)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v22))) = l0
			return l1
		}
	}
}
func Fn14332(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 float64
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 float64
	_ = v56
	var v58 float64
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 float64
	_ = v61
	var v64 int32
	_ = v64
	var v65 float64
	_ = v65
	var v70 float64
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v86 float64
	_ = v86
	var v88 float64
	_ = v88
	var v91 int32
	_ = v91
	var v92 float64
	_ = v92
	var v103 float64
	_ = v103
	var v105 float64
	_ = v105
	var v106 float64
	_ = v106
	var v107 float64
	_ = v107
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v202 float64
	_ = v202
	var v204 float64
	_ = v204
	var v206 float64
	_ = v206
	var v219 float64
	_ = v219
	var v229 float64
	_ = v229
	var v240 float64
	_ = v240
	var v252 float64
	_ = v252
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v267 int64
	_ = v267
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int64
	_ = v280
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int64
	_ = v298
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int64
	_ = v324
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	v11 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(96)
	m.G0 = v18
	if l0 < int32(_a_Fn14332_0) {
		v336 = v11
		m.G0 = v18 + int32(96)
		return v336
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v23 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
		if v23 < int64(10000) {
			v336 = v11
			m.G0 = v18 + int32(96)
			return v336
		} else {
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)))
			if v26 != int32(1) {
				v336 = v11
				m.G0 = v18 + int32(96)
				return v336
			} else {
				v30 = v22 + int32(16)
				v31 = int32(0)
				v38 = float64(0)
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
				if v40 != 0 {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
					if v40 != int32(1) {
						v50 = v31
						v51 = v31
						v56 = v38
						for {
							v58 = float64(1)
							v59 = v50 + v41
							v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+1)))
							v61 = F_scalbn(m, v58, v60)
							mBase = m.M
							v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
							v65 = F_scalbn(m, v58, v64)
							mBase = m.M
							v70 = base.F64_add(base.F64_add(v56, base.F64_div(v58, v65)), base.F64_div(v58, v61))
							v71 = int32(2)
							v72 = v50 + v71
							v74 = v51 + v71
							if v74 != v40&int32(-2) {
								v50 = v72
								v51 = v74
								v56 = v70
								continue
							} else {
								break
							}
							break
						}
						if v40&int32(1) == int32(0) {
							v103 = v70
						} else {
							v80 = v72
							v86 = v70
							v88 = float64(1)
							v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+v41))))
							v92 = F_scalbn(m, v88, v91)
							mBase = m.M
							v103 = base.F64_add(v86, base.F64_div(v88, v92))
						}
					} else {
						v80 = v31
						v86 = v38
						v88 = float64(1)
						v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+v41))))
						v92 = F_scalbn(m, v88, v91)
						mBase = m.M
						v103 = base.F64_add(v86, base.F64_div(v88, v92))
					}
					v105 = *(*float64)(unsafe.Add(mBase, uint32(v30)+8))
					v106 = base.F64_div(v105, v103)
					v107 = base.F64_convert_i32_u(v40)
					if base.F64_le(v106, base.F64_mul(v107, float64(2.5))) == int32(0) {
						v219 = v106
						if base.F64_gt(v219, float64(1.4316557653333333e+08)) == int32(0) {
							v240 = v219
						} else {
							v229 = F_log(m, base.F64_add(base.F64_mul(v219, float64(-2.3283064365386963e-10)), float64(1)))
							mBase = m.M
							v240 = base.F64_mul(v229, float64(-4.294967296e+09))
						}
						v252 = v240
					} else {
						v114 = v40 & int32(3)
						v115 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
						v116 = int32(0)
						if base.Ui32(int32(4)) <= base.Ui32(v40) {
							v125 = int32(0)
							v126 = v116
							v127 = v116
							for {
								v134 = v126 + v115
								v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
								v136 = int32(0)
								v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+1)))
								v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+2)))
								v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+3)))
								v150 = v127 + base.B2i32(v135 == v136) + base.B2i32(v139 == v136) + base.B2i32(v143 == v136) + base.B2i32(v147 == v136)
								v151 = int32(4)
								v152 = v126 + v151
								v154 = v125 + v151
								if v154 != v40&int32(-4) {
									v125 = v154
									v126 = v152
									v127 = v150
									continue
								} else {
									break
								}
								break
							}
							if v114 == int32(0) {
								v191 = v150
							} else {
								v160 = v152
								v161 = v150
								v170 = v160
								v171 = v161
								v174 = v116
								for {
									v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170+v115))))
									v182 = v171 + base.B2i32(v179 == int32(0))
									v183 = int32(1)
									v186 = v174 + v183
									if v186 != v114 {
										v170 = v170 + v183
										v171 = v182
										v174 = v186
										continue
									} else {
										break
									}
									break
								}
								v191 = v182
							}
						} else {
							v160 = v116
							v161 = v116
							v170 = v160
							v171 = v161
							v174 = v116
							for {
								v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170+v115))))
								v182 = v171 + base.B2i32(v179 == int32(0))
								v183 = int32(1)
								v186 = v174 + v183
								if v186 != v114 {
									v170 = v170 + v183
									v171 = v182
									v174 = v186
									continue
								} else {
									break
								}
								break
							}
							v191 = v182
						}
						if v191 == int32(0) {
							v240 = v106
							v252 = v240
						} else {
							v202 = F_log(m, base.F64_div(v107, base.F64_convert_i32_s(v191)))
							mBase = m.M
							v252 = base.F64_mul(v202, v107)
						}
					}
				} else {
					v204 = *(*float64)(unsafe.Add(mBase, uint32(v30)+8))
					v206 = base.F64_div(v204, float64(0))
					if base.F64_le(v206, base.F64_mul(base.F64_convert_i32_u(v40), float64(2.5))) != 0 {
						v240 = v206
					} else {
						v219 = v206
						if base.F64_gt(v219, float64(1.4316557653333333e+08)) == int32(0) {
							v240 = v219
						} else {
							v229 = F_log(m, base.F64_add(base.F64_mul(v219, float64(-2.3283064365386963e-10)), float64(1)))
							mBase = m.M
							v240 = base.F64_mul(v229, float64(-4.294967296e+09))
						}
					}
					v252 = v240
				}
				if base.F64_gt(v252, float64(100000)) != 0 {
					v256 = int32(*(*uint8)(unsafe.Add(mBase, _c_Fn14332[0])))
					if v256 != int32(1) {
						v276 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)) = uint8(v276)
						v336 = v11
						m.G0 = v18 + int32(96)
						return v336
					} else {
						v261 = F_errstart(m, int32(15), int32(0))
						mBase = m.M
						v264 = m.ExcPending
						if v264 != 0 {
							return int32(0)
						} else {
							if v261 == int32(0) {
								v276 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)) = uint8(v276)
								v336 = v11
								m.G0 = v18 + int32(96)
								return v336
							} else {
								v267 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
								*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = l0
								*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v267
								*(*float64)(unsafe.Add(mBase, uint32(v18))) = v252
								F_errmsg_internal(m, l9, v18)
								mBase = m.M
								v272 = m.ExcPending
								if v272 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, l4, l8, l2)
									mBase = m.M
									v274 = m.ExcPending
									if v274 != 0 {
										return int32(0)
									} else {
										v276 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)) = uint8(v276)
										v336 = v11
										m.G0 = v18 + int32(96)
										return v336
									}
								}
							}
						}
					}
				} else {
					v279 = int32(*(*uint8)(unsafe.Add(mBase, _c_Fn14332[0])))
					v280 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
					if base.F64_gt(base.F64_add(base.F64_div(base.F64_convert_i64_s(v280), float64(2000)), float64(0.5)), v252) != 0 {
						v287 = int32(1)
						if v279&v287 == int32(0) {
							v336 = v287
							m.G0 = v18 + int32(96)
							return v336
						} else {
							v294 = F_errstart(m, int32(15), int32(0))
							mBase = m.M
							v295 = m.ExcPending
							if v295 != 0 {
								return int32(0)
							} else {
								if v294 == int32(0) {
									v336 = v287
									m.G0 = v18 + int32(96)
									return v336
								} else {
									v298 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
									*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = l0
									*(*int64)(unsafe.Add(mBase, uint32(v18)+48)) = v298
									*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v252
									*(*float64)(unsafe.Add(mBase, uint32(v18)+40)) = base.F64_add(base.F64_div(base.F64_convert_i64_s(v298), float64(2000)), float64(0.5))
									F_errmsg_internal(m, l7, v18+int32(32))
									mBase = m.M
									v311 = m.ExcPending
									if v311 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, l4, l6, l2)
										mBase = m.M
										v313 = m.ExcPending
										if v313 != 0 {
											return int32(0)
										} else {
											v336 = v287
											m.G0 = v18 + int32(96)
											return v336
										}
									}
								}
							}
						}
					} else {
						if v279&int32(1) == int32(0) {
							v336 = v11
							m.G0 = v18 + int32(96)
							return v336
						} else {
							v320 = F_errstart(m, int32(15), int32(0))
							mBase = m.M
							v321 = m.ExcPending
							if v321 != 0 {
								return int32(0)
							} else {
								if v320 == int32(0) {
									v336 = v11
									m.G0 = v18 + int32(96)
									return v336
								} else {
									v324 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
									*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = l0
									*(*int64)(unsafe.Add(mBase, uint32(v18)+72)) = v324
									*(*float64)(unsafe.Add(mBase, uint32(v18)+64)) = v252
									F_errmsg_internal(m, l5, v18-int32(-64))
									mBase = m.M
									v331 = m.ExcPending
									if v331 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, l4, l3, l2)
										mBase = m.M
										v333 = m.ExcPending
										if v333 != 0 {
											return int32(0)
										} else {
											v336 = v11
											m.G0 = v18 + int32(96)
											return v336
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
func Fn14345(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	switch v4 - int32(142) {
	case 0:
		v15 = l1
		return v15
	case 1:
		return int32(3)
	default:
		if int32(0) <= base.I32_extend8_s(v4) {
			v14 = int32(1)
		} else {
			v14 = int32(2)
		}
		v15 = v14
		return v15
	}
}
func Fn14352(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if int32(3) <= v16 {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				v20 = F_pg_detoast_datum_packed(m, v19)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int64(0)
				} else {
					v22 = v20
					v23 = F_encrypt_internal(m, l2, l1, v9, v14, v22)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int64(0)
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						if v25 != v9 {
							F_pfree(m, v9)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int64(0)
							} else {
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								if v29 != v14 {
									F_pfree(m, v14)
									mBase = m.M
									v32 = m.ExcPending
									if v32 != 0 {
										return int64(0)
									} else {
										v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
										if v33 < int32(3) {
											return base.I64_extend_i32_u(v23)
										} else {
											v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
											if v22 == v36 {
												return base.I64_extend_i32_u(v23)
											} else {
												F_pfree(m, v22)
												mBase = m.M
												v39 = m.ExcPending
												if v39 != 0 {
													return int64(0)
												} else {
													return base.I64_extend_i32_u(v23)
												}
											}
										}
									}
								} else {
									v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v33 < int32(3) {
										return base.I64_extend_i32_u(v23)
									} else {
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
										if v22 == v36 {
											return base.I64_extend_i32_u(v23)
										} else {
											F_pfree(m, v22)
											mBase = m.M
											v39 = m.ExcPending
											if v39 != 0 {
												return int64(0)
											} else {
												return base.I64_extend_i32_u(v23)
											}
										}
									}
								}
							}
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							if v29 != v14 {
								F_pfree(m, v14)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return int64(0)
								} else {
									v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v33 < int32(3) {
										return base.I64_extend_i32_u(v23)
									} else {
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
										if v22 == v36 {
											return base.I64_extend_i32_u(v23)
										} else {
											F_pfree(m, v22)
											mBase = m.M
											v39 = m.ExcPending
											if v39 != 0 {
												return int64(0)
											} else {
												return base.I64_extend_i32_u(v23)
											}
										}
									}
								}
							} else {
								v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v33 < int32(3) {
									return base.I64_extend_i32_u(v23)
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
									if v22 == v36 {
										return base.I64_extend_i32_u(v23)
									} else {
										F_pfree(m, v22)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int64(0)
										} else {
											return base.I64_extend_i32_u(v23)
										}
									}
								}
							}
						}
					}
				}
			} else {
				v22 = int32(0)
				v23 = F_encrypt_internal(m, l2, l1, v9, v14, v22)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int64(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v25 != v9 {
						F_pfree(m, v9)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int64(0)
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							if v29 != v14 {
								F_pfree(m, v14)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return int64(0)
								} else {
									v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v33 < int32(3) {
										return base.I64_extend_i32_u(v23)
									} else {
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
										if v22 == v36 {
											return base.I64_extend_i32_u(v23)
										} else {
											F_pfree(m, v22)
											mBase = m.M
											v39 = m.ExcPending
											if v39 != 0 {
												return int64(0)
											} else {
												return base.I64_extend_i32_u(v23)
											}
										}
									}
								}
							} else {
								v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v33 < int32(3) {
									return base.I64_extend_i32_u(v23)
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
									if v22 == v36 {
										return base.I64_extend_i32_u(v23)
									} else {
										F_pfree(m, v22)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int64(0)
										} else {
											return base.I64_extend_i32_u(v23)
										}
									}
								}
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v29 != v14 {
							F_pfree(m, v14)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int64(0)
							} else {
								v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v33 < int32(3) {
									return base.I64_extend_i32_u(v23)
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
									if v22 == v36 {
										return base.I64_extend_i32_u(v23)
									} else {
										F_pfree(m, v22)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int64(0)
										} else {
											return base.I64_extend_i32_u(v23)
										}
									}
								}
							}
						} else {
							v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
							if v33 < int32(3) {
								return base.I64_extend_i32_u(v23)
							} else {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
								if v22 == v36 {
									return base.I64_extend_i32_u(v23)
								} else {
									F_pfree(m, v22)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return int64(0)
									} else {
										return base.I64_extend_i32_u(v23)
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
func Fn14354(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int64
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	v6 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(176)
	m.G0 = v14
	v23 = v6
	v24 = int32(-1)
	v25 = v6
	v26 = v6
	goto L1
L1:
	;
	if v24 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_Fn14354[0])) = v47
	*(*int32)(unsafe.Add(mBase, _c_Fn14354[1])) = v48
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v98 - int32(1)
	m.G0 = v14 + int32(176)
	return
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v31 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v30 + v31
	v35 = *(*int32)(unsafe.Add(mBase, _c_Fn14354[0]))
	v37 = *(*int32)(unsafe.Add(mBase, _c_Fn14354[1]))
	v39 = v14 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v14 + int32(12)
	goto L6
L4:
	;
	v46 = v23
	v47 = v25
	v48 = v26
	goto L5
L5:
	;
	goto L8
L6:
	;
	v46 = int32(0)
	v47 = v35
	v48 = v37
	goto L5
L7:
	;
	goto L2
L8:
	;
	if v46 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L7
L10:
	;
	v73 = int32(m.ExcTag)
	v74 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v73 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	F_standard_ExecutorRun(m, l0, l1, l2)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L10
	} else {
		goto L18
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_Fn14354[0])) = v14 + int32(16)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v55 == int32(0) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_Fn14354[1])) = v48
	*(*int32)(unsafe.Add(mBase, _c_Fn14354[0])) = v47
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v64 - int32(1)
	F_pg_re_throw(m)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L10
	} else {
		goto L17
	}
L15:
	;
	m.T0[v55].(func(*base.Module, int32, int32, int64))(m, l0, l1, l2)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	goto L7
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	goto L9
L19:
	;
	v78 = int32(v74)
	m.G0 = v14
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	if v14+int32(12) == v84 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	m.ExcPending = 1
	goto L28
L21:
	;
	if v88 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v88 = v86
	goto L24
L23:
	;
	v88 = int32(0)
	goto L24
L24:
	;
	goto L21
L25:
	;
	F___wasm_longjmp(m, v81, v80)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v23 = v80
	v24 = v88
	v25 = v47
	v26 = v48
	goto L1
L28:
	;
	return
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func Fn14361(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	v6 = F_palloc0(m, int32(24))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l2
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(0)
		v14 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v14
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = v14
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v6
		return v14
	}
}
func Fn14370(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 float64
	_ = v10
	var v13 int32
	_ = v13
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v10 = F_patternsel_common(m, v3, v4, v5, v6, v7, v8, l1, v5)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		return base.I64_reinterpret_f64(v10)
	}
}
func Fn14383(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v75 int64
	_ = v75
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int64
	_ = v92
	var v93 int64
	_ = v93
	var v103 int64
	_ = v103
	var v106 int64
	_ = v106
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v118 int64
	_ = v118
	var v119 int64
	_ = v119
	var v120 int64
	_ = v120
	var v125 int64
	_ = v125
	var v130 int64
	_ = v130
	var v135 int64
	_ = v135
	var v141 int32
	_ = v141
	var v149 int64
	_ = v149
	var v152 int64
	_ = v152
	var v154 int64
	_ = v154
	var v156 int64
	_ = v156
	var v158 int64
	_ = v158
	var v162 int64
	_ = v162
	var v164 int64
	_ = v164
	var v172 int64
	_ = v172
	var v177 int64
	_ = v177
	var v191 int64
	_ = v191
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	if v7 <= v8 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L40
	} else {
		goto L45
	}
L2:
	;
	if base.B2i32(v7 == int64(-9223372036854775807-1))|base.B2i32(v8 == int64(9223372036854775807)) != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L40
	} else {
		goto L41
	}
L5:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _c_Fn14383[0])))
	if v16 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v20 = int32(16)
	v21 = int32(0)
	v25 = m.G0
	v27 = v25 - v20
	m.G0 = v27
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v21
	v33 = F_open(m, int32(_a_Fn14383_0), v21, v27)
	mBase = m.M
	if v33 != int32(-1) {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	goto L8
L8:
	;
	if v7 < v8 {
		goto L34
	} else {
		goto L35
	}
L9:
	;
	v141 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_Fn14383[0])) = uint8(v141)
	goto L8
L10:
	;
	if v66 != 0 {
		goto L23
	} else {
		goto L24
	}
L11:
	;
	goto L15
L12:
	;
	v66 = v21
	goto L13
L13:
	;
	m.G0 = v27 + int32(16)
	goto L10
L14:
	;
	v61 = F_close(m, v33)
	mBase = m.M
	v66 = v59
	goto L13
L15:
	;
	v39 = int32(_a_Fn14383_1)
	v40 = v20
	goto L16
L16:
	;
	v45 = F_read(m, v33, v39, v40)
	mBase = m.M
	if v45 <= int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v59 = int32(1)
	goto L14
L18:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_Fn14383[1]))
	if v49 == int32(27) {
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v54 = v40 - v45
	if v54 != 0 {
		v39 = v39 + v45
		v40 = v54
		goto L16
	} else {
		goto L22
	}
L21:
	;
	v59 = int32(0)
	goto L14
L22:
	;
	goto L17
L23:
	;
	v71 = int32(_a_Fn14383_1)
	v72 = *(*int64)(unsafe.Add(mBase, _c_Fn14383[2]))
	if v72 != int64(0) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	goto L25
L25:
	;
	v83 = int32(_a_Fn14383_1)
	v87 = m.G0
	v88 = int32(16)
	v89 = v87 - v88
	m.G0 = v89
	F_gettimeofday(m, v89)
	mBase = m.M
	v92 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
	v93 = int64(*(*int32)(unsafe.Add(mBase, uint32(v89)+8)))
	m.G0 = v89 + v88
	goto L31
L26:
	;
	goto L9
L27:
	;
	goto L26
L28:
	;
	v75 = *(*int64)(unsafe.Add(mBase, _c_Fn14383[3]))
	if v75 != int64(0) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, _c_Fn14383[3])) = int64(1442695040888963407)
	*(*int64)(unsafe.Add(mBase, _c_Fn14383[2])) = int64(6364136223846793005)
	goto L27
L31:
	;
	v103 = int64(*(*uint32)(unsafe.Add(mBase, _c_Fn14383[4])))
	v106 = v93 + v92*int64(1000000) - int64(946684800000000) ^ v103<<(uint(int64(32))%64)
	v109 = v106 + int64(4354685564936845354)
	v110 = int64(30)
	v113 = int64(-4658895280553007687)
	v114 = (int64(base.Ui64(v109)>>(uint(v110)%64)) ^ v109) * v113
	v115 = int64(27)
	v118 = int64(-7723592293110705685)
	v119 = (int64(base.Ui64(v114)>>(uint(v115)%64)) ^ v114) * v118
	v120 = int64(31)
	*(*int64)(unsafe.Add(mBase, _c_Fn14383[3])) = int64(base.Ui64(v119)>>(uint(v120)%64)) ^ v119
	v125 = v106 - int64(7046029254386353131)
	v130 = (int64(base.Ui64(v125)>>(uint(v110)%64)) ^ v125) * v113
	v135 = (int64(base.Ui64(v130)>>(uint(v115)%64)) ^ v130) * v118
	*(*int64)(unsafe.Add(mBase, _c_Fn14383[2])) = int64(base.Ui64(v135)>>(uint(v120)%64)) ^ v135
	goto L32
L32:
	;
	goto L9
L33:
	;
	return v191
L34:
	;
	v149 = v8 - v7
	v152 = *(*int64)(unsafe.Add(mBase, _c_Fn14383[3]))
	v154 = *(*int64)(unsafe.Add(mBase, _c_Fn14383[2]))
	v156 = v154
	v158 = v152
	goto L37
L35:
	;
	v191 = v7
	goto L36
L36:
	;
	goto L33
L37:
	;
	v162 = v156 ^ v158
	v164 = base.I64_rotl(v162, int64(37))
	v172 = v162 ^ (v162<<(uint(int64(16))%64) ^ base.I64_rotl(v156, int64(24)))
	v177 = int64(base.Ui64(base.I64_rotl(v156*int64(5), int64(7))*int64(9)) >> (uint(base.I64_clz(v149)) % 64))
	if base.Ui64(v149) < base.Ui64(v177) {
		v156 = v172
		v158 = v164
		goto L37
	} else {
		goto L39
	}
L38:
	;
	*(*int64)(unsafe.Add(mBase, _c_Fn14383[3])) = v164
	*(*int64)(unsafe.Add(mBase, _c_Fn14383[2])) = v172
	v191 = v7 + v177
	goto L36
L39:
	;
	goto L38
L40:
	;
	return int64(0)
L41:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	F_errmsg(m, int32(_a_Fn14383_2), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_Fn14383_3), l3, l1)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L40
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L40
	} else {
		goto L46
	}
L46:
	;
	F_errmsg(m, int32(_a_Fn14383_4), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L40
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_Fn14383_3), l2, l1)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L40
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func Fn14385(m *base.Module, l0 int32, l1 int32) int64 {
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v15 = F_ArrayGetIntegerTypmods(m, v9, v6+int32(12))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			if v17 != int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_Fn14385_0), int32(0))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_Fn14385_1), int32(58), int32(_a_Fn14385_2))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				v37 = F_anytime_typmod_check(m, l1, v36)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int64(0)
				} else {
					m.G0 = v6 + int32(16)
					return base.I64_extend_i32_u(v37)
				}
			}
		}
	}
}
func Fn14389(m *base.Module, l0 int32, l1 int32) int64 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = F_text_to_cstring(m, v11)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _c_Fn14389[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v18
			v21 = *(*int64)(unsafe.Add(mBase, _c_Fn14389[1]))
			*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v21
			v28 = F_DirectInputFunctionCallSafe(m, l1, v15, int32(-1), v8+int32(8), v8+int32(24))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int64(0)
			} else {
				if v28 == int32(0) {
					v32 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v32)
					v36 = int64(0)
				} else {
					v35 = *(*int64)(unsafe.Add(mBase, uint32(v8)+24))
					v36 = v35
				}
				m.G0 = v8 + int32(32)
				return v36
			}
		}
	}
}
func Fn14394(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v7 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_Fn14394_0), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_Fn14394_1), l4, l3)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v27 = F_heap_getsysattr(m, v7, l1, l2)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int64(0)
		} else {
			return v27
		}
	}
}
func Fn14398(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_check_encoding_conversion_args(m, v13, v14, v15, int32(6), l1)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v23 = F_UtfToLocal(m, v11, v15, v10, l5, l4, l3, l2, l1, base.B2i32(v12 != int64(0)))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_s(v23)
		}
	}
}
func Fn14406(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 float32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = int32(1)
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			v20 = v18 & v16
			if v20 != 0 {
				v21 = v16
			} else {
				v21 = int32(4)
			}
			if v18 == int32(1) {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
				if v28 == int32(18) {
					v31 = int32(16)
				} else {
					v31 = int32(0)
				}
				if base.Ui32((v28-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v38 = int32(4)
				} else {
					v38 = v31
				}
				v49 = v38
			} else {
				v39 = int32(1)
				if v20 != 0 {
					v49 = int32(base.Ui32(v18)>>(uint(v39)%32)) - v39
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v49 = int32(base.Ui32(v43)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v50 = int32(1)
			v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			v54 = v52 & v50
			if v54 != 0 {
				v55 = v50
			} else {
				v55 = int32(4)
			}
			if v52 == int32(1) {
				v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
				if v62 == int32(18) {
					v65 = int32(16)
				} else {
					v65 = int32(0)
				}
				if base.Ui32((v62-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v72 = int32(4)
				} else {
					v72 = v65
				}
				v83 = v72
			} else {
				v73 = int32(1)
				if v54 != 0 {
					v83 = int32(base.Ui32(v52)>>(uint(v73)%32)) - v73
				} else {
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v83 = int32(base.Ui32(v77)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v84 = F_calc_word_similarity(m, v21+v9, v49, v14+v55, v83, l1)
			mBase = m.M
			v85 = m.ExcPending
			if v85 != 0 {
				return int64(0)
			} else {
				v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v86 != v9 {
					F_pfree(m, v9)
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return int64(0)
					} else {
						v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v90 != v14 {
							F_pfree(m, v14)
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_s(base.I32_reinterpret_f32(v84))
							}
						} else {
							return base.I64_extend_i32_s(base.I32_reinterpret_f32(v84))
						}
					}
				} else {
					v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v90 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_s(base.I32_reinterpret_f32(v84))
						}
					} else {
						return base.I64_extend_i32_s(base.I32_reinterpret_f32(v84))
					}
				}
			}
		}
	}
}
func Fn14408(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 float32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v17 = F_pg_detoast_datum_packed(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
			v21 = v19 & int32(1)
			if v21 != 0 {
				v22 = int32(1)
			} else {
				v22 = int32(4)
			}
			if v19 == int32(1) {
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
				if v29 == int32(18) {
					v32 = int32(16)
				} else {
					v32 = int32(0)
				}
				if base.Ui32((v29-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v39 = int32(4)
				} else {
					v39 = v32
				}
				v50 = v39
			} else {
				v40 = int32(1)
				if v21 != 0 {
					v50 = int32(base.Ui32(v19)>>(uint(v40)%32)) - v40
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
					v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v51 = int32(1)
			v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
			v55 = v53 & v51
			if v55 != 0 {
				v56 = v51
			} else {
				v56 = int32(4)
			}
			if v53 == int32(1) {
				v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
				if v63 == int32(18) {
					v66 = int32(16)
				} else {
					v66 = int32(0)
				}
				if base.Ui32((v63-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v73 = int32(4)
				} else {
					v73 = v66
				}
				v84 = v73
			} else {
				v74 = int32(1)
				if v55 != 0 {
					v84 = int32(base.Ui32(v53)>>(uint(v74)%32)) - v74
				} else {
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v84 = int32(base.Ui32(v78)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v85 = F_calc_word_similarity(m, v17+v22, v50, v10+v56, v84, l1)
			mBase = m.M
			v86 = m.ExcPending
			if v86 != 0 {
				return int64(0)
			} else {
				v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v87 != v10 {
					F_pfree(m, v10)
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int64(0)
					} else {
						v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v91 != v17 {
							F_pfree(m, v17)
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_s(base.I32_reinterpret_f32(base.F32_sub(float32(1), v85)))
							}
						} else {
							return base.I64_extend_i32_s(base.I32_reinterpret_f32(base.F32_sub(float32(1), v85)))
						}
					}
				} else {
					v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v91 != v17 {
						F_pfree(m, v17)
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_s(base.I32_reinterpret_f32(base.F32_sub(float32(1), v85)))
						}
					} else {
						return base.I64_extend_i32_s(base.I32_reinterpret_f32(base.F32_sub(float32(1), v85)))
					}
				}
			}
		}
	}
}
func Fn14413(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
		F_errmsg_internal(m, l4, v9)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			F_errfinish(m, l3, l2, l1)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
