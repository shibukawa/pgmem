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
func Fn13826(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	v7 = m.G0
	v9 = v7 - int32(96)
	m.G0 = v9
	F_ScanKeyInit(m, v9, int32(1), int32(3), int32(184), l0)
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_ScanKeyInit(m, v9+int32(48), int32(2), int32(3), int32(184), l1)
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = F_table_open(m, l3, int32(3))
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v29 = F_systable_beginscan(m, v24, l2, int32(1), int32(0), int32(2), v9)
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v31 = F_systable_getnext(m, v29)
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v31 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v33 = v31
	goto L10
L8:
	;
	goto L9
L9:
	;
	F_systable_endscan(m, v29)
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L15
	}
L10:
	;
	F_simple_heap_delete(m, v24, v33+int32(4))
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	v43 = F_systable_getnext(m, v29)
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v43 != 0 {
		v33 = v43
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	F_relation_close(m, v24, int32(3))
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	m.G0 = v9 + int32(96)
	return
}
func Fn13833(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	v10 = int32(0)
	v12 = m.G0
	v14 = v12 + int32(-64)
	m.G0 = v14
	if l1 == v10 {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
		if v18 != 0 {
			v76 = v10
			m.G0 = v14 - int32(-64)
			return v76
		} else {
			v19 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v19
			*(*int64)(unsafe.Add(mBase, uint32(v14)+29)) = v19
			v23 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v14)+60)) = uint8(v23)
			*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = l3
			*(*uint8)(unsafe.Add(mBase, uint32(v14)+52)) = uint8(v23)
			*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = l2
			*(*uint8)(unsafe.Add(mBase, uint32(v14)+44)) = uint8(v23)
			*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = l1
			v32 = int32(3)
			*(*uint16)(unsafe.Add(mBase, uint32(v14)+38)) = uint16(v32)
			*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = l0
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v38 = m.T0[v37].(func(*base.Module, int32) int32)(m, v12+int32(-44))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+36)))
				if l1 == int32(0) {
					if v42&int32(1) != 0 {
						v76 = v38
						m.G0 = v14 - int32(-64)
						return v76
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v14))) = v51
							F_errmsg_internal(m, l8, v14)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_Fn13833_0), l7, l4)
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
					if v42&int32(1) == int32(0) {
						v76 = v38
						m.G0 = v14 - int32(-64)
						return v76
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v66
							F_errmsg_internal(m, l6, v12+int32(-48))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_Fn13833_0), l5, l4)
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
					}
				}
			}
		}
	} else {
		v19 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v19
		*(*int64)(unsafe.Add(mBase, uint32(v14)+29)) = v19
		v23 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+60)) = uint8(v23)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = l3
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+52)) = uint8(v23)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = l2
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+44)) = uint8(v23)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = l1
		v32 = int32(3)
		*(*uint16)(unsafe.Add(mBase, uint32(v14)+38)) = uint16(v32)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = l0
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v38 = m.T0[v37].(func(*base.Module, int32) int32)(m, v12+int32(-44))
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return int32(0)
		} else {
			v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+36)))
			if l1 == int32(0) {
				if v42&int32(1) != 0 {
					v76 = v38
					m.G0 = v14 - int32(-64)
					return v76
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v14))) = v51
						F_errmsg_internal(m, l8, v14)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_Fn13833_0), l7, l4)
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
				if v42&int32(1) == int32(0) {
					v76 = v38
					m.G0 = v14 - int32(-64)
					return v76
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v66
						F_errmsg_internal(m, l6, v12+int32(-48))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_Fn13833_0), l5, l4)
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
				}
			}
		}
	}
}
func Fn13835(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	v8 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = F_SearchSysCache1(m, l6, l0)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v18 + int32(16)
	return v174
L2:
	;
	return int32(0)
L3:
	;
	if v20 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if l1 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+22)))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	v26 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v26)
	v174 = int32(0)
	goto L1
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = l0
	F_errmsg_internal(m, l5, v18)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_Fn13835_0), l4, l3)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	v43 = v39 + v40
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+72))
	if v44 != int32(11) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F_ReleaseCatCache(m, v20)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L2
	} else {
		goto L44
	}
L15:
	;
	v47 = int32(0)
	v49 = *(*int32)(unsafe.Add(mBase, _c_Fn13835[0]))
	if v49 == v47 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L2
	} else {
		goto L32
	}
L18:
	;
	if v88 == int32(0) {
		v157 = v47
		goto L14
	} else {
		goto L31
	}
