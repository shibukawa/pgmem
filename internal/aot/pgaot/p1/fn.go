package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_coerce_fn_result_column(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	if l3 == int32(0) {
		v32 = int32(0)
		v35 = F_makeVarFromTargetEntry(m, int32(1), l0)
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
			v41 = F_coerce_to_target_type(m, v32, v35, v37, l1, l2, int32(1), int32(2), int32(-1))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				if v41 == int32(0) {
					v75 = v32
					return v75
				} else {
					F_assign_expr_collations(m, int32(0), v41)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						if v41 == v35 {
							v52 = v41
						} else {
							v49 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v49)
							v52 = v41
						}
						v57 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
						if v57 != 0 {
							v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)))
							v62 = v58 + int32(1)
						} else {
							v62 = int32(1)
						}
						v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v66 = F_makeTargetEntry(m, v52, base.I32_extend16_s(v62), v64, int32(0))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							v68 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
							v69 = F_lappend(m, v68, v66)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v69
								v75 = int32(1)
								return v75
							}
						}
					}
				}
			}
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v10 != 0 {
			v32 = int32(0)
			v35 = F_makeVarFromTargetEntry(m, int32(1), l0)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
				v41 = F_coerce_to_target_type(m, v32, v35, v37, l1, l2, int32(1), int32(2), int32(-1))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					if v41 == int32(0) {
						v75 = v32
						return v75
					} else {
						F_assign_expr_collations(m, int32(0), v41)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							if v41 == v35 {
								v52 = v41
							} else {
								v49 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v49)
								v52 = v41
							}
							v57 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
							if v57 != 0 {
								v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)))
								v62 = v58 + int32(1)
							} else {
								v62 = int32(1)
							}
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v66 = F_makeTargetEntry(m, v52, base.I32_extend16_s(v62), v64, int32(0))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								v68 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
								v69 = F_lappend(m, v68, v66)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l4))) = v69
									v75 = int32(1)
									return v75
								}
							}
						}
					}
				}
			}
		} else {
			v11 = int32(0)
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v14 = F_exprType(m, v13)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v21 = F_coerce_to_target_type(m, v11, v13, v14, l1, l2, int32(1), int32(2), int32(-1))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					if v21 == int32(0) {
						v75 = v11
						return v75
					} else {
						F_assign_expr_collations(m, int32(0), v21)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v21
							v30 = F_makeVarFromTargetEntry(m, int32(1), l0)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								v52 = v30
								v57 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
								if v57 != 0 {
									v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)))
									v62 = v58 + int32(1)
								} else {
									v62 = int32(1)
								}
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v66 = F_makeTargetEntry(m, v52, base.I32_extend16_s(v62), v64, int32(0))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									v68 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
									v69 = F_lappend(m, v68, v66)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l4))) = v69
										v75 = int32(1)
										return v75
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
func F_get_fn_expr_variadic(m *base.Module, l0 int32) int32 {
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
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v5 == int32(0) {
			v13 = v2
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			if v8 != int32(15) {
				v13 = v2
			} else {
				v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+13)))
				v13 = v11
			}
		}
	}
	return v13 & int32(1)
}
func Fn14210(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
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
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v11 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v148
	v154 = int32(*(*int8)(unsafe.Add(mBase, uint32(v148)+11)))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v148)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v155
	return v154
L2:
	;
	goto L7
L3:
	;
	goto L4
L4:
	;
	v60 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	v67 = l4
	v69 = l3
	goto L19
L5:
	;
	if v49-v50 == int32(0) {
		v148 = v11
		goto L1
	} else {
		goto L18
	}
L7:
	;
	goto L8
L8:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v18 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v19 = l0
	v20 = v11
	v21 = int32(10)
	v22 = v18
	goto L13
L10:
	;
	v45 = v11
	v49 = int32(0)
	goto L11
L11:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	goto L5
L12:
	;
	v45 = v40
	v49 = v42
	goto L11
L13:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if base.B2i32(v22 != v24)|base.B2i32(v24 == int32(0)) != 0 {
		v40 = v20
		v42 = v22
		goto L12
	} else {
		goto L15
	}
L14:
	;
	v40 = v34
	v42 = int32(0)
	goto L12
