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
	var v10 int32
	_ = v10
	v2 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[1230])) = v2
	*(*int32)(unsafe.Add(mBase, _consts[1])) = v2
	v8 = *(*int32)(unsafe.Add(mBase, _consts[99]))
	F_SetLatch(m, v8)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		return
	}
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
	var v13 int32
	_ = v13
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
	var v43 int32
	_ = v43
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2 < v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = v2
	v12 = v2
	v13 = v7
	goto L4
L2:
	;
	v43 = v2
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v43
	return
L4:
	;
	v17 = v12 << (uint(int32(2)) % 32)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17+v18)))
	if v20 == int32(0) {
		v37 = v11
		v38 = v13
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v43 = v37
	goto L3
L6:
	;
	v40 = v12 + int32(1)
	if v40 < v38 {
		v11 = v37
		v12 = v40
		v13 = v38
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
		v38 = v13
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
		v38 = v13
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	v34 = F_lappend_oid(m, v11, v33)
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
	v37 = v34
	v38 = v36
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
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = int32(0)
	v9 = F_pgl_shmctl(m, l1, v7, v7)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		if int32(0) <= v9 {
			m.G0 = v5 + int32(16)
			return
		} else {
			v15 = F_errstart(m, int32(15), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				if v15 == int32(0) {
					m.G0 = v5 + int32(16)
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v5))) = l1
					F_errmsg_internal(m, int32(309986), v5)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						F_errfinish(m, int32(522134), int32(302), int32(369107))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
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
}
func F_IpcMemoryDetach(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v9 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	if v9 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v5 + int32(16)
	return
L2:
	;
	if int32(0) <= v26 {
		goto L1
	} else {
		goto L11
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[158])) = int32(28)
	v26 = int32(-1)
	goto L2
L4:
	;
	v13 = v9
	goto L5
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if l1 != v14 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v26 = int32(0)
	goto L2
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	if v16 != 0 {
		v13 = v16
		goto L5
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	goto L6
L10:
	;
	goto L3
L11:
	;
	v31 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	if v31 == int32(0) {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = l1
	F_errmsg_internal(m, int32(309941), v5)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(522134), int32(290), int32(342639))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L1
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
				F_errmsg(m, int32(421859), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(520045), int32(197), int32(241876))
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
				F_errmsg(m, int32(421806), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(525865), int32(1360), int32(454545))
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
		v12 = F_ArrayGetNItems(m, v9, v5+int32(16))
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 float64
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_get_negator(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(220627), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(516413), int32(773), int32(321747))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v33 = F_patternsel_common(m, v9, v11, int32(0), v7, v6, v8, int32(3), int32(1))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v35 = F_Float8GetDatum(m, v33)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					return v35
				}
			}
		}
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
								F_errcontext_msg(m, int32(753667), v6)
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
					F_errcontext_msg(m, int32(753667), v6)
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
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
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
	v85 = l3*int32(28) + v19
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
				v138 = m.ExcPending
				if v138 != 0 {
					return int32(0)
				} else {
					v139 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
					*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v97
					*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v139
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = l3
					F_errmsg_internal(m, int32(42479), v11)
					mBase = m.M
					v146 = m.ExcPending
					if v146 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(520903), int32(648), int32(253782))
						mBase = m.M
						v151 = m.ExcPending
						if v151 != 0 {
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
	var v27 float64
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 float64
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v93 int32
	_ = v93
	var v120 int32
	_ = v120
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
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v153 int32
	_ = v153
	var v176 int32
	_ = v176
	var v187 int32
	_ = v187
	var v209 int32
	_ = v209
	var v220 int32
	_ = v220
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 float64
	_ = v318
	var v323 float64
	_ = v323
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v356 float64
	_ = v356
	var v358 float64
	_ = v358
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v375 int32
	_ = v375
	var v388 int32
	_ = v388
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 float64
	_ = v456
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 float32
	_ = v464
	var v466 float32
	_ = v466
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v493 float64
	_ = v493
	var v494 float64
	_ = v494
	var v498 int32
	_ = v498
	var v499 float64
	_ = v499
	var v502 float64
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v509 float64
	_ = v509
	var v512 int32
	_ = v512
	var v519 float64
	_ = v519
	var v522 float64
	_ = v522
	var v523 int32
	_ = v523
	var v528 float64
	_ = v528
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v543 int32
	_ = v543
	var v544 float64
	_ = v544
	var v545 float64
	_ = v545
	var v550 float64
	_ = v550
	var v552 float64
	_ = v552
	var v553 float64
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v585 int32
	_ = v585
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v625 int32
	_ = v625
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v656 int32
	_ = v656
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v687 int32
	_ = v687
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v748 int32
	_ = v748
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v826 int32
	_ = v826
	var v845 float64
	_ = v845
	var v849 float64
	_ = v849
	var v851 int32
	_ = v851
	var v854 float64
	_ = v854
	var v855 int32
	_ = v855
	var v893 float64
	_ = v893
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v906 int32
	_ = v906
	var v910 int32
	_ = v910
	var v929 float64
	_ = v929
	var v932 float64
	_ = v932
	var v935 int32
	_ = v935
	var v938 float64
	_ = v938
	var v939 int32
	_ = v939
	var v976 float64
	_ = v976
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v1016 float64
	_ = v1016
	var v1019 float64
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1025 float64
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1113 float64
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1118 float64
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1123 float64
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1180 int32
	_ = v1180
	var v1182 int32
	_ = v1182
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1188 int64
	_ = v1188
	var v1199 float64
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1204 int32
	_ = v1204
	var v1208 int64
	_ = v1208
	var v1212 int64
	_ = v1212
	var v1214 int32
	_ = v1214
	var v1218 int64
	_ = v1218
	var v1223 int64
	_ = v1223
	var v1225 float64
	_ = v1225
	var v1233 int64
	_ = v1233
	var v1237 int64
	_ = v1237
	var v1249 float64
	_ = v1249
	var v1261 float64
	_ = v1261
	var v1263 int32
	_ = v1263
	var v1265 float64
	_ = v1265
	var v1267 int32
	_ = v1267
	var v1269 float64
	_ = v1269
	var v1271 int64
	_ = v1271
	var v1277 int32
	_ = v1277
	var v1281 int32
	_ = v1281
	var v1285 int64
	_ = v1285
	var v1290 int64
	_ = v1290
	var v1294 int64
	_ = v1294
	var v1299 int32
	_ = v1299
	var v1301 float64
	_ = v1301
	var v1303 int64
	_ = v1303
	var v1308 int32
	_ = v1308
	var v1312 int64
	_ = v1312
	var v1323 int64
	_ = v1323
	var v1327 int64
	_ = v1327
	var v1331 int32
	_ = v1331
	var v1338 int32
	_ = v1338
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1370 int32
	_ = v1370
	var v1395 int32
	_ = v1395
	var v1397 int32
	_ = v1397
	var v1399 int32
	_ = v1399
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1404 int32
	_ = v1404
	var v1411 int32
	_ = v1411
	var v1437 int32
	_ = v1437
	var v1440 int32
	_ = v1440
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1482 int32
	_ = v1482
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1521 int32
	_ = v1521
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1526 int32
	_ = v1526
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1537 int32
	_ = v1537
	var v1563 int32
	_ = v1563
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1569 int32
	_ = v1569
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1580 float64
	_ = v1580
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1607 float64
	_ = v1607
	var v1611 float64
	_ = v1611
	var v1613 int32
	_ = v1613
	var v1615 int32
	_ = v1615
	var v1617 int32
	_ = v1617
	var v1621 float64
	_ = v1621
	var v1623 int32
	_ = v1623
	var v1659 float64
	_ = v1659
	var v1662 int32
	_ = v1662
	var v1665 int32
	_ = v1665
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1672 int32
	_ = v1672
	var v1674 float64
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1682 int32
	_ = v1682
	var v1701 float64
	_ = v1701
	var v1704 float64
	_ = v1704
	var v1707 int32
	_ = v1707
	var v1709 int32
	_ = v1709
	var v1711 int32
	_ = v1711
	var v1715 float64
	_ = v1715
	var v1717 int32
	_ = v1717
	var v1752 float64
	_ = v1752
	var v1756 int32
	_ = v1756
	var v1760 int32
	_ = v1760
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1768 int32
	_ = v1768
	var v1770 float64
	_ = v1770
	var v1772 int32
	_ = v1772
	var v1778 int32
	_ = v1778
	var v1797 float64
	_ = v1797
	var v1800 float64
	_ = v1800
	var v1803 int32
	_ = v1803
	var v1805 int32
	_ = v1805
	var v1807 int32
	_ = v1807
	var v1811 float64
	_ = v1811
	var v1813 int32
	_ = v1813
	var v1848 float64
	_ = v1848
	var v1853 int32
	_ = v1853
	var v1855 int32
	_ = v1855
	var v1857 int32
	_ = v1857
	var v1898 float64
	_ = v1898
	var v1899 int32
	_ = v1899
	var v1903 float64
	_ = v1903
	var v1904 int32
	_ = v1904
	var v1908 float64
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1911 int32
	_ = v1911
	var v1915 int64
	_ = v1915
	var v1930 int32
	_ = v1930
	var v1961 float64
	_ = v1961
	var v1962 float64
	_ = v1962
	var v1964 float64
	_ = v1964
	var v1969 float64
	_ = v1969
	var v1974 float64
	_ = v1974
	var v1980 float64
	_ = v1980
	var v1983 float64
	_ = v1983
	var v1986 float64
	_ = v1986
	var v1988 float64
	_ = v1988
	var v1990 float64
	_ = v1990
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v1997 float64
	_ = v1997
	var v2004 float64
	_ = v2004
	var v2006 float64
	_ = v2006
	var v2008 float64
	_ = v2008
	var v2013 float64
	_ = v2013
	var v2016 float64
	_ = v2016
	var v2048 float64
	_ = v2048
	var v2053 float64
	_ = v2053
	var v2088 float64
	_ = v2088
	var v2094 int32
	_ = v2094
	var v2098 float64
	_ = v2098
	var v2101 float64
	_ = v2101
	var v2133 float64
	_ = v2133
	var v2140 int32
	_ = v2140
	var v2169 float64
	_ = v2169
	v10 = int32(0)
	v27 = float64(0)
	v33 = m.G0
	v35 = v33 - int32(112)
	m.G0 = v35
	v37 = float64(-1)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v38 == v10 {
		v2169 = v37
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v35 + int32(112)
	return v2169
L2:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+29)))
	if v43 != 0 {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	F_free_attstatsslot(m, v35+int32(76))
	mBase = m.M
	v2140 = m.ExcPending
	if v2140 != 0 {
		goto L13
	} else {
		goto L470
	}
L4:
	;
	v2094 = *(*int32)(unsafe.Add(mBase, uint32(v35)+92))
	v2098 = base.F64_div(float64(0.01), base.F64_convert_i32_s(v2094-int32(1)))
	if base.F64_lt(v2088, v2098) != 0 {
		v2133 = v2098
		goto L3
	} else {
		goto L468
	}
L5:
	;
	v367 = int32(0)
	v375 = v209
	v388 = v363
	goto L57
L6:
	;
	v358 = base.F64_sub(float64(1), v323)
	if base.F64_gt(v318, v358) == int32(0) {
		v2133 = v318
		goto L3
	} else {
		goto L56
	}
L7:
	;
	if v209 == int32(2) {
		goto L48
	} else {
		goto L49
	}
L8:
	;
	v329 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L13
	} else {
		goto L43
	}
L9:
	;
	v55 = v38
	goto L11
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v44 == int32(0) {
		v2169 = v37
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v59 = F_get_attstatsslot(m, v35+int32(76), v55, int32(2), int32(0), int32(1))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L13
	} else {
		goto L16
	}
L12:
	;
	v47 = F_get_func_leakproof(m, v44)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return float64(0)
L14:
	;
	if v47 == int32(0) {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v55 = v53
	goto L11
L16:
	;
	if v59 == int32(0) {
		v2169 = v37
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v35)+92))
	if v63 < int32(2) {
		v220 = v63
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if v220 < int32(2) {
		v2133 = v37
		goto L3
	} else {
		goto L37
	}
L19:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v35)+80))
	if v66 != l6 {
		v220 = v63
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v35)+76))
	if l2 != v69 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v71 = int32(0)
	v76 = F_SearchSysCacheList(m, int32(3), int32(1), v69, v71, v71)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L13
	} else {
		goto L25
	}
L22:
	;
	v187 = int32(1)
	goto L23
L23:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v35)+92))
	if v187 != 0 {
		goto L7
	} else {
		goto L36
	}
L24:
	;
	F_ReleaseCatCacheList(m, v76)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L13
	} else {
		goto L35
	}
L25:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v76)+40))
	if v78 <= int32(0) {
		v153 = v71
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v93 = v71
	goto L27
L27:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v76+int32(48)+v93<<(uint(int32(2))%32))))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+56))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+22)))
	v123 = v121 + v122
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	v126 = F_SearchSysCacheExists(m, int32(3), l2, int32(115), v124, int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L13
	} else {
		goto L30
	}
L28:
	;
	v153 = int32(0)
	goto L24
L29:
	;
	v139 = v93 + int32(1)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v76)+40))
	if v139 < v140 {
		v93 = v139
		goto L27
	} else {
		goto L34
	}
L30:
	;
	if v126 == int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v123)+24))
	v132 = F_GetIndexAmRoutineByAmId(m, v130, int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L13
	} else {
		goto L32
	}
L32:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+14)))
	if v134 == int32(0) {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	v153 = int32(1)
	goto L24
L34:
	;
	goto L28
L35:
	;
	v187 = v153
	goto L23
L36:
	;
	v220 = v209
	goto L18
L37:
	;
	v244 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+48)) = uint8(v244)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+40)) = uint8(v244)
	v248 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+34)) = uint16(v248)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+32)) = uint8(v244)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = l6
	*(*int64)(unsafe.Add(mBase, uint32(v35)+20)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v35)+44)) = l7
	v269 = v244
	v270 = v244
	goto L38
L38:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v35)+88))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v291+v270<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+36)) = v295
	v297 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+32)) = uint8(v297)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v301)))
	v303 = m.T0[v302].(func(*base.Module, int32) int32)(m, v35+int32(16))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L13
	} else {
		goto L40
	}
L39:
	;
	v318 = base.F64_div(base.F64_convert_i32_s(v311), base.F64_convert_i32_s(v314))
	v323 = base.F64_div(float64(0.01), base.F64_convert_i32_s(v314-int32(1)))
	if base.F64_lt(v318, v323) == int32(0) {
		goto L6
	} else {
		goto L42
	}
L40:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+32)))
	v311 = v269 + (v305^int32(-1))&base.B2i32(v303 != int32(0))
	v313 = v270 + int32(1)
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v35)+92))
	if v313 < v314 {
		v269 = v311
		v270 = v313
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v2133 = v323
	goto L3
L43:
	;
	if v329 == int32(0) {
		v2169 = v37
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v333 = F_get_func_name(m, v44)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v333
	F_errmsg_internal(m, int32(355508), v35)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L13
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(519114), int32(6242), int32(333732))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	v2169 = v37
	goto L1
L48:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v35)+76))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v35)+88))
	v350 = F_get_actual_variable_range(m, l0, l1, v346, l6, v347, v347+int32(4))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L13
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	if int32(0) < v209 {
		v363 = v10
		goto L5
	} else {
		goto L52
	}
L51:
	;
	v363 = v350
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
	v356 = float64(1)
	goto L55
L54:
	;
	v356 = float64(0)
	goto L55
L55:
	;
	v2088 = v356
	goto L4
L56:
	;
	v2133 = v358
	goto L3
