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
func Fn13843(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int32 {
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	v9 = int32(4)
	v10 = base.I32_wrap_i64(l0)<<(uint(l3)%32) | v9
	v12 = base.I32_wrap_i64(l1) << (uint(l3) % 32)
	v14 = v12 | v9
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v14))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v10)) == int32(0) {
		v26 = base.B2i32(base.Ui32(v10) < base.Ui32(v14))
	} else {
		v26 = int32(base.Ui32(v10-v14) >> (uint(int32(31)) % 32))
	}
	if v26 != 0 {
		v27 = v12 + l2
		if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v27))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v10)) == int32(0) {
			v39 = base.B2i32(base.Ui32(v10) < base.Ui32(v27))
		} else {
			v39 = int32(base.Ui32(v10-v27) >> (uint(int32(31)) % 32))
		}
		v41 = v39
	} else {
		v41 = int32(0)
	}
	return v41
}
func Fn13849(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = F_SearchSysCache1(m, l5, l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
				F_errmsg_internal(m, l4, v11)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, l3, l2, l1)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v28+v29)+4))
			F_ReleaseCatCache(m, v13)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				m.G0 = v11 + int32(16)
				return v31
			}
		}
	}
}
func Fn13850(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	v3 = int32(0)
	v9 = F_SearchSysCacheList(m, l1, int32(1), l0, v3, v3)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
	if int32(0) < v13 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v19 = int32(0)
	v21 = v3
	goto L6
L4:
	;
	v40 = v3
	goto L5
L5:
	;
	F_ReleaseCatCacheList(m, v9)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L10
	}
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(48)+v19<<(uint(int32(2))%32))))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+56))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28+v29)+4))
	v32 = F_lappend_oid(m, v21, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v40 = v32
	goto L5
L8:
	;
	v35 = v19 + int32(1)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
	if v35 < v36 {
		v19 = v35
		v21 = v32
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	return v40
}
func Fn13856(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
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
	F_errfinish(m, int32(_a_Fn13856_0), l4, l3)
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
	v49 = *(*int32)(unsafe.Add(mBase, _c_Fn13856[0]))
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
	v96 = *(*int32)(unsafe.Add(mBase, _c_Fn13856[0]))
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
	v105 = *(*int32)(unsafe.Add(mBase, _c_Fn13856[1]))
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
	v133 = *(*int32)(unsafe.Add(mBase, _c_Fn13856[1]))
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
func Fn13867(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	if int32(0) <= l1 {
		if base.Ui32(l1) < base.Ui32(int32(7)) {
			v44 = l1
			m.G0 = v13 + int32(32)
			return v44
		} else {
			v19 = int32(6)
			v22 = F_errstart(m, int32(19), int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				if v22 == int32(0) {
					v44 = v19
					m.G0 = v13 + int32(32)
					return v44
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = int32(6)
						if l0 != 0 {
							v35 = int32(_a_Fn13867_0)
						} else {
							v35 = int32(_a_Fn13867_1)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v35
						*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l1
						F_errmsg(m, l7, v13+int32(16))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, l4, l6, l2)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								v44 = v19
								m.G0 = v13 + int32(32)
								return v44
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return int32(0)
			} else {
				if l0 != 0 {
					v58 = int32(_a_Fn13867_0)
				} else {
					v58 = int32(_a_Fn13867_1)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v58
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
				F_errmsg(m, l5, v13)
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, l4, l3, l2)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
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
func Fn13872(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
					F_errmsg_internal(m, int32(_a_Fn13872_0), int32(0))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(11)
						F_errdetail_internal(m, int32(_a_Fn13872_1), v11)
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
func Fn13887(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
func Fn13892(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	if v11 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v23
				F_errmsg(m, l3, v8)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_Fn13892_0), l2, l1)
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
		}
	} else {
		m.G0 = v8 + int32(16)
		return int32(0)
	}
}
func Fn13894(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	v14 = int32(_a_Fn13894_0)
	v16 = *(*int32)(unsafe.Add(mBase, _c_Fn13894[0]))
	*(*int32)(unsafe.Add(mBase, _c_Fn13894[0])) = v16 + int32(1)
	v21 = *(*int32)(unsafe.Add(mBase, _c_Fn13894[1]))
	if int32(0) <= v21 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = int32(_a_Fn13894_1)
	v25 = *(*int32)(unsafe.Add(mBase, _c_Fn13894[2]))
	v28 = v21 * int32(100)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_Fn13894[3])))
	*(*int32)(unsafe.Add(mBase, _c_Fn13894[2])) = v31
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
	*(*int32)(unsafe.Add(mBase, _c_Fn13894[1])) = int32(-1)
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
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_Fn13894[4])))
	*(*int32)(unsafe.Add(mBase, _c_Fn13894[5])) = v38
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
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_Fn13894[6])))
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
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_Fn13894[4])))
	*(*int32)(unsafe.Add(mBase, _c_Fn13894[5])) = v57
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
	*(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_Fn13894[6]))) = v75
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
	*(*int32)(unsafe.Add(mBase, _c_Fn13894[2])) = v25
	v83 = int32(_a_Fn13894_0)
	v85 = *(*int32)(unsafe.Add(mBase, _c_Fn13894[0]))
	*(*int32)(unsafe.Add(mBase, _c_Fn13894[0])) = v85 - int32(1)
	m.G0 = v12 + int32(32)
	return