L15:
	;
	v30 = v21 - int32(1)
	if v30 == int32(0) {
		v40 = v20
		v42 = v22
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v33 = int32(1)
	v34 = v20 + v33
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v35 != 0 {
		v19 = v19 + v33
		v20 = v34
		v21 = v30
		v22 = v35
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	goto L4
L19:
	;
	v76 = v67 + (v69-v67)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v77 = int32(*(*int8)(unsafe.Add(mBase, uint32(v76))))
	v78 = v60 - v77
	if v78 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	return int32(31)
L21:
	;
	goto L26
L22:
	;
	v129 = v78
	goto L23
L23:
	;
	v133 = base.B2i32(v129 < int32(0))
	if v129 < int32(0) {
		goto L38
	} else {
		goto L39
	}
L24:
	;
	if v120 == int32(0) {
		v148 = v76
		goto L1
	} else {
		goto L37
	}
L26:
	;
	goto L27
L27:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v87 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v88 = l0
	v89 = v76
	v90 = int32(10)
	v91 = v87
	goto L32
L29:
	;
	v114 = v76
	v118 = int32(0)
	goto L30
L30:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	v120 = v118 - v119
	goto L24
L31:
	;
	v114 = v109
	v118 = v111
	goto L30
L32:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	if base.B2i32(v91 != v93)|base.B2i32(v93 == int32(0)) != 0 {
		v109 = v89
		v111 = v91
		goto L31
	} else {
		goto L34
	}
L33:
	;
	v109 = v103
	v111 = int32(0)
	goto L31
L34:
	;
	v99 = v90 - int32(1)
	if v99 == int32(0) {
		v109 = v89
		v111 = v91
		goto L31
	} else {
		goto L35
	}
L35:
	;
	v102 = int32(1)
	v103 = v89 + v102
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
	if v104 != 0 {
		v88 = v88 + v102
		v89 = v103
		v90 = v99
		v91 = v104
		goto L32
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	v129 = v120
	goto L23
L38:
	;
	v134 = v76 - int32(16)
	goto L40
L39:
	;
	v134 = v69
	goto L40
L40:
	;
	if v129 < int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v137 = v67
	goto L43
L42:
	;
	v137 = v76 + int32(16)
	goto L43
L43:
	;
	if base.Ui32(v137) <= base.Ui32(v134) {
		v67 = v137
		v69 = v134
		goto L19
	} else {
		goto L44
	}
L44:
	;
	goto L20
}
func Fn14223(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
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
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
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
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v124 int64
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v18 = F_SearchSysCache1(m, l6, base.I64_extend_i32_u(l0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v148
L2:
	;
	return int32(0)
L3:
	;
	if v18 == int32(0) {
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
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+22)))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	v24 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v24)
	v148 = int32(0)
	goto L1
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l0
	F_errmsg_internal(m, l5, v15)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_Fn14223_0), l4, l3)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
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
	v41 = v37 + v38
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+68))
	if v42 != int32(11) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F_ReleaseCatCache(m, v18)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L2
	} else {
		goto L46
	}
L15:
	;
	v45 = int32(0)
	v47 = *(*int32)(unsafe.Add(mBase, _c_Fn14223[0]))
	if v47 == v45 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	v90 = int32(0)
	v92 = *(*int32)(unsafe.Add(mBase, _c_Fn14223[0]))
	if v92 == v90 {
		v134 = v90
		goto L14
	} else {
		goto L32
	}
L18:
	;
	if v86 == int32(0) {
		v134 = v45
		goto L14
	} else {
		goto L31
	}
L19:
	;
	v86 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v54 <= int32(0) {
		v80 = v45
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v86 = v80
	goto L18
L23:
	;
	v57 = int32(0)
	if v57 < v54 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v60 = v54
	goto L26
L25:
	;
	v60 = v57
	goto L26
L26:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	v63 = int32(0)
	goto L27
L27:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v61+v63<<(uint(int32(2))%32))))
	v72 = base.B2i32(v71 == v42)
	if v71 == v42 {
		v80 = v72
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v80 = v72
	goto L22
L29:
	;
	v74 = v63 + int32(1)
	if v74 != v60 {
		v63 = v74
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
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	if v95 <= int32(0) {
		v134 = v90
		goto L14
	} else {
		goto L33
	}
L33:
	;
	v101 = v90
	goto L34
L34:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113+v101<<(uint(int32(2))%32))))
	v119 = *(*int32)(unsafe.Add(mBase, _c_Fn14223[1]))
	if v117 != v119 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v134 = int32(0)
	goto L14
L36:
	;
	goto L35
L37:
	;
	if v117 == v42 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	v129 = v101 + int32(1)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	if v129 < v130 {
		v101 = v129
		goto L34
	} else {
		goto L45
	}
L40:
	;
	v134 = int32(1)
	goto L14
L41:
	;
	goto L42
L42:
	;
	v124 = int64(0)
	v126 = F_SearchSysCacheExists(m, l2, base.I64_extend_i32_u(v41+int32(4)), base.I64_extend_i32_u(v117), v124, v124)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	if v126 != 0 {
		goto L36
	} else {
		goto L44
	}
L44:
	;
	goto L39
L45:
	;
	goto L36
