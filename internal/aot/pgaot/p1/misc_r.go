package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReadDir(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_ReadDirExtended(m, l0, l1, int32(21))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_ReadNextFullTransactionId(m *base.Module) int64 {
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
	var v13 int64
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v7 = F_LWLockAcquire(m, v3+int32(384), int32(1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _consts[140]))
		v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
		v15 = *(*int32)(unsafe.Add(mBase, _consts[24]))
		F_LWLockRelease(m, v15+int32(384))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			return v13
		}
	}
}
func F_RegisterCatcacheInvalidation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
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
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	v1 = l0
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _consts[213]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v18 = *(*int32)(unsafe.Add(mBase, _consts[1161]))
	if v18 <= v16 {
		if v15 == int32(0) {
			v24 = *(*int32)(unsafe.Add(mBase, _consts[190]))
			v26 = F_MemoryContextAlloc(m, v24, int32(512))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				v34 = int32(32)
				v35 = v26
				*(*int32)(unsafe.Add(mBase, _consts[1161])) = v34
				*(*int32)(unsafe.Add(mBase, _consts[213])) = v35
				v40 = v35
				v44 = v40 + v16<<(uint(int32(4))%32)
				*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v1)
				v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+13)))
				*(*uint16)(unsafe.Add(mBase, uint32(v44)+1)) = uint16(v46)
				v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
				*(*uint8)(unsafe.Add(mBase, uint32(v44)+3)) = uint8(v48)
				*(*int32)(unsafe.Add(mBase, uint32(v44)+8)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = l2
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v52 + int32(1)
				m.G0 = v12 + int32(16)
				return
			}
		} else {
			v32 = F_repalloc(m, v15, v18<<(uint(int32(5))%32))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				v34 = v18 << (uint(int32(1)) % 32)
				v35 = v32
				*(*int32)(unsafe.Add(mBase, _consts[1161])) = v34
				*(*int32)(unsafe.Add(mBase, _consts[213])) = v35
				v40 = v35
				v44 = v40 + v16<<(uint(int32(4))%32)
				*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v1)
				v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+13)))
				*(*uint16)(unsafe.Add(mBase, uint32(v44)+1)) = uint16(v46)
				v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
				*(*uint8)(unsafe.Add(mBase, uint32(v44)+3)) = uint8(v48)
				*(*int32)(unsafe.Add(mBase, uint32(v44)+8)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = l2
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v52 + int32(1)
				m.G0 = v12 + int32(16)
				return
			}
		}
	} else {
		v40 = v15
		v44 = v40 + v16<<(uint(int32(4))%32)
		*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v1)
		v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+13)))
		*(*uint16)(unsafe.Add(mBase, uint32(v44)+1)) = uint16(v46)
		v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
		*(*uint8)(unsafe.Add(mBase, uint32(v44)+3)) = uint8(v48)
		*(*int32)(unsafe.Add(mBase, uint32(v44)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = l2
		v52 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v52 + int32(1)
		m.G0 = v12 + int32(16)
		return
	}
}
func F_ReleaseAuxProcessResources(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v3 = *(*int32)(unsafe.Add(mBase, _consts[1231]))
	v4 = int32(1)
	F_ResourceOwnerReleaseInternal(m, v3, v4, l0, v4)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[1231]))
		F_ResourceOwnerReleaseInternal(m, v9, int32(2), l0, int32(1))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, _consts[1231]))
			F_ResourceOwnerReleaseInternal(m, v15, int32(3), l0, int32(1))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, _consts[1231]))
				v22 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v21)+16)) = uint16(v22)
				return
			}
		}
	}
}
func F_ReplaceVarsFromTargetList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
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
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l2
	v23 = F_replace_rte_variables(m, l0, l1, int32(0), int32(1057), v12+int32(12), l7)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return int32(0)
	} else {
		m.G0 = v12 + int32(32)
		return v23
	}
}
func F_ReplaceVarsFromTargetList_callback(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v14 = F_ReplaceVarFromTargetList(m, l0, v9, v10, v11, v12, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v18 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v18
			v26 = F_query_or_expression_tree_walker_impl(m, v14, int32(1049), v6+int32(8), int32(16))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v14
			}
		} else {
			m.G0 = v6 + int32(16)
			return v14
		}
	}
}
func F_ReportChangedGUCOptions(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1208])))
	if v4 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1210])))
	if v8 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[1211]))
	if v37 == int32(0) {
		goto L1
	} else {
		goto L12
	}
L4:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, _consts[29])))
	if v13 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v23 != 0 {
		goto L3
	} else {
		goto L9
	}
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+316))
	v21 = base.B2i32(v19 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[29])) = uint8(v21)
	v23 = v21
	goto L8
L7:
	;
	v23 = int32(0)
	goto L8
L8:
	;
	goto L5
L9:
	;
	v25 = int32(0)
	v28 = int32(10)
	v34 = F_set_config_with_handle(m, int32(22062), v25, int32(347024), v25, v28, v28, v25, int32(1), v25, v25)
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
	goto L3
L12:
	;
	v40 = v37
	goto L13
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	F_ReportGUCOption(m, v40-int32(76))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L10
	} else {
		goto L15
	}
L14:
	;
	goto L1
L15:
	;
	v48 = v40 - int32(48)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v49 & int32(-5)
	*(*int32)(unsafe.Add(mBase, _consts[1211])) = v42
	if v42 != 0 {
		v40 = v42
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
}
func F_RequestCheckpoint(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v4 = F_CreateCheckPoint(m, l0|int32(4))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		F_smgrdestroyall(m)
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			return
		}
	}
}
func F_ReserveExternalFD(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	v4 = *(*int32)(unsafe.Add(mBase, _consts[754]))
	v6 = *(*int32)(unsafe.Add(mBase, _consts[755]))
	if v6 <= int32(0) {
		v36 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _consts[754])) = v36 + int32(1)
	return
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, _consts[756]))
	v12 = *(*int32)(unsafe.Add(mBase, _consts[752]))
	if v12+v6+v4 < v10 {
		v36 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	goto L4
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[749]))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	F_LruDelete(m, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v36 = v24
	goto L1
L6:
	;
	return
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _consts[754]))
	v26 = *(*int32)(unsafe.Add(mBase, _consts[755]))
	if v26 <= int32(0) {
		v36 = v24
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[756]))
	v32 = *(*int32)(unsafe.Add(mBase, _consts[752]))
	if v30 <= v32+v26+v24 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
}
func F_r_LONG_1(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_find_among_b(m, l0, int32(4149440), int32(7))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v4 != int32(0))
	}
}
func F_r_SUFFIX_AN_OK(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	return base.B2i32(v3 != int32(1))
}
func F_r_Step_5b(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 <= v9 {
		v45 = v2
		return v45
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+v7-int32(1)))))
		if v15 != int32(108) {
			v45 = v2
			return v45
		} else {
			v19 = v7 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v19
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v19
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
			if v7 <= v23 {
				v45 = v2
				return v45
			} else {
				if v19 <= v9 {
					v45 = v2
					return v45
				} else {
					v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+v11-int32(1)))))
					if v29 != int32(108) {
						v45 = v2
						return v45
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7 - int32(2)
						v36 = F_slice_del(m, l0)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							if int32(0) <= v36 {
								v42 = int32(1)
							} else {
								v42 = v36
							}
							v45 = v42
							return v45
						}
					}
				}
			}
		}
	}
}
func F_r_VI_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v4 <= v5 {
		v148 = v2
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+v4-int32(1)))))
		if v11 != int32(105) {
			v148 = v2
		} else {
			v15 = v4 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v15
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v15 <= v34 {
				v138 = int32(-1)
				v145 = v138
			} else {
				v51 = int32(1)
				v52 = v15 - v51
				v54 = int32(*(*int8)(unsafe.Add(mBase, uint32(v30+v52))))
				v56 = v54 & int32(255)
				if v52 == v34 {
					v111 = v56
					v112 = v51
				} else {
					if int32(0) <= v54 {
						v111 = v56
						v112 = v51
					} else {
						v62 = v56 & int32(63)
						v64 = v15 - int32(2)
						v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+v64))))
						v68 = v66 << (uint(int32(6)) % 32)
						if base.B2i32(v64 != v34)&base.B2i32(base.Ui32(v66) < base.Ui32(int32(192))) == int32(0) {
							v111 = v68&int32(1984) | v62
							v112 = int32(2)
						} else {
							v81 = v68&int32(4032) | v62
							v83 = v15 - int32(3)
							v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+v83))))
							if base.B2i32(v83 != v34)&base.B2i32(base.Ui32(v85) < base.Ui32(int32(224))) == int32(0) {
								v111 = v85<<(uint(int32(12))%32)&int32(61440) | v81
								v112 = int32(3)
							} else {
								v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+(v30-int32(4))))))
								v111 = v85<<(uint(int32(12))%32)&int32(258048) | v103&int32(7)<<(uint(int32(18))%32) | v81
								v112 = int32(4)
							}
						}
					}
				}
				if int32(246) < v111 {
					v145 = v112
				} else {
					v116 = v111 - int32(97)
					if v116 < int32(0) {
						v145 = v112
					} else {
						v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v116)>>(uint(int32(3))%32)))+uint32(_consts[1298]))))
						if int32(base.Ui32(v122)>>(uint(v116&int32(7))%32))&int32(1) == int32(0) {
							v145 = v112
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v15 - v112
							v138 = int32(0)
							v145 = v138
						}
					}
				}
			}
			v148 = base.B2i32(v145 == int32(0))
		}
	}
	return v148
}
func F_r_VOWEL_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v74 int32
	_ = v74
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v14 <= v13 {
		v119 = int32(-1)
	} else {
		v30 = int32(1)
		v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v15))))
		if base.Ui32(v32) < base.Ui32(int32(192)) {
			v89 = v32
			v90 = v30
		} else {
			v36 = v13 + int32(1)
			if v36 == v14 {
				v89 = v32
				v90 = v30
			} else {
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v15))))
				v41 = v39 & int32(63)
				if base.Ui32(int32(224)) <= base.Ui32(v32) {
					v45 = v13 + int32(2)
					if v45 != v14 {
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v15))))
						v57 = v55 & int32(63)
						if base.Ui32(int32(240)) <= base.Ui32(v32) {
							v61 = v13 + int32(3)
							if v61 != v14 {
								v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v61))))
								v89 = v74&int32(63) | (v32<<(uint(int32(18))%32)&int32(1835008) | v41<<(uint(int32(12))%32) | v57<<(uint(int32(6))%32))
								v90 = int32(4)
							} else {
								v89 = v32<<(uint(int32(12))%32)&int32(61440) | v41<<(uint(int32(6))%32) | v57
								v90 = int32(3)
							}
						} else {
							v89 = v32<<(uint(int32(12))%32)&int32(61440) | v41<<(uint(int32(6))%32) | v57
							v90 = int32(3)
						}
					} else {
						v89 = v32<<(uint(int32(6))%32)&int32(1984) | v41
						v90 = int32(2)
					}
				} else {
					v89 = v32<<(uint(int32(6))%32)&int32(1984) | v41
					v90 = int32(2)
				}
			}
		}
		if int32(117) < v89 {
			v112 = v90
		} else {
			v94 = v89 - int32(97)
			if v94 < int32(0) {
				v112 = v90
			} else {
				v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v94)>>(uint(int32(3))%32)))+uint32(_consts[1299]))))
				if int32(base.Ui32(v100)>>(uint(v94&int32(7))%32))&int32(1) == int32(0) {
					v112 = v90
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90 + v13
					v112 = int32(0)
				}
			}
		}
		v119 = v112
	}
	return base.B2i32(v119 == int32(0))
}
func F_r_e_ending_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v2
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v9 <= v11 {
		v172 = v2
		return v172
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v9-int32(1)))))
		if v17 != int32(101) {
			v172 = v2
			return v172
		} else {
			v21 = v9 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v21
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v21
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			if v9 <= v24 {
				v172 = v2
				return v172
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v43 <= v44 {
					v149 = int32(-1)
					v156 = v149
				} else {
					v61 = int32(1)
					v62 = v43 - v61
					v64 = int32(*(*int8)(unsafe.Add(mBase, uint32(v40+v62))))
					v66 = v64 & int32(255)
					if v62 == v44 {
						v121 = v66
						v122 = v61
					} else {
						if int32(0) <= v64 {
							v121 = v66
							v122 = v61
						} else {
							v72 = v66 & int32(63)
							v74 = v43 - int32(2)
							v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+v74))))
							v78 = v76 << (uint(int32(6)) % 32)
							if base.B2i32(v74 != v44)&base.B2i32(base.Ui32(v76) < base.Ui32(int32(192))) == int32(0) {
								v121 = v78&int32(1984) | v72
								v122 = int32(2)
							} else {
								v91 = v78&int32(4032) | v72
								v93 = v43 - int32(3)
								v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+v93))))
								if base.B2i32(v93 != v44)&base.B2i32(base.Ui32(v95) < base.Ui32(int32(224))) == int32(0) {
									v121 = v95<<(uint(int32(12))%32)&int32(61440) | v91
									v122 = int32(3)
								} else {
									v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+(v40-int32(4))))))
									v121 = v95<<(uint(int32(12))%32)&int32(258048) | v113&int32(7)<<(uint(int32(18))%32) | v91
									v122 = int32(4)
								}
							}
						}
					}
					if int32(232) < v121 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v43 - v122
						v149 = int32(0)
						v156 = v149
					} else {
						v126 = v121 - int32(97)
						if v126 < int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v43 - v122
							v149 = int32(0)
							v156 = v149
						} else {
							v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v126)>>(uint(int32(3))%32)))+uint32(_consts[1294]))))
							if int32(base.Ui32(v132)>>(uint(v126&int32(7))%32))&int32(1) == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v43 - v122
								v149 = int32(0)
								v156 = v149
							} else {
								v156 = v122
							}
						}
					}
				}
				if v156 != 0 {
					v172 = v2
					return v172
				} else {
					v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v157 + (v21 - v26)
					v161 = F_slice_del(m, l0)
					mBase = m.M
					v164 = m.ExcPending
					if v164 != 0 {
						return int32(0)
					} else {
						if v161 < int32(0) {
							v172 = v161
							return v172
						} else {
							v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							*(*int32)(unsafe.Add(mBase, uint32(v167)+12)) = int32(1)
							v170 = F_r_undouble(m, l0)
							mBase = m.M
							v171 = m.ExcPending
							if v171 != 0 {
								return int32(0)
							} else {
								v172 = v170
								return v172
							}
						}
					}
				}
			}
		}
	}
}
func F_r_remove_suffix_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
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
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v4
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v4 <= v6 {
		v36 = v2
		return v36
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+v4-int32(1)))))
		switch v12 - int32(105) {
		case 0, 5:
			v17 = F_find_among_b(m, l0, int32(4154080), int32(3))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if v17 == int32(0) {
					v36 = v2
					return v36
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v23
					v25 = F_slice_del(m, l0)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						if v25 < int32(0) {
							v36 = v25
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
							v31 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v30 - v31
							v36 = v31
						}
						return v36
					}
				}
			}
		default:
			v36 = v2
			return v36
		}
	}
}
func F_r_stem_suffix_chain_before_ki(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
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
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
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
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v379 int32
	_ = v379
	var v386 int32
	_ = v386
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v515 int32
	_ = v515
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v580 int32
	_ = v580
	var v587 int32
	_ = v587
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
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
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v6
	v8 = int32(2)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6-v13 < v8 {
		v23 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v23 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L1
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = F_memcmp(m, v16+v6-v8, int32(2180396), v8)
	mBase = m.M
	if v19 != 0 {
		v23 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6 - v8
	v23 = int32(1)
	goto L2
L5:
	;
	return int32(0)
L6:
	;
	goto L7
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v30 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v30 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	return v661
L9:
	;
	if v91 < int32(0) {
		goto L187
	} else {
		goto L188
	}
L10:
	;
	if v160 < int32(0) {
		goto L181
	} else {
		goto L182
	}
L11:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v169 = v29 - v28
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v168 - v169
	v172 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v172 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v35 = v33 - int32(1)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v35 <= v36 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v35))))
	switch v40 - int32(97) {
	case 0, 4:
		goto L14
	default:
		goto L11
	}
L14:
	;
	v45 = F_find_among_b(m, l0, int32(4331744), int32(4))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	if v45 == int32(0) {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v51
	v53 = F_slice_del(m, l0)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	if v53 < int32(0) {
		v661 = v53
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v57 = int32(0)
	v58 = base.B2i32(v23 != v57)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v59
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v62 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v62 == v57 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v106 = v59 - v61
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v106 + v107
	v110 = F_r_mark_possessives(m, l0)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L15
	} else {
		goto L33
	}
L21:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v66-int32(2) <= v65 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+v66-int32(1)))))
	if v74 != int32(114) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v79 = F_find_among_b(m, l0, int32(4330480), int32(2))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L15
	} else {
		goto L24
	}
L24:
	;
	if v79 == int32(0) {
		goto L20
	} else {
		goto L25
	}
L25:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v83
	v85 = F_slice_del(m, l0)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L15
	} else {
		goto L26
	}
L26:
	;
	if v85 < int32(0) {
		v661 = v85
		goto L8
	} else {
		goto L27
	}
L27:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v91 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L15
	} else {
		goto L28
	}
L28:
	;
	if v91 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v95 + (v89 - v90)
	return int32(1)
L30:
	;
	goto L31
L31:
	;
	v102 = base.B2i32(v91 < int32(0))
	if v91 < int32(0) {
		goto L9
	} else {
		goto L32
	}
L32:
	;
	return int32(1)
L33:
	;
	if v110 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v114 + v106
	return int32(1)
L35:
	;
	goto L36
L36:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v119
	v121 = F_slice_del(m, l0)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L15
	} else {
		goto L37
	}
L37:
	;
	if v121 < int32(0) {
		v661 = v121
		goto L8
	} else {
		goto L38
	}
L38:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v125
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v128 = v127 - v125
	v129 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v129 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v154
	v156 = F_slice_del(m, l0)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L15
	} else {
		goto L46
	}
L40:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v149 - v128
	return int32(1)
L41:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v133-int32(2) <= v132 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137+v133-int32(1)))))
	if v141 != int32(114) {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v146 = F_find_among_b(m, l0, int32(4330480), int32(2))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L15
	} else {
		goto L44
	}
L44:
	;
	if v146 != 0 {
		goto L39
	} else {
		goto L45
	}
L45:
	;
	goto L40
L46:
	;
	if v156 < int32(0) {
		v661 = v156
		goto L8
	} else {
		goto L47
	}
L47:
	;
	v160 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L15
	} else {
		goto L48
	}
L48:
	;
	if v160 != 0 {
		goto L10
	} else {
		goto L49
	}
L49:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v162 - v128
	return int32(1)
L50:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v459 - v169
	v462 = int32(0)
	v463 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v463 == v462 {
		v661 = v462
		goto L8
	} else {
		goto L129
	}
L51:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v177 = v175 - int32(1)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v177 <= v178 {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180+v177))))
	if v182 != int32(110) {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v187 = F_find_among_b(m, l0, int32(4331616), int32(4))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L15
	} else {
		goto L54
	}
L54:
	;
	if v187 == int32(0) {
		goto L50
	} else {
		goto L55
	}
L55:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v199 <= v200 {
		goto L61
	} else {
		goto L62
	}
L56:
	;
	if v261 == int32(0) {
		goto L50
	} else {
		goto L72
	}
L57:
	;
	goto L56
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v257
	v261 = int32(1)
	goto L57
L59:
	;
	v257 = v209 - v197 + v216
	goto L58
L60:
	;
	v224 = v199 - v197
	v225 = v223 + v224
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v225
	if v222 < v225 {
		goto L66
	} else {
		goto L67
	}
L61:
	;
	v221 = v198
	v222 = v200
	v223 = v197
	goto L60
L62:
	;
	goto L63
L63:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199+v198-int32(1)))))
	if v205 != int32(110) {
		v221 = v198
		v222 = v200
		v223 = v197
		goto L60
	} else {
		goto L64
	}
L64:
	;
	v209 = v199 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v209
	v214 = int32(0)
	v215 = F_in_grouping_b_U(m, l0, int32(2180000), int32(97), int32(305), v214)
	mBase = m.M
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v215 == v214 {
		goto L59
	} else {
		goto L65
	}
L65:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v221 = v219
	v222 = v220
	v223 = v216
	goto L60
L66:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225+v221-int32(1)))))
	if v232 == int32(110) {
		v261 = int32(0)
		goto L57
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v236 = int32(0)
	v238 = F_skip_b_utf8(m, v221, v225, v222, int32(1))
	mBase = m.M
	if v238 < v236 {
		v261 = v236
		goto L57
	} else {
		goto L70
	}
L69:
	;
	goto L68
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v238
	v246 = F_in_grouping_b_U(m, l0, int32(2180000), int32(97), int32(305), int32(0))
	mBase = m.M
	if v246 != 0 {
		v261 = v236
		goto L57
	} else {
		goto L71
	}
L71:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v257 = v247 + v224
	goto L58
L72:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v268
	v270 = F_slice_del(m, l0)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L15
	} else {
		goto L73
	}
L73:
	;
	if v270 < int32(0) {
		v661 = v270
		goto L8
	} else {
		goto L74
	}
L74:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v274
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v274-int32(3) <= v277 {
		v296 = v276
		goto L78
	} else {
		goto L79
	}
L75:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v441 - v297
	v444 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L15
	} else {
		goto L124
	}
L76:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v422
	v424 = F_slice_del(m, l0)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L15
	} else {
		goto L117
	}
L77:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v414
	v416 = F_slice_del(m, l0)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L15
	} else {
		goto L115
	}
L78:
	;
	v297 = v276 - v274
	v298 = v296 - v297
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v298
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v298
	v301 = F_r_mark_possessives(m, l0)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L15
	} else {
		goto L86
	}
L79:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281+v274-int32(1)))))
	if v285 != int32(177) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	if v285 != int32(105) {
		v296 = v276
		goto L78
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v292 = F_find_among_b(m, l0, int32(4331056), int32(2))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L15
	} else {
		goto L84
	}
L83:
	;
	goto L82
L84:
	;
	if v292 != 0 {
		goto L77
	} else {
		goto L85
	}
L85:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v296 = v294
	goto L78
L86:
	;
	if v301 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v305 - v297
	v308 = int32(0)
	v314 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v314 == v308 {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	goto L89
L89:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v397
	v399 = F_slice_del(m, l0)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L15
	} else {
		goto L111
	}
L90:
	;
	if v394 == int32(0) {
		goto L75
	} else {
		goto L110
	}
L91:
	;
	v394 = int32(0)
	goto L90
L92:
	;
	goto L93
