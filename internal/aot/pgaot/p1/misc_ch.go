package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ChangeVarNodes_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
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
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	v3 = int32(0)
	if l0 == v3 {
		v166 = v3
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v168 + int32(1)
	v174 = F_query_tree_walker_impl(m, l0, int32(1049), l1, int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L7
	} else {
		goto L71
	}
L2:
	;
	return v166
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v9 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v10 = m.T0[v9].(func(*base.Module, int32, int32) int32)(m, l0, l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v14 - int32(58) {
	case 0:
		goto L15
	case 1, 2, 3, 4:
		v121 = v14
		goto L11
	case 5:
		goto L14
	case 6:
		goto L13
	default:
		goto L16
	}
L7:
	;
	return int32(0)
L8:
	;
	if v10 != 0 {
		v166 = v3
		goto L2
	} else {
		goto L9
	}
L9:
	;
	goto L6
L10:
	;
	v161 = F_expression_tree_walker_impl(m, l0, int32(1049), l1)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L7
	} else {
		goto L70
	}
L11:
	;
	if v121 == int32(67) {
		goto L1
	} else {
		goto L55
	}
L12:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v77 != v78 {
		goto L10
	} else {
		goto L38
	}
L13:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v71 != 0 {
		goto L10
	} else {
		goto L36
	}
L14:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v63 != 0 {
		v166 = v3
		goto L2
	} else {
		goto L34
	}
L15:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v55 != 0 {
		v166 = v3
		goto L2
	} else {
		goto L32
	}
L16:
	;
	if v14 == int32(319) {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	if v14 != int32(6) {
		v121 = v14
		goto L11
	} else {
		goto L18
	}
L18:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v21 != v22 {
		v166 = v3
		goto L2
	} else {
		goto L19
	}
L19:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v25 == v26 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v24
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v30 = v29
	goto L22
L21:
	;
	v30 = v25
	goto L22
L22:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v30 < int32(0) {
		v46 = v31
		goto L23
	} else {
		goto L24
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v46
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v48 != v49 {
		v166 = v3
		goto L2
	} else {
		goto L31
	}
L24:
	;
	v34 = F_bms_is_member(m, v30, v31)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	if v34 == int32(0) {
		v46 = v31
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v38 = F_bms_copy(m, v31)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	v40 = F_bms_del_member(m, v38, v30)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	if v24 < int32(0) {
		v46 = v40
		goto L23
	} else {
		goto L29
	}
L29:
	;
	v44 = F_bms_add_member(m, v40, v24)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	v46 = v44
	goto L23
L31:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v51
	return int32(0)
L32:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v56 != v57 {
		v166 = v3
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v59
	return int32(0)
L34:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v64 != v65 {
		v166 = v3
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v67
	return int32(0)
L36:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v72 != v73 {
		goto L10
	} else {
		goto L37
	}
L37:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v75
	goto L10
L38:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v81 < int32(0) {
		v97 = v80
		goto L39
	} else {
		goto L40
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v97
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v101 < int32(0) {
		v117 = v100
		goto L47
	} else {
		goto L48
	}
L40:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v85 = F_bms_is_member(m, v81, v80)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	if v85 == int32(0) {
		v97 = v80
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v89 = F_bms_copy(m, v80)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	v91 = F_bms_del_member(m, v89, v81)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	if v84 < int32(0) {
		v97 = v91
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v95 = F_bms_add_member(m, v91, v84)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	v97 = v95
	goto L39
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v117
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v121 = v120
	goto L11
L48:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v105 = F_bms_is_member(m, v101, v100)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L7
	} else {
		goto L49
	}
L49:
	;
	if v105 == int32(0) {
		v117 = v100
		goto L47
	} else {
		goto L50
	}
L50:
	;
	v109 = F_bms_copy(m, v100)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L7
	} else {
		goto L51
	}
L51:
	;
	v111 = F_bms_del_member(m, v109, v101)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L7
	} else {
		goto L52
	}
L52:
	;
	if v104 < int32(0) {
		v117 = v111
		goto L47
	} else {
		goto L53
	}
L53:
	;
	v115 = F_bms_add_member(m, v111, v104)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L7
	} else {
		goto L54
	}
L54:
	;
	v117 = v115
	goto L47
L55:
	;
	if v121 != int32(322) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	if v121 != int32(374) {
		goto L10
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v144 != 0 {
		goto L10
	} else {
		goto L65
	}
L59:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v130 != 0 {
		v166 = v3
		goto L2
	} else {
		goto L60
	}
L60:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v131 == v132 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v134
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v137 = v136
	goto L63
L62:
	;
	v137 = v131
	goto L63
L63:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v137 != v138 {
		v166 = v3
		goto L2
	} else {
		goto L64
	}
L64:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v140
	return int32(0)
L65:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v145 == v146 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v148
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v151 = v150
	goto L68
L67:
	;
	v151 = v145
	goto L68
L68:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v151 != v152 {
		goto L10
	} else {
		goto L69
	}
L69:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v154
	goto L10
L70:
	;
	v166 = v161
	goto L2
L71:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v176 - int32(1)
	return v174
}
func F_char_bpchar(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_palloc(m, int32(5))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+4)) = uint8(v3)
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(20)
		return v5
	}
}
func F_chargt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(base.Ui32(v3) < base.Ui32(v2))
}
func F_charle(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(base.Ui32(v2) <= base.Ui32(v3))
}
func F_checkMembershipInCurrentExtension(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
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
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
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
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_checkMembershipInCurrentExtension[0])))
	if v8 != int32(1) {
		m.G0 = v5 + int32(16)
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v13 = F_getExtensionOfObject(m, v11, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_checkMembershipInCurrentExtension[1]))
			if v13 == v16 {
				m.G0 = v5 + int32(16)
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						v26 = F_getObjectDescription(m, l0, int32(0))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, _c_F_checkMembershipInCurrentExtension[1]))
							v30 = F_get_extension_name(m, v29)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v30
								*(*int32)(unsafe.Add(mBase, uint32(v5))) = v26
								F_errmsg(m, int32(_a_F_checkMembershipInCurrentExtension_0), v5)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return
								} else {
									F_errdetail(m, int32(_a_F_checkMembershipInCurrentExtension_1), int32(0))
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_checkMembershipInCurrentExtension_2), int32(286), int32(_a_F_checkMembershipInCurrentExtension_3))
										mBase = m.M
										v45 = m.ExcPending
										if v45 != 0 {
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
			}
		}
	}
}
func F_check_application_name(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = F_pg_clean_ascii(m, v5, int32(2))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 != 0 {
			v12 = F_guc_strdup(m, int32(15), v7)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				if v12 == int32(0) {
					F_pfree(m, v7)
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					F_bms_free(m, v20)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						F_pfree(m, v7)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v12
							v28 = int32(1)
							return v28
						}
					}
				}
			}
		} else {
			v28 = int32(0)
			return v28
		}
	}
}
func F_check_canonical_path(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v4 != 0 {
		F_canonicalize_path_enc(m, v4)
		mBase = m.M
	} else {
	}
	return int32(1)
}
func F_check_db_file_conflict(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	v14 = F_table_open(m, int32(1213), int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+188))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	m.T0[v60].(func(*base.Module, int32))(m, v20)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L2
	} else {
		goto L17
	}