L46:
	;
	v148 = v134
	goto L1
}
func Fn14225(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int64
	_ = v61
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v20 = F_pg_detoast_datum(m, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			v24 = F_array_iterator(m, v15, l1, v20, v12+int32(12))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int64(0)
			} else {
				if v24 == int32(0) {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v28 != v15 {
						F_pfree(m, v15)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int64(0)
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							if v32 != v20 {
								F_pfree(m, v20)
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return int64(0)
								} else {
									v36 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v36)
									v61 = int64(0)
									m.G0 = v12 + int32(16)
									return v61
								}
							} else {
								v36 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v36)
								v61 = int64(0)
								m.G0 = v12 + int32(16)
								return v61
							}
						}
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v32 != v20 {
							F_pfree(m, v20)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int64(0)
							} else {
								v36 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v36)
								v61 = int64(0)
								m.G0 = v12 + int32(16)
								return v61
							}
						} else {
							v36 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v36)
							v61 = int64(0)
							m.G0 = v12 + int32(16)
							return v61
						}
					}
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
					v43 = F_palloc0(m, int32(base.Ui32(v40)>>(uint(int32(2))%32)))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int64(0)
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
						v47 = int32(base.Ui32(v45) >> (uint(int32(2)) % 32))
						if v47 != 0 {
							base.MemoryCopy(m, v43, v39, v47)
						} else {
						}
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						if v49 != v15 {
							F_pfree(m, v15)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int64(0)
							} else {
								v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								if v53 != v20 {
									F_pfree(m, v20)
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int64(0)
									} else {
										v61 = base.I64_extend_i32_u(v43)
										m.G0 = v12 + int32(16)
										return v61
									}
								} else {
									v61 = base.I64_extend_i32_u(v43)
									m.G0 = v12 + int32(16)
									return v61
								}
							}
						} else {
							v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							if v53 != v20 {
								F_pfree(m, v20)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int64(0)
								} else {
									v61 = base.I64_extend_i32_u(v43)
									m.G0 = v12 + int32(16)
									return v61
								}
							} else {
								v61 = base.I64_extend_i32_u(v43)
								m.G0 = v12 + int32(16)
								return v61
							}
						}
					}
				}
			}
		}
	}
}
func Fn14229(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_DatumGetAnyArrayP(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v12 = F_DatumGetAnyArrayP(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v18 = F_array_contain_compare(m, v7, v12, v14, l1, v15+int32(16))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int64(0)
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				if v20 == int32(-1) {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if v27 == int32(-1) {
						return base.I64_extend_i32_u(v18)
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v12 == v30 {
							return base.I64_extend_i32_u(v18)
						} else {
							F_pfree(m, v12)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(v18)
							}
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v7 == v23 {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						if v27 == int32(-1) {
							return base.I64_extend_i32_u(v18)
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							if v12 == v30 {
								return base.I64_extend_i32_u(v18)
							} else {
								F_pfree(m, v12)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return int64(0)
								} else {
									return base.I64_extend_i32_u(v18)
								}
							}
						}
					} else {
						F_pfree(m, v7)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int64(0)
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
							if v27 == int32(-1) {
								return base.I64_extend_i32_u(v18)
							} else {
								v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								if v12 == v30 {
									return base.I64_extend_i32_u(v18)
								} else {
									F_pfree(m, v12)
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return int64(0)
									} else {
										return base.I64_extend_i32_u(v18)
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
func Fn14234(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v4 = F_palloc0(m, l2)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v4)+2)) = uint8(v8)
		*(*uint16)(unsafe.Add(mBase, uint32(v4))) = uint16(v8)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = (v4 + int32(19)) & int32(-8)
		v18 = F_lookup_type_cache(m, l1, int32(0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = v18
			return base.I64_extend_i32_u(v4)
		}
	}
}
func Fn14238(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	v5 = int32(_a_Fn14238_0)
	v6 = *(*int32)(unsafe.Add(mBase, _c_Fn14238[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	*(*int32)(unsafe.Add(mBase, _c_Fn14238[0])) = v10
	F_varstr_sortsupport(m, v7, l1, v8)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_Fn14238[0])) = v6
		return int64(0)
	}
}
func Fn14243(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = v11 - v8
	if base.B2i32(int64(0) < v8) != base.B2i32(v12 < v11) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, l4, int32(0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, l3, l2, l1)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		return v12
	}
}
func Fn14249(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
	v10 = F_errdetail(m, l1, v6)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return v10
	}
}
func Fn14254(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
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
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v17 - int32(1) {
	case 0:
		v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v55 == int32(0) {
			v63 = int32(23)
			v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v65 = F_errsave_start(m, v64)
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return int32(0)
			} else {
				if v65 == int32(0) {
					v84 = v63
					m.G0 = v15 - int32(-64)
					return v84
				} else {
					F_errcode(m, int32(33685634))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v72
						F_errmsg(m, l6, v13+int32(-16))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return int32(0)
						} else {
							v80 = F_errdetail(m, int32(_a_Fn14254_0), int32(0))
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								F_errsave_finish(m, v64, l3, l5, l1)
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return int32(0)
								} else {
									v84 = v63
									m.G0 = v15 - int32(-64)
									return v84
								}
							}
						}
					}
				}
			}
		} else {
			v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
			if v58 <= int32(0) {
				v63 = int32(23)
				v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v65 = F_errsave_start(m, v64)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return int32(0)
				} else {
					if v65 == int32(0) {
						v84 = v63
						m.G0 = v15 - int32(-64)
						return v84
					} else {
						F_errcode(m, int32(33685634))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v72
							F_errmsg(m, l6, v13+int32(-16))
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return int32(0)
							} else {
								v80 = F_errdetail(m, int32(_a_Fn14254_0), int32(0))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									F_errsave_finish(m, v64, l3, l5, l1)
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return int32(0)
									} else {
										v84 = v63
										m.G0 = v15 - int32(-64)
										return v84
									}
								}
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l7
				v84 = int32(0)
				m.G0 = v15 - int32(-64)
				return v84
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v93 = m.ExcPending
		if v93 != 0 {
			return int32(0)
		} else {
			v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v94
			*(*int32)(unsafe.Add(mBase, uint32(v15))) = l4
			F_errmsg_internal(m, int32(_a_Fn14254_1), v15)
			mBase = m.M
			v99 = m.ExcPending
			if v99 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, l3, l2, l1)
				mBase = m.M
				v101 = m.ExcPending
				if v101 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 3:
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v20 == int32(0) {
			v29 = int32(23)
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v31 = F_errsave_start(m, v30)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				if v31 == int32(0) {
					v84 = v29
					m.G0 = v15 - int32(-64)
					return v84
				} else {
					F_errcode(m, int32(33685634))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v40
						F_errmsg(m, l6, v13+int32(-32))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(_a_Fn14254_2)
							v51 = F_errdetail(m, int32(_a_Fn14254_3), v13+int32(-48))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								F_errsave_finish(m, v30, l3, l8, l1)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									v84 = v29
									m.G0 = v15 - int32(-64)
									return v84
								}
							}
						}
					}
				}
			}
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
			if v23 <= int32(0) {
				v29 = int32(23)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v31 = F_errsave_start(m, v30)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					if v31 == int32(0) {
						v84 = v29
						m.G0 = v15 - int32(-64)
						return v84
					} else {
						F_errcode(m, int32(33685634))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v40
							F_errmsg(m, l6, v13+int32(-32))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(_a_Fn14254_2)
								v51 = F_errdetail(m, int32(_a_Fn14254_3), v13+int32(-48))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									F_errsave_finish(m, v30, l3, l8, l1)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										v84 = v29
										m.G0 = v15 - int32(-64)
										return v84
									}
								}
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(2)
				v84 = int32(0)
				m.G0 = v15 - int32(-64)
				return v84
			}
		}
	}
}
func Fn14261(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = int32(_a_Fn14261_0)
	v16 = *(*int32)(unsafe.Add(mBase, _c_Fn14261[0]))
	*(*int32)(unsafe.Add(mBase, _c_Fn14261[0])) = v16 + int32(1)
	v21 = *(*int32)(unsafe.Add(mBase, _c_Fn14261[1]))
	if int32(0) <= v21 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = int32(_a_Fn14261_1)
	v25 = *(*int32)(unsafe.Add(mBase, _c_Fn14261[2]))
	v28 = v21 * int32(100)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_Fn14261[3])))
	*(*int32)(unsafe.Add(mBase, _c_Fn14261[2])) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_Fn14261[4]))) = l0
	v35 = v12 + int32(16)
	F_initStringInfo(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_Fn14261[1])) = int32(-1)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L21
	}
