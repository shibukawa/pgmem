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
	v174 = F_query_tree_walker_impl(m, l0, int32(1048), l1, int32(0))
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
	v161 = F_expression_tree_walker_impl(m, l0, int32(1048), l1)
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
func F_CleanUpLock(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v7 == int32(0) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v11
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
		*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v16
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v16))) = v18
		v21 = *(*int32)(unsafe.Add(mBase, _c_F_CleanUpLock[0]))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v28 = F_hash_search_with_hash_value(m, v21, l1, v22<<(uint(int32(4))%32)^l3, int32(2), int32(0))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return
		} else {
			if v28 == int32(0) {
				F_errstart_cold(m, int32(23), int32(0))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_CleanUpLock_0), int32(0))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_CleanUpLock_1), int32(1758), int32(_a_F_CleanUpLock_2))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
				if v33 == int32(0) {
					v37 = *(*int32)(unsafe.Add(mBase, _c_F_CleanUpLock[1]))
					v40 = F_hash_search_with_hash_value(m, v37, l0, l3, int32(2), int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						if v40 != 0 {
							return
						} else {
							F_errstart_cold(m, int32(23), int32(0))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(_a_F_CleanUpLock_3), int32(0))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_CleanUpLock_1), int32(1774), int32(_a_F_CleanUpLock_2))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
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
					if l4 == int32(0) {
						return
					} else {
						F_ProcLockWakeup(m, l2, l0)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	} else {
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
		if v33 == int32(0) {
			v37 = *(*int32)(unsafe.Add(mBase, _c_F_CleanUpLock[1]))
			v40 = F_hash_search_with_hash_value(m, v37, l0, l3, int32(2), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				if v40 != 0 {
					return
				} else {
					F_errstart_cold(m, int32(23), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_CleanUpLock_3), int32(0))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_CleanUpLock_1), int32(1774), int32(_a_F_CleanUpLock_2))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
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
			if l4 == int32(0) {
				return
			} else {
				F_ProcLockWakeup(m, l2, l0)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_CleanupInvalidationState(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupInvalidationState[0]))
	v15 = F_LWLockAcquire(m, v11+int32(768), int32(0))
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
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupInvalidationState[1]))
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupInvalidationState[2]))
	v23 = l1 + v20<<(uint(int32(4))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_CleanupInvalidationState[3]))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_CleanupInvalidationState[4]))) = v18
	v33 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_CleanupInvalidationState[5]))) = uint16(v33)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_CleanupInvalidationState[6])))
	v37 = v35 - int32(1)
	if v37 < v33 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L13
	}
L4:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_CleanupInvalidationState[7])))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+v37<<(uint(int32(2))%32))))
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupInvalidationState[2]))
	if v44 == v46 {
		v73 = v37
		goto L5
	} else {
		goto L6
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_CleanupInvalidationState[6]))) = v73
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_CleanupInvalidationState[0]))
	F_LWLockRelease(m, v82+int32(768))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L12
	}
L6:
	;
	v48 = v37
	goto L7
L7:
	;
	if v48 <= int32(0) {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	if v48 == v35 {
		v73 = v37
		goto L5
	} else {
		goto L11
	}
L9:
	;
	v60 = v48 - int32(1)
	v63 = v40 + v60<<(uint(int32(2))%32)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if v64 != v46 {
		v48 = v60
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v44
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_CleanupInvalidationState[6])))
	v73 = v68 - int32(1)
	goto L5
L12:
	;
	return
L13:
	;
	F_errmsg_internal(m, int32(_a_F_CleanupInvalidationState_0), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_CleanupInvalidationState_1), int32(359), int32(_a_F_CleanupInvalidationState_2))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CloneRowTriggersToPartition(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
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
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
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
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
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
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int64
	_ = v310
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	v15 = m.G0
	v17 = v15 - int32(96)
	m.G0 = v17
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ScanKeyInit(m, v17+int32(48), int32(2), int32(3), int32(184), v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v29 = F_table_open(m, int32(2620), int32(3))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = int32(1)
	v37 = F_systable_beginscan(m, v29, int32(2701), v32, int32(0), v32, v17+int32(48))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_CloneRowTriggersToPartition[0]))
	v45 = F_AllocSetContextCreateInternal(m, v40, int32(_a_F_CloneRowTriggersToPartition_0), int32(0), int32(1024), int32(_a_F_CloneRowTriggersToPartition_1))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v47 = F_systable_getnext(m, v37)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L79
	}
L7:
	;
	if v47 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v57 = v47
	goto L11
L9:
	;
	goto L10
L10:
	;
	F_MemoryContextDelete(m, v45)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L76
	}
L11:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+22)))
	v65 = v63 + v64
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+80)))
	if v66&int32(1) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L10
L13:
	;
	v367 = F_systable_getnext(m, v37)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L74
	}
L14:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+83)))
	if v71 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	switch v66 & int32(66) {
	case 0, 2:
		goto L16
	default:
		goto L17
	}
L16:
	;
	v89 = int32(0)
	v90 = int32(_a_F_CloneRowTriggersToPartition_2)
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_CloneRowTriggersToPartition[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_CloneRowTriggersToPartition[0])) = v45
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	v98 = F_heap_getattr_6(m, v57, int32(17), v95, v17+int32(47))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L21
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v65 + int32(12)
	F_errmsg_internal(m, int32(_a_F_CloneRowTriggersToPartition_3), v17)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_CloneRowTriggersToPartition_4), int32(_a_F_CloneRowTriggersToPartition_5), int32(_a_F_CloneRowTriggersToPartition_6))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L21:
	;
	v100 = int32(0)
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+47)))
	if v101 == v100 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v104 = F_text_to_cstring(m, v98)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	v114 = v100
	goto L24
L24:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v65)+116))
	if int32(0) < v115 {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v106 = F_stringToNode(m, v104)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v109 = F_map_partition_varattnos(m, v106, int32(1), l1, l0)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v112 = F_map_partition_varattnos(m, v109, int32(2), l1, l0)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v114 = v112
	goto L24
L29:
	;
	v123 = int32(0)
	v128 = v89
	goto L32
L30:
	;
	v166 = v89
	goto L31
L31:
	;
	v173 = int32(0)
	v174 = int32(*(*int16)(unsafe.Add(mBase, uint32(v65)+98)))
	if v174 <= v173 {
		v291 = v173
		goto L38
	} else {
		goto L39
	}
L32:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	v143 = int32(*(*int16)(unsafe.Add(mBase, uint32(v65+int32(124)+v123<<(uint(int32(1))%32)))))
	v149 = F_pstrdup(m, v135+v136<<(uint(int32(4))%32)+v143*int32(100)-int32(76))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L34
	}
L33:
	;
	v166 = v153
	goto L31
L34:
	;
	v151 = F_makeString(m, v149)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v153 = F_lappend(m, v128, v151)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v156 = v123 + int32(1)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v65)+116))
	if v156 < v157 {
		v123 = v156
		v128 = v153
		goto L32
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	v300 = F_palloc0(m, int32(52))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L71
	}
L39:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	v181 = F_heap_getattr_6(m, v57, int32(16), v178, v17+int32(47))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+47)))
	if v183 == int32(1) {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	v186 = F_pg_detoast_datum_packed(m, v181)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	v189 = F_pg_detoast_datum_packed(m, v181)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v191 = int32(*(*int16)(unsafe.Add(mBase, uint32(v65)+98)))
	if v191 <= int32(0) {
		v291 = v173
		goto L38
	} else {
		goto L44
	}
L44:
	;
	v194 = int32(1)
	if v188&v194 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v198 = v194
	goto L47
L46:
	;
	v198 = int32(4)
	goto L47
L47:
	;
	v203 = v189 + v198
	v206 = int32(0)
	v207 = v173
	goto L48
L48:
	;
	v215 = F_pstrdup(m, v203)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L50
	}
L49:
	;
	v291 = v219
	goto L38
L50:
	;
	v217 = F_makeString(m, v215)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v219 = F_lappend(m, v207, v217)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	if v203&int32(3) == int32(0) {
		v244 = v203
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v279 = int32(1)
	v282 = v206 + v279
	v283 = int32(*(*int16)(unsafe.Add(mBase, uint32(v65)+98)))
	if v282 < v283 {
		v203 = v277 + v203 + v279
		v206 = v282
		v207 = v219
		goto L48
	} else {
		goto L70
	}
L54:
	;
	v277 = v269 - v203
	goto L53
L55:
	;
	v248 = v244
	goto L64
L56:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	if v228 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v277 = int32(0)
	goto L53
L58:
	;
	goto L59
L59:
	;
	v233 = v203
	goto L60
L60:
	;
	v237 = v233 + int32(1)
	if v237&int32(3) == int32(0) {
		v244 = v237
		goto L55
	} else {
		goto L62
	}
L61:
	;
	v269 = v237
	goto L54
L62:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237))))
	if v242 != 0 {
		v233 = v237
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
	v257 = int32(-2139062144)
	if (int32(16843008)-v254|v254)&v257 == v257 {
		v248 = v248 + int32(4)
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v263 = v248
	goto L67
L66:
	;
	goto L65
L67:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263))))
	if v267 != 0 {
		v263 = v263 + int32(1)
		goto L67
	} else {
		goto L69
	}
L68:
	;
	v269 = v263
	goto L54
L69:
	;
	goto L68
L70:
	;
	goto L49
L71:
	;
	v302 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v300)+4)) = uint8(v302)
	*(*int32)(unsafe.Add(mBase, uint32(v300))) = int32(181)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v65)+92))
	v307 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v300)+24)) = uint8(v307)
	*(*int32)(unsafe.Add(mBase, uint32(v300)+20)) = v291
	v310 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v300)+12)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v300)+8)) = v65 + int32(12)
	*(*uint8)(unsafe.Add(mBase, uint32(v300)+5)) = uint8(base.B2i32(v306 != v302))
	v318 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+80)))
	v320 = v318 & int32(66)
	*(*uint16)(unsafe.Add(mBase, uint32(v300)+26)) = uint16(v320)
	v322 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+80)))
	*(*int64)(unsafe.Add(mBase, uint32(v300)+36)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v300)+32)) = v166
	v327 = v322 & int32(60)
	*(*uint16)(unsafe.Add(mBase, uint32(v300)+28)) = uint16(v327)
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v300)+44)) = uint8(v329)
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+97)))
	*(*int32)(unsafe.Add(mBase, uint32(v300)+48)) = v302
	*(*uint8)(unsafe.Add(mBase, uint32(v300)+45)) = uint8(v331)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v65)+84))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v65)+76))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v346 = int32(*(*int8)(unsafe.Add(mBase, uint32(v65)+82)))
	F_CreateTriggerFiringOn(m, v17+int32(32), v300, v302, v338, v339, v302, v302, v342, v343, v114, v302, v307, v346)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CloneRowTriggersToPartition[0])) = v91
	F_MemoryContextReset(m, v45)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	goto L13
L74:
	;
	if v367 != 0 {
		v57 = v367
		goto L11
	} else {
		goto L75
	}
L75:
	;
	goto L12
L76:
	;
	F_systable_endscan(m, v37)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	F_sequence_close(m, v29, int32(3))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	m.G0 = v17 + int32(96)
	return
L79:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v65 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v397 + int32(4)
	F_errmsg_internal(m, int32(_a_F_CloneRowTriggersToPartition_7), v17+int32(16))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_CloneRowTriggersToPartition_4), int32(_a_F_CloneRowTriggersToPartition_8), int32(_a_F_CloneRowTriggersToPartition_6))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CompareTSQ(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
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
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v6 != v7 {
		if v6 < v7 {
			v12 = int32(-1)
		} else {
			v12 = int32(1)
		}
		return v12
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v15 = int32(2)
		v16 = int32(base.Ui32(v14) >> (uint(v15) % 32))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v19 = int32(base.Ui32(v17) >> (uint(v15) % 32))
		if v16 != v19 {
			if base.Ui32(v16) < base.Ui32(v19) {
				v24 = int32(-1)
			} else {
				v24 = int32(1)
			}
			return v24
		} else {
			if v6 == int32(0) {
				return int32(0)
			} else {
				v31 = l0 + int32(8)
				v35 = F_QT2QTN(m, v31, v31+v6*int32(12))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					v40 = l1 + int32(8)
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v45 = F_QT2QTN(m, v40, v40+v41*int32(12))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						v47 = F_QTNodeCompare(m, v35, v45)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							F_QTNFree(m, v35)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								F_QTNFree(m, v45)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									return v47
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_ConditionVariableInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(-4294967296)
	return
}
func F_ConditionalLockRelation(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	v2 = int32(0)
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+24)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+20)) = v9
	v21 = F_LockAcquireExtended(m, v5+int32(16), int32(8), v2, int32(1), v5+int32(12), v2)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		switch v21 {
		case 0, 3:
			m.G0 = v5 + int32(32)
			return base.B2i32(v21 != int32(0))
		default:
			F_ReceiveSharedInvalidMessages(m)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
				v28 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v27)+53)) = uint8(v28)
				m.G0 = v5 + int32(32)
				return base.B2i32(v21 != int32(0))
			}
		}
	}
}
func F_ConfigurePostmasterWaitSet(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_ConfigurePostmasterWaitSet[0]))
	if v4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_FreeWaitEventSet(m, v4)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v7 = int32(_a_F_ConfigurePostmasterWaitSet_0)
	v8 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ConfigurePostmasterWaitSet[0])) = v8
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_ConfigurePostmasterWaitSet[1]))
	v16 = F_CreateWaitEventSet(m, v8, v13+int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ConfigurePostmasterWaitSet[0])) = v16
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_ConfigurePostmasterWaitSet[2]))
	F_AddWaitEventToSet(m, v16, int32(1), int32(-1), v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_ConfigurePostmasterWaitSet[1]))
	if int32(0) < v26 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v30 = int32(0)
	goto L11
L9:
	;
	goto L10
L10:
	;
	return
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_ConfigurePostmasterWaitSet[0]))
	v33 = int32(2)
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_ConfigurePostmasterWaitSet[3]))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v30<<(uint(v33)%32))))
	F_AddWaitEventToSet(m, v32, v33, v39, int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L13
	}
L12:
	;
	goto L10
L13:
	;
	v44 = v30 + int32(1)
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_ConfigurePostmasterWaitSet[1]))
	if v44 < v46 {
		v30 = v44
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
}
func F_CreateEmptyBlockRefTable(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
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
	v5 = F_palloc(m, int32(8))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_CreateEmptyBlockRefTable[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v10
		v13 = F_MemoryContextAllocZero(m, v10, int32(32))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v10
			v20 = F_MemoryContextAllocExtended(m, v10, int32(_a_F_CreateEmptyBlockRefTable_0), int32(5))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v13)+12)) = int64(31662498914303)
				*(*int64)(unsafe.Add(mBase, uint32(v13))) = int64(8192)
				*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v20
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = v13
				return v5
			}
		}
	}
}
func F___cos(m *base.Module, l0 float64, l1 float64) float64 {
	var v6 float64
	_ = v6
	var v7 float64
	_ = v7
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v22 float64
	_ = v22
	v6 = float64(1)
	v7 = base.F64_mul(l0, l0)
	v9 = base.F64_mul(v7, float64(0.5))
	v10 = base.F64_sub(v6, v9)
	v22 = base.F64_mul(v7, v7)
	return base.F64_add(v10, base.F64_add(base.F64_sub(base.F64_sub(v6, v10), v9), base.F64_sub(base.F64_mul(v7, base.F64_add(base.F64_mul(v7, base.F64_add(base.F64_mul(v7, base.F64_add(base.F64_mul(v7, float64(2.480158728947673e-05)), float64(-0.001388888888887411))), float64(0.0416666666666666))), base.F64_mul(base.F64_mul(v22, v22), base.F64_add(base.F64_mul(v7, base.F64_add(base.F64_mul(v7, float64(-1.1359647557788195e-11)), float64(2.087572321298175e-09))), float64(-2.7557314351390663e-07))))), base.F64_mul(l0, l1))))
}
func F_calc_hist_selectivity_scalar(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64 {
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 float64
	_ = v54
	var v55 float64
	_ = v55
	var v61 int32
	_ = v61
	var v64 float64
	_ = v64
	var v65 int32
	_ = v65
	var v69 float64
	_ = v69
	v14 = l3 - int32(1)
	v20 = int32(-1)
	v21 = v14
	goto L1
L1:
	;
	v30 = base.I32_div_s(v20+v21+int32(1), int32(2))
	v34 = F_range_cmp_bounds(m, l0, l2+v30<<(uint(int32(3))%32), l1)
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v49 = int32(0)
	if v49 < v44 {
		goto L12
	} else {
		goto L13
	}
L3:
	;
	return float64(0)
L4:
	;
	v38 = int32(0)
	v43 = base.B2i32(v34 < v38) | l4&base.B2i32(v34 == v38)
	if v43 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v44 = v30
	goto L7
L6:
	;
	v44 = v20
	goto L7
L7:
	;
	if v43 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v47 = v21
	goto L10
L9:
	;
	v47 = v30 - int32(1)
	goto L10
L10:
	;
	if v44 < v47 {
		v20 = v44
		v21 = v47
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L2
L12:
	;
	v52 = v44
	goto L14
L13:
	;
	v52 = v49
	goto L14
L14:
	;
	v54 = base.F64_convert_i32_u(v14)
	v55 = base.F64_div(base.F64_convert_i32_u(v52), v54)
	if v44 < int32(0) {
		v69 = v55
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return v69
L16:
	;
	if v14 <= v44 {
		v69 = v55
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v61 = l2 + v44<<(uint(int32(3))%32)
	v64 = F_get_position(m, l0, l1, v61, v61+int32(8))
	v65 = m.ExcPending
	if v65 != 0 {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	v69 = base.F64_add(v55, base.F64_div(v64, v54))
	goto L15
}
func F_calc_key_id(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
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
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
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
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int64
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v13 = F_pgp_load_digest(m, int32(2), v8+int32(28))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if int32(0) <= v13 {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
			switch v22 - int32(1) {
			case 0, 1, 2:
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
				v55 = v36 + v38 + int32(10)
			default:
				v55 = int32(6)
			case 15:
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
				v55 = v26 + (v28 + v30) + int32(12)
			case 16:
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
				v55 = v43 + (v45 + (v47 + v49)) + int32(14)
			}
			v56 = int32(153)
			*(*uint8)(unsafe.Add(mBase, uint32(v8)+25)) = uint8(v56)
			*(*uint8)(unsafe.Add(mBase, uint32(v8)+27)) = uint8(v55)
			v59 = int32(8)
			v61 = int32(base.Ui32(v55) >> (uint(v59) % 32))
			*(*uint8)(unsafe.Add(mBase, uint32(v8)+26)) = uint8(v61)
			v63 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
			v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
			m.T0[v67].(func(*base.Module, int32, int32, int32))(m, v63, v8+int32(25), int32(3))
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return int32(0)
			} else {
				v70 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
				v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
				m.T0[v72].(func(*base.Module, int32, int32, int32))(m, v70, l0, int32(1))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int32(0)
				} else {
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
					m.T0[v79].(func(*base.Module, int32, int32, int32))(m, v75, l0+int32(1), int32(4))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
						m.T0[v84].(func(*base.Module, int32, int32, int32))(m, v82, l0+int32(5), int32(1))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int32(0)
						} else {
							v87 = int32(12)
							v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
							switch v88 - int32(1) {
							case 0, 1, 2:
								v106 = v87
								v108 = v59
								v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
								v111 = *(*int32)(unsafe.Add(mBase, uint32(l0+v108)))
								v112 = F_pgp_mpi_hash(m, v109, v111)
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
									return int32(0)
								} else {
									v114 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
									v116 = *(*int32)(unsafe.Add(mBase, uint32(l0+v106)))
									v117 = F_pgp_mpi_hash(m, v114, v116)
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return int32(0)
									} else {
										v122 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
										v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
										m.T0[v123].(func(*base.Module, int32, int32))(m, v122, v8)
										mBase = m.M
										v125 = m.ExcPending
										if v125 != 0 {
											return int32(0)
										} else {
											v126 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
											v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+20))
											m.T0[v127].(func(*base.Module, int32))(m, v126)
											mBase = m.M
											v129 = m.ExcPending
											if v129 != 0 {
												return int32(0)
											} else {
												v130 = *(*int64)(unsafe.Add(mBase, uint32(v8)+12))
												*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v130
												v132 = int32(0)
												v135 = F___memset(m, v8, v132, int32(20))
												mBase = m.M
												v136 = v132
												m.G0 = v8 + int32(32)
												return v136
											}
										}
									}
								}
							default:
								v122 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
								v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
								m.T0[v123].(func(*base.Module, int32, int32))(m, v122, v8)
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return int32(0)
								} else {
									v126 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
									v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+20))
									m.T0[v127].(func(*base.Module, int32))(m, v126)
									mBase = m.M
									v129 = m.ExcPending
									if v129 != 0 {
										return int32(0)
									} else {
										v130 = *(*int64)(unsafe.Add(mBase, uint32(v8)+12))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v130
										v132 = int32(0)
										v135 = F___memset(m, v8, v132, int32(20))
										mBase = m.M
										v136 = v132
										m.G0 = v8 + int32(32)
										return v136
									}
								}
							case 15:
								v98 = v87
								v99 = v88
								v100 = v59
								v101 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
								v103 = *(*int32)(unsafe.Add(mBase, uint32(l0+v100)))
								v104 = F_pgp_mpi_hash(m, v101, v103)
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return int32(0)
								} else {
									v106 = v99
									v108 = v98
									v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
									v111 = *(*int32)(unsafe.Add(mBase, uint32(l0+v108)))
									v112 = F_pgp_mpi_hash(m, v109, v111)
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return int32(0)
									} else {
										v114 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
										v116 = *(*int32)(unsafe.Add(mBase, uint32(l0+v106)))
										v117 = F_pgp_mpi_hash(m, v114, v116)
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int32(0)
										} else {
											v122 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
											v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
											m.T0[v123].(func(*base.Module, int32, int32))(m, v122, v8)
											mBase = m.M
											v125 = m.ExcPending
											if v125 != 0 {
												return int32(0)
											} else {
												v126 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
												v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+20))
												m.T0[v127].(func(*base.Module, int32))(m, v126)
												mBase = m.M
												v129 = m.ExcPending
												if v129 != 0 {
													return int32(0)
												} else {
													v130 = *(*int64)(unsafe.Add(mBase, uint32(v8)+12))
													*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v130
													v132 = int32(0)
													v135 = F___memset(m, v8, v132, int32(20))
													mBase = m.M
													v136 = v132
													m.G0 = v8 + int32(32)
													return v136
												}
											}
										}
									}
								}
							case 16:
								v91 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
								v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v93 = F_pgp_mpi_hash(m, v91, v92)
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return int32(0)
								} else {
									v98 = int32(16)
									v99 = int32(20)
									v100 = int32(12)
									v101 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
									v103 = *(*int32)(unsafe.Add(mBase, uint32(l0+v100)))
									v104 = F_pgp_mpi_hash(m, v101, v103)
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return int32(0)
									} else {
										v106 = v99
										v108 = v98
										v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
										v111 = *(*int32)(unsafe.Add(mBase, uint32(l0+v108)))
										v112 = F_pgp_mpi_hash(m, v109, v111)
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return int32(0)
										} else {
											v114 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
											v116 = *(*int32)(unsafe.Add(mBase, uint32(l0+v106)))
											v117 = F_pgp_mpi_hash(m, v114, v116)
											mBase = m.M
											v118 = m.ExcPending
											if v118 != 0 {
												return int32(0)
											} else {
												v122 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
												v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
												m.T0[v123].(func(*base.Module, int32, int32))(m, v122, v8)
												mBase = m.M
												v125 = m.ExcPending
												if v125 != 0 {
													return int32(0)
												} else {
													v126 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
													v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+20))
													m.T0[v127].(func(*base.Module, int32))(m, v126)
													mBase = m.M
													v129 = m.ExcPending
													if v129 != 0 {
														return int32(0)
													} else {
														v130 = *(*int64)(unsafe.Add(mBase, uint32(v8)+12))
														*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v130
														v132 = int32(0)
														v135 = F___memset(m, v8, v132, int32(20))
														mBase = m.M
														v136 = v132
														m.G0 = v8 + int32(32)
														return v136
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
			v136 = v13
			m.G0 = v8 + int32(32)
			return v136
		}
	}
}
func F_case_index(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	v2 = int32(0)
	if base.Ui32(l0) < base.Ui32(int32(1416)) {
		v89 = l0
		v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
		v96 = v94
	} else {
		if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_0)) {
			if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_1)) {
				if base.Ui32(l0-int32(_a_F_case_index_2)) <= base.Ui32(int32(95)) {
					v89 = l0 - int32(2840)
					v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
					v96 = v94
				} else {
					if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_3)) {
						v96 = v2
					} else {
						if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_4)) {
							v89 = l0 - int32(3512)
							v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
							v96 = v94
						} else {
							if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_5)) {
								v96 = v2
							} else {
								v89 = l0 - int32(_a_F_case_index_6)
								v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
								v96 = v94
							}
						}
					}
				}
			} else {
				if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_7)) {
					v96 = v2
				} else {
					if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_8)) {
						if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_9)) {
							v89 = l0 - int32(_a_F_case_index_10)
							v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
							v96 = v94
						} else {
							if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_11)) {
								v96 = v2
							} else {
								v89 = l0 - int32(_a_F_case_index_12)
								v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
								v96 = v94
							}
						}
					} else {
						if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_13)) {
							v96 = v2
						} else {
							if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_14)) {
								v89 = l0 - int32(_a_F_case_index_15)
								v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
								v96 = v94
							} else {
								if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_16)) {
									v96 = v2
								} else {
									v89 = l0 - int32(_a_F_case_index_17)
									v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
									v96 = v94
								}
							}
						}
					}
				}
			}
		} else {
			if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_18)) {
				v96 = v2
			} else {
				if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_19)) {
					if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_20)) {
						if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_21)) {
							v89 = l0 - int32(_a_F_case_index_22)
							v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
							v96 = v94
						} else {
							if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_23)) {
								v96 = v2
							} else {
								v89 = l0 - int32(_a_F_case_index_24)
								v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
								v96 = v94
							}
						}
					} else {
						if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_25)) {
							v96 = v2
						} else {
							if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_26)) {
								v89 = l0 - int32(_a_F_case_index_27)
								v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
								v96 = v94
							} else {
								if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_28)) {
									v96 = v2
								} else {
									v89 = l0 - int32(_a_F_case_index_29)
									v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
									v96 = v94
								}
							}
						}
					}
				} else {
					if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_30)) {
						v96 = v2
					} else {
						if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_31)) {
							if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_32)) {
								v89 = l0 - int32(_a_F_case_index_33)
								v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
								v96 = v94
							} else {
								if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_34)) {
									v96 = v2
								} else {
									v89 = l0 - int32(_a_F_case_index_35)
									v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
									v96 = v94
								}
							}
						} else {
							if base.Ui32(int32(67)) < base.Ui32(l0-int32(_a_F_case_index_36)) {
								v96 = v2
							} else {
								v89 = l0 - int32(_a_F_case_index_37)
								v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
								v96 = v94
							}
						}
					}
				}
			}
		}
	}
	return v96
}
func F_catalan_ISO_8859_1_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v221 int32
	_ = v221
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
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
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v6
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v18 < v9 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v9
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v239
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v239
	v243 = v239 - int32(1)
	if v243 <= v9 {
		goto L65
	} else {
		goto L66
	}