L93:
	;
	v322 = F_in_grouping_b_U(m, l0, int32(2180448), int32(105), int32(305), int32(0))
	mBase = m.M
	if v322 != 0 {
		v386 = v308
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v394 = v386
	goto L90
L95:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v325 <= v326 {
		goto L99
	} else {
		goto L100
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v379
	v386 = int32(1)
	goto L94
L97:
	;
	v379 = v335 - v323 + v342
	goto L96
L98:
	;
	v350 = v325 - v323
	v351 = v349 + v350
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v351
	if v348 < v351 {
		goto L104
	} else {
		goto L105
	}
L99:
	;
	v347 = v324
	v348 = v326
	v349 = v323
	goto L98
L100:
	;
	goto L101
L101:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324+v325-int32(1)))))
	if v331 != int32(115) {
		v347 = v324
		v348 = v326
		v349 = v323
		goto L98
	} else {
		goto L102
	}
L102:
	;
	v335 = v325 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v335
	v340 = int32(0)
	v341 = F_in_grouping_b_U(m, l0, int32(2180000), int32(97), int32(305), v340)
	mBase = m.M
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v341 == v340 {
		goto L97
	} else {
		goto L103
	}
L103:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v347 = v345
	v348 = v346
	v349 = v342
	goto L98
L104:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351+v347-int32(1)))))
	if v357 == int32(115) {
		v386 = v308
		goto L94
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v361 = F_skip_b_utf8(m, v347, v351, v348, int32(1))
	mBase = m.M
	if v361 < int32(0) {
		v386 = v308
		goto L94
	} else {
		goto L108
	}
L107:
	;
	goto L106
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v361
	v369 = F_in_grouping_b_U(m, l0, int32(2180000), int32(97), int32(305), int32(0))
	mBase = m.M
	if v369 != 0 {
		v386 = v308
		goto L94
	} else {
		goto L109
	}
L109:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v379 = v370 + v350
	goto L96
L110:
	;
	goto L89
L111:
	;
	if v399 < int32(0) {
		v661 = v399
		goto L8
	} else {
		goto L112
	}
L112:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v403
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v406 = v405 - v403
	v407 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L15
	} else {
		goto L113
	}
L113:
	;
	if v407 != 0 {
		goto L76
	} else {
		goto L114
	}
L114:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v409 - v406
	return int32(1)
L115:
	;
	if v416 < int32(0) {
		v661 = v416
		goto L8
	} else {
		goto L116
	}
L116:
	;
	return int32(1)
L117:
	;
	if v424 < int32(0) {
		v661 = v424
		goto L8
	} else {
		goto L118
	}
L118:
	;
	v428 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L15
	} else {
		goto L119
	}
L119:
	;
	if v428 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v432 - v406
	return int32(1)
L121:
	;
	goto L122
L122:
	;
	if v428 < int32(0) {
		v661 = v428
		goto L8
	} else {
		goto L123
	}
L123:
	;
	return int32(1)
L124:
	;
	if v444 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v448 + (v274 - v276)
	return int32(1)
L126:
	;
	goto L127
L127:
	;
	if v444 < int32(0) {
		v661 = v444
		goto L8
	} else {
		goto L128
	}
L128:
	;
	return int32(1)
L129:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v467-int32(2) <= v466 {
		v661 = v462
		goto L8
	} else {
		goto L130
	}
L130:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471+v467-int32(1)))))
	switch v475 - int32(97) {
	case 0, 4:
		goto L131
	default:
		v661 = v462
		goto L8
	}
L131:
	;
	v480 = F_find_among_b(m, l0, int32(4331312), int32(2))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L15
	} else {
		goto L132
	}
L132:
	;
	if v480 == int32(0) {
		v661 = v462
		goto L8
	} else {
		goto L133
	}
L133:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v486-int32(3) <= v485 {
		v505 = v484
		goto L137
	} else {
		goto L138
	}
L134:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v642 - v506
	v645 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L15
	} else {
		goto L179
	}
L135:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v623
	v625 = F_slice_del(m, l0)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L15
	} else {
		goto L172
	}
L136:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v615
	v617 = F_slice_del(m, l0)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L15
	} else {
		goto L170
	}
L137:
	;
	v506 = v484 - v486
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v505 - v506
	v509 = int32(0)
	v515 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v515 == v509 {
		goto L146
	} else {
		goto L147
	}
L138:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v490+v486-int32(1)))))
	if v494 != int32(177) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	if v494 != int32(105) {
		v505 = v484
		goto L137
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v501 = F_find_among_b(m, l0, int32(4331056), int32(2))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L15
	} else {
		goto L143
	}
L142:
	;
	goto L141
L143:
	;
	if v501 != 0 {
		goto L136
	} else {
		goto L144
	}
L144:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v505 = v503
	goto L137
L145:
	;
	if v595 == int32(0) {
		goto L134
	} else {
		goto L165
	}
L146:
	;
	v595 = int32(0)
	goto L145
L147:
	;
	goto L148
L148:
	;
	v523 = F_in_grouping_b_U(m, l0, int32(2180448), int32(105), int32(305), int32(0))
	mBase = m.M
	if v523 != 0 {
		v587 = v509
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v595 = v587
	goto L145
L150:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v527 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v526 <= v527 {
		goto L154
	} else {
		goto L155
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v580
	v587 = int32(1)
	goto L149
L152:
	;
	v580 = v536 - v524 + v543
	goto L151
L153:
	;
	v551 = v526 - v524
	v552 = v550 + v551
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v552
	if v549 < v552 {
		goto L159
	} else {
		goto L160
	}
L154:
	;
	v548 = v525
	v549 = v527
	v550 = v524
	goto L153
L155:
	;
	goto L156
L156:
	;
	v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525+v526-int32(1)))))
	if v532 != int32(115) {
		v548 = v525
		v549 = v527
		v550 = v524
		goto L153
	} else {
		goto L157
	}
L157:
	;
	v536 = v526 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v536
	v541 = int32(0)
	v542 = F_in_grouping_b_U(m, l0, int32(2180000), int32(97), int32(305), v541)
	mBase = m.M
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v542 == v541 {
		goto L152
	} else {
		goto L158
	}
L158:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v548 = v546
	v549 = v547
	v550 = v543
	goto L153
L159:
	;
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552+v548-int32(1)))))
	if v558 == int32(115) {
		v587 = v509
		goto L149
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v562 = F_skip_b_utf8(m, v548, v552, v549, int32(1))
	mBase = m.M
	if v562 < int32(0) {
		v587 = v509
		goto L149
	} else {
		goto L163
	}
L162:
	;
	goto L161
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v562
	v570 = F_in_grouping_b_U(m, l0, int32(2180000), int32(97), int32(305), int32(0))
	mBase = m.M
	if v570 != 0 {
		v587 = v509
		goto L149
	} else {
		goto L164
	}
L164:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v580 = v571 + v551
	goto L151
L165:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v598
	v600 = F_slice_del(m, l0)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L15
	} else {
		goto L166
	}
L166:
	;
	if v600 < int32(0) {
		v661 = v600
		goto L8
	} else {
		goto L167
	}
L167:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v604
	v606 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v607 = v606 - v604
	v608 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L15
	} else {
		goto L168
	}
L168:
	;
	if v608 != 0 {
		goto L135
	} else {
		goto L169
	}
L169:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v610 - v607
	return int32(1)
L170:
	;
	if v617 < int32(0) {
		v661 = v617
		goto L8
	} else {
		goto L171
	}
L171:
	;
	return int32(1)
L172:
	;
	if v625 < int32(0) {
		v661 = v625
		goto L8
	} else {
		goto L173
	}
L173:
	;
	v629 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L15
	} else {
		goto L174
	}
L174:
	;
	if v629 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v633 - v607
	return int32(1)
L176:
	;
	goto L177
L177:
	;
	if v629 < int32(0) {
		v661 = v629
		goto L8
	} else {
		goto L178
	}
L178:
	;
	return int32(1)
L179:
	;
	if v645 <= int32(0) {
		v661 = v645
		goto L8
	} else {
		goto L180
	}
L180:
	;
	return int32(1)
L181:
	;
	if v160 < int32(0) {
		goto L184
	} else {
		goto L185
	}
L182:
	;
	goto L183
L183:
	;
	return int32(1)
L184:
	;
	v655 = v160
	goto L186
L185:
	;
	v655 = v58
	goto L186
L186:
	;
	return v655
L187:
	;
	v659 = v91
	goto L189
L188:
	;
	v659 = v58
	goto L189
L189:
	;
	v661 = v659
	goto L8
}
func F_r_undouble(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = v5 - int32(1)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 <= v8 {
		v105 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v105
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v7))))
	if v12&int32(224) != int32(96) {
		v105 = v2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if int32(1)<<(uint(v12)%32)&int32(1050640) == int32(0) {
		v105 = v2
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v26 = F_find_among_b(m, l0, int32(4206912), int32(3))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	if v26 == int32(0) {
		v105 = v2
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v34 = v32 + (v5 - v23)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v34
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L10
L8:
	;
	if v89 < int32(0) {
		v105 = v2
		goto L1
	} else {
		goto L28
	}
L10:
	;
	goto L11
L11:
	;
	goto L12
L12:
	;
	v45 = v34
	v47 = int32(1)
	goto L15
L14:
	;
	v89 = v71
	goto L8
L15:
	;
	if v45 <= v38 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L14
L17:
	;
	v89 = int32(-1)
	goto L8
L18:
	;
	goto L19
L19:
	;
	v52 = v45 - int32(1)
	v54 = int32(*(*int8)(unsafe.Add(mBase, uint32(v37+v52))))
	if int32(0) <= v54 {
		v71 = v52
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v75 = int32(1)
	if v75 < v47 {
		v45 = v71
		v47 = v47 - v75
		goto L15
	} else {
		goto L27
	}
L21:
	;
	if v52 <= v38 {
		v71 = v52
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v59 = v52
	goto L23
L23:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v59))))
	if base.Ui32(int32(191)) < base.Ui32(v64) {
		v71 = v59
		goto L20
	} else {
		goto L25
	}
L24:
	;
	v71 = v38
	goto L20
L25:
	;
	v68 = v59 - int32(1)
	if v38 < v68 {
		v59 = v68
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	goto L16
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v89
	v95 = F_slice_del(m, l0)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	if int32(0) <= v95 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v102 = int32(1)
	goto L32
L31:
	;
	v102 = v95 >> (uint(int32(31)) % 32) & v95
	goto L32
L32:
	;
	v105 = v102
	goto L1
}
func F_rangeTableEntry_used(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v3 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l1
	v15 = F_query_or_expression_tree_walker_impl(m, l0, int32(1050), v6+int32(8), v3)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return v15
	}
}
func F_raw_heap_insert(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+119)))
	if v14 == int32(116) {
		v30 = l1
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L8
	} else {
		goto L65
	}
L2:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v35 = (v31 + int32(7)) & int32(-8)
	if base.Ui32(v35) < base.Ui32(int32(8161)) {
		goto L10
	} else {
		goto L11
	}
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+20)))
	if v18&int32(4) == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v23) < base.Ui32(int32(2033)) {
		v30 = l1
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v28 = F_heap_toast_insert_or_update(m, v12, l1, int32(0), int32(10))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L6
L8:
	;
	return
L9:
	;
	v30 = v28
	goto L2
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+180))
	if v39 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L8
	} else {
		goto L61
	}
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v46 = base.I32_div_s(int32(819200)-v41<<(uint(int32(13))%32), int32(100))
	v48 = v46
	goto L15
L14:
	;
	v48 = int32(0)
	goto L15
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v49 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v183 = F_PageAddItemExtended(m, v178, v179, v180, int32(0), int32(2))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L8
	} else {
		goto L52
	}
L17:
	;
	v53 = int32(4)
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49)+14)))
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49)+12)))
	v56 = v54 - v55
	if v56 <= v53 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	goto L19
L19:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v134 = F_smgr_bulk_get_buf(m, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L8
	} else {
		goto L41
	}
L20:
	;
	if base.Ui32(v35+v48) <= base.Ui32(v118) {
		v178 = v49
		goto L16
	} else {
		goto L39
	}
L21:
	;
	v59 = v53
	goto L23
L22:
	;
	v59 = v56
	goto L23
L23:
	;
	v61 = v59 - int32(4)
	if v61 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v118 = int32(0)
	goto L20
L25:
	;
	goto L26
L26:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v55) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v118 = v61
	goto L20
L28:
	;
	v72 = int32(base.Ui32(v55+int32(262120)) >> (uint(int32(2)) % 32))
	goto L30
L29:
	;
	v72 = int32(0)
	goto L30
L30:
	;
	if base.Ui32(v72&int32(65535)) < base.Ui32(int32(291)) {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+10)))
	if v77&int32(1) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v118 = int32(0)
	goto L20
L33:
	;
	goto L34
L34:
	;
	v86 = int32(1)
	goto L35
L35:
	;
	v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86&int32(65535)<<(uint(int32(2))%32)+(v49+int32(24))-int32(3)))))
	if v97&int32(384) == int32(0) {
		goto L27
	} else {
		goto L37
	}
L36:
	;
	v118 = int32(0)
	goto L20
L37:
	;
	v103 = v86 + int32(1)
	v104 = int32(65535)
	if base.Ui32(v103&v104) <= base.Ui32(v72&v104) {
		v86 = v103
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_smgr_bulk_write(m, v121, v122, v123, int32(1))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L8
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v129 + int32(1)
	goto L19
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v134
	if v134&int32(3) != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v178 = v134
	goto L16
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134)+10)) = int32(1572864)
	v169 = int32(8196)
	*(*uint16)(unsafe.Add(mBase, uint32(v134)+18)) = uint16(v169)
	v175 = int32(8192)
	*(*uint16)(unsafe.Add(mBase, uint32(v134)+16)) = uint16(v175)
	*(*uint16)(unsafe.Add(mBase, uint32(v134)+14)) = uint16(v175)
	goto L42
L44:
	;
	v163 = F___memset(m, v134, int32(0), int32(8192))
	mBase = m.M
	goto L43
L45:
	;
	goto L44
L52:
	;
	if v183 == int32(0) {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v183)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v187)
	v191 = int32(base.Ui32(v187) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v191)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v193)+16)))
	if v194 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v183<<(uint(int32(2))%32)+v178)+20))
	v203 = v178 + v200&int32(32767)
	v205 = l1 + int32(4)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+12)) = v206
	v208 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v205)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v203)+16)) = uint16(v208)
	goto L56
L55:
	;
	goto L56
L56:
	;
	if l1 != v30 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	F_pfree(m, v30)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L8
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	m.G0 = v10 + int32(16)
	return
L60:
	;
	goto L59
L61:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L8
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(8160)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v35
	F_errmsg(m, int32(35467), v10)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L8
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(476880), int32(641), int32(77994))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L8
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	F_errmsg_internal(m, int32(369008), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L8
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(476880), int32(679), int32(77994))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L8
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_read_gucstate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v6) < base.Ui32(l1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = v6
	goto L5
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L9
	} else {
		goto L13
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v14
	return v6
L5:
	;
	v14 = v10 + int32(1)
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v15 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	if l1 != v14 {
		v10 = v14
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	return int32(0)
L10:
	;
	F_errmsg_internal(m, int32(339318), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(480467), int32(6155), int32(337416))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L9
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
	F_errmsg_internal(m, int32(339402), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(480467), int32(6148), int32(337416))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_readdir(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int64
	_ = v32
	var v37 int32
	_ = v37
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v5 <= v4 {
		v7 = int32(0)
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v12 = m.Env.X__syscall_getdents64(m, v8, l0+int32(24), int32(2048))
		mBase = m.M
		if v12 <= v7 {
			if v12 == int32(-44) {
				v37 = v7
				return v37
			} else {
				if v12 == int32(0) {
					v37 = v7
					return v37
				} else {
					v20 = int32(0)
					*(*int32)(unsafe.Add(mBase, _consts[137])) = v20 - v12
					return v20
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v12
			v27 = v7
			v28 = l0 + v27
			v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+40)))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v27 + v29
			v32 = *(*int64)(unsafe.Add(mBase, uint32(v28)+32))
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v32
			v37 = v28 + int32(24)
			return v37
		}
	} else {
		v27 = v4
		v28 = l0 + v27
		v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+40)))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v27 + v29
		v32 = *(*int64)(unsafe.Add(mBase, uint32(v28)+32))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v32
		v37 = v28 + int32(24)
		return v37
	}
}
func F_record_larger(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v4 = F_record_cmp(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if int32(0) < v4 {
			v10 = int32(20)
		} else {
			v10 = int32(28)
		}
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0+v10)))
		return v12
	}
}
func F_record_le(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_record_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v2 <= int32(0))
	}
}
func F_record_ne(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_record_eq(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2 ^ int32(1)
	}
}
func F_reduce_expanded_ranges(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v19 = base.I32_div_s(l3, int32(2))
	v21 = l1 - int32(1)
	if v19 <= v21 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = l5
	v27 = F_palloc(m, l3<<(uint(int32(2))%32))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v211 = l1
	goto L3
L3:
	;
	m.G0 = v16 + int32(16)
	return v211
L4:
	;
	return int32(0)
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v31
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0+v21*int32(12))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v36
	if l3 <= int32(3) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v169 = int32(0)
	goto L25
L7:
	;
	F_qsort_arg(m, v27, int32(2), int32(4), int32(21), v16+int32(8))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v49 = l0 + int32(12)
	v50 = int32(2)
	if v19 <= v50 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v164 = int32(1)
	goto L6
L11:
	;
	v54 = v50
	goto L13
L12:
	;
	v54 = v19
	goto L13
L13:
	;
	v55 = int32(1)
	v56 = v54 - v55
	if l3 < int32(6) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v56&v55 != 0 {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	v114 = int32(0)
	v115 = v50
	goto L14
L16:
	;
	goto L17
L17:
	;
	v64 = int32(0)
	v70 = v64
	v71 = v50
	v73 = v64
	goto L18
L18:
	;
	v79 = int32(2)
	v81 = v27 + v71<<(uint(v79)%32)
	v82 = int32(4)
	v84 = l2 + v70<<(uint(v82)%32)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v86 = int32(12)
	v87 = v85 * v86
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0+v87)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = v89
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v87+v49)))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+4)) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
	v96 = v94 * v86
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0+v96)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+8)) = v98
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v49+v96)))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+12)) = v101
	v104 = v70 + v79
	v106 = v71 + v82
	v108 = v73 + v79
	if v108 != v56&int32(-2) {
		v70 = v104
		v71 = v106
		v73 = v108
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v114 = v104
	v115 = v106
	goto L14
L20:
	;
	goto L19
L21:
	;
	v123 = int32(2)
	v125 = v27 + v115<<(uint(v123)%32)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l2+v114<<(uint(int32(4))%32))))
	v131 = v129 * int32(12)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0+v131)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v125))) = v133
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v131+v49)))
	*(*int32)(unsafe.Add(mBase, uint32(v125)+4)) = v136
	v142 = v115 + v123
	goto L23
L22:
	;
	v142 = v115
	goto L23
L23:
	;
	F_qsort_arg(m, v27, v142, int32(4), int32(21), v16+int32(8))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v164 = int32(base.Ui32(v142) >> (uint(int32(1)) % 32))
	goto L6
L25:
	;
	v181 = l0 + v169*int32(12)
	v184 = v27 + v169<<(uint(int32(3))%32)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	*(*int32)(unsafe.Add(mBase, uint32(v181))) = v185
	v188 = v184 + int32(4)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	*(*int32)(unsafe.Add(mBase, uint32(v181)+4)) = v189
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v195 = F_FunctionCall2Coll(m, v192, v193, v194, v189)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L4
	} else {
		goto L27
	}
L26:
	;
	v211 = v164
	goto L3
L27:
	;
	if v195 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v205 = int32(0)
	goto L30
L29:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v201 = F_FunctionCall2Coll(m, v197, v198, v199, v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L31
	}
L30:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v181)+8)) = uint8(v205)
	v208 = v169 + int32(1)
	if v208 != v164 {
		v169 = v208
		goto L25
	} else {
		goto L32
	}
L31:
	;
	v205 = base.B2i32(v201 == int32(0))
	goto L30
L32:
	;
	goto L26
}
func F_reduce_outer_joins_pass1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = F_palloc(m, int32(12))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v15
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)) = uint8(v15)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v15
	if l0 == v15 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L27
	}
L4:
	;
	m.G0 = v8 + int32(16)
	return v11
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v23 - int32(63) {
	case 0:
		goto L6
	case 1:
		goto L7
	case 2:
		goto L8
	default:
		goto L3
	}
L6:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v101 = F_bms_make_singleton(m, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L26
	}
L7:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(1)<<(uint(v62)%32)&int32(174) != 0 {
		goto L17
	} else {
		goto L18
	}
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v26 == int32(0) {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v29 <= int32(0) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v34 = int32(0)
	goto L11
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37+v34<<(uint(int32(2))%32))))
	v42 = F_reduce_outer_joins_pass1(m, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L4
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v46 = F_bms_add_members(m, v44, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v46
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+4)))
	v51 = v49 | v50
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)) = uint8(v51)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v54 = F_lappend(m, v53, v42)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v54
	v58 = v34 + int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v58 < v59 {
		v34 = v58
		goto L11
	} else {
		goto L16
	}
L16:
	;
	goto L12
L17:
	;
	v66 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)) = uint8(v66)
	goto L19
L18:
	;
	goto L19
L19:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v69 = F_reduce_outer_joins_pass1(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v73 = F_bms_add_members(m, v71, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v73
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+4)))
	v78 = v76 | v77
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)) = uint8(v78)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v81 = F_lappend(m, v80, v69)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v81
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v85 = F_reduce_outer_joins_pass1(m, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v89 = F_bms_add_members(m, v87, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v89
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+4)))
	v94 = v92 | v93
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)) = uint8(v94)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v97 = F_lappend(m, v96, v85)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v97
	goto L4
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v101
	goto L4
L27:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v117
	F_errmsg_internal(m, int32(467381), v8)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(480094), int32(3232), int32(532879))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_regclassin(m *base.Module, l0 int32) int32 {
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v12 == int32(45) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L49
	} else {
		goto L65
	}
L2:
	;
	m.G0 = v8 + int32(16)
	return v214
L3:
	;
	v214 = int32(0)
	goto L2
L4:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _consts[298]))
	if v176 == int32(0) {
		goto L1
	} else {
		goto L51
	}
L5:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v15 == int32(0) {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	if base.Ui32(int32(9)) < base.Ui32((v12-int32(48))&int32(255)) {
		goto L4
	} else {
		goto L9
	}
L8:
	;
	goto L4
L9:
	;
	v24 = int32(527523)
	v28 = m.G0
	v30 = v28 - int32(32)
	v31 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+24)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v30)+16)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v30)+8)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v30))) = v31
	v39 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1069])))
	if v39 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v11&int32(3) == int32(0) {
		v131 = v11
		goto L33
	} else {
		goto L34
	}
