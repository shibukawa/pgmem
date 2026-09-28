package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_fn_expr_rettype(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	v2 = int32(0)
	if l0 == v2 {
		v13 = v2
		return v13
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v5 == int32(0) {
			v13 = v2
			return v13
		} else {
			v8 = F_exprType(m, v5)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				v13 = v8
				return v13
			}
		}
	}
}
func Fn14207(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32, l5 int64, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
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
	var v51 int64
	_ = v51
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int64
	_ = v76
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v104 int32
	_ = v104
	v14 = m.G0
	v16 = v14 - int32(112)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	if v18 != 0 {
		v74 = v18
		v76 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v76
		*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v76
		*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = l8
		*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = l7
		v82 = v74
		v84 = int32(16)
		if v82 <= v84 {
			v87 = v84
		} else {
			v87 = v82
		}
		if l6 <= v87 {
			v89 = l6
		} else {
			v89 = v87
		}
		v92 = v89
		*(*int64)(unsafe.Add(mBase, uint32(v16)+68)) = l5
		*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = l2
		*(*int64)(unsafe.Add(mBase, uint32(v16)+44)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v92
		F_SimpleLruRequestWithOpts(m, v16+int32(16))
		mBase = m.M
		v104 = m.ExcPending
		if v104 != 0 {
			return
		} else {
			m.G0 = v16 + int32(112)
			return
		}
	} else {
		v21 = int32(16)
		v23 = *(*int32)(unsafe.Add(mBase, _c_Fn14207[0]))
		v25 = base.I32_div_s(v23, int32(512))
		v27 = base.I32_rem_s(v25, v21)
		v28 = v25 - v27
		if v28 <= v21 {
			v31 = v21
		} else {
			v31 = v28
		}
		if int32(1024) < v31 {
			v34 = int32(1024)
		} else {
			v34 = v31
		}
		*(*int32)(unsafe.Add(mBase, uint32(v16))) = v34
		v37 = v16 + int32(80)
		v40 = F_pg_snprintf(m, v37, int32(32), int32(_a_Fn14207_0), v16)
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return
		} else {
			v42 = int32(1)
			F_SetConfigOption(m, l10, v37, v42, v42)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return
			} else {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
				if v46 != 0 {
					v74 = v46
					v76 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v76
					*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v76
					*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = l8
					*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = l7
					v82 = v74
					v84 = int32(16)
					if v82 <= v84 {
						v87 = v84
					} else {
						v87 = v82
					}
					if l6 <= v87 {
						v89 = l6
					} else {
						v89 = v87
					}
					v92 = v89
					*(*int64)(unsafe.Add(mBase, uint32(v16)+68)) = l5
					*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = l4
					*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = l2
					*(*int64)(unsafe.Add(mBase, uint32(v16)+44)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v92
					F_SimpleLruRequestWithOpts(m, v16+int32(16))
					mBase = m.M
					v104 = m.ExcPending
					if v104 != 0 {
						return
					} else {
						m.G0 = v16 + int32(112)
						return
					}
				} else {
					F_SetConfigOption(m, l10, v37, int32(1), int32(10))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						v51 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v51
						*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v51
						*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = l8
						*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = l7
						v57 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
						if v57 != 0 {
							v82 = v57
							v84 = int32(16)
							if v82 <= v84 {
								v87 = v84
							} else {
								v87 = v82
							}
							if l6 <= v87 {
								v89 = l6
							} else {
								v89 = v87
							}
							v92 = v89
						} else {
							v60 = int32(16)
							v62 = *(*int32)(unsafe.Add(mBase, _c_Fn14207[0]))
							v64 = base.I32_div_s(v62, int32(512))
							v66 = base.I32_rem_s(v64, v60)
							v67 = v64 - v66
							if v67 <= v60 {
								v70 = v60
							} else {
								v70 = v67
							}
							if int32(1024) < v70 {
								v73 = int32(1024)
							} else {
								v73 = v70
							}
							v92 = v73
						}
						*(*int64)(unsafe.Add(mBase, uint32(v16)+68)) = l5
						*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = l4
						*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = l3
						*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = l2
						*(*int64)(unsafe.Add(mBase, uint32(v16)+44)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v92
						F_SimpleLruRequestWithOpts(m, v16+int32(16))
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return
						} else {
							m.G0 = v16 + int32(112)
							return
						}
					}
				}
			}
		}
	}
}
func Fn14212(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = l2
	v14 = v10 + int32(16)
	v16 = F_pg_snprintf(m, v14, int32(32), l4, v10)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		F_ExplainProperty(m, l0, l1, v14, int32(1), l3)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			m.G0 = v10 + int32(48)
			return
		}
	}
}
func Fn14214(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+6)))
	v12 = int32(2)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v6+v8*int32(0)<<(uint(v12)%32)+l1<<(uint(v12)%32)-int32(4))))
	if v20 == int32(0) {
		v34 = l2
		return v34
	} else {
		v24 = F_index_getprocinfo(m, l0, int32(1), l1)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			if v24 == int32(0) {
				v34 = l2
				return v34
			} else {
				v30 = F_FunctionCall0Coll(m, v24)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v34 = base.I32_wrap_i64(v30)
					return v34
				}
			}
		}
	}
}
func Fn14227(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
	v16 = F_query_or_expression_tree_mutator_impl(m, l0, l3, v8+int32(4))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		m.G0 = v8 + int32(16)
		return v16
	}
}
func Fn14232(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64 {
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
	var v25 int64
	_ = v25
	v6 = int32(*(*uint8)(unsafe.Add(mBase, _c_Fn14232[0])))
	if v6 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(33685829))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_Fn14232_0), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_Fn14232_1), l3, l2)
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
	} else {
		v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		*(*uint32)(unsafe.Add(mBase, uint32(l1))) = uint32(v25)
		return int64(0)
	}
}
func Fn14241(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
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
	v8 = l1 << (uint(l3) % 32)
	v10 = v8 + int32(24)
	v11 = F_palloc0(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = int32(0)
		if base.B2i32(v8 == v15)|(base.B2i32(l0 == v15)|base.B2i32(l1 <= v15)) == v15 {
			base.MemoryCopy(m, v11+int32(24), l0, v8)
		} else {
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l2
		*(*int64)(unsafe.Add(mBase, uint32(v11)+4)) = int64(1)
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = v10 << (uint(int32(2)) % 32)
		return v11
	}
}
func Fn14247(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = v8 + v11
	if base.B2i32(v8 < int64(0)) != base.B2i32(v12 < v11) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, l4, int32(0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, l3, l2, l1)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
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
		return v12
	}
}
func Fn14250(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l1
	v13 = F_query_or_expression_tree_walker_impl(m, l0, l2, v7+int32(12), int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v13
	}
}
func Fn14256(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 float64, l5 float64) int64 {
	mBase := m.M
	_ = mBase
	var v9 float64
	_ = v9
	var v10 float64
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
	v9 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = base.F64_nearest(v9)
	v17 = int32(0)
	if base.B2i32(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v10)&int64(9223372036854775807)))|base.B2i32(base.F64_ge(v10, l5) == v17) == v17)&base.F64_lt(v10, l4) == v17 {
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
						F_errsave_finish(m, v26, int32(_a_Fn14256_0), l2, l1)
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
		v47 = base.I64_extend_i32_s(base.I32_trunc_sat_f64_s(v10))
		return v47
	}
}
func Fn14263(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int64 {
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
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v76 int64
	_ = v76
	var v77 int32
	_ = v77
	var v82 int64
	_ = v82
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
			v76 = F_nocachegetattr(m, l0, l1, l2)
			mBase = m.M
			v77 = m.ExcPending
			if v77 != 0 {
				return int64(0)
			} else {
				v82 = v76
				m.G0 = v12 + int32(16)
				return v82
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
						F_errmsg_internal(m, int32(_a_Fn14263_0), v12)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, l4, int32(123), int32(_a_Fn14263_1))
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
						v82 = v41
						m.G0 = v12 + int32(16)
						return v82
					case 1:
						v42 = int64(*(*int16)(unsafe.Add(mBase, uint32(v32))))
						v82 = v42
						m.G0 = v12 + int32(16)
						return v82
					case 2:
						v43 = int64(*(*int32)(unsafe.Add(mBase, uint32(v32))))
						v82 = v43
						m.G0 = v12 + int32(16)
						return v82
					case 3:
						v44 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
						v82 = v44
						m.G0 = v12 + int32(16)
						return v82
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12))) = v36
							F_errmsg_internal(m, int32(_a_Fn14263_0), v12)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, l4, int32(123), int32(_a_Fn14263_1))
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
				v82 = base.I64_extend_i32_u(v32)
				m.G0 = v12 + int32(16)
				return v82
			}
		}
	} else {
		v60 = int32(1)
		v61 = l1 - v60
		v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(base.Ui32(v61)>>(uint(int32(3))%32)))+23)))
		if int32(base.Ui32(v65)>>(uint(v61&int32(7))%32))&v60 != 0 {
			v76 = F_nocachegetattr(m, l0, l1, l2)
			mBase = m.M
			v77 = m.ExcPending
			if v77 != 0 {
				return int64(0)
			} else {
				v82 = v76
				m.G0 = v12 + int32(16)
				return v82
			}
		} else {
			v71 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v71)
			v82 = int64(0)
			m.G0 = v12 + int32(16)
			return v82
		}
	}
}
func Fn14276(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v13 int64
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v11
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+22)) = uint16(v13)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v15 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v15
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+16)))
	v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29+v30)+12)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v36 = F_gbt_num_consistent(m, v8+int32(12), v8+int32(24), v8+int32(22), v32&int32(1), l1, v35)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		return int64(0)
	} else {
		m.G0 = v8 + int32(32)
		return base.I64_extend_i32_u(v36)
	}
}
func Fn14283(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
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
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
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
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
		v22 = v11 + int32(8)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v25 = int32(4)
		v26 = v23 + v25
		*(*int32)(unsafe.Add(mBase, uint32(v22))) = v26
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
		v29 = int32(2)
		v30 = int32(base.Ui32(v28) >> (uint(v29) % 32))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
		if base.Ui32(v30+v25) < base.Ui32(int32(base.Ui32(v38)>>(uint(v29)%32))) {
			v42 = v26 + (v30+int32(3))&int32(2147483644)
		} else {
			v42 = v26
		}
		*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v42
		v44 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v44)
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		if v46 == v44 {
			v50 = *(*int32)(unsafe.Add(mBase, _c_Fn14283[0]))
			v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
			v56 = *(*int32)(unsafe.Add(mBase, uint32(v51*int32(28))+uint32(_c_Fn14283[1])))
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v56
		} else {
		}
		v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
		v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v64)+16)))
		v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v64+v65)+12)))
		v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v71 = F_gbt_var_consistent(m, v11+int32(8), v15, base.I32_wrap_i64(v19)&int32(_a_Fn14283_0), v63, v67&int32(1), l1, v70)
		mBase = m.M
		v72 = m.ExcPending
		if v72 != 0 {
			return int64(0)
		} else {
			m.G0 = v11 + int32(16)
			return base.I64_extend_i32_u(v71)
		}
	}
}
func Fn14285(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v10 = F_DirectFunctionCall2Coll(m, l3, int32(0), v7, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = base.I32_wrap_i64(v10)
		if v14 != 0 {
			v21 = v14
			return v21
		} else {
			v16 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
			v17 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
			v18 = F_DirectFunctionCall2Coll(m, l3, int32(0), v16, v17)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v21 = base.I32_wrap_i64(v18)
				return v21
			}
		}
	}
}
func Fn14296(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
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
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	if l2 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l3 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v20 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = int32(0)
	goto L4
L4:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37+v25<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = l0
	v43 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v41)+24)) = uint16(v43)
	v46 = v25 + int32(1)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v46 < v47 {
		v25 = v46
		goto L4
	} else {
		goto L6
	}
