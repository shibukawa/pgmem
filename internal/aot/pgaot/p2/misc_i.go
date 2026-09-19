package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_IdleInTransactionSessionTimeoutHandler(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	v2 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_IdleInTransactionSessionTimeoutHandler[0])) = v2
	*(*int32)(unsafe.Add(mBase, _c_F_IdleInTransactionSessionTimeoutHandler[1])) = v2
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_IdleInTransactionSessionTimeoutHandler[2]))
	F_SetLatch(m, v8)
	mBase = m.M
	return
}
func F_IncrementVarSublevelsUp(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v5 = m.G0
	v6 = int32(16)
	v7 = v5 - v6
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l1
	v15 = F_query_or_expression_tree_walker_impl(m, l0, int32(1050), v7+int32(8), v6)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		m.G0 = v7 + int32(16)
		return
	}
}
func F_InitConflictIndexes(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
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
	var v23 int32
	_ = v23
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2 < v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = v7
	v12 = v2
	v14 = v2
	goto L4
L2:
	;
	v46 = v2
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v46
	return
L4:
	;
	v17 = v12 << (uint(int32(2)) % 32)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17+v18)))
	if v20 == int32(0) {
		v37 = v11
		v38 = v14
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v46 = v38
	goto L3
L6:
	;
	v40 = v12 + int32(1)
	if v40 < v37 {
		v11 = v37
		v12 = v40
		v14 = v38
		goto L4
	} else {
		goto L12
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v23+v17)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+116)))
	if v26 != int32(1) {
		v37 = v11
		v38 = v14
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v20)+192))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+16)))
	if v30 != int32(1) {
		v37 = v11
		v38 = v14
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	v34 = F_lappend_oid(m, v14, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v37 = v36
	v38 = v34
	goto L6
L12:
	;
	goto L5
}
func F_IpcMemoryDelete(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = int32(0)
	v9 = F_pgmem_shmctl(m, l1, v7, v7)
	mBase = m.M
	if v7 <= v9 {
		m.G0 = v5 + int32(16)
		return
	} else {
		v14 = F_errstart(m, int32(15), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			if v14 == int32(0) {
				m.G0 = v5 + int32(16)
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = l1
				F_errmsg_internal(m, int32(_a_F_IpcMemoryDelete_0), v5)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_IpcMemoryDelete_1), int32(302), int32(_a_F_IpcMemoryDelete_2))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						m.G0 = v5 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_IpcMemoryDetach(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = F_pgmem_shmdt(m, l1)
	mBase = m.M
	if int32(0) <= v7 {
		m.G0 = v5 + int32(16)
		return
	} else {
		v12 = F_errstart(m, int32(15), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			if v12 == int32(0) {
				m.G0 = v5 + int32(16)
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = l1
				F_errmsg_internal(m, int32(_a_F_IpcMemoryDetach_0), v5)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_IpcMemoryDetach_1), int32(290), int32(_a_F_IpcMemoryDetach_2))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						m.G0 = v5 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_i2tof(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	return base.I32_reinterpret_f32(base.F32_convert_i32_s(v2))
}
func F_i4tochar(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(v2-int32(128)) <= base.Ui32(int32(-257)) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_i4tochar_0), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_i4tochar_1), int32(197), int32(_a_F_i4tochar_2))
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
		return v2
	}
}
func F_i8tod(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v5 = F_Float8GetDatum(m, base.F64_convert_i64_s(v3))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_i8tooid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	if base.Ui64(int64(4294967296)) <= base.Ui64(v4) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_i8tooid_0), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_i8tooid_1), int32(1360), int32(_a_F_i8tooid_2))
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
		return base.I32_wrap_i64(v4)
	}
}
func F_icount(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
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
	var v17 int32
	_ = v17
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
		v12 = F_ArrayGetNItemsSafe(m, v9, v5+int32(16))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v14 != v5 {
				F_pfree(m, v5)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					return v12
				}
			} else {
				return v12
			}
		}
	}
}
func F_icregexnesel(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14011(m, l0, int32(3))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_import_error_callback(m *base.Module, l0 int32) {
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
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
						v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						if v20 != 0 {
							F_set_errcontext_domain(m, int32(0))
							mBase = m.M
							v23 = m.ExcPending
							if v23 != 0 {
								return
							} else {
								v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v6))) = v24
								F_errcontext_msg(m, int32(_a_F_import_error_callback_0), v6)
								mBase = m.M
								v28 = m.ExcPending
								if v28 != 0 {
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
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v20 != 0 {
				F_set_errcontext_domain(m, int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v24
					F_errcontext_msg(m, int32(_a_F_import_error_callback_0), v6)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
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
}
func F_inclusion_get_strategy_procinfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = l1 - int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0+v14<<(uint(int32(2))%32))+20))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+116))
	if l2 != v20 {
		v22 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v19)+936)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v19)+908)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v19)+880)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v19)+852)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v19)+824)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v19)+796)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v19)+768)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v19)+740)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v19)+712)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v19)+684)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v19)+656)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v19)+628)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v19)+600)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v19)+572)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v19)+544)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v19)+516)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v19)+488)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v19)+460)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v19)+432)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v19)+404)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v19)+376)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v19)+348)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v19)+320)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v19)+292)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v19)+264)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v19)+236)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v19)+208)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v19)+180)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v19)+152)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v19)+124)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v19)+116)) = l2
	} else {
	}
	v85 = v19 + l3*int32(28)
	v87 = v85 + int32(92)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+96))
	if v88 == int32(0) {
		v91 = int32(4)
		v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+208))
		v97 = *(*int32)(unsafe.Add(mBase, uint32(v93+v14<<(uint(int32(2))%32))))
		v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
		v107 = v98 + v99<<(uint(v91)%32) + v14*int32(100) + int32(88)
		v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
		v110 = F_SearchSysCache4(m, v91, v97, v108, l2, base.I32_extend16_s(l3))
		mBase = m.M
		v113 = m.ExcPending
		if v113 != 0 {
			return int32(0)
		} else {
			if v110 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v137 = m.ExcPending
				if v137 != 0 {
					return int32(0)
				} else {
					v138 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
					*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v97
					*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v138
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = l3
					F_errmsg_internal(m, int32(_a_F_inclusion_get_strategy_procinfo_0), v11)
					mBase = m.M
					v145 = m.ExcPending
					if v145 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_inclusion_get_strategy_procinfo_1), int32(648), int32(_a_F_inclusion_get_strategy_procinfo_2))
						mBase = m.M
						v150 = m.ExcPending
						if v150 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v118 = F_SysCacheGetAttrNotNull(m, int32(4), v110, int32(7))
				mBase = m.M
				v119 = m.ExcPending
				if v119 != 0 {
					return int32(0)
				} else {
					F_ReleaseCatCache(m, v110)
					mBase = m.M
					v121 = m.ExcPending
					if v121 != 0 {
						return int32(0)
					} else {
						v122 = F_get_opcode(m, v118)
						mBase = m.M
						v123 = m.ExcPending
						if v123 != 0 {
							return int32(0)
						} else {
							v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							F_fmgr_info_cxt(m, v122, v87, v124)
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return int32(0)
							} else {
								m.G0 = v11 + int32(16)
								return v87
							}
						}
					}
				}
			}
		}
	} else {
		m.G0 = v11 + int32(16)
		return v87
	}
}
func F_ineq_histogram_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) float64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v28 float64
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 float64
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v94 int32
	_ = v94
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v155 int32
	_ = v155
	var v179 int32
	_ = v179
	var v190 int32
	_ = v190
	var v213 int32
	_ = v213
	var v224 int32
	_ = v224
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v265 int32
	_ = v265
	var v274 int32
	_ = v274
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 float64
	_ = v324
	var v329 float64
	_ = v329
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v362 float64
	_ = v362
	var v364 float64
	_ = v364
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v381 int32
	_ = v381
	var v395 int32
	_ = v395
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 float64
	_ = v467
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 float32
	_ = v475
	var v477 float32
	_ = v477
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v504 float64
	_ = v504
	var v505 float64
	_ = v505
	var v509 int32
	_ = v509
	var v510 float64
	_ = v510
	var v513 float64
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v520 float64
	_ = v520
	var v523 int32
	_ = v523
	var v530 float64
	_ = v530
	var v533 float64
	_ = v533
	var v534 int32
	_ = v534
	var v539 float64
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v553 float64
	_ = v553
	var v554 float64
	_ = v554
	var v559 float64
	_ = v559
	var v562 float64
	_ = v562
	var v563 float64
	_ = v563
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v595 int32
	_ = v595
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v634 int32
	_ = v634
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v664 int32
	_ = v664
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v694 int32
	_ = v694
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v754 int32
	_ = v754
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v835 int32
	_ = v835
	var v856 float64
	_ = v856
	var v859 float64
	_ = v859
	var v861 int32
	_ = v861
	var v864 float64
	_ = v864
	var v865 int32
	_ = v865
	var v904 float64
	_ = v904
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v917 int32
	_ = v917
	var v921 int32
	_ = v921
	var v942 float64
	_ = v942
	var v943 float64
	_ = v943
	var v947 int32
	_ = v947
	var v950 float64
	_ = v950
	var v951 int32
	_ = v951
	var v988 float64
	_ = v988
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1031 float64
	_ = v1031
	var v1032 float64
	_ = v1032
	var v1036 int32
	_ = v1036
	var v1039 float64
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1128 int32
	_ = v1128
	var v1129 float64
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1132 float64
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1135 float64
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1138 int32
	_ = v1138
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1197 int64
	_ = v1197
	var v1208 float64
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1213 int32
	_ = v1213
	var v1217 int64
	_ = v1217
	var v1221 int64
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1227 int64
	_ = v1227
	var v1232 int64
	_ = v1232
	var v1234 float64
	_ = v1234
	var v1242 int64
	_ = v1242
	var v1246 int64
	_ = v1246
	var v1258 float64
	_ = v1258
	var v1270 float64
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1274 float64
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1278 float64
	_ = v1278
	var v1280 int64
	_ = v1280
	var v1286 int32
	_ = v1286
	var v1290 int32
	_ = v1290
	var v1294 int64
	_ = v1294
	var v1299 int64
	_ = v1299
	var v1303 int64
	_ = v1303
	var v1308 int32
	_ = v1308
	var v1310 float64
	_ = v1310
	var v1312 int64
	_ = v1312
	var v1317 int32
	_ = v1317
	var v1321 int64
	_ = v1321
	var v1332 int64
	_ = v1332
	var v1336 int64
	_ = v1336
	var v1340 int32
	_ = v1340
	var v1347 int32
	_ = v1347
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1375 int32
	_ = v1375
	var v1380 int32
	_ = v1380
	var v1406 int32
	_ = v1406
	var v1408 int32
	_ = v1408
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1415 int32
	_ = v1415
	var v1422 int32
	_ = v1422
	var v1449 int32
	_ = v1449
	var v1452 int32
	_ = v1452
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1494 int32
	_ = v1494
	var v1498 int32
	_ = v1498
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1534 int32
	_ = v1534
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1550 int32
	_ = v1550
	var v1577 int32
	_ = v1577
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1583 int32
	_ = v1583
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1594 float64
	_ = v1594
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1623 float64
	_ = v1623
	var v1626 float64
	_ = v1626
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1632 int32
	_ = v1632
	var v1636 float64
	_ = v1636
	var v1638 int32
	_ = v1638
	var v1675 float64
	_ = v1675
	var v1678 int32
	_ = v1678
	var v1681 int32
	_ = v1681
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1688 int32
	_ = v1688
	var v1690 float64
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1698 int32
	_ = v1698
	var v1719 float64
	_ = v1719
	var v1720 float64
	_ = v1720
	var v1724 int32
	_ = v1724
	var v1726 int32
	_ = v1726
	var v1728 int32
	_ = v1728
	var v1732 float64
	_ = v1732
	var v1734 int32
	_ = v1734
	var v1769 float64
	_ = v1769
	var v1774 int32
	_ = v1774
	var v1778 int32
	_ = v1778
	var v1781 int32
	_ = v1781
	var v1782 int32
	_ = v1782
	var v1786 int32
	_ = v1786
	var v1788 float64
	_ = v1788
	var v1790 int32
	_ = v1790
	var v1796 int32
	_ = v1796
	var v1817 float64
	_ = v1817
	var v1818 float64
	_ = v1818
	var v1822 int32
	_ = v1822
	var v1824 int32
	_ = v1824
	var v1826 int32
	_ = v1826
	var v1830 float64
	_ = v1830
	var v1832 int32
	_ = v1832
	var v1867 float64
	_ = v1867
	var v1873 int32
	_ = v1873
	var v1875 int32
	_ = v1875
	var v1877 int32
	_ = v1877
	var v1918 int32
	_ = v1918
	var v1919 float64
	_ = v1919
	var v1920 int32
	_ = v1920
	var v1922 float64
	_ = v1922
	var v1923 int32
	_ = v1923
	var v1925 float64
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1928 int32
	_ = v1928
	var v1932 int64
	_ = v1932
	var v1947 int32
	_ = v1947
	var v1979 float64
	_ = v1979
	var v1980 float64
	_ = v1980
	var v1982 float64
	_ = v1982
	var v1987 float64
	_ = v1987
	var v1992 float64
	_ = v1992
	var v1998 float64
	_ = v1998
	var v2001 float64
	_ = v2001
	var v2004 float64
	_ = v2004
	var v2007 float64
	_ = v2007
	var v2009 float64
	_ = v2009
	var v2012 int32
	_ = v2012
	var v2013 int32
	_ = v2013
	var v2016 float64
	_ = v2016
	var v2023 float64
	_ = v2023
	var v2025 float64
	_ = v2025
	var v2027 float64
	_ = v2027
	var v2028 float64
	_ = v2028
	var v2031 float64
	_ = v2031
	var v2062 float64
	_ = v2062
	var v2069 float64
	_ = v2069
	var v2107 float64
	_ = v2107
	var v2111 int32
	_ = v2111
	var v2115 float64
	_ = v2115
	var v2118 float64
	_ = v2118
	var v2149 float64
	_ = v2149
	var v2158 int32
	_ = v2158
	var v2186 float64
	_ = v2186
	v10 = int32(0)
	v28 = float64(0)
	v34 = m.G0
	v36 = v34 - int32(112)
	m.G0 = v36
	v38 = float64(-1)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v39 == v10 {
		v2186 = v38
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v36 + int32(112)
	return v2186
L2:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+29)))
	if v44 != 0 {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	F_free_attstatsslot(m, v36+int32(76))
	mBase = m.M
	v2158 = m.ExcPending
	if v2158 != 0 {
		goto L13
	} else {
		goto L471
	}
L4:
	;
	v2111 = *(*int32)(unsafe.Add(mBase, uint32(v36)+92))
	v2115 = base.F64_div(float64(0.01), base.F64_convert_i32_s(v2111-int32(1)))
	if base.F64_lt(v2107, v2115) != 0 {
		v2149 = v2115
		goto L3
	} else {
		goto L469
	}
L5:
	;
	v373 = int32(0)
	v381 = v213
	v395 = v369
	goto L57
L6:
	;
	v364 = base.F64_sub(float64(1), v329)
	if base.F64_gt(v324, v364) == int32(0) {
		v2149 = v324
		goto L3
	} else {
		goto L56
	}
L7:
	;
	if v213 == int32(2) {
		goto L48
	} else {
		goto L49
	}
L8:
	;
	v335 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L13
	} else {
		goto L43
	}
L9:
	;
	v56 = v39
	goto L11
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v45 == int32(0) {
		v2186 = v38
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v60 = F_get_attstatsslot(m, v36+int32(76), v56, int32(2), int32(0), int32(1))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L13
	} else {
		goto L16
	}
L12:
	;
	v48 = F_get_func_leakproof(m, v45)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return float64(0)
L14:
	;
	if v48 == int32(0) {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v56 = v54
	goto L11
L16:
	;
	if v60 == int32(0) {
		v2186 = v38
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v36)+92))
	if v64 < int32(2) {
		v224 = v64
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if v224 < int32(2) {
		v2149 = v38
		goto L3
	} else {
		goto L37
	}
L19:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v36)+80))
	if v67 != l6 {
		v224 = v64
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v36)+76))
	if l2 != v70 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v72 = int32(0)
	v77 = F_SearchSysCacheList(m, int32(3), int32(1), v70, v72, v72)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L13
	} else {
		goto L25
	}
L22:
	;
	v190 = int32(1)
	goto L23
L23:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v36)+92))
	if v190 != 0 {
		goto L7
	} else {
		goto L36
	}
L24:
	;
	F_ReleaseCatCacheList(m, v77)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L13
	} else {
		goto L35
	}
L25:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v77)+40))
	if v79 <= int32(0) {
		v155 = v72
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v94 = v72
	goto L27
L27:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v77+int32(48)+v94<<(uint(int32(2))%32))))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+56))
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+22)))
	v125 = v123 + v124
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	v128 = F_SearchSysCacheExists(m, int32(3), l2, int32(115), v126, int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L13
	} else {
		goto L30
	}
L28:
	;
	v155 = int32(0)
	goto L24
L29:
	;
	v141 = v94 + int32(1)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v77)+40))
	if v141 < v142 {
		v94 = v141
		goto L27
	} else {
		goto L34
	}
L30:
	;
	if v128 == int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v125)+24))
	v134 = F_GetIndexAmRoutineByAmId(m, v132, int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L13
	} else {
		goto L32
	}
L32:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+14)))
	if v136 == int32(0) {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	v155 = int32(1)
	goto L24
L34:
	;
	goto L28
L35:
	;
	v190 = v155
	goto L23
L36:
	;
	v224 = v213
	goto L18
L37:
	;
	v249 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+48)) = uint8(v249)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+40)) = uint8(v249)
	v253 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v36)+34)) = uint16(v253)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+32)) = uint8(v249)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+28)) = l6
	*(*int64)(unsafe.Add(mBase, uint32(v36)+20)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v36)+44)) = l7
	v265 = v249
	v274 = v249
	goto L38
L38:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v36)+88))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v297+v265<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+36)) = v301
	v303 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+32)) = uint8(v303)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
	v309 = m.T0[v308].(func(*base.Module, int32) int32)(m, v36+int32(16))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L13
	} else {
		goto L40
	}
L39:
	;
	v324 = base.F64_div(base.F64_convert_i32_s(v317), base.F64_convert_i32_s(v320))
	v329 = base.F64_div(float64(0.01), base.F64_convert_i32_s(v320-int32(1)))
	if base.F64_lt(v324, v329) == int32(0) {
		goto L6
	} else {
		goto L42
	}
L40:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+32)))
	v317 = v274 + (v311^int32(-1))&base.B2i32(v309 != int32(0))
	v319 = v265 + int32(1)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v36)+92))
	if v319 < v320 {
		v265 = v319
		v274 = v317
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v2149 = v329
	goto L3
L43:
	;
	if v335 == int32(0) {
		v2186 = v38
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v339 = F_get_func_name(m, v45)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v339
	F_errmsg_internal(m, int32(_a_F_ineq_histogram_selectivity_0), v36)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L13
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_ineq_histogram_selectivity_1), int32(_a_F_ineq_histogram_selectivity_2), int32(_a_F_ineq_histogram_selectivity_3))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	v2186 = v38
	goto L1