L11:
	;
	v107 = int32(0)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1070])))
	if v43 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v47 = v11
	goto L17
L15:
	;
	goto L16
L16:
	;
	v57 = v24
	v58 = v39
	goto L20
L17:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v53 == v39 {
		v47 = v47 + int32(1)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v107 = v47 - v11
	goto L10
L19:
	;
	goto L18
L20:
	;
	v65 = v30 + int32(base.Ui32(v58)>>(uint(int32(3))%32))&int32(28)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v67 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v66 | v67<<(uint(v58)%32)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+1)))
	if v71 != 0 {
		v57 = v57 + v67
		v58 = v71
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v74 == int32(0) {
		v99 = v11
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L21
L23:
	;
	v107 = v99 - v11
	goto L10
L24:
	;
	v78 = v11
	v79 = v74
	goto L25
L25:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v30+int32(base.Ui32(v79)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v87)>>(uint(v79)%32))&int32(1) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v99 = v95
	goto L23
L27:
	;
	v99 = v78
	goto L23
L28:
	;
	goto L29
L29:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
	v95 = v78 + int32(1)
	if v93 != 0 {
		v78 = v95
		v79 = v93
		goto L25
	} else {
		goto L30
	}
L30:
	;
	goto L26
L31:
	;
	if v107 != v164 {
		goto L4
	} else {
		goto L48
	}
L32:
	;
	v164 = v156 - v11
	goto L31
L33:
	;
	v135 = v131
	goto L42
L34:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v115 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v164 = int32(0)
	goto L31
L36:
	;
	goto L37
L37:
	;
	v120 = v11
	goto L38
L38:
	;
	v124 = v120 + int32(1)
	if v124&int32(3) == int32(0) {
		v131 = v124
		goto L33
	} else {
		goto L40
	}
L39:
	;
	v156 = v124
	goto L32
L40:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if v129 != 0 {
		v120 = v124
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	v144 = int32(-2139062144)
	if (int32(16843008)-v141|v141)&v144 == v144 {
		v135 = v135 + int32(4)
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v150 = v135
	goto L45
L44:
	;
	goto L43
L45:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	if v154 != 0 {
		v150 = v150 + int32(1)
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v156 = v150
	goto L32
L47:
	;
	goto L46
L48:
	;
	v170 = F_DirectInputFunctionCallSafe(m, int32(547), v11, int32(-1), v10, v8+int32(12))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	return int32(0)
L50:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v214 = v174
	goto L2
L51:
	;
	v179 = F_stringToQualifiedNameList(m, v11, v10)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	if v179 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v183 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v183)
	goto L3
L54:
	;
	goto L55
L55:
	;
	v185 = F_makeRangeVarFromNameList(m, v179)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L49
	} else {
		goto L56
	}
L56:
	;
	v187 = int32(0)
	v191 = F_RangeVarGetRelidExtended(m, v185, v187, int32(1), v187, v187)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L49
	} else {
		goto L57
	}
L57:
	;
	if v191 != 0 {
		v214 = v191
		goto L2
	} else {
		goto L58
	}
L58:
	;
	v193 = F_errsave_start(m, v10)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L49
	} else {
		goto L59
	}
L59:
	;
	if v193 == int32(0) {
		goto L3
	} else {
		goto L60
	}
L60:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L49
	} else {
		goto L61
	}
L61:
	;
	v200 = F_NameListToString(m, v179)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L49
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v200
	F_errmsg(m, int32(68954), v8)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L49
	} else {
		goto L63
	}
L63:
	;
	F_errsave_finish(m, v10, int32(480557), int32(914), int32(263940))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L49
	} else {
		goto L64
	}
L64:
	;
	goto L3
L65:
	;
	F_errmsg_internal(m, int32(395533), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L49
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(480557), int32(897), int32(263940))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L49
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_regcollationout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v10 == int32(0) {
		v14 = F_pstrdup(m, int32(632695))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v53 = v14
			m.G0 = v8 + int32(16)
			return v53
		}
	} else {
		v19 = F_SearchSysCache1(m, int32(16), v10)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			if v19 != 0 {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+22)))
				v23 = v21 + v22
				v25 = v23 + int32(4)
				v27 = *(*int32)(unsafe.Add(mBase, _consts[298]))
				if v27 == int32(0) {
					v30 = F_pstrdup(m, v25)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						F_ReleaseCatCache(m, v19)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v53 = v30
							m.G0 = v8 + int32(16)
							return v53
						}
					}
				} else {
					v34 = F_CollationIsVisible(m, v10)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						if v34 != 0 {
							v40 = int32(0)
							v41 = F_quote_qualified_identifier(m, v40, v25)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								F_ReleaseCatCache(m, v19)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int32(0)
								} else {
									v53 = v41
									m.G0 = v8 + int32(16)
									return v53
								}
							}
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v23)+68))
							v38 = F_get_namespace_name(m, v37)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v40 = v38
								v41 = F_quote_qualified_identifier(m, v40, v25)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									F_ReleaseCatCache(m, v19)
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return int32(0)
									} else {
										v53 = v41
										m.G0 = v8 + int32(16)
										return v53
									}
								}
							}
						}
					}
				}
			} else {
				v46 = F_palloc(m, int32(64))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
					v51 = F_pg_snprintf(m, v46, int32(64), int32(57422), v8)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						v53 = v46
						m.G0 = v8 + int32(16)
						return v53
					}
				}
			}
		}
	}
}
func F_register_ENR(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = F_lappend(m, v3, l1)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v4
		return
	}
}
func F_regoperatorin(m *base.Module, l0 int32) int32 {
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
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
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
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	v5 = m.G0
	v7 = v5 - int32(432)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if base.Ui32(int32(9)) < base.Ui32((v11-int32(48))&int32(255)) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L44
	} else {
		goto L73
	}
L2:
	;
	m.G0 = v7 + int32(432)
	return v254
L3:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _consts[298]))
	if v170 == int32(0) {
		goto L1
	} else {
		goto L46
	}
L4:
	;
	v18 = int32(527523)
	v22 = m.G0
	v24 = v22 - int32(32)
	v25 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+24)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(v24)+16)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(v24))) = v25
	v33 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1069])))
	if v33 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v10&int32(3) == int32(0) {
		v125 = v10
		goto L28
	} else {
		goto L29
	}
L6:
	;
	v101 = int32(0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1070])))
	if v37 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v41 = v10
	goto L12
L10:
	;
	goto L11
L11:
	;
	v51 = v18
	v52 = v33
	goto L15
L12:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v47 == v33 {
		v41 = v41 + int32(1)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v101 = v41 - v10
	goto L5
L14:
	;
	goto L13
L15:
	;
	v59 = v24 + int32(base.Ui32(v52)>>(uint(int32(3))%32))&int32(28)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v61 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = v60 | v61<<(uint(v52)%32)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	if v65 != 0 {
		v51 = v51 + v61
		v52 = v65
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v68 == int32(0) {
		v93 = v10
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	v101 = v93 - v10
	goto L5
L19:
	;
	v72 = v10
	v73 = v68
	goto L20
L20:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(base.Ui32(v73)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v81)>>(uint(v73)%32))&int32(1) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v93 = v89
	goto L18
L22:
	;
	v93 = v72
	goto L18
L23:
	;
	goto L24
L24:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
	v89 = v72 + int32(1)
	if v87 != 0 {
		v72 = v89
		v73 = v87
		goto L20
	} else {
		goto L25
	}
L25:
	;
	goto L21
L26:
	;
	if v101 != v158 {
		goto L3
	} else {
		goto L43
	}
L27:
	;
	v158 = v150 - v10
	goto L26
L28:
	;
	v129 = v125
	goto L37
L29:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v109 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v158 = int32(0)
	goto L26
L31:
	;
	goto L32
L32:
	;
	v114 = v10
	goto L33
L33:
	;
	v118 = v114 + int32(1)
	if v118&int32(3) == int32(0) {
		v125 = v118
		goto L28
	} else {
		goto L35
	}
L34:
	;
	v150 = v118
	goto L27
L35:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	if v123 != 0 {
		v114 = v118
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v138 = int32(-2139062144)
	if (int32(16843008)-v135|v135)&v138 == v138 {
		v129 = v129 + int32(4)
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v144 = v129
	goto L40
L39:
	;
	goto L38
L40:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	if v148 != 0 {
		v144 = v144 + int32(1)
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v150 = v144
	goto L27
L42:
	;
	goto L41
L43:
	;
	v164 = F_DirectInputFunctionCallSafe(m, int32(547), v10, int32(-1), v9, v7+int32(16))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	return int32(0)
L45:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v254 = v168
	goto L2
L46:
	;
	v180 = F_parseNameAndArgTypes(m, v10, int32(1), v7+int32(428), v7+int32(424), v7+int32(16), v9)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	if v180 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v184 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v184)
	v254 = int32(0)
	goto L2
L49:
	;
	goto L50
L50:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v7)+424))
	switch v187 - int32(1) {
	case 0:
		goto L53
	case 1:
		goto L51
	default:
		goto L52
	}
L51:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v7)+428))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	v235 = F_OpernameGetOprid(m, v232, v233, v234)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L44
	} else {
		goto L66
	}
L52:
	;
	v211 = int32(0)
	v212 = F_errsave_start(m, v9)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L44
	} else {
		goto L60
	}
L53:
	;
	v190 = int32(0)
	v191 = F_errsave_start(m, v9)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L44
	} else {
		goto L54
	}
L54:
	;
	if v191 == int32(0) {
		v254 = v190
		goto L2
	} else {
		goto L55
	}
L55:
	;
	F_errcode(m, int32(33685636))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L44
	} else {
		goto L56
	}
L56:
	;
	F_errmsg(m, int32(89038), int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L44
	} else {
		goto L57
	}
L57:
	;
	F_errhint(m, int32(570097), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L44
	} else {
		goto L58
	}
L58:
	;
	F_errsave_finish(m, v9, int32(480557), int32(671), int32(263999))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L44
	} else {
		goto L59
	}
L59:
	;
	v254 = v190
	goto L2
L60:
	;
	if v212 == int32(0) {
		v254 = v211
		goto L2
	} else {
		goto L61
	}
L61:
	;
	F_errcode(m, int32(50856197))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L44
	} else {
		goto L62
	}
L62:
	;
	F_errmsg(m, int32(114110), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L44
	} else {
		goto L63
	}
L63:
	;
	F_errhint(m, int32(570158), int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L44
	} else {
		goto L64
	}
L64:
	;
	F_errsave_finish(m, v9, int32(480557), int32(676), int32(263999))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L44
	} else {
		goto L65
	}
L65:
	;
	v254 = v211
	goto L2
L66:
	;
	if v235 != 0 {
		v254 = v235
		goto L2
	} else {
		goto L67
	}
L67:
	;
	v237 = int32(0)
	v238 = F_errsave_start(m, v9)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L44
	} else {
		goto L68
	}
L68:
	;
	if v238 == int32(0) {
		v254 = v237
		goto L2
	} else {
		goto L69
	}
L69:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L44
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10
	F_errmsg(m, int32(190516), v7)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L44
	} else {
		goto L71
	}
L71:
	;
	F_errsave_finish(m, v9, int32(480557), int32(683), int32(263999))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L44
	} else {
		goto L72
	}
L72:
	;
	v254 = v237
	goto L2
L73:
	;
	F_errmsg_internal(m, int32(395580), int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L44
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(480557), int32(654), int32(263999))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L44
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_regoperin(m *base.Module, l0 int32) int32 {
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
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if base.Ui32(int32(9)) < base.Ui32((v12-int32(48))&int32(255)) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L44
	} else {
		goto L68
	}
L2:
	;
	m.G0 = v8 + int32(32)
	return v224
L3:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _consts[298]))
	if v171 == int32(0) {
		goto L1
	} else {
		goto L46
	}
L4:
	;
	v19 = int32(527523)
	v23 = m.G0
	v25 = v23 - int32(32)
	v26 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+24)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v25)+16)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v25)+8)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v25))) = v26
	v34 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1069])))
	if v34 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v11&int32(3) == int32(0) {
		v126 = v11
		goto L28
	} else {
		goto L29
	}
L6:
	;
	v102 = int32(0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1070])))
	if v38 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v42 = v11
	goto L12
L10:
	;
	goto L11
L11:
	;
	v52 = v19
	v53 = v34
	goto L15
L12:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v48 == v34 {
		v42 = v42 + int32(1)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v102 = v42 - v11
	goto L5
L14:
	;
	goto L13
L15:
	;
	v60 = v25 + int32(base.Ui32(v53)>>(uint(int32(3))%32))&int32(28)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v62 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v61 | v62<<(uint(v53)%32)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	if v66 != 0 {
		v52 = v52 + v62
		v53 = v66
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v69 == int32(0) {
		v94 = v11
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	v102 = v94 - v11
	goto L5
L19:
	;
	v73 = v11
	v74 = v69
	goto L20
L20:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v25+int32(base.Ui32(v74)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v82)>>(uint(v74)%32))&int32(1) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v94 = v90
	goto L18
L22:
	;
	v94 = v73
	goto L18
L23:
	;
	goto L24
L24:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
	v90 = v73 + int32(1)
	if v88 != 0 {
		v73 = v90
		v74 = v88
		goto L20
	} else {
		goto L25
	}
L25:
	;
	goto L21
L26:
	;
	if v102 != v159 {
		goto L3
	} else {
		goto L43
	}
L27:
	;
	v159 = v151 - v11
	goto L26
L28:
	;
	v130 = v126
	goto L37
L29:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v110 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v159 = int32(0)
	goto L26
L31:
	;
	goto L32
L32:
	;
	v115 = v11
	goto L33
L33:
	;
	v119 = v115 + int32(1)
	if v119&int32(3) == int32(0) {
		v126 = v119
		goto L28
	} else {
		goto L35
	}
L34:
	;
	v151 = v119
	goto L27
L35:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	if v124 != 0 {
		v115 = v119
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v139 = int32(-2139062144)
	if (int32(16843008)-v136|v136)&v139 == v139 {
		v130 = v130 + int32(4)
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v145 = v130
	goto L40
L39:
	;
	goto L38
L40:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	if v149 != 0 {
		v145 = v145 + int32(1)
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v151 = v145
	goto L27
L42:
	;
	goto L41
L43:
	;
	v165 = F_DirectInputFunctionCallSafe(m, int32(547), v11, int32(-1), v10, v8+int32(28))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	return int32(0)
L45:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	v224 = v169
	goto L2
L46:
	;
	v174 = F_stringToQualifiedNameList(m, v11, v10)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	if v174 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v178 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v178)
	v224 = int32(0)
	goto L2
L49:
	;
	goto L50
L50:
	;
	v181 = int32(0)
	v184 = F_OpernameGetCandidates(m, v174, v181, int32(1))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L44
	} else {
		goto L51
	}
L51:
	;
	if v184 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v188 = F_errsave_start(m, v10)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L44
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	if v204 != 0 {
		goto L60
	} else {
		goto L61
	}
L55:
	;
	if v188 == int32(0) {
		v224 = v181
		goto L2
	} else {
		goto L56
	}
L56:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L44
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11
	F_errmsg(m, int32(190516), v8)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L44
	} else {
		goto L58
	}
L58:
	;
	F_errsave_finish(m, v10, int32(480557), int32(509), int32(264013))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L44
	} else {
		goto L59
	}
L59:
	;
	v224 = v181
	goto L2
L60:
	;
	v205 = F_errsave_start(m, v10)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L44
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v184)+8))
	v224 = v223
	goto L2
L63:
	;
	if v205 == int32(0) {
		v224 = v181
		goto L2
	} else {
		goto L64
	}
L64:
	;
	F_errcode(m, int32(84439172))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L44
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v11
	F_errmsg(m, int32(188441), v8+int32(16))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L44
	} else {
		goto L66
	}
L66:
	;
	F_errsave_finish(m, v10, int32(480557), int32(514), int32(264013))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L44
	} else {
		goto L67
	}
L67:
	;
	v224 = v181
	goto L2
L68:
	;
	F_errmsg_internal(m, int32(395630), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L44
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(480557), int32(494), int32(264013))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L44
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_regprocedureout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v2 == int32(0) {
		v6 = F_pstrdup(m, int32(632695))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v6
		}
	} else {
		v12 = F_format_procedure_extended(m, v2, int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			return v12
		}
	}
}
func F_regprocin(m *base.Module, l0 int32) int32 {
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v12 == int32(45) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L48
	} else {
		goto L72
	}
L2:
	;
	m.G0 = v8 + int32(32)
	return v232
L3:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _consts[298]))
	if v175 == int32(0) {
		goto L1
	} else {
		goto L50
	}
L4:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v15 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if base.Ui32(int32(9)) < base.Ui32((v12-int32(48))&int32(255)) {
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v232 = int32(0)
	goto L2
L8:
	;
	v23 = int32(527523)
	v27 = m.G0
	v29 = v27 - int32(32)
	v30 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+24)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = v30
	v38 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1069])))
	if v38 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v11&int32(3) == int32(0) {
		v130 = v11
		goto L32
	} else {
		goto L33
	}
L10:
	;
	v106 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1070])))
	if v42 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v46 = v11
	goto L16
L14:
	;
	goto L15
L15:
	;
	v56 = v23
	v57 = v38
	goto L19
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v52 == v38 {
		v46 = v46 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v106 = v46 - v11
	goto L9
L18:
	;
	goto L17
L19:
	;
	v64 = v29 + int32(base.Ui32(v57)>>(uint(int32(3))%32))&int32(28)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v66 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v65 | v66<<(uint(v57)%32)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	if v70 != 0 {
		v56 = v56 + v66
		v57 = v70
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v73 == int32(0) {
		v98 = v11
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v106 = v98 - v11
	goto L9
L23:
	;
	v77 = v11
	v78 = v73
	goto L24
L24:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(base.Ui32(v78)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v86)>>(uint(v78)%32))&int32(1) == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v98 = v94
	goto L22
L26:
	;
	v98 = v77
	goto L22
L27:
	;
	goto L28
L28:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	v94 = v77 + int32(1)
	if v92 != 0 {
		v77 = v94
		v78 = v92
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	if v106 != v163 {
		goto L3
	} else {
		goto L47
	}
L31:
	;
	v163 = v155 - v11
	goto L30
L32:
	;
	v134 = v130
	goto L41
L33:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v114 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v163 = int32(0)
	goto L30
L35:
	;
	goto L36
L36:
	;
	v119 = v11
	goto L37
L37:
	;
	v123 = v119 + int32(1)
	if v123&int32(3) == int32(0) {
		v130 = v123
		goto L32
	} else {
		goto L39
	}
L38:
	;
	v155 = v123
	goto L31
L39:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	if v128 != 0 {
		v119 = v123
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	v143 = int32(-2139062144)
	if (int32(16843008)-v140|v140)&v143 == v143 {
		v134 = v134 + int32(4)
		goto L41
	} else {
		goto L43
	}
L42:
	;
	v149 = v134
	goto L44
L43:
	;
	goto L42
L44:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v153 != 0 {
		v149 = v149 + int32(1)
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v155 = v149
	goto L31
L46:
	;
	goto L45
L47:
	;
	v169 = F_DirectInputFunctionCallSafe(m, int32(547), v11, int32(-1), v10, v8+int32(28))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	return int32(0)
L49:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	v232 = v173
	goto L2
L50:
	;
	v178 = F_stringToQualifiedNameList(m, v11, v10)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L48
	} else {
		goto L51
	}
L51:
	;
	if v178 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v182 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v182)
	v232 = int32(0)
	goto L2
L53:
	;
	goto L54
L54:
	;
	v185 = int32(0)
	v192 = F_FuncnameGetCandidates(m, v178, int32(-1), v185, v185, v185, v185, int32(1))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L48
	} else {
		goto L55
	}
L55:
	;
	if v192 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v196 = F_errsave_start(m, v10)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L48
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	if v212 != 0 {
		goto L64
	} else {
		goto L65
	}
L59:
	;
	if v196 == int32(0) {
		v232 = v185
		goto L2
	} else {
		goto L60
	}
L60:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L48
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11
	F_errmsg(m, int32(68619), v8)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L48
	} else {
		goto L62
	}
L62:
	;
	F_errsave_finish(m, v10, int32(480557), int32(100), int32(265956))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L48
	} else {
		goto L63
	}
L63:
	;
	v232 = v185
	goto L2
L64:
	;
	v213 = F_errsave_start(m, v10)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L48
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v192)+8))
	v232 = v231
	goto L2
L67:
	;
	if v213 == int32(0) {
		v232 = v185
		goto L2
	} else {
		goto L68
	}
L68:
	;
	F_errcode(m, int32(84439172))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L48
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v11
	F_errmsg(m, int32(685699), v8+int32(16))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L48
	} else {
		goto L70
	}
L70:
	;
	F_errsave_finish(m, v10, int32(480557), int32(105), int32(265956))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L48
	} else {
		goto L71
	}
L71:
	;
	v232 = v185
	goto L2