L19:
	;
	v88 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v56 <= int32(0) {
		v82 = v47
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v88 = v82
	goto L18
L23:
	;
	v59 = int32(0)
	if v59 < v56 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v62 = v56
	goto L26
L25:
	;
	v62 = v59
	goto L26
L26:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v65 = int32(0)
	goto L27
L27:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v63+v65<<(uint(int32(2))%32))))
	v74 = base.B2i32(v73 == v44)
	if v73 == v44 {
		v82 = v74
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v82 = v74
	goto L22
L29:
	;
	v76 = v65 + int32(1)
	if v76 != v62 {
		v65 = v76
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L17
L32:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _c_Fn13835[0]))
	if v96 == int32(0) {
		v148 = v8
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v157 = base.B2i32(l0 == v148)
	goto L14
L34:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v99 <= int32(0) {
		v148 = v8
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_Fn13835[1]))
	v108 = int32(0)
	v115 = v105
	v119 = v99
	goto L36
L36:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122+v108<<(uint(int32(2))%32))))
	if v115 != v126 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v148 = int32(0)
	goto L33
L38:
	;
	v129 = F_GetSysCacheOid(m, l2, v92, v43+int32(8), v126, int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L2
	} else {
		goto L41
	}
L39:
	;
	v134 = v115
	v135 = v119
	goto L40
L40:
	;
	v137 = v108 + int32(1)
	if v137 < v135 {
		v108 = v137
		v115 = v134
		v119 = v135
		goto L36
	} else {
		goto L43
	}
L41:
	;
	if v129 != 0 {
		v148 = v129
		goto L33
	} else {
		goto L42
	}
L42:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v133 = *(*int32)(unsafe.Add(mBase, _c_Fn13835[1]))
	v134 = v133
	v135 = v131
	goto L40
L43:
	;
	goto L37
