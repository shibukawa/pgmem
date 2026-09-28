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
	v15 = F_query_or_expression_tree_walker_impl(m, l0, int32(1128), v7+int32(8), v6)
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
func F_IpcMemoryDelete(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	v3 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = base.I32_wrap_i64(l1)
	v11 = F_pgmem_shmctl(m, v8, v3, v3)
	mBase = m.M
	if v3 <= v11 {
		m.G0 = v6 + int32(16)
		return
	} else {
		v16 = F_errstart(m, int32(15), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			if v16 == int32(0) {
				m.G0 = v6 + int32(16)
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
				F_errmsg_internal(m, int32(_a_F_IpcMemoryDelete_0), v6)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_IpcMemoryDelete_1), int32(303), int32(_a_F_IpcMemoryDelete_2))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						m.G0 = v6 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_IpcMemoryDetach(m *base.Module, l0 int32, l1 int64) {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = base.I32_wrap_i64(l1)
	v9 = F_pgmem_shmdt(m, v8)
	mBase = m.M
	if int32(0) <= v9 {
		m.G0 = v6 + int32(16)
		return
	} else {
		v14 = F_errstart(m, int32(15), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			if v14 == int32(0) {
				m.G0 = v6 + int32(16)
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
				F_errmsg_internal(m, int32(_a_F_IpcMemoryDetach_0), v6)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_IpcMemoryDetach_1), int32(291), int32(_a_F_IpcMemoryDetach_2))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						m.G0 = v6 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_i2tof(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	return base.I64_extend_i32_s(base.I32_reinterpret_f32(base.F32_convert_i32_s(v2)))
}
func F_i4tochar(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v33 int64
	_ = v33
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui32(base.I32_wrap_i64(v3)-int32(128)) <= base.Ui32(int32(-257)) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v11 = F_errsave_start(m, v10)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			if v11 == int32(0) {
				v33 = int64(0)
				return v33
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_i4tochar_0), int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int64(0)
					} else {
						F_errsave_finish(m, v10, int32(_a_F_i4tochar_1), int32(197), int32(_a_F_i4tochar_2))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int64(0)
						} else {
							return int64(0)
						}
					}
				}
			}
		}
	} else {
		v33 = base.I64_extend8_s(v3)
		return v33
	}
}
func F_i8tod(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	return base.I64_reinterpret_f64(base.F64_convert_i64_s(v2))
}
func F_i8tooid(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(v3) < base.Ui64(int64(4294967296)) {
		v27 = v3
		return v27
	} else {
		v6 = int64(0)
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v8 = F_errsave_start(m, v7)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			if v8 == int32(0) {
				v27 = v6
				return v27
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_i8tooid_0), int32(0))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int64(0)
					} else {
						F_errsave_finish(m, v7, int32(_a_F_i8tooid_1), int32(1402), int32(_a_F_i8tooid_2))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int64(0)
						} else {
							v27 = v6
							return v27
						}
					}
				}
			}
		}
	}
}
func F_icount(m *base.Module, l0 int32) int64 {
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pg_detoast_datum(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
		v12 = F_ArrayGetNItemsSafe(m, v9, v5+int32(16))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v14 != v5 {
				F_pfree(m, v5)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int64(0)
				} else {
					return base.I64_extend_i32_s(v12)
				}
			} else {
				return base.I64_extend_i32_s(v12)
			}
		}
	}
}
func F_icregexnesel(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14371(m, l0, int32(3))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int64
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v119 int64
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = l1 - int32(1)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0+v15<<(uint(int32(2))%32))+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+116))
	if l2 != v21 {
		v23 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v20)+936)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v20)+908)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v20)+880)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v20)+852)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v20)+824)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v20)+796)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v20)+768)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v20)+740)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v20)+712)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v20)+684)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v20)+656)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v20)+628)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v20)+600)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v20)+572)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v20)+544)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v20)+516)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v20)+488)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v20)+460)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v20)+432)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v20)+404)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v20)+376)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v20)+348)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v20)+320)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v20)+292)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v20)+264)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v20)+208)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v20)+180)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v20)+152)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v20)+124)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v20)+116)) = l2
	} else {
	}
	v86 = v20 + l3*int32(28)
	v88 = v86 + int32(92)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v86)+96))
	if v89 == int32(0) {
		v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+208))
		v98 = *(*int32)(unsafe.Add(mBase, uint32(v94+v15<<(uint(int32(2))%32))))
		v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
		v107 = v100 + v101<<(uint(int32(3))%32) + v15*int32(100)
		v108 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v107)+96)))
		v111 = F_SearchSysCache4(m, int32(4), base.I64_extend_i32_u(v98), v108, base.I64_extend_i32_u(l2), base.I64_extend_i32_u(l3))
		mBase = m.M
		v114 = m.ExcPending
		if v114 != 0 {
			return int32(0)
		} else {
			if v111 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v139 = m.ExcPending
				if v139 != 0 {
					return int32(0)
				} else {
					v140 = *(*int32)(unsafe.Add(mBase, uint32(v107)+96))
					*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v98
					*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v140
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = l3
					F_errmsg_internal(m, int32(_a_F_inclusion_get_strategy_procinfo_0), v12)
					mBase = m.M
					v147 = m.ExcPending
					if v147 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_inclusion_get_strategy_procinfo_1), int32(648), int32(_a_F_inclusion_get_strategy_procinfo_2))
						mBase = m.M
						v152 = m.ExcPending
						if v152 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v119 = F_SysCacheGetAttrNotNull(m, int32(4), v111, int32(7))
				mBase = m.M
				v120 = m.ExcPending
				if v120 != 0 {
					return int32(0)
				} else {
					F_ReleaseCatCache(m, v111)
					mBase = m.M
					v122 = m.ExcPending
					if v122 != 0 {
						return int32(0)
					} else {
						v124 = F_get_opcode(m, base.I32_wrap_i64(v119))
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return int32(0)
						} else {
							v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							F_fmgr_info_cxt(m, v124, v88, v126)
							mBase = m.M
							v128 = m.ExcPending
							if v128 != 0 {
								return int32(0)
							} else {
								m.G0 = v12 + int32(16)
								return v88
							}
						}
					}
				}
			}
		}
	} else {
		m.G0 = v12 + int32(16)
		return v88
	}
}
func F_ineq_histogram_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int64, l8 int32) float64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v29 float64
	_ = v29
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 float64
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v78 int64
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v99 int32
	_ = v99
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int64
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
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
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v237 int32
	_ = v237
	var v263 int32
	_ = v263
	var v274 int32
	_ = v274
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v318 int32
	_ = v318
	var v327 int32
	_ = v327
	var v353 int32
	_ = v353
	var v357 int64
	_ = v357
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int64
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v380 float64
	_ = v380
	var v385 float64
	_ = v385
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v418 float64
	_ = v418
	var v420 float64
	_ = v420
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v437 int32
	_ = v437
	var v451 int32
	_ = v451
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v502 int64
	_ = v502
	var v503 int64
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 float64
	_ = v526
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 float32
	_ = v534
	var v536 float32
	_ = v536
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v563 float64
	_ = v563
	var v564 float64
	_ = v564
	var v568 int32
	_ = v568
	var v569 float64
	_ = v569
	var v572 float64
	_ = v572
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v579 float64
	_ = v579
	var v582 int32
	_ = v582
	var v589 float64
	_ = v589
	var v592 float64
	_ = v592
	var v593 int32
	_ = v593
	var v598 float64
	_ = v598
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v612 float64
	_ = v612
	var v613 float64
	_ = v613
	var v618 float64
	_ = v618
	var v621 float64
	_ = v621
	var v622 float64
	_ = v622
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int64
	_ = v629
	var v633 int64
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v654 int32
	_ = v654
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v696 int32
	_ = v696
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v726 int32
	_ = v726
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v756 int32
	_ = v756
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v816 int32
	_ = v816
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v925 float64
	_ = v925
	var v928 float64
	_ = v928
	var v932 int32
	_ = v932
	var v935 float64
	_ = v935
	var v936 int32
	_ = v936
	var v976 float64
	_ = v976
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v993 int32
	_ = v993
	var v1017 float64
	_ = v1017
	var v1018 float64
	_ = v1018
	var v1024 int32
	_ = v1024
	var v1027 float64
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1066 float64
	_ = v1066
	var v1077 int32
	_ = v1077
	var v1080 int32
	_ = v1080
	var v1084 int32
	_ = v1084
	var v1088 int32
	_ = v1088
	var v1112 float64
	_ = v1112
	var v1113 float64
	_ = v1113
	var v1119 int32
	_ = v1119
	var v1122 float64
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1218 int32
	_ = v1218
	var v1219 float64
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1222 float64
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1225 float64
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1246 int32
	_ = v1246
	var v1248 int32
	_ = v1248
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1282 int32
	_ = v1282
	var v1284 int32
	_ = v1284
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1290 int32
	_ = v1290
	var v1300 float64
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1306 int32
	_ = v1306
	var v1310 int64
	_ = v1310
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1320 int64
	_ = v1320
	var v1327 float64
	_ = v1327
	var v1340 int32
	_ = v1340
	var v1350 float64
	_ = v1350
	var v1353 int32
	_ = v1353
	var v1363 float64
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1368 float64
	_ = v1368
	var v1370 int32
	_ = v1370
	var v1372 float64
	_ = v1372
	var v1374 int64
	_ = v1374
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1385 int32
	_ = v1385
	var v1389 int64
	_ = v1389
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1404 float64
	_ = v1404
	var v1406 int64
	_ = v1406
	var v1411 int32
	_ = v1411
	var v1412 int32
	_ = v1412
	var v1416 int64
	_ = v1416
	var v1433 int32
	_ = v1433
	var v1440 int32
	_ = v1440
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1476 int32
	_ = v1476
	var v1505 int32
	_ = v1505
	var v1507 int32
	_ = v1507
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1514 int32
	_ = v1514
	var v1521 int32
	_ = v1521
	var v1551 int32
	_ = v1551
	var v1554 int32
	_ = v1554
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1589 int32
	_ = v1589
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1596 int32
	_ = v1596
	var v1600 int32
	_ = v1600
	var v1605 int32
	_ = v1605
	var v1607 int32
	_ = v1607
	var v1610 int32
	_ = v1610
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1639 int32
	_ = v1639
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1654 int32
	_ = v1654
	var v1685 int32
	_ = v1685
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1691 int32
	_ = v1691
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1702 float64
	_ = v1702
	var v1709 int32
	_ = v1709
	var v1714 int32
	_ = v1714
	var v1732 float64
	_ = v1732
	var v1735 float64
	_ = v1735
	var v1739 int32
	_ = v1739
	var v1741 int32
	_ = v1741
	var v1743 int32
	_ = v1743
	var v1747 float64
	_ = v1747
	var v1749 int32
	_ = v1749
	var v1787 float64
	_ = v1787
	var v1792 int32
	_ = v1792
	var v1795 int32
	_ = v1795
	var v1798 int32
	_ = v1798
	var v1799 int32
	_ = v1799
	var v1802 int32
	_ = v1802
	var v1804 float64
	_ = v1804
	var v1805 int32
	_ = v1805
	var v1811 int32
	_ = v1811
	var v1834 float64
	_ = v1834
	var v1835 float64
	_ = v1835
	var v1841 int32
	_ = v1841
	var v1843 int32
	_ = v1843
	var v1845 int32
	_ = v1845
	var v1849 float64
	_ = v1849
	var v1851 int32
	_ = v1851
	var v1887 float64
	_ = v1887
	var v1894 int32
	_ = v1894
	var v1898 int32
	_ = v1898
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1906 int32
	_ = v1906
	var v1908 float64
	_ = v1908
	var v1910 int32
	_ = v1910
	var v1915 int32
	_ = v1915
	var v1938 float64
	_ = v1938
	var v1939 float64
	_ = v1939
	var v1945 int32
	_ = v1945
	var v1947 int32
	_ = v1947
	var v1949 int32
	_ = v1949
	var v1953 float64
	_ = v1953
	var v1955 int32
	_ = v1955
	var v1991 float64
	_ = v1991
	var v1999 int32
	_ = v1999
	var v2001 int32
	_ = v2001
	var v2003 int32
	_ = v2003
	var v2047 int32
	_ = v2047
	var v2048 float64
	_ = v2048
	var v2049 int32
	_ = v2049
	var v2051 float64
	_ = v2051
	var v2052 int32
	_ = v2052
	var v2054 float64
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2057 int32
	_ = v2057
	var v2061 int64
	_ = v2061
	var v2076 int32
	_ = v2076
	var v2111 float64
	_ = v2111
	var v2112 float64
	_ = v2112
	var v2114 float64
	_ = v2114
	var v2119 float64
	_ = v2119
	var v2124 float64
	_ = v2124
	var v2130 float64
	_ = v2130
	var v2133 float64
	_ = v2133
	var v2136 float64
	_ = v2136
	var v2139 float64
	_ = v2139
	var v2141 float64
	_ = v2141
	var v2144 int32
	_ = v2144
	var v2145 int32
	_ = v2145
	var v2148 float64
	_ = v2148
	var v2155 float64
	_ = v2155
	var v2157 float64
	_ = v2157
	var v2159 float64
	_ = v2159
	var v2160 float64
	_ = v2160
	var v2163 float64
	_ = v2163
	var v2195 float64
	_ = v2195
	var v2204 float64
	_ = v2204
	var v2243 float64
	_ = v2243
	var v2249 int32
	_ = v2249
	var v2253 float64
	_ = v2253
	var v2256 float64
	_ = v2256
	var v2288 float64
	_ = v2288
	var v2299 int32
	_ = v2299
	var v2328 float64
	_ = v2328
	v10 = int32(0)
	v29 = float64(0)
	v37 = m.G0
	v39 = v37 - int32(128)
	m.G0 = v39
	v41 = float64(-1)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v42 == v10 {
		v2328 = v41
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v39 + int32(128)
	return v2328
L2:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+29)))
	if v47 != 0 {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	F_free_attstatsslot(m, v39+int32(92))
	mBase = m.M
	v2299 = m.ExcPending
	if v2299 != 0 {
		goto L13
	} else {
		goto L477
	}
L4:
	;
	v2249 = *(*int32)(unsafe.Add(mBase, uint32(v39)+108))
	v2253 = base.F64_div(float64(0.01), base.F64_convert_i32_s(v2249-int32(1)))
	if base.F64_lt(v2243, v2253) != 0 {
		v2288 = v2253
		goto L3
	} else {
		goto L475
	}
L5:
	;
	v429 = int32(0)
	v437 = v263
	v451 = v425
	goto L59
L6:
	;
	v420 = base.F64_sub(float64(1), v385)
	if base.F64_gt(v380, v420) == int32(0) {
		v2288 = v380
		goto L3
	} else {
		goto L58
	}
L7:
	;
	if v263 == int32(2) {
		goto L50
	} else {
		goto L51
	}
L8:
	;
	v391 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L13
	} else {
		goto L45
	}
