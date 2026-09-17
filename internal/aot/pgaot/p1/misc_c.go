package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
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
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
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
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	v3 = int32(0)
	if l0 == v3 {
		v157 = v3
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v172
	return int32(0)
L2:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v159 + int32(1)
	v165 = F_query_tree_walker_impl(m, l0, int32(1049), l1, int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L8
	} else {
		goto L72
	}
L3:
	;
	return v157
L4:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v9 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v10 = m.T0[v9].(func(*base.Module, int32, int32) int32)(m, l0, l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v14 - int32(58) {
	case 0:
		goto L16
	case 1, 2, 3, 4:
		v113 = v14
		goto L12
	case 5:
		goto L15
	case 6:
		goto L14
	default:
		goto L17
	}
L8:
	;
	return int32(0)
L9:
	;
	if v10 != 0 {
		v157 = v3
		goto L3
	} else {
		goto L10
	}
L10:
	;
	goto L7
L11:
	;
	v152 = F_expression_tree_walker_impl(m, l0, int32(1049), l1)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L8
	} else {
		goto L71
	}
L12:
	;
	if v113 == int32(67) {
		goto L2
	} else {
		goto L56
	}
L13:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v69 != v70 {
		goto L11
	} else {
		goto L39
	}
L14:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v63 != 0 {
		goto L11
	} else {
		goto L37
	}
L15:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v59 != 0 {
		v157 = v3
		goto L3
	} else {
		goto L35
	}
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v55 != 0 {
		v157 = v3
		goto L3
	} else {
		goto L33
	}
L17:
	;
	if v14 == int32(319) {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	if v14 != int32(6) {
		v113 = v14
		goto L12
	} else {
		goto L19
	}
L19:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v21 != v22 {
		v157 = v3
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v25 == v26 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v24
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v30 = v29
	goto L23
L22:
	;
	v30 = v25
	goto L23
L23:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v30 < int32(0) {
		v46 = v31
		goto L24
	} else {
		goto L25
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v46
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v48 != v49 {
		v157 = v3
		goto L3
	} else {
		goto L32
	}
L25:
	;
	v34 = F_bms_is_member(m, v30, v31)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	if v34 == int32(0) {
		v46 = v31
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v38 = F_bms_copy(m, v31)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L8
	} else {
		goto L28
	}
L28:
	;
	v40 = F_bms_del_member(m, v38, v30)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L8
	} else {
		goto L29
	}
L29:
	;
	if v24 < int32(0) {
		v46 = v40
		goto L24
	} else {
		goto L30
	}
L30:
	;
	v44 = F_bms_add_member(m, v40, v24)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L8
	} else {
		goto L31
	}
L31:
	;
	v46 = v44
	goto L24
L32:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v51
	return int32(0)
L33:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v56 != v57 {
		v157 = v3
		goto L3
	} else {
		goto L34
	}
L34:
	;
	goto L1
L35:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v60 != v61 {
		v157 = v3
		goto L3
	} else {
		goto L36
	}
L36:
	;
	goto L1
L37:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v64 != v65 {
		goto L11
	} else {
		goto L38
	}
L38:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v67
	goto L11
L39:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v73 < int32(0) {
		v89 = v72
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v89
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v93 < int32(0) {
		v109 = v92
		goto L48
	} else {
		goto L49
	}
L41:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v77 = F_bms_is_member(m, v73, v72)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L8
	} else {
		goto L42
	}
L42:
	;
	if v77 == int32(0) {
		v89 = v72
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v81 = F_bms_copy(m, v72)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L8
	} else {
		goto L44
	}
L44:
	;
	v83 = F_bms_del_member(m, v81, v73)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L8
	} else {
		goto L45
	}
L45:
	;
	if v76 < int32(0) {
		v89 = v83
		goto L40
	} else {
		goto L46
	}
L46:
	;
	v87 = F_bms_add_member(m, v83, v76)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L8
	} else {
		goto L47
	}
L47:
	;
	v89 = v87
	goto L40
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v109
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v113 = v112
	goto L12
L49:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v97 = F_bms_is_member(m, v93, v92)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L8
	} else {
		goto L50
	}
L50:
	;
	if v97 == int32(0) {
		v109 = v92
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v101 = F_bms_copy(m, v92)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L8
	} else {
		goto L52
	}
L52:
	;
	v103 = F_bms_del_member(m, v101, v93)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L8
	} else {
		goto L53
	}
L53:
	;
	if v96 < int32(0) {
		v109 = v103
		goto L48
	} else {
		goto L54
	}
L54:
	;
	v107 = F_bms_add_member(m, v103, v96)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L8
	} else {
		goto L55
	}
L55:
	;
	v109 = v107
	goto L48
L56:
	;
	if v113 != int32(322) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	if v113 != int32(374) {
		goto L11
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v136 != 0 {
		goto L11
	} else {
		goto L66
	}
L60:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v122 != 0 {
		v157 = v3
		goto L3
	} else {
		goto L61
	}
L61:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v123 == v124 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v126
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v129 = v128
	goto L64
L63:
	;
	v129 = v123
	goto L64
L64:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v129 != v130 {
		v157 = v3
		goto L3
	} else {
		goto L65
	}
L65:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v132
	return int32(0)
L66:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v137 == v138 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v140
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v143 = v142
	goto L69
L68:
	;
	v143 = v137
	goto L69
L69:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v143 != v144 {
		goto L11
	} else {
		goto L70
	}
L70:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v146
	goto L11
L71:
	;
	v157 = v152
	goto L3
L72:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v167 - int32(1)
	return v165
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
	if v35 == v48 {
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
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
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int64
	_ = v246
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	v14 = m.G0
	v16 = v14 - int32(96)
	m.G0 = v16
	v19 = v16 + int32(48)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ScanKeyInit(m, v19, int32(2), int32(3), int32(184), v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v28 = F_table_open(m, int32(2620), int32(3))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = int32(1)
	v34 = F_systable_beginscan(m, v28, int32(2701), v31, int32(0), v31, v19)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_CloneRowTriggersToPartition[0]))
	v42 = F_AllocSetContextCreateInternal(m, v37, int32(_a_F_CloneRowTriggersToPartition_0), int32(0), int32(1024), int32(_a_F_CloneRowTriggersToPartition_1))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v44 = F_systable_getnext(m, v34)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L62
	}
L7:
	;
	if v44 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v50 = v44
	goto L11
L9:
	;
	goto L10
L10:
	;
	F_MemoryContextDelete(m, v42)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L59
	}
L11:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+22)))
	v61 = v59 + v60
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+80)))
	if v62&int32(1) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L10
L13:
	;
	v302 = F_systable_getnext(m, v34)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L57
	}
L14:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+83)))
	if v67 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	switch v62 & int32(66) {
	case 0, 2:
		goto L16
	default:
		goto L17
	}
L16:
	;
	v85 = int32(0)
	v86 = int32(_a_F_CloneRowTriggersToPartition_2)
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_CloneRowTriggersToPartition[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_CloneRowTriggersToPartition[0])) = v42
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v28)+52))
	v94 = F_heap_getattr_6(m, v50, int32(17), v91, v16+int32(47))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L21
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v61 + int32(12)
	F_errmsg_internal(m, int32(_a_F_CloneRowTriggersToPartition_3), v16)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_CloneRowTriggersToPartition_4), int32(_a_F_CloneRowTriggersToPartition_5), int32(_a_F_CloneRowTriggersToPartition_6))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
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
	v96 = int32(0)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+47)))
	if v97 == v96 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v100 = F_text_to_cstring(m, v94)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	v110 = v96
	goto L24
L24:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v61)+116))
	if int32(0) < v111 {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v102 = F_stringToNode(m, v100)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v105 = F_map_partition_varattnos(m, v102, int32(1), l1, l0)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v108 = F_map_partition_varattnos(m, v105, int32(2), l1, l0)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v110 = v108
	goto L24
L29:
	;
	v119 = int32(0)
	v124 = v85
	goto L32
L30:
	;
	v161 = v85
	goto L31
L31:
	;
	v167 = int32(0)
	v168 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+98)))
	if v168 <= v167 {
		v230 = v167
		goto L38
	} else {
		goto L39
	}
L32:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v138 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61+int32(124)+v119<<(uint(int32(1))%32)))))
	v144 = F_pstrdup(m, v130+v131<<(uint(int32(4))%32)+v138*int32(100)-int32(76))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L34
	}
L33:
	;
	v161 = v148
	goto L31
L34:
	;
	v146 = F_makeString(m, v144)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v148 = F_lappend(m, v124, v146)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v151 = v119 + int32(1)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v61)+116))
	if v151 < v152 {
		v119 = v151
		v124 = v148
		goto L32
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	v236 = F_palloc0(m, int32(52))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L54
	}
L39:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v28)+52))
	v175 = F_heap_getattr_6(m, v50, int32(16), v172, v16+int32(47))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+47)))
	if v177 == int32(1) {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	v180 = F_pg_detoast_datum_packed(m, v175)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	v183 = F_pg_detoast_datum_packed(m, v175)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v185 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+98)))
	if v185 <= int32(0) {
		v230 = v167
		goto L38
	} else {
		goto L44
	}
L44:
	;
	v188 = int32(1)
	if v182&v188 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v192 = v188
	goto L47
L46:
	;
	v192 = int32(4)
	goto L47
L47:
	;
	v197 = v183 + v192
	v199 = int32(0)
	v203 = v167
	goto L48
L48:
	;
	v208 = F_pstrdup(m, v197)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L50
	}
L49:
	;
	v230 = v212
	goto L38
L50:
	;
	v210 = F_makeString(m, v208)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v212 = F_lappend(m, v203, v210)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v214 = F_strlen(m, v197)
	mBase = m.M
	v216 = int32(1)
	v219 = v199 + v216
	v220 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+98)))
	if v219 < v220 {
		v197 = v214 + v197 + v216
		v199 = v219
		v203 = v212
		goto L48
	} else {
		goto L53
	}
L53:
	;
	goto L49
L54:
	;
	v238 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v236)+4)) = uint8(v238)
	*(*int32)(unsafe.Add(mBase, uint32(v236))) = int32(181)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v61)+92))
	v243 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v236)+24)) = uint8(v243)
	*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v230
	v246 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v236)+12)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v236)+8)) = v61 + int32(12)
	*(*uint8)(unsafe.Add(mBase, uint32(v236)+5)) = uint8(base.B2i32(v242 != v238))
	v254 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+80)))
	v256 = v254 & int32(66)
	*(*uint16)(unsafe.Add(mBase, uint32(v236)+26)) = uint16(v256)
	v258 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+80)))
	*(*int64)(unsafe.Add(mBase, uint32(v236)+36)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v236)+32)) = v161
	v263 = v258 & int32(60)
	*(*uint16)(unsafe.Add(mBase, uint32(v236)+28)) = uint16(v263)
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v236)+44)) = uint8(v265)
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+97)))
	*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v238
	*(*uint8)(unsafe.Add(mBase, uint32(v236)+45)) = uint8(v267)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v61)+84))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v61)+76))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v282 = int32(*(*int8)(unsafe.Add(mBase, uint32(v61)+82)))
	F_CreateTriggerFiringOn(m, v16+int32(32), v236, v238, v274, v275, v238, v238, v278, v279, v110, v238, v243, v282)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CloneRowTriggersToPartition[0])) = v87
	F_MemoryContextReset(m, v42)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L13
L57:
	;
	if v302 != 0 {
		v50 = v302
		goto L11
	} else {
		goto L58
	}
L58:
	;
	goto L12
L59:
	;
	F_systable_endscan(m, v34)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_relation_close(m, v28, int32(3))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	m.G0 = v16 + int32(96)
	return
L62:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v61 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v331 + int32(4)
	F_errmsg_internal(m, int32(_a_F_CloneRowTriggersToPartition_7), v16+int32(16))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_CloneRowTriggersToPartition_4), int32(_a_F_CloneRowTriggersToPartition_8), int32(_a_F_CloneRowTriggersToPartition_6))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v62 int32
	_ = v62
	var v65 float64
	_ = v65
	var v66 int32
	_ = v66
	var v70 float64
	_ = v70
	v14 = l3 - int32(1)
	v18 = v14
	v20 = int32(-1)
	goto L1
L1:
	;
	v30 = base.I32_div_s(v18+v20+int32(1), int32(2))
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
	v47 = v18
	goto L10
L9:
	;
	v47 = v30 - int32(1)
	goto L10
L10:
	;
	if v44 < v47 {
		v18 = v47
		v20 = v44
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
	if base.B2i32(v44 < int32(0))|base.B2i32(v14 <= v44) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v70 = v55
	goto L17
L16:
	;
	v62 = l2 + v44<<(uint(int32(3))%32)
	v65 = F_get_position(m, l0, l1, v62, v62+int32(8))
	v66 = m.ExcPending
	if v66 != 0 {
		goto L3
	} else {
		goto L18
	}
L17:
	;
	return v70
L18:
	;
	v70 = base.F64_add(v55, base.F64_div(v65, v54))
	goto L17
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
												base.MemoryFill(m, v8, v132, int32(20))
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
										base.MemoryFill(m, v8, v132, int32(20))
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
													base.MemoryFill(m, v8, v132, int32(20))
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
														base.MemoryFill(m, v8, v132, int32(20))
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
	var v7 int32
	_ = v7
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	v2 = int32(0)
	if base.Ui32(l0) <= base.Ui32(int32(1415)) {
		v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_case_index[0]))))
		v134 = v7
	} else {
		if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_0)) {
			if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_1)) {
				if base.Ui32(l0-int32(_a_F_case_index_2)) <= base.Ui32(int32(95)) {
					v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_case_index[1]))))
					v134 = v20
				} else {
					if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_3)) {
						v134 = v2
					} else {
						if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_4)) {
							v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_case_index[2]))))
							v134 = v29
						} else {
							if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_5)) {
								v134 = v2
							} else {
								v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_case_index[3]))))
								v134 = v36
							}
						}
					}
				}
			} else {
				if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_6)) {
					v134 = v2
				} else {
					if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_7)) {
						if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_8)) {
							v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_case_index[4]))))
							v134 = v47
						} else {
							if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_9)) {
								v134 = v2
							} else {
								v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_case_index[5]))))
								v134 = v54
							}
						}
					} else {
						if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_10)) {
							v134 = v2
						} else {
							if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_11)) {
								v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_case_index[6]))))
								v134 = v63
							} else {
								if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_12)) {
									v134 = v2
								} else {
									v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_case_index[7]))))
									v134 = v70
								}
							}
						}
					}
				}
			}
		} else {
			if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_13)) {
				v134 = v2
			} else {
				if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_14)) {
					if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_15)) {
						if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_16)) {
							v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_case_index[8]))))
							v134 = v83
						} else {
							if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_17)) {
								v134 = v2
							} else {
								v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_case_index[9]))))
								v134 = v90
							}
						}
					} else {
						if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_18)) {
							v134 = v2
						} else {
							if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_19)) {
								v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_case_index[10]))))
								v134 = v99
							} else {
								if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_20)) {
									v134 = v2
								} else {
									v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_case_index[11]))))
									v134 = v106
								}
							}
						}
					}
				} else {
					if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_21)) {
						v134 = v2
					} else {
						if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_22)) {
							if base.Ui32(l0) <= base.Ui32(int32(_a_F_case_index_23)) {
								v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_case_index[12]))))
								v134 = v117
							} else {
								if base.Ui32(l0) < base.Ui32(int32(_a_F_case_index_24)) {
									v134 = v2
								} else {
									v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_case_index[13]))))
									v134 = v124
								}
							}
						} else {
							if base.Ui32(int32(67)) < base.Ui32(l0-int32(_a_F_case_index_25)) {
								v134 = v2
							} else {
								v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_case_index[14]))))
								v134 = v133
							}
						}
					}
				}
			}
		}
	}
	return v134
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
	var v81 int32
	_ = v81
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
	var v195 int32
	_ = v195
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
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
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
	v81 = v62
	goto L23
L22:
	;
	v116 = int32(1)
	goto L18
L23:
	;
	if v81 == v75 {
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
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88+v81))))
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
	v107 = v81 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v107
	v81 = v107
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
	v195 = v176
	goto L54
L53:
	;
	v230 = int32(1)
	goto L49
L54:
	;
	if v195 == v189 {
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
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202+v195))))
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
	v221 = v195 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v221
	v195 = v221
	goto L54
L63:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v233))) = v234 + v230
	goto L1
L64:
	;
	return v453
L65:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v278
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v278
	v283 = F_find_among_b(m, l0, int32(_a_F_catalan_ISO_8859_1_stem_0), int32(200))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L68
	} else {
		goto L76
	}
L66:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245+v243))))
	if base.B2i32(v247&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v247)%32)&int32(_a_F_catalan_ISO_8859_1_stem_1) == int32(0)) != 0 {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v261 = F_find_among_b(m, l0, int32(_a_F_catalan_ISO_8859_1_stem_2), int32(39))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	return int32(0)
L69:
	;
	if v261 == int32(0) {
		goto L65
	} else {
		goto L70
	}
L70:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v267
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	if v267 < v270 {
		goto L65
	} else {
		goto L71
	}
L71:
	;
	v272 = F_slice_del(m, l0)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L68
	} else {
		goto L72
	}
L72:
	;
	if v272 < int32(0) {
		v453 = v272
		goto L64
	} else {
		goto L73
	}
L73:
	;
	goto L65
L74:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v362
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v362
	v367 = F_find_among_b(m, l0, int32(_a_F_catalan_ISO_8859_1_stem_3), int32(22))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L68
	} else {
		goto L109
	}
L75:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v333
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v333
	v338 = F_find_among_b(m, l0, int32(_a_F_catalan_ISO_8859_1_stem_4), int32(283))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L68
	} else {
		goto L98
	}
L76:
	;
	if v283 == int32(0) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v287
	switch v283 - int32(1) {
	case 0:
		goto L82
	case 1:
		goto L81
	case 2:
		goto L80
	case 3:
		goto L79
	case 4:
		goto L78
	default:
		goto L74
	}
L78:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)+4))
	if v287 < v324 {
		goto L75
	} else {
		goto L95
	}
L79:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)))
	if v287 < v315 {
		goto L75
	} else {
		goto L92
	}
L80:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	if v287 < v306 {
		goto L75
	} else {
		goto L89
	}
L81:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)))
	if v287 < v299 {
		goto L75
	} else {
		goto L86
	}
L82:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)+4))
	if v287 < v292 {
		goto L75
	} else {
		goto L83
	}
L83:
	;
	v294 = F_slice_del(m, l0)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L68
	} else {
		goto L84
	}
L84:
	;
	if int32(0) <= v294 {
		goto L74
	} else {
		goto L85
	}
L85:
	;
	v453 = v294
	goto L64
L86:
	;
	v301 = F_slice_del(m, l0)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L68
	} else {
		goto L87
	}
L87:
	;
	if int32(0) <= v301 {
		goto L74
	} else {
		goto L88
	}
L88:
	;
	v453 = v301
	goto L64
L89:
	;
	v310 = F_slice_from_s(m, l0, int32(3), int32(_a_F_catalan_ISO_8859_1_stem_5))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L68
	} else {
		goto L90
	}
L90:
	;
	if int32(0) <= v310 {
		goto L74
	} else {
		goto L91
	}
L91:
	;
	v453 = v310
	goto L64
L92:
	;
	v319 = F_slice_from_s(m, l0, int32(2), int32(_a_F_catalan_ISO_8859_1_stem_6))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L68
	} else {
		goto L93
	}
L93:
	;
	if int32(0) <= v319 {
		goto L74
	} else {
		goto L94
	}
L94:
	;
	v453 = v319
	goto L64
L95:
	;
	v328 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_ISO_8859_1_stem_7))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L68
	} else {
		goto L96
	}
L96:
	;
	if int32(0) <= v328 {
		goto L74
	} else {
		goto L97
	}
L97:
	;
	v453 = v328
	goto L64
L98:
	;
	if v338 == int32(0) {
		goto L74
	} else {
		goto L99
	}
L99:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v342
	switch v338 - int32(1) {
	case 0:
		goto L101
	case 1:
		goto L100
	default:
		goto L74
	}
L100:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v353)))
	if v342 < v354 {
		goto L74
	} else {
		goto L105
	}
L101:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v346)+4))
	if v342 < v347 {
		goto L74
	} else {
		goto L102
	}
L102:
	;
	v349 = F_slice_del(m, l0)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L68
	} else {
		goto L103
	}
L103:
	;
	if int32(0) <= v349 {
		goto L74
	} else {
		goto L104
	}
L104:
	;
	v453 = v349
	goto L64
L105:
	;
	v356 = F_slice_del(m, l0)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L68
	} else {
		goto L106
	}
L106:
	;
	if v356 < int32(0) {
		v453 = v356
		goto L64
	} else {
		goto L107
	}
L107:
	;
	goto L74
L108:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v392
	v395 = v392
	goto L119
L109:
	;
	if v367 == int32(0) {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v371
	switch v367 - int32(1) {
	case 0:
		goto L112
	case 1:
		goto L111
	default:
		goto L108
	}
L111:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v382)+4))
	if v371 < v383 {
		goto L108
	} else {
		goto L116
	}
L112:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v375)+4))
	if v371 < v376 {
		goto L108
	} else {
		goto L113
	}
L113:
	;
	v378 = F_slice_del(m, l0)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L68
	} else {
		goto L114
	}
L114:
	;
	if int32(0) <= v378 {
		goto L108
	} else {
		goto L115
	}
L115:
	;
	v453 = v378
	goto L64
L116:
	;
	v387 = F_slice_from_s(m, l0, int32(2), int32(_a_F_catalan_ISO_8859_1_stem_8))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L68
	} else {
		goto L117
	}
L117:
	;
	if v387 < int32(0) {
		v453 = v387
		goto L64
	} else {
		goto L118
	}
L118:
	;
	goto L108
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v395
	v401 = F_find_among(m, l0, int32(_a_F_catalan_ISO_8859_1_stem_9), int32(13))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L68
	} else {
		goto L122
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v392
	v453 = int32(1)
	goto L64
L121:
	;
	goto L120
L122:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v403
	switch v401 - int32(1) {
	case 0:
		goto L130
	case 1:
		goto L129
	case 2:
		goto L128
	case 3:
		goto L127
	case 4:
		goto L126
	case 5:
		goto L125
	case 6:
		goto L124
	default:
		goto L123
	}
L123:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v395 = v449
	goto L119
L124:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v443 <= v403 {
		goto L121
	} else {
		goto L143
	}
L125:
	;
	v439 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_ISO_8859_1_stem_10))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L68
	} else {
		goto L141
	}
L126:
	;
	v433 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_ISO_8859_1_stem_11))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L68
	} else {
		goto L139
	}
L127:
	;
	v427 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_ISO_8859_1_stem_12))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L68
	} else {
		goto L137
	}
L128:
	;
	v421 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_ISO_8859_1_stem_13))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L68
	} else {
		goto L135
	}
L129:
	;
	v415 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_ISO_8859_1_stem_14))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L68
	} else {
		goto L133
	}
L130:
	;
	v409 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_ISO_8859_1_stem_15))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L68
	} else {
		goto L131
	}
L131:
	;
	if int32(0) <= v409 {
		goto L123
	} else {
		goto L132
	}
L132:
	;
	v453 = v409
	goto L64
L133:
	;
	if int32(0) <= v415 {
		goto L123
	} else {
		goto L134
	}
L134:
	;
	v453 = v415
	goto L64
L135:
	;
	if int32(0) <= v421 {
		goto L123
	} else {
		goto L136
	}
L136:
	;
	v453 = v421
	goto L64
L137:
	;
	if int32(0) <= v427 {
		goto L123
	} else {
		goto L138
	}
L138:
	;
	v453 = v427
	goto L64
L139:
	;
	if int32(0) <= v433 {
		goto L123
	} else {
		goto L140
	}
L140:
	;
	v453 = v433
	goto L64
L141:
	;
	if int32(0) <= v439 {
		goto L123
	} else {
		goto L142
	}
L142:
	;
	v453 = v439
	goto L64
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v403 + int32(1)
	goto L123
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
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
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
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v784 int32
	_ = v784
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
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
	return v796
L105:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v545
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v545
	v550 = F_find_among_b(m, l0, int32(_a_F_catalan_UTF_8_stem_2), int32(200))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L108
	} else {
		goto L116
	}
L106:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512+v510))))
	if base.B2i32(v514&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v514)%32)&int32(_a_F_catalan_UTF_8_stem_3) == int32(0)) != 0 {
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v528 = F_find_among_b(m, l0, int32(_a_F_catalan_UTF_8_stem_4), int32(39))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	return int32(0)
L109:
	;
	if v528 == int32(0) {
		goto L105
	} else {
		goto L110
	}
L110:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v534
	v536 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v536)+4))
	if v534 < v537 {
		goto L105
	} else {
		goto L111
	}
L111:
	;
	v539 = F_slice_del(m, l0)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L108
	} else {
		goto L112
	}
L112:
	;
	if v539 < int32(0) {
		v796 = v539
		goto L104
	} else {
		goto L113
	}
L113:
	;
	goto L105
L114:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v629
	v634 = F_find_among_b(m, l0, int32(_a_F_catalan_UTF_8_stem_5), int32(22))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L108
	} else {
		goto L149
	}
L115:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v600
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v600
	v605 = F_find_among_b(m, l0, int32(_a_F_catalan_UTF_8_stem_6), int32(283))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L108
	} else {
		goto L138
	}
L116:
	;
	if v550 == int32(0) {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v554
	switch v550 - int32(1) {
	case 0:
		goto L122
	case 1:
		goto L121
	case 2:
		goto L120
	case 3:
		goto L119
	case 4:
		goto L118
	default:
		goto L114
	}
L118:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v590)+4))
	if v554 < v591 {
		goto L115
	} else {
		goto L135
	}
L119:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v581)))
	if v554 < v582 {
		goto L115
	} else {
		goto L132
	}
L120:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v572)))
	if v554 < v573 {
		goto L115
	} else {
		goto L129
	}
L121:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v565)))
	if v554 < v566 {
		goto L115
	} else {
		goto L126
	}
L122:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v558)+4))
	if v554 < v559 {
		goto L115
	} else {
		goto L123
	}
L123:
	;
	v561 = F_slice_del(m, l0)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L108
	} else {
		goto L124
	}
L124:
	;
	if int32(0) <= v561 {
		goto L114
	} else {
		goto L125
	}
L125:
	;
	v796 = v561
	goto L104
L126:
	;
	v568 = F_slice_del(m, l0)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L108
	} else {
		goto L127
	}
L127:
	;
	if int32(0) <= v568 {
		goto L114
	} else {
		goto L128
	}
L128:
	;
	v796 = v568
	goto L104
L129:
	;
	v577 = F_slice_from_s(m, l0, int32(3), int32(_a_F_catalan_UTF_8_stem_7))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L108
	} else {
		goto L130
	}
L130:
	;
	if int32(0) <= v577 {
		goto L114
	} else {
		goto L131
	}
L131:
	;
	v796 = v577
	goto L104
L132:
	;
	v586 = F_slice_from_s(m, l0, int32(2), int32(_a_F_catalan_UTF_8_stem_8))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L108
	} else {
		goto L133
	}
L133:
	;
	if int32(0) <= v586 {
		goto L114
	} else {
		goto L134
	}
L134:
	;
	v796 = v586
	goto L104