L4:
	;
	return
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_Fn14261[5])))
	*(*int32)(unsafe.Add(mBase, _c_Fn14261[6])) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l1
	v42 = F_appendStringInfoVA(m, v35, l0, l1)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v42 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v50 = v42
	goto L10
L8:
	;
	goto L9
L9:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_Fn14261[7])))
	if v72 != 0 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	v54 = v12 + int32(16)
	F_enlargeStringInfo(m, v54, v50)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_Fn14261[5])))
	*(*int32)(unsafe.Add(mBase, _c_Fn14261[6])) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l1
	v61 = F_appendStringInfoVA(m, v54, l0, l1)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	if v61 != 0 {
		v50 = v61
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	F_pfree(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v76 = F_pstrdup(m, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_Fn14261[7]))) = v76
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	F_pfree(m, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_Fn14261[2])) = v25
	v84 = int32(_a_Fn14261_0)
	v86 = *(*int32)(unsafe.Add(mBase, _c_Fn14261[0]))
	*(*int32)(unsafe.Add(mBase, _c_Fn14261[0])) = v86 - int32(1)
	m.G0 = v12 + int32(32)
	return
L21:
	;
	F_errmsg_internal(m, int32(_a_Fn14261_2), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_Fn14261_3), l3, l2)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
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
func Fn14267(m *base.Module, l0 int32, l1 int32) int64 {
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
	var v11 int32
	_ = v11
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = F_gbt_var_penalty(m, v3, v4, v5, v6, l1, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v8)
	}
}
func Fn14270(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
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
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = base.I32_wrap_i64(l0)
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = base.I32_wrap_i64(l1)
		v19 = F_pg_detoast_datum(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v22 = v11 + int32(8)
			v24 = int32(4)
			v25 = v14 + v24
			*(*int32)(unsafe.Add(mBase, uint32(v22))) = v25
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
			v28 = int32(2)
			v29 = int32(base.Ui32(v27) >> (uint(v28) % 32))
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			if base.Ui32(v29+v24) < base.Ui32(int32(base.Ui32(v37)>>(uint(v28)%32))) {
				v41 = v25 + (v29+int32(3))&int32(2147483644)
			} else {
				v41 = v25
			}
			*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v41
			v44 = int32(4)
			v45 = v19 + v44
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = v45
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
			v48 = int32(2)
			v49 = int32(base.Ui32(v47) >> (uint(v48) % 32))
			v57 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			if base.Ui32(v49+v44) < base.Ui32(int32(base.Ui32(v57)>>(uint(v48)%32))) {
				v61 = v45 + (v49+int32(3))&int32(2147483644)
			} else {
				v61 = v45
			}
			*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v61
			v64 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v11)+8)))
			v65 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v11))))
			v66 = F_DirectFunctionCall2Coll(m, l3, int32(0), v64, v65)
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int32(0)
			} else {
				if v14 != v13 {
					F_pfree(m, v14)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						if v19 != v18 {
							F_pfree(m, v19)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								m.G0 = v11 + int32(16)
								return base.I32_wrap_i64(v66)
							}
						} else {
							m.G0 = v11 + int32(16)
							return base.I32_wrap_i64(v66)
						}
					}
				} else {
					if v19 != v18 {
						F_pfree(m, v19)
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							m.G0 = v11 + int32(16)
							return base.I32_wrap_i64(v66)
						}
					} else {
						m.G0 = v11 + int32(16)
						return base.I32_wrap_i64(v66)
					}
				}
			}
		}
	}
}
func Fn14287(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v5 = F_SearchSysCache1(m, l1, base.I64_extend_i32_u(l0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v13+v14)+92))
			F_ReleaseCatCache(m, v5)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return v16
			}
		}
	}
}
func Fn14289(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v5 = F_SearchSysCache1(m, l1, base.I64_extend_i32_u(l0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v13+v14)+68))
			F_ReleaseCatCache(m, v5)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return v16
			}
		}
	}
}
func Fn14292(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
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
	v10 = F_SearchSysCache4(m, l4, base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1), base.I64_extend_i32_u(l2), base.I64_extend_i32_s(l3))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+22)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v18+v19)+20))
			F_ReleaseCatCache(m, v10)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				return v21
			}
		}
	}
}
func Fn14294(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
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
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int64
	_ = v45
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int64
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
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
	return v120
L4:
	;
	if l1|v95 != 0 {
		v120 = v95
		goto L3
	} else {
		goto L26
	}
L5:
	;
	v95 = int32(0)
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
	v35 = m.ExcPending
	if v35 != 0 {
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
	v28 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v13)+8)))
	v30 = int64(0)
	v32 = F_GetSysCacheOid(m, l5, v28, base.I64_extend_i32_u(v25), v30, v30)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v95 = v32
	goto L4