L21:
	;
	F_errmsg_internal(m, int32(_a_Fn13894_2), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_Fn13894_3), l3, l2)
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
func Fn13900(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	v4 = l0
	goto L1
L1:
	;
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	if v7 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	return v14
L3:
	;
	goto L2
L4:
	;
	v14 = int32(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	if v7 == l1 {
		v14 = v4
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v4 = v4 + int32(1)
	goto L1
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 float64
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v14 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v13 + v14
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v13
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+16)))
	v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22+v23)+12)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = F_gbt_num_distance(m, v8+v14, v8+int32(12), v25&int32(1), l1, v28)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		return int32(0)
	} else {
		v33 = F_Float8GetDatum(m, v29)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			m.G0 = v8 + int32(16)
			return v33
		}
	}
}
func Fn13917(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
			v17 = F_pstrdup(m, v12+v13+int32(4))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_ReleaseCatCache(m, v4)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					return v17
				}
			}
		}
	}
}
func Fn13922(m *base.Module, l0 int32, l1 int32) int32 {
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
func Fn13924(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = F_SearchSysCache1(m, l6, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if v15 == int32(0) {
			if l1 != 0 {
				v40 = int32(0)
				m.G0 = v13 + int32(16)
				return v40
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
					F_errmsg_internal(m, l5, v13)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_Fn13924_0), l4, l3)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
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
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+22)))
			v35 = F_pstrdup(m, v31+v32+l2)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				F_ReleaseCatCache(m, v15)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					v40 = v35
					m.G0 = v13 + int32(16)
					return v40
				}
			}
		}
	}
}
func Fn13937(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
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
		v22 = F_get_role_oid_or_public(m, v14)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = F_convert_any_priv_string(m, v16, l1)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v28 = F_object_aclcheck_ext(m, l2, v13, v22, v24, v11+int32(15))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
					if v30 == int32(1) {
						v33 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v33)
						v37 = int32(0)
					} else {
						v37 = base.B2i32(v28 == int32(0))
					}
					m.G0 = v11 + int32(16)
					return v37
				}
			}
		}
	}
}
func Fn13939(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
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
			v26 = *(*int32)(unsafe.Add(mBase, _c_Fn13939[0]))
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
									F_errfinish(m, int32(_a_Fn13939_0), l4, l3)
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
func Fn13944(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
						switch v46&int32(_a_Fn13944_0) - int32(1) {
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
								F_errmsg_internal(m, int32(_a_Fn13944_1), v11)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, l4, int32(70), int32(_a_Fn13944_2))
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
func Fn13953(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
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
		v15 = F_pg_detoast_datum_packed(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = F_text_to_cstring(m, v15)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v17
				v24 = F_get_worker(m, v10, v7+int32(12), int32(0), int32(1), l1)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					if v24 == int32(0) {
						v28 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
						v31 = int32(0)
					} else {
						v31 = v24
					}
					m.G0 = v7 + int32(16)
					return v31
				}
			}
		}
	}
}
func Fn13955(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v9
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v9
	v14 = F_pushJsonbValue(m, v7, l2, int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v19 = F_pushJsonbValue(m, v7, l1, int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v19
			v22 = F_JsonbValueToJsonb(m, v19)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				m.G0 = v7 + int32(16)
				return v22
			}
		}
	}
}
func Fn13964(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l3
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)))
	if v11 == int32(1) {
		v14 = int32(_a_Fn13964_0)
		v15 = *(*int32)(unsafe.Add(mBase, _c_Fn13964[0]))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		*(*int32)(unsafe.Add(mBase, _c_Fn13964[0])) = v17
		v20 = F_palloc(m, int32(40))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v20)+8)) = uint8(v24)
			*(*int64)(unsafe.Add(mBase, uint32(v20))) = int64(0)
			F_initHyperLogLog(m, v20+int32(16), int32(10))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(116)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v20
				*(*int32)(unsafe.Add(mBase, _c_Fn13964[0])) = v15
				return int32(0)
			}
		}
	} else {
		return int32(0)
	}
}
func Fn13973(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v16 = F_numeric_stddev_internal(m, v13, l2, l1, v8+int32(15))
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
func Fn13979(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(34209794)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v9
	v16 = *(*int32)(unsafe.Add(mBase, _c_Fn13979[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v16
	v19 = F_LockRelease(m, v7, l1, int32(1))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v19
	}
}
func Fn13980(m *base.Module, l0 int32, l1 int32) int32 {
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
	v18 = *(*int32)(unsafe.Add(mBase, _c_Fn13980[0]))
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
func Fn13991(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
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
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_Fn13991[0]))))
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
func Fn13997(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	if l3 < v12 {
		m.G0 = v9 + int32(16)
		return int32(0)
	} else {
		v14 = F_strlen(m, l1)
		mBase = m.M
		if base.Ui32(int32(63)) < base.Ui32(v14) {
			m.G0 = v9 + int32(16)
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v21 = F_hash_search(m, v17, l1, int32(1), v9+int32(15))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v25
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v27 != 0 {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
					v30 = v29 - v27
					v33 = F_palloc(m, v30+int32(1))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						if v30 != 0 {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							base.MemoryCopy(m, v33, v35, v30)
						} else {
						}
						v38 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v33+v30))) = uint8(v38)
						v41 = v33
						*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v41
						m.G0 = v9 + int32(16)
						return int32(0)
					}
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v41 = v40
					*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v41
					m.G0 = v9 + int32(16)
					return int32(0)
				}
			}
		}
	}
}
func Fn14004(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v4 = int32(0)
	v6 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v6 == v4 {
		v28 = v4
		return v28
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v11 = v9 - int32(1)
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v11 <= v12 {
			v28 = v4
			return v28
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v11))))
			if v16 != l2 {
				v28 = v4
				return v28
			} else {
				v19 = F_find_among_b(m, l0, l1, int32(4))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					if v19 == int32(0) {
						v28 = v4
					} else {
						v26 = Fn14003(m, l0, int32(121))
						mBase = m.M
						v28 = v26
					}
					return v28
				}
			}
		}
	}
}
func Fn14008(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = l1
	return v4
}
func Fn14013(m *base.Module, l0 int32, l1 int32) {
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
		v20 = int32(_a_Fn14013_0)
	} else {
		v20 = int32(_a_Fn14013_1)
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
func Fn14015(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	v2 = l1
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	if v9 <= v6+int32(1) {
		F_appendStringInfoChar(m, v5, v2)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		*(*uint8)(unsafe.Add(mBase, uint32(v17+v6))) = uint8(v2)
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
		v23 = v21 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v23
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
		v27 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v25+v23))) = uint8(v27)
		return v27
	}
}
func Fn14019(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
						F_errmsg_internal(m, int32(_a_Fn14019_0), int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_Fn14019_1), int32(327), l3)
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
				F_errmsg_internal(m, int32(_a_Fn14019_0), int32(0))
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_Fn14019_1), int32(327), l3)
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
func Fn14024(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if int32(0) <= v8 {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
		v14 = F_psprintf(m, int32(_a_Fn14024_0), v6)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v20 = v14
			m.G0 = v6 + int32(16)
			return v20
		}
	} else {
		v18 = F_pstrdup(m, l1)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = v18
			m.G0 = v6 + int32(16)
			return v20
		}
	}
}
func Fn14035(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
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
func Fn14042(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
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
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = F_pg_detoast_datum_packed(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v20 = int32(1)
			v21 = v16 + v20
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
			v24 = v22 & v20
			if v24 != 0 {
				v25 = v21
			} else {
				v25 = v16 + int32(4)
			}
			if v22 == int32(1) {
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
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
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					v52 = int32(base.Ui32(v46)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v53 = int32(1)
			v54 = v11 + v53
			v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
			v59 = v57 & v53
			if v59 != 0 {
				v60 = v54
			} else {
				v60 = v11 + int32(4)
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
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
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
						if v94 != v16 {
							F_pfree(m, v16)
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
					if v94 != v16 {
						F_pfree(m, v16)
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
func Fn14044(m *base.Module, l0 int32, l1 int32) int32 {
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
								return base.I32_reinterpret_f32(base.F32_sub(float32(1), v87))
							}
						} else {
							return base.I32_reinterpret_f32(base.F32_sub(float32(1), v87))
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
							return base.I32_reinterpret_f32(base.F32_sub(float32(1), v87))
						}
					} else {
						return base.I32_reinterpret_f32(base.F32_sub(float32(1), v87))
					}
				}
			}
		}
	}
}