L2:
	;
	return int32(0)
L3:
	;
	v18 = int32(0)
	v20 = F_table_beginscan_catalog(m, v14, v18, v18)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v22 = F_heap_getnext(m, v20)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v22 == int32(0) {
		v55 = v2
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v27 = v22
	goto L7
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+22)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33+v34)))
	if v36 == int32(1664) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v55 = v2
	goto L1
L9:
	;
	v49 = F_heap_getnext(m, v20)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L2
	} else {
		goto L15
	}
L10:
	;
	v39 = F_GetDatabasePath(m, l0, v36)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v43 = F___fstatat(m, int32(-100), v39, v10, int32(256))
	mBase = m.M
	goto L12
L12:
	;
	F_pfree(m, v39)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	if v43 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v55 = int32(1)
	goto L1
L15:
	;
	if v49 != 0 {
		v27 = v49
		goto L7
	} else {
		goto L16
	}
L16:
	;
	goto L8
L17:
	;
	F_sequence_close(m, v14, int32(1))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	m.G0 = v10 + int32(96)
	return v55
}
func F_check_encoding_locale_matches(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v12 = F_pg_get_encoding_from_locale(m, l2, int32(1))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v15 = F_pg_get_encoding_from_locale(m, l1, int32(1))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			if l0 == v12 {
				if l0 == v15 {
					m.G0 = v9 - int32(-64)
					return
				} else {
					if base.Ui32(v15+int32(1)) < base.Ui32(int32(2)) {
						m.G0 = v9 - int32(-64)
						return
					} else {
						if l0 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return
								} else {
									if base.Ui32(l0) <= base.Ui32(int32(41)) {
										v97 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
										v98 = v97
									} else {
										v98 = int32(_a_F_check_encoding_locale_matches_0)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l1
									*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v98
									F_errmsg(m, int32(_a_F_check_encoding_locale_matches_1), v7+int32(-48))
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return
									} else {
										if base.Ui32(v15) <= base.Ui32(int32(41)) {
											v114 = *(*int32)(unsafe.Add(mBase, uint32(v15<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
											v115 = v114
										} else {
											v115 = int32(_a_F_check_encoding_locale_matches_0)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = v115
										F_errdetail(m, int32(_a_F_check_encoding_locale_matches_2), v9)
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_check_encoding_locale_matches_3), int32(1629), int32(_a_F_check_encoding_locale_matches_4))
											mBase = m.M
											v124 = m.ExcPending
											if v124 != 0 {
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
							v80 = F_superuser(m)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return
							} else {
								if v80 != 0 {
									m.G0 = v9 - int32(-64)
									return
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return
										} else {
											if base.Ui32(l0) <= base.Ui32(int32(41)) {
												v97 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
												v98 = v97
											} else {
												v98 = int32(_a_F_check_encoding_locale_matches_0)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l1
											*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v98
											F_errmsg(m, int32(_a_F_check_encoding_locale_matches_1), v7+int32(-48))
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return
											} else {
												if base.Ui32(v15) <= base.Ui32(int32(41)) {
													v114 = *(*int32)(unsafe.Add(mBase, uint32(v15<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
													v115 = v114
												} else {
													v115 = int32(_a_F_check_encoding_locale_matches_0)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v9))) = v115
												F_errdetail(m, int32(_a_F_check_encoding_locale_matches_2), v9)
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_check_encoding_locale_matches_3), int32(1629), int32(_a_F_check_encoding_locale_matches_4))
													mBase = m.M
													v124 = m.ExcPending
													if v124 != 0 {
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
						}
					}
				}
			} else {
				if base.Ui32(v12+int32(1)) < base.Ui32(int32(2)) {
					if l0 == v15 {
						m.G0 = v9 - int32(-64)
						return
					} else {
						if base.Ui32(v15+int32(1)) < base.Ui32(int32(2)) {
							m.G0 = v9 - int32(-64)
							return
						} else {
							if l0 != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return
									} else {
										if base.Ui32(l0) <= base.Ui32(int32(41)) {
											v97 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
											v98 = v97
										} else {
											v98 = int32(_a_F_check_encoding_locale_matches_0)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l1
										*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v98
										F_errmsg(m, int32(_a_F_check_encoding_locale_matches_1), v7+int32(-48))
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return
										} else {
											if base.Ui32(v15) <= base.Ui32(int32(41)) {
												v114 = *(*int32)(unsafe.Add(mBase, uint32(v15<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
												v115 = v114
											} else {
												v115 = int32(_a_F_check_encoding_locale_matches_0)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v9))) = v115
											F_errdetail(m, int32(_a_F_check_encoding_locale_matches_2), v9)
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_check_encoding_locale_matches_3), int32(1629), int32(_a_F_check_encoding_locale_matches_4))
												mBase = m.M
												v124 = m.ExcPending
												if v124 != 0 {
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
								v80 = F_superuser(m)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return
								} else {
									if v80 != 0 {
										m.G0 = v9 - int32(-64)
										return
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return
											} else {
												if base.Ui32(l0) <= base.Ui32(int32(41)) {
													v97 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
													v98 = v97
												} else {
													v98 = int32(_a_F_check_encoding_locale_matches_0)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l1
												*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v98
												F_errmsg(m, int32(_a_F_check_encoding_locale_matches_1), v7+int32(-48))
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return
												} else {
													if base.Ui32(v15) <= base.Ui32(int32(41)) {
														v114 = *(*int32)(unsafe.Add(mBase, uint32(v15<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
														v115 = v114
													} else {
														v115 = int32(_a_F_check_encoding_locale_matches_0)
													}
													*(*int32)(unsafe.Add(mBase, uint32(v9))) = v115
													F_errdetail(m, int32(_a_F_check_encoding_locale_matches_2), v9)
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_check_encoding_locale_matches_3), int32(1629), int32(_a_F_check_encoding_locale_matches_4))
														mBase = m.M
														v124 = m.ExcPending
														if v124 != 0 {
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
							}
						}
					}
				} else {
					if l0 == int32(0) {
						v24 = F_superuser(m)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							if v24 != 0 {
								if base.Ui32(v15+int32(1)) < base.Ui32(int32(2)) {
									m.G0 = v9 - int32(-64)
									return
								} else {
									v80 = F_superuser(m)
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return
									} else {
										if v80 != 0 {
											m.G0 = v9 - int32(-64)
											return
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v88 = m.ExcPending
												if v88 != 0 {
													return
												} else {
													if base.Ui32(l0) <= base.Ui32(int32(41)) {
														v97 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
														v98 = v97
													} else {
														v98 = int32(_a_F_check_encoding_locale_matches_0)
													}
													*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l1
													*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v98
													F_errmsg(m, int32(_a_F_check_encoding_locale_matches_1), v7+int32(-48))
													mBase = m.M
													v105 = m.ExcPending
													if v105 != 0 {
														return
													} else {
														if base.Ui32(v15) <= base.Ui32(int32(41)) {
															v114 = *(*int32)(unsafe.Add(mBase, uint32(v15<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
															v115 = v114
														} else {
															v115 = int32(_a_F_check_encoding_locale_matches_0)
														}
														*(*int32)(unsafe.Add(mBase, uint32(v9))) = v115
														F_errdetail(m, int32(_a_F_check_encoding_locale_matches_2), v9)
														mBase = m.M
														v119 = m.ExcPending
														if v119 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_check_encoding_locale_matches_3), int32(1629), int32(_a_F_check_encoding_locale_matches_4))
															mBase = m.M
															v124 = m.ExcPending
															if v124 != 0 {
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
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v29 = m.ExcPending
								if v29 != 0 {
									return
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v32 = m.ExcPending
									if v32 != 0 {
										return
									} else {
										if base.Ui32(l0) <= base.Ui32(int32(41)) {
											v41 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
											v42 = v41
										} else {
											v42 = int32(_a_F_check_encoding_locale_matches_0)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l2
										*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v42
										F_errmsg(m, int32(_a_F_check_encoding_locale_matches_1), v7+int32(-16))
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return
										} else {
											if base.Ui32(v12) <= base.Ui32(int32(41)) {
												v58 = *(*int32)(unsafe.Add(mBase, uint32(v12<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
												v59 = v58
											} else {
												v59 = int32(_a_F_check_encoding_locale_matches_0)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v59
											F_errdetail(m, int32(_a_F_check_encoding_locale_matches_5), v7+int32(-32))
											mBase = m.M
											v65 = m.ExcPending
											if v65 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_check_encoding_locale_matches_3), int32(1614), int32(_a_F_check_encoding_locale_matches_4))
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
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
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								if base.Ui32(l0) <= base.Ui32(int32(41)) {
									v41 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
									v42 = v41
								} else {
									v42 = int32(_a_F_check_encoding_locale_matches_0)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l2
								*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v42
								F_errmsg(m, int32(_a_F_check_encoding_locale_matches_1), v7+int32(-16))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									if base.Ui32(v12) <= base.Ui32(int32(41)) {
										v58 = *(*int32)(unsafe.Add(mBase, uint32(v12<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
										v59 = v58
									} else {
										v59 = int32(_a_F_check_encoding_locale_matches_0)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v59
									F_errdetail(m, int32(_a_F_check_encoding_locale_matches_5), v7+int32(-32))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_check_encoding_locale_matches_3), int32(1614), int32(_a_F_check_encoding_locale_matches_4))
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
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
			}
		}
	}
}
func F_check_escape_warning(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+56)))
	if v4 != int32(1) {
		v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v46 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v45)+56)) = uint8(v46)
		return
	} else {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+20)))
		if v7 != int32(1) {
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v46 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v45)+56)) = uint8(v46)
			return
		} else {
			v12 = F_errstart(m, int32(19), int32(0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				if v12 == int32(0) {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v46 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v45)+56)) = uint8(v46)
					return
				} else {
					F_errcode(m, int32(100794498))
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_check_escape_warning_0), int32(0))
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return
						} else {
							F_errhint(m, int32(_a_F_check_escape_warning_1), int32(0))
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return
							} else {
								v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
								if int32(0) <= v28 {
									v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
									v33 = F_pg_mbstrlen_with_len(m, v32, v28)
									mBase = m.M
									v34 = m.ExcPending
									if v34 != 0 {
										return
									} else {
										v37 = F_errposition(m, v33+int32(1))
										mBase = m.M
										v38 = m.ExcPending
										if v38 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_check_escape_warning_2), int32(1458), int32(_a_F_check_escape_warning_3))
											mBase = m.M
											v43 = m.ExcPending
											if v43 != 0 {
												return
											} else {
												v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v46 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v45)+56)) = uint8(v46)
												return
											}
										}
									}
								} else {
									F_errfinish(m, int32(_a_F_check_escape_warning_2), int32(1458), int32(_a_F_check_escape_warning_3))
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return
									} else {
										v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v46 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v45)+56)) = uint8(v46)
										return
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
func F_check_of_type(m *base.Module, l0 int32) {
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
	var v15 int32
	_ = v15
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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
	v11 = v9 + v10
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+79)))
	if v12 == int32(99) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+84))
		v17 = F_relation_open(m, v15, int32(1))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+119)))
			F_relation_close(m, v17, int32(0))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				if v20 != int32(99) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						F_errcode(m, int32(151027844))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							v58 = F_format_type_be(m, v57)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v58
								F_errmsg(m, int32(_a_F_check_of_type_0), v7)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									F_errdetail(m, int32(_a_F_check_of_type_1), int32(0))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_check_of_type_2), int32(_a_F_check_of_type_3), int32(_a_F_check_of_type_4))
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
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
				} else {
					m.G0 = v7 + int32(32)
					return
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			F_errcode(m, int32(151027844))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v37 = F_format_type_be(m, v36)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v37
					F_errmsg(m, int32(_a_F_check_of_type_5), v7+int32(16))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_check_of_type_2), int32(_a_F_check_of_type_6), int32(_a_F_check_of_type_4))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
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
func F_check_serial_buffers(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_check_slru_buffers(m, int32(_a_F_check_serial_buffers_0), l0)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_check_subtrans_buffers(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_check_slru_buffers(m, int32(_a_F_check_subtrans_buffers_0), l0)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_choose_next_subplan_for_leader(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v8 = F_LWLockAcquire(m, v6, int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v12 != int32(-1) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v96 = int32(1)
	v98 = v6 + int32(20)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98+v99))))
	if v101 == v96 {
		goto L31
	} else {
		goto L32
	}
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v17 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15+v12)+20)) = uint8(v17)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v19 - int32(1)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+172)))
	if v23 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v24 = int32(0)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v28 = F_ExecFindMatchingSubPlans(m, v25, v24, v24)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v30 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+172)) = uint8(v30)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v28
	v33 = int32(0)
	if v28 == v33 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v68 == v69 {
		goto L3
	} else {
		goto L22
	}