L2:
	;
	if v58 < int32(0) {
		goto L1
	} else {
		goto L17
	}
L3:
	;
	v20 = v9
	goto L5
L4:
	;
	v20 = v18
	goto L5
L5:
	;
	v27 = v9
	goto L7
L6:
	;
	v58 = v38
	goto L2
L7:
	;
	if v27 == v20 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v58 = int32(-1)
	goto L2
L10:
	;
	goto L11
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v27))))
	if int32(252) < v33 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v50 = v27 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v50
	v27 = v50
	goto L7
L13:
	;
	v35 = v33 - int32(97)
	if v35 < int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v38 = int32(1)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v35)>>(uint(int32(3))%32)))+uint32(_c_F_catalan_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v42)>>(uint(v35&int32(7))%32))&v38 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	goto L12
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v62 = v61 + v58
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v62
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v73 < v62 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v116 < int32(0) {
		goto L1
	} else {
		goto L32
	}
L19:
	;
	v75 = v62
	goto L21
L20:
	;
	v75 = v73
	goto L21
L21:
	;
	v82 = v62
	goto L23
L22:
	;
	v116 = int32(1)
	goto L18
L23:
	;
	if v82 == v75 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v116 = int32(-1)
	goto L18
L26:
	;
	goto L27
L27:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88+v82))))
	if int32(252) < v90 {
		goto L22
	} else {
		goto L28
	}
L28:
	;
	v92 = v90 - int32(97)
	if v92 < int32(0) {
		goto L22
	} else {
		goto L29
	}
L29:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v92)>>(uint(int32(3))%32)))+uint32(_c_F_catalan_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v98)>>(uint(v92&int32(7))%32))&int32(1) == int32(0) {
		goto L22
	} else {
		goto L30
	}
L30:
	;
	v107 = v82 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v107
	v82 = v107
	goto L23
L32:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v120 = v119 + v116
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v120
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v122)+4)) = v120
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v132 < v131 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v172 < int32(0) {
		goto L1
	} else {
		goto L48
	}
L34:
	;
	v134 = v131
	goto L36
L35:
	;
	v134 = v132
	goto L36
L36:
	;
	v141 = v131
	goto L38
L37:
	;
	v172 = v152
	goto L33
L38:
	;
	if v141 == v134 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v172 = int32(-1)
	goto L33
L41:
	;
	goto L42
L42:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145+v141))))
	if int32(252) < v147 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v164 = v141 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v164
	v141 = v164
	goto L38
L44:
	;
	v149 = v147 - int32(97)
	if v149 < int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v152 = int32(1)
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v149)>>(uint(int32(3))%32)))+uint32(_c_F_catalan_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v156)>>(uint(v149&int32(7))%32))&v152 != 0 {
		goto L37
	} else {
		goto L46
	}
L46:
	;
	goto L43
L48:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v176 = v175 + v172
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v176
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v187 < v176 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v230 < int32(0) {
		goto L1
	} else {
		goto L63
	}
L50:
	;
	v189 = v176
	goto L52
L51:
	;
	v189 = v187
	goto L52
L52:
	;
	v196 = v176
	goto L54
L53:
	;
	v230 = int32(1)
	goto L49
L54:
	;
	if v196 == v189 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v230 = int32(-1)
	goto L49
L57:
	;
	goto L58
L58:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202+v196))))
	if int32(252) < v204 {
		goto L53
	} else {
		goto L59
	}
L59:
	;
	v206 = v204 - int32(97)
	if v206 < int32(0) {
		goto L53
	} else {
		goto L60
	}
L60:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v206)>>(uint(int32(3))%32)))+uint32(_c_F_catalan_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v212)>>(uint(v206&int32(7))%32))&int32(1) == int32(0) {
		goto L53
	} else {
		goto L61
	}
L61:
	;
	v221 = v196 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v221
	v196 = v221
	goto L54
L63:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v233))) = v234 + v230
	goto L1
L64:
	;
	return v452
L65:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v276
	v281 = F_find_among_b(m, l0, int32(_a_F_catalan_ISO_8859_1_stem_0), int32(200))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L69
	} else {
		goto L77
	}
L66:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245+v243))))
	if v247&int32(224) != int32(96) {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	if int32(1)<<(uint(v247)%32)&int32(_a_F_catalan_ISO_8859_1_stem_1) == int32(0) {
		goto L65
	} else {
		goto L68
	}
L68:
	;
	v260 = F_find_among_b(m, l0, int32(_a_F_catalan_ISO_8859_1_stem_2), int32(39))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	return int32(0)
L70:
	;
	if v260 == int32(0) {
		goto L65
	} else {
		goto L71
	}
L71:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v266
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	if v266 < v269 {
		goto L65
	} else {
		goto L72
	}
L72:
	;
	v271 = F_slice_del(m, l0)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L69
	} else {
		goto L73
	}
L73:
	;
	if v271 < int32(0) {
		v452 = v271
		goto L64
	} else {
		goto L74
	}
L74:
	;
	goto L65
L75:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v360
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v360
	v365 = F_find_among_b(m, l0, int32(_a_F_catalan_ISO_8859_1_stem_3), int32(22))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L69
	} else {
		goto L110
	}
L76:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v331
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v331
	v336 = F_find_among_b(m, l0, int32(_a_F_catalan_ISO_8859_1_stem_4), int32(283))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L69
	} else {
		goto L99
	}
L77:
	;
	if v281 == int32(0) {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v285
	switch v281 - int32(1) {
	case 0:
		goto L83
	case 1:
		goto L82
	case 2:
		goto L81
	case 3:
		goto L80
	case 4:
		goto L79
	default:
		goto L75
	}
L79:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v321)+4))
	if v285 < v322 {
		goto L76
	} else {
		goto L96
	}
L80:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
	if v285 < v313 {
		goto L76
	} else {
		goto L93
	}
L81:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v303)))
	if v285 < v304 {
		goto L76
	} else {
		goto L90
	}
L82:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v296)))
	if v285 < v297 {
		goto L76
	} else {
		goto L87
	}
L83:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+4))
	if v285 < v290 {
		goto L76
	} else {
		goto L84
	}
L84:
	;
	v292 = F_slice_del(m, l0)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L69
	} else {
		goto L85
	}
L85:
	;
	if int32(0) <= v292 {
		goto L75
	} else {
		goto L86
	}
L86:
	;
	v452 = v292
	goto L64
L87:
	;
	v299 = F_slice_del(m, l0)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L69
	} else {
		goto L88
	}
L88:
	;
	if int32(0) <= v299 {
		goto L75
	} else {
		goto L89
	}
L89:
	;
	v452 = v299
	goto L64
L90:
	;
	v308 = F_slice_from_s(m, l0, int32(3), int32(_a_F_catalan_ISO_8859_1_stem_5))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L69
	} else {
		goto L91
	}
L91:
	;
	if int32(0) <= v308 {
		goto L75
	} else {
		goto L92
	}
L92:
	;
	v452 = v308
	goto L64
L93:
	;
	v317 = F_slice_from_s(m, l0, int32(2), int32(_a_F_catalan_ISO_8859_1_stem_6))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L69
	} else {
		goto L94
	}
L94:
	;
	if int32(0) <= v317 {
		goto L75
	} else {
		goto L95
	}
L95:
	;
	v452 = v317
	goto L64
L96:
	;
	v326 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_ISO_8859_1_stem_7))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L69
	} else {
		goto L97
	}
L97:
	;
	if int32(0) <= v326 {
		goto L75
	} else {
		goto L98
	}
L98:
	;
	v452 = v326
	goto L64
L99:
	;
	if v336 == int32(0) {
		goto L75
	} else {
		goto L100
	}
L100:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v340
	switch v336 - int32(1) {
	case 0:
		goto L102
	case 1:
		goto L101
	default:
		goto L75
	}
L101:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v351)))
	if v340 < v352 {
		goto L75
	} else {
		goto L106
	}
L102:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)+4))
	if v340 < v345 {
		goto L75
	} else {
		goto L103
	}
L103:
	;
	v347 = F_slice_del(m, l0)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L69
	} else {
		goto L104
	}
L104:
	;
	if int32(0) <= v347 {
		goto L75
	} else {
		goto L105
	}
L105:
	;
	v452 = v347
	goto L64
L106:
	;
	v354 = F_slice_del(m, l0)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L69
	} else {
		goto L107
	}
L107:
	;
	if v354 < int32(0) {
		v452 = v354
		goto L64
	} else {
		goto L108
	}
L108:
	;
	goto L75
L109:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v391
	v394 = v391
	goto L120
L110:
	;
	if v365 == int32(0) {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v369
	switch v365 - int32(1) {
	case 0:
		goto L113
	case 1:
		goto L112
	default:
		goto L109
	}
L112:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)+4))
	if v369 < v381 {
		goto L109
	} else {
		goto L117
	}
L113:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
	if v369 < v374 {
		goto L109
	} else {
		goto L114
	}
L114:
	;
	v376 = F_slice_del(m, l0)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L69
	} else {
		goto L115
	}
L115:
	;
	if int32(0) <= v376 {
		goto L109
	} else {
		goto L116
	}
L116:
	;
	v452 = v376
	goto L64
L117:
	;
	v385 = F_slice_from_s(m, l0, int32(2), int32(_a_F_catalan_ISO_8859_1_stem_8))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L69
	} else {
		goto L118
	}
L118:
	;
	if v385 < int32(0) {
		v452 = v385
		goto L64
	} else {
		goto L119
	}
L119:
	;
	goto L109
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v394
	v400 = F_find_among(m, l0, int32(_a_F_catalan_ISO_8859_1_stem_9), int32(13))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L69
	} else {
		goto L123
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v391
	v452 = int32(1)
	goto L64
L122:
	;
	goto L121
L123:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v402
	switch v400 - int32(1) {
	case 0:
		goto L131
	case 1:
		goto L130
	case 2:
		goto L129
	case 3:
		goto L128
	case 4:
		goto L127
	case 5:
		goto L126
	case 6:
		goto L125
	default:
		goto L124
	}
L124:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v394 = v448
	goto L120
L125:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v442 <= v402 {
		goto L122
	} else {
		goto L144
	}
L126:
	;
	v438 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_ISO_8859_1_stem_10))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L69
	} else {
		goto L142
	}
L127:
	;
	v432 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_ISO_8859_1_stem_11))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L69
	} else {
		goto L140
	}
L128:
	;
	v426 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_ISO_8859_1_stem_12))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L69
	} else {
		goto L138
	}
L129:
	;
	v420 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_ISO_8859_1_stem_13))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L69
	} else {
		goto L136
	}
L130:
	;
	v414 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_ISO_8859_1_stem_14))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L69
	} else {
		goto L134
	}
L131:
	;
	v408 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_ISO_8859_1_stem_15))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L69
	} else {
		goto L132
	}
L132:
	;
	if int32(0) <= v408 {
		goto L124
	} else {
		goto L133
	}
L133:
	;
	v452 = v408
	goto L64
L134:
	;
	if int32(0) <= v414 {
		goto L124
	} else {
		goto L135
	}
L135:
	;
	v452 = v414
	goto L64
L136:
	;
	if int32(0) <= v420 {
		goto L124
	} else {
		goto L137
	}
L137:
	;
	v452 = v420
	goto L64
L138:
	;
	if int32(0) <= v426 {
		goto L124
	} else {
		goto L139
	}
L139:
	;
	v452 = v426
	goto L64
L140:
	;
	if int32(0) <= v432 {
		goto L124
	} else {
		goto L141
	}
L141:
	;
	v452 = v432
	goto L64
L142:
	;
	if int32(0) <= v438 {
		goto L124
	} else {
		goto L143
	}
L143:
	;
	v452 = v438
	goto L64
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v402 + int32(1)
	goto L124
}
func F_catalan_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v83 int32
	_ = v83
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v205 int32
	_ = v205
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v330 int32
	_ = v330
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v363 int32
	_ = v363
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v401 int32
	_ = v401
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v452 int32
	_ = v452
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v478 int32
	_ = v478
	var v486 int32
	_ = v486
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v783 int32
	_ = v783
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v7
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = v10
	goto L4
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v10
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v506
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v506
	v510 = v506 - int32(1)
	if v510 <= v10 {
		goto L105
	} else {
		goto L106
	}
L2:
	;
	if v127 < int32(0) {
		goto L1
	} else {
		goto L27
	}
L3:
	;
	v127 = v99
	goto L2
L4:
	;
	if v23 <= v32 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v127 = int32(-1)
	goto L2
L7:
	;
	goto L8
L8:
	;
	v39 = int32(1)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+v24))))
	if base.Ui32(v41) < base.Ui32(int32(192)) {
		v98 = v41
		v99 = v39
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if int32(252) < v98 {
		goto L22
	} else {
		goto L23
	}
L10:
	;
	v45 = v32 + int32(1)
	if v45 == v23 {
		v98 = v41
		v99 = v39
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v24))))
	v50 = v48 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v41) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+v24))))
	v66 = v64 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v41) {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v54 = v32 + int32(2)
	if v54 != v23 {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v98 = v41<<(uint(int32(6))%32)&int32(1984) | v50
	v99 = int32(2)
	goto L9
L16:
	;
	goto L15
L17:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v70))))
	v98 = v83&int32(63) | (v41<<(uint(int32(18))%32)&int32(_a_F_catalan_UTF_8_stem_0) | v50<<(uint(int32(12))%32) | v66<<(uint(int32(6))%32))
	v99 = int32(4)
	goto L9
L18:
	;
	v70 = v32 + int32(3)
	if v70 != v23 {
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v98 = v41<<(uint(int32(12))%32)&int32(_a_F_catalan_UTF_8_stem_1) | v50<<(uint(int32(6))%32) | v66
	v99 = int32(3)
	goto L9
L21:
	;
	goto L20
L22:
	;
	v116 = v99 + v32
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v116
	v32 = v116
	goto L4
L23:
	;
	v103 = v98 - int32(97)
	if v103 < int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v103)>>(uint(int32(3))%32)))+uint32(_c_F_catalan_UTF_8_stem[0]))))
	if int32(base.Ui32(v109)>>(uint(v103&int32(7))%32))&int32(1) != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	goto L22
L27:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v131 = v130 + v127
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v131
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v154 = v131
	goto L30
L28:
	;
	if v250 < int32(0) {
		goto L1
	} else {
		goto L52
	}
L29:
	;
	v250 = v221
	goto L28