L57:
	;
	v397 = v367 + v375
	v398 = int32(2)
	v399 = base.I32_div_s(v397, v398)
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v35)+92))
	if base.Ui32(v398) < base.Ui32(v397+int32(1)) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	if v438 <= int32(0) {
		goto L77
	} else {
		goto L78
	}
L59:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v35)+88))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v428+v399<<(uint(int32(2))%32))))
	v433 = F_FunctionCall2Coll(m, l3, l6, v432, l7)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L13
	} else {
		goto L67
	}
L60:
	;
	if v400 < int32(3) {
		v425 = v388
		goto L59
	} else {
		goto L64
	}
L61:
	;
	if v400 < int32(3) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v35)+76))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v35)+88))
	v410 = F_get_actual_variable_range(m, l0, l1, v407, l6, v408, int32(0))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L13
	} else {
		goto L63
	}
L63:
	;
	v425 = v410
	goto L59
L64:
	;
	if v399 != v400-int32(1) {
		v425 = v388
		goto L59
	} else {
		goto L65
	}
L65:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v35)+76))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v35)+88))
	v423 = F_get_actual_variable_range(m, l0, l1, v417, l6, int32(0), v419+v399<<(uint(int32(2))%32))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L13
	} else {
		goto L66
	}
L66:
	;
	v425 = v423
	goto L59
L67:
	;
	v437 = l4 ^ base.B2i32(v433 != int32(0))
	if v437 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v438 = v399 + int32(1)
	goto L70
L69:
	;
	v438 = v367
	goto L70
L70:
	;
	if v437 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v439 = v375
	goto L73
L72:
	;
	v439 = v399
	goto L73
L73:
	;
	if v438 < v439 {
		v367 = v438
		v375 = v439
		v388 = v425
		goto L57
	} else {
		goto L74
	}
L74:
	;
	goto L58
L75:
	;
	v2053 = float64(0)
	if base.F64_lt(v2048, v2053) != 0 {
		v2133 = v2053
		goto L3
	} else {
		goto L466
	}
L76:
	;
	if l4 != 0 {
		goto L462
	} else {
		goto L463
	}
L77:
	;
	v2013 = float64(0)
	goto L76
L78:
	;
	goto L79
L79:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v35)+92))
	if v444 <= v438 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v2013 = float64(1)
	goto L76
L81:
	;
	goto L82
L82:
	;
	v447 = l4 ^ l5
	if v447 == int32(1) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v553 = float64(0.5)
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v35)+88))
	v556 = v438 - int32(1)
	v557 = int32(2)
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v554+v556<<(uint(v557)%32))))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v554+v438<<(uint(v557)%32))))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v567 = v35 + int32(56)
	v569 = v35 - int32(-64)
	v570 = m.G0
	v571 = int32(16)
	v572 = v570 - v571
	m.G0 = v572
	v574 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v572)+15)) = uint8(v574)
	v577 = v35 + v571
	if l8 <= int32(1081) {
		goto L144
	} else {
		goto L145
	}
L84:
	;
	if v438 != int32(1) {
		v552 = float64(0)
		goto L83
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v454 = v35 - int32(-64)
	v455 = int32(0)
	v456 = float64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v454))) = uint8(v455)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v460 != 0 {
		goto L90
	} else {
		goto L91
	}
L87:
	;
	goto L86
L88:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v535 = F_get_attstatsslot(m, v35+int32(16), v531, int32(1), int32(0), int32(2))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L13
	} else {
		goto L121
	}
L89:
	;
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+28)))
	if v498 != 0 {
		goto L103
	} else {
		goto L104
	}
L90:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v460)+16))
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461)+22)))
	v463 = v461 + v462
	v464 = *(*float32)(unsafe.Add(mBase, uint32(v463)+8))
	v466 = *(*float32)(unsafe.Add(mBase, uint32(v463)+16))
	v493 = base.F64_promote_f32(v466)
	v494 = base.F64_promote_f32(v464)
	goto L89
L91:
	;
	goto L92
L92:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v468 == int32(16) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v493 = float64(2)
	v494 = v456
	goto L89
L94:
	;
	goto L95
L95:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v472 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v479 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v472)+76))
	if v475 != int32(5) {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v493 = float64(-1)
	v494 = v456
	goto L89
L99:
	;
	v493 = float64(0)
	v494 = v456
	goto L89
L100:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v479)))
	if v482 != int32(6) {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v486 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v479)+8)))
	switch v486 - int32(65530) {
	case 0:
		goto L102
	default:
		goto L99
	case 5:
		v493 = float64(-1)
		v494 = v456
		goto L89
	}
L102:
	;
	v493 = float64(1)
	v494 = v456
	goto L89
L103:
	;
	v499 = base.F64_neg(base.F64_sub(float64(1), v494))
	goto L105
L104:
	;
	v499 = v493
	goto L105
L105:
	;
	if base.F64_gt(v499, float64(0)) != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v502 = F_clamp_row_est(m, v499)
	mBase = m.M
	v528 = v502
	goto L88
L107:
	;
	goto L108
L108:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v503 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v506 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v454))) = uint8(v506)
	v528 = float64(200)
	goto L88
L110:
	;
	goto L111
L111:
	;
	v509 = *(*float64)(unsafe.Add(mBase, uint32(v503)+120))
	if base.F64_le(v509, float64(0)) != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v512 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v454))) = uint8(v512)
	v528 = float64(200)
	goto L88
L113:
	;
	goto L114
L114:
	;
	if base.F64_lt(v499, float64(0)) != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v519 = F_clamp_row_est(m, base.F64_mul(v509, base.F64_neg(v499)))
	mBase = m.M
	v528 = v519
	goto L88
L116:
	;
	goto L117
L117:
	;
	if base.F64_lt(v509, float64(200)) != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v522 = F_clamp_row_est(m, v509)
	mBase = m.M
	v528 = v522
	goto L88
L119:
	;
	goto L120
L120:
	;
	v523 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v454))) = uint8(v523)
	v528 = float64(200)
	goto L88
L121:
	;
	if v535 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v35)+40))
	F_free_attstatsslot(m, v35+int32(16))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L13
	} else {
		goto L125
	}
L123:
	;
	v544 = v528
	goto L124
L124:
	;
	v545 = float64(1)
	if base.F64_gt(v544, v545) != 0 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v544 = base.F64_sub(v528, base.F64_convert_i32_s(v537))
	goto L124
L126:
	;
	v550 = base.F64_div(v545, v544)
	goto L128
L127:
	;
	v550 = float64(0)
	goto L128
L128:
	;
	v552 = v550
	goto L83
L129:
	;
	m.G0 = v572 + int32(16)
	if v1930&int32(1) == int32(0) {
		v1988 = v553
		goto L434
	} else {
		goto L435
	}
L130:
	;
	v1915 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v569))) = v1915
	*(*int64)(unsafe.Add(mBase, uint32(v567))) = v1915
	*(*int64)(unsafe.Add(mBase, uint32(v577))) = v1915
	v1930 = int32(0)
	goto L129
L131:
	;
	v1898 = F_convert_network_to_scalar(m, l7, l8, v572+int32(15))
	mBase = m.M
	v1899 = m.ExcPending
	if v1899 != 0 {
		goto L13
	} else {
		goto L431
	}
L132:
	;
	if l8 != int32(650) {
		goto L130
	} else {
		goto L430
	}
L133:
	;
	v1930 = v1142 ^ int32(1)
	goto L129
L134:
	;
	if v1143 != 0 {
		goto L320
	} else {
		goto L321
	}
L135:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v577))) = v1225
	if v565 <= int32(1183) {
		goto L300
	} else {
		goto L301
	}
L136:
	;
	if l8 != int32(1114) {
		goto L130
	} else {
		goto L292
	}
L137:
	;
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(l7)+8))
	v1218 = *(*int64)(unsafe.Add(mBase, uint32(l7)))
	v1225 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1214), float64(1e+06)), base.F64_convert_i64_s(v1218))
	goto L135
L138:
	;
	v1212 = *(*int64)(unsafe.Add(mBase, uint32(l7)))
	v1225 = base.F64_convert_i64_s(v1212)
	goto L135
L139:
	;
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(l7)+8))
	v1208 = *(*int64)(unsafe.Add(mBase, uint32(l7)))
	v1225 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1200), float64(2.6298e+12)), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1204), float64(8.64e+10)), base.F64_convert_i64_s(v1208)))
	goto L135
L140:
	;
	if l7 == int32(-2147483648) {
		goto L286
	} else {
		goto L287
	}
L141:
	;
	v1188 = *(*int64)(unsafe.Add(mBase, uint32(l7)))
	v1225 = base.F64_convert_i64_s(v1188)
	goto L135
L142:
	;
	v1132 = F_convert_string_datum(m, l7, l8, l6, v572+int32(15))
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L13
	} else {
		goto L269
	}
L143:
	;
	v1113 = F_convert_numeric_to_scalar(m, l7, l8, v572+int32(15))
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L13
	} else {
		goto L266
	}
L144:
	;
	if l8 <= int32(699) {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	goto L146
L146:
	;
	if l8 <= int32(2201) {
		goto L249
	} else {
		goto L250
	}
L147:
	;
	if base.Ui32(int32(26)) < base.Ui32(l8) {
		goto L132
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	if l8 <= int32(828) {
		goto L241
	} else {
		goto L242
	}
L150:
	;
	v585 = int32(1) << (uint(l8) % 32)
	if v585&int32(95485952) != 0 {
		goto L143
	} else {
		goto L151
	}
L151:
	;
	if v585&int32(34340864) != 0 {
		goto L142
	} else {
		goto L152
	}
L152:
	;
	if l8 != int32(17) {
		goto L132
	} else {
		goto L153
	}
L153:
	;
	if v565 != int32(17) {
		v1930 = int32(0)
		goto L129
	} else {
		goto L154
	}
L154:
	;
	v595 = F_pg_detoast_datum_packed(m, l7)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L13
	} else {
		goto L155
	}
L155:
	;
	v597 = F_pg_detoast_datum_packed(m, v560)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L13
	} else {
		goto L156
	}
L156:
	;
	v599 = F_pg_detoast_datum_packed(m, v564)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L13
	} else {
		goto L157
	}
L157:
	;
	v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v595))))
	if v601 == int32(1) {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v597))))
	if v632 == int32(1) {
		goto L170
	} else {
		goto L171
	}
L159:
	;
	v604 = int32(4)
	v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v595)+1)))
	if v606&int32(254) == int32(2) {
		goto L162
	} else {
		goto L163
	}
L160:
	;
	goto L161
L161:
	;
	v619 = int32(1)
	if v601&v619 != 0 {
		v631 = int32(base.Ui32(v601)>>(uint(v619)%32)) - v619
		goto L158
	} else {
		goto L168
	}
L162:
	;
	v615 = v604
	goto L164
L163:
	;
	v615 = base.B2i32(v606 == int32(18)) << (uint(v604) % 32)
	goto L164
L164:
	;
	if v606 == int32(1) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v618 = v604
	goto L167
L166:
	;
	v618 = v615
	goto L167
L167:
	;
	v631 = v618
	goto L158
L168:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v595)))
	v631 = int32(base.Ui32(v625)>>(uint(int32(2))%32)) - int32(4)
	goto L158
L169:
	;
	v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599))))
	if v663 == int32(1) {
		goto L181
	} else {
		goto L182
	}
L170:
	;
	v635 = int32(4)
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v597)+1)))
	if v637&int32(254) == int32(2) {
		goto L173
	} else {
		goto L174
	}
L171:
	;
	goto L172
L172:
	;
	v650 = int32(1)
	if v632&v650 != 0 {
		v662 = int32(base.Ui32(v632)>>(uint(v650)%32)) - v650
		goto L169
	} else {
		goto L179
	}
L173:
	;
	v646 = v635
	goto L175
L174:
	;
	v646 = base.B2i32(v637 == int32(18)) << (uint(v635) % 32)
	goto L175
L175:
	;
	if v637 == int32(1) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v649 = v635
	goto L178
L177:
	;
	v649 = v646
	goto L178
L178:
	;
	v662 = v649
	goto L169
L179:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v597)))
	v662 = int32(base.Ui32(v656)>>(uint(int32(2))%32)) - int32(4)
	goto L169
L180:
	;
	v694 = int32(1)
	if v663&v694 != 0 {
		goto L191
	} else {
		goto L192
	}
L181:
	;
	v666 = int32(4)
	v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599)+1)))
	if v668&int32(254) == int32(2) {
		goto L184
	} else {
		goto L185
	}
L182:
	;
	goto L183
L183:
	;
	v681 = int32(1)
	if v663&v681 != 0 {
		v693 = int32(base.Ui32(v663)>>(uint(v681)%32)) - v681
		goto L180
	} else {
		goto L190
	}
L184:
	;
	v677 = v666
	goto L186
L185:
	;
	v677 = base.B2i32(v668 == int32(18)) << (uint(v666) % 32)
	goto L186
L186:
	;
	if v668 == int32(1) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v680 = v666
	goto L189
L188:
	;
	v680 = v677
	goto L189
L189:
	;
	v693 = v680
	goto L180
L190:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v599)))
	v693 = int32(base.Ui32(v687)>>(uint(int32(2))%32)) - int32(4)
	goto L180
L191:
	;
	v698 = v694
	goto L193
L192:
	;
	v698 = int32(4)
	goto L193
L193:
	;
	v699 = v599 + v698
	v700 = int32(1)
	if v632&v700 != 0 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v704 = v700
	goto L196
L195:
	;
	v704 = int32(4)
	goto L196
L196:
	;
	v705 = v597 + v704
	v706 = int32(1)
	if v601&v706 != 0 {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v710 = v706
	goto L199
L198:
	;
	v710 = int32(4)
	goto L199
L199:
	;
	v711 = v595 + v710
	if v631 < v662 {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	if int32(0) < v781 {
		goto L213
	} else {
		goto L214
	}
L201:
	;
	v713 = v631
	goto L203
L202:
	;
	v713 = v662
	goto L203
L203:
	;
	if v713 < v693 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v715 = v713
	goto L206
L205:
	;
	v715 = v693
	goto L206
L206:
	;
	if v715 <= int32(0) {
		v780 = v699
		v781 = v631
		v783 = v705
		v785 = v693
		v786 = v662
		v787 = v711
		goto L200
	} else {
		goto L207
	}
L207:
	;
	v728 = v699
	v729 = v631
	v731 = v705
	v733 = v693
	v734 = v662
	v735 = v711
	v748 = int32(0)
	goto L208
L208:
	;
	v760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731))))
	v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v728))))
	if v760 != v761 {
		v780 = v728
		v781 = v729
		v783 = v731
		v785 = v733
		v786 = v734
		v787 = v735
		goto L200
	} else {
		goto L210
	}
L209:
	;
	v780 = v599 + v698 + v715
	v781 = v631 - v715
	v783 = v704 + v597 + v715
	v785 = v693 - v715
	v786 = v662 - v715
	v787 = v595 + v710 + v715
	goto L200
L210:
	;
	v763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v735))))
	if v760 != v763 {
		v780 = v728
		v781 = v729
		v783 = v731
		v785 = v733
		v786 = v734
		v787 = v735
		goto L200
	} else {
		goto L211
	}
L211:
	;
	v765 = int32(1)
	v778 = v748 + v765
	if v778 != v715 {
		v728 = v728 + v765
		v729 = v729 - v765
		v731 = v731 + v765
		v733 = v733 - v765
		v734 = v734 - v765
		v735 = v735 + v765
		v748 = v778
		goto L208
	} else {
		goto L212
	}
L212:
	;
	goto L209
L213:
	;
	v814 = int32(10)
	if base.Ui32(v814) <= base.Ui32(v781) {
		goto L216
	} else {
		goto L217
	}