L72:
	;
	F_errmsg_internal(m, int32(395969), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L48
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(480557), int32(85), int32(265956))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L48
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_remove_self_joins_recurse(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v52 int32
	_ = v52
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v134 int32
	_ = v134
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v344 int32
	_ = v344
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v670 int32
	_ = v670
	var v677 int32
	_ = v677
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v810 int32
	_ = v810
	var v818 int32
	_ = v818
	var v825 int32
	_ = v825
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v877 int32
	_ = v877
	var v884 int32
	_ = v884
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v939 int32
	_ = v939
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1026 int32
	_ = v1026
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1069 int32
	_ = v1069
	var v1072 int32
	_ = v1072
	var v1076 int32
	_ = v1076
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1090 int32
	_ = v1090
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1185 int32
	_ = v1185
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1192 int32
	_ = v1192
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1246 int32
	_ = v1246
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1265 int32
	_ = v1265
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1279 int32
	_ = v1279
	var v1304 int32
	_ = v1304
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1319 int32
	_ = v1319
	var v1321 int32
	_ = v1321
	var v1323 int32
	_ = v1323
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1393 int32
	_ = v1393
	var v1395 int32
	_ = v1395
	var v1398 int32
	_ = v1398
	var v1402 int32
	_ = v1402
	var v1405 int32
	_ = v1405
	var v1408 int32
	_ = v1408
	var v1429 int32
	_ = v1429
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1455 int32
	_ = v1455
	var v1458 int32
	_ = v1458
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1475 int32
	_ = v1475
	var v1478 int32
	_ = v1478
	var v1482 int32
	_ = v1482
	var v1489 int32
	_ = v1489
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1506 int32
	_ = v1506
	var v1509 int32
	_ = v1509
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1547 int32
	_ = v1547
	var v1568 int32
	_ = v1568
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1591 int32
	_ = v1591
	var v1594 int32
	_ = v1594
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1611 int32
	_ = v1611
	var v1614 int32
	_ = v1614
	var v1618 int32
	_ = v1618
	var v1625 int32
	_ = v1625
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1635 int32
	_ = v1635
	var v1636 int32
	_ = v1636
	var v1642 int32
	_ = v1642
	var v1645 int32
	_ = v1645
	var v1669 int32
	_ = v1669
	var v1673 int32
	_ = v1673
	var v1675 int32
	_ = v1675
	var v1677 int32
	_ = v1677
	var v1678 int32
	_ = v1678
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1692 int32
	_ = v1692
	var v1696 int32
	_ = v1696
	var v1699 int32
	_ = v1699
	var v1701 int32
	_ = v1701
	var v1704 int32
	_ = v1704
	var v1711 int32
	_ = v1711
	var v1713 int32
	_ = v1713
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1735 int32
	_ = v1735
	var v1750 int32
	_ = v1750
	var v1766 int32
	_ = v1766
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1771 int32
	_ = v1771
	var v1772 int32
	_ = v1772
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1780 int32
	_ = v1780
	var v1781 int32
	_ = v1781
	var v1785 int32
	_ = v1785
	var v1793 int32
	_ = v1793
	var v1812 int32
	_ = v1812
	var v1816 int32
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1832 int32
	_ = v1832
	var v1835 int32
	_ = v1835
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1845 int32
	_ = v1845
	var v1870 int32
	_ = v1870
	var v1871 int32
	_ = v1871
	var v1875 int32
	_ = v1875
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1879 int32
	_ = v1879
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1882 int32
	_ = v1882
	var v1884 int32
	_ = v1884
	var v1885 int32
	_ = v1885
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1918 int32
	_ = v1918
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1949 int32
	_ = v1949
	var v1951 int32
	_ = v1951
	var v1978 int32
	_ = v1978
	var v1980 int32
	_ = v1980
	var v1983 int32
	_ = v1983
	var v1984 int32
	_ = v1984
	var v1985 int32
	_ = v1985
	var v1988 int32
	_ = v1988
	var v1989 int32
	_ = v1989
	var v1993 int32
	_ = v1993
	var v2001 int32
	_ = v2001
	var v2020 int32
	_ = v2020
	var v2024 int32
	_ = v2024
	var v2025 int32
	_ = v2025
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2032 int32
	_ = v2032
	var v2035 int32
	_ = v2035
	var v2036 int32
	_ = v2036
	var v2042 int32
	_ = v2042
	var v2067 int32
	_ = v2067
	var v2068 int32
	_ = v2068
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2074 int32
	_ = v2074
	var v2075 int32
	_ = v2075
	var v2076 int32
	_ = v2076
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2079 int32
	_ = v2079
	var v2081 int32
	_ = v2081
	var v2082 int32
	_ = v2082
	var v2112 int32
	_ = v2112
	var v2113 int32
	_ = v2113
	var v2115 int32
	_ = v2115
	var v2143 int32
	_ = v2143
	var v2144 int32
	_ = v2144
	var v2146 int32
	_ = v2146
	var v2148 int32
	_ = v2148
	var v2167 int32
	_ = v2167
	var v2176 int32
	_ = v2176
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2180 int32
	_ = v2180
	var v2182 int32
	_ = v2182
	var v2183 int32
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2186 int32
	_ = v2186
	var v2193 int32
	_ = v2193
	var v2195 int32
	_ = v2195
	var v2196 int32
	_ = v2196
	var v2199 int32
	_ = v2199
	var v2203 int32
	_ = v2203
	var v2206 int32
	_ = v2206
	var v2208 int32
	_ = v2208
	var v2211 int32
	_ = v2211
	var v2218 int32
	_ = v2218
	var v2220 int32
	_ = v2220
	var v2228 int32
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2242 int32
	_ = v2242
	var v2273 int32
	_ = v2273
	var v2274 int32
	_ = v2274
	var v2277 int32
	_ = v2277
	var v2278 int32
	_ = v2278
	var v2284 int32
	_ = v2284
	var v2309 int32
	_ = v2309
	var v2313 int32
	_ = v2313
	var v2314 int32
	_ = v2314
	var v2315 int32
	_ = v2315
	var v2318 int32
	_ = v2318
	var v2319 int32
	_ = v2319
	var v2320 int32
	_ = v2320
	var v2321 int32
	_ = v2321
	var v2322 int32
	_ = v2322
	var v2325 int32
	_ = v2325
	var v2326 int32
	_ = v2326
	var v2327 int32
	_ = v2327
	var v2328 int32
	_ = v2328
	var v2329 int32
	_ = v2329
	var v2333 int32
	_ = v2333
	var v2334 int32
	_ = v2334
	var v2364 int32
	_ = v2364
	var v2365 int32
	_ = v2365
	var v2368 int32
	_ = v2368
	var v2395 int32
	_ = v2395
	var v2398 int32
	_ = v2398
	var v2399 int32
	_ = v2399
	var v2401 int32
	_ = v2401
	var v2402 int32
	_ = v2402
	var v2403 int32
	_ = v2403
	var v2404 int32
	_ = v2404
	var v2405 int32
	_ = v2405
	var v2406 int32
	_ = v2406
	var v2409 int32
	_ = v2409
	var v2411 int32
	_ = v2411
	var v2412 int32
	_ = v2412
	var v2414 int32
	_ = v2414
	var v2415 int32
	_ = v2415
	var v2416 int32
	_ = v2416
	var v2417 int32
	_ = v2417
	var v2420 int32
	_ = v2420
	var v2454 int32
	_ = v2454
	var v2455 int32
	_ = v2455
	var v2456 int32
	_ = v2456
	var v2458 int32
	_ = v2458
	var v2462 int32
	_ = v2462
	var v2463 int32
	_ = v2463
	var v2464 int32
	_ = v2464
	var v2467 int32
	_ = v2467
	var v2468 int32
	_ = v2468
	var v2469 int32
	_ = v2469
	var v2472 int32
	_ = v2472
	var v2473 int32
	_ = v2473
	var v2474 int32
	_ = v2474
	var v2475 int32
	_ = v2475
	var v2478 int32
	_ = v2478
	var v2479 int32
	_ = v2479
	var v2480 int32
	_ = v2480
	var v2481 int32
	_ = v2481
	var v2484 int32
	_ = v2484
	var v2485 int32
	_ = v2485
	var v2486 int32
	_ = v2486
	var v2487 int32
	_ = v2487
	var v2488 int32
	_ = v2488
	var v2489 int32
	_ = v2489
	var v2490 int32
	_ = v2490
	var v2491 int32
	_ = v2491
	var v2492 int32
	_ = v2492
	var v2493 int32
	_ = v2493
	var v2494 int32
	_ = v2494
	var v2495 int32
	_ = v2495
	var v2496 int32
	_ = v2496
	var v2497 int32
	_ = v2497
	var v2500 int32
	_ = v2500
	var v2502 int32
	_ = v2502
	var v2503 int32
	_ = v2503
	var v2510 int32
	_ = v2510
	var v2512 int32
	_ = v2512
	var v2514 int32
	_ = v2514
	var v2516 int32
	_ = v2516
	var v2518 int32
	_ = v2518
	var v2519 int32
	_ = v2519
	var v2520 int32
	_ = v2520
	var v2555 int32
	_ = v2555
	var v2557 int32
	_ = v2557
	var v2558 int32
	_ = v2558
	var v2561 int32
	_ = v2561
	var v2565 int32
	_ = v2565
	var v2568 int32
	_ = v2568
	var v2570 int32
	_ = v2570
	var v2573 int32
	_ = v2573
	var v2580 int32
	_ = v2580
	var v2582 int32
	_ = v2582
	var v2590 int32
	_ = v2590
	var v2591 int32
	_ = v2591
	var v2604 int32
	_ = v2604
	var v2631 int32
	_ = v2631
	var v2641 int32
	_ = v2641
	var v2643 int32
	_ = v2643
	var v2644 int32
	_ = v2644
	var v2647 int32
	_ = v2647
	var v2651 int32
	_ = v2651
	var v2654 int32
	_ = v2654
	var v2656 int32
	_ = v2656
	var v2659 int32
	_ = v2659
	var v2666 int32
	_ = v2666
	var v2668 int32
	_ = v2668
	var v2676 int32
	_ = v2676
	var v2677 int32
	_ = v2677
	var v2690 int32
	_ = v2690
	var v2693 int32
	_ = v2693
	var v2695 int32
	_ = v2695
	var v2708 int32
	_ = v2708
	var v2709 int32
	_ = v2709
	var v2710 int32
	_ = v2710
	var v2714 int32
	_ = v2714
	var v2715 int32
	_ = v2715
	var v2716 int32
	_ = v2716
	var v2717 int32
	_ = v2717
	var v2718 int32
	_ = v2718
	var v2721 int32
	_ = v2721
	var v2722 int32
	_ = v2722
	var v2723 int32
	_ = v2723
	var v2724 int32
	_ = v2724
	var v2733 int32
	_ = v2733
	var v2734 int32
	_ = v2734
	var v2737 int32
	_ = v2737
	var v2740 int32
	_ = v2740
	var v2742 int32
	_ = v2742
	var v2743 int32
	_ = v2743
	var v2751 int32
	_ = v2751
	var v2752 int32
	_ = v2752
	var v2753 int32
	_ = v2753
	var v2757 int32
	_ = v2757
	var v2760 int32
	_ = v2760
	var v2764 int32
	_ = v2764
	var v2771 int32
	_ = v2771
	var v2775 int32
	_ = v2775
	var v2777 int32
	_ = v2777
	var v2781 int32
	_ = v2781
	var v2782 int32
	_ = v2782
	var v2783 int32
	_ = v2783
	var v2784 int32
	_ = v2784
	var v2786 int32
	_ = v2786
	var v2800 int32
	_ = v2800
	var v2801 int32
	_ = v2801
	var v2805 int32
	_ = v2805
	var v2806 int32
	_ = v2806
	var v2807 int32
	_ = v2807
	var v2809 int32
	_ = v2809
	var v2817 int32
	_ = v2817
	var v2837 int32
	_ = v2837
	v4 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(16)
	m.G0 = v31
	if l1 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v146 = int32(0)
	if v134 == v146 {
		goto L30
	} else {
		goto L31
	}
L2:
	;
	v38 = l2
	v39 = v4
	v52 = v4
	goto L7
L3:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v33 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v120 = l2
	v134 = v4
	goto L1
L6:
	;
	goto L5
L7:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+v39<<(uint(int32(2))%32))))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v69 != int32(1) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v120 = v110
	v134 = v113
	goto L1
L9:
	;
	v115 = v39 + int32(1)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v115 < v116 {
		v38 = v110
		v39 = v115
		v52 = v113
		goto L7
	} else {
		goto L27
	}
L10:
	;
	if v69 == int32(63) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v108 = F_remove_self_joins_recurse(m, l0, v68, v38)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L21
	} else {
		goto L26
	}
L13:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v74+v75<<(uint(int32(2))%32))))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	if v80 != 0 {
		v110 = v38
		v113 = v52
		goto L9
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L21
	} else {
		goto L23
	}
L16:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+21)))
	if v81 != int32(114) {
		v110 = v38
		v113 = v52
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v79)+32))
	if v84 != 0 {
		v110 = v38
		v113 = v52
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+32))
	if v75 == v86 {
		v110 = v38
		v113 = v52
		goto L9
	} else {
		goto L19
	}
L19:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+68))
	if v75 == v88 {
		v110 = v38
		v113 = v52
		goto L9
	} else {
		goto L20
	}
L20:
	;
	v90 = F_bms_add_member(m, v52, v75)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(0)
L22:
	;
	v110 = v38
	v113 = v90
	goto L9
L23:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v98
	F_errmsg_internal(m, int32(467113), v31)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(474751), int32(2352), int32(346090))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	v110 = v108
	v113 = v52
	goto L9
L27:
	;
	goto L8
L28:
	;
	m.G0 = v2837 + int32(16)
	return v2817
L29:
	;
	if v181 < int32(2) {
		v2817 = v120
		v2837 = v31
		goto L28
	} else {
		goto L42
	}
L30:
	;
	v181 = int32(0)
	goto L29
L31:
	;
	goto L32
L32:
	;
	v153 = int32(1)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v134)+4))
	if v154 <= v153 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v157 = v153
	goto L35
L34:
	;
	v157 = v154
	goto L35
L35:
	;
	v161 = int32(0)
	v163 = v146
	goto L36
L36:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v134+int32(8)+v161<<(uint(int32(2))%32))))
	if v169 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v181 = v172
	goto L29
L38:
	;
	v172 = v163 + base.I32_popcnt(v169)
	goto L40
L39:
	;
	v172 = v163
	goto L40
L40:
	;
	v174 = v161 + int32(1)
	if v174 != v157 {
		v161 = v174
		v163 = v172
		goto L36
	} else {
		goto L41
	}
L41:
	;
	goto L37
L42:
	;
	v187 = F_palloc(m, v181<<(uint(int32(3))%32))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L21
	} else {
		goto L43
	}
L43:
	;
	if v134 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	if int32(0) <= v245 {
		goto L55
	} else {
		goto L56
	}
L45:
	;
	v245 = base.I32_ctz(v231) | v232<<(uint(int32(5))%32)
	goto L44
L46:
	;
	v245 = int32(-2)
	goto L44
L47:
	;
	v198 = base.I32_div_s(int32(0), int32(32))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v134)+4))
	if v199 <= v198 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v202 = v134 + int32(8)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v202+v198<<(uint(int32(2))%32))))
	v209 = v206 & int32(-1)
	if v209 != 0 {
		v231 = v209
		v232 = v198
		goto L45
	} else {
		goto L49
	}
L49:
	;
	v211 = v198 + int32(1)
	if v211 == v199 {
		goto L46
	} else {
		goto L50
	}
L50:
	;
	v214 = v211
	goto L51
L51:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v202+v214<<(uint(int32(2))%32))))
	if v221 != 0 {
		v231 = v221
		v232 = v214
		goto L45
	} else {
		goto L53
	}
L52:
	;
	goto L46
L53:
	;
	v223 = v214 + int32(1)
	if v223 != v199 {
		v214 = v223
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v249 = int32(0)
	v251 = v245
	goto L58
L56:
	;
	goto L57
L57:
	;
	F_pg_qsort(m, v187, v181, int32(8), int32(828))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L21
	} else {
		goto L72
	}
L58:
	;
	v278 = v187 + v249<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v278))) = v251
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v280+v251<<(uint(int32(2))%32))))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v278)+4)) = v285
	if v134 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	goto L57
L60:
	;
	if int32(0) <= v344 {
		v249 = v249 + int32(1)
		v251 = v344
		goto L58
	} else {
		goto L71
	}
L61:
	;
	v344 = base.I32_ctz(v330) | v331<<(uint(int32(5))%32)
	goto L60
L62:
	;
	v344 = int32(-2)
	goto L60
L63:
	;
	v295 = v251 + int32(1)
	v297 = base.I32_div_s(v295, int32(32))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v134)+4))
	if v298 <= v297 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v301 = v134 + int32(8)
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v301+v297<<(uint(int32(2))%32))))
	v308 = v305 & (int32(-1) << (uint(v295) % 32))
	if v308 != 0 {
		v330 = v308
		v331 = v297
		goto L61
	} else {
		goto L65
	}
L65:
	;
	v310 = v297 + int32(1)
	if v310 == v298 {
		goto L62
	} else {
		goto L66
	}
L66:
	;
	v313 = v310
	goto L67
L67:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v301+v313<<(uint(int32(2))%32))))
	if v320 != 0 {
		v330 = v320
		v331 = v313
		goto L61
	} else {
		goto L69
	}
L68:
	;
	goto L62
L69:
	;
	v322 = v313 + int32(1)
	if v322 != v298 {
		v313 = v322
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	goto L59
L72:
	;
	if v181+int32(1) < int32(2) {
		v2817 = v120
		v2837 = v31
		goto L28
	} else {
		goto L73
	}
L73:
	;
	v384 = l0
	v386 = v120
	v387 = int32(1)
	v400 = v134
	v401 = v4
	v406 = v31
	v407 = v187
	v409 = v181
	goto L74
L74:
	;
	if v409 != v387 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	v2817 = v2786
	v2837 = v2806
	goto L28
L76:
	;
	if v2805 != v2809 {
		v384 = v2784
		v386 = v2786
		v387 = v2805 + int32(1)
		v400 = v2800
		v401 = v2801
		v406 = v2806
		v407 = v2807
		v409 = v2809
		goto L74
	} else {
		goto L537
	}
L77:
	;
	v413 = int32(3)
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v407+v387<<(uint(v413)%32))+4))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v407+v401<<(uint(v413)%32))+4))
	if v416 == v420 {
		v2784 = v384
		v2786 = v386
		v2800 = v400
		v2801 = v401
		v2805 = v387
		v2806 = v406
		v2807 = v407
		v2809 = v409
		goto L76
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	if int32(2) <= v387-v401 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	goto L79
L81:
	;
	v425 = int32(0)
	if v401 < v387 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	goto L83
L83:
	;
	v2781 = *(*int32)(unsafe.Add(mBase, uint32(v407+v401<<(uint(int32(3))%32))))
	v2782 = F_bms_del_member(m, v400, v2781)
	mBase = m.M
	v2783 = m.ExcPending
	if v2783 != 0 {
		goto L21
	} else {
		goto L536
	}
L84:
	;
	v442 = v425
	v444 = v401
	goto L87
L85:
	;
	v479 = v425
	v481 = v401
	goto L86
L86:
	;
	v492 = F_bms_del_members(m, v400, v479)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L21
	} else {
		goto L91
	}
L87:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v407+v444<<(uint(int32(3))%32))))
	v459 = F_bms_add_member(m, v442, v458)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L21
	} else {
		goto L89
	}
L88:
	;
	v479 = v459
	v481 = v387
	goto L86
L89:
	;
	v462 = v444 + int32(1)
	if v462 != v387 {
		v442 = v459
		v444 = v462
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	v494 = v384
	v496 = v386
	v509 = v479
	v510 = v492
	v511 = v481
	v515 = v387
	v516 = v406
	v517 = v407
	v519 = v409
	goto L92
L92:
	;
	v522 = int32(0)
	if v509 == v522 {
		goto L96
	} else {
		goto L97
	}
L93:
	;
	F_bms_free(m, v2717)
	mBase = m.M
	v2775 = m.ExcPending
	if v2775 != 0 {
		goto L21
	} else {
		goto L534
	}
L94:
	;
	if int32(0) < v579 {
		goto L105
	} else {
		goto L106
	}
L95:
	;
	v579 = base.I32_ctz(v565) | v566<<(uint(int32(5))%32)
	goto L94
L96:
	;
	v579 = int32(-2)
	goto L94
L97:
	;
	v532 = base.I32_div_s(int32(0), int32(32))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v509)+4))
	if v533 <= v532 {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v536 = v509 + int32(8)
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v536+v532<<(uint(int32(2))%32))))
	v543 = v540 & int32(-1)
	if v543 != 0 {
		v565 = v543
		v566 = v532
		goto L95
	} else {
		goto L99
	}
L99:
	;
	v545 = v532 + int32(1)
	if v545 == v533 {
		goto L96
	} else {
		goto L100
	}
L100:
	;
	v548 = v545
	goto L101
L101:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v536+v548<<(uint(int32(2))%32))))
	if v555 != 0 {
		v565 = v555
		v566 = v548
		goto L95
	} else {
		goto L103
	}
L102:
	;
	goto L96
L103:
	;
	v557 = v548 + int32(1)
	if v557 != v533 {
		v548 = v557
		goto L101
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	v582 = v494
	v584 = v496
	v597 = v509
	v598 = v510
	v599 = v511
	v601 = v579
	v603 = v515
	v604 = v516
	v605 = v517
	v606 = v522
	v607 = v519
	goto L108
L106:
	;
	v2693 = v494
	v2695 = v496
	v2708 = v509
	v2709 = v510
	v2710 = v511
	v2714 = v515
	v2715 = v516
	v2716 = v517
	v2717 = v522
	v2718 = v519
	goto L107
L107:
	;
	v2721 = F_bms_add_members(m, v2695, v2717)
	mBase = m.M
	v2722 = m.ExcPending
	if v2722 != 0 {
		goto L21
	} else {
		goto L512
	}
L108:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v582)+28))
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v610+v601<<(uint(int32(2))%32))))
	if v597 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L109:
	;
	v2693 = v582
	v2695 = v584
	v2708 = v597
	v2709 = v598
	v2710 = v599
	v2714 = v603
	v2715 = v604
	v2716 = v605
	v2717 = v2631
	v2718 = v607
	goto L107
L110:
	;
	if v597 == int32(0) {
		goto L502
	} else {
		goto L503
	}
L111:
	;
	if v670 <= int32(0) {
		v2631 = v606
		goto L110
	} else {
		goto L122
	}
L112:
	;
	v670 = base.I32_ctz(v656) | v657<<(uint(int32(5))%32)
	goto L111
L113:
	;
	v670 = int32(-2)
	goto L111
L114:
	;
	v621 = v601 + int32(1)
	v623 = base.I32_div_s(v621, int32(32))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v597)+4))
	if v624 <= v623 {
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v627 = v597 + int32(8)
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v627+v623<<(uint(int32(2))%32))))
	v634 = v631 & (int32(-1) << (uint(v621) % 32))
	if v634 != 0 {
		v656 = v634
		v657 = v623
		goto L112
	} else {
		goto L116
	}
L116:
	;
	v636 = v623 + int32(1)
	if v636 == v624 {
		goto L113
	} else {
		goto L117
	}
L117:
	;
	v639 = v636
	goto L118
L118:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v627+v639<<(uint(int32(2))%32))))
	if v646 != 0 {
		v656 = v646
		v657 = v639
		goto L112
	} else {
		goto L120
	}
L119:
	;
	goto L113
L120:
	;
	v648 = v639 + int32(1)
	if v648 != v624 {
		v639 = v648
		goto L118
	} else {
		goto L121
	}
L121:
	;
	goto L119
L122:
	;
	v677 = v670
	goto L123
L123:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v582)+28))
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v701+v677<<(uint(int32(2))%32))))
	v706 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v604)+12)) = v706
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v582)+112))
	if v708 == v706 {
		goto L126
	} else {
		goto L127
	}
L124:
	;
	v2631 = v606
	goto L110
L125:
	;
	if v597 == int32(0) {
		goto L490
	} else {
		goto L491
	}
L126:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v582)+136))
	if v794 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L127:
	;
	v711 = int32(0)
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v708)+4))
	if v712 <= v711 {
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v716 = v711
	goto L129
L129:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v708)+12))
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v743+v716<<(uint(int32(2))%32))))
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v747)+12))
	v749 = F_bms_is_member(m, v677, v748)
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L21
	} else {
		goto L131
	}