L9:
	;
	v59 = v42
	goto L11
L10:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v48 == int32(0) {
		v2328 = v41
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v63 = F_get_attstatsslot(m, v39+int32(92), v59, int32(2), int32(0), int32(1))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L13
	} else {
		goto L16
	}
L12:
	;
	v51 = F_get_func_leakproof(m, v48)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return float64(0)
L14:
	;
	if v51 == int32(0) {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v59 = v57
	goto L11
L16:
	;
	if v63 == int32(0) {
		v2328 = v41
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v39)+108))
	if v67 < int32(2) {
		v274 = v67
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if v274 < int32(2) {
		v2288 = v41
		goto L3
	} else {
		goto L39
	}
L19:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v39)+96))
	if v70 != l6 {
		v274 = v67
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v39)+92))
	if l2 != v73 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v78 = int64(0)
	v80 = F_SearchSysCacheList(m, int32(3), int32(1), base.I64_extend_i32_u(v73), v78, v78)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L13
	} else {
		goto L25
	}
L22:
	;
	v237 = int32(1)
	goto L23
L23:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v39)+108))
	if v237 != 0 {
		goto L7
	} else {
		goto L38
	}
L24:
	;
	F_ReleaseCatCacheList(m, v80)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L13
	} else {
		goto L37
	}
L25:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v80)+56))
	if int32(0) < v82 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v99 = v10
	goto L29
L27:
	;
	goto L28
L28:
	;
	v224 = int32(0)
	goto L24
L29:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v80-int32(-64)+v99<<(uint(int32(2))%32))))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+72))
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+22)))
	v132 = v130 + v131
	v133 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v132)+4)))
	v135 = F_SearchSysCacheExists(m, int32(3), base.I64_extend_i32_u(l2), int64(115), v133, int64(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L13
	} else {
		goto L32
	}
L30:
	;
	goto L28
L31:
	;
	v148 = v99 + int32(1)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v80)+56))
	if v148 < v149 {
		v99 = v148
		goto L29
	} else {
		goto L36
	}
L32:
	;
	if v135 == int32(0) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v132)+24))
	v141 = F_GetIndexAmRoutineByAmId(m, v139, int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L13
	} else {
		goto L34
	}
L34:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+14)))
	if v143 == int32(0) {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	v224 = int32(1)
	goto L24
L36:
	;
	goto L30
L37:
	;
	v237 = v224
	goto L23
L38:
	;
	v274 = v263
	goto L18
L39:
	;
	v302 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+64)) = uint8(v302)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+48)) = uint8(v302)
	v306 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+34)) = uint16(v306)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+32)) = uint8(v302)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+28)) = l6
	*(*int64)(unsafe.Add(mBase, uint32(v39)+20)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = l3
	*(*int64)(unsafe.Add(mBase, uint32(v39)+56)) = l7
	v318 = v302
	v327 = v302
	goto L40
L40:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v39)+104))
	v357 = *(*int64)(unsafe.Add(mBase, uint32(v353+v318<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v39)+40)) = v357
	v359 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+32)) = uint8(v359)
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v363)))
	v365 = m.T0[v364].(func(*base.Module, int32) int64)(m, v39+int32(16))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L13
	} else {
		goto L42
	}
L41:
	;
	v380 = base.F64_div(base.F64_convert_i32_s(v373), base.F64_convert_i32_s(v376))
	v385 = base.F64_div(float64(0.01), base.F64_convert_i32_s(v376-int32(1)))
	if base.F64_lt(v380, v385) == int32(0) {
		goto L6
	} else {
		goto L44
	}
L42:
	;
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+32)))
	v373 = v327 + (v367^int32(-1))&base.B2i32(v365 != int64(0))
	v375 = v318 + int32(1)
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v39)+108))
	if v375 < v376 {
		v318 = v375
		v327 = v373
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v2288 = v385
	goto L3
L45:
	;
	if v391 == int32(0) {
		v2328 = v41
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v395 = F_get_func_name(m, v48)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v395
	F_errmsg_internal(m, int32(_a_F_ineq_histogram_selectivity_0), v39)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L13
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_ineq_histogram_selectivity_1), int32(_a_F_ineq_histogram_selectivity_2), int32(_a_F_ineq_histogram_selectivity_3))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L13
	} else {
		goto L49
	}
L49:
	;
	v2328 = v41
	goto L1
L50:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v39)+92))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v39)+104))
	v412 = F_get_actual_variable_range(m, l0, l1, v408, l6, v409, v409+int32(8))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L13
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	if int32(0) < v263 {
		v425 = v10
		goto L5
	} else {
		goto L54
	}
L53:
	;
	v425 = v412
	goto L5
L54:
	;
	if l4 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v418 = float64(1)
	goto L57
L56:
	;
	v418 = float64(0)
	goto L57
L57:
	;
	v2243 = v418
	goto L4
L58:
	;
	v2288 = v420
	goto L3
L59:
	;
	v463 = v429 + v437
	v464 = int32(2)
	v465 = base.I32_div_s(v463, v464)
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v39)+108))
	if base.B2i32(v466 < int32(3))|base.B2i32(base.Ui32(v464) < base.Ui32(v463+int32(1))) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	if v508 <= int32(0) {
		goto L78
	} else {
		goto L79
	}
L61:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v39)+104))
	v502 = *(*int64)(unsafe.Add(mBase, uint32(v498+v465<<(uint(int32(3))%32))))
	v503 = F_FunctionCall2Coll(m, l3, l6, v502, l7)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L13
	} else {
		goto L68
	}
L62:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v39)+92))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v39)+104))
	v479 = F_get_actual_variable_range(m, l0, l1, v476, l6, v477, int32(0))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L13
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	if base.B2i32(v465 != v466-int32(1))|base.B2i32(v466 < int32(3)) != 0 {
		v495 = v451
		goto L61
	} else {
		goto L66
	}
L65:
	;
	v495 = v479
	goto L61
L66:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v39)+92))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v39)+104))
	v493 = F_get_actual_variable_range(m, l0, l1, v487, l6, int32(0), v489+v465<<(uint(int32(3))%32))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L13
	} else {
		goto L67
	}
L67:
	;
	v495 = v493
	goto L61
L68:
	;
	v507 = l4 ^ base.B2i32(v503 != int64(0))
	if v507 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v508 = v465 + int32(1)
	goto L71
L70:
	;
	v508 = v429
	goto L71
L71:
	;
	if v507 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v509 = v437
	goto L74
L73:
	;
	v509 = v465
	goto L74
L74:
	;
	if v508 < v509 {
		v429 = v508
		v437 = v509
		v451 = v495
		goto L59
	} else {
		goto L75
	}
L75:
	;
	goto L60
L76:
	;
	if v495&int32(1) == int32(0) {
		v2243 = v2195
		goto L4
	} else {
		goto L472
	}
L77:
	;
	if l4 != 0 {
		goto L469
	} else {
		goto L470
	}
L78:
	;
	v2160 = float64(0)
	goto L77
L79:
	;
	goto L80
L80:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v39)+108))
	if v514 <= v508 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v2160 = float64(1)
	goto L77
L82:
	;
	goto L83
L83:
	;
	v517 = l4 ^ l5
	if v517 == int32(1) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v622 = float64(0.5)
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v39)+104))
	v625 = v508 - int32(1)
	v626 = int32(3)
	v629 = *(*int64)(unsafe.Add(mBase, uint32(v623+v625<<(uint(v626)%32))))
	v633 = *(*int64)(unsafe.Add(mBase, uint32(v623+v508<<(uint(v626)%32))))
	v634 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v636 = v39 + int32(72)
	v638 = v39 + int32(80)
	v639 = m.G0
	v640 = int32(16)
	v641 = v639 - v640
	m.G0 = v641
	v643 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v641)+15)) = uint8(v643)
	v646 = v39 + v640
	if l8 <= int32(1081) {
		goto L145
	} else {
		goto L146
	}
L85:
	;
	if v508 != int32(1) {
		v621 = float64(0)
		goto L84
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v524 = v39 + int32(80)
	v525 = int32(0)
	v526 = float64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v524))) = uint8(v525)
	v530 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v530 != 0 {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	goto L87
L89:
	;
	v600 = v39 + int32(16)
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v605 = F_get_attstatsslot(m, v600, v601, int32(1), int32(0), int32(2))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L13
	} else {
		goto L122
	}
L90:
	;
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+28)))
	if v568 != 0 {
		goto L104
	} else {
		goto L105
	}
L91:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v530)+16))
	v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v531)+22)))
	v533 = v531 + v532
	v534 = *(*float32)(unsafe.Add(mBase, uint32(v533)+8))
	v536 = *(*float32)(unsafe.Add(mBase, uint32(v533)+16))
	v563 = base.F64_promote_f32(v536)
	v564 = base.F64_promote_f32(v534)
	goto L90
L92:
	;
	goto L93
L93:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v538 == int32(16) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v563 = float64(2)
	v564 = v526
	goto L90
L95:
	;
	goto L96
L96:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v542 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v549 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v542)+84))
	if v545 != int32(5) {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v563 = float64(-1)
	v564 = v526
	goto L90
L100:
	;
	v563 = float64(0)
	v564 = v526
	goto L90
L101:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v549)))
	if v552 != int32(6) {
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v556 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v549)+8)))
	switch v556 - int32(_a_F_ineq_histogram_selectivity_4) {
	case 0:
		goto L103
	default:
		goto L100
	case 5:
		v563 = float64(-1)
		v564 = v526
		goto L90
	}
L103:
	;
	v563 = float64(1)
	v564 = v526
	goto L90
