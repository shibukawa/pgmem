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
					v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+24)))
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
func Fn13823(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if int32(0) < l0 {
		if base.Ui32(l7) <= base.Ui32(l0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				F_errcode(m, int32(261))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = l4
					F_errmsg(m, l3, v12)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						F_errfinish(m, l2, l1, int32(_a_Fn13823_0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
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
			m.G0 = v12 + int32(16)
			return
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			F_errcode(m, int32(130))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				F_errmsg(m, l6, int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					F_errfinish(m, l2, l5, int32(_a_Fn13823_0))
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
		}
	}
}
func Fn13830(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+6)))
	v12 = int32(2)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v6+v8*int32(0)<<(uint(v12)%32)+l1<<(uint(v12)%32)-int32(4))))
	if v20 == int32(0) {
		v33 = l2
		return v33
	} else {
		v24 = F_index_getprocinfo(m, l0, int32(1), l1)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			if v24 == int32(0) {
				v33 = l2
				return v33
			} else {
				v30 = F_FunctionCall0Coll(m, v24)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v33 = v30
					return v33
				}
			}
		}
	}
}
func Fn13843(m *base.Module, l0 int32, l1 int32) int32 {
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = F_pg_detoast_datum(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v15 = F_array_iterator(m, v7, l1, v12, int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v17 != v7 {
					F_pfree(m, v7)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v21 != v12 {
							F_pfree(m, v12)
							mBase = m.M
							v24 = m.ExcPending
							if v24 != 0 {
								return int32(0)
							} else {
								return v15
							}
						} else {
							return v15
						}
					}
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v21 != v12 {
						F_pfree(m, v12)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							return v15
						}
					} else {
						return v15
					}
				}
			}
		}
	}
}
func Fn13849(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v6 = int32(*(*uint8)(unsafe.Add(mBase, _c_Fn13849[0])))
	if v6 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33685829))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_Fn13849_0), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_Fn13849_1), l3, l2)
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
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v25
		return int32(0)
	}
}
func Fn13850(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
					F_errmsg_internal(m, int32(_a_Fn13850_0), int32(0))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(11)
						F_errdetail_internal(m, int32(_a_Fn13850_1), v11)
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
func Fn13856(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = int32(1)
		v13 = v8 + v12
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
		v18 = v16 & v12
		if v18 != 0 {
			v19 = v13
		} else {
			v19 = v8 + int32(4)
		}
		if v16 == int32(1) {
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			if v25 == int32(18) {
				v28 = int32(16)
			} else {
				v28 = int32(0)
			}
			if base.Ui32((v25-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v35 = int32(4)
			} else {
				v35 = v28
			}
			v46 = v35
		} else {
			v36 = int32(1)
			if v18 != 0 {
				v46 = int32(base.Ui32(v16)>>(uint(v36)%32)) - v36
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v49 = F_dotrim(m, v19, v46, int32(_a_Fn13856_0), int32(1), l2, l1)
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return int32(0)
		} else {
			return v49
		}
	}
}
func Fn13867(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v7 == int32(-2147483648) {
		v18 = int64(-9223372036854775807 - 1)
		v19 = F_Int64GetDatum(m, v18)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = F_DirectFunctionCall2Coll(m, l1, int32(0), v19, v4)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				return v23
			}
		}
	} else {
		if v7 == int32(2147483647) {
			v18 = int64(9223372036854775807)
			v19 = F_Int64GetDatum(m, v18)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = F_DirectFunctionCall2Coll(m, l1, int32(0), v19, v4)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					return v23
				}
			}
		} else {
			if int32(106751983) <= v7 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_Fn13867_0), int32(0))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_Fn13867_1), int32(658), int32(_a_Fn13867_2))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
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
				v18 = base.I64_extend_i32_s(v7) * int64(86400000000)
				v19 = F_Int64GetDatum(m, v18)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v23 = F_DirectFunctionCall2Coll(m, l1, int32(0), v19, v4)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						return v23
					}
				}
			}
		}
	}
}
func Fn13872(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	v14 = int32(_a_Fn13872_0)
	v16 = *(*int32)(unsafe.Add(mBase, _c_Fn13872[0]))
	*(*int32)(unsafe.Add(mBase, _c_Fn13872[0])) = v16 + int32(1)
	v21 = *(*int32)(unsafe.Add(mBase, _c_Fn13872[1]))
	if int32(0) <= v21 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = int32(_a_Fn13872_1)
	v25 = *(*int32)(unsafe.Add(mBase, _c_Fn13872[2]))
	v28 = v21 * int32(100)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_Fn13872[3])))
	*(*int32)(unsafe.Add(mBase, _c_Fn13872[2])) = v31
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
	*(*int32)(unsafe.Add(mBase, _c_Fn13872[1])) = int32(-1)
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
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_Fn13872[4])))
	*(*int32)(unsafe.Add(mBase, _c_Fn13872[5])) = v38
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
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_Fn13872[6])))
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
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_Fn13872[4])))
	*(*int32)(unsafe.Add(mBase, _c_Fn13872[5])) = v57
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
	*(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_Fn13872[6]))) = v75
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
	*(*int32)(unsafe.Add(mBase, _c_Fn13872[2])) = v25
	v83 = int32(_a_Fn13872_0)
	v85 = *(*int32)(unsafe.Add(mBase, _c_Fn13872[0]))
	*(*int32)(unsafe.Add(mBase, _c_Fn13872[0])) = v85 - int32(1)
	m.G0 = v12 + int32(32)
	return