L214:
	;
	v893 = v27
	goto L215
L215:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v577))) = v893
	if int32(0) < v786 {
		goto L222
	} else {
		goto L223
	}
L216:
	;
	v817 = v814
	goto L218
L217:
	;
	v817 = v781
	goto L218
L218:
	;
	v820 = v817
	v826 = v787
	v845 = float64(256)
	v849 = v27
	goto L219
L219:
	;
	v851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v826))))
	v854 = base.F64_add(v849, base.F64_div(base.F64_convert_i32_u(v851), v845))
	v855 = int32(1)
	if base.Ui32(v855) < base.Ui32(v820) {
		v820 = v820 - v855
		v826 = v826 + v855
		v845 = base.F64_mul(v845, float64(256))
		v849 = v854
		goto L219
	} else {
		goto L221
	}
L220:
	;
	v893 = v854
	goto L215
L221:
	;
	goto L220
L222:
	;
	v898 = int32(10)
	if base.Ui32(v898) <= base.Ui32(v786) {
		goto L225
	} else {
		goto L226
	}
L223:
	;
	v976 = v27
	goto L224
L224:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v567))) = v976
	if v785 <= int32(0) {
		goto L232
	} else {
		goto L233
	}
L225:
	;
	v901 = v898
	goto L227
L226:
	;
	v901 = v786
	goto L227
L227:
	;
	v906 = v783
	v910 = v901
	v929 = float64(256)
	v932 = v27
	goto L228
L228:
	;
	v935 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v906))))
	v938 = base.F64_add(v932, base.F64_div(base.F64_convert_i32_u(v935), v929))
	v939 = int32(1)
	if base.Ui32(v939) < base.Ui32(v910) {
		v906 = v906 + v939
		v910 = v910 - v939
		v929 = base.F64_mul(v929, float64(256))
		v932 = v938
		goto L228
	} else {
		goto L230
	}
L229:
	;
	v976 = v938
	goto L224
L230:
	;
	goto L229
L231:
	;
	v1930 = int32(1)
	goto L129
L232:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v569))) = float64(0)
	goto L231
L233:
	;
	goto L234
L234:
	;
	v984 = int32(10)
	if base.Ui32(v984) <= base.Ui32(v785) {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v987 = v984
	goto L237
L236:
	;
	v987 = v785
	goto L237
L237:
	;
	v990 = v780
	v993 = v987
	v1016 = float64(256)
	v1019 = float64(0)
	goto L238
L238:
	;
	v1022 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v990))))
	v1025 = base.F64_add(v1019, base.F64_div(base.F64_convert_i32_u(v1022), v1016))
	v1026 = int32(1)
	if base.Ui32(v1026) < base.Ui32(v993) {
		v990 = v990 + v1026
		v993 = v993 - v1026
		v1016 = base.F64_mul(v1016, float64(256))
		v1019 = v1025
		goto L238
	} else {
		goto L240
	}
L239:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v569))) = v1025
	goto L231
L240:
	;
	goto L239
L241:
	;
	if base.Ui32(l8-int32(700)) < base.Ui32(int32(2)) {
		goto L143
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	if base.Ui32(l8-int32(1042)) < base.Ui32(int32(2)) {
		goto L142
	} else {
		goto L246
	}
L244:
	;
	if l8 == int32(774) {
		goto L131
	} else {
		goto L245
	}
L245:
	;
	goto L130
L246:
	;
	if l8 == int32(829) {
		goto L131
	} else {
		goto L247
	}
L247:
	;
	if l8 != int32(869) {
		goto L130
	} else {
		goto L248
	}
L248:
	;
	goto L131
L249:
	;
	if l8 <= int32(1183) {
		goto L252
	} else {
		goto L253
	}
L250:
	;
	goto L251
L251:
	;
	if l8 <= int32(3768) {
		goto L258
	} else {
		goto L259
	}
L252:
	;
	switch l8 - int32(1082) {
	case 0:
		goto L140
	case 1:
		goto L138
	default:
		goto L136
	}
L253:
	;
	goto L254
L254:
	;
	switch l8 - int32(1184) {
	case 0:
		goto L141
	case 1:
		goto L130
	case 2:
		goto L139
	default:
		goto L255
	}
L255:
	;
	if l8 == int32(1266) {
		goto L137
	} else {
		goto L256
	}
L256:
	;
	if l8 == int32(1700) {
		goto L143
	} else {
		goto L257
	}
L257:
	;
	goto L130
L258:
	;
	if base.Ui32(l8-int32(2202)) < base.Ui32(int32(5)) {
		goto L143
	} else {
		goto L261
	}
L259:
	;
	goto L260
L260:
	;
	switch l8 - int32(4089) {
	case 0, 7:
		goto L143
	case 1, 2, 3, 4, 5, 6:
		goto L130
	default:
		goto L263
	}
L261:
	;
	if l8 == int32(3734) {
		goto L143
	} else {
		goto L262
	}
L262:
	;
	goto L130
L263:
	;
	if l8 == int32(4191) {
		goto L143
	} else {
		goto L264
	}
L264:
	;
	if l8 != int32(3769) {
		goto L130
	} else {
		goto L265
	}
L265:
	;
	goto L143
L266:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v577))) = v1113
	v1118 = F_convert_numeric_to_scalar(m, v560, v565, v572+int32(15))
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		goto L13
	} else {
		goto L267
	}
L267:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v567))) = v1118
	v1123 = F_convert_numeric_to_scalar(m, v564, v565, v572+int32(15))
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L13
	} else {
		goto L268
	}
L268:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v569))) = v1123
	v1126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572)+15)))
	v1930 = v1126 ^ int32(1)
	goto L129
L269:
	;
	v1136 = F_convert_string_datum(m, v560, v565, l6, v572+int32(15))
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L13
	} else {
		goto L270
	}
L270:
	;
	v1140 = F_convert_string_datum(m, v564, v565, l6, v572+int32(15))
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L13
	} else {
		goto L271
	}
L271:
	;
	v1142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572)+15)))
	if v1142 != 0 {
		goto L133
	} else {
		goto L272
	}
L272:
	;
	v1143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1140))))
	v1144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1136))))
	if v1144 == int32(0) {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v1331 = v1143
	v1338 = v1143
	goto L134
L274:
	;
	goto L275
L275:
	;
	v1148 = v1143
	v1150 = v1136
	v1154 = v1144
	v1155 = v1143
	goto L276
L276:
	;
	v1180 = v1154 & int32(255)
	if base.Ui32(v1180) < base.Ui32(v1148) {
		goto L278
	} else {
		goto L279
	}
L277:
	;
	v1331 = v1182
	v1338 = v1184
	goto L134
L278:
	;
	v1182 = v1148
	goto L280
L279:
	;
	v1182 = v1180
	goto L280
L280:
	;
	if base.Ui32(v1155) < base.Ui32(v1180) {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1184 = v1155
	goto L283
L282:
	;
	v1184 = v1180
	goto L283
L283:
	;
	v1186 = v1150 + int32(1)
	v1187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1186))))
	if v1187 != 0 {
		v1148 = v1182
		v1150 = v1186
		v1154 = v1187
		v1155 = v1184
		goto L276
	} else {
		goto L284
	}
L284:
	;
	goto L277
L285:
	;
	v1225 = v1199
	goto L135
L286:
	;
	v1199 = float64(-1.7976931348623157e+308)
	goto L285
L287:
	;
	goto L288
L288:
	;
	if l7 == int32(2147483647) {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	v1199 = float64(1.7976931348623157e+308)
	goto L285
L290:
	;
	goto L291
L291:
	;
	v1199 = base.F64_mul(base.F64_convert_i32_s(l7), float64(8.64e+10))
	goto L285
L292:
	;
	v1223 = *(*int64)(unsafe.Add(mBase, uint32(l7)))
	v1225 = base.F64_convert_i64_s(v1223)
	goto L135
L293:
	;
	v1323 = *(*int64)(unsafe.Add(mBase, uint32(v560)))
	*(*float64)(unsafe.Add(mBase, uint32(v567))) = base.F64_convert_i64_s(v1323)
	v1327 = *(*int64)(unsafe.Add(mBase, uint32(v564)))
	*(*float64)(unsafe.Add(mBase, uint32(v569))) = base.F64_convert_i64_s(v1327)
	v1930 = int32(1)
	goto L129
L294:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v567))) = int64(0)
	*(*float64)(unsafe.Add(mBase, uint32(v569))) = float64(0)
	v1930 = int32(0)
	goto L129
L295:
	;
	if v565 == int32(1114) {
		goto L293
	} else {
		goto L319
	}
L296:
	;
	if v565 != int32(1266) {
		goto L294
	} else {
		goto L318
	}
L297:
	;
	v1290 = *(*int64)(unsafe.Add(mBase, uint32(v560)))
	*(*float64)(unsafe.Add(mBase, uint32(v567))) = base.F64_convert_i64_s(v1290)
	v1294 = *(*int64)(unsafe.Add(mBase, uint32(v564)))
	*(*float64)(unsafe.Add(mBase, uint32(v569))) = base.F64_convert_i64_s(v1294)
	v1930 = int32(1)
	goto L129
L298:
	;
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(v560)+12))
	v1265 = float64(2.6298e+12)
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v560)+8))
	v1269 = float64(8.64e+10)
	v1271 = *(*int64)(unsafe.Add(mBase, uint32(v560)))
	*(*float64)(unsafe.Add(mBase, uint32(v567))) = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1263), v1265), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1267), v1269), base.F64_convert_i64_s(v1271)))
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v564)+12))
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v564)+8))
	v1285 = *(*int64)(unsafe.Add(mBase, uint32(v564)))
	*(*float64)(unsafe.Add(mBase, uint32(v569))) = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1277), v1265), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1281), v1269), base.F64_convert_i64_s(v1285)))
	v1930 = int32(1)
	goto L129
L299:
	;
	if v560 == int32(-2147483648) {
		goto L305
	} else {
		goto L306
	}
L300:
	;
	switch v565 - int32(1082) {
	case 0:
		goto L299
	case 1:
		goto L297
	default:
		goto L295
	}
L301:
	;
	goto L302
L302:
	;
	switch v565 - int32(1184) {
	case 0:
		goto L303
	case 1:
		goto L294
	case 2:
		goto L298
	default:
		goto L296
	}
L303:
	;
	v1233 = *(*int64)(unsafe.Add(mBase, uint32(v560)))
	*(*float64)(unsafe.Add(mBase, uint32(v567))) = base.F64_convert_i64_s(v1233)
	v1237 = *(*int64)(unsafe.Add(mBase, uint32(v564)))
	*(*float64)(unsafe.Add(mBase, uint32(v569))) = base.F64_convert_i64_s(v1237)
	v1930 = int32(1)
	goto L129
L304:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v567))) = v1249
	if v564 == int32(-2147483648) {
		goto L312
	} else {
		goto L313
	}
L305:
	;
	v1249 = float64(-1.7976931348623157e+308)
	goto L304
L306:
	;
	goto L307
L307:
	;
	if v560 == int32(2147483647) {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	v1249 = float64(1.7976931348623157e+308)
	goto L304
L309:
	;
	goto L310
L310:
	;
	v1249 = base.F64_mul(base.F64_convert_i32_s(v560), float64(8.64e+10))
	goto L304
L311:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v569))) = v1261
	v1930 = int32(1)
	goto L129
L312:
	;
	v1261 = float64(-1.7976931348623157e+308)
	goto L311
L313:
	;
	goto L314
L314:
	;
	if v564 == int32(2147483647) {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	v1261 = float64(1.7976931348623157e+308)
	goto L311
L316:
	;
	goto L317
L317:
	;
	v1261 = base.F64_mul(base.F64_convert_i32_s(v564), float64(8.64e+10))
	goto L311
L318:
	;
	v1299 = *(*int32)(unsafe.Add(mBase, uint32(v560)+8))
	v1301 = float64(1e+06)
	v1303 = *(*int64)(unsafe.Add(mBase, uint32(v560)))
	*(*float64)(unsafe.Add(mBase, uint32(v567))) = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1299), v1301), base.F64_convert_i64_s(v1303))
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v564)+8))
	v1312 = *(*int64)(unsafe.Add(mBase, uint32(v564)))
	*(*float64)(unsafe.Add(mBase, uint32(v569))) = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1308), v1301), base.F64_convert_i64_s(v1312))
	v1930 = int32(1)
	goto L129
L319:
	;
	goto L294
L320:
	;
	v1362 = v1143
	v1363 = v1331
	v1365 = v1140
	v1370 = v1338
	goto L323
L321:
	;
	v1404 = v1331
	v1411 = v1338
	goto L322
L322:
	;
	v1437 = int32(90)
	if v1404 <= v1437 {
		goto L332
	} else {
		goto L333
	}
L323:
	;
	v1395 = v1362 & int32(255)
	if v1395 < v1363 {
		goto L325
	} else {
		goto L326
	}
L324:
	;
	v1404 = v1397
	v1411 = v1399
	goto L322
L325:
	;
	v1397 = v1363
	goto L327
L326:
	;
	v1397 = v1395
	goto L327
L327:
	;
	if v1370 < v1395 {
		goto L328
	} else {
		goto L329
	}
L328:
	;
	v1399 = v1370
	goto L330
L329:
	;
	v1399 = v1395
	goto L330
L330:
	;
	v1401 = v1365 + int32(1)
	v1402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1401))))
	if v1402 != 0 {
		v1362 = v1402
		v1363 = v1397
		v1365 = v1401
		v1370 = v1399
		goto L323
	} else {
		goto L331
	}
L331:
	;
	goto L324
L332:
	;
	v1440 = v1437
	goto L334
L333:
	;
	v1440 = v1404
	goto L334
L334:
	;
	v1445 = base.B2i32(v1411 < int32(91)) & base.B2i32(int32(64) < v1404)
	if v1445 != 0 {
		goto L335
	} else {
		goto L336
	}
L335:
	;
	v1446 = v1440
	goto L337
L336:
	;
	v1446 = v1404
	goto L337
L337:
	;
	if v1446 <= int32(122) {
		goto L338
	} else {
		goto L339
	}
L338:
	;
	v1449 = int32(122)
	goto L340
L339:
	;
	v1449 = v1446
	goto L340
L340:
	;
	v1450 = int32(65)
	if v1450 <= v1411 {
		goto L341
	} else {
		goto L342
	}
L341:
	;
	v1453 = v1450
	goto L343
L342:
	;
	v1453 = v1411
	goto L343
L343:
	;
	if v1445 != 0 {
		goto L344
	} else {
		goto L345
	}
L344:
	;
	v1454 = v1453
	goto L346
L345:
	;
	v1454 = v1411
	goto L346
L346:
	;
	v1459 = base.B2i32(v1454 < int32(123)) & base.B2i32(int32(96) < v1446)
	if v1459 != 0 {
		goto L347
	} else {
		goto L348
	}
L347:
	;
	v1460 = v1449
	goto L349
L348:
	;
	v1460 = v1446
	goto L349
L349:
	;
	if v1460 <= int32(57) {
		goto L350
	} else {
		goto L351
	}
L350:
	;
	v1463 = int32(57)
	goto L352
L351:
	;
	v1463 = v1460
	goto L352
L352:
	;
	v1464 = int32(97)
	if v1464 <= v1454 {
		goto L353
	} else {
		goto L354
	}
L353:
	;
	v1467 = v1464
	goto L355
L354:
	;
	v1467 = v1454
	goto L355
L355:
	;
	if v1459 != 0 {
		goto L356
	} else {
		goto L357
	}
L356:
	;
	v1468 = v1467
	goto L358
L357:
	;
	v1468 = v1454
	goto L358