L104:
	;
	v569 = base.F64_neg(base.F64_sub(float64(1), v564))
	goto L106
L105:
	;
	v569 = v563
	goto L106
L106:
	;
	if base.F64_gt(v569, float64(0)) != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v572 = F_clamp_row_est(m, v569)
	mBase = m.M
	v598 = v572
	goto L89
L108:
	;
	goto L109
L109:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v573 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v576 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v524))) = uint8(v576)
	v598 = float64(200)
	goto L89
L111:
	;
	goto L112
L112:
	;
	v579 = *(*float64)(unsafe.Add(mBase, uint32(v573)+128))
	if base.F64_le(v579, float64(0)) != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v582 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v524))) = uint8(v582)
	v598 = float64(200)
	goto L89
L114:
	;
	goto L115
L115:
	;
	if base.F64_lt(v569, float64(0)) != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v589 = F_clamp_row_est(m, base.F64_mul(v579, base.F64_neg(v569)))
	mBase = m.M
	v598 = v589
	goto L89
L117:
	;
	goto L118
L118:
	;
	if base.F64_lt(v579, float64(200)) != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v592 = F_clamp_row_est(m, v579)
	mBase = m.M
	v598 = v592
	goto L89
L120:
	;
	goto L121
L121:
	;
	v593 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v524))) = uint8(v593)
	v598 = float64(200)
	goto L89
L122:
	;
	if v605 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v39)+40))
	F_free_attstatsslot(m, v600)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L13
	} else {
		goto L126
	}
L124:
	;
	v612 = v598
	goto L125
L125:
	;
	v613 = float64(1)
	if base.F64_gt(v612, v613) != 0 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v612 = base.F64_sub(v598, base.F64_convert_i32_s(v607))
	goto L125
L127:
	;
	v618 = base.F64_div(v613, v612)
	goto L129
L128:
	;
	v618 = float64(0)
	goto L129
L129:
	;
	v621 = v618
	goto L84
L130:
	;
	m.G0 = v641 + int32(16)
	if v2076&int32(1) == int32(0) {
		v2139 = v622
		goto L442
	} else {
		goto L443
	}
L131:
	;
	v2061 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v638))) = v2061
	*(*int64)(unsafe.Add(mBase, uint32(v636))) = v2061
	*(*int64)(unsafe.Add(mBase, uint32(v646))) = v2061
	v2076 = int32(0)
	goto L130
L132:
	;
	v2047 = v641 + int32(15)
	v2048 = F_convert_network_to_scalar(m, l7, l8, v2047)
	mBase = m.M
	v2049 = m.ExcPending
	if v2049 != 0 {
		goto L13
	} else {
		goto L439
	}
L133:
	;
	if l8 != int32(650) {
		goto L131
	} else {
		goto L438
	}
L134:
	;
	v2076 = v1240 ^ int32(1)
	goto L130
L135:
	;
	if v1241 != 0 {
		goto L324
	} else {
		goto L325
	}
L136:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v646))) = v1327
	if v634 <= int32(1183) {
		goto L304
	} else {
		goto L305
	}
L137:
	;
	v1327 = base.F64_convert_i64_s(l7)
	goto L136
L138:
	;
	if l8 != int32(1114) {
		goto L131
	} else {
		goto L296
	}
L139:
	;
	v1315 = base.I32_wrap_i64(l7)
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v1315)+8))
	v1320 = *(*int64)(unsafe.Add(mBase, uint32(v1315)))
	v1327 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1316), float64(1e+06)), base.F64_convert_i64_s(v1320))
	goto L136
L140:
	;
	v1327 = base.F64_convert_i64_s(l7)
	goto L136
L141:
	;
	v1301 = base.I32_wrap_i64(l7)
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(v1301)+12))
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v1301)+8))
	v1310 = *(*int64)(unsafe.Add(mBase, uint32(v1301)))
	v1327 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1302), float64(2.6298e+12)), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1306), float64(8.64e+10)), base.F64_convert_i64_s(v1310)))
	goto L136
L142:
	;
	v1290 = base.I32_wrap_i64(l7)
	if v1290 == int32(-2147483648) {
		goto L290
	} else {
		goto L291
	}
L143:
	;
	v1233 = v641 + int32(15)
	v1234 = F_convert_string_datum(m, l7, l8, l6, v1233)
	mBase = m.M
	v1235 = m.ExcPending
	if v1235 != 0 {
		goto L13
	} else {
		goto L273
	}
L144:
	;
	v1218 = v641 + int32(15)
	v1219 = F_convert_numeric_to_scalar(m, l7, l8, v1218)
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L13
	} else {
		goto L270
	}
L145:
	;
	if l8 <= int32(699) {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	goto L147
L147:
	;
	if l8 <= int32(2201) {
		goto L250
	} else {
		goto L251
	}
L148:
	;
	if base.Ui32(int32(26)) < base.Ui32(l8) {
		goto L133
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	if l8 <= int32(828) {
		goto L242
	} else {
		goto L243
	}
L151:
	;
	v654 = int32(1) << (uint(l8) % 32)
	if v654&int32(95485952) != 0 {
		goto L144
	} else {
		goto L152
	}
L152:
	;
	if v654&int32(34340864) != 0 {
		goto L143
	} else {
		goto L153
	}
L153:
	;
	if l8 != int32(17) {
		goto L133
	} else {
		goto L154
	}
L154:
	;
	if v634 != int32(17) {
		v2076 = int32(0)
		goto L130
	} else {
		goto L155
	}
L155:
	;
	v665 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(l7))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L13
	} else {
		goto L156
	}
L156:
	;
	v668 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v629))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L13
	} else {
		goto L157
	}
L157:
	;
	v671 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v633))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L13
	} else {
		goto L158
	}
L158:
	;
	v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v665))))
	if v673 == int32(1) {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	v703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v668))))
	if v703 == int32(1) {
		goto L171
	} else {
		goto L172
	}
L160:
	;
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v665)+1)))
	if v679 == int32(18) {
		goto L163
	} else {
		goto L164
	}
L161:
	;
	goto L162
L162:
	;
	v690 = int32(1)
	if v673&v690 != 0 {
		v702 = int32(base.Ui32(v673)>>(uint(v690)%32)) - v690
		goto L159
	} else {
		goto L169
	}
L163:
	;
	v682 = int32(16)
	goto L165
L164:
	;
	v682 = int32(0)
	goto L165
L165:
	;
	if base.Ui32((v679-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v689 = int32(4)
	goto L168
L167:
	;
	v689 = v682
	goto L168
L168:
	;
	v702 = v689
	goto L159
L169:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v665)))
	v702 = int32(base.Ui32(v696)>>(uint(int32(2))%32)) - int32(4)
	goto L159
L170:
	;
	v733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v671))))
	if v733 == int32(1) {
		goto L182
	} else {
		goto L183
	}
L171:
	;
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v668)+1)))
	if v709 == int32(18) {
		goto L174
	} else {
		goto L175
	}
L172:
	;
	goto L173
L173:
	;
	v720 = int32(1)
	if v703&v720 != 0 {
		v732 = int32(base.Ui32(v703)>>(uint(v720)%32)) - v720
		goto L170
	} else {
		goto L180
	}
L174:
	;
	v712 = int32(16)
	goto L176
L175:
	;
	v712 = int32(0)
	goto L176
L176:
	;
	if base.Ui32((v709-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v719 = int32(4)
	goto L179
L178:
	;
	v719 = v712
	goto L179
L179:
	;
	v732 = v719
	goto L170
L180:
	;
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v668)))
	v732 = int32(base.Ui32(v726)>>(uint(int32(2))%32)) - int32(4)
	goto L170
L181:
	;
	v763 = int32(1)
	if v733&v763 != 0 {
		goto L192
	} else {
		goto L193
	}
L182:
	;
	v739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v671)+1)))
	if v739 == int32(18) {
		goto L185
	} else {
		goto L186
	}
L183:
	;
	goto L184
L184:
	;
	v750 = int32(1)
	if v733&v750 != 0 {
		v762 = int32(base.Ui32(v733)>>(uint(v750)%32)) - v750
		goto L181
	} else {
		goto L191
	}
L185:
	;
	v742 = int32(16)
	goto L187
L186:
	;
	v742 = int32(0)
	goto L187
L187:
	;
	if base.Ui32((v739-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v749 = int32(4)
	goto L190
L189:
	;
	v749 = v742
	goto L190
L190:
	;
	v762 = v749
	goto L181
L191:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v671)))
	v762 = int32(base.Ui32(v756)>>(uint(int32(2))%32)) - int32(4)
	goto L181
L192:
	;
	v767 = v763
	goto L194
L193:
	;
	v767 = int32(4)
	goto L194
L194:
	;
	v768 = v671 + v767
	v769 = int32(1)
	if v703&v769 != 0 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v773 = v769
	goto L197
L196:
	;
	v773 = int32(4)
	goto L197
L197:
	;
	v774 = v668 + v773
	v775 = int32(1)
	if v673&v775 != 0 {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v779 = v775
	goto L200
L199:
	;
	v779 = int32(4)
	goto L200
L200:
	;
	v780 = v665 + v779
	if v702 < v732 {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	if int32(0) < v856 {
		goto L214
	} else {
		goto L215
	}
L202:
	;
	v782 = v702
	goto L204
L203:
	;
	v782 = v732
	goto L204
L204:
	;
	if v782 < v762 {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v784 = v782
	goto L207
L206:
	;
	v784 = v762
	goto L207
L207:
	;
	if v784 <= int32(0) {
		v853 = v780
		v854 = v768
		v856 = v702
		v858 = v774
		v859 = v762
		v861 = v732
		goto L201
	} else {
		goto L208
	}
L208:
	;
	v797 = v780
	v798 = v768
	v800 = v702
	v802 = v774
	v803 = v762
	v805 = v732
	v816 = int32(0)
	goto L209
L209:
	;
	v833 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v802))))
	v834 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v798))))
	if v833 != v834 {
		v853 = v797
		v854 = v798
		v856 = v800
		v858 = v802
		v859 = v803
		v861 = v805
		goto L201
	} else {
		goto L211
	}
L210:
	;
	v853 = v779 + v665 + v784
	v854 = v767 + v671 + v784
	v856 = v702 - v784
	v858 = v773 + v668 + v784
	v859 = v762 - v784
	v861 = v732 - v784
	goto L201
L211:
	;
	v836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v797))))
	if v833 != v836 {
		v853 = v797
		v854 = v798
		v856 = v800
		v858 = v802
		v859 = v803
		v861 = v805
		goto L201
	} else {
		goto L212
	}
L212:
	;
	v838 = int32(1)
	v851 = v816 + v838
	if v851 != v784 {
		v797 = v797 + v838
		v798 = v798 + v838
		v800 = v800 - v838
		v802 = v802 + v838
		v803 = v803 - v838
		v805 = v805 - v838
		v816 = v851
		goto L209
	} else {
		goto L213
	}
L213:
	;
	goto L210
L214:
	;
	v891 = int32(10)
	if base.Ui32(v891) <= base.Ui32(v856) {
		goto L217
	} else {
		goto L218
	}
L215:
	;
	v976 = v29
	goto L216
L216:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v646))) = v976
	if int32(0) < v861 {
		goto L223
	} else {
		goto L224
	}
L217:
	;
	v894 = v891
	goto L219
L218:
	;
	v894 = v856
	goto L219
L219:
	;
	v896 = v853
	v899 = v894
	v925 = float64(256)
	v928 = v29
	goto L220
L220:
	;
	v932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v896))))
	v935 = base.F64_add(v928, base.F64_div(base.F64_convert_i32_u(v932), v925))
	v936 = int32(1)
	if base.Ui32(v936) < base.Ui32(v899) {
		v896 = v896 + v936
		v899 = v899 - v936
		v925 = base.F64_mul(v925, float64(256))
		v928 = v935
		goto L220
	} else {
		goto L222
	}
L221:
	;
	v976 = v935
	goto L216
L222:
	;
	goto L221
L223:
	;
	v983 = int32(10)
	if base.Ui32(v983) <= base.Ui32(v861) {
		goto L226
	} else {
		goto L227
	}
L224:
	;
	v1066 = v29
	goto L225
L225:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v636))) = v1066
	if v859 <= int32(0) {
		goto L233
	} else {
		goto L234
	}
L226:
	;
	v986 = v983
	goto L228