L48:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v36)+76))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v36)+88))
	v356 = F_get_actual_variable_range(m, l0, l1, v352, l6, v353, v353+int32(4))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L13
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	if int32(0) < v213 {
		v369 = v10
		goto L5
	} else {
		goto L52
	}
L51:
	;
	v369 = v356
	goto L5
L52:
	;
	if l4 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v362 = float64(1)
	goto L55
L54:
	;
	v362 = float64(0)
	goto L55
L55:
	;
	v2107 = v362
	goto L4
L56:
	;
	v2149 = v364
	goto L3
L57:
	;
	v404 = v373 + v381
	v405 = int32(2)
	v406 = base.I32_div_s(v404, v405)
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v36)+92))
	if base.B2i32(v407 < int32(3))|base.B2i32(base.Ui32(v405) < base.Ui32(v404+int32(1))) == int32(0) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	if v449 <= int32(0) {
		goto L76
	} else {
		goto L77
	}
L59:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v36)+88))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v439+v406<<(uint(int32(2))%32))))
	v444 = F_FunctionCall2Coll(m, l3, l6, v443, l7)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L13
	} else {
		goto L66
	}
L60:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v36)+76))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v36)+88))
	v420 = F_get_actual_variable_range(m, l0, l1, v417, l6, v418, int32(0))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L13
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	if base.B2i32(v406 != v407-int32(1))|base.B2i32(v407 < int32(3)) != 0 {
		v436 = v395
		goto L59
	} else {
		goto L64
	}
L63:
	;
	v436 = v420
	goto L59
L64:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v36)+76))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v36)+88))
	v434 = F_get_actual_variable_range(m, l0, l1, v428, l6, int32(0), v430+v406<<(uint(int32(2))%32))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L13
	} else {
		goto L65
	}
L65:
	;
	v436 = v434
	goto L59
L66:
	;
	v448 = l4 ^ base.B2i32(v444 != int32(0))
	if v448 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v449 = v406 + int32(1)
	goto L69
L68:
	;
	v449 = v373
	goto L69
L69:
	;
	if v448 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v450 = v381
	goto L72
L71:
	;
	v450 = v406
	goto L72
L72:
	;
	if v449 < v450 {
		v373 = v449
		v381 = v450
		v395 = v436
		goto L57
	} else {
		goto L73
	}
L73:
	;
	goto L58
L74:
	;
	if v436&int32(1) == int32(0) {
		v2107 = v2062
		goto L4
	} else {
		goto L466
	}
L75:
	;
	if l4 != 0 {
		goto L463
	} else {
		goto L464
	}
L76:
	;
	v2028 = float64(0)
	goto L75
L77:
	;
	goto L78
L78:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v36)+92))
	if v455 <= v449 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v2028 = float64(1)
	goto L75
L80:
	;
	goto L81
L81:
	;
	v458 = l4 ^ l5
	if v458 == int32(1) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v563 = float64(0.5)
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v36)+88))
	v566 = v449 - int32(1)
	v567 = int32(2)
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v564+v566<<(uint(v567)%32))))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v564+v449<<(uint(v567)%32))))
	v575 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v577 = v36 + int32(56)
	v579 = v36 - int32(-64)
	v580 = m.G0
	v581 = int32(16)
	v582 = v580 - v581
	m.G0 = v582
	v584 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v582)+15)) = uint8(v584)
	v587 = v36 + v581
	if l8 <= int32(1081) {
		goto L143
	} else {
		goto L144
	}
L83:
	;
	if v449 != int32(1) {
		v562 = float64(0)
		goto L82
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v465 = v36 - int32(-64)
	v466 = int32(0)
	v467 = float64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v465))) = uint8(v466)
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v471 != 0 {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	goto L85
L87:
	;
	v541 = v36 + int32(16)
	v542 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v546 = F_get_attstatsslot(m, v541, v542, int32(1), int32(0), int32(2))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L13
	} else {
		goto L120
	}
L88:
	;
	v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+28)))
	if v509 != 0 {
		goto L102
	} else {
		goto L103
	}
L89:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v471)+16))
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472)+22)))
	v474 = v472 + v473
	v475 = *(*float32)(unsafe.Add(mBase, uint32(v474)+8))
	v477 = *(*float32)(unsafe.Add(mBase, uint32(v474)+16))
	v504 = base.F64_promote_f32(v477)
	v505 = base.F64_promote_f32(v475)
	goto L88
L90:
	;
	goto L91
L91:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v479 == int32(16) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v504 = float64(2)
	v505 = v467
	goto L88
L93:
	;
	goto L94
L94:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v483 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v490 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v483)+76))
	if v486 != int32(5) {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v504 = float64(-1)
	v505 = v467
	goto L88
L98:
	;
	v504 = float64(0)
	v505 = v467
	goto L88
L99:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v490)))
	if v493 != int32(6) {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v497 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v490)+8)))
	switch v497 - int32(_a_F_ineq_histogram_selectivity_4) {
	case 0:
		goto L101
	default:
		goto L98
	case 5:
		v504 = float64(-1)
		v505 = v467
		goto L88
	}
L101:
	;
	v504 = float64(1)
	v505 = v467
	goto L88
L102:
	;
	v510 = base.F64_neg(base.F64_sub(float64(1), v505))
	goto L104
L103:
	;
	v510 = v504
	goto L104
L104:
	;
	if base.F64_gt(v510, float64(0)) != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v513 = F_clamp_row_est(m, v510)
	mBase = m.M
	v539 = v513
	goto L87
L106:
	;
	goto L107
L107:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v514 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v517 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v465))) = uint8(v517)
	v539 = float64(200)
	goto L87
L109:
	;
	goto L110
L110:
	;
	v520 = *(*float64)(unsafe.Add(mBase, uint32(v514)+120))
	if base.F64_le(v520, float64(0)) != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v523 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v465))) = uint8(v523)
	v539 = float64(200)
	goto L87
L112:
	;
	goto L113
L113:
	;
	if base.F64_lt(v510, float64(0)) != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v530 = F_clamp_row_est(m, base.F64_mul(v520, base.F64_neg(v510)))
	mBase = m.M
	v539 = v530
	goto L87
L115:
	;
	goto L116
L116:
	;
	if base.F64_lt(v520, float64(200)) != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v533 = F_clamp_row_est(m, v520)
	mBase = m.M
	v539 = v533
	goto L87
L118:
	;
	goto L119
L119:
	;
	v534 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v465))) = uint8(v534)
	v539 = float64(200)
	goto L87
L120:
	;
	if v546 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v36)+40))
	F_free_attstatsslot(m, v541)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L13
	} else {
		goto L124
	}
L122:
	;
	v553 = v539
	goto L123
L123:
	;
	v554 = float64(1)
	if base.F64_gt(v553, v554) != 0 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v553 = base.F64_sub(v539, base.F64_convert_i32_s(v548))
	goto L123
L125:
	;
	v559 = base.F64_div(v554, v553)
	goto L127
L126:
	;
	v559 = float64(0)
	goto L127
L127:
	;
	v562 = v559
	goto L82
L128:
	;
	m.G0 = v582 + int32(16)
	if v1947&int32(1) == int32(0) {
		v2007 = v563
		goto L436
	} else {
		goto L437
	}
L129:
	;
	v1932 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v579))) = v1932
	*(*int64)(unsafe.Add(mBase, uint32(v577))) = v1932
	*(*int64)(unsafe.Add(mBase, uint32(v587))) = v1932
	v1947 = int32(0)
	goto L128
L130:
	;
	v1918 = v582 + int32(15)
	v1919 = F_convert_network_to_scalar(m, l7, l8, v1918)
	mBase = m.M
	v1920 = m.ExcPending
	if v1920 != 0 {
		goto L13
	} else {
		goto L433
	}
L131:
	;
	if l8 != int32(650) {
		goto L129
	} else {
		goto L432
	}
L132:
	;
	v1947 = v1150 ^ int32(1)
	goto L128
L133:
	;
	if v1151 != 0 {
		goto L318
	} else {
		goto L319
	}
L134:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v587))) = v1234
	if v575 <= int32(1183) {
		goto L298
	} else {
		goto L299
	}
L135:
	;
	if l8 != int32(1114) {
		goto L129
	} else {
		goto L290
	}
L136:
	;
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(l7)+8))
	v1227 = *(*int64)(unsafe.Add(mBase, uint32(l7)))
	v1234 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1223), float64(1e+06)), base.F64_convert_i64_s(v1227))
	goto L134
L137:
	;
	v1221 = *(*int64)(unsafe.Add(mBase, uint32(l7)))
	v1234 = base.F64_convert_i64_s(v1221)
	goto L134
L138:
	;
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(l7)+8))
	v1217 = *(*int64)(unsafe.Add(mBase, uint32(l7)))
	v1234 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1209), float64(2.6298e+12)), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1213), float64(8.64e+10)), base.F64_convert_i64_s(v1217)))
	goto L134
L139:
	;
	if l7 == int32(-2147483648) {
		goto L284
	} else {
		goto L285
	}
L140:
	;
	v1197 = *(*int64)(unsafe.Add(mBase, uint32(l7)))
	v1234 = base.F64_convert_i64_s(v1197)
	goto L134
L141:
	;
	v1143 = v582 + int32(15)
	v1144 = F_convert_string_datum(m, l7, l8, l6, v1143)
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		goto L13
	} else {
		goto L267
	}
L142:
	;
	v1128 = v582 + int32(15)
	v1129 = F_convert_numeric_to_scalar(m, l7, l8, v1128)
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L13
	} else {
		goto L264
	}
L143:
	;
	if l8 <= int32(699) {
		goto L146
	} else {
		goto L147
	}
L144:
	;
	goto L145
L145:
	;
	if l8 <= int32(2201) {
		goto L248
	} else {
		goto L249
	}
L146:
	;
	if base.Ui32(int32(26)) < base.Ui32(l8) {
		goto L131
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	if l8 <= int32(828) {
		goto L240
	} else {
		goto L241
	}
L149:
	;
	v595 = int32(1) << (uint(l8) % 32)
	if v595&int32(95485952) != 0 {
		goto L142
	} else {
		goto L150
	}
L150:
	;
	if v595&int32(34340864) != 0 {
		goto L141
	} else {
		goto L151
	}
L151:
	;
	if l8 != int32(17) {
		goto L131
	} else {
		goto L152
	}
L152:
	;
	if v575 != int32(17) {
		v1947 = int32(0)
		goto L128
	} else {
		goto L153
	}
L153:
	;
	v605 = F_pg_detoast_datum_packed(m, l7)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L13
	} else {
		goto L154
	}
L154:
	;
	v607 = F_pg_detoast_datum_packed(m, v570)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L13
	} else {
		goto L155
	}
L155:
	;
	v609 = F_pg_detoast_datum_packed(m, v574)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L13
	} else {
		goto L156
	}
L156:
	;
	v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v605))))
	if v611 == int32(1) {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v607))))
	if v641 == int32(1) {
		goto L169
	} else {
		goto L170
	}
L158:
	;
	v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v605)+1)))
	if v617 == int32(18) {
		goto L161
	} else {
		goto L162
	}
L159:
	;
	goto L160
L160:
	;
	v628 = int32(1)
	if v611&v628 != 0 {
		v640 = int32(base.Ui32(v611)>>(uint(v628)%32)) - v628
		goto L157
	} else {
		goto L167
	}
L161:
	;
	v620 = int32(16)
	goto L163
L162:
	;
	v620 = int32(0)
	goto L163
L163:
	;
	if base.Ui32((v617-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v627 = int32(4)
	goto L166
L165:
	;
	v627 = v620
	goto L166
L166:
	;
	v640 = v627
	goto L157
L167:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v605)))
	v640 = int32(base.Ui32(v634)>>(uint(int32(2))%32)) - int32(4)
	goto L157
L168:
	;
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609))))
	if v671 == int32(1) {
		goto L180
	} else {
		goto L181
	}
L169:
	;
	v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v607)+1)))
	if v647 == int32(18) {
		goto L172
	} else {
		goto L173
	}
L170:
	;
	goto L171
L171:
	;
	v658 = int32(1)
	if v641&v658 != 0 {
		v670 = int32(base.Ui32(v641)>>(uint(v658)%32)) - v658
		goto L168
	} else {
		goto L178
	}
L172:
	;
	v650 = int32(16)
	goto L174
L173:
	;
	v650 = int32(0)
	goto L174
L174:
	;
	if base.Ui32((v647-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v657 = int32(4)
	goto L177
L176:
	;
	v657 = v650
	goto L177
L177:
	;
	v670 = v657
	goto L168
L178:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v607)))
	v670 = int32(base.Ui32(v664)>>(uint(int32(2))%32)) - int32(4)
	goto L168
L179:
	;
	v701 = int32(1)
	if v671&v701 != 0 {
		goto L190
	} else {
		goto L191
	}
L180:
	;
	v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609)+1)))
	if v677 == int32(18) {
		goto L183
	} else {
		goto L184
	}
L181:
	;
	goto L182
L182:
	;
	v688 = int32(1)
	if v671&v688 != 0 {
		v700 = int32(base.Ui32(v671)>>(uint(v688)%32)) - v688
		goto L179
	} else {
		goto L189
	}
L183:
	;
	v680 = int32(16)
	goto L185
L184:
	;
	v680 = int32(0)
	goto L185
L185:
	;
	if base.Ui32((v677-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v687 = int32(4)
	goto L188
L187:
	;
	v687 = v680
	goto L188
L188:
	;
	v700 = v687
	goto L179
L189:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v609)))
	v700 = int32(base.Ui32(v694)>>(uint(int32(2))%32)) - int32(4)
	goto L179
L190:
	;
	v705 = v701
	goto L192
L191:
	;
	v705 = int32(4)
	goto L192
L192:
	;
	v706 = v609 + v705
	v707 = int32(1)
	if v641&v707 != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v711 = v707
	goto L195
L194:
	;
	v711 = int32(4)
	goto L195
L195:
	;
	v712 = v607 + v711
	v713 = int32(1)
	if v611&v713 != 0 {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v717 = v713
	goto L198
L197:
	;
	v717 = int32(4)
	goto L198
L198:
	;
	v718 = v605 + v717
	if v640 < v670 {
		goto L200
	} else {
		goto L201
	}
L199:
	;
	if int32(0) < v789 {
		goto L212
	} else {
		goto L213
	}
L200:
	;
	v720 = v640
	goto L202
L201:
	;
	v720 = v670
	goto L202
L202:
	;
	if v720 < v700 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v722 = v720
	goto L205
L204:
	;
	v722 = v700
	goto L205
L205:
	;
	if v722 <= int32(0) {
		v788 = v706
		v789 = v640
		v791 = v712
		v793 = v700
		v794 = v670
		v795 = v718
		goto L199
	} else {
		goto L206
	}
L206:
	;
	v735 = v706
	v736 = v640
	v738 = v712
	v740 = v700
	v741 = v670
	v742 = v718
	v754 = int32(0)
	goto L207
L207:
	;
	v768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v738))))
	v769 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v735))))
	if v768 != v769 {
		v788 = v735
		v789 = v736
		v791 = v738
		v793 = v740
		v794 = v741
		v795 = v742
		goto L199
	} else {
		goto L209
	}
L208:
	;
	v788 = v705 + v609 + v722
	v789 = v640 - v722
	v791 = v711 + v607 + v722
	v793 = v700 - v722
	v794 = v670 - v722
	v795 = v717 + v605 + v722
	goto L199
L209:
	;
	v771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v742))))
	if v768 != v771 {
		v788 = v735
		v789 = v736
		v791 = v738
		v793 = v740
		v794 = v741
		v795 = v742
		goto L199
	} else {
		goto L210
	}
L210:
	;
	v773 = int32(1)
	v786 = v754 + v773
	if v786 != v722 {
		v735 = v735 + v773
		v736 = v736 - v773
		v738 = v738 + v773
		v740 = v740 - v773
		v741 = v741 - v773
		v742 = v742 + v773
		v754 = v786
		goto L207
	} else {
		goto L211
	}
L211:
	;
	goto L208
L212:
	;
	v823 = int32(10)
	if base.Ui32(v823) <= base.Ui32(v789) {
		goto L215
	} else {
		goto L216
	}
L213:
	;
	v904 = v28
	goto L214
L214:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v587))) = v904
	if int32(0) < v794 {
		goto L221
	} else {
		goto L222
	}
L215:
	;
	v826 = v823
	goto L217
L216:
	;
	v826 = v789
	goto L217
L217:
	;
	v829 = v826
	v835 = v795
	v856 = float64(256)
	v859 = v28
	goto L218
L218:
	;
	v861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v835))))
	v864 = base.F64_add(v859, base.F64_div(base.F64_convert_i32_u(v861), v856))
	v865 = int32(1)
	if base.Ui32(v865) < base.Ui32(v829) {
		v829 = v829 - v865
		v835 = v835 + v865
		v856 = base.F64_mul(v856, float64(256))
		v859 = v864
		goto L218
	} else {
		goto L220
	}
L219:
	;
	v904 = v864
	goto L214
L220:
	;
	goto L219
L221:
	;
	v909 = int32(10)
	if base.Ui32(v909) <= base.Ui32(v794) {
		goto L224
	} else {
		goto L225
	}
L222:
	;
	v988 = v28
	goto L223
L223:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v577))) = v988
	if v793 <= int32(0) {
		goto L231
	} else {
		goto L232
	}
L224:
	;
	v912 = v909
	goto L226
L225:
	;
	v912 = v794
	goto L226
L226:
	;
	v917 = v791
	v921 = v912
	v942 = float64(256)
	v943 = v28
	goto L227
L227:
	;
	v947 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v917))))
	v950 = base.F64_add(v943, base.F64_div(base.F64_convert_i32_u(v947), v942))
	v951 = int32(1)
	if base.Ui32(v951) < base.Ui32(v921) {
		v917 = v917 + v951
		v921 = v921 - v951
		v942 = base.F64_mul(v942, float64(256))
		v943 = v950
		goto L227
	} else {
		goto L229
	}
L228:
	;
	v988 = v950
	goto L223
L229:
	;
	goto L228
L230:
	;
	v1947 = int32(1)
	goto L128
L231:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v579))) = float64(0)
	goto L230
L232:
	;
	goto L233
L233:
	;
	v997 = int32(10)
	if base.Ui32(v997) <= base.Ui32(v793) {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v1000 = v997
	goto L236
L235:
	;
	v1000 = v793
	goto L236
L236:
	;
	v1003 = v788
	v1006 = v1000
	v1031 = float64(256)
	v1032 = float64(0)
	goto L237
L237:
	;
	v1036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1003))))
	v1039 = base.F64_add(v1032, base.F64_div(base.F64_convert_i32_u(v1036), v1031))
	v1040 = int32(1)
	if base.Ui32(v1040) < base.Ui32(v1006) {
		v1003 = v1003 + v1040
		v1006 = v1006 - v1040
		v1031 = base.F64_mul(v1031, float64(256))
		v1032 = v1039
		goto L237
	} else {
		goto L239
	}