L30:
	;
	if v145 <= v154 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v250 = int32(-1)
	goto L28
L33:
	;
	goto L34
L34:
	;
	v161 = int32(1)
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154+v146))))
	if base.Ui32(v163) < base.Ui32(int32(192)) {
		v220 = v163
		v221 = v161
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if int32(252) < v220 {
		goto L29
	} else {
		goto L48
	}
L36:
	;
	v167 = v154 + int32(1)
	if v167 == v145 {
		v220 = v163
		v221 = v161
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167+v146))))
	v172 = v170 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v163) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176+v146))))
	v188 = v186 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v163) {
		goto L44
	} else {
		goto L45
	}
L39:
	;
	v176 = v154 + int32(2)
	if v176 != v145 {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v220 = v163<<(uint(int32(6))%32)&int32(1984) | v172
	v221 = int32(2)
	goto L35
L42:
	;
	goto L41
L43:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146+v192))))
	v220 = v205&int32(63) | (v163<<(uint(int32(18))%32)&int32(_a_F_catalan_UTF_8_stem_0) | v172<<(uint(int32(12))%32) | v188<<(uint(int32(6))%32))
	v221 = int32(4)
	goto L35
L44:
	;
	v192 = v154 + int32(3)
	if v192 != v145 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v220 = v163<<(uint(int32(12))%32)&int32(_a_F_catalan_UTF_8_stem_1) | v172<<(uint(int32(6))%32) | v188
	v221 = int32(3)
	goto L35
L47:
	;
	goto L46
L48:
	;
	v225 = v220 - int32(97)
	if v225 < int32(0) {
		goto L29
	} else {
		goto L49
	}
L49:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v225)>>(uint(int32(3))%32)))+uint32(_c_F_catalan_UTF_8_stem[0]))))
	if int32(base.Ui32(v231)>>(uint(v225&int32(7))%32))&int32(1) == int32(0) {
		goto L29
	} else {
		goto L50
	}
L50:
	;
	v239 = v221 + v154
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v239
	v154 = v239
	goto L30
L52:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v254 = v253 + v250
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v254
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v256)+4)) = v254
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v279 = v269
	goto L55
L53:
	;
	if v374 < int32(0) {
		goto L1
	} else {
		goto L78
	}
L54:
	;
	v374 = v346
	goto L53
L55:
	;
	if v270 <= v279 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v374 = int32(-1)
	goto L53
L58:
	;
	goto L59
L59:
	;
	v286 = int32(1)
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279+v271))))
	if base.Ui32(v288) < base.Ui32(int32(192)) {
		v345 = v288
		v346 = v286
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if int32(252) < v345 {
		goto L73
	} else {
		goto L74
	}
L61:
	;
	v292 = v279 + int32(1)
	if v292 == v270 {
		v345 = v288
		v346 = v286
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292+v271))))
	v297 = v295 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v288) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301+v271))))
	v313 = v311 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v288) {
		goto L69
	} else {
		goto L70
	}
L64:
	;
	v301 = v279 + int32(2)
	if v301 != v270 {
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v345 = v288<<(uint(int32(6))%32)&int32(1984) | v297
	v346 = int32(2)
	goto L60
L67:
	;
	goto L66
L68:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271+v317))))
	v345 = v330&int32(63) | (v288<<(uint(int32(18))%32)&int32(_a_F_catalan_UTF_8_stem_0) | v297<<(uint(int32(12))%32) | v313<<(uint(int32(6))%32))
	v346 = int32(4)
	goto L60
L69:
	;
	v317 = v279 + int32(3)
	if v317 != v270 {
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v345 = v288<<(uint(int32(12))%32)&int32(_a_F_catalan_UTF_8_stem_1) | v297<<(uint(int32(6))%32) | v313
	v346 = int32(3)
	goto L60
L72:
	;
	goto L71
L73:
	;
	v363 = v346 + v279
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v363
	v279 = v363
	goto L55
L74:
	;
	v350 = v345 - int32(97)
	if v350 < int32(0) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v350)>>(uint(int32(3))%32)))+uint32(_c_F_catalan_UTF_8_stem[0]))))
	if int32(base.Ui32(v356)>>(uint(v350&int32(7))%32))&int32(1) != 0 {
		goto L54
	} else {
		goto L76
	}
L76:
	;
	goto L73
L78:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v378 = v377 + v374
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v378
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v401 = v378
	goto L81
L79:
	;
	if v497 < int32(0) {
		goto L1
	} else {
		goto L103
	}
L80:
	;
	v497 = v468
	goto L79
L81:
	;
	if v392 <= v401 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v497 = int32(-1)
	goto L79
L84:
	;
	goto L85
L85:
	;
	v408 = int32(1)
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401+v393))))
	if base.Ui32(v410) < base.Ui32(int32(192)) {
		v467 = v410
		v468 = v408
		goto L86
	} else {
		goto L87
	}
L86:
	;
	if int32(252) < v467 {
		goto L80
	} else {
		goto L99
	}
L87:
	;
	v414 = v401 + int32(1)
	if v414 == v392 {
		v467 = v410
		v468 = v408
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v414+v393))))
	v419 = v417 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v410) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423+v393))))
	v435 = v433 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v410) {
		goto L95
	} else {
		goto L96
	}
L90:
	;
	v423 = v401 + int32(2)
	if v423 != v392 {
		goto L89
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v467 = v410<<(uint(int32(6))%32)&int32(1984) | v419
	v468 = int32(2)
	goto L86
L93:
	;
	goto L92
L94:
	;
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393+v439))))
	v467 = v452&int32(63) | (v410<<(uint(int32(18))%32)&int32(_a_F_catalan_UTF_8_stem_0) | v419<<(uint(int32(12))%32) | v435<<(uint(int32(6))%32))
	v468 = int32(4)
	goto L86
L95:
	;
	v439 = v401 + int32(3)
	if v439 != v392 {
		goto L94
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v467 = v410<<(uint(int32(12))%32)&int32(_a_F_catalan_UTF_8_stem_1) | v419<<(uint(int32(6))%32) | v435
	v468 = int32(3)
	goto L86
L98:
	;
	goto L97
L99:
	;
	v472 = v467 - int32(97)
	if v472 < int32(0) {
		goto L80
	} else {
		goto L100
	}
L100:
	;
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v472)>>(uint(int32(3))%32)))+uint32(_c_F_catalan_UTF_8_stem[0]))))
	if int32(base.Ui32(v478)>>(uint(v472&int32(7))%32))&int32(1) == int32(0) {
		goto L80
	} else {
		goto L101
	}
L101:
	;
	v486 = v468 + v401
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v486
	v401 = v486
	goto L81
L103:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v500))) = v501 + v497
	goto L1
L104:
	;
	return v795
L105:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v543
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v543
	v548 = F_find_among_b(m, l0, int32(_a_F_catalan_UTF_8_stem_2), int32(200))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L109
	} else {
		goto L117
	}
L106:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512+v510))))
	if v514&int32(224) != int32(96) {
		goto L105
	} else {
		goto L107
	}
L107:
	;
	if int32(1)<<(uint(v514)%32)&int32(_a_F_catalan_UTF_8_stem_3) == int32(0) {
		goto L105
	} else {
		goto L108
	}
L108:
	;
	v527 = F_find_among_b(m, l0, int32(_a_F_catalan_UTF_8_stem_4), int32(39))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	return int32(0)
L110:
	;
	if v527 == int32(0) {
		goto L105
	} else {
		goto L111
	}
L111:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v533
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v535)+4))
	if v533 < v536 {
		goto L105
	} else {
		goto L112
	}
L112:
	;
	v538 = F_slice_del(m, l0)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L109
	} else {
		goto L113
	}
L113:
	;
	if v538 < int32(0) {
		v795 = v538
		goto L104
	} else {
		goto L114
	}
L114:
	;
	goto L105
L115:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v627
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v627
	v632 = F_find_among_b(m, l0, int32(_a_F_catalan_UTF_8_stem_5), int32(22))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L109
	} else {
		goto L150
	}
L116:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v598
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v598
	v603 = F_find_among_b(m, l0, int32(_a_F_catalan_UTF_8_stem_6), int32(283))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L109
	} else {
		goto L139
	}
L117:
	;
	if v548 == int32(0) {
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v552
	switch v548 - int32(1) {
	case 0:
		goto L123
	case 1:
		goto L122
	case 2:
		goto L121
	case 3:
		goto L120
	case 4:
		goto L119
	default:
		goto L115
	}
L119:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v588)+4))
	if v552 < v589 {
		goto L116
	} else {
		goto L136
	}
L120:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v579)))
	if v552 < v580 {
		goto L116
	} else {
		goto L133
	}
L121:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v570)))
	if v552 < v571 {
		goto L116
	} else {
		goto L130
	}
L122:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v563)))
	if v552 < v564 {
		goto L116
	} else {
		goto L127
	}
L123:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v556)+4))
	if v552 < v557 {
		goto L116
	} else {
		goto L124
	}
L124:
	;
	v559 = F_slice_del(m, l0)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L109
	} else {
		goto L125
	}
L125:
	;
	if int32(0) <= v559 {
		goto L115
	} else {
		goto L126
	}
L126:
	;
	v795 = v559
	goto L104
L127:
	;
	v566 = F_slice_del(m, l0)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L109
	} else {
		goto L128
	}
L128:
	;
	if int32(0) <= v566 {
		goto L115
	} else {
		goto L129
	}
L129:
	;
	v795 = v566
	goto L104
L130:
	;
	v575 = F_slice_from_s(m, l0, int32(3), int32(_a_F_catalan_UTF_8_stem_7))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L109
	} else {
		goto L131
	}
L131:
	;
	if int32(0) <= v575 {
		goto L115
	} else {
		goto L132
	}
L132:
	;
	v795 = v575
	goto L104
L133:
	;
	v584 = F_slice_from_s(m, l0, int32(2), int32(_a_F_catalan_UTF_8_stem_8))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L109
	} else {
		goto L134
	}
L134:
	;
	if int32(0) <= v584 {
		goto L115
	} else {
		goto L135
	}
L135:
	;
	v795 = v584
	goto L104
L136:
	;
	v593 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_UTF_8_stem_9))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L109
	} else {
		goto L137
	}
L137:
	;
	if int32(0) <= v593 {
		goto L115
	} else {
		goto L138
	}
L138:
	;
	v795 = v593
	goto L104
L139:
	;
	if v603 == int32(0) {
		goto L115
	} else {
		goto L140
	}
L140:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v607
	switch v603 - int32(1) {
	case 0:
		goto L142
	case 1:
		goto L141
	default:
		goto L115
	}
L141:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v618)))
	if v607 < v619 {
		goto L115
	} else {
		goto L146
	}
L142:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v611)+4))
	if v607 < v612 {
		goto L115
	} else {
		goto L143
	}
L143:
	;
	v614 = F_slice_del(m, l0)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L109
	} else {
		goto L144
	}
L144:
	;
	if int32(0) <= v614 {
		goto L115
	} else {
		goto L145
	}
L145:
	;
	v795 = v614
	goto L104
L146:
	;
	v621 = F_slice_del(m, l0)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L109
	} else {
		goto L147
	}
L147:
	;
	if v621 < int32(0) {
		v795 = v621
		goto L104
	} else {
		goto L148
	}
L148:
	;
	goto L115
L149:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v658
	v660 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v662 = v658
	v663 = v660
	goto L160
L150:
	;
	if v632 == int32(0) {
		goto L149
	} else {
		goto L151
	}
L151:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v636
	switch v632 - int32(1) {
	case 0:
		goto L153
	case 1:
		goto L152
	default:
		goto L149
	}
L152:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v647)+4))
	if v636 < v648 {
		goto L149
	} else {
		goto L157
	}
L153:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v640)+4))
	if v636 < v641 {
		goto L149
	} else {
		goto L154
	}
L154:
	;
	v643 = F_slice_del(m, l0)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L109
	} else {
		goto L155
	}
L155:
	;
	if int32(0) <= v643 {
		goto L149
	} else {
		goto L156
	}
L156:
	;
	v795 = v643
	goto L104
L157:
	;
	v652 = F_slice_from_s(m, l0, int32(2), int32(_a_F_catalan_UTF_8_stem_10))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L109
	} else {
		goto L158
	}
L158:
	;
	if v652 < int32(0) {
		v795 = v652
		goto L104
	} else {
		goto L159
	}
L159:
	;
	goto L149
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v662
	v668 = v662 + int32(1)
	if v663 <= v668 {
		goto L166
	} else {
		goto L167
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v658
	v795 = int32(1)
	goto L104
L162:
	;
	goto L161
L163:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v791 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v662 = v791
	v663 = v790
	goto L160
L164:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L192
L165:
	;
	v685 = F_find_among(m, l0, int32(_a_F_catalan_UTF_8_stem_11), int32(13))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L109
	} else {
		goto L170
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v662
	v728 = v662
	v729 = v663
	goto L164
L167:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v670+v668))))
	if v672&int32(224) != int32(160) {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	if int32(1)<<(uint(v672)%32)&int32(344765187) != 0 {
		goto L165
	} else {
		goto L169
	}
L169:
	;
	goto L166
L170:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v687
	switch v685 - int32(1) {
	case 0:
		goto L177
	case 1:
		goto L176
	case 2:
		goto L175
	case 3:
		goto L174
	case 4:
		goto L173
	case 5:
		goto L172
	case 6:
		goto L171
	default:
		goto L163
	}
L171:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v728 = v687
	v729 = v727
	goto L164
L172:
	;
	v723 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_UTF_8_stem_12))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L109
	} else {
		goto L188
	}
L173:
	;
	v717 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_UTF_8_stem_13))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L109
	} else {
		goto L186
	}
L174:
	;
	v711 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_UTF_8_stem_14))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L109
	} else {
		goto L184
	}
L175:
	;
	v705 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_UTF_8_stem_15))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L109
	} else {
		goto L182
	}
L176:
	;
	v699 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_UTF_8_stem_16))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L109
	} else {
		goto L180
	}
L177:
	;
	v693 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_UTF_8_stem_17))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L109
	} else {
		goto L178
	}
L178:
	;
	if int32(0) <= v693 {
		goto L163
	} else {
		goto L179
	}
L179:
	;
	v795 = v693
	goto L104
L180:
	;
	if int32(0) <= v699 {
		goto L163
	} else {
		goto L181
	}
L181:
	;
	v795 = v699
	goto L104
L182:
	;
	if int32(0) <= v705 {
		goto L163
	} else {
		goto L183
	}
L183:
	;
	v795 = v705
	goto L104
L184:
	;
	if int32(0) <= v711 {
		goto L163
	} else {
		goto L185
	}
L185:
	;
	v795 = v711
	goto L104
L186:
	;
	if int32(0) <= v717 {
		goto L163
	} else {
		goto L187
	}
L187:
	;
	v795 = v717
	goto L104
L188:
	;
	if int32(0) <= v723 {
		goto L163
	} else {
		goto L189
	}
L189:
	;
	v795 = v723
	goto L104
L190:
	;
	if v783 < int32(0) {
		goto L162
	} else {
		goto L210
	}
L192:
	;
	goto L193
L193:
	;
	goto L194
L194:
	;
	v738 = v728
	v740 = int32(1)
	goto L197
L196:
	;
	v783 = v768
	goto L190
L197:
	;
	if v729 <= v738 {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	goto L196
L199:
	;
	v783 = int32(-1)
	goto L190
L200:
	;
	goto L201
L201:
	;
	v745 = v738 + int32(1)
	v747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731+v738))))
	if base.Ui32(v747) < base.Ui32(int32(192)) {
		v768 = v745
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v769 = int32(1)
	if v769 < v740 {
		v738 = v768
		v740 = v740 - v769
		goto L197
	} else {
		goto L209
	}
L203:
	;
	if v729 <= v745 {
		v768 = v745
		goto L202
	} else {
		goto L204
	}
L204:
	;
	v754 = v745
	goto L205
L205:
	;
	v757 = int32(*(*int8)(unsafe.Add(mBase, uint32(v731+v754))))
	if int32(-65) < v757 {
		v768 = v754
		goto L202
	} else {
		goto L207
	}
L206:
	;
	v768 = v729
	goto L202
L207:
	;
	v761 = v754 + int32(1)
	if v761 != v729 {
		v754 = v761
		goto L205
	} else {
		goto L208
	}
L208:
	;
	goto L206
L209:
	;
	goto L198
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v783
	goto L163
}
func F_cfb_process(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
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
	var v99 int32
	_ = v99
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	if l2 <= int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v25 = l1
	v26 = l2
	v27 = l3
	v29 = v23
	goto L3
L3:
	;
	if int32(0) < v29 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v63 = l0 + int32(20)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v66 = v25
	v67 = v26
	v68 = v27
	v71 = v64
	goto L21
L5:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v39 = v38 - v29
	if v26 < v39 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	goto L4
L8:
	;
	v41 = v26
	goto L10
L9:
	;
	v41 = v39
	goto L10
L10:
	;
	v42 = m.T0[l4].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v25, v41, v27)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return
L12:
	;
	v44 = v26 - v42
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v45 == v46 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v45 != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v53 = v45
	goto L15
L15:
	;
	if int32(0) < v44 {
		v25 = v25 + v42
		v26 = v44
		v27 = v27 + v42
		v29 = v53
		goto L3
	} else {
		goto L20
	}
L16:
	;
	v50 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v50
	v53 = v50
	goto L15
L17:
	;
	v48 = F__emscripten_memcpy_bulkmem(m, l0+int32(20), l0+int32(84), v45)
	mBase = m.M
	goto L19
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	goto L1
L21:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	v82 = m.T0[v81].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v77, int32(0), v63, v71, l0+int32(52), v15+int32(12))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L11
	} else {
		goto L23
	}
L22:
	;
	goto L1
L23:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v84 <= int32(4) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v84 + int32(1)
	goto L26
L25:
	;
	goto L26
L26:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v67 < v90 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v92 = v67
	goto L29
L28:
	;
	v92 = v90
	goto L29
L29:
	;
	v93 = m.T0[l4].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v66, v92, v68)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L11
	} else {
		goto L30
	}
L30:
	;
	v95 = v67 - v93
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v96 == v97 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if v96 != 0 {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	goto L33
L33:
	;
	if int32(0) < v95 {
		v66 = v66 + v93
		v67 = v95
		v68 = v68 + v93
		v71 = v97
		goto L21
	} else {
		goto L38
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	goto L33
L35:
	;
	v99 = F__emscripten_memcpy_bulkmem(m, v63, l0+int32(84), v96)
	mBase = m.M
	goto L37
L36:
	;
	goto L37
L37:
	;
	goto L34
L38:
	;
	goto L22
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
func F_checkcondition_QueryOperand(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = base.I32_div_s(l1-v7-int32(8), int32(12))
	v15 = v6 + v12*int32(_a_F_checkcondition_QueryOperand_0)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v16 != int32(1) {
		v41 = int32(0)
	} else {
		if l2 == int32(0) {
			v41 = int32(1)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			v24 = v15 + int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v24
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v22
			v27 = int32(1)
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
			if v28 != v27 {
				v41 = v27
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
				v34 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v24 + (int32(_a_F_checkcondition_QueryOperand_1)-v32)<<(uint(v34)%32)
				v41 = v34
			}
		}
	}
	return v41
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
func F_cidr_abbrev(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
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
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
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
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
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
	var v475 int32
	_ = v475
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v585 int32
	_ = v585
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 + int32(-64)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = F_pg_detoast_datum_packed(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = int32(1)
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v26&v24 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v29 = v24
	goto L5
L4:
	;
	v29 = int32(4)
	goto L5
L5:
	;
	v30 = v20 + v29
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	v32 = int32(2)
	v33 = v30 + v32
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
	v35 = int32(50)
	v36 = m.G0
	v38 = v36 - int32(192)
	m.G0 = v38
	switch v31 - v32 {
	case 0:
		goto L10
	case 1:
		goto L9
	default:
		goto L8
	}
L6:
	;
	m.G0 = v38 + int32(192)
	if v585 == int32(0) {
		goto L158
	} else {
		goto L159
	}
L7:
	;
	v585 = int32(0)
	goto L6
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_cidr_abbrev[0])) = int32(5)
	goto L7
L9:
	;
	if base.Ui32(int32(129)) <= base.Ui32(v34) {
		goto L39
	} else {
		goto L40
	}
L10:
	;
	if base.Ui32(int32(33)) <= base.Ui32(v34) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_cidr_abbrev[0])) = int32(28)
	goto L7
L12:
	;
	goto L13
L13:
	;
	if v34 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_cidr_abbrev[0])) = int32(35)
	goto L7
L15:
	;
	if base.Ui32(v149) < base.Ui32(int32(5)) {
		goto L14
	} else {
		goto L37
	}
L16:
	;
	v49 = int32(48)
	*(*uint16)(unsafe.Add(mBase, uint32(v17))) = uint16(v49)
	v146 = v15 + int32(-63)
	v149 = int32(49)
	goto L15