L227:
	;
	v986 = v861
	goto L228
L228:
	;
	v988 = v986
	v993 = v858
	v1017 = float64(256)
	v1018 = v29
	goto L229
L229:
	;
	v1024 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v993))))
	v1027 = base.F64_add(v1018, base.F64_div(base.F64_convert_i32_u(v1024), v1017))
	v1028 = int32(1)
	if base.Ui32(v1028) < base.Ui32(v988) {
		v988 = v988 - v1028
		v993 = v993 + v1028
		v1017 = base.F64_mul(v1017, float64(256))
		v1018 = v1027
		goto L229
	} else {
		goto L231
	}
L230:
	;
	v1066 = v1027
	goto L225
L231:
	;
	goto L230
L232:
	;
	v2076 = int32(1)
	goto L130
L233:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v638))) = float64(0)
	goto L232
L234:
	;
	goto L235
L235:
	;
	v1077 = int32(10)
	if base.Ui32(v1077) <= base.Ui32(v859) {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v1080 = v1077
	goto L238
L237:
	;
	v1080 = v859
	goto L238
L238:
	;
	v1084 = v854
	v1088 = v1080
	v1112 = float64(256)
	v1113 = float64(0)
	goto L239
L239:
	;
	v1119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1084))))
	v1122 = base.F64_add(v1113, base.F64_div(base.F64_convert_i32_u(v1119), v1112))
	v1123 = int32(1)
	if base.Ui32(v1123) < base.Ui32(v1088) {
		v1084 = v1084 + v1123
		v1088 = v1088 - v1123
		v1112 = base.F64_mul(v1112, float64(256))
		v1113 = v1122
		goto L239
	} else {
		goto L241
	}
L240:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v638))) = v1122
	goto L232
L241:
	;
	goto L240
L242:
	;
	if base.Ui32(l8-int32(700)) < base.Ui32(int32(2)) {
		goto L144
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	if base.Ui32(l8-int32(1042)) < base.Ui32(int32(2)) {
		goto L143
	} else {
		goto L247
	}
L245:
	;
	if l8 == int32(774) {
		goto L132
	} else {
		goto L246
	}
L246:
	;
	goto L131
L247:
	;
	if l8 == int32(829) {
		goto L132
	} else {
		goto L248
	}
L248:
	;
	if l8 != int32(869) {
		goto L131
	} else {
		goto L249
	}
L249:
	;
	goto L132
L250:
	;
	if l8 <= int32(1183) {
		goto L253
	} else {
		goto L254
	}
L251:
	;
	goto L252
L252:
	;
	if l8 <= int32(3768) {
		goto L259
	} else {
		goto L260
	}
L253:
	;
	switch l8 - int32(1082) {
	case 0:
		goto L142
	case 1:
		goto L140
	default:
		goto L138
	}
L254:
	;
	goto L255
L255:
	;
	switch l8 - int32(1184) {
	case 0:
		goto L137
	case 1:
		goto L131
	case 2:
		goto L141
	default:
		goto L256
	}
L256:
	;
	if l8 == int32(1266) {
		goto L139
	} else {
		goto L257
	}
L257:
	;
	if l8 == int32(1700) {
		goto L144
	} else {
		goto L258
	}
L258:
	;
	goto L131
L259:
	;
	if base.B2i32(l8 == int32(3734))|base.B2i32(base.Ui32(l8-int32(2202)) < base.Ui32(int32(5))) != 0 {
		goto L144
	} else {
		goto L262
	}
L260:
	;
	goto L261
L261:
	;
	if l8 <= int32(_a_F_ineq_histogram_selectivity_5) {
		goto L263
	} else {
		goto L264
	}
L262:
	;
	goto L131
L263:
	;
	switch l8 - int32(4089) {
	case 0, 7:
		goto L144
	case 1, 2, 3, 4, 5, 6:
		goto L131
	default:
		goto L266
	}
L264:
	;
	goto L265
L265:
	;
	if l8 == int32(_a_F_ineq_histogram_selectivity_6) {
		goto L144
	} else {
		goto L268
	}
L266:
	;
	if l8 == int32(3769) {
		goto L144
	} else {
		goto L267
	}
L267:
	;
	goto L131
L268:
	;
	if l8 != int32(_a_F_ineq_histogram_selectivity_7) {
		goto L131
	} else {
		goto L269
	}
L269:
	;
	goto L144
L270:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v646))) = v1219
	v1222 = F_convert_numeric_to_scalar(m, v629, v634, v1218)
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L13
	} else {
		goto L271
	}
L271:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v636))) = v1222
	v1225 = F_convert_numeric_to_scalar(m, v633, v634, v1218)
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L13
	} else {
		goto L272
	}
L272:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v638))) = v1225
	v1228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641)+15)))
	v2076 = v1228 ^ int32(1)
	goto L130
L273:
	;
	v1236 = F_convert_string_datum(m, v629, v634, l6, v1233)
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		goto L13
	} else {
		goto L274
	}
L274:
	;
	v1238 = F_convert_string_datum(m, v633, v634, l6, v1233)
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
		goto L13
	} else {
		goto L275
	}
L275:
	;
	v1240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641)+15)))
	if v1240 != 0 {
		goto L134
	} else {
		goto L276
	}
L276:
	;
	v1241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1238))))
	v1242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1236))))
	if v1242 == int32(0) {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	v1433 = v1241
	v1440 = v1241
	goto L135
L278:
	;
	goto L279
L279:
	;
	v1246 = v1241
	v1248 = v1236
	v1251 = v1242
	v1253 = v1241
	goto L280
L280:
	;
	v1282 = v1251 & int32(255)
	if base.Ui32(v1282) < base.Ui32(v1246) {
		goto L282
	} else {
		goto L283
	}
L281:
	;
	v1433 = v1284
	v1440 = v1286
	goto L135
L282:
	;
	v1284 = v1246
	goto L284
L283:
	;
	v1284 = v1282
	goto L284
L284:
	;
	if base.Ui32(v1253) < base.Ui32(v1282) {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1286 = v1253
	goto L287
L286:
	;
	v1286 = v1282
	goto L287
L287:
	;
	v1287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1248)+1)))
	if v1287 != 0 {
		v1246 = v1284
		v1248 = v1248 + int32(1)
		v1251 = v1287
		v1253 = v1286
		goto L280
	} else {
		goto L288
	}
L288:
	;
	goto L281
L289:
	;
	v1327 = v1300
	goto L136
L290:
	;
	v1300 = float64(-1.7976931348623157e+308)
	goto L289
L291:
	;
	goto L292
L292:
	;
	if v1290 == int32(2147483647) {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	v1300 = float64(1.7976931348623157e+308)
	goto L289
L294:
	;
	goto L295
L295:
	;
	v1300 = base.F64_mul(base.F64_convert_i32_s(v1290), float64(8.64e+10))
	goto L289
L296:
	;
	goto L137
L297:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v636))) = base.F64_convert_i64_s(v629)
	*(*float64)(unsafe.Add(mBase, uint32(v638))) = base.F64_convert_i64_s(v633)
	v2076 = int32(1)
	goto L130
L298:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v636))) = int64(0)
	*(*float64)(unsafe.Add(mBase, uint32(v638))) = float64(0)
	v2076 = int32(0)
	goto L130
L299:
	;
	if v634 == int32(1114) {
		goto L297
	} else {
		goto L323
	}
L300:
	;
	if v634 != int32(1266) {
		goto L298
	} else {
		goto L322
	}
L301:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v636))) = base.F64_convert_i64_s(v629)
	*(*float64)(unsafe.Add(mBase, uint32(v638))) = base.F64_convert_i64_s(v633)
	v2076 = int32(1)
	goto L130
L302:
	;
	v1365 = base.I32_wrap_i64(v629)
	v1366 = *(*int32)(unsafe.Add(mBase, uint32(v1365)+12))
	v1368 = float64(2.6298e+12)
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v1365)+8))
	v1372 = float64(8.64e+10)
	v1374 = *(*int64)(unsafe.Add(mBase, uint32(v1365)))
	*(*float64)(unsafe.Add(mBase, uint32(v636))) = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1366), v1368), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1370), v1372), base.F64_convert_i64_s(v1374)))
	v1380 = base.I32_wrap_i64(v633)
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v1380)+12))
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v1380)+8))
	v1389 = *(*int64)(unsafe.Add(mBase, uint32(v1380)))
	*(*float64)(unsafe.Add(mBase, uint32(v638))) = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1381), v1368), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1385), v1372), base.F64_convert_i64_s(v1389)))
	v2076 = int32(1)
	goto L130
L303:
	;
	v1340 = base.I32_wrap_i64(v629)
	if v1340 == int32(-2147483648) {
		goto L309
	} else {
		goto L310
	}
L304:
	;
	switch v634 - int32(1082) {
	case 0:
		goto L303
	case 1:
		goto L301
	default:
		goto L299
	}
L305:
	;
	goto L306
L306:
	;
	switch v634 - int32(1184) {
	case 0:
		goto L307
	case 1:
		goto L298
	case 2:
		goto L302
	default:
		goto L300
	}
L307:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v636))) = base.F64_convert_i64_s(v629)
	*(*float64)(unsafe.Add(mBase, uint32(v638))) = base.F64_convert_i64_s(v633)
	v2076 = int32(1)
	goto L130
L308:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v636))) = v1350
	v1353 = base.I32_wrap_i64(v633)
	if v1353 == int32(-2147483648) {
		goto L316
	} else {
		goto L317
	}
L309:
	;
	v1350 = float64(-1.7976931348623157e+308)
	goto L308
L310:
	;
	goto L311
L311:
	;
	if v1340 == int32(2147483647) {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v1350 = float64(1.7976931348623157e+308)
	goto L308
L313:
	;
	goto L314
L314:
	;
	v1350 = base.F64_mul(base.F64_convert_i32_s(v1340), float64(8.64e+10))
	goto L308
L315:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v638))) = v1363
	v2076 = int32(1)
	goto L130
L316:
	;
	v1363 = float64(-1.7976931348623157e+308)
	goto L315
L317:
	;
	goto L318
L318:
	;
	if v1353 == int32(2147483647) {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	v1363 = float64(1.7976931348623157e+308)
	goto L315
L320:
	;
	goto L321
L321:
	;
	v1363 = base.F64_mul(base.F64_convert_i32_s(v1353), float64(8.64e+10))
	goto L315
L322:
	;
	v1401 = base.I32_wrap_i64(v629)
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(v1401)+8))
	v1404 = float64(1e+06)
	v1406 = *(*int64)(unsafe.Add(mBase, uint32(v1401)))
	*(*float64)(unsafe.Add(mBase, uint32(v636))) = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1402), v1404), base.F64_convert_i64_s(v1406))
	v1411 = base.I32_wrap_i64(v633)
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(v1411)+8))
	v1416 = *(*int64)(unsafe.Add(mBase, uint32(v1411)))
	*(*float64)(unsafe.Add(mBase, uint32(v638))) = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1412), v1404), base.F64_convert_i64_s(v1416))
	v2076 = int32(1)
	goto L130
L323:
	;
	goto L298
L324:
	;
	v1468 = v1241
	v1469 = v1433
	v1471 = v1238
	v1476 = v1440
	goto L327
L325:
	;
	v1514 = v1433
	v1521 = v1440
	goto L326
L326:
	;
	v1551 = int32(90)
	if v1514 <= v1551 {
		goto L336
	} else {
		goto L337
	}
L327:
	;
	v1505 = v1468 & int32(255)
	if v1505 < v1469 {
		goto L329
	} else {
		goto L330
	}
L328:
	;
	v1514 = v1507
	v1521 = v1509
	goto L326
L329:
	;
	v1507 = v1469
	goto L331
L330:
	;
	v1507 = v1505
	goto L331
L331:
	;
	if v1476 < v1505 {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	v1509 = v1476
	goto L334
L333:
	;
	v1509 = v1505
	goto L334
L334:
	;
	v1510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1471)+1)))
	if v1510 != 0 {
		v1468 = v1510
		v1469 = v1507
		v1471 = v1471 + int32(1)
		v1476 = v1509
		goto L327
	} else {
		goto L335
	}
L335:
	;
	goto L328
L336:
	;
	v1554 = v1551
	goto L338
L337:
	;
	v1554 = v1514
	goto L338