L10:
	;
	v68 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v40 = int32(1)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v41 <= v40 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v44 = v40
	goto L15
L14:
	;
	v44 = v41
	goto L15
L15:
	;
	v48 = int32(0)
	v50 = v33
	goto L16
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(8)+v48<<(uint(int32(2))%32))))
	if v56 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v68 = v59
	goto L9
L18:
	;
	v59 = v50 + base.I32_popcnt(v56)
	goto L20
L19:
	;
	v59 = v50
	goto L20
L20:
	;
	v61 = v48 + int32(1)
	if v61 != v44 {
		v48 = v61
		v50 = v59
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L17
L22:
	;
	if v69 <= int32(0) {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	v74 = v24
	goto L24
L24:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v79 = F_bms_is_member(m, v74, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	goto L3
L26:
	;
	if v79 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v85 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v83+v74)+20)) = uint8(v85)
	goto L29
L28:
	;
	goto L29
L29:
	;
	v88 = v74 + int32(1)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v88 < v89 {
		v74 = v88
		goto L24
	} else {
		goto L30
	}
L30:
	;
	goto L25
L31:
	;
	v105 = v99
	goto L34
L32:
	;
	v125 = v99
	goto L33
L33:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v125 < v129 {
		goto L41
	} else {
		goto L42
	}
L34:
	;
	if v105 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v125 = v120
	goto L33
L36:
	;
	v111 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v111
	F_LWLockRelease(m, v6)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v120 = v105 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v120
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120+v98))))
	if v123 != 0 {
		v105 = v120
		goto L34
	} else {
		goto L40
	}
L39:
	;
	return int32(0)
L40:
	;
	goto L35
L41:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v133 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v131+v125)+20)) = uint8(v133)
	goto L43
L42:
	;
	goto L43
L43:
	;
	F_LWLockRelease(m, v6)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	return v96
}