L15:
	;
	v36 = int32(0)
	v38 = *(*int32)(unsafe.Add(mBase, _c_Fn14294[0]))
	if v38 == v36 {
		v95 = v36
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v41 = int32(0)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v42 <= v41 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v45 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v13)+8)))
	v53 = v41
	goto L18
L18:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56+v53<<(uint(int32(2))%32))))
	v62 = *(*int32)(unsafe.Add(mBase, _c_Fn14294[1]))
	if v60 != v62 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L5
L20:
	;
	v65 = int64(0)
	v67 = F_GetSysCacheOid(m, l5, v45, base.I64_extend_i32_u(v60), v65, v65)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v71 = v53 + int32(1)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v71 < v72 {
		v53 = v71
		goto L18
	} else {
		goto L25
	}
L23:
	;
	if v67 != 0 {
		v120 = v67
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
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v104 = F_NameListToString(m, l0)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v104
	F_errmsg(m, l4, v13)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_Fn14294_0), l3, l2)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
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
func Fn14298(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		if v16 == int32(-1) {
			m.G0 = v9 + int32(16)
			return base.I64_extend_i32_u(v12)
		} else {
			v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
			if v16 == v19 {
				m.G0 = v9 + int32(16)
				return base.I64_extend_i32_u(v12)
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v19
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v16
						F_errmsg(m, int32(_a_Fn14298_0), v9)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, l2, l1, int32(_a_Fn14298_1))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int64(0)
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
func Fn14302(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
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
	var v21 int32
	_ = v21
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
	var v37 int64
	_ = v37
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v18 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v18)
		v21 = *(*int32)(unsafe.Add(mBase, _c_Fn14302[0]))
		v22 = F_convert_any_priv_string(m, v14, l1)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int64(0)
		} else {
			v26 = F_object_aclcheck_ext(m, l2, v12, v21, v22, v10+int32(15))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int64(0)
			} else {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
				if v28 == int32(1) {
					v31 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
					v37 = int64(0)
				} else {
					v37 = base.I64_extend_i32_u(base.B2i32(v26 == int32(0)))
				}
				m.G0 = v10 + int32(16)
				return v37
			}
		}
	}
}
func Fn14304(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
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
	var v39 int64
	_ = v39
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int64(0)
	} else {
		v20 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v20)
		v22 = F_get_role_oid_or_public(m, v14)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int64(0)
		} else {
			v24 = F_convert_any_priv_string(m, v16, l1)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int64(0)
			} else {
				v28 = F_object_aclcheck_ext(m, l2, v13, v22, v24, v11+int32(15))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int64(0)
				} else {
					v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
					if v30 == int32(1) {
						v33 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v33)
						v39 = int64(0)
					} else {
						v39 = base.I64_extend_i32_u(base.B2i32(v28 == int32(0)))
					}
					m.G0 = v11 + int32(16)
					return v39
				}
			}
		}
	}
}
func Fn14315(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	F_ean2isn(m, v9, v7+int32(8), l1)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
		m.G0 = v7 + int32(16)
		return v16
	}
}
func Fn14322(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
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
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	v8 = int64(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v20 = F_JsonbExtractScalar(m, v14+int32(4), v11)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			if v20 == int32(0) {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				F_cannotCastJsonbValue(m, v22, l2, v25)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int64(0)
				} else {
					v45 = v8
					m.G0 = v11 + int32(32)
					return v45
				}
			} else {
				switch v22 {
				case 0:
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v28 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int64(0)
						} else {
							v32 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v32)
							v45 = v8
							m.G0 = v11 + int32(32)
							return v45
						}
					} else {
						v32 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v32)
						v45 = v8
						m.G0 = v11 + int32(32)
						return v45
					}
				default:
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					F_cannotCastJsonbValue(m, v22, l2, v34)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int64(0)
					} else {
						v45 = v8
						m.G0 = v11 + int32(32)
						return v45
					}
				case 2:
					v38 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v11)+8)))
					v39 = F_DirectFunctionCall1Coll(m, l1, int32(0), v38)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int64(0)
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						if v14 == v41 {
							v45 = v39
							m.G0 = v11 + int32(32)
							return v45
						} else {
							F_pfree(m, v14)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int64(0)
							} else {
								v45 = v39
								m.G0 = v11 + int32(32)
								return v45
							}
						}
					}
				}
			}
		}
	}
}
func Fn14335(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v28 int64
	_ = v28
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v11 != 0 {
		v14 = int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v14 = v13
	}
	v17 = F_numeric_poly_stddev_internal(m, v14, l2, l1, v9+int32(15))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
		if v21 == int32(1) {
			v24 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v24)
			v28 = int64(0)
		} else {
			v28 = base.I64_extend_i32_u(v17)
		}
		m.G0 = v9 + int32(16)
		return v28
	}
}
func Fn14339(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
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
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v16 = F_strlen(m, l0)
	mBase = m.M
	if v16 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L19
	} else {
		goto L35
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L19
	} else {
		goto L31
	}