L338:
	;
	v1559 = base.B2i32(v1521 < int32(91)) & base.B2i32(int32(64) < v1514)
	if v1559 != 0 {
		goto L339
	} else {
		goto L340
	}
L339:
	;
	v1560 = v1554
	goto L341
L340:
	;
	v1560 = v1514
	goto L341
L341:
	;
	if v1560 <= int32(122) {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	v1563 = int32(122)
	goto L344
L343:
	;
	v1563 = v1560
	goto L344
L344:
	;
	v1564 = int32(65)
	if v1564 <= v1521 {
		goto L345
	} else {
		goto L346
	}
L345:
	;
	v1567 = v1564
	goto L347
L346:
	;
	v1567 = v1521
	goto L347
L347:
	;
	if v1559 != 0 {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	v1568 = v1567
	goto L350
L349:
	;
	v1568 = v1521
	goto L350
L350:
	;
	v1573 = base.B2i32(v1568 < int32(123)) & base.B2i32(int32(96) < v1560)
	if v1573 != 0 {
		goto L351
	} else {
		goto L352
	}
L351:
	;
	v1574 = v1563
	goto L353
L352:
	;
	v1574 = v1560
	goto L353
L353:
	;
	if v1574 <= int32(57) {
		goto L354
	} else {
		goto L355
	}
L354:
	;
	v1577 = int32(57)
	goto L356
L355:
	;
	v1577 = v1574
	goto L356
L356:
	;
	v1578 = int32(97)
	if v1578 <= v1568 {
		goto L357
	} else {
		goto L358
	}
L357:
	;
	v1581 = v1578
	goto L359
L358:
	;
	v1581 = v1568
	goto L359
L359:
	;
	if v1573 != 0 {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	v1582 = v1581
	goto L362
L361:
	;
	v1582 = v1568
	goto L362
L362:
	;
	v1587 = base.B2i32(v1582 < int32(58)) & base.B2i32(int32(47) < v1574)
	if v1587 != 0 {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	v1588 = v1577
	goto L365
L364:
	;
	v1588 = v1574
	goto L365
L365:
	;
	v1589 = int32(48)
	if v1589 <= v1582 {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	v1592 = v1589
	goto L368
L367:
	;
	v1592 = v1582
	goto L368
L368:
	;
	if v1587 != 0 {
		goto L369
	} else {
		goto L370
	}
L369:
	;
	v1593 = v1592
	goto L371
L370:
	;
	v1593 = v1582
	goto L371
L371:
	;
	v1596 = base.B2i32(v1588-v1593 < int32(9))
	if v1242 == int32(0) {
		v1648 = v1236
		v1649 = v1238
		v1654 = v1234
		goto L372
	} else {
		goto L373
	}
L372:
	;
	if v1588-v1593 < int32(9) {
		goto L383
	} else {
		goto L384
	}
L373:
	;
	v1600 = v1238
	v1605 = v1234
	v1607 = v1236
	v1610 = v1242
	goto L374
L374:
	;
	v1636 = v1610 & int32(255)
	v1637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1600))))
	if v1636 != v1637 {
		goto L376
	} else {
		goto L377
	}
L375:
	;
	v1648 = v1647
	v1649 = v1644
	v1654 = v1642
	goto L372
L376:
	;
	v1648 = v1607
	v1649 = v1600
	v1654 = v1605
	goto L372
L377:
	;
	goto L378
L378:
	;
	v1639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1605))))
	if v1639 != v1636 {
		goto L379
	} else {
		goto L380
	}
L379:
	;
	v1648 = v1607
	v1649 = v1600
	v1654 = v1605
	goto L372
L380:
	;
	goto L381
L381:
	;
	v1641 = int32(1)
	v1642 = v1605 + v1641
	v1644 = v1600 + v1641
	v1645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1607)+1)))
	v1647 = v1607 + v1641
	if v1645 != 0 {
		v1600 = v1644
		v1605 = v1642
		v1607 = v1647
		v1610 = v1645
		goto L374
	} else {
		goto L382
	}
L382:
	;
	goto L375
L383:
	;
	v1685 = int32(127)
	goto L385
L384:
	;
	v1685 = v1588
	goto L385
L385:
	;
	if v1588-v1593 < int32(9) {
		goto L386
	} else {
		goto L387
	}
L386:
	;
	v1687 = int32(32)
	goto L388
L387:
	;
	v1687 = v1593
	goto L388
L388:
	;
	v1688 = F_strlen(m, v1654)
	mBase = m.M
	if int32(0) < v1688 {
		goto L389
	} else {
		goto L390
	}
L389:
	;
	v1691 = int32(12)
	if base.Ui32(v1691) <= base.Ui32(v1688) {
		goto L392
	} else {
		goto L393
	}
L390:
	;
	v1787 = v29
	goto L391
L391:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v646))) = v1787
	v1792 = F_strlen(m, v1648)
	mBase = m.M
	if int32(0) < v1792 {
		goto L404
	} else {
		goto L405
	}
L392:
	;
	v1694 = v1691
	goto L394
L393:
	;
	v1694 = v1688
	goto L394
L394:
	;
	v1695 = int32(1)
	v1702 = base.F64_convert_i32_s(v1685 - v1687 + v1695)
	v1709 = v1654
	v1714 = v1694
	v1732 = v1702
	v1735 = v29
	goto L395
L395:
	;
	v1739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1709))))
	if base.Ui32(v1685) < base.Ui32(v1739) {
		goto L397
	} else {
		goto L398
	}
L396:
	;
	v1787 = v1747
	goto L391
L397:
	;
	v1741 = v1685 + v1695
	goto L399
L398:
	;
	v1741 = v1739
	goto L399
L399:
	;
	if base.Ui32(v1739) < base.Ui32(v1687) {
		goto L400
	} else {
		goto L401
	}
L400:
	;
	v1743 = v1687 - v1695
	goto L402
L401:
	;
	v1743 = v1741
	goto L402
L402:
	;
	v1747 = base.F64_add(v1735, base.F64_div(base.F64_convert_i32_s(v1743-v1687), v1732))
	v1749 = int32(1)
	if base.Ui32(v1749) < base.Ui32(v1714) {
		v1709 = v1709 + v1749
		v1714 = v1714 - v1749
		v1732 = base.F64_mul(v1732, v1702)
		v1735 = v1747
		goto L395
	} else {
		goto L403
	}
L403:
	;
	goto L396
L404:
	;
	v1795 = int32(12)
	if base.Ui32(v1795) <= base.Ui32(v1792) {
		goto L407
	} else {
		goto L408
	}
L405:
	;
	v1887 = v29
	goto L406
L406:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v636))) = v1887
	v1894 = F_strlen(m, v1649)
	mBase = m.M
	if v1894 <= int32(0) {
		goto L420
	} else {
		goto L421
	}
L407:
	;
	v1798 = v1795
	goto L409
L408:
	;
	v1798 = v1792
	goto L409
L409:
	;
	v1799 = int32(1)
	v1802 = v1685 + v1799
	v1804 = base.F64_convert_i32_s(v1802 - v1687)
	v1805 = v1648
	v1811 = v1798
	v1834 = v1804
	v1835 = v29
	goto L410
L410:
	;
	v1841 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1805))))
	if base.Ui32(v1685) < base.Ui32(v1841) {
		goto L412
	} else {
		goto L413
	}
L411:
	;
	v1887 = v1849
	goto L406
L412:
	;
	v1843 = v1802
	goto L414
L413:
	;
	v1843 = v1841
	goto L414
L414:
	;
	if base.Ui32(v1841) < base.Ui32(v1687) {
		goto L415
	} else {
		goto L416
	}
L415:
	;
	v1845 = v1687 - v1799
	goto L417
L416:
	;
	v1845 = v1843
	goto L417
L417:
	;
	v1849 = base.F64_add(v1835, base.F64_div(base.F64_convert_i32_s(v1845-v1687), v1834))
	v1851 = int32(1)
	if base.Ui32(v1851) < base.Ui32(v1811) {
		v1805 = v1805 + v1851
		v1811 = v1811 - v1851
		v1834 = base.F64_mul(v1834, v1804)
		v1835 = v1849
		goto L410
	} else {
		goto L418
	}
L418:
	;
	goto L411
L419:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v638))) = v1991
	F_pfree(m, v1234)
	mBase = m.M
	v1999 = m.ExcPending
	if v1999 != 0 {
		goto L13
	} else {
		goto L435
	}
L420:
	;
	v1991 = float64(0)
	goto L419
L421:
	;
	goto L422
L422:
	;
	v1898 = int32(12)
	if base.Ui32(v1898) <= base.Ui32(v1894) {
		goto L423
	} else {
		goto L424
	}
L423:
	;
	v1901 = v1898
	goto L425
L424:
	;
	v1901 = v1894
	goto L425
L425:
	;
	v1902 = int32(1)
	v1906 = v1685 + v1902
	v1908 = base.F64_convert_i32_s(v1906 - v1687)
	v1910 = v1649
	v1915 = v1901
	v1938 = v1908
	v1939 = float64(0)
	goto L426
L426:
	;
	v1945 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1910))))
	if base.Ui32(v1685) < base.Ui32(v1945) {
		goto L428
	} else {
		goto L429
	}
L427:
	;
	v1991 = v1953
	goto L419
L428:
	;
	v1947 = v1906
	goto L430
L429:
	;
	v1947 = v1945
	goto L430
L430:
	;
	if base.Ui32(v1945) < base.Ui32(v1687) {
		goto L431
	} else {
		goto L432
	}
L431:
	;
	v1949 = v1687 - v1902
	goto L433
L432:
	;
	v1949 = v1947
	goto L433
L433:
	;
	v1953 = base.F64_add(v1939, base.F64_div(base.F64_convert_i32_s(v1949-v1687), v1938))
	v1955 = int32(1)
	if base.Ui32(v1955) < base.Ui32(v1915) {
		v1910 = v1910 + v1955
		v1915 = v1915 - v1955
		v1938 = base.F64_mul(v1938, v1908)
		v1939 = v1953
		goto L426
	} else {
		goto L434
	}
L434:
	;
	goto L427
L435:
	;
	F_pfree(m, v1236)
	mBase = m.M
	v2001 = m.ExcPending
	if v2001 != 0 {
		goto L13
	} else {
		goto L436
	}
L436:
	;
	F_pfree(m, v1238)
	mBase = m.M
	v2003 = m.ExcPending
	if v2003 != 0 {
		goto L13
	} else {
		goto L437
	}
L437:
	;
	goto L134
L438:
	;
	goto L132
L439:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v646))) = v2048
	v2051 = F_convert_network_to_scalar(m, v629, v634, v2047)
	mBase = m.M
	v2052 = m.ExcPending
	if v2052 != 0 {
		goto L13
	} else {
		goto L440
	}
L440:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v636))) = v2051
	v2054 = F_convert_network_to_scalar(m, v633, v634, v2047)
	mBase = m.M
	v2055 = m.ExcPending
	if v2055 != 0 {
		goto L13
	} else {
		goto L441
	}
L441:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v638))) = v2054
	v2057 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641)+15)))
	v2076 = v2057 ^ int32(1)
	goto L130
L442:
	;
	v2141 = float64(1)
	v2144 = *(*int32)(unsafe.Add(mBase, uint32(v39)+108))
	v2145 = int32(1)
	v2148 = base.F64_div(base.F64_add(v2139, base.F64_convert_i32_u(v625)), base.F64_convert_i32_s(v2144-v2145))
	if v508 != v2145 {
		goto L460
	} else {
		goto L461
	}
L443:
	;
	v2111 = *(*float64)(unsafe.Add(mBase, uint32(v39)+80))
	v2112 = *(*float64)(unsafe.Add(mBase, uint32(v39)+72))
	if base.F64_le(v2111, v2112) != 0 {
		v2139 = v622
		goto L442
	} else {
		goto L444
	}
L444:
	;
	v2114 = *(*float64)(unsafe.Add(mBase, uint32(v39)+16))
	if base.F64_ge(v2112, v2114) != 0 {
		goto L445
	} else {
		goto L446
	}
L445:
	;
	v2139 = float64(0)
	goto L442
L446:
	;
	goto L447
L447:
	;
	if base.F64_ge(v2114, v2111) != 0 {
		goto L448
	} else {
		goto L449
	}