L21:
	;
	F_errmsg_internal(m, int32(_a_Fn13872_2), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_Fn13872_3), l3, l2)
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
func Fn13887(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v15 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v14 + v15
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v14
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+16)))
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21+v22)+12)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = F_gbt_num_distance(m, v8, v8+v15, v24&int32(1), l1, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		return int32(0)
	} else {
		v32 = F_Float8GetDatum(m, v28)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			m.G0 = v8 + int32(16)
			return v32
		}
	}
}
func Fn13892(m *base.Module, l0 int32, l1 int32) int32 {
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_palloc0(m, int32(16))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(16)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v15 = F_gbt_num_union(m, v7, v5, l1, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v15
		}
	}
}
func Fn13894(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = F_pg_detoast_datum(m, l1)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v19 = v10 + int32(8)
			v21 = int32(4)
			v22 = v12 + v21
			*(*int32)(unsafe.Add(mBase, uint32(v19))) = v22
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
			v25 = int32(2)
			v26 = int32(base.Ui32(v24) >> (uint(v25) % 32))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			if base.Ui32(v26+v21) < base.Ui32(int32(base.Ui32(v34)>>(uint(v25)%32))) {
				v38 = v22 + (v26+int32(3))&int32(2147483644)
			} else {
				v38 = v22
			}
			*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v38
			v41 = int32(4)
			v42 = v16 + v41
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v42
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
			v45 = int32(2)
			v46 = int32(base.Ui32(v44) >> (uint(v45) % 32))
			v54 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			if base.Ui32(v46+v41) < base.Ui32(int32(base.Ui32(v54)>>(uint(v45)%32))) {
				v58 = v42 + (v46+int32(3))&int32(2147483644)
			} else {
				v58 = v42
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v58
			v60 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			v61 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
			v62 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			v63 = F_DirectFunctionCall2Coll(m, l3, v60, v61, v62)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				if l0 != v12 {
					F_pfree(m, v12)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						if l1 != v16 {
							F_pfree(m, v16)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								m.G0 = v10 + int32(16)
								return v63
							}
						} else {
							m.G0 = v10 + int32(16)
							return v63
						}
					}
				} else {
					if l1 != v16 {
						F_pfree(m, v16)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							m.G0 = v10 + int32(16)
							return v63
						}
					} else {
						m.G0 = v10 + int32(16)
						return v63
					}
				}
			}
		}
	}
}
func Fn13900(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = F_SearchSysCache1(m, l1, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13)+96))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func Fn13911(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 float32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = F_pg_detoast_datum(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
			v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)))
			if v18 != v19 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+4)))
						v29 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v29
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v28
						F_errmsg(m, int32(_a_Fn13911_0), v8)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_Fn13911_1), int32(80), int32(_a_Fn13911_2))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
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
				v41 = int32(8)
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v46 = m.T0[v45].(func(*base.Module, int32, int32, int32) float32)(m, base.I32_extend16_s(v18), v11+v41, v16+v41)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					v49 = F_Float8GetDatum(m, base.F64_promote_f32(v46))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(16)
						return v49
					}
				}
			}
		}
	}
}
func Fn13917(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = F_pg_detoast_datum_packed(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v23 = F_pg_detoast_datum_packed(m, v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, _c_Fn13917[0]))
			v28 = F_text_to_cstring(m, v18)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = F_DirectFunctionCall1Coll(m, l7, int32(0), v28)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					if v30 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							F_errcode(m, l6)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v15))) = v28
								F_errmsg(m, l5, v15)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_Fn13917_0), l4, l3)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
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
						v46 = F_convert_any_priv_string(m, v23, l1)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							v48 = F_object_aclcheck(m, l2, v30, v26, v46)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								m.G0 = v15 + int32(16)
								return base.B2i32(v48 == int32(0))
							}
						}
					}
				}
			}
		}
	}
}
func Fn13922(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if int32(0) < l1 {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+18)))
		if base.Ui32(v16&int32(2047)) < base.Ui32(l1) {
			v20 = F_getmissingattr(m, l2, l1, l3)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v93 = v20
				m.G0 = v11 + int32(16)
				return v93
			}
		} else {
			v24 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v24)
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+20)))
			if v27&int32(1) == v24 {
				v32 = int32(4)
				v36 = l2 + l1<<(uint(v32)%32) + v32
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
				if int32(0) <= v37 {
					v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+22)))
					v42 = v26 + v40 + v37
					v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+6)))
					if v43 != int32(1) {
						v93 = v42
						m.G0 = v11 + int32(16)
						return v93
					} else {
						v46 = int32(*(*int16)(unsafe.Add(mBase, uint32(v36)+4)))
						switch v46&int32(_a_Fn13922_0) - int32(1) {
						case 0:
							v51 = int32(*(*int8)(unsafe.Add(mBase, uint32(v42))))
							v93 = v51
							m.G0 = v11 + int32(16)
							return v93
						case 1:
							v52 = int32(*(*int16)(unsafe.Add(mBase, uint32(v42))))
							v93 = v52
							m.G0 = v11 + int32(16)
							return v93
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = v46
								F_errmsg_internal(m, int32(_a_Fn13922_1), v11)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, l4, int32(70), int32(_a_Fn13922_2))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						case 3:
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
							v93 = v53
							m.G0 = v11 + int32(16)
							return v93
						}
					}
				} else {
					v66 = F_nocachegetattr(m, l0, l1, l2)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						v93 = v66
						m.G0 = v11 + int32(16)
						return v93
					}
				}
			} else {
				v68 = int32(1)
				v69 = l1 - v68
				v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+int32(base.Ui32(v69)>>(uint(int32(3))%32)))+23)))
				if int32(base.Ui32(v73)>>(uint(v69&int32(7))%32))&v68 == int32(0) {
					v81 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v81)
					v93 = int32(0)
					m.G0 = v11 + int32(16)
					return v93
				} else {
					v84 = F_nocachegetattr(m, l0, l1, l2)
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return int32(0)
					} else {
						v93 = v84
						m.G0 = v11 + int32(16)
						return v93
					}
				}
			}
		}
	} else {
		v86 = F_heap_getsysattr(m, l0, l1, l3)
		mBase = m.M
		v87 = m.ExcPending
		if v87 != 0 {
			return int32(0)
		} else {
			v93 = v86
			m.G0 = v11 + int32(16)
			return v93
		}
	}
}
func Fn13924(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int64) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v67 int64
	_ = v67
	v17 = F_palloc0(m, int32(140))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v21 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v21
		*(*int64)(unsafe.Add(mBase, uint32(v17)+4)) = l14
		*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(438)
		*(*uint16)(unsafe.Add(mBase, uint32(v17)+16)) = uint16(v21)
		*(*int32)(unsafe.Add(mBase, uint32(v17)+19)) = v21
		v30 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v17)+18)) = uint8(v30)
		*(*uint16)(unsafe.Add(mBase, uint32(v17)+23)) = uint16(v21)
		*(*int32)(unsafe.Add(mBase, uint32(v17)+108)) = l13
		*(*int32)(unsafe.Add(mBase, uint32(v17)+104)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v17)+100)) = l12
		*(*int32)(unsafe.Add(mBase, uint32(v17)+96)) = l11
		*(*int32)(unsafe.Add(mBase, uint32(v17)+92)) = l10
		*(*int32)(unsafe.Add(mBase, uint32(v17)+88)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v17)+84)) = l9
		*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = l8
		*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = l7
		*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = l6
		*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v17)+56)) = l5
		*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v21
		*(*uint8)(unsafe.Add(mBase, uint32(v17)+29)) = uint8(v30)
		*(*int32)(unsafe.Add(mBase, uint32(v17)+25)) = v30
		*(*int32)(unsafe.Add(mBase, uint32(v17)+136)) = v21
		v67 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v17)+128)) = v67
		*(*int64)(unsafe.Add(mBase, uint32(v17)+120)) = v67
		*(*int64)(unsafe.Add(mBase, uint32(v17)+112)) = v67
		return v17
	}
}
func Fn13937(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	if l0 == int32(0) {
		v11 = F_palloc(m, int32(32))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(4)
			*(*int64)(unsafe.Add(mBase, uint32(v11))) = l2
			v19 = v11 + int32(16)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v19
			v75 = v11
			v76 = v19
			v77 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v76+v77<<(uint(int32(2))%32)-int32(4)))) = l1
			return v75
		}
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v23 <= v22 {
			v25 = int32(1)
			v27 = int32(16)
			v29 = v22 + v25
			if v29 <= v27 {
				v32 = v27
			} else {
				v32 = v29
			}
			if v32&(v32-int32(1)) != 0 {
				v39 = v25 << (uint(int32(32)-base.I32_clz(v32)) % 32)
			} else {
				v39 = v32
			}
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v42 = l0 + int32(16)
			if v40 == v42 {
				v44 = F_GetMemoryChunkContext(m, l0)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					v48 = F_MemoryContextAlloc(m, v44, v39<<(uint(int32(2))%32))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v48
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v53 = v51 << (uint(int32(2)) % 32)
						if v53 == int32(0) {
						} else {
							base.MemoryCopy(m, v48, v42, v53)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v39
						v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v70 = v65
						v72 = v70 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v72
						v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v75 = l0
						v76 = v74
						v77 = v72
						*(*int32)(unsafe.Add(mBase, uint32(v76+v77<<(uint(int32(2))%32)-int32(4)))) = l1
						return v75
					}
				}
			} else {
				v59 = F_repalloc(m, v40, v39<<(uint(int32(2))%32))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v59
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v39
					v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v70 = v65
					v72 = v70 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v72
					v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v75 = l0
					v76 = v74
					v77 = v72
					*(*int32)(unsafe.Add(mBase, uint32(v76+v77<<(uint(int32(2))%32)-int32(4)))) = l1
					return v75
				}
			}
		} else {
			v70 = v22
			v72 = v70 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v72
			v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v75 = l0
			v76 = v74
			v77 = v72
			*(*int32)(unsafe.Add(mBase, uint32(v76+v77<<(uint(int32(2))%32)-int32(4)))) = l1
			return v75
		}
	}
}
func Fn13939(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v8 = int32(16711935)
	v10 = int32(8)
	v12 = int32(24)
	v16 = base.I32_rotr(v7&v8, v10) | base.I32_rotr(v7, v12)&v8
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v27 = base.I32_rotr(v18&v8, v10) | base.I32_rotr(v18, v12)&v8
	if base.Ui32(v16) < base.Ui32(v27) {
		v55 = l1
	} else {
		if base.Ui32(v27) < base.Ui32(v16) {
			v55 = int32(1)
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			v32 = int32(16711935)
			v34 = int32(8)
			v36 = int32(24)
			v40 = base.I32_rotr(v31&v32, v34) | base.I32_rotr(v31, v36)&v32
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
			v50 = base.I32_rotr(v41&v32, v34) | base.I32_rotr(v41, v36)&v32
			if base.Ui32(v40) < base.Ui32(v50) {
				v55 = l1
			} else {
				v55 = base.B2i32(base.Ui32(v50) < base.Ui32(v40))
			}
		}
	}
	return v55
}
func Fn13944(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
		if v18 == int32(1) {
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
			if v24 == int32(18) {
				v27 = int32(16)
			} else {
				v27 = int32(0)
			}
			if base.Ui32((v24-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v34 = int32(4)
			} else {
				v34 = v27
			}
			v47 = v34
		} else {
			v35 = int32(1)
			if v18&v35 != 0 {
				v47 = int32(base.Ui32(v18)>>(uint(v35)%32)) - v35
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v47 = int32(base.Ui32(v41)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v48 = int32(1)
		if v18&v48 != 0 {
			v52 = v48
		} else {
			v52 = int32(4)
		}
		v58 = F_pg_md5_hash(m, v12+v52, v47, v7+int32(-48), v7+int32(-52))
		mBase = m.M
		v59 = m.ExcPending
		if v59 != 0 {
			return int32(0)
		} else {
			if v58 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(2600))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(_a_Fn13944_0)
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v71
						F_errmsg(m, int32(_a_Fn13944_1), v9)
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_Fn13944_2), l2, l1)
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
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
				v81 = F_cstring_to_text(m, v7+int32(-48))
				mBase = m.M
				v82 = m.ExcPending
				if v82 != 0 {
					return int32(0)
				} else {
					m.G0 = v9 - int32(-64)
					return v81
				}
			}
		}
	}
}
func Fn13953(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
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
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v105 int32
	_ = v105
	var v115 int32
	_ = v115
	v13 = int32(1)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v20 == v13 {
		if v16&int32(1) == int32(0) {
			v40 = v19
			v41 = v19
			v42 = v13
			if v15&int32(1) != 0 {
				if v14&int32(1) == int32(0) {
					v53 = v17
					v55 = F_DirectFunctionCall2Coll(m, l2, int32(0), v41, v53)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						if v55 != 0 {
							v105 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
							v115 = int32(0)
							return v115
						} else {
							v79 = v53
							v80 = int32(1)
							v83 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v79)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								if v83 != 0 {
									if v42 != 0 {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v115 = int32(0)
										return v115
									} else {
										v86 = F_DirectFunctionCall2Coll(m, l1, int32(0), v79, v40)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int32(0)
										} else {
											if v80&base.B2i32(v86 == int32(0)) != 0 {
												v105 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
												v115 = int32(0)
												return v115
											} else {
												return base.B2i32(v86 != int32(0))
											}
										}
									}
								} else {
									if v80|v42 == int32(0) {
										v115 = int32(1)
									} else {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v115 = int32(0)
									}
									return v115
								}
							}
						}
					}
				} else {
					v105 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
					v115 = int32(0)
					return v115
				}
			} else {
				if v14&int32(1) == int32(0) {
					v58 = int32(0)
					v61 = F_DirectFunctionCall2Coll(m, l2, v58, v18, v17)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						if v61 != 0 {
							v63 = v17
						} else {
							v63 = v18
						}
						v64 = F_DirectFunctionCall2Coll(m, l2, v58, v41, v63)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							if v64 == int32(0) {
								v79 = v63
								v80 = v58
								v83 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v79)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									if v83 != 0 {
										if v42 != 0 {
											v105 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
											v115 = int32(0)
											return v115
										} else {
											v86 = F_DirectFunctionCall2Coll(m, l1, int32(0), v79, v40)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int32(0)
											} else {
												if v80&base.B2i32(v86 == int32(0)) != 0 {
													v105 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
													v115 = int32(0)
													return v115
												} else {
													return base.B2i32(v86 != int32(0))
												}
											}
										}
									} else {
										if v80|v42 == int32(0) {
											v115 = int32(1)
										} else {
											v105 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
											v115 = int32(0)
										}
										return v115
									}
								}
							} else {
								if v61 != 0 {
									v69 = v18
								} else {
									v69 = v17
								}
								v70 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v69)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									if v42&base.B2i32(v70 == int32(0)) != 0 {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v115 = int32(0)
										return v115
									} else {
										return base.B2i32(v70 != int32(0))
									}
								}
							}
						}
					}
				} else {
					v53 = v18
					v55 = F_DirectFunctionCall2Coll(m, l2, int32(0), v41, v53)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						if v55 != 0 {
							v105 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
							v115 = int32(0)
							return v115
						} else {
							v79 = v53
							v80 = int32(1)
							v83 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v79)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								if v83 != 0 {
									if v42 != 0 {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v115 = int32(0)
										return v115
									} else {
										v86 = F_DirectFunctionCall2Coll(m, l1, int32(0), v79, v40)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int32(0)
										} else {
											if v80&base.B2i32(v86 == int32(0)) != 0 {
												v105 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
												v115 = int32(0)
												return v115
											} else {
												return base.B2i32(v86 != int32(0))
											}
										}
									}
								} else {
									if v80|v42 == int32(0) {
										v115 = int32(1)
									} else {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v115 = int32(0)
									}
									return v115
								}
							}
						}
					}
				}
			}
		} else {
			v105 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
			v115 = int32(0)
			return v115
		}
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v16&int32(1) != 0 {
			v40 = v19
			v41 = v27
			v42 = v13
			if v15&int32(1) != 0 {
				if v14&int32(1) == int32(0) {
					v53 = v17
					v55 = F_DirectFunctionCall2Coll(m, l2, int32(0), v41, v53)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						if v55 != 0 {
							v105 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
							v115 = int32(0)
							return v115
						} else {
							v79 = v53
							v80 = int32(1)
							v83 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v79)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								if v83 != 0 {
									if v42 != 0 {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v115 = int32(0)
										return v115
									} else {
										v86 = F_DirectFunctionCall2Coll(m, l1, int32(0), v79, v40)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int32(0)
										} else {
											if v80&base.B2i32(v86 == int32(0)) != 0 {
												v105 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
												v115 = int32(0)
												return v115
											} else {
												return base.B2i32(v86 != int32(0))
											}
										}
									}
								} else {
									if v80|v42 == int32(0) {
										v115 = int32(1)
									} else {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v115 = int32(0)
									}
									return v115
								}
							}
						}
					}
				} else {
					v105 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
					v115 = int32(0)
					return v115
				}
			} else {
				if v14&int32(1) == int32(0) {
					v58 = int32(0)
					v61 = F_DirectFunctionCall2Coll(m, l2, v58, v18, v17)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						if v61 != 0 {
							v63 = v17
						} else {
							v63 = v18
						}
						v64 = F_DirectFunctionCall2Coll(m, l2, v58, v41, v63)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							if v64 == int32(0) {
								v79 = v63
								v80 = v58
								v83 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v79)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									if v83 != 0 {
										if v42 != 0 {
											v105 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
											v115 = int32(0)
											return v115
										} else {
											v86 = F_DirectFunctionCall2Coll(m, l1, int32(0), v79, v40)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int32(0)
											} else {
												if v80&base.B2i32(v86 == int32(0)) != 0 {
													v105 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
													v115 = int32(0)
													return v115
												} else {
													return base.B2i32(v86 != int32(0))
												}
											}
										}
									} else {
										if v80|v42 == int32(0) {
											v115 = int32(1)
										} else {
											v105 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
											v115 = int32(0)
										}
										return v115
									}
								}
							} else {
								if v61 != 0 {
									v69 = v18
								} else {
									v69 = v17
								}
								v70 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v69)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									if v42&base.B2i32(v70 == int32(0)) != 0 {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v115 = int32(0)
										return v115
									} else {
										return base.B2i32(v70 != int32(0))
									}
								}
							}
						}
					}
				} else {
					v53 = v18
					v55 = F_DirectFunctionCall2Coll(m, l2, int32(0), v41, v53)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						if v55 != 0 {
							v105 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
							v115 = int32(0)
							return v115
						} else {
							v79 = v53
							v80 = int32(1)
							v83 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v79)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								if v83 != 0 {
									if v42 != 0 {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v115 = int32(0)
										return v115
									} else {
										v86 = F_DirectFunctionCall2Coll(m, l1, int32(0), v79, v40)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int32(0)
										} else {
											if v80&base.B2i32(v86 == int32(0)) != 0 {
												v105 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
												v115 = int32(0)
												return v115
											} else {
												return base.B2i32(v86 != int32(0))
											}
										}
									}
								} else {
									if v80|v42 == int32(0) {
										v115 = int32(1)
									} else {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v115 = int32(0)
									}
									return v115
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
				return int32(0)
			} else {
				if v32 != 0 {
					v36 = v19
				} else {
					v36 = v27
				}
				if v32 != 0 {
					v37 = v27
				} else {
					v37 = v19
				}
				v40 = v37
				v41 = v36
				v42 = v30
				if v15&int32(1) != 0 {
					if v14&int32(1) == int32(0) {
						v53 = v17
						v55 = F_DirectFunctionCall2Coll(m, l2, int32(0), v41, v53)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							if v55 != 0 {
								v105 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
								v115 = int32(0)
								return v115
							} else {
								v79 = v53
								v80 = int32(1)
								v83 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v79)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									if v83 != 0 {
										if v42 != 0 {
											v105 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
											v115 = int32(0)
											return v115
										} else {
											v86 = F_DirectFunctionCall2Coll(m, l1, int32(0), v79, v40)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int32(0)
											} else {
												if v80&base.B2i32(v86 == int32(0)) != 0 {
													v105 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
													v115 = int32(0)
													return v115
												} else {
													return base.B2i32(v86 != int32(0))
												}
											}
										}
									} else {
										if v80|v42 == int32(0) {
											v115 = int32(1)
										} else {
											v105 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
											v115 = int32(0)
										}
										return v115
									}
								}
							}
						}
					} else {
						v105 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
						v115 = int32(0)
						return v115
					}
				} else {
					if v14&int32(1) == int32(0) {
						v58 = int32(0)
						v61 = F_DirectFunctionCall2Coll(m, l2, v58, v18, v17)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							if v61 != 0 {
								v63 = v17
							} else {
								v63 = v18
							}
							v64 = F_DirectFunctionCall2Coll(m, l2, v58, v41, v63)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								if v64 == int32(0) {
									v79 = v63
									v80 = v58
									v83 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v79)
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return int32(0)
									} else {
										if v83 != 0 {
											if v42 != 0 {
												v105 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
												v115 = int32(0)
												return v115
											} else {
												v86 = F_DirectFunctionCall2Coll(m, l1, int32(0), v79, v40)
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return int32(0)
												} else {
													if v80&base.B2i32(v86 == int32(0)) != 0 {
														v105 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
														v115 = int32(0)
														return v115
													} else {
														return base.B2i32(v86 != int32(0))
													}
												}
											}
										} else {
											if v80|v42 == int32(0) {
												v115 = int32(1)
											} else {
												v105 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
												v115 = int32(0)
											}
											return v115
										}
									}
								} else {
									if v61 != 0 {
										v69 = v18
									} else {
										v69 = v17
									}
									v70 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v69)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										if v42&base.B2i32(v70 == int32(0)) != 0 {
											v105 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
											v115 = int32(0)
											return v115
										} else {
											return base.B2i32(v70 != int32(0))
										}
									}
								}
							}
						}
					} else {
						v53 = v18
						v55 = F_DirectFunctionCall2Coll(m, l2, int32(0), v41, v53)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							if v55 != 0 {
								v105 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
								v115 = int32(0)
								return v115
							} else {
								v79 = v53
								v80 = int32(1)
								v83 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v79)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									if v83 != 0 {
										if v42 != 0 {
											v105 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
											v115 = int32(0)
											return v115
										} else {
											v86 = F_DirectFunctionCall2Coll(m, l1, int32(0), v79, v40)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int32(0)
											} else {
												if v80&base.B2i32(v86 == int32(0)) != 0 {
													v105 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
													v115 = int32(0)
													return v115
												} else {
													return base.B2i32(v86 != int32(0))
												}
											}
										}
									} else {
										if v80|v42 == int32(0) {
											v115 = int32(1)
										} else {
											v105 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
											v115 = int32(0)
										}
										return v115
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
func Fn13955(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(34209794)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v11
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v10
	v17 = *(*int32)(unsafe.Add(mBase, _c_Fn13955[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v17
	v20 = F_LockAcquire(m, v8, l2, l1, int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		m.G0 = v8 + int32(16)
		return int32(0)
	}
}
func Fn13964(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	v4 = int32(_a_Fn13964_0)
	v5 = int32(_a_Fn13964_1)
	v6 = *(*int64)(unsafe.Add(mBase, _c_Fn13964[0]))
	v8 = *(*int64)(unsafe.Add(mBase, _c_Fn13964[1]))
	v9 = v6 ^ v8
	*(*int64)(unsafe.Add(mBase, _c_Fn13964[1])) = base.I64_rotl(v9, int64(37))
	*(*int64)(unsafe.Add(mBase, _c_Fn13964[0])) = v9<<(uint(int64(16))%64) ^ base.I64_rotl(v6, int64(24)) ^ v9
	return base.I32_wrap_i64(int64(base.Ui64(base.I64_rotl(v6*int64(5), int64(7))*int64(9)) >> (uint(l0) % 64)))
}
func Fn13973(m *base.Module, l0 int32, l1 int64, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int64
	_ = v13
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v27 int64
	_ = v27
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v52 int64
	_ = v52
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int64
	_ = v166
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	v4 = int32(0)
	v13 = int64(2)
	if base.Ui64(l1) <= base.Ui64(v13) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = v13
	goto L3
L2:
	;
	v16 = l1
	goto L3
L3:
	;
	v17 = int64(1)
	if v16&(v16-v17) == int64(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v27 = v16
	goto L6
L5:
	;
	v27 = v17 << (uint(int64(64)-base.I64_clz(v16)) % 64)
	goto L6
L6:
	;
	if base.Ui64(v27<<(uint(int64(3))%64)) < base.Ui64(int64(2147483647)) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v39 = F_MemoryContextAllocExtended(m, v34, base.I32_wrap_i64(v27)<<(uint(int32(3))%32), int32(5))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L11
	} else {
		goto L42
	}
L10:
	;
	goto L9
L11:
	;
	return
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v39
	v42 = int64(1)
	if v27&(v27-v42) == int64(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v52 = v27
	goto L15
L14:
	;
	v52 = v42 << (uint(int64(64)-base.I64_clz(v27)) % 64)
	goto L15
L15:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v52<<(uint(int64(3))%64)) {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v52
	v60 = base.I32_wrap_i64(v52) - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v60
	if v52 == int64(4294967296) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v69 = int32(-85899346)
	goto L19
L18:
	;
	v69 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i64_u(v52), float64(0.9)))
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v69
	if v33 != int64(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v77 = v4
	goto L24
L21:
	;
	goto L22
L22:
	;
	F_pfree(m, v32)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L11
	} else {
		goto L41
	}
L23:
	;
	v118 = v113
	v124 = v4
	goto L29
L24:
	;
	v87 = v32 + v77<<(uint(int32(3))%32)
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+4)))
	if v88 != int32(1) {
		v113 = v77
		goto L23
	} else {
		goto L26
	}
L25:
	;
	v113 = int32(0)
	goto L23
L26:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v92 = int32(16)
	v96 = (int32(base.Ui32(v91)>>(uint(v92)%32)) ^ v91) * int32(-2048144789)
	v101 = (int32(base.Ui32(v96)>>(uint(int32(13))%32)) ^ v96) * int32(-1028477387)
	if (int32(base.Ui32(v101)>>(uint(v92)%32))^v101)&v60 == v77 {
		v113 = v77
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v108 = v77 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v108)) < base.Ui64(v33) {
		v77 = v108
		goto L24
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	v128 = v32 + v118<<(uint(int32(3))%32)
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+4)))
	if v129 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L22