L238:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v579))) = v1039
	goto L230
L239:
	;
	goto L238
L240:
	;
	if base.Ui32(l8-int32(700)) < base.Ui32(int32(2)) {
		goto L142
	} else {
		goto L243
	}
L241:
	;
	goto L242
L242:
	;
	if base.Ui32(l8-int32(1042)) < base.Ui32(int32(2)) {
		goto L141
	} else {
		goto L245
	}
L243:
	;
	if l8 == int32(774) {
		goto L130
	} else {
		goto L244
	}
L244:
	;
	goto L129
L245:
	;
	if l8 == int32(829) {
		goto L130
	} else {
		goto L246
	}
L246:
	;
	if l8 != int32(869) {
		goto L129
	} else {
		goto L247
	}
L247:
	;
	goto L130
L248:
	;
	if l8 <= int32(1183) {
		goto L251
	} else {
		goto L252
	}
L249:
	;
	goto L250
L250:
	;
	if l8 <= int32(3768) {
		goto L257
	} else {
		goto L258
	}
L251:
	;
	switch l8 - int32(1082) {
	case 0:
		goto L139
	case 1:
		goto L137
	default:
		goto L135
	}
L252:
	;
	goto L253
L253:
	;
	switch l8 - int32(1184) {
	case 0:
		goto L140
	case 1:
		goto L129
	case 2:
		goto L138
	default:
		goto L254
	}
L254:
	;
	if l8 == int32(1266) {
		goto L136
	} else {
		goto L255
	}
L255:
	;
	if l8 == int32(1700) {
		goto L142
	} else {
		goto L256
	}
L256:
	;
	goto L129
L257:
	;
	if base.B2i32(l8 == int32(3734))|base.B2i32(base.Ui32(l8-int32(2202)) < base.Ui32(int32(5))) != 0 {
		goto L142
	} else {
		goto L260
	}
L258:
	;
	goto L259
L259:
	;
	switch l8 - int32(4089) {
	case 0, 7:
		goto L142
	case 1, 2, 3, 4, 5, 6:
		goto L129
	default:
		goto L261
	}
L260:
	;
	goto L129
L261:
	;
	if l8 == int32(_a_F_ineq_histogram_selectivity_5) {
		goto L142
	} else {
		goto L262
	}
L262:
	;
	if l8 != int32(3769) {
		goto L129
	} else {
		goto L263
	}
L263:
	;
	goto L142
L264:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v587))) = v1129
	v1132 = F_convert_numeric_to_scalar(m, v570, v575, v1128)
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L13
	} else {
		goto L265
	}
L265:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v577))) = v1132
	v1135 = F_convert_numeric_to_scalar(m, v574, v575, v1128)
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L13
	} else {
		goto L266
	}
L266:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v579))) = v1135
	v1138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582)+15)))
	v1947 = v1138 ^ int32(1)
	goto L128
L267:
	;
	v1146 = F_convert_string_datum(m, v570, v575, l6, v1143)
	mBase = m.M
	v1147 = m.ExcPending
	if v1147 != 0 {
		goto L13
	} else {
		goto L268
	}
L268:
	;
	v1148 = F_convert_string_datum(m, v574, v575, l6, v1143)
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L13
	} else {
		goto L269
	}
L269:
	;
	v1150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582)+15)))
	if v1150 != 0 {
		goto L132
	} else {
		goto L270
	}
L270:
	;
	v1151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1148))))
	v1152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1146))))
	if v1152 == int32(0) {
		goto L271
	} else {
		goto L272
	}
L271:
	;
	v1340 = v1151
	v1347 = v1151
	goto L133
L272:
	;
	goto L273
L273:
	;
	v1156 = v1151
	v1158 = v1146
	v1162 = v1152
	v1163 = v1151
	goto L274
L274:
	;
	v1189 = v1162 & int32(255)
	if base.Ui32(v1189) < base.Ui32(v1156) {
		goto L276
	} else {
		goto L277
	}
L275:
	;
	v1340 = v1191
	v1347 = v1193
	goto L133
L276:
	;
	v1191 = v1156
	goto L278
L277:
	;
	v1191 = v1189
	goto L278
L278:
	;
	if base.Ui32(v1163) < base.Ui32(v1189) {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	v1193 = v1163
	goto L281
L280:
	;
	v1193 = v1189
	goto L281
L281:
	;
	v1194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1158)+1)))
	if v1194 != 0 {
		v1156 = v1191
		v1158 = v1158 + int32(1)
		v1162 = v1194
		v1163 = v1193
		goto L274
	} else {
		goto L282
	}
L282:
	;
	goto L275
L283:
	;
	v1234 = v1208
	goto L134
L284:
	;
	v1208 = float64(-1.7976931348623157e+308)
	goto L283
L285:
	;
	goto L286
L286:
	;
	if l7 == int32(2147483647) {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v1208 = float64(1.7976931348623157e+308)
	goto L283
L288:
	;
	goto L289
L289:
	;
	v1208 = base.F64_mul(base.F64_convert_i32_s(l7), float64(8.64e+10))
	goto L283
L290:
	;
	v1232 = *(*int64)(unsafe.Add(mBase, uint32(l7)))
	v1234 = base.F64_convert_i64_s(v1232)
	goto L134
L291:
	;
	v1332 = *(*int64)(unsafe.Add(mBase, uint32(v570)))
	*(*float64)(unsafe.Add(mBase, uint32(v577))) = base.F64_convert_i64_s(v1332)
	v1336 = *(*int64)(unsafe.Add(mBase, uint32(v574)))
	*(*float64)(unsafe.Add(mBase, uint32(v579))) = base.F64_convert_i64_s(v1336)
	v1947 = int32(1)
	goto L128
L292:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v577))) = int64(0)
	*(*float64)(unsafe.Add(mBase, uint32(v579))) = float64(0)
	v1947 = int32(0)
	goto L128
L293:
	;
	if v575 == int32(1114) {
		goto L291
	} else {
		goto L317
	}
L294:
	;
	if v575 != int32(1266) {
		goto L292
	} else {
		goto L316
	}
L295:
	;
	v1299 = *(*int64)(unsafe.Add(mBase, uint32(v570)))
	*(*float64)(unsafe.Add(mBase, uint32(v577))) = base.F64_convert_i64_s(v1299)
	v1303 = *(*int64)(unsafe.Add(mBase, uint32(v574)))
	*(*float64)(unsafe.Add(mBase, uint32(v579))) = base.F64_convert_i64_s(v1303)
	v1947 = int32(1)
	goto L128
L296:
	;
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v570)+12))
	v1274 = float64(2.6298e+12)
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v570)+8))
	v1278 = float64(8.64e+10)
	v1280 = *(*int64)(unsafe.Add(mBase, uint32(v570)))
	*(*float64)(unsafe.Add(mBase, uint32(v577))) = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1272), v1274), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1276), v1278), base.F64_convert_i64_s(v1280)))
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(v574)+12))
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(v574)+8))
	v1294 = *(*int64)(unsafe.Add(mBase, uint32(v574)))
	*(*float64)(unsafe.Add(mBase, uint32(v579))) = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1286), v1274), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1290), v1278), base.F64_convert_i64_s(v1294)))
	v1947 = int32(1)
	goto L128
L297:
	;
	if v570 == int32(-2147483648) {
		goto L303
	} else {
		goto L304
	}
L298:
	;
	switch v575 - int32(1082) {
	case 0:
		goto L297
	case 1:
		goto L295
	default:
		goto L293
	}
L299:
	;
	goto L300
L300:
	;
	switch v575 - int32(1184) {
	case 0:
		goto L301
	case 1:
		goto L292
	case 2:
		goto L296
	default:
		goto L294
	}
L301:
	;
	v1242 = *(*int64)(unsafe.Add(mBase, uint32(v570)))
	*(*float64)(unsafe.Add(mBase, uint32(v577))) = base.F64_convert_i64_s(v1242)
	v1246 = *(*int64)(unsafe.Add(mBase, uint32(v574)))
	*(*float64)(unsafe.Add(mBase, uint32(v579))) = base.F64_convert_i64_s(v1246)
	v1947 = int32(1)
	goto L128
L302:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v577))) = v1258
	if v574 == int32(-2147483648) {
		goto L310
	} else {
		goto L311
	}
L303:
	;
	v1258 = float64(-1.7976931348623157e+308)
	goto L302
L304:
	;
	goto L305
L305:
	;
	if v570 == int32(2147483647) {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	v1258 = float64(1.7976931348623157e+308)
	goto L302
L307:
	;
	goto L308
L308:
	;
	v1258 = base.F64_mul(base.F64_convert_i32_s(v570), float64(8.64e+10))
	goto L302
L309:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v579))) = v1270
	v1947 = int32(1)
	goto L128
L310:
	;
	v1270 = float64(-1.7976931348623157e+308)
	goto L309
L311:
	;
	goto L312
L312:
	;
	if v574 == int32(2147483647) {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	v1270 = float64(1.7976931348623157e+308)
	goto L309
L314:
	;
	goto L315
L315:
	;
	v1270 = base.F64_mul(base.F64_convert_i32_s(v574), float64(8.64e+10))
	goto L309
L316:
	;
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v570)+8))
	v1310 = float64(1e+06)
	v1312 = *(*int64)(unsafe.Add(mBase, uint32(v570)))
	*(*float64)(unsafe.Add(mBase, uint32(v577))) = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1308), v1310), base.F64_convert_i64_s(v1312))
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v574)+8))
	v1321 = *(*int64)(unsafe.Add(mBase, uint32(v574)))
	*(*float64)(unsafe.Add(mBase, uint32(v579))) = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1317), v1310), base.F64_convert_i64_s(v1321))
	v1947 = int32(1)
	goto L128
L317:
	;
	goto L292
L318:
	;
	v1372 = v1151
	v1373 = v1340
	v1375 = v1148
	v1380 = v1347
	goto L321
L319:
	;
	v1415 = v1340
	v1422 = v1347
	goto L320
L320:
	;
	v1449 = int32(90)
	if v1415 <= v1449 {
		goto L330
	} else {
		goto L331
	}
L321:
	;
	v1406 = v1372 & int32(255)
	if v1406 < v1373 {
		goto L323
	} else {
		goto L324
	}
L322:
	;
	v1415 = v1408
	v1422 = v1410
	goto L320
L323:
	;
	v1408 = v1373
	goto L325
L324:
	;
	v1408 = v1406
	goto L325
L325:
	;
	if v1380 < v1406 {
		goto L326
	} else {
		goto L327
	}
L326:
	;
	v1410 = v1380
	goto L328
L327:
	;
	v1410 = v1406
	goto L328
L328:
	;
	v1411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1375)+1)))
	if v1411 != 0 {
		v1372 = v1411
		v1373 = v1408
		v1375 = v1375 + int32(1)
		v1380 = v1410
		goto L321
	} else {
		goto L329
	}
L329:
	;
	goto L322
L330:
	;
	v1452 = v1449
	goto L332
L331:
	;
	v1452 = v1415
	goto L332
L332:
	;
	v1457 = base.B2i32(v1422 < int32(91)) & base.B2i32(int32(64) < v1415)
	if v1457 != 0 {
		goto L333
	} else {
		goto L334
	}
L333:
	;
	v1458 = v1452
	goto L335
L334:
	;
	v1458 = v1415
	goto L335
L335:
	;
	if v1458 <= int32(122) {
		goto L336
	} else {
		goto L337
	}
L336:
	;
	v1461 = int32(122)
	goto L338
L337:
	;
	v1461 = v1458
	goto L338
L338:
	;
	v1462 = int32(65)
	if v1462 <= v1422 {
		goto L339
	} else {
		goto L340
	}
L339:
	;
	v1465 = v1462
	goto L341
L340:
	;
	v1465 = v1422
	goto L341
L341:
	;
	if v1457 != 0 {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	v1466 = v1465
	goto L344
L343:
	;
	v1466 = v1422
	goto L344
L344:
	;
	v1471 = base.B2i32(v1466 < int32(123)) & base.B2i32(int32(96) < v1458)
	if v1471 != 0 {
		goto L345
	} else {
		goto L346
	}
L345:
	;
	v1472 = v1461
	goto L347
L346:
	;
	v1472 = v1458
	goto L347
L347:
	;
	if v1472 <= int32(57) {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	v1475 = int32(57)
	goto L350
L349:
	;
	v1475 = v1472
	goto L350
L350:
	;
	v1476 = int32(97)
	if v1476 <= v1466 {
		goto L351
	} else {
		goto L352
	}
L351:
	;
	v1479 = v1476
	goto L353
L352:
	;
	v1479 = v1466
	goto L353
L353:
	;
	if v1471 != 0 {
		goto L354
	} else {
		goto L355
	}
L354:
	;
	v1480 = v1479
	goto L356
L355:
	;
	v1480 = v1466
	goto L356
L356:
	;
	v1485 = base.B2i32(v1480 < int32(58)) & base.B2i32(int32(47) < v1472)
	if v1485 != 0 {
		goto L357
	} else {
		goto L358
	}
L357:
	;
	v1486 = v1475
	goto L359
L358:
	;
	v1486 = v1472
	goto L359
L359:
	;
	v1487 = int32(48)
	if v1487 <= v1480 {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	v1490 = v1487
	goto L362
L361:
	;
	v1490 = v1480
	goto L362
L362:
	;
	if v1485 != 0 {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	v1491 = v1490
	goto L365
L364:
	;
	v1491 = v1480
	goto L365
L365:
	;
	v1494 = base.B2i32(v1486-v1491 < int32(9))
	if v1152 == int32(0) {
		v1543 = v1146
		v1544 = v1148
		v1550 = v1144
		goto L366
	} else {
		goto L367
	}
L366:
	;
	if v1486-v1491 < int32(9) {
		goto L377
	} else {
		goto L378
	}
L367:
	;
	v1498 = v1148
	v1503 = v1152
	v1504 = v1144
	v1505 = v1146
	goto L368
L368:
	;
	v1531 = v1503 & int32(255)
	v1532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1498))))
	if v1531 != v1532 {
		goto L370
	} else {
		goto L371
	}
L369:
	;
	v1543 = v1542
	v1544 = v1539
	v1550 = v1537
	goto L366
L370:
	;
	v1543 = v1505
	v1544 = v1498
	v1550 = v1504
	goto L366
L371:
	;
	goto L372
L372:
	;
	v1534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1504))))
	if v1534 != v1531 {
		goto L373
	} else {
		goto L374
	}
L373:
	;
	v1543 = v1505
	v1544 = v1498
	v1550 = v1504
	goto L366
L374:
	;
	goto L375
L375:
	;
	v1536 = int32(1)
	v1537 = v1504 + v1536
	v1539 = v1498 + v1536
	v1540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1505)+1)))
	v1542 = v1505 + v1536
	if v1540 != 0 {
		v1498 = v1539
		v1503 = v1540
		v1504 = v1537
		v1505 = v1542
		goto L368
	} else {
		goto L376
	}
L376:
	;
	goto L369
L377:
	;
	v1577 = int32(127)
	goto L379
L378:
	;
	v1577 = v1486
	goto L379
L379:
	;
	if v1486-v1491 < int32(9) {
		goto L380
	} else {
		goto L381
	}
L380:
	;
	v1579 = int32(32)
	goto L382
L381:
	;
	v1579 = v1491
	goto L382
L382:
	;
	v1580 = F_strlen(m, v1550)
	mBase = m.M
	if int32(0) < v1580 {
		goto L383
	} else {
		goto L384
	}
L383:
	;
	v1583 = int32(12)
	if base.Ui32(v1583) <= base.Ui32(v1580) {
		goto L386
	} else {
		goto L387
	}
L384:
	;
	v1675 = v28
	goto L385
L385:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v587))) = v1675
	v1678 = F_strlen(m, v1543)
	mBase = m.M
	if int32(0) < v1678 {
		goto L398
	} else {
		goto L399
	}
L386:
	;
	v1586 = v1583
	goto L388
L387:
	;
	v1586 = v1580
	goto L388
L388:
	;
	v1587 = int32(1)
	v1594 = base.F64_convert_i32_s(v1577 - v1579 + v1587)
	v1601 = v1586
	v1602 = v1550
	v1623 = v1594
	v1626 = v28
	goto L389
L389:
	;
	v1628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1602))))
	if base.Ui32(v1577) < base.Ui32(v1628) {
		goto L391
	} else {
		goto L392
	}
L390:
	;
	v1675 = v1636
	goto L385
L391:
	;
	v1630 = v1577 + v1587
	goto L393
L392:
	;
	v1630 = v1628
	goto L393
L393:
	;
	if base.Ui32(v1628) < base.Ui32(v1579) {
		goto L394
	} else {
		goto L395
	}
L394:
	;
	v1632 = v1579 - v1587
	goto L396
L395:
	;
	v1632 = v1630
	goto L396
L396:
	;
	v1636 = base.F64_add(v1626, base.F64_div(base.F64_convert_i32_s(v1632-v1579), v1623))
	v1638 = int32(1)
	if base.Ui32(v1638) < base.Ui32(v1601) {
		v1601 = v1601 - v1638
		v1602 = v1602 + v1638
		v1623 = base.F64_mul(v1623, v1594)
		v1626 = v1636
		goto L389
	} else {
		goto L397
	}
L397:
	;
	goto L390
L398:
	;
	v1681 = int32(12)
	if base.Ui32(v1681) <= base.Ui32(v1678) {
		goto L401
	} else {
		goto L402
	}
L399:
	;
	v1769 = v28
	goto L400
L400:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v577))) = v1769
	v1774 = F_strlen(m, v1544)
	mBase = m.M
	if v1774 <= int32(0) {
		goto L414
	} else {
		goto L415
	}
L401:
	;
	v1684 = v1681
	goto L403
L402:
	;
	v1684 = v1678
	goto L403
L403:
	;
	v1685 = int32(1)
	v1688 = v1577 + v1685
	v1690 = base.F64_convert_i32_s(v1688 - v1579)
	v1691 = v1543
	v1698 = v1684
	v1719 = v1690
	v1720 = v28
	goto L404
L404:
	;
	v1724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1691))))
	if base.Ui32(v1577) < base.Ui32(v1724) {
		goto L406
	} else {
		goto L407
	}
L405:
	;
	v1769 = v1732
	goto L400
L406:
	;
	v1726 = v1688
	goto L408
L407:
	;
	v1726 = v1724
	goto L408
L408:
	;
	if base.Ui32(v1724) < base.Ui32(v1579) {
		goto L409
	} else {
		goto L410
	}
L409:
	;
	v1728 = v1579 - v1685
	goto L411
L410:
	;
	v1728 = v1726
	goto L411
L411:
	;
	v1732 = base.F64_add(v1720, base.F64_div(base.F64_convert_i32_s(v1728-v1579), v1719))
	v1734 = int32(1)
	if base.Ui32(v1734) < base.Ui32(v1698) {
		v1691 = v1691 + v1734
		v1698 = v1698 - v1734
		v1719 = base.F64_mul(v1719, v1690)
		v1720 = v1732
		goto L404
	} else {
		goto L412
	}
L412:
	;
	goto L405
