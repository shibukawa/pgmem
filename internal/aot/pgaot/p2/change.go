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
					F_errfinish(m, int32(_a_F_ChangeToDataDir_1), int32(468), int32(_a_F_ChangeToDataDir_2))
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
	var v6 int32
	_ = v6
	F_ChangeVarNodesExtended(m, l0, l1, l2, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
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
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	v9 = m.G0
	v11 = v9 - int32(112)
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
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L26
	}
L7:
	;
	v61 = v11 + int32(16)
	F_ScanKeyInit(m, v61, int32(4), int32(3), int32(184), int32(1259))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	F_ScanKeyInit(m, v11-int32(-64), int32(5), int32(3), int32(184), l0)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v79 = F_systable_beginscan(m, v15, int32(2674), int32(1), int32(0), int32(2), v61)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	goto L11
L11:
	;
	v89 = F_systable_getnext(m, v79)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	F_systable_endscan(m, v79)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L24
	}
L13:
	;
	if v89 != 0 {
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
	F_simple_heap_delete(m, v15, v89+int32(4))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v97 = F_heap_copytuple(m, v89)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L11
L21:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v97)+16))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v99+v100)+16)) = l1
	F_CatalogTupleUpdate(m, v15, v97+int32(4), v97)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_pfree(m, v97)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
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
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	m.G0 = v11 + int32(112)
	return
L26:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v129 = F_getObjectDescription(m, v11+int32(4), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v129
	F_errmsg(m, int32(_a_F_changeDependenciesOn_1), v11)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_changeDependenciesOn_2), int32(650), int32(_a_F_changeDependenciesOn_3))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
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
func F_change_plan_targetlist(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v33 float64
	_ = v33
	var v35 float64
	_ = v35
	var v37 float64
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v6 - int32(332) {
	case 0, 1, 2, 3, 4, 28, 29, 30, 35, 38, 39, 40, 41:
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v13 = F_tlist_same_exprs(m, l1, v12)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			if v13 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l1
				v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)))
				v48 = l2 & v47
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)) = uint8(v48)
				return l0
			} else {
				v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)))
				v19 = F_palloc0(m, int32(80))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v21 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v21
					*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v21
					*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v21
					*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(331)
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v31
					v33 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
					*(*float64)(unsafe.Add(mBase, uint32(v19)+8)) = v33
					v35 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
					*(*float64)(unsafe.Add(mBase, uint32(v19)+16)) = v35
					v37 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
					*(*float64)(unsafe.Add(mBase, uint32(v19)+24)) = v37
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v40 = l2 & v17
					*(*uint8)(unsafe.Add(mBase, uint32(v19)+37)) = uint8(v40)
					*(*uint8)(unsafe.Add(mBase, uint32(v19)+36)) = uint8(v21)
					*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v39
					return v19
				}
			}
		}
	default:
		*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l1
		v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)))
		v48 = l2 & v47
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)) = uint8(v48)
		return l0
	case 23:
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+80)))
		if v9&int32(4) != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l1
			v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)))
			v48 = l2 & v47
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)) = uint8(v48)
			return l0
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v13 = F_tlist_same_exprs(m, l1, v12)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				if v13 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l1
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)))
					v48 = l2 & v47
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)) = uint8(v48)
					return l0
				} else {
					v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)))
					v19 = F_palloc0(m, int32(80))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						v21 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v21
						*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v21
						*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v21
						*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(331)
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v31
						v33 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
						*(*float64)(unsafe.Add(mBase, uint32(v19)+8)) = v33
						v35 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
						*(*float64)(unsafe.Add(mBase, uint32(v19)+16)) = v35
						v37 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
						*(*float64)(unsafe.Add(mBase, uint32(v19)+24)) = v37
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						v40 = l2 & v17
						*(*uint8)(unsafe.Add(mBase, uint32(v19)+37)) = uint8(v40)
						*(*uint8)(unsafe.Add(mBase, uint32(v19)+36)) = uint8(v21)
						*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v39
						return v19
					}
				}
			}
		}
	}
}
