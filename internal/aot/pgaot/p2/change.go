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
	v7 = *(*int32)(unsafe.Add(mBase, _consts[204]))
	v8 = m.Env.X__syscall_chdir(m, v7)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v8) {
		*(*int32)(unsafe.Add(mBase, _consts[159])) = int32(0) - v8
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
				v26 = *(*int32)(unsafe.Add(mBase, _consts[204]))
				*(*int32)(unsafe.Add(mBase, uint32(v4))) = v26
				F_errmsg(m, int32(297591), v4)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					F_errfinish(m, int32(493597), int32(468), int32(213670))
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
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
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
	goto L5
L3:
	;
	if v36 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L3
L5:
	;
	if base.Ui32(int32(11999)) < base.Ui32(l0) {
		v36 = v17
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v29 = int32(1)
	v36 = (v29 | base.B2i32(l0 != int32(2200))) & v29
	goto L4
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = l1
	goto L12
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L32
	}
L10:
	;
	F_ScanKeyInit(m, v11+int32(16), int32(4), int32(3), int32(184), int32(1259))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L14
	}
L11:
	;
	goto L10
L12:
	;
	if base.Ui32(int32(11999)) < base.Ui32(l1) {
		v54 = int32(0)
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v47 = int32(1)
	v54 = (v47 | base.B2i32(l1 != int32(2200))) & v47
	goto L11
L14:
	;
	F_ScanKeyInit(m, v11-int32(-64), int32(5), int32(3), int32(184), l0)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v76 = F_systable_beginscan(m, v15, int32(2674), int32(1), int32(0), int32(2), v11+int32(16))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	goto L17
L17:
	;
	v86 = F_systable_getnext(m, v76)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L19
	}
L18:
	;
	F_systable_endscan(m, v76)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L30
	}
L19:
	;
	if v86 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v54 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	goto L18
L23:
	;
	F_CatalogTupleDelete(m, v15, v86+int32(4))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v94 = F_heap_copytuple(m, v86)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L27
	}
L26:
	;
	goto L17
L27:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v94)+16))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v96+v97)+16)) = l1
	F_CatalogTupleUpdate(m, v15, v94+int32(4), v94)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_pfree(m, v94)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	goto L17
L30:
	;
	F_sequence_close(m, v15, int32(3))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	m.G0 = v11 + int32(112)
	return
L32:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v126 = F_getObjectDescription(m, v11+int32(4), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v126
	F_errmsg(m, int32(111138), v11)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(500070), int32(650), int32(285848))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
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