L413:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v579))) = v1867
	F_pfree(m, v1144)
	mBase = m.M
	v1873 = m.ExcPending
	if v1873 != 0 {
		goto L13
	} else {
		goto L429
	}
L414:
	;
	v1867 = float64(0)
	goto L413
L415:
	;
	goto L416
L416:
	;
	v1778 = int32(12)
	if base.Ui32(v1778) <= base.Ui32(v1774) {
		goto L417
	} else {
		goto L418
	}
L417:
	;
	v1781 = v1778
	goto L419
L418:
	;
	v1781 = v1774
	goto L419
L419:
	;
	v1782 = int32(1)
	v1786 = v1577 + v1782
	v1788 = base.F64_convert_i32_s(v1786 - v1579)
	v1790 = v1544
	v1796 = v1781
	v1817 = v1788
	v1818 = float64(0)
	goto L420
L420:
	;
	v1822 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1790))))
	if base.Ui32(v1577) < base.Ui32(v1822) {
		goto L422
	} else {
		goto L423
	}
L421:
	;
	v1867 = v1830
	goto L413
L422:
	;
	v1824 = v1786
	goto L424
L423:
	;
	v1824 = v1822
	goto L424
L424:
	;
	if base.Ui32(v1822) < base.Ui32(v1579) {
		goto L425
	} else {
		goto L426
	}
L425:
	;
	v1826 = v1579 - v1782
	goto L427
L426:
	;
	v1826 = v1824
	goto L427
L427:
	;
	v1830 = base.F64_add(v1818, base.F64_div(base.F64_convert_i32_s(v1826-v1579), v1817))
	v1832 = int32(1)
	if base.Ui32(v1832) < base.Ui32(v1796) {
		v1790 = v1790 + v1832
		v1796 = v1796 - v1832
		v1817 = base.F64_mul(v1817, v1788)
		v1818 = v1830
		goto L420
	} else {
		goto L428
	}
L428:
	;
	goto L421
L429:
	;
	F_pfree(m, v1146)
	mBase = m.M
	v1875 = m.ExcPending
	if v1875 != 0 {
		goto L13
	} else {
		goto L430
	}
L430:
	;
	F_pfree(m, v1148)
	mBase = m.M
	v1877 = m.ExcPending
	if v1877 != 0 {
		goto L13
	} else {
		goto L431
	}
L431:
	;
	goto L132
L432:
	;
	goto L130
L433:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v587))) = v1919
	v1922 = F_convert_network_to_scalar(m, v570, v575, v1918)
	mBase = m.M
	v1923 = m.ExcPending
	if v1923 != 0 {
		goto L13
	} else {
		goto L434
	}
L434:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v577))) = v1922
	v1925 = F_convert_network_to_scalar(m, v574, v575, v1918)
	mBase = m.M
	v1926 = m.ExcPending
	if v1926 != 0 {
		goto L13
	} else {
		goto L435
	}
L435:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v579))) = v1925
	v1928 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582)+15)))
	v1947 = v1928 ^ int32(1)
	goto L128
L436:
	;
	v2009 = float64(1)
	v2012 = *(*int32)(unsafe.Add(mBase, uint32(v36)+92))
	v2013 = int32(1)
	v2016 = base.F64_div(base.F64_add(v2007, base.F64_convert_i32_u(v566)), base.F64_convert_i32_s(v2012-v2013))
	if v449 != v2013 {
		goto L454
	} else {
		goto L455
	}
L437:
	;
	v1979 = *(*float64)(unsafe.Add(mBase, uint32(v36)+64))
	v1980 = *(*float64)(unsafe.Add(mBase, uint32(v36)+56))
	if base.F64_le(v1979, v1980) != 0 {
		v2007 = v563
		goto L436
	} else {
		goto L438
	}
L438:
	;
	v1982 = *(*float64)(unsafe.Add(mBase, uint32(v36)+16))
	if base.F64_ge(v1980, v1982) != 0 {
		goto L439
	} else {
		goto L440
	}
L439:
	;
	v2007 = float64(0)
	goto L436
L440:
	;
	goto L441
L441:
	;
	if base.F64_ge(v1982, v1979) != 0 {
		goto L442
	} else {
		goto L443
	}
L442:
	;
	v2007 = float64(1)
	goto L436
L443:
	;
	goto L444
L444:
	;
	v1987 = float64(0.5)
	v1992 = base.F64_div(base.F64_sub(v1982, v1980), base.F64_sub(v1979, v1980))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1992)&int64(9223372036854775807)) {
		goto L445
	} else {
		goto L446
	}
L445:
	;
	v1998 = v1987
	goto L447
L446:
	;
	v1998 = v1992
	goto L447
L447:
	;
	if base.F64_lt(v1992, float64(0)) != 0 {
		goto L448
	} else {
		goto L449
	}
L448:
	;
	v2001 = v1987
	goto L450
L449:
	;
	v2001 = v1998
	goto L450
L450:
	;
	if base.F64_gt(v1992, float64(1)) != 0 {
		goto L451
	} else {
		goto L452
	}
L451:
	;
	v2004 = v1987
	goto L453
L452:
	;
	v2004 = v2001
	goto L453
L453:
	;
	v2007 = v2004
	goto L436
L454:
	;
	v2023 = v2016
	goto L456
L455:
	;
	v2023 = base.F64_add(v2016, base.F64_mul(v562, base.F64_sub(v2009, v2007)))
	goto L456
L456:
	;
	if v458 != 0 {
		goto L457
	} else {
		goto L458
	}
L457:
	;
	v2025 = v2023
	goto L459
L458:
	;
	v2025 = base.F64_sub(v2023, v562)
	goto L459
L459:
	;
	if l4 != 0 {
		goto L460
	} else {
		goto L461
	}
L460:
	;
	v2027 = base.F64_sub(v2009, v2025)
	goto L462
L461:
	;
	v2027 = v2025
	goto L462
L462:
	;
	v2062 = v2027
	goto L74
L463:
	;
	v2031 = base.F64_sub(float64(1), v2028)
	goto L465
L464:
	;
	v2031 = v2028
	goto L465
L465:
	;
	v2062 = v2031
	goto L74
L466:
	;
	v2069 = float64(0)
	if base.F64_lt(v2062, v2069) != 0 {
		v2149 = v2069
		goto L3
	} else {
		goto L467
	}
L467:
	;
	if base.F64_gt(v2062, float64(1)) == int32(0) {
		v2149 = v2062
		goto L3
	} else {
		goto L468
	}
L468:
	;
	v2149 = float64(1)
	goto L3
L469:
	;
	v2118 = base.F64_sub(float64(1), v2115)
	if base.F64_lt(v2118, v2107) == int32(0) {
		v2149 = v2107
		goto L3
	} else {
		goto L470
	}
L470:
	;
	v2149 = v2118
	goto L3
L471:
	;
	v2186 = v2149
	goto L1
}
func F_inetmi_int8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
		v11 = F_internal_inetpl(m, v3, int64(0)-v9)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return v11
		}
	}
}
func F_infix_1(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
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
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v390 int32
	_ = v390
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v455 int32
	_ = v455
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v644 int32
	_ = v644
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	v14 = m.G0
	v16 = v14 - int32(96)
	m.G0 = v16
	F_check_stack_depth(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v21 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v16 + int32(96)
	return
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_infix_1[0]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v33*int32(28))+uint32(_c_F_infix_1[1])))
	goto L7
L5:
	;
	goto L6
L6:
	;
	v224 = int32(*(*int8)(unsafe.Add(mBase, uint32(v20)+1)))
	v226 = v224 & int32(255)
	if v226 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L7:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v47 <= v28-v29+(v38+int32(1))*(v25&int32(4095))+int32(9) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v51 = v47
	goto L11
L9:
	;
	goto L10
L10:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v107 = int32(39)
	*(*uint8)(unsafe.Add(mBase, uint32(v106))) = uint8(v107)
	v112 = v24 + int32(base.Ui32(v25)>>(uint(int32(12))%32))
	v113 = int32(1)
	goto L16
L11:
	;
	v63 = v51 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v63
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v67 = F_repalloc(m, v66, v63)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L10
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v67
	v70 = v65 - v66
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v67 + v70
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_infix_1[0]))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v76*int32(28))+uint32(_c_F_infix_1[1])))
	goto L14
L14:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v90 <= v70+(v81+int32(1))*(v73&int32(4095))+int32(9) {
		v51 = v90
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v124 = v123 + v113
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v124
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if base.B2i32(v126 == int32(39))|base.B2i32(v126 == int32(92)) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v220 = F_pg_mblen_cstr(m, v112)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L41
	}
L19:
	;
	if v126 != 0 {
		v219 = v124
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v124))) = uint8(v126)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v217 = v215 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v217
	v219 = v217
	goto L18
L22:
	;
	v134 = int32(39)
	*(*uint8)(unsafe.Add(mBase, uint32(v124))) = uint8(v134)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v138 = v136 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v138
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v140 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v208 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v206))) = uint8(v208)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v210 + int32(12)
	goto L3
L24:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+2)))
	if v143 != int32(1) {
		v206 = v138
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v146 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v138))) = uint8(v146)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v149 = int32(1)
	v150 = v148 + v149
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v150
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+2)))
	if v152 == v149 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L26
L28:
	;
	v155 = int32(42)
	*(*uint8)(unsafe.Add(mBase, uint32(v150))) = uint8(v155)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v159 = v157 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v159
	v161 = v159
	goto L30
L29:
	;
	v161 = v150
	goto L30
L30:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v162&int32(8) != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v165 = int32(65)
	*(*uint8)(unsafe.Add(mBase, uint32(v161))) = uint8(v165)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v169 = v167 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v169
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	v172 = v169
	v173 = v171
	goto L33
L32:
	;
	v172 = v161
	v173 = v162
	goto L33
L33:
	;
	if v173&int32(4) != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v176 = int32(66)
	*(*uint8)(unsafe.Add(mBase, uint32(v172))) = uint8(v176)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v180 = v178 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v180
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	v183 = v180
	v184 = v182
	goto L36
L35:
	;
	v183 = v172
	v184 = v173
	goto L36
L36:
	;
	if v184&int32(2) != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v187 = int32(67)
	*(*uint8)(unsafe.Add(mBase, uint32(v183))) = uint8(v187)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v191 = v189 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v191
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	v194 = v191
	v195 = v193
	goto L39
L38:
	;
	v194 = v183
	v195 = v184
	goto L39
L39:
	;
	if v195&int32(1) == int32(0) {
		v206 = v194
		goto L23
	} else {
		goto L40
	}
L40:
	;
	v200 = int32(68)
	*(*uint8)(unsafe.Add(mBase, uint32(v194))) = uint8(v200)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v204 = v202 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v204
	v206 = v204
	goto L23
L41:
	;
	if v220 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	base.MemoryCopy(m, v219, v112, v220)
	goto L44
L43:
	;
	goto L44
L44:
	;
	v112 = v112 + v220
	v113 = v220
	goto L16
L45:
	;
	if l1 <= int32(4) {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	goto L47
L47:
	;
	v407 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20)+2)))
	v409 = v20 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v409
	v412 = base.B2i32(v224 == int32(4))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v224<<(uint(int32(2))%32))+uint32(_c_F_infix_1[2])))
	v420 = l2&v412 | base.B2i32(v418 < l1)
	if v420 != 0 {
		goto L77
	} else {
		goto L78
	}
L48:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v297 = v287 - v296
	v299 = v297 + int32(2)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v300 <= v299 {
		goto L60
	} else {
		goto L61
	}
L49:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v287 = v231
	goto L48
L50:
	;
	goto L51
L51:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v234 = v232 - v233
	v236 = v234 + int32(3)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v237 <= v236 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v241 = v233
	v242 = v237
	goto L55
L53:
	;
	v266 = v232
	goto L54
L54:
	;
	v277 = F_pg_sprintf(m, v266, int32(_a_F_infix_1_0), int32(0))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L59
	}
L55:
	;
	v253 = v242 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v253
	v255 = F_repalloc(m, v241, v253)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	v266 = v258
	goto L54
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v255
	v258 = v255 + v234
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v258
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v260 <= v236 {
		v241 = v255
		v242 = v260
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v280 = F_strlen(m, v279)
	mBase = m.M
	v281 = v280 + v279
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v281
	v287 = v281
	goto L48
L60:
	;
	v304 = v296
	v305 = v300
	goto L63
L61:
	;
	v329 = v287
	goto L62
L62:
	;
	v338 = int32(33)
	*(*uint8)(unsafe.Add(mBase, uint32(v329))) = uint8(v338)
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v340 + int32(1)
	v344 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v340)+1)) = uint8(v344)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v346 + int32(12)
	F_infix_1(m, l0, int32(4), v344)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L67
	}
L63:
	;
	v316 = v305 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v316
	v318 = F_repalloc(m, v304, v316)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L65
	}
L64:
	;
	v329 = v321
	goto L62
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v318
	v321 = v318 + v297
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v321
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v323 <= v299 {
		v304 = v318
		v305 = v323
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	if l1 < int32(5) {
		goto L3
	} else {
		goto L68
	}
L68:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v358 = v356 - v357
	v360 = v358 + int32(3)
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v361 <= v360 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v365 = v357
	v366 = v361
	goto L72
L70:
	;
	v390 = v356
	goto L71
L71:
	;
	v401 = F_pg_sprintf(m, v390, int32(_a_F_infix_1_1), int32(0))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L76
	}
L72:
	;
	v377 = v366 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v377
	v379 = F_repalloc(m, v365, v377)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L74
	}
L73:
	;
	v390 = v382
	goto L71
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v379
	v382 = v358 + v379
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v382
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v384 <= v360 {
		v365 = v379
		v366 = v384
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v404 = F_strlen(m, v403)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v404 + v403
	goto L3
L77:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v423 = v421 - v422
	v425 = v423 + int32(3)
	v426 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v426 <= v425 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	v486 = v409
	goto L79
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v486
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v489 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+92)) = v489
	*(*int32)(unsafe.Add(mBase, uint32(v16)+88)) = v488
	v493 = F_palloc(m, v489)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L88
	}
L80:
	;
	v430 = v422
	v431 = v426
	goto L83
L81:
	;
	v455 = v421
	goto L82
L82:
	;
	v466 = F_pg_sprintf(m, v455, int32(_a_F_infix_1_0), int32(0))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L87
	}
L83:
	;
	v442 = v431 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v442
	v444 = F_repalloc(m, v430, v442)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L85
	}
L84:
	;
	v455 = v447
	goto L82
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v444
	v447 = v423 + v444
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v447
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v449 <= v425 {
		v430 = v444
		v431 = v449
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v469 = F_strlen(m, v468)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v469 + v468
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v486 = v472
	goto L79
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+84)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v493
	F_infix_1(m, v16+int32(76), v418, v412)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v16)+76))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v501
	F_infix_1(m, l0, v418, int32(0))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v509 = v507 - v508
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v16)+84))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v16)+80))
	if v506 <= v509+v510-v512+int32(16) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v521 = v508
	v522 = v506
	goto L94
L92:
	;
	v547 = v512
	v550 = v507
	goto L93
L93:
	;
	switch v226 - int32(2) {
	case 0:
		goto L102
	case 1:
		goto L99
	case 2:
		goto L101
	default:
		goto L100
	}
L94:
	;
	v533 = v522 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v533
	v535 = F_repalloc(m, v521, v533)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L96
	}
L95:
	;
	v547 = v543
	v550 = v538
	goto L93
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v535
	v538 = v535 + v509
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v538
	v540 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v16)+84))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v16)+80))
	if v540 <= v509+int32(16)+v541-v543 {
		v521 = v535
		v522 = v540
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v602 = F_strlen(m, v601)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v602 + v601
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v16)+80))
	F_pfree(m, v605)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L1
	} else {
		goto L113
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v547
	v599 = F_pg_sprintf(m, v550, int32(_a_F_infix_1_2), v16+int32(16))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
	} else {
		goto L112
	}
L100:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L109
	}
L101:
	;
	if v407 != int32(1) {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v547
	v565 = F_pg_sprintf(m, v550, int32(_a_F_infix_1_3), v16+int32(32))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	goto L98
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v547
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v407
	v574 = F_pg_sprintf(m, v550, int32(_a_F_infix_1_4), v16-int32(-64))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v547
	v580 = F_pg_sprintf(m, v550, int32(_a_F_infix_1_5), v16+int32(48))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L1
	} else {
		goto L108
	}
L107:
	;
	goto L98
L108:
	;
	goto L98
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v224
	F_errmsg_internal(m, int32(_a_F_infix_1_6), v16)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_infix_1_7), int32(1130), int32(_a_F_infix_1_8))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	goto L98
L113:
	;
	if v420 == int32(0) {
		goto L3
	} else {
		goto L114
	}
L114:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v611 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v612 = v610 - v611
	v614 = v612 + int32(3)
	v615 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v615 <= v614 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v619 = v611
	v620 = v615
	goto L118
L116:
	;
	v644 = v610
	goto L117
L117:
	;
	v655 = F_pg_sprintf(m, v644, int32(_a_F_infix_1_1), int32(0))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L1
	} else {
		goto L122
	}
L118:
	;
	v631 = v620 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v631
	v633 = F_repalloc(m, v619, v631)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L1
	} else {
		goto L120
	}
L119:
	;
	v644 = v636
	goto L117
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v633
	v636 = v612 + v633
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v636
	v638 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v638 <= v614 {
		v619 = v633
		v620 = v638
		goto L118
	} else {
		goto L121
	}
L121:
	;
	goto L119
L122:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v658 = F_strlen(m, v657)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v658 + v657
	goto L3
}
func F_initClosestMatch(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	return
}
func F_initGISTstate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v80 int32
	_ = v80
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int64
	_ = v106
	var v108 int32
	_ = v108
	var v110 int64
	_ = v110
	var v112 int64
	_ = v112
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int64
	_ = v121
	var v123 int32
	_ = v123
	var v125 int64
	_ = v125
	var v127 int64
	_ = v127
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int64
	_ = v152
	var v154 int32
	_ = v154
	var v156 int64
	_ = v156
	var v158 int64
	_ = v158
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int64
	_ = v186
	var v188 int32
	_ = v188
	var v190 int64
	_ = v190
	var v192 int64
	_ = v192
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int64
	_ = v204
	var v206 int32
	_ = v206
	var v208 int64
	_ = v208
	var v210 int64
	_ = v210
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int64
	_ = v219
	var v221 int32
	_ = v221
	var v223 int64
	_ = v223
	var v225 int64
	_ = v225
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int64
	_ = v234
	var v236 int32
	_ = v236
	var v238 int64
	_ = v238
	var v240 int64
	_ = v240
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int64
	_ = v265
	var v267 int32
	_ = v267
	var v269 int64
	_ = v269
	var v271 int64
	_ = v271
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int64
	_ = v299
	var v301 int32
	_ = v301
	var v303 int64
	_ = v303
	var v305 int64
	_ = v305
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v372 int32
	_ = v372
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	v2 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(16)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v25 < int32(33) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	if v348 < v362 {
		goto L57
	} else {
		goto L58
	}
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_initGISTstate[0]))
	v34 = F_AllocSetContextCreateInternal(m, v29, int32(_a_F_initGISTstate_0), int32(0), int32(_a_F_initGISTstate_1), int32(_a_F_initGISTstate_2))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L5
	} else {
		goto L54
	}