L130:
	;
	goto L126
L131:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v747)+12))
	v752 = F_bms_is_member(m, v601, v751)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L21
	} else {
		goto L132
	}
L132:
	;
	if v749 != v752 {
		goto L125
	} else {
		goto L133
	}
L133:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v747)+16))
	v756 = F_bms_is_member(m, v677, v755)
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L21
	} else {
		goto L134
	}
L134:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v747)+16))
	v759 = F_bms_is_member(m, v601, v758)
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L21
	} else {
		goto L135
	}
L135:
	;
	if v756 != v759 {
		goto L125
	} else {
		goto L136
	}
L136:
	;
	v763 = v716 + int32(1)
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v708)+4))
	if v763 < v764 {
		v716 = v763
		goto L129
	} else {
		goto L137
	}
L137:
	;
	goto L130
L138:
	;
	v895 = F_bms_add_member(m, int32(0), v601)
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L21
	} else {
		goto L164
	}
L139:
	;
	v797 = int32(0)
	v877 = v797
	v884 = v797
	goto L138
L140:
	;
	goto L141
L141:
	;
	v799 = int32(0)
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v794)+4))
	if v799 < v800 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v804 = v800
	goto L144
L143:
	;
	v804 = v799
	goto L144
L144:
	;
	v805 = int32(0)
	v810 = v799
	v818 = v805
	v825 = v805
	goto L145
L145:
	;
	if v810 != v804 {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	if v856 == int32(0) {
		v877 = v856
		v884 = v858
		goto L138
	} else {
		goto L161
	}
L147:
	;
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v794)+12))
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v836+v810<<(uint(int32(2))%32))))
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v840)+4))
	v842 = base.B2i32(v841 == v601)
	if v841 == v601 {
		goto L150
	} else {
		goto L151
	}
L148:
	;
	v856 = v818
	v858 = v825
	goto L149
L149:
	;
	goto L146
L150:
	;
	v843 = v840
	goto L152
L151:
	;
	v843 = v825
	goto L152
L152:
	;
	v845 = v810 + int32(1)
	if v677 == v841 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v847 = v840
	goto L155
L154:
	;
	v847 = v818
	goto L155
L155:
	;
	if v841 == v601 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v848 = v818
	goto L158
L157:
	;
	v848 = v847
	goto L158
L158:
	;
	if v848 == int32(0) {
		v810 = v845
		v818 = v848
		v825 = v843
		goto L145
	} else {
		goto L159
	}
L159:
	;
	if v843 == int32(0) {
		v810 = v845
		v818 = v848
		v825 = v843
		goto L145
	} else {
		goto L160
	}
L160:
	;
	v856 = v848
	v858 = v843
	goto L149
L161:
	;
	if v858 == int32(0) {
		v877 = v856
		v884 = v858
		goto L138
	} else {
		goto L162
	}
L162:
	;
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v856)+16))
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v858)+16))
	if v863 != v864 {
		goto L125
	} else {
		goto L163
	}
L163:
	;
	v877 = v856
	v884 = v858
	goto L138
L164:
	;
	v897 = F_bms_add_member(m, v895, v677)
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L21
	} else {
		goto L165
	}
L165:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v614)+8))
	v901 = F_generate_join_implied_equalities(m, v582, v897, v899, v705, int32(0))
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L21
	} else {
		goto L166
	}
L166:
	;
	if v901 == int32(0) {
		goto L125
	} else {
		goto L167
	}
L167:
	;
	v905 = int32(0)
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v901)+4))
	if v905 < v908 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v912 = v905
	v918 = v905
	v921 = v905
	goto L171
L169:
	;
	v1155 = v905
	v1158 = v905
	goto L170
L170:
	;
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v705)+184))
	v1177 = F_list_concat(m, v1158, v1176)
	mBase = m.M
	v1178 = m.ExcPending
	if v1178 != 0 {
		goto L21
	} else {
		goto L244
	}
L171:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v901)+12))
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v939+v912<<(uint(int32(2))%32))))
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v943)+96))
	if v944 == int32(0) {
		goto L174
	} else {
		goto L175
	}
L172:
	;
	v1155 = v1138
	v1158 = v1140
	goto L170
L173:
	;
	v1145 = v912 + int32(1)
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v901)+4))
	if v1145 < v1146 {
		v912 = v1145
		v918 = v1138
		v921 = v1140
		goto L171
	} else {
		goto L243
	}
L174:
	;
	v1136 = F_lappend(m, v918, v943)
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L21
	} else {
		goto L242
	}
L175:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v943)+28))
	v948 = int32(0)
	if v947 == v948 {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	if v983 != int32(2) {
		goto L174
	} else {
		goto L189
	}
L177:
	;
	v983 = int32(0)
	goto L176
L178:
	;
	goto L179
L179:
	;
	v955 = int32(1)
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v947)+4))
	if v956 <= v955 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v959 = v955
	goto L182
L181:
	;
	v959 = v956
	goto L182
L182:
	;
	v963 = int32(0)
	v965 = v948
	goto L183
L183:
	;
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v947+int32(8)+v963<<(uint(int32(2))%32))))
	if v971 != 0 {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	v983 = v974
	goto L176
L185:
	;
	v974 = v965 + base.I32_popcnt(v971)
	goto L187
L186:
	;
	v974 = v965
	goto L187
L187:
	;
	v976 = v963 + int32(1)
	if v976 != v959 {
		v963 = v976
		v965 = v974
		goto L183
	} else {
		goto L188
	}
L188:
	;
	goto L184
L189:
	;
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v943)+44))
	if v986 == int32(0) {
		goto L191
	} else {
		goto L192
	}
L190:
	;
	if v1033 != int32(1) {
		goto L174
	} else {
		goto L206
	}
L191:
	;
	v1033 = int32(0)
	goto L190
L192:
	;
	goto L193
L193:
	;
	v995 = int32(1)
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v986)+4))
	if v996 <= v995 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v999 = v995
	goto L196
L195:
	;
	v999 = v996
	goto L196
L196:
	;
	v1002 = int32(0)
	v1004 = v1002
	v1005 = v1002
	goto L197
L197:
	;
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v986+int32(8)+v1004<<(uint(int32(2))%32))))
	if v1013 != 0 {
		goto L200
	} else {
		goto L201
	}
L198:
	;
	v1033 = v1026
	goto L190
L199:
	;
	goto L198
L200:
	;
	v1014 = int32(2)
	if v1005 != 0 {
		v1026 = v1014
		goto L199
	} else {
		goto L203
	}
L201:
	;
	v1019 = v1005
	goto L202
L202:
	;
	v1022 = v1004 + int32(1)
	if v1022 != v999 {
		v1004 = v1022
		v1005 = v1019
		goto L197
	} else {
		goto L205
	}
L203:
	;
	v1015 = int32(1)
	if base.Ui32(v1015) < base.Ui32(base.I32_popcnt(v1013)) {
		v1026 = v1014
		goto L199
	} else {
		goto L204
	}
L204:
	;
	v1019 = v1015
	goto L202
L205:
	;
	v1026 = v1019
	goto L199
L206:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v943)+48))
	if v1036 == int32(0) {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	if v1083 != int32(1) {
		goto L174
	} else {
		goto L223
	}
L208:
	;
	v1083 = int32(0)
	goto L207
L209:
	;
	goto L210
L210:
	;
	v1045 = int32(1)
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v1036)+4))
	if v1046 <= v1045 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v1049 = v1045
	goto L213
L212:
	;
	v1049 = v1046
	goto L213
L213:
	;
	v1052 = int32(0)
	v1054 = v1052
	v1055 = v1052
	goto L214
L214:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v1036+int32(8)+v1054<<(uint(int32(2))%32))))
	if v1063 != 0 {
		goto L217
	} else {
		goto L218
	}
L215:
	;
	v1083 = v1076
	goto L207
L216:
	;
	goto L215
L217:
	;
	v1064 = int32(2)
	if v1055 != 0 {
		v1076 = v1064
		goto L216
	} else {
		goto L220
	}
L218:
	;
	v1069 = v1055
	goto L219
L219:
	;
	v1072 = v1054 + int32(1)
	if v1072 != v1049 {
		v1054 = v1072
		v1055 = v1069
		goto L214
	} else {
		goto L222
	}
L220:
	;
	v1065 = int32(1)
	if base.Ui32(v1065) < base.Ui32(base.I32_popcnt(v1063)) {
		v1076 = v1064
		goto L216
	} else {
		goto L221
	}
L221:
	;
	v1069 = v1065
	goto L219
L222:
	;
	v1076 = v1069
	goto L216
L223:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v943)+4))
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1086)))
	if v1087 != int32(17) {
		goto L174
	} else {
		goto L224
	}
L224:
	;
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v1086)+28))
	if v1090 == int32(0) {
		goto L174
	} else {
		goto L225
	}
L225:
	;
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+4))
	if v1093 != int32(2) {
		goto L174
	} else {
		goto L226
	}
L226:
	;
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+12))
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1096)))
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1096)+4))
	v1100 = F_copyObjectImpl(m, v1099)
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L21
	} else {
		goto L227
	}
L227:
	;
	v1102 = int32(0)
	if v1097 == v1102 {
		v1109 = v1102
		goto L228
	} else {
		goto L229
	}
L228:
	;
	if v1100 == int32(0) {
		v1116 = int32(0)
		goto L231
	} else {
		goto L232
	}
L229:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v1097)))
	if v1105 != int32(27) {
		v1109 = v1097
		goto L228
	} else {
		goto L230
	}
L230:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v1097)+4))
	v1109 = v1108
	goto L228
L231:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v943)+48))
	v1118 = F_bms_singleton_member(m, v1117)
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		goto L21
	} else {
		goto L236
	}
L232:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1100)))
	if v1112 != int32(27) {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v1116 = v1100
	goto L231
L234:
	;
	goto L235
L235:
	;
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v1100)+4))
	v1116 = v1115
	goto L231
L236:
	;
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v943)+44))
	v1121 = F_bms_singleton_member(m, v1120)
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L21
	} else {
		goto L237
	}
L237:
	;
	F_ChangeVarNodesExtended(m, v1116, v1118, v1121, int32(827))
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L21
	} else {
		goto L238
	}
L238:
	;
	v1126 = F_equal(m, v1109, v1116)
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L21
	} else {
		goto L239
	}
L239:
	;
	if v1126 == int32(0) {
		goto L174
	} else {
		goto L240
	}
L240:
	;
	v1130 = F_lappend(m, v921, v943)
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L21
	} else {
		goto L241
	}
L241:
	;
	v1138 = v918
	v1140 = v1130
	goto L173
L242:
	;
	v1138 = v1136
	v1140 = v921
	goto L173
L243:
	;
	goto L172
L244:
	;
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v614)+8))
	if v1155 != 0 {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v1155)+4))
	v1185 = base.B2i32(v1181 == int32(0))
	goto L247
L246:
	;
	v1185 = int32(1)
	goto L247
L247:
	;
	v1188 = F_innerrel_is_unique_ext(m, v582, v897, v1179, v705, int32(0), v1177, v1185, v604+int32(12))
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
		goto L21
	} else {
		goto L248
	}
L248:
	;
	if v1188 == int32(0) {
		goto L125
	} else {
		goto L249
	}
L249:
	;
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v604)+12))
	if v1192 == int32(0) {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v1388 = *(*int32)(unsafe.Add(mBase, uint32(v614)+212))
	v1389 = F_list_copy(m, v1388)
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L21
	} else {
		goto L299
	}
L251:
	;
	v1195 = int32(0)
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(v1192)+4))
	if v1196 <= v1195 {
		goto L250
	} else {
		goto L252
	}
L252:
	;
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v705)+68))
	v1226 = v1195
	goto L253
L253:
	;
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v1192)+12))
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v1228+v1226<<(uint(int32(2))%32))))
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v1232)+4))
	v1234 = F_copyObjectImpl(m, v1233)
	mBase = m.M
	v1235 = m.ExcPending
	if v1235 != 0 {
		goto L21
	} else {
		goto L255
	}
L254:
	;
	goto L250
L255:
	;
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(v614)+68))
	F_ChangeVarNodesExtended(m, v1234, v1199, v1236, int32(827))
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
		goto L21
	} else {
		goto L256
	}
L256:
	;
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v1234)+28))
	v1241 = *(*int32)(unsafe.Add(mBase, uint32(v1232)+44))
	if v1241 == int32(0) {
		goto L258
	} else {
		goto L259
	}
L257:
	;
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(v614)+184))
	if v1269 == int32(0) {
		goto L125
	} else {
		goto L271
	}
L258:
	;
	if v1240 == int32(0) {
		goto L261
	} else {
		goto L262
	}
L259:
	;
	goto L260
L260:
	;
	v1256 = int32(0)
	if v1240 == v1256 {
		goto L267
	} else {
		goto L268
	}
L261:
	;
	v1246 = int32(0)
	v1267 = v1246
	v1268 = v1246
	goto L257
L262:
	;
	goto L263
L263:
	;
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v1240)+12))
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(v1240)+4))
	if int32(2) <= v1250 {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v1248)+4))
	v1254 = v1253
	goto L266
L265:
	;
	v1254 = int32(0)
	goto L266
L266:
	;
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(v1248)))
	v1267 = v1254
	v1268 = v1255
	goto L257
L267:
	;
	v1267 = int32(0)
	v1268 = v1256
	goto L257
L268:
	;
	goto L269
L269:
	;
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v1240)+12))
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(v1260)))
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(v1240)+4))
	if v1262 < int32(2) {
		v1267 = v1261
		v1268 = v1256
		goto L257
	} else {
		goto L270
	}
L270:
	;
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v1260)+4))
	v1267 = v1261
	v1268 = v1265
	goto L257
L271:
	;
	v1272 = int32(0)
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(v1269)+4))
	if v1273 <= v1272 {
		goto L125
	} else {
		goto L272
	}
L272:
	;
	v1279 = v1272
	goto L273
L273:
	;
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(v1269)+12))
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v1304+v1279<<(uint(int32(2))%32))))
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v1308)+96))
	if v1309 == int32(0) {
		goto L276
	} else {
		goto L277
	}
L274:
	;
	v1357 = v1226 + int32(1)
	v1358 = *(*int32)(unsafe.Add(mBase, uint32(v1192)+4))
	if v1357 < v1358 {
		v1226 = v1357
		goto L253
	} else {
		goto L297
	}
L275:
	;
	goto L274
L276:
	;
	v1353 = v1279 + int32(1)
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v1269)+4))
	if v1353 < v1354 {
		v1279 = v1353
		goto L273
	} else {
		goto L296
	}
L277:
	;
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(v1308)+4))
	v1313 = *(*int32)(unsafe.Add(mBase, uint32(v1312)+28))
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(v1308)+44))
	if v1314 == int32(0) {
		goto L279
	} else {
		goto L280
	}
L278:
	;
	v1342 = F_equal(m, v1267, v1339)
	mBase = m.M
	v1343 = m.ExcPending
	if v1343 != 0 {
		goto L21
	} else {
		goto L292
	}
L279:
	;
	if v1313 == int32(0) {
		goto L282
	} else {
		goto L283
	}
L280:
	;
	goto L281
L281:
	;
	v1329 = int32(0)
	if v1313 == v1329 {
		goto L288
	} else {
		goto L289
	}
L282:
	;
	v1319 = int32(0)
	v1339 = v1319
	v1340 = v1319
	goto L278
L283:
	;
	goto L284
L284:
	;
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(v1313)+12))
	v1323 = *(*int32)(unsafe.Add(mBase, uint32(v1313)+4))
	if int32(2) <= v1323 {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v1321)+4))
	v1327 = v1326
	goto L287
L286:
	;
	v1327 = int32(0)
	goto L287
L287:
	;
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(v1321)))
	v1339 = v1327
	v1340 = v1328
	goto L278
L288:
	;
	v1339 = int32(0)
	v1340 = v1329
	goto L278
L289:
	;
	goto L290
L290:
	;
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v1313)+12))
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(v1333)))
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v1313)+4))
	if v1335 < int32(2) {
		v1339 = v1334
		v1340 = v1329
		goto L278
	} else {
		goto L291
	}
L291:
	;
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v1333)+4))
	v1339 = v1334
	v1340 = v1338
	goto L278
L292:
	;
	if v1342 == int32(0) {
		goto L276
	} else {
		goto L293
	}
L293:
	;
	v1346 = F_equal(m, v1268, v1340)
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L21
	} else {
		goto L294
	}
L294:
	;
	if v1346 != 0 {
		goto L275
	} else {
		goto L295
	}
L295:
	;
	goto L276
L296:
	;
	goto L125
L297:
	;
	goto L254
L298:
	;
	v1530 = *(*int32)(unsafe.Add(mBase, uint32(v614)+184))
	v1531 = F_list_concat(m, v1530, v901)
	mBase = m.M
	v1532 = m.ExcPending
	if v1532 != 0 {
		goto L21
	} else {
		goto L331
	}
L299:
	;
	if v1389 == int32(0) {
		goto L300
	} else {
		goto L301
	}
L300:
	;
	v1393 = int32(0)
	v1506 = v1393
	v1509 = v1393
	goto L298
L301:
	;
	goto L302
L302:
	;
	v1395 = int32(0)
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(v1389)+4))
	if v1398 <= v1395 {
		v1506 = v1395
		v1509 = v1395
		goto L298
	} else {
		goto L303
	}
L303:
	;
	v1402 = v1395
	v1405 = v1395
	v1408 = v1395
	goto L304
L304:
	;
	v1429 = *(*int32)(unsafe.Add(mBase, uint32(v1389)+12))
	v1433 = *(*int32)(unsafe.Add(mBase, uint32(v1429+v1402<<(uint(int32(2))%32))))
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(v1433)+32))
	F_remove_join_clause_from_rels(m, v582, v1433, v1434)
	mBase = m.M
	v1436 = m.ExcPending
	if v1436 != 0 {
		goto L21
	} else {
		goto L306
	}
L305:
	;
	v1506 = v1496
	v1509 = v1497
	goto L298
L306:
	;
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(v614)+68))
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v705)+68))
	F_ChangeVarNodesExtended(m, v1433, v1437, v1438, int32(827))
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L21
	} else {
		goto L307
	}
L307:
	;
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v1433)+32))
	if v1442 == int32(0) {
		goto L310
	} else {
		goto L311
	}
L308:
	;
	v1499 = v1402 + int32(1)
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(v1389)+4))
	if v1499 < v1500 {
		v1402 = v1499
		v1405 = v1496
		v1408 = v1497
		goto L304
	} else {
		goto L330
	}
L309:
	;
	if v1489 == int32(2) {
		goto L325
	} else {
		goto L326
	}
L310:
	;
	v1489 = int32(0)
	goto L309
L311:
	;
	goto L312
L312:
	;
	v1451 = int32(1)
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(v1442)+4))
	if v1452 <= v1451 {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	v1455 = v1451
	goto L315
L314:
	;
	v1455 = v1452
	goto L315
L315:
	;
	v1458 = int32(0)
	v1460 = v1458
	v1461 = v1458
	goto L316
L316:
	;
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(v1442+int32(8)+v1460<<(uint(int32(2))%32))))
	if v1469 != 0 {
		goto L319
	} else {
		goto L320
	}
L317:
	;
	v1489 = v1482
	goto L309
L318:
	;
	goto L317
L319:
	;
	v1470 = int32(2)
	if v1461 != 0 {
		v1482 = v1470
		goto L318
	} else {
		goto L322
	}
L320:
	;
	v1475 = v1461
	goto L321
L321:
	;
	v1478 = v1460 + int32(1)
	if v1478 != v1455 {
		v1460 = v1478
		v1461 = v1475
		goto L316
	} else {
		goto L324
	}
L322:
	;
	v1471 = int32(1)
	if base.Ui32(v1471) < base.Ui32(base.I32_popcnt(v1469)) {
		v1482 = v1470
		goto L318
	} else {
		goto L323
	}
L323:
	;
	v1475 = v1471
	goto L321
L324:
	;
	v1482 = v1475
	goto L318
L325:
	;
	v1492 = F_lappend(m, v1408, v1433)
	mBase = m.M
	v1493 = m.ExcPending
	if v1493 != 0 {
		goto L21
	} else {
		goto L328
	}
L326:
	;
	goto L327
L327:
	;
	v1494 = F_lappend(m, v1405, v1433)
	mBase = m.M
	v1495 = m.ExcPending
	if v1495 != 0 {
		goto L21
	} else {
		goto L329
	}
L328:
	;
	v1496 = v1405
	v1497 = v1492
	goto L308
L329:
	;
	v1496 = v1494
	v1497 = v1408
	goto L308
L330:
	;
	goto L305
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v614)+184)) = v1531
	if v1531 == int32(0) {
		v1642 = v1506
		v1645 = v1509
		goto L332
	} else {
		goto L333
	}
L332:
	;
	F_add_non_redundant_clauses(m, v582, v1642, v705+int32(184))
	mBase = m.M
	v1669 = m.ExcPending
	if v1669 != 0 {
		goto L21
	} else {
		goto L361
	}
L333:
	;
	v1536 = int32(0)
	v1537 = *(*int32)(unsafe.Add(mBase, uint32(v1531)+4))
	if v1537 <= v1536 {
		v1642 = v1506
		v1645 = v1509
		goto L332
	} else {
		goto L334
	}
L334:
	;
	v1543 = v1536
	v1544 = v1506
	v1547 = v1509
	goto L335
L335:
	;
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(v1531)+12))
	v1572 = *(*int32)(unsafe.Add(mBase, uint32(v1568+v1543<<(uint(int32(2))%32))))
	v1573 = *(*int32)(unsafe.Add(mBase, uint32(v614)+68))
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(v705)+68))
	F_ChangeVarNodesExtended(m, v1572, v1573, v1574, int32(827))
	mBase = m.M
	v1577 = m.ExcPending
	if v1577 != 0 {
		goto L21
	} else {
		goto L337
	}
L336:
	;
	v1642 = v1632
	v1645 = v1633
	goto L332
L337:
	;
	v1578 = *(*int32)(unsafe.Add(mBase, uint32(v1572)+32))
	if v1578 == int32(0) {
		goto L340
	} else {
		goto L341
	}
L338:
	;
	v1635 = v1543 + int32(1)
	v1636 = *(*int32)(unsafe.Add(mBase, uint32(v1531)+4))
	if v1635 < v1636 {
		v1543 = v1635
		v1544 = v1632
		v1547 = v1633
		goto L335
	} else {
		goto L360
	}
L339:
	;
	if v1625 == int32(2) {
		goto L355
	} else {
		goto L356
	}
L340:
	;
	v1625 = int32(0)
	goto L339
L341:
	;
	goto L342