L17:
	;
	goto L18
L18:
	;
	if base.Ui32(v34) < base.Ui32(int32(8)) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v113 = v34 & int32(7)
	if v113 == int32(0) {
		v146 = v102
		v149 = v105
		goto L15
	} else {
		goto L31
	}
L20:
	;
	v99 = v33
	v102 = v17
	v105 = v35
	goto L19
L21:
	;
	goto L22
L22:
	;
	v59 = v33
	v61 = v17
	v63 = int32(base.Ui32(v34) >> (uint(int32(3)) % 32))
	v65 = v35
	goto L23
L23:
	;
	if base.Ui32(v65) < base.Ui32(int32(6)) {
		goto L14
	} else {
		goto L25
	}
L24:
	;
	v99 = v77
	v102 = v92
	v105 = v93
	goto L19
L25:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+32)) = v74
	v77 = v59 + int32(1)
	v81 = F_pg_sprintf(m, v61, int32(_a_F_cidr_abbrev_0), v38+int32(32))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v83 = v81 + v61
	if v63 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v99 = v77
	v102 = v83
	v105 = v61 + v65 - v83
	goto L19
L28:
	;
	goto L29
L29:
	;
	v88 = int32(46)
	*(*uint16)(unsafe.Add(mBase, uint32(v83))) = uint16(v88)
	v91 = int32(1)
	v92 = v83 + v91
	v93 = v61 + v65 - v92
	if v91 < v63 {
		v59 = v77
		v61 = v92
		v63 = v63 - v91
		v65 = v93
		goto L23
	} else {
		goto L30
	}
L30:
	;
	goto L24
L31:
	;
	if base.Ui32(v105) < base.Ui32(int32(6)) {
		goto L14
	} else {
		goto L32
	}
L32:
	;
	if v17 != v102 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v119 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v102))) = uint8(v119)
	v123 = v102 + int32(1)
	goto L35
L34:
	;
	v123 = v17
	goto L35
L35:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	v125 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v124 & ((v125<<(uint(v113)%32) ^ v125) << (uint(int32(8)-v113) % 32))
	v138 = F_pg_sprintf(m, v123, int32(_a_F_cidr_abbrev_0), v38+int32(16))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v140 = v138 + v123
	v146 = v140
	v149 = v102 + v105 - v140
	goto L15
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v34
	v160 = F_pg_sprintf(m, v146, int32(_a_F_cidr_abbrev_1), v38)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v585 = v17
	goto L6
L39:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_cidr_abbrev[0])) = int32(28)
	goto L7
L40:
	;
	goto L41
L41:
	;
	if v34 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v34
	v423 = F_pg_sprintf(m, v406, int32(_a_F_cidr_abbrev_1), v38+int32(48))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L116
	}
L43:
	;
	v186 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+114)) = uint8(v186)
	v188 = int32(_a_F_cidr_abbrev_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v38)+112)) = uint16(v188)
	v406 = v38 + int32(112) | int32(2)
	goto L42
L44:
	;
	goto L45
L45:
	;
	v199 = int32(base.Ui32(v34+int32(7)) >> (uint(int32(3)) % 32))
	if v199 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v202 = int32(0)
	v210 = F__emscripten_memset_bulkmem(m, v38+int32(176)+v199, base.I32_extend8_s(v202), int32(16)-v199)
	mBase = m.M
	goto L50
L47:
	;
	v200 = F__emscripten_memcpy_bulkmem(m, v38+int32(176), v33, v199)
	mBase = m.M
	goto L49
L48:
	;
	goto L49
L49:
	;
	goto L46
L50:
	;
	v212 = v34 & int32(7)
	if v212 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v215 = v38 + v199 + int32(175)
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	v221 = v216 & (int32(-1) << (uint(int32(8)-v212) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v215))) = uint8(v221)
	goto L53
L52:
	;
	goto L53
L53:
	;
	v228 = int32(base.Ui32(v34+int32(15)) >> (uint(int32(4)) % 32))
	if v228 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v231 = int32(2)
	goto L56
L55:
	;
	v231 = v228
	goto L56
L56:
	;
	v236 = int32(0)
	v239 = v202
	v243 = v2
	v244 = v2
	v245 = v2
	goto L57
L57:
	;
	v251 = v38 + int32(176) + v239
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+1)))
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
	if v252|v253 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v273 = int32(0)
	v277 = base.B2i32(v266 != v273) & base.B2i32(v268 < v266)
	if v277 != 0 {
		goto L71
	} else {
		goto L72
	}
L59:
	;
	v271 = v239 + int32(2)
	if base.Ui32(v271) < base.Ui32(v231<<(uint(int32(1))%32)) {
		v236 = v266
		v239 = v271
		v243 = v267
		v244 = v268
		v245 = v269
		goto L57
	} else {
		goto L70
	}
L60:
	;
	if v236 != 0 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	if v236 != 0 {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	v259 = v243
	goto L65
L64:
	;
	v259 = int32(base.Ui32(v239) >> (uint(int32(1)) % 32))
	goto L65
L65:
	;
	v266 = v236 + int32(1)
	v267 = v259
	v268 = v244
	v269 = v245
	goto L59
L66:
	;
	if v236 <= v244 {
		v266 = v236
		v267 = v243
		v268 = v244
		v269 = v245
		goto L59
	} else {
		goto L69
	}
L67:
	;
	v263 = v244
	v264 = v245
	goto L68
L68:
	;
	v266 = int32(0)
	v267 = v243
	v268 = v263
	v269 = v264
	goto L59
L69:
	;
	v263 = v236
	v264 = v243
	goto L68
L70:
	;
	goto L58
L71:
	;
	v278 = v267
	goto L73
L72:
	;
	v278 = v269
	goto L73
L73:
	;
	v279 = int32(0)
	if v277 != 0 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v308 = v38 + int32(112)
	v310 = v38 + int32(176)
	v311 = v273
	goto L87
L75:
	;
	v280 = v266
	goto L77
L76:
	;
	v280 = v268
	goto L77
L77:
	;
	if v280 == v231 {
		v297 = v279
		goto L74
	} else {
		goto L78
	}
L78:
	;
	if v278 != 0 {
		v297 = v279
		goto L74
	} else {
		goto L79
	}
L79:
	;
	switch v280 - int32(5) {
	case 0:
		goto L82
	case 1:
		goto L80
	case 2:
		goto L81
	default:
		v297 = v279
		goto L74
	}
L80:
	;
	v297 = int32(1)
	goto L74
L81:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+190)))
	if v290 == int32(0) {
		v297 = v279
		goto L74
	} else {
		goto L85
	}
L82:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+186)))
	if v284 != int32(255) {
		v297 = v279
		goto L74
	} else {
		goto L83
	}
L83:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+187)))
	if v287 == int32(255) {
		goto L80
	} else {
		goto L84
	}
L84:
	;
	v297 = v279
	goto L74
L85:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+191)))
	if v293 == int32(1) {
		v297 = v279
		goto L74
	} else {
		goto L86
	}
L86:
	;
	goto L80
L87:
	;
	if v280 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	v406 = v400
	goto L42
L89:
	;
	v403 = v311 + int32(1)
	if v403 != v231 {
		v308 = v400
		v310 = v401
		v311 = v403
		goto L87
	} else {
		goto L115
	}
L90:
	;
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v311))&v297 != 0 {
		goto L100
	} else {
		goto L101
	}
L91:
	;
	if v311 < v278 {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	if v278+v280 <= v311 {
		goto L90
	} else {
		goto L93
	}
L93:
	;
	if v311 == v278 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v326 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v308))) = uint8(v326)
	v330 = v308 + int32(1)
	goto L96
L95:
	;
	v330 = v308
	goto L96
L96:
	;
	if v311 == v231-int32(1) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v332 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v330))) = uint8(v332)
	v336 = v330 + int32(1)
	goto L99
L98:
	;
	v336 = v330
	goto L99
L99:
	;
	v400 = v336
	v401 = v310 + int32(2)
	goto L89
L100:
	;
	if v311 == int32(6) {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	goto L102
L102:
	;
	if v38+int32(112) == v308 {
		goto L111
	} else {
		goto L112
	}
L103:
	;
	v346 = int32(58)
	goto L105
L104:
	;
	v346 = int32(46)
	goto L105
L105:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v308))) = uint8(v346)
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v348
	v351 = v308 + int32(1)
	v355 = F_pg_sprintf(m, v351, int32(_a_F_cidr_abbrev_0), v38+int32(80))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v357 = v355 + v351
	if base.Ui32(int32(120)) < base.Ui32(v34) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v362 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v357))) = uint8(v362)
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v364
	v367 = v357 + int32(1)
	v371 = F_pg_sprintf(m, v367, int32(_a_F_cidr_abbrev_0), v38-int32(-64))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L110
	}
L108:
	;
	if v311 != int32(7) {
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v400 = v357
	v401 = v310 + int32(1)
	goto L89
L110:
	;
	v400 = v371 + v367
	v401 = v310 + int32(2)
	goto L89
L111:
	;
	v385 = v38 + int32(112)
	goto L113
L112:
	;
	v381 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v308))) = uint8(v381)
	v385 = v308 + int32(1)
	goto L113
L113:
	;
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v386 | v387<<(uint(int32(8))%32)
	v397 = F_pg_sprintf(m, v385, int32(_a_F_cidr_abbrev_3), v38+int32(96))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v400 = v397 + v385
	v401 = v310 + int32(2)
	goto L89
L115:
	;
	goto L88
L116:
	;
	v426 = v38 + int32(112)
	if v426&int32(3) == int32(0) {
		v450 = v426
		goto L119
	} else {
		goto L120
	}
L117:
	;
	if base.Ui32(v483+int32(1)) <= base.Ui32(int32(50)) {
		goto L134
	} else {
		goto L135
	}
L118:
	;
	v483 = v475 - v426
	goto L117
L119:
	;
	v454 = v450
	goto L128
L120:
	;
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426))))
	if v434 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v483 = int32(0)
	goto L117
L122:
	;
	goto L123
L123:
	;
	v439 = v426
	goto L124
L124:
	;
	v443 = v439 + int32(1)
	if v443&int32(3) == int32(0) {
		v450 = v443
		goto L119
	} else {
		goto L126
	}
L125:
	;
	v475 = v443
	goto L118
L126:
	;
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	if v448 != 0 {
		v439 = v443
		goto L124
	} else {
		goto L127
	}
L127:
	;
	goto L125
L128:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v454)))
	v463 = int32(-2139062144)
	if (int32(16843008)-v460|v460)&v463 == v463 {
		v454 = v454 + int32(4)
		goto L128
	} else {
		goto L130
	}
L129:
	;
	v469 = v454
	goto L131
L130:
	;
	goto L129
L131:
	;
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469))))
	if v473 != 0 {
		v469 = v469 + int32(1)
		goto L131
	} else {
		goto L133
	}
L132:
	;
	v475 = v469
	goto L118
L133:
	;
	goto L132
L134:
	;
	v489 = v38 + int32(112)
	if (v489^v17)&int32(3) != 0 {
		goto L140
	} else {
		goto L141
	}
L135:
	;
	goto L136
L136:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_cidr_abbrev[0])) = int32(35)
	goto L7
L137:
	;
	v585 = v17
	goto L6
L138:
	;
	goto L137
L139:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v544))) = uint8(v543)
	if v543&int32(255) == int32(0) {
		goto L138
	} else {
		goto L154
	}
L140:
	;
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489))))
	v542 = v489
	v543 = v495
	v544 = v17
	goto L139
L141:
	;
	goto L142
L142:
	;
	if v489&int32(3) != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v499 = v489
	v501 = v17
	goto L146
L144:
	;
	v513 = v489
	v515 = v17
	goto L145
L145:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v513)))
	v520 = int32(-2139062144)
	if (int32(16843008)-v517|v517)&v520 != v520 {
		v542 = v513
		v543 = v517
		v544 = v515
		goto L139
	} else {
		goto L150
	}
L146:
	;
	v502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499))))
	*(*uint8)(unsafe.Add(mBase, uint32(v501))) = uint8(v502)
	if v502 == int32(0) {
		goto L138
	} else {
		goto L148
	}
L147:
	;
	v513 = v509
	v515 = v507
	goto L145
L148:
	;
	v506 = int32(1)
	v507 = v501 + v506
	v509 = v499 + v506
	if v509&int32(3) != 0 {
		v499 = v509
		v501 = v507
		goto L146
	} else {
		goto L149
	}
L149:
	;
	goto L147
L150:
	;
	v525 = v513
	v526 = v517
	v527 = v515
	goto L151
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v527))) = v526
	v529 = int32(4)
	v530 = v527 + v529
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v525)+4))
	v533 = v525 + v529
	v537 = int32(-2139062144)
	if (v531|(int32(16843008)-v531))&v537 == v537 {
		v525 = v533
		v526 = v531
		v527 = v530
		goto L151
	} else {
		goto L153
	}
L152:
	;
	v542 = v533
	v543 = v531
	v544 = v530
	goto L139
L153:
	;
	goto L152
L154:
	;
	v551 = v542
	v553 = v544
	goto L155
L155:
	;
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v553)+1)) = uint8(v554)
	v556 = int32(1)
	if v554 != 0 {
		v551 = v551 + v556
		v553 = v553 + v556
		goto L155
	} else {
		goto L157
	}
L156:
	;
	goto L138
L157:
	;
	goto L156
L158:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L1
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	v620 = F_cstring_to_text(m, v17)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L165
	}
L161:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	F_errmsg(m, int32(_a_F_cidr_abbrev_4), int32(0))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(_a_F_cidr_abbrev_5), int32(1217), int32(_a_F_cidr_abbrev_6))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L165:
	;
	m.G0 = v17 - int32(-64)
	return v620
}
func F_cidr_out(m *base.Module, l0 int32) int32 {
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
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = F_network_out(m, v3, int32(1))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_cidr_send(m *base.Module, l0 int32) int32 {
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
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = F_network_send(m, v3, int32(1))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_cidr_set_masklen(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v17 == int32(-1) {
			v22 = int32(1)
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			v26 = v24 & v22
			if v26 != 0 {
				v27 = v22
			} else {
				v27 = int32(4)
			}
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v27))))
			if v29 == int32(2) {
				v32 = int32(32)
			} else {
				v32 = int32(128)
			}
			v38 = v26
			v39 = v32
			if v38 != 0 {
				v44 = int32(1)
			} else {
				v44 = int32(4)
			}
			v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v44))))
			if v46 == int32(2) {
				v49 = int32(32)
			} else {
				v49 = int32(128)
			}
			if base.Ui32(v49) < base.Ui32(v39) {
				v139 = v39
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v143 = m.ExcPending
				if v143 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v146 = m.ExcPending
					if v146 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v139
						F_errmsg(m, int32(_a_F_cidr_set_masklen_0), v10)
						mBase = m.M
						v150 = m.ExcPending
						if v150 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_cidr_set_masklen_1), int32(357), int32(_a_F_cidr_set_masklen_2))
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
				v52 = F_palloc0(m, int32(22))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					v54 = int32(1)
					v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
					v58 = v56 & v54
					if v58 != 0 {
						v59 = v54
					} else {
						v59 = int32(4)
					}
					v61 = int32(1)
					v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
					if v63&v61 != 0 {
						v66 = v61
					} else {
						v66 = int32(4)
					}
					v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v66))))
					*(*uint8)(unsafe.Add(mBase, uint32(v52+v59))) = uint8(v68)
					v71 = v52 + int32(1)
					v73 = v52 + int32(4)
					if v58 != 0 {
						v74 = v71
					} else {
						v74 = v73
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)) = uint8(v39)
					if v39 == int32(0) {
					} else {
						v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
						if v78&int32(1) != 0 {
							v81 = v71
						} else {
							v81 = v73
						}
						v84 = int32(1)
						v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
						if v88&v84 != 0 {
							v91 = v13 + v84
						} else {
							v91 = v13 + int32(4)
						}
						v97 = int32(base.Ui32(v39+int32(7)) >> (uint(int32(3)) % 32))
						if v97 != 0 {
							v98 = F__emscripten_memcpy_bulkmem(m, v81+int32(2), v91+int32(2), v97)
							mBase = m.M
						} else {
						}
						v101 = v39 & int32(7)
						if v101 == int32(0) {
						} else {
							v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
							if v106&int32(1) != 0 {
								v109 = v71
							} else {
								v109 = v73
							}
							v112 = int32(base.Ui32(v39)>>(uint(int32(3))%32)) + v109 + int32(2)
							v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
							v116 = v113 & (int32(-256) >> (uint(v101) % 32))
							*(*uint8)(unsafe.Add(mBase, uint32(v112))) = uint8(v116)
						}
					}
					v122 = int32(1)
					v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
					if v124&v122 != 0 {
						v127 = v122
					} else {
						v127 = int32(4)
					}
					v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+v127))))
					if v129 == int32(2) {
						v132 = int32(40)
					} else {
						v132 = int32(88)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v52))) = v132
					m.G0 = v10 + int32(16)
					return v52
				}
			}
		} else {
			if v17 < int32(0) {
				v139 = v17
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v143 = m.ExcPending
				if v143 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v146 = m.ExcPending
					if v146 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v139
						F_errmsg(m, int32(_a_F_cidr_set_masklen_0), v10)
						mBase = m.M
						v150 = m.ExcPending
						if v150 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_cidr_set_masklen_1), int32(357), int32(_a_F_cidr_set_masklen_2))
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
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
				v38 = v35 & int32(1)
				v39 = v17
				if v38 != 0 {
					v44 = int32(1)
				} else {
					v44 = int32(4)
				}
				v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v44))))
				if v46 == int32(2) {
					v49 = int32(32)
				} else {
					v49 = int32(128)
				}
				if base.Ui32(v49) < base.Ui32(v39) {
					v139 = v39
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v143 = m.ExcPending
					if v143 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v146 = m.ExcPending
						if v146 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v139
							F_errmsg(m, int32(_a_F_cidr_set_masklen_0), v10)
							mBase = m.M
							v150 = m.ExcPending
							if v150 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_cidr_set_masklen_1), int32(357), int32(_a_F_cidr_set_masklen_2))
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
					v52 = F_palloc0(m, int32(22))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						v54 = int32(1)
						v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
						v58 = v56 & v54
						if v58 != 0 {
							v59 = v54
						} else {
							v59 = int32(4)
						}
						v61 = int32(1)
						v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
						if v63&v61 != 0 {
							v66 = v61
						} else {
							v66 = int32(4)
						}
						v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v66))))
						*(*uint8)(unsafe.Add(mBase, uint32(v52+v59))) = uint8(v68)
						v71 = v52 + int32(1)
						v73 = v52 + int32(4)
						if v58 != 0 {
							v74 = v71
						} else {
							v74 = v73
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)) = uint8(v39)
						if v39 == int32(0) {
						} else {
							v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
							if v78&int32(1) != 0 {
								v81 = v71
							} else {
								v81 = v73
							}
							v84 = int32(1)
							v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
							if v88&v84 != 0 {
								v91 = v13 + v84
							} else {
								v91 = v13 + int32(4)
							}
							v97 = int32(base.Ui32(v39+int32(7)) >> (uint(int32(3)) % 32))
							if v97 != 0 {
								v98 = F__emscripten_memcpy_bulkmem(m, v81+int32(2), v91+int32(2), v97)
								mBase = m.M
							} else {
							}
							v101 = v39 & int32(7)
							if v101 == int32(0) {
							} else {
								v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
								if v106&int32(1) != 0 {
									v109 = v71
								} else {
									v109 = v73
								}
								v112 = int32(base.Ui32(v39)>>(uint(int32(3))%32)) + v109 + int32(2)
								v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
								v116 = v113 & (int32(-256) >> (uint(v101) % 32))
								*(*uint8)(unsafe.Add(mBase, uint32(v112))) = uint8(v116)
							}
						}
						v122 = int32(1)
						v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
						if v124&v122 != 0 {
							v127 = v122
						} else {
							v127 = int32(4)
						}
						v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+v127))))
						if v129 == int32(2) {
							v132 = int32(40)
						} else {
							v132 = int32(88)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v52))) = v132
						m.G0 = v10 + int32(16)
						return v52
					}
				}
			}
		}
	}
}
func F_clamp_cardinality_to_long(m *base.Module, l0 float64) int32 {
	var v13 float64
	_ = v13
	var v16 float64
	_ = v16
	var v20 int32
	_ = v20
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(l0)&int64(9223372036854775807)) {
		return int32(2147483647)
	} else {
		if base.F64_le(l0, float64(0)) != 0 {
			return int32(0)
		} else {
			v13 = float64(2.147483647e+09)
			if base.F64_lt(l0, v13) != 0 {
				v16 = l0
			} else {
				v16 = v13
			}
			if base.F64_lt(base.F64_abs(v16), float64(2.147483648e+09)) != 0 {
				v20 = base.I32_trunc_f64_s(v16)
				return v20
			} else {
				return int32(-2147483648)
			}
		}
	}
}
func F_clamp_row_est(m *base.Module, l0 float64) float64 {
	var v3 float64
	_ = v3
	var v11 float64
	_ = v11
	var v15 float64
	_ = v15
	v3 = float64(1e+100)
	if base.F64_gt(l0, v3) != 0 {
		v15 = v3
	} else {
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(l0)&int64(9223372036854775807)) {
			v15 = v3
		} else {
			v11 = float64(1)
			if base.F64_le(l0, v11) != 0 {
				v15 = v11
			} else {
				v15 = base.F64_nearest(l0)
			}
		}
	}
	return v15
}
func F_clamp_width_est(m *base.Module, l0 int64) int32 {
	var v2 int64
	_ = v2
	var v5 int64
	_ = v5
	v2 = int64(1073741823)
	if v2 <= l0 {
		v5 = v2
	} else {
		v5 = l0
	}
	return base.I32_wrap_i64(v5)
}
func F_clauselist_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64 {
	var v7 float64
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_clauselist_selectivity_ext(m, l0, l1, l2, l3, l4, int32(1))
	v10 = m.ExcPending
	if v10 != 0 {
		return float64(0)
	} else {
		return v7
	}
}
func F_clauselist_selectivity_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) float64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 float64
	_ = v33
	var v36 int32
	_ = v36
	var v37 float64
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 float64
	_ = v51
	var v52 int32
	_ = v52
	var v53 float64
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v75 float64
	_ = v75
	var v78 int32
	_ = v78
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
	var v88 float64
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
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
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
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
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
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
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v173 float64
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v192 float64
	_ = v192
	var v195 int32
	_ = v195
	var v205 int32
	_ = v205
	var v212 float64
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 float64
	_ = v221
	var v222 float64
	_ = v222
	var v225 float64
	_ = v225
	var v232 int32
	_ = v232
	var v233 float64
	_ = v233
	var v234 int32
	_ = v234
	var v235 float64
	_ = v235
	var v244 float64
	_ = v244
	var v245 float64
	_ = v245
	var v246 float64
	_ = v246
	var v247 float64
	_ = v247
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 float64
	_ = v253
	var v268 float64
	_ = v268
	v7 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v7
	if l1 == v7 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v20 + int32(16)
	return v268
