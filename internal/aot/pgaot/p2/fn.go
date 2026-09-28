package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_fn_expr_arg_stable(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	v3 = int32(0)
	if l0 == v3 {
		v49 = v3
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v7 == int32(0) {
			v49 = v3
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			v12 = v10 - int32(11)
			v19 = int32(0)
			if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v12))|base.B2i32(int32(base.Ui32(int32(977))>>(uint(v12)%32))&int32(1) == v19)|base.B2i32(l1 < v19) != 0 {
				v49 = v3
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v12<<(uint(int32(2))%32))+uint32(_c_F_get_fn_expr_arg_stable[0])))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v7+v27)))
				if v29 == int32(0) {
					v49 = v3
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
					if v32 <= l1 {
						v49 = v3
					} else {
						v34 = int32(1)
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+l1<<(uint(int32(2))%32))))
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
						switch v40 - int32(7) {
						case 0:
							v49 = v34
						case 1:
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
							if v43 == int32(0) {
								v49 = v34
							} else {
								v49 = int32(0)
							}
						default:
							v49 = int32(0)
						}
					}
				}
			}
		}
	}
	return v49
}
func F_has_fn_opclass_options(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	v2 = int32(0)
	if l0 == v2 {
		v18 = v2
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v5 == int32(0) {
			v18 = v2
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			if v8 != int32(7) {
				v18 = v2
			} else {
				v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
				if v11 != int32(17) {
					v18 = v2
				} else {
					v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+32)))
					v18 = v14 ^ int32(1)
				}
			}
		}
	}
	return v18 & int32(1)
}
func F_make_fn_arguments(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	v5 = int32(0)
	if l1 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v12 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = v5
	goto L4
L4:
	;
	v25 = v20 << (uint(int32(2)) % 32)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l2+v25)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l3+v25)))
	if v27 == v29 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	v55 = v20 + int32(1)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v55 < v56 {
		v20 = v55
		goto L4
	} else {
		goto L14
	}
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v32 = v31 + v25
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if v34 == int32(16) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v38 = int32(-1)
	v42 = F_coerce_type(m, l0, v37, v27, v29, v38, int32(0), int32(2), v38)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v45 = int32(-1)
	v49 = F_coerce_type(m, l0, v33, v27, v29, v45, int32(0), int32(2), v45)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L11
	} else {
		goto L13
	}
L11:
	;
	return
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v42
	goto L6
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v49
	goto L6
L14:
	;
	goto L5
}
func Fn14211(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	v7 = m.G0
	v9 = v7 - int32(112)
	m.G0 = v9
	F_ScanKeyInit(m, v9, int32(1), int32(3), int32(184), base.I64_extend_i32_u(l0))
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_ScanKeyInit(m, v9+int32(56), int32(2), int32(3), int32(184), base.I64_extend_i32_u(l1))
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = F_table_open(m, l3, int32(3))
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v31 = F_systable_beginscan(m, v26, l2, int32(1), int32(0), int32(2), v9)
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v33 = F_systable_getnext(m, v31)
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v33 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v35 = v33
	goto L10
L8:
	;
	goto L9
L9:
	;
	F_systable_endscan(m, v31)
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L15
	}