L135:
	;
	v595 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_UTF_8_stem_9))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L108
	} else {
		goto L136
	}
L136:
	;
	if int32(0) <= v595 {
		goto L114
	} else {
		goto L137
	}
L137:
	;
	v796 = v595
	goto L104
L138:
	;
	if v605 == int32(0) {
		goto L114
	} else {
		goto L139
	}
L139:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v609
	switch v605 - int32(1) {
	case 0:
		goto L141
	case 1:
		goto L140
	default:
		goto L114
	}
L140:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v620)))
	if v609 < v621 {
		goto L114
	} else {
		goto L145
	}
L141:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v613)+4))
	if v609 < v614 {
		goto L114
	} else {
		goto L142
	}
L142:
	;
	v616 = F_slice_del(m, l0)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L108
	} else {
		goto L143
	}
L143:
	;
	if int32(0) <= v616 {
		goto L114
	} else {
		goto L144
	}
L144:
	;
	v796 = v616
	goto L104
L145:
	;
	v623 = F_slice_del(m, l0)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L108
	} else {
		goto L146
	}
L146:
	;
	if v623 < int32(0) {
		v796 = v623
		goto L104
	} else {
		goto L147
	}
L147:
	;
	goto L114
L148:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v659
	v661 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v663 = v659
	v666 = v661
	goto L159
L149:
	;
	if v634 == int32(0) {
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v638
	switch v634 - int32(1) {
	case 0:
		goto L152
	case 1:
		goto L151
	default:
		goto L148
	}
L151:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v649)+4))
	if v638 < v650 {
		goto L148
	} else {
		goto L156
	}
L152:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v642)+4))
	if v638 < v643 {
		goto L148
	} else {
		goto L153
	}
L153:
	;
	v645 = F_slice_del(m, l0)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L108
	} else {
		goto L154
	}
L154:
	;
	if int32(0) <= v645 {
		goto L148
	} else {
		goto L155
	}
L155:
	;
	v796 = v645
	goto L104
L156:
	;
	v654 = F_slice_from_s(m, l0, int32(2), int32(_a_F_catalan_UTF_8_stem_10))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L108
	} else {
		goto L157
	}
L157:
	;
	if v654 < int32(0) {
		v796 = v654
		goto L104
	} else {
		goto L158
	}
L158:
	;
	goto L148
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v663
	v669 = v663 + int32(1)
	if v666 <= v669 {
		goto L165
	} else {
		goto L166
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v659
	v796 = int32(1)
	goto L104
L161:
	;
	goto L160
L162:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v792 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v663 = v792
	v666 = v791
	goto L159
L163:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L191
L164:
	;
	v686 = F_find_among(m, l0, int32(_a_F_catalan_UTF_8_stem_11), int32(13))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L108
	} else {
		goto L169
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v663
	v729 = v663
	v731 = v666
	goto L163
L166:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v671+v669))))
	if v673&int32(224) != int32(160) {
		goto L165
	} else {
		goto L167
	}
L167:
	;
	if int32(1)<<(uint(v673)%32)&int32(344765187) != 0 {
		goto L164
	} else {
		goto L168
	}
L168:
	;
	goto L165
L169:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v688
	switch v686 - int32(1) {
	case 0:
		goto L176
	case 1:
		goto L175
	case 2:
		goto L174
	case 3:
		goto L173
	case 4:
		goto L172
	case 5:
		goto L171
	case 6:
		goto L170
	default:
		goto L162
	}
L170:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v729 = v688
	v731 = v728
	goto L163
L171:
	;
	v724 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_UTF_8_stem_12))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L108
	} else {
		goto L187
	}
L172:
	;
	v718 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_UTF_8_stem_13))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L108
	} else {
		goto L185
	}
L173:
	;
	v712 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_UTF_8_stem_14))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L108
	} else {
		goto L183
	}
L174:
	;
	v706 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_UTF_8_stem_15))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L108
	} else {
		goto L181
	}
L175:
	;
	v700 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_UTF_8_stem_16))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L108
	} else {
		goto L179
	}
L176:
	;
	v694 = F_slice_from_s(m, l0, int32(1), int32(_a_F_catalan_UTF_8_stem_17))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L108
	} else {
		goto L177
	}
L177:
	;
	if int32(0) <= v694 {
		goto L162
	} else {
		goto L178
	}
L178:
	;
	v796 = v694
	goto L104
L179:
	;
	if int32(0) <= v700 {
		goto L162
	} else {
		goto L180
	}
L180:
	;
	v796 = v700
	goto L104
L181:
	;
	if int32(0) <= v706 {
		goto L162
	} else {
		goto L182
	}
L182:
	;
	v796 = v706
	goto L104
L183:
	;
	if int32(0) <= v712 {
		goto L162
	} else {
		goto L184
	}
L184:
	;
	v796 = v712
	goto L104
L185:
	;
	if int32(0) <= v718 {
		goto L162
	} else {
		goto L186
	}
L186:
	;
	v796 = v718
	goto L104
L187:
	;
	if int32(0) <= v724 {
		goto L162
	} else {
		goto L188
	}
L188:
	;
	v796 = v724
	goto L104
L189:
	;
	if v784 < int32(0) {
		goto L161
	} else {
		goto L209
	}
L191:
	;
	goto L192
L192:
	;
	goto L193
L193:
	;
	v739 = v729
	v741 = int32(1)
	goto L196
L195:
	;
	v784 = v769
	goto L189
L196:
	;
	if v731 <= v739 {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	goto L195
L198:
	;
	v784 = int32(-1)
	goto L189
L199:
	;
	goto L200
L200:
	;
	v746 = v739 + int32(1)
	v748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v732+v739))))
	if base.Ui32(v748) < base.Ui32(int32(192)) {
		v769 = v746
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v770 = int32(1)
	if v770 < v741 {
		v739 = v769
		v741 = v741 - v770
		goto L196
	} else {
		goto L208
	}
L202:
	;
	if v731 <= v746 {
		v769 = v746
		goto L201
	} else {
		goto L203
	}
L203:
	;
	v755 = v746
	goto L204
L204:
	;
	v758 = int32(*(*int8)(unsafe.Add(mBase, uint32(v732+v755))))
	if int32(-65) < v758 {
		v769 = v755
		goto L201
	} else {
		goto L206
	}
L205:
	;
	v769 = v731
	goto L201
L206:
	;
	v762 = v755 + int32(1)
	if v762 != v731 {
		v755 = v762
		goto L204
	} else {
		goto L207
	}
L207:
	;
	goto L205
L208:
	;
	goto L197
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v784
	goto L162
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
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
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
	v62 = l0 + int32(20)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v65 = v25
	v66 = v26
	v67 = v27
	v70 = v63
	goto L20
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
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v52 = v45
	goto L15
L15:
	;
	if int32(0) < v44 {
		v25 = v25 + v42
		v26 = v44
		v27 = v27 + v42
		v29 = v52
		goto L3
	} else {
		goto L19
	}
L16:
	;
	base.MemoryCopy(m, l0+int32(20), l0+int32(84), v45)
	goto L18
L17:
	;
	goto L18
L18:
	;
	v49 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v49
	v52 = v49
	goto L15
L19:
	;
	goto L1
L20:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)+16))
	v81 = m.T0[v80].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v76, int32(0), v62, v70, l0+int32(52), v15+int32(12))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L11
	} else {
		goto L22
	}
L21:
	;
	goto L1
L22:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v83 <= int32(4) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v83 + int32(1)
	goto L25
L24:
	;
	goto L25
L25:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v66 < v89 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v91 = v66
	goto L28
L27:
	;
	v91 = v89
	goto L28
L28:
	;
	v92 = m.T0[l4].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v65, v91, v67)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L11
	} else {
		goto L29
	}
L29:
	;
	v94 = v66 - v92
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v95 == v96 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if v95 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	if int32(0) < v94 {
		v65 = v65 + v92
		v66 = v94
		v67 = v67 + v92
		v70 = v96
		goto L20
	} else {
		goto L36
	}
L33:
	;
	base.MemoryCopy(m, v62, l0+int32(84), v95)
	goto L35
L34:
	;
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	goto L32
L36:
	;
	goto L21
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = F_pg_clean_ascii(m, v4, int32(2))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 != 0 {
			v11 = F_guc_strdup(m, int32(15), v6)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				if v11 == int32(0) {
					F_pfree(m, v6)
					mBase = m.M
					v16 = m.ExcPending
					if v16 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					F_bms_free(m, v19)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						F_pfree(m, v6)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v11
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
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
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+188))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	m.T0[v59].(func(*base.Module, int32))(m, v20)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
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
		v56 = v2
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
	v56 = v2
	goto L1
L9:
	;
	v48 = F_heap_getnext(m, v20)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
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
	v56 = int32(1)
	goto L1
L15:
	;
	if v48 != 0 {
		v27 = v48
		goto L7
	} else {
		goto L16
	}
L16:
	;
	goto L8
L17:
	;
	F_relation_close(m, v14, int32(1))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	m.G0 = v10 + int32(96)
	return v56
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
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
			if base.B2i32(l0 == v12)|base.B2i32(base.Ui32(v12+int32(1)) < base.Ui32(int32(2))) == int32(0) {
				if l0 == int32(0) {
					v27 = F_superuser(m)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						if v27 != 0 {
							if base.Ui32(v15+int32(1)) < base.Ui32(int32(2)) {
								m.G0 = v9 - int32(-64)
								return
							} else {
								v78 = F_superuser(m)
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return
								} else {
									if v78 != 0 {
										m.G0 = v9 - int32(-64)
										return
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return
											} else {
												if base.Ui32(l0) <= base.Ui32(int32(41)) {
													v91 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
													v93 = v91
												} else {
													v93 = int32(_a_F_check_encoding_locale_matches_0)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l1
												*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v93
												F_errmsg(m, int32(_a_F_check_encoding_locale_matches_1), v7+int32(-48))
												mBase = m.M
												v100 = m.ExcPending
												if v100 != 0 {
													return
												} else {
													if base.Ui32(v15) <= base.Ui32(int32(41)) {
														v105 = *(*int32)(unsafe.Add(mBase, uint32(v15<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
														v107 = v105
													} else {
														v107 = int32(_a_F_check_encoding_locale_matches_0)
													}
													*(*int32)(unsafe.Add(mBase, uint32(v9))) = v107
													F_errdetail(m, int32(_a_F_check_encoding_locale_matches_2), v9)
													mBase = m.M
													v111 = m.ExcPending
													if v111 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_check_encoding_locale_matches_3), int32(1629), int32(_a_F_check_encoding_locale_matches_4))
														mBase = m.M
														v116 = m.ExcPending
														if v116 != 0 {
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
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return
								} else {
									if base.Ui32(l0) <= base.Ui32(int32(41)) {
										v40 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
										v42 = v40
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
											v54 = *(*int32)(unsafe.Add(mBase, uint32(v12<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
											v56 = v54
										} else {
											v56 = int32(_a_F_check_encoding_locale_matches_0)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v56
										F_errdetail(m, int32(_a_F_check_encoding_locale_matches_5), v7+int32(-32))
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_check_encoding_locale_matches_3), int32(1614), int32(_a_F_check_encoding_locale_matches_4))
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
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
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							if base.Ui32(l0) <= base.Ui32(int32(41)) {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
								v42 = v40
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
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v12<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
									v56 = v54
								} else {
									v56 = int32(_a_F_check_encoding_locale_matches_0)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v56
								F_errdetail(m, int32(_a_F_check_encoding_locale_matches_5), v7+int32(-32))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_check_encoding_locale_matches_3), int32(1614), int32(_a_F_check_encoding_locale_matches_4))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
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
				if base.B2i32(l0 == v15)|base.B2i32(base.Ui32(v15+int32(1)) < base.Ui32(int32(2))) != 0 {
					m.G0 = v9 - int32(-64)
					return
				} else {
					if l0 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v86 = m.ExcPending
							if v86 != 0 {
								return
							} else {
								if base.Ui32(l0) <= base.Ui32(int32(41)) {
									v91 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
									v93 = v91
								} else {
									v93 = int32(_a_F_check_encoding_locale_matches_0)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l1
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v93
								F_errmsg(m, int32(_a_F_check_encoding_locale_matches_1), v7+int32(-48))
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return
								} else {
									if base.Ui32(v15) <= base.Ui32(int32(41)) {
										v105 = *(*int32)(unsafe.Add(mBase, uint32(v15<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
										v107 = v105
									} else {
										v107 = int32(_a_F_check_encoding_locale_matches_0)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v107
									F_errdetail(m, int32(_a_F_check_encoding_locale_matches_2), v9)
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_check_encoding_locale_matches_3), int32(1629), int32(_a_F_check_encoding_locale_matches_4))
										mBase = m.M
										v116 = m.ExcPending
										if v116 != 0 {
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
						v78 = F_superuser(m)
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return
						} else {
							if v78 != 0 {
								m.G0 = v9 - int32(-64)
								return
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return
									} else {
										if base.Ui32(l0) <= base.Ui32(int32(41)) {
											v91 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
											v93 = v91
										} else {
											v93 = int32(_a_F_check_encoding_locale_matches_0)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l1
										*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v93
										F_errmsg(m, int32(_a_F_check_encoding_locale_matches_1), v7+int32(-48))
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return
										} else {
											if base.Ui32(v15) <= base.Ui32(int32(41)) {
												v105 = *(*int32)(unsafe.Add(mBase, uint32(v15<<(uint(int32(3))%32))+uint32(_c_F_check_encoding_locale_matches[0])))
												v107 = v105
											} else {
												v107 = int32(_a_F_check_encoding_locale_matches_0)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v9))) = v107
											F_errdetail(m, int32(_a_F_check_encoding_locale_matches_2), v9)
											mBase = m.M
											v111 = m.ExcPending
											if v111 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_check_encoding_locale_matches_3), int32(1629), int32(_a_F_check_encoding_locale_matches_4))
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
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
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v7 = F_LWLockAcquire(m, v5, int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v11 != int32(-1) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v95 = v5 + int32(20)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v96))))
	if v98 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v16 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14+v11)+20)) = uint8(v16)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v18 - int32(1)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+172)))
	if v22 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v23 = int32(0)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v27 = F_ExecFindMatchingSubPlans(m, v24, v23, v23)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v29 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+172)) = uint8(v29)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v27
	v32 = int32(0)
	if v27 == v32 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if base.B2i32(v67 == v68)|base.B2i32(v68 <= int32(0)) != 0 {
		goto L3
	} else {
		goto L22
	}
L10:
	;
	v67 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v39 = int32(1)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v40 <= v39 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v43 = v39
	goto L15
L14:
	;
	v43 = v40
	goto L15
L15:
	;
	v47 = int32(0)
	v49 = v32
	goto L16
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(8)+v47<<(uint(int32(2))%32))))
	if v55 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v67 = v58
	goto L9
L18:
	;
	v58 = v49 + base.I32_popcnt(v55)
	goto L20
L19:
	;
	v58 = v49
	goto L20
L20:
	;
	v60 = v47 + int32(1)
	if v60 != v43 {
		v47 = v60
		v49 = v58
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L17
L22:
	;
	v74 = v23
	goto L23
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v78 = F_bms_is_member(m, v74, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	goto L3
L25:
	;
	if v78 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v84 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v82+v74)+20)) = uint8(v84)
	goto L28
L27:
	;
	goto L28
L28:
	;
	v87 = v74 + int32(1)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v87 < v88 {
		v74 = v87
		goto L23
	} else {
		goto L29
	}
L29:
	;
	goto L24
L30:
	;
	v102 = v96
	goto L33
L31:
	;
	v121 = v96
	goto L32
L32:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v121 < v124 {
		goto L40
	} else {
		goto L41
	}
L33:
	;
	if v102 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v121 = v116
	goto L32
L35:
	;
	v107 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v107
	F_LWLockRelease(m, v5)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v116 = v102 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v116
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116+v95))))
	if v119 != 0 {
		v102 = v116
		goto L33
	} else {
		goto L39
	}
L38:
	;
	return int32(0)
L39:
	;
	goto L34
L40:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v128 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v126+v121)+20)) = uint8(v128)
	goto L42
L41:
	;
	goto L42
L42:
	;
	F_LWLockRelease(m, v5)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	return int32(1)
}
func F_cidr_abbrev(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
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
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v346 int32
	_ = v346
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v634 int32
	_ = v634
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	v2 = int32(0)
	v16 = m.G0
	v18 = v16 + int32(-64)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = F_pg_detoast_datum_packed(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v25 = int32(1)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v27&v25 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v30 = v25
	goto L5
L4:
	;
	v30 = int32(4)
	goto L5
L5:
	;
	v31 = v21 + v30
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	v33 = int32(2)
	v34 = v31 + v33
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	v36 = int32(50)
	v37 = m.G0
	v39 = v37 - int32(240)
	m.G0 = v39
	switch v32 - v33 {
	case 0:
		goto L13
	case 1:
		goto L12
	default:
		goto L11
	}
L6:
	;
	m.G0 = v39 + int32(240)
	if v634 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L7:
	;
	v634 = int32(0)
	goto L6
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_cidr_abbrev[0])) = int32(35)
	goto L7
L9:
	;
	if base.Ui32(v585) < base.Ui32(int32(5)) {
		goto L8
	} else {
		goto L158
	}
L10:
	;
	v550 = v35 & int32(7)
	if v550 == int32(0) {
		v582 = v537
		v585 = v540
		goto L9
	} else {
		goto L152
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_cidr_abbrev[0])) = int32(5)
	goto L7
L12:
	;
	if base.Ui32(int32(129)) <= base.Ui32(v35) {
		goto L30
	} else {
		goto L31
	}
L13:
	;
	if base.Ui32(int32(33)) <= base.Ui32(v35) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_cidr_abbrev[0])) = int32(28)
	goto L7
L15:
	;
	goto L16
L16:
	;
	if v35 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v50 = int32(48)
	*(*uint16)(unsafe.Add(mBase, uint32(v18))) = uint16(v50)
	v582 = v16 + int32(-63)
	v585 = int32(49)
	goto L9
L18:
	;
	goto L19
L19:
	;
	v56 = int32(base.Ui32(v35) >> (uint(int32(3)) % 32))
	if v56 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v537 = v18
	v539 = v34
	v540 = v36
	goto L10
L21:
	;
	goto L22
L22:
	;
	v60 = v18
	v63 = v56
	v64 = v34
	v65 = v36
	goto L23
L23:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v74
	v77 = v64 + int32(1)
	v81 = F_pg_sprintf(m, v60, int32(_a_F_cidr_abbrev_0), v39+int32(32))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	goto L8
L25:
	;
	v83 = v81 + v60
	if v63 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v537 = v83
	v539 = v77
	v540 = v60 + v65 - v83
	goto L10
L27:
	;
	goto L28
L28:
	;
	v88 = int32(46)
	*(*uint16)(unsafe.Add(mBase, uint32(v83))) = uint16(v88)
	v90 = int32(1)
	v94 = v83 + v90
	v95 = v60 + v65 - v94
	if base.Ui32(int32(6)) <= base.Ui32(v95) {
		v60 = v94
		v63 = v63 - v90
		v64 = v77
		v65 = v95
		goto L23
	} else {
		goto L29
	}
L29:
	;
	goto L24
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_cidr_abbrev[0])) = int32(28)
	goto L7
L31:
	;
	goto L32
L32:
	;
	if v35 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v35
	v445 = F_pg_sprintf(m, v440, int32(_a_F_cidr_abbrev_1), v39+int32(48))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L127
	}
L34:
	;
	v105 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+162)) = uint8(v105)
	v107 = int32(_a_F_cidr_abbrev_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+160)) = uint16(v107)
	v440 = v39 + int32(160) | int32(2)
	goto L33
L35:
	;
	goto L36
L36:
	;
	v116 = int32(base.Ui32(v35+int32(7)) >> (uint(int32(3)) % 32))
	if v116 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	base.MemoryCopy(m, v39+int32(224), v34, v116)
	goto L39
L38:
	;
	goto L39
L39:
	;
	v122 = v39 + int32(224) + v116
	v124 = int32(16) - v116
	if v124 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	base.MemoryFill(m, v122, int32(0), v124)
	goto L42
L41:
	;
	goto L42
L42:
	;
	v128 = v35 & int32(7)
	if v128 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v130 = v122 - int32(1)
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	v136 = v131 & (int32(-1) << (uint(int32(8)-v128) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v136)
	goto L45
L44:
	;
	goto L45
L45:
	;
	v143 = int32(base.Ui32(v35+int32(15)) >> (uint(int32(4)) % 32))
	if v143 == int32(1) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v146 = int32(2)
	goto L48
L47:
	;
	v146 = v143
	goto L48
L48:
	;
	v149 = int32(0)
	v152 = v149
	v154 = v2
	v156 = v149
	v159 = v2
	v160 = v2
	goto L49
L49:
	;
	v168 = v39 + int32(224) + v156
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+1)))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	if v169|v170 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v190 = int32(0)
	v194 = base.B2i32(v183 != v190) & base.B2i32(v185 < v183)
	if v194 != 0 {
		goto L64
	} else {
		goto L65
	}
L51:
	;
	v188 = v156 + int32(2)
	if base.Ui32(v188) < base.Ui32(v146<<(uint(int32(1))%32)) {
		v152 = v183
		v154 = v184
		v156 = v188
		v159 = v185
		v160 = v186
		goto L49
	} else {
		goto L62
	}
L52:
	;
	if v152 != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	if v152 != 0 {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v176 = v154
	goto L57
L56:
	;
	v176 = int32(base.Ui32(v156) >> (uint(int32(1)) % 32))
	goto L57
L57:
	;
	v183 = v152 + int32(1)
	v184 = v176
	v185 = v159
	v186 = v160
	goto L51
L58:
	;
	if v152 <= v159 {
		v183 = v152
		v184 = v154
		v185 = v159
		v186 = v160
		goto L51
	} else {
		goto L61
	}
L59:
	;
	v180 = v159
	v181 = v160
	goto L60
L60:
	;
	v183 = int32(0)
	v184 = v154
	v185 = v180
	v186 = v181
	goto L51
L61:
	;
	v180 = v152
	v181 = v154
	goto L60
L62:
	;
	goto L50
L63:
	;
	v215 = v196 + v195
	v217 = v146 - int32(1)
	if v217 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L64:
	;
	v195 = v184
	goto L66
L65:
	;
	v195 = v186
	goto L66
L66:
	;
	if v194 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v196 = v183
	goto L69
L68:
	;
	v196 = v185
	goto L69
L69:
	;
	if v195|base.B2i32(v196 == v146) != 0 {
		v214 = v2
		goto L63
	} else {
		goto L70
	}
L70:
	;
	switch v196 - int32(5) {
	case 0:
		goto L73
	case 1:
		goto L71
	case 2:
		goto L72
	default:
		v214 = v2
		goto L63
	}
L71:
	;
	v214 = int32(1)
	goto L63
L72:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+238)))
	if v207 == int32(0) {
		v214 = v2
		goto L63
	} else {
		goto L76
	}
L73:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+234)))
	if v201 != int32(255) {
		v214 = v2
		goto L63
	} else {
		goto L74
	}
L74:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+235)))
	if v204 == int32(255) {
		goto L71
	} else {
		goto L75
	}
L75:
	;
	v214 = v2
	goto L63
L76:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+239)))
	if v210 == int32(1) {
		v214 = v2
		goto L63
	} else {
		goto L77
	}
L77:
	;
	goto L71
L78:
	;
	v346 = int32(0)
	if base.B2i32(base.B2i32(v196 == v346)|base.B2i32(v335 < v195) == v346)&base.B2i32(v335 < v215) == v346 {
		goto L107
	} else {
		goto L108
	}
L79:
	;
	v332 = v39 + int32(160)
	v334 = v39 + int32(224)
	v335 = v190
	goto L78
L80:
	;
	goto L81
L81:
	;
	v233 = v39 + int32(160)
	v235 = v39 + int32(224)
	v236 = v190
	goto L82
L82:
	;
	v247 = int32(0)
	if base.B2i32(v196 == v247)|base.B2i32(v236 < v195)|base.B2i32(v215 <= v236) == v247 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	v332 = v325
	v334 = v326
	v335 = v329
	goto L78
L84:
	;
	v329 = v236 + int32(1)
	if v146-int32(2) != v236 {
		v233 = v325
		v235 = v326
		v236 = v329
		goto L82
	} else {
		goto L106
	}
L85:
	;
	if v236 == v195 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v236))&v214 != 0 {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	v256 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v233))) = uint8(v256)
	v260 = v233 + int32(1)
	goto L90
L89:
	;
	v260 = v233
	goto L90
L90:
	;
	v325 = v260
	v326 = v235 + int32(2)
	goto L84
L91:
	;
	if v236 == int32(6) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	goto L93
L93:
	;
	v304 = v39 + int32(160)
	if v304 == v233 {
		goto L102
	} else {
		goto L103
	}
L94:
	;
	v270 = int32(58)
	goto L96
L95:
	;
	v270 = int32(46)
	goto L96
L96:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v233))) = uint8(v270)
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235))))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+128)) = v272
	v275 = v233 + int32(1)
	v279 = F_pg_sprintf(m, v275, int32(_a_F_cidr_abbrev_0), v39+int32(128))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v281 = v279 + v275
	if base.B2i32(base.Ui32(int32(120)) < base.Ui32(v35))|base.B2i32(v236 != int32(7)) == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v325 = v281
	v326 = v235 + int32(1)
	goto L84
L99:
	;
	goto L100
L100:
	;
	v289 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v281))) = uint8(v289)
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+112)) = v291
	v294 = v281 + int32(1)
	v298 = F_pg_sprintf(m, v294, int32(_a_F_cidr_abbrev_0), v39+int32(112))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v325 = v298 + v294
	v326 = v235 + int32(2)
	goto L84
L102:
	;
	v310 = v304
	goto L104
L103:
	;
	v306 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v233))) = uint8(v306)
	v310 = v233 + int32(1)
	goto L104
L104:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+1)))
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235))))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+144)) = v311 | v312<<(uint(int32(8))%32)
	v322 = F_pg_sprintf(m, v310, int32(_a_F_cidr_abbrev_3), v39+int32(144))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v325 = v322 + v310
	v326 = v235 + int32(2)
	goto L84
L106:
	;
	goto L83
L107:
	;
	if v214&base.B2i32(base.Ui32(int32(5)) < base.Ui32(v335)) == int32(0) {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	goto L109
L109:
	;
	if v335 == v195 {
		goto L123
	} else {
		goto L124
	}
L110:
	;
	v362 = v39 + int32(160)
	if v362 == v332 {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L112
L112:
	;
	if v335 == int32(6) {
		goto L117
	} else {
		goto L118
	}
L113:
	;
	v368 = v362
	goto L115
L114:
	;
	v364 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v332))) = uint8(v364)
	v368 = v332 + int32(1)
	goto L115
L115:
	;
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334)+1)))
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334))))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+96)) = v369 | v370<<(uint(int32(8))%32)
	v378 = F_pg_sprintf(m, v368, int32(_a_F_cidr_abbrev_3), v39+int32(96))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v440 = v378 + v368
	goto L33
L117:
	;
	v385 = int32(58)
	goto L119
L118:
	;
	v385 = int32(46)
	goto L119