L342:
	;
	v1587 = int32(1)
	v1588 = *(*int32)(unsafe.Add(mBase, uint32(v1578)+4))
	if v1588 <= v1587 {
		goto L343
	} else {
		goto L344
	}
L343:
	;
	v1591 = v1587
	goto L345
L344:
	;
	v1591 = v1588
	goto L345
L345:
	;
	v1594 = int32(0)
	v1596 = v1594
	v1597 = v1594
	goto L346
L346:
	;
	v1605 = *(*int32)(unsafe.Add(mBase, uint32(v1578+int32(8)+v1596<<(uint(int32(2))%32))))
	if v1605 != 0 {
		goto L349
	} else {
		goto L350
	}
L347:
	;
	v1625 = v1618
	goto L339
L348:
	;
	goto L347
L349:
	;
	v1606 = int32(2)
	if v1597 != 0 {
		v1618 = v1606
		goto L348
	} else {
		goto L352
	}
L350:
	;
	v1611 = v1597
	goto L351
L351:
	;
	v1614 = v1596 + int32(1)
	if v1614 != v1591 {
		v1596 = v1614
		v1597 = v1611
		goto L346
	} else {
		goto L354
	}
L352:
	;
	v1607 = int32(1)
	if base.Ui32(v1607) < base.Ui32(base.I32_popcnt(v1605)) {
		v1618 = v1606
		goto L348
	} else {
		goto L353
	}
L353:
	;
	v1611 = v1607
	goto L351
L354:
	;
	v1618 = v1611
	goto L348
L355:
	;
	v1628 = F_lappend(m, v1547, v1572)
	mBase = m.M
	v1629 = m.ExcPending
	if v1629 != 0 {
		goto L21
	} else {
		goto L358
	}
L356:
	;
	goto L357
L357:
	;
	v1630 = F_lappend(m, v1544, v1572)
	mBase = m.M
	v1631 = m.ExcPending
	if v1631 != 0 {
		goto L21
	} else {
		goto L359
	}
L358:
	;
	v1632 = v1544
	v1633 = v1628
	goto L338
L359:
	;
	v1632 = v1630
	v1633 = v1547
	goto L338
L360:
	;
	goto L336
L361:
	;
	F_add_non_redundant_clauses(m, v582, v1645, v705+int32(212))
	mBase = m.M
	v1673 = m.ExcPending
	if v1673 != 0 {
		goto L21
	} else {
		goto L362
	}
L362:
	;
	F_list_free(m, v1642)
	mBase = m.M
	v1675 = m.ExcPending
	if v1675 != 0 {
		goto L21
	} else {
		goto L363
	}
L363:
	;
	F_list_free(m, v1645)
	mBase = m.M
	v1677 = m.ExcPending
	if v1677 != 0 {
		goto L21
	} else {
		goto L364
	}
L364:
	;
	v1678 = *(*int32)(unsafe.Add(mBase, uint32(v614)+136))
	if v1678 == int32(0) {
		goto L367
	} else {
		goto L368
	}
L365:
	;
	if int32(0) <= v1735 {
		goto L376
	} else {
		goto L377
	}
L366:
	;
	v1735 = base.I32_ctz(v1721) | v1722<<(uint(int32(5))%32)
	goto L365
L367:
	;
	v1735 = int32(-2)
	goto L365
L368:
	;
	v1688 = base.I32_div_s(int32(0), int32(32))
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(v1678)+4))
	if v1689 <= v1688 {
		goto L367
	} else {
		goto L369
	}
L369:
	;
	v1692 = v1678 + int32(8)
	v1696 = *(*int32)(unsafe.Add(mBase, uint32(v1692+v1688<<(uint(int32(2))%32))))
	v1699 = v1696 & int32(-1)
	if v1699 != 0 {
		v1721 = v1699
		v1722 = v1688
		goto L366
	} else {
		goto L370
	}
L370:
	;
	v1701 = v1688 + int32(1)
	if v1701 == v1689 {
		goto L367
	} else {
		goto L371
	}
L371:
	;
	v1704 = v1701
	goto L372
L372:
	;
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(v1692+v1704<<(uint(int32(2))%32))))
	if v1711 != 0 {
		v1721 = v1711
		v1722 = v1704
		goto L366
	} else {
		goto L374
	}
L373:
	;
	goto L367
L374:
	;
	v1713 = v1704 + int32(1)
	if v1713 != v1689 {
		v1704 = v1713
		goto L372
	} else {
		goto L375
	}
L375:
	;
	goto L373
L376:
	;
	v1750 = v1735
	goto L379
L377:
	;
	goto L378
L378:
	;
	v2273 = *(*int32)(unsafe.Add(mBase, uint32(v614)+28))
	v2274 = *(*int32)(unsafe.Add(mBase, uint32(v2273)+4))
	if v2274 == int32(0) {
		goto L450
	} else {
		goto L451
	}
L379:
	;
	v1766 = *(*int32)(unsafe.Add(mBase, uint32(v705)+68))
	v1767 = *(*int32)(unsafe.Add(mBase, uint32(v614)+68))
	v1768 = int32(0)
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(v582)+88))
	v1772 = *(*int32)(unsafe.Add(mBase, uint32(v1771)+12))
	v1776 = *(*int32)(unsafe.Add(mBase, uint32(v1772+v1750<<(uint(int32(2))%32))))
	v1777 = *(*int32)(unsafe.Add(mBase, uint32(v1776)+16))
	if v1777 == v1768 {
		v1951 = v1768
		v1978 = v1768
		goto L381
	} else {
		goto L382
	}
L380:
	;
	goto L378
L381:
	;
	F_list_free(m, v1978)
	mBase = m.M
	v1980 = m.ExcPending
	if v1980 != 0 {
		goto L21
	} else {
		goto L408
	}
L382:
	;
	v1780 = int32(0)
	v1781 = *(*int32)(unsafe.Add(mBase, uint32(v1777)+4))
	if v1781 <= v1780 {
		goto L383
	} else {
		goto L384
	}
L383:
	;
	v1951 = v1768
	v1978 = v1777
	goto L381
L384:
	;
	goto L385
L385:
	;
	v1785 = v1768
	v1793 = v1780
	goto L386
L386:
	;
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(v1777)+12))
	v1816 = *(*int32)(unsafe.Add(mBase, uint32(v1812+v1793<<(uint(int32(2))%32))))
	v1817 = *(*int32)(unsafe.Add(mBase, uint32(v1816)+8))
	v1818 = F_bms_is_member(m, v1767, v1817)
	mBase = m.M
	v1819 = m.ExcPending
	if v1819 != 0 {
		goto L21
	} else {
		goto L390
	}
L387:
	;
	v1949 = *(*int32)(unsafe.Add(mBase, uint32(v1776)+16))
	v1951 = v1918
	v1978 = v1949
	goto L381
L388:
	;
	v1946 = v1793 + int32(1)
	v1947 = *(*int32)(unsafe.Add(mBase, uint32(v1777)+4))
	if v1946 < v1947 {
		v1785 = v1918
		v1793 = v1946
		goto L386
	} else {
		goto L407
	}
L389:
	;
	v1915 = F_lappend(m, v1785, v1816)
	mBase = m.M
	v1916 = m.ExcPending
	if v1916 != 0 {
		goto L21
	} else {
		goto L406
	}
L390:
	;
	if v1818 == int32(0) {
		goto L389
	} else {
		goto L391
	}
L391:
	;
	v1822 = *(*int32)(unsafe.Add(mBase, uint32(v1816)+8))
	v1823 = F_adjust_relid_set(m, v1822, v1767, v1766)
	mBase = m.M
	v1824 = m.ExcPending
	if v1824 != 0 {
		goto L21
	} else {
		goto L392
	}
L392:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1816)+8)) = v1823
	v1826 = *(*int32)(unsafe.Add(mBase, uint32(v1816)+20))
	v1827 = *(*int32)(unsafe.Add(mBase, uint32(v1826)+4))
	v1828 = F_adjust_relid_set(m, v1827, v1767, v1766)
	mBase = m.M
	v1829 = m.ExcPending
	if v1829 != 0 {
		goto L21
	} else {
		goto L393
	}
L393:
	;
	v1830 = *(*int32)(unsafe.Add(mBase, uint32(v1816)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1830)+4)) = v1828
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(v1816)+4))
	F_ChangeVarNodesExtended(m, v1832, v1767, v1766, int32(827))
	mBase = m.M
	v1835 = m.ExcPending
	if v1835 != 0 {
		goto L21
	} else {
		goto L394
	}
L394:
	;
	if v1785 == int32(0) {
		goto L389
	} else {
		goto L395
	}
L395:
	;
	v1838 = int32(0)
	v1839 = *(*int32)(unsafe.Add(mBase, uint32(v1785)+4))
	if v1839 <= v1838 {
		goto L389
	} else {
		goto L396
	}
L396:
	;
	v1845 = v1838
	goto L397
L397:
	;
	v1870 = *(*int32)(unsafe.Add(mBase, uint32(v1816)+8))
	v1871 = *(*int32)(unsafe.Add(mBase, uint32(v1785)+12))
	v1875 = *(*int32)(unsafe.Add(mBase, uint32(v1871+v1845<<(uint(int32(2))%32))))
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(v1875)+8))
	v1877 = F_equal(m, v1870, v1876)
	mBase = m.M
	v1878 = m.ExcPending
	if v1878 != 0 {
		goto L21
	} else {
		goto L399
	}
L398:
	;
	goto L389
L399:
	;
	if v1877 != 0 {
		goto L400
	} else {
		goto L401
	}
L400:
	;
	v1879 = *(*int32)(unsafe.Add(mBase, uint32(v1816)+4))
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(v1875)+4))
	v1881 = F_equal(m, v1879, v1880)
	mBase = m.M
	v1882 = m.ExcPending
	if v1882 != 0 {
		goto L21
	} else {
		goto L403
	}
L401:
	;
	goto L402
L402:
	;
	v1884 = v1845 + int32(1)
	v1885 = *(*int32)(unsafe.Add(mBase, uint32(v1785)+4))
	if v1884 < v1885 {
		v1845 = v1884
		goto L397
	} else {
		goto L405
	}
L403:
	;
	if v1881 != 0 {
		v1918 = v1785
		goto L388
	} else {
		goto L404
	}
L404:
	;
	goto L402
L405:
	;
	goto L398
L406:
	;
	v1918 = v1915
	goto L388
L407:
	;
	goto L387
L408:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1776)+16)) = v1951
	F_ec_clear_derived_clauses(m, v1776)
	mBase = m.M
	v1983 = m.ExcPending
	if v1983 != 0 {
		goto L21
	} else {
		goto L409
	}
L409:
	;
	v1984 = int32(0)
	v1985 = *(*int32)(unsafe.Add(mBase, uint32(v1776)+24))
	if v1985 == v1984 {
		v2148 = v1984
		v2167 = v1768
		goto L410
	} else {
		goto L411
	}
L410:
	;
	F_list_free(m, v2167)
	mBase = m.M
	v2176 = m.ExcPending
	if v2176 != 0 {
		goto L21
	} else {
		goto L435
	}
L411:
	;
	v1988 = int32(0)
	v1989 = *(*int32)(unsafe.Add(mBase, uint32(v1985)+4))
	if v1989 <= v1988 {
		goto L412
	} else {
		goto L413
	}
L412:
	;
	v2148 = v1984
	v2167 = v1985
	goto L410
L413:
	;
	goto L414
L414:
	;
	v1993 = v1984
	v2001 = v1988
	goto L415
L415:
	;
	v2020 = *(*int32)(unsafe.Add(mBase, uint32(v1985)+12))
	v2024 = *(*int32)(unsafe.Add(mBase, uint32(v2020+v2001<<(uint(int32(2))%32))))
	v2025 = *(*int32)(unsafe.Add(mBase, uint32(v2024)+32))
	v2026 = F_bms_is_member(m, v1767, v2025)
	mBase = m.M
	v2027 = m.ExcPending
	if v2027 != 0 {
		goto L21
	} else {
		goto L419
	}
L416:
	;
	v2146 = *(*int32)(unsafe.Add(mBase, uint32(v1776)+24))
	v2148 = v2115
	v2167 = v2146
	goto L410
L417:
	;
	v2143 = v2001 + int32(1)
	v2144 = *(*int32)(unsafe.Add(mBase, uint32(v1985)+4))
	if v2143 < v2144 {
		v1993 = v2115
		v2001 = v2143
		goto L415
	} else {
		goto L434
	}
L418:
	;
	v2112 = F_lappend(m, v1993, v2024)
	mBase = m.M
	v2113 = m.ExcPending
	if v2113 != 0 {
		goto L21
	} else {
		goto L433
	}
L419:
	;
	if v2026 == int32(0) {
		goto L418
	} else {
		goto L420
	}
L420:
	;
	F_ChangeVarNodesExtended(m, v2024, v1767, v1766, int32(827))
	mBase = m.M
	v2032 = m.ExcPending
	if v2032 != 0 {
		goto L21
	} else {
		goto L421
	}
L421:
	;
	if v1993 == int32(0) {
		goto L418
	} else {
		goto L422
	}
L422:
	;
	v2035 = int32(0)
	v2036 = *(*int32)(unsafe.Add(mBase, uint32(v1993)+4))
	if v2036 <= v2035 {
		goto L418
	} else {
		goto L423
	}
L423:
	;
	v2042 = v2035
	goto L424
L424:
	;
	v2067 = *(*int32)(unsafe.Add(mBase, uint32(v2024)+28))
	v2068 = *(*int32)(unsafe.Add(mBase, uint32(v1993)+12))
	v2072 = *(*int32)(unsafe.Add(mBase, uint32(v2068+v2042<<(uint(int32(2))%32))))
	v2073 = *(*int32)(unsafe.Add(mBase, uint32(v2072)+28))
	v2074 = F_equal(m, v2067, v2073)
	mBase = m.M
	v2075 = m.ExcPending
	if v2075 != 0 {
		goto L21
	} else {
		goto L426
	}
L425:
	;
	goto L418
L426:
	;
	if v2074 != 0 {
		goto L427
	} else {
		goto L428
	}
L427:
	;
	v2076 = *(*int32)(unsafe.Add(mBase, uint32(v2024)+4))
	v2077 = *(*int32)(unsafe.Add(mBase, uint32(v2072)+4))
	v2078 = F_equal(m, v2076, v2077)
	mBase = m.M
	v2079 = m.ExcPending
	if v2079 != 0 {
		goto L21
	} else {
		goto L430
	}
L428:
	;
	goto L429
L429:
	;
	v2081 = v2042 + int32(1)
	v2082 = *(*int32)(unsafe.Add(mBase, uint32(v1993)+4))
	if v2081 < v2082 {
		v2042 = v2081
		goto L424
	} else {
		goto L432
	}
L430:
	;
	if v2078 != 0 {
		v2115 = v1993
		goto L417
	} else {
		goto L431
	}
L431:
	;
	goto L429
L432:
	;
	goto L425
L433:
	;
	v2115 = v2112
	goto L417
L434:
	;
	goto L416
L435:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1776)+24)) = v2148
	v2178 = *(*int32)(unsafe.Add(mBase, uint32(v1776)+36))
	v2179 = F_adjust_relid_set(m, v2178, v1767, v1766)
	mBase = m.M
	v2180 = m.ExcPending
	if v2180 != 0 {
		goto L21
	} else {
		goto L436
	}
L436:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1776)+36)) = v2179
	v2182 = *(*int32)(unsafe.Add(mBase, uint32(v705)+136))
	v2183 = F_bms_add_member(m, v2182, v1750)
	mBase = m.M
	v2184 = m.ExcPending
	if v2184 != 0 {
		goto L21
	} else {
		goto L437
	}
L437:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v705)+136)) = v2183
	v2186 = *(*int32)(unsafe.Add(mBase, uint32(v614)+136))
	if v2186 == int32(0) {
		goto L440
	} else {
		goto L441
	}
L438:
	;
	if int32(0) <= v2242 {
		v1750 = v2242
		goto L379
	} else {
		goto L449
	}
L439:
	;
	v2242 = base.I32_ctz(v2228) | v2229<<(uint(int32(5))%32)
	goto L438
L440:
	;
	v2242 = int32(-2)
	goto L438
L441:
	;
	v2193 = v1750 + int32(1)
	v2195 = base.I32_div_s(v2193, int32(32))
	v2196 = *(*int32)(unsafe.Add(mBase, uint32(v2186)+4))
	if v2196 <= v2195 {
		goto L440
	} else {
		goto L442
	}
L442:
	;
	v2199 = v2186 + int32(8)
	v2203 = *(*int32)(unsafe.Add(mBase, uint32(v2199+v2195<<(uint(int32(2))%32))))
	v2206 = v2203 & (int32(-1) << (uint(v2193) % 32))
	if v2206 != 0 {
		v2228 = v2206
		v2229 = v2195
		goto L439
	} else {
		goto L443
	}
L443:
	;
	v2208 = v2195 + int32(1)
	if v2208 == v2196 {
		goto L440
	} else {
		goto L444
	}
L444:
	;
	v2211 = v2208
	goto L445
L445:
	;
	v2218 = *(*int32)(unsafe.Add(mBase, uint32(v2199+v2211<<(uint(int32(2))%32))))
	if v2218 != 0 {
		v2228 = v2218
		v2229 = v2211
		goto L439
	} else {
		goto L447
	}
L446:
	;
	goto L440
L447:
	;
	v2220 = v2211 + int32(1)
	if v2220 != v2196 {
		v2211 = v2220
		goto L445
	} else {
		goto L448
	}
L448:
	;
	goto L446
L449:
	;
	goto L380
L450:
	;
	v2364 = int32(*(*int16)(unsafe.Add(mBase, uint32(v705)+80)))
	v2365 = int32(*(*int16)(unsafe.Add(mBase, uint32(v705)+82)))
	if v2364 <= v2365 {
		goto L462
	} else {
		goto L463
	}
L451:
	;
	v2277 = int32(0)
	v2278 = *(*int32)(unsafe.Add(mBase, uint32(v2274)+4))
	if v2278 <= v2277 {
		goto L450
	} else {
		goto L452
	}
L452:
	;
	v2284 = v2277
	goto L453
L453:
	;
	v2309 = *(*int32)(unsafe.Add(mBase, uint32(v2274)+12))
	v2313 = *(*int32)(unsafe.Add(mBase, uint32(v2309+v2284<<(uint(int32(2))%32))))
	v2314 = *(*int32)(unsafe.Add(mBase, uint32(v614)+68))
	v2315 = *(*int32)(unsafe.Add(mBase, uint32(v705)+68))
	F_ChangeVarNodesExtended(m, v2313, v2314, v2315, int32(827))
	mBase = m.M
	v2318 = m.ExcPending
	if v2318 != 0 {
		goto L21
	} else {
		goto L455
	}
L454:
	;
	goto L450
L455:
	;
	v2319 = *(*int32)(unsafe.Add(mBase, uint32(v705)+28))
	v2320 = *(*int32)(unsafe.Add(mBase, uint32(v2319)+4))
	v2321 = F_list_member(m, v2320, v2313)
	mBase = m.M
	v2322 = m.ExcPending
	if v2322 != 0 {
		goto L21
	} else {
		goto L456
	}
L456:
	;
	if v2321 == int32(0) {
		goto L457
	} else {
		goto L458
	}
L457:
	;
	v2325 = *(*int32)(unsafe.Add(mBase, uint32(v705)+28))
	v2326 = *(*int32)(unsafe.Add(mBase, uint32(v2325)+4))
	v2327 = F_lappend(m, v2326, v2313)
	mBase = m.M
	v2328 = m.ExcPending
	if v2328 != 0 {
		goto L21
	} else {
		goto L460
	}
L458:
	;
	goto L459
L459:
	;
	v2333 = v2284 + int32(1)
	v2334 = *(*int32)(unsafe.Add(mBase, uint32(v2274)+4))
	if v2333 < v2334 {
		v2284 = v2333
		goto L453
	} else {
		goto L461
	}
L460:
	;
	v2329 = *(*int32)(unsafe.Add(mBase, uint32(v705)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v2329)+4)) = v2327
	goto L459
L461:
	;
	goto L454
L462:
	;
	v2368 = v2364
	goto L465
L463:
	;
	goto L464
L464:
	;
	if v884 == int32(0) {
		goto L470
	} else {
		goto L471
	}
L465:
	;
	v2395 = int32(*(*int16)(unsafe.Add(mBase, uint32(v705)+80)))
	v2398 = (v2368 - v2395) << (uint(int32(2)) % 32)
	v2399 = *(*int32)(unsafe.Add(mBase, uint32(v614)+84))
	v2401 = *(*int32)(unsafe.Add(mBase, uint32(v2398+v2399)))
	v2402 = *(*int32)(unsafe.Add(mBase, uint32(v614)+68))
	v2403 = *(*int32)(unsafe.Add(mBase, uint32(v705)+68))
	v2404 = F_adjust_relid_set(m, v2401, v2402, v2403)
	mBase = m.M
	v2405 = m.ExcPending
	if v2405 != 0 {
		goto L21
	} else {
		goto L467
	}
L466:
	;
	goto L464
L467:
	;
	v2406 = *(*int32)(unsafe.Add(mBase, uint32(v614)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v2406+v2398))) = v2404
	v2409 = *(*int32)(unsafe.Add(mBase, uint32(v705)+84))
	v2411 = *(*int32)(unsafe.Add(mBase, uint32(v2409+v2398)))
	v2412 = *(*int32)(unsafe.Add(mBase, uint32(v614)+84))
	v2414 = *(*int32)(unsafe.Add(mBase, uint32(v2412+v2398)))
	v2415 = F_bms_add_members(m, v2411, v2414)
	mBase = m.M
	v2416 = m.ExcPending
	if v2416 != 0 {
		goto L21
	} else {
		goto L468
	}
L468:
	;
	v2417 = *(*int32)(unsafe.Add(mBase, uint32(v705)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v2417+v2398))) = v2415
	v2420 = int32(*(*int16)(unsafe.Add(mBase, uint32(v705)+82)))
	if v2368 < v2420 {
		v2368 = v2368 + int32(1)
		goto L465
	} else {
		goto L469
	}
L469:
	;
	goto L466
L470:
	;
	v2462 = *(*int32)(unsafe.Add(mBase, uint32(v582)+4))
	v2463 = *(*int32)(unsafe.Add(mBase, uint32(v614)+68))
	v2464 = *(*int32)(unsafe.Add(mBase, uint32(v705)+68))
	F_ChangeVarNodesExtended(m, v2462, v2463, v2464, int32(827))
	mBase = m.M
	v2467 = m.ExcPending
	if v2467 != 0 {
		goto L21
	} else {
		goto L476
	}
L471:
	;
	if v877 != 0 {
		goto L472
	} else {
		goto L473
	}