L10:
	;
	F_simple_heap_delete(m, v26, v35+int32(4))
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	v45 = F_systable_getnext(m, v31)
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v45 != 0 {
		v35 = v45
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	F_relation_close(m, v26, int32(3))
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	m.G0 = v9 + int32(112)
	return
}
func Fn14222(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l3
	*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
	F_ShmemRequestStructWithOpts(m, v7)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		m.G0 = v7 + int32(16)
		return
	}
}
func Fn14224(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v11
	v15 = F_LockRelease(m, v7, l1, l2)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		m.G0 = v7 + int32(16)
		return
	}
}
func Fn14228(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	if int32(0) <= l1 {
		if base.Ui32(l1) < base.Ui32(int32(7)) {
			v44 = l1
			m.G0 = v13 + int32(32)
			return v44
		} else {
			v19 = int32(6)
			v22 = F_errstart(m, int32(19), int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				if v22 == int32(0) {
					v44 = v19
					m.G0 = v13 + int32(32)
					return v44
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = int32(6)
						if l0 != 0 {
							v35 = int32(_a_Fn14228_0)
						} else {
							v35 = int32(_a_Fn14228_1)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v35
						*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l1
						F_errmsg(m, l7, v13+int32(16))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, l4, l6, l2)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								v44 = v19
								m.G0 = v13 + int32(32)
								return v44
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return int32(0)
			} else {
				if l0 != 0 {
					v58 = int32(_a_Fn14228_0)
				} else {
					v58 = int32(_a_Fn14228_1)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v58
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
				F_errmsg(m, l5, v13)
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, l4, l3, l2)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
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
func Fn14235(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = l5
			F_errmsg(m, l4, v9)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, l3, l2, l1)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func Fn14239(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v15 = F_pg_detoast_datum_packed(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
			if v17 == int32(1) {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
				if v23 == int32(18) {
					v26 = int32(16)
				} else {
					v26 = int32(0)
				}
				if base.Ui32((v23-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v33 = int32(4)
				} else {
					v33 = v26
				}
				v46 = v33
			} else {
				v34 = int32(1)
				if v17&v34 != 0 {
					v46 = int32(base.Ui32(v17)>>(uint(v34)%32)) - v34
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
			if v47 == int32(1) {
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
				if v53 == int32(18) {
					v56 = int32(16)
				} else {
					v56 = int32(0)
				}
				if base.Ui32((v53-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v63 = int32(4)
				} else {
					v63 = v56
				}
				v76 = v63
			} else {
				v64 = int32(1)
				if v47&v64 != 0 {
					v76 = int32(base.Ui32(v47)>>(uint(v64)%32)) - v64
				} else {
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
					v76 = int32(base.Ui32(v70)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v77 = int32(1)
			if v17&v77 != 0 {
				v81 = v77
			} else {
				v81 = int32(4)
			}
			v83 = int32(1)
			if v47&v83 != 0 {
				v87 = v83
			} else {
				v87 = int32(4)
			}
			v89 = F_dotrim(m, v10+v81, v46, v15+v87, v76, l2, l1)
			mBase = m.M
			v90 = m.ExcPending
			if v90 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v89)
			}
		}
	}
}
func Fn14242(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int64
	_ = v37
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = v10 - v7
	v14 = int32(0)
	if base.B2i32(base.B2i32(int64(0) < v7)^base.B2i32(v11 < v10) == v14)&base.B2i32(v11 != int64(-9223372036854775807-1)) == v14 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, l3, int32(0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, l2, int32(108), l1)
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
		v37 = v11 >> (uint(int64(63)) % 64)
		return v11 ^ v37 - v37
	}
}
func Fn14248(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	v14 = m.G0
	v16 = v14 + int32(-64)
	m.G0 = v16
	v18 = F_strlen(m, l0)
	mBase = m.M
	if v18 != 0 {
		v20 = F_strstr(m, l0, int32(_a_Fn14248_0))
		mBase = m.M
		if v20 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = l0
					F_errmsg(m, l4, v14+int32(-16))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						v64 = F_errdetail(m, l8, int32(0))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_Fn14248_1), l7, l1)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
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
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v21 == int32(45) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l0
						F_errmsg(m, l4, v14+int32(-48))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							v82 = F_errdetail(m, l6, int32(0))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_Fn14248_1), l5, l1)
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
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
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v18-int32(1)))))
				if v27 == int32(45) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l0
							F_errmsg(m, l4, v14+int32(-48))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return
							} else {
								v82 = F_errdetail(m, l6, int32(0))
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_Fn14248_1), l5, l1)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
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
					v31 = Fn14265(m, l0, int32(47))
					mBase = m.M
					if v31 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = l0
								F_errmsg(m, l4, v14+int32(-32))
								mBase = m.M
								v98 = m.ExcPending
								if v98 != 0 {
									return
								} else {
									v100 = F_errdetail(m, l3, int32(0))
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_Fn14248_1), l2, l1)
										mBase = m.M
										v104 = m.ExcPending
										if v104 != 0 {
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
						m.G0 = v16 - int32(-64)
						return
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = l0
				F_errmsg(m, l4, v16)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					v46 = F_errdetail(m, l10, int32(0))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_Fn14248_1), l9, l1)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
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
func Fn14255(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v14 {
	case 0:
		v37 = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v37
		v40 = int32(0)
		m.G0 = v11 + int32(16)
		return v40
	default:
		v15 = int32(23)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v17 = F_errsave_start(m, v16)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			if v17 == int32(0) {
				v40 = v15
				m.G0 = v11 + int32(16)
				return v40
			} else {
				F_errcode(m, int32(33685634))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = v26
					F_errmsg(m, l4, v11)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						v32 = F_errdetail(m, int32(_a_Fn14255_0), int32(0))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							F_errsave_finish(m, v16, l3, l2, l1)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								v40 = v15
								m.G0 = v11 + int32(16)
								return v40
							}
						}
					}
				}
			}
		}
	case 3:
		v37 = int32(4)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v37
		v40 = int32(0)
		m.G0 = v11 + int32(16)
		return v40
	}
}
func Fn14260(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = int32(_a_Fn14260_0)
	v16 = *(*int32)(unsafe.Add(mBase, _c_Fn14260[0]))
	*(*int32)(unsafe.Add(mBase, _c_Fn14260[0])) = v16 + int32(1)
	v21 = *(*int32)(unsafe.Add(mBase, _c_Fn14260[1]))
	if int32(0) <= v21 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = int32(_a_Fn14260_1)
	v25 = *(*int32)(unsafe.Add(mBase, _c_Fn14260[2]))
	v28 = v21 * int32(100)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_Fn14260[3])))
	*(*int32)(unsafe.Add(mBase, _c_Fn14260[2])) = v31
	v34 = v12 + int32(16)
	F_initStringInfo(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_Fn14260[1])) = int32(-1)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L21
	}
L4:
	;
	return
L5:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_Fn14260[4])))
	*(*int32)(unsafe.Add(mBase, _c_Fn14260[5])) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l1
	v41 = F_appendStringInfoVA(m, v34, l0, l1)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v41 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v49 = v41
	goto L10
L8:
	;
	goto L9
L9:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_Fn14260[6])))
	if v71 != 0 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	v53 = v12 + int32(16)
	F_enlargeStringInfo(m, v53, v49)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_Fn14260[4])))
	*(*int32)(unsafe.Add(mBase, _c_Fn14260[5])) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l1
	v60 = F_appendStringInfoVA(m, v53, l0, l1)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	if v60 != 0 {
		v49 = v60
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	F_pfree(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v75 = F_pstrdup(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_Fn14260[6]))) = v75
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	F_pfree(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_Fn14260[2])) = v25
	v83 = int32(_a_Fn14260_0)
	v85 = *(*int32)(unsafe.Add(mBase, _c_Fn14260[0]))
	*(*int32)(unsafe.Add(mBase, _c_Fn14260[0])) = v85 - int32(1)
	m.G0 = v12 + int32(32)
	return
L21:
	;
	F_errmsg_internal(m, int32(_a_Fn14260_2), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_Fn14260_3), l3, l2)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func Fn14266(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 float32, l5 float32) int64 {
	mBase := m.M
	_ = mBase
	var v9 float32
	_ = v9
	var v10 float32
	_ = v10
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int64
	_ = v47
	v9 = *(*float32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = base.F32_nearest(v9)
	v17 = int32(0)
	if base.B2i32(base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(base.I32_reinterpret_f32(v10)&int32(2147483647)))|base.B2i32(base.F32_ge(v10, l5) == v17) == v17)&base.F32_lt(v10, l4) == v17 {
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v27 = F_errsave_start(m, v26)
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int64(0)
		} else {
			if v27 == int32(0) {
				v47 = int64(0)
				return v47
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, l3, int32(0))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int64(0)
					} else {
						F_errsave_finish(m, v26, int32(_a_Fn14266_0), l2, l1)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int64(0)
						} else {
							return int64(0)
						}
					}
				}
			}
		}
	} else {
		v47 = base.I64_extend_i32_s(base.I32_trunc_sat_f32_s(v10))
		return v47
	}
}
func Fn14271(m *base.Module, l0 int32, l1 int32) int64 {
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
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = F_gbt_var_union(m, v3, v4, v5, l1, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v7)
	}
}
func Fn14286(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
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
	v5 = F_SearchSysCache1(m, l1, base.I64_extend_i32_u(l0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)))
			v18 = F_pstrdup(m, v13+v14+int32(4))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_ReleaseCatCache(m, v5)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					return v18
				}
			}
		}
	}
}
func Fn14288(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v5 = F_SearchSysCache1(m, l1, base.I64_extend_i32_u(l0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v13+v14)+8))
			F_ReleaseCatCache(m, v5)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return v16
			}
		}
	}
}
func Fn14293(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = F_SearchSysCache1(m, l6, base.I64_extend_i32_u(l0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		if v16 == int32(0) {
			if l1 != 0 {
				v41 = int32(0)
				m.G0 = v13 + int32(16)
				return v41
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
					F_errmsg_internal(m, l5, v13)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_Fn14293_0), l4, l3)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
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
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+22)))
			v36 = F_pstrdup(m, v32+v33+l2)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				F_ReleaseCatCache(m, v16)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v41 = v36
					m.G0 = v13 + int32(16)
					return v41
				}
			}
		}
	}
}
func Fn14295(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v97 int32
	_ = v97
	v5 = int32(0)
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v14 == v5 {
		v31 = v5
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v31&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	if v18 == int32(0) {
		v31 = v5
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v21 != int32(7) {
		v31 = v5
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v24 != int32(17) {
		v31 = v5
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+32)))
	v31 = v27 ^ int32(1)
	goto L2
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v35 = F_get_fn_opclass_options(m, v34)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v40 = l3
	goto L9
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v44 = v43 & l2
	v45 = base.I32_wrap_i64(v11)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+4)))
	if v46&l2 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	return int64(0)
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v40 = v39
	goto L9