L119:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v332))) = uint8(v385)
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334))))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+80)) = v387
	v390 = v332 + int32(1)
	v394 = F_pg_sprintf(m, v390, int32(_a_F_cidr_abbrev_0), v39+int32(80))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v396 = v394 + v390
	if base.B2i32(v335 == int32(7))&base.B2i32(base.Ui32(v35) <= base.Ui32(int32(120))) != 0 {
		v440 = v396
		goto L33
	} else {
		goto L121
	}
L121:
	;
	v402 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v396))) = uint8(v402)
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v404
	v407 = v396 + int32(1)
	v411 = F_pg_sprintf(m, v407, int32(_a_F_cidr_abbrev_0), v39-int32(-64))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v440 = v411 + v407
	goto L33
L123:
	;
	v415 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v332))) = uint8(v415)
	v419 = v332 + int32(1)
	goto L125
L124:
	;
	v419 = v332
	goto L125
L125:
	;
	if v335 != v217 {
		v440 = v419
		goto L33
	} else {
		goto L126
	}
L126:
	;
	v421 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v419))) = uint8(v421)
	v440 = v419 + int32(1)
	goto L33
L127:
	;
	v448 = v39 + int32(160)
	v449 = F_strlen(m, v448)
	mBase = m.M
	if base.Ui32(v449+int32(1)) <= base.Ui32(int32(50)) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	if (v448^v18)&int32(3) != 0 {
		goto L134
	} else {
		goto L135
	}
L129:
	;
	goto L130
L130:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_cidr_abbrev[0])) = int32(35)
	goto L7
L131:
	;
	v634 = v18
	goto L6
L132:
	;
	goto L131
L133:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v508))) = uint8(v507)
	if v507&int32(255) == int32(0) {
		goto L132
	} else {
		goto L148
	}
L134:
	;
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448))))
	v506 = v448
	v507 = v459
	v508 = v18
	goto L133
L135:
	;
	goto L136
L136:
	;
	if v448&int32(3) != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v463 = v448
	v465 = v18
	goto L140
L138:
	;
	v477 = v448
	v479 = v18
	goto L139
L139:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v477)))
	v484 = int32(-2139062144)
	if (int32(16843008)-v481|v481)&v484 != v484 {
		v506 = v477
		v507 = v481
		v508 = v479
		goto L133
	} else {
		goto L144
	}
L140:
	;
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463))))
	*(*uint8)(unsafe.Add(mBase, uint32(v465))) = uint8(v466)
	if v466 == int32(0) {
		goto L132
	} else {
		goto L142
	}
L141:
	;
	v477 = v473
	v479 = v471
	goto L139
L142:
	;
	v470 = int32(1)
	v471 = v465 + v470
	v473 = v463 + v470
	if v473&int32(3) != 0 {
		v463 = v473
		v465 = v471
		goto L140
	} else {
		goto L143
	}
L143:
	;
	goto L141
L144:
	;
	v489 = v477
	v490 = v481
	v491 = v479
	goto L145
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v491))) = v490
	v493 = int32(4)
	v494 = v491 + v493
	v496 = v489 + v493
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v489)+4))
	v501 = int32(-2139062144)
	if (int32(16843008)-v498|v498)&v501 == v501 {
		v489 = v496
		v490 = v498
		v491 = v494
		goto L145
	} else {
		goto L147
	}
L146:
	;
	v506 = v496
	v507 = v498
	v508 = v494
	goto L133
L147:
	;
	goto L146
L148:
	;
	v515 = v506
	v517 = v508
	goto L149
L149:
	;
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v517)+1)) = uint8(v518)
	v520 = int32(1)
	if v518 != 0 {
		v515 = v515 + v520
		v517 = v517 + v520
		goto L149
	} else {
		goto L151
	}
L150:
	;
	goto L132
L151:
	;
	goto L150
L152:
	;
	if base.Ui32(v540) < base.Ui32(int32(6)) {
		goto L8
	} else {
		goto L153
	}
L153:
	;
	if v537 != v18 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v556 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v537))) = uint8(v556)
	v560 = v537 + int32(1)
	goto L156
L155:
	;
	v560 = v18
	goto L156
L156:
	;
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v539))))
	v562 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v561 & ((v562<<(uint(v550)%32) ^ v562) << (uint(int32(8)-v550) % 32))
	v575 = F_pg_sprintf(m, v560, int32(_a_F_cidr_abbrev_0), v39+int32(16))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	v577 = v575 + v560
	v582 = v577
	v585 = v537 + v540 - v577
	goto L9
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v35
	v598 = F_pg_sprintf(m, v582, int32(_a_F_cidr_abbrev_1), v39)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	v634 = v18
	goto L6
L160:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L1
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	v670 = F_cstring_to_text(m, v18)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L1
	} else {
		goto L167
	}
L163:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	F_errmsg(m, int32(_a_F_cidr_abbrev_4), int32(0))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(_a_F_cidr_abbrev_5), int32(1217), int32(_a_F_cidr_abbrev_6))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L167:
	;
	m.G0 = v18 - int32(-64)
	return v670
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
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
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
				v137 = v39
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v141 = m.ExcPending
				if v141 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v144 = m.ExcPending
					if v144 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v137
						F_errmsg(m, int32(_a_F_cidr_set_masklen_0), v10)
						mBase = m.M
						v148 = m.ExcPending
						if v148 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_cidr_set_masklen_1), int32(357), int32(_a_F_cidr_set_masklen_2))
							mBase = m.M
							v153 = m.ExcPending
							if v153 != 0 {
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
						v81 = int32(base.Ui32(v39+int32(7)) >> (uint(int32(3)) % 32))
						if v81 != 0 {
							v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
							if v82&int32(1) != 0 {
								v85 = v71
							} else {
								v85 = v73
							}
							v88 = int32(1)
							v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
							if v92&v88 != 0 {
								v95 = v13 + v88
							} else {
								v95 = v13 + int32(4)
							}
							base.MemoryCopy(m, v85+int32(2), v95+int32(2), v81)
						} else {
						}
						v100 = v39 & int32(7)
						if v100 == int32(0) {
						} else {
							v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
							if v105&int32(1) != 0 {
								v108 = v71
							} else {
								v108 = v73
							}
							v109 = int32(base.Ui32(v39)>>(uint(int32(3))%32)) + v108
							v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+2)))
							v113 = v110 & (int32(-256) >> (uint(v100) % 32))
							*(*uint8)(unsafe.Add(mBase, uint32(v109)+2)) = uint8(v113)
						}
					}
					v120 = int32(1)
					v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
					if v122&v120 != 0 {
						v125 = v120
					} else {
						v125 = int32(4)
					}
					v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+v125))))
					if v127 == int32(2) {
						v130 = int32(40)
					} else {
						v130 = int32(88)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v52))) = v130
					m.G0 = v10 + int32(16)
					return v52
				}
			}
		} else {
			if v17 < int32(0) {
				v137 = v17
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v141 = m.ExcPending
				if v141 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v144 = m.ExcPending
					if v144 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v137
						F_errmsg(m, int32(_a_F_cidr_set_masklen_0), v10)
						mBase = m.M
						v148 = m.ExcPending
						if v148 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_cidr_set_masklen_1), int32(357), int32(_a_F_cidr_set_masklen_2))
							mBase = m.M
							v153 = m.ExcPending
							if v153 != 0 {
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
					v137 = v39
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v141 = m.ExcPending
					if v141 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v144 = m.ExcPending
						if v144 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v137
							F_errmsg(m, int32(_a_F_cidr_set_masklen_0), v10)
							mBase = m.M
							v148 = m.ExcPending
							if v148 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_cidr_set_masklen_1), int32(357), int32(_a_F_cidr_set_masklen_2))
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
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
							v81 = int32(base.Ui32(v39+int32(7)) >> (uint(int32(3)) % 32))
							if v81 != 0 {
								v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
								if v82&int32(1) != 0 {
									v85 = v71
								} else {
									v85 = v73
								}
								v88 = int32(1)
								v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
								if v92&v88 != 0 {
									v95 = v13 + v88
								} else {
									v95 = v13 + int32(4)
								}
								base.MemoryCopy(m, v85+int32(2), v95+int32(2), v81)
							} else {
							}
							v100 = v39 & int32(7)
							if v100 == int32(0) {
							} else {
								v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
								if v105&int32(1) != 0 {
									v108 = v71
								} else {
									v108 = v73
								}
								v109 = int32(base.Ui32(v39)>>(uint(int32(3))%32)) + v108
								v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+2)))
								v113 = v110 & (int32(-256) >> (uint(v100) % 32))
								*(*uint8)(unsafe.Add(mBase, uint32(v109)+2)) = uint8(v113)
							}
						}
						v120 = int32(1)
						v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
						if v122&v120 != 0 {
							v125 = v120
						} else {
							v125 = int32(4)
						}
						v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+v125))))
						if v127 == int32(2) {
							v130 = int32(40)
						} else {
							v130 = int32(88)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v52))) = v130
						m.G0 = v10 + int32(16)
						return v52
					}
				}
			}
		}
	}
}
func F_citextcmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
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
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	v8 = int32(1)
	v9 = l1 + v8
	v11 = l0 + v8
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v16 = v14 & v8
	if v16 != 0 {
		v17 = v11
	} else {
		v17 = l0 + int32(4)
	}
	if v14 == int32(1) {
		v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
		if v23 == int32(18) {
			v26 = int32(16)
		} else {
			v26 = int32(0)
		}
		if base.Ui32((v23-int32(1))&int32(255)) < base.Ui32(int32(3)) {
			v33 = int32(4)
		} else {
			v33 = v26
		}
		v44 = v33
	} else {
		v34 = int32(1)
		if v16 != 0 {
			v44 = int32(base.Ui32(v14)>>(uint(v34)%32)) - v34
		} else {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v44 = int32(base.Ui32(v38)>>(uint(int32(2))%32)) - int32(4)
		}
	}
	v46 = F_str_tolower(m, v17, v44, int32(100))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		return int32(0)
	} else {
		v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		v54 = v52 & int32(1)
		if v54 != 0 {
			v55 = v9
		} else {
			v55 = l1 + int32(4)
		}
		if v52 == int32(1) {
			v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			if v61 == int32(18) {
				v64 = int32(16)
			} else {
				v64 = int32(0)
			}
			if base.Ui32((v61-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v71 = int32(4)
			} else {
				v71 = v64
			}
			v82 = v71
		} else {
			v72 = int32(1)
			if v54 != 0 {
				v82 = int32(base.Ui32(v52)>>(uint(v72)%32)) - v72
			} else {
				v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v82 = int32(base.Ui32(v76)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v84 = F_str_tolower(m, v55, v82, int32(100))
		mBase = m.M
		v85 = m.ExcPending
		if v85 != 0 {
			return int32(0)
		} else {
			v86 = F_strlen(m, v46)
			mBase = m.M
			v87 = F_strlen(m, v84)
			mBase = m.M
			v88 = F_varstr_cmp(m, v46, v86, v84, v87, l2)
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v46)
				mBase = m.M
				v91 = m.ExcPending
				if v91 != 0 {
					return int32(0)
				} else {
					F_pfree(m, v84)
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return int32(0)
					} else {
						return v88
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
			return base.I32_trunc_sat_f64_s(v16)
		}
	}
}
func F_clamp_row_est(m *base.Module, l0 float64) float64 {
	var v2 float64
	_ = v2
	var v11 float64
	_ = v11
	var v15 float64
	_ = v15
	v2 = float64(1e+100)
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(l0)&int64(9223372036854775807)))|base.F64_gt(l0, v2) != 0 {
		v15 = v2
	} else {
		v11 = float64(1)
		if base.F64_le(l0, v11) != 0 {
			v15 = v11
		} else {
			v15 = base.F64_nearest(l0)
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 float64
	_ = v52
	var v53 int32
	_ = v53
	var v54 float64
	_ = v54
	var v57 int32
	_ = v57
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
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
	var v201 int32
	_ = v201
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
	v40 = F_find_single_rel_for_clauses(m, l0, l1)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L5
	} else {
		goto L8
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
	if l1 == int32(0) {
		v268 = v54
		goto L1
	} else {
		goto L13
	}
L8:
	;
	if base.B2i32(l5 == int32(0))|base.B2i32(v40 == int32(0)) != 0 {
		v54 = v37
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)+76))
	if v45 != 0 {
		v54 = v37
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v40)+112))
	if v46 == int32(0) {
		v54 = v37
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v52 = F_statext_clauselist_selectivity(m, l0, l1, l2, l3, l4, v40, v20+int32(12), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v54 = v52
	goto L7
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v57 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v71 = v7
	v72 = int32(-1)
	v75 = v54
	goto L17
L15:
	;
	v192 = v54
	goto L16
L16:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v195 == int32(0) {
		v268 = v192
		goto L1
	} else {
		goto L56
	}
L17:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78+v71<<(uint(int32(2))%32))))
	v84 = v72 + int32(1)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v86 = F_bms_is_member(m, v84, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L20
	}
L18:
	;
	v192 = v173
	goto L16
L19:
	;
	v175 = v71 + int32(1)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v175 < v176 {
		v71 = v175
		v72 = v84
		v75 = v173
		goto L17
	} else {
		goto L55
	}
L20:
	;
	if v86 != 0 {
		v173 = v75
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v88 = F_clause_selectivity_ext(m, l0, v82, l2, l3, l4, l5)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v91 != int32(318) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v173 = base.F64_mul(v75, v88)
	goto L19
L24:
	;
	if v103 != int32(17) {
		goto L23
	} else {
		goto L30
	}
L25:
	;
	v101 = v82
	v102 = int32(0)
	v103 = v91
	goto L24
L26:
	;
	goto L27
L27:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+10)))
	if v94 == int32(1) {
		goto L23
	} else {
		goto L28
	}
L28:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v97 == int32(0) {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v101 = v97
	v102 = v82
	v103 = v100
	goto L24
L30:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v101)+28))
	if v106 == int32(0) {
		goto L23
	} else {
		goto L31
	}
L31:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	if v109 != int32(2) {
		goto L23
	} else {
		goto L32
	}
L32:
	;
	if v102 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v147 = F_get_oprrest(m, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L5
	} else {
		goto L52
	}
L34:
	;
	v145 = int32(0)
	goto L33
L35:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v102)+24))
	if v112 != int32(1) {
		goto L23
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v127 = F_NumRelids(m, l0, v101)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L5
	} else {
		goto L43
	}
L38:
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
		goto L39
	}
L39:
	;
	if v119 != 0 {
		v145 = int32(1)
		goto L33
	} else {
		goto L40
	}
L40:
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
		goto L41
	}
L41:
	;
	if v125 != 0 {
		goto L34
	} else {
		goto L42
	}
L42:
	;
	goto L23
L43:
	;
	if v127 != int32(1) {
		goto L23
	} else {
		goto L44
	}
L44:
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
		goto L45
	}
L45:
	;
	if v135 != 0 {
		v145 = int32(1)
		goto L33
	} else {
		goto L46
	}
L46:
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
		goto L47
	}
L47:
	;
	if v140 == int32(0) {
		goto L23
	} else {
		goto L48
	}
L48:
	;
	goto L34
L49:
	;
	F_addRangeClause(m, v20+int32(8), v101, v145, int32(0), v88)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L5
	} else {
		goto L54
	}
L50:
	;
	F_addRangeClause(m, v20+int32(8), v101, v145, int32(1), v88)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L5
	} else {
		goto L53
	}
L51:
	;
	switch v147 - int32(336) {
	case 0:
		goto L50
	case 1:
		goto L49
	default:
		goto L23
	}
L52:
	;
	switch v147 - int32(103) {
	case 0:
		goto L50
	case 1:
		goto L49
	default:
		goto L51
	}
L53:
	;
	v173 = v75
	goto L19
L54:
	;
	v173 = v75
	goto L19
L55:
	;
	goto L18
L56:
	;
	v201 = v195
	v212 = v192
	goto L57
L57:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+8)))
	if v215 == int32(1) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v268 = v253
	goto L1
L59:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	F_pfree(m, v201)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L5
	} else {
		goto L73
	}
L60:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+9)))
	if v218 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	v246 = *(*float64)(unsafe.Add(mBase, uint32(v201)+24))
	v247 = v246
	goto L59
L63:
	;
	v221 = float64(0.005)
	v222 = *(*float64)(unsafe.Add(mBase, uint32(v201)+24))
	if base.F64_eq(v222, float64(0.3333333333333333)) != 0 {
		v247 = v221
		goto L59
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v245 = *(*float64)(unsafe.Add(mBase, uint32(v201)+16))
	v247 = v245
	goto L59
L66:
	;
	v225 = *(*float64)(unsafe.Add(mBase, uint32(v201)+16))
	if base.F64_eq(v225, float64(0.3333333333333333)) != 0 {
		v247 = v221
		goto L59
	} else {
		goto L67
	}
L67:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v201)+4))
	v233 = F_nulltestsel(m, l0, int32(0), v232, l2)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L5
	} else {
		goto L68
	}
L68:
	;
	v235 = base.F64_add(base.F64_add(base.F64_add(v222, v225), float64(-1)), v233)
	if base.F64_le(v235, float64(0)) == int32(0) {
		v247 = v235
		goto L59
	} else {
		goto L69
	}
L69:
	;
	if base.F64_lt(v235, float64(-0.01)) != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v244 = float64(0.005)
	goto L72
L71:
	;
	v244 = float64(1e-10)
	goto L72
L72:
	;
	v247 = v244
	goto L59
L73:
	;
	v253 = base.F64_mul(v212, v247)
	if v250 != 0 {
		v201 = v250
		v212 = v253
		goto L57
	} else {
		goto L74
	}
L74:
	;
	goto L58
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
	var v69 int32
	_ = v69
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
	var v179 int32
	_ = v179
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
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
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
	var v304 int32
	_ = v304
	var v305 int64
	_ = v305
	var v311 int32
	_ = v311
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
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
	var v394 int32
	_ = v394
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
		goto L25
	} else {
		goto L26
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
	if l7 == int32(0) {
		v51 = v30
		goto L9
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if l7 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	base.MemoryCopy(m, v30, l6, l7)
	v51 = v30
	goto L9
L22:
	;
	base.MemoryFill(m, v30, int32(0), l7)
	goto L24
L23:
	;
	goto L24
L24:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v49 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v30+v47))) = uint8(v49)
	v51 = v30
	goto L9
L25:
	;
	if l5 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L26:
	;
	v69 = v56
	goto L27
L27:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
	if v74 != 0 {
		goto L25
	} else {
		goto L29
	}
L28:
	;
	goto L25
L29:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	switch v76 - int32(76) {
	case 0, 18, 21, 38:
		goto L33
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37:
		goto L32
	default:
		goto L34
	}
L30:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	if v457 != 0 {
		v69 = v457
		goto L27
	} else {
		goto L148
	}
L31:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+v120))))
	if v122 != 0 {
		goto L30
	} else {
		goto L44
	}
L32:
	;
	F_cparc(m, l0, v69, l2, v75)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L43
	}
L33:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
	if v81 == int32(0) {
		goto L32
	} else {
		goto L36
	}
L34:
	;
	if v76 != int32(36) {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v85 = v81
	goto L37
L37:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	switch v98 - int32(76) {
	case 0, 18, 21, 38:
		goto L31
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37:
		goto L39
	default:
		goto L40
	}
L38:
	;
	goto L32
L39:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
	if v103 != 0 {
		v85 = v103
		goto L37
	} else {
		goto L42
	}
L40:
	;
	if v98 == int32(36) {
		goto L31
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	goto L38
L43:
	;
	goto L30
L44:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v123 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if l4 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L46:
	;
	v154 = int32(0)
	goto L45
L47:
	;
	goto L48
L48:
	;
	v128 = v123
	goto L49
L49:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+24))
	if v142 == v75 {
		v154 = v141
		goto L45
	} else {
		goto L51
	}
L50:
	;
	v154 = int32(0)
	goto L45
L51:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
	if v144 != 0 {
		v128 = v144
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v436 = F_newstate(m, l0)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L145
	}
L54:
	;
	if v154 != 0 {
		goto L71
	} else {
		goto L72
	}
L55:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v167 != 0 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v76 != v162 {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)))
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+4)))
	if v164 == v165 {
		goto L54
	} else {
		goto L58
	}
L58:
	;
	goto L55
L59:
	;
	v169 = v167
	v179 = l2
	goto L62
L60:
	;
	goto L61
L61:
	;
	if v154 == int32(0) {
		goto L53
	} else {
		goto L69
	}
L62:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v179)+8))
	if v182 != int32(1) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	goto L61
L64:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v169)+8))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)+16))
	if v191 != 0 {
		v169 = v191
		v179 = v190
		goto L62
	} else {
		goto L68
	}
L65:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	if v76 != v185 {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)))
	v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v169)+4)))
	if v187 == v188 {
		goto L54
	} else {
		goto L67
	}
L67:
	;
	goto L64
L68:
	;
	goto L63
L69:
	;
	F_cparc(m, l0, v69, l2, v154)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	goto L30
L71:
	;
	goto L74
L72:
	;
	goto L73
L73:
	;
	F_clonesuccessorstates(m, l0, v75, l2, l3, l4, v51, l6, l7)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L144
	}
L74:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v154)+16))
	if v238 != 0 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	goto L105
L76:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	v245 = int32(*(*int16)(unsafe.Add(mBase, uint32(v238)+4)))
	if v245 < int32(0) {
		goto L80
	} else {
		goto L81
	}
L77:
	;
	goto L78
L78:
	;
	goto L75
L79:
	;
	goto L74
L80:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v238)+16))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v238)+20))
	if v280 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L81:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	v250 = v248 - int32(97)
	if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v250))|base.B2i32(int32(1)<<(uint(v250)%32)&int32(_a_F_clonesuccessorstates_0) == int32(0)) != 0 {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v260 != 0 {
		goto L80
	} else {
		goto L83
	}
L83:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v238)+36))
	if v261 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	if v273 != 0 {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+20))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v238)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v265+v245*int32(24))+12)) = v269
	v273 = v269
	goto L84
L86:
	;
	goto L87
L87:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v238)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v261)+32)) = v271
	v273 = v271
	goto L84
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v273)+36)) = v261
	goto L90
L89:
	;
	goto L90
L90:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v238)+32)) = int64(0)
	goto L80
L91:
	;
	if v279 != 0 {
		goto L95
	} else {
		goto L96
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v244)+20)) = v279
	goto L91
L93:
	;
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v280)+16)) = v279
	goto L91
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v279)+20)) = v280
	goto L97
L96:
	;
	goto L97
L97:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v244)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v244)+12)) = v286 - int32(1)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v238)+24))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v238)+28))
	if v291 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	if v290 != 0 {
		goto L102
	} else {
		goto L103
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v243)+16)) = v290
	goto L98
L100:
	;
	goto L101
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291)+24)) = v290
	goto L98
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v290)+28)) = v291
	goto L104
L103:
	;
	goto L104
L104:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v243)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v243)+8)) = v297 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v238))) = int32(0)
	v304 = v238 + int32(8)
	v305 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v304)+16)) = v305
	*(*int64)(unsafe.Add(mBase, uint32(v304)+8)) = v305
	*(*int64)(unsafe.Add(mBase, uint32(v304))) = v305
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v238)+16)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v238
	goto L79
L105:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v154)+20))
	if v328 != 0 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v404 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v154)+4)) = uint8(v404)
	*(*int32)(unsafe.Add(mBase, uint32(v154))) = int32(-1)
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v154)+32))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v154)+28))
	if v409 != 0 {
		goto L137
	} else {
		goto L138
	}
L107:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v328)+12))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v328)+8))
	v335 = int32(*(*int16)(unsafe.Add(mBase, uint32(v328)+4)))
	if v335 < int32(0) {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	goto L109
L109:
	;
	goto L106
L110:
	;
	goto L105
L111:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v328)+16))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v328)+20))
	if v370 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L112:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v328)))
	v340 = v338 - int32(97)
	if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v340))|base.B2i32(int32(1)<<(uint(v340)%32)&int32(_a_F_clonesuccessorstates_0) == int32(0)) != 0 {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v350 != 0 {
		goto L111
	} else {
		goto L114
	}
L114:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v328)+36))
	if v351 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	if v363 != 0 {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v354)+20))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v328)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v355+v335*int32(24))+12)) = v359
	v363 = v359
	goto L115
L117:
	;
	goto L118
L118:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v328)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v351)+32)) = v361
	v363 = v361
	goto L115
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v363)+36)) = v351
	goto L121
L120:
	;
	goto L121
L121:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v328)+32)) = int64(0)
	goto L111
L122:
	;
	if v369 != 0 {
		goto L126
	} else {
		goto L127
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v334)+20)) = v369
	goto L122
L124:
	;
	goto L125
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v370)+16)) = v369
	goto L122
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v369)+20)) = v370
	goto L128
L127:
	;
	goto L128
L128:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v334)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v334)+12)) = v376 - int32(1)
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v328)+24))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v328)+28))
	if v381 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	if v380 != 0 {
		goto L133
	} else {
		goto L134
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v333)+16)) = v380
	goto L129
L131:
	;
	goto L132
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v381)+24)) = v380
	goto L129
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v380)+28)) = v381
	goto L135
L134:
	;
	goto L135
L135:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v333)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v333)+8)) = v387 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v328))) = int32(0)
	v394 = v328 + int32(8)
	v395 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v394)+16)) = v395
	*(*int64)(unsafe.Add(mBase, uint32(v394)+8)) = v395
	*(*int64)(unsafe.Add(mBase, uint32(v394))) = v395
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v328)+16)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v328
	goto L110
L136:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v154)+28))
	if v408 != 0 {
		goto L141
	} else {
		goto L142
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v409)+32)) = v408
	goto L136
L138:
	;
	goto L139
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v408
	goto L136
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154)+32)) = int32(0)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v154)+28)) = v417
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v154
	goto L73
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v408)+28)) = v412
	goto L140
L142:
	;
	goto L143
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v412
	goto L140
L144:
	;
	goto L30
L145:
	;
	if v436 == int32(0) {
		goto L25
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v436)+24)) = v75
	F_cparc(m, l0, v69, l2, v436)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	goto L30
L148:
	;
	goto L28
L149:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v474 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L150:
	;
	goto L151
L151:
	;
	return
L152:
	;
	F_pfree(m, v51)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L162
	}
L153:
	;
	v478 = v474
	goto L154
L154:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v491)+12))
	if v492 != 0 {
		goto L152
	} else {
		goto L156
	}
L155:
	;
	goto L152
L156:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v478)+12))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v493)+24))
	if v494 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v495 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v493)+24)) = v495
	F_clonesuccessorstates(m, l0, v494, v493, l3, l4, v495, v51, l7)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L1
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v478)+16))
	if v500 != 0 {
		v478 = v500
		goto L154
	} else {
		goto L161
	}
L160:
	;
	goto L159
L161:
	;
	goto L155
L162:
	;
	goto L151
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
	var v17 int32
	_ = v17
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
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
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v7 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v73
L2:
	;
	if v6 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v40 = base.B2i32(v6 != int32(0))
	goto L4
L4:
	;
	if v40 != 0 {
		v73 = v40
		goto L1
	} else {
		goto L15
	}
L5:
	;
	return int32(-1)
L6:
	;
	goto L7