L44:
	;
	v174 = v157
	goto L1
}
func Fn13839(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
func Fn13844(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
func Fn13855(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
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
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
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
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v14 = v9 + int32(1)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = F_pg_detoast_datum_packed(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			v22 = v20 & int32(1)
			if v22 != 0 {
				v23 = v14
			} else {
				v23 = v9 + int32(4)
			}
			if v20 == int32(1) {
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
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
				if v22 != 0 {
					v50 = int32(base.Ui32(v20)>>(uint(v40)%32)) - v40
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v51 = int32(1)
			v52 = v16 + v51
			v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
			v57 = v55 & v51
			if v57 != 0 {
				v58 = v52
			} else {
				v58 = v16 + int32(4)
			}
			if v55 == int32(1) {
				v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
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
				if v57 != 0 {
					v85 = int32(base.Ui32(v55)>>(uint(v75)%32)) - v75
				} else {
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					v85 = int32(base.Ui32(v79)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v86 = F_dotrim(m, v23, v50, v58, v85, l2, l1)
			mBase = m.M
			v87 = m.ExcPending
			if v87 != 0 {
				return int32(0)
			} else {
				return v86
			}
		}
	}
}
func Fn13862(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v38 int64
	_ = v38
	var v45 int64
	_ = v45
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	v14 = int64(63)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	v24 = int64(32)
	v25 = int64(base.Ui64(v17) >> (uint(v24) % 64))
	v27 = int64(base.Ui64(v13) >> (uint(v24) % 64))
	v30 = int64(4294967295)
	v31 = v17 & v30
	v33 = v13 & v30
	v34 = v31 * v33
	v38 = int64(base.Ui64(v34)>>(uint(v24)%64)) + v31*v27
	v45 = v33*v25 + v38&v30
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v13*(v17>>(uint(v14)%64)) + v13>>(uint(v14)%64)*v17 + v25*v27 + int64(base.Ui64(v38)>>(uint(v24)%64)) + int64(base.Ui64(v45)>>(uint(v24)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v34&v30 | v45<<(uint(v24)%64)
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	if v56 != v57>>(uint(int64(63))%64) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v66 = m.ExcPending
		if v66 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, l4, int32(0))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, l3, l2, l1)
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
		}
	} else {
		v75 = F_Int64GetDatum(m, v57)
		mBase = m.M
		v76 = m.ExcPending
		if v76 != 0 {
			return int32(0)
		} else {
			m.G0 = v10 + int32(16)
			return v75
		}
	}
}
func Fn13864(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) {
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
		v20 = F_strstr(m, l0, int32(_a_Fn13864_0))
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
						F_errdetail(m, l8, int32(0))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_Fn13864_1), l7, l1)
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
							F_errdetail(m, l6, int32(0))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_Fn13864_1), l5, l1)
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
								F_errdetail(m, l6, int32(0))
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_Fn13864_1), l5, l1)
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
					v31 = Fn13878(m, l0, int32(47))
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
									F_errdetail(m, l3, int32(0))
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_Fn13864_1), l2, l1)
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
					F_errdetail(m, l10, int32(0))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_Fn13864_1), l9, l1)
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
func Fn13868(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 float64, l5 float64) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v17 int32
	_ = v17
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
	v10 = base.F64_nearest(v9)
	v17 = int32(0)
	if base.B2i32(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v10)&int64(9223372036854775807)))|base.B2i32(base.F64_ge(v10, l5) == v17) == v17)&base.F64_lt(v10, l4) == v17 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, l3, int32(0))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_Fn13868_0), l2, l1)
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
	} else {
		return base.I32_trunc_sat_f64_s(v10)
	}
}
func Fn13875(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	v6 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v6)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+20)))
	if v16&int32(1) == v6 {
		v21 = int32(4)
		v25 = l2 + l1<<(uint(v21)%32) + v21
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
		if v26 < int32(0) {
			v68 = F_nocachegetattr(m, l0, l1, l2)
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return int32(0)
			} else {
				v74 = v68
				m.G0 = v11 + int32(16)
				return v74
			}
		} else {
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+22)))
			v31 = v15 + v29 + v26
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+6)))
			if v32 != int32(1) {
				v74 = v31
				m.G0 = v11 + int32(16)
				return v74
			} else {
				v35 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+4)))
				switch v35&int32(_a_Fn13875_0) - int32(1) {
				case 0:
					v40 = int32(*(*int8)(unsafe.Add(mBase, uint32(v31))))
					v74 = v40
					m.G0 = v11 + int32(16)
					return v74
				case 1:
					v41 = int32(*(*int16)(unsafe.Add(mBase, uint32(v31))))
					v74 = v41
					m.G0 = v11 + int32(16)
					return v74
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v35
						F_errmsg_internal(m, int32(_a_Fn13875_1), v11)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, l4, int32(70), int32(_a_Fn13875_2))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 3:
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
					v74 = v42
					m.G0 = v11 + int32(16)
					return v74
				}
			}
		}
	} else {
		v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+23)))
		v58 = int32(1)
		if int32(base.Ui32(v57)>>(uint(l1-v58)%32))&v58 != 0 {
			v68 = F_nocachegetattr(m, l0, l1, l2)
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return int32(0)
			} else {
				v74 = v68
				m.G0 = v11 + int32(16)
				return v74
			}
		} else {
			v63 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v63)
			v74 = int32(0)
			m.G0 = v11 + int32(16)
			return v74
		}
	}
}
func Fn13879(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 float32, l5 float32) int32 {
	mBase := m.M
	_ = mBase
	var v8 float32
	_ = v8
	var v9 float32
	_ = v9
	var v16 int32
	_ = v16
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	v8 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = base.F32_nearest(v8)
	v16 = int32(0)
	if base.B2i32(base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(base.I32_reinterpret_f32(v9)&int32(2147483647)))|base.B2i32(base.F32_ge(v9, l5) == v16) == v16)&base.F32_lt(v9, l4) == v16 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, l3, int32(0))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_Fn13879_0), l2, l1)
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
		return base.I32_trunc_sat_f32_s(v9)
	}
}
func Fn13880(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v11 int32
	_ = v11
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = F_gbt_var_picksplit(m, v4, v5, v6, l1, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func Fn13891(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v9 = F_DirectFunctionCall2Coll(m, l4, int32(0), v7, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 != 0 {
			v18 = v9
			return v18
		} else {
			v16 = F_DirectFunctionCall2Coll(m, l4, int32(0), v7+l3, v8+l3)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = v16
				return v18
			}
		}
	}
}
func Fn13903(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	F_DeconstructQualifiedName(m, l0, v13+int32(12), v13+int32(8))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v23 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	m.G0 = v13 + int32(16)
	return v118
L4:
	;
	if l1|v93 != 0 {
		v118 = v93
		goto L3
	} else {
		goto L26
	}
L5:
	;
	v93 = int32(0)
	goto L4
L6:
	;
	v25 = F_LookupExplicitNamespace(m, v23, l1)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_recomputeNamespacePath(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L15
	}
L9:
	;
	if v25 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v27 = int32(0)
	goto L12
L11:
	;
	v27 = l1
	goto L12
L12:
	;
	if v27 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v29 = int32(0)
	v31 = F_GetSysCacheOid(m, l5, v28, v25, v29, v29)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v93 = v31
	goto L4
L15:
	;
	v35 = int32(0)
	v37 = *(*int32)(unsafe.Add(mBase, _c_Fn13903[0]))
	if v37 == v35 {
		v93 = v35
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v40 = int32(0)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v41 <= v40 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v52 = v40
	goto L18
L18:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55+v52<<(uint(int32(2))%32))))
	v61 = *(*int32)(unsafe.Add(mBase, _c_Fn13903[1]))
	if v59 != v61 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L5
L20:
	;
	v63 = int32(0)
	v65 = F_GetSysCacheOid(m, l5, v44, v59, v63, v63)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v69 = v52 + int32(1)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v69 < v70 {
		v52 = v69
		goto L18
	} else {
		goto L25
	}