L448:
	;
	v2139 = float64(1)
	goto L442
L449:
	;
	goto L450
L450:
	;
	v2119 = float64(0.5)
	v2124 = base.F64_div(base.F64_sub(v2114, v2112), base.F64_sub(v2111, v2112))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v2124)&int64(9223372036854775807)) {
		goto L451
	} else {
		goto L452
	}
L451:
	;
	v2130 = v2119
	goto L453
L452:
	;
	v2130 = v2124
	goto L453
L453:
	;
	if base.F64_lt(v2124, float64(0)) != 0 {
		goto L454
	} else {
		goto L455
	}
L454:
	;
	v2133 = v2119
	goto L456
L455:
	;
	v2133 = v2130
	goto L456
L456:
	;
	if base.F64_gt(v2124, float64(1)) != 0 {
		goto L457
	} else {
		goto L458
	}
L457:
	;
	v2136 = v2119
	goto L459
L458:
	;
	v2136 = v2133
	goto L459
L459:
	;
	v2139 = v2136
	goto L442
L460:
	;
	v2155 = v2148
	goto L462
L461:
	;
	v2155 = base.F64_add(v2148, base.F64_mul(v621, base.F64_sub(v2141, v2139)))
	goto L462
L462:
	;
	if v517 != 0 {
		goto L463
	} else {
		goto L464
	}
L463:
	;
	v2157 = v2155
	goto L465
L464:
	;
	v2157 = base.F64_sub(v2155, v621)
	goto L465
L465:
	;
	if l4 != 0 {
		goto L466
	} else {
		goto L467
	}
L466:
	;
	v2159 = base.F64_sub(v2141, v2157)
	goto L468
L467:
	;
	v2159 = v2157
	goto L468
L468:
	;
	v2195 = v2159
	goto L76
L469:
	;
	v2163 = base.F64_sub(float64(1), v2160)
	goto L471
L470:
	;
	v2163 = v2160
	goto L471
L471:
	;
	v2195 = v2163
	goto L76
L472:
	;
	v2204 = float64(0)
	if base.F64_lt(v2195, v2204) != 0 {
		v2288 = v2204
		goto L3
	} else {
		goto L473
	}
L473:
	;
	if base.F64_gt(v2195, float64(1)) == int32(0) {
		v2288 = v2195
		goto L3
	} else {
		goto L474
	}
L474:
	;
	v2288 = float64(1)
	goto L3
L475:
	;
	v2256 = base.F64_sub(float64(1), v2253)
	if base.F64_lt(v2256, v2243) == int32(0) {
		v2288 = v2243
		goto L3
	} else {
		goto L476
	}
L476:
	;
	v2288 = v2256
	goto L3
L477:
	;
	v2328 = v2288
	goto L1
}
func F_inetmi_int8(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v10 = F_internal_inetpl(m, v3, int64(0)-v8)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v10)
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
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
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
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v645 int32
	_ = v645
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
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
	v494 = F_palloc_mul(m, int32(1), v489)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+84)) = v494
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v494
	F_infix_1(m, v16+int32(76), v418, v412)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v16)+76))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v502
	F_infix_1(m, l0, v418, int32(0))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v510 = v508 - v509
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v16)+84))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v16)+80))
	if v507 <= v510+v511-v513+int32(16) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v522 = v509
	v523 = v507
	goto L94
L92:
	;
	v548 = v513
	v551 = v508
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
	v534 = v523 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v534
	v536 = F_repalloc(m, v522, v534)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L1
	} else {
		goto L96
	}
L95:
	;
	v548 = v544
	v551 = v539
	goto L93
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v536
	v539 = v536 + v510
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v539
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v16)+84))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v16)+80))
	if v541 <= v510+int32(16)+v542-v544 {
		v522 = v536
		v523 = v541
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v603 = F_strlen(m, v602)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v603 + v602
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v16)+80))
	F_pfree(m, v606)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L1
	} else {
		goto L113
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v548
	v600 = F_pg_sprintf(m, v551, int32(_a_F_infix_1_2), v16+int32(16))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L1
	} else {
		goto L112
	}
L100:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v548
	v566 = F_pg_sprintf(m, v551, int32(_a_F_infix_1_3), v16+int32(32))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	goto L98
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v548
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v407
	v575 = F_pg_sprintf(m, v551, int32(_a_F_infix_1_4), v16-int32(-64))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v548
	v581 = F_pg_sprintf(m, v551, int32(_a_F_infix_1_5), v16+int32(48))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
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
	v590 = m.ExcPending
	if v590 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_infix_1_7), int32(1130), int32(_a_F_infix_1_8))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
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
	v611 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v613 = v611 - v612
	v615 = v613 + int32(3)
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v616 <= v615 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v620 = v612
	v621 = v616
	goto L118
L116:
	;
	v645 = v611
	goto L117
L117:
	;
	v656 = F_pg_sprintf(m, v645, int32(_a_F_infix_1_1), int32(0))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L1
	} else {
		goto L122
	}
L118:
	;
	v632 = v621 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v632
	v634 = F_repalloc(m, v620, v632)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L1
	} else {
		goto L120
	}
L119:
	;
	v645 = v637
	goto L117
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v634
	v637 = v613 + v634
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v637
	v639 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v639 <= v615 {
		v620 = v634
		v621 = v639
		goto L118
	} else {
		goto L121
	}
L121:
	;
	goto L119
L122:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v659 = F_strlen(m, v658)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v659 + v658
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
func F_int24le(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	return base.I64_extend_i32_u(base.B2i32(v3 <= v2))
}
func F_int28le(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	return base.I64_extend_i32_u(base.B2i32(v3 <= v2))
}
func F_int2in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = F_pg_strtoint16_safe(m, v2, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_s(v4)
	}
}
func F_int2int4_sum(m *base.Module, l0 int32) int64 {
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
	var v29 int64
	_ = v29
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
		if v8 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int64(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_int2int4_sum_0), int32(0))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int2int4_sum_1), int32(_a_F_int2int4_sum_2), int32(_a_F_int2int4_sum_3))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int64(0)
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
				v34 = m.ExcPending
				if v34 != 0 {
					return int64(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_int2int4_sum_0), int32(0))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_int2int4_sum_1), int32(_a_F_int2int4_sum_2), int32(_a_F_int2int4_sum_3))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int64(0)
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
					return int64(0)
				} else {
					v29 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
					return v29
				}
			}
		}
	}
}
func F_int2larger(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	v4 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
	if v4 < v3 {
		v6 = v3
	} else {
		v6 = v4
	}
	return base.I64_extend_i32_s(v6)
}
func F_int2ne(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)))
	v3 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+40)))
	return base.I64_extend_i32_u(base.B2i32(v2 != v3))
}
func F_int2not(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	return base.I64_extend16_s(v2 ^ int64(-1))
}
func F_int2shl(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	return base.I64_extend16_s(base.I64_extend_i32_u(v2 << (uint(v3) % 32)))
}
func F_int42gt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
	return base.I64_extend_i32_u(base.B2i32(v3 < v2))
}
func F_int42le(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
	return base.I64_extend_i32_u(base.B2i32(v2 <= v3))
}
func F_int48ne(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+24)))
	return base.I64_extend_i32_u(base.B2i32(v2 != v3))
}
func F_int4xor(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	return base.I64_extend32_s(v2 ^ v3)
}
func F_int82div(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
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
	var v35 int64
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = int32(_a_F_int82div_0)
	v8 = base.I32_wrap_i64(v5) & v7
	if v8 != v7 {
		if v8 != 0 {
			v35 = base.I64_div_s(v4, base.I64_extend16_s(v5))
			return v35
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(33816706))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_int82div_1), int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_int82div_2), int32(1123), int32(_a_F_int82div_3))
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
		}
	} else {
		if v4 == int64(-9223372036854775807-1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_int82div_4), int32(0))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_int82div_2), int32(1139), int32(_a_F_int82div_3))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
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
			return int64(0) - v4
		}
	}
}
func F_int82ne(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
	return base.I64_extend_i32_u(base.B2i32(v2 != v3))
}
func F_int84gt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+40)))
	return base.I64_extend_i32_u(base.B2i32(v3 < v2))
}
func F_int84ne(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+40)))
	return base.I64_extend_i32_u(base.B2i32(v2 != v3))
}
func F_int8div(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v37 int64
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = int64(1)
	v8 = v6 + v7
	if base.Ui64(v8) <= base.Ui64(v7) {
		if base.I32_wrap_i64(v8) == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(33816706))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_int8div_0), int32(0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_int8div_1), int32(521), int32(_a_F_int8div_2))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
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
			if v5 == int64(-9223372036854775807-1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_int8div_3), int32(0))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_int8div_1), int32(537), int32(_a_F_int8div_2))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
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
				return int64(0) - v5
			}
		}
	} else {
		v37 = base.I64_div_s(v5, v6)
		return v37
	}
}
func F_int8range_subdiff(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v4 int64
	_ = v4
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	return base.I64_reinterpret_f64(base.F64_sub(base.F64_convert_i64_s(v2), base.F64_convert_i64_s(v4)))
}
func F_int8shl(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+40)))
	return v2 << (uint(v3) % 64)
}
func F_int8shr(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+40)))
	return v2 >> (uint(v3) % 64)
}
func F_intarray_del_elem(m *base.Module, l0 int32) int64 {
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
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	v2 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum_copy(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
		if v18 != 0 {
			v19 = F_array_contains_nulls(m, v13)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int64(0)
			} else {
				if v19 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v144 = m.ExcPending
					if v144 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(67108994))
						mBase = m.M
						v147 = m.ExcPending
						if v147 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_intarray_del_elem_0), int32(0))
							mBase = m.M
							v151 = m.ExcPending
							if v151 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_intarray_del_elem_1), int32(360), int32(_a_F_intarray_del_elem_2))
								mBase = m.M
								v156 = m.ExcPending
								if v156 != 0 {
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
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v23 = v13 + int32(16)
					v24 = F_ArrayGetNItemsSafe(m, v21, v23)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int64(0)
					} else {
						if v24 != 0 {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
							v27 = F_ArrayGetNItemsSafe(m, v26, v23)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int64(0)
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
									return int64(0)
								} else {
									v138 = v125
									return base.I64_extend_i32_u(v138)
								}
							}
						} else {
							v138 = v13
							return base.I64_extend_i32_u(v138)
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
				return int64(0)
			} else {
				if v24 != 0 {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v27 = F_ArrayGetNItemsSafe(m, v26, v23)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int64(0)
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
							return int64(0)
						} else {
							v138 = v125
							return base.I64_extend_i32_u(v138)
						}
					}
				} else {
					v138 = v13
					return base.I64_extend_i32_u(v138)
				}
			}
		}
	}
}
func F_inter_sl(m *base.Module, l0 int32) int64 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = F_lseg_interpt_line(m, int32(0), v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v5)
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
				F_errfinish(m, int32(_a_F_internalerrquery_1), int32(1700), int32(_a_F_internalerrquery_2))
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
	F_errfinish(m, int32(_a_F_interpret_func_volatility_4), int32(644), int32(_a_F_interpret_func_volatility_5))
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
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
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
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v154 int32
	_ = v154
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v277 int32
	_ = v277
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v311 int32
	_ = v311
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v401 int32
	_ = v401
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	var v434 int32
	_ = v434
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v472 int32
	_ = v472
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v523 int32
	_ = v523
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v549 int32
	_ = v549
	var v557 int32
	_ = v557
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v4
	v9 = F_find_among(m, l0, int32(_a_F_irish_UTF_8_stem_0), int32(24), int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v701
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v4
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v78
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v103 = v4
	goto L38
L3:
	;
	return int32(0)
L4:
	;
	if v9 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v15
	switch v9 - int32(1) {
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
		goto L33
	}
L7:
	;
	v66 = F_slice_from_s(m, l0, int32(1), int32(_a_F_irish_UTF_8_stem_2))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L3
	} else {
		goto L31
	}
L8:
	;
	v60 = F_slice_from_s(m, l0, int32(1), int32(_a_F_irish_UTF_8_stem_3))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L3
	} else {
		goto L29
	}
L9:
	;
	v54 = F_slice_from_s(m, l0, int32(1), int32(_a_F_irish_UTF_8_stem_4))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L3
	} else {
		goto L27
	}