L358:
	;
	v1473 = base.B2i32(v1468 < int32(58)) & base.B2i32(int32(47) < v1460)
	if v1473 != 0 {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	v1474 = v1463
	goto L361
L360:
	;
	v1474 = v1460
	goto L361
L361:
	;
	v1475 = int32(48)
	if v1475 <= v1468 {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	v1478 = v1475
	goto L364
L363:
	;
	v1478 = v1468
	goto L364
L364:
	;
	if v1473 != 0 {
		goto L365
	} else {
		goto L366
	}
L365:
	;
	v1479 = v1478
	goto L367
L366:
	;
	v1479 = v1468
	goto L367
L367:
	;
	v1482 = base.B2i32(v1474-v1479 < int32(9))
	if v1144 == int32(0) {
		v1530 = v1136
		v1531 = v1140
		v1537 = v1132
		goto L368
	} else {
		goto L369
	}
L368:
	;
	if v1474-v1479 < int32(9) {
		goto L375
	} else {
		goto L376
	}
L369:
	;
	v1485 = v1136
	v1486 = v1140
	v1491 = v1144
	v1492 = v1132
	goto L370
L370:
	;
	v1518 = v1491 & int32(255)
	v1519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1486))))
	if v1518 != v1519 {
		v1530 = v1485
		v1531 = v1486
		v1537 = v1492
		goto L368
	} else {
		goto L372
	}
L371:
	;
	v1530 = v1528
	v1531 = v1526
	v1537 = v1524
	goto L368
L372:
	;
	v1521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1492))))
	if v1518 != v1521 {
		v1530 = v1485
		v1531 = v1486
		v1537 = v1492
		goto L368
	} else {
		goto L373
	}
L373:
	;
	v1523 = int32(1)
	v1524 = v1492 + v1523
	v1526 = v1486 + v1523
	v1528 = v1485 + v1523
	v1529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1528))))
	if v1529 != 0 {
		v1485 = v1528
		v1486 = v1526
		v1491 = v1529
		v1492 = v1524
		goto L370
	} else {
		goto L374
	}
L374:
	;
	goto L371
L375:
	;
	v1563 = int32(127)
	goto L377
L376:
	;
	v1563 = v1474
	goto L377
L377:
	;
	if v1474-v1479 < int32(9) {
		goto L378
	} else {
		goto L379
	}
L378:
	;
	v1565 = int32(32)
	goto L380
L379:
	;
	v1565 = v1479
	goto L380
L380:
	;
	v1566 = F_strlen(m, v1537)
	mBase = m.M
	if int32(0) < v1566 {
		goto L381
	} else {
		goto L382
	}
L381:
	;
	v1569 = int32(12)
	if base.Ui32(v1569) <= base.Ui32(v1566) {
		goto L384
	} else {
		goto L385
	}
L382:
	;
	v1659 = v27
	goto L383
L383:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v577))) = v1659
	v1662 = F_strlen(m, v1530)
	mBase = m.M
	if int32(0) < v1662 {
		goto L396
	} else {
		goto L397
	}
L384:
	;
	v1572 = v1569
	goto L386
L385:
	;
	v1572 = v1566
	goto L386
L386:
	;
	v1573 = int32(1)
	v1580 = base.F64_convert_i32_s(v1563 - v1565 + v1573)
	v1587 = v1572
	v1588 = v1537
	v1607 = v1580
	v1611 = v27
	goto L387
L387:
	;
	v1613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1588))))
	if base.Ui32(v1563) < base.Ui32(v1613) {
		goto L389
	} else {
		goto L390
	}
L388:
	;
	v1659 = v1621
	goto L383
L389:
	;
	v1615 = v1563 + v1573
	goto L391
L390:
	;
	v1615 = v1613
	goto L391
L391:
	;
	if base.Ui32(v1613) < base.Ui32(v1565) {
		goto L392
	} else {
		goto L393
	}
L392:
	;
	v1617 = v1565 - v1573
	goto L394
L393:
	;
	v1617 = v1615
	goto L394
L394:
	;
	v1621 = base.F64_add(v1611, base.F64_div(base.F64_convert_i32_s(v1617-v1565), v1607))
	v1623 = int32(1)
	if base.Ui32(v1623) < base.Ui32(v1587) {
		v1587 = v1587 - v1623
		v1588 = v1588 + v1623
		v1607 = base.F64_mul(v1607, v1580)
		v1611 = v1621
		goto L387
	} else {
		goto L395
	}
L395:
	;
	goto L388
L396:
	;
	v1665 = int32(12)
	if base.Ui32(v1665) <= base.Ui32(v1662) {
		goto L399
	} else {
		goto L400
	}
L397:
	;
	v1752 = v27
	goto L398
L398:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v567))) = v1752
	v1756 = F_strlen(m, v1531)
	mBase = m.M
	if v1756 <= int32(0) {
		goto L412
	} else {
		goto L413
	}
L399:
	;
	v1668 = v1665
	goto L401
L400:
	;
	v1668 = v1662
	goto L401
L401:
	;
	v1669 = int32(1)
	v1672 = v1563 + v1669
	v1674 = base.F64_convert_i32_s(v1672 - v1565)
	v1675 = v1530
	v1682 = v1668
	v1701 = v1674
	v1704 = v27
	goto L402
L402:
	;
	v1707 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1675))))
	if base.Ui32(v1563) < base.Ui32(v1707) {
		goto L404
	} else {
		goto L405
	}
L403:
	;
	v1752 = v1715
	goto L398
L404:
	;
	v1709 = v1672
	goto L406
L405:
	;
	v1709 = v1707
	goto L406
L406:
	;
	if base.Ui32(v1707) < base.Ui32(v1565) {
		goto L407
	} else {
		goto L408
	}
L407:
	;
	v1711 = v1565 - v1669
	goto L409
L408:
	;
	v1711 = v1709
	goto L409
L409:
	;
	v1715 = base.F64_add(v1704, base.F64_div(base.F64_convert_i32_s(v1711-v1565), v1701))
	v1717 = int32(1)
	if base.Ui32(v1717) < base.Ui32(v1682) {
		v1675 = v1675 + v1717
		v1682 = v1682 - v1717
		v1701 = base.F64_mul(v1701, v1674)
		v1704 = v1715
		goto L402
	} else {
		goto L410
	}
L410:
	;
	goto L403
L411:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v569))) = v1848
	F_pfree(m, v1132)
	mBase = m.M
	v1853 = m.ExcPending
	if v1853 != 0 {
		goto L13
	} else {
		goto L427
	}
L412:
	;
	v1848 = float64(0)
	goto L411
L413:
	;
	goto L414
L414:
	;
	v1760 = int32(12)
	if base.Ui32(v1760) <= base.Ui32(v1756) {
		goto L415
	} else {
		goto L416
	}
L415:
	;
	v1763 = v1760
	goto L417
L416:
	;
	v1763 = v1756
	goto L417
L417:
	;
	v1764 = int32(1)
	v1768 = v1563 + v1764
	v1770 = base.F64_convert_i32_s(v1768 - v1565)
	v1772 = v1531
	v1778 = v1763
	v1797 = v1770
	v1800 = float64(0)
	goto L418
L418:
	;
	v1803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1772))))
	if base.Ui32(v1563) < base.Ui32(v1803) {
		goto L420
	} else {
		goto L421
	}
L419:
	;
	v1848 = v1811
	goto L411
L420:
	;
	v1805 = v1768
	goto L422
L421:
	;
	v1805 = v1803
	goto L422
L422:
	;
	if base.Ui32(v1803) < base.Ui32(v1565) {
		goto L423
	} else {
		goto L424
	}
L423:
	;
	v1807 = v1565 - v1764
	goto L425
L424:
	;
	v1807 = v1805
	goto L425
L425:
	;
	v1811 = base.F64_add(v1800, base.F64_div(base.F64_convert_i32_s(v1807-v1565), v1797))
	v1813 = int32(1)
	if base.Ui32(v1813) < base.Ui32(v1778) {
		v1772 = v1772 + v1813
		v1778 = v1778 - v1813
		v1797 = base.F64_mul(v1797, v1770)
		v1800 = v1811
		goto L418
	} else {
		goto L426
	}
L426:
	;
	goto L419
L427:
	;
	F_pfree(m, v1136)
	mBase = m.M
	v1855 = m.ExcPending
	if v1855 != 0 {
		goto L13
	} else {
		goto L428
	}
L428:
	;
	F_pfree(m, v1140)
	mBase = m.M
	v1857 = m.ExcPending
	if v1857 != 0 {
		goto L13
	} else {
		goto L429
	}
L429:
	;
	goto L133
L430:
	;
	goto L131
L431:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v577))) = v1898
	v1903 = F_convert_network_to_scalar(m, v560, v565, v572+int32(15))
	mBase = m.M
	v1904 = m.ExcPending
	if v1904 != 0 {
		goto L13
	} else {
		goto L432
	}
L432:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v567))) = v1903
	v1908 = F_convert_network_to_scalar(m, v564, v565, v572+int32(15))
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		goto L13
	} else {
		goto L433
	}
L433:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v569))) = v1908
	v1911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572)+15)))
	v1930 = v1911 ^ int32(1)
	goto L129
L434:
	;
	v1990 = float64(1)
	v1993 = *(*int32)(unsafe.Add(mBase, uint32(v35)+92))
	v1994 = int32(1)
	v1997 = base.F64_div(base.F64_add(v1988, base.F64_convert_i32_u(v556)), base.F64_convert_i32_s(v1993-v1994))
	if v438 != v1994 {
		goto L452
	} else {
		goto L453
	}
L435:
	;
	v1961 = *(*float64)(unsafe.Add(mBase, uint32(v35)+64))
	v1962 = *(*float64)(unsafe.Add(mBase, uint32(v35)+56))
	if base.F64_le(v1961, v1962) != 0 {
		v1988 = v553
		goto L434
	} else {
		goto L436
	}
L436:
	;
	v1964 = *(*float64)(unsafe.Add(mBase, uint32(v35)+16))
	if base.F64_ge(v1962, v1964) != 0 {
		goto L437
	} else {
		goto L438
	}
L437:
	;
	v1988 = float64(0)
	goto L434
L438:
	;
	goto L439
L439:
	;
	if base.F64_ge(v1964, v1961) != 0 {
		goto L440
	} else {
		goto L441
	}
L440:
	;
	v1988 = float64(1)
	goto L434
L441:
	;
	goto L442
L442:
	;
	v1969 = float64(0.5)
	v1974 = base.F64_div(base.F64_sub(v1964, v1962), base.F64_sub(v1961, v1962))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1974)&int64(9223372036854775807)) {
		goto L443
	} else {
		goto L444
	}
L443:
	;
	v1980 = v1969
	goto L445
L444:
	;
	v1980 = v1974
	goto L445
L445:
	;
	if base.F64_lt(v1974, float64(0)) != 0 {
		goto L446
	} else {
		goto L447
	}
L446:
	;
	v1983 = v1969
	goto L448
L447:
	;
	v1983 = v1980
	goto L448
L448:
	;
	if base.F64_gt(v1974, float64(1)) != 0 {
		goto L449
	} else {
		goto L450
	}
L449:
	;
	v1986 = v1969
	goto L451
L450:
	;
	v1986 = v1983
	goto L451
L451:
	;
	v1988 = v1986
	goto L434
L452:
	;
	v2004 = v1997
	goto L454
L453:
	;
	v2004 = base.F64_add(v1997, base.F64_mul(v552, base.F64_sub(v1990, v1988)))
	goto L454
L454:
	;
	if v447 != 0 {
		goto L455
	} else {
		goto L456
	}
L455:
	;
	v2006 = v2004
	goto L457
L456:
	;
	v2006 = base.F64_sub(v2004, v552)
	goto L457
L457:
	;
	if l4 != 0 {
		goto L458
	} else {
		goto L459
	}
L458:
	;
	v2008 = base.F64_sub(v1990, v2006)
	goto L460
L459:
	;
	v2008 = v2006
	goto L460
L460:
	;
	if v425&int32(1) == int32(0) {
		v2088 = v2008
		goto L4
	} else {
		goto L461
	}
L461:
	;
	v2048 = v2008
	goto L75
L462:
	;
	v2016 = base.F64_sub(float64(1), v2013)
	goto L464
L463:
	;
	v2016 = v2013
	goto L464
L464:
	;
	if v425&int32(1) == int32(0) {
		v2088 = v2016
		goto L4
	} else {
		goto L465
	}
L465:
	;
	v2048 = v2016
	goto L75
L466:
	;
	if base.F64_gt(v2048, float64(1)) == int32(0) {
		v2133 = v2048
		goto L3
	} else {
		goto L467
	}
L467:
	;
	v2133 = float64(1)
	goto L3
L468:
	;
	v2101 = base.F64_sub(float64(1), v2098)
	if base.F64_lt(v2101, v2088) == int32(0) {
		v2133 = v2088
		goto L3
	} else {
		goto L469
	}
L469:
	;
	v2133 = v2101
	goto L3
L470:
	;
	v2169 = v2133
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
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v453 int32
	_ = v453
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v642 int32
	_ = v642
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
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
	v32 = *(*int32)(unsafe.Add(mBase, _consts[361]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v33*int32(28))+uint32(_consts[1015])))
	goto L7
L5:
	;
	goto L6
L6:
	;
	v222 = int32(*(*int8)(unsafe.Add(mBase, uint32(v20)+1)))
	v224 = v222 & int32(255)
	if v224 == int32(1) {
		goto L46
	} else {
		goto L47
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
	v112 = int32(base.Ui32(v25)>>(uint(int32(12))%32)) + v24
	v114 = int32(1)
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
	v75 = *(*int32)(unsafe.Add(mBase, _consts[361]))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v76*int32(28))+uint32(_consts[1015])))
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
	v124 = v123 + v114
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v124
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v126 == int32(39) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v217 = F_pg_mblen_cstr(m, v112)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L41
	}
L19:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v124))) = uint8(v126)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v214 = v212 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v214
	v216 = v214
	goto L18
L20:
	;
	if v126 == int32(92) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	if v126 != 0 {
		v216 = v124
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v131 = int32(39)
	*(*uint8)(unsafe.Add(mBase, uint32(v124))) = uint8(v131)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v135 = v133 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v135
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v137 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v205 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v203))) = uint8(v205)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v207 + int32(12)
	goto L3
L24:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+2)))
	if v140 != int32(1) {
		v203 = v135
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v143 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v135))) = uint8(v143)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v146 = int32(1)
	v147 = v145 + v146
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v147
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+2)))
	if v149 == v146 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L26
L28:
	;
	v152 = int32(42)
	*(*uint8)(unsafe.Add(mBase, uint32(v147))) = uint8(v152)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v156 = v154 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v156
	v158 = v156
	goto L30
L29:
	;
	v158 = v147
	goto L30
L30:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v159&int32(8) != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v162 = int32(65)
	*(*uint8)(unsafe.Add(mBase, uint32(v158))) = uint8(v162)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v166 = v164 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v166
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	v169 = v166
	v170 = v168
	goto L33
L32:
	;
	v169 = v158
	v170 = v159
	goto L33
L33:
	;
	if v170&int32(4) != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v173 = int32(66)
	*(*uint8)(unsafe.Add(mBase, uint32(v169))) = uint8(v173)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v177 = v175 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v177
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	v180 = v177
	v181 = v179
	goto L36
L35:
	;
	v180 = v169
	v181 = v170
	goto L36
L36:
	;
	if v181&int32(2) != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v184 = int32(67)
	*(*uint8)(unsafe.Add(mBase, uint32(v180))) = uint8(v184)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v188 = v186 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v188
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	v191 = v188
	v192 = v190
	goto L39