L2:
	;
	v37 = float64(1)
	v38 = F_find_single_rel_for_clauses(m, l0, l1)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L7
	}
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v28 != int32(1) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v33 = F_clause_selectivity_ext(m, l0, v32, l2, l3, l4, l5)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return float64(0)
L6:
	;
	v268 = v33
	goto L1
L7:
	;
	if l5 == int32(0) {
		v53 = v37
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if l1 == int32(0) {
		v268 = v53
		goto L1
	} else {
		goto L14
	}
L9:
	;
	if v38 == int32(0) {
		v53 = v37
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)+76))
	if v44 != 0 {
		v53 = v37
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v38)+112))
	if v45 == int32(0) {
		v53 = v37
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v51 = F_statext_clauselist_selectivity(m, l0, l1, l2, l3, l4, v38, v20+int32(12), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v53 = v51
	goto L8
L14:
	;
	v56 = int32(0)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v56 < v57 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v67 = v56
	v73 = int32(-1)
	v75 = v53
	goto L18
L16:
	;
	v192 = v53
	goto L17
L17:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v195 == int32(0) {
		v268 = v192
		goto L1
	} else {
		goto L57
	}
L18:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78+v67<<(uint(int32(2))%32))))
	v84 = v73 + int32(1)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v86 = F_bms_is_member(m, v84, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L21
	}
L19:
	;
	v192 = v173
	goto L17
L20:
	;
	v175 = v67 + int32(1)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v175 < v176 {
		v67 = v175
		v73 = v84
		v75 = v173
		goto L18
	} else {
		goto L56
	}
L21:
	;
	if v86 != 0 {
		v173 = v75
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v88 = F_clause_selectivity_ext(m, l0, v82, l2, l3, l4, l5)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v91 != int32(318) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v173 = base.F64_mul(v75, v88)
	goto L20
L25:
	;
	if v103 != int32(17) {
		goto L24
	} else {
		goto L31
	}
L26:
	;
	v101 = v82
	v102 = int32(0)
	v103 = v91
	goto L25
L27:
	;
	goto L28
L28:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+10)))
	if v94 == int32(1) {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v97 == int32(0) {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v101 = v97
	v102 = v82
	v103 = v100
	goto L25
L31:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v101)+28))
	if v106 == int32(0) {
		goto L24
	} else {
		goto L32
	}
L32:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	if v109 != int32(2) {
		goto L24
	} else {
		goto L33
	}
L33:
	;
	if v102 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v147 = F_get_oprrest(m, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L5
	} else {
		goto L53
	}
L35:
	;
	v145 = int32(0)
	goto L34
L36:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v102)+24))
	if v112 != int32(1) {
		goto L24
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v127 = F_NumRelids(m, l0, v101)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L5
	} else {
		goto L44
	}
L39:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v102)+48))
	v119 = F_is_pseudo_constant_clause_relids(m, v117, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	if v119 != 0 {
		v145 = int32(1)
		goto L34
	} else {
		goto L41
	}
L41:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v101)+28))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+12))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v102)+44))
	v125 = F_is_pseudo_constant_clause_relids(m, v123, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	if v125 != 0 {
		goto L35
	} else {
		goto L43
	}
L43:
	;
	goto L24
L44:
	;
	if v127 != int32(1) {
		goto L24
	} else {
		goto L45
	}
L45:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v101)+28))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+12))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	v135 = F_is_pseudo_constant_clause(m, v134)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	if v135 != 0 {
		v145 = int32(1)
		goto L34
	} else {
		goto L47
	}
L47:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v101)+28))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v140 = F_is_pseudo_constant_clause(m, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	if v140 == int32(0) {
		goto L24
	} else {
		goto L49
	}
L49:
	;
	goto L35
L50:
	;
	F_addRangeClause(m, v20+int32(8), v101, v145, int32(0), v88)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L5
	} else {
		goto L55
	}
L51:
	;
	F_addRangeClause(m, v20+int32(8), v101, v145, int32(1), v88)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L5
	} else {
		goto L54
	}
L52:
	;
	switch v147 - int32(336) {
	case 0:
		goto L51
	case 1:
		goto L50
	default:
		goto L24
	}
L53:
	;
	switch v147 - int32(103) {
	case 0:
		goto L51
	case 1:
		goto L50
	default:
		goto L52
	}
L54:
	;
	v173 = v75
	goto L20
L55:
	;
	v173 = v75
	goto L20
L56:
	;
	goto L19
L57:
	;
	v205 = v195
	v212 = v192
	goto L58
L58:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+8)))
	if v215 == int32(1) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v268 = v253
	goto L1
L60:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	F_pfree(m, v205)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L5
	} else {
		goto L74
	}
L61:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+9)))
	if v218 == int32(1) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L63
L63:
	;
	v246 = *(*float64)(unsafe.Add(mBase, uint32(v205)+24))
	v247 = v246
	goto L60
L64:
	;
	v221 = float64(0.005)
	v222 = *(*float64)(unsafe.Add(mBase, uint32(v205)+24))
	if base.F64_eq(v222, float64(0.3333333333333333)) != 0 {
		v247 = v221
		goto L60
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v245 = *(*float64)(unsafe.Add(mBase, uint32(v205)+16))
	v247 = v245
	goto L60
L67:
	;
	v225 = *(*float64)(unsafe.Add(mBase, uint32(v205)+16))
	if base.F64_eq(v225, float64(0.3333333333333333)) != 0 {
		v247 = v221
		goto L60
	} else {
		goto L68
	}
L68:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v205)+4))
	v233 = F_nulltestsel(m, l0, int32(0), v232, l2)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	v235 = base.F64_add(base.F64_add(base.F64_add(v222, v225), float64(-1)), v233)
	if base.F64_le(v235, float64(0)) == int32(0) {
		v247 = v235
		goto L60
	} else {
		goto L70
	}
L70:
	;
	if base.F64_lt(v235, float64(-0.01)) != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v244 = float64(0.005)
	goto L73
L72:
	;
	v244 = float64(1e-10)
	goto L73
L73:
	;
	v247 = v244
	goto L60
L74:
	;
	v253 = base.F64_mul(v212, v247)
	if v250 != 0 {
		v205 = v250
		v212 = v253
		goto L58
	} else {
		goto L75
	}
L75:
	;
	goto L59
}
func F_cleartraverse(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v7 = m.T0[v6].(func(*base.Module) int32)(m)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v7 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(101)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v13 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v17 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v15 = v13
	goto L8
L7:
	;
	v15 = int32(19)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v15
	return
L9:
	;
	return
L10:
	;
	v20 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v22 == v20 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v26 = v22
	goto L12
L12:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	F_cleartraverse(m, l0, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	goto L9
L14:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	if v30 != 0 {
		v26 = v30
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
}
func F_clonesuccessorstates(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v209 int32
	_ = v209
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v305 int64
	_ = v305
	var v311 int32
	_ = v311
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v395 int64
	_ = v395
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v457 int32
	_ = v457
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v516 int32
	_ = v516
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v19 = m.T0[v18].(func(*base.Module) int32)(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v19 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = int32(101)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	if v25 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	if l5 != 0 {
		v51 = l5
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v27 = v25
	goto L8
L7:
	;
	v27 = int32(19)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v27
	return
L9:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v54 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v51+v52))) = uint8(v54)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v56 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L10:
	;
	v30 = F_palloc_extended(m, l7, int32(2))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if v30 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = int32(101)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	if v38 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	if l6 != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v40 = v38
	goto L17
L16:
	;
	v40 = int32(12)
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = v40
	return
L18:
	;
	if l7 != 0 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	goto L20
L20:
	;
	v46 = F__emscripten_memset_bulkmem(m, v30, base.I32_extend8_s(int32(0)), l7)
	mBase = m.M
	goto L25
L21:
	;
	v51 = v30
	goto L9
L22:
	;
	v42 = F__emscripten_memcpy_bulkmem(m, v30, l6, l7)
	mBase = m.M
	goto L24
L23:
	;
	goto L24
L24:
	;
	goto L21
L25:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v49 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v46+v47))) = uint8(v49)
	v51 = v30
	goto L9
L26:
	;
	if l5 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L27:
	;
	v68 = v56
	goto L28
L28:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
	if v74 != 0 {
		goto L26
	} else {
		goto L30
	}
L29:
	;
	goto L26
L30:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	switch v76 - int32(76) {
	case 0, 18, 21, 38:
		goto L34
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37:
		goto L33
	default:
		goto L35
	}
L31:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
	if v457 != 0 {
		v68 = v457
		goto L28
	} else {
		goto L151
	}
L32:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+v120))))
	if v122 != 0 {
		goto L31
	} else {
		goto L45
	}
L33:
	;
	F_cparc(m, l0, v68, l2, v75)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L44
	}
L34:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
	if v81 == int32(0) {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	if v76 != int32(36) {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v85 = v81
	goto L38
L38:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	switch v98 - int32(76) {
	case 0, 18, 21, 38:
		goto L32
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37:
		goto L40
	default:
		goto L41
	}
L39:
	;
	goto L33
L40:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
	if v103 != 0 {
		v85 = v103
		goto L38
	} else {
		goto L43
	}
L41:
	;
	if v98 == int32(36) {
		goto L32
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	goto L39
L44:
	;
	goto L31
L45:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v123 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if l4 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L47:
	;
	v154 = int32(0)
	goto L46
L48:
	;
	goto L49
L49:
	;
	v128 = v123
	goto L50
L50:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+24))
	if v142 == v75 {
		v154 = v141
		goto L46
	} else {
		goto L52
	}
L51:
	;
	v154 = int32(0)
	goto L46
L52:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
	if v144 != 0 {
		v128 = v144
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v436 = F_newstate(m, l0)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L148
	}
L55:
	;
	if v154 != 0 {
		goto L72
	} else {
		goto L73
	}
L56:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v167 != 0 {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v76 != v162 {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+4)))
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+4)))
	if v164 == v165 {
		goto L55
	} else {
		goto L59
	}
L59:
	;
	goto L56
L60:
	;
	v169 = v167
	v178 = l2
	goto L63
L61:
	;
	goto L62
L62:
	;
	if v154 == int32(0) {
		goto L54
	} else {
		goto L70
	}
L63:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v178)+8))
	if v182 != int32(1) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	goto L62
L65:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v169)+8))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)+16))
	if v191 != 0 {
		v169 = v191
		v178 = v190
		goto L63
	} else {
		goto L69
	}
L66:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	if v76 != v185 {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+4)))
	v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v169)+4)))
	if v187 == v188 {
		goto L55
	} else {
		goto L68
	}
L68:
	;
	goto L65
L69:
	;
	goto L64
L70:
	;
	F_cparc(m, l0, v68, l2, v154)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	goto L31
L72:
	;
	goto L75
L73:
	;
	goto L74
L74:
	;
	F_clonesuccessorstates(m, l0, v75, l2, l3, l4, v51, l6, l7)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L147
	}
L75:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v154)+16))
	if v238 != 0 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	goto L107
L77:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	v246 = int32(*(*int16)(unsafe.Add(mBase, uint32(v238)+4)))
	if v246 < int32(0) {
		goto L81
	} else {
		goto L82
	}
L78:
	;
	goto L79
L79:
	;
	goto L76
L80:
	;
	goto L75
L81:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v238)+16))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v238)+20))
	if v280 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L82:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	v251 = v249 - int32(97)
	if base.Ui32(int32(17)) < base.Ui32(v251) {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	if int32(1)<<(uint(v251)%32)&int32(_a_F_clonesuccessorstates_0) == int32(0) {
		goto L81
	} else {
		goto L84
	}
L84:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v260 != 0 {
		goto L81
	} else {
		goto L85
	}
L85:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v238)+36))
	if v261 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	if v273 != 0 {
		goto L90
	} else {
		goto L91
	}
L87:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+20))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v238)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v265+v246*int32(24))+12)) = v269
	v273 = v269
	goto L86
L88:
	;
	goto L89
L89:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v238)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v261)+32)) = v271
	v273 = v271
	goto L86
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v273)+36)) = v261
	goto L92
L91:
	;
	goto L92
L92:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v238)+32)) = int64(0)
	goto L81
L93:
	;
	if v279 != 0 {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v245)+20)) = v279
	goto L93
L95:
	;
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v280)+16)) = v279
	goto L93
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v279)+20)) = v280
	goto L99
L98:
	;
	goto L99
L99:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v245)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v245)+12)) = v286 - int32(1)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v238)+24))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v238)+28))
	if v291 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v297 = v238 + int32(8)
	if v290 != 0 {
		goto L104
	} else {
		goto L105
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v244)+16)) = v290
	goto L100
L102:
	;
	goto L103
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291)+24)) = v290
	goto L100
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v290)+28)) = v291
	goto L106
L105:
	;
	goto L106
L106:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v244)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v244)+8)) = v299 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v238))) = int32(0)
	v305 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v297)+16)) = v305
	*(*int64)(unsafe.Add(mBase, uint32(v297)+8)) = v305
	*(*int64)(unsafe.Add(mBase, uint32(v297))) = v305
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v238)+16)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v238
	goto L80
L107:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v154)+20))
	if v328 != 0 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v404 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v154)+4)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v154))) = int32(-1)
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v154)+32))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v154)+28))
	if v409 != 0 {
		goto L140
	} else {
		goto L141
	}
L109:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v328)+12))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v328)+8))
	v336 = int32(*(*int16)(unsafe.Add(mBase, uint32(v328)+4)))
	if v336 < int32(0) {
		goto L113
	} else {
		goto L114
	}
L110:
	;
	goto L111
L111:
	;
	goto L108
L112:
	;
	goto L107
L113:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v328)+16))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v328)+20))
	if v370 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L114:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v328)))
	v341 = v339 - int32(97)
	if base.Ui32(int32(17)) < base.Ui32(v341) {
		goto L113
	} else {
		goto L115
	}
L115:
	;
	if int32(1)<<(uint(v341)%32)&int32(_a_F_clonesuccessorstates_0) == int32(0) {
		goto L113
	} else {
		goto L116
	}
L116:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v350 != 0 {
		goto L113
	} else {
		goto L117
	}
L117:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v328)+36))
	if v351 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	if v363 != 0 {
		goto L122
	} else {
		goto L123
	}
L119:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v354)+20))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v328)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v355+v336*int32(24))+12)) = v359
	v363 = v359
	goto L118
L120:
	;
	goto L121
L121:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v328)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v351)+32)) = v361
	v363 = v361
	goto L118
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v363)+36)) = v351
	goto L124
L123:
	;
	goto L124
L124:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v328)+32)) = int64(0)
	goto L113
L125:
	;
	if v369 != 0 {
		goto L129
	} else {
		goto L130
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v335)+20)) = v369
	goto L125
L127:
	;
	goto L128
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v370)+16)) = v369
	goto L125
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v369)+20)) = v370
	goto L131
L130:
	;
	goto L131
L131:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v335)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v335)+12)) = v376 - int32(1)
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v328)+24))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v328)+28))
	if v381 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	v387 = v328 + int32(8)
	if v380 != 0 {
		goto L136
	} else {
		goto L137
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v334)+16)) = v380
	goto L132
L134:
	;
	goto L135
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v381)+24)) = v380
	goto L132
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v380)+28)) = v381
	goto L138
L137:
	;
	goto L138
L138:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v334)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v334)+8)) = v389 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v328))) = int32(0)
	v395 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v387)+16)) = v395
	*(*int64)(unsafe.Add(mBase, uint32(v387)+8)) = v395
	*(*int64)(unsafe.Add(mBase, uint32(v387))) = v395
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v328)+16)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v328
	goto L112
L139:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v154)+28))
	if v408 != 0 {
		goto L144
	} else {
		goto L145
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v409)+32)) = v408
	goto L139
L141:
	;
	goto L142
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v408
	goto L139
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154)+32)) = int32(0)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v154)+28)) = v417
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v154
	goto L74
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v408)+28)) = v412
	goto L143
L145:
	;
	goto L146
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v412
	goto L143
L147:
	;
	goto L31
L148:
	;
	if v436 == int32(0) {
		goto L26
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v436)+24)) = v75
	F_cparc(m, l0, v68, l2, v436)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	goto L31
L151:
	;
	goto L29
L152:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v474 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L153:
	;
	goto L154
L154:
	;
	return
L155:
	;
	F_pfree(m, v51)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L165
	}
L156:
	;
	v478 = v474
	goto L157
L157:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v491)+12))
	if v492 != 0 {
		goto L155
	} else {
		goto L159
	}
L158:
	;
	goto L155
L159:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v478)+12))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v493)+24))
	if v494 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v495 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v493)+24)) = v495
	F_clonesuccessorstates(m, l0, v494, v493, l3, l4, v495, v51, l7)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L1
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v478)+16))
	if v500 != 0 {
		v478 = v500
		goto L157
	} else {
		goto L164
	}
L163:
	;
	goto L162
L164:
	;
	goto L158
L165:
	;
	goto L154
}
func F_cmpOffsetNumbers(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	v4 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	return v3 - v4
}
func F_cmpTheLexeme(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v7 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v72
L2:
	;
	if v6 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v39 = base.B2i32(v6 != int32(0))
	goto L4
L4:
	;
	if v39 != 0 {
		v72 = v39
		goto L1
	} else {
		goto L16
	}
L5:
	;
	return int32(-1)
L6:
	;
	goto L7
L7:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v15 == int32(0) {
		v34 = v14
		v35 = v15
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v39 = v35 - v34
	goto L4
L9:
	;
	goto L8
L10:
	;
	if v14 != v15 {
		v34 = v14
		v35 = v15
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v19 = v7
	v20 = v6
	goto L12
L12:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v24 == int32(0) {
		v34 = v23
		v35 = v24
		goto L9
	} else {
		goto L14
	}
L13:
	;
	v34 = v23
	v35 = v24
	goto L9
L14:
	;
	v27 = int32(1)
	if v23 == v24 {
		v19 = v19 + v27
		v20 = v20 + v27
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v40 = int32(0)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v41 == v40 {
		v72 = v40
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v44 == int32(0) {
		v72 = v40
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v47 == v48 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+4)))
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+4)))
	if v50 == v51 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	if base.Ui32(v48) < base.Ui32(v47) {
		goto L32
	} else {
		goto L33
	}