L472:
	;
	v2454 = *(*int32)(unsafe.Add(mBase, uint32(v582)+136))
	v2455 = F_list_delete_ptr(m, v2454, v884)
	mBase = m.M
	v2456 = m.ExcPending
	if v2456 != 0 {
		goto L21
	} else {
		goto L475
	}
L473:
	;
	goto L474
L474:
	;
	v2458 = *(*int32)(unsafe.Add(mBase, uint32(v705)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v884)+4)) = v2458
	*(*int32)(unsafe.Add(mBase, uint32(v884)+8)) = v2458
	goto L470
L475:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v582)+136)) = v2455
	goto L470
L476:
	;
	v2468 = *(*int32)(unsafe.Add(mBase, uint32(v705)+68))
	v2469 = int32(0)
	F_remove_rel_from_query(m, v582, v614, v2468, v2469, v2469)
	mBase = m.M
	v2472 = m.ExcPending
	if v2472 != 0 {
		goto L21
	} else {
		goto L477
	}
L477:
	;
	v2473 = *(*int32)(unsafe.Add(mBase, uint32(v582)+264))
	v2474 = *(*int32)(unsafe.Add(mBase, uint32(v614)+68))
	v2475 = *(*int32)(unsafe.Add(mBase, uint32(v705)+68))
	F_ChangeVarNodesExtended(m, v2473, v2474, v2475, int32(827))
	mBase = m.M
	v2478 = m.ExcPending
	if v2478 != 0 {
		goto L21
	} else {
		goto L478
	}
L478:
	;
	v2479 = *(*int32)(unsafe.Add(mBase, uint32(v582)+256))
	v2480 = *(*int32)(unsafe.Add(mBase, uint32(v614)+68))
	v2481 = *(*int32)(unsafe.Add(mBase, uint32(v705)+68))
	F_ChangeVarNodesExtended(m, v2479, v2480, v2481, int32(827))
	mBase = m.M
	v2484 = m.ExcPending
	if v2484 != 0 {
		goto L21
	} else {
		goto L479
	}
L479:
	;
	v2485 = *(*int32)(unsafe.Add(mBase, uint32(v582)+120))
	v2486 = *(*int32)(unsafe.Add(mBase, uint32(v614)+68))
	v2487 = *(*int32)(unsafe.Add(mBase, uint32(v705)+68))
	v2488 = F_adjust_relid_set(m, v2485, v2486, v2487)
	mBase = m.M
	v2489 = m.ExcPending
	if v2489 != 0 {
		goto L21
	} else {
		goto L480
	}
L480:
	;
	v2490 = *(*int32)(unsafe.Add(mBase, uint32(v582)+124))
	v2491 = *(*int32)(unsafe.Add(mBase, uint32(v614)+68))
	v2492 = *(*int32)(unsafe.Add(mBase, uint32(v705)+68))
	v2493 = F_adjust_relid_set(m, v2490, v2491, v2492)
	mBase = m.M
	v2494 = m.ExcPending
	if v2494 != 0 {
		goto L21
	} else {
		goto L481
	}
L481:
	;
	v2495 = *(*int32)(unsafe.Add(mBase, uint32(v582)+28))
	v2496 = *(*int32)(unsafe.Add(mBase, uint32(v614)+68))
	v2497 = int32(2)
	v2500 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2495+v2496<<(uint(v2497)%32)))) = v2500
	v2502 = *(*int32)(unsafe.Add(mBase, uint32(v582)+36))
	v2503 = *(*int32)(unsafe.Add(mBase, uint32(v614)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v2502+v2503<<(uint(v2497)%32)))) = v2500
	F_pfree(m, v614)
	mBase = m.M
	v2510 = m.ExcPending
	if v2510 != 0 {
		goto L21
	} else {
		goto L482
	}
L482:
	;
	F_rebuild_placeholder_attr_needed(m, v582)
	mBase = m.M
	v2512 = m.ExcPending
	if v2512 != 0 {
		goto L21
	} else {
		goto L483
	}
L483:
	;
	F_rebuild_joinclause_attr_needed(m, v582)
	mBase = m.M
	v2514 = m.ExcPending
	if v2514 != 0 {
		goto L21
	} else {
		goto L484
	}
L484:
	;
	F_rebuild_eclass_attr_needed(m, v582)
	mBase = m.M
	v2516 = m.ExcPending
	if v2516 != 0 {
		goto L21
	} else {
		goto L485
	}
L485:
	;
	F_rebuild_lateral_attr_needed(m, v582)
	mBase = m.M
	v2518 = m.ExcPending
	if v2518 != 0 {
		goto L21
	} else {
		goto L486
	}
L486:
	;
	v2519 = F_bms_add_member(m, v606, v601)
	mBase = m.M
	v2520 = m.ExcPending
	if v2520 != 0 {
		goto L21
	} else {
		goto L487
	}
L487:
	;
	v2631 = v2519
	goto L110
L488:
	;
	if int32(0) < v2604 {
		v677 = v2604
		goto L123
	} else {
		goto L499
	}
L489:
	;
	v2604 = base.I32_ctz(v2590) | v2591<<(uint(int32(5))%32)
	goto L488
L490:
	;
	v2604 = int32(-2)
	goto L488
L491:
	;
	v2555 = v677 + int32(1)
	v2557 = base.I32_div_s(v2555, int32(32))
	v2558 = *(*int32)(unsafe.Add(mBase, uint32(v597)+4))
	if v2558 <= v2557 {
		goto L490
	} else {
		goto L492
	}
L492:
	;
	v2561 = v597 + int32(8)
	v2565 = *(*int32)(unsafe.Add(mBase, uint32(v2561+v2557<<(uint(int32(2))%32))))
	v2568 = v2565 & (int32(-1) << (uint(v2555) % 32))
	if v2568 != 0 {
		v2590 = v2568
		v2591 = v2557
		goto L489
	} else {
		goto L493
	}
L493:
	;
	v2570 = v2557 + int32(1)
	if v2570 == v2558 {
		goto L490
	} else {
		goto L494
	}
L494:
	;
	v2573 = v2570
	goto L495
L495:
	;
	v2580 = *(*int32)(unsafe.Add(mBase, uint32(v2561+v2573<<(uint(int32(2))%32))))
	if v2580 != 0 {
		v2590 = v2580
		v2591 = v2573
		goto L489
	} else {
		goto L497
	}
L496:
	;
	goto L490
L497:
	;
	v2582 = v2573 + int32(1)
	if v2582 != v2558 {
		v2573 = v2582
		goto L495
	} else {
		goto L498
	}
L498:
	;
	goto L496
L499:
	;
	goto L124
L500:
	;
	if int32(0) < v2690 {
		v601 = v2690
		v606 = v2631
		goto L108
	} else {
		goto L511
	}
L501:
	;
	v2690 = base.I32_ctz(v2676) | v2677<<(uint(int32(5))%32)
	goto L500
L502:
	;
	v2690 = int32(-2)
	goto L500
L503:
	;
	v2641 = v601 + int32(1)
	v2643 = base.I32_div_s(v2641, int32(32))
	v2644 = *(*int32)(unsafe.Add(mBase, uint32(v597)+4))
	if v2644 <= v2643 {
		goto L502
	} else {
		goto L504
	}
L504:
	;
	v2647 = v597 + int32(8)
	v2651 = *(*int32)(unsafe.Add(mBase, uint32(v2647+v2643<<(uint(int32(2))%32))))
	v2654 = v2651 & (int32(-1) << (uint(v2641) % 32))
	if v2654 != 0 {
		v2676 = v2654
		v2677 = v2643
		goto L501
	} else {
		goto L505
	}
L505:
	;
	v2656 = v2643 + int32(1)
	if v2656 == v2644 {
		goto L502
	} else {
		goto L506
	}
L506:
	;
	v2659 = v2656
	goto L507
L507:
	;
	v2666 = *(*int32)(unsafe.Add(mBase, uint32(v2647+v2659<<(uint(int32(2))%32))))
	if v2666 != 0 {
		v2676 = v2666
		v2677 = v2659
		goto L501
	} else {
		goto L509
	}
L508:
	;
	goto L502
L509:
	;
	v2668 = v2659 + int32(1)
	if v2668 != v2644 {
		v2659 = v2668
		goto L507
	} else {
		goto L510
	}
L510:
	;
	goto L508
L511:
	;
	goto L109
L512:
	;
	v2723 = F_bms_del_members(m, v2708, v2717)
	mBase = m.M
	v2724 = m.ExcPending
	if v2724 != 0 {
		goto L21
	} else {
		goto L513
	}
L513:
	;
	if v2717 != 0 {
		goto L514
	} else {
		goto L515
	}
L514:
	;
	if v2723 == int32(0) {
		goto L518
	} else {
		goto L519
	}
L515:
	;
	goto L516
L516:
	;
	goto L93
L517:
	;
	if v2771 == int32(2) {
		v494 = v2693
		v496 = v2721
		v509 = v2723
		v510 = v2709
		v511 = v2710
		v515 = v2714
		v516 = v2715
		v517 = v2716
		v519 = v2718
		goto L92
	} else {
		goto L533
	}
L518:
	;
	v2771 = int32(0)
	goto L517
L519:
	;
	goto L520
L520:
	;
	v2733 = int32(1)
	v2734 = *(*int32)(unsafe.Add(mBase, uint32(v2723)+4))
	if v2734 <= v2733 {
		goto L521
	} else {
		goto L522
	}
L521:
	;
	v2737 = v2733
	goto L523
L522:
	;
	v2737 = v2734
	goto L523
L523:
	;
	v2740 = int32(0)
	v2742 = v2740
	v2743 = v2740
	goto L524
L524:
	;
	v2751 = *(*int32)(unsafe.Add(mBase, uint32(v2723+int32(8)+v2742<<(uint(int32(2))%32))))
	if v2751 != 0 {
		goto L527
	} else {
		goto L528
	}
L525:
	;
	v2771 = v2764
	goto L517
L526:
	;
	goto L525
L527:
	;
	v2752 = int32(2)
	if v2743 != 0 {
		v2764 = v2752
		goto L526
	} else {
		goto L530
	}
L528:
	;
	v2757 = v2743
	goto L529
L529:
	;
	v2760 = v2742 + int32(1)
	if v2760 != v2737 {
		v2742 = v2760
		v2743 = v2757
		goto L524
	} else {
		goto L532
	}
L530:
	;
	v2753 = int32(1)
	if base.Ui32(v2753) < base.Ui32(base.I32_popcnt(v2751)) {
		v2764 = v2752
		goto L526
	} else {
		goto L531
	}
L531:
	;
	v2757 = v2753
	goto L529
L532:
	;
	v2764 = v2757
	goto L526
L533:
	;
	goto L516
L534:
	;
	F_bms_free(m, v2723)
	mBase = m.M
	v2777 = m.ExcPending
	if v2777 != 0 {
		goto L21
	} else {
		goto L535
	}
L535:
	;
	v2784 = v2693
	v2786 = v2721
	v2800 = v2709
	v2801 = v2710
	v2805 = v2714
	v2806 = v2715
	v2807 = v2716
	v2809 = v2718
	goto L76
L536:
	;
	v2784 = v384
	v2786 = v386
	v2800 = v2782
	v2801 = v387
	v2805 = v387
	v2806 = v406
	v2807 = v407
	v2809 = v409
	goto L76
L537:
	;
	goto L75
}
func F_removeabbrev_cluster(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
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
	var v27 int32
	_ = v27
	v4 = int32(0)
	if v4 < l2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v12 = v4
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v16 = l1 + v12<<(uint(int32(4))%32)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+12)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v23 = F_heap_getattr_1(m, v17, v19, v20, v16+int32(8))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v23
	v27 = v12 + int32(1)
	if v27 != l2 {
		v12 = v27
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
}
func F_removecaptures(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v5&int32(32) == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(0)
	v13 = v5 & int32(215)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v13)
	v15 = v13
	goto L3
L2:
	;
	v15 = v5
	goto L3
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v20 = v16
	goto L7
L5:
	;
	v34 = v15
	goto L6
L6:
	;
	if v34&int32(24) == int32(0) {
		goto L15
	} else {
		goto L16
	}
L7:
	;
	F_removecaptures(m, l0, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	v34 = v31
	goto L6
L9:
	;
	return
L10:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v23&int32(8) != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	v28 = v26 | int32(8)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v28)
	goto L13
L12:
	;
	goto L13
L13:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	if v30 != 0 {
		v20 = v30
		goto L7
	} else {
		goto L14
	}
L14:
	;
	goto L8
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v40 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	return
L18:
	;
	v44 = v40
	goto L21
L19:
	;
	v51 = v34
	goto L20
L20:
	;
	v53 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v53)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = int32(0)
	v58 = v51 & int32(251)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v58)
	goto L17
L21:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+24))
	F_freesubre(m, l0, v44)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L9
	} else {
		goto L23
	}
L22:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	v51 = v48
	goto L20
L23:
	;
	if v45 != 0 {
		v44 = v45
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
}
func F_renameatt_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	v18 = F_relation_open(m, l0, int32(8))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	F_renameatt_check(m, l0, v22, l4)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l3 != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L65
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L61
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L57
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L53
	}
L8:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+119)))
	if v92 != int32(99) {
		goto L29
	} else {
		goto L30
	}
L9:
	;
	v28 = F_find_all_inheritors(m, l0, int32(8), v13+int32(-4))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	if l5 != 0 {
		goto L8
	} else {
		goto L26
	}
L12:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	v36 = int32(0)
	goto L13
L13:
	;
	v44 = int32(0)
	if v28 == v44 {
		v54 = v44
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if v30 == int32(0) {
		goto L8
	} else {
		goto L18
	}
L16:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v48 <= v36 {
		v54 = int32(0)
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v54 = v50 + v36<<(uint(int32(2))%32)
	goto L15
L18:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v57 <= v36 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	if v54 == int32(0) {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v64 = v61 + v36<<(uint(int32(2))%32)
	if v64 == int32(0) {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if l0 != v67 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v72 = F_renameatt_internal(m, v67, l1, l2, int32(0), int32(1), v71, l6)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v36 = v36 + int32(1)
	goto L13
L25:
	;
	goto L24
L26:
	;
	v77 = F_find_inheritance_children(m, l0, int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v77 != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	goto L8
L29:
	;
	v146 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L38
	}
L30:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91)+72))
	v98 = F_find_typed_table_dependencies(m, v95, v91+int32(4), l6)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v98 == int32(0) {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	if v102 <= int32(0) {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	v110 = int32(0)
	goto L34
L34:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v98)+12))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118+v110<<(uint(int32(2))%32))))
	v123 = int32(1)
	v126 = F_renameatt_internal(m, v122, l1, l2, v123, v123, int32(0), l6)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L36
	}
L35:
	;
	goto L29
L36:
	;
	v129 = v110 + int32(1)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	if v129 < v130 {
		v110 = v129
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v148 = F_SearchSysCacheCopyAttName(m, l0, l1)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if v148 == int32(0) {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v148)+16))
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+22)))
	v154 = v152 + v153
	v155 = int32(*(*int16)(unsafe.Add(mBase, uint32(v154)+74)))
	if v155 <= int32(0) {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	v158 = int32(*(*int16)(unsafe.Add(mBase, uint32(v154)+94)))
	if l5 < v158 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v161 = F_check_for_column_name_collision(m, v18, l2, int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v166 = F_strncpy(m, v154+int32(4), l2, int32(64))
	mBase = m.M
	v167 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v166)+63)) = uint8(v167)
	goto L44
L44:
	;
	F_CatalogTupleUpdate(m, v146, v148+int32(4), v148)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _consts[297]))
	if v174 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v176 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1259), l0, v155, v176, v176)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	F_pfree(m, v148)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L50
	}
L49:
	;
	goto L48
L50:
	;
	F_sequence_close(m, v146, int32(3))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_relation_close(m, v18, int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	m.G0 = v15 - int32(-64)
	return v155
L53:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l1
	F_errmsg(m, int32(230447), v13+int32(-16))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(475396), int32(3917), int32(297956))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
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
	F_errcode(m, int32(50360452))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l1
	F_errmsg(m, int32(69214), v15)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(475396), int32(3941), int32(297956))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l1
	F_errmsg(m, int32(673886), v13+int32(-48))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(475396), int32(3949), int32(297956))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l1
	F_errmsg(m, int32(674223), v13+int32(-32))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(475396), int32(3964), int32(297956))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_renametrig_partition(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v97 int32
	_ = v97
	var v109 int32
	_ = v109
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	F_ScanKeyInit(m, v13, int32(2), int32(3), int32(184), l1)
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
	v21 = int32(1)
	v24 = F_systable_beginscan(m, l0, int32(2701), v21, int32(0), v21, v13)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	goto L5
L4:
	;
	F_systable_endscan(m, v24)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L21
	}
L5:
	;
	v36 = F_systable_getnext(m, v24)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v46 = F_table_open(m, l1, int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L10
	}
L7:
	;
	if v36 == int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+22)))
	v42 = v40 + v41
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v43 != l2 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	goto L6
L10:
	;
	F_renametrig_internal(m, l0, v46, v36, l3, l4)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+48))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+119)))
	if v51 != int32(112) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_sequence_close(m, v46, int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L20
	}
L13:
	;
	v55 = F_RelationGetPartitionDesc(m, v46, int32(1))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v57 <= int32(0) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v68 = int32(0)
	goto L16
L16:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73+v68<<(uint(int32(2))%32))))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	F_renametrig_partition(m, l0, v77, v78, l3, v42+int32(12))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	goto L12
L18:
	;
	v82 = v68 + int32(1)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v82 < v83 {
		v68 = v82
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	goto L4
L21:
	;
	m.G0 = v13 + int32(48)
	return
}
func F_replace_correlation_vars_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
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
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
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
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	v3 = int32(0)
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v12 - int32(6) {
	case 0:
		goto L9
	case 1, 2, 5, 6:
		goto L4
	case 3:
		goto L8
	case 4:
		goto L7
	case 7:
		goto L6
	default:
		goto L10
	}
L4:
	;
	v780 = F_expression_tree_mutator_impl(m, l0, int32(845), l1)
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L39
	} else {
		goto L196
	}
L5:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v670 <= int32(0) {
		goto L4
	} else {
		goto L170
	}
L6:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v600)+4))
	if v601 == int32(5) {
		goto L4
	} else {
		goto L151
	}
L7:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v499 == int32(0) {
		goto L4
	} else {
		goto L127
	}
L8:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v396 == int32(0) {
		goto L4
	} else {
		goto L104
	}
L9:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v180 == int32(0) {
		goto L4
	} else {
		goto L53
	}
L10:
	;
	if v12 == int32(61) {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	if v12 != int32(319) {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v19 == int32(0) {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v22 == int32(0) {
		v70 = l1
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	if v75 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L15:
	;
	v26 = v22 & int32(7)
	if v26 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if base.Ui32(v22) < base.Ui32(int32(8)) {
		v70 = v44
		goto L14
	} else {
		goto L23
	}
L17:
	;
	v42 = v22
	v44 = l1
	goto L16
L18:
	;
	goto L19
L19:
	;
	v29 = v22
	v31 = l1
	v32 = v3
	goto L20
L20:
	;
	v36 = int32(1)
	v37 = v29 - v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v40 = v32 + v36
	if v40 != v26 {
		v29 = v37
		v31 = v38
		v32 = v40
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v42 = v37
	v44 = v38
	goto L16
L22:
	;
	goto L21
L23:
	;
	v51 = v42
	v53 = v44
	goto L24
L24:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
	v67 = v51 - int32(8)
	if v67 != 0 {
		v51 = v67
		v53 = v65
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v70 = v65
	goto L14
L26:
	;
	goto L25
L27:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	v160 = F_palloc0(m, int32(28))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L39
	} else {
		goto L49
	}
L28:
	;
	v157 = v93 + int32(8)
	goto L27
L29:
	;
	v111 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L39
	} else {
		goto L40
	}
L30:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v78 <= int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	v83 = int32(0)
	goto L32
L32:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v81+v83<<(uint(int32(2))%32))))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if v95 == int32(319) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L29
L34:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94)+16))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v98 == v99 {
		goto L28
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v102 = v83 + int32(1)
	if v78 != v102 {
		v83 = v102
		goto L32
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	goto L33
L39:
	;
	return int32(0)
L40:
	;
	v115 = int32(0)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	F_IncrementVarSublevelsUp(m, v111, v115-v116, v115)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v122 = F_palloc0(m, int32(12))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L39
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122)+4)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = int32(326)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+64))
	if v128 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	v131 = v129
	goto L45
L44:
	;
	v131 = int32(0)
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122)+8)) = v131
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+64))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v136 = F_exprType(m, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L39
	} else {
		goto L46
	}
L46:
	;
	v138 = F_lappend_oid(m, v134, v136)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L39
	} else {
		goto L47
	}
L47:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v140)+64)) = v138
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	v143 = F_lappend(m, v142, v122)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L39
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+20)) = v143
	v157 = v122 + int32(8)
	goto L27
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+8)) = v158
	*(*int64)(unsafe.Add(mBase, uint32(v160))) = int64(4294967304)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v166 = F_exprType(m, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L39
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+12)) = v166
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v170 = F_exprTypmod(m, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L39
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+16)) = v170
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v174 = F_exprCollation(m, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L39
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+24)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v160)+20)) = v174
	return v160
L53:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v183 == int32(0) {
		v231 = l1
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v231)+20))
	if v236 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L55:
	;
	v187 = v183 & int32(7)
	if v187 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	if base.Ui32(v183) < base.Ui32(int32(8)) {
		v231 = v205
		goto L54
	} else {
		goto L63
	}
L57:
	;
	v203 = v183
	v205 = l1
	goto L56
L58:
	;
	goto L59
L59:
	;
	v190 = v183
	v192 = l1
	v193 = v3
	goto L60
L60:
	;
	v197 = int32(1)
	v198 = v190 - v197
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v192)+16))
	v201 = v193 + v197
	if v201 != v187 {
		v190 = v198
		v192 = v199
		v193 = v201
		goto L60
	} else {
		goto L62
	}
L61:
	;
	v203 = v198
	v205 = v199
	goto L56
L62:
	;
	goto L61
L63:
	;
	v212 = v203
	v214 = v205
	goto L64
L64:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v214)+16))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+16))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+16))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+16))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+16))
	v228 = v212 - int32(8)
	if v228 != 0 {
		v212 = v228
		v214 = v226
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v231 = v226
	goto L54
L66:
	;
	goto L65
L67:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v379)))
	v382 = F_palloc0(m, int32(28))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L39
	} else {
		goto L103
	}
L68:
	;
	v379 = v254 + int32(8)
	goto L67
L69:
	;
	v341 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L39
	} else {
		goto L96
	}
L70:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v236)+4))
	if v239 <= int32(0) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v243 = int32(0)
	v244 = v239
	goto L72