L38:
	;
	v191 = v180
	v192 = v181
	goto L39
L39:
	;
	if v192&int32(1) == int32(0) {
		v203 = v191
		goto L23
	} else {
		goto L40
	}
L40:
	;
	v197 = int32(68)
	*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v197)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v201 = v199 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v201
	v203 = v201
	goto L23
L41:
	;
	if v217 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v112 = v112 + v217
	v114 = v217
	goto L16
L43:
	;
	v219 = F__emscripten_memcpy_bulkmem(m, v216, v112, v217)
	mBase = m.M
	goto L45
L44:
	;
	goto L45
L45:
	;
	goto L42
L46:
	;
	if l1 <= int32(4) {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	goto L48
L48:
	;
	v405 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20)+2)))
	v407 = v20 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v407
	v410 = base.B2i32(v222 == int32(4))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v222<<(uint(int32(2))%32))+uint32(_consts[1158])))
	v418 = l2&v410 | base.B2i32(v416 < l1)
	if v418 != 0 {
		goto L78
	} else {
		goto L79
	}
L49:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v295 = v285 - v294
	v297 = v295 + int32(2)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v298 <= v297 {
		goto L61
	} else {
		goto L62
	}
L50:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v285 = v229
	goto L49
L51:
	;
	goto L52
L52:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v232 = v230 - v231
	v234 = v232 + int32(3)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v235 <= v234 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v239 = v231
	v240 = v235
	goto L56
L54:
	;
	v264 = v230
	goto L55
L55:
	;
	v275 = F_pg_sprintf(m, v264, int32(780668), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L60
	}
L56:
	;
	v251 = v240 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v251
	v253 = F_repalloc(m, v239, v251)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L58
	}
L57:
	;
	v264 = v256
	goto L55
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v253
	v256 = v253 + v232
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v256
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v258 <= v234 {
		v239 = v253
		v240 = v258
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v278 = F_strlen(m, v277)
	mBase = m.M
	v279 = v278 + v277
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v279
	v285 = v279
	goto L49
L61:
	;
	v302 = v294
	v303 = v298
	goto L64
L62:
	;
	v327 = v285
	goto L63
L63:
	;
	v336 = int32(33)
	*(*uint8)(unsafe.Add(mBase, uint32(v327))) = uint8(v336)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v340 = v338 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v340
	v342 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v340))) = uint8(v342)
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v344 + int32(12)
	F_infix_1(m, l0, int32(4), v342)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L68
	}
L64:
	;
	v314 = v303 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v314
	v316 = F_repalloc(m, v302, v314)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L66
	}
L65:
	;
	v327 = v319
	goto L63
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v316
	v319 = v316 + v295
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v319
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v321 <= v297 {
		v302 = v316
		v303 = v321
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	if l1 < int32(5) {
		goto L3
	} else {
		goto L69
	}
L69:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v356 = v354 - v355
	v358 = v356 + int32(3)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v359 <= v358 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v363 = v355
	v364 = v359
	goto L73
L71:
	;
	v388 = v354
	goto L72
L72:
	;
	v399 = F_pg_sprintf(m, v388, int32(717031), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L77
	}
L73:
	;
	v375 = v364 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v375
	v377 = F_repalloc(m, v363, v375)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L75
	}
L74:
	;
	v388 = v380
	goto L72
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v377
	v380 = v377 + v356
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v380
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v382 <= v358 {
		v363 = v377
		v364 = v382
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v402 = F_strlen(m, v401)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v402 + v401
	goto L3
L78:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v421 = v419 - v420
	v423 = v421 + int32(3)
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v424 <= v423 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	v484 = v407
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v484
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v487 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+92)) = v487
	*(*int32)(unsafe.Add(mBase, uint32(v16)+88)) = v486
	v491 = F_palloc(m, v487)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L1
	} else {
		goto L89
	}
L81:
	;
	v428 = v420
	v429 = v424
	goto L84
L82:
	;
	v453 = v419
	goto L83
L83:
	;
	v464 = F_pg_sprintf(m, v453, int32(780668), int32(0))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L1
	} else {
		goto L88
	}
L84:
	;
	v440 = v429 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v440
	v442 = F_repalloc(m, v428, v440)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L86
	}
L85:
	;
	v453 = v445
	goto L83
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v442
	v445 = v442 + v421
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v445
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v447 <= v423 {
		v428 = v442
		v429 = v447
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v467 = F_strlen(m, v466)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v467 + v466
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v484 = v470
	goto L80
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+84)) = v491
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v491
	F_infix_1(m, v16+int32(76), v416, v410)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v16)+76))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v499
	F_infix_1(m, l0, v416, int32(0))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v507 = v505 - v506
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v16)+84))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v16)+80))
	if v504 <= v507+v508-v510+int32(16) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v519 = v506
	v520 = v504
	goto L95
L93:
	;
	v548 = v505
	v549 = v510
	goto L94
L94:
	;
	switch v224 - int32(2) {
	case 0:
		goto L103
	case 1:
		goto L100
	case 2:
		goto L102
	default:
		goto L101
	}
L95:
	;
	v531 = v520 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v531
	v533 = F_repalloc(m, v519, v531)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L97
	}
L96:
	;
	v548 = v536
	v549 = v541
	goto L94
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v533
	v536 = v533 + v507
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v536
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v16)+84))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v16)+80))
	if v538 <= v507+int32(16)+v539-v541 {
		v519 = v533
		v520 = v538
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v600 = F_strlen(m, v599)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v600 + v599
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v16)+80))
	F_pfree(m, v603)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L1
	} else {
		goto L114
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v549
	v597 = F_pg_sprintf(m, v548, int32(188002), v16+int32(16))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L1
	} else {
		goto L113
	}
L101:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L1
	} else {
		goto L110
	}
L102:
	;
	if v405 != int32(1) {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v549
	v563 = F_pg_sprintf(m, v548, int32(216885), v16+int32(32))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	goto L99
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v549
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v405
	v572 = F_pg_sprintf(m, v548, int32(209050), v16-int32(-64))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L1
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v549
	v578 = F_pg_sprintf(m, v548, int32(209127), v16+int32(48))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
	} else {
		goto L109
	}
L108:
	;
	goto L99
L109:
	;
	goto L99
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v222
	F_errmsg_internal(m, int32(507975), v16)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(515884), int32(1130), int32(27861))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	goto L99
L114:
	;
	if v418 == int32(0) {
		goto L3
	} else {
		goto L115
	}
L115:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v609 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v610 = v608 - v609
	v612 = v610 + int32(3)
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v613 <= v612 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v617 = v609
	v618 = v613
	goto L119
L117:
	;
	v642 = v608
	goto L118
L118:
	;
	v653 = F_pg_sprintf(m, v642, int32(717031), int32(0))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L1
	} else {
		goto L123
	}
L119:
	;
	v629 = v618 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v629
	v631 = F_repalloc(m, v617, v629)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L1
	} else {
		goto L121
	}
L120:
	;
	v642 = v634
	goto L118
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v631
	v634 = v631 + v610
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v634
	v636 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v636 <= v612 {
		v617 = v631
		v618 = v636
		goto L119
	} else {
		goto L122
	}
L122:
	;
	goto L120
L123:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v656 = F_strlen(m, v655)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v656 + v655
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v96 int32
	_ = v96
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v127 int64
	_ = v127
	var v129 int32
	_ = v129
	var v131 int64
	_ = v131
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int64
	_ = v143
	var v145 int64
	_ = v145
	var v147 int32
	_ = v147
	var v149 int64
	_ = v149
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int64
	_ = v177
	var v179 int64
	_ = v179
	var v181 int32
	_ = v181
	var v183 int64
	_ = v183
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int64
	_ = v214
	var v216 int64
	_ = v216
	var v218 int32
	_ = v218
	var v220 int64
	_ = v220
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int64
	_ = v235
	var v237 int64
	_ = v237
	var v239 int32
	_ = v239
	var v241 int64
	_ = v241
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int64
	_ = v253
	var v255 int64
	_ = v255
	var v257 int32
	_ = v257
	var v259 int64
	_ = v259
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int64
	_ = v271
	var v273 int64
	_ = v273
	var v275 int32
	_ = v275
	var v277 int64
	_ = v277
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int64
	_ = v305
	var v307 int64
	_ = v307
	var v309 int32
	_ = v309
	var v311 int64
	_ = v311
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int64
	_ = v342
	var v344 int64
	_ = v344
	var v346 int32
	_ = v346
	var v348 int64
	_ = v348
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v390 int32
	_ = v390
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v436 int32
	_ = v436
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	v2 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(16)
	m.G0 = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v29 < int32(33) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v408)))
	if v390 < v409 {
		goto L57
	} else {
		goto L58
	}
L2:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v38 = F_AllocSetContextCreateInternal(m, v33, int32(66549), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
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
	v371 = m.ExcPending
	if v371 != 0 {
		goto L5
	} else {
		goto L54
	}
L5:
	;
	return int32(0)
L6:
	;
	v42 = int32(4562080)
	v43 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v38
	v47 = F_palloc(m, int32(8212))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v38
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v51
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+10)))
	v55 = F_CreateTupleDescTruncatedCopy(m, v51, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v55
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v59 = int32(*(*int16)(unsafe.Add(mBase, uint32(v58)+10)))
	if v59 <= int32(0) {
		v390 = v2
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v96 = v2
	goto L10
L10:
	;
	v114 = v96 * int32(28)
	v115 = v47 + int32(20) + v114
	v116 = int32(1)
	v117 = v96 + v116
	v118 = base.I32_extend16_s(v117)
	v120 = F_index_getprocinfo(m, l0, v118, v116)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L5
	} else {
		goto L12
	}
L11:
	;
	v390 = v117
	goto L1
L12:
	;
	v124 = v115 + int32(16)
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v120)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v124))) = v125
	v127 = *(*int64)(unsafe.Add(mBase, uint32(v120)))
	*(*int64)(unsafe.Add(mBase, uint32(v115))) = v127
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v120)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v115)+24)) = v129
	v131 = *(*int64)(unsafe.Add(mBase, uint32(v120)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v115)+8)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v115)+20)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v124))) = int32(0)
	goto L13
L13:
	;
	v136 = v114 + (v47 + int32(916))
	v138 = F_index_getprocinfo(m, l0, v118, int32(2))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v142 = v136 + int32(16)
	v143 = *(*int64)(unsafe.Add(mBase, uint32(v138)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v142))) = v143
	v145 = *(*int64)(unsafe.Add(mBase, uint32(v138)))
	*(*int64)(unsafe.Add(mBase, uint32(v136))) = v145
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v138)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v136)+24)) = v147
	v149 = *(*int64)(unsafe.Add(mBase, uint32(v138)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v136)+8)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v136)+20)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v142))) = int32(0)
	goto L15
L15:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+6)))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v155+v157*(v118-int32(1))<<(uint(int32(2))%32)+int32(12)-int32(4))))
	goto L17
L16:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v193)+6)))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v192+v194*(v118-int32(1))<<(uint(int32(2))%32)+int32(16)-int32(4))))
	goto L24
L17:
	;
	if v169 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v170 = v114 + (v47 + int32(1812))
	v172 = F_index_getprocinfo(m, l0, v118, int32(3))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L5
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114+(v47+int32(1816))))) = int32(0)
	goto L16
L21:
	;
	v176 = v170 + int32(16)
	v177 = *(*int64)(unsafe.Add(mBase, uint32(v172)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v176))) = v177
	v179 = *(*int64)(unsafe.Add(mBase, uint32(v172)))
	*(*int64)(unsafe.Add(mBase, uint32(v170))) = v179
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v172)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v170)+24)) = v181
	v183 = *(*int64)(unsafe.Add(mBase, uint32(v172)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v170)+8)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v170)+20)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v176))) = int32(0)
	goto L22
L22:
	;
	goto L16
L23:
	;
	v228 = v114 + (v47 + int32(3604))
	v230 = F_index_getprocinfo(m, l0, v118, int32(5))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L5
	} else {
		goto L30
	}
L24:
	;
	if v206 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v207 = v114 + (v47 + int32(2708))
	v209 = F_index_getprocinfo(m, l0, v118, int32(4))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L5
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114+(v47+int32(2712))))) = int32(0)
	goto L23
L28:
	;
	v213 = v207 + int32(16)
	v214 = *(*int64)(unsafe.Add(mBase, uint32(v209)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v213))) = v214
	v216 = *(*int64)(unsafe.Add(mBase, uint32(v209)))
	*(*int64)(unsafe.Add(mBase, uint32(v207))) = v216
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v209)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v207)+24)) = v218
	v220 = *(*int64)(unsafe.Add(mBase, uint32(v209)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v207)+8)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v207)+20)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v213))) = int32(0)
	goto L29
L29:
	;
	goto L23
L30:
	;
	v234 = v228 + int32(16)
	v235 = *(*int64)(unsafe.Add(mBase, uint32(v230)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v234))) = v235
	v237 = *(*int64)(unsafe.Add(mBase, uint32(v230)))
	*(*int64)(unsafe.Add(mBase, uint32(v228))) = v237
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v230)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v228)+24)) = v239
	v241 = *(*int64)(unsafe.Add(mBase, uint32(v230)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v228)+8)) = v241
	*(*int32)(unsafe.Add(mBase, uint32(v228)+20)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = int32(0)
	goto L31
L31:
	;
	v246 = v114 + (v47 + int32(4500))
	v248 = F_index_getprocinfo(m, l0, v118, int32(6))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v252 = v246 + int32(16)
	v253 = *(*int64)(unsafe.Add(mBase, uint32(v248)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v252))) = v253
	v255 = *(*int64)(unsafe.Add(mBase, uint32(v248)))
	*(*int64)(unsafe.Add(mBase, uint32(v246))) = v255
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v248)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v246)+24)) = v257
	v259 = *(*int64)(unsafe.Add(mBase, uint32(v248)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v246)+8)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v246)+20)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v252))) = int32(0)
	goto L33
L33:
	;
	v264 = v114 + (v47 + int32(5396))
	v266 = F_index_getprocinfo(m, l0, v118, int32(7))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v270 = v264 + int32(16)
	v271 = *(*int64)(unsafe.Add(mBase, uint32(v266)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v270))) = v271
	v273 = *(*int64)(unsafe.Add(mBase, uint32(v266)))
	*(*int64)(unsafe.Add(mBase, uint32(v264))) = v273
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v266)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v264)+24)) = v275
	v277 = *(*int64)(unsafe.Add(mBase, uint32(v266)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v264)+8)) = v277
	*(*int32)(unsafe.Add(mBase, uint32(v264)+20)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v270))) = int32(0)
	goto L35
L35:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v284)+6)))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v283+v285*(v118-int32(1))<<(uint(int32(2))%32)+int32(32)-int32(4))))
	goto L37
L36:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v322 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v321)+6)))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v320+v322*(v118-int32(1))<<(uint(int32(2))%32)+int32(36)-int32(4))))
	goto L44
L37:
	;
	if v297 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v298 = v114 + (v47 + int32(6292))
	v300 = F_index_getprocinfo(m, l0, v118, int32(8))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L5
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114+(v47+int32(6296))))) = int32(0)
	goto L36
L41:
	;
	v304 = v298 + int32(16)
	v305 = *(*int64)(unsafe.Add(mBase, uint32(v300)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v304))) = v305
	v307 = *(*int64)(unsafe.Add(mBase, uint32(v300)))
	*(*int64)(unsafe.Add(mBase, uint32(v298))) = v307
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v300)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v298)+24)) = v309
	v311 = *(*int64)(unsafe.Add(mBase, uint32(v300)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v298)+8)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(v298)+20)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v304))) = int32(0)
	goto L42