L22:
	;
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+6)))
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+6)))
	if v53 == v54 {
		v72 = v40
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if base.Ui32(v51) < base.Ui32(v50) {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	if base.Ui32(v54) < base.Ui32(v53) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v59 = int32(-1)
	goto L28
L27:
	;
	v59 = int32(1)
	goto L28
L28:
	;
	return v59
L29:
	;
	v64 = int32(-1)
	goto L31
L30:
	;
	v64 = int32(1)
	goto L31
L31:
	;
	return v64
L32:
	;
	v69 = int32(-1)
	goto L34
L33:
	;
	v69 = int32(1)
	goto L34
L34:
	;
	v72 = v69
	goto L1
}
func F_cmp_abs_common(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v170 int32
	_ = v170
	v7 = int32(0)
	if base.B2i32(l5 < l2)&base.B2i32(v7 < l1) == v7 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v170
L2:
	;
	if l5 <= v38 {
		v73 = l5
		v75 = v7
		goto L11
	} else {
		goto L12
	}
L3:
	;
	v38 = l2
	v42 = v7
	goto L2
L4:
	;
	goto L5
L5:
	;
	v19 = l2
	v23 = v7
	goto L6
L6:
	;
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+v23<<(uint(int32(1))%32)))))
	if v29 != 0 {
		v170 = int32(1)
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v38 = v33
	v42 = v31
	goto L2
L8:
	;
	v30 = int32(1)
	v31 = v23 + v30
	v33 = v19 - v30
	if v33 <= l5 {
		v38 = v33
		v42 = v31
		goto L2
	} else {
		goto L9
	}
L9:
	;
	if v31 < l1 {
		v19 = v33
		v23 = v31
		goto L6
	} else {
		goto L10
	}
L10:
	;
	goto L7
L11:
	;
	if v38 != v73 {
		v115 = v42
		v116 = v75
		goto L19
	} else {
		goto L20
	}
L12:
	;
	if l4 <= int32(0) {
		v73 = l5
		v75 = v7
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v54 = l5
	v56 = v7
	goto L14
L14:
	;
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3+v56<<(uint(int32(1))%32)))))
	if v61 != 0 {
		v170 = int32(-1)
		goto L1
	} else {
		goto L16
	}
L15:
	;
	v73 = v65
	v75 = v63
	goto L11
L16:
	;
	v62 = int32(1)
	v63 = v56 + v62
	v65 = v54 - v62
	if v65 <= v38 {
		v73 = v65
		v75 = v63
		goto L11
	} else {
		goto L17
	}
L17:
	;
	if v63 < l4 {
		v54 = v65
		v56 = v63
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	if l1 < v115 {
		goto L29
	} else {
		goto L30
	}
L20:
	;
	v84 = v42
	v85 = v75
	goto L21
L21:
	;
	if l1 <= v84 {
		v115 = v84
		v116 = v85
		goto L19
	} else {
		goto L23
	}
L22:
	;
	if base.I32_extend16_s(v100) < base.I32_extend16_s(v98) {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	if l4 <= v85 {
		v115 = v84
		v116 = v85
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v89 = int32(1)
	v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+v84<<(uint(v89)%32)))))
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85<<(uint(v89)%32)+l3))))
	if v98 == v100 {
		v84 = v84 + v89
		v85 = v85 + v89
		goto L21
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	v107 = int32(1)
	goto L28
L27:
	;
	v107 = int32(-1)
	goto L28
L28:
	;
	return v107
L29:
	;
	v119 = v115
	goto L31
L30:
	;
	v119 = l1
	goto L31
L31:
	;
	v126 = v115
	goto L32
L32:
	;
	if v119 == v126 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v170 = v153
	goto L1
L34:
	;
	if l4 < v116 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v153 = int32(1)
	v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+v126<<(uint(v153)%32)))))
	if v159 == int32(0) {
		v126 = v126 + v153
		goto L32
	} else {
		goto L46
	}
L37:
	;
	v131 = v116
	goto L39
L38:
	;
	v131 = l4
	goto L39
L39:
	;
	v139 = v116
	goto L40
L40:
	;
	if v131 == v139 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v170 = int32(-1)
	goto L1
L42:
	;
	return int32(0)
L43:
	;
	goto L44
L44:
	;
	v144 = int32(1)
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3+v139<<(uint(v144)%32)))))
	if v149 == int32(0) {
		v139 = v139 + v144
		goto L40
	} else {
		goto L45
	}
L45:
	;
	goto L41
L46:
	;
	goto L33
}
func F_cmp_numerics(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v306 int32
	_ = v306
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v483 int32
	_ = v483
	var v493 int32
	_ = v493
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v17 = int32(_a_F_cmp_numerics_0)
	v18 = v16 & v17
	if v18 == v17 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v16 != int32(_a_F_cmp_numerics_1) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	if base.Ui32(int32(_a_F_cmp_numerics_0)) <= base.Ui32(v15) {
		goto L15
	} else {
		goto L16
	}
L4:
	;
	if v15 != int32(_a_F_cmp_numerics_2) {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	if v16 != int32(_a_F_cmp_numerics_0) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	if v15 == int32(_a_F_cmp_numerics_0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	return base.B2i32(v15 != int32(_a_F_cmp_numerics_0))
L9:
	;
	v33 = int32(-1)
	goto L11
L10:
	;
	v33 = base.B2i32(v15 != int32(_a_F_cmp_numerics_1))
	goto L11
L11:
	;
	return v33
L12:
	;
	v39 = int32(-1)
	goto L14
L13:
	;
	v39 = int32(0)
	goto L14
L14:
	;
	return v39
L15:
	;
	if v15 == int32(_a_F_cmp_numerics_2) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v50 = l0 + int32(6)
	v55 = base.B2i32(int32(0) <= base.I32_extend16_s(v16))
	if int32(0) <= base.I32_extend16_s(v16) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v47 = int32(1)
	goto L20
L19:
	;
	v47 = int32(-1)
	goto L20
L20:
	;
	return v47
L21:
	;
	v56 = int32(-8)
	goto L23
L22:
	;
	v56 = int32(-6)
	goto L23
L23:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if int32(0) <= base.I32_extend16_s(v16) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v60 = int32(*(*int16)(unsafe.Add(mBase, uint32(v50))))
	v70 = v60
	goto L26
L25:
	;
	v70 = v16<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v16&int32(63)
	goto L26
L26:
	;
	v71 = v56 + int32(base.Ui32(v57)>>(uint(int32(2))%32))
	v73 = l1 + int32(6)
	v78 = base.B2i32(int32(0) <= base.I32_extend16_s(v15))
	if int32(0) <= base.I32_extend16_s(v15) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v79 = int32(-8)
	goto L29
L28:
	;
	v79 = int32(-6)
	goto L29
L29:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if int32(0) <= base.I32_extend16_s(v15) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(v73))))
	v93 = v83
	goto L32
L31:
	;
	v93 = v15<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v15&int32(63)
	goto L32
L32:
	;
	v94 = v79 + int32(base.Ui32(v80)>>(uint(int32(2))%32))
	v100 = v15 & int32(_a_F_cmp_numerics_0)
	if v100 == int32(_a_F_cmp_numerics_3) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v103 = v15 << (uint(int32(1)) % 32) & int32(_a_F_cmp_numerics_4)
	goto L35
L34:
	;
	v103 = v100
	goto L35
L35:
	;
	if base.Ui32(v71) <= base.Ui32(int32(1)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if base.Ui32(v94) < base.Ui32(int32(2)) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	if v18 == int32(_a_F_cmp_numerics_3) {
		goto L45
	} else {
		goto L46
	}
L39:
	;
	return int32(0)
L40:
	;
	goto L41
L41:
	;
	if v103 == int32(_a_F_cmp_numerics_4) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v114 = int32(1)
	goto L44
L43:
	;
	v114 = int32(-1)
	goto L44
L44:
	;
	return v114
L45:
	;
	v122 = v16 << (uint(int32(1)) % 32) & int32(_a_F_cmp_numerics_4)
	goto L47
L46:
	;
	v122 = v18
	goto L47
L47:
	;
	if base.Ui32(v94) <= base.Ui32(int32(1)) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	if v122 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	if int32(0) <= base.I32_extend16_s(v16) {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v127 = int32(-1)
	goto L53
L52:
	;
	v127 = int32(1)
	goto L53
L53:
	;
	return v127
L54:
	;
	v131 = l0 + int32(8)
	goto L56
L55:
	;
	v131 = v50
	goto L56
L56:
	;
	v133 = int32(base.Ui32(v71) >> (uint(int32(1)) % 32))
	if int32(0) <= base.I32_extend16_s(v15) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v136 = l1 + int32(8)
	goto L59
L58:
	;
	v136 = v73
	goto L59
L59:
	;
	v138 = int32(base.Ui32(v94) >> (uint(int32(1)) % 32))
	if v122 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if v103 == int32(_a_F_cmp_numerics_4) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	if v103 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L63:
	;
	return int32(1)
L64:
	;
	goto L65
L65:
	;
	v145 = int32(0)
	if base.B2i32(v93 < v70)&base.B2i32(v145 < v133) == v145 {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	return v316
L67:
	;
	v316 = v306
	goto L66
L68:
	;
	if v93 <= v176 {
		v211 = v93
		v213 = v145
		goto L77
	} else {
		goto L78
	}
L69:
	;
	v176 = v70
	v180 = v145
	goto L68
L70:
	;
	goto L71
L71:
	;
	v157 = v70
	v161 = v145
	goto L72
L72:
	;
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131+v161<<(uint(int32(1))%32)))))
	if v167 != 0 {
		v306 = int32(1)
		goto L67
	} else {
		goto L74
	}
L73:
	;
	v176 = v171
	v180 = v169
	goto L68
L74:
	;
	v168 = int32(1)
	v169 = v161 + v168
	v171 = v157 - v168
	if v171 <= v93 {
		v176 = v171
		v180 = v169
		goto L68
	} else {
		goto L75
	}
L75:
	;
	if v169 < v133 {
		v157 = v171
		v161 = v169
		goto L72
	} else {
		goto L76
	}
L76:
	;
	goto L73
L77:
	;
	if v176 != v211 {
		v252 = v180
		v253 = v213
		goto L85
	} else {
		goto L86
	}
L78:
	;
	if v138 <= int32(0) {
		v211 = v93
		v213 = v145
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v192 = v93
	v194 = v145
	goto L80
L80:
	;
	v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136+v194<<(uint(int32(1))%32)))))
	if v199 != 0 {
		v306 = int32(-1)
		goto L67
	} else {
		goto L82
	}
L81:
	;
	v211 = v203
	v213 = v201
	goto L77
L82:
	;
	v200 = int32(1)
	v201 = v194 + v200
	v203 = v192 - v200
	if v203 <= v176 {
		v211 = v203
		v213 = v201
		goto L77
	} else {
		goto L83
	}
L83:
	;
	if v201 < v138 {
		v192 = v203
		v194 = v201
		goto L80
	} else {
		goto L84
	}
L84:
	;
	goto L81
L85:
	;
	if v133 < v252 {
		goto L95
	} else {
		goto L96
	}
L86:
	;
	v222 = v180
	v223 = v213
	goto L87
L87:
	;
	if v133 <= v222 {
		v252 = v222
		v253 = v223
		goto L85
	} else {
		goto L89
	}
L88:
	;
	if base.I32_extend16_s(v238) < base.I32_extend16_s(v236) {
		goto L92
	} else {
		goto L93
	}
L89:
	;
	if v138 <= v223 {
		v252 = v222
		v253 = v223
		goto L85
	} else {
		goto L90
	}
L90:
	;
	v227 = int32(1)
	v236 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131+v222<<(uint(v227)%32)))))
	v238 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v223<<(uint(v227)%32)+v136))))
	if v236 == v238 {
		v222 = v222 + v227
		v223 = v223 + v227
		goto L87
	} else {
		goto L91
	}
L91:
	;
	goto L88
L92:
	;
	v245 = int32(1)
	goto L94
L93:
	;
	v245 = int32(-1)
	goto L94
L94:
	;
	v316 = v245
	goto L66
L95:
	;
	v256 = v252
	goto L97
L96:
	;
	v256 = v133
	goto L97
L97:
	;
	v263 = v252
	goto L98
L98:
	;
	if v256 == v263 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v306 = v289
	goto L67
L100:
	;
	if v138 < v253 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	goto L102
L102:
	;
	v289 = int32(1)
	v295 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131+v263<<(uint(v289)%32)))))
	if v295 == int32(0) {
		v263 = v263 + v289
		goto L98
	} else {
		goto L112
	}
L103:
	;
	v268 = v253
	goto L105
L104:
	;
	v268 = v138
	goto L105
L105:
	;
	v276 = v253
	goto L106
L106:
	;
	if v268 == v276 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v306 = int32(-1)
	goto L67
L108:
	;
	v316 = int32(0)
	goto L66
L109:
	;
	goto L110
L110:
	;
	v280 = int32(1)
	v285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136+v276<<(uint(v280)%32)))))
	if v285 == int32(0) {
		v276 = v276 + v280
		goto L106
	} else {
		goto L111
	}
L111:
	;
	goto L107
L112:
	;
	goto L99
L113:
	;
	return int32(-1)
L114:
	;
	goto L115
L115:
	;
	v322 = int32(0)
	if base.B2i32(v70 < v93)&base.B2i32(v322 < v138) == v322 {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	return v493
L117:
	;
	v493 = v483
	goto L116
L118:
	;
	if v70 <= v353 {
		v388 = v70
		v390 = v322
		goto L127
	} else {
		goto L128
	}
L119:
	;
	v353 = v93
	v357 = v322
	goto L118
L120:
	;
	goto L121
L121:
	;
	v334 = v93
	v338 = v322
	goto L122
L122:
	;
	v344 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136+v338<<(uint(int32(1))%32)))))
	if v344 != 0 {
		v483 = int32(1)
		goto L117
	} else {
		goto L124
	}
L123:
	;
	v353 = v348
	v357 = v346
	goto L118
L124:
	;
	v345 = int32(1)
	v346 = v338 + v345
	v348 = v334 - v345
	if v348 <= v70 {
		v353 = v348
		v357 = v346
		goto L118
	} else {
		goto L125
	}
L125:
	;
	if v346 < v138 {
		v334 = v348
		v338 = v346
		goto L122
	} else {
		goto L126
	}
L126:
	;
	goto L123
L127:
	;
	if v353 != v388 {
		v429 = v357
		v430 = v390
		goto L135
	} else {
		goto L136
	}
L128:
	;
	if v133 <= int32(0) {
		v388 = v70
		v390 = v322
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v369 = v70
	v371 = v322
	goto L130
L130:
	;
	v376 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131+v371<<(uint(int32(1))%32)))))
	if v376 != 0 {
		v483 = int32(-1)
		goto L117
	} else {
		goto L132
	}
L131:
	;
	v388 = v380
	v390 = v378
	goto L127
L132:
	;
	v377 = int32(1)
	v378 = v371 + v377
	v380 = v369 - v377
	if v380 <= v353 {
		v388 = v380
		v390 = v378
		goto L127
	} else {
		goto L133
	}
L133:
	;
	if v378 < v133 {
		v369 = v380
		v371 = v378
		goto L130
	} else {
		goto L134
	}
L134:
	;
	goto L131
L135:
	;
	if v138 < v429 {
		goto L145
	} else {
		goto L146
	}
L136:
	;
	v399 = v357
	v400 = v390
	goto L137
L137:
	;
	if v138 <= v399 {
		v429 = v399
		v430 = v400
		goto L135
	} else {
		goto L139
	}
L138:
	;
	if base.I32_extend16_s(v415) < base.I32_extend16_s(v413) {
		goto L142
	} else {
		goto L143
	}
L139:
	;
	if v133 <= v400 {
		v429 = v399
		v430 = v400
		goto L135
	} else {
		goto L140
	}
L140:
	;
	v404 = int32(1)
	v413 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136+v399<<(uint(v404)%32)))))
	v415 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v400<<(uint(v404)%32)+v131))))
	if v413 == v415 {
		v399 = v399 + v404
		v400 = v400 + v404
		goto L137
	} else {
		goto L141
	}
L141:
	;
	goto L138
L142:
	;
	v422 = int32(1)
	goto L144
L143:
	;
	v422 = int32(-1)
	goto L144
L144:
	;
	v493 = v422
	goto L116
L145:
	;
	v433 = v429
	goto L147
L146:
	;
	v433 = v138
	goto L147
L147:
	;
	v440 = v429
	goto L148
L148:
	;
	if v433 == v440 {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	v483 = v466
	goto L117
L150:
	;
	if v133 < v430 {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	goto L152
L152:
	;
	v466 = int32(1)
	v472 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136+v440<<(uint(v466)%32)))))
	if v472 == int32(0) {
		v440 = v440 + v466
		goto L148
	} else {
		goto L162
	}
L153:
	;
	v445 = v430
	goto L155
L154:
	;
	v445 = v133
	goto L155
L155:
	;
	v453 = v430
	goto L156
L156:
	;
	if v445 == v453 {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	v483 = int32(-1)
	goto L117
L158:
	;
	v493 = int32(0)
	goto L116
L159:
	;
	goto L160
L160:
	;
	v457 = int32(1)
	v462 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131+v453<<(uint(v457)%32)))))
	if v462 == int32(0) {
		v453 = v453 + v457
		goto L156
	} else {
		goto L161
	}
L161:
	;
	goto L157