L12:
	;
	return v11 & int64(4294967295)
L13:
	;
	v97 = int32(base.Ui32(v44) >> (uint(l1) % 32))
	goto L15
L14:
	;
	if v44 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v45))) = uint8(v97)
	goto L12
L16:
	;
	v97 = int32(0)
	goto L15
L17:
	;
	v49 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v45))) = uint8(v49)
	v51 = int32(0)
	if v40 <= v51 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v54 = int32(8)
	v58 = v51
	goto L19
L19:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+(v13+v54)))))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+(v12+v54)))))
	if v69 != v71 {
		goto L16
	} else {
		goto L21
	}
L20:
	;
	goto L12
L21:
	;
	v74 = v58 + int32(1)
	if v40 != v74 {
		v58 = v74
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
}
func Fn14299(m *base.Module, l0 int32, l1 int64) int64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 float32
	_ = v135
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 float32
	_ = v227
	var v234 int32
	_ = v234
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
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
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+4)))
	v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+4)))
	v22 = base.B2i32(v20 < v21)
	if v20 < v21 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v20 < v21 {
		goto L104
	} else {
		goto L105
	}
L5:
	;
	v23 = v20
	goto L7
L6:
	;
	v23 = v21
	goto L7
L7:
	;
	if v23 <= int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v26 = int32(8)
	v31 = int32(0)
	goto L9
L9:
	;
	v43 = v31 << (uint(int32(1)) % 32)
	v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+v26+v43))))
	v50 = v45 & int32(1023)
	v54 = v45 << (uint(int32(16)) % 32) & int32(-2147483648)
	v57 = int32(31)
	v58 = int32(base.Ui32(v45)>>(uint(int32(10))%32)) & v57
	if v58 != v57 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	return int64(1)
L11:
	;
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18+v26+v43))))
	v142 = v137 & int32(1023)
	v146 = v137 << (uint(int32(16)) % 32) & int32(-2147483648)
	v149 = int32(31)
	v150 = int32(base.Ui32(v137)>>(uint(int32(10))%32)) & v149
	if v150 != v149 {
		goto L58
	} else {
		goto L59
	}
L12:
	;
	v135 = base.F32_reinterpret_i32(v131 | v130<<(uint(int32(13))%32))
	goto L11
L13:
	;
	v130 = v50
	v131 = v58<<(uint(int32(23))%32) + v54 + int32(939524096)
	goto L12
L14:
	;
	if v45&int32(512) != 0 {
		goto L24
	} else {
		goto L25
	}