L5:
	;
	return int32(0)
L6:
	;
	v38 = int32(_a_F_initGISTstate_3)
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_initGISTstate[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_initGISTstate[0])) = v34
	v43 = F_palloc(m, int32(_a_F_initGISTstate_4))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v34
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+8)) = v47
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v50 = int32(*(*int16)(unsafe.Add(mBase, uint32(v49)+10)))
	v51 = F_CreateTupleDescTruncatedCopy(m, v47, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+12)) = v51
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v55 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54)+10)))
	if v55 <= int32(0) {
		v348 = v2
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v80 = v2
	goto L10
L10:
	;
	v98 = v80 * int32(28)
	v99 = v43 + int32(20) + v98
	v100 = int32(1)
	v101 = v80 + v100
	v102 = base.I32_extend16_s(v101)
	v104 = F_index_getprocinfo(m, l0, v102, v100)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L5
	} else {
		goto L12
	}
L11:
	;
	v348 = v101
	goto L1
L12:
	;
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v104)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v99)+16)) = v106
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v99)+24)) = v108
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v104)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v99)+8)) = v110
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v104)))
	*(*int64)(unsafe.Add(mBase, uint32(v99))) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v99)+20)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v99)+16)) = int32(0)
	goto L13
L13:
	;
	v117 = v98 + (v43 + int32(916))
	v119 = F_index_getprocinfo(m, l0, v102, int32(2))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v121 = *(*int64)(unsafe.Add(mBase, uint32(v119)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v117)+16)) = v121
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v119)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v117)+24)) = v123
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v119)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v117)+8)) = v125
	v127 = *(*int64)(unsafe.Add(mBase, uint32(v119)))
	*(*int64)(unsafe.Add(mBase, uint32(v117))) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v117)+20)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v117)+16)) = int32(0)
	goto L15
L15:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v134)+6)))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v133+v135*(v102-int32(1))<<(uint(int32(2))%32)+int32(12)-int32(4))))
	goto L17
L16:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v168)+6)))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v167+v169*(v102-int32(1))<<(uint(int32(2))%32)+int32(16)-int32(4))))
	goto L24
L17:
	;
	if v147 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v148 = v98 + (v43 + int32(1812))
	v150 = F_index_getprocinfo(m, l0, v102, int32(3))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L5
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43+v98)+1816)) = int32(0)
	goto L16
L21:
	;
	v152 = *(*int64)(unsafe.Add(mBase, uint32(v150)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v148)+16)) = v152
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v150)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v148)+24)) = v154
	v156 = *(*int64)(unsafe.Add(mBase, uint32(v150)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v148)+8)) = v156
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v150)))
	*(*int64)(unsafe.Add(mBase, uint32(v148))) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v148)+20)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v148)+16)) = int32(0)
	goto L22
L22:
	;
	goto L16
L23:
	;
	v200 = v98 + (v43 + int32(3604))
	v202 = F_index_getprocinfo(m, l0, v102, int32(5))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L5
	} else {
		goto L30
	}
L24:
	;
	if v181 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v182 = v98 + (v43 + int32(2708))
	v184 = F_index_getprocinfo(m, l0, v102, int32(4))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L5
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43+v98)+2712)) = int32(0)
	goto L23
L28:
	;
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v184)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v182)+16)) = v186
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v184)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v182)+24)) = v188
	v190 = *(*int64)(unsafe.Add(mBase, uint32(v184)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v182)+8)) = v190
	v192 = *(*int64)(unsafe.Add(mBase, uint32(v184)))
	*(*int64)(unsafe.Add(mBase, uint32(v182))) = v192
	*(*int32)(unsafe.Add(mBase, uint32(v182)+20)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v182)+16)) = int32(0)
	goto L29
L29:
	;
	goto L23
L30:
	;
	v204 = *(*int64)(unsafe.Add(mBase, uint32(v202)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v200)+16)) = v204
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v202)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v200)+24)) = v206
	v208 = *(*int64)(unsafe.Add(mBase, uint32(v202)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v200)+8)) = v208
	v210 = *(*int64)(unsafe.Add(mBase, uint32(v202)))
	*(*int64)(unsafe.Add(mBase, uint32(v200))) = v210
	*(*int32)(unsafe.Add(mBase, uint32(v200)+20)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v200)+16)) = int32(0)
	goto L31
L31:
	;
	v215 = v98 + (v43 + int32(_a_F_initGISTstate_5))
	v217 = F_index_getprocinfo(m, l0, v102, int32(6))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v219 = *(*int64)(unsafe.Add(mBase, uint32(v217)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v215)+16)) = v219
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v217)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v215)+24)) = v221
	v223 = *(*int64)(unsafe.Add(mBase, uint32(v217)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v215)+8)) = v223
	v225 = *(*int64)(unsafe.Add(mBase, uint32(v217)))
	*(*int64)(unsafe.Add(mBase, uint32(v215))) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v215)+20)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v215)+16)) = int32(0)
	goto L33
L33:
	;
	v230 = v98 + (v43 + int32(_a_F_initGISTstate_6))
	v232 = F_index_getprocinfo(m, l0, v102, int32(7))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v234 = *(*int64)(unsafe.Add(mBase, uint32(v232)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v230)+16)) = v234
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v232)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v230)+24)) = v236
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v232)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v230)+8)) = v238
	v240 = *(*int64)(unsafe.Add(mBase, uint32(v232)))
	*(*int64)(unsafe.Add(mBase, uint32(v230))) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v230)+20)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v230)+16)) = int32(0)
	goto L35
L35:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v248 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v247)+6)))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v246+v248*(v102-int32(1))<<(uint(int32(2))%32)+int32(32)-int32(4))))
	goto L37
L36:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v282 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v281)+6)))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v280+v282*(v102-int32(1))<<(uint(int32(2))%32)+int32(36)-int32(4))))
	goto L44
L37:
	;
	if v260 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v261 = v98 + (v43 + int32(_a_F_initGISTstate_7))
	v263 = F_index_getprocinfo(m, l0, v102, int32(8))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L5
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43+v98)+uint32(_c_F_initGISTstate[1]))) = int32(0)
	goto L36
L41:
	;
	v265 = *(*int64)(unsafe.Add(mBase, uint32(v263)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v261)+16)) = v265
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v263)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v261)+24)) = v267
	v269 = *(*int64)(unsafe.Add(mBase, uint32(v263)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v261)+8)) = v269
	v271 = *(*int64)(unsafe.Add(mBase, uint32(v263)))
	*(*int64)(unsafe.Add(mBase, uint32(v261))) = v271
	*(*int32)(unsafe.Add(mBase, uint32(v261)+20)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v261)+16)) = int32(0)
	goto L42
L42:
	;
	goto L36
L43:
	;
	v314 = v80 << (uint(int32(2)) % 32)
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v316+v314)))
	if v318 != 0 {
		goto L50
	} else {
		goto L51
	}
L44:
	;
	if v294 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v295 = v98 + (v43 + int32(_a_F_initGISTstate_8))
	v297 = F_index_getprocinfo(m, l0, v102, int32(9))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L5
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43+v98)+uint32(_c_F_initGISTstate[2]))) = int32(0)
	goto L43
L48:
	;
	v299 = *(*int64)(unsafe.Add(mBase, uint32(v297)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v295)+16)) = v299
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v297)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v295)+24)) = v301
	v303 = *(*int64)(unsafe.Add(mBase, uint32(v297)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v295)+8)) = v303
	v305 = *(*int64)(unsafe.Add(mBase, uint32(v297)))
	*(*int64)(unsafe.Add(mBase, uint32(v295))) = v305
	*(*int32)(unsafe.Add(mBase, uint32(v295)+20)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v295)+16)) = int32(0)
	goto L49
L49:
	;
	goto L43
L50:
	;
	v320 = v318
	goto L52
L51:
	;
	v320 = int32(100)
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43+int32(_a_F_initGISTstate_9)+v314))) = v320
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v323 = int32(*(*int16)(unsafe.Add(mBase, uint32(v322)+10)))
	if v101 < v323 {
		v80 = v101
		goto L10
	} else {
		goto L53
	}
L53:
	;
	goto L11
L54:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v329)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v330
	F_errmsg_internal(m, int32(_a_F_initGISTstate_10), v22)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_initGISTstate_11), int32(1547), int32(_a_F_initGISTstate_12))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	v372 = v348
	goto L60
L58:
	;
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_initGISTstate[0])) = v39
	m.G0 = v22 + int32(16)
	return v43
L60:
	;
	v387 = v43 + v372*int32(28)
	v390 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v387)+uint32(_c_F_initGISTstate[2]))) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v387)+uint32(_c_F_initGISTstate[1]))) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v387)+uint32(_c_F_initGISTstate[3]))) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v387)+uint32(_c_F_initGISTstate[4]))) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v387+int32(3608)))) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v387+int32(2712)))) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v387+int32(1816)))) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v387)+920)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v387)+24)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v43+int32(_a_F_initGISTstate_9)+v372<<(uint(int32(2))%32)))) = v390
	v426 = v372 + int32(1)
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v427)))
	if v426 < v428 {
		v372 = v426
		goto L60
	} else {
		goto L62
	}
L61:
	;
	goto L59
L62:
	;
	goto L61
}
func F_initHyperLogLog(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 float64
	_ = v48
	var v49 float64
	_ = v49
	v2 = l1
	if base.Ui32(int32(242)) < base.Ui32((v2-int32(17))&int32(255)) {
		*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v2)
		v11 = int32(1)
		v12 = v11 << (uint(v2) % 32)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12
		v15 = v12 + v11
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v15
		v17 = F_palloc0(m, v15)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v17
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			switch v21 - int32(16) {
			case 0:
				v48 = float64(0.673)
			case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
				v48 = base.F64_div(float64(0.7213), base.F64_add(base.F64_div(float64(1.079), base.F64_convert_i32_u(v21)), float64(1)))
			case 16:
				v48 = float64(0.697)
			default:
				if v21 == int32(64) {
					v48 = float64(0.709)
				} else {
					v48 = base.F64_div(float64(0.7213), base.F64_add(base.F64_div(float64(1.079), base.F64_convert_i32_u(v21)), float64(1)))
				}
			}
			v49 = base.F64_convert_i32_u(v21)
			*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = base.F64_mul(base.F64_mul(v48, v49), v49)
			return
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_initHyperLogLog_0), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_initHyperLogLog_1), int32(71), int32(_a_F_initHyperLogLog_2))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
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
func F_initialize(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v194 int32
	_ = v194
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v268 int64
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v303 int64
	_ = v303
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v337 int32
	_ = v337
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v535 int32
	_ = v535
	var v542 int32
	_ = v542
	var v550 int32
	_ = v550
	var v561 int32
	_ = v561
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v575 int32
	_ = v575
	v4 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v4 < v12 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v575)+20)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(l1)+48)) = int64(0)
	return v575
L2:
	;
	v550 = int32(0)
	goto L111
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
	if v16&int32(1) != 0 {
		v542 = v15
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v26 < v27 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L5
L7:
	;
	if v390 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L8:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v194)+12))
	if v202 != 0 {
		goto L48
	} else {
		goto L49
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v26 + int32(1)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v34 = int32(0)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v38 = v35 + v26<<(uint(int32(5))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v38)+16)) = uint16(v34)
	*(*int64)(unsafe.Add(mBase, uint32(v38)+8)) = int64(0)
	v44 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v32 + v26*v33<<(uint(v44)%32)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+24)) = v48 + v49*v26<<(uint(v44)%32)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+28)) = v55 + v56*v26<<(uint(int32(3))%32)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v62 <= v34 {
		v194 = v38
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v94 = base.I32_div_s(v27<<(uint(int32(1))%32), int32(3))
	v95 = int32(2)
	if v94 < (l2-l2)>>(uint(v95)%32) {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v68 = v34
	goto L13
L13:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	v79 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v75+v68<<(uint(int32(2))%32)))) = v79
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v38)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v81+v68<<(uint(int32(3))%32)))) = v79
	v88 = v68 + int32(1)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v88 < v89 {
		v68 = v88
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v194 = v38
	goto L8
L15:
	;
	goto L14
L16:
	;
	v102 = l2 - v94<<(uint(v95)%32)
	goto L18
L17:
	;
	v102 = l2
	goto L18
L18:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v107 = v104 + v27<<(uint(int32(5))%32)
	if base.Ui32(v103) < base.Ui32(v107) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+56)) = v181 + int32(32)
	v194 = v181
	goto L8
L20:
	;
	v111 = v103
	goto L23
L21:
	;
	goto L22
L22:
	;
	if base.Ui32(v104) < base.Ui32(v103) {
		goto L33
	} else {
		goto L34
	}
L23:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	if base.Ui32(v102) <= base.Ui32(v119) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L22
L25:
	;
	v122 = v119
	goto L27
L26:
	;
	v122 = int32(0)
	goto L27
L27:
	;
	if v122 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+8)))
	if v125&int32(4) == int32(0) {
		v181 = v111
		goto L19
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v131 = v111 + int32(32)
	if base.Ui32(v131) < base.Ui32(v107) {
		v111 = v131
		goto L23
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	goto L24
L33:
	;
	v147 = v104
	goto L36
L34:
	;
	goto L35
L35:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v174 != 0 {
		goto L45
	} else {
		goto L46
	}
L36:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v147)+20))
	if base.Ui32(v102) <= base.Ui32(v154) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L35
L38:
	;
	v162 = v147 + int32(32)
	if base.Ui32(v162) < base.Ui32(v103) {
		v147 = v162
		goto L36
	} else {
		goto L44
	}
L39:
	;
	v157 = v154
	goto L41
L40:
	;
	v157 = int32(0)
	goto L41
L41:
	;
	if v157 != 0 {
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+8)))
	if v158&int32(4) != 0 {
		goto L38
	} else {
		goto L43
	}
L43:
	;
	v181 = v147
	goto L19
L44:
	;
	goto L37
L45:
	;
	v176 = v174
	goto L47
L46:
	;
	v176 = int32(15)
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v176
	v390 = int32(0)
	goto L7
L48:
	;
	v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v194)+16)))
	v204 = v203
	v207 = v202
	goto L51
L49:
	;
	goto L50
L50:
	;
	v239 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v194)+12)) = v239
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v239 < v241 {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v207)+24))
	v215 = base.I32_extend16_s(v204)
	v219 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v214+v215<<(uint(int32(2))%32)))) = v219
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v207)+28))
	v224 = v221 + v215<<(uint(int32(3))%32)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	*(*int32)(unsafe.Add(mBase, uint32(v224))) = v219
	v228 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v224)+4)))
	if v225 != 0 {
		v204 = v228
		v207 = v225
		goto L51
	} else {
		goto L53
	}
L52:
	;
	goto L50
L53:
	;
	goto L52
L54:
	;
	v245 = v241
	v250 = int32(0)
	goto L57
L55:
	;
	goto L56
L56:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v194)+8))
	if v349&int32(2) == int32(0) {
		v364 = v349
		goto L72
	} else {
		goto L73
	}
L57:
	;
	v256 = v250 << (uint(int32(2)) % 32)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v194)+24))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v256+v257)))
	if v259 != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L56
L59:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)+12))
	if v260 != v194 {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v326 = v245
	goto L61
L61:
	;
	v337 = v250 + int32(1)
	if v337 < v326 {
		v245 = v326
		v250 = v337
		goto L57
	} else {
		goto L71
	}
L62:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v194)+24))
	v317 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v315+v256))) = v317
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v194)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v319+v250<<(uint(int32(3))%32)))) = v317
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v326 = v325
	goto L61
L63:
	;
	v270 = int32(*(*int16)(unsafe.Add(mBase, uint32(v259)+16)))
	v271 = v260
	v274 = v270
	goto L67
L64:
	;
	v262 = int32(*(*int16)(unsafe.Add(mBase, uint32(v259)+16)))
	if v250 != v262 {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v194)+28))
	v268 = *(*int64)(unsafe.Add(mBase, uint32(v264+v250<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v259)+12)) = v268
	goto L62
L66:
	;
	v296 = int32(3)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v194)+28))
	v303 = *(*int64)(unsafe.Add(mBase, uint32(v299+v250<<(uint(v296)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v295+v294<<(uint(v296)%32)))) = v303
	goto L62
L67:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v271)+28))
	v284 = v281 + v274<<(uint(int32(3))%32)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)))
	if v285 == int32(0) {
		v294 = v274
		v295 = v281
		goto L66
	} else {
		goto L69
	}
L68:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v271)+28))
	v294 = base.I32_extend16_s(v274)
	v295 = v292
	goto L66
L69:
	;
	v289 = int32(*(*int16)(unsafe.Add(mBase, uint32(v284)+4)))
	if base.B2i32(v285 != v194)|base.B2i32(v250 != v289) != 0 {
		v271 = v285
		v274 = v289
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	goto L58
L72:
	;
	if v364&int32(8) == int32(0) {
		goto L78
	} else {
		goto L79
	}
L73:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v194)+20))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if base.Ui32(v354) <= base.Ui32(v355) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v359 = v355
	goto L76
L75:
	;
	v359 = int32(0)
	goto L76
L76:
	;
	if base.B2i32(v354 == v355)|v359 != 0 {
		v364 = v349
		goto L72
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v354
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v194)+8))
	v364 = v362
	goto L72
L78:
	;
	v390 = v194
	goto L7
L79:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v194)+20))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if base.Ui32(v370) <= base.Ui32(v371) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v375 = v371
	goto L82
L81:
	;
	v375 = int32(0)
	goto L82
L82:
	;
	if base.B2i32(v370 == v371)|v375 != 0 {
		goto L78
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = v370
	goto L78
L84:
	;
	return int32(0)
L85:
	;
	goto L86
L86:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if int32(0) < v395 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v399 = int32(0)
	goto L90
L88:
	;
	goto L89
L89:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v432)+12))
	v438 = v431 + int32(base.Ui32(v433)>>(uint(int32(3))%32))&int32(536870908)
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v438)))
	v440 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v438))) = v439 | v440<<(uint(v433)%32)
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v445 == v440 {
		goto L94
	} else {
		goto L95
	}
L90:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
	*(*int32)(unsafe.Add(mBase, uint32(v410+v399<<(uint(int32(2))%32)))) = int32(0)
	v417 = v399 + int32(1)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v417 < v418 {
		v399 = v417
		goto L90
	} else {
		goto L92
	}
L91:
	;
	goto L89
L92:
	;
	goto L91
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v390)+8)) = int32(13)
	*(*int32)(unsafe.Add(mBase, uint32(v390)+4)) = v524
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v535 <= int32(0) {
		v575 = v390
		goto L1
	} else {
		goto L110
	}
L94:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v444)))
	v524 = v448
	goto L93
L95:
	;
	goto L96
L96:
	;
	if v445 <= int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v524 = int32(0)
	goto L93
L98:
	;
	goto L99
L99:
	;
	v453 = v445 & int32(3)
	v454 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v445) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v460 = v454
	v463 = v454
	v469 = v4
	goto L103