L42:
	;
	goto L36
L43:
	;
	v357 = v96 << (uint(int32(2)) % 32)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v359+v357)))
	if v361 != 0 {
		goto L50
	} else {
		goto L51
	}
L44:
	;
	if v334 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v335 = v114 + (v47 + int32(7188))
	v337 = F_index_getprocinfo(m, l0, v118, int32(9))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L5
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114+(v47+int32(7192))))) = int32(0)
	goto L43
L48:
	;
	v341 = v335 + int32(16)
	v342 = *(*int64)(unsafe.Add(mBase, uint32(v337)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v341))) = v342
	v344 = *(*int64)(unsafe.Add(mBase, uint32(v337)))
	*(*int64)(unsafe.Add(mBase, uint32(v335))) = v344
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v337)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v335)+24)) = v346
	v348 = *(*int64)(unsafe.Add(mBase, uint32(v337)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v335)+8)) = v348
	*(*int32)(unsafe.Add(mBase, uint32(v335)+20)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v341))) = int32(0)
	goto L49
L49:
	;
	goto L43
L50:
	;
	v363 = v361
	goto L52
L51:
	;
	v363 = int32(100)
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47+int32(8084)+v357))) = v363
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v366 = int32(*(*int16)(unsafe.Add(mBase, uint32(v365)+10)))
	if v117 < v366 {
		v96 = v117
		goto L10
	} else {
		goto L53
	}
L53:
	;
	goto L11
L54:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v373
	F_errmsg_internal(m, int32(503028), v26)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(516299), int32(1547), int32(369677))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
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
	v436 = v390
	goto L60
L58:
	;
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v43
	m.G0 = v26 + int32(16)
	return v47
L60:
	;
	v455 = v436 * int32(28)
	v457 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v47+int32(24)+v455))) = v457
	*(*int32)(unsafe.Add(mBase, uint32(v455+(v47+int32(920))))) = v457
	*(*int32)(unsafe.Add(mBase, uint32(v455+(v47+int32(1816))))) = v457
	*(*int32)(unsafe.Add(mBase, uint32(v455+(v47+int32(2712))))) = v457
	*(*int32)(unsafe.Add(mBase, uint32(v455+(v47+int32(3608))))) = v457
	*(*int32)(unsafe.Add(mBase, uint32(v455+(v47+int32(4504))))) = v457
	*(*int32)(unsafe.Add(mBase, uint32(v455+(v47+int32(5400))))) = v457
	*(*int32)(unsafe.Add(mBase, uint32(v455+(v47+int32(6296))))) = v457
	*(*int32)(unsafe.Add(mBase, uint32(v455+(v47+int32(7192))))) = v457
	*(*int32)(unsafe.Add(mBase, uint32(v47+int32(8084)+v436<<(uint(int32(2))%32)))) = v457
	v489 = v436 + int32(1)
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v490)))
	if v489 < v491 {
		v436 = v489
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
			F_errmsg_internal(m, int32(361302), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				F_errfinish(m, int32(523084), int32(71), int32(344292))
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
	var v114 int32
	_ = v114
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
	var v184 int32
	_ = v184
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
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
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
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
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v302 int64
	_ = v302
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v521 int32
	_ = v521
	var v532 int32
	_ = v532
	var v539 int32
	_ = v539
	var v547 int32
	_ = v547
	var v558 int32
	_ = v558
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v572 int32
	_ = v572
	v4 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v4 < v12 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572)+20)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(l1)+48)) = int64(0)
	return v572
L2:
	;
	v547 = int32(0)
	goto L114
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
	if v16&int32(1) != 0 {
		v539 = v15
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
	if v387 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L8:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v197)+12))
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
		v197 = v38
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
	v197 = v38
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
	*(*int32)(unsafe.Add(mBase, uint32(l1)+56)) = v184 + int32(32)
	v197 = v184
	goto L8
L20:
	;
	v114 = v103
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
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v114)+20))
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
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+8)))
	if v125&int32(4) == int32(0) {
		v184 = v114
		goto L19
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v131 = v114 + int32(32)
	if base.Ui32(v131) < base.Ui32(v107) {
		v114 = v131
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
	v184 = v147
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
	v387 = int32(0)
	goto L7
L48:
	;
	v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v197)+16)))
	v207 = v202
	v208 = v203
	goto L51
L49:
	;
	goto L50
L50:
	;
	v239 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v197)+12)) = v239
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v239 < v241 {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v207)+24))
	v215 = base.I32_extend16_s(v208)
	v219 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v214+v215<<(uint(int32(2))%32)))) = v219
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v207)+28))
	v224 = v221 + v215<<(uint(int32(3))%32)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	*(*int32)(unsafe.Add(mBase, uint32(v224))) = v219
	v228 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v224)+4)))
	if v225 != 0 {
		v207 = v225
		v208 = v228
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
	v249 = v241
	v251 = int32(0)
	goto L57
L55:
	;
	goto L56
L56:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v197)+8))
	if v348&int32(2) == int32(0) {
		v362 = v348
		goto L73
	} else {
		goto L74
	}
L57:
	;
	v256 = v251 << (uint(int32(2)) % 32)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v197)+24))
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
	if v260 != v197 {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v329 = v249
	goto L61
L61:
	;
	v336 = v251 + int32(1)
	if v336 < v329 {
		v249 = v329
		v251 = v336
		goto L57
	} else {
		goto L72
	}
L62:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v197)+24))
	v316 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v314+v256))) = v316
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v197)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v318+v251<<(uint(int32(3))%32)))) = v316
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v329 = v324
	goto L61
L63:
	;
	v270 = int32(*(*int16)(unsafe.Add(mBase, uint32(v259)+16)))
	v274 = v270
	v275 = v260
	goto L67
L64:
	;
	v262 = int32(*(*int16)(unsafe.Add(mBase, uint32(v259)+16)))
	if v251 != v262 {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v197)+28))
	v268 = *(*int64)(unsafe.Add(mBase, uint32(v264+v251<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v259)+12)) = v268
	goto L62
L66:
	;
	v295 = int32(3)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v197)+28))
	v302 = *(*int64)(unsafe.Add(mBase, uint32(v298+v251<<(uint(v295)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v294+v293<<(uint(v295)%32)))) = v302
	goto L62
L67:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v275)+28))
	v284 = v281 + v274<<(uint(int32(3))%32)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)))
	if v285 == int32(0) {
		v293 = v274
		v294 = v281
		goto L66
	} else {
		goto L69
	}
L68:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v275)+28))
	v293 = base.I32_extend16_s(v274)
	v294 = v291
	goto L66
L69:
	;
	v288 = int32(*(*int16)(unsafe.Add(mBase, uint32(v284)+4)))
	if v285 != v197 {
		v274 = v288
		v275 = v285
		goto L67
	} else {
		goto L70
	}
L70:
	;
	if v288 != v251 {
		v274 = v288
		v275 = v285
		goto L67
	} else {
		goto L71
	}
L71:
	;
	goto L68
L72:
	;
	goto L58
L73:
	;
	if v362&int32(8) == int32(0) {
		goto L80
	} else {
		goto L81
	}
L74:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v197)+20))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if v353 == v354 {
		v362 = v348
		goto L73
	} else {
		goto L75
	}
L75:
	;
	if base.Ui32(v353) <= base.Ui32(v354) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v358 = v354
	goto L78
L77:
	;
	v358 = int32(0)
	goto L78
L78:
	;
	if v358 != 0 {
		v362 = v348
		goto L73
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v353
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v197)+8))
	v362 = v360
	goto L73
L80:
	;
	v387 = v197
	goto L7
L81:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v197)+20))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v368 == v369 {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	if base.Ui32(v368) <= base.Ui32(v369) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v373 = v369
	goto L85
L84:
	;
	v373 = int32(0)
	goto L85
L85:
	;
	if v373 != 0 {
		goto L80
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = v368
	goto L80
L87:
	;
	return int32(0)
L88:
	;
	goto L89
L89:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if int32(0) < v392 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v396 = int32(0)
	goto L93
L91:
	;
	goto L92
L92:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v387)))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v429)+12))
	v435 = v428 + int32(base.Ui32(v430)>>(uint(int32(3))%32))&int32(536870908)
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v435)))
	v437 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v435))) = v436 | v437<<(uint(v430)%32)
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v387)))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v442 == v437 {
		goto L97
	} else {
		goto L98
	}
L93:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v387)))
	*(*int32)(unsafe.Add(mBase, uint32(v407+v396<<(uint(int32(2))%32)))) = int32(0)
	v414 = v396 + int32(1)
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v414 < v415 {
		v396 = v414
		goto L93
	} else {
		goto L95
	}
L94:
	;
	goto L92
L95:
	;
	goto L94
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v387)+8)) = int32(13)
	*(*int32)(unsafe.Add(mBase, uint32(v387)+4)) = v521
	v532 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v532 <= int32(0) {
		v572 = v387
		goto L1
	} else {
		goto L113
	}
L97:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v441)))
	v521 = v445
	goto L96
L98:
	;
	goto L99
L99:
	;
	if v442 <= int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v521 = int32(0)
	goto L96
L101:
	;
	goto L102
L102:
	;
	v450 = v442 & int32(3)
	v451 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v442) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v457 = v451
	v460 = v451
	v466 = v4
	goto L106
L104:
	;
	v484 = v451
	v487 = v451
	goto L105
L105:
	;
	if v450 == int32(0) {
		v521 = v487
		goto L96
	} else {
		goto L109
	}
L106:
	;
	v470 = v441 + v457<<(uint(int32(2))%32)
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v470)+12))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v470)+8))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v470)+4))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v470)))
	v478 = v471 ^ (v472 ^ (v473 ^ (v474 ^ v460)))
	v479 = int32(4)
	v480 = v457 + v479
	v482 = v466 + v479
	if v482 != v442&int32(2147483644) {
		v457 = v480
		v460 = v478
		v466 = v482
		goto L106
	} else {
		goto L108
	}
L107:
	;
	v484 = v480
	v487 = v478
	goto L105
L108:
	;
	goto L107
L109:
	;
	v497 = v484
	v500 = v487
	v505 = v4
	goto L110
L110:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v441+v497<<(uint(int32(2))%32))))
	v512 = v511 ^ v500
	v513 = int32(1)
	v516 = v505 + v513
	if v516 != v450 {
		v497 = v497 + v513
		v500 = v512
		v505 = v516
		goto L110
	} else {
		goto L112
	}
L111:
	;
	v521 = v512
	goto L96
L112:
	;
	goto L111
L113:
	;
	v539 = v387
	goto L2
L114:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v558+v547<<(uint(int32(5))%32))+20)) = int32(0)
	v565 = v547 + int32(1)
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v565 < v566 {
		v547 = v565
		goto L114
	} else {
		goto L116
	}
L115:
	;
	v572 = v539
	goto L1
L116:
	;
	goto L115
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
								v40 = *(*int32)(unsafe.Add(mBase, _consts[130]))
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
							v40 = *(*int32)(unsafe.Add(mBase, _consts[130]))
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
						v40 = *(*int32)(unsafe.Add(mBase, _consts[130]))
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
							v40 = *(*int32)(unsafe.Add(mBase, _consts[130]))
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
						v40 = *(*int32)(unsafe.Add(mBase, _consts[130]))
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
					v40 = *(*int32)(unsafe.Add(mBase, _consts[130]))
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
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
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v246 int32
	_ = v246
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v275 int32
	_ = v275
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
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
	return base.I32_extend16_s(v405)
L2:
	;
	F_errsave_finish(m, v11, int32(518001), v400, int32(429606))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L81
	} else {
		goto L90
	}
L3:
	;
	v374 = int32(0)
	v375 = F_errsave_start(m, v11)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L81
	} else {
		goto L86
	}
L4:
	;
	v348 = int32(0)
	v349 = F_errsave_start(m, v11)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L81
	} else {
		goto L82
	}
L5:
	;
	v95 = v10
	v101 = v17
	goto L22
L6:
	;
	v29 = v20 + int32(1)
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	v32 = v30 - int32(48)
	if base.Ui32(v32&int32(255)) <= base.Ui32(int32(9)) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v37 = v29
	v38 = v25
	v39 = v32
	goto L10
L8:
	;
	v65 = v25
	v67 = v30
	goto L9
L9:
	;
	if v67 != 0 {
		goto L5
	} else {
		goto L14
	}
L10:
	;
	if base.Ui32(int32(3276)) < base.Ui32(v38&int32(65535)) {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	v65 = v54
	v67 = v57
	goto L9
L12:
	;
	v50 = int32(10)
	v52 = int32(255)
	v54 = v38*v50 + v39&v52
	v56 = v37 + int32(1)
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	v59 = v57 - int32(48)
	if base.Ui32(v59&v52) < base.Ui32(v50) {
		v37 = v56
		v38 = v54
		v39 = v59
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
	if base.Ui32(v65&int32(65535)-int32(32769)) < base.Ui32(int32(-65536)) {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if base.I32_extend16_s(v65) < int32(0) {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	v405 = int32(0) - v65
	goto L1
L19:
	;
	v405 = v65
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
	if base.Ui32(v101-int32(9)) < base.Ui32(int32(5)) {
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
	v111 = v95 + int32(1)
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	v95 = v111
	v101 = v112
	goto L22
L26:
	;
	switch v101 - int32(32) {
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
	if v300 == v302 {
		goto L3
	} else {
		goto L69
	}
L28:
	;
	v264 = int32(0)
	v265 = v118
	v266 = v120
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
	v223 = v118 + int32(2)
	v225 = int32(0)
	v226 = v223
	goto L52
L31:
	;
	v182 = v118 + int32(2)
	v184 = int32(0)
	v185 = v182
	goto L44
L32:
	;
	v128 = v118 + int32(2)
	v130 = int32(0)
	v131 = v128
	goto L33
L33:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
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
	if base.Ui32(int32(2048)) < base.Ui32(v130&int32(65535)) {
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
	v156 = int32(*(*int8)(unsafe.Add(mBase, uint32(v138)+uint32(_consts[1099]))))
	v130 = v156 + v130<<(uint(int32(4))%32)
	v131 = v131 + int32(1)
	goto L33
L40:
	;
	v165 = v131 + int32(1)
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
	if v166 == int32(0) {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	goto L42
L42:
	;
	if base.B2i32(base.Ui32(v166-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v166|int32(32)-int32(97)) < base.Ui32(int32(6))) != 0 {
		v131 = v165
		goto L33
	} else {
		goto L43
	}
L43:
	;
	goto L34
L44:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
	if v192&int32(248) == int32(48) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L3
L46:
	;
	if base.Ui32(int32(4096)) < base.Ui32(v184&int32(65535)) {
		goto L4
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if v192 != int32(95) {
		v300 = v182
		v301 = v184
		v302 = v185
		v303 = v192
		goto L27
	} else {
		goto L50
	}
L49:
	;
	v184 = (v192-int32(48))&int32(255) | v184<<(uint(int32(3))%32)
	v185 = v185 + int32(1)
	goto L44
L50:
	;
	v213 = v185 + int32(1)
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
	if base.Ui32(int32(248)) <= base.Ui32((v214-int32(56))&int32(255)) {
		v185 = v213
		goto L44
	} else {
		goto L51
	}
L51:
	;
	goto L45
L52:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226))))
	if v233&int32(254) == int32(48) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L3
L54:
	;
	if base.Ui32(int32(16384)) < base.Ui32(v225&int32(65535)) {
		goto L4
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v233 != int32(95) {
		v300 = v223
		v301 = v225
		v302 = v226
		v303 = v233
		goto L27
	} else {
		goto L58
	}
L57:
	;
	v246 = int32(1)
	v225 = (v233-int32(48))&int32(255) | v225<<(uint(v246)%32)
	v226 = v226 + v246
	goto L52
L58:
	;
	v254 = v226 + int32(1)
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254))))
	if base.Ui32(int32(254)) <= base.Ui32((v255-int32(50))&int32(255)) {
		v226 = v254
		goto L52
	} else {
		goto L59
	}
L59:
	;
	goto L53
L60:
	;
	v275 = (v266 - int32(48)) & int32(255)
	if base.Ui32(v275) <= base.Ui32(int32(9)) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	goto L3
L62:
	;
	if base.Ui32(int32(3276)) < base.Ui32(v264&int32(65535)) {
		goto L4
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	if v266 != int32(95) {
		v300 = v118
		v301 = v264
		v302 = v265
		v303 = v266
		goto L27
	} else {
		goto L66
	}
L65:
	;
	v286 = v265 + int32(1)
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286))))
	v264 = v264*int32(10) + v275
	v265 = v286
	v266 = v287
	goto L60