L72:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v236)+12))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v250+v243<<(uint(int32(2))%32))))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)+4))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	if v256 != int32(6) {
		v330 = v244
		goto L74
	} else {
		goto L75
	}
L73:
	;
	goto L69
L74:
	;
	v332 = v243 + int32(1)
	if v332 < v330 {
		v243 = v332
		v244 = v330
		goto L72
	} else {
		goto L95
	}
L75:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v259 != v260 {
		v330 = v244
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v262 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v255)+8)))
	v263 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
	if v262 != v263 {
		v330 = v244
		goto L74
	} else {
		goto L77
	}
L77:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v255)+12))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v265 != v266 {
		v330 = v244
		goto L74
	} else {
		goto L78
	}
L78:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v255)+16))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v268 != v269 {
		v330 = v244
		goto L74
	} else {
		goto L79
	}
L79:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v255)+20))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v271 != v272 {
		v330 = v244
		goto L74
	} else {
		goto L80
	}
L80:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v255)+32))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v274 != v275 {
		v330 = v244
		goto L74
	} else {
		goto L81
	}
L81:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v255)+24))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v279 = int32(0)
	v286 = base.B2i32(v277|v278 == v279)
	if v277 == v279 {
		v325 = v286
		goto L83
	} else {
		goto L84
	}
L82:
	;
	if v325 != 0 {
		goto L68
	} else {
		goto L94
	}
L83:
	;
	goto L82
L84:
	;
	if v278 == int32(0) {
		v325 = v286
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v278)+4))
	if v292 != v293 {
		v325 = int32(0)
		goto L83
	} else {
		goto L86
	}
L86:
	;
	v295 = int32(1)
	if v292 <= v295 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v298 = v295
	goto L89
L88:
	;
	v298 = v292
	goto L89
L89:
	;
	v299 = int32(8)
	v304 = int32(0)
	goto L90
L90:
	;
	v312 = v304 << (uint(int32(2)) % 32)
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v277+v299+v312)))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v312+(v278+v299))))
	v317 = base.B2i32(v314 == v316)
	if v316 != v314 {
		v325 = v317
		goto L83
	} else {
		goto L92
	}
L91:
	;
	v325 = v317
	goto L83
L92:
	;
	v320 = v304 + int32(1)
	if v320 != v298 {
		v304 = v320
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v236)+4))
	v330 = v329
	goto L74
L95:
	;
	goto L73
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v341)+28)) = int32(0)
	v346 = F_palloc0(m, int32(12))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L39
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v346)+4)) = v341
	*(*int32)(unsafe.Add(mBase, uint32(v346))) = int32(326)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v231)+8))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v351)+64))
	if v352 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)+4))
	v355 = v353
	goto L100
L99:
	;
	v355 = int32(0)
	goto L100
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v346)+8)) = v355
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v231)+8))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v357)+64))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v341)+12))
	v360 = F_lappend_oid(m, v358, v359)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L39
	} else {
		goto L101
	}
L101:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v231)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v362)+64)) = v360
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v231)+20))
	v365 = F_lappend(m, v364, v346)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L39
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+20)) = v365
	v379 = v346 + int32(8)
	goto L67
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v382)+8)) = v380
	*(*int64)(unsafe.Add(mBase, uint32(v382))) = int64(4294967304)
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v382)+12)) = v387
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v382)+16)) = v389
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v382)+20)) = v391
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v382)+24)) = v393
	return v382
L104:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v399 == int32(0) {
		v447 = l1
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v452 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L39
	} else {
		goto L118
	}
L106:
	;
	v403 = v399 & int32(7)
	if v403 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	if base.Ui32(v399) < base.Ui32(int32(8)) {
		v447 = v421
		goto L105
	} else {
		goto L114
	}
L108:
	;
	v419 = v399
	v421 = l1
	goto L107
L109:
	;
	goto L110
L110:
	;
	v406 = v399
	v408 = l1
	v409 = v3
	goto L111
L111:
	;
	v413 = int32(1)
	v414 = v406 - v413
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v408)+16))
	v417 = v409 + v413
	if v417 != v403 {
		v406 = v414
		v408 = v415
		v409 = v417
		goto L111
	} else {
		goto L113
	}
L112:
	;
	v419 = v414
	v421 = v415
	goto L107
L113:
	;
	goto L112
L114:
	;
	v428 = v419
	v430 = v421
	goto L115
L115:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v430)+16))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v435)+16))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v436)+16))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v437)+16))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v438)+16))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v439)+16))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v440)+16))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v441)+16))
	v444 = v428 - int32(8)
	if v444 != 0 {
		v428 = v444
		v430 = v442
		goto L115
	} else {
		goto L117
	}
L116:
	;
	v447 = v442
	goto L105
L117:
	;
	goto L116
L118:
	;
	v454 = int32(0)
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v452)+52))
	F_IncrementVarSublevelsUp(m, v452, v454-v455, v454)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L39
	} else {
		goto L119
	}
L119:
	;
	v461 = F_palloc0(m, int32(12))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L39
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v461)+4)) = v452
	*(*int32)(unsafe.Add(mBase, uint32(v461))) = int32(326)
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v447)+8))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v466)+64))
	if v467 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v467)+4))
	v470 = v468
	goto L123
L122:
	;
	v470 = int32(0)
	goto L123
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v461)+8)) = v470
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v447)+8))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v472)+64))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v452)+8))
	v475 = F_lappend_oid(m, v473, v474)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L39
	} else {
		goto L124
	}
L124:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v447)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v477)+64)) = v475
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v447)+20))
	v480 = F_lappend(m, v479, v461)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L39
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v447)+20)) = v480
	v484 = F_palloc0(m, int32(28))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L39
	} else {
		goto L126
	}
L126:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v484))) = int64(4294967304)
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v461)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v484)+8)) = v488
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v452)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v484)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v484)+12)) = v490
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v452)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v484)+20)) = v494
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v452)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v484)+24)) = v496
	return v484
L127:
	;
	v502 = F_exprType(m, l0)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L39
	} else {
		goto L128
	}
L128:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v504 == int32(0) {
		v552 = l1
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v557 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L39
	} else {
		goto L142
	}
L130:
	;
	v508 = v504 & int32(7)
	if v508 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	if base.Ui32(v504) < base.Ui32(int32(8)) {
		v552 = v526
		goto L129
	} else {
		goto L138
	}
L132:
	;
	v524 = v504
	v526 = l1
	goto L131
L133:
	;
	goto L134
L134:
	;
	v511 = v504
	v513 = l1
	v514 = v3
	goto L135
L135:
	;
	v518 = int32(1)
	v519 = v511 - v518
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v513)+16))
	v522 = v514 + v518
	if v522 != v508 {
		v511 = v519
		v513 = v520
		v514 = v522
		goto L135
	} else {
		goto L137
	}
L136:
	;
	v524 = v519
	v526 = v520
	goto L131
L137:
	;
	goto L136
L138:
	;
	v533 = v524
	v535 = v526
	goto L139
L139:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v535)+16))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v540)+16))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v541)+16))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v542)+16))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v543)+16))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v544)+16))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v545)+16))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v546)+16))
	v549 = v533 - int32(8)
	if v549 != 0 {
		v533 = v549
		v535 = v547
		goto L139
	} else {
		goto L141
	}
L140:
	;
	v552 = v547
	goto L129
L141:
	;
	goto L140
L142:
	;
	v559 = int32(0)
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v557)+16))
	F_IncrementVarSublevelsUp(m, v557, v559-v560, v559)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L39
	} else {
		goto L143
	}
L143:
	;
	v566 = F_palloc0(m, int32(12))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L39
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v566)+4)) = v557
	*(*int32)(unsafe.Add(mBase, uint32(v566))) = int32(326)
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v552)+8))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v571)+64))
	if v572 != 0 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	v575 = v573
	goto L147
L146:
	;
	v575 = int32(0)
	goto L147
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v566)+8)) = v575
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v552)+8))
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v577)+64))
	v579 = F_lappend_oid(m, v578, v502)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L39
	} else {
		goto L148
	}
L148:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v552)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v581)+64)) = v579
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v552)+20))
	v584 = F_lappend(m, v583, v566)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L39
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v552)+20)) = v584
	v588 = F_palloc0(m, int32(28))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L39
	} else {
		goto L150
	}
L150:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v588))) = int64(4294967304)
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v566)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v588)+16)) = int64(4294967295)
	*(*int32)(unsafe.Add(mBase, uint32(v588)+12)) = v502
	*(*int32)(unsafe.Add(mBase, uint32(v588)+8)) = v592
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v557)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v588)+24)) = v597
	return v588
L151:
	;
	v604 = F_exprType(m, l0)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L39
	} else {
		goto L153
	}
L152:
	;
	return v645
L153:
	;
	v607 = l1
	goto L155
L154:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L39
	} else {
		goto L167
	}
L155:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v607)+16))
	if v613 == int32(0) {
		goto L154
	} else {
		goto L157
	}
L156:
	;
	v620 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L39
	} else {
		goto L159
	}
L157:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v613)+4))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v616)+4))
	if v617 != int32(5) {
		v607 = v613
		goto L155
	} else {
		goto L158
	}
L158:
	;
	goto L156
L159:
	;
	v623 = F_palloc0(m, int32(12))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L39
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+4)) = v620
	*(*int32)(unsafe.Add(mBase, uint32(v623))) = int32(326)
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v613)+8))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v628)+64))
	if v629 != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v629)+4))
	v632 = v630
	goto L163
L162:
	;
	v632 = int32(0)
	goto L163
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+8)) = v632
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v613)+8))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v634)+64))
	v636 = F_lappend_oid(m, v635, v604)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L39
	} else {
		goto L164
	}
L164:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v613)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v638)+64)) = v636
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v613)+20))
	v641 = F_lappend(m, v640, v623)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L39
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v613)+20)) = v641
	v645 = F_palloc0(m, int32(28))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L39
	} else {
		goto L166
	}
L166:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v645))) = int64(4294967304)
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v623)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v645)+16)) = int64(4294967295)
	*(*int32)(unsafe.Add(mBase, uint32(v645)+12)) = v604
	*(*int32)(unsafe.Add(mBase, uint32(v645)+8)) = v649
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v620)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v645)+24)) = v654
	goto L152
L167:
	;
	F_errmsg_internal(m, int32(521320), int32(0))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L39
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(477473), int32(334), int32(76938))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L39
	} else {
		goto L169
	}
L169:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L170:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v674 = F_exprType(m, v673)
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L39
	} else {
		goto L171
	}
L171:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v676 == int32(0) {
		v724 = l1
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v729 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L39
	} else {
		goto L185
	}
L173:
	;
	v680 = v676 & int32(7)
	if v680 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	if base.Ui32(v676) < base.Ui32(int32(8)) {
		v724 = v698
		goto L172
	} else {
		goto L181
	}
L175:
	;
	v696 = v676
	v698 = l1
	goto L174
L176:
	;
	goto L177
L177:
	;
	v683 = v676
	v685 = l1
	v686 = v3
	goto L178
L178:
	;
	v690 = int32(1)
	v691 = v683 - v690
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v685)+16))
	v694 = v686 + v690
	if v694 != v680 {
		v683 = v691
		v685 = v692
		v686 = v694
		goto L178
	} else {
		goto L180
	}
L179:
	;
	v696 = v691
	v698 = v692
	goto L174
L180:
	;
	goto L179
L181:
	;
	v705 = v696
	v707 = v698
	goto L182
L182:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v707)+16))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v712)+16))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v713)+16))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v714)+16))
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v715)+16))
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v716)+16))
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v717)+16))
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v718)+16))
	v721 = v705 - int32(8)
	if v721 != 0 {
		v705 = v721
		v707 = v719
		goto L182
	} else {
		goto L184
	}
L183:
	;
	v724 = v719
	goto L172
L184:
	;
	goto L183
L185:
	;
	v731 = int32(0)
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v729)+4))
	F_IncrementVarSublevelsUp(m, v729, v731-v732, v731)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L39
	} else {
		goto L186
	}
L186:
	;
	v738 = F_palloc0(m, int32(12))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L39
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v738)+4)) = v729
	*(*int32)(unsafe.Add(mBase, uint32(v738))) = int32(326)
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v724)+8))
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v743)+64))
	if v744 != 0 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v744)+4))
	v747 = v745
	goto L190
L189:
	;
	v747 = int32(0)
	goto L190
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v738)+8)) = v747
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v724)+8))
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v749)+64))
	v751 = F_lappend_oid(m, v750, v674)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L39
	} else {
		goto L191
	}
L191:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v724)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v753)+64)) = v751
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v724)+20))
	v756 = F_lappend(m, v755, v738)
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L39
	} else {
		goto L192
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v724)+20)) = v756
	v760 = F_palloc0(m, int32(28))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L39
	} else {
		goto L193
	}
L193:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v760))) = int64(4294967304)
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v738)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v760)+12)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v760)+8)) = v764
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v729)+12))
	v768 = F_exprTypmod(m, v767)
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L39
	} else {
		goto L194
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v760)+16)) = v768
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v729)+12))
	v772 = F_exprCollation(m, v771)
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L39
	} else {
		goto L195
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v760)+20)) = v772
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v729)+12))
	v776 = F_exprLocation(m, v775)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v760)+24)) = v776
	return v760
L196:
	;
	return v780
}
func F_replace_nestloop_params_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
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
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int64
	_ = v194
	var v196 int64
	_ = v196
	var v198 int64
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12 != int32(319) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v316 = F_expression_tree_mutator_impl(m, l0, int32(829), l1)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L12
	} else {
		goto L83
	}
L5:
	;
	if v12 != int32(6) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v129 = F_find_placeholder_info(m, l1, l0)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L12
	} else {
		goto L37
	}
L8:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v17 < int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return l0
L10:
	;
	goto L11
L11:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+352))
	v22 = F_bms_is_member(m, v17, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	if v22 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return l0
L15:
	;
	goto L16
L16:
	;
	v29 = int32(0)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+356))
	if v30 == v29 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	return v127
L18:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v83 = F_palloc0(m, int32(28))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L12
	} else {
		goto L29
	}
L19:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v33 <= int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v38 = v29
	goto L21
L21:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43+v38<<(uint(int32(2))%32))))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v49 = F_equal(m, l0, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L12
	} else {
		goto L23
	}
L22:
	;
	v58 = F_palloc0(m, int32(28))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L12
	} else {
		goto L28
	}
L23:
	;
	if v49 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v54 = v38 + int32(1)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v54 < v55 {
		v38 = v54
		goto L21
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	goto L22
L27:
	;
	goto L18
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v58))) = int64(4294967304)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+12)) = v64
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+16)) = v66
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+20)) = v68
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = v70
	v127 = v58
	goto L17
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = int64(4294967304)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+64))
	if v88 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v91 = v89
	goto L32
L31:
	;
	v91 = int32(0)
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v91
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+64))
	v95 = F_lappend_oid(m, v94, v81)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+64)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v83)+24)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v83)+20)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v83)+12)) = v81
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v83)+24)) = v104
	v107 = F_palloc0(m, int32(12))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L12
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = int32(357)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v107)+4)) = v111
	v113 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L12
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+8)) = v113
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+356))
	v117 = F_lappend(m, v116, v107)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L12
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+356)) = v117
	v127 = v83
	goto L17
L37:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+12))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+352))
	v133 = int32(0)
	if v131 == v133 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v186 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L39:
	;
	v186 = int32(1)
	goto L38
L40:
	;
	goto L41
L41:
	;
	if v132 == int32(0) {
		v177 = v133
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v186 = v177
	goto L38
L43:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	if v143 < v142 {
		v177 = v133
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v145 = int32(1)
	if v142 <= v145 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v148 = v145
	goto L47
L46:
	;
	v148 = v142
	goto L47
L47:
	;
	v149 = int32(8)
	v154 = int32(0)
	goto L48
L48:
	;
	v161 = v154 << (uint(int32(2)) % 32)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v131+v149+v161)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v161+(v132+v149))))
	v168 = v163 & (v165 ^ int32(-1))
	v170 = base.B2i32(v168 == int32(0))
	if v168 != 0 {
		v177 = v170
		goto L42
	} else {
		goto L50
	}
L49:
	;
	v177 = v170
	goto L42
L50:
	;
	v172 = v154 + int32(1)
	if v172 != v148 {
		v154 = v172
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v190 = F_palloc0(m, int32(24))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L12
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v205 = int32(0)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)+356))
	if v206 == v205 {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190))) = int32(319)
	v194 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v190)+8)) = v194
	v196 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v190)+16)) = v196
	v198 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v190))) = v198
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v201 = F_replace_nestloop_params_mutator(m, v200, l1)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L12
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190)+4)) = v201
	return v190
L57:
	;
	return v313
L58:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v262 = F_exprType(m, v261)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L12
	} else {
		goto L72
	}
L59:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	if v209 <= int32(0) {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v214 = v205
	goto L61
L61:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v206)+12))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v219+v214<<(uint(int32(2))%32))))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
	v225 = F_equal(m, l0, v224)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L12
	} else {
		goto L63
	}
L62:
	;
	v234 = F_palloc0(m, int32(28))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L12
	} else {
		goto L68
	}
L63:
	;
	if v225 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v230 = v214 + int32(1)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	if v230 < v231 {
		v214 = v230
		goto L61
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	goto L62
L67:
	;
	goto L58
L68:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v234))) = int64(4294967304)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v234)+8)) = v238
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v241 = F_exprType(m, v240)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L12
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234)+12)) = v241
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v245 = F_exprTypmod(m, v244)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L12
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234)+16)) = v245
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v249 = F_exprCollation(m, v248)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L12
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234)+24)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v234)+20)) = v249
	v313 = v234
	goto L57
L72:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v265 = F_exprTypmod(m, v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L12
	} else {
		goto L73
	}
L73:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v268 = F_exprCollation(m, v267)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L12
	} else {
		goto L74
	}
L74:
	;
	v271 = F_palloc0(m, int32(28))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L12
	} else {
		goto L75
	}
L75:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v271))) = int64(4294967304)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)+64))
	if v276 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	v279 = v277
	goto L78
L77:
	;
	v279 = int32(0)
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v271)+8)) = v279
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)+64))
	v283 = F_lappend_oid(m, v282, v262)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L12
	} else {
		goto L79
	}
L79:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v285)+64)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v271)+24)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v271)+20)) = v268
	*(*int32)(unsafe.Add(mBase, uint32(v271)+16)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v271)+12)) = v262
	v293 = F_palloc0(m, int32(12))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L12
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v293))) = int32(357)
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v271)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v293)+4)) = v297
	v299 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L12
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v293)+8)) = v299
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l1)+356))
	v303 = F_lappend(m, v302, v293)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L12
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+356)) = v303
	v313 = v271
	goto L57
L83:
	;
	return v316
}
func F_report_invalid_encoding_db(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v5 = *(*int32)(unsafe.Add(mBase, _consts[356]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	F_report_invalid_encoding_int(m, v6, l0, l1, l2)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		base.Wasm_trap_unreachable()
		for {
		}
	}
}
func F_reset_formatted_start_time(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[1185])) = uint8(v2)
	return
}
func F_resolve_anyrange_from_others(m *base.Module, l0 int32) {
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v9 != 0 {
		v10 = F_getBaseType(m, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v12 = F_get_multirange_range(m, v10)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				if v12 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						F_errcode(m, int32(67141764))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							v40 = F_format_type_be(m, v10)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v40
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(384839)
								F_errmsg(m, int32(179860), v7)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return
								} else {
									F_errfinish(m, int32(478493), int32(699), int32(126375))
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
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
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v12
					m.G0 = v7 + int32(16)
					return
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(356155), int32(0))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				F_errfinish(m, int32(478493), int32(703), int32(126375))
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
func F_rowtype_field_matches(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	if l0 == int32(2249) {
		v54 = int32(1)
		return v54
	} else {
		v11 = int32(0)
		v13 = F_lookup_rowtype_tupdesc_domain(m, l0, int32(-1))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			if int32(0) < l1 {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				if l1 <= v19 {
					v30 = v13 + v19<<(uint(int32(4))%32) + l1*int32(100)
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+11)))
					if v31 != 0 {
						v41 = int32(0)
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
						if v41 <= v42 {
							v49 = v41
							F_DecrTupleDescRefCount(m, v13)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								v54 = v49
								return v54
							}
						} else {
							v54 = v41
							return v54
						}
					} else {
						v33 = v30 - int32(80)
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+68))
						if v34 != l2 {
							v41 = int32(0)
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
							if v41 <= v42 {
								v49 = v41
								F_DecrTupleDescRefCount(m, v13)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									v54 = v49
									return v54
								}
							} else {
								v54 = v41
								return v54
							}
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+76))
							if v36 != l3 {
								v41 = int32(0)
								v42 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
								if v41 <= v42 {
									v49 = v41
									F_DecrTupleDescRefCount(m, v13)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										v54 = v49
										return v54
									}
								} else {
									v54 = v41
									return v54
								}
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+96))
								if v38 == l4 {
									v45 = int32(1)
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
									if v46 < int32(0) {
										v54 = v45
										return v54
									} else {
										v49 = v45
										F_DecrTupleDescRefCount(m, v13)
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int32(0)
										} else {
											v54 = v49
											return v54
										}
									}
								} else {
									v41 = int32(0)
									v42 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
									if v41 <= v42 {
										v49 = v41
										F_DecrTupleDescRefCount(m, v13)
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int32(0)
										} else {
											v54 = v49
											return v54
										}
									} else {
										v54 = v41
										return v54
									}
								}
							}
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
					if int32(0) <= v22 {
						v49 = v11
						F_DecrTupleDescRefCount(m, v13)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							v54 = v49
							return v54
						}
					} else {
						v54 = v11
						return v54
					}
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
				if int32(0) <= v22 {
					v49 = v11
					F_DecrTupleDescRefCount(m, v13)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						v54 = v49
						return v54
					}
				} else {
					v54 = v11
					return v54
				}
			}
		}
	}
}
func F_rtrim1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = int32(1)
		v11 = v6 + v10
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
		v16 = v14 & v10
		if v16 != 0 {
			v17 = v11
		} else {
			v17 = v6 + int32(4)
		}
		if v14 == int32(1) {
			v20 = int32(4)
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
			if v22&int32(254) == int32(2) {
				v31 = v20
			} else {
				v31 = base.B2i32(v22 == int32(18)) << (uint(v20) % 32)
			}
			if v22 == int32(1) {
				v34 = v20
			} else {
				v34 = v31
			}
			v45 = v34
		} else {
			v35 = int32(1)
			if v16 != 0 {
				v45 = int32(base.Ui32(v14)>>(uint(v35)%32)) - v35
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
				v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v47 = int32(1)
		v50 = F_dotrim(m, v17, v45, int32(708501), v47, int32(0), v47)
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			return v50
		}
	}
}