L101:
	;
	v489 = v454
	v492 = v454
	goto L102
L102:
	;
	v500 = v489
	v503 = v492
	v510 = v4
	goto L107
L103:
	;
	v473 = v444 + v460<<(uint(int32(2))%32)
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v473)+12))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v473)+8))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v473)+4))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v473)))
	v481 = v474 ^ (v475 ^ (v476 ^ (v477 ^ v463)))
	v482 = int32(4)
	v483 = v460 + v482
	v485 = v469 + v482
	if v485 != v445&int32(2147483644) {
		v460 = v483
		v463 = v481
		v469 = v485
		goto L103
	} else {
		goto L105
	}
L104:
	;
	if v453 == int32(0) {
		v524 = v481
		goto L93
	} else {
		goto L106
	}
L105:
	;
	goto L104
L106:
	;
	v489 = v483
	v492 = v481
	goto L102
L107:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v444+v500<<(uint(int32(2))%32))))
	v515 = v514 ^ v503
	v516 = int32(1)
	v519 = v510 + v516
	if v519 != v453 {
		v500 = v500 + v516
		v503 = v515
		v510 = v519
		goto L107
	} else {
		goto L109
	}
L108:
	;
	v524 = v515
	goto L93
L109:
	;
	goto L108
L110:
	;
	v542 = v390
	goto L2
L111:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v561+v550<<(uint(int32(5))%32))+20)) = int32(0)
	v568 = v550 + int32(1)
	v569 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v568 < v569 {
		v550 = v568
		goto L111
	} else {
		goto L113
	}
L112:
	;
	v575 = v542
	goto L1
L113:
	;
	goto L112
}
func F_initialize_phase(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	if v4 != 0 {
		F_tuplesort_end(m, v4)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+220)) = int32(0)
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
			if l1 <= int32(1) {
				if v9 != 0 {
					F_tuplesort_end(m, v9)
					mBase = m.M
					v13 = m.ExcPending
					if v13 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+224)) = int32(0)
						if l1 != int32(1) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = l1
							v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v48 + l1*int32(48)
							return
						} else {
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
							if v23-int32(1) <= l1 {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = l1
								v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v48 + l1*int32(48)
								return
							} else {
								v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+56))
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
								v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+l1*int32(48))+72))
								v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+72))
								v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+76))
								v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+80))
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+84))
								v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+88))
								v40 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_phase[0]))
								v41 = int32(0)
								v43 = F_tuplesort_begin_heap(m, v28, v34, v35, v36, v37, v38, v40, v41, v41)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+224)) = v43
									*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = l1
									v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v48 + l1*int32(48)
									return
								}
							}
						}
					}
				} else {
					if l1 != int32(1) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = l1
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v48 + l1*int32(48)
						return
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
						if v23-int32(1) <= l1 {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = l1
							v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v48 + l1*int32(48)
							return
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+56))
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+l1*int32(48))+72))
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+72))
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+76))
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+80))
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+84))
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+88))
							v40 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_phase[0]))
							v41 = int32(0)
							v43 = F_tuplesort_begin_heap(m, v28, v34, v35, v36, v37, v38, v40, v41, v41)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+224)) = v43
								*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = l1
								v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v48 + l1*int32(48)
								return
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+224)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+220)) = v9
				F_tuplesort_performsort(m, v9)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
					if v23-int32(1) <= l1 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = l1
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v48 + l1*int32(48)
						return
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+56))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+l1*int32(48))+72))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+72))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+76))
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+80))
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+84))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+88))
						v40 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_phase[0]))
						v41 = int32(0)
						v43 = F_tuplesort_begin_heap(m, v28, v34, v35, v36, v37, v38, v40, v41, v41)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+224)) = v43
							*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = l1
							v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v48 + l1*int32(48)
							return
						}
					}
				}
			}
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
		if l1 <= int32(1) {
			if v9 != 0 {
				F_tuplesort_end(m, v9)
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+224)) = int32(0)
					if l1 != int32(1) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = l1
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v48 + l1*int32(48)
						return
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
						if v23-int32(1) <= l1 {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = l1
							v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v48 + l1*int32(48)
							return
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+56))
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+l1*int32(48))+72))
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+72))
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+76))
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+80))
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+84))
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+88))
							v40 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_phase[0]))
							v41 = int32(0)
							v43 = F_tuplesort_begin_heap(m, v28, v34, v35, v36, v37, v38, v40, v41, v41)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+224)) = v43
								*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = l1
								v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v48 + l1*int32(48)
								return
							}
						}
					}
				}
			} else {
				if l1 != int32(1) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = l1
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v48 + l1*int32(48)
					return
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
					if v23-int32(1) <= l1 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = l1
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v48 + l1*int32(48)
						return
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+56))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+l1*int32(48))+72))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+72))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+76))
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+80))
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+84))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+88))
						v40 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_phase[0]))
						v41 = int32(0)
						v43 = F_tuplesort_begin_heap(m, v28, v34, v35, v36, v37, v38, v40, v41, v41)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+224)) = v43
							*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = l1
							v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v48 + l1*int32(48)
							return
						}
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+224)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+220)) = v9
			F_tuplesort_performsort(m, v9)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
				if v23-int32(1) <= l1 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = l1
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v48 + l1*int32(48)
					return
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+56))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+l1*int32(48))+72))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+72))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+76))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+80))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+84))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+88))
					v40 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_phase[0]))
					v41 = int32(0)
					v43 = F_tuplesort_begin_heap(m, v28, v34, v35, v36, v37, v38, v40, v41, v41)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+224)) = v43
						*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = l1
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v48 + l1*int32(48)
						return
					}
				}
			}
		}
	}
}
func F_int24le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	return base.B2i32(v3 <= v2)
}
func F_int28le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	return base.B2i32(v4 <= v3)
}
func F_int2in(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v210 int32
	_ = v210
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v273 int32
	_ = v273
	var v283 int32
	_ = v283
	var v291 int32
	_ = v291
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v328 int32
	_ = v328
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	v19 = base.B2i32(v17 == int32(45))
	v20 = v10 + v19
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	v25 = (v21 - int32(48)) & int32(255)
	if base.Ui32(int32(10)) <= base.Ui32(v25) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v14 + int32(32)
	return base.I32_extend16_s(v407)
L2:
	;
	F_errsave_finish(m, v11, int32(_a_F_int2in_0), v401, int32(_a_F_int2in_1))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L81
	} else {
		goto L90
	}
L3:
	;
	v375 = int32(0)
	v376 = F_errsave_start(m, v11)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L81
	} else {
		goto L86
	}
L4:
	;
	v349 = int32(0)
	v350 = F_errsave_start(m, v11)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L81
	} else {
		goto L82
	}
L5:
	;
	v95 = v10
	v100 = v17
	goto L22
L6:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	v30 = v28 - int32(48)
	if base.Ui32(v30&int32(255)) <= base.Ui32(int32(9)) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v37 = v20 + int32(1)
	v39 = v25
	v40 = v30
	goto L10
L8:
	;
	v65 = v28
	v66 = v25
	goto L9
L9:
	;
	if v65&int32(255) != 0 {
		goto L5
	} else {
		goto L14
	}
L10:
	;
	if base.Ui32(int32(3276)) < base.Ui32(v39&int32(_a_F_int2in_2)) {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	v65 = v55
	v66 = v54
	goto L9
L12:
	;
	v50 = int32(10)
	v52 = int32(255)
	v54 = v39*v50 + v40&v52
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+1)))
	v59 = v55 - int32(48)
	if base.Ui32(v59&v52) < base.Ui32(v50) {
		v37 = v37 + int32(1)
		v39 = v54
		v40 = v59
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	if v17 == int32(45) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if base.Ui32(int32(_a_F_int2in_3)) < base.Ui32(v66&int32(_a_F_int2in_2)) {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if base.I32_extend16_s(v66) < int32(0) {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	v407 = int32(0) - v66
	goto L1
L19:
	;
	v407 = v66
	goto L1
L20:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	if v120 != int32(48) {
		goto L28
	} else {
		goto L29
	}
L21:
	;
	v118 = v95 + int32(1)
	v119 = v19
	goto L20
L22:
	;
	if base.Ui32(v100-int32(9)) < base.Ui32(int32(5)) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v113 = int32(1)
	v118 = v95 + v113
	v119 = v113
	goto L20
L24:
	;
	goto L23
L25:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+1)))
	v95 = v95 + int32(1)
	v100 = v110
	goto L22
L26:
	;
	switch v100 - int32(32) {
	case 0:
		goto L25
	default:
		v118 = v95
		v119 = v19
		goto L20
	case 11:
		goto L21
	case 13:
		goto L24
	}
L27:
	;
	if v300 == v301 {
		goto L3
	} else {
		goto L69
	}
L28:
	;
	v262 = v118
	v263 = int32(0)
	v264 = v120
	goto L60
L29:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
	switch v123 - int32(66) {
	case 0, 32:
		goto L30
	default:
		goto L28
	case 13, 45:
		goto L31
	case 22, 54:
		goto L32
	}
L30:
	;
	v221 = v118 + int32(2)
	v223 = v221
	v224 = int32(0)
	goto L52
L31:
	;
	v180 = v118 + int32(2)
	v182 = v180
	v183 = int32(0)
	goto L44
L32:
	;
	v128 = v118 + int32(2)
	v130 = v128
	v131 = int32(0)
	goto L33
L33:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	goto L35
L34:
	;
	goto L3
L35:
	;
	if base.B2i32(base.Ui32(v138-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v138|int32(32)-int32(97)) < base.Ui32(int32(6))) != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if base.Ui32(int32(2048)) < base.Ui32(v131&int32(_a_F_int2in_2)) {
		goto L4
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	if v138 != int32(95) {
		v300 = v128
		v301 = v130
		v302 = v131
		v303 = v138
		goto L27
	} else {
		goto L40
	}
L39:
	;
	v156 = int32(*(*int8)(unsafe.Add(mBase, uint32(v138)+uint32(_c_F_int2in[0]))))
	v130 = v130 + int32(1)
	v131 = v156 + v131<<(uint(int32(4))%32)
	goto L33
L40:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+1)))
	if v162 == int32(0) {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	goto L42
L42:
	;
	if base.B2i32(base.Ui32(v162-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v162|int32(32)-int32(97)) < base.Ui32(int32(6))) != 0 {
		v130 = v130 + int32(1)
		goto L33
	} else {
		goto L43
	}
L43:
	;
	goto L34
L44:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	if v190&int32(248) == int32(48) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L3
L46:
	;
	if base.Ui32(int32(_a_F_int2in_4)) < base.Ui32(v183&int32(_a_F_int2in_2)) {
		goto L4
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if v190 != int32(95) {
		v300 = v180
		v301 = v182
		v302 = v183
		v303 = v190
		goto L27
	} else {
		goto L50
	}
L49:
	;
	v182 = v182 + int32(1)
	v183 = (v190-int32(48))&int32(255) | v183<<(uint(int32(3))%32)
	goto L44
L50:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+1)))
	if base.Ui32(int32(248)) <= base.Ui32((v210-int32(56))&int32(255)) {
		v182 = v182 + int32(1)
		goto L44
	} else {
		goto L51
	}
L51:
	;
	goto L45
L52:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223))))
	if v231&int32(254) == int32(48) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L3
L54:
	;
	if base.Ui32(int32(_a_F_int2in_5)) < base.Ui32(v224&int32(_a_F_int2in_2)) {
		goto L4
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v231 != int32(95) {
		v300 = v221
		v301 = v223
		v302 = v224
		v303 = v231
		goto L27
	} else {
		goto L58
	}
L57:
	;
	v244 = int32(1)
	v223 = v223 + v244
	v224 = (v231-int32(48))&int32(255) | v224<<(uint(v244)%32)
	goto L52
L58:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+1)))
	if base.Ui32(int32(254)) <= base.Ui32((v251-int32(50))&int32(255)) {
		v223 = v223 + int32(1)
		goto L52
	} else {
		goto L59
	}
L59:
	;
	goto L53
L60:
	;
	v273 = (v264 - int32(48)) & int32(255)
	if base.Ui32(v273) <= base.Ui32(int32(9)) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	goto L3
L62:
	;
	if base.Ui32(int32(3276)) < base.Ui32(v263&int32(_a_F_int2in_2)) {
		goto L4
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	if v264&int32(255) != int32(95) {
		v300 = v118
		v301 = v262
		v302 = v263
		v303 = v264
		goto L27
	} else {
		goto L66
	}
L65:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+1)))
	v262 = v262 + int32(1)
	v263 = v263*int32(10) + v273
	v264 = v283
	goto L60
L66:
	;
	if v118 == v262 {
		goto L3
	} else {
		goto L67
	}
L67:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+1)))
	if base.Ui32((v291-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v262 = v262 + int32(1)
		v264 = v291
		goto L60
	} else {
		goto L68
	}
L68:
	;
	goto L61
L69:
	;
	v311 = v301
	v313 = v303
	goto L70
L70:
	;
	v320 = v313 & int32(255)
	if base.B2i32(base.Ui32(v320-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v320 == int32(32)) != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	goto L4
L72:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311)+1)))
	v311 = v311 + int32(1)
	v313 = v328
	goto L70
L73:
	;
	if v320 != 0 {
		goto L3
	} else {
		goto L75
	}
L74:
	;
	goto L71
L75:
	;
	if v119 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	if base.Ui32(int32(_a_F_int2in_3)) < base.Ui32(v302&int32(_a_F_int2in_2)) {
		goto L4
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	if int32(0) <= base.I32_extend16_s(v302) {
		v407 = v302
		goto L1
	} else {
		goto L80
	}
L79:
	;
	v407 = int32(0) - v302
	goto L1
L80:
	;
	goto L74
L81:
	;
	return int32(0)
L82:
	;
	if v350 == int32(0) {
		v407 = v349
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L81
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = int32(_a_F_int2in_6)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v10
	F_errmsg(m, int32(_a_F_int2in_7), v14)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L81
	} else {
		goto L85
	}
L85:
	;
	v394 = v349
	v401 = int32(351)
	goto L2
L86:
	;
	if v376 == int32(0) {
		v407 = v375
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L81
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(_a_F_int2in_6)
	F_errmsg(m, int32(_a_F_int2in_8), v14+int32(16))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L81
	} else {
		goto L89
	}
L89:
	;
	v394 = v375
	v401 = int32(357)
	goto L2
L90:
	;
	v407 = v394
	goto L1
}
func F_int2int4_sum(m *base.Module, l0 int32) int32 {
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
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
		if v8 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_int2int4_sum_0), int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int2int4_sum_1), int32(_a_F_int2int4_sum_2), int32(_a_F_int2int4_sum_3))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
			if v9&int32(-4) != int32(160) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_int2int4_sum_0), int32(0))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_int2int4_sum_1), int32(_a_F_int2int4_sum_2), int32(_a_F_int2int4_sum_3))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
				v21 = v4 + (v14<<(uint(int32(3))%32)+int32(23))&int32(-8)
				v22 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
				if v22 == int64(0) {
					v25 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v25)
					return int32(0)
				} else {
					return v21 + int32(8)
				}
			}
		}
	}
}
func F_int2larger(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v4 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	if v4 < v3 {
		v6 = v3
	} else {
		v6 = v4
	}
	return v6
}
func F_int2ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	v3 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v2 != v3)
}
func F_int2not(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	return base.I32_extend16_s(v2 ^ int32(-1))
}
func F_int2shl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return base.I32_extend16_s(v2 << (uint(v3) % 32))
}
func F_int42gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v3 < v2)
}
func F_int42le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v2 <= v3)
}
func F_int48ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+20)))
	return base.B2i32(v3 != v4)
}
func F_int4xor(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	return v2 ^ v3
}
func F_int82div(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = int32(_a_F_int82div_0)
	v8 = v6 & v7
	if v8 != v7 {
		if v8 != 0 {
			v38 = base.I64_div_s(v5, base.I64_extend16_s(base.I64_extend_i32_u(v6)))
			v39 = F_Int64GetDatum(m, v38)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				return v39
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(33816706))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_int82div_1), int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_int82div_2), int32(1084), int32(_a_F_int82div_3))
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
			}
		}
	} else {
		if v5 == int64(-9223372036854775807-1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_int82div_4), int32(0))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_int82div_2), int32(1100), int32(_a_F_int82div_3))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
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
			v33 = F_Int64GetDatum(m, int64(0)-v5)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				return v33
			}
		}
	}
}
func F_int82ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v3 != v4)
}
func F_int84gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v4 < v3)
}
func F_int84ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v3 != v4)
}
func F_int8dec_any(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_int8dec(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_int8div(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	v9 = int64(1)
	v10 = v8 + v9
	if base.Ui64(v10) <= base.Ui64(v9) {
		if base.I32_wrap_i64(v10) == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(33816706))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_int8div_0), int32(0))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_int8div_1), int32(514), int32(_a_F_int8div_2))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
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
			if v6 == int64(-9223372036854775807-1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_int8div_3), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_int8div_1), int32(530), int32(_a_F_int8div_2))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
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
				v38 = F_Int64GetDatum(m, int64(0)-v6)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					return v38
				}
			}
		}
	} else {
		v41 = base.I64_div_s(v6, v8)
		v42 = F_Int64GetDatum(m, v41)
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return int32(0)
		} else {
			return v42
		}
	}
}
func F_int8range_subdiff(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v9 = F_Float8GetDatum(m, base.F64_sub(base.F64_convert_i64_s(v3), base.F64_convert_i64_s(v6)))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_int8shl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+28)))
	v6 = F_Int64GetDatum(m, v3<<(uint(v4)%64))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_int8shr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+28)))
	v6 = F_Int64GetDatum(m, v3>>(uint(v4)%64))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_intarray_del_elem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v105 int32
	_ = v105
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	v2 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_copy(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
		if v18 != 0 {
			v19 = F_array_contains_nulls(m, v13)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if v19 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v143 = m.ExcPending
					if v143 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67108994))
						mBase = m.M
						v146 = m.ExcPending
						if v146 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_intarray_del_elem_0), int32(0))
							mBase = m.M
							v150 = m.ExcPending
							if v150 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_intarray_del_elem_1), int32(359), int32(_a_F_intarray_del_elem_2))
								mBase = m.M
								v155 = m.ExcPending
								if v155 != 0 {
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
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v23 = v13 + int32(16)
					v24 = F_ArrayGetNItemsSafe(m, v21, v23)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						if v24 != 0 {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
							v27 = F_ArrayGetNItemsSafe(m, v26, v23)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
								if v29 == int32(0) {
									v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
									v39 = (v32<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								} else {
									v39 = v29
								}
								if v27 <= int32(0) {
									v117 = v2
								} else {
									v42 = v39 + v13
									v43 = int32(0)
									if v27 != int32(1) {
										v50 = v43
										v53 = v2
										v60 = v2
										for {
											v63 = v42 + v50<<(uint(int32(2))%32)
											v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
											if v17 == v64 {
												v73 = v53
											} else {
												v67 = v53 + int32(1)
												if v50 <= v53 {
													v73 = v67
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v42+v53<<(uint(int32(2))%32)))) = v64
													v73 = v67
												}
											}
											v74 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
											if v17 == v74 {
												v83 = v73
											} else {
												v77 = v73 + int32(1)
												if v50 < v73 {
													v83 = v77
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v42+v73<<(uint(int32(2))%32)))) = v74
													v83 = v77
												}
											}
											v84 = int32(2)
											v85 = v50 + v84
											v87 = v60 + v84
											if v87 != v27&int32(2147483646) {
												v50 = v85
												v53 = v83
												v60 = v87
												continue
											} else {
												break
											}
											break
										}
										if v27&int32(1) == int32(0) {
											v117 = v83
										} else {
											v91 = v85
											v94 = v83
											v105 = *(*int32)(unsafe.Add(mBase, uint32(v42+v91<<(uint(int32(2))%32))))
											if v105 == v17 {
												v117 = v94
											} else {
												if v94 < v91 {
													*(*int32)(unsafe.Add(mBase, uint32(v42+v94<<(uint(int32(2))%32)))) = v105
												} else {
												}
												v117 = v94 + int32(1)
											}
										}
									} else {
										v91 = v43
										v94 = v2
										v105 = *(*int32)(unsafe.Add(mBase, uint32(v42+v91<<(uint(int32(2))%32))))
										if v105 == v17 {
											v117 = v94
										} else {
											if v94 < v91 {
												*(*int32)(unsafe.Add(mBase, uint32(v42+v94<<(uint(int32(2))%32)))) = v105
											} else {
											}
											v117 = v94 + int32(1)
										}
									}
								}
								v125 = F_resize_intArrayType(m, v13, v117)
								mBase = m.M
								v126 = m.ExcPending
								if v126 != 0 {
									return int32(0)
								} else {
									v138 = v125
									return v138
								}
							}
						} else {
							v138 = v13
							return v138
						}
					}
				}
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			v23 = v13 + int32(16)
			v24 = F_ArrayGetNItemsSafe(m, v21, v23)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				if v24 != 0 {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v27 = F_ArrayGetNItemsSafe(m, v26, v23)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
						if v29 == int32(0) {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
							v39 = (v32<<(uint(int32(3))%32) + int32(23)) & int32(-8)
						} else {
							v39 = v29
						}
						if v27 <= int32(0) {
							v117 = v2
						} else {
							v42 = v39 + v13
							v43 = int32(0)
							if v27 != int32(1) {
								v50 = v43
								v53 = v2
								v60 = v2
								for {
									v63 = v42 + v50<<(uint(int32(2))%32)
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
									if v17 == v64 {
										v73 = v53
									} else {
										v67 = v53 + int32(1)
										if v50 <= v53 {
											v73 = v67
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v42+v53<<(uint(int32(2))%32)))) = v64
											v73 = v67
										}
									}
									v74 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
									if v17 == v74 {
										v83 = v73
									} else {
										v77 = v73 + int32(1)
										if v50 < v73 {
											v83 = v77
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v42+v73<<(uint(int32(2))%32)))) = v74
											v83 = v77
										}
									}
									v84 = int32(2)
									v85 = v50 + v84
									v87 = v60 + v84
									if v87 != v27&int32(2147483646) {
										v50 = v85
										v53 = v83
										v60 = v87
										continue
									} else {
										break
									}
									break
								}
								if v27&int32(1) == int32(0) {
									v117 = v83
								} else {
									v91 = v85
									v94 = v83
									v105 = *(*int32)(unsafe.Add(mBase, uint32(v42+v91<<(uint(int32(2))%32))))
									if v105 == v17 {
										v117 = v94
									} else {
										if v94 < v91 {
											*(*int32)(unsafe.Add(mBase, uint32(v42+v94<<(uint(int32(2))%32)))) = v105
										} else {
										}
										v117 = v94 + int32(1)
									}
								}
							} else {
								v91 = v43
								v94 = v2
								v105 = *(*int32)(unsafe.Add(mBase, uint32(v42+v91<<(uint(int32(2))%32))))
								if v105 == v17 {
									v117 = v94
								} else {
									if v94 < v91 {
										*(*int32)(unsafe.Add(mBase, uint32(v42+v94<<(uint(int32(2))%32)))) = v105
									} else {
									}
									v117 = v94 + int32(1)
								}
							}
						}
						v125 = F_resize_intArrayType(m, v13, v117)
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return int32(0)
						} else {
							v138 = v125
							return v138
						}
					}
				} else {
					v138 = v13
					return v138
				}
			}
		}
	}
}
func F_inter_sl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = F_lseg_interpt_line(m, int32(0), v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_internalerrquery(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_internalerrquery[0]))
	if int32(0) <= v5 {
		v9 = v5 * int32(100)
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_internalerrquery[1])))
		if v12 != 0 {
			F_pfree(m, v12)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_internalerrquery[1]))) = int32(0)
				if l0 != 0 {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_internalerrquery[2])))
					v20 = F_MemoryContextStrdup(m, v19, l0)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_internalerrquery[1]))) = v20
						return int32(0)
					}
				} else {
					return int32(0)
				}
			}
		} else {
			if l0 != 0 {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_internalerrquery[2])))
				v20 = F_MemoryContextStrdup(m, v19, l0)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_internalerrquery[1]))) = v20
					return int32(0)
				}
			} else {
				return int32(0)
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_internalerrquery[0])) = int32(-1)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_internalerrquery_0), int32(0))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_internalerrquery_1), int32(1509), int32(_a_F_internalerrquery_2))
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
func F_interpret_func_volatility(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v10 = int32(_a_F_interpret_func_volatility_0)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_interpret_func_volatility[0])))
	if base.B2i32(v13 == int32(0))|base.B2i32(v13 != v16) != 0 {
		v34 = v13
		v35 = v16
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L27
	} else {
		goto L28
	}