L7:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if base.B2i32(v14 == int32(0))|base.B2i32(v14 != v17) != 0 {
		v35 = v14
		v36 = v17
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v40 = v35 - v36
	goto L4
L9:
	;
	goto L8
L10:
	;
	v20 = v7
	v21 = v6
	goto L11
L11:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v25 == int32(0) {
		v35 = v25
		v36 = v24
		goto L9
	} else {
		goto L13
	}
L12:
	;
	v35 = v25
	v36 = v24
	goto L9
L13:
	;
	v28 = int32(1)
	if v25 == v24 {
		v20 = v20 + v28
		v21 = v21 + v28
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v41 = int32(0)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v42 == v41 {
		v73 = v41
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v45 == int32(0) {
		v73 = v41
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v48 == v49 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+4)))
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+4)))
	if v51 == v52 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	if base.Ui32(v49) < base.Ui32(v48) {
		goto L31
	} else {
		goto L32
	}
L21:
	;
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+6)))
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+6)))
	if v54 == v55 {
		v73 = v41
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if base.Ui32(v52) < base.Ui32(v51) {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	if base.Ui32(v55) < base.Ui32(v54) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v60 = int32(-1)
	goto L27
L26:
	;
	v60 = int32(1)
	goto L27
L27:
	;
	return v60
L28:
	;
	v65 = int32(-1)
	goto L30
L29:
	;
	v65 = int32(1)
	goto L30
L30:
	;
	return v65
L31:
	;
	v70 = int32(-1)
	goto L33
L32:
	;
	v70 = int32(1)
	goto L33
L33:
	;
	v73 = v70
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
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v172 int32
	_ = v172
	v7 = int32(0)
	if base.B2i32(l5 < l2)&base.B2i32(v7 < l1) == v7 {
		v38 = l2
		v42 = v7
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v172
L2:
	;
	if base.B2i32(l4 <= int32(0))|base.B2i32(l5 <= v38) != 0 {
		v74 = l5
		v76 = v7
		goto L9
	} else {
		goto L10
	}
L3:
	;
	v19 = l2
	v23 = v7
	goto L4
L4:
	;
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+v23<<(uint(int32(1))%32)))))
	if v29 != 0 {
		v172 = int32(1)
		goto L1
	} else {
		goto L6
	}
L5:
	;
	v38 = v33
	v42 = v31
	goto L2
L6:
	;
	v30 = int32(1)
	v31 = v23 + v30
	v33 = v19 - v30
	if v33 <= l5 {
		v38 = v33
		v42 = v31
		goto L2
	} else {
		goto L7
	}
L7:
	;
	if v31 < l1 {
		v19 = v33
		v23 = v31
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	if v38 != v74 {
		v117 = v42
		v118 = v76
		goto L16
	} else {
		goto L17
	}
L10:
	;
	v55 = l5
	v57 = v7
	goto L11
L11:
	;
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3+v57<<(uint(int32(1))%32)))))
	if v62 != 0 {
		v172 = int32(-1)
		goto L1
	} else {
		goto L13
	}
L12:
	;
	v74 = v66
	v76 = v64
	goto L9
L13:
	;
	v63 = int32(1)
	v64 = v57 + v63
	v66 = v55 - v63
	if v66 <= v38 {
		v74 = v66
		v76 = v64
		goto L9
	} else {
		goto L14
	}
L14:
	;
	if v64 < l4 {
		v55 = v66
		v57 = v64
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	if l1 < v117 {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v85 = v42
	v86 = v76
	goto L18
L18:
	;
	if base.B2i32(l1 <= v85)|base.B2i32(l4 <= v86) != 0 {
		v117 = v85
		v118 = v86
		goto L16
	} else {
		goto L20
	}
L19:
	;
	if base.I32_extend16_s(v102) < base.I32_extend16_s(v100) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v91 = int32(1)
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+v85<<(uint(v91)%32)))))
	v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86<<(uint(v91)%32)+l3))))
	if v100 == v102 {
		v85 = v85 + v91
		v86 = v86 + v91
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v109 = int32(1)
	goto L24
L23:
	;
	v109 = int32(-1)
	goto L24
L24:
	;
	return v109
L25:
	;
	v121 = v117
	goto L27
L26:
	;
	v121 = l1
	goto L27
L27:
	;
	v128 = v117
	goto L28
L28:
	;
	if v121 == v128 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v172 = v155
	goto L1
L30:
	;
	if l4 < v118 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v155 = int32(1)
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+v128<<(uint(v155)%32)))))
	if v161 == int32(0) {
		v128 = v128 + v155
		goto L28
	} else {
		goto L42
	}
L33:
	;
	v133 = v118
	goto L35
L34:
	;
	v133 = l4
	goto L35
L35:
	;
	v141 = v118
	goto L36
L36:
	;
	if v133 == v141 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v172 = int32(-1)
	goto L1
L38:
	;
	return int32(0)
L39:
	;
	goto L40
L40:
	;
	v146 = int32(1)
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141<<(uint(v146)%32)+l3))))
	if v151 == int32(0) {
		v141 = v141 + v146
		goto L36
	} else {
		goto L41
	}
L41:
	;
	goto L37
L42:
	;
	goto L29
}
func F_cmp_numerics(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
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
	var v228 int32
	_ = v228
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v307 int32
	_ = v307
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v486 int32
	_ = v486
	var v496 int32
	_ = v496
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v16 = int32(_a_F_cmp_numerics_0)
	v17 = v15 & v16
	if v17 == v16 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v15 != int32(_a_F_cmp_numerics_1) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	if base.Ui32(int32(_a_F_cmp_numerics_0)) <= base.Ui32(v14) {
		goto L15
	} else {
		goto L16
	}
L4:
	;
	if v14 != int32(_a_F_cmp_numerics_2) {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	if v15 != int32(_a_F_cmp_numerics_0) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	if v14 == int32(_a_F_cmp_numerics_0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	return base.B2i32(v14 != int32(_a_F_cmp_numerics_0))
L9:
	;
	v32 = int32(-1)
	goto L11
L10:
	;
	v32 = base.B2i32(v14 != int32(_a_F_cmp_numerics_1))
	goto L11
L11:
	;
	return v32
L12:
	;
	v38 = int32(-1)
	goto L14
L13:
	;
	v38 = int32(0)
	goto L14
L14:
	;
	return v38
L15:
	;
	if v14 == int32(_a_F_cmp_numerics_2) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v49 = l0 + int32(6)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v57 = base.B2i32(int32(0) <= base.I32_extend16_s(v15))
	if int32(0) <= base.I32_extend16_s(v15) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v46 = int32(1)
	goto L20
L19:
	;
	v46 = int32(-1)
	goto L20
L20:
	;
	return v46
L21:
	;
	v58 = int32(-8)
	goto L23
L22:
	;
	v58 = int32(-6)
	goto L23
L23:
	;
	if int32(0) <= base.I32_extend16_s(v15) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v60 = int32(*(*int16)(unsafe.Add(mBase, uint32(v49))))
	v70 = v60
	goto L26
L25:
	;
	v70 = v15<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v15&int32(63)
	goto L26
L26:
	;
	v72 = int32(base.Ui32(int32(base.Ui32(v50)>>(uint(int32(2))%32))+v58) >> (uint(int32(1)) % 32))
	v74 = l1 + int32(6)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v82 = base.B2i32(int32(0) <= base.I32_extend16_s(v14))
	if int32(0) <= base.I32_extend16_s(v14) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v83 = int32(-8)
	goto L29
L28:
	;
	v83 = int32(-6)
	goto L29
L29:
	;
	if int32(0) <= base.I32_extend16_s(v14) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v85 = int32(*(*int16)(unsafe.Add(mBase, uint32(v74))))
	v95 = v85
	goto L32
L31:
	;
	v95 = v14<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v14&int32(63)
	goto L32
L32:
	;
	v96 = int32(1)
	v97 = int32(base.Ui32(int32(base.Ui32(v75)>>(uint(int32(2))%32))+v83) >> (uint(v96) % 32))
	v103 = v14 & int32(_a_F_cmp_numerics_0)
	if v103 == int32(_a_F_cmp_numerics_3) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v106 = v14 << (uint(v96) % 32) & int32(_a_F_cmp_numerics_4)
	goto L35
L34:
	;
	v106 = v103
	goto L35
L35:
	;
	if v72 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if v97 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	if v17 == int32(_a_F_cmp_numerics_3) {
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
	if v106 == int32(_a_F_cmp_numerics_4) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v117 = int32(1)
	goto L44
L43:
	;
	v117 = int32(-1)
	goto L44
L44:
	;
	return v117
L45:
	;
	v125 = v15 << (uint(int32(1)) % 32) & int32(_a_F_cmp_numerics_4)
	goto L47
L46:
	;
	v125 = v17
	goto L47
L47:
	;
	if v97 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	if v125 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	if int32(0) <= base.I32_extend16_s(v15) {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v130 = int32(-1)
	goto L53
L52:
	;
	v130 = int32(1)
	goto L53
L53:
	;
	return v130
L54:
	;
	v134 = l0 + int32(8)
	goto L56
L55:
	;
	v134 = v49
	goto L56
L56:
	;
	if int32(0) <= base.I32_extend16_s(v14) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v137 = l1 + int32(8)
	goto L59
L58:
	;
	v137 = v74
	goto L59
L59:
	;
	if v125 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if v106 == int32(_a_F_cmp_numerics_4) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	if v106 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L63:
	;
	return int32(1)
L64:
	;
	goto L65
L65:
	;
	v144 = int32(0)
	if base.B2i32(v95 < v70)&base.B2i32(v144 < v72) == v144 {
		v175 = v70
		v179 = v144
		goto L68
	} else {
		goto L69
	}
L66:
	;
	return v317
L67:
	;
	v317 = v307
	goto L66
L68:
	;
	if base.B2i32(v97 <= int32(0))|base.B2i32(v95 <= v175) != 0 {
		v211 = v95
		v213 = v144
		goto L75
	} else {
		goto L76
	}
L69:
	;
	v156 = v70
	v160 = v144
	goto L70
L70:
	;
	v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v134+v160<<(uint(int32(1))%32)))))
	if v166 != 0 {
		v307 = int32(1)
		goto L67
	} else {
		goto L72
	}
L71:
	;
	v175 = v170
	v179 = v168
	goto L68
L72:
	;
	v167 = int32(1)
	v168 = v160 + v167
	v170 = v156 - v167
	if v170 <= v95 {
		v175 = v170
		v179 = v168
		goto L68
	} else {
		goto L73
	}
L73:
	;
	if v168 < v72 {
		v156 = v170
		v160 = v168
		goto L70
	} else {
		goto L74
	}
L74:
	;
	goto L71
L75:
	;
	if v175 != v211 {
		v253 = v179
		v254 = v213
		goto L82
	} else {
		goto L83
	}
L76:
	;
	v192 = v95
	v194 = v144
	goto L77
L77:
	;
	v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v137+v194<<(uint(int32(1))%32)))))
	if v199 != 0 {
		v307 = int32(-1)
		goto L67
	} else {
		goto L79
	}
L78:
	;
	v211 = v203
	v213 = v201
	goto L75
L79:
	;
	v200 = int32(1)
	v201 = v194 + v200
	v203 = v192 - v200
	if v203 <= v175 {
		v211 = v203
		v213 = v201
		goto L75
	} else {
		goto L80
	}
L80:
	;
	if v201 < v97 {
		v192 = v203
		v194 = v201
		goto L77
	} else {
		goto L81
	}
L81:
	;
	goto L78
L82:
	;
	if v72 < v253 {
		goto L91
	} else {
		goto L92
	}
L83:
	;
	v222 = v179
	v223 = v213
	goto L84
L84:
	;
	if base.B2i32(v72 <= v222)|base.B2i32(v97 <= v223) != 0 {
		v253 = v222
		v254 = v223
		goto L82
	} else {
		goto L86
	}
L85:
	;
	if base.I32_extend16_s(v239) < base.I32_extend16_s(v237) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v228 = int32(1)
	v237 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v134+v222<<(uint(v228)%32)))))
	v239 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v223<<(uint(v228)%32)+v137))))
	if v237 == v239 {
		v222 = v222 + v228
		v223 = v223 + v228
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v246 = int32(1)
	goto L90
L89:
	;
	v246 = int32(-1)
	goto L90
L90:
	;
	v317 = v246
	goto L66
L91:
	;
	v257 = v253
	goto L93
L92:
	;
	v257 = v72
	goto L93
L93:
	;
	v264 = v253
	goto L94
L94:
	;
	if v257 == v264 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v307 = v290
	goto L67
L96:
	;
	if v97 < v254 {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	goto L98
L98:
	;
	v290 = int32(1)
	v296 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v134+v264<<(uint(v290)%32)))))
	if v296 == int32(0) {
		v264 = v264 + v290
		goto L94
	} else {
		goto L108
	}
L99:
	;
	v269 = v254
	goto L101
L100:
	;
	v269 = v97
	goto L101
L101:
	;
	v277 = v254
	goto L102
L102:
	;
	if v269 == v277 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v307 = int32(-1)
	goto L67
L104:
	;
	v317 = int32(0)
	goto L66
L105:
	;
	goto L106
L106:
	;
	v281 = int32(1)
	v286 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v277<<(uint(v281)%32)+v137))))
	if v286 == int32(0) {
		v277 = v277 + v281
		goto L102
	} else {
		goto L107
	}
L107:
	;
	goto L103
L108:
	;
	goto L95
L109:
	;
	return int32(-1)
L110:
	;
	goto L111
L111:
	;
	v323 = int32(0)
	if base.B2i32(v70 < v95)&base.B2i32(v323 < v97) == v323 {
		v354 = v95
		v358 = v323
		goto L114
	} else {
		goto L115
	}
L112:
	;
	return v496
L113:
	;
	v496 = v486
	goto L112
L114:
	;
	if base.B2i32(v72 <= int32(0))|base.B2i32(v70 <= v354) != 0 {
		v390 = v70
		v392 = v323
		goto L121
	} else {
		goto L122
	}
L115:
	;
	v335 = v95
	v339 = v323
	goto L116
L116:
	;
	v345 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v137+v339<<(uint(int32(1))%32)))))
	if v345 != 0 {
		v486 = int32(1)
		goto L113
	} else {
		goto L118
	}
L117:
	;
	v354 = v349
	v358 = v347
	goto L114
L118:
	;
	v346 = int32(1)
	v347 = v339 + v346
	v349 = v335 - v346
	if v349 <= v70 {
		v354 = v349
		v358 = v347
		goto L114
	} else {
		goto L119
	}
L119:
	;
	if v347 < v97 {
		v335 = v349
		v339 = v347
		goto L116
	} else {
		goto L120
	}
L120:
	;
	goto L117
L121:
	;
	if v354 != v390 {
		v432 = v358
		v433 = v392
		goto L128
	} else {
		goto L129
	}
L122:
	;
	v371 = v70
	v373 = v323
	goto L123
L123:
	;
	v378 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v134+v373<<(uint(int32(1))%32)))))
	if v378 != 0 {
		v486 = int32(-1)
		goto L113
	} else {
		goto L125
	}
L124:
	;
	v390 = v382
	v392 = v380
	goto L121
L125:
	;
	v379 = int32(1)
	v380 = v373 + v379
	v382 = v371 - v379
	if v382 <= v354 {
		v390 = v382
		v392 = v380
		goto L121
	} else {
		goto L126
	}
L126:
	;
	if v380 < v72 {
		v371 = v382
		v373 = v380
		goto L123
	} else {
		goto L127
	}
L127:
	;
	goto L124
L128:
	;
	if v97 < v432 {
		goto L137
	} else {
		goto L138
	}
L129:
	;
	v401 = v358
	v402 = v392
	goto L130
L130:
	;
	if base.B2i32(v97 <= v401)|base.B2i32(v72 <= v402) != 0 {
		v432 = v401
		v433 = v402
		goto L128
	} else {
		goto L132
	}
L131:
	;
	if base.I32_extend16_s(v418) < base.I32_extend16_s(v416) {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	v407 = int32(1)
	v416 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v137+v401<<(uint(v407)%32)))))
	v418 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v402<<(uint(v407)%32)+v134))))
	if v416 == v418 {
		v401 = v401 + v407
		v402 = v402 + v407
		goto L130
	} else {
		goto L133
	}
L133:
	;
	goto L131
L134:
	;
	v425 = int32(1)
	goto L136
L135:
	;
	v425 = int32(-1)
	goto L136
L136:
	;
	v496 = v425
	goto L112
L137:
	;
	v436 = v432
	goto L139
L138:
	;
	v436 = v97
	goto L139
L139:
	;
	v443 = v432
	goto L140
L140:
	;
	if v436 == v443 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v486 = v469
	goto L113
L142:
	;
	if v72 < v433 {
		goto L145
	} else {
		goto L146
	}
L143:
	;
	goto L144
L144:
	;
	v469 = int32(1)
	v475 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v137+v443<<(uint(v469)%32)))))
	if v475 == int32(0) {
		v443 = v443 + v469
		goto L140
	} else {
		goto L154
	}
L145:
	;
	v448 = v433
	goto L147
L146:
	;
	v448 = v72
	goto L147
L147:
	;
	v456 = v433
	goto L148
L148:
	;
	if v448 == v456 {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	v486 = int32(-1)
	goto L113
L150:
	;
	v496 = int32(0)
	goto L112
L151:
	;
	goto L152
L152:
	;
	v460 = int32(1)
	v465 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v456<<(uint(v460)%32)+v134))))
	if v465 == int32(0) {
		v456 = v456 + v460
		goto L148
	} else {
		goto L153
	}
L153:
	;
	goto L149
L154:
	;
	goto L141
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
func F_colorTrgmInfoCmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v8 int64
	_ = v8
	var v10 int64
	_ = v10
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v76 int64
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int64
	_ = v89
	var v90 int32
	_ = v90
	var v100 int64
	_ = v100
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v6 = int64(56)
	v8 = int64(65280)
	v10 = int64(40)
	v13 = int64(16711680)
	v15 = int64(24)
	v17 = int64(4278190080)
	v19 = int64(8)
	v40 = v5<<(uint(v6)%64) | v5&v8<<(uint(v10)%64) | (v5&v13<<(uint(v15)%64) | v5&v17<<(uint(v19)%64)) | (int64(base.Ui64(v5)>>(uint(v19)%64))&v17 | int64(base.Ui64(v5)>>(uint(v15)%64))&v13 | (int64(base.Ui64(v5)>>(uint(v10)%64))&v8 | int64(base.Ui64(v5)>>(uint(v6)%64))))
	v41 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v76 = v41<<(uint(v6)%64) | v41&v8<<(uint(v10)%64) | (v41&v13<<(uint(v15)%64) | v41&v17<<(uint(v19)%64)) | (int64(base.Ui64(v41)>>(uint(v19)%64))&v17 | int64(base.Ui64(v41)>>(uint(v15)%64))&v13 | (int64(base.Ui64(v41)>>(uint(v10)%64))&v8 | int64(base.Ui64(v41)>>(uint(v6)%64))))
	if v40 == v76 {
		v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v80 = int32(16711935)
		v82 = int32(8)
		v84 = int32(24)
		v89 = base.I64_extend_i32_u(base.I32_rotr(v79&v80, v82) | base.I32_rotr(v79, v84)&v80)
		v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v100 = base.I64_extend_i32_u(base.I32_rotr(v90&v80, v82) | base.I32_rotr(v90, v84)&v80)
		if v89 == v100 {
			v112 = int32(0)
		} else {
			v103 = v100
			v104 = v89
			if base.Ui64(v104) < base.Ui64(v103) {
				v108 = int32(-1)
			} else {
				v108 = int32(1)
			}
			v112 = v108
		}
	} else {
		v103 = v76
		v104 = v40
		if base.Ui64(v104) < base.Ui64(v103) {
			v108 = int32(-1)
		} else {
			v108 = int32(1)
		}
		v112 = v108
	}
	return v112
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
			*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v7
			*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v7
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v7
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v7
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
		*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v7
		*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v7
		*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v7
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v7
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
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(993)
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
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(993)
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
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
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
	v12 = int32(0)
	goto L6
L4:
	;
	if v53 != 0 {
		goto L21
	} else {
		goto L22
	}
L5:
	;
	v48 = int32(0)
	if v34 != 0 {
		goto L18
	} else {
		goto L19
	}
L6:
	;
	v16 = int32(0)
	if l0 == v16 {
		v26 = v16
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
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v20 <= v12 {
		v26 = int32(0)
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v26 = v22 + v12<<(uint(int32(2))%32)
	goto L8
L11:
	;
	v32 = int32(0)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if base.B2i32(v26 == v32)|base.B2i32(v34 == v32) != 0 {
		goto L5
	} else {
		goto L16
	}
L12:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v12 < v27 {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v29 = int32(0)
	v53 = base.B2i32(v26 == v29)
	v55 = v29
	goto L4
L15:
	;
	goto L14
L16:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v34+v12<<(uint(int32(2))%32))))
	if v42 == v44 {
		v12 = v12 + int32(1)
		goto L6
	} else {
		goto L17
	}
L17:
	;
	goto L7
L18:
	;
	v52 = int32(2)
	goto L20
L19:
	;
	v52 = v48
	goto L20
L20:
	;
	v53 = base.B2i32(v26 == v48)
	v55 = v52
	goto L4
L21:
	;
	v57 = v55
	goto L23
L22:
	;
	v57 = int32(1)
	goto L23
L23:
	;
	return v57
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
func F_comparecost_2(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	return v3 - v4
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
					base.MemoryFill(m, l0, int32(0), int32(_a_F_compress_free_1))
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
				base.MemoryFill(m, l0, int32(0), int32(_a_F_compress_free_1))
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
		base.MemoryFill(m, l0, int32(0), int32(_a_F_compress_free_1))
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
func F_contains_required_value(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	F_check_stack_depth(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v6 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	if v6 == int32(2) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(1)
L4:
	;
	goto L5
L5:
	;
	v11 = l0
	goto L6
L6:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	switch v12 - int32(33) {
	case 0:
		goto L9
	default:
		goto L10
	case 5:
		goto L11
	}
L7:
	;
	return int32(1)
L8:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L16
	}
L9:
	;
	return int32(0)
L10:
	;
	v25 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+2)))
	v29 = F_contains_required_value(m, v11+v25<<(uint(int32(3))%32))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L14
	}
L11:
	;
	v15 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+2)))
	v19 = F_contains_required_value(m, v11+v15<<(uint(int32(3))%32))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v19 == int32(0) {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	return int32(1)
L14:
	;
	if v29 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	goto L9
L16:
	;
	v36 = v11 - int32(8)
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36))))
	if v37 != int32(2) {
		v11 = v36
		goto L6
	} else {
		goto L17
	}
L17:
	;
	goto L7
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
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
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
	var v125 int32
	_ = v125
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	v4 = l3
	v6 = l5
	if base.Ui32(v4) <= base.Ui32(int32(24)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+11)) = uint8(v4)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v6)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = l4
	v12 = int32(17)
	*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v12)
	v15 = l1 + int32(12)
	if base.Ui32(int32(512)) <= base.Ui32(v4) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	goto L3
L5:
	;
	if v4 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v22 = v15 + v4
	if (v15^l2)&int32(3) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	base.MemoryCopy(m, v15, l2, v4)
	goto L10
L9:
	;
	goto L10
L10:
	;
	goto L4
L11:
	;
	if base.Ui32(v154) < base.Ui32(v22) {
		goto L45
	} else {
		goto L46
	}
L12:
	;
	if v15&int32(3) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	goto L14
L14:
	;
	if base.Ui32(v22) < base.Ui32(int32(4)) {
		goto L36
	} else {
		goto L37
	}