L5:
	;
	goto L1
L6:
	;
	goto L5
L7:
	;
	m.G0 = v16 + int32(16)
	return
L8:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v64 <= int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v69 = int32(0)
	goto L10
L10:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81+v69<<(uint(int32(2))%32))))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	if base.Ui32(l10) < base.Ui32(v86) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L7
L12:
	;
	v116 = v69 + int32(1)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v116 < v117 {
		v69 = v116
		goto L10
	} else {
		goto L23
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85)+28)) = l0
	v113 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v85)+24)) = uint16(v113)
	goto L12
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v89 = int32(1) << (uint(v86) % 32)
	if v89&l9 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	if v89&l8 == int32(0) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v94 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v85)+24)) = uint8(v94)
	goto L12
L18:
	;
	return
L19:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v104
	F_errmsg(m, int32(_a_Fn14296_0), v16)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, l6, l5, l4)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	goto L11
}
func Fn14300(m *base.Module, l0 int32, l1 int32) int64 {
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 float32
	_ = v47
	var v48 int32
	_ = v48
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)))
			v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
			if v19 != v20 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int64(0)
					} else {
						v29 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
						v30 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v30
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v29
						F_errmsg(m, int32(_a_Fn14300_0), v9)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_Fn14300_1), int32(80), int32(_a_Fn14300_2))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
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
				v42 = int32(8)
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v47 = m.T0[v46].(func(*base.Module, int32, int32, int32) float32)(m, base.I32_extend16_s(v19), v12+v42, v17+v42)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int64(0)
				} else {
					m.G0 = v9 + int32(16)
					return base.I64_reinterpret_f64(base.F64_promote_f32(v47))
				}
			}
		}
	}
}
func Fn14311(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int64
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int64
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v87 int64
	_ = v87
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v90 int32
	_ = v90
	var v96 int64
	_ = v96
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if int32(0) < l1 {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+18)))
		if base.Ui32(v17&int32(2047)) < base.Ui32(l1) {
			v21 = F_getmissingattr(m, l2, l1, l3)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int64(0)
			} else {
				v96 = v21
				m.G0 = v12 + int32(16)
				return v96
			}
		} else {
			v25 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v25)
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+20)))
			if v28&int32(1) == v25 {
				v37 = l2 + l1<<(uint(int32(3))%32) + int32(20)
				v38 = int32(*(*int16)(unsafe.Add(mBase, uint32(v37))))
				if int32(0) <= v38 {
					v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+22)))
					v43 = v27 + v41 + v38
					v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+4)))
					if v44 == int32(1) {
						v47 = int32(*(*int16)(unsafe.Add(mBase, uint32(v37)+2)))
						if base.I32_popcnt(v47) != int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12))) = v47
								F_errmsg_internal(m, int32(_a_Fn14311_0), v12)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, l4, int32(123), int32(_a_Fn14311_1))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							switch base.I32_ctz(v47) {
							case 0:
								v52 = int64(*(*int8)(unsafe.Add(mBase, uint32(v43))))
								v96 = v52
								m.G0 = v12 + int32(16)
								return v96
							case 1:
								v53 = int64(*(*int16)(unsafe.Add(mBase, uint32(v43))))
								v96 = v53
								m.G0 = v12 + int32(16)
								return v96
							case 2:
								v54 = int64(*(*int32)(unsafe.Add(mBase, uint32(v43))))
								v96 = v54
								m.G0 = v12 + int32(16)
								return v96
							case 3:
								v55 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
								v96 = v55
								m.G0 = v12 + int32(16)
								return v96
							default:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12))) = v47
									F_errmsg_internal(m, int32(_a_Fn14311_0), v12)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, l4, int32(123), int32(_a_Fn14311_1))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
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
						v96 = base.I64_extend_i32_u(v43)
						m.G0 = v12 + int32(16)
						return v96
					}
				} else {
					v69 = F_nocachegetattr(m, l0, l1, l2)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int64(0)
					} else {
						v96 = v69
						m.G0 = v12 + int32(16)
						return v96
					}
				}
			} else {
				v71 = int32(1)
				v72 = l1 - v71
				v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(base.Ui32(v72)>>(uint(int32(3))%32)))+23)))
				if int32(base.Ui32(v76)>>(uint(v72&int32(7))%32))&v71 == int32(0) {
					v84 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v84)
					v96 = int64(0)
					m.G0 = v12 + int32(16)
					return v96
				} else {
					v87 = F_nocachegetattr(m, l0, l1, l2)
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return int64(0)
					} else {
						v96 = v87
						m.G0 = v12 + int32(16)
						return v96
					}
				}
			}
		}
	} else {
		v89 = F_heap_getsysattr(m, l0, l1, l3)
		mBase = m.M
		v90 = m.ExcPending
		if v90 != 0 {
			return int64(0)
		} else {
			v96 = v89
			m.G0 = v12 + int32(16)
			return v96
		}
	}
}
func Fn14317(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int64 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	v22 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_check_encoding_conversion_args(m, v23, v24, v25, int32(-1), int32(6))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		return int64(0)
	} else {
		v32 = v23 - l8
		if base.B2i32(base.Ui32(v32) <= base.Ui32(l7))&(int32(base.Ui32(l6)>>(uint(v32)%32))&int32(1)) == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(2600))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v18))) = v23
					F_errmsg(m, l5, v18)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, l4, l3, l2)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
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
			v57 = *(*int32)(unsafe.Add(mBase, uint32(v32<<(uint(int32(2))%32)+l1)))
			v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
			v59 = int32(0)
			v64 = F_LocalToUtf(m, base.I32_wrap_i64(v22), v25, base.I32_wrap_i64(v21), v58, v59, v59, v59, v23, base.B2i32(v20 != int64(0)))
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return int64(0)
			} else {
				m.G0 = v18 + int32(16)
				return base.I64_extend_i32_s(v64)
			}
		}
	}
}
func Fn14324(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
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
func Fn14333(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l3
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)))
	if v11 == int32(1) {
		v14 = int32(_a_Fn14333_0)
		v15 = *(*int32)(unsafe.Add(mBase, _c_Fn14333[0]))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		*(*int32)(unsafe.Add(mBase, _c_Fn14333[0])) = v17
		v20 = F_palloc(m, int32(40))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int64(0)
		} else {
			v24 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v20)+8)) = uint8(v24)
			*(*int64)(unsafe.Add(mBase, uint32(v20))) = int64(0)
			F_initHyperLogLog(m, v20+int32(16), int32(10))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(118)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v20
				*(*int32)(unsafe.Add(mBase, _c_Fn14333[0])) = v15
				return int64(0)
			}
		}
	} else {
		return int64(0)
	}
}
func Fn14344(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int64
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	if v8 == int32(0) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v12 = F_init_MultiFuncCall(m, l0)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = int32(_a_Fn14344_0)
			v17 = *(*int32)(unsafe.Add(mBase, _c_Fn14344[0]))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
			*(*int32)(unsafe.Add(mBase, _c_Fn14344[0])) = v19
			v21 = F_collect_corrupt_items(m, v11, l2, l1)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v21
				*(*int32)(unsafe.Add(mBase, _c_Fn14344[0])) = v17
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
				if base.Ui32(v32) < base.Ui32(v33) {
					v35 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
					*(*int64)(unsafe.Add(mBase, uint32(v30))) = v35 + int64(1)
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v40 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v39)+20)) = v40
					*(*int32)(unsafe.Add(mBase, uint32(v31))) = v32 + v40
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
					return base.I64_extend_i32_u(v45 + v32*int32(6))
				} else {
					F_end_MultiFuncCall(m, l0)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int64(0)
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v53)+20)) = int32(2)
						v56 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v56)
						return int64(0)
					}
				}
			}
		}
	} else {
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
		if base.Ui32(v32) < base.Ui32(v33) {
			v35 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
			*(*int64)(unsafe.Add(mBase, uint32(v30))) = v35 + int64(1)
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v40 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v39)+20)) = v40
			*(*int32)(unsafe.Add(mBase, uint32(v31))) = v32 + v40
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
			return base.I64_extend_i32_u(v45 + v32*int32(6))
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int64(0)
			} else {
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v53)+20)) = int32(2)
				v56 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v56)
				return int64(0)
			}
		}
	}
}
func Fn14353(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int64
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(176)
	m.G0 = v12
	v19 = v4
	v20 = int32(-1)
	v21 = v4
	v22 = v4
	goto L1
L1:
	;
	if v20 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_Fn14353[0])) = v43
	*(*int32)(unsafe.Add(mBase, _c_Fn14353[1])) = v44
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v94 - int32(1)
	m.G0 = v12 + int32(176)
	return
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v27 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v26 + v27
	v31 = *(*int32)(unsafe.Add(mBase, _c_Fn14353[0]))
	v33 = *(*int32)(unsafe.Add(mBase, _c_Fn14353[1]))
	v35 = v12 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v12 + int32(12)
	goto L6