L23:
	;
	if v65 != 0 {
		v118 = v65
		goto L3
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	goto L19
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v102 = F_NameListToString(m, l0)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v102
	F_errmsg(m, l4, v13)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_Fn13903_0), l3, l2)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func Fn13905(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v21 = F_palloc(m, int32(4))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		v26 = F_palloc(m, int32(16))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v28 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v19))) = v28
			v31 = F_palloc(m, v28)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = v31
				v34 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v34)
				*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = l2
				*(*uint8)(unsafe.Add(mBase, uint32(v26)+8)) = uint8(v34)
				*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v17
				*(*uint16)(unsafe.Add(mBase, uint32(v26))) = uint16(v16)
				v42 = F_palloc(m, int32(4))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = v42
					*(*int32)(unsafe.Add(mBase, uint32(v42))) = v26
					v47 = v16 & int32(_a_Fn13905_0)
					switch v47 - int32(1) {
					case 0, 1:
						*(*int32)(unsafe.Add(mBase, uint32(v21))) = l1
						v67 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v67)
						m.G0 = v13 + int32(16)
						return v21
					case 2:
						*(*int32)(unsafe.Add(mBase, uint32(v21))) = v17
						m.G0 = v13 + int32(16)
						return v21
					case 3, 4:
						v50 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v50)
						*(*int32)(unsafe.Add(mBase, uint32(v21))) = v17
						m.G0 = v13 + int32(16)
						return v21
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v13))) = v47
							F_errmsg_internal(m, int32(_a_Fn13905_1), v13)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_Fn13905_2), int32(97), int32(_a_Fn13905_3))
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
					}
				}
			}
		}
	}
}
func Fn13909(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
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
		v16 = *(*int32)(unsafe.Add(mBase, _c_Fn13909[0]))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v16+(l1^int32(-1))<<(uint(int32(2))%32))))
		v30 = v22
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, _c_Fn13909[1]))
		v30 = v24 + l1<<(uint(int32(13))%32) + int32(-8192)
	}
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+14)))
	if v31 != 0 {
		v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+19)))
		v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+16)))
		if (v32<<(uint(int32(8))%32)-v35)&int32(_a_Fn13909_0) != int32(16) {
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
						v95 = *(*int32)(unsafe.Add(mBase, _c_Fn13909[2]))
						v101 = *(*int32)(unsafe.Add(mBase, uint32(v95+(l1^int32(-1))<<(uint(int32(6))%32))+16))
						v110 = v101
					} else {
						v103 = *(*int32)(unsafe.Add(mBase, _c_Fn13909[3]))
						v109 = *(*int32)(unsafe.Add(mBase, uint32(v103+l1<<(uint(int32(6))%32)+int32(-64))+16))
						v110 = v109
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v110
					*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v91 + int32(4)
					F_errmsg(m, int32(_a_Fn13909_1), v11+int32(16))
					mBase = m.M
					v119 = m.ExcPending
					if v119 != 0 {
						return
					} else {
						F_errhint(m, int32(_a_Fn13909_2), int32(0))
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
					v55 = *(*int32)(unsafe.Add(mBase, _c_Fn13909[2]))
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v55+(l1^int32(-1))<<(uint(int32(6))%32))+16))
					v70 = v61
				} else {
					v63 = *(*int32)(unsafe.Add(mBase, _c_Fn13909[3]))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v63+l1<<(uint(int32(6))%32)+int32(-64))+16))
					v70 = v69
				}
				*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v70
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v51 + int32(4)
				F_errmsg(m, int32(_a_Fn13909_3), v11)
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return
				} else {
					F_errhint(m, int32(_a_Fn13909_2), int32(0))
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
func Fn13914(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v35 int32
	_ = v35
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v20)
		v22 = F_convert_any_priv_string(m, v16, l1)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v26 = F_object_aclcheck_ext(m, l2, v13, v14, v22, v11+int32(15))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
				if v28 == int32(1) {
					v31 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
					v35 = int32(0)
				} else {
					v35 = base.B2i32(v26 == int32(0))
				}
				m.G0 = v11 + int32(16)
				return v35
			}
		}
	}
}
func Fn13918(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
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
	var v26 int32
	_ = v26
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_pg_detoast_datum_packed(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v24 = F_pg_detoast_datum_packed(m, v23)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v26 = F_get_role_oid_or_public(m, v17)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v29 = F_text_to_cstring(m, v19)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v31 = F_DirectFunctionCall1Coll(m, l7, int32(0), v29)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						if v31 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								F_errcode(m, l6)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v15))) = v29
									F_errmsg(m, l5, v15)
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_Fn13918_0), l4, l3)
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
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
							v47 = F_convert_any_priv_string(m, v24, l1)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								v49 = F_object_aclcheck(m, l2, v31, v26, v47)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									m.G0 = v15 + int32(16)
									return base.B2i32(v49 == int32(0))
								}
							}
						}
					}
				}
			}
		}
	}
}
func Fn13921(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
func Fn13929(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v11, v12, v13, l2, int32(7))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v21 = F_latin2mic_with_table(m, v9, v8, v13, l3, l2, l1, base.B2i32(v10 != int32(0)))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			return v21
		}
	}
}
func Fn13930(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v14
		v20 = F_get_worker(m, v10, int32(0), v7+int32(12), int32(1), l1)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			if v20 == int32(0) {
				v24 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v24)
				v27 = int32(0)
			} else {
				v27 = v20
			}
			m.G0 = v7 + int32(16)
			return v27
		}
	}
}
func Fn13941(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
	v8 = int32(8)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v11 = int32(16)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+2)))
	v15 = v7<<(uint(v8)%32) | v10<<(uint(v11)%32) | v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+2)))
	v25 = v17<<(uint(v8)%32) | v20<<(uint(v11)%32) | v24
	if base.Ui32(v15) < base.Ui32(v25) {
		v51 = l1
	} else {
		if base.Ui32(v25) < base.Ui32(v15) {
			v51 = int32(1)
		} else {
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+5)))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+4)))
			v31 = int32(8)
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+3)))
			v34 = int32(16)
			v37 = v29 | (v30<<(uint(v31)%32) | v33<<(uint(v34)%32))
			v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+5)))
			v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+4)))
			v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+3)))
			v46 = v38 | (v39<<(uint(v31)%32) | v42<<(uint(v34)%32))
			if base.Ui32(v37) < base.Ui32(v46) {
				v51 = l1
			} else {
				v51 = base.B2i32(base.Ui32(v46) < base.Ui32(v37))
			}
		}
	}
	return v51
}
func Fn13947(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v14 int32
	_ = v14
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = F_strlen(m, v6)
		mBase = m.M
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v14 = F_RE_compile_and_cache(m, v8, l1, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v20 = F_palloc(m, v12<<(uint(int32(2))%32)+int32(4))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = F_pg_mb2wchar_with_len(m, v6, v20, v12)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = int32(0)
					v27 = F_RE_wchar_execute(m, v20, v22, v24, v24, v24)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						F_pfree(m, v20)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							return v27
						}
					}
				}
			}
		}
	}
}
func Fn13949(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	v2 = l1
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v10 == int32(0) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v13 != 0 {
			v64 = v13
			v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v66 == int32(0) {
				v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v70 = F_pg_detoast_datum(m, v69)
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					F_do_numeric_accum(m, v64, v70)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(16)
						return v64
					}
				}
			} else {
				m.G0 = v8 + int32(16)
				return v64
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
				switch v21 - int32(429) {
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
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)+368))
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
				v81 = m.ExcPending
				if v81 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_Fn13949_0), int32(0))
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_Fn13949_1), int32(_a_Fn13949_2), int32(_a_Fn13949_3))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v49 = int32(_a_Fn13949_4)
				v50 = *(*int32)(unsafe.Add(mBase, _c_Fn13949[0]))
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
				*(*int32)(unsafe.Add(mBase, _c_Fn13949[0])) = v52
				v55 = F_palloc0(m, int32(112))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v2)
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v60
					*(*int32)(unsafe.Add(mBase, _c_Fn13949[0])) = v50
					v64 = v55
					v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
					if v66 == int32(0) {
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v70 = F_pg_detoast_datum(m, v69)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							F_do_numeric_accum(m, v64, v70)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(16)
								return v64
							}
						}
					} else {
						m.G0 = v8 + int32(16)
						return v64
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
			switch v21 - int32(429) {
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
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)+368))
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
			v81 = m.ExcPending
			if v81 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_Fn13949_0), int32(0))
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_Fn13949_1), int32(_a_Fn13949_2), int32(_a_Fn13949_3))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v49 = int32(_a_Fn13949_4)
			v50 = *(*int32)(unsafe.Add(mBase, _c_Fn13949[0]))
			v52 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			*(*int32)(unsafe.Add(mBase, _c_Fn13949[0])) = v52
			v55 = F_palloc0(m, int32(112))
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int32(0)
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v2)
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v60
				*(*int32)(unsafe.Add(mBase, _c_Fn13949[0])) = v50
				v64 = v55
				v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
				if v66 == int32(0) {
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v70 = F_pg_detoast_datum(m, v69)
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						F_do_numeric_accum(m, v64, v70)
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(16)
							return v64
						}
					}
				} else {
					m.G0 = v8 + int32(16)
					return v64
				}
			}
		}
	}
}
func Fn13950(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v26 int32
	_ = v26
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v10 != 0 {
		v13 = int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v13 = v12
	}
	v16 = F_numeric_poly_stddev_internal(m, v13, l2, l1, v8+int32(15))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
		if v20 == int32(1) {
			v23 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v23)
			v26 = int32(0)
		} else {
			v26 = v16
		}
		m.G0 = v8 + int32(16)
		return v26
	}
}
func Fn13958(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(34209793)
	*(*uint32)(unsafe.Add(mBase, uint32(v7)+8)) = uint32(v10)
	v15 = int64(base.Ui64(v10) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v7)+4)) = uint32(v15)
	v18 = *(*int32)(unsafe.Add(mBase, _c_Fn13958[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v18
	v21 = F_LockRelease(m, v7, l1, int32(1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v21
	}
}
func Fn13963(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)))
		if v20 == int32(_a_Fn13963_0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, l4, int32(0))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_Fn13963_1), l3, l2)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
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
			*(*int64)(unsafe.Add(mBase, uint32(v11))) = v14
			v38 = v11 + int32(16)
			v41 = F_pg_snprintf(m, v38, int32(32), int32(_a_Fn13963_2), v11)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				v44 = int32(0)
				v50 = F_DirectFunctionCall3Coll(m, int32(408), v44, v38, v44, int32(-1))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					v52 = F_DirectFunctionCall2Coll(m, l1, v44, v50, v16)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						v54 = F_DirectFunctionCall1Coll(m, int32(1465), v44, v52)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							m.G0 = v11 + int32(48)
							return v54
						}
					}
				}
			}
		}
	}
}
func Fn13969(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = l5
	v20 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_Fn13969[0]))))
	return int32(base.Ui32(v47&l2) >> (uint(l1) % 32))