L10:
	;
	v48 = F_slice_from_s(m, l0, int32(1), int32(_a_F_irish_UTF_8_stem_5))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L3
	} else {
		goto L25
	}
L11:
	;
	v42 = F_slice_from_s(m, l0, int32(1), int32(_a_F_irish_UTF_8_stem_6))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L3
	} else {
		goto L23
	}
L12:
	;
	v36 = F_slice_from_s(m, l0, int32(1), int32(_a_F_irish_UTF_8_stem_7))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L21
	}
L13:
	;
	v30 = F_slice_from_s(m, l0, int32(1), int32(_a_F_irish_UTF_8_stem_8))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L3
	} else {
		goto L19
	}
L14:
	;
	v24 = F_slice_from_s(m, l0, int32(1), int32(_a_F_irish_UTF_8_stem_9))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L3
	} else {
		goto L17
	}
L15:
	;
	v19 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v19 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	v701 = v19
	goto L1
L17:
	;
	if int32(0) <= v24 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	v701 = v24
	goto L1
L19:
	;
	if int32(0) <= v30 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v701 = v30
	goto L1
L21:
	;
	if int32(0) <= v36 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v701 = v36
	goto L1
L23:
	;
	if int32(0) <= v42 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	v701 = v42
	goto L1
L25:
	;
	if int32(0) <= v48 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v701 = v48
	goto L1
L27:
	;
	if int32(0) <= v54 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	v701 = v54
	goto L1
L29:
	;
	if int32(0) <= v60 {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	v701 = v60
	goto L1
L31:
	;
	if int32(0) <= v66 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v701 = v66
	goto L1
L33:
	;
	if v72 < int32(0) {
		v701 = v72
		goto L1
	} else {
		goto L34
	}
L34:
	;
	goto L2
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v4
	v576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v576
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v576
	v582 = F_find_among_b(m, l0, int32(_a_F_irish_UTF_8_stem_10), int32(16), int32(0))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L3
	} else {
		goto L139
	}
L36:
	;
	if v198 < int32(0) {
		goto L35
	} else {
		goto L61
	}
L37:
	;
	v198 = v170
	goto L36
L38:
	;
	if v78 <= v103 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v198 = int32(-1)
	goto L36
L41:
	;
	goto L42
L42:
	;
	v110 = int32(1)
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103+v95))))
	if base.Ui32(v112) < base.Ui32(int32(192)) {
		v169 = v112
		v170 = v110
		goto L43
	} else {
		goto L44
	}
L43:
	;
	if int32(250) < v169 {
		goto L56
	} else {
		goto L57
	}
L44:
	;
	v116 = v103 + int32(1)
	if v116 == v78 {
		v169 = v112
		v170 = v110
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116+v95))))
	v121 = v119 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v112) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125+v95))))
	v137 = v135 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v112) {
		goto L52
	} else {
		goto L53
	}
L47:
	;
	v125 = v103 + int32(2)
	if v125 != v78 {
		goto L46
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v169 = v112<<(uint(int32(6))%32)&int32(1984) | v121
	v170 = int32(2)
	goto L43
L50:
	;
	goto L49
L51:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v141))))
	v169 = v154&int32(63) | (v112<<(uint(int32(18))%32)&int32(_a_F_irish_UTF_8_stem_11) | v121<<(uint(int32(12))%32) | v137<<(uint(int32(6))%32))
	v170 = int32(4)
	goto L43
L52:
	;
	v141 = v103 + int32(3)
	if v141 != v78 {
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v169 = v112<<(uint(int32(12))%32)&int32(_a_F_irish_UTF_8_stem_12) | v121<<(uint(int32(6))%32) | v137
	v170 = int32(3)
	goto L43
L55:
	;
	goto L54
L56:
	;
	v187 = v170 + v103
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v187
	v103 = v187
	goto L38
L57:
	;
	v174 = v169 - int32(97)
	if v174 < int32(0) {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v174)>>(uint(int32(3))%32)))+uint32(_c_F_irish_UTF_8_stem[0]))))
	if int32(base.Ui32(v180)>>(uint(v174&int32(7))%32))&int32(1) != 0 {
		goto L37
	} else {
		goto L59
	}
L59:
	;
	goto L56
L61:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v202 = v201 + v198
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v202
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v202
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v226 = v202
	goto L64
L62:
	;
	if v322 < int32(0) {
		goto L35
	} else {
		goto L86
	}
L63:
	;
	v322 = v293
	goto L62
L64:
	;
	if v217 <= v226 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v322 = int32(-1)
	goto L62
L67:
	;
	goto L68
L68:
	;
	v233 = int32(1)
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226+v218))))
	if base.Ui32(v235) < base.Ui32(int32(192)) {
		v292 = v235
		v293 = v233
		goto L69
	} else {
		goto L70
	}
L69:
	;
	if int32(250) < v292 {
		goto L63
	} else {
		goto L82
	}
L70:
	;
	v239 = v226 + int32(1)
	if v239 == v217 {
		v292 = v235
		v293 = v233
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239+v218))))
	v244 = v242 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v235) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248+v218))))
	v260 = v258 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v235) {
		goto L78
	} else {
		goto L79
	}
L73:
	;
	v248 = v226 + int32(2)
	if v248 != v217 {
		goto L72
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v292 = v235<<(uint(int32(6))%32)&int32(1984) | v244
	v293 = int32(2)
	goto L69
L76:
	;
	goto L75
L77:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218+v264))))
	v292 = v277&int32(63) | (v235<<(uint(int32(18))%32)&int32(_a_F_irish_UTF_8_stem_11) | v244<<(uint(int32(12))%32) | v260<<(uint(int32(6))%32))
	v293 = int32(4)
	goto L69
L78:
	;
	v264 = v226 + int32(3)
	if v264 != v217 {
		goto L77
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v292 = v235<<(uint(int32(12))%32)&int32(_a_F_irish_UTF_8_stem_12) | v244<<(uint(int32(6))%32) | v260
	v293 = int32(3)
	goto L69
L81:
	;
	goto L80
L82:
	;
	v297 = v292 - int32(97)
	if v297 < int32(0) {
		goto L63
	} else {
		goto L83
	}
L83:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v297)>>(uint(int32(3))%32)))+uint32(_c_F_irish_UTF_8_stem[0]))))
	if int32(base.Ui32(v303)>>(uint(v297&int32(7))%32))&int32(1) == int32(0) {
		goto L63
	} else {
		goto L84
	}
L84:
	;
	v311 = v293 + v226
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v311
	v226 = v311
	goto L64
L86:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v326 = v325 + v322
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v326
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v326
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v350 = v326
	goto L89
L87:
	;
	if v445 < int32(0) {
		goto L35
	} else {
		goto L112
	}
L88:
	;
	v445 = v417
	goto L87
L89:
	;
	if v341 <= v350 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v445 = int32(-1)
	goto L87
L92:
	;
	goto L93
L93:
	;
	v357 = int32(1)
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350+v342))))
	if base.Ui32(v359) < base.Ui32(int32(192)) {
		v416 = v359
		v417 = v357
		goto L94
	} else {
		goto L95
	}
L94:
	;
	if int32(250) < v416 {
		goto L107
	} else {
		goto L108
	}
L95:
	;
	v363 = v350 + int32(1)
	if v363 == v341 {
		v416 = v359
		v417 = v357
		goto L94
	} else {
		goto L96
	}
L96:
	;
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+v342))))
	v368 = v366 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v359) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372+v342))))
	v384 = v382 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v359) {
		goto L103
	} else {
		goto L104
	}
L98:
	;
	v372 = v350 + int32(2)
	if v372 != v341 {
		goto L97
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v416 = v359<<(uint(int32(6))%32)&int32(1984) | v368
	v417 = int32(2)
	goto L94
L101:
	;
	goto L100
L102:
	;
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342+v388))))
	v416 = v401&int32(63) | (v359<<(uint(int32(18))%32)&int32(_a_F_irish_UTF_8_stem_11) | v368<<(uint(int32(12))%32) | v384<<(uint(int32(6))%32))
	v417 = int32(4)
	goto L94
L103:
	;
	v388 = v350 + int32(3)
	if v388 != v341 {
		goto L102
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v416 = v359<<(uint(int32(12))%32)&int32(_a_F_irish_UTF_8_stem_12) | v368<<(uint(int32(6))%32) | v384
	v417 = int32(3)
	goto L94
L106:
	;
	goto L105
L107:
	;
	v434 = v417 + v350
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v434
	v350 = v434
	goto L89
L108:
	;
	v421 = v416 - int32(97)
	if v421 < int32(0) {
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v421)>>(uint(int32(3))%32)))+uint32(_c_F_irish_UTF_8_stem[0]))))
	if int32(base.Ui32(v427)>>(uint(v421&int32(7))%32))&int32(1) != 0 {
		goto L88
	} else {
		goto L110
	}
L110:
	;
	goto L107
L112:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v449 = v448 + v445
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v449
	v463 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v472 = v449
	goto L115
L113:
	;
	if v568 < int32(0) {
		goto L35
	} else {
		goto L137
	}
L114:
	;
	v568 = v539
	goto L113
L115:
	;
	if v463 <= v472 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v568 = int32(-1)
	goto L113
L118:
	;
	goto L119
L119:
	;
	v479 = int32(1)
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472+v464))))
	if base.Ui32(v481) < base.Ui32(int32(192)) {
		v538 = v481
		v539 = v479
		goto L120
	} else {
		goto L121
	}
L120:
	;
	if int32(250) < v538 {
		goto L114
	} else {
		goto L133
	}
L121:
	;
	v485 = v472 + int32(1)
	if v485 == v463 {
		v538 = v481
		v539 = v479
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485+v464))))
	v490 = v488 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v481) {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494+v464))))
	v506 = v504 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v481) {
		goto L129
	} else {
		goto L130
	}
L124:
	;
	v494 = v472 + int32(2)
	if v494 != v463 {
		goto L123
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v538 = v481<<(uint(int32(6))%32)&int32(1984) | v490
	v539 = int32(2)
	goto L120
L127:
	;
	goto L126
L128:
	;
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464+v510))))
	v538 = v523&int32(63) | (v481<<(uint(int32(18))%32)&int32(_a_F_irish_UTF_8_stem_11) | v490<<(uint(int32(12))%32) | v506<<(uint(int32(6))%32))
	v539 = int32(4)
	goto L120
L129:
	;
	v510 = v472 + int32(3)
	if v510 != v463 {
		goto L128
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v538 = v481<<(uint(int32(12))%32)&int32(_a_F_irish_UTF_8_stem_12) | v490<<(uint(int32(6))%32) | v506
	v539 = int32(3)
	goto L120
L132:
	;
	goto L131
L133:
	;
	v543 = v538 - int32(97)
	if v543 < int32(0) {
		goto L114
	} else {
		goto L134
	}
L134:
	;
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v543)>>(uint(int32(3))%32)))+uint32(_c_F_irish_UTF_8_stem[0]))))
	if int32(base.Ui32(v549)>>(uint(v543&int32(7))%32))&int32(1) == int32(0) {
		goto L114
	} else {
		goto L135
	}
L135:
	;
	v557 = v539 + v472
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v557
	v472 = v557
	goto L115
L137:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v571 + v568
	goto L35
L138:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v601
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v601
	v607 = F_find_among_b(m, l0, int32(_a_F_irish_UTF_8_stem_13), int32(25), int32(0))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L3
	} else {
		goto L148
	}
L139:
	;
	if v582 == int32(0) {
		goto L138
	} else {
		goto L140
	}
L140:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v586
	switch v582 - int32(1) {
	case 0:
		goto L142
	case 1:
		goto L141
	default:
		goto L138
	}
L141:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v586 < v595 {
		goto L138
	} else {
		goto L145
	}
L142:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v586 < v590 {
		goto L138
	} else {
		goto L143
	}
L143:
	;
	v592 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v592 {
		goto L138
	} else {
		goto L144
	}
L144:
	;
	v701 = v592
	goto L1
L145:
	;
	v597 = F_slice_del(m, l0)
	mBase = m.M
	if v597 < int32(0) {
		v701 = v597
		goto L1
	} else {
		goto L146
	}
L146:
	;
	goto L138
L147:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v652
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v652
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v652-int32(2) <= v655 {
		goto L168
	} else {
		goto L169
	}