L4:
	;
	v42 = v19
	v43 = v21
	v44 = v22
	goto L5
L5:
	;
	goto L8
L6:
	;
	v42 = int32(0)
	v43 = v31
	v44 = v33
	goto L5
L7:
	;
	goto L2
L8:
	;
	if v42 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L7
L10:
	;
	v69 = int32(m.ExcTag)
	v70 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v69 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	F_standard_ExecutorFinish(m, l0)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L10
	} else {
		goto L18
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_Fn14353[0])) = v12 + int32(16)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v51 == int32(0) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_Fn14353[1])) = v44
	*(*int32)(unsafe.Add(mBase, _c_Fn14353[0])) = v43
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v60 - int32(1)
	F_pg_re_throw(m)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L10
	} else {
		goto L17
	}
L15:
	;
	m.T0[v51].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
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
	v74 = int32(v70)
	m.G0 = v12
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	if v12+int32(12) == v80 {
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
	if v84 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	v84 = v82
	goto L24
L23:
	;
	v84 = int32(0)
	goto L24
L24:
	;
	goto L21
L25:
	;
	F___wasm_longjmp(m, v77, v76)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v19 = v76
	v20 = v84
	v21 = v43
	v22 = v44
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
func Fn14355(m *base.Module, l0 int32, l1 int64, l2 int32) {
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
	F_errmsg_internal(m, int32(_a_Fn14355_0), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L11
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_Fn14355_1), int32(332), l2)
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
func Fn14360(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l3
	v15 = F_query_or_expression_tree_walker_impl(m, l0, l2, v8+int32(8), int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
		m.G0 = v8 + int32(16)
		return v19
	}
}
func Fn14371(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v36 float64
	_ = v36
	var v37 int32
	_ = v37
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = F_get_negator(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		if v12 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int64(0)
			} else {
				F_errmsg_internal(m, int32(_a_Fn14371_0), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_Fn14371_1), int32(783), int32(_a_Fn14371_2))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v36 = F_patternsel_common(m, base.I32_wrap_i64(v10), v12, int32(0), base.I32_wrap_i64(v9), base.I32_wrap_i64(v8), v7, l1, int32(1))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int64(0)
			} else {
				return base.I64_reinterpret_f64(v36)
			}
		}
	}
}
func Fn14382(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32) int64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int64
	_ = v33
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v70 int64
	_ = v70
	var v77 int64
	_ = v77
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int64
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	if base.Ui64(int64(2)) <= base.Ui64(v16-int64(9223372036854775807)) {
		v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v23 = v21 & int64(4294967295)
		v24 = base.I32_wrap_i64(v21)
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
		if v25 != 0 {
			if v25 != int32(2147483647) {
				if v25 != int32(-2147483648) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int64(0)
						} else {
							v107 = F_DirectFunctionCall1Coll(m, int32(1401), int32(0), v23)
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return int64(0)
							} else {
								*(*uint32)(unsafe.Add(mBase, uint32(v14))) = uint32(v107)
								F_errmsg(m, int32(_a_Fn14382_0), v14)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_Fn14382_1), l2, l1)
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
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
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
					if v30 != int32(-2147483648) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int64(0)
							} else {
								v107 = F_DirectFunctionCall1Coll(m, int32(1401), int32(0), v23)
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int64(0)
								} else {
									*(*uint32)(unsafe.Add(mBase, uint32(v14))) = uint32(v107)
									F_errmsg(m, int32(_a_Fn14382_0), v14)
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_Fn14382_1), l2, l1)
										mBase = m.M
										v115 = m.ExcPending
										if v115 != 0 {
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
						v33 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
						if v33 == int64(-9223372036854775807-1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int64(0)
								} else {
									v53 = F_DirectFunctionCall1Coll(m, int32(1401), int32(0), v23)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int64(0)
									} else {
										*(*uint32)(unsafe.Add(mBase, uint32(v14)+16)) = uint32(v53)
										F_errmsg(m, int32(_a_Fn14382_2), v14+int32(16))
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_Fn14382_1), l5, l1)
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
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
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return int64(0)
								} else {
									v107 = F_DirectFunctionCall1Coll(m, int32(1401), int32(0), v23)
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return int64(0)
									} else {
										*(*uint32)(unsafe.Add(mBase, uint32(v14))) = uint32(v107)
										F_errmsg(m, int32(_a_Fn14382_0), v14)
										mBase = m.M
										v112 = m.ExcPending
										if v112 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_Fn14382_1), l2, l1)
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
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
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
				if v36 != int32(2147483647) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int64(0)
						} else {
							v107 = F_DirectFunctionCall1Coll(m, int32(1401), int32(0), v23)
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return int64(0)
							} else {
								*(*uint32)(unsafe.Add(mBase, uint32(v14))) = uint32(v107)
								F_errmsg(m, int32(_a_Fn14382_0), v14)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_Fn14382_1), l2, l1)
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
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
					v39 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
					if v39 != int64(9223372036854775807) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int64(0)
							} else {
								v107 = F_DirectFunctionCall1Coll(m, int32(1401), int32(0), v23)
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int64(0)
								} else {
									*(*uint32)(unsafe.Add(mBase, uint32(v14))) = uint32(v107)
									F_errmsg(m, int32(_a_Fn14382_0), v14)
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_Fn14382_1), l2, l1)
										mBase = m.M
										v115 = m.ExcPending
										if v115 != 0 {
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
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int64(0)
							} else {
								v53 = F_DirectFunctionCall1Coll(m, int32(1401), int32(0), v23)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int64(0)
								} else {
									*(*uint32)(unsafe.Add(mBase, uint32(v14)+16)) = uint32(v53)
									F_errmsg(m, int32(_a_Fn14382_2), v14+int32(16))
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_Fn14382_1), l5, l1)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
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
		} else {
			v64 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
			if v64 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v101 = m.ExcPending
				if v101 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v104 = m.ExcPending
					if v104 != 0 {
						return int64(0)
					} else {
						v107 = F_DirectFunctionCall1Coll(m, int32(1401), int32(0), v23)
						mBase = m.M
						v108 = m.ExcPending
						if v108 != 0 {
							return int64(0)
						} else {
							*(*uint32)(unsafe.Add(mBase, uint32(v14))) = uint32(v107)
							F_errmsg(m, int32(_a_Fn14382_0), v14)
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_Fn14382_1), l2, l1)
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
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
				v65 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
				v66 = base.I64_div_s(v65, l4)
				v70 = base.I64_extend32_s(v66)*int64(-1000000) + v16
				if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v70+int64(211813488000000000)) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_Fn14382_3), int32(0))
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_Fn14382_1), l3, l1)
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
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
					v77 = v70
					m.G0 = v14 + int32(32)
					return v77
				}
			}
		}
	} else {
		v77 = v16
		m.G0 = v14 + int32(32)
		return v77
	}
}
func Fn14384(m *base.Module, l0 int32, l1 int32) int64 {
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
						F_errmsg(m, int32(_a_Fn14384_0), int32(0))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_Fn14384_1), int32(109), int32(_a_Fn14384_2))
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
				v37 = F_anytimestamp_typmod_check(m, l1, v36)
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
func Fn14388(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int64 {
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
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = v12 - v11
	v17 = v16
	v21 = v14
	goto L1
L1:
	;
	v27 = v17 - int32(1)
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v21)&l3)+uint32(_c_Fn14388[0]))))
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
func Fn14395(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v156 int32
	_ = v156
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v4 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if base.Ui32(v24) < base.Ui32(int32(3)) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	return int32(0)
L5:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	F_errmsg(m, int32(_a_Fn14395_0), int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	F_errfinish(m, int32(_a_Fn14395_1), l2, l1)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L9:
	;
	return v156
L10:
	;
	v156 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_Fn14395[0]))
	if v36 == v24 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v156 = int32(1)
	goto L9