L4:
	;
	v25 = base.I32_div_s(v19+v20, int32(2))
	v27 = v25 << (uint(int32(3)) % 32)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v27+l4)))
	if base.Ui32(v29) < base.Ui32(l0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	return int32(0)
L6:
	;
	if v41 <= v40 {
		v19 = v40
		v20 = v41
		goto L4
	} else {
		goto L13
	}
L7:
	;
	v40 = v19
	v41 = v25 + int32(1)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v27+l3)))
	if base.Ui32(v34) <= base.Ui32(l0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(1)
L11:
	;
	goto L12
L12:
	;
	v40 = v25 - int32(1)
	v41 = v20
	goto L6
L13:
	;
	goto L5
}
func Fn13974(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 float64
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 float64
	_ = v146
	var v150 int64
	_ = v150
	var v155 float64
	_ = v155
	var v156 float64
	_ = v156
	var v157 float64
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int64
	_ = v162
	var v163 int32
	_ = v163
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	v3 = int32(0)
	v13 = float64(0)
	v16 = m.G0
	v18 = v16 + int32(-64)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = F_palloc0(m, v20<<(uint(int32(2))%32))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return
	} else {
		v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		if base.B2i32(v25 == int64(0)) == int32(0) {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v36 = v3
			v40 = v3
			v42 = v3
			for {
				v49 = v31 + v36<<(uint(int32(3))%32)
				v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+4)))
				if v50 == int32(1) {
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
					v55 = int32(16)
					v59 = (int32(base.Ui32(v54)>>(uint(v55)%32)) ^ v54) * int32(-2048144789)
					v64 = (int32(base.Ui32(v59)>>(uint(int32(13))%32)) ^ v59) * int32(-1028477387)
					v68 = v53 & (int32(base.Ui32(v64)>>(uint(v55)%32)) ^ v64)
					v71 = v23 + v68<<(uint(int32(2))%32)
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
					*(*int32)(unsafe.Add(mBase, uint32(v71))) = v72 + int32(1)
					if base.Ui32(v36) < base.Ui32(v68) {
						v79 = base.I32_wrap_i64(v25)
					} else {
						v79 = int32(0)
					}
					v80 = v36 - v68 + v79
					if base.Ui32(v42) < base.Ui32(v80) {
						v82 = v80
					} else {
						v82 = v42
					}
					v85 = v80 + v40
					v87 = v82
				} else {
					v85 = v40
					v87 = v42
				}
				v89 = v36 + int32(1)
				if base.Ui64(base.I64_extend_i32_u(v89)) < base.Ui64(v25) {
					v36 = v89
					v40 = v85
					v42 = v87
					continue
				} else {
					break
				}
				break
			}
			v92 = int32(0)
			v97 = v92
			v99 = v92
			v101 = v92
			for {
				v113 = *(*int32)(unsafe.Add(mBase, uint32(v23+v101<<(uint(int32(2))%32))))
				v115 = v113 - int32(1)
				if base.Ui32(v99) < base.Ui32(v115) {
					v117 = v115
				} else {
					v117 = v99
				}
				if v113 != 0 {
					v118 = v117
				} else {
					v118 = v99
				}
				v122 = v113 - base.B2i32(v113 != int32(0)) + v97
				v124 = v101 + int32(1)
				if base.Ui64(base.I64_extend_i32_u(v124)) < base.Ui64(v25) {
					v97 = v122
					v99 = v118
					v101 = v124
					continue
				} else {
					break
				}
				break
			}
			v129 = v122
			v131 = v118
			v135 = v85
			v137 = v87
		} else {
			v129 = v3
			v131 = v3
			v135 = v3
			v137 = v3
		}
		F_pfree(m, v23)
		mBase = m.M
		v143 = m.ExcPending
		if v143 != 0 {
			return
		} else {
			v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v144 != 0 {
				v146 = base.F64_convert_i32_u(v144)
				v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				v155 = base.F64_div(base.F64_convert_i32_u(v129), v146)
				v156 = base.F64_div(base.F64_convert_i32_u(v135), v146)
				v157 = base.F64_div(v146, base.F64_convert_i64_u(v150))
			} else {
				v155 = v13
				v156 = v13
				v157 = float64(0)
			}
			v160 = F_errstart(m, int32(15), int32(0))
			mBase = m.M
			v161 = m.ExcPending
			if v161 != 0 {
				return
			} else {
				if v160 != 0 {
					v162 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
					v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*float64)(unsafe.Add(mBase, uint32(v18)+48)) = v155
					*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v131
					*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v129
					*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v156
					*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v137
					*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v135
					*(*float64)(unsafe.Add(mBase, uint32(v18)+16)) = v157
					*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v163
					*(*int64)(unsafe.Add(mBase, uint32(v18))) = v162
					F_errmsg_internal(m, int32(_a_Fn13974_0), v18)
					mBase = m.M
					v175 = m.ExcPending
					if v175 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_Fn13974_1), int32(1144), l1)
						mBase = m.M
						v179 = m.ExcPending
						if v179 != 0 {
							return
						} else {
							m.G0 = v18 - int32(-64)
							return
						}
					}
				} else {
					m.G0 = v18 - int32(-64)
					return
				}
			}
		}
	}
}
func Fn13983(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	v5 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v7
	v10 = v7 + int32(1)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v11 <= v10 {
		v66 = v5
		return v66
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v10))))
		if v15 != int32(101) {
			v66 = v5
			return v66
		} else {
			v19 = F_find_among(m, l0, l3, int32(6))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				if v19 == int32(0) {
					v66 = v5
					return v66
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v25
					v27 = int32(1)
					switch v19 - v27 {
					case 0:
						v30 = F_slice_del(m, l0)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							if v30 < int32(0) {
								v66 = v30
							} else {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(2)
								v58 = v34
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
								v61 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v60 - v61
								v66 = v61
							}
							return v66
						}
					case 1:
						v38 = F_slice_from_s(m, l0, int32(4), l2)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							if v38 < int32(0) {
								v66 = v38
							} else {
								v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								v58 = v42
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
								v61 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v60 - v61
								v66 = v61
							}
							return v66
						}
					case 2:
						v43 = F_slice_del(m, l0)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							if v43 < int32(0) {
								v66 = v43
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v47))) = int32(4)
								v58 = v47
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
								v61 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v60 - v61
								v66 = v61
							}
							return v66
						}
					case 3:
						v51 = F_slice_from_s(m, l0, int32(4), l1)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							if v51 < int32(0) {
								v66 = v51
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v55))) = int32(4)
								v58 = v55
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
								v61 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v60 - v61
								v66 = v61
							}
							return v66
						}
					default:
						v66 = v27
						return v66
					}
				}
			}
		}
	}
}
func Fn13985(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v12 = int32(1)
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4+int32(base.Ui32(v8)>>(uint(int32(2))%32))-v12))))
		return int32(base.Ui32(v14)>>(uint(l1)%32)) & v12
	}
}
func Fn13994(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	v4 = l3
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_copy(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
		if v17 != 0 {
			v18 = F_array_contains_nulls(m, v13)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v18 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67108994))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_Fn13994_0), int32(0))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_Fn13994_1), l2, l1)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
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
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v23 = F_ArrayGetNItemsSafe(m, v20, v13+int32(16))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v4)
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
						if v26 != 0 {
							v34 = v26
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
							v34 = (v27<<(uint(int32(3))%32) + int32(23)) & int32(-8)
						}
						F_isort(m, v34+v13, v23, v10+int32(15))
						mBase = m.M
						m.G0 = v10 + int32(16)
						return v13
					}
				}
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			v23 = F_ArrayGetNItemsSafe(m, v20, v13+int32(16))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v4)
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
				if v26 != 0 {
					v34 = v26
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v34 = (v27<<(uint(int32(3))%32) + int32(23)) & int32(-8)
				}
				F_isort(m, v34+v13, v23, v10+int32(15))
				mBase = m.M
				m.G0 = v10 + int32(16)
				return v13
			}
		}
	}
}
func Fn14001(m *base.Module, l0 int32, l1 int32) int32 {
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
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v15 = F_ArrayGetIntegerTypmods(m, v9, v6+int32(12))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			if v17 != int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_Fn14001_0), int32(0))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_Fn14001_1), int32(65), int32(_a_Fn14001_2))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
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
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				v37 = F_anytime_typmod_check(m, l1, v36)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					m.G0 = v6 + int32(16)
					return v37
				}
			}
		}
	}
}
func Fn14007(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
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
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(440)
		*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = l7
		*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = l7
		v27 = F_list_make1_impl(m, int32(472), v13+int32(8))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
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
			return v16
		}
	}
}
func Fn14010(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	F_errmsg(m, int32(_a_Fn14010_0), int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	F_errfinish(m, int32(_a_Fn14010_1), l2, l1)
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
	return v144
L10:
	;
	v144 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_Fn14010[0]))
	if v35 == v24 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v144 = int32(1)
	goto L9