L15:
	;
	v58 = v22 & int32(-4)
	if base.Ui32(v22) < base.Ui32(int32(64)) {
		v108 = v52
		v109 = v53
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v52 = l2
	v53 = v15
	goto L15
L17:
	;
	goto L18
L18:
	;
	if v4 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v52 = l2
	v53 = v15
	goto L15
L20:
	;
	goto L21
L21:
	;
	v35 = l2
	v36 = v15
	goto L22
L22:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	*(*uint8)(unsafe.Add(mBase, uint32(v36))) = uint8(v40)
	v42 = int32(1)
	v43 = v35 + v42
	v45 = v36 + v42
	if v45&int32(3) == int32(0) {
		v52 = v43
		v53 = v45
		goto L15
	} else {
		goto L24
	}
L23:
	;
	v52 = v43
	v53 = v45
	goto L15
L24:
	;
	if base.Ui32(v45) < base.Ui32(v22) {
		v35 = v43
		v36 = v45
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	if base.Ui32(v58) <= base.Ui32(v109) {
		v153 = v108
		v154 = v109
		goto L11
	} else {
		goto L32
	}
L27:
	;
	v62 = v58 + int32(-64)
	if base.Ui32(v62) < base.Ui32(v53) {
		v108 = v52
		v109 = v53
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v65 = v52
	v66 = v53
	goto L29
L29:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v70
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = v72
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+8)) = v74
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+12)) = v76
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+16)) = v78
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+20)) = v80
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v65)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+24)) = v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v65)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+28)) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v65)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+32)) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v65)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+36)) = v88
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v65)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+40)) = v90
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v65)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+44)) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v65)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+48)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v65)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+52)) = v96
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v65)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+56)) = v98
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v65)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+60)) = v100
	v102 = int32(-64)
	v103 = v65 - v102
	v105 = v66 - v102
	if base.Ui32(v105) <= base.Ui32(v62) {
		v65 = v103
		v66 = v105
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v108 = v103
	v109 = v105
	goto L26
L31:
	;
	goto L30
L32:
	;
	v115 = v108
	v116 = v109
	goto L33
L33:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v120
	v122 = int32(4)
	v123 = v115 + v122
	v125 = v116 + v122
	if base.Ui32(v125) < base.Ui32(v58) {
		v115 = v123
		v116 = v125
		goto L33
	} else {
		goto L35
	}
L34:
	;
	v153 = v123
	v154 = v125
	goto L11
L35:
	;
	goto L34
L36:
	;
	v153 = l2
	v154 = v15
	goto L11
L37:
	;
	goto L38
L38:
	;
	if base.Ui32(v4) < base.Ui32(int32(4)) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v153 = l2
	v154 = v15
	goto L11
L40:
	;
	goto L41
L41:
	;
	v134 = l2
	v135 = v15
	goto L42
L42:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	*(*uint8)(unsafe.Add(mBase, uint32(v135))) = uint8(v139)
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)) = uint8(v141)
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v135)+2)) = uint8(v143)
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v135)+3)) = uint8(v145)
	v147 = int32(4)
	v148 = v134 + v147
	v150 = v135 + v147
	if base.Ui32(v150) <= base.Ui32(v22-int32(4)) {
		v134 = v148
		v135 = v150
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v153 = v148
	v154 = v150
	goto L11
L44:
	;
	goto L43
L45:
	;
	v160 = v153
	v161 = v154
	goto L48
L46:
	;
	goto L47
L47:
	;
	goto L4
L48:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	*(*uint8)(unsafe.Add(mBase, uint32(v161))) = uint8(v165)
	v167 = int32(1)
	v170 = v161 + v167
	if v170 != v22 {
		v160 = v160 + v167
		v161 = v170
		goto L48
	} else {
		goto L50
	}
L49:
	;
	goto L47
L50:
	;
	goto L49
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
func F_cosine_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 float32
	_ = v10
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
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
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v59 float32
	_ = v59
	var v60 float32
	_ = v60
	var v61 float32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 float32
	_ = v70
	var v72 float32
	_ = v72
	var v75 float32
	_ = v75
	var v77 float32
	_ = v77
	var v80 float32
	_ = v80
	var v84 float32
	_ = v84
	var v88 float32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v109 float32
	_ = v109
	var v110 float32
	_ = v110
	var v111 float32
	_ = v111
	var v116 int32
	_ = v116
	var v118 float32
	_ = v118
	var v120 float32
	_ = v120
	var v138 float32
	_ = v138
	var v139 float32
	_ = v139
	var v140 float32
	_ = v140
	var v144 float64
	_ = v144
	var v150 float64
	_ = v150
	var v175 float64
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	v10 = float32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = F_pg_detoast_datum(m, v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return int32(0)
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v28 = F_pg_detoast_datum(m, v27)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
			v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+4)))
			if v30 == v31 {
				v35 = base.I32_extend16_s(v30)
				if v35 <= int32(0) {
					v175 = math.Float64frombits(uint64(0x7ff8000000000000))
				} else {
					v38 = int32(8)
					v39 = v28 + v38
					v41 = v23 + v38
					if v35 == int32(1) {
						v98 = int32(0)
						v109 = v10
						v110 = v10
						v111 = v10
						v116 = v98 << (uint(int32(2)) % 32)
						v118 = *(*float32)(unsafe.Add(mBase, uint32(v41+v116)))
						v120 = *(*float32)(unsafe.Add(mBase, uint32(v116+v39)))
						v138 = base.F32_add(base.F32_mul(v118, v120), v109)
						v139 = base.F32_add(base.F32_mul(v120, v120), v110)
						v140 = base.F32_add(base.F32_mul(v118, v118), v111)
					} else {
						v48 = int32(0)
						v56 = int32(0)
						v59 = v10
						v60 = v10
						v61 = v10
						for {
							v65 = int32(2)
							v66 = v48 << (uint(v65) % 32)
							v68 = v66 | int32(4)
							v70 = *(*float32)(unsafe.Add(mBase, uint32(v41+v68)))
							v72 = *(*float32)(unsafe.Add(mBase, uint32(v39+v68)))
							v75 = *(*float32)(unsafe.Add(mBase, uint32(v41+v66)))
							v77 = *(*float32)(unsafe.Add(mBase, uint32(v39+v66)))
							v80 = base.F32_add(base.F32_mul(v70, v72), base.F32_add(base.F32_mul(v75, v77), v59))
							v84 = base.F32_add(base.F32_mul(v72, v72), base.F32_add(base.F32_mul(v77, v77), v60))
							v88 = base.F32_add(base.F32_mul(v70, v70), base.F32_add(base.F32_mul(v75, v75), v61))
							v90 = v48 + v65
							v92 = v56 + v65
							if v92 != v35&int32(_a_F_cosine_distance_0) {
								v48 = v90
								v56 = v92
								v59 = v80
								v60 = v84
								v61 = v88
								continue
							} else {
								break
							}
							break
						}
						if v35&int32(1) == int32(0) {
							v138 = v80
							v139 = v84
							v140 = v88
						} else {
							v98 = v90
							v109 = v80
							v110 = v84
							v111 = v88
							v116 = v98 << (uint(int32(2)) % 32)
							v118 = *(*float32)(unsafe.Add(mBase, uint32(v41+v116)))
							v120 = *(*float32)(unsafe.Add(mBase, uint32(v116+v39)))
							v138 = base.F32_add(base.F32_mul(v118, v120), v109)
							v139 = base.F32_add(base.F32_mul(v120, v120), v110)
							v140 = base.F32_add(base.F32_mul(v118, v118), v111)
						}
					}
					v144 = float64(1)
					v150 = base.F64_div(base.F64_promote_f32(v138), base.F64_sqrt(base.F64_mul(base.F64_promote_f32(v139), base.F64_promote_f32(v140))))
					if base.F64_gt(v150, v144) != 0 {
						v175 = v144
					} else {
						if base.F64_lt(v150, float64(-1)) == int32(0) {
							v175 = v150
						} else {
							v175 = float64(-1)
						}
					}
				}
				v177 = F_Float8GetDatum(m, base.F64_sub(float64(1), v175))
				mBase = m.M
				v178 = m.ExcPending
				if v178 != 0 {
					return int32(0)
				} else {
					m.G0 = v20 + int32(16)
					return v177
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v186 = m.ExcPending
				if v186 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v189 = m.ExcPending
					if v189 != 0 {
						return int32(0)
					} else {
						v190 = int32(*(*int16)(unsafe.Add(mBase, uint32(v23)+4)))
						v191 = int32(*(*int16)(unsafe.Add(mBase, uint32(v28)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v191
						*(*int32)(unsafe.Add(mBase, uint32(v20))) = v190
						F_errmsg(m, int32(_a_F_cosine_distance_1), v20)
						mBase = m.M
						v196 = m.ExcPending
						if v196 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_cosine_distance_2), int32(76), int32(_a_F_cosine_distance_3))
							mBase = m.M
							v201 = m.ExcPending
							if v201 != 0 {
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
func F_cost_subplan(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v52 float64
	_ = v52
	var v53 float64
	_ = v53
	var v56 float64
	_ = v56
	var v63 float64
	_ = v63
	var v64 float64
	_ = v64
	var v65 int32
	_ = v65
	var v69 float64
	_ = v69
	var v70 float64
	_ = v70
	var v74 float64
	_ = v74
	var v75 float64
	_ = v75
	var v76 int32
	_ = v76
	var v77 float64
	_ = v77
	var v78 float64
	_ = v78
	var v87 float64
	_ = v87
	var v91 float64
	_ = v91
	var v94 float64
	_ = v94
	var v95 float64
	_ = v95
	var v98 float64
	_ = v98
	var v106 float64
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v116 float64
	_ = v116
	var v120 float64
	_ = v120
	var v122 float64
	_ = v122
	var v124 float64
	_ = v124
	v3 = float64(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v15 = F_make_ands_implicit(m, v14)
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
	v17 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v17
	v19 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v17
	if v15 == v19 {
		v56 = v3
		v63 = float64(0)
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v64 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v65 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v27 <= int32(0) {
		v56 = v3
		v63 = float64(0)
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v38 = int32(0)
	goto L6
L6:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+v38<<(uint(int32(2))%32))))
	v46 = F_cost_qual_eval_walker(m, v43, v12+int32(8))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v52 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
	v53 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
	v56 = v52
	v63 = v53
	goto L3
L8:
	;
	v49 = v38 + int32(1)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v49 < v50 {
		v38 = v49
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+64)) = v122
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = v124
	m.G0 = v12 + int32(32)
	return
L11:
	;
	v69 = *(*float64)(unsafe.Add(mBase, _c_F_cost_subplan[0]))
	v70 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	v122 = v56
	v124 = base.F64_add(v63, base.F64_add(base.F64_mul(v69, v70), v64))
	goto L10
L12:
	;
	goto L13
L13:
	;
	v74 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v75 = base.F64_sub(v64, v74)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v76 {
	case 0:
		goto L17
	case 1, 2:
		goto L16
	default:
		goto L15
	}
L14:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v107 != 0 {
		v120 = v74
		goto L21
	} else {
		goto L22
	}
L15:
	;
	v106 = base.F64_add(v56, v75)
	goto L14
L16:
	;
	v94 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	v95 = float64(0.5)
	v98 = *(*float64)(unsafe.Add(mBase, _c_F_cost_subplan[0]))
	v106 = base.F64_add(base.F64_mul(base.F64_mul(v94, v95), v98), base.F64_add(base.F64_mul(v75, v95), v56))
	goto L14
L17:
	;
	v77 = float64(1e+100)
	v78 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	if base.F64_gt(v78, v77)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v78)&int64(9223372036854775807))) != 0 {
		v91 = v77
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v106 = base.F64_add(v56, base.F64_div(v75, v91))
	goto L14
L19:
	;
	v87 = float64(1)
	if base.F64_le(v78, v87) != 0 {
		v91 = v87
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v91 = base.F64_nearest(v78)
	goto L18
L21:
	;
	v122 = base.F64_add(v106, v120)
	v124 = v63
	goto L10
L22:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v110 = v108 - int32(348)
	goto L23
L23:
	;
	v116 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	if base.B2i32(base.Ui32(v110) < base.Ui32(int32(15)))&int32(base.Ui32(int32(_a_F_cost_subplan_0))>>(uint(v110)%32)) == int32(0) {
		v120 = v116
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v122 = v106
	v124 = base.F64_add(v63, v116)
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
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
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
				v47 = int32(4)
				if v15 != 0 {
					v51 = v12
				} else {
					v51 = v7 + int32(4)
				}
				v52 = int32(-1)
				if v47 != int32(1) {
					v60 = v51
					v61 = v52
					v62 = int32(0)
					for {
						v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
						v67 = int32(255)
						v69 = int32(2)
						v71 = *(*int32)(unsafe.Add(mBase, uint32((v65^v61)&v67<<(uint(v69)%32))+uint32(_c_F_crc32_bytea[0])))
						v72 = int32(8)
						v74 = v71 ^ int32(base.Ui32(v61)>>(uint(v72)%32))
						v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
						v81 = *(*int32)(unsafe.Add(mBase, uint32((v74^v75)&v67<<(uint(v69)%32))+uint32(_c_F_crc32_bytea[0])))
						v84 = v81 ^ int32(base.Ui32(v74)>>(uint(v72)%32))
						v86 = v60 + v69
						v88 = v62 + v69
						if v88 != v47&int32(-2) {
							v60 = v86
							v61 = v84
							v62 = v88
							continue
						} else {
							break
						}
						break
					}
					if v47&int32(1) == int32(0) {
						v108 = v84
					} else {
						v92 = v86
						v93 = v84
						v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
						v103 = *(*int32)(unsafe.Add(mBase, uint32((v97^v93)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_crc32_bytea[0])))
						v108 = v103 ^ int32(base.Ui32(v93)>>(uint(int32(8))%32))
					}
				} else {
					v92 = v51
					v93 = v52
					v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
					v103 = *(*int32)(unsafe.Add(mBase, uint32((v97^v93)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_crc32_bytea[0])))
					v108 = v103 ^ int32(base.Ui32(v93)>>(uint(int32(8))%32))
				}
				v115 = F_Int64GetDatum(m, base.I64_extend_i32_u(v108^int32(-1)))
				mBase = m.M
				v116 = m.ExcPending
				if v116 != 0 {
					return int32(0)
				} else {
					return v115
				}
			} else {
				if v19 == int32(18) {
					v30 = int32(16)
				} else {
					v30 = int32(0)
				}
				v42 = v30
				if v42 != 0 {
					v47 = v42
					if v15 != 0 {
						v51 = v12
					} else {
						v51 = v7 + int32(4)
					}
					v52 = int32(-1)
					if v47 != int32(1) {
						v60 = v51
						v61 = v52
						v62 = int32(0)
						for {
							v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
							v67 = int32(255)
							v69 = int32(2)
							v71 = *(*int32)(unsafe.Add(mBase, uint32((v65^v61)&v67<<(uint(v69)%32))+uint32(_c_F_crc32_bytea[0])))
							v72 = int32(8)
							v74 = v71 ^ int32(base.Ui32(v61)>>(uint(v72)%32))
							v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
							v81 = *(*int32)(unsafe.Add(mBase, uint32((v74^v75)&v67<<(uint(v69)%32))+uint32(_c_F_crc32_bytea[0])))
							v84 = v81 ^ int32(base.Ui32(v74)>>(uint(v72)%32))
							v86 = v60 + v69
							v88 = v62 + v69
							if v88 != v47&int32(-2) {
								v60 = v86
								v61 = v84
								v62 = v88
								continue
							} else {
								break
							}
							break
						}
						if v47&int32(1) == int32(0) {
							v108 = v84
						} else {
							v92 = v86
							v93 = v84
							v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
							v103 = *(*int32)(unsafe.Add(mBase, uint32((v97^v93)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_crc32_bytea[0])))
							v108 = v103 ^ int32(base.Ui32(v93)>>(uint(int32(8))%32))
						}
					} else {
						v92 = v51
						v93 = v52
						v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
						v103 = *(*int32)(unsafe.Add(mBase, uint32((v97^v93)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_crc32_bytea[0])))
						v108 = v103 ^ int32(base.Ui32(v93)>>(uint(int32(8))%32))
					}
					v115 = F_Int64GetDatum(m, base.I64_extend_i32_u(v108^int32(-1)))
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return int32(0)
					} else {
						return v115
					}
				} else {
					v44 = F_Int64GetDatum(m, int64(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						return v44
					}
				}
			}
		} else {
			v31 = int32(1)
			if v15 != 0 {
				v42 = int32(base.Ui32(v13)>>(uint(v31)%32)) - v31
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v42 = int32(base.Ui32(v35)>>(uint(int32(2))%32)) - int32(4)
			}
			if v42 != 0 {
				v47 = v42
				if v15 != 0 {
					v51 = v12
				} else {
					v51 = v7 + int32(4)
				}
				v52 = int32(-1)
				if v47 != int32(1) {
					v60 = v51
					v61 = v52
					v62 = int32(0)
					for {
						v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
						v67 = int32(255)
						v69 = int32(2)
						v71 = *(*int32)(unsafe.Add(mBase, uint32((v65^v61)&v67<<(uint(v69)%32))+uint32(_c_F_crc32_bytea[0])))
						v72 = int32(8)
						v74 = v71 ^ int32(base.Ui32(v61)>>(uint(v72)%32))
						v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
						v81 = *(*int32)(unsafe.Add(mBase, uint32((v74^v75)&v67<<(uint(v69)%32))+uint32(_c_F_crc32_bytea[0])))
						v84 = v81 ^ int32(base.Ui32(v74)>>(uint(v72)%32))
						v86 = v60 + v69
						v88 = v62 + v69
						if v88 != v47&int32(-2) {
							v60 = v86
							v61 = v84
							v62 = v88
							continue
						} else {
							break
						}
						break
					}
					if v47&int32(1) == int32(0) {
						v108 = v84
					} else {
						v92 = v86
						v93 = v84
						v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
						v103 = *(*int32)(unsafe.Add(mBase, uint32((v97^v93)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_crc32_bytea[0])))
						v108 = v103 ^ int32(base.Ui32(v93)>>(uint(int32(8))%32))
					}
				} else {
					v92 = v51
					v93 = v52
					v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
					v103 = *(*int32)(unsafe.Add(mBase, uint32((v97^v93)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_crc32_bytea[0])))
					v108 = v103 ^ int32(base.Ui32(v93)>>(uint(int32(8))%32))
				}
				v115 = F_Int64GetDatum(m, base.I64_extend_i32_u(v108^int32(-1)))
				mBase = m.M
				v116 = m.ExcPending
				if v116 != 0 {
					return int32(0)
				} else {
					return v115
				}
			} else {
				v44 = F_Int64GetDatum(m, int64(0))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					return v44
				}
			}
		}
	}
}
func F_createTrgmNFA(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
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
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v134 int32
	_ = v134
	var v167 int32
	_ = v167
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
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v339 int32
	_ = v339
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v447 int32
	_ = v447
	var v489 int32
	_ = v489
	var v530 int32
	_ = v530
	var v571 int64
	_ = v571
	var v578 int32
	_ = v578
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v614 int64
	_ = v614
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v633 int32
	_ = v633
	var v650 int32
	_ = v650
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v704 int32
	_ = v704
	var v739 int32
	_ = v739
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v763 int32
	_ = v763
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v828 int32
	_ = v828
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v878 int32
	_ = v878
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v933 int32
	_ = v933
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v957 int32
	_ = v957
	var v961 int32
	_ = v961
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1131 int32
	_ = v1131
	var v1138 int32
	_ = v1138
	var v1166 int32
	_ = v1166
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1190 int32
	_ = v1190
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1229 int32
	_ = v1229
	var v1235 int32
	_ = v1235
	var v1267 int32
	_ = v1267
	var v1271 int32
	_ = v1271
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1283 int64
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1304 int32
	_ = v1304
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1390 int32
	_ = v1390
	var v1394 int32
	_ = v1394
	var __phi1394 int32
	_ = __phi1394
	var v1398 int32
	_ = v1398
	var __phi1398 int32
	_ = __phi1398
	var v1400 int32
	_ = v1400
	var __phi1400 int32
	_ = __phi1400
	var v1433 int64
	_ = v1433
	var v1434 int64
	_ = v1434
	var v1436 int64
	_ = v1436
	var v1438 int64
	_ = v1438
	var v1441 int64
	_ = v1441
	var v1443 int64
	_ = v1443
	var v1445 int64
	_ = v1445
	var v1447 int64
	_ = v1447
	var v1468 int64
	_ = v1468
	var v1469 int64
	_ = v1469
	var v1504 int64
	_ = v1504
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1510 int32
	_ = v1510
	var v1512 int32
	_ = v1512
	var v1517 int64
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1528 int64
	_ = v1528
	var v1531 int64
	_ = v1531
	var v1532 int64
	_ = v1532
	var v1536 int32
	_ = v1536
	var v1540 int32
	_ = v1540
	var v1543 int64
	_ = v1543
	var v1545 int64
	_ = v1545
	var v1547 int64
	_ = v1547
	var v1549 int64
	_ = v1549
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1558 int32
	_ = v1558
	var v1560 int32
	_ = v1560
	var v1578 int32
	_ = v1578
	var v1611 int32
	_ = v1611
	var v1619 int32
	_ = v1619
	var v1647 int64
	_ = v1647
	var v1649 float32
	_ = v1649
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1672 int32
	_ = v1672
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1686 int32
	_ = v1686
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1697 float32
	_ = v1697
	var v1699 float32
	_ = v1699
	var v1701 float32
	_ = v1701
	var v1703 int64
	_ = v1703
	var v1705 int32
	_ = v1705
	var v1710 int32
	_ = v1710
	var v1728 int32
	_ = v1728
	var v1746 int64
	_ = v1746
	var v1748 float32
	_ = v1748
	var v1756 int32
	_ = v1756
	var v1757 int32
	_ = v1757
	var v1760 int32
	_ = v1760
	var v1763 int32
	_ = v1763
	var v1773 int32
	_ = v1773
	var v1777 int32
	_ = v1777
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1849 int32
	_ = v1849
	var v1856 int32
	_ = v1856
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
	var v1897 int32
	_ = v1897
	var v1900 int32
	_ = v1900
	var v1902 int32
	_ = v1902
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1935 int32
	_ = v1935
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1940 int32
	_ = v1940
	var v1946 int32
	_ = v1946
	var v1979 int32
	_ = v1979
	var v1980 int32
	_ = v1980
	var v1982 int32
	_ = v1982
	var v1985 int32
	_ = v1985
	var v1988 int32
	_ = v1988
	var v1995 int32
	_ = v1995
	var v1996 int32
	_ = v1996
	var v1998 int32
	_ = v1998
	var v2000 int32
	_ = v2000
	var v2002 int32
	_ = v2002
	var v2012 int32
	_ = v2012
	var v2047 int32
	_ = v2047
	var v2048 int32
	_ = v2048
	var v2049 int32
	_ = v2049
	var v2050 int32
	_ = v2050
	var v2089 int32
	_ = v2089
	var v2096 int32
	_ = v2096
	var v2129 int32
	_ = v2129
	var v2134 int32
	_ = v2134
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2211 int32
	_ = v2211
	var v2215 int32
	_ = v2215
	var v2216 int32
	_ = v2216
	var v2226 int32
	_ = v2226
	var v2257 int32
	_ = v2257
	var v2267 int32
	_ = v2267
	var v2268 int32
	_ = v2268
	var v2302 int32
	_ = v2302
	var v2303 int32
	_ = v2303
	var v2304 int32
	_ = v2304
	var v2305 int32
	_ = v2305
	var v2344 int32
	_ = v2344
	var v2349 int32
	_ = v2349
	var v2384 int32
	_ = v2384
	var v2386 int32
	_ = v2386
	var v2387 int32
	_ = v2387
	var v2391 int32
	_ = v2391
	var v2392 int32
	_ = v2392
	var v2394 int32
	_ = v2394
	var v2435 int32
	_ = v2435
	var v2437 float32
	_ = v2437
	var v2439 int64
	_ = v2439
	var v2476 int64
	_ = v2476
	var v2478 float32
	_ = v2478
	var v2481 int32
	_ = v2481
	var v2518 int64
	_ = v2518
	var v2524 int32
	_ = v2524
	var v2529 int32
	_ = v2529
	var v2531 int32
	_ = v2531
	var v2532 int32
	_ = v2532
	var v2540 int32
	_ = v2540
	var v2541 int32
	_ = v2541
	var v2544 int32
	_ = v2544
	var v2581 int32
	_ = v2581
	var v2582 int32
	_ = v2582
	var v2588 int32
	_ = v2588
	var v2589 int32
	_ = v2589
	var v2595 int32
	_ = v2595
	var v2596 int32
	_ = v2596
	var v2602 int32
	_ = v2602
	var v2603 int32
	_ = v2603
	var v2609 int32
	_ = v2609
	var v2610 int32
	_ = v2610
	var v2611 int32
	_ = v2611
	var v2613 int32
	_ = v2613
	var v2617 int32
	_ = v2617
	var v2621 int32
	_ = v2621
	var v2656 int32
	_ = v2656
	var v2660 int32
	_ = v2660
	var v2663 int32
	_ = v2663
	var v2697 int32
	_ = v2697
	var v2698 int32
	_ = v2698
	var v2704 int32
	_ = v2704
	var v2705 int32
	_ = v2705
	var v2708 int32
	_ = v2708
	var v2713 int32
	_ = v2713
	var v2714 int32
	_ = v2714
	var v2720 int32
	_ = v2720
	var v2722 int32
	_ = v2722
	var v2724 int32
	_ = v2724
	var v2728 int32
	_ = v2728
	var v2731 int32
	_ = v2731
	var v2737 int32
	_ = v2737
	var v2742 int32
	_ = v2742
	var v2755 int32
	_ = v2755
	var v2788 int32
	_ = v2788
	var v2789 int32
	_ = v2789
	var v2795 int32
	_ = v2795
	var v2808 int32
	_ = v2808
	var v2816 int32
	_ = v2816
	var v2820 int32
	_ = v2820
	var v2830 int32
	_ = v2830
	var v2848 int32
	_ = v2848
	var v2851 int32
	_ = v2851
	var v2852 int32
	_ = v2852
	var v2855 int32
	_ = v2855
	var v2856 int32
	_ = v2856
	var v2859 int32
	_ = v2859
	var v2860 int32
	_ = v2860
	var v2863 int32
	_ = v2863
	var v2866 int32
	_ = v2866
	var v2869 int32
	_ = v2869
	var v2870 int32
	_ = v2870
	var v2871 int32
	_ = v2871
	var v2873 int32
	_ = v2873
	var v2874 int32
	_ = v2874
	var v2878 int32
	_ = v2878
	var v2881 int32
	_ = v2881
	var v2882 int32
	_ = v2882
	var v2883 int32
	_ = v2883
	var v2885 int32
	_ = v2885
	var v2889 int32
	_ = v2889
	var v2892 int32
	_ = v2892
	var v2898 int32
	_ = v2898
	var v2928 int32
	_ = v2928
	var v2929 int32
	_ = v2929
	var v2933 int32
	_ = v2933
	var v2936 int32
	_ = v2936
	var v2939 int32
	_ = v2939
	var v2940 int32
	_ = v2940
	var v2942 int32
	_ = v2942
	var v2944 int32
	_ = v2944
	var v2946 int32
	_ = v2946
	var v2950 int32
	_ = v2950
	var v2954 int32
	_ = v2954
	var v2957 int32
	_ = v2957
	var v2964 int32
	_ = v2964
	var v2993 int32
	_ = v2993
	var v2994 int32
	_ = v2994
	var v2998 int32
	_ = v2998
	var v3000 int32
	_ = v3000
	var v3002 int32
	_ = v3002
	var v3004 int32
	_ = v3004
	var v3007 int32
	_ = v3007
	var v3013 int32
	_ = v3013
	var v3047 int32
	_ = v3047
	var v3048 int32
	_ = v3048
	var v3052 int32
	_ = v3052
	var v3067 int32
	_ = v3067
	var v3069 int32
	_ = v3069
	var v3071 int32
	_ = v3071
	var v3094 int32
	_ = v3094
	var v3096 int32
	_ = v3096
	var v3098 int32
	_ = v3098
	var v3103 int32
	_ = v3103
	var v3110 int32
	_ = v3110
	var v3119 int32
	_ = v3119
	var v3127 int32
	_ = v3127
	var v3129 int32
	_ = v3129
	var v3132 int32
	_ = v3132
	var v3133 int32
	_ = v3133
	var v3139 int32
	_ = v3139
	var v3143 int32
	_ = v3143
	var v3145 int32
	_ = v3145
	var v3147 int32
	_ = v3147
	var v3150 int32
	_ = v3150
	var v3151 int32
	_ = v3151
	var v3153 int32
	_ = v3153
	var v3154 int32
	_ = v3154
	var v3155 int32
	_ = v3155
	var v3156 int32
	_ = v3156
	var v3157 int32
	_ = v3157
	var v3159 int32
	_ = v3159
	var v3165 int32
	_ = v3165
	var v3168 int32
	_ = v3168
	var v3173 int32
	_ = v3173
	var v3176 int32
	_ = v3176
	var v3183 int32
	_ = v3183
	var v3186 int32
	_ = v3186
	var v3191 int32
	_ = v3191
	var v3199 int32
	_ = v3199
	var v3210 int32
	_ = v3210
	var v3214 int32
	_ = v3214
	var v3215 int32
	_ = v3215
	var v3217 int32
	_ = v3217
	var v3219 int32
	_ = v3219
	var v3220 int32
	_ = v3220
	var v3224 int32
	_ = v3224
	var v3227 int32
	_ = v3227
	var v3260 int32
	_ = v3260
	var v3266 int32
	_ = v3266
	var v3269 int32
	_ = v3269
	var v3301 int32
	_ = v3301
	var v3302 int32
	_ = v3302
	var v3306 int32
	_ = v3306
	var v3309 int32
	_ = v3309
	var v3342 int32
	_ = v3342
	var v3351 int32
	_ = v3351
	var v3383 int32
	_ = v3383
	var v3391 int32
	_ = v3391
	var v3395 int32
	_ = v3395
	var v3424 int32
	_ = v3424
	var v3426 int32
	_ = v3426
	var v3441 int32
	_ = v3441
	var v3467 int32
	_ = v3467
	var v3469 int32
	_ = v3469
	var v3470 int32
	_ = v3470
	var v3471 int32
	_ = v3471
	var v3472 int32
	_ = v3472
	var v3477 int32
	_ = v3477
	var v3484 int32
	_ = v3484
	var v3512 int32
	_ = v3512
	var v3513 int32
	_ = v3513
	var v3516 int32
	_ = v3516
	var v3529 int32
	_ = v3529
	var v3532 int32
	_ = v3532
	var v3533 int32
	_ = v3533
	var v3545 int32
	_ = v3545
	var v3573 int32
	_ = v3573
	var v3576 int32
	_ = v3576
	var v3577 int32
	_ = v3577
	var v3579 int32
	_ = v3579
	var v3581 int32
	_ = v3581
	var v3582 int32
	_ = v3582
	var v3583 int32
	_ = v3583
	var v3585 int32
	_ = v3585
	var v3586 int32
	_ = v3586
	var v3587 int32
	_ = v3587
	var v3588 int32
	_ = v3588
	var v3630 int32
	_ = v3630
	var v3665 int32
	_ = v3665
	var v3666 int32
	_ = v3666
	var v3669 int32
	_ = v3669
	var v3670 int32
	_ = v3670
	var v3674 int32
	_ = v3674
	var v3680 int32
	_ = v3680
	var v3682 int32
	_ = v3682
	var v3712 int32
	_ = v3712
	var v3716 int32
	_ = v3716
	var v3717 int32
	_ = v3717
	var v3722 int32
	_ = v3722
	var v3757 int32
	_ = v3757
	var v3758 int32
	_ = v3758
	var v3759 int32
	_ = v3759
	var v3763 int32
	_ = v3763
	var v3764 int32
	_ = v3764
	var v3767 int32
	_ = v3767
	var v3768 int32
	_ = v3768
	var v3770 int32
	_ = v3770
	var v3772 int32
	_ = v3772
	var v3774 int32
	_ = v3774
	var v3777 int32
	_ = v3777
	var v3779 int32
	_ = v3779
	var v3781 int32
	_ = v3781
	var v3784 int32
	_ = v3784
	var v3824 int32
	_ = v3824
	var v3825 int32
	_ = v3825
	var v3829 int32
	_ = v3829
	var v3832 int32
	_ = v3832
	var v3837 int32
	_ = v3837
	var __phi3837 int32
	_ = __phi3837
	var v3841 int32
	_ = v3841
	var __phi3841 int32
	_ = __phi3841
	var v3844 int32
	_ = v3844
	var __phi3844 int32
	_ = __phi3844
	var v3876 int32
	_ = v3876
	var v3877 int32
	_ = v3877
	var v3880 int32
	_ = v3880
	var v3881 int32
	_ = v3881
	var v3884 int32
	_ = v3884
	var v3885 int32
	_ = v3885
	var v3889 int32
	_ = v3889
	var v3891 int64
	_ = v3891
	var v3895 int32
	_ = v3895
	var v3899 int32
	_ = v3899
	var v3903 int32
	_ = v3903
	var v3906 int32
	_ = v3906
	var v3911 int32
	_ = v3911
	var v3913 int32
	_ = v3913
	var v3951 int32
	_ = v3951
	var v3953 int32
	_ = v3953
	var v3954 int32
	_ = v3954
	var v3955 int32
	_ = v3955
	var v3957 int32
	_ = v3957
	var v3961 int32
	_ = v3961
	var v3962 int32
	_ = v3962
	var v3965 int32
	_ = v3965
	var v3966 int32
	_ = v3966
	var v3967 int32
	_ = v3967
	var v3973 int32
	_ = v3973
	var v3977 int32
	_ = v3977
	var v3982 int32
	_ = v3982
	var v4014 int32
	_ = v4014
	var v4015 int32
	_ = v4015
	var v4019 int32
	_ = v4019
	var v4021 int32
	_ = v4021
	var v4022 int32
	_ = v4022
	var v4026 int32
	_ = v4026
	var v4028 int32
	_ = v4028
	var v4029 int32
	_ = v4029
	var v4033 int32
	_ = v4033
	var v4035 int32
	_ = v4035
	var v4036 int32
	_ = v4036
	var v4040 int32
	_ = v4040
	var v4042 int32
	_ = v4042
	var v4043 int32
	_ = v4043
	var v4044 int32
	_ = v4044
	var v4046 int32
	_ = v4046
	var v4050 int32
	_ = v4050
	var v4054 int32
	_ = v4054
	var v4090 int32
	_ = v4090
	var v4094 int32
	_ = v4094
	var v4096 int32
	_ = v4096
	var v4132 int32
	_ = v4132
	var v4136 int32
	_ = v4136
	var v4138 int32
	_ = v4138
	var v4139 int32
	_ = v4139
	var v4142 int32
	_ = v4142
	var v4144 int32
	_ = v4144
	var v4185 int32
	_ = v4185
	var v4186 int32
	_ = v4186
	var v4188 int32
	_ = v4188
	var v4197 int32
	_ = v4197
	var v4201 int32
	_ = v4201
	var v4206 int32
	_ = v4206
	var v4238 int32
	_ = v4238
	var v4239 int32
	_ = v4239
	var v4245 int32
	_ = v4245
	var v4249 int32
	_ = v4249
	var v4250 int32
	_ = v4250
	var v4256 int32
	_ = v4256
	var v4260 int32
	_ = v4260
	var v4261 int32
	_ = v4261
	var v4262 int32
	_ = v4262
	var v4264 int32
	_ = v4264
	var v4268 int32
	_ = v4268
	var v4272 int32
	_ = v4272
	var v4309 int32
	_ = v4309
	var v4310 int32
	_ = v4310
	var v4316 int32
	_ = v4316
	var v4360 int32
	_ = v4360
	var v4361 int32
	_ = v4361
	var v4365 int32
	_ = v4365
	var v4366 int32
	_ = v4366
	var v4369 int32
	_ = v4369
	var v4370 int32
	_ = v4370
	var v4376 int32
	_ = v4376
	var v4381 int32
	_ = v4381
	var v4411 int32
	_ = v4411
	var v4413 int32
	_ = v4413
	var v4420 int32
	_ = v4420
	var v4426 int32
	_ = v4426
	var v4428 int32
	_ = v4428
	var v4463 int32
	_ = v4463
	var v4464 int32
	_ = v4464
	var v4468 int32
	_ = v4468
	var v4469 int32
	_ = v4469
	var v4471 int32
	_ = v4471
	var v4473 int32
	_ = v4473
	var v4476 int32
	_ = v4476
	var v4482 int32
	_ = v4482
	var v4484 int32
	_ = v4484
	var v4519 int32
	_ = v4519
	var v4560 int32
	_ = v4560
	var v4561 int32
	_ = v4561
	var v4562 int32
	_ = v4562
	var v4564 int32
	_ = v4564
	var v4565 int32
	_ = v4565
	var v4566 int32
	_ = v4566
	var v4568 int32
	_ = v4568
	var v4571 int32
	_ = v4571
	var v4572 int32
	_ = v4572
	var v4594 int32
	_ = v4594
	var v4617 int32
	_ = v4617
	v5 = int32(0)
	v40 = m.G0
	v42 = v40 - int32(240)
	m.G0 = v42
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_createTrgmNFA[0]))
	v50 = F_AllocSetContextCreateInternal(m, v45, int32(_a_F_createTrgmNFA_0), v5, int32(_a_F_createTrgmNFA_1), int32(_a_F_createTrgmNFA_2))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v54 = int32(_a_F_createTrgmNFA_3)
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_createTrgmNFA[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_createTrgmNFA[0])) = v50
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v60 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v94 = F_palloc(m, v89<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L14
	}