L14:
	;
	goto L15
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_Fn14395[1]))
	if v40 <= int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v156 = v146
	goto L9
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_Fn14395[2]))
	if v44 == int32(0) {
		v146 = int32(0)
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _c_Fn14395[3]))
	v116 = int32(0)
	v119 = v40 - int32(1)
	goto L39
L20:
	;
	v49 = v44
	goto L21
L21:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v49)+20))
	if v55 == int32(4) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v146 = int32(0)
	goto L16
L23:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v49)+80))
	if v109 != 0 {
		v49 = v109
		goto L21
	} else {
		goto L38
	}
L24:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v58 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v61 = int32(1)
	if v24 == v58 {
		v146 = v61
		goto L16
	} else {
		goto L26
	}
L26:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v49)+52))
	v65 = v63 - int32(1)
	if v65 < int32(0) {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v49)+48))
	v71 = int32(0)
	v74 = v65
	goto L28
L28:
	;
	v79 = int32(2)
	v80 = base.I32_div_s(v74-v71, v79)
	v81 = v80 + v71
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v68+v81<<(uint(v79)%32))))
	if v85 == v24 {
		v146 = v61
		goto L16
	} else {
		goto L30
	}
L29:
	;
	goto L23
L30:
	;
	v94 = base.B2i32(v85-v24 < int32(0)) | base.B2i32(base.Ui32(v85) < base.Ui32(int32(3)))
	if v94 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v95 = v81 + int32(1)
	goto L33