L162:
	;
	goto L149
}
func F_collprovider_name(m *base.Module, l0 int32) int32 {
	var v11 int32
	_ = v11
	switch l0 - int32(98) {
	case 0:
		v11 = int32(_a_F_collprovider_name_0)
		return v11
	case 1:
		return int32(_a_F_collprovider_name_1)
	default:
		v11 = int32(_a_F_collprovider_name_2)
		return v11
	case 7:
		return int32(_a_F_collprovider_name_3)
	}
}
func F_combo_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v16 int32
	_ = v16
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v3 != 0 {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+24))
		m.T0[v4].(func(*base.Module, int32))(m, v3)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			v7 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v7
			*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v7
			*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v7
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v7
			F_pfree(m, l0)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		v7 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v7
		*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v7
		*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v7
		*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v7
		F_pfree(m, l0)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			return
		}
	}
}
func F_commit_cb_wrapper(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int64
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(_a_F_commit_cb_wrapper_0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v11
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(992)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v15
	v19 = int32(_a_F_commit_cb_wrapper_1)
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_commit_cb_wrapper[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_commit_cb_wrapper[0])) = v9 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v9 + int32(16)
	v29 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+147)) = uint8(v29)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+160)) = v31
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+164)) = uint8(v29)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+152)) = v33
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	m.T0[v37].(func(*base.Module, int32, int32, int64))(m, v11, l1, l2)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		return
	} else {
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		*(*int32)(unsafe.Add(mBase, _c_F_commit_cb_wrapper[0])) = v41
		m.G0 = v9 + int32(32)
		return
	}
}
func F_commit_prepared_cb_wrapper(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int64
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(_a_F_commit_prepared_cb_wrapper_0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v11
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(992)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v15
	v19 = int32(_a_F_commit_prepared_cb_wrapper_1)
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_commit_prepared_cb_wrapper[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_commit_prepared_cb_wrapper[0])) = v9 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v9 + int32(16)
	v29 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+147)) = uint8(v29)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+160)) = v31
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+164)) = uint8(v29)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+152)) = v33
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+68))
	if v37 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(_a_F_commit_prepared_cb_wrapper_2)
				F_errmsg(m, int32(_a_F_commit_prepared_cb_wrapper_3), v9)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_commit_prepared_cb_wrapper_4), int32(1031), int32(_a_F_commit_prepared_cb_wrapper_5))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
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
		m.T0[v37].(func(*base.Module, int32, int32, int64))(m, v11, l1, l2)
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return
		} else {
			v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			*(*int32)(unsafe.Add(mBase, _c_F_commit_prepared_cb_wrapper[0])) = v60
			m.G0 = v9 + int32(32)
			return
		}
	}
}
func F_committssyncfiletag(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_SlruSyncFileTag(m, int32(_a_F_committssyncfiletag_0), l0, l1)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_compare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v8 float64
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v7 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	if base.F64_gt(v7, v8) != 0 {
		v10 = int32(1)
	} else {
		v10 = int32(-1)
	}
	if base.F64_ne(v7, v8) != 0 {
		v13 = v10
	} else {
		v13 = int32(0)
	}
	return v13
}
func F_compareWORD(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v5 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	if v5 == int32(0) {
		v13 = int32(0)
		if v13 < v7 {
			v16 = int32(-1)
		} else {
			v16 = v13
		}
		v33 = v16
	} else {
		if v7 == int32(0) {
			v33 = base.B2i32(int32(0) < v5)
		} else {
			if base.Ui32(v5) < base.Ui32(v7) {
				v22 = v5
			} else {
				v22 = v7
			}
			v23 = F_memcmp(m, v4, v6, v22)
			mBase = m.M
			if v23 != 0 {
				v31 = v23
				v33 = v31
			} else {
				if v5 == v7 {
					v33 = int32(0)
				} else {
					if v5 < v7 {
						v30 = int32(-1)
					} else {
						v30 = int32(1)
					}
					v31 = v30
					v33 = v31
				}
			}
		}
	}
	if v33 != 0 {
		v44 = v33
	} else {
		v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
		v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
		if v35 == v36 {
			v44 = int32(0)
		} else {
			if base.Ui32(v36) < base.Ui32(v35) {
				v41 = int32(1)
			} else {
				v41 = int32(-1)
			}
			v44 = v41
		}
	}
	return v44
}
func F_compareWordEntryPos(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v3 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	v4 = int32(_a_F_compareWordEntryPos_0)
	v5 = v3 & v4
	v6 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v8 = v6 & v4
	return base.B2i32(base.Ui32(v8) < base.Ui32(v5)) - base.B2i32(base.Ui32(v5) < base.Ui32(v8))
}
func F_compare_mcvs(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	return v4 - v5
}
func F_compare_pathkeys(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
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
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	if l0 == l1 {
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
	v11 = int32(0)
	goto L6
L4:
	;
	if v51 != 0 {
		goto L19
	} else {
		goto L20
	}
L5:
	;
	v45 = int32(0)
	v51 = base.B2i32(v24 == v45)
	v53 = base.B2i32(v33 != v45) << (uint(int32(1)) % 32)
	goto L4
L6:
	;
	v14 = int32(0)
	if l0 == v14 {
		v24 = v14
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return int32(3)
L8:
	;
	if l1 != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v18 <= v11 {
		v24 = int32(0)
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v24 = v20 + v11<<(uint(int32(2))%32)
	goto L8
L11:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v33 = v30 + v11<<(uint(int32(2))%32)
	if v24 == int32(0) {
		goto L5
	} else {
		goto L16
	}
L12:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v11 < v25 {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v27 = int32(0)
	v51 = base.B2i32(v24 == v27)
	v53 = v27
	goto L4
L15:
	;
	goto L14
L16:
	;
	if v33 == int32(0) {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if v40 == v41 {
		v11 = v11 + int32(1)
		goto L6
	} else {
		goto L18
	}
L18:
	;
	goto L7
L19:
	;
	v55 = v53
	goto L21
L20:
	;
	v55 = int32(1)
	goto L21
L21:
	;
	return v55
}
func F_compare_rows(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+4)))
	v8 = int32(16)
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+6)))
	v11 = v7<<(uint(v8)%32) | v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)))
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+6)))
	v17 = v13<<(uint(v8)%32) | v16
	if base.Ui32(v11) < base.Ui32(v17) {
		v28 = int32(-1)
	} else {
		if base.Ui32(v17) < base.Ui32(v11) {
			v28 = int32(1)
		} else {
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+8)))
			v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+8)))
			if base.Ui32(v22) < base.Ui32(v23) {
				v28 = int32(-1)
			} else {
				v28 = base.B2i32(base.Ui32(v23) < base.Ui32(v22))
			}
		}
	}
	return v28
}
func F_complete_direction(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	v6 = int32(0)
	v8 = F_plpgsql_yylex(m, l2, l3, l4)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		switch v8 - int32(324) {
		case 0, 5:
			v38 = v6
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v38)
			return
		case 1, 2, 3, 4:
			F_plpgsql_push_back_token(m, v8, l2, l3, l4)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v25 = int32(0)
				v28 = int32(1)
				v32 = F_read_sql_construct(m, int32(324), int32(329), v25, int32(_a_F_complete_direction_0), int32(2), v28, v28, v25, v25, l2, l3, l4)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v32
					v35 = v6
					v36 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)) = uint8(v36)
					v38 = v35
					*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v38)
					return
				}
			}
		default:
			if v8 != int32(282) {
				if v8 != 0 {
					F_plpgsql_push_back_token(m, v8, l2, l3, l4)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						v25 = int32(0)
						v28 = int32(1)
						v32 = F_read_sql_construct(m, int32(324), int32(329), v25, int32(_a_F_complete_direction_0), int32(2), v28, v28, v25, v25, l2, l3, l4)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v32
							v35 = v6
							v36 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)) = uint8(v36)
							v38 = v35
							*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v38)
							return
						}
					}
				} else {
					F_plpgsql_yyerror(m, l3, int32(0), l4, int32(_a_F_complete_direction_1))
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(2147483647)
				v35 = int32(1)
				v36 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)) = uint8(v36)
				v38 = v35
				*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v38)
				return
			}
		}
	}
}
func F_compress_flush(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v8 = m.Env.Pgmem_deflate_finish(m, v7)
	mBase = m.M
	if v8 < int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(-105)
L2:
	;
	goto L3
L3:
	;
	v14 = l1 + int32(8)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	goto L4
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v23 = m.Env.Pgmem_zstream_read(m, v22, v14, v15)
	mBase = m.M
	if v23 < int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	return v32
L6:
	;
	return int32(-105)
L7:
	;
	goto L8
L8:
	;
	if v23 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	goto L11
L11:
	;
	v32 = F_pushf_write(m, l0, v14, v23)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	if int32(0) <= v32 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	goto L5
}
func F_compress_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v4 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		m.Env.Pgmem_zstream_free(m, v5)
		mBase = m.M
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		if v7 != 0 {
			F_ResourceOwnerForget(m, v7, v4, int32(_a_F_compress_free_0))
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				F_pfree(m, v4)
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					v16 = F___memset(m, l0, int32(0), int32(_a_F_compress_free_1))
					mBase = m.M
					F_pfree(m, l0)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			F_pfree(m, v4)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				v16 = F___memset(m, l0, int32(0), int32(_a_F_compress_free_1))
				mBase = m.M
				F_pfree(m, l0)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		v16 = F___memset(m, l0, int32(0), int32(_a_F_compress_free_1))
		mBase = m.M
		F_pfree(m, l0)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			return
		}
	}
}
func F_contains_user_functions_checker(m *base.Module, l0 int32, l1 int32) int32 {
	return base.B2i32(base.Ui32(int32(_a_F_contains_user_functions_checker_0)) < base.Ui32(l0))
}
func F_convert_tuples_by_name(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
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
	v3 = int32(0)
	v7 = F_build_attrmap_by_name_if_req(m, l0, l1, v3)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 != 0 {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v13 = F_palloc(m, int32(28))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v7
				*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
				v20 = F_palloc(m, v11<<(uint(int32(2))%32))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v20
					v23 = F_palloc(m, v11)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v23
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v28 = v26 + int32(1)
						v31 = F_palloc(m, v28<<(uint(int32(2))%32))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v31
							v34 = F_palloc(m, v28)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v34
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v37))) = int32(0)
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
								v41 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v41)
								v44 = v13
								return v44
							}
						}
					}
				}
			}
		} else {
			v44 = v3
			return v44
		}
	}
}
func F_cookDefault(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
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
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l5 != 0 {
		v14 = int32(43)
	} else {
		v14 = int32(30)
	}
	v15 = F_transformExpr(m, l0, l1, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if l5 == int32(0) {
			if l2 != 0 {
				v29 = F_exprType(m, v15)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v34 = F_coerce_to_target_type(m, l0, v15, v29, l2, l3, int32(1), int32(2), int32(-1))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						if v34 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(67141764))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									v69 = F_format_type_be(m, l2)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										v71 = F_format_type_be(m, v29)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v71
											*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v69
											*(*int32)(unsafe.Add(mBase, uint32(v10))) = l4
											F_errmsg(m, int32(_a_F_cookDefault_0), v10)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												F_errhint(m, int32(_a_F_cookDefault_1), int32(0))
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_cookDefault_2), int32(3385), int32(_a_F_cookDefault_3))
													mBase = m.M
													v87 = m.ExcPending
													if v87 != 0 {
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
							v38 = v34
							F_assign_expr_collations(m, l0, v38)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								m.G0 = v10 + int32(16)
								return v38
							}
						}
					}
				}
			} else {
				v38 = v15
				F_assign_expr_collations(m, l0, v38)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					m.G0 = v10 + int32(16)
					return v38
				}
			}
		} else {
			v21 = F_check_nested_generated_walker(m, v15, l0)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = F_contain_mutable_functions_after_planning(m, v15)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					if v23 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(117833860))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_cookDefault_4), int32(0))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_cookDefault_2), int32(3348), int32(_a_F_cookDefault_3))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
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
						if l5 != int32(118) {
							if l2 != 0 {
								v29 = F_exprType(m, v15)
								mBase = m.M
								v30 = m.ExcPending
								if v30 != 0 {
									return int32(0)
								} else {
									v34 = F_coerce_to_target_type(m, l0, v15, v29, l2, l3, int32(1), int32(2), int32(-1))
									mBase = m.M
									v35 = m.ExcPending
									if v35 != 0 {
										return int32(0)
									} else {
										if v34 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v65 = m.ExcPending
											if v65 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(67141764))
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return int32(0)
												} else {
													v69 = F_format_type_be(m, l2)
													mBase = m.M
													v70 = m.ExcPending
													if v70 != 0 {
														return int32(0)
													} else {
														v71 = F_format_type_be(m, v29)
														mBase = m.M
														v72 = m.ExcPending
														if v72 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v71
															*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v69
															*(*int32)(unsafe.Add(mBase, uint32(v10))) = l4
															F_errmsg(m, int32(_a_F_cookDefault_0), v10)
															mBase = m.M
															v78 = m.ExcPending
															if v78 != 0 {
																return int32(0)
															} else {
																F_errhint(m, int32(_a_F_cookDefault_1), int32(0))
																mBase = m.M
																v82 = m.ExcPending
																if v82 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_cookDefault_2), int32(3385), int32(_a_F_cookDefault_3))
																	mBase = m.M
																	v87 = m.ExcPending
																	if v87 != 0 {
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
											v38 = v34
											F_assign_expr_collations(m, l0, v38)
											mBase = m.M
											v41 = m.ExcPending
											if v41 != 0 {
												return int32(0)
											} else {
												m.G0 = v10 + int32(16)
												return v38
											}
										}
									}
								}
							} else {
								v38 = v15
								F_assign_expr_collations(m, l0, v38)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									m.G0 = v10 + int32(16)
									return v38
								}
							}
						} else {
							v27 = F_check_virtual_generated_security_walker(m, v15, l0)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								if l2 != 0 {
									v29 = F_exprType(m, v15)
									mBase = m.M
									v30 = m.ExcPending
									if v30 != 0 {
										return int32(0)
									} else {
										v34 = F_coerce_to_target_type(m, l0, v15, v29, l2, l3, int32(1), int32(2), int32(-1))
										mBase = m.M
										v35 = m.ExcPending
										if v35 != 0 {
											return int32(0)
										} else {
											if v34 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v65 = m.ExcPending
												if v65 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(67141764))
													mBase = m.M
													v68 = m.ExcPending
													if v68 != 0 {
														return int32(0)
													} else {
														v69 = F_format_type_be(m, l2)
														mBase = m.M
														v70 = m.ExcPending
														if v70 != 0 {
															return int32(0)
														} else {
															v71 = F_format_type_be(m, v29)
															mBase = m.M
															v72 = m.ExcPending
															if v72 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v71
																*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v69
																*(*int32)(unsafe.Add(mBase, uint32(v10))) = l4
																F_errmsg(m, int32(_a_F_cookDefault_0), v10)
																mBase = m.M
																v78 = m.ExcPending
																if v78 != 0 {
																	return int32(0)
																} else {
																	F_errhint(m, int32(_a_F_cookDefault_1), int32(0))
																	mBase = m.M
																	v82 = m.ExcPending
																	if v82 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_cookDefault_2), int32(3385), int32(_a_F_cookDefault_3))
																		mBase = m.M
																		v87 = m.ExcPending
																		if v87 != 0 {
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
												v38 = v34
												F_assign_expr_collations(m, l0, v38)
												mBase = m.M
												v41 = m.ExcPending
												if v41 != 0 {
													return int32(0)
												} else {
													m.G0 = v10 + int32(16)
													return v38
												}
											}
										}
									}
								} else {
									v38 = v15
									F_assign_expr_collations(m, l0, v38)
									mBase = m.M
									v41 = m.ExcPending
									if v41 != 0 {
										return int32(0)
									} else {
										m.G0 = v10 + int32(16)
										return v38
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
func F_copy_lladdr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	v4 = l3
	v6 = l5
	if base.Ui32(v4) <= base.Ui32(int32(24)) {
		*(*uint8)(unsafe.Add(mBase, uint32(l1)+11)) = uint8(v4)
		*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v6)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = l4
		v12 = int32(17)
		*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v12)
		if v4 != 0 {
			v16 = F__emscripten_memcpy_bulkmem(m, l1+int32(12), l2, v4)
			mBase = m.M
		} else {
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	} else {
	}
	return
}
func F_cos(m *base.Module, l0 float64) float64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v24 float64
	_ = v24
	var v25 float64
	_ = v25
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v40 float64
	_ = v40
	var v59 int32
	_ = v59
	var v60 float64
	_ = v60
	var v61 float64
	_ = v61
	var v69 float64
	_ = v69
	var v70 float64
	_ = v70
	var v72 float64
	_ = v72
	var v73 float64
	_ = v73
	var v85 float64
	_ = v85
	var v105 float64
	_ = v105
	var v121 float64
	_ = v121
	var v144 float64
	_ = v144
	var v145 float64
	_ = v145
	var v147 float64
	_ = v147
	var v148 float64
	_ = v148
	var v160 float64
	_ = v160
	var v181 float64
	_ = v181
	var v197 float64
	_ = v197
	var v219 float64
	_ = v219
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v14 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(l0))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v14) <= base.Ui32(int32(1072243195)) {
		if base.Ui32(v14) < base.Ui32(int32(1044816030)) {
			v219 = float64(1)
		} else {
			v24 = float64(1)
			v25 = base.F64_mul(l0, l0)
			v27 = base.F64_mul(v25, float64(0.5))
			v28 = base.F64_sub(v24, v27)
			v40 = base.F64_mul(v25, v25)
			v219 = base.F64_add(v28, base.F64_add(base.F64_sub(base.F64_sub(v24, v28), v27), base.F64_sub(base.F64_mul(v25, base.F64_add(base.F64_mul(v25, base.F64_add(base.F64_mul(v25, base.F64_add(base.F64_mul(v25, float64(2.480158728947673e-05)), float64(-0.001388888888887411))), float64(0.0416666666666666))), base.F64_mul(base.F64_mul(v40, v40), base.F64_add(base.F64_mul(v25, base.F64_add(base.F64_mul(v25, float64(-1.1359647557788195e-11)), float64(2.087572321298175e-09))), float64(-2.7557314351390663e-07))))), base.F64_mul(l0, float64(0)))))
		}
	} else {
		if base.Ui32(int32(2146435072)) <= base.Ui32(v14) {
			v219 = base.F64_sub(l0, l0)
		} else {
			v59 = F___rem_pio2(m, l0, v7)
			mBase = m.M
			v60 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
			v61 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
			switch v59&int32(3) - int32(1) {
			case 0:
				v105 = base.F64_mul(v61, v61)
				v121 = base.F64_mul(v61, v105)
				v219 = base.F64_neg(base.F64_sub(v61, base.F64_add(base.F64_sub(base.F64_mul(v105, base.F64_sub(base.F64_mul(v60, float64(0.5)), base.F64_mul(v121, base.F64_add(base.F64_mul(base.F64_mul(v105, base.F64_mul(v105, v105)), base.F64_add(base.F64_mul(v105, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v105, base.F64_add(base.F64_mul(v105, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))))), v60), base.F64_mul(v121, float64(0.16666666666666632)))))
			case 1:
				v144 = float64(1)
				v145 = base.F64_mul(v61, v61)
				v147 = base.F64_mul(v145, float64(0.5))
				v148 = base.F64_sub(v144, v147)
				v160 = base.F64_mul(v145, v145)
				v219 = base.F64_neg(base.F64_add(v148, base.F64_add(base.F64_sub(base.F64_sub(v144, v148), v147), base.F64_sub(base.F64_mul(v145, base.F64_add(base.F64_mul(v145, base.F64_add(base.F64_mul(v145, base.F64_add(base.F64_mul(v145, float64(2.480158728947673e-05)), float64(-0.001388888888887411))), float64(0.0416666666666666))), base.F64_mul(base.F64_mul(v160, v160), base.F64_add(base.F64_mul(v145, base.F64_add(base.F64_mul(v145, float64(-1.1359647557788195e-11)), float64(2.087572321298175e-09))), float64(-2.7557314351390663e-07))))), base.F64_mul(v61, v60)))))
			case 2:
				v181 = base.F64_mul(v61, v61)
				v197 = base.F64_mul(v61, v181)
				v219 = base.F64_sub(v61, base.F64_add(base.F64_sub(base.F64_mul(v181, base.F64_sub(base.F64_mul(v60, float64(0.5)), base.F64_mul(v197, base.F64_add(base.F64_mul(base.F64_mul(v181, base.F64_mul(v181, v181)), base.F64_add(base.F64_mul(v181, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v181, base.F64_add(base.F64_mul(v181, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))))), v60), base.F64_mul(v197, float64(0.16666666666666632))))
			default:
				v69 = float64(1)
				v70 = base.F64_mul(v61, v61)
				v72 = base.F64_mul(v70, float64(0.5))
				v73 = base.F64_sub(v69, v72)
				v85 = base.F64_mul(v70, v70)
				v219 = base.F64_add(v73, base.F64_add(base.F64_sub(base.F64_sub(v69, v73), v72), base.F64_sub(base.F64_mul(v70, base.F64_add(base.F64_mul(v70, base.F64_add(base.F64_mul(v70, base.F64_add(base.F64_mul(v70, float64(2.480158728947673e-05)), float64(-0.001388888888887411))), float64(0.0416666666666666))), base.F64_mul(base.F64_mul(v85, v85), base.F64_add(base.F64_mul(v70, base.F64_add(base.F64_mul(v70, float64(-1.1359647557788195e-11)), float64(2.087572321298175e-09))), float64(-2.7557314351390663e-07))))), base.F64_mul(v61, v60))))
			}
		}
	}
	m.G0 = v7 + int32(16)
	return v219
}
func F_cost_subplan(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 float64
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 float64
	_ = v56
	var v57 float64
	_ = v57
	var v63 float64
	_ = v63
	var v69 float64
	_ = v69
	var v70 float64
	_ = v70
	var v71 int32
	_ = v71
	var v75 float64
	_ = v75
	var v76 float64
	_ = v76
	var v80 float64
	_ = v80
	var v81 float64
	_ = v81
	var v82 int32
	_ = v82
	var v83 float64
	_ = v83
	var v84 float64
	_ = v84
	var v92 float64
	_ = v92
	var v96 float64
	_ = v96
	var v99 float64
	_ = v99
	var v100 float64
	_ = v100
	var v103 float64
	_ = v103
	var v112 float64
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v122 float64
	_ = v122
	var v127 float64
	_ = v127
	var v130 float64
	_ = v130
	var v132 float64
	_ = v132
	v6 = float64(0)
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v17 = F_make_ands_implicit(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v19 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v19
	v23 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v23
	if v17 == v23 {
		v63 = v6
		v69 = float64(0)
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v70 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v71 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v29 <= int32(0) {
		v63 = v6
		v69 = float64(0)
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v36 = int32(0)
	goto L6
L6:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43+v36<<(uint(int32(2))%32))))
	v50 = F_cost_qual_eval_walker(m, v47, v14+int32(8))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v56 = *(*float64)(unsafe.Add(mBase, uint32(v14)+24))
	v57 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
	v63 = v56
	v69 = v57
	goto L3
L8:
	;
	v53 = v36 + int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v53 < v54 {
		v36 = v53
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+64)) = v130
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = v132
	m.G0 = v14 + int32(32)
	return
L11:
	;
	v75 = *(*float64)(unsafe.Add(mBase, _c_F_cost_subplan[0]))
	v76 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	v130 = v63
	v132 = base.F64_add(v69, base.F64_add(base.F64_mul(v75, v76), v70))
	goto L10
L12:
	;
	goto L13
L13:
	;
	v80 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v81 = base.F64_sub(v70, v80)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v82 {
	case 0:
		goto L17
	case 1, 2:
		goto L16
	default:
		goto L15
	}
L14:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v113 != 0 {
		v127 = v80
		goto L22
	} else {
		goto L23
	}
L15:
	;
	v112 = base.F64_add(v63, v81)
	goto L14
L16:
	;
	v99 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	v100 = float64(0.5)
	v103 = *(*float64)(unsafe.Add(mBase, _c_F_cost_subplan[0]))
	v112 = base.F64_add(base.F64_mul(base.F64_mul(v99, v100), v103), base.F64_add(base.F64_mul(v81, v100), v63))
	goto L14
L17:
	;
	v83 = float64(1e+100)
	v84 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	if base.F64_gt(v84, v83) != 0 {
		v96 = v83
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v112 = base.F64_add(v63, base.F64_div(v81, v96))
	goto L14
L19:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v84)&int64(9223372036854775807)) {
		v96 = v83
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v92 = float64(1)
	if base.F64_le(v84, v92) != 0 {
		v96 = v92
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v96 = base.F64_nearest(v84)
	goto L18
L22:
	;
	v130 = base.F64_add(v112, v127)
	v132 = v69
	goto L10
L23:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v116 = v114 - int32(348)
	goto L24
L24:
	;
	v122 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	if base.B2i32(base.Ui32(v116) < base.Ui32(int32(15)))&int32(base.Ui32(int32(_a_F_cost_subplan_0))>>(uint(v116)%32)) == int32(0) {
		v127 = v122
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v130 = v112
	v132 = base.F64_add(v69, v122)
	goto L10
}
func F_cparc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	v8 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_cparc[0]))
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v14 <= v15 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	F_createarc(m, l0, v9, v8, l2, l3)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L27
	}
L8:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v17 == int32(0) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v36 == int32(0) {
		goto L7
	} else {
		goto L19
	}
L11:
	;
	v23 = v17
	goto L12
L12:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if v29 != l3 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L7
L14:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	if v35 != 0 {
		v23 = v35
		goto L12
	} else {
		goto L18
	}
L15:
	;
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
	if v31 != v8&int32(_a_F_cparc_0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v33 == v9 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	goto L13
L19:
	;
	v42 = v36
	goto L20
L20:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v48 != l2 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L7
L22:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
	if v54 != 0 {
		v42 = v54
		goto L20
	} else {
		goto L26
	}
L23:
	;
	v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+4)))
	if v50 != v8&int32(_a_F_cparc_0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v52 == v9 {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	goto L21
L27:
	;
	goto L6
}
func F_crc32_bytea(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = int32(1)
		v12 = v7 + v11
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
		v15 = v13 & v11
		if v13 == v11 {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			if base.Ui32((v19-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v46 = int32(4)
				if v15 != 0 {
					v50 = v12
				} else {
					v50 = v7 + int32(4)
				}
				v51 = int32(1)
				if v46 == v51 {
					v94 = v50
					v95 = int32(-1)
				} else {
					v60 = v50
					v61 = int32(-1)
					v62 = int32(0)
					for {
						v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
						v67 = int32(255)
						v69 = int32(2)
						v73 = *(*int32)(unsafe.Add(mBase, uint32((v65^v61)&v67<<(uint(v69)%32))+uint32(_c_F_crc32_bytea[0])))
						v74 = int32(8)
						v76 = v73 ^ int32(base.Ui32(v61)>>(uint(v74)%32))
						v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
						v85 = *(*int32)(unsafe.Add(mBase, uint32((v76^v77)&v67<<(uint(v69)%32))+uint32(_c_F_crc32_bytea[0])))
						v88 = v85 ^ int32(base.Ui32(v76)>>(uint(v74)%32))
						v90 = v60 + v69
						v92 = v62 + v69
						if v92 != v46&int32(-2) {
							v60 = v90
							v61 = v88
							v62 = v92
							continue
						} else {
							break
						}
						break
					}
					v94 = v90
					v95 = v88
				}
				if v46&v51 != 0 {
					v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
					v107 = *(*int32)(unsafe.Add(mBase, uint32((v99^v95)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_crc32_bytea[0])))
					v111 = v107 ^ int32(base.Ui32(v95)>>(uint(int32(8))%32))
				} else {
					v111 = v95
				}
				v115 = F_Int64GetDatum(m, base.I64_extend_i32_u(v111^int32(-1)))
				mBase = m.M
				v116 = m.ExcPending
				if v116 != 0 {
					return int32(0)
				} else {
					return v115
				}
			} else {
				v41 = base.B2i32(v19 == int32(18)) << (uint(int32(4)) % 32)
				if v41 != 0 {
					v46 = v41
					if v15 != 0 {
						v50 = v12
					} else {
						v50 = v7 + int32(4)
					}
					v51 = int32(1)
					if v46 == v51 {
						v94 = v50
						v95 = int32(-1)
					} else {
						v60 = v50
						v61 = int32(-1)
						v62 = int32(0)
						for {
							v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
							v67 = int32(255)
							v69 = int32(2)
							v73 = *(*int32)(unsafe.Add(mBase, uint32((v65^v61)&v67<<(uint(v69)%32))+uint32(_c_F_crc32_bytea[0])))
							v74 = int32(8)
							v76 = v73 ^ int32(base.Ui32(v61)>>(uint(v74)%32))
							v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
							v85 = *(*int32)(unsafe.Add(mBase, uint32((v76^v77)&v67<<(uint(v69)%32))+uint32(_c_F_crc32_bytea[0])))
							v88 = v85 ^ int32(base.Ui32(v76)>>(uint(v74)%32))
							v90 = v60 + v69
							v92 = v62 + v69
							if v92 != v46&int32(-2) {
								v60 = v90
								v61 = v88
								v62 = v92
								continue
							} else {
								break
							}
							break
						}
						v94 = v90
						v95 = v88
					}
					if v46&v51 != 0 {
						v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
						v107 = *(*int32)(unsafe.Add(mBase, uint32((v99^v95)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_crc32_bytea[0])))
						v111 = v107 ^ int32(base.Ui32(v95)>>(uint(int32(8))%32))
					} else {
						v111 = v95
					}
					v115 = F_Int64GetDatum(m, base.I64_extend_i32_u(v111^int32(-1)))
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return int32(0)
					} else {
						return v115
					}
				} else {
					v43 = F_Int64GetDatum(m, int64(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						return v43
					}
				}
			}
		} else {
			v30 = int32(1)
			if v15 != 0 {
				v41 = int32(base.Ui32(v13)>>(uint(v30)%32)) - v30
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v41 = int32(base.Ui32(v34)>>(uint(int32(2))%32)) - int32(4)
			}
			if v41 != 0 {
				v46 = v41
				if v15 != 0 {
					v50 = v12
				} else {
					v50 = v7 + int32(4)
				}
				v51 = int32(1)
				if v46 == v51 {
					v94 = v50
					v95 = int32(-1)
				} else {
					v60 = v50
					v61 = int32(-1)
					v62 = int32(0)
					for {
						v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
						v67 = int32(255)
						v69 = int32(2)
						v73 = *(*int32)(unsafe.Add(mBase, uint32((v65^v61)&v67<<(uint(v69)%32))+uint32(_c_F_crc32_bytea[0])))
						v74 = int32(8)
						v76 = v73 ^ int32(base.Ui32(v61)>>(uint(v74)%32))
						v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
						v85 = *(*int32)(unsafe.Add(mBase, uint32((v76^v77)&v67<<(uint(v69)%32))+uint32(_c_F_crc32_bytea[0])))
						v88 = v85 ^ int32(base.Ui32(v76)>>(uint(v74)%32))
						v90 = v60 + v69
						v92 = v62 + v69
						if v92 != v46&int32(-2) {
							v60 = v90
							v61 = v88
							v62 = v92
							continue
						} else {
							break
						}
						break
					}
					v94 = v90
					v95 = v88
				}
				if v46&v51 != 0 {
					v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
					v107 = *(*int32)(unsafe.Add(mBase, uint32((v99^v95)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_crc32_bytea[0])))
					v111 = v107 ^ int32(base.Ui32(v95)>>(uint(int32(8))%32))
				} else {
					v111 = v95
				}
				v115 = F_Int64GetDatum(m, base.I64_extend_i32_u(v111^int32(-1)))
				mBase = m.M
				v116 = m.ExcPending
				if v116 != 0 {
					return int32(0)
				} else {
					return v115
				}
			} else {
				v43 = F_Int64GetDatum(m, int64(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					return v43
				}
			}
		}
	}
}
func F_create_edata_for_relation(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
	v14 = F_palloc0(m, int32(20))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = l0
		v19 = F_CreateExecutorState(m)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v14))) = v19
			v23 = F_palloc0(m, int32(136))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(101)
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+56))
				*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v30
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+48))
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+119)))
				*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v23)+21)) = uint8(v34)
				v40 = F_addRTEPermissionInfo(m, v9+int32(12), v23)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v23
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v23
					v47 = F_list_make1_impl(m, int32(1), v9+int32(4))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
						v51 = F_bms_make_singleton(m, int32(1))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							F_ExecInitRangeTable(m, v19, v47, v49, v51)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								v56 = F_palloc0(m, int32(216))
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v56))) = int32(388)
									*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v56
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									v63 = int32(0)
									F_InitResultRelInfo(m, v56, v61, int32(1), v63, v63)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return int32(0)
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
										v68 = F_lappend(m, v67, v56)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v68
											v72 = F_GetCurrentCommandId(m, int32(1))
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v72
												v75 = int32(_a_F_create_edata_for_relation_0)
												v77 = *(*int32)(unsafe.Add(mBase, _c_F_create_edata_for_relation[0]))
												*(*int32)(unsafe.Add(mBase, _c_F_create_edata_for_relation[0])) = v77 + int32(1)
												m.G0 = v9 + int32(16)
												return v14
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
func F_create_seqscan_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	v8 = F_palloc0(m, int32(72))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(1455993913623)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v15
		v17 = F_get_baserel_parampathinfo(m, l0, l1, l2)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)) = uint8(base.B2i32(v19 < l3))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v17
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v19
			*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = l3
			*(*uint8)(unsafe.Add(mBase, uint32(v8)+21)) = uint8(v23)
			F_cost_seqscan(m, v8, l0, l1, v17)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				return v8
			}
		}
	}
}
func F_create_tidscan_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 float64
	_ = v10
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
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
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 float64
	_ = v49
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v65 float64
	_ = v65
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 float64
	_ = v83
	var v84 int32
	_ = v84
	var v86 float64
	_ = v86
	var v87 float64
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v105 float64
	_ = v105
	var v111 int64
	_ = v111
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 float64
	_ = v148
	var v161 float64
	_ = v161
	var v177 float64
	_ = v177
	var v178 float64
	_ = v178
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 float64
	_ = v187
	var v188 int32
	_ = v188
	var v189 int64
	_ = v189
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 float64
	_ = v232
	var v233 float64
	_ = v233
	var v243 float64
	_ = v243
	var v251 float64
	_ = v251
	var v252 float64
	_ = v252
	var v254 float64
	_ = v254
	var v256 float64
	_ = v256
	var v257 float64
	_ = v257
	var v267 float64
	_ = v267
	var v275 float64
	_ = v275
	var v276 int32
	_ = v276
	var v277 float64
	_ = v277
	var v279 float64
	_ = v279
	var v280 float64
	_ = v280
	var v284 float64
	_ = v284
	var v286 float64
	_ = v286
	var v288 float64
	_ = v288
	v5 = int32(0)
	v10 = float64(0)
	v19 = F_palloc0(m, int32(80))
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
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = int64(1481763717405)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v26
	v28 = F_get_baserel_parampathinfo(m, l0, l1, l3)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v30 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+20)) = uint8(v30)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v28
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v30
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+21)) = uint8(v33)
	v40 = m.G0
	v42 = v40 - int32(32)
	m.G0 = v42
	if v28 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v48 = v28 + int32(8)
	goto L6