L2:
	;
	m.G0 = v5 + int32(16)
	return v97
L3:
	;
	if v34-v35 == int32(0) {
		v97 = int32(105)
		goto L2
	} else {
		goto L10
	}
L4:
	;
	goto L3
L5:
	;
	v19 = v9
	v20 = v10
	goto L6
L6:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v24 == int32(0) {
		v34 = v24
		v35 = v23
		goto L4
	} else {
		goto L8
	}
L7:
	;
	v34 = v24
	v35 = v23
	goto L4
L8:
	;
	v27 = int32(1)
	if v24 == v23 {
		v19 = v19 + v27
		v20 = v20 + v27
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v40 = int32(_a_F_interpret_func_volatility_1)
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_interpret_func_volatility[1])))
	if base.B2i32(v43 == int32(0))|base.B2i32(v43 != v46) != 0 {
		v64 = v43
		v65 = v46
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v64-v65 == int32(0) {
		v97 = int32(115)
		goto L2
	} else {
		goto L18
	}
L12:
	;
	goto L11
L13:
	;
	v49 = v9
	v50 = v40
	goto L14
L14:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
	if v54 == int32(0) {
		v64 = v54
		v65 = v53
		goto L12
	} else {
		goto L16
	}
L15:
	;
	v64 = v54
	v65 = v53
	goto L12
L16:
	;
	v57 = int32(1)
	if v54 == v53 {
		v49 = v49 + v57
		v50 = v50 + v57
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v69 = int32(_a_F_interpret_func_volatility_2)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_interpret_func_volatility[2])))
	if base.B2i32(v72 == int32(0))|base.B2i32(v72 != v75) != 0 {
		v93 = v72
		v94 = v75
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v93-v94 != 0 {
		goto L1
	} else {
		goto L26
	}
L20:
	;
	goto L19
L21:
	;
	v78 = v9
	v79 = v69
	goto L22
L22:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+1)))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
	if v83 == int32(0) {
		v93 = v83
		v94 = v82
		goto L20
	} else {
		goto L24
	}
L23:
	;
	v93 = v83
	v94 = v82
	goto L20
L24:
	;
	v86 = int32(1)
	if v83 == v82 {
		v78 = v78 + v86
		v79 = v79 + v86
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v97 = int32(118)
	goto L2
L27:
	;
	return int32(0)
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v9
	F_errmsg_internal(m, int32(_a_F_interpret_func_volatility_3), v5)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_interpret_func_volatility_4), int32(629), int32(_a_F_interpret_func_volatility_5))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_intorel_shutdown(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+32)))
	if v5 != 0 {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		F_relation_close(m, v21, int32(0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
			return
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		F_FreeBulkInsertState(m, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+188))
			if v10 == int32(0) {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				F_relation_close(m, v21, int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
					return
				}
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)+108))
				if v13 == int32(0) {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					F_relation_close(m, v21, int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
						return
					}
				} else {
					v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					m.T0[v13].(func(*base.Module, int32, int32))(m, v9, v16)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						F_relation_close(m, v21, int32(0))
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
							return
						}
					}
				}
			}
		}
	}
}
func F_irish_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v157 int32
	_ = v157
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v281 int32
	_ = v281
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v315 int32
	_ = v315
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v406 int32
	_ = v406
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v432 int32
	_ = v432
	var v439 int32
	_ = v439
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v477 int32
	_ = v477
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v528 int32
	_ = v528
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v554 int32
	_ = v554
	var v562 int32
	_ = v562
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v4
	v8 = F_find_among(m, l0, int32(_a_F_irish_UTF_8_stem_0), int32(24))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v714
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v4
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = v79
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v106 = v84
	goto L39
L3:
	;
	return int32(0)
L4:
	;
	if v8 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v14
	switch v8 - int32(1) {
	case 0:
		goto L15
	case 1:
		goto L14
	case 2:
		goto L13
	case 3:
		goto L12
	case 4:
		goto L11
	case 5:
		goto L10
	case 6:
		goto L9
	case 7:
		goto L8
	case 8:
		goto L7
	case 9:
		goto L6
	default:
		goto L2
	}
L6:
	;
	v72 = F_slice_from_s(m, l0, int32(1), int32(_a_F_irish_UTF_8_stem_1))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L3
	} else {
		goto L34
	}
L7:
	;
	v66 = F_slice_from_s(m, l0, int32(1), int32(_a_F_irish_UTF_8_stem_2))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L3
	} else {
		goto L32
	}
L8:
	;
	v60 = F_slice_from_s(m, l0, int32(1), int32(_a_F_irish_UTF_8_stem_3))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L3
	} else {
		goto L30
	}
L9:
	;
	v54 = F_slice_from_s(m, l0, int32(1), int32(_a_F_irish_UTF_8_stem_4))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L3
	} else {
		goto L28
	}
L10:
	;
	v48 = F_slice_from_s(m, l0, int32(1), int32(_a_F_irish_UTF_8_stem_5))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L3
	} else {
		goto L26
	}
L11:
	;
	v42 = F_slice_from_s(m, l0, int32(1), int32(_a_F_irish_UTF_8_stem_6))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L3
	} else {
		goto L24
	}
L12:
	;
	v36 = F_slice_from_s(m, l0, int32(1), int32(_a_F_irish_UTF_8_stem_7))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L22
	}
L13:
	;
	v30 = F_slice_from_s(m, l0, int32(1), int32(_a_F_irish_UTF_8_stem_8))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L3
	} else {
		goto L20
	}
L14:
	;
	v24 = F_slice_from_s(m, l0, int32(1), int32(_a_F_irish_UTF_8_stem_9))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L3
	} else {
		goto L18
	}
L15:
	;
	v18 = F_slice_del(m, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	if int32(0) <= v18 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v714 = v18
	goto L1
L18:
	;
	if int32(0) <= v24 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	v714 = v24
	goto L1
L20:
	;
	if int32(0) <= v30 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v714 = v30
	goto L1
L22:
	;
	if int32(0) <= v36 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v714 = v36
	goto L1
L24:
	;
	if int32(0) <= v42 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v714 = v42
	goto L1
L26:
	;
	if int32(0) <= v48 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	v714 = v48
	goto L1
L28:
	;
	if int32(0) <= v54 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	v714 = v54
	goto L1
L30:
	;
	if int32(0) <= v60 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v714 = v60
	goto L1
L32:
	;
	if int32(0) <= v66 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v714 = v66
	goto L1
L34:
	;
	if v72 < int32(0) {
		v714 = v72
		goto L1
	} else {
		goto L35
	}
L35:
	;
	goto L2
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v84
	v582 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v582
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v582
	v587 = F_find_among_b(m, l0, int32(_a_F_irish_UTF_8_stem_10), int32(16))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L3
	} else {
		goto L140
	}
L37:
	;
	if v201 < int32(0) {
		goto L36
	} else {
		goto L62
	}
L38:
	;
	v201 = v173
	goto L37
L39:
	;
	if v97 <= v106 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v201 = int32(-1)
	goto L37
L42:
	;
	goto L43
L43:
	;
	v113 = int32(1)
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106+v98))))
	if base.Ui32(v115) < base.Ui32(int32(192)) {
		v172 = v115
		v173 = v113
		goto L44
	} else {
		goto L45
	}
L44:
	;
	if int32(250) < v172 {
		goto L57
	} else {
		goto L58
	}
L45:
	;
	v119 = v106 + int32(1)
	if v119 == v97 {
		v172 = v115
		v173 = v113
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119+v98))))
	v124 = v122 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v115) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128+v98))))
	v140 = v138 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v115) {
		goto L53
	} else {
		goto L54
	}
L48:
	;
	v128 = v106 + int32(2)
	if v128 != v97 {
		goto L47
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v172 = v115<<(uint(int32(6))%32)&int32(1984) | v124
	v173 = int32(2)
	goto L44
L51:
	;
	goto L50
L52:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98+v144))))
	v172 = v157&int32(63) | (v115<<(uint(int32(18))%32)&int32(_a_F_irish_UTF_8_stem_11) | v124<<(uint(int32(12))%32) | v140<<(uint(int32(6))%32))
	v173 = int32(4)
	goto L44
L53:
	;
	v144 = v106 + int32(3)
	if v144 != v97 {
		goto L52
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v172 = v115<<(uint(int32(12))%32)&int32(_a_F_irish_UTF_8_stem_12) | v124<<(uint(int32(6))%32) | v140
	v173 = int32(3)
	goto L44
L56:
	;
	goto L55
L57:
	;
	v190 = v173 + v106
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v190
	v106 = v190
	goto L39
L58:
	;
	v177 = v172 - int32(97)
	if v177 < int32(0) {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v177)>>(uint(int32(3))%32)))+uint32(_c_F_irish_UTF_8_stem[0]))))
	if int32(base.Ui32(v183)>>(uint(v177&int32(7))%32))&int32(1) != 0 {
		goto L38
	} else {
		goto L60
	}
L60:
	;
	goto L57
L62:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v205 = v204 + v201
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v205
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v207)+8)) = v205
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v230 = v220
	goto L65
L63:
	;
	if v326 < int32(0) {
		goto L36
	} else {
		goto L87
	}
L64:
	;
	v326 = v297
	goto L63
L65:
	;
	if v221 <= v230 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v326 = int32(-1)
	goto L63
L68:
	;
	goto L69
L69:
	;
	v237 = int32(1)
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230+v222))))
	if base.Ui32(v239) < base.Ui32(int32(192)) {
		v296 = v239
		v297 = v237
		goto L70
	} else {
		goto L71
	}
L70:
	;
	if int32(250) < v296 {
		goto L64
	} else {
		goto L83
	}
L71:
	;
	v243 = v230 + int32(1)
	if v243 == v221 {
		v296 = v239
		v297 = v237
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243+v222))))
	v248 = v246 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v239) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252+v222))))
	v264 = v262 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v239) {
		goto L79
	} else {
		goto L80
	}
L74:
	;
	v252 = v230 + int32(2)
	if v252 != v221 {
		goto L73
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v296 = v239<<(uint(int32(6))%32)&int32(1984) | v248
	v297 = int32(2)
	goto L70
L77:
	;
	goto L76
L78:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222+v268))))
	v296 = v281&int32(63) | (v239<<(uint(int32(18))%32)&int32(_a_F_irish_UTF_8_stem_11) | v248<<(uint(int32(12))%32) | v264<<(uint(int32(6))%32))
	v297 = int32(4)
	goto L70
L79:
	;
	v268 = v230 + int32(3)
	if v268 != v221 {
		goto L78
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v296 = v239<<(uint(int32(12))%32)&int32(_a_F_irish_UTF_8_stem_12) | v248<<(uint(int32(6))%32) | v264
	v297 = int32(3)
	goto L70
L82:
	;
	goto L81
L83:
	;
	v301 = v296 - int32(97)
	if v301 < int32(0) {
		goto L64
	} else {
		goto L84
	}
L84:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v301)>>(uint(int32(3))%32)))+uint32(_c_F_irish_UTF_8_stem[0]))))
	if int32(base.Ui32(v307)>>(uint(v301&int32(7))%32))&int32(1) == int32(0) {
		goto L64
	} else {
		goto L85
	}
L85:
	;
	v315 = v297 + v230
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v315
	v230 = v315
	goto L65
L87:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v330 = v329 + v326
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v330
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v332)+4)) = v330
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v355 = v345
	goto L90
L88:
	;
	if v450 < int32(0) {
		goto L36
	} else {
		goto L113
	}
L89:
	;
	v450 = v422
	goto L88
L90:
	;
	if v346 <= v355 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v450 = int32(-1)
	goto L88
L93:
	;
	goto L94
L94:
	;
	v362 = int32(1)
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v355+v347))))
	if base.Ui32(v364) < base.Ui32(int32(192)) {
		v421 = v364
		v422 = v362
		goto L95
	} else {
		goto L96
	}
L95:
	;
	if int32(250) < v421 {
		goto L108
	} else {
		goto L109
	}
L96:
	;
	v368 = v355 + int32(1)
	if v368 == v346 {
		v421 = v364
		v422 = v362
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368+v347))))
	v373 = v371 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v364) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377+v347))))
	v389 = v387 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v364) {
		goto L104
	} else {
		goto L105
	}
L99:
	;
	v377 = v355 + int32(2)
	if v377 != v346 {
		goto L98
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v421 = v364<<(uint(int32(6))%32)&int32(1984) | v373
	v422 = int32(2)
	goto L95
L102:
	;
	goto L101
L103:
	;
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347+v393))))
	v421 = v406&int32(63) | (v364<<(uint(int32(18))%32)&int32(_a_F_irish_UTF_8_stem_11) | v373<<(uint(int32(12))%32) | v389<<(uint(int32(6))%32))
	v422 = int32(4)
	goto L95
L104:
	;
	v393 = v355 + int32(3)
	if v393 != v346 {
		goto L103
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v421 = v364<<(uint(int32(12))%32)&int32(_a_F_irish_UTF_8_stem_12) | v373<<(uint(int32(6))%32) | v389
	v422 = int32(3)
	goto L95
L107:
	;
	goto L106
L108:
	;
	v439 = v422 + v355
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v439
	v355 = v439
	goto L90
L109:
	;
	v426 = v421 - int32(97)
	if v426 < int32(0) {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v426)>>(uint(int32(3))%32)))+uint32(_c_F_irish_UTF_8_stem[0]))))
	if int32(base.Ui32(v432)>>(uint(v426&int32(7))%32))&int32(1) != 0 {
		goto L89
	} else {
		goto L111
	}
