package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ChangeToDataDir(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_ChangeToDataDir[0]))
	v8 = m.Env.X__syscall_chdir(m, v7)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v8) {
		*(*int32)(unsafe.Add(mBase, _c_F_ChangeToDataDir[1])) = int32(0) - v8
		v16 = int32(-1)
	} else {
		v16 = v8
	}
	if v16 < int32(0) {
		F_errstart_cold(m, int32(22), int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			F_errcode_for_file_access(m)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, _c_F_ChangeToDataDir[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v4))) = v26
				F_errmsg(m, int32(_a_F_ChangeToDataDir_0), v4)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_ChangeToDataDir_1), int32(418), int32(_a_F_ChangeToDataDir_2))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
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
		m.G0 = v4 + int32(16)
		return
	}
}
func F_ChangeVarNodes(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l1
	if l0 == v4 {
		v73 = F_ChangeVarNodes_walker(m, l0, v10+int32(4))
		mBase = m.M
		v74 = m.ExcPending
		if v74 != 0 {
			return
		} else {
			m.G0 = v10 + int32(16)
			return
		}
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v18 != int32(67) {
			v73 = F_ChangeVarNodes_walker(m, l0, v10+int32(4))
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return
			} else {
				m.G0 = v10 + int32(16)
				return
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			if l1 == v21 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l2
			} else {
			}
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
			if l1 == v24 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = l2
			} else {
			}
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
			if v27 == int32(0) {
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+32))
				if v30 != l1 {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = l2
				}
			}
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
			if v33 == int32(0) {
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
				if v36 <= int32(0) {
				} else {
					v44 = v4
					for {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+v44<<(uint(int32(2))%32))))
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
						if l1 == v51 {
							*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = l2
						} else {
						}
						v55 = v44 + int32(1)
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
						if v55 < v56 {
							v44 = v55
							continue
						} else {
							break
						}
						break
					}
				}
			}
			v69 = F_query_tree_walker_impl(m, l0, int32(1127), v10+int32(4), int32(0))
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return
			} else {
				m.G0 = v10 + int32(16)
				return
			}
		}
	}
}
func F_changeDependenciesOn(m *base.Module, l0 int32, l1 int32) {
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
	var v17 int32
	_ = v17
	var v32 int32
	_ = v32
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	v9 = m.G0
	v11 = v9 - int32(128)
	m.G0 = v11
	v15 = F_table_open(m, int32(2608), int32(3))
	mBase = m.M
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
	v17 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(1259)
	v32 = int32(1)
	goto L3
L3:
	;
	if base.B2i32(int32(0)|base.B2i32(base.Ui32(int32(_a_F_changeDependenciesOn_0)) < base.Ui32(l0)) == v17)&((v32|base.B2i32(l0 != int32(2200)))&v32) == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v52 = int32(1)
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L26
	}
L7:
	;
	v61 = v11 + int32(16)
	F_ScanKeyInit(m, v61, int32(4), int32(3), int32(184), int64(1259))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	F_ScanKeyInit(m, v11+int32(72), int32(5), int32(3), int32(184), base.I64_extend_i32_u(l0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v80 = F_systable_beginscan(m, v15, int32(2674), int32(1), int32(0), int32(2), v61)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	goto L11
L11:
	;
	v90 = F_systable_getnext(m, v80)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	F_systable_endscan(m, v80)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L24
	}
L13:
	;
	if v90 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if base.B2i32(int32(0)|base.B2i32(base.Ui32(int32(_a_F_changeDependenciesOn_0)) < base.Ui32(l1)) == int32(0))&((v52|base.B2i32(l1 != int32(2200)))&v52) != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	goto L12
L17:
	;
	F_simple_heap_delete(m, v15, v90+int32(4))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v98 = F_heap_copytuple(m, v90)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L11
L21:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v98)+16))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v100+v101)+16)) = l1
	F_CatalogTupleUpdate(m, v15, v98+int32(4), v98)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_pfree(m, v98)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	goto L11
L24:
	;
	F_relation_close(m, v15, int32(3))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	m.G0 = v11 + int32(128)
	return
L26:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v130 = F_getObjectDescription(m, v11+int32(4), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v130
	F_errmsg(m, int32(_a_F_changeDependenciesOn_1), v11)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_changeDependenciesOn_2), int32(668), int32(_a_F_changeDependenciesOn_3))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