L4:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v66 == int32(18) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v77 = int32(1)
	if v60&v77 != 0 {
		v89 = int32(base.Ui32(v60)>>(uint(v77)%32)) - v77
		goto L3
	} else {
		goto L13
	}
L7:
	;
	v69 = int32(16)
	goto L9
L8:
	;
	v69 = int32(0)
	goto L9
L9:
	;
	if base.Ui32((v66-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v76 = int32(4)
	goto L12
L11:
	;
	v76 = v69
	goto L12
L12:
	;
	v89 = v76
	goto L3
L13:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v89 = int32(base.Ui32(v83)>>(uint(int32(2))%32)) - int32(4)
	goto L3
L14:
	;
	v96 = int32(1)
	if v60&v96 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v100 = v96
	goto L17
L16:
	;
	v100 = int32(4)
	goto L17
L17:
	;
	v102 = F_pg_mb2wchar_with_len(m, l0+v100, v94, v89)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v105 = F_pg_regcomp(m, v42+int32(32), v94, v102, int32(27), l1)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_pfree(m, v94)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if v105 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_createTrgmNFA[0])) = v55
	F_MemoryContextDelete(m, v50)
	mBase = m.M
	v4617 = m.ExcPending
	if v4617 != 0 {
		goto L1
	} else {
		goto L521
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+224)) = int32(0)
	v2788 = F_MemoryContextAllocZero(m, l3, v2755*int32(3)+int32(5))
	mBase = m.M
	v2789 = m.ExcPending
	if v2789 != 0 {
		goto L1
	} else {
		goto L310
	}
L23:
	;
	v112 = v42 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+64)) = v112
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v112)+24))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+84))
	v117 = v115 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+72)) = v117
	v121 = F_palloc0(m, v117*int32(12))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v2722 = v42 - int32(-64)
	F_pg_regerror(m, v105, v2722)
	mBase = m.M
	v2724 = m.ExcPending
	if v2724 != 0 {
		goto L1
	} else {
		goto L305
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+68)) = v121
	if int32(0) < v117 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v134 = v5
	goto L30
L28:
	;
	goto L29
L29:
	;
	v571 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v42)+88)) = v571
	*(*int64)(unsafe.Add(mBase, uint32(v42)+93)) = v571
	*(*int64)(unsafe.Add(mBase, uint32(v42)+192)) = int64(171798691852)
	v578 = *(*int32)(unsafe.Add(mBase, _c_F_createTrgmNFA[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+216)) = v578
	v585 = F_hash_create(m, int32(_a_F_createTrgmNFA_4), int32(1024), v42+int32(176), int32(1064))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L1
	} else {
		goto L84
	}
L30:
	;
	v167 = v121 + v134*int32(12)
	v170 = int32(-1)
	if v134 <= int32(0) {
		v186 = v170
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L29
L32:
	;
	if base.Ui32(int32(257)) <= base.Ui32(v186) {
		goto L38
	} else {
		goto L39
	}
L33:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v42+int32(32))+24))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+84))
	if base.Ui32(v174) < base.Ui32(v134) {
		v186 = v170
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v173)+92))
	v179 = v176 + v134*int32(24)
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+20)))
	if v180&int32(2) != 0 {
		v186 = v170
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	if v183 != 0 {
		v186 = v170
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	v186 = v184
	goto L32
L37:
	;
	v530 = v134 + int32(1)
	if v530 != v117 {
		v134 = v530
		goto L30
	} else {
		goto L83
	}
L38:
	;
	v189 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v167))) = uint8(v189)
	goto L37
L39:
	;
	goto L40
L40:
	;
	v191 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v167))) = uint16(v191)
	v195 = v186 << (uint(int32(2)) % 32)
	v196 = F_palloc(m, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v167)+8)) = v196
	v203 = F_palloc(m, v195)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v205 = int32(0)
	if base.B2i32(v134 <= v205)|base.B2i32(v186 <= v205) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	if v186 != 0 {
		goto L58
	} else {
		goto L59
	}
L44:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v42+int32(32))+24))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+84))
	if base.Ui32(v211) < base.Ui32(v134) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v210)+92))
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213+v134*int32(24))+20)))
	if v217&int32(2) != 0 {
		goto L43
	} else {
		goto L46
	}
L46:
	;
	v221 = v186
	v228 = int32(0)
	v230 = v203
	goto L47
L47:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v210)+96))
	v264 = int32(*(*int16)(unsafe.Add(mBase, uint32(v260+v228<<(uint(int32(1))%32)))))
	if v264 == v134 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L43
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v230))) = v228
	v268 = v221 - int32(1)
	if v268 == int32(0) {
		goto L43
	} else {
		goto L52
	}
L50:
	;
	v273 = v221
	v274 = v230
	goto L51
L51:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v210)+96))
	v276 = int32(1)
	v277 = v228 | v276
	v281 = int32(*(*int16)(unsafe.Add(mBase, uint32(v275+v277<<(uint(v276)%32)))))
	if v281 == v134 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v273 = v268
	v274 = v230 + int32(4)
	goto L51
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v274))) = v277
	v285 = v273 - int32(1)
	if v285 == int32(0) {
		goto L43
	} else {
		goto L56
	}
L54:
	;
	v290 = v273
	v291 = v274
	goto L55
L55:
	;
	v293 = v228 + int32(2)
	if v293 != int32(2048) {
		v221 = v290
		v228 = v293
		v230 = v291
		goto L47
	} else {
		goto L57
	}
L56:
	;
	v290 = v285
	v291 = v274 + int32(4)
	goto L55
L57:
	;
	goto L48
L58:
	;
	v339 = int32(0)
	goto L61
L59:
	;
	goto L60
L60:
	;
	F_pfree(m, v203)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L82
	}
L61:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v203+v339<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+224)) = v377
	if v377 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L60
L63:
	;
	v447 = v339 + int32(1)
	if v447 != v186 {
		v339 = v447
		goto L61
	} else {
		goto L81
	}
L64:
	;
	v381 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v42)+180)) = uint8(v381)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+176)) = v381
	v386 = v42 + int32(176)
	v390 = F_pg_wchar2mb_with_len(m, v42+int32(224), v386, int32(1))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v393 = F_str_tolower(m, v386, v390, int32(100))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393))))
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386))))
	if base.B2i32(v397 == int32(0))|base.B2i32(v397 != v400) != 0 {
		v418 = v397
		v419 = v400
		goto L68
	} else {
		goto L69
	}
L67:
	;
	F_pfree(m, v393)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L74
	}
L68:
	;
	goto L67
L69:
	;
	v403 = v393
	v404 = v386
	goto L70
L70:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404)+1)))
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403)+1)))
	if v408 == int32(0) {
		v418 = v408
		v419 = v407
		goto L68
	} else {
		goto L72
	}
L71:
	;
	v418 = v408
	v419 = v407
	goto L68
L72:
	;
	v411 = int32(1)
	if v408 == v407 {
		v403 = v403 + v411
		v404 = v404 + v411
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	if v418-v419 != 0 {
		goto L63
	} else {
		goto L75
	}
L75:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v42)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+164)) = v423
	if v390 == int32(0) {
		goto L63
	} else {
		goto L76
	}
L76:
	;
	v429 = F_t_isalnum_with_len(m, v42+int32(164), v390)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	if v429 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v167)+4)) = v431 + int32(1)
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v167)+8))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v42)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v435+v431<<(uint(int32(2))%32)))) = v439
	goto L63
L79:
	;
	goto L80
L80:
	;
	v441 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)) = uint8(v441)
	goto L63
L81:
	;
	goto L62
L82:
	;
	goto L37
L83:
	;
	goto L31
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+84)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+76)) = v585
	*(*int64)(unsafe.Add(mBase, uint32(v42)+164)) = int64(-8589934595)
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v42+int32(32))+24))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v594)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+172)) = v595
	v602 = F_hash_search(m, v585, v42+int32(164), int32(1), v42+int32(224))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+224)))
	if v604 == int32(1) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v1166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1131)+20)))
	if v1166&int32(2) != 0 {
		v4594 = v5
		goto L21
	} else {
		goto L148
	}
L87:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v602)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v602)+20)) = v607 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+80)) = v602
	v1131 = v602
	v1138 = v5
	goto L86
L88:
	;
	goto L89
L89:
	;
	v612 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v602)+20)) = v612
	v614 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v602)+12)) = v614
	*(*int32)(unsafe.Add(mBase, uint32(v42)+84)) = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v602)+32)) = v614
	*(*int64)(unsafe.Add(mBase, uint32(v602)+24)) = int64(4294967295)
	v623 = F_lappend(m, v612, v602)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+88)) = v623
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v602)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v602)+20)) = v626 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+80)) = v602
	if v623 == int32(0) {
		v1131 = v602
		v1138 = v5
		goto L86
	} else {
		goto L91
	}
L91:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v623)+4))
	if v633 <= int32(0) {
		v1131 = v602
		v1138 = v5
		goto L86
	} else {
		goto L92
	}
L92:
	;
	v650 = v5
	goto L93
L93:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v623)+12))
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v675+v650<<(uint(int32(2))%32))))
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+100)))
	if v680 == int32(1) {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v42)+80))
	v1131 = v1126
	v1138 = v1047
	goto L86
L95:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v42)+96))
	if v1047 <= int32(1024) {
		goto L139
	} else {
		goto L140
	}
L96:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v679)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v679)+20)) = v683 | int32(2)
	goto L95
L97:
	;
	goto L98
L98:
	;
	v687 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+92)) = v687
	F_addKey(m, v42-int32(-64), v679, v679)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v42)+92))
	if v694 == int32(0) {
		v763 = v687
		goto L100
	} else {
		goto L101
	}
L100:
	;
	F_list_free(m, v763)
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L1
	} else {
		goto L112
	}
L101:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v694)+4))
	if v697 <= int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v763 = v694
	goto L100
L103:
	;
	goto L104
L104:
	;
	v704 = v687
	goto L105
L105:
	;
	v739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v679)+20)))
	if v739&int32(2) == int32(0) {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v42)+92))
	v763 = v758
	goto L100
L107:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v694)+12))
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v746+v704<<(uint(int32(2))%32))))
	F_addKey(m, v42-int32(-64), v679, v750)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L1
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	goto L106
L110:
	;
	v754 = v704 + int32(1)
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v694)+4))
	if v754 < v755 {
		v704 = v754
		goto L105
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+92)) = int32(0)
	v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v679)+20)))
	if v802&int32(2) != 0 {
		goto L95
	} else {
		goto L113
	}
L113:
	;
	v805 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+232)) = v805
	*(*int64)(unsafe.Add(mBase, uint32(v42)+224)) = int64(0)
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v679)+16))
	if v809 == v805 {
		goto L95
	} else {
		goto L114
	}
L114:
	;
	v812 = int32(0)
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v809)+4))
	if v813 <= v812 {
		goto L95
	} else {
		goto L115
	}
L115:
	;
	v828 = v812
	goto L116
L116:
	;
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v42)+64))
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v809)+12))
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v856+v828<<(uint(int32(2))%32))))
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v860)+8))
	v862 = F_pg_reg_getnumoutarcs(m, v855, v861)
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L1
	} else {
		goto L118
	}
L117:
	;
	goto L95
L118:
	;
	v866 = F_palloc(m, v862<<(uint(int32(3))%32))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v860)+8))
	F_pg_reg_getoutarcs(m, v855, v868, v866, v862)
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	if int32(0) < v862 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v878 = int32(0)
	goto L124
L122:
	;
	goto L123
L123:
	;
	F_pfree(m, v866)
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L1
	} else {
		goto L136
	}
L124:
	;
	v915 = v866 + v878<<(uint(int32(3))%32)
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v915)))
	if v916 < int32(0) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	goto L123
L126:
	;
	v961 = v878 + int32(1)
	if v961 != v862 {
		v878 = v961
		goto L124
	} else {
		goto L135
	}
L127:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v42)+68))
	v922 = v919 + v916*int32(12)
	v923 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v922))))
	if v923 != int32(1) {
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v926 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v922)+1)))
	if v926 == int32(1) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v860)+4))
	v930 = int32(-4)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+228)) = v930
	*(*int32)(unsafe.Add(mBase, uint32(v42)+224)) = v929
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v915)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+232)) = v933
	F_addArc(m, v42-int32(-64), v679, v860, v930, v42+int32(224))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L1
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v922)+4))
	if v943 <= int32(0) {
		goto L126
	} else {
		goto L133
	}
L132:
	;
	goto L131
L133:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v860)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+224)) = v946
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v915)))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+228)) = v948
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v915)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+232)) = v950
	F_addArc(m, v42-int32(-64), v679, v860, v948, v42+int32(224))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	goto L126
L135:
	;
	goto L125
L136:
	;
	v1005 = v828 + int32(1)
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v809)+4))
	if v1005 < v1006 {
		v828 = v1005
		goto L116
	} else {
		goto L137
	}
L137:
	;
	goto L117
L138:
	;
	v1123 = v650 + int32(1)
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v623)+4))
	if v1123 < v1124 {
		v650 = v1123
		goto L93
	} else {
		goto L147
	}
L139:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v42)+76))
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v1050)))
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+4))
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+412))
	if v1054 != 0 {
		goto L143
	} else {
		goto L144
	}
L140:
	;
	goto L141
L141:
	;
	v1120 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v42)+100)) = uint8(v1120)
	goto L138
L142:
	;
	if v1117 < int32(129) {
		goto L138
	} else {
		goto L146
	}
L143:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+376))
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+364))
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+352))
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+340))
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+328))
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+316))
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+304))
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+292))
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+280))
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+268))
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+256))
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+244))
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+232))
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+220))
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+208))
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+196))
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+184))
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+172))
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+160))
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+148))
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+136))
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+124))
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+112))
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+100))
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+88))
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+76))
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+64))
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+52))
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+40))
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+28))
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+16))
	v1117 = v1055 + (v1056 + (v1057 + (v1058 + (v1059 + (v1060 + (v1061 + (v1062 + (v1063 + (v1064 + (v1065 + (v1066 + (v1067 + (v1068 + (v1069 + (v1070 + (v1071 + (v1072 + (v1073 + (v1074 + (v1075 + (v1076 + (v1077 + (v1078 + (v1079 + (v1080 + (v1081 + (v1082 + (v1083 + (v1084 + (v1085 + v1053))))))))))))))))))))))))))))))
	goto L145
L144:
	;
	v1117 = v1053
	goto L145
L145:
	;
	goto L142
L146:
	;
	goto L141
L147:
	;
	goto L94
L148:
	;
	v1170 = v1138 << (uint(int32(5)) % 32)
	v1171 = F_palloc0(m, v1170)
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+104)) = v1171
	v1175 = v42 + int32(176)
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v42)+76))
	F_hash_seq_init(m, v1175, v1176)
	mBase = m.M
	v1178 = m.ExcPending
	if v1178 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v1179 = F_hash_seq_search(m, v1175)
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	if v1179 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v1183 = int32(0)
	v1190 = v1179
	goto L155
L153:
	;
	goto L154
L154:
	;
	if int32(2) <= v1138 {
		goto L167
	} else {
		goto L168
	}
L155:
	;
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(v1190)+12))
	if v1221 == int32(0) {
		v1304 = v1183
		goto L157
	} else {
		goto L158
	}
L156:
	;
	goto L154
L157:
	;
	v1344 = F_hash_seq_search(m, v42+int32(176))
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L1
	} else {
		goto L165
	}
L158:
	;
	v1224 = int32(0)
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v1221)+4))
	if v1225 <= v1224 {
		v1304 = v1183
		goto L157
	} else {
		goto L159
	}
L159:
	;
	v1229 = v1183
	v1235 = v1224
	goto L160
L160:
	;
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v1221)+12))
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v1267+v1235<<(uint(int32(2))%32))))
	v1273 = F_palloc(m, int32(8))
	mBase = m.M
	v1274 = m.ExcPending
	if v1274 != 0 {
		goto L1
	} else {
		goto L162
	}
L161:
	;
	v1304 = v1298
	goto L157
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1273))) = v1190
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v1271)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1273)+4)) = v1276
	v1280 = v1171 + v1229<<(uint(int32(5))%32)
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v1271)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1280)+8)) = v1281
	v1283 = *(*int64)(unsafe.Add(mBase, uint32(v1271)))
	*(*int64)(unsafe.Add(mBase, uint32(v1280))) = v1283
	v1285 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1280)+24)) = uint8(v1285)
	*(*int32)(unsafe.Add(mBase, uint32(v1280)+12)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+12)) = v1273
	*(*int32)(unsafe.Add(mBase, uint32(v42)+224)) = v1273
	v1294 = F_list_make1_impl(m, v1285, v42+int32(12))
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1280)+28)) = v1294
	v1297 = int32(1)
	v1298 = v1229 + v1297
	v1300 = v1235 + v1297
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(v1221)+4))
	if v1300 < v1301 {
		v1229 = v1298
		v1235 = v1300
		goto L160
	} else {
		goto L164
	}
L164:
	;
	goto L161
L165:
	;
	if v1344 != 0 {
		v1183 = v1304
		v1190 = v1344
		goto L155
	} else {
		goto L166
	}
L166:
	;
	goto L156
L167:
	;
	F_pg_qsort(m, v1171, v1138, int32(32), int32(_a_F_createTrgmNFA_5))
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L1
	} else {
		goto L170
	}
L168:
	;
	v1578 = v1138
	goto L169
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+108)) = v1578
	if int32(0) < v1578 {
		goto L187
	} else {
		goto L188
	}
L170:
	;
	__phi1394 = v1171 + int32(32)
	__phi1398 = v1171
	__phi1400 = v1171
	v1394 = __phi1394
	v1398 = __phi1398
	v1400 = __phi1400
	goto L171
L171:
	;
	v1433 = *(*int64)(unsafe.Add(mBase, uint32(v1400)+32))
	v1434 = int64(56)
	v1436 = int64(65280)
	v1438 = int64(40)
	v1441 = int64(16711680)
	v1443 = int64(24)
	v1445 = int64(4278190080)
	v1447 = int64(8)
	v1468 = v1433<<(uint(v1434)%64) | v1433&v1436<<(uint(v1438)%64) | (v1433&v1441<<(uint(v1443)%64) | v1433&v1445<<(uint(v1447)%64)) | (int64(base.Ui64(v1433)>>(uint(v1447)%64))&v1445 | int64(base.Ui64(v1433)>>(uint(v1443)%64))&v1441 | (int64(base.Ui64(v1433)>>(uint(v1438)%64))&v1436 | int64(base.Ui64(v1433)>>(uint(v1434)%64))))
	v1469 = *(*int64)(unsafe.Add(mBase, uint32(v1398)))
	v1504 = v1469<<(uint(v1434)%64) | v1469&v1436<<(uint(v1438)%64) | (v1469&v1441<<(uint(v1443)%64) | v1469&v1445<<(uint(v1447)%64)) | (int64(base.Ui64(v1469)>>(uint(v1447)%64))&v1445 | int64(base.Ui64(v1469)>>(uint(v1443)%64))&v1441 | (int64(base.Ui64(v1469)>>(uint(v1438)%64))&v1436 | int64(base.Ui64(v1469)>>(uint(v1434)%64))))
	if v1468 == v1504 {
		goto L175
	} else {
		goto L176
	}
L172:
	;
	v1578 = (v1558-v1171)>>(uint(int32(5))%32) + int32(1)
	goto L169
L173:
	;
	v1560 = v1394 + int32(32)
	if base.Ui32(v1560) < base.Ui32(v1171+v1170) {
		__phi1394 = v1560
		__phi1398 = v1558
		__phi1400 = v1394
		v1394 = __phi1394
		v1398 = __phi1398
		v1400 = __phi1400
		goto L171
	} else {
		goto L186
	}
L174:
	;
	if int32(0) < v1540 {
		goto L182
	} else {
		goto L183
	}
L175:
	;
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(v1394)+8))
	v1508 = int32(16711935)
	v1510 = int32(8)
	v1512 = int32(24)
	v1517 = base.I64_extend_i32_u(base.I32_rotr(v1507&v1508, v1510) | base.I32_rotr(v1507, v1512)&v1508)
	v1518 = *(*int32)(unsafe.Add(mBase, uint32(v1398)+8))
	v1528 = base.I64_extend_i32_u(base.I32_rotr(v1518&v1508, v1510) | base.I32_rotr(v1518, v1512)&v1508)
	if v1517 == v1528 {
		v1540 = int32(0)
		goto L174
	} else {
		goto L178
	}
L176:
	;
	v1531 = v1504
	v1532 = v1468
	goto L177
L177:
	;
	if base.Ui64(v1532) < base.Ui64(v1531) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	v1531 = v1528
	v1532 = v1517
	goto L177
L179:
	;
	v1536 = int32(-1)
	goto L181