L66:
	;
	if v118 == v265 {
		goto L3
	} else {
		goto L67
	}
L67:
	;
	v292 = v265 + int32(1)
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
	if base.Ui32((v293-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v265 = v292
		v266 = v293
		goto L60
	} else {
		goto L68
	}
L68:
	;
	goto L61
L69:
	;
	v312 = v302
	v313 = v303
	goto L70
L70:
	;
	if base.Ui32(v313-int32(9)) < base.Ui32(int32(5)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v337 = v312 + int32(1)
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337))))
	v312 = v337
	v313 = v338
	goto L70
L73:
	;
	if v313 == int32(32) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	if v313 != 0 {
		goto L3
	} else {
		goto L75
	}
L75:
	;
	if v119 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	if base.Ui32(v301&int32(65535)-int32(32769)) < base.Ui32(int32(-65536)) {
		goto L4
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	if int32(0) <= base.I32_extend16_s(v301) {
		v405 = v301
		goto L1
	} else {
		goto L80
	}
L79:
	;
	v405 = int32(0) - v301
	goto L1
L80:
	;
	goto L4
L81:
	;
	return int32(0)
L82:
	;
	if v349 == int32(0) {
		v405 = v348
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L81
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = int32(95552)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v10
	F_errmsg(m, int32(200611), v14)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L81
	} else {
		goto L85
	}
L85:
	;
	v392 = v348
	v400 = int32(351)
	goto L2
L86:
	;
	if v375 == int32(0) {
		v405 = v374
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L81
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(95552)
	F_errmsg(m, int32(759399), v14+int32(16))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L81
	} else {
		goto L89
	}
L89:
	;
	v392 = v374
	v400 = int32(357)
	goto L2
L90:
	;
	v405 = v392
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
				F_errmsg_internal(m, int32(26664), int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(525301), int32(6953), int32(300520))
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
					F_errmsg_internal(m, int32(26664), int32(0))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(525301), int32(6953), int32(300520))
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
	v7 = int32(65535)
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
					F_errmsg(m, int32(251839), int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(525865), int32(1084), int32(37056))
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
					F_errmsg(m, int32(421289), int32(0))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(525865), int32(1100), int32(37056))
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
					F_errmsg(m, int32(251839), int32(0))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(525865), int32(514), int32(36954))
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
						F_errmsg(m, int32(421289), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(525865), int32(530), int32(36954))
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
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
					v147 = m.ExcPending
					if v147 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67108994))
						mBase = m.M
						v150 = m.ExcPending
						if v150 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(162284), int32(0))
							mBase = m.M
							v156 = m.ExcPending
							if v156 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(520256), int32(359), int32(305872))
								mBase = m.M
								v163 = m.ExcPending
								if v163 != 0 {
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
					v24 = F_ArrayGetNItems(m, v21, v23)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						if v24 != 0 {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
							v27 = F_ArrayGetNItems(m, v26, v23)
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
									v119 = int32(0)
								} else {
									v43 = v39 + v13
									v44 = int32(1)
									if v27 == v44 {
										v48 = int32(0)
										v93 = v48
										v94 = v48
									} else {
										v52 = int32(0)
										v54 = v52
										v55 = v52
										v61 = int32(0)
										for {
											v67 = v43 + v54<<(uint(int32(2))%32)
											v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
											if v17 == v68 {
												v77 = v55
											} else {
												v71 = v55 + int32(1)
												if v54 <= v55 {
													v77 = v71
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v43+v55<<(uint(int32(2))%32)))) = v68
													v77 = v71
												}
											}
											v78 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
											if v17 == v78 {
												v87 = v77
											} else {
												v81 = v77 + int32(1)
												if v54 < v77 {
													v87 = v81
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v43+v77<<(uint(int32(2))%32)))) = v78
													v87 = v81
												}
											}
											v88 = int32(2)
											v89 = v54 + v88
											v91 = v61 + v88
											if v91 != v27&int32(2147483646) {
												v54 = v89
												v55 = v87
												v61 = v91
												continue
											} else {
												break
											}
											break
										}
										v93 = v89
										v94 = v87
									}
									if v27&v44 == int32(0) {
										v119 = v94
									} else {
										v109 = *(*int32)(unsafe.Add(mBase, uint32(v43+v93<<(uint(int32(2))%32))))
										if v109 == v17 {
											v119 = v94
										} else {
											if v94 < v93 {
												*(*int32)(unsafe.Add(mBase, uint32(v43+v94<<(uint(int32(2))%32)))) = v109
											} else {
											}
											v119 = v94 + int32(1)
										}
									}
								}
								v129 = F_resize_intArrayType(m, v13, v119)
								mBase = m.M
								v130 = m.ExcPending
								if v130 != 0 {
									return int32(0)
								} else {
									v142 = v129
									return v142
								}
							}
						} else {
							v142 = v13
							return v142
						}
					}
				}
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			v23 = v13 + int32(16)
			v24 = F_ArrayGetNItems(m, v21, v23)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				if v24 != 0 {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v27 = F_ArrayGetNItems(m, v26, v23)
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
							v119 = int32(0)
						} else {
							v43 = v39 + v13
							v44 = int32(1)
							if v27 == v44 {
								v48 = int32(0)
								v93 = v48
								v94 = v48
							} else {
								v52 = int32(0)
								v54 = v52
								v55 = v52
								v61 = int32(0)
								for {
									v67 = v43 + v54<<(uint(int32(2))%32)
									v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
									if v17 == v68 {
										v77 = v55
									} else {
										v71 = v55 + int32(1)
										if v54 <= v55 {
											v77 = v71
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v43+v55<<(uint(int32(2))%32)))) = v68
											v77 = v71
										}
									}
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
									if v17 == v78 {
										v87 = v77
									} else {
										v81 = v77 + int32(1)
										if v54 < v77 {
											v87 = v81
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v43+v77<<(uint(int32(2))%32)))) = v78
											v87 = v81
										}
									}
									v88 = int32(2)
									v89 = v54 + v88
									v91 = v61 + v88
									if v91 != v27&int32(2147483646) {
										v54 = v89
										v55 = v87
										v61 = v91
										continue
									} else {
										break
									}
									break
								}
								v93 = v89
								v94 = v87
							}
							if v27&v44 == int32(0) {
								v119 = v94
							} else {
								v109 = *(*int32)(unsafe.Add(mBase, uint32(v43+v93<<(uint(int32(2))%32))))
								if v109 == v17 {
									v119 = v94
								} else {
									if v94 < v93 {
										*(*int32)(unsafe.Add(mBase, uint32(v43+v94<<(uint(int32(2))%32)))) = v109
									} else {
									}
									v119 = v94 + int32(1)
								}
							}
						}
						v129 = F_resize_intArrayType(m, v13, v119)
						mBase = m.M
						v130 = m.ExcPending
						if v130 != 0 {
							return int32(0)
						} else {
							v142 = v129
							return v142
						}
					}
				} else {
					v142 = v13
					return v142
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
	v5 = *(*int32)(unsafe.Add(mBase, _consts[1194]))
	if int32(0) <= v5 {
		v9 = v5 * int32(100)
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_consts[1211])))
		if v12 != 0 {
			F_pfree(m, v12)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_consts[1211]))) = int32(0)
				if l0 != 0 {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_consts[1206])))
					v20 = F_MemoryContextStrdup(m, v19, l0)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_consts[1211]))) = v20
						return int32(0)
					}
				} else {
					return int32(0)
				}
			}
		} else {
			if l0 != 0 {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_consts[1206])))
				v20 = F_MemoryContextStrdup(m, v19, l0)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_consts[1211]))) = v20
					return int32(0)
				}
			} else {
				return int32(0)
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[1194])) = int32(-1)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(475542), int32(0))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(523098), int32(1509), int32(15505))
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
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v10 = int32(410550)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, _consts[400])))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v14 == int32(0) {
		v33 = v13
		v34 = v14
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L30
	} else {
		goto L31
	}
L2:
	;
	m.G0 = v5 + int32(16)
	return v94
L3:
	;
	if v34-v33 == int32(0) {
		v94 = int32(105)
		goto L2
	} else {
		goto L11
	}
L4:
	;
	goto L3
L5:
	;
	if v13 != v14 {
		v33 = v13
		v34 = v14
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v18 = v9
	v19 = v10
	goto L7
L7:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	if v23 == int32(0) {
		v33 = v22
		v34 = v23
		goto L4
	} else {
		goto L9
	}
L8:
	;
	v33 = v22
	v34 = v23
	goto L4
L9:
	;
	v26 = int32(1)
	if v22 == v23 {
		v18 = v18 + v26
		v19 = v19 + v26
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v39 = int32(410583)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, _consts[401])))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v43 == int32(0) {
		v62 = v42
		v63 = v43
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v63-v62 == int32(0) {
		v94 = int32(115)
		goto L2
	} else {
		goto L20
	}
L13:
	;
	goto L12
L14:
	;
	if v42 != v43 {
		v62 = v42
		v63 = v43
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v47 = v9
	v48 = v39
	goto L16
L16:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
	if v52 == int32(0) {
		v62 = v51
		v63 = v52
		goto L13
	} else {
		goto L18
	}
L17:
	;
	v62 = v51
	v63 = v52
	goto L13
L18:
	;
	v55 = int32(1)
	if v51 == v52 {
		v47 = v47 + v55
		v48 = v48 + v55
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v67 = int32(405033)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, _consts[402])))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v71 == int32(0) {
		v90 = v70
		v91 = v71
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v91-v90 != 0 {
		goto L1
	} else {
		goto L29
	}
L22:
	;
	goto L21
L23:
	;
	if v70 != v71 {
		v90 = v70
		v91 = v71
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v75 = v9
	v76 = v67
	goto L25
L25:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+1)))
	if v80 == int32(0) {
		v90 = v79
		v91 = v80
		goto L22
	} else {
		goto L27
	}
L26:
	;
	v90 = v79
	v91 = v80
	goto L22
L27:
	;
	v83 = int32(1)
	if v79 == v80 {
		v75 = v75 + v83
		v76 = v76 + v83
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v94 = int32(118)
	goto L2
L30:
	;
	return int32(0)
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v9
	F_errmsg_internal(m, int32(723742), v5)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(518591), int32(629), int32(11549))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L30
	} else {
		goto L33
	}
L33:
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
		F_sequence_close(m, v21, int32(0))
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
				F_sequence_close(m, v21, int32(0))
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
					F_sequence_close(m, v21, int32(0))
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
						F_sequence_close(m, v21, int32(0))
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
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
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
	v8 = F_find_among(m, l0, int32(4342912), int32(24))
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
	v72 = F_slice_from_s(m, l0, int32(1), int32(2245871))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L3
	} else {
		goto L34
	}
L7:
	;
	v66 = F_slice_from_s(m, l0, int32(1), int32(2245870))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L3
	} else {
		goto L32
	}
L8:
	;
	v60 = F_slice_from_s(m, l0, int32(1), int32(2245869))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L3
	} else {
		goto L30
	}
L9:
	;
	v54 = F_slice_from_s(m, l0, int32(1), int32(2245868))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L3
	} else {
		goto L28
	}
L10:
	;
	v48 = F_slice_from_s(m, l0, int32(1), int32(2245867))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L3
	} else {
		goto L26
	}
L11:
	;
	v42 = F_slice_from_s(m, l0, int32(1), int32(2245866))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L3
	} else {
		goto L24
	}
L12:
	;
	v36 = F_slice_from_s(m, l0, int32(1), int32(2245865))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L22
	}
L13:
	;
	v30 = F_slice_from_s(m, l0, int32(1), int32(2245864))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L3
	} else {
		goto L20
	}
L14:
	;
	v24 = F_slice_from_s(m, l0, int32(1), int32(2245863))
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
	v587 = F_find_among_b(m, l0, int32(4343392), int32(16))
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
	v172 = v157&int32(63) | (v115<<(uint(int32(18))%32)&int32(1835008) | v124<<(uint(int32(12))%32) | v140<<(uint(int32(6))%32))
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
	v172 = v115<<(uint(int32(12))%32)&int32(61440) | v124<<(uint(int32(6))%32) | v140
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
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v177)>>(uint(int32(3))%32)))+uint32(_consts[1334]))))
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
	v296 = v281&int32(63) | (v239<<(uint(int32(18))%32)&int32(1835008) | v248<<(uint(int32(12))%32) | v264<<(uint(int32(6))%32))
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
	v296 = v239<<(uint(int32(12))%32)&int32(61440) | v248<<(uint(int32(6))%32) | v264
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
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v301)>>(uint(int32(3))%32)))+uint32(_consts[1334]))))
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
	v421 = v406&int32(63) | (v364<<(uint(int32(18))%32)&int32(1835008) | v373<<(uint(int32(12))%32) | v389<<(uint(int32(6))%32))
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
	v421 = v364<<(uint(int32(12))%32)&int32(61440) | v373<<(uint(int32(6))%32) | v389
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
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v426)>>(uint(int32(3))%32)))+uint32(_consts[1334]))))
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
	v543 = v528&int32(63) | (v486<<(uint(int32(18))%32)&int32(1835008) | v495<<(uint(int32(12))%32) | v511<<(uint(int32(6))%32))
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
	v543 = v486<<(uint(int32(12))%32)&int32(61440) | v495<<(uint(int32(6))%32) | v511
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
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v548)>>(uint(int32(3))%32)))+uint32(_consts[1334]))))
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
	v611 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v611
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v611
	v616 = F_find_among_b(m, l0, int32(4343712), int32(25))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
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
	v663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v663
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v663
	v666 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v663-int32(2) <= v666 {
		goto L172
	} else {
		goto L173
	}
L151:
	;
	if v616 == int32(0) {
		goto L150
	} else {
		goto L152
	}
L152:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v620
	switch v616 - int32(1) {
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
	v657 = F_slice_from_s(m, l0, int32(4), int32(2246043))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L3
	} else {
		goto L170
	}
L154:
	;
	v651 = F_slice_from_s(m, l0, int32(5), int32(2246038))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L3
	} else {
		goto L168
	}
L155:
	;
	v645 = F_slice_from_s(m, l0, int32(4), int32(2246034))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L3
	} else {
		goto L166
	}
L156:
	;
	v639 = F_slice_from_s(m, l0, int32(3), int32(2246031))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L3
	} else {
		goto L164
	}
L157:
	;
	v633 = F_slice_from_s(m, l0, int32(3), int32(2246028))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L3
	} else {
		goto L162
	}
L158:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v624)))
	if v620 < v625 {
		goto L150
	} else {
		goto L159
	}
L159:
	;
	v627 = F_slice_del(m, l0)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L3
	} else {
		goto L160
	}
L160:
	;
	if int32(0) <= v627 {
		goto L150
	} else {
		goto L161
	}