L148:
	;
	if v607 == int32(0) {
		goto L147
	} else {
		goto L149
	}
L149:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v611
	switch v607 - int32(1) {
	case 0:
		goto L155
	case 1:
		goto L154
	case 2:
		goto L153
	case 3:
		goto L152
	case 4:
		goto L151
	case 5:
		goto L150
	default:
		goto L147
	}
L150:
	;
	v646 = F_slice_from_s(m, l0, int32(4), int32(_a_F_irish_UTF_8_stem_14))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L3
	} else {
		goto L166
	}
L151:
	;
	v640 = F_slice_from_s(m, l0, int32(5), int32(_a_F_irish_UTF_8_stem_15))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L3
	} else {
		goto L164
	}
L152:
	;
	v634 = F_slice_from_s(m, l0, int32(4), int32(_a_F_irish_UTF_8_stem_16))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L3
	} else {
		goto L162
	}
L153:
	;
	v628 = F_slice_from_s(m, l0, int32(3), int32(_a_F_irish_UTF_8_stem_17))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L3
	} else {
		goto L160
	}
L154:
	;
	v622 = F_slice_from_s(m, l0, int32(3), int32(_a_F_irish_UTF_8_stem_18))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L3
	} else {
		goto L158
	}
L155:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v611 < v615 {
		goto L147
	} else {
		goto L156
	}
L156:
	;
	v617 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v617 {
		goto L147
	} else {
		goto L157
	}
L157:
	;
	v701 = v617
	goto L1
L158:
	;
	if int32(0) <= v622 {
		goto L147
	} else {
		goto L159
	}
L159:
	;
	v701 = v622
	goto L1
L160:
	;
	if int32(0) <= v628 {
		goto L147
	} else {
		goto L161
	}
L161:
	;
	v701 = v628
	goto L1
L162:
	;
	if int32(0) <= v634 {
		goto L147
	} else {
		goto L163
	}
L163:
	;
	v701 = v634
	goto L1
L164:
	;
	if int32(0) <= v640 {
		goto L147
	} else {
		goto L165
	}
L165:
	;
	v701 = v640
	goto L1
L166:
	;
	if v646 < int32(0) {
		v701 = v646
		goto L1
	} else {
		goto L167
	}
L167:
	;
	goto L147
L168:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v698
	v701 = int32(1)
	goto L1
L169:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v661 = int32(1)
	v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v659+v652-v661))))
	if base.B2i32(v663&int32(224) != int32(96))|base.B2i32(v661<<(uint(v663)%32)&int32(_a_F_irish_UTF_8_stem_19) == int32(0)) != 0 {
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v678 = F_find_among_b(m, l0, int32(_a_F_irish_UTF_8_stem_20), int32(12), int32(0))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L3
	} else {
		goto L171
	}
L171:
	;
	if v678 == int32(0) {
		goto L168
	} else {
		goto L172
	}
L172:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v682
	switch v678 - int32(1) {
	case 0:
		goto L174
	case 1:
		goto L173
	default:
		goto L168
	}
L173:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v682 < v691 {
		goto L168
	} else {
		goto L177
	}
L174:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v682 < v686 {
		goto L168
	} else {
		goto L175
	}
L175:
	;
	v688 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v688 {
		goto L168
	} else {
		goto L176
	}
L176:
	;
	v701 = v688
	goto L1
L177:
	;
	v693 = F_slice_del(m, l0)
	mBase = m.M
	if v693 < int32(0) {
		v701 = v693
		goto L1
	} else {
		goto L178
	}
L178:
	;
	goto L168
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
func F_is_strict_saop(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
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
				v50 = int32(0)
				return v50
			} else {
				if l1 != 0 {
					v12 = int32(1)
					v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
					if v13&v12 != 0 {
						v50 = v12
						return v50
					} else {
						v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
						v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
						if v19 == int32(0) {
							v50 = int32(0)
							return v50
						} else {
							v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
							if v22 != int32(35) {
								if v22 != int32(7) {
									v50 = int32(0)
									return v50
								} else {
									v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+32)))
									if v27 != 0 {
										v50 = int32(0)
										return v50
									} else {
										v29 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
										v30 = F_pg_detoast_datum(m, v29)
										mBase = m.M
										v31 = m.ExcPending
										if v31 != 0 {
											return int32(0)
										} else {
											v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
											v35 = F_ArrayGetNItemsSafe(m, v32, v30+int32(16))
											mBase = m.M
											v36 = m.ExcPending
											if v36 != 0 {
												return int32(0)
											} else {
												if v35 <= int32(0) {
													v50 = int32(0)
												} else {
													v50 = int32(1)
												}
												return v50
											}
										}
									}
								}
							} else {
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
								if v39 == int32(0) {
									v50 = int32(0)
								} else {
									v42 = int32(1)
									v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+20)))
									if v43 != v42 {
										v50 = v42
									} else {
										v50 = int32(0)
									}
								}
								return v50
							}
						}
					}
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
					if v19 == int32(0) {
						v50 = int32(0)
						return v50
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
						if v22 != int32(35) {
							if v22 != int32(7) {
								v50 = int32(0)
								return v50
							} else {
								v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+32)))
								if v27 != 0 {
									v50 = int32(0)
									return v50
								} else {
									v29 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
									v30 = F_pg_detoast_datum(m, v29)
									mBase = m.M
									v31 = m.ExcPending
									if v31 != 0 {
										return int32(0)
									} else {
										v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
										v35 = F_ArrayGetNItemsSafe(m, v32, v30+int32(16))
										mBase = m.M
										v36 = m.ExcPending
										if v36 != 0 {
											return int32(0)
										} else {
											if v35 <= int32(0) {
												v50 = int32(0)
											} else {
												v50 = int32(1)
											}
											return v50
										}
									}
								}
							}
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
							if v39 == int32(0) {
								v50 = int32(0)
							} else {
								v42 = int32(1)
								v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+20)))
								if v43 != v42 {
									v50 = v42
								} else {
									v50 = int32(0)
								}
							}
							return v50
						}
					}
				}
			}
		}
	}
}
func F_isbn_cast_from_ean13(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14315(m, l0, int32(3))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_isbn_in(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14257(m, l0, int32(3))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
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
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_ivfflatbeginscan[0]))
	v16 = F_RelationGetIndexScan(m, l0, l1, l2)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_IvfflatGetMetaPageInfo(m, l0, v12+int32(12), v12+int32(8))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_ivfflatbeginscan[1]))
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_ivfflatbeginscan[2]))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v32 = F_palloc(m, int32(88))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v34 = F_IvfflatGetTypeInfo(m, l0)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v36 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+16)) = uint8(v36)
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v34
	if v15 < v30 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v40 = v15
	goto L8
L7:
	;
	v40 = v30
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v40
	if v15 < v29 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v43 = v29
	goto L11
L10:
	;
	v43 = v15
	goto L11
L11:
	;
	if v27 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v44 = v43
	goto L14
L13:
	;
	v44 = v15
	goto L14
L14:
	;
	if v44 < v30 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v46 = v44
	goto L17
L16:
	;
	v46 = v30
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v46
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v48
	v52 = int32(1)
	v54 = F_index_getprocinfo(m, l0, v52, v52)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+56)) = v54
	v58 = F_HnswOptionalProcInfo(m, l0, int32(2))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+60)) = v58
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = v62
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_ivfflatbeginscan[3]))
	v70 = F_AllocSetContextCreateInternal(m, v65, int32(_a_F_ivfflatbeginscan_0), int32(0), int32(_a_F_ivfflatbeginscan_1), int32(_a_F_ivfflatbeginscan_2))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = v70
	v73 = int32(_a_F_ivfflatbeginscan_3)
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_ivfflatbeginscan[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ivfflatbeginscan[3])) = v70
	v78 = F_CreateTemplateTupleDesc(m, int32(2))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+40)) = v78
	F_TupleDescInitEntry(m, v78, int32(1), int32(_a_F_ivfflatbeginscan_4), int32(701), int32(-1), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v32)+40))
	F_TupleDescInitEntry(m, v88, int32(2), int32(_a_F_ivfflatbeginscan_5), int32(27), int32(-1), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v32)+40))
	v97 = int32(0)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	if v97 < v106 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v32)+40))
	v185 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+30)) = uint16(v185)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(672)
	v189 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v189
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+19)) = uint8(v189)
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_ivfflatbeginscan[4]))
	v206 = F_tuplesort_begin_heap(m, v184, v185, v12+int32(30), v12+int32(24), v12+int32(20), v12+int32(19), v203, v189, v189)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L43
	}
L25:
	;
	v110 = v96 + int32(28)
	v117 = v97
	v118 = v106
	v120 = v97
	goto L29
L26:
	;
	v174 = v97
	v181 = v106
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+20)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v96)+16)) = v174
	goto L24
L28:
	;
	v174 = v168
	v181 = v147
	goto L27
L29:
	;
	v126 = v110 + v106<<(uint(int32(3))%32) + v117*int32(100)
	v129 = v110 + v117<<(uint(int32(3))%32)
	if v106 != v118 {
		v147 = v118
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v168 = v106
	goto L28
L31:
	;
	v148 = int32(*(*int16)(unsafe.Add(mBase, uint32(v129)+2)))
	if v148 <= int32(0) {
		v168 = v117
		goto L28
	} else {
		goto L39
	}
L32:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+7)))
	if v131 != int32(118) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v147 = v117
	goto L31
L34:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+4)))
	if v134 != int32(1) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+6)))
	if v137&int32(6) != 0 {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v140 = int32(*(*int16)(unsafe.Add(mBase, uint32(v129)+2)))
	if v140 <= int32(0) {
		goto L33
	} else {
		goto L37
	}
L37:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+90)))
	if v143 != int32(118) {
		v147 = v106
		goto L31
	} else {
		goto L38
	}
L38:
	;
	goto L33
L39:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+90)))
	if v151 == int32(118) {
		v168 = v117
		goto L28
	} else {
		goto L40
	}
L40:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+5)))
	v160 = (v120 + v154 - int32(1)) & (int32(0) - v154)
	if int32(_a_F_ivfflatbeginscan_6) < v160 {
		v168 = v117
		goto L28
	} else {
		goto L41
	}
L41:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v129))) = uint16(v160)
	v166 = v117 + int32(1)
	if v166 != v106 {
		v117 = v166
		v118 = v147
		v120 = v160 + v148
		goto L29
	} else {
		goto L42
	}
L42:
	;
	goto L30
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+36)) = v206
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v32)+40))
	v211 = F_MakeSingleTupleTableSlot(m, v209, int32(_a_F_ivfflatbeginscan_7))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+44)) = v211
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v32)+40))
	v216 = F_MakeSingleTupleTableSlot(m, v214, int32(_a_F_ivfflatbeginscan_8))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v216
	v220 = F_GetAccessStrategy(m, int32(1))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+52)) = v220
	v224 = F_pairingheap_allocate(m, int32(_a_F_ivfflatbeginscan_9), v16)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+72)) = v224
	v228 = F_palloc_mul(m, int32(4), v46)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+80)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+76)) = v228
	v234 = F_palloc_mul(m, int32(24), v46)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v234
	*(*int32)(unsafe.Add(mBase, _c_F_ivfflatbeginscan[3])) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v32
	m.G0 = v12 + int32(32)
	return v16
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
func F_ivfflathandler(m *base.Module, l0 int32) int64 {
	return int64(4171880)
}
func F_ivfflatoptions(m *base.Module, l0 int64, l1 int32) int32 {
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
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+72))
	v10 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(v6)+80)) = v10
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
	if v14 == v10 {
		if l1 == int32(0) {
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v27 <= int32(0) {
			} else {
				v31 = v27 * int32(56)
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
				v43 = v39 * int32(56)
				if v43 == int32(0) {
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					base.MemoryCopy(m, v46, l3, v43)
				}
			}
		}
		return
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
		if v17 == int32(0) {
			if l1 == int32(0) {
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v27 <= int32(0) {
				} else {
					v31 = v27 * int32(56)
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
					v43 = v39 * int32(56)
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
				*(*int64)(unsafe.Add(mBase, uint32(v6)+24)) = int64(0)
				if l1 == int32(0) {
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v27 <= int32(0) {
					} else {
						v31 = v27 * int32(56)
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
						v43 = v39 * int32(56)
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