L5:
	;
	v48 = l1 + int32(16)
	goto L6
L6:
	;
	v49 = *(*float64)(unsafe.Add(mBase, uint32(v48)))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+32)) = v49
	if l2 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if int32(0) < v51 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v177 = v10
	v178 = v10
	goto L9
L9:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	F_get_tablespace_page_costs(m, v183, v42, int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L26
	}
L10:
	;
	v177 = v105
	v178 = v161
	goto L9
L11:
	;
	v60 = v5
	v65 = v10
	goto L14
L12:
	;
	v100 = v5
	v105 = v10
	goto L13
L13:
	;
	v111 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v42)+24)) = v111
	*(*int64)(unsafe.Add(mBase, uint32(v42)+16)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = l0
	if v100 == int32(0) {
		v161 = v10
		goto L10
	} else {
		goto L21
	}
L14:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71+v60<<(uint(int32(2))%32))))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	if v77 == int32(20) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v100 = base.B2i32(int32(0) < v90)
	v105 = v87
	goto L13
L16:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)+28))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v83 = F_estimate_array_length(m, l0, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	v86 = float64(1)
	goto L18
L18:
	;
	v87 = base.F64_add(v65, v86)
	v89 = v60 + int32(1)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v89 < v90 {
		v60 = v89
		v65 = v87
		goto L14
	} else {
		goto L20
	}
L19:
	;
	v86 = v83
	goto L18
L20:
	;
	goto L15
L21:
	;
	v125 = v5
	goto L22
L22:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v135+v125<<(uint(int32(2))%32))))
	v142 = F_cost_qual_eval_walker(m, v139, v42+int32(8))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	v148 = *(*float64)(unsafe.Add(mBase, uint32(v42)+24))
	v161 = v148
	goto L10
L24:
	;
	v145 = v125 + int32(1)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v145 < v146 {
		v125 = v145
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v187 = *(*float64)(unsafe.Add(mBase, uint32(v42)))
	if v28 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v277 = *(*float64)(unsafe.Add(mBase, uint32(v276)+24))
	v279 = *(*float64)(unsafe.Add(mBase, _c_F_create_tidscan_path[0]))
	v280 = *(*float64)(unsafe.Add(mBase, uint32(v276)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = int32(0)
	v284 = float64(0)
	v286 = base.F64_add(v280, base.F64_add(base.F64_add(v178, v275), v284))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+48)) = v286
	v288 = *(*float64)(unsafe.Add(mBase, uint32(v19)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+56)) = base.F64_add(v286, base.F64_add(base.F64_mul(v277, v288), base.F64_add(base.F64_mul(base.F64_sub(base.F64_add(v267, v279), v178), v177), base.F64_add(base.F64_mul(v187, v177), v284))))
	m.G0 = v42 + int32(32)
	return v19
L28:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v189 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v42)+24)) = v189
	*(*int64)(unsafe.Add(mBase, uint32(v42)+16)) = v189
	*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = l0
	if v188 == int32(0) {
		v243 = v10
		v251 = float64(0)
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v256 = *(*float64)(unsafe.Add(mBase, uint32(l1)+200))
	v257 = *(*float64)(unsafe.Add(mBase, uint32(l1)+192))
	v267 = v256
	v275 = v257
	goto L27
L31:
	;
	v252 = *(*float64)(unsafe.Add(mBase, uint32(l1)+200))
	v254 = *(*float64)(unsafe.Add(mBase, uint32(l1)+192))
	v267 = base.F64_add(v243, v252)
	v275 = base.F64_add(v251, v254)
	goto L27
L32:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v198 <= int32(0) {
		v243 = v10
		v251 = float64(0)
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v204 = int32(0)
	goto L34
L34:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v188)+12))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v219+v204<<(uint(int32(2))%32))))
	v226 = F_cost_qual_eval_walker(m, v223, v42+int32(8))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L36
	}
L35:
	;
	v232 = *(*float64)(unsafe.Add(mBase, uint32(v42)+24))
	v233 = *(*float64)(unsafe.Add(mBase, uint32(v42)+16))
	v243 = v232
	v251 = v233
	goto L31
L36:
	;
	v229 = v204 + int32(1)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v229 < v230 {
		v204 = v229
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
}
func F_create_upper_unique_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) int32 {
	mBase := m.M
	_ = mBase
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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 float64
	_ = v41
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v46 float64
	_ = v46
	v11 = F_palloc0(m, int32(80))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v11))) = int64(1576252997938)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v19 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v19)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v19
		*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v18
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
		if v24 == int32(1) {
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
			v29 = v27
		} else {
			v29 = int32(0)
		}
		v31 = v29 & int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(v31)
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v33
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+76)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v35
		v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v39
		v41 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
		*(*float64)(unsafe.Add(mBase, uint32(v11)+48)) = v41
		v43 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
		v44 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
		v46 = *(*float64)(unsafe.Add(mBase, _c_F_create_upper_unique_path[0]))
		*(*float64)(unsafe.Add(mBase, uint32(v11)+32)) = l3
		*(*float64)(unsafe.Add(mBase, uint32(v11)+56)) = base.F64_add(v43, base.F64_mul(base.F64_mul(v46, v44), base.F64_convert_i32_s(l2)))
		return v11
	}
}
func F_createdb_failure_callback(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v3 == int32(0) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		F_DropDatabaseBuffers(m, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			F_ForgetDatabaseSyncRequests(m, v9)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				F_UnlockSharedObject(m, int32(1262), v13, int32(1))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					F_UnlockSharedObject(m, int32(1262), v18, int32(5))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						F_remove_dbtablespaces(m, v22)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		F_UnlockSharedObject(m, int32(1262), v18, int32(5))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			F_remove_dbtablespaces(m, v22)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_cryptohash_internal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v15 = l0 << (uint(int32(2)) % 32)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_c_F_cryptohash_internal[0])))
	v20 = v18 + int32(4)
	v21 = F_palloc0(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		if v27 == int32(1) {
			v30 = int32(4)
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
			if v32&int32(254) == int32(2) {
				v41 = v30
			} else {
				v41 = base.B2i32(v32 == int32(18)) << (uint(v30) % 32)
			}
			if v32 == int32(1) {
				v44 = v30
			} else {
				v44 = v41
			}
			v57 = v44
		} else {
			v45 = int32(1)
			if v27&v45 != 0 {
				v57 = int32(base.Ui32(v27)>>(uint(v45)%32)) - v45
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v57 = int32(base.Ui32(v51)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v58 = *(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_c_F_cryptohash_internal[1])))
		v59 = F_pg_cryptohash_create(m, l0)
		mBase = m.M
		v60 = m.ExcPending
		if v60 != 0 {
			return int32(0)
		} else {
			v61 = F_pg_cryptohash_init(m, v59)
			mBase = m.M
			if int32(0) <= v61 {
				v64 = int32(1)
				if v27&v64 != 0 {
					v68 = v64
				} else {
					v68 = int32(4)
				}
				v70 = F_pg_cryptohash_update(m, v59, l1+v68, v57)
				mBase = m.M
				if v70 < int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v119 = m.ExcPending
					if v119 != 0 {
						return int32(0)
					} else {
						if v59 == int32(0) {
							v134 = int32(_a_F_cryptohash_internal_0)
						} else {
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
							if v126 == int32(1) {
								v129 = int32(_a_F_cryptohash_internal_1)
							} else {
								v129 = int32(_a_F_cryptohash_internal_2)
							}
							if v126 == int32(2) {
								v132 = int32(_a_F_cryptohash_internal_0)
							} else {
								v132 = v129
							}
							v134 = v132
						}
						*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v134
						*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v58
						F_errmsg_internal(m, int32(_a_F_cryptohash_internal_3), v12+int32(16))
						mBase = m.M
						v141 = m.ExcPending
						if v141 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_cryptohash_internal_4), int32(123), int32(_a_F_cryptohash_internal_5))
							mBase = m.M
							v146 = m.ExcPending
							if v146 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v75 = F_pg_cryptohash_final(m, v59, v21+int32(4), v18)
					mBase = m.M
					if v75 < int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v150 = m.ExcPending
						if v150 != 0 {
							return int32(0)
						} else {
							if v59 == int32(0) {
								v165 = int32(_a_F_cryptohash_internal_0)
							} else {
								v157 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
								if v157 == int32(1) {
									v160 = int32(_a_F_cryptohash_internal_1)
								} else {
									v160 = int32(_a_F_cryptohash_internal_2)
								}
								if v157 == int32(2) {
									v163 = int32(_a_F_cryptohash_internal_0)
								} else {
									v163 = v160
								}
								v165 = v163
							}
							*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v165
							*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v58
							F_errmsg_internal(m, int32(_a_F_cryptohash_internal_6), v12+int32(32))
							mBase = m.M
							v172 = m.ExcPending
							if v172 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_cryptohash_internal_4), int32(127), int32(_a_F_cryptohash_internal_5))
								mBase = m.M
								v177 = m.ExcPending
								if v177 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						F_pg_cryptohash_free(m, v59)
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v21))) = v20 << (uint(int32(2)) % 32)
							m.G0 = v12 + int32(48)
							return v21
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return int32(0)
				} else {
					if v59 == int32(0) {
						v105 = int32(_a_F_cryptohash_internal_0)
					} else {
						v97 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
						if v97 == int32(1) {
							v100 = int32(_a_F_cryptohash_internal_1)
						} else {
							v100 = int32(_a_F_cryptohash_internal_2)
						}
						if v97 == int32(2) {
							v103 = int32(_a_F_cryptohash_internal_0)
						} else {
							v103 = v100
						}
						v105 = v103
					}
					*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v105
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = v58
					F_errmsg_internal(m, int32(_a_F_cryptohash_internal_7), v12)
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_cryptohash_internal_4), int32(120), int32(_a_F_cryptohash_internal_5))
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
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
func F_cstring_recv(m *base.Module, l0 int32) int32 {
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
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v13 = F_pq_getmsgtext(m, v7, v8-v9, v5+int32(12))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		m.G0 = v5 + int32(16)
		return v13
	}
}
func F_currtid_byrelname(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v9 = F_textToQualifiedNameList(m, v4)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = F_makeRangeVarFromNameList(m, v9)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				v14 = F_table_openrv(m, v11, int32(1))
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return int32(0)
				} else {
					v16 = F_currtid_internal(m, v14, v8)
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return int32(0)
					} else {
						F_sequence_close(m, v14, int32(1))
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
							return int32(0)
						} else {
							return v16
						}
					}
				}
			}
		}
	}
}