L161:
	;
	v714 = v627
	goto L1
L162:
	;
	if int32(0) <= v633 {
		goto L150
	} else {
		goto L163
	}
L163:
	;
	v714 = v633
	goto L1
L164:
	;
	if int32(0) <= v639 {
		goto L150
	} else {
		goto L165
	}
L165:
	;
	v714 = v639
	goto L1
L166:
	;
	if int32(0) <= v645 {
		goto L150
	} else {
		goto L167
	}
L167:
	;
	v714 = v645
	goto L1
L168:
	;
	if int32(0) <= v651 {
		goto L150
	} else {
		goto L169
	}
L169:
	;
	v714 = v651
	goto L1
L170:
	;
	if v657 < int32(0) {
		v714 = v657
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
	v670 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v670+v663-int32(1)))))
	if v674&int32(224) != int32(96) {
		goto L172
	} else {
		goto L174
	}
L174:
	;
	if int32(1)<<(uint(v674)%32)&int32(282896) == int32(0) {
		goto L172
	} else {
		goto L175
	}
L175:
	;
	v687 = F_find_among_b(m, l0, int32(4344224), int32(12))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L3
	} else {
		goto L176
	}
L176:
	;
	if v687 == int32(0) {
		goto L172
	} else {
		goto L177
	}
L177:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v691
	switch v687 - int32(1) {
	case 0:
		goto L179
	case 1:
		goto L178
	default:
		goto L172
	}
L178:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v702)+4))
	if v691 < v703 {
		goto L172
	} else {
		goto L183
	}
L179:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v695)+8))
	if v691 < v696 {
		goto L172
	} else {
		goto L180
	}
L180:
	;
	v698 = F_slice_del(m, l0)
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L3
	} else {
		goto L181
	}
L181:
	;
	if int32(0) <= v698 {
		goto L172
	} else {
		goto L182
	}
L182:
	;
	v714 = v698
	goto L1
L183:
	;
	v705 = F_slice_del(m, l0)
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L3
	} else {
		goto L184
	}
L184:
	;
	if v705 < int32(0) {
		v714 = v705
		goto L1
	} else {
		goto L185
	}
L185:
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
										v32 = F_ArrayGetNItems(m, v29, v27+int32(16))
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
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	F_ean2isn(m, v8, v5+int32(8), int32(3))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
		v17 = F_Int64GetDatum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			m.G0 = v5 + int32(16)
			return v17
		}
	}
}
func F_isbn_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = F_string2ean(m, v7, v8, v5+int32(8), int32(3))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			v18 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v18)
			v24 = int32(0)
			m.G0 = v5 + int32(16)
			return v24
		} else {
			v21 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
			v22 = F_Int64GetDatum(m, v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = v22
				m.G0 = v5 + int32(16)
				return v24
			}
		}
	}
}
func F_isgraph(m *base.Module, l0 int32) int32 {
	return base.B2i32(base.Ui32(l0-int32(33)) < base.Ui32(int32(94)))
}
func F_ispunct(m *base.Module, l0 int32) int32 {
	var v20 int32
	_ = v20
	if base.Ui32(l0-int32(33)) <= base.Ui32(int32(93)) {
		v20 = base.B2i32(base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(l0|int32(32)-int32(97)) < base.Ui32(int32(26))) == int32(0))
	} else {
		v20 = int32(0)
	}
	return v20
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
	v15 = *(*int32)(unsafe.Add(mBase, _consts[1463]))
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
			v27 = *(*int32)(unsafe.Add(mBase, _consts[1464]))
			v29 = *(*int32)(unsafe.Add(mBase, _consts[1465]))
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
							v65 = *(*int32)(unsafe.Add(mBase, _consts[0]))
							v70 = F_AllocSetContextCreateInternal(m, v65, int32(65352), int32(0), int32(8192), int32(8388608))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v70
								v73 = int32(4562080)
								v74 = *(*int32)(unsafe.Add(mBase, _consts[0]))
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v70
								v78 = F_CreateTemplateTupleDesc(m, int32(2))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = v78
									F_TupleDescInitEntry(m, v78, int32(1), int32(438289), int32(701), int32(-1), int32(0))
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return int32(0)
									} else {
										v88 = *(*int32)(unsafe.Add(mBase, uint32(v32)+32))
										F_TupleDescInitEntry(m, v88, int32(2), int32(453920), int32(27), int32(-1), int32(0))
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
											v115 = *(*int32)(unsafe.Add(mBase, _consts[130]))
											v118 = F_tuplesort_begin_heap(m, v96, v97, v12+int32(30), v12+int32(24), v12+int32(20), v12+int32(19), v115, v101, v101)
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v118
												v121 = *(*int32)(unsafe.Add(mBase, uint32(v32)+32))
												v123 = F_MakeSingleTupleTableSlot(m, v121, int32(1652372))
												mBase = m.M
												v124 = m.ExcPending
												if v124 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v32)+36)) = v123
													v126 = *(*int32)(unsafe.Add(mBase, uint32(v32)+32))
													v128 = F_MakeSingleTupleTableSlot(m, v126, int32(1652476))
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
															v136 = F_pairingheap_allocate(m, int32(7689), v16)
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
																				*(*int32)(unsafe.Add(mBase, _consts[0])) = v74
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
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v67 int64
	_ = v67
	v3 = F_palloc0(m, int32(140))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = v7
		*(*int64)(unsafe.Add(mBase, uint32(v3)+4)) = int64(72057594038255616)
		*(*int32)(unsafe.Add(mBase, uint32(v3))) = int32(438)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+19)) = v7
		v15 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v3)+18)) = uint8(v15)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+108)) = int32(7113)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+104)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+100)) = int32(7114)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+96)) = int32(7117)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+92)) = int32(7109)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+88)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+84)) = int32(7685)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+80)) = int32(7686)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+76)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+72)) = int32(7687)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+68)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+64)) = int32(7688)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+60)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+56)) = int32(7118)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+52)) = int32(7112)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+48)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+44)) = int32(7116)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+40)) = int32(7111)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+36)) = int32(7110)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+32)) = v7
		*(*uint8)(unsafe.Add(mBase, uint32(v3)+29)) = uint8(v15)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+25)) = v15
		*(*uint16)(unsafe.Add(mBase, uint32(v3)+16)) = uint16(v7)
		*(*uint16)(unsafe.Add(mBase, uint32(v3)+23)) = uint16(v7)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+136)) = v7
		v67 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v3)+128)) = v67
		*(*int64)(unsafe.Add(mBase, uint32(v3)+120)) = v67
		*(*int64)(unsafe.Add(mBase, uint32(v3)+112)) = v67
		return v3
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
	v4 = *(*int32)(unsafe.Add(mBase, _consts[1462]))
	v8 = F_build_reloptions(m, l0, l1, v4, int32(8), int32(4120704), int32(1))
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
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
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+16)) = uint8(v7)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+64))
	v10 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(v6)+72)) = v10
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+52))
	if v14 == v10 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l1 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
	if v17 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_pfree(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(0)
	goto L1
L6:
	;
	if l3 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v27 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v32 = v27 * int32(48)
	if v30 == l1 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L6
L10:
	;
	goto L9
L11:
	;
	v36 = v30 + v32
	if base.Ui32(l1-v36) <= base.Ui32(int32(0)-v32<<(uint(int32(1))%32)) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v43 = F___memcpy(m, v30, l1, v32)
	mBase = m.M
	goto L9
L13:
	;
	goto L14
L14:
	;
	v46 = (v30 ^ l1) & int32(3)
	if base.Ui32(v30) < base.Ui32(l1) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	if v148 == int32(0) {
		goto L10
	} else {
		goto L51
	}
L16:
	;
	if base.Ui32(v126) <= base.Ui32(int32(3)) {
		v147 = v125
		v148 = v126
		v149 = v127
		goto L15
	} else {
		goto L47
	}
L17:
	;
	if v46 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	if v46 != 0 {
		v108 = v32
		goto L30
	} else {
		goto L31
	}
L20:
	;
	v147 = l1
	v148 = v32
	v149 = v30
	goto L15
L21:
	;
	goto L22
L22:
	;
	if v30&int32(3) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v125 = l1
	v126 = v32
	v127 = v30
	goto L16
L24:
	;
	goto L25
L25:
	;
	v53 = l1
	v54 = v32
	v55 = v30
	goto L26
L26:
	;
	if v54 == int32(0) {
		goto L10
	} else {
		goto L28
	}
L27:
	;
	v125 = v62
	v126 = v64
	v127 = v66
	goto L16
L28:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v59)
	v61 = int32(1)
	v62 = v53 + v61
	v64 = v54 - v61
	v66 = v55 + v61
	if v66&int32(3) != 0 {
		v53 = v62
		v54 = v64
		v55 = v66
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	if v108 == int32(0) {
		goto L10
	} else {
		goto L43
	}
L31:
	;
	if v36&int32(3) != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v73 = v32
	goto L35
L33:
	;
	v88 = v32
	goto L34
L34:
	;
	if base.Ui32(v88) <= base.Ui32(int32(3)) {
		v108 = v88
		goto L30
	} else {
		goto L39
	}
L35:
	;
	if v73 == int32(0) {
		goto L10
	} else {
		goto L37
	}
L36:
	;
	v88 = v79
	goto L34
L37:
	;
	v79 = v73 - int32(1)
	v80 = v30 + v79
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v79))))
	*(*uint8)(unsafe.Add(mBase, uint32(v80))) = uint8(v82)
	if v80&int32(3) != 0 {
		v73 = v79
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v95 = v88
	goto L40
L40:
	;
	v99 = v95 - int32(4)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1+v99)))
	*(*int32)(unsafe.Add(mBase, uint32(v30+v99))) = v102
	if base.Ui32(int32(3)) < base.Ui32(v99) {
		v95 = v99
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v108 = v99
	goto L30
L42:
	;
	goto L41
L43:
	;
	v115 = v108
	goto L44
L44:
	;
	v119 = v115 - int32(1)
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v119))))
	*(*uint8)(unsafe.Add(mBase, uint32(v30+v119))) = uint8(v122)
	if v119 != 0 {
		v115 = v119
		goto L44
	} else {
		goto L46
	}
L45:
	;
	goto L10
L46:
	;
	goto L45
L47:
	;
	v132 = v125
	v133 = v126
	v134 = v127
	goto L48
L48:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = v136
	v138 = int32(4)
	v139 = v132 + v138
	v141 = v134 + v138
	v143 = v133 - v138
	if base.Ui32(int32(3)) < base.Ui32(v143) {
		v132 = v139
		v133 = v143
		v134 = v141
		goto L48
	} else {
		goto L50
	}
L49:
	;
	v147 = v139
	v148 = v143
	v149 = v141
	goto L15
L50:
	;
	goto L49
L51:
	;
	v154 = v147
	v155 = v148
	v156 = v149
	goto L52
L52:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	*(*uint8)(unsafe.Add(mBase, uint32(v156))) = uint8(v158)
	v160 = int32(1)
	v165 = v155 - v160
	if v165 != 0 {
		v154 = v154 + v160
		v155 = v165
		v156 = v156 + v160
		goto L52
	} else {
		goto L54
	}
L53:
	;
	goto L10
L54:
	;
	goto L53
L55:
	;
	return
L56:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v180 <= int32(0) {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v185 = v180 * int32(48)
	if v183 == l3 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L55
L59:
	;
	goto L58
L60:
	;
	v189 = v183 + v185
	if base.Ui32(l3-v189) <= base.Ui32(int32(0)-v185<<(uint(int32(1))%32)) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v196 = F___memcpy(m, v183, l3, v185)
	mBase = m.M
	goto L58
L62:
	;
	goto L63
L63:
	;
	v199 = (v183 ^ l3) & int32(3)
	if base.Ui32(v183) < base.Ui32(l3) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	if v301 == int32(0) {
		goto L59
	} else {
		goto L100
	}
L65:
	;
	if base.Ui32(v279) <= base.Ui32(int32(3)) {
		v300 = v278
		v301 = v279
		v302 = v280
		goto L64
	} else {
		goto L96
	}
L66:
	;
	if v199 != 0 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L68
L68:
	;
	if v199 != 0 {
		v261 = v185
		goto L79
	} else {
		goto L80
	}
L69:
	;
	v300 = l3
	v301 = v185
	v302 = v183
	goto L64
L70:
	;
	goto L71
L71:
	;
	if v183&int32(3) == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v278 = l3
	v279 = v185
	v280 = v183
	goto L65
L73:
	;
	goto L74
L74:
	;
	v206 = l3
	v207 = v185
	v208 = v183
	goto L75
L75:
	;
	if v207 == int32(0) {
		goto L59
	} else {
		goto L77
	}
L76:
	;
	v278 = v215
	v279 = v217
	v280 = v219
	goto L65
L77:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
	*(*uint8)(unsafe.Add(mBase, uint32(v208))) = uint8(v212)
	v214 = int32(1)
	v215 = v206 + v214
	v217 = v207 - v214
	v219 = v208 + v214
	if v219&int32(3) != 0 {
		v206 = v215
		v207 = v217
		v208 = v219
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	if v261 == int32(0) {
		goto L59
	} else {
		goto L92
	}
L80:
	;
	if v189&int32(3) != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v226 = v185
	goto L84
L82:
	;
	v241 = v185
	goto L83
L83:
	;
	if base.Ui32(v241) <= base.Ui32(int32(3)) {
		v261 = v241
		goto L79
	} else {
		goto L88
	}
L84:
	;
	if v226 == int32(0) {
		goto L59
	} else {
		goto L86
	}
L85:
	;
	v241 = v232
	goto L83
L86:
	;
	v232 = v226 - int32(1)
	v233 = v183 + v232
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v232))))
	*(*uint8)(unsafe.Add(mBase, uint32(v233))) = uint8(v235)
	if v233&int32(3) != 0 {
		v226 = v232
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v248 = v241
	goto L89
L89:
	;
	v252 = v248 - int32(4)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l3+v252)))
	*(*int32)(unsafe.Add(mBase, uint32(v183+v252))) = v255
	if base.Ui32(int32(3)) < base.Ui32(v252) {
		v248 = v252
		goto L89
	} else {
		goto L91
	}
L90:
	;
	v261 = v252
	goto L79
L91:
	;
	goto L90
L92:
	;
	v268 = v261
	goto L93
L93:
	;
	v272 = v268 - int32(1)
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v272))))
	*(*uint8)(unsafe.Add(mBase, uint32(v183+v272))) = uint8(v275)
	if v272 != 0 {
		v268 = v272
		goto L93
	} else {
		goto L95
	}
L94:
	;
	goto L59
L95:
	;
	goto L94
L96:
	;
	v285 = v278
	v286 = v279
	v287 = v280
	goto L97
L97:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	*(*int32)(unsafe.Add(mBase, uint32(v287))) = v289
	v291 = int32(4)
	v292 = v285 + v291
	v294 = v287 + v291
	v296 = v286 - v291
	if base.Ui32(int32(3)) < base.Ui32(v296) {
		v285 = v292
		v286 = v296
		v287 = v294
		goto L97
	} else {
		goto L99
	}
L98:
	;
	v300 = v292
	v301 = v296
	v302 = v294
	goto L64
L99:
	;
	goto L98
L100:
	;
	v307 = v300
	v308 = v301
	v309 = v302
	goto L101
L101:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307))))
	*(*uint8)(unsafe.Add(mBase, uint32(v309))) = uint8(v311)
	v313 = int32(1)
	v318 = v308 - v313
	if v318 != 0 {
		v307 = v307 + v313
		v308 = v318
		v309 = v309 + v313
		goto L101
	} else {
		goto L103
	}
L102:
	;
	goto L59
L103:
	;
	goto L102
}