L14:
	;
	goto L15
L15:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_Fn14010[1]))
	if v39 <= int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v144 = v136
	goto L9
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_Fn14010[2]))
	if v43 == int32(0) {
		v136 = int32(0)
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_Fn14010[3]))
	v107 = int32(0)
	v109 = v39 - int32(1)
	goto L39
L20:
	;
	v48 = v43
	goto L21
L21:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	if v53 == int32(4) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v136 = int32(0)
	goto L16
L23:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v48)+80))
	if v100 != 0 {
		v48 = v100
		goto L21
	} else {
		goto L38
	}
L24:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v56 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v59 = int32(1)
	if v24 == v56 {
		v136 = v59
		goto L16
	} else {
		goto L26
	}
L26:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
	v63 = v61 - int32(1)
	if v63 < int32(0) {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v68 = int32(0)
	v70 = v63
	goto L28
L28:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v76 = int32(2)
	v77 = base.I32_div_s(v70-v68, v76)
	v78 = v77 + v68
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v74+v78<<(uint(v76)%32))))
	if v82 == v24 {
		v136 = v59
		goto L16
	} else {
		goto L30
	}
L29:
	;
	goto L23
L30:
	;
	v86 = F_TransactionIdPrecedes(m, v82, v24)
	mBase = m.M
	if v86 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v87 = v78 + int32(1)
	goto L33