L180:
	;
	v1536 = int32(1)
	goto L181
L181:
	;
	v1540 = v1536
	goto L174
L182:
	;
	v1543 = *(*int64)(unsafe.Add(mBase, uint32(v1394)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1398)+56)) = v1543
	v1545 = *(*int64)(unsafe.Add(mBase, uint32(v1394)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1398)+48)) = v1545
	v1547 = *(*int64)(unsafe.Add(mBase, uint32(v1394)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1398)+40)) = v1547
	v1549 = *(*int64)(unsafe.Add(mBase, uint32(v1394)))
	*(*int64)(unsafe.Add(mBase, uint32(v1398)+32)) = v1549
	v1558 = v1398 + int32(32)
	goto L173
L183:
	;
	goto L184
L184:
	;
	v1553 = *(*int32)(unsafe.Add(mBase, uint32(v1398)+28))
	v1554 = *(*int32)(unsafe.Add(mBase, uint32(v1400)+60))
	v1555 = F_list_concat(m, v1553, v1554)
	mBase = m.M
	v1556 = m.ExcPending
	if v1556 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1398)+28)) = v1555
	v1558 = v1398
	goto L173
L186:
	;
	goto L172
L187:
	;
	v1611 = *(*int32)(unsafe.Add(mBase, uint32(v42)+68))
	v1619 = int32(0)
	v1647 = int64(0)
	v1649 = float32(0)
	goto L190
L188:
	;
	goto L189
L189:
	;
	F_pg_qsort(m, v1171, v1578, int32(32), int32(_a_F_createTrgmNFA_6))
	mBase = m.M
	v2713 = m.ExcPending
	if v2713 != 0 {
		goto L1
	} else {
		goto L303
	}
L190:
	;
	v1655 = v1171 + v1619<<(uint(int32(5))%32)
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(v1655)))
	if v1656 != int32(-4) {
		goto L192
	} else {
		goto L193
	}
L191:
	;
	F_pg_qsort(m, v1171, v1578, int32(32), int32(_a_F_createTrgmNFA_6))
	mBase = m.M
	v1710 = m.ExcPending
	if v1710 != 0 {
		goto L1
	} else {
		goto L204
	}
L192:
	;
	v1663 = *(*int32)(unsafe.Add(mBase, uint32(v1611+v1656*int32(12))+4))
	v1664 = v1663
	v1665 = int32(0)
	goto L194
L193:
	;
	v1664 = int32(1)
	v1665 = int32(2)
	goto L194
L194:
	;
	v1666 = *(*int32)(unsafe.Add(mBase, uint32(v1655)+4))
	if v1666 != int32(-4) {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	v1679 = v1677 << (uint(int32(1)) % 32)
	v1680 = *(*int32)(unsafe.Add(mBase, uint32(v1655)+8))
	if v1680 != int32(-4) {
		goto L200
	} else {
		goto L201
	}
L196:
	;
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(v1611+v1666*int32(12))+4))
	v1676 = v1672 * v1664
	v1677 = v1665
	goto L195
L197:
	;
	goto L198
L198:
	;
	v1676 = v1664
	v1677 = v1665 | int32(1)
	goto L195
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1655)+16)) = v1690
	v1697 = *(*float32)(unsafe.Add(mBase, uint32(v1691<<(uint(int32(2))%32))+uint32(_c_F_createTrgmNFA[1])))
	v1699 = base.F32_mul(v1697, base.F32_convert_i32_s(v1690))
	*(*float32)(unsafe.Add(mBase, uint32(v1655)+20)) = v1699
	v1701 = base.F32_add(v1649, v1699)
	v1703 = v1647 + base.I64_extend_i32_s(v1690)
	v1705 = v1619 + int32(1)
	if v1705 != v1578 {
		v1619 = v1705
		v1647 = v1703
		v1649 = v1701
		goto L190
	} else {
		goto L203
	}
L200:
	;
	v1686 = *(*int32)(unsafe.Add(mBase, uint32(v1611+v1680*int32(12))+4))
	v1690 = v1686 * v1676
	v1691 = v1679
	goto L199
L201:
	;
	goto L202
L202:
	;
	v1690 = v1676
	v1691 = v1679 | int32(1)
	goto L199
L203:
	;
	goto L191
L204:
	;
	v1728 = v5
	v1746 = v1703
	v1748 = v1701
	goto L205
L205:
	;
	if base.F32_le(v1748, float32(16)) == int32(0) {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	if v2518 <= int64(256) {
		goto L274
	} else {
		goto L275
	}
L207:
	;
	v1756 = v1171 + v1728<<(uint(int32(5))%32)
	v1757 = *(*int32)(unsafe.Add(mBase, uint32(v1756)+28))
	if v1757 == int32(0) {
		goto L211
	} else {
		goto L212
	}
L208:
	;
	v2518 = v1746
	goto L209
L209:
	;
	goto L206
L210:
	;
	v2481 = v1728 + int32(1)
	if v2481 != v1578 {
		v1728 = v2481
		v1746 = v2476
		v1748 = v2478
		goto L205
	} else {
		goto L273
	}
L211:
	;
	v2435 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1756)+24)) = uint8(v2435)
	v2437 = *(*float32)(unsafe.Add(mBase, uint32(v1756)+20))
	v2439 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1756)+16)))
	v2476 = v1746 - v2439
	v2478 = base.F32_sub(v1748, v2437)
	goto L210
L212:
	;
	v1760 = *(*int32)(unsafe.Add(mBase, uint32(v1757)+4))
	if v1760 <= int32(0) {
		goto L211
	} else {
		goto L213
	}
L213:
	;
	v1763 = *(*int32)(unsafe.Add(mBase, uint32(v1757)+12))
	v1773 = v1760
	v1777 = int32(0)
	goto L214
L214:
	;
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(v1763+v1777<<(uint(int32(2))%32))))
	v1808 = *(*int32)(unsafe.Add(mBase, uint32(v1807)+4))
	v1809 = *(*int32)(unsafe.Add(mBase, uint32(v1807)))
	v1810 = v1809
	goto L216
L215:
	;
	v2002 = int32(0)
	if v2002 < v2000 {
		goto L239
	} else {
		goto L240
	}
L216:
	;
	v1849 = *(*int32)(unsafe.Add(mBase, uint32(v1810)+28))
	if v1849 != 0 {
		v1810 = v1849
		goto L216
	} else {
		goto L218
	}
L217:
	;
	v1856 = v1808
	goto L219
L218:
	;
	goto L217
L219:
	;
	v1889 = *(*int32)(unsafe.Add(mBase, uint32(v1856)+28))
	if v1889 != 0 {
		v1856 = v1889
		goto L219
	} else {
		goto L221
	}
L220:
	;
	v1890 = *(*int32)(unsafe.Add(mBase, uint32(v1810)+32))
	v1891 = *(*int32)(unsafe.Add(mBase, uint32(v1810)+20))
	v1897 = v1810
	v1900 = v1890 | v1891
	v1902 = v1890
	goto L222
L221:
	;
	goto L220
L222:
	;
	v1932 = *(*int32)(unsafe.Add(mBase, uint32(v1897)+36))
	if v1932 != 0 {
		goto L224
	} else {
		goto L225
	}
L223:
	;
	v1937 = *(*int32)(unsafe.Add(mBase, uint32(v1856)+32))
	v1938 = *(*int32)(unsafe.Add(mBase, uint32(v1856)+20))
	v1940 = v1856
	v1946 = v1937 | v1938
	goto L227
L224:
	;
	v1933 = *(*int32)(unsafe.Add(mBase, uint32(v1932)+20))
	v1935 = *(*int32)(unsafe.Add(mBase, uint32(v1932)+32))
	v1897 = v1932
	v1900 = v1933 | v1900 | v1935
	v1902 = v1935
	goto L222
L225:
	;
	goto L226
L226:
	;
	goto L223
L227:
	;
	v1979 = *(*int32)(unsafe.Add(mBase, uint32(v1940)+36))
	if v1979 != 0 {
		goto L229
	} else {
		goto L230
	}
L228:
	;
	v1985 = int32(3)
	v1988 = base.B2i32((v1946|v1900)&v1985 == v1985)
	if v1988 == int32(0) {
		goto L232
	} else {
		goto L233
	}
L229:
	;
	v1980 = *(*int32)(unsafe.Add(mBase, uint32(v1979)+20))
	v1982 = *(*int32)(unsafe.Add(mBase, uint32(v1979)+32))
	v1940 = v1979
	v1946 = v1980 | v1946 | v1982
	goto L227
L230:
	;
	goto L231
L231:
	;
	goto L228
L232:
	;
	if v1940 != v1897 {
		goto L235
	} else {
		goto L236
	}
L233:
	;
	v2000 = v1773
	goto L234
L234:
	;
	goto L215
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1940)+36)) = v1897
	*(*int32)(unsafe.Add(mBase, uint32(v1897)+32)) = v1946 | v1902
	v1995 = *(*int32)(unsafe.Add(mBase, uint32(v1757)+4))
	v1996 = v1995
	goto L237
L236:
	;
	v1996 = v1773
	goto L237
L237:
	;
	v1998 = v1777 + int32(1)
	if v1998 < v1996 {
		v1773 = v1996
		v1777 = v1998
		goto L214
	} else {
		goto L238
	}
L238:
	;
	v2000 = v1996
	goto L234
L239:
	;
	v2012 = v2002
	goto L242
L240:
	;
	v2226 = v2000
	goto L241
L241:
	;
	if (v1946|v1900)&v1985 == v1985 {
		v2476 = v1746
		v2478 = v1748
		goto L210
	} else {
		goto L259
	}
L242:
	;
	v2047 = *(*int32)(unsafe.Add(mBase, uint32(v1763+v2012<<(uint(int32(2))%32))))
	v2048 = *(*int32)(unsafe.Add(mBase, uint32(v2047)+4))
	v2049 = *(*int32)(unsafe.Add(mBase, uint32(v2047)))
	v2050 = v2049
	goto L244
L243:
	;
	v2226 = v2216
	goto L241
L244:
	;
	v2089 = *(*int32)(unsafe.Add(mBase, uint32(v2050)+28))
	if v2089 != 0 {
		v2050 = v2089
		goto L244
	} else {
		goto L246
	}
L245:
	;
	v2096 = v2048
	goto L247
L246:
	;
	goto L245
L247:
	;
	v2129 = *(*int32)(unsafe.Add(mBase, uint32(v2096)+28))
	if v2129 != 0 {
		v2096 = v2129
		goto L247
	} else {
		goto L249
	}
L248:
	;
	v2134 = v2050
	goto L250
L249:
	;
	goto L248
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2134)+32)) = int32(0)
	v2171 = *(*int32)(unsafe.Add(mBase, uint32(v2134)+36))
	if v2171 != 0 {
		v2134 = v2171
		goto L250
	} else {
		goto L252
	}
L251:
	;
	v2172 = v2096
	goto L253
L252:
	;
	goto L251
L253:
	;
	v2211 = *(*int32)(unsafe.Add(mBase, uint32(v2172)+36))
	if v2211 != 0 {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	v2215 = v2012 + int32(1)
	v2216 = *(*int32)(unsafe.Add(mBase, uint32(v1757)+4))
	if v2215 < v2216 {
		v2012 = v2215
		goto L242
	} else {
		goto L258
	}
L255:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2172)+32)) = int64(0)
	v2172 = v2211
	goto L253
L256:
	;
	goto L257
L257:
	;
	goto L254
L258:
	;
	goto L243
L259:
	;
	v2257 = int32(0)
	if v2226 <= v2257 {
		goto L211
	} else {
		goto L260
	}
L260:
	;
	v2267 = v2257
	v2268 = v2226
	goto L261
L261:
	;
	v2302 = *(*int32)(unsafe.Add(mBase, uint32(v1763+v2267<<(uint(int32(2))%32))))
	v2303 = *(*int32)(unsafe.Add(mBase, uint32(v2302)+4))
	v2304 = *(*int32)(unsafe.Add(mBase, uint32(v2302)))
	v2305 = v2304
	goto L263
L262:
	;
	goto L211
L263:
	;
	v2344 = *(*int32)(unsafe.Add(mBase, uint32(v2305)+28))
	if v2344 != 0 {
		v2305 = v2344
		goto L263
	} else {
		goto L265
	}
L264:
	;
	v2349 = v2303
	goto L266
L265:
	;
	goto L264
L266:
	;
	v2384 = *(*int32)(unsafe.Add(mBase, uint32(v2349)+28))
	if v2384 != 0 {
		v2349 = v2384
		goto L266
	} else {
		goto L268
	}
L267:
	;
	if v2349 != v2305 {
		goto L269
	} else {
		goto L270
	}
L268:
	;
	goto L267
L269:
	;
	v2386 = *(*int32)(unsafe.Add(mBase, uint32(v2305)+20))
	v2387 = *(*int32)(unsafe.Add(mBase, uint32(v2349)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2305)+20)) = v2386 | v2387
	*(*int32)(unsafe.Add(mBase, uint32(v2349)+28)) = v2305
	v2391 = *(*int32)(unsafe.Add(mBase, uint32(v1757)+4))
	v2392 = v2391
	goto L271
L270:
	;
	v2392 = v2268
	goto L271
L271:
	;
	v2394 = v2267 + int32(1)
	if v2394 < v2392 {
		v2267 = v2394
		v2268 = v2392
		goto L261
	} else {
		goto L272
	}
L272:
	;
	goto L262
L273:
	;
	v2518 = v2476
	goto L209
L274:
	;
	v2524 = base.I32_wrap_i64(v2518)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+112)) = v2524
	F_pg_qsort(m, v1171, v1578, int32(32), int32(_a_F_createTrgmNFA_5))
	mBase = m.M
	v2529 = m.ExcPending
	if v2529 != 0 {
		goto L1
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	v4594 = v5
	goto L21
L277:
	;
	v2531 = v1578 & int32(3)
	v2532 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v1578) {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v2540 = v2532
	v2541 = int32(0)
	v2544 = v2532
	goto L281
L279:
	;
	v2617 = v2532
	v2621 = v2532
	goto L280
L280:
	;
	v2656 = v2617
	v2660 = v2621
	v2663 = v2532
	goto L297
L281:
	;
	v2581 = v1171 + v2540<<(uint(int32(5))%32)
	v2582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2581)+24)))
	if v2582 == int32(1) {
		goto L283
	} else {
		goto L284
	}
L282:
	;
	if v2531 == int32(0) {
		v2755 = v2524
		goto L22
	} else {
		goto L296
	}
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2581)+12)) = v2544
	v2588 = v2544 + int32(1)
	goto L285
L284:
	;
	v2588 = v2544
	goto L285
L285:
	;
	v2589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2581)+56)))
	if v2589 == int32(1) {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2581)+44)) = v2588
	v2595 = v2588 + int32(1)
	goto L288
L287:
	;
	v2595 = v2588
	goto L288
L288:
	;
	v2596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2581)+88)))
	if v2596 == int32(1) {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2581)+76)) = v2595
	v2602 = v2595 + int32(1)
	goto L291
L290:
	;
	v2602 = v2595
	goto L291
L291:
	;
	v2603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2581)+120)))
	if v2603 == int32(1) {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2581)+108)) = v2602
	v2609 = v2602 + int32(1)
	goto L294
L293:
	;
	v2609 = v2602
	goto L294
L294:
	;
	v2610 = int32(4)
	v2611 = v2540 + v2610
	v2613 = v2541 + v2610
	if v2613 != v1578&int32(2147483644) {
		v2540 = v2611
		v2541 = v2613
		v2544 = v2609
		goto L281
	} else {
		goto L295
	}
L295:
	;
	goto L282
L296:
	;
	v2617 = v2611
	v2621 = v2609
	goto L280
L297:
	;
	v2697 = v1171 + v2656<<(uint(int32(5))%32)
	v2698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2697)+24)))
	if v2698 == int32(1) {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	v2755 = v2524
	goto L22
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2697)+12)) = v2660
	v2704 = v2660 + int32(1)
	goto L301
L300:
	;
	v2704 = v2660
	goto L301
L301:
	;
	v2705 = int32(1)
	v2708 = v2663 + v2705
	if v2708 != v2531 {
		v2656 = v2656 + v2705
		v2660 = v2704
		v2663 = v2708
		goto L297
	} else {
		goto L302
	}
L302:
	;
	goto L298
L303:
	;
	v2714 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+112)) = v2714
	F_pg_qsort(m, v1171, v1578, int32(32), int32(_a_F_createTrgmNFA_5))
	mBase = m.M
	v2720 = m.ExcPending
	if v2720 != 0 {
		goto L1
	} else {
		goto L304
	}
L304:
	;
	v2755 = v2714
	goto L22
L305:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2728 = m.ExcPending
	if v2728 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	F_errcode(m, int32(302252162))
	mBase = m.M
	v2731 = m.ExcPending
	if v2731 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = v2722
	F_errmsg(m, int32(_a_F_createTrgmNFA_7), v42+int32(16))
	mBase = m.M
	v2737 = m.ExcPending
	if v2737 != 0 {
		goto L1
	} else {
		goto L308
	}
L308:
	;
	F_errfinish(m, int32(_a_F_createTrgmNFA_8), int32(751), int32(_a_F_createTrgmNFA_9))
	mBase = m.M
	v2742 = m.ExcPending
	if v2742 != 0 {
		goto L1
	} else {
		goto L309
	}
L309:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2788))) = v2755*int32(12) + int32(20)
	v2795 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2788)+4)) = uint8(v2795)
	if int32(0) < v1578 {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	v2808 = v42 + int32(177)
	v2816 = v2788 + int32(5)
	v2820 = v1578
	v2830 = v5
	goto L314
L312:
	;
	v3441 = v1176
	goto L313
L313:
	;
	v3467 = v42 + int32(176)
	F_hash_seq_init(m, v3467, v3441)
	mBase = m.M
	v3469 = m.ExcPending
	if v3469 != 0 {
		goto L1
	} else {
		goto L397
	}
L314:
	;
	v2848 = *(*int32)(unsafe.Add(mBase, uint32(v42)+104))
	v2851 = v2848 + v2830<<(uint(int32(5))%32)
	v2852 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2851)+24)))
	if v2852 != int32(1) {
		v3391 = v2816
		v3395 = v2820
		goto L316
	} else {
		goto L317
	}
L315:
	;
	v3426 = *(*int32)(unsafe.Add(mBase, uint32(v42)+76))
	v3441 = v3426
	goto L313
L316:
	;
	v3424 = v2830 + int32(1)
	if v3424 < v3395 {
		v2816 = v3391
		v2820 = v3395
		v2830 = v3424
		goto L314
	} else {
		goto L396
	}
L317:
	;
	v2855 = *(*int32)(unsafe.Add(mBase, uint32(v42)+68))
	v2856 = *(*int32)(unsafe.Add(mBase, uint32(v2851)))
	v2859 = v2855 + v2856*int32(12)
	v2860 = int32(1)
	v2863 = base.B2i32(v2856 == int32(-4))
	if v2863 == int32(0) {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v2866 = *(*int32)(unsafe.Add(mBase, uint32(v2859)+4))
	if v2866 <= int32(0) {
		v3391 = v2816
		v3395 = v2820
		goto L316
	} else {
		goto L321
	}
L319:
	;
	v2869 = v2860
	goto L320
L320:
	;
	v2870 = *(*int32)(unsafe.Add(mBase, uint32(v2851)+4))
	v2871 = int32(12)
	v2873 = v2855 + v2870*v2871
	v2874 = *(*int32)(unsafe.Add(mBase, uint32(v2851)+8))
	v2878 = base.B2i32(v2870 == int32(-4))
	if v2878 == int32(0) {
		goto L322
	} else {
		goto L323
	}
L321:
	;
	v2869 = v2866
	goto L320
L322:
	;
	v2881 = *(*int32)(unsafe.Add(mBase, uint32(v2873)+4))
	v2882 = v2881
	goto L324
L323:
	;
	v2882 = v2860
	goto L324
L324:
	;
	v2883 = v2874*v2871 + v2855
	v2885 = v2869
	v2889 = v2882
	v2892 = v2816
	v2898 = int32(0)
	goto L325
L325:
	;
	if int32(0) < v2889 {
		goto L328
	} else {
		goto L329
	}
L326:
	;
	v3383 = *(*int32)(unsafe.Add(mBase, uint32(v42)+108))
	v3391 = v3351
	v3395 = v3383
	goto L316
L327:
	;
	goto L326
L328:
	;
	if v2856 == int32(-4) {
		goto L331
	} else {
		goto L332
	}
L329:
	;
	v3302 = v2885
	v3306 = v2889
	v3309 = v2892
	goto L330
L330:
	;
	v3342 = v2898 + int32(1)
	if v3342 < v3302 {
		v2885 = v3302
		v2889 = v3306
		v2892 = v3309
		v2898 = v3342
		goto L325
	} else {
		goto L395
	}
L331:
	;
	v2929 = v42 + int32(224)
	goto L333
L332:
	;
	v2928 = *(*int32)(unsafe.Add(mBase, uint32(v2859)+8))
	v2929 = v2928
	goto L333
L333:
	;
	v2933 = *(*int32)(unsafe.Add(mBase, uint32(v2929+v2898<<(uint(int32(2))%32))))
	v2936 = base.B2i32(v2874 == int32(-4))
	if v2936 == int32(0) {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	v2939 = *(*int32)(unsafe.Add(mBase, uint32(v2883)+4))
	v2940 = v2939
	goto L336
L335:
	;
	v2940 = int32(1)
	goto L336
L336:
	;
	v2942 = int32(base.Ui32(v2933) >> (uint(int32(24)) % 32))
	v2944 = int32(base.Ui32(v2933) >> (uint(int32(16)) % 32))
	v2946 = int32(base.Ui32(v2933) >> (uint(int32(8)) % 32))
	v2950 = v2940
	v2954 = v2889
	v2957 = v2892
	v2964 = int32(0)
	goto L337
L337:
	;
	if int32(0) < v2950 {
		goto L340
	} else {
		goto L341
	}
L338:
	;
	if v2856 == int32(-4) {
		v3351 = v3269
		goto L327
	} else {
		goto L394
	}
L339:
	;
	goto L338
L340:
	;
	if v2870 == int32(-4) {
		goto L343
	} else {
		goto L344
	}
L341:
	;
	v3220 = v2950
	v3224 = v2954
	v3227 = v2957
	goto L342
L342:
	;
	v3260 = v2964 + int32(1)
	if v3260 < v3224 {
		v2950 = v3220
		v2954 = v3224
		v2957 = v3227
		v2964 = v3260
		goto L337
	} else {
		goto L393
	}
L343:
	;
	v2994 = v42 + int32(224)
	goto L345
L344:
	;
	v2993 = *(*int32)(unsafe.Add(mBase, uint32(v2873)+8))
	v2994 = v2993
	goto L345
L345:
	;
	v2998 = *(*int32)(unsafe.Add(mBase, uint32(v2994+v2964<<(uint(int32(2))%32))))
	v3000 = int32(base.Ui32(v2998) >> (uint(int32(24)) % 32))
	v3002 = int32(base.Ui32(v2998) >> (uint(int32(16)) % 32))
	v3004 = int32(base.Ui32(v2998) >> (uint(int32(8)) % 32))
	v3007 = int32(0)
	v3013 = v2957
	goto L346
L346:
	;
	if v2874 == int32(-4) {
		goto L348
	} else {
		goto L349
	}
L347:
	;
	if v2870 == int32(-4) {
		v3266 = int32(1)
		v3269 = v3210
		goto L339
	} else {
		goto L392
	}
L348:
	;
	v3048 = v42 + int32(224)
	goto L350
L349:
	;
	v3047 = *(*int32)(unsafe.Add(mBase, uint32(v2883)+8))
	v3048 = v3047
	goto L350
L350:
	;
	v3052 = *(*int32)(unsafe.Add(mBase, uint32(v3048+v3007<<(uint(int32(2))%32))))
	if v2933&int32(255) != 0 {
		goto L352
	} else {
		goto L353
	}
L351:
	;
	v3071 = v3069 + int32(1)
	if v2998&int32(255) != 0 {
		goto L359
	} else {
		goto L360
	}
L352:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v42)+176)) = uint8(v2933)
	if v2946&int32(255) == int32(0) {
		v3069 = v2808
		goto L351
	} else {
		goto L355
	}
L353:
	;
	goto L354
L354:
	;
	v3067 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v42)+176)) = uint8(v3067)
	v3069 = v2808
	goto L351
L355:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v42)+177)) = uint8(v2946)
	if v2944&int32(255) == int32(0) {
		v3069 = v42 + int32(178)
		goto L351
	} else {
		goto L356
	}
L356:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v42)+178)) = uint8(v2944)
	if v2942 == int32(0) {
		v3069 = v42 + int32(179)
		goto L351
	} else {
		goto L357
	}
L357:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v42)+179)) = uint8(v2942)
	v3069 = v42 + int32(180)
	goto L351
L358:
	;
	v3098 = v3096 + int32(1)
	if v3052&int32(255) != 0 {
		goto L370
	} else {
		goto L371
	}
L359:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3069))) = uint8(v2998)
	if v3004&int32(255) == int32(0) {
		v3096 = v3071
		goto L358
	} else {
		goto L362
	}
L360:
	;
	goto L361
L361:
	;
	v3094 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v3069))) = uint8(v3094)
	v3096 = v3071
	goto L358
L362:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3069)+1)) = uint8(v3004)
	if v3002&int32(255) == int32(0) {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	v3096 = v3069 + int32(2)
	goto L358
L364:
	;
	goto L365
L365:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3069)+2)) = uint8(v3002)
	if v3000 == int32(0) {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	v3096 = v3069 + int32(3)
	goto L358
L367:
	;
	goto L368
L368:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3069)+3)) = uint8(v3000)
	v3096 = v3069 + int32(4)
	goto L358
L369:
	;
	v3132 = v42 + int32(176)
	v3133 = v3129 - v3132
	v3139 = int32(255)
	switch v3133 {
	case 0:
		v3191 = v3133
		goto L381
	default:
		goto L382
	case 3:
		goto L383
	}
L370:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3096))) = uint8(v3052)
	v3103 = int32(base.Ui32(v3052) >> (uint(int32(8)) % 32))
	if v3103&int32(255) == int32(0) {
		v3129 = v3098
		goto L369
	} else {
		goto L373
	}
L371:
	;
	goto L372
L372:
	;
	v3127 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v3096))) = uint8(v3127)
	v3129 = v3098
	goto L369
L373:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3096)+1)) = uint8(v3103)
	v3110 = int32(base.Ui32(v3052) >> (uint(int32(16)) % 32))
	if v3110&int32(255) == int32(0) {
		goto L374
	} else {
		goto L375
	}
L374:
	;
	v3129 = v3096 + int32(2)
	goto L369
L375:
	;
	goto L376
L376:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3096)+2)) = uint8(v3110)
	v3119 = int32(base.Ui32(v3052) >> (uint(int32(24)) % 32))
	if v3119 == int32(0) {
		goto L377
	} else {
		goto L378
	}
L377:
	;
	v3129 = v3096 + int32(3)
	goto L369
L378:
	;
	goto L379
L379:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3096)+3)) = uint8(v3119)
	v3129 = v3096 + int32(4)
	goto L369
L380:
	;
	v3210 = v3013 + int32(3)
	if v2874 == int32(-4) {
		goto L388
	} else {
		goto L389
	}