L31:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	v133 = int32(16)
	v137 = (int32(base.Ui32(v132)>>(uint(v133)%32)) ^ v132) * int32(-2048144789)
	v142 = (int32(base.Ui32(v137)>>(uint(int32(13))%32)) ^ v137) * int32(-1028477387)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v150 = int32(base.Ui32(v142)>>(uint(v133)%32)) ^ v142
	goto L34
L32:
	;
	goto L33
L33:
	;
	v181 = v118 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v181)) < base.Ui64(v33) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v159 = v150 & v146
	v164 = v39 + v159<<(uint(int32(3))%32)
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+4)))
	if v165 != 0 {
		v150 = v159 + int32(1)
		goto L34
	} else {
		goto L36
	}
L35:
	;
	v166 = *(*int64)(unsafe.Add(mBase, uint32(v128)))
	*(*int64)(unsafe.Add(mBase, uint32(v164))) = v166
	goto L33
L36:
	;
	goto L35
L37:
	;
	v185 = v181
	goto L39
L38:
	;
	v185 = int32(0)
	goto L39
L39:
	;
	v187 = v124 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v187)) < base.Ui64(v33) {
		v118 = v185
		v124 = v187
		goto L29
	} else {
		goto L40
	}
L40:
	;
	goto L30
L41:
	;
	return