L32:
	;
	v95 = v71
	goto L33
L33:
	;
	if v94 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v98 = v74
	goto L36
L35:
	;
	v98 = v81 - int32(1)
	goto L36
L36:
	;
	if v95 <= v98 {
		v71 = v95
		v74 = v98
		goto L28
	} else {
		goto L37
	}
L37:
	;
	goto L29
L38:
	;
	goto L22
L39:
	;
	v124 = int32(2)
	v125 = base.I32_div_s(v119-v116, v124)
	v126 = v125 + v116
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v114+v126<<(uint(v124)%32))))
	v131 = base.B2i32(v130 == v24)
	if v130 == v24 {
		v146 = v131
		goto L16
	} else {
		goto L41
	}
L40:
	;
	v146 = v131
	goto L16
L41:
	;
	v134 = base.B2i32(base.Ui32(v130) < base.Ui32(v24))
	if base.Ui32(v130) < base.Ui32(v24) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v135 = v126 + int32(1)
	goto L44
L43:
	;
	v135 = v116
	goto L44
L44:
	;
	if base.Ui32(v130) < base.Ui32(v24) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v138 = v119
	goto L47
L46:
	;
	v138 = v126 - int32(1)
	goto L47
L47:
	;
	if v135 <= v138 {
		v116 = v135
		v119 = v138
		goto L39
	} else {
		goto L48
	}