L381:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v3013))) = uint16(v3191)
	v3199 = int32(base.Ui32(v3191) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v3013)+2)) = uint8(v3199)
	goto L380
L382:
	;
	v3150 = v3132
	v3151 = v3133
	v3153 = v3139
	v3154 = v3139
	v3155 = v3139
	v3156 = v3139
	goto L384
L383:
	;
	v3143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3132))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3013))) = uint8(v3143)
	v3145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3132)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3013)+1)) = uint8(v3145)
	v3147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3132)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3013)+2)) = uint8(v3147)
	goto L380
L384:
	;
	v3157 = int32(8)
	v3159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3150))))
	v3165 = *(*int32)(unsafe.Add(mBase, uint32((v3159^v3153)<<(uint(int32(2))%32))+uint32(_c_F_createTrgmNFA[2])))
	v3168 = int32(16)
	v3173 = int32(24)
	v3176 = v3165 ^ (v3154<<(uint(v3157)%32)&int32(_a_F_createTrgmNFA_10) | v3155<<(uint(v3168)%32)&int32(16711680) | v3156<<(uint(v3173)%32))
	v3183 = int32(1)
	v3186 = v3151 - v3183
	if v3186 != 0 {
		v3150 = v3150 + v3183
		v3151 = v3186
		v3153 = int32(base.Ui32(v3176) >> (uint(v3173) % 32))
		v3154 = v3165
		v3155 = int32(base.Ui32(v3176) >> (uint(v3157) % 32))
		v3156 = int32(base.Ui32(v3176) >> (uint(v3168) % 32))
		goto L384
	} else {
		goto L386
	}
L385:
	;
	v3191 = v3176 ^ int32(-1)
	goto L381
L386:
	;
	goto L385
L387:
	;
	goto L347
L388:
	;
	v3217 = int32(1)
	goto L387
L389:
	;
	goto L390
L390:
	;
	v3214 = v3007 + int32(1)
	v3215 = *(*int32)(unsafe.Add(mBase, uint32(v2883)+4))
	if v3214 < v3215 {
		v3007 = v3214
		v3013 = v3210
		goto L346
	} else {
		goto L391
	}
L391:
	;
	v3217 = v3215
	goto L387
L392:
	;
	v3219 = *(*int32)(unsafe.Add(mBase, uint32(v2873)+4))
	v3220 = v3217
	v3224 = v3219
	v3227 = v3210
	goto L342
L393:
	;
	v3266 = v3224
	v3269 = v3227
	goto L339
L394:
	;
	v3301 = *(*int32)(unsafe.Add(mBase, uint32(v2859)+4))
	v3302 = v3301
	v3306 = v3266
	v3309 = v3269
	goto L330
L395:
	;
	v3351 = v3309
	goto L327
L396:
	;
	goto L315
L397:
	;
	v3470 = int32(2)
	v3471 = F_hash_seq_search(m, v3467)
	mBase = m.M
	v3472 = m.ExcPending
	if v3472 != 0 {
		goto L1
	} else {
		goto L398
	}
L398:
	;
	if v3471 != 0 {
		goto L399
	} else {
		goto L400
	}
L399:
	;
	v3477 = v3471
	v3484 = v3470
	goto L402
L400:
	;
	v3545 = v3470
	goto L401
L401:
	;
	v3573 = *(*int32)(unsafe.Add(mBase, uint32(v42)+96))
	v3576 = F_palloc(m, v3573*int32(12))
	mBase = m.M
	v3577 = m.ExcPending
	if v3577 != 0 {
		goto L1
	} else {
		goto L415
	}
L402:
	;
	v3512 = *(*int32)(unsafe.Add(mBase, uint32(v3477)+28))
	if v3512 != 0 {
		v3477 = v3512
		goto L402
	} else {
		goto L404
	}
L403:
	;
	v3545 = v3529
	goto L401
L404:
	;
	v3513 = *(*int32)(unsafe.Add(mBase, uint32(v3477)+24))
	if int32(0) <= v3513 {
		v3529 = v3484
		goto L405
	} else {
		goto L406
	}
L405:
	;
	v3532 = F_hash_seq_search(m, v42+int32(176))
	mBase = m.M
	v3533 = m.ExcPending
	if v3533 != 0 {
		goto L1
	} else {
		goto L413
	}
L406:
	;
	v3516 = *(*int32)(unsafe.Add(mBase, uint32(v3477)+20))
	if v3516&int32(1) != 0 {
		goto L407
	} else {
		goto L408
	}
L407:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3477)+24)) = int32(0)
	v3529 = v3484
	goto L405
L408:
	;
	goto L409
L409:
	;
	if v3516&int32(2) != 0 {
		goto L410
	} else {
		goto L411
	}
L410:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3477)+24)) = int32(1)
	v3529 = v3484
	goto L405
L411:
	;
	goto L412
L412:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3477)+24)) = v3484
	v3529 = v3484 + int32(1)
	goto L405
L413:
	;
	if v3532 != 0 {
		v3477 = v3532
		v3484 = v3529
		goto L402
	} else {
		goto L414
	}
L414:
	;
	goto L403
L415:
	;
	v3579 = v42 + int32(176)
	F_hash_seq_init(m, v3579, v3441)
	mBase = m.M
	v3581 = m.ExcPending
	if v3581 != 0 {
		goto L1
	} else {
		goto L416
	}
L416:
	;
	v3582 = F_hash_seq_search(m, v3579)
	mBase = m.M
	v3583 = m.ExcPending
	if v3583 != 0 {
		goto L1
	} else {
		goto L418
	}
L417:
	;
	v3951 = int32(0)
	v3953 = F_MemoryContextAlloc(m, l3, int32(28))
	mBase = m.M
	v3954 = m.ExcPending
	if v3954 != 0 {
		goto L1
	} else {
		goto L455
	}
L418:
	;
	if v3582 != 0 {
		goto L419
	} else {
		goto L420
	}
L419:
	;
	v3585 = *(*int32)(unsafe.Add(mBase, uint32(v42)+108))
	v3586 = *(*int32)(unsafe.Add(mBase, uint32(v42)+104))
	v3587 = v3582
	v3588 = int32(0)
	goto L422
L420:
	;
	goto L421
L421:
	;
	v3906 = int32(0)
	F_pg_qsort(m, v3576, v3906, int32(12), int32(_a_F_createTrgmNFA_11))
	mBase = m.M
	v3911 = m.ExcPending
	if v3911 != 0 {
		goto L1
	} else {
		goto L454
	}
L422:
	;
	v3630 = v3587
	goto L424
L423:
	;
	F_pg_qsort(m, v3576, v3784, int32(12), int32(_a_F_createTrgmNFA_11))
	mBase = m.M
	v3829 = m.ExcPending
	if v3829 != 0 {
		goto L1
	} else {
		goto L442
	}
L424:
	;
	v3665 = *(*int32)(unsafe.Add(mBase, uint32(v3630)+28))
	if v3665 != 0 {
		v3630 = v3665
		goto L424
	} else {
		goto L426
	}
L425:
	;
	v3666 = *(*int32)(unsafe.Add(mBase, uint32(v3587)+12))
	if v3666 == int32(0) {
		v3784 = v3588
		goto L427
	} else {
		goto L428
	}
L426:
	;
	goto L425
L427:
	;
	v3824 = F_hash_seq_search(m, v42+int32(176))
	mBase = m.M
	v3825 = m.ExcPending
	if v3825 != 0 {
		goto L1
	} else {
		goto L440
	}
L428:
	;
	v3669 = int32(0)
	v3670 = *(*int32)(unsafe.Add(mBase, uint32(v3666)+4))
	if v3670 <= v3669 {
		v3784 = v3588
		goto L427
	} else {
		goto L429
	}
L429:
	;
	v3674 = v3588
	v3680 = v3669
	v3682 = v3670
	goto L430
L430:
	;
	v3712 = *(*int32)(unsafe.Add(mBase, uint32(v3666)+12))
	v3716 = *(*int32)(unsafe.Add(mBase, uint32(v3712+v3680<<(uint(int32(2))%32))))
	v3717 = *(*int32)(unsafe.Add(mBase, uint32(v3716)+12))
	v3722 = v3717
	goto L432
L431:
	;
	v3784 = v3777
	goto L427
L432:
	;
	v3757 = *(*int32)(unsafe.Add(mBase, uint32(v3722)+28))
	if v3757 != 0 {
		v3722 = v3757
		goto L432
	} else {
		goto L434
	}
L433:
	;
	v3758 = *(*int32)(unsafe.Add(mBase, uint32(v3630)+24))
	v3759 = *(*int32)(unsafe.Add(mBase, uint32(v3722)+24))
	if v3758 != v3759 {
		goto L435
	} else {
		goto L436
	}
L434:
	;
	goto L433
L435:
	;
	v3763 = F_bsearch(m, v3716, v3586, v3585, int32(32), int32(_a_F_createTrgmNFA_5))
	mBase = m.M
	v3764 = m.ExcPending
	if v3764 != 0 {
		goto L1
	} else {
		goto L438
	}
L436:
	;
	v3777 = v3674
	v3779 = v3682
	goto L437
L437:
	;
	v3781 = v3680 + int32(1)
	if v3781 < v3779 {
		v3674 = v3777
		v3680 = v3781
		v3682 = v3779
		goto L430
	} else {
		goto L439
	}
L438:
	;
	v3767 = v3576 + v3674*int32(12)
	v3768 = *(*int32)(unsafe.Add(mBase, uint32(v3630)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3767))) = v3768
	v3770 = *(*int32)(unsafe.Add(mBase, uint32(v3722)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3767)+4)) = v3770
	v3772 = *(*int32)(unsafe.Add(mBase, uint32(v3763)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v3767)+8)) = v3772
	v3774 = *(*int32)(unsafe.Add(mBase, uint32(v3666)+4))
	v3777 = v3674 + int32(1)
	v3779 = v3774
	goto L437
L439:
	;
	goto L431
L440:
	;
	if v3824 != 0 {
		v3587 = v3824
		v3588 = v3784
		goto L422
	} else {
		goto L441
	}
L441:
	;
	goto L423
L442:
	;
	if v3784 < int32(2) {
		v3913 = v3784
		goto L417
	} else {
		goto L443
	}
L443:
	;
	v3832 = int32(12)
	__phi3837 = v3576
	__phi3841 = v3576
	__phi3844 = v3576 + v3832
	v3837 = __phi3837
	v3841 = __phi3841
	v3844 = __phi3844
	goto L444
L444:
	;
	v3876 = *(*int32)(unsafe.Add(mBase, uint32(v3841)+12))
	v3877 = *(*int32)(unsafe.Add(mBase, uint32(v3837)))
	if v3876 < v3877 {
		v3895 = v3837
		goto L446
	} else {
		goto L447
	}
L445:
	;
	v3903 = base.I32_div_s(v3895-v3576, int32(12))
	v3913 = v3903 + int32(1)
	goto L417
L446:
	;
	v3899 = v3844 + int32(12)
	if base.Ui32(v3899) < base.Ui32(v3576+v3784*v3832) {
		__phi3837 = v3895
		__phi3841 = v3844
		__phi3844 = v3899
		v3837 = __phi3837
		v3841 = __phi3841
		v3844 = __phi3844
		goto L444
	} else {
		goto L453
	}
L447:
	;
	if v3877 < v3876 {
		goto L448
	} else {
		goto L449
	}
L448:
	;
	v3889 = *(*int32)(unsafe.Add(mBase, uint32(v3844)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v3837)+20)) = v3889
	v3891 = *(*int64)(unsafe.Add(mBase, uint32(v3844)))
	*(*int64)(unsafe.Add(mBase, uint32(v3837)+12)) = v3891
	v3895 = v3837 + int32(12)
	goto L446
L449:
	;
	v3880 = *(*int32)(unsafe.Add(mBase, uint32(v3841)+20))
	v3881 = *(*int32)(unsafe.Add(mBase, uint32(v3837)+8))
	if v3880 < v3881 {
		v3895 = v3837
		goto L446
	} else {
		goto L450
	}
L450:
	;
	if v3881 < v3880 {
		goto L448
	} else {
		goto L451
	}
L451:
	;
	v3884 = *(*int32)(unsafe.Add(mBase, uint32(v3841)+16))
	v3885 = *(*int32)(unsafe.Add(mBase, uint32(v3837)+4))
	if v3884 <= v3885 {
		v3895 = v3837
		goto L446
	} else {
		goto L452
	}
L452:
	;
	goto L448
L453:
	;
	goto L445
L454:
	;
	v3913 = v3906
	goto L417
L455:
	;
	v3955 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3953))) = v3955
	v3957 = *(*int32)(unsafe.Add(mBase, uint32(v42)+108))
	if v3957 <= v3955 {
		goto L457
	} else {
		goto L458
	}
L456:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3953)+8)) = v3545
	v4360 = F_MemoryContextAlloc(m, l3, v3545<<(uint(int32(3))%32))
	mBase = m.M
	v4361 = m.ExcPending
	if v4361 != 0 {
		goto L1
	} else {
		goto L502
	}
L457:
	;
	v3961 = F_MemoryContextAlloc(m, l3, int32(0))
	mBase = m.M
	v3962 = m.ExcPending
	if v3962 != 0 {
		goto L1
	} else {
		goto L460
	}
L458:
	;
	goto L459
L459:
	;
	v3965 = v3957 & int32(3)
	v3966 = *(*int32)(unsafe.Add(mBase, uint32(v42)+104))
	v3967 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v3957) {
		goto L462
	} else {
		goto L463
	}
L460:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3953)+4)) = v3961
	goto L456
L461:
	;
	v4185 = F_MemoryContextAlloc(m, l3, v4144<<(uint(int32(2))%32))
	mBase = m.M
	v4186 = m.ExcPending
	if v4186 != 0 {
		goto L1
	} else {
		goto L487
	}
L462:
	;
	v3973 = v3951
	v3977 = v3967
	v3982 = int32(0)
	goto L465
L463:
	;
	v4050 = v3951
	v4054 = v3967
	goto L464
L464:
	;
	v4090 = v4050
	v4094 = v4054
	v4096 = int32(0)
	goto L481
L465:
	;
	v4014 = v3966 + v3977<<(uint(int32(5))%32)
	v4015 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4014)+24)))
	if v4015 == int32(1) {
		goto L467
	} else {
		goto L468
	}
L466:
	;
	if v3965 == int32(0) {
		v4144 = v4042
		goto L461
	} else {
		goto L480
	}
L467:
	;
	v4019 = v3973 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3953))) = v4019
	v4021 = v4019
	goto L469
L468:
	;
	v4021 = v3973
	goto L469
L469:
	;
	v4022 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4014)+56)))
	if v4022 == int32(1) {
		goto L470
	} else {
		goto L471
	}
L470:
	;
	v4026 = v4021 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3953))) = v4026
	v4028 = v4026
	goto L472
L471:
	;
	v4028 = v4021
	goto L472
L472:
	;
	v4029 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4014)+88)))
	if v4029 == int32(1) {
		goto L473
	} else {
		goto L474
	}
L473:
	;
	v4033 = v4028 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3953))) = v4033
	v4035 = v4033
	goto L475
L474:
	;
	v4035 = v4028
	goto L475
L475:
	;
	v4036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4014)+120)))
	if v4036 == int32(1) {
		goto L476
	} else {
		goto L477
	}
L476:
	;
	v4040 = v4035 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3953))) = v4040
	v4042 = v4040
	goto L478
L477:
	;
	v4042 = v4035
	goto L478
L478:
	;
	v4043 = int32(4)
	v4044 = v3977 + v4043
	v4046 = v3982 + v4043
	if v4046 != v3957&int32(2147483644) {
		v3973 = v4042
		v3977 = v4044
		v3982 = v4046
		goto L465
	} else {
		goto L479
	}
L479:
	;
	goto L466
L480:
	;
	v4050 = v4042
	v4054 = v4044
	goto L464
L481:
	;
	v4132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3966+v4094<<(uint(int32(5))%32))+24)))
	if v4132 == int32(1) {
		goto L483
	} else {
		goto L484
	}
L482:
	;
	v4144 = v4138
	goto L461
L483:
	;
	v4136 = v4090 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3953))) = v4136
	v4138 = v4136
	goto L485
L484:
	;
	v4138 = v4090
	goto L485
L485:
	;
	v4139 = int32(1)
	v4142 = v4096 + v4139
	if v4142 != v3965 {
		v4090 = v4138
		v4094 = v4094 + v4139
		v4096 = v4142
		goto L481
	} else {
		goto L486
	}
L486:
	;
	goto L482
L487:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3953)+4)) = v4185
	v4188 = int32(0)
	if v3957 != int32(1) {
		goto L488
	} else {
		goto L489
	}
L488:
	;
	v4197 = v4188
	v4201 = v4188
	v4206 = int32(0)
	goto L491
L489:
	;
	v4268 = v4188
	v4272 = v4188
	goto L490
L490:
	;
	v4309 = v3966 + v4272<<(uint(int32(5))%32)
	v4310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4309)+24)))
	if v4310 != int32(1) {
		goto L456
	} else {
		goto L501
	}
L491:
	;
	v4238 = v3966 + v4201<<(uint(int32(5))%32)
	v4239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4238)+24)))
	if v4239 == int32(1) {
		goto L493
	} else {
		goto L494
	}
L492:
	;
	if v3957&int32(1) == int32(0) {
		goto L456
	} else {
		goto L500
	}
L493:
	;
	v4245 = *(*int32)(unsafe.Add(mBase, uint32(v4238)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v4185+v4197<<(uint(int32(2))%32)))) = v4245
	v4249 = v4197 + int32(1)
	goto L495
L494:
	;
	v4249 = v4197
	goto L495
L495:
	;
	v4250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4238)+56)))
	if v4250 == int32(1) {
		goto L496
	} else {
		goto L497
	}
L496:
	;
	v4256 = *(*int32)(unsafe.Add(mBase, uint32(v4238)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v4185+v4249<<(uint(int32(2))%32)))) = v4256
	v4260 = v4249 + int32(1)
	goto L498
L497:
	;
	v4260 = v4249
	goto L498
L498:
	;
	v4261 = int32(2)
	v4262 = v4201 + v4261
	v4264 = v4206 + v4261
	if v4264 != v3957&int32(2147483646) {
		v4197 = v4260
		v4201 = v4262
		v4206 = v4264
		goto L491
	} else {
		goto L499
	}
L499:
	;
	goto L492
L500:
	;
	v4268 = v4260
	v4272 = v4262
	goto L490
L501:
	;
	v4316 = *(*int32)(unsafe.Add(mBase, uint32(v4309)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v4185+v4268<<(uint(int32(2))%32)))) = v4316
	goto L456
L502:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3953)+12)) = v4360
	v4365 = F_MemoryContextAlloc(m, l3, v3913<<(uint(int32(3))%32))
	mBase = m.M
	v4366 = m.ExcPending
	if v4366 != 0 {
		goto L1
	} else {
		goto L503
	}
L503:
	;
	if int32(0) < v3545 {
		goto L504
	} else {
		goto L505
	}
L504:
	;
	v4369 = *(*int32)(unsafe.Add(mBase, uint32(v3953)+12))
	v4370 = int32(0)
	v4376 = v4370
	v4381 = v4370
	goto L507
L505:
	;
	goto L506
L506:
	;
	v4560 = *(*int32)(unsafe.Add(mBase, uint32(v3953)))
	v4561 = F_MemoryContextAlloc(m, l3, v4560)
	mBase = m.M
	v4562 = m.ExcPending
	if v4562 != 0 {
		goto L1
	} else {
		goto L518
	}
L507:
	;
	v4411 = int32(3)
	v4413 = v4369 + v4381<<(uint(v4411)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v4413)+4)) = v4365 + v4376<<(uint(v4411)%32)
	if v3913 <= v4376 {
		goto L510
	} else {
		goto L511
	}
L508:
	;
	goto L506
L509:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4413))) = v4484
	v4519 = v4381 + int32(1)
	if v4519 != v3545 {
		v4376 = v4482
		v4381 = v4519
		goto L507
	} else {
		goto L517
	}
L510:
	;
	v4482 = v4376
	v4484 = int32(0)
	goto L509
L511:
	;
	goto L512
L512:
	;
	v4420 = v3913 - v4376
	v4426 = v4376
	v4428 = int32(0)
	goto L513
L513:
	;
	v4463 = v3576 + v4426*int32(12)
	v4464 = *(*int32)(unsafe.Add(mBase, uint32(v4463)))
	if v4464 != v4381 {
		v4482 = v4426
		v4484 = v4428
		goto L509
	} else {
		goto L515
	}
L514:
	;
	v4482 = v3913
	v4484 = v4420
	goto L509
L515:
	;
	v4468 = v4365 + v4426<<(uint(int32(3))%32)
	v4469 = *(*int32)(unsafe.Add(mBase, uint32(v4463)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v4468))) = v4469
	v4471 = *(*int32)(unsafe.Add(mBase, uint32(v4463)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v4468)+4)) = v4471
	v4473 = int32(1)
	v4476 = v4428 + v4473
	if v4476 != v4420 {
		v4426 = v4426 + v4473
		v4428 = v4476
		goto L513
	} else {
		goto L516
	}
L516:
	;
	goto L514
L517:
	;
	goto L508
L518:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3953)+16)) = v4561
	v4564 = *(*int32)(unsafe.Add(mBase, uint32(v3953)+8))
	v4565 = F_MemoryContextAlloc(m, l3, v4564)
	mBase = m.M
	v4566 = m.ExcPending
	if v4566 != 0 {
		goto L1
	} else {
		goto L519
	}
L519:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3953)+20)) = v4565
	v4568 = *(*int32)(unsafe.Add(mBase, uint32(v3953)+8))
	v4571 = F_MemoryContextAlloc(m, l3, v4568<<(uint(int32(2))%32))
	mBase = m.M
	v4572 = m.ExcPending
	if v4572 != 0 {
		goto L1
	} else {
		goto L520
	}
L520:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3953)+24)) = v4571
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v3953
	v4594 = v2788
	goto L21
L521:
	;
	m.G0 = v42 + int32(240)
	return v4594
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
	var v126 int32
	_ = v126
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
	var v244 float64
	_ = v244
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
	var v268 float64
	_ = v268
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
	*(*int64)(unsafe.Add(mBase, uint32(v42)+16)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v42)+24)) = v111
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
	v126 = v5
	goto L22
L22:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v135+v126<<(uint(int32(2))%32))))
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
	v145 = v126 + int32(1)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v145 < v146 {
		v126 = v145
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
	*(*float64)(unsafe.Add(mBase, uint32(v19)+56)) = base.F64_add(v286, base.F64_add(base.F64_mul(v277, v288), base.F64_add(base.F64_mul(base.F64_sub(base.F64_add(v268, v279), v178), v177), base.F64_add(base.F64_mul(v187, v177), v284))))
	m.G0 = v42 + int32(32)
	return v19
L28:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v189 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v42)+16)) = v189
	*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v42)+24)) = v189
	if v188 == int32(0) {
		v244 = v10
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
	v268 = v256
	v275 = v257
	goto L27
L31:
	;
	v252 = *(*float64)(unsafe.Add(mBase, uint32(l1)+200))
	v254 = *(*float64)(unsafe.Add(mBase, uint32(l1)+192))
	v268 = base.F64_add(v244, v252)
	v275 = base.F64_add(v251, v254)
	goto L27
L32:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v198 <= int32(0) {
		v244 = v10
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
	v244 = v232
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
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
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
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
			if v33 == int32(18) {
				v36 = int32(16)
			} else {
				v36 = int32(0)
			}
			if base.Ui32((v33-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v43 = int32(4)
			} else {
				v43 = v36
			}
			v56 = v43
		} else {
			v44 = int32(1)
			if v27&v44 != 0 {
				v56 = int32(base.Ui32(v27)>>(uint(v44)%32)) - v44
			} else {
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v56 = int32(base.Ui32(v50)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v57 = *(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_c_F_cryptohash_internal[1])))
		v58 = F_pg_cryptohash_create(m, l0)
		mBase = m.M
		v59 = m.ExcPending
		if v59 != 0 {
			return int32(0)
		} else {
			v60 = F_pg_cryptohash_init(m, v58)
			mBase = m.M
			if int32(0) <= v60 {
				v63 = int32(1)
				if v27&v63 != 0 {
					v67 = v63
				} else {
					v67 = int32(4)
				}
				v69 = F_pg_cryptohash_update(m, v58, l1+v67, v56)
				mBase = m.M
				if v69 < int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
						return int32(0)
					} else {
						if v58 == int32(0) {
							v133 = int32(_a_F_cryptohash_internal_0)
						} else {
							v125 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
							if v125 == int32(1) {
								v128 = int32(_a_F_cryptohash_internal_1)
							} else {
								v128 = int32(_a_F_cryptohash_internal_2)
							}
							if v125 == int32(2) {
								v131 = int32(_a_F_cryptohash_internal_0)
							} else {
								v131 = v128
							}
							v133 = v131
						}
						*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v133
						*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v57
						F_errmsg_internal(m, int32(_a_F_cryptohash_internal_3), v12+int32(16))
						mBase = m.M
						v140 = m.ExcPending
						if v140 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_cryptohash_internal_4), int32(123), int32(_a_F_cryptohash_internal_5))
							mBase = m.M
							v145 = m.ExcPending
							if v145 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v74 = F_pg_cryptohash_final(m, v58, v21+int32(4), v18)
					mBase = m.M
					if v74 < int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v149 = m.ExcPending
						if v149 != 0 {
							return int32(0)
						} else {
							if v58 == int32(0) {
								v164 = int32(_a_F_cryptohash_internal_0)
							} else {
								v156 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
								if v156 == int32(1) {
									v159 = int32(_a_F_cryptohash_internal_1)
								} else {
									v159 = int32(_a_F_cryptohash_internal_2)
								}
								if v156 == int32(2) {
									v162 = int32(_a_F_cryptohash_internal_0)
								} else {
									v162 = v159
								}
								v164 = v162
							}
							*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v164
							*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v57
							F_errmsg_internal(m, int32(_a_F_cryptohash_internal_6), v12+int32(32))
							mBase = m.M
							v171 = m.ExcPending
							if v171 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_cryptohash_internal_4), int32(127), int32(_a_F_cryptohash_internal_5))
								mBase = m.M
								v176 = m.ExcPending
								if v176 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						F_pg_cryptohash_free(m, v58)
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
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
				v89 = m.ExcPending
				if v89 != 0 {
					return int32(0)
				} else {
					if v58 == int32(0) {
						v104 = int32(_a_F_cryptohash_internal_0)
					} else {
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
						if v96 == int32(1) {
							v99 = int32(_a_F_cryptohash_internal_1)
						} else {
							v99 = int32(_a_F_cryptohash_internal_2)
						}
						if v96 == int32(2) {
							v102 = int32(_a_F_cryptohash_internal_0)
						} else {
							v102 = v99
						}
						v104 = v102
					}
					*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v104
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = v57
					F_errmsg_internal(m, int32(_a_F_cryptohash_internal_7), v12)
					mBase = m.M
					v109 = m.ExcPending
					if v109 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_cryptohash_internal_4), int32(120), int32(_a_F_cryptohash_internal_5))
						mBase = m.M
						v114 = m.ExcPending
						if v114 != 0 {
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
						F_relation_close(m, v14, int32(1))
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