L3:
	;
	if v61 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L4:
	;
	v61 = int32(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v22 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v23 = v15
	v24 = l0
	v25 = v16
	v26 = v22
	goto L11
L8:
	;
	v49 = l0
	v53 = int32(0)
	goto L9
L9:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	v61 = v53 - v54
	goto L3
L10:
	;
	v49 = v44
	v53 = v46
	goto L9
L11:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if base.B2i32(v26 != v28)|base.B2i32(v28 == int32(0)) != 0 {
		v44 = v24
		v46 = v26
		goto L10
	} else {
		goto L13
	}
L12:
	;
	v44 = v38
	v46 = int32(0)
	goto L10
L13:
	;
	v34 = v25 - int32(1)
	if v34 == int32(0) {
		v44 = v24
		v46 = v26
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v37 = int32(1)
	v38 = v24 + v37
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v39 != 0 {
		v23 = v23 + v37
		v24 = v38
		v25 = v34
		v26 = v39
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v11 + int32(-4)
	v67 = v16 + v15
	v70 = F_sscanf(m, v67, l7, v11+int32(-32))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L19
	} else {
		goto L27
	}
L19:
	;
	return int32(0)
L20:
	;
	if v70 != int32(1) {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v76 = int32(10)
	v77 = F___strchrnul(m, v67, v76)
	mBase = m.M
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v79 == v76 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v83 == int32(0) {
		goto L1
	} else {
		goto L26
	}
L23:
	;
	v83 = v77
	goto L25
L24:
	;
	v83 = int32(0)
	goto L25
L25:
	;
	goto L22
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v83 + int32(1)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
	m.G0 = v13 - int32(-64)
	return v89
L27:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L19
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = l2
	F_errmsg(m, int32(_a_Fn14339_0), v11+int32(-16))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L19
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_Fn14339_1), l6, l3)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L19
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L19
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l2
	F_errmsg(m, int32(_a_Fn14339_0), v11+int32(-48))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L19
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_Fn14339_1), l5, l3)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L19
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L19
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l2
	F_errmsg(m, int32(_a_Fn14339_0), v13)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L19
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_Fn14339_1), l4, l3)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L19
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func Fn14340(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(34209794)
	*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)) = uint32(v12)
	*(*uint32)(unsafe.Add(mBase, uint32(v9)+4)) = uint32(v11)
	v18 = *(*int32)(unsafe.Add(mBase, _c_Fn14340[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v18
	v21 = F_LockAcquire(m, v9, l2, l1, int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int64(0)
	} else {
		m.G0 = v9 + int32(16)
		return int64(0)
	}
}
func Fn14346(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
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
	var v32 int32
	_ = v32
	var v35 int64
	_ = v35
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
	v12 = F_SearchSysCache1(m, int32(47), v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		if v12 == int32(0) {
			v18 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v18)
			v35 = int64(0)
			m.G0 = v8 + int32(16)
			return v35
		} else {
			F_initStringInfo(m, v8)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int64(0)
			} else {
				v24 = F_print_function_arguments(m, v8, v12, int32(0), l1)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					F_ReleaseCatCache(m, v12)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int64(0)
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
						v29 = F_cstring_to_text(m, v28)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int64(0)
						} else {
							F_pfree(m, v28)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int64(0)
							} else {
								v35 = base.I64_extend_i32_u(v29)
								m.G0 = v8 + int32(16)
								return v35
							}
						}
					}
				}
			}
		}
	}
}
func Fn14351(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(34209793)
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+8)) = uint32(v10)
	v15 = int64(base.Ui64(v10) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+4)) = uint32(v15)
	v18 = *(*int32)(unsafe.Add(mBase, _c_Fn14351[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v18
	v21 = F_LockAcquire(m, v8, l2, l1, int32(1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int64(0)
	} else {
		m.G0 = v8 + int32(16)
		return base.I64_extend_i32_u(base.B2i32(v21 != int32(0)))
	}
}
func Fn14366(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 <= v8 {
		v78 = v3
		return v78
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v6-int32(1)))))
		switch v14 - int32(105) {
		case 0, 5:
			v19 = F_find_among_b(m, l0, l1, int32(2), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				if v19 == int32(0) {
					v78 = v3
					return v78
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v25
					switch v19 - int32(1) {
					case 0:
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v29&int32(-2) == int32(2) {
							v68 = F_slice_del(m, l0)
							mBase = m.M
							if v68 < int32(0) {
								v78 = v68
							} else {
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								v72 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v71 - v72
								v78 = v72
							}
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v25 <= v34 {
								if v29 == int32(1) {
									v78 = v3
								} else {
									v68 = F_slice_del(m, l0)
									mBase = m.M
									if v68 < int32(0) {
										v78 = v68
									} else {
										v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										v72 = int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v71 - v72
										v78 = v72
									}
								}
							} else {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v25-int32(1)))))
								if v40 != int32(107) {
									if v29 == int32(1) {
										v78 = v3
									} else {
										v68 = F_slice_del(m, l0)
										mBase = m.M
										if v68 < int32(0) {
											v78 = v68
										} else {
											v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											v72 = int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v71 - v72
											v78 = v72
										}
									}
								} else {
									v44 = v25 - int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v44
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v44
									v68 = F_slice_del(m, l0)
									mBase = m.M
									if v68 < int32(0) {
										v78 = v68
									} else {
										v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										v72 = int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v71 - v72
										v78 = v72
									}
								}
							}
						}
						return v78
					case 1:
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if int32(2) < v47 {
							v78 = v3
							return v78
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v25 <= v50 {
								v68 = F_slice_del(m, l0)
								mBase = m.M
								if v68 < int32(0) {
									v78 = v68
								} else {
									v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									v72 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v71 - v72
									v78 = v72
								}
								return v78
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+v25-int32(1)))))
								if v56 != int32(115) {
									v68 = F_slice_del(m, l0)
									mBase = m.M
									if v68 < int32(0) {
										v78 = v68
									} else {
										v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										v72 = int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v71 - v72
										v78 = v72
									}
									return v78
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v25 - int32(1)
									return int32(0)
								}
							}
						}
					default:
						v68 = F_slice_del(m, l0)
						mBase = m.M
						if v68 < int32(0) {
							v78 = v68
						} else {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							v72 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v71 - v72
							v78 = v72
						}
						return v78
					}
				}
			}
		default:
			v78 = v3
			return v78
		}
	}
}
func Fn14373(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v8 = base.I64_extend_i32_u(l1) * base.I64_extend_i32_u(l2)
	if int64(base.Ui64(v8)>>(uint(int64(32))%64)) != int64(0) {
		F_mul_size_error(m, l1, l2)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(8))))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v20&int32(15)*int32(36))+uint32(_c_Fn14373[0])))
		v26 = m.T0[v25].(func(*base.Module, int32, int32, int32) int32)(m, l0, base.I32_wrap_i64(v8), l3)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			return v26
		}
	}
}
func Fn14375(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
	v12 = F_pg_snprintf(m, l0, int32(12), int32(_a_Fn14375_0), v6)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return l0
	}
}
func Fn14386(m *base.Module, l0 int32, l1 int32) int64 {
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
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if int32(0) <= v8 {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
		v14 = F_psprintf(m, int32(_a_Fn14386_0), v6)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			v20 = v14
			m.G0 = v6 + int32(16)
			return base.I64_extend_i32_u(v20)
		}
	} else {
		v18 = F_pstrdup(m, l1)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v20 = v18
			m.G0 = v6 + int32(16)
			return base.I64_extend_i32_u(v20)
		}
	}
}
func Fn14393(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	v2 = l1
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum_copy(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v13 = F_pg_detoast_datum_copy(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			if v15 == int32(0) {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v8 != v18 {
					v77 = v13
					v79 = v8
					F_pfree(m, v79)
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int64(0)
					} else {
						v82 = v77
						return base.I64_extend_i32_u(v82)
					}
				} else {
					v82 = v13
					return base.I64_extend_i32_u(v82)
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
				if v20 == int32(0) {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v23 != v13 {
						v77 = v8
						v79 = v13
						F_pfree(m, v79)
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int64(0)
						} else {
							v82 = v77
							return base.I64_extend_i32_u(v82)
						}
					} else {
						v82 = v8
						return base.I64_extend_i32_u(v82)
					}
				} else {
					v26 = F_palloc0(m, int32(24))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int64(0)
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v28 | int32(1)
						v33 = F_palloc0(m, int32(12))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v26))) = v33
							v36 = int32(2)
							*(*uint8)(unsafe.Add(mBase, uint32(v33))) = uint8(v36)
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
							*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)) = uint8(v2)
							v42 = F_palloc0_mul(m, int32(4), v36)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v42
								v46 = v13 + int32(8)
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
								v51 = F_QT2QTN(m, v46, v46+v47*int32(12))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int64(0)
								} else {
									v53 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
									*(*int32)(unsafe.Add(mBase, uint32(v53))) = v51
									v56 = v8 + int32(8)
									v57 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
									v61 = F_QT2QTN(m, v56, v56+v57*int32(12))
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int64(0)
									} else {
										v63 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
										*(*int32)(unsafe.Add(mBase, uint32(v63)+4)) = v61
										*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = int32(2)
										v67 = F_QTN2QT(m, v26)
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return int64(0)
										} else {
											F_QTNFree(m, v26)
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return int64(0)
											} else {
												v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
												if v71 != v8 {
													F_pfree(m, v8)
													mBase = m.M
													v74 = m.ExcPending
													if v74 != 0 {
														return int64(0)
													} else {
														v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
														if v75 == v13 {
															v82 = v67
															return base.I64_extend_i32_u(v82)
														} else {
															v77 = v67
															v79 = v13
															F_pfree(m, v79)
															mBase = m.M
															v81 = m.ExcPending
															if v81 != 0 {
																return int64(0)
															} else {
																v82 = v77
																return base.I64_extend_i32_u(v82)
															}
														}
													}
												} else {
													v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
													if v75 == v13 {
														v82 = v67
														return base.I64_extend_i32_u(v82)
													} else {
														v77 = v67
														v79 = v13
														F_pfree(m, v79)
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int64(0)
														} else {
															v82 = v77
															return base.I64_extend_i32_u(v82)
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
func Fn14403(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v53 int32
	_ = v53
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v53
L2:
	;
	v13 = l5
	v18 = int32(0)
	goto L5
L3:
	;
	goto L4
L4:
	;
	v53 = base.B2i32(base.Ui32(l0-l2) < base.Ui32(int32(26)))
	goto L1
L5:
	;
	v23 = base.I32_div_s(v13+v18, int32(2))
	v25 = v23 << (uint(int32(3)) % 32)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v25+l4)))
	if base.Ui32(v27) < base.Ui32(l0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v53 = int32(0)
	goto L1
L7:
	;
	if v38 <= v37 {
		v13 = v37
		v18 = v38
		goto L5
	} else {
		goto L12
	}
L8:
	;
	v37 = v13
	v38 = v23 + int32(1)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v25+l3)))
	if base.Ui32(v33) <= base.Ui32(l0) {
		v53 = int32(1)
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v37 = v23 - int32(1)
	v38 = v18
	goto L7
L12:
	;
	goto L6
}
func Fn14405(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int64
	_ = v29
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_WinCheckAndInitializeNullTreatment(m, v10, int32(1), l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v21 = F_WinGetFuncArgInPartition(m, v10, l1, int32(1), v8+int32(15), v8+int32(14))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int64(0)
		} else {
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
			if v23 == int32(1) {
				v26 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
				v29 = int64(0)
			} else {
				v29 = v21
			}
			m.G0 = v8 + int32(16)
			return v29
		}
	}
}