L42:
	;
	F_errmsg_internal(m, int32(_a_Fn13973_0), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L11
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_Fn13973_1), int32(327), l2)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L11
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func Fn13979(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
func Fn13980(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	v5 = int32(0)
	v7 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v7 == v5 {
		v28 = v5
		return v28
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v11-int32(2) <= v10 {
			v28 = v5
			return v28
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v11-int32(1)))))
			if v19 != l3 {
				v28 = v5
				return v28
			} else {
				v21 = F_find_among_b(m, l0, l2, l1)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v28 = base.B2i32(v21 != int32(0))
					return v28
				}
			}
		}
	}
}
func Fn13991(m *base.Module, l0 int32, l1 int32) {
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
		v20 = int32(_a_Fn13991_0)
	} else {
		v20 = int32(_a_Fn13991_1)
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
func Fn13997(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 float64
	_ = v16
	var v19 float64
	_ = v19
	var v22 float64
	_ = v22
	var v23 int64
	_ = v23
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v37 int64
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
	var v59 int64
	_ = v59
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	v9 = F_MemoryContextAllocZero(m, l0, int32(32))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = l0
		v16 = float64(4.294967296e+09)
		v19 = base.F64_div(base.F64_convert_i32_u(l1), float64(0.9))
		if base.F64_ge(v19, v16) != 0 {
			v22 = v16
		} else {
			v22 = v19
		}
		v23 = base.I64_trunc_sat_f64_u(v22)
		if base.Ui64(v23) <= base.Ui64(int64(2)) {
			v26 = int64(2)
		} else {
			v26 = v23
		}
		v27 = int64(1)
		if v26&(v26-v27) == int64(0) {
			v37 = v26
		} else {
			v37 = v27 << (uint(int64(64)-base.I64_clz(v26)) % 64)
		}
		if base.Ui64(v37<<(uint(int64(3))%64)) < base.Ui64(int64(2147483647)) {
			v46 = F_MemoryContextAllocExtended(m, l0, base.I32_wrap_i64(v37)<<(uint(int32(3))%32), int32(5))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v46
				v49 = int64(1)
				if v37&(v37-v49) == int64(0) {
					v59 = v37
				} else {
					v59 = v49 << (uint(int64(64)-base.I64_clz(v37)) % 64)
				}
				if base.Ui64(int64(2147483647)) <= base.Ui64(v59<<(uint(int64(3))%64)) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_Fn13997_0), int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_Fn13997_1), int32(327), l3)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v9))) = v59
					*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = base.I32_wrap_i64(v59) - int32(1)
					if v59 == int64(4294967296) {
						v76 = int32(-85899346)
					} else {
						v76 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i64_u(v59), float64(0.9)))
					}
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v76
					return v9
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_Fn13997_0), int32(0))
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_Fn13997_1), int32(327), l3)
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
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
func Fn14004(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32 {
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
	var v15 int64
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int64
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	v10 = m.G0
	v11 = int32(-64)
	v12 = v10 + v11
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v17 = v12 - v11
	v18 = v17
	v22 = v15
	goto L1
L1:
	;
	v28 = v18 - int32(1)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v22)&l3)+uint32(_c_Fn14004[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v31)
	if base.Ui64(v22) < base.Ui64(l2) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v37 = v17 - v28
	v39 = v37 + int32(4)
	v40 = F_palloc(m, v39)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L2
L4:
	;
	if base.Ui32(v12) < base.Ui32(v28) {
		v18 = v28
		v22 = int64(base.Ui64(v22) >> (uint(l1) % 64))
		goto L1
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	return int32(0)
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v39 << (uint(int32(2)) % 32)
	if v37 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	base.MemoryCopy(m, v40+int32(4), v28, v37)
	goto L10
L9:
	;
	goto L10
L10:
	;
	m.G0 = v12 - int32(-64)
	return v40
}
func Fn14008(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	v2 = l1
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum_copy(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v13 = F_pg_detoast_datum_copy(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			if v15 == int32(0) {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v8 != v18 {
					v76 = v13
					v78 = v8
					F_pfree(m, v78)
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return int32(0)
					} else {
						v81 = v76
						return v81
					}
				} else {
					v81 = v13
					return v81
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
				if v20 == int32(0) {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v23 != v13 {
						v76 = v8
						v78 = v13
						F_pfree(m, v78)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							v81 = v76
							return v81
						}
					} else {
						v81 = v8
						return v81
					}
				} else {
					v26 = F_palloc0(m, int32(24))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v28 | int32(1)
						v33 = F_palloc0(m, int32(12))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v26))) = v33
							v36 = int32(2)
							*(*uint8)(unsafe.Add(mBase, uint32(v33))) = uint8(v36)
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
							*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)) = uint8(v2)
							v41 = F_palloc0(m, int32(8))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v41
								v45 = v13 + int32(8)
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
								v50 = F_QT2QTN(m, v45, v45+v46*int32(12))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
									*(*int32)(unsafe.Add(mBase, uint32(v52))) = v50
									v55 = v8 + int32(8)
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
									v60 = F_QT2QTN(m, v55, v55+v56*int32(12))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										v62 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
										*(*int32)(unsafe.Add(mBase, uint32(v62)+4)) = v60
										*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = int32(2)
										v66 = F_QTN2QT(m, v26)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											F_QTNFree(m, v26)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int32(0)
											} else {
												v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												if v70 != v8 {
													F_pfree(m, v8)
													mBase = m.M
													v73 = m.ExcPending
													if v73 != 0 {
														return int32(0)
													} else {
														v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
														if v74 == v13 {
															v81 = v66
															return v81
														} else {
															v76 = v66
															v78 = v13
															F_pfree(m, v78)
															mBase = m.M
															v80 = m.ExcPending
															if v80 != 0 {
																return int32(0)
															} else {
																v81 = v76
																return v81
															}
														}
													}
												} else {
													v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													if v74 == v13 {
														v81 = v66
														return v81
													} else {
														v76 = v66
														v78 = v13
														F_pfree(m, v78)
														mBase = m.M
														v80 = m.ExcPending
														if v80 != 0 {
															return int32(0)
														} else {
															v81 = v76
															return v81
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
		}
	}
}
func Fn14013(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v13, v14, v15, int32(6), l1)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v23 = F_UtfToLocal(m, v11, v15, v10, l5, l4, l3, l2, l1, base.B2i32(v12 != int32(0)))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			return v23
		}
	}
}
func Fn14015(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = int32(1)
		v13 = v8 + v12
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
		v18 = v16 & v12
		if v18 != 0 {
			v19 = v13
		} else {
			v19 = v8 + int32(4)
		}
		if v16 == int32(1) {
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			if v25 == int32(18) {
				v28 = int32(16)
			} else {
				v28 = int32(0)
			}
			if base.Ui32((v25-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v35 = int32(4)
			} else {
				v35 = v28
			}
			v46 = v35
		} else {
			v36 = int32(1)
			if v18 != 0 {
				v46 = int32(base.Ui32(v16)>>(uint(v36)%32)) - v36
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v47 = F_uuid_generate_internal(m, l1, v6, v19, v46)
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return int32(0)
		} else {
			return v47
		}
	}
}
func Fn14019(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 float32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v17 = v10 + int32(1)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v19 = F_pg_detoast_datum_packed(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
			v23 = v21 & int32(1)
			if v23 != 0 {
				v24 = v17
			} else {
				v24 = v10 + int32(4)
			}
			if v21 == int32(1) {
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
				if v30 == int32(18) {
					v33 = int32(16)
				} else {
					v33 = int32(0)
				}
				if base.Ui32((v30-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v40 = int32(4)
				} else {
					v40 = v33
				}
				v51 = v40
			} else {
				v41 = int32(1)
				if v23 != 0 {
					v51 = int32(base.Ui32(v21)>>(uint(v41)%32)) - v41
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v52 = int32(1)
			v53 = v19 + v52
			v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
			v58 = v56 & v52
			if v58 != 0 {
				v59 = v53
			} else {
				v59 = v19 + int32(4)
			}
			if v56 == int32(1) {
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
				if v65 == int32(18) {
					v68 = int32(16)
				} else {
					v68 = int32(0)
				}
				if base.Ui32((v65-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v75 = int32(4)
				} else {
					v75 = v68
				}
				v86 = v75
			} else {
				v76 = int32(1)
				if v58 != 0 {
					v86 = int32(base.Ui32(v56)>>(uint(v76)%32)) - v76
				} else {
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
					v86 = int32(base.Ui32(v80)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v87 = F_calc_word_similarity(m, v24, v51, v59, v86, l1)
			mBase = m.M
			v88 = m.ExcPending
			if v88 != 0 {
				return int32(0)
			} else {
				v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v89 != v10 {
					F_pfree(m, v10)
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int32(0)
					} else {
						v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v93 != v19 {
							F_pfree(m, v19)
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return int32(0)
							} else {
								return base.I32_reinterpret_f32(v87)
							}
						} else {
							return base.I32_reinterpret_f32(v87)
						}
					}
				} else {
					v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v93 != v19 {
						F_pfree(m, v19)
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return int32(0)
						} else {
							return base.I32_reinterpret_f32(v87)
						}
					} else {
						return base.I32_reinterpret_f32(v87)
					}
				}
			}
		}
	}
}
func Fn14024(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