L111:
	;
	goto L108
L113:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v454 = v453 + v450
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v454
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v477 = v454
	goto L116
L114:
	;
	if v573 < int32(0) {
		goto L36
	} else {
		goto L138
	}
L115:
	;
	v573 = v544
	goto L114
L116:
	;
	if v468 <= v477 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v573 = int32(-1)
	goto L114
L119:
	;
	goto L120
L120:
	;
	v484 = int32(1)
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477+v469))))
	if base.Ui32(v486) < base.Ui32(int32(192)) {
		v543 = v486
		v544 = v484
		goto L121
	} else {
		goto L122
	}
L121:
	;
	if int32(250) < v543 {
		goto L115
	} else {
		goto L134
	}
L122:
	;
	v490 = v477 + int32(1)
	if v490 == v468 {
		v543 = v486
		v544 = v484
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v490+v469))))
	v495 = v493 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v486) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499+v469))))
	v511 = v509 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v486) {
		goto L130
	} else {
		goto L131
	}
L125:
	;
	v499 = v477 + int32(2)
	if v499 != v468 {
		goto L124
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v543 = v486<<(uint(int32(6))%32)&int32(1984) | v495
	v544 = int32(2)
	goto L121
L128:
	;
	goto L127
L129:
	;
	v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469+v515))))
	v543 = v528&int32(63) | (v486<<(uint(int32(18))%32)&int32(_a_F_irish_UTF_8_stem_11) | v495<<(uint(int32(12))%32) | v511<<(uint(int32(6))%32))
	v544 = int32(4)
	goto L121
L130:
	;
	v515 = v477 + int32(3)
	if v515 != v468 {
		goto L129
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v543 = v486<<(uint(int32(12))%32)&int32(_a_F_irish_UTF_8_stem_12) | v495<<(uint(int32(6))%32) | v511
	v544 = int32(3)
	goto L121
L133:
	;
	goto L132
L134:
	;
	v548 = v543 - int32(97)
	if v548 < int32(0) {
		goto L115
	} else {
		goto L135
	}
L135:
	;
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v548)>>(uint(int32(3))%32)))+uint32(_c_F_irish_UTF_8_stem[0]))))
	if int32(base.Ui32(v554)>>(uint(v548&int32(7))%32))&int32(1) == int32(0) {
		goto L115
	} else {
		goto L136
	}
L136:
	;
	v562 = v544 + v477
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v562
	v477 = v562
	goto L116
L138:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v577 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v576))) = v577 + v573
	goto L36
L139:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v610
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v610
	v615 = F_find_among_b(m, l0, int32(_a_F_irish_UTF_8_stem_13), int32(25))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L3
	} else {
		goto L151
	}
L140:
	;
	if v587 == int32(0) {
		goto L139
	} else {
		goto L141
	}
L141:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v591
	switch v587 - int32(1) {
	case 0:
		goto L143
	case 1:
		goto L142
	default:
		goto L139
	}
L142:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v602)))
	if v591 < v603 {
		goto L139
	} else {
		goto L147
	}
L143:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v595)+4))
	if v591 < v596 {
		goto L139
	} else {
		goto L144
	}
L144:
	;
	v598 = F_slice_del(m, l0)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L3
	} else {
		goto L145
	}
L145:
	;
	if int32(0) <= v598 {
		goto L139
	} else {
		goto L146
	}
L146:
	;
	v714 = v598
	goto L1
L147:
	;
	v605 = F_slice_del(m, l0)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L3
	} else {
		goto L148
	}
L148:
	;
	if v605 < int32(0) {
		v714 = v605
		goto L1
	} else {
		goto L149
	}
L149:
	;
	goto L139
L150:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v662
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v662
	v665 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v662-int32(2) <= v665 {
		goto L172
	} else {
		goto L173
	}
L151:
	;
	if v615 == int32(0) {
		goto L150
	} else {
		goto L152
	}
L152:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v619
	switch v615 - int32(1) {
	case 0:
		goto L158
	case 1:
		goto L157
	case 2:
		goto L156
	case 3:
		goto L155
	case 4:
		goto L154
	case 5:
		goto L153
	default:
		goto L150
	}
L153:
	;
	v656 = F_slice_from_s(m, l0, int32(4), int32(_a_F_irish_UTF_8_stem_14))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L3
	} else {
		goto L170
	}
L154:
	;
	v650 = F_slice_from_s(m, l0, int32(5), int32(_a_F_irish_UTF_8_stem_15))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L3
	} else {
		goto L168
	}
L155:
	;
	v644 = F_slice_from_s(m, l0, int32(4), int32(_a_F_irish_UTF_8_stem_16))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L3
	} else {
		goto L166
	}
L156:
	;
	v638 = F_slice_from_s(m, l0, int32(3), int32(_a_F_irish_UTF_8_stem_17))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L3
	} else {
		goto L164
	}
L157:
	;
	v632 = F_slice_from_s(m, l0, int32(3), int32(_a_F_irish_UTF_8_stem_18))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L3
	} else {
		goto L162
	}
L158:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v623)))
	if v619 < v624 {
		goto L150
	} else {
		goto L159
	}
L159:
	;
	v626 = F_slice_del(m, l0)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L3
	} else {
		goto L160
	}
L160:
	;
	if int32(0) <= v626 {
		goto L150
	} else {
		goto L161
	}
L161:
	;
	v714 = v626
	goto L1
L162:
	;
	if int32(0) <= v632 {
		goto L150
	} else {
		goto L163
	}
L163:
	;
	v714 = v632
	goto L1
L164:
	;
	if int32(0) <= v638 {
		goto L150
	} else {
		goto L165
	}
L165:
	;
	v714 = v638
	goto L1
L166:
	;
	if int32(0) <= v644 {
		goto L150
	} else {
		goto L167
	}
L167:
	;
	v714 = v644
	goto L1
L168:
	;
	if int32(0) <= v650 {
		goto L150
	} else {
		goto L169
	}
L169:
	;
	v714 = v650
	goto L1
L170:
	;
	if v656 < int32(0) {
		v714 = v656
		goto L1
	} else {
		goto L171
	}
L171:
	;
	goto L150
L172:
	;
	v711 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v711
	v714 = int32(1)
	goto L1
L173:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v671 = int32(1)
	v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v669+v662-v671))))
	if base.B2i32(v673&int32(224) != int32(96))|base.B2i32(v671<<(uint(v673)%32)&int32(_a_F_irish_UTF_8_stem_19) == int32(0)) != 0 {
		goto L172
	} else {
		goto L174
	}
L174:
	;
	v687 = F_find_among_b(m, l0, int32(_a_F_irish_UTF_8_stem_20), int32(12))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L3
	} else {
		goto L175
	}
L175:
	;
	if v687 == int32(0) {
		goto L172
	} else {
		goto L176
	}
L176:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v691
	switch v687 - int32(1) {
	case 0:
		goto L178
	case 1:
		goto L177
	default:
		goto L172
	}
L177:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v702)+4))
	if v691 < v703 {
		goto L172
	} else {
		goto L182
	}
L178:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v695)+8))
	if v691 < v696 {
		goto L172
	} else {
		goto L179
	}
L179:
	;
	v698 = F_slice_del(m, l0)
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L3
	} else {
		goto L180
	}
L180:
	;
	if int32(0) <= v698 {
		goto L172
	} else {
		goto L181
	}
L181:
	;
	v714 = v698
	goto L1
L182:
	;
	v705 = F_slice_del(m, l0)
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L3
	} else {
		goto L183
	}
L183:
	;
	if v705 < int32(0) {
		v714 = v705
		goto L1
	} else {
		goto L184
	}
L184:
	;
	goto L172
}
func F_is_pseudo_constant_clause(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	v3 = F_contain_var_clause(m, l0)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v11 = F_contain_volatile_functions_walker(m, l0, int32(0))
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				if v11 == int32(0) {
					v17 = int32(1)
				} else {
					v17 = int32(0)
				}
				return v17
			}
		} else {
			v17 = int32(0)
			return v17
		}
	}
}
func F_is_strict_saop(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	F_set_opfuncid(m, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v8 = F_func_strict(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			if v8 == int32(0) {
				v47 = int32(0)
				return v47
			} else {
				v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
				if v13 != 0 {
					v47 = int32(1)
					return v47
				} else {
					v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
					v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
					if v16 == int32(0) {
						v47 = int32(0)
						return v47
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
						if v19 != int32(35) {
							if v19 != int32(7) {
								v47 = int32(0)
								return v47
							} else {
								v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+24)))
								if v24 != 0 {
									v47 = int32(0)
									return v47
								} else {
									v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
									v27 = F_pg_detoast_datum(m, v26)
									mBase = m.M
									v28 = m.ExcPending
									if v28 != 0 {
										return int32(0)
									} else {
										v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
										v32 = F_ArrayGetNItemsSafe(m, v29, v27+int32(16))
										mBase = m.M
										v33 = m.ExcPending
										if v33 != 0 {
											return int32(0)
										} else {
											if v32 <= int32(0) {
												v47 = int32(0)
											} else {
												v47 = int32(1)
											}
											return v47
										}
									}
								}
							}
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
							if v36 == int32(0) {
								v47 = int32(0)
							} else {
								v39 = int32(1)
								v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+20)))
								if v40 != v39 {
									v47 = v39
								} else {
									v47 = int32(0)
								}
							}
							return v47
						}
					}
				}
			}
		}
	}
}
func F_isbn_cast_from_ean13(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13948(m, l0, int32(3))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_isbn_in(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13891(m, l0, int32(3))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_isgraph(m *base.Module, l0 int32) int32 {
	return base.B2i32(base.Ui32(l0-int32(33)) < base.Ui32(int32(94)))
}
func F_ispunct(m *base.Module, l0 int32) int32 {
	var v18 int32
	_ = v18
	if base.Ui32(l0-int32(33)) <= base.Ui32(int32(93)) {
		v18 = base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(l0|int32(32)-int32(97)) < base.Ui32(int32(26)))
	} else {
		v18 = int32(1)
	}
	return base.B2i32(v18 == int32(0))
}
func F_iterate_values_scalar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	switch l2 - int32(1) {
	case 0:
		v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
		if v6&int32(2) != 0 {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v18 = F_strlen(m, l1)
			mBase = m.M
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			m.T0[v19].(func(*base.Module, int32, int32, int32))(m, v17, l1, v18)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		} else {
			return int32(0)
		}
	case 1:
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
		if v9&int32(4) != 0 {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v18 = F_strlen(m, l1)
			mBase = m.M
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			m.T0[v19].(func(*base.Module, int32, int32, int32))(m, v17, l1, v18)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		} else {
			return int32(0)
		}
	default:
		return int32(0)
	case 8, 9:
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
		if v12&int32(8) == int32(0) {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v18 = F_strlen(m, l1)
			mBase = m.M
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			m.T0[v19].(func(*base.Module, int32, int32, int32))(m, v17, l1, v18)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		}
	}
}
func F_ivfflatbeginscan(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_ivfflatbeginscan[0]))
	v16 = F_RelationGetIndexScan(m, l0, l1, l2)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		F_IvfflatGetMetaPageInfo(m, l0, v12+int32(12), v12+int32(8))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, _c_F_ivfflatbeginscan[1]))
			v29 = *(*int32)(unsafe.Add(mBase, _c_F_ivfflatbeginscan[2]))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
			v32 = F_palloc(m, int32(80))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v34 = F_IvfflatGetTypeInfo(m, l0)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					v36 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v32)+16)) = uint8(v36)
					*(*int32)(unsafe.Add(mBase, uint32(v32))) = v34
					if v15 < v30 {
						v40 = v15
					} else {
						v40 = v30
					}
					*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v40
					if v15 < v29 {
						v43 = v29
					} else {
						v43 = v15
					}
					if v27 != 0 {
						v44 = v43
					} else {
						v44 = v15
					}
					if v44 < v30 {
						v46 = v44
					} else {
						v46 = v30
					}
					*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v46
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v48
					v52 = int32(1)
					v54 = F_index_getprocinfo(m, l0, v52, v52)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v54
						v58 = F_HnswOptionalProcInfo(m, l0, int32(2))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v32)+52)) = v58
							v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
							*(*int32)(unsafe.Add(mBase, uint32(v32)+56)) = v62
							v65 = *(*int32)(unsafe.Add(mBase, _c_F_ivfflatbeginscan[3]))
							v70 = F_AllocSetContextCreateInternal(m, v65, int32(_a_F_ivfflatbeginscan_0), int32(0), int32(_a_F_ivfflatbeginscan_1), int32(_a_F_ivfflatbeginscan_2))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v70
								v73 = int32(_a_F_ivfflatbeginscan_3)
								v74 = *(*int32)(unsafe.Add(mBase, _c_F_ivfflatbeginscan[3]))
								*(*int32)(unsafe.Add(mBase, _c_F_ivfflatbeginscan[3])) = v70
								v78 = F_CreateTemplateTupleDesc(m, int32(2))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = v78
									F_TupleDescInitEntry(m, v78, int32(1), int32(_a_F_ivfflatbeginscan_4), int32(701), int32(-1), int32(0))
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return int32(0)
									} else {
										v88 = *(*int32)(unsafe.Add(mBase, uint32(v32)+32))
										F_TupleDescInitEntry(m, v88, int32(2), int32(_a_F_ivfflatbeginscan_5), int32(27), int32(-1), int32(0))
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return int32(0)
										} else {
											v96 = *(*int32)(unsafe.Add(mBase, uint32(v32)+32))
											v97 = int32(1)
											*(*uint16)(unsafe.Add(mBase, uint32(v12)+30)) = uint16(v97)
											*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(672)
											v101 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v101
											*(*uint8)(unsafe.Add(mBase, uint32(v12)+19)) = uint8(v101)
											v115 = *(*int32)(unsafe.Add(mBase, _c_F_ivfflatbeginscan[4]))
											v118 = F_tuplesort_begin_heap(m, v96, v97, v12+int32(30), v12+int32(24), v12+int32(20), v12+int32(19), v115, v101, v101)
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v118
												v121 = *(*int32)(unsafe.Add(mBase, uint32(v32)+32))
												v123 = F_MakeTupleTableSlot(m, v121, int32(_a_F_ivfflatbeginscan_6))
												mBase = m.M
												v124 = m.ExcPending
												if v124 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v32)+36)) = v123
													v126 = *(*int32)(unsafe.Add(mBase, uint32(v32)+32))
													v128 = F_MakeTupleTableSlot(m, v126, int32(_a_F_ivfflatbeginscan_7))
													mBase = m.M
													v129 = m.ExcPending
													if v129 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v32)+40)) = v128
														v132 = F_GetAccessStrategy(m, int32(1))
														mBase = m.M
														v133 = m.ExcPending
														if v133 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v32)+44)) = v132
															v136 = F_pairingheap_allocate(m, int32(_a_F_ivfflatbeginscan_8), v16)
															mBase = m.M
															v137 = m.ExcPending
															if v137 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = v136
																v140 = F_mul_size(m, int32(4), v46)
																mBase = m.M
																v141 = m.ExcPending
																if v141 != 0 {
																	return int32(0)
																} else {
																	v142 = F_palloc(m, v140)
																	mBase = m.M
																	v143 = m.ExcPending
																	if v143 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v32)+72)) = int32(0)
																		*(*int32)(unsafe.Add(mBase, uint32(v32)+68)) = v142
																		v148 = F_mul_size(m, int32(24), v46)
																		mBase = m.M
																		v149 = m.ExcPending
																		if v149 != 0 {
																			return int32(0)
																		} else {
																			v150 = F_palloc(m, v148)
																			mBase = m.M
																			v151 = m.ExcPending
																			if v151 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v32)+76)) = v150
																				*(*int32)(unsafe.Add(mBase, _c_F_ivfflatbeginscan[3])) = v74
																				*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v32
																				m.G0 = v12 + int32(32)
																				return v16
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
				}
			}
		}
	}
}
func F_ivfflatbuild(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 float64
	_ = v17
	var v19 float64
	_ = v19
	v5 = m.G0
	v7 = v5 - int32(176)
	m.G0 = v7
	F_BuildIndex_2(m, l0, l1, l2, v7, int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v15 = F_palloc(m, int32(16))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*float64)(unsafe.Add(mBase, uint32(v7)+40))
			*(*float64)(unsafe.Add(mBase, uint32(v15))) = v17
			v19 = *(*float64)(unsafe.Add(mBase, uint32(v7)+32))
			*(*float64)(unsafe.Add(mBase, uint32(v15)+8)) = v19
			m.G0 = v7 + int32(176)
			return v15
		}
	}
}
func F_ivfflathandler(m *base.Module, l0 int32) int32 {
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	v16 = Fn13946(m, l0, int32(_a_F_ivfflathandler_0), int32(_a_F_ivfflathandler_1), int32(_a_F_ivfflathandler_2), int32(_a_F_ivfflathandler_3), int32(_a_F_ivfflathandler_4), int32(_a_F_ivfflathandler_5), int32(_a_F_ivfflathandler_6), int32(_a_F_ivfflathandler_7), int32(_a_F_ivfflathandler_8), int32(_a_F_ivfflathandler_9), int32(_a_F_ivfflathandler_10), int32(_a_F_ivfflathandler_11), int32(_a_F_ivfflathandler_12), int64(72057594038255616))
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		return v16
	}
}
func F_ivfflatoptions(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_ivfflatoptions[0]))
	v8 = F_build_reloptions(m, l0, l1, v4, int32(8), int32(_a_F_ivfflatoptions_0), int32(1))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_ivfflatrescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+16)) = uint8(v7)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+64))
	v10 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(v6)+72)) = v10
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+52))
	if v14 == v10 {
		if l1 == int32(0) {
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v27 <= int32(0) {
			} else {
				v31 = v27 * int32(48)
				if v31 == int32(0) {
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					base.MemoryCopy(m, v34, l1, v31)
				}
			}
		}
		if l3 == int32(0) {
		} else {
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if v39 <= int32(0) {
			} else {
				v43 = v39 * int32(48)
				if v43 == int32(0) {
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					base.MemoryCopy(m, v46, l3, v43)
				}
			}
		}
		return
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
		if v17 == int32(0) {
			if l1 == int32(0) {
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v27 <= int32(0) {
				} else {
					v31 = v27 * int32(48)
					if v31 == int32(0) {
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						base.MemoryCopy(m, v34, l1, v31)
					}
				}
			}
			if l3 == int32(0) {
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v39 <= int32(0) {
				} else {
					v43 = v39 * int32(48)
					if v43 == int32(0) {
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						base.MemoryCopy(m, v46, l3, v43)
					}
				}
			}
			return
		} else {
			F_pfree(m, v17)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(0)
				if l1 == int32(0) {
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v27 <= int32(0) {
					} else {
						v31 = v27 * int32(48)
						if v31 == int32(0) {
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							base.MemoryCopy(m, v34, l1, v31)
						}
					}
				}
				if l3 == int32(0) {
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					if v39 <= int32(0) {
					} else {
						v43 = v39 * int32(48)
						if v43 == int32(0) {
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							base.MemoryCopy(m, v46, l3, v43)
						}
					}
				}
				return
			}
		}
	}
}