L32:
	;
	v87 = v68
	goto L33
L33:
	;
	if v86 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v90 = v70
	goto L36
L35:
	;
	v90 = v78 - int32(1)
	goto L36
L36:
	;
	if v87 <= v90 {
		v68 = v87
		v70 = v90
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
	v114 = int32(2)
	v115 = base.I32_div_s(v109-v107, v114)
	v116 = v115 + v107
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v105+v116<<(uint(v114)%32))))
	v121 = base.B2i32(v120 == v24)
	if v120 == v24 {
		v136 = v121
		goto L16
	} else {
		goto L41
	}
L40:
	;
	v136 = v121
	goto L16
L41:
	;
	v124 = base.B2i32(base.Ui32(v120) < base.Ui32(v24))
	if base.Ui32(v120) < base.Ui32(v24) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v125 = v116 + int32(1)
	goto L44
L43:
	;
	v125 = v107
	goto L44
L44:
	;
	if base.Ui32(v120) < base.Ui32(v24) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v128 = v109
	goto L47
L46:
	;
	v128 = v116 - int32(1)
	goto L47
L47:
	;
	if v125 <= v128 {
		v107 = v125
		v109 = v128
		goto L39
	} else {
		goto L48
	}
L48:
	;
	goto L40
}
func Fn14023(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 float32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 float64
	_ = v98
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v18 = v11 + int32(1)
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v20 = F_pg_detoast_datum_packed(m, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
			v24 = v22 & int32(1)
			if v24 != 0 {
				v25 = v18
			} else {
				v25 = v11 + int32(4)
			}
			if v22 == int32(1) {
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
				if v31 == int32(18) {
					v34 = int32(16)
				} else {
					v34 = int32(0)
				}
				if base.Ui32((v31-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v41 = int32(4)
				} else {
					v41 = v34
				}
				v52 = v41
			} else {
				v42 = int32(1)
				if v24 != 0 {
					v52 = int32(base.Ui32(v22)>>(uint(v42)%32)) - v42
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v52 = int32(base.Ui32(v46)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v53 = int32(1)
			v54 = v20 + v53
			v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
			v59 = v57 & v53
			if v59 != 0 {
				v60 = v54
			} else {
				v60 = v20 + int32(4)
			}
			if v57 == int32(1) {
				v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
				if v66 == int32(18) {
					v69 = int32(16)
				} else {
					v69 = int32(0)
				}
				if base.Ui32((v66-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v76 = int32(4)
				} else {
					v76 = v69
				}
				v87 = v76
			} else {
				v77 = int32(1)
				if v59 != 0 {
					v87 = int32(base.Ui32(v57)>>(uint(v77)%32)) - v77
				} else {
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
					v87 = int32(base.Ui32(v81)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v88 = F_calc_word_similarity(m, v25, v52, v60, v87, l2)
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return int32(0)
			} else {
				v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v90 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return int32(0)
					} else {
						v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v94 != v20 {
							F_pfree(m, v20)
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int32(0)
							} else {
								v98 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
								return base.F64_le(v98, base.F64_promote_f32(v88))
							}
						} else {
							v98 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
							return base.F64_le(v98, base.F64_promote_f32(v88))
						}
					}
				} else {
					v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v94 != v20 {
						F_pfree(m, v20)
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int32(0)
						} else {
							v98 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
							return base.F64_le(v98, base.F64_promote_f32(v88))
						}
					} else {
						v98 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
						return base.F64_le(v98, base.F64_promote_f32(v88))
					}
				}
			}
		}
	}
}