L15:
	;
	if v58 != 0 {
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if v50 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	if v50 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v130 = int32(0)
	v131 = v54
	goto L12
L20:
	;
	v130 = int32(0)
	v131 = v54 | int32(2139095040)
	goto L12
L21:
	;
	goto L22
L22:
	;
	v130 = v50
	v131 = v54 | int32(2143289344)
	goto L12
L23:
	;
	v130 = v118 & int32(1022)
	v131 = v120 | v54
	goto L12
L24:
	;
	v118 = v50 << (uint(int32(1)) % 32)
	v120 = int32(939524096)
	goto L23
L25:
	;
	goto L26
L26:
	;
	if base.Ui32(int32(255)) < base.Ui32(v50) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v118 = v50 << (uint(int32(2)) % 32)
	v120 = int32(931135488)
	goto L23
L28:
	;
	goto L29
L29:
	;
	if base.Ui32(int32(127)) < base.Ui32(v50) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v118 = v50 << (uint(int32(3)) % 32)
	v120 = int32(922746880)
	goto L23
L31:
	;
	goto L32
L32:
	;
	if base.Ui32(int32(63)) < base.Ui32(v50) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v118 = v50 << (uint(int32(4)) % 32)
	v120 = int32(914358272)
	goto L23
L34:
	;
	goto L35
L35:
	;
	if base.Ui32(int32(31)) < base.Ui32(v50) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v118 = v50 << (uint(int32(5)) % 32)
	v120 = int32(905969664)
	goto L23
L37:
	;
	goto L38
L38:
	;
	if base.Ui32(int32(15)) < base.Ui32(v50) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v118 = v50 << (uint(int32(6)) % 32)
	v120 = int32(897581056)
	goto L23
L40:
	;
	goto L41
L41:
	;
	if base.Ui32(int32(7)) < base.Ui32(v50) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v118 = v50 << (uint(int32(7)) % 32)
	v120 = int32(889192448)
	goto L23
L43:
	;
	goto L44
L44:
	;
	if base.Ui32(int32(3)) < base.Ui32(v50) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v118 = v50 << (uint(int32(8)) % 32)
	v120 = int32(880803840)
	goto L23
L46:
	;
	goto L47
L47:
	;
	v113 = base.B2i32(v50 == int32(1))
	if v50 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v114 = int32(1024)
	goto L50
L49:
	;
	v114 = v50 << (uint(int32(9)) % 32)
	goto L50
L50:
	;
	if v50 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v117 = int32(864026624)
	goto L53
L52:
	;
	v117 = int32(872415232)
	goto L53
L53:
	;
	v118 = v114
	v120 = v117
	goto L23
L54:
	;
	if base.F32_lt(v135, v227) != 0 {
		goto L97
	} else {
		goto L98
	}
L55:
	;
	v227 = base.F32_reinterpret_i32(v223 | v222<<(uint(int32(13))%32))
	goto L54
L56:
	;
	v222 = v142
	v223 = v150<<(uint(int32(23))%32) + v146 + int32(939524096)
	goto L55
L57:
	;
	if v137&int32(512) != 0 {
		goto L67
	} else {
		goto L68
	}
L58:
	;
	if v150 != 0 {
		goto L56
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if v142 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	if v142 != 0 {
		goto L57
	} else {
		goto L62
	}
L62:
	;
	v222 = int32(0)
	v223 = v146
	goto L55
L63:
	;
	v222 = int32(0)
	v223 = v146 | int32(2139095040)
	goto L55
L64:
	;
	goto L65
L65:
	;
	v222 = v142
	v223 = v146 | int32(2143289344)
	goto L55
L66:
	;
	v222 = v210 & int32(1022)
	v223 = v212 | v146
	goto L55
L67:
	;
	v210 = v142 << (uint(int32(1)) % 32)
	v212 = int32(939524096)
	goto L66
L68:
	;
	goto L69
L69:
	;
	if base.Ui32(int32(255)) < base.Ui32(v142) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v210 = v142 << (uint(int32(2)) % 32)
	v212 = int32(931135488)
	goto L66
L71:
	;
	goto L72
L72:
	;
	if base.Ui32(int32(127)) < base.Ui32(v142) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v210 = v142 << (uint(int32(3)) % 32)
	v212 = int32(922746880)
	goto L66
L74:
	;
	goto L75
L75:
	;
	if base.Ui32(int32(63)) < base.Ui32(v142) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v210 = v142 << (uint(int32(4)) % 32)
	v212 = int32(914358272)
	goto L66
L77:
	;
	goto L78
L78:
	;
	if base.Ui32(int32(31)) < base.Ui32(v142) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v210 = v142 << (uint(int32(5)) % 32)
	v212 = int32(905969664)
	goto L66
L80:
	;
	goto L81
L81:
	;
	if base.Ui32(int32(15)) < base.Ui32(v142) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v210 = v142 << (uint(int32(6)) % 32)
	v212 = int32(897581056)
	goto L66
L83:
	;
	goto L84
L84:
	;
	if base.Ui32(int32(7)) < base.Ui32(v142) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v210 = v142 << (uint(int32(7)) % 32)
	v212 = int32(889192448)
	goto L66
L86:
	;
	goto L87
L87:
	;
	if base.Ui32(int32(3)) < base.Ui32(v142) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v210 = v142 << (uint(int32(8)) % 32)
	v212 = int32(880803840)
	goto L66
L89:
	;
	goto L90
L90:
	;
	v205 = base.B2i32(v142 == int32(1))
	if v142 == int32(1) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v206 = int32(1024)
	goto L93
L92:
	;
	v206 = v142 << (uint(int32(9)) % 32)
	goto L93
L93:
	;
	if v142 == int32(1) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v209 = int32(864026624)
	goto L96
L95:
	;
	v209 = int32(872415232)
	goto L96
L96:
	;
	v210 = v206
	v212 = v209
	goto L66
L97:
	;
	return l1
L98:
	;
	goto L99
L99:
	;
	if base.F32_gt(v135, v227) == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v234 = v31 + int32(1)
	if v234 == v23 {
		goto L4
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	goto L10
L103:
	;
	v31 = v234
	goto L9
L104:
	;
	return l1
L105:
	;
	goto L106
L106:
	;
	return base.I64_extend_i32_u(base.B2i32(v21 < v20))
}
func Fn14303(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
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
	var v37 int64
	_ = v37
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int64(0)
	} else {
		v20 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v20)
		v22 = F_convert_any_priv_string(m, v16, l1)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int64(0)
		} else {
			v26 = F_object_aclcheck_ext(m, l2, v13, v14, v22, v11+int32(15))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int64(0)
			} else {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
				if v28 == int32(1) {
					v31 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
					v37 = int64(0)
				} else {
					v37 = base.I64_extend_i32_u(base.B2i32(v26 == int32(0)))
				}
				m.G0 = v11 + int32(16)
				return v37
			}
		}
	}
}
func Fn14305(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int64 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int64
	_ = v17
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v19 = F_pg_detoast_datum_packed(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int64(0)
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v24 = F_pg_detoast_datum_packed(m, v23)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int64(0)
		} else {
			v27 = F_text_to_cstring(m, v19)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int64(0)
			} else {
				v30 = F_DirectFunctionCall1Coll(m, l7, int32(0), base.I64_extend_i32_u(v27))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int64(0)
				} else {
					v32 = base.I32_wrap_i64(v30)
					if v32 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int64(0)
						} else {
							F_errcode(m, l6)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v15))) = v27
								F_errmsg(m, l5, v15)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_Fn14305_0), l4, l3)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
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
						v48 = F_convert_any_priv_string(m, v24, l1)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int64(0)
						} else {
							v50 = F_object_aclcheck(m, l2, v32, base.I32_wrap_i64(v17), v48)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int64(0)
							} else {
								m.G0 = v15 + int32(16)
								return base.I64_extend_i32_u(base.B2i32(v50 == int32(0)))
							}
						}
					}
				}
			}
		}
	}
}
func Fn14314(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = F_superuser_arg(m, l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 != 0 {
			v23 = int32(1)
			m.G0 = v7 + int32(16)
			return v23
		} else {
			if l0 == l1 {
				v23 = int32(0)
				m.G0 = v7 + int32(16)
				return v23
			} else {
				v18 = F_roles_is_member_of(m, l0, l2, l1, v7+int32(12))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					v23 = base.B2i32(v20 != int32(0))
					m.G0 = v7 + int32(16)
					return v23
				}
			}
		}
	}
}
func Fn14323(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_check_encoding_conversion_args(m, v11, v12, v13, l3, l2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v20 = F_local2local(m, v9, v8, v13, l3, l2, l1, base.B2i32(v10 != int64(0)))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_s(v20)
		}
	}
}
func Fn14334(m *base.Module, l0 int32, l1 int32) int64 {
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	v2 = l1
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v10 == int32(0) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v13 != 0 {
			v64 = v13
			v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
			if v66 == int32(0) {
				v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v70 = F_pg_detoast_datum(m, v69)
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int64(0)
				} else {
					F_do_numeric_accum(m, v64, v70)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int64(0)
					} else {
						m.G0 = v8 + int32(16)
						return base.I64_extend_i32_u(v64)
					}
				}
			} else {
				m.G0 = v8 + int32(16)
				return base.I64_extend_i32_u(v64)
			}
		} else {
			v16 = v8 + int32(12)
			v17 = int32(0)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v18 == v17 {
				v35 = int32(0)
				if v16 == v35 {
					v43 = v35
				} else {
					v38 = v35
					v39 = v17
					*(*int32)(unsafe.Add(mBase, uint32(v16))) = v38
					v43 = v39
				}
				v46 = v43
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				switch v21 - int32(435) {
				case 0:
					if v16 == int32(0) {
						v46 = int32(1)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+168))
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
						v38 = v28
						v39 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v16))) = v38
						v43 = v39
						v46 = v43
					}
				case 1:
					if v16 == int32(0) {
						v46 = int32(2)
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)+376))
						v38 = v33
						v39 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v16))) = v38
						v43 = v39
						v46 = v43
					}
				default:
					v35 = int32(0)
					if v16 == v35 {
						v43 = v35
					} else {
						v38 = v35
						v39 = v17
						*(*int32)(unsafe.Add(mBase, uint32(v16))) = v38
						v43 = v39
					}
					v46 = v43
				}
			}
			if v46 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v82 = m.ExcPending
				if v82 != 0 {
					return int64(0)
				} else {
					F_errmsg_internal(m, int32(_a_Fn14334_0), int32(0))
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_Fn14334_1), int32(_a_Fn14334_2), int32(_a_Fn14334_3))
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v49 = int32(_a_Fn14334_4)
				v50 = *(*int32)(unsafe.Add(mBase, _c_Fn14334[0]))
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
				*(*int32)(unsafe.Add(mBase, _c_Fn14334[0])) = v52
				v55 = F_palloc0(m, int32(112))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int64(0)
				} else {
					*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v2)
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v60
					*(*int32)(unsafe.Add(mBase, _c_Fn14334[0])) = v50
					v64 = v55
					v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
					if v66 == int32(0) {
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						v70 = F_pg_detoast_datum(m, v69)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int64(0)
						} else {
							F_do_numeric_accum(m, v64, v70)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int64(0)
							} else {
								m.G0 = v8 + int32(16)
								return base.I64_extend_i32_u(v64)
							}
						}
					} else {
						m.G0 = v8 + int32(16)
						return base.I64_extend_i32_u(v64)
					}
				}
			}
		}
	} else {
		v16 = v8 + int32(12)
		v17 = int32(0)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v18 == v17 {
			v35 = int32(0)
			if v16 == v35 {
				v43 = v35
			} else {
				v38 = v35
				v39 = v17
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = v38
				v43 = v39
			}
			v46 = v43
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
			switch v21 - int32(435) {
			case 0:
				if v16 == int32(0) {
					v46 = int32(1)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+168))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
					v38 = v28
					v39 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v16))) = v38
					v43 = v39
					v46 = v43
				}
			case 1:
				if v16 == int32(0) {
					v46 = int32(2)
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)+376))
					v38 = v33
					v39 = int32(2)
					*(*int32)(unsafe.Add(mBase, uint32(v16))) = v38
					v43 = v39
					v46 = v43
				}
			default:
				v35 = int32(0)
				if v16 == v35 {
					v43 = v35
				} else {
					v38 = v35
					v39 = v17
					*(*int32)(unsafe.Add(mBase, uint32(v16))) = v38
					v43 = v39
				}
				v46 = v43
			}
		}
		if v46 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return int64(0)
			} else {
				F_errmsg_internal(m, int32(_a_Fn14334_0), int32(0))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_Fn14334_1), int32(_a_Fn14334_2), int32(_a_Fn14334_3))
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v49 = int32(_a_Fn14334_4)
			v50 = *(*int32)(unsafe.Add(mBase, _c_Fn14334[0]))
			v52 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			*(*int32)(unsafe.Add(mBase, _c_Fn14334[0])) = v52
			v55 = F_palloc0(m, int32(112))
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int64(0)
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v2)
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v60
				*(*int32)(unsafe.Add(mBase, _c_Fn14334[0])) = v50
				v64 = v55
				v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
				if v66 == int32(0) {
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					v70 = F_pg_detoast_datum(m, v69)
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int64(0)
					} else {
						F_do_numeric_accum(m, v64, v70)
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int64(0)
						} else {
							m.G0 = v8 + int32(16)
							return base.I64_extend_i32_u(v64)
						}
					}
				} else {
					m.G0 = v8 + int32(16)
					return base.I64_extend_i32_u(v64)
				}
			}
		}
	}
}
func Fn14338(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v44 int32
	_ = v44
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v71 int32
	_ = v71
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v77 int32
	_ = v77
	var v86 int64
	_ = v86
	var v87 int32
	_ = v87
	var v90 int64
	_ = v90
	var v91 int32
	_ = v91
	var v95 int64
	_ = v95
	var v96 int32
	_ = v96
	var v115 int32
	_ = v115
	var v125 int32
	_ = v125
	v13 = int32(1)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+80)))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v20 == v13 {
		if v16&int32(1) == int32(0) {
			v41 = v19
			v42 = v19
			v44 = v13
			if v15&int32(1) != 0 {
				if v14&int32(1) == int32(0) {
					v55 = v17
					v57 = F_DirectFunctionCall2Coll(m, l2, int32(0), v42, v55)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int64(0)
					} else {
						if v57 != int64(0) {
							v115 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
							v125 = int32(0)
							return base.I64_extend_i32_u(v125)
						} else {
							v86 = v55
							v87 = int32(1)
							v90 = F_DirectFunctionCall2Coll(m, l1, int32(0), v42, v86)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int64(0)
							} else {
								if v90 != int64(0) {
									if v44 != 0 {
										v115 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
										v125 = int32(0)
										return base.I64_extend_i32_u(v125)
									} else {
										v95 = F_DirectFunctionCall2Coll(m, l1, int32(0), v86, v41)
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
											return int64(0)
										} else {
											if v87&base.B2i32(v95 == int64(0)) != 0 {
												v115 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
												v125 = int32(0)
												return base.I64_extend_i32_u(v125)
											} else {
												return base.I64_extend_i32_u(base.B2i32(v95 != int64(0)))
											}
										}
									}
								} else {
									if v87|v44 == int32(0) {
										v125 = int32(1)
									} else {
										v115 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
										v125 = int32(0)
									}
									return base.I64_extend_i32_u(v125)
								}
							}
						}
					}
				} else {
					v115 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
					v125 = int32(0)
					return base.I64_extend_i32_u(v125)
				}
			} else {
				if v14&int32(1) == int32(0) {
					v62 = int32(0)
					v65 = F_DirectFunctionCall2Coll(m, l2, v62, v18, v17)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int64(0)
					} else {
						v68 = base.B2i32(v65 == int64(0))
						if v65 == int64(0) {
							v69 = v18
						} else {
							v69 = v17
						}
						v70 = F_DirectFunctionCall2Coll(m, l2, v62, v42, v69)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int64(0)
						} else {
							if v70 == int64(0) {
								v86 = v69
								v87 = v62
								v90 = F_DirectFunctionCall2Coll(m, l1, int32(0), v42, v86)
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int64(0)
								} else {
									if v90 != int64(0) {
										if v44 != 0 {
											v115 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
											v125 = int32(0)
											return base.I64_extend_i32_u(v125)
										} else {
											v95 = F_DirectFunctionCall2Coll(m, l1, int32(0), v86, v41)
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return int64(0)
											} else {
												if v87&base.B2i32(v95 == int64(0)) != 0 {
													v115 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
													v125 = int32(0)
													return base.I64_extend_i32_u(v125)
												} else {
													return base.I64_extend_i32_u(base.B2i32(v95 != int64(0)))
												}
											}
										}
									} else {
										if v87|v44 == int32(0) {
											v125 = int32(1)
										} else {
											v115 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
											v125 = int32(0)
										}
										return base.I64_extend_i32_u(v125)
									}
								}
							} else {
								if v65 == int64(0) {
									v75 = v17
								} else {
									v75 = v18
								}
								v76 = F_DirectFunctionCall2Coll(m, l1, int32(0), v42, v75)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int64(0)
								} else {
									if v44&base.B2i32(v76 == int64(0)) != 0 {
										v115 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
										v125 = int32(0)
										return base.I64_extend_i32_u(v125)
									} else {
										return base.I64_extend_i32_u(base.B2i32(v76 != int64(0)))
									}
								}
							}
						}
					}
				} else {
					v55 = v18
					v57 = F_DirectFunctionCall2Coll(m, l2, int32(0), v42, v55)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int64(0)
					} else {
						if v57 != int64(0) {
							v115 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
							v125 = int32(0)
							return base.I64_extend_i32_u(v125)
						} else {
							v86 = v55
							v87 = int32(1)
							v90 = F_DirectFunctionCall2Coll(m, l1, int32(0), v42, v86)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int64(0)
							} else {
								if v90 != int64(0) {
									if v44 != 0 {
										v115 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
										v125 = int32(0)
										return base.I64_extend_i32_u(v125)
									} else {
										v95 = F_DirectFunctionCall2Coll(m, l1, int32(0), v86, v41)
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
											return int64(0)
										} else {
											if v87&base.B2i32(v95 == int64(0)) != 0 {
												v115 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
												v125 = int32(0)
												return base.I64_extend_i32_u(v125)
											} else {
												return base.I64_extend_i32_u(base.B2i32(v95 != int64(0)))
											}
										}
									}
								} else {
									if v87|v44 == int32(0) {
										v125 = int32(1)
									} else {
										v115 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
										v125 = int32(0)
									}
									return base.I64_extend_i32_u(v125)
								}
							}
						}
					}
				}
			}
		} else {
			v115 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
			v125 = int32(0)
			return base.I64_extend_i32_u(v125)
		}
	} else {
		v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		if v16&int32(1) != 0 {
			v41 = v19
			v42 = v27
			v44 = v13
			if v15&int32(1) != 0 {
				if v14&int32(1) == int32(0) {
					v55 = v17
					v57 = F_DirectFunctionCall2Coll(m, l2, int32(0), v42, v55)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int64(0)
					} else {
						if v57 != int64(0) {
							v115 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
							v125 = int32(0)
							return base.I64_extend_i32_u(v125)
						} else {
							v86 = v55
							v87 = int32(1)
							v90 = F_DirectFunctionCall2Coll(m, l1, int32(0), v42, v86)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int64(0)
							} else {
								if v90 != int64(0) {
									if v44 != 0 {
										v115 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
										v125 = int32(0)
										return base.I64_extend_i32_u(v125)
									} else {
										v95 = F_DirectFunctionCall2Coll(m, l1, int32(0), v86, v41)
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
											return int64(0)
										} else {
											if v87&base.B2i32(v95 == int64(0)) != 0 {
												v115 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
												v125 = int32(0)
												return base.I64_extend_i32_u(v125)
											} else {
												return base.I64_extend_i32_u(base.B2i32(v95 != int64(0)))
											}
										}
									}
								} else {
									if v87|v44 == int32(0) {
										v125 = int32(1)
									} else {
										v115 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
										v125 = int32(0)
									}
									return base.I64_extend_i32_u(v125)
								}
							}
						}
					}
				} else {
					v115 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
					v125 = int32(0)
					return base.I64_extend_i32_u(v125)
				}
			} else {
				if v14&int32(1) == int32(0) {
					v62 = int32(0)
					v65 = F_DirectFunctionCall2Coll(m, l2, v62, v18, v17)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int64(0)
					} else {
						v68 = base.B2i32(v65 == int64(0))
						if v65 == int64(0) {
							v69 = v18
						} else {
							v69 = v17
						}
						v70 = F_DirectFunctionCall2Coll(m, l2, v62, v42, v69)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int64(0)
						} else {
							if v70 == int64(0) {
								v86 = v69
								v87 = v62
								v90 = F_DirectFunctionCall2Coll(m, l1, int32(0), v42, v86)
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int64(0)
								} else {
									if v90 != int64(0) {
										if v44 != 0 {
											v115 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
											v125 = int32(0)
											return base.I64_extend_i32_u(v125)
										} else {
											v95 = F_DirectFunctionCall2Coll(m, l1, int32(0), v86, v41)
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return int64(0)
											} else {
												if v87&base.B2i32(v95 == int64(0)) != 0 {
													v115 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
													v125 = int32(0)
													return base.I64_extend_i32_u(v125)
												} else {
													return base.I64_extend_i32_u(base.B2i32(v95 != int64(0)))
												}
											}
										}
									} else {
										if v87|v44 == int32(0) {
											v125 = int32(1)
										} else {
											v115 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
											v125 = int32(0)
										}
										return base.I64_extend_i32_u(v125)
									}
								}
							} else {
								if v65 == int64(0) {
									v75 = v17
								} else {
									v75 = v18
								}
								v76 = F_DirectFunctionCall2Coll(m, l1, int32(0), v42, v75)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int64(0)
								} else {
									if v44&base.B2i32(v76 == int64(0)) != 0 {
										v115 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
										v125 = int32(0)
										return base.I64_extend_i32_u(v125)
									} else {
										return base.I64_extend_i32_u(base.B2i32(v76 != int64(0)))
									}
								}
							}
						}
					}
				} else {
					v55 = v18
					v57 = F_DirectFunctionCall2Coll(m, l2, int32(0), v42, v55)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int64(0)
					} else {
						if v57 != int64(0) {
							v115 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
							v125 = int32(0)
							return base.I64_extend_i32_u(v125)
						} else {
							v86 = v55
							v87 = int32(1)
							v90 = F_DirectFunctionCall2Coll(m, l1, int32(0), v42, v86)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int64(0)
							} else {
								if v90 != int64(0) {
									if v44 != 0 {
										v115 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
										v125 = int32(0)
										return base.I64_extend_i32_u(v125)
									} else {
										v95 = F_DirectFunctionCall2Coll(m, l1, int32(0), v86, v41)
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
											return int64(0)
										} else {
											if v87&base.B2i32(v95 == int64(0)) != 0 {
												v115 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
												v125 = int32(0)
												return base.I64_extend_i32_u(v125)
											} else {
												return base.I64_extend_i32_u(base.B2i32(v95 != int64(0)))
											}
										}
									}
								} else {
									if v87|v44 == int32(0) {
										v125 = int32(1)
									} else {
										v115 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
										v125 = int32(0)
									}
									return base.I64_extend_i32_u(v125)
								}
							}
						}
					}
				}
			}
		} else {
			v30 = int32(0)
			v32 = F_DirectFunctionCall2Coll(m, l2, v30, v27, v19)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int64(0)
			} else {
				v37 = base.B2i32(v32 == int64(0))
				if v32 == int64(0) {
					v38 = v27
				} else {
					v38 = v19
				}
				if v32 == int64(0) {
					v39 = v19
				} else {
					v39 = v27
				}
				v41 = v39
				v42 = v38
				v44 = v30
				if v15&int32(1) != 0 {
					if v14&int32(1) == int32(0) {
						v55 = v17
						v57 = F_DirectFunctionCall2Coll(m, l2, int32(0), v42, v55)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int64(0)
						} else {
							if v57 != int64(0) {
								v115 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
								v125 = int32(0)
								return base.I64_extend_i32_u(v125)
							} else {
								v86 = v55
								v87 = int32(1)
								v90 = F_DirectFunctionCall2Coll(m, l1, int32(0), v42, v86)
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int64(0)
								} else {
									if v90 != int64(0) {
										if v44 != 0 {
											v115 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
											v125 = int32(0)
											return base.I64_extend_i32_u(v125)
										} else {
											v95 = F_DirectFunctionCall2Coll(m, l1, int32(0), v86, v41)
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return int64(0)
											} else {
												if v87&base.B2i32(v95 == int64(0)) != 0 {
													v115 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
													v125 = int32(0)
													return base.I64_extend_i32_u(v125)
												} else {
													return base.I64_extend_i32_u(base.B2i32(v95 != int64(0)))
												}
											}
										}
									} else {
										if v87|v44 == int32(0) {
											v125 = int32(1)
										} else {
											v115 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
											v125 = int32(0)
										}
										return base.I64_extend_i32_u(v125)
									}
								}
							}
						}
					} else {
						v115 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
						v125 = int32(0)
						return base.I64_extend_i32_u(v125)
					}
				} else {
					if v14&int32(1) == int32(0) {
						v62 = int32(0)
						v65 = F_DirectFunctionCall2Coll(m, l2, v62, v18, v17)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int64(0)
						} else {
							v68 = base.B2i32(v65 == int64(0))
							if v65 == int64(0) {
								v69 = v18
							} else {
								v69 = v17
							}
							v70 = F_DirectFunctionCall2Coll(m, l2, v62, v42, v69)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int64(0)
							} else {
								if v70 == int64(0) {
									v86 = v69
									v87 = v62
									v90 = F_DirectFunctionCall2Coll(m, l1, int32(0), v42, v86)
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int64(0)
									} else {
										if v90 != int64(0) {
											if v44 != 0 {
												v115 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
												v125 = int32(0)
												return base.I64_extend_i32_u(v125)
											} else {
												v95 = F_DirectFunctionCall2Coll(m, l1, int32(0), v86, v41)
												mBase = m.M
												v96 = m.ExcPending
												if v96 != 0 {
													return int64(0)
												} else {
													if v87&base.B2i32(v95 == int64(0)) != 0 {
														v115 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
														v125 = int32(0)
														return base.I64_extend_i32_u(v125)
													} else {
														return base.I64_extend_i32_u(base.B2i32(v95 != int64(0)))
													}
												}
											}
										} else {
											if v87|v44 == int32(0) {
												v125 = int32(1)
											} else {
												v115 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
												v125 = int32(0)
											}
											return base.I64_extend_i32_u(v125)
										}
									}
								} else {
									if v65 == int64(0) {
										v75 = v17
									} else {
										v75 = v18
									}
									v76 = F_DirectFunctionCall2Coll(m, l1, int32(0), v42, v75)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return int64(0)
									} else {
										if v44&base.B2i32(v76 == int64(0)) != 0 {
											v115 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
											v125 = int32(0)
											return base.I64_extend_i32_u(v125)
										} else {
											return base.I64_extend_i32_u(base.B2i32(v76 != int64(0)))
										}
									}
								}
							}
						}
					} else {
						v55 = v18
						v57 = F_DirectFunctionCall2Coll(m, l2, int32(0), v42, v55)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int64(0)
						} else {
							if v57 != int64(0) {
								v115 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
								v125 = int32(0)
								return base.I64_extend_i32_u(v125)
							} else {
								v86 = v55
								v87 = int32(1)
								v90 = F_DirectFunctionCall2Coll(m, l1, int32(0), v42, v86)
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int64(0)
								} else {
									if v90 != int64(0) {
										if v44 != 0 {
											v115 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
											v125 = int32(0)
											return base.I64_extend_i32_u(v125)
										} else {
											v95 = F_DirectFunctionCall2Coll(m, l1, int32(0), v86, v41)
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return int64(0)
											} else {
												if v87&base.B2i32(v95 == int64(0)) != 0 {
													v115 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
													v125 = int32(0)
													return base.I64_extend_i32_u(v125)
												} else {
													return base.I64_extend_i32_u(base.B2i32(v95 != int64(0)))
												}
											}
										}
									} else {
										if v87|v44 == int32(0) {
											v125 = int32(1)
										} else {
											v115 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
											v125 = int32(0)
										}
										return base.I64_extend_i32_u(v125)
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
func Fn14341(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(34209793)
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+8)) = uint32(v10)
	v15 = int64(base.Ui64(v10) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+4)) = uint32(v15)
	v18 = *(*int32)(unsafe.Add(mBase, _c_Fn14341[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v18
	v21 = F_LockAcquire(m, v8, l2, l1, int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int64(0)
	} else {
		m.G0 = v8 + int32(16)
		return int64(0)
	}
}
func Fn14347(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_get_statisticsobj_worker(m, v4, l1, int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		if v6 == int32(0) {
			v12 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v12)
			return int64(0)
		} else {
			v16 = F_cstring_to_text(m, v6)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int64(0)
			} else {
				F_pfree(m, v6)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					return base.I64_extend_i32_u(v16)
				}
			}
		}
	}
}
func Fn14350(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(34209794)
	*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)) = uint32(v12)
	*(*uint32)(unsafe.Add(mBase, uint32(v9)+4)) = uint32(v11)
	v18 = *(*int32)(unsafe.Add(mBase, _c_Fn14350[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v18
	v21 = F_LockAcquire(m, v9, l2, l1, int32(1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int64(0)
	} else {
		m.G0 = v9 + int32(16)
		return base.I64_extend_i32_u(base.B2i32(v21 != int32(0)))
	}
}
func Fn14367(m *base.Module, l0 int32, l1 int64) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v14 int64
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v14 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v4+int32(base.Ui32(v8)>>(uint(int32(2))%32))-int32(1)))))
		return int64(base.Ui64(v14)>>(uint(l1)%64)) & int64(1)
	}
}
func Fn14372(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+48)))
	if base.Ui32(v11) <= base.Ui32(int32(15)) {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
		v15 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v16
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = v15
		F_appendStringInfo(m, l0, l2, v8)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			m.G0 = v8 + int32(16)
			return
		}
	} else {
		m.G0 = v8 + int32(16)
		return
	}
}
func Fn14374(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(255)
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0 & v8
	if base.Ui32((l0-int32(33))&v8) < base.Ui32(int32(94)) {
		v20 = int32(_a_Fn14374_0)
	} else {
		v20 = int32(_a_Fn14374_1)
	}
	v21 = F_pg_snprintf(m, l1, int32(5), v20, v6)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return
	} else {
		m.G0 = v6 + int32(16)
		return
	}
}
func Fn14387(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int64
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	v10 = m.G0
	v11 = int32(-64)
	v12 = v10 + v11
	m.G0 = v12
	v14 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
	v16 = v12 - v11
	v17 = v16
	v21 = v14
	goto L1
L1:
	;
	v27 = v17 - int32(1)
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v21)&l3)+uint32(_c_Fn14387[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v30)
	if base.Ui64(v21) < base.Ui64(l2) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v36 = v16 - v27
	v38 = v36 + int32(4)
	v39 = F_palloc(m, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L2
L4:
	;
	if base.Ui32(v12) < base.Ui32(v27) {
		v17 = v27
		v21 = int64(base.Ui64(v21) >> (uint(l1) % 64))
		goto L1
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	return int64(0)
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v38 << (uint(int32(2)) % 32)
	if v36 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	base.MemoryCopy(m, v39+int32(4), v27, v36)
	goto L10
L9:
	;
	goto L10
L10:
	;
	m.G0 = v12 - int32(-64)
	return base.I64_extend_i32_u(v39)
}
func Fn14392(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v2 = l1
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = F_palloc0(m, int32(36))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int64(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(446)
		*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = l7
		*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = l7
		v27 = F_list_make1_impl(m, int32(480), v13+int32(8))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = l6
			*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = l5
			*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = l4
			*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = l2
			*(*uint16)(unsafe.Add(mBase, uint32(v16)+8)) = uint16(v2)
			*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v27
			m.G0 = v13 + int32(16)
			return base.I64_extend_i32_u(v16)
		}
	}
}
func Fn14402(m *base.Module, l0 int32, l1 int64) int64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v43 int32
	_ = v43
	var v45 float32
	_ = v45
	var v47 float32
	_ = v47
	var v54 int32
	_ = v54
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
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
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+4)))
	v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+4)))
	v22 = base.B2i32(v20 < v21)
	if v20 < v21 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v20 < v21 {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	v23 = v20
	goto L7
L6:
	;
	v23 = v21
	goto L7
L7:
	;
	if v23 <= int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v26 = int32(8)
	v31 = int32(0)
	goto L9
L9:
	;
	v43 = v31 << (uint(int32(2)) % 32)
	v45 = *(*float32)(unsafe.Add(mBase, uint32(v13+v26+v43)))
	v47 = *(*float32)(unsafe.Add(mBase, uint32(v18+v26+v43)))
	if base.F32_lt(v45, v47) != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	return int64(1)
L11:
	;
	return l1
L12:
	;
	goto L13
L13:
	;
	if base.F32_gt(v45, v47) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v54 = v31 + int32(1)
	if v54 == v23 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	goto L10
L17:
	;
	v31 = v54
	goto L9
L18:
	;
	return l1
L19:
	;
	goto L20
L20:
	;
	return base.I64_extend_i32_u(base.B2i32(v21 < v20))
}
func Fn14404(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_WinCheckAndInitializeNullTreatment(m, v10, int32(1), l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v20 = F_WinGetFuncArgInFrame(m, v10, int32(0), l1, int32(1), v8+int32(15))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
			if v22 == int32(1) {
				v25 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v25)
				v28 = int64(0)
			} else {
				v28 = v20
			}
			m.G0 = v8 + int32(16)
			return v28
		}
	}
}