L48:
	;
	goto L40
}
func Fn14399(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int64 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	v22 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_check_encoding_conversion_args(m, v23, v24, v25, int32(6), int32(-1))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		return int64(0)
	} else {
		v32 = v24 - l8
		if base.B2i32(base.Ui32(v32) <= base.Ui32(l7))&(int32(base.Ui32(l6)>>(uint(v32)%32))&int32(1)) == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(2600))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v18))) = v24
					F_errmsg(m, l5, v18)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, l4, l3, l2)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
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
			v57 = *(*int32)(unsafe.Add(mBase, uint32(v32<<(uint(int32(2))%32)+l1)))
			v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
			v59 = int32(0)
			v64 = F_UtfToLocal(m, base.I32_wrap_i64(v22), v25, base.I32_wrap_i64(v21), v58, v59, v59, v59, v24, base.B2i32(v20 != int64(0)))
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return int64(0)
			} else {
				m.G0 = v18 + int32(16)
				return base.I64_extend_i32_s(v64)
			}
		}
	}
}
func Fn14407(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
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
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 float32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 float64
	_ = v96
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v18 = F_pg_detoast_datum_packed(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
			v22 = v20 & int32(1)
			if v22 != 0 {
				v23 = int32(1)
			} else {
				v23 = int32(4)
			}
			if v20 == int32(1) {
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
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
				if v22 != 0 {
					v51 = int32(base.Ui32(v20)>>(uint(v41)%32)) - v41
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
					v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v52 = int32(1)
			v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
			v56 = v54 & v52
			if v56 != 0 {
				v57 = v52
			} else {
				v57 = int32(4)
			}
			if v54 == int32(1) {
				v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
				if v64 == int32(18) {
					v67 = int32(16)
				} else {
					v67 = int32(0)
				}
				if base.Ui32((v64-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v74 = int32(4)
				} else {
					v74 = v67
				}
				v85 = v74
			} else {
				v75 = int32(1)
				if v56 != 0 {
					v85 = int32(base.Ui32(v54)>>(uint(v75)%32)) - v75
				} else {
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v85 = int32(base.Ui32(v79)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v86 = F_calc_word_similarity(m, v18+v23, v51, v11+v57, v85, l2)
			mBase = m.M
			v87 = m.ExcPending
			if v87 != 0 {
				return int64(0)
			} else {
				v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v88 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int64(0)
					} else {
						v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v92 != v18 {
							F_pfree(m, v18)
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int64(0)
							} else {
								v96 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
								return base.I64_extend_i32_u(base.F64_le(v96, base.F64_promote_f32(v86)))
							}
						} else {
							v96 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
							return base.I64_extend_i32_u(base.F64_le(v96, base.F64_promote_f32(v86)))
						}
					}
				} else {
					v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v92 != v18 {
						F_pfree(m, v18)
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int64(0)
						} else {
							v96 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
							return base.I64_extend_i32_u(base.F64_le(v96, base.F64_promote_f32(v86)))
						}
					} else {
						v96 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
						return base.I64_extend_i32_u(base.F64_le(v96, base.F64_promote_f32(v86)))
					}
				}
			}
		}
	}
}
func Fn14409(m *base.Module, l0 int32, l1 int32) int64 {
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
								return base.I64_extend_i32_s(base.I32_reinterpret_f32(base.F32_sub(float32(1), v84)))
							}
						} else {
							return base.I64_extend_i32_s(base.I32_reinterpret_f32(base.F32_sub(float32(1), v84)))
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
							return base.I64_extend_i32_s(base.I32_reinterpret_f32(base.F32_sub(float32(1), v84)))
						}
					} else {
						return base.I64_extend_i32_s(base.I32_reinterpret_f32(base.F32_sub(float32(1), v84)))
					}
				}
			}
		}
	}
}
func Fn14412(m *base.Module, l0 int32, l1 int32) int32 {
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	F_errstart_cold(m, int32(21), int32(0))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errmsg(m, int32(_a_Fn14412_0), int32(0))
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v18 = F_errdetail(m, int32(_a_Fn14412_1), int32(0))
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_Fn14412_2), l1, l0)
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
